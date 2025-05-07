package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"study.co.jp/go-rest-gin-gorm/controllers"
	"study.co.jp/go-rest-gin-gorm/database"
	"study.co.jp/go-rest-gin-gorm/models"
)

var ID int

const TESTNUMBER = "TESTME"

func TestingRoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateTestStudent() {
	student := models.Student{Name: "Student test", Number: TESTNUMBER, SSID: "987654321"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteTestStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestSaluteIsOK(t *testing.T) {
	// Setup
	r := TestingRoutesSetup()
	r.GET("/salute/:name", controllers.Salute)

	req, _ := http.NewRequest(http.MethodGet, "/salute/TestName", nil)
	response := httptest.NewRecorder()
	expectedResponse := `{"response":"Hello TestName, how you doing?"}`

	// Call
	r.ServeHTTP(response, req)

	actualResponse, _ := io.ReadAll(response.Body)

	// Assertions
	assert.Equal(t, http.StatusOK, response.Code, "/salute - Incorrect status code")
	assert.Equal(t, expectedResponse, string(actualResponse))
}

func TestGetAllStudents(t *testing.T) {
	//TODO: This has to have something different, e.g.: an in-memory DB just for the tests
	// Maybe create a test package with the configuration for testing DB?
	database.Connect()
	CreateTestStudent()
	defer DeleteTestStudent()

	r := TestingRoutesSetup()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest(http.MethodGet, "/students", nil)
	response := httptest.NewRecorder()

	//Call
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindStudentByNumber(t *testing.T) {
	database.Connect()
	CreateTestStudent()
	defer DeleteTestStudent()

	r := TestingRoutesSetup()
	r.GET("/students/number/:number", controllers.GetStudentByNumber)

	req, _ := http.NewRequest(http.MethodGet, "/students/number/"+TESTNUMBER, nil)
	response := httptest.NewRecorder()

	//Call
	r.ServeHTTP(response, req)

	//Assertions
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestGetStudentById(t *testing.T) {
	database.Connect()
	CreateTestStudent()
	defer DeleteTestStudent()

	r := TestingRoutesSetup()
	r.GET("/students/:id", controllers.GetStudentByID)

	req, _ := http.NewRequest(http.MethodGet, "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()

	//Call
	r.ServeHTTP(response, req)

	var actualResponse models.Student
	json.Unmarshal(response.Body.Bytes(), &actualResponse)

	//Assertions
	assert.Equal(t, "Student test", actualResponse.Name)
	assert.Equal(t, TESTNUMBER, actualResponse.Number)
	assert.Equal(t, "987654321", actualResponse.SSID)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.Connect()
	CreateTestStudent()

	r := TestingRoutesSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	req, _ := http.NewRequest(http.MethodDelete, "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()

	//Call
	r.ServeHTTP(response, req)

	//Assert
	assert.Equal(t, http.StatusOK, response.Code)
	//TODO: It would be better to have an actual check if the Student doesn't exist in the DB

}

// TODO: Refactor this once structs for DTO and internal are separated
func TestUpdateStudent(t *testing.T) {
	database.Connect()
	CreateTestStudent()
	// defer DeleteTestStudent()

	r := TestingRoutesSetup()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	expectedResponse := models.Student{Name: "Updated student", Number: "123456789", SSID: "22346534534"}
	jsonRequestBody, _ := json.Marshal(expectedResponse)
	fmt.Println("Request: ", string(jsonRequestBody))

	req, _ := http.NewRequest(http.MethodPatch, "/students/"+strconv.Itoa(ID), bytes.NewBuffer(jsonRequestBody))
	response := httptest.NewRecorder()

	//Call
	r.ServeHTTP(response, req)

	var actualResponse models.Student
	json.Unmarshal(response.Body.Bytes(), &actualResponse)

	fmt.Println("RESPONSE: ")
	fmt.Println(actualResponse)

	//Assertions
	assert.Equal(t, "Updated student", actualResponse.Name)
	assert.Equal(t, "123456789", actualResponse.Number)
	assert.Equal(t, "22346534534", actualResponse.SSID)
	assert.Equal(t, http.StatusOK, response.Code)

}
