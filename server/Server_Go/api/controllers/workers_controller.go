package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/victorsteven/forum/api/auth"
	"github.com/victorsteven/forum/api/models"
	"github.com/victorsteven/forum/api/security"
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
	fmt.Print(userCreated.ID)
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
			"response": userCreated,
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

	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	user := models.Worker{}

	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
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
		"response": userGotten,
	})
}

func (server *Server) UpdateUser(c *gin.Context) {

	errList = map[string]string{}

	userID := c.Param("id")

	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

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
	err = server.DB.Debug().Model(models.Worker{}).Where("id = ?", uid).Take(&formerUser).Error
	if err != nil {
		errList["User_invalid"] = "The user is does not exist"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	newUser := models.Worker{}

	//When current password has content.
	if requestBody["current_password"] == "" && requestBody["new_password"] != "" {
		errList["Empty_current"] = "Please Provide current password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	if requestBody["current_password"] != "" && requestBody["new_password"] == "" {
		errList["Empty_new"] = "Please Provide new password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	if requestBody["current_password"] != "" && requestBody["new_password"] != "" {
		//Also check if the new password
		if len(requestBody["new_password"]) < 6 {
			errList["Invalid_password"] = "Password should be atleast 6 characters"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		//if they do, check that the former password is correct
		err = security.VerifyPassword(formerUser.Password, requestBody["current_password"])
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			errList["Password_mismatch"] = "The password not correct"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		//update both the password and the email
		newUser.Name = formerUser.Name //remember, you cannot update the username
		newUser.Email = requestBody["email"]
		newUser.Password = requestBody["new_password"]
	}

	//The password fields not entered, so update only the email
	newUser.Name = formerUser.Name
	newUser.Email = requestBody["email"]

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

	updatedUser, err := newUser.UpdateAUser(server.DB, uint32(uid))
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
		"response": updatedUser,
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

/*
func CheckAdmin(c *gin.Context) {

	workerModel := models.Worker{}
	var jsonData = []byte(`{
	   		"email": "` + workerModel.Email + `"
	   	}`)
	isAdmin := workerModel.Admin

	if isAdmin == "true" {
		if data := postForm(jsonData); data != true {
			c.JSON(http.StatusNotFound, gin.H{"message": "No es un administrador"})
			return
		}
		serializer := workerModel.SerializeWorkerLogin()
		c.JSON(http.StatusOK, gin.H{"user": serializer})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No es un administrador"})

	}
} */
