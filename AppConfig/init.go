/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 08.06.2020  18:05
   Notes   :
*/
package AppConfig

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB                *gorm.DB
	appDBName         = "gorm"
	appDBHost         = "localhost"
	appDBUserName     = "postgres"
	appDBUserPassword = "mglo0704"
)

func InitDB() {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBUserPassword,
		appDBName)

	var err error
	DB, err = gorm.Open("postgres", cnnString)
	if err != nil {
		log.Println("DB Error", err)
	}
	log.Println("DB Connected")

}
