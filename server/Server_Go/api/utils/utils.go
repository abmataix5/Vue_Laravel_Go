package utils

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAdmin(c *gin.Context) {

	/* 	myUserModel := c.MustGet("my_user_model").(UserModel)
	   	var jsonData = []byte(`{
	   		"email": "` + myUserModel.Email + `"
	   	}`)
	   	appointment := myUserModel.Appointment

	   	if appointment == "gerente" {
	   		if data := postForm(jsonData); data != true {
	   			c.JSON(http.StatusNotFound, gin.H{"message": "no authorized"})
	   			return
	   		}
	   		serializer := RegisterSerializer{c}
	   		c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
	   		return
	   	} else {
	   		c.JSON(http.StatusNotFound, gin.H{"message": "no authorized"})

	   	} */
}

func postForm(jsonData []byte) bool {

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

	if response.Status == "200 OK" {
		/* AÇÒ NO VA, HI HAURIA QUE COMPROBAR SI TORNA ALGO O QUÉ */
		fmt.Println("TOT OK")
		return true
	} else {
		fmt.Println("FAIL")
		// err:= nil!=error
		return false
	}

}
