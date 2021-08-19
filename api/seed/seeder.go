package seed

import (
	"log"

	"github.com/RaphaelNagato/goresume-api/api/models"
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.ContactMe{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.ContactMe{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}
