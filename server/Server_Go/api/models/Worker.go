package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
	"github.com/victorsteven/forum/api/security"
)

type Worker struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primary_key;json:"id""`
	Name        string    `gorm:"size:255;not null;unique" json:"username"`
	Email       string    `gorm:"size:100;not null;unique" json:"email"`
	Phone       string    `gorm:"size:100;not null;unique" json:"phone"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	Address     string    `gorm:"size:255;not null;unique" json:"adress"`
	Worker_type string    `gorm:"size:255;not null;unique" json:"worker_type"`
	Antiguedad  string    `gorm:"size:255;not null;unique" json:"antiguedad"`
	Admin       string    `gorm:"size:255;not null;unique" json:"admin"`
}

func (worker *Worker) BeforeSave() error {
	hashedPassword, err := security.Hash(worker.Password)
	if err != nil {
		return err
	}
	worker.Password = string(hashedPassword)
	worker.ID = uuid.NewV4()

	return nil
}

func (worker *Worker) Prepare() {
	worker.Name = html.EscapeString(strings.TrimSpace(worker.Name))
	worker.Email = html.EscapeString(strings.TrimSpace(worker.Email))

}

func (worker *Worker) SaveUser(db *gorm.DB) (*Worker, error) {

	var err error
	err = db.Debug().Create(&worker).Error
	if err != nil {
		return &Worker{}, err
	}
	return worker, nil
}

func (u *Worker) FindAllUsers(db *gorm.DB) (*[]Worker, error) {
	var err error
	worker := []Worker{}
	err = db.Debug().Model(&Worker{}).Limit(100).Find(&worker).Error
	if err != nil {
		return &[]Worker{}, err
	}
	return &worker, err
}

func (worker *Worker) FindUserByID(db *gorm.DB, uid uint32) (*Worker, error) {
	var err error
	err = db.Debug().Model(Worker{}).Where("id = ?", uid).Take(&worker).Error
	if err != nil {
		return &Worker{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Worker{}, errors.New("usuario no encontrado")
	}
	return worker, err
}

func (worker *Worker) UpdateAUser(db *gorm.DB, uid uint32) (*Worker, error) {
	if worker.Password != "" {
		// To hash the password
		err := worker.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).UpdateColumns(
			map[string]interface{}{
				"password":  worker.Password,
				"email":     worker.Email,
				"update_at": time.Now(),
			},
		)
	}

	db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).UpdateColumns(
		map[string]interface{}{
			"email":     worker.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Worker{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&worker).Error
	if err != nil {
		return &Worker{}, err
	}
	return worker, nil
}

func (u *Worker) DeleteAUser(db *gorm.DB, uid string) (int64, error) {
	/* fmt.Println(uid) */
	db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).Delete(&Worker{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
