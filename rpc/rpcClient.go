package rpc

import (
	"equb1/DistributedEqub/config"
	db2 "equb1/DistributedEqub/db"
	"log"
	"net/rpc"
)

func GetClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", config.ServerIP+":"+config.Port)
	if err != nil {
		log.Fatal(err)
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
