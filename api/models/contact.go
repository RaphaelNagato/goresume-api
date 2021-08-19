package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//Message Struct(Model)
type ContactMe struct {
	ID          string
	Subject     string `gorm:"not null" json:"subject"`
	Message     string `gorm:"not null" json:"message"`
	SenderEmail string `gorm:"not null" json:"sender"`
}

func (m *ContactMe) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid.String())
}

func (m *ContactMe) SaveMessage(db *gorm.DB) (*ContactMe, error) {

	err := db.Debug().Create(&m).Error
	if err != nil {
		return &ContactMe{}, err
	}
	return m, nil
}

func (m *ContactMe) FindAllMessages(db *gorm.DB) (*[]ContactMe, error) {
	var err error
	messages := []ContactMe{}
	err = db.Debug().Model(&ContactMe{}).Limit(100).Find(&messages).Error
	if err != nil {
		return &[]ContactMe{}, err
	}
	return &messages, err
}
