package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Get all workers
func GetAllWorkers(worker *[]Worker) (err error) {
	if err = Config.DB.Find(worker).Error; err != nil {
		return err
	}
	return nil
}

//Create one worker
func CreateUser(worker *Worker) (err error) {
	if err = Config.DB.Create(worker).Error; err != nil {
		return err
	}
	return nil
}

//Get on worker by ID
func GetUserByID(worker *Worker, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(worker).Error; err != nil {
		return err
	}
	return nil
}

//Update worker by ID
func UpdateUser(worker *Worker, id string) (err error) {
	fmt.Println(worker)
	Config.DB.Save(worker)
	return nil
}

//Delete worker by ID
func DeleteUser(worker *Worker, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(worker)
	return nil
}
