package rpc

import (
	"equb2/DistributedEqub/config"
	db2 "equb2/DistributedEqub/db"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type CLIENT int

//func ClientServe() {
//	serverPort := config.ClientPort
//
//	server := new(CLIENT)
//	err := rpc.Register(server)
//	if err != nil {
//		log.Fatal("Error Registering RPC", err)
//	}
//
//	rpc.HandleHTTP()
//
//	listener, err := net.Listen("tcp", ":"+serverPort)
//
//	if err != nil {
//		log.Fatal("Listener Error", err)
//	}
//	log.Printf("Serving RPC on port \"%s\"", serverPort)
//	err = http.Serve(listener, nil)
//
//	if err != nil {
//		log.Fatal("Error Serving: ", err)
//	}
//}

func GetClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", config.ServerIP+":"+config.ServerPort)
	db := db2.GetDatabase()
	equb := db2.FindEqub(db)[0]
	defer db.Close()
	for err != nil {
		if equb.NextServerID == config.Me.ID {
			//tell clients
		} else {
			for _, member := range equb.Members {
				time.Sleep(time.Second)
				client, err = rpc.DialHTTP("tcp", member.IP+":"+config.ServerPort)
				fmt.Print("In GetClient: ")
				log.Println(err)
			}
		}
	}
	return client
}

func StartClient(member db2.Member) db2.Equb {
	client := GetClient()

	var result Result
	err2 := client.Call("SERVER.StartClient", member, &result)
	if err2 != nil {
		log.Println(err2)
	}

	if result.Equb.Name == "" {
		log.Fatal(result.Message)
	}

	return result.Equb
}

func GetEqub() db2.Equb {
	client := GetClient()

	var result Result
	err2 := client.Call("SERVER.GetEqub", config.Me, &result)
	if err2 != nil {
		log.Println(err2)
	}

	if result.Equb.Name == "" {
		log.Fatal(result.Message)
	}

	return result.Equb

}

func MakePayment() (string, db2.Equb) {
	client := GetClient()

	var result Result
	err2 := client.Call("SERVER.MakePayment", config.Me, &result)
	if err2 != nil {
		log.Println(err2)
	}

	return result.Message, result.Equb

}

func CollectWinnings() (string, db2.Equb) {
	client := GetClient()

	var result Result
	err2 := client.Call("SERVER.CollectWinnings", config.Me, &result)
	if err2 != nil {
		log.Println(err2)
	}

	return result.Message, result.Equb

}

func (SERVER) Notify(equb db2.Equb, result *string) error {
	db := db2.GetDatabase()
	defer db.Close()
	db2.UpdateEqub(db, equb)
	*result = "Successfully Updated"
	return nil
}
