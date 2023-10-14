package service

import (
	"accountservice/app/constants"
	"accountservice/app/domain/dao"
	"accountservice/app/domain/dto"
	"accountservice/app/pkg"
	"accountservice/app/repository"
	"accountservice/app/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type UserService interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
	RefreshAuthToken(c *gin.Context)
	VerifyUserEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
	GetUserToken(c *gin.Context)
	SaveUserToken(c *gin.Context)
	CodeExists(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository      repository.UserRepository
	userTokenRepository repository.UserTokenRepository
}

func (u UserServiceImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute program Login")

	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		log.Errorf("Invalid data %s", err)
		pkg.PanicException(constants.InvalidRequest)
	}

	user, err := u.userRepository.GetUser(loginRequest.Email)
	if err != nil {
		log.Errorf("User not found  %s", err)
		pkg.PanicException(constants.DataNotFound)
	}

	if !verifyPassword(loginRequest.Password, user.Password) {
		log.Errorf("Incorrect Password %s", err)
		pkg.PanicException(constants.Unauthorized)
	}

	header := dto.Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		log.Errorf("Error marshaling header to JSON: %s", err)
		pkg.PanicException(constants.UnknownError)
	}

	// Marshal JWTClaims into a map
	payloadMap := map[string]interface{}{
		"UserID":   user.ID,
		"Username": user.Username,
		"Exp":      time.Now().Add(1 * time.Hour),
	}

	// generate jwt token

	token, err := utils.GenerateToken(string(headerBytes), payloadMap, os.Getenv("JWT_SECRET"))
	if err != nil {
		// Handle token generation error
		log.Errorf("Invalid Token Generated:  %s", err)
		pkg.PanicException(constants.UnknownError)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, os.Getenv("JWT_SECRET"))

	if err != nil {
		// Handle token generation error
		log.Errorf("Internal Server Error: %s", err)
		pkg.PanicException(constants.UnknownError)
	}

	// Return the token in the response
	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	}))

}
func (u UserServiceImpl) CreateUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var userRequest dto.UserRegistrationRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Errorf("Invalid data %s", err)
		pkg.PanicException(constants.InvalidRequest)
	}

	if u.userRepository.UserExists(userRequest.Email) {
		log.Errorf("User with the provided email address already exist")
		pkg.PanicException(constants.InvalidRequest)
	}

	user := dao.User{
		Email:         userRequest.Email,
		Username:      userRequest.Email,
		Fullname:      userRequest.Fullname,
		LastIp:        c.ClientIP(),
		EmailVerified: false,
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 15)
	user.Password = string(hash)

	data, err := u.userRepository.CreateUser(&user)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constants.UnknownError)
	}
	fmt.Printf("%s", data)

	// create and save token details for the user
	code, err := utils.GenOtpCode()
	userToken := &dao.UserToken{
		Id:    user.ID,
		Email: user.Email,
		Exp:   time.Now().Add(1 * time.Hour),
		Code:  code,
	}

	userTokenData, err := u.userTokenRepository.SaveUserToken(userToken)

	// send verification token to the user
	emailData := utils.EmailData{
		URL:       userTokenData.Code,
		FirstName: data.Fullname,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(data, &emailData)
	c.JSON(http.StatusCreated, pkg.BuildResponse(constants.Success, data))
}

func (u UserServiceImpl) UpdateUser(c *gin.Context) {

}
func (u UserServiceImpl) GetUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by email")
	email := c.PostForm("email")
	data, err := u.userRepository.GetUser(email)
	if err != nil {
		log.Error("Happened error when getting data from database. Error ", err)
		pkg.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))

}
func (u UserServiceImpl) GetUserToken(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by email")
	email := c.PostForm("email")
	data, err := u.userTokenRepository.GetUserToken(email)
	if err != nil {
		log.Error("Happened error when getting data from database. Error ", err)
		pkg.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))

}
func (u UserServiceImpl) RefreshAuthToken(c *gin.Context) {

}
func (u UserServiceImpl) VerifyUserEmail(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var verificationCodeRequest dto.EmailVerification

	if err := c.ShouldBindJSON(&verificationCodeRequest); err != nil {
		log.Errorf("Invalid data %s", err)
		pkg.PanicException(constants.InvalidRequest)
	}

	// check if the code exist in the database
	if !u.userTokenRepository.CodeExists(verificationCodeRequest.Code) {
		log.Errorf("Provided Code Does Not Exist")
		pkg.PanicException(constants.InvalidRequest)
	}

	// get token code using Email
	userTokenDetails, err := u.userTokenRepository.GetUserToken(verificationCodeRequest.Email)
	if err != nil {
		log.Errorf("Invalid data %s", err)
		pkg.PanicException(constants.DataNotFound)
	}
	if !utils.VerifyOTP(verificationCodeRequest.Code, userTokenDetails.Code) {
		log.Errorf("The Code Provided Does not match")
		pkg.PanicException(constants.InvalidRequest)
	}

	user, err := u.userRepository.GetUser(userTokenDetails.Email)
	if err != nil {
		log.Errorf("Invalid data %s", err)
		pkg.PanicException(constants.DataNotFound)
	}
	user.EmailVerified = true
	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, user))

}

func (u UserServiceImpl) SaveUserToken(c *gin.Context) {

}
func (u UserServiceImpl) CodeExists(c *gin.Context) {

}
func (u UserServiceImpl) DeleteUser(c *gin.Context) {

}

func UserServiceInit(userRepository repository.UserRepository, tokenRepository repository.UserTokenRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository:      userRepository,
		userTokenRepository: tokenRepository,
	}
}

func verifyPassword(plainPassword string, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil

}
