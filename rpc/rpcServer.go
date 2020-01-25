package rpc

import (
	"equb2/DistributedEqub/config"
	db2 "equb2/DistributedEqub/db"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type SERVER int

type Result struct {
	Message string
	Equb    db2.Equb
}

func Serve() {
	serverPort := config.Port

	server := new(SERVER)
	err := rpc.Register(server)
	if err != nil {
		log.Fatal("Error Registering RPC", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+serverPort)

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

func (SERVER) StartClient(member db2.Member, result *Result) error {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	defer db.Close()

	if equb.Status == "started" {
		*result = Result{
			Message: "Cannot Join This Equb",
			Equb:    db2.Equb{},
		}
	} else {
		member.EqubID = equb.ID
		member.CreateMember(db, equb)

		equb = db2.FindEqub(db)[0]
		equb.SetNextServer(db, member)

		*result = Result{
			Message: "Successfully Joined",
			Equb:    equb,
		}
		log.Println(member.Name + " connected")
	}

	return nil
}
