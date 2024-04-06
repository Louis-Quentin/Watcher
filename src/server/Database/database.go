package Database

import (
	"Area/Models"
	"fmt"
	"gorm.io/driver/postgres"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func (database *Database) Init_database() (err error) {
	dsn := "host=db user=dev_user password=Watcher dbname=watcher port=5432 sslmode=disable TimeZone=Europe/Paris"
	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("ON A PAS OPEN LA DB")
		return err
	}
	fmt.Println("ON A RÉUSSI A OPEN LA DB")
	err = database.DB.AutoMigrate(&Models.User{}, &Models.ServiceState{}, &Models.AreaState{})
	return err
}
