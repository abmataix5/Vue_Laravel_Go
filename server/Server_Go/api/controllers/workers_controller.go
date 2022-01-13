package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/victorsteven/forum/api/models"
	"github.com/victorsteven/forum/api/utils/formaterror"
)

func (server *Server) CreateUser(c *gin.Context) {
	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	user := models.Worker{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	user.Prepare()
	errorMessages := user.Validate("")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	serialized := userCreated.SerializeWorkerInfo()

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	} else { /* Envia un email ala dirección del nuevo trabajador registrado por el administrador */

		from := mail.NewEmail("Administrador", user.Email)
		subject := "Alta Nuevo Trabajador"
		to := mail.NewEmail("Administrador MyPadel", user.Email)
		plainTextContent := ("Has sido registrado como nuevo trabajador! Esta es tu información:")
		htmlContent := ("<strong>Has sido registrado como nuevo trabajador! Esta es tu información:</strong><br>Email del trabajador: " + user.Email + "<br> Username: " + user.Name + "<br> Para iniciar sesión:  " + "http://192.168.8.161:8080/login")
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   http.StatusCreated,
			"response": serialized,
		})
	}

}

func (server *Server) GetUsers(c *gin.Context) {

	errList = map[string]string{}
	user := models.Worker{}
	users, err := user.FindAllUsers(server.DB)
	fmt.Println(c.Request.Body)
	if err != nil {
		errList["No_user"] = "No hemos encontrado usuarios disponibles"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	fmt.Println(users)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": users})

}

func (server *Server) GetUser(c *gin.Context) {

	errList = map[string]string{}
	userID := c.Param("id")

	user := models.Worker{}

	userGotten, err := user.FindUserByID(server.DB, userID)
	serialized := userGotten.SerializeWorkerInfo()
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": serialized,
	})
}

func (server *Server) UpdateUser(c *gin.Context) {

	errList = map[string]string{}

	userID := c.Param("id")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	formerUser := models.Worker{}
	err = server.DB.Debug().Model(models.Worker{}).Where("id = ?", userID).Take(&formerUser).Error
	if err != nil {
		errList["User_invalid"] = "The user is does not exist"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	newUser := models.Worker{}

	//Fields update

	newUser.Name = formerUser.Name
	newUser.Email = requestBody["email"]
	newUser.Name = requestBody["username"]
	newUser.Phone = requestBody["phone"]
	newUser.Address = requestBody["address"]
	newUser.Worker_type = requestBody["worker_type"]

	newUser.Prepare()
	errorMessages := newUser.Validate("update")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	updatedUser, err := newUser.UpdateAUser(server.DB, string(userID))
	serialized := updatedUser.SerializeWorkerInfo()

	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": serialized,
	})
}

func (server *Server) DeleteUser(c *gin.Context) {

	/*Limpia cadena de errores si existe alguno */

	errList = map[string]string{}

	userID := c.Params.ByName("id")

	user := models.Worker{}
	_, err := user.DeleteAUser(server.DB, userID)
	if err != nil {
		fmt.Println(err)
		errList["Other_error"] = "UUID no encontrado!! Error en borrar usuario"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

}

func CheckAdmin(c *gin.Context) {

	workerModel := models.Worker{}
	var jsonData = []byte(`{
	   		"email": "` + workerModel.Email + `"
	   	}`)
	isAdmin := workerModel.Admin

	if isAdmin == "true" {
		if data := httpPost(jsonData); data != true {
			c.JSON(http.StatusNotFound, gin.H{"message": "No es un administrador"})
			return
		}
		serializer := workerModel.SerializeWorkerLogin()
		c.JSON(http.StatusOK, gin.H{"user": serializer})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No es un administrador"})

	}
}

func httpPost(jsonData []byte) bool {

	fmt.Println("http://0.0.0.0:8001/api/users/")

	url := "http://0.0.0.0:8001/api/users/"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("http.Status", http.StatusOK)

	if response.Status == "200" {

		fmt.Println("succes")
		return true
	} else {
		fmt.Println("error")

		return false
	}

}
