package main

import (
	"equb/config"
	db2 "equb/db"
	"equb/rpc"
)

func StartServer(name string, month string){
	db := db2.GetDatabase()

	if db2.FindAllEqub(db) == nil {
		equb := db2.Equb{Name: name, CurrentMonth: month}
		equb.CreateEqub(db)
	}

	defer db.Close()

	rpc.Serve()
}

func StartClient(address string, name string, amount int){
	config.ServerIP = address
	db := db2.GetDatabase()
	defer db.Close()

	member:= db2.Member{
		Name:       name,
		HasPaid:    false,
		Amount:     amount,
		IP:         config.IP,
	}

	equb := rpc.StartClient(member)

	equb.CreateEqub(db)
	defer db.Close()
}