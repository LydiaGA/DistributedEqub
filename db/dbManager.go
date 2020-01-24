package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type Equb struct {
	gorm.Model

	Name         string
	CurrentMonth string
	Members      []Member
	Winner       Member
	Status       string

	NextServerID uint
}

type Member struct {
	gorm.Model

	EqubID uint

	Name    string
	HasPaid bool
	Amount  int

	IP string
}

func (model *Equb) CreateEqub(database *gorm.DB) {
	if len(FindEqub(database)) == 0 {
		database.Create(model)
	}

	for _, member := range model.Members {
		memberFound := FindMember(database, member.ID)
		if memberFound.ID != 0 {
			database.Create(&member)
		}

	}
}

func FindEqub(database *gorm.DB) []Equb {
	equbs := []Equb{}
	database.Preload("Members").First(&equbs)

	return equbs
}

func (model *Equb) SetNextServer(database *gorm.DB, member Member) {
	model.NextServerID = member.ID
	database.Save(&model)
}

func (model *Member) CreateMember(database *gorm.DB, equb Equb) {
	model.EqubID = equb.ID
	database.Create(model)
}

func FindMember(database *gorm.DB, id uint) Member {
	var member Member
	database.First(&member, "ID = ?", id)
	return member
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
