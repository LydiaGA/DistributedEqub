package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type Equb struct {
	gorm.Model

	Name     string
	CurrentMonth  string
	Members []Member `gorm:"foreignkey:EqubRefer"`
	Winner Member

	NextServer Member
}

type Member struct{
	gorm.Model

	EqubID uint

	Name string
	HasPaid bool
	Amount int

	IP string
}

func (model *Equb) CreateEqub(database *gorm.DB) {
	database.Create(model)
	for _, member := range model.Members {
		database.Create(&member)
	}
}

func FindAllEqub(database *gorm.DB) []Equb{
	equbs := []Equb {}
	database.Preload("Members").Find(&equbs)

	return equbs
}

func (model *Member) CreateMember(database *gorm.DB) {
	database.Create(model)
}

func GetDatabase() *gorm.DB {
	workingDirectory, _ := os.Getwd()
	db, err := gorm.Open("sqlite3", workingDirectory+"/db/equb_database.db")

	if err != nil {
		panic(err)
	}
	return db
}

func Migrate() {
	db := GetDatabase()
	db.AutoMigrate(&Equb{}, &Member{})
	db.Close()
}