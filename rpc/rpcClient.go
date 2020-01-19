package rpc

import (
	"equb/config"
	db2 "equb/db"
	"log"
	"net/rpc"
)

func GetClient() *rpc.Client{
	client, err := rpc.DialHTTP("tcp", config.ServerIP + ":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func StartClient(member db2.Member) db2.Equb{
	client := GetClient()

	var result db2.Equb
	err2 := client.Call("SERVER.StartClient", member, &result)
	if err2 != nil {
		log.Println(err2)
	}

	return result
}