package models

import (
	"log"

	"github.com/victorsteven/forum/api/auth"
	"github.com/victorsteven/forum/api/common"
)

/* Serializer para devolver datos en el inicio de sesion */

func (worker *Worker) SerializeWorkerLogin() common.JSON {

	token, err := auth.CreateToken(worker.ID)
	if err != nil {
		log.Fatal("Error generando el token", err)
	}

	return common.JSON{
		"id":       worker.ID,
		"username": worker.Name,
		"email":    worker.Email,
		"token":    token,
		"admin":    worker.Admin,
		"admin2":   true,
	}
}

/* Serializer para devolver la información del trabajador */

func (worker *Worker) SerializeWorkerInfo() common.JSON {

	return common.JSON{
		"id":          worker.ID,
		"username":    worker.Name,
		"email":       worker.Email,
		"pone":        worker.Phone,
		"address":     worker.Address,
		"worker_type": worker.Worker_type,
		"antiguedad":  worker.Antiguedad,
		"admin":       worker.Admin,
	}
}
