package main

import (
	"bufio"
	db2 "equb1/DistributedEqub/db"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//Serve()

	db2.Migrate()
	//db := db2.GetDatabase()
	//member1 := db2.Member{
	//	Name:       "Member 1",
	//	ServerTurn: 1,
	//	HasPaid: false,
	//	Amount: 2000,
	//	IP: config.IP,
	//}
	//
	//member2 := db2.Member{
	//	Name:       "Member 2",
	//	ServerTurn: 1,
	//	HasPaid: false,
	//	Amount: 1000,
	//	IP: config.IP,
	//}
	//
	//members := make([]db2.Member, 2)
	//members[0] = member1
	//members[1] = member2
	//
	//equb := db2.Equb{Name: "Equb 2", CurrentMonth: "January", Members: members}
	//equb.CreateEqub(db)
	//db.Close()

	fmt.Println("Do you want to create a new equb or join an existing one?")
	fmt.Println("(1) Create")
	fmt.Println("(2) Join")

	in := bufio.NewReader(os.Stdin)

	var role int
	_, _ = fmt.Scanf("%d", &role)

	if role == 1 {
		fmt.Println("Enter the Name of the Equb")

		var name string
		name, _ = in.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")

		fmt.Println("Enter the Starting Month")

		var month string
		month, _ = in.ReadString('\n')
		month = strings.TrimSuffix(month, "\n")

		fmt.Println("Enter 'start' to Start the Equb or 'exit' to Exit")

		time.Sleep(2 * time.Second)

		go StartServer(name, month)

		var command string
		command, _ = in.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		if command == "start" {
			StartEqub()
		} else if command == "exit" {

		}

	} else if role == 2 {
		fmt.Println("Enter Address of Server")

		var address string
		address, _ = in.ReadString('\n')
		address = strings.TrimSuffix(address, "\n")

		fmt.Println("Enter your name")

		var name string
		name, _ = in.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")

		fmt.Println("Enter the Amount you are Going to Pay")

		var amount int
		_, _ = fmt.Scanf("%d", &amount)

		StartClient(address, name, amount)

		for true {
			fmt.Println("What would you like to do?")
			fmt.Println("(1) Make Payment")
			fmt.Println("(2) Collect Winnings")
			fmt.Println("(3) View Total")
			fmt.Println("(4) View Member List")
			fmt.Println("(5) Exit")

			var action int
			_, _ = fmt.Scanf("%d", &action)

			switch action {
			case 1:
				fmt.Println(action)
			case 2:
				fmt.Println(action)
			case 3:
				fmt.Println(action)
			case 4:
				fmt.Println(action)
			case 5:
				fmt.Println(action)
			}
		}

		//var command string
		//command, _ = in.ReadString('\n')
	}
}
