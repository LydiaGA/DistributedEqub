package main

import (
	db2 "equb/db"
	"fmt"
)

func main(){
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

	var role int
	_, _ = fmt.Scanf("%d", &role)

	if role == 1{
		fmt.Println("Enter the Name of the Equb")

		var name string
		_, _ = fmt.Scanln(&name)

		fmt.Println("Enter the Starting Month")

		var month string
		_, _ = fmt.Scanln(&month)

		go StartServer(name, month)

		fmt.Println("Enter 'exit' to Exit")

		var command string
		_, _ = fmt.Scanln(&command)
	}else if role == 2{
		fmt.Println("Enter Address of Server")

		var address string
		_, _ = fmt.Scanln(&address)

		fmt.Println("Enter your name")

		var name string
		_, _ = fmt.Scanln(&name)

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

		var command string
		_, _ = fmt.Scanln(&command)
	}

	fmt.Println(role)

	//fmt.Println(config.IP)
}