package setup

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/masato25/micro_ico_ethereum/app/model/tokens"
	"github.com/masato25/micro_ico_ethereum/app/model/users"
	"github.com/masato25/micro_ico_ethereum/config"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func GetConn() *gorm.DB {
	return db
}

func ConnDB(migrate ...bool) (err error) {
	dbconf := config.MyConfig().Database
	connpath := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		dbconf.Host, dbconf.Port, dbconf.User, dbconf.DBName, dbconf.Password)
	db, err = gorm.Open("postgres", connpath)
	if err != nil {
		log.Error(err)
		return
	}
	if dbconf.Debug {
		db.LogMode(true)
	}
	if len(migrate) != 0 {
		log.Debug(migrate[0])
		if migrate[0] {
			Migration()
		}
	}
	return
}

func Migration() {
	db.DropTable(&users.Session{})
	db.DropTable(&users.Account{})
	db.DropTable(&tokens.Token{})
	db.AutoMigrate(
		users.Session{},
		users.Account{},
		tokens.Token{},
	)
	db.Model(&users.Account{}).AddUniqueIndex("idx_user_name", "user_name")
	db.Model(&users.Account{}).Related(&users.Session{})
}

func CloseDB() error {
	return db.Close()
}
