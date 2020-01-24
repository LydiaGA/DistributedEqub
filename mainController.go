package main

import (
	"equb1/DistributedEqub/config"
	db2 "equb1/DistributedEqub/db"
	"equb1/DistributedEqub/rpc"
)

func StartServer(name string, month string) {
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
}
