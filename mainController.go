package main

import (
	"equb1/DistributedEqub/config"
	db2 "equb1/DistributedEqub/db"
	"equb1/DistributedEqub/rpc"
)

func StartServer(name string, month int) {
	db := db2.GetDatabase()
	equb := db2.Equb{Name: name, CurrentMonth: month, Status: "created"}
	equb.CreateEqub(db)

	defer db.Close()

	rpc.Serve()
}

func StartEqub() {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	equb.Status = "started"
	db.Save(&equb)
}

func StartClient(address string, name string, amount int) {
	config.ServerIP = address
	db := db2.GetDatabase()
	defer db.Close()

	member := db2.Member{
		Name:    name,
		HasPaid: false,
		Amount:  amount,
		IP:      config.IP,
	}

	equb := rpc.StartClient(member)
	equb.CreateEqub(db)

	config.Me = db2.FindMember(db, equb.NextServerID)
}

func GetTotal() int {
	db := db2.GetDatabase()
	defer db.Close()
	equb := rpc.GetEqub()
	db2.UpdateEqub(db, equb)
	return equb.Total
}

func GetList() []db2.Member {
	db := db2.GetDatabase()
	defer db.Close()
	equb := rpc.GetEqub()
	db2.UpdateEqub(db, equb)
	return equb.Members
}

func MakePayment() string {
	db := db2.GetDatabase()
	defer db.Close()
	message, equb := rpc.MakePayment()
	db2.UpdateEqub(db, equb)
	return message
}

func CollectWinnings() string {
	db := db2.GetDatabase()
	defer db.Close()
	message, equb := rpc.CollectWinnings()
	db2.UpdateEqub(db, equb)
	return message
}

func ChangeMonth() string {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	defer db.Close()
	equb.CurrentMonth = equb.CurrentMonth + 1
	db.Save(&equb)

	for _, member := range equb.Members {
		memberFound := db2.FindMember(db, member.ID)
		memberFound.HasPaid = false
		db.Save(&memberFound)
	}

	return "Successfully Changed"
}
