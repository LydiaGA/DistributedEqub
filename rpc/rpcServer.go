package rpc

import (
	"equb/config"
	db2 "equb/db"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type SERVER int

func Serve() {
	serverPort := config.Port

	server := new(SERVER)
	err := rpc.Register(server)
	if err != nil {
		log.Fatal("Error Registering RPC", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":" + serverPort)

	if err != nil {
		log.Fatal("Listener Error", err)
	}
	log.Printf("Serving RPC on port \"%s\"", serverPort)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error Serving: ", err)
	}
}

func (SERVER) Try(input string, result *string) error {
	*result = input + "received"

	return nil
}

func (SERVER) StartClient(member db2.Member, result *db2.Equb) error {
	db := db2.GetDatabase()
	equb := db2.FindAllEqub(db)[0]
	defer db.Close()

	member.EqubID = equb.ID
	member.CreateMember(db)

	result = &db2.FindAllEqub(db)[0]

	return nil
}