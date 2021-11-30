package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all workers
func GetWorkers(c *gin.Context) {
	var worker []Models.Worker
	err := Models.GetAllWorkers(&worker)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, worker)
	}
}

//Create one worker
func CreateWorker(c *gin.Context) {
	var worker Models.Worker
	c.BindJSON(&worker)
	err := Models.CreateUser(&worker)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, worker)
	}
}

//Get one worker by id
func GetWorkerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var worker Models.Worker
	err := Models.GetUserByID(&worker, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, worker)
	}
}

//Update worker information
func UpdateWorker(c *gin.Context) {
	var worker Models.Worker
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&worker, id)
	if err != nil {
		c.JSON(http.StatusNotFound, worker)
	}
	c.BindJSON(&worker)
	err = Models.UpdateUser(&worker, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, worker)
	}
}

//Delte worker by id
func DeleteWorker(c *gin.Context) {
	var worker Models.Worker
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&worker, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"El trabajador con id :" + id: " ha sido eliminado"})
	}
}
