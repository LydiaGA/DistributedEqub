package rpc

import (
	"equb1/DistributedEqub/config"
	db2 "equb1/DistributedEqub/db"
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
	serverPort := config.ServerPort

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

	NotifyAll(equb.Members, equb)
	return nil
}

func (SERVER) GetEqub(member db2.Member, result *Result) error {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	defer db.Close()

	equb.SetNextServer(db, member)

	*result = Result{
		Message: "Successfully Retrieved",
		Equb:    equb,
	}

	NotifyAll(equb.Members, equb)
	return nil
}

func (SERVER) MakePayment(member db2.Member, result *Result) error {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	defer db.Close()

	equb.Total = equb.Total + member.Amount

	memberFound := db2.FindMember(db, member.ID)
	memberFound.HasPaid = true
	db.Save(&memberFound)

	equb.SetNextServer(db, member)

	*result = Result{
		Message: "Successfully Retrieved",
		Equb:    equb,
	}
	NotifyAll(equb.Members, equb)
	return nil
}

func (SERVER) CollectWinnings(member db2.Member, result *Result) error {
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	defer db.Close()

	if equb.Winner == member {
		equb.Total = equb.Total - (12 * member.Amount)
		equb.SetNextServer(db, member)

		*result = Result{
			Message: "Successfully Retrieved",
			Equb:    equb,
		}
	} else {
		equb.SetNextServer(db, member)

		*result = Result{
			Message: "You are not this month's winner",
			Equb:    db2.Equb{},
		}
	}
	NotifyAll(equb.Members, equb)
	return nil
}

func NotifyAll(members []db2.Member, equb db2.Equb) {
	for _, member := range members {
		client, err := rpc.DialHTTP("tcp", member.IP+":"+member.Port)
		if err != nil {
			log.Println(err)
		}

		var result Result
		err2 := client.Call("CLIENT.Notify", equb, &result)
		if err2 != nil {
			log.Println(err2)
		}
	}
}
