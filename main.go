package main

import (
	"bufio"
	db2 "equb2/DistributedEqub/db"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	db2.Migrate()

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

		fmt.Println("Enter the Starting Month(Use Numbers)")

		var month int
		_, _ = fmt.Scanf("%d", &month)

		time.Sleep(2 * time.Second)

		go StartServer(name, month)

		for true {
			fmt.Println("What would you like to do?")
			fmt.Println("(1) Make Payment")
			fmt.Println("(2) Collect Winnings")
			fmt.Println("(3) View Total")
			fmt.Println("(4) View Member List")
			fmt.Println("(5) Change Month")
			fmt.Println("(6) Exit")

			var action int
			_, _ = fmt.Scanf("%d", &action)

			switch action {
			case 1:
				fmt.Println(action)
				fmt.Println(MakePayment())
			case 2:
				fmt.Println(action)
				fmt.Println(CollectWinnings())
			case 3:
				fmt.Println(action)
				total := GetTotal()
				fmt.Println("Total: " + strconv.FormatInt(int64(total), 10))
			case 4:
				fmt.Println(action)
				members := GetList()
				for _, member := range members {
					fmt.Println(member.Name + "\t" + strconv.FormatInt(int64(member.Amount), 10) + "\t" + strconv.FormatBool(member.HasPaid))
				}
			case 5:
				fmt.Println(action)
				fmt.Println(ChangeMonth())
			case 6:
				fmt.Println(action)
			}
		}

		//var command string
		//command, _ = in.ReadString('\n')
		//command = strings.TrimSuffix(command, "\n")
		//
		//if command == "start" {
		//	StartEqub()
		//} else if command == "exit" {
		//
		//}

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
				fmt.Println(MakePayment())
			case 2:
				fmt.Println(action)
				fmt.Println(CollectWinnings())
			case 3:
				fmt.Println(action)
				total := GetTotal()
				fmt.Println("Total: " + strconv.FormatInt(int64(total), 10))
			case 4:
				fmt.Println(action)
				members := GetList()
				for _, member := range members {
					fmt.Println(member.Name + "\t" + strconv.FormatInt(int64(member.Amount), 10) + "\t" + strconv.FormatBool(member.HasPaid))
				}
			case 5:
				fmt.Println(action)
			}
		}

		//var command string
		//command, _ = in.ReadString('\n')
	}
}
