package main

import (
	"equb1/DistributedEqub/config"
	db2 "equb1/DistributedEqub/db"
	"equb1/DistributedEqub/rpc"
	"fmt"
)

func StartServer(name string, month int, port string, memberName string, amount int) {
	db := db2.GetDatabase()
	equb := db2.Equb{Name: name, CurrentMonth: month, Status: "created", NextServerID: 2}
	equb.CreateEqub(db)

	member := db2.Member{
		Name:    memberName,
		HasPaid: false,
		Amount:  amount,
		IP:      config.IP,
		Port:    port,
	}

	member.CreateMember(db, equb)

	defer db.Close()

	rpc.Serve(port)
}

func StartEqub() {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	equb.Status = "started"
	db.Save(&equb)
}

func StartClient(address string, serverPort string, port string, name string, amount int) {
	config.ServerIP = address
	config.ServerPort = serverPort
	db := db2.GetDatabase()
	defer db.Close()

	config.ClientPort = port
	go rpc.Serve(port)
	fmt.Println("After Serve")
	equb := db2.Equb{Name: "", CurrentMonth: 0, Status: ""}
	equb.CreateEqub(db)

	member := db2.Member{
		Name:    name,
		HasPaid: false,
		Amount:  amount,
		IP:      config.IP,
		Port:    config.ClientPort,
	}

	//rpc.StartClient(member)
	equb = rpc.StartClient(member)
	//equb.CreateEqub(db)

	config.Me = db2.FindMember(db, equb.Members[len(equb.Members)-1].ID)
}

func GetTotal() int {
	db := db2.GetDatabase()
	defer db.Close()
	equb := rpc.GetEqub()
	//db2.UpdateEqub(db, equb)
	return equb.Total
}

func GetList() []db2.Member {
	db := db2.GetDatabase()
	defer db.Close()
	equb := rpc.GetEqub()
	//db2.UpdateEqub(db, equb)
	return equb.Members
}

func MakePayment() string {
	db := db2.GetDatabase()
	defer db.Close()
	message, _ := rpc.MakePayment()
	//db2.UpdateEqub(db, equb)
	return message
}

func CollectWinnings() string {
	db := db2.GetDatabase()
	defer db.Close()
	message, _ := rpc.CollectWinnings()
	//db2.UpdateEqub(db, equb)
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

	equb = db2.FindEqub(db)[0]
	rpc.NotifyAll(equb.Members, equb)

	return "Successfully Changed"
}
