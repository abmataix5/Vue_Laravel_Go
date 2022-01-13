package models

import (
	"errors"
	"fmt"
	"html"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

func (worker *Worker) FindUserByID(db *gorm.DB, uid string) (*Worker, error) {
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

func (worker *Worker) UpdateAUser(db *gorm.DB, uid string) (*Worker, error) {

	db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).UpdateColumns(
		map[string]interface{}{
			"email":       worker.Email,
			"phone":       worker.Phone,
			"address":     worker.Address,
			"username":    worker.Name,
			"worker_type": worker.Worker_type,
		},
	)

	if db.Error != nil {
		return &Worker{}, db.Error
	}

	err := db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&worker).Error
	if err != nil {
		return &Worker{}, err
	}
	return worker, nil
}

func (u *Worker) DeleteAUser(db *gorm.DB, uid string) (int64, error) {

	db = db.Debug().Model(&Worker{}).Where("id = ?", uid).Take(&Worker{}).Delete(&Worker{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (worker *Worker) send_email_register(*Worker) {
	fmt.Print("Entra email")
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
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

}
