package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/victorsteven/forum/api/security"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type Worker struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null;unique" json:"username"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Password    string `gorm:"size:100;not null;" json:"password"`
	Address     string `gorm:"size:255;not null;unique" json:"adress"`
	Worker_type string `gorm:"size:255;not null;unique" json:"worker_type"`
	Antiguedad  string `gorm:"size:255;not null;unique" json:"antiguedad"`
	/* 	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	   	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` */
}

func (u *Worker) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Worker) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	/* 	u.CreatedAt = time.Now()
	   	u.UpdatedAt = time.Now() */
}

func (u *Worker) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			err = errors.New("required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	case "login":
		if u.Password == "" {
			err = errors.New("required password")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("required email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if u.Name == "" {
			err = errors.New("required username")
			errorMessages["Required_username"] = err.Error()
		}
		if u.Password == "" {
			err = errors.New("required password")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Password != "" && len(u.Password) < 6 {
			err = errors.New("password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("required email")
			errorMessages["Required_email"] = err.Error()

		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages
}

func (u *Worker) SaveUser(db *gorm.DB) (*Worker, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Worker{}, err
	}
	return u, nil
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

func (u *Worker) FindUserByID(db *gorm.DB, uid uint32) (*Worker, error) {
	var err error
	err = db.Debug().Model(Worker{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Worker{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Worker{}, errors.New("usuario no encontrado")
	}
	return u, err
}

func (u *Worker) UpdateAUser(db *gorm.DB, uid uint32) (*Worker, error) {
	if u.Password != "" {
		// To hash the password
		err := u.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).UpdateColumns(
			map[string]interface{}{
				"password":  u.Password,
				"email":     u.Email,
				"update_at": time.Now(),
			},
		)
	}

	db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).UpdateColumns(
		map[string]interface{}{
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Worker{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Worker{}, err
	}
	return u, nil
}

func (u *Worker) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).Delete(&Worker{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
