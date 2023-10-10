package controller

import (
	"accountservice/app/controller"
	"accountservice/app/domain/dao"
	"accountservice/app/repository"
	"accountservice/app/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	RegisterUserEndpoint = "/api/user/register"
	LoginEndpoint        = "/api/user/login"
)

func setupTestEnvironment() (*gin.Engine, *gorm.DB) {
	// Initialize a new SQLite in-memory database.
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	// Auto-migrate the database.
	if err := db.AutoMigrate(&dao.User{}); err != nil {
		log.Fatalf("Unable to migrate User %s", err)
	}

	// Initialize a Gin router.
	router := gin.Default()

	// Inject the database and router into the UserController.
	userRepo := repository.UserRepositoryInit(db)
	userService := service.UserServiceInit(userRepo)
	userController := controller.UserControllerInit(userService)

	// Set up routes.
	router.POST(RegisterUserEndpoint, userController.RegisterUser)
	router.POST(LoginEndpoint, userController.LoginUser)
	return router, db
}

func TestUserController_RegisterUserSuccess(t *testing.T) {
	// Set up the test environment.
	router, db := setupTestEnvironment()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	req := httptest.NewRequest("POST", RegisterUserEndpoint, strings.NewReader(`{"email":"testuser@gmail.com","password":"TestUser@2023","fullname":"Paul Odhiambo"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"response_key":"SUCCESS"`)
	var user dao.User
	if err := json.Unmarshal(w.Body.Bytes(), &user); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	assert.NotNil(t, user)
}

func TestUserController_TestRegisterUserBadRequest(t *testing.T) {
	// Set up the test environment.
	router, db := setupTestEnvironment()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	req := httptest.NewRequest("POST", RegisterUserEndpoint, strings.NewReader(`{"email":"testuser@gmail.com"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"response_key":"INVALID_REQUEST"`)
	var user dao.User
	if err := json.Unmarshal(w.Body.Bytes(), &user); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	assert.NotNil(t, user)
}

func TestUserController_TestRegisterUserEmailExists(t *testing.T) {
	// Set up the test environment.
	router, db := setupTestEnvironment()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
	//create first user
	httptest.NewRequest("POST", RegisterUserEndpoint, strings.NewReader(`{"email":"testuser@gmail.com","password":"TestUser@2023","fullname":"Paul Odhiambo"}`))
	req := httptest.NewRequest("POST", RegisterUserEndpoint, strings.NewReader(`{"email":"testuser@gmail.com","password":"TestUser@2023"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"response_key":"INVALID_REQUEST"`)
	var user dao.User
	if err := json.Unmarshal(w.Body.Bytes(), &user); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	assert.NotNil(t, user)
}
func TestUserController_LoginUserSuccess(t *testing.T) {
	// Set up the test environment.
	router, db := setupTestEnvironment()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	httptest.NewRequest("POST", RegisterUserEndpoint, strings.NewReader(`{"email":"testuser@gmail.com","password":"TestUser@2023","fullname":"Paul Odhiambo"}`))
	req := httptest.NewRequest("POST", LoginEndpoint, strings.NewReader(`{"email":"testuser@gmail.com","password":"TestUser@2023"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	log.Printf("response body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	var resp LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, w.Body.String(), `"response_key":"SUCCESS"`)
	assert.Contains(t, w.Body.String(), `"response_message":"Success"`)
	// Check the response body.
	assert.Equal(t, "SUCCESS", resp.ResponseKey)
	assert.Equal(t, "Success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.Data.Token)
	assert.NotEmpty(t, resp.Data.RefreshToken)

}

type LoginResponse struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            struct {
		RefreshToken string `json:"refresh_token"`
		Token        string `json:"token"`
	} `json:"data"`
}
