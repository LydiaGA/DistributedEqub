package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type Equb struct {
	gorm.Model

	Name         string
	CurrentMonth int
	Total        int
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

	IP   string
	Port string
}

type Me struct {
	gorm.Model
	MyId uint
}

func (model *Equb) CreateEqub(database *gorm.DB) {
	if len(FindEqub(database)) == 0 {
		database.Create(model)
	}

	for _, member := range model.Members {
		memberFound := FindMember(database, member.ID)
		if memberFound.ID == 0 {
			database.Create(&member)
		}

	}
}

func FindEqub(database *gorm.DB) []Equb {
	equbs := []Equb{}
	database.Preload("Members").First(&equbs)

	return equbs
}

func UpdateEqub(database *gorm.DB, model Equb) {
	equb := FindEqub(database)[0]

	equb.Winner = model.Winner
	equb.Total = model.Total
	equb.Name = model.Name
	equb.Status = model.Status
	equb.CurrentMonth = model.CurrentMonth
	equb.NextServerID = model.NextServerID

	database.Save(&equb)

	for _, member := range model.Members {
		memberFound := FindMember(database, member.ID)
		if memberFound.ID == 0 {
			database.Create(&member)
		} else {
			memberFound.HasPaid = member.HasPaid
		}
	}
}

func (model *Equb) SetNextServer(database *gorm.DB, member Member) {
	model.NextServerID = member.ID
	database.Save(&model)
}

func SaveMe(database *gorm.DB, member Member) {
	me := Me{
		MyId: member.ID,
	}
	if len(FindMe(database)) == 0 {
		database.Create(&me)
	}
}

func FindMe(database *gorm.DB) []Me {
	me := []Me{}
	database.First(&me)
	return me
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
	db.AutoMigrate(&Equb{}, &Member{}, &Me{})
	db.Close()
}
