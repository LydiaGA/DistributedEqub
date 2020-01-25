package config

import (
	"equb1/DistributedEqub/db"
	"log"
	"net"
)

var Port = "8081"
var IP = getIp().String()
var ServerIP = getIp().String()
var Me = db.Member{}

func getIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
