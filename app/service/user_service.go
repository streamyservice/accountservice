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
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute program Login")

	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		log.Error("JSON binding error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//if u.userRepository.UserExists(loginRequest.Email) {
	//	log.Fatalf("User with the provided email address does not exist")
	//	pkg.PanicException(constants.InvalidRequest)
	//}

	user, err := u.userRepository.GetUser(loginRequest.Email)
	if err != nil {
		// Handle the error (e.g., user not found)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !verifyPassword(loginRequest.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	header := dto.Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		fmt.Println("Error marshaling header to JSON:", err)
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})

		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, os.Getenv("JWT_SECRET"))

	if err != nil {
		// Handle token generation error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})

		return
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
		log.Fatalf("User with the provided email address already exist")
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
func (u UserServiceImpl) RefreshAuthToken(c *gin.Context) {

}
func (u UserServiceImpl) VerifyUserEmail(c *gin.Context) {

}
func (u UserServiceImpl) DeleteUser(c *gin.Context) {

}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func verifyPassword(plainPassword string, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil

}
