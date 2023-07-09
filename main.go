package main

import (
	"fmt"
	"os"
)

func main() {
	for true {
		var cmdln string
		fmt.Println("请输入命令：")
		fmt.Scanln(&cmdln)

		switch cmdln {
		case "help":
			fmt.Println("There are commands as follows")
			fmt.Println("help、show、laugh")
			break
		case "show":
			fmt.Println("You R Sooooo Good!")
			break
		case "laugh":
			fmt.Println("HHHHHHHHHHHHHHHHHH")
			break
		case "quit":
			fmt.Println("Awwwww, R U gonna leave me?")
			os.Exit(0)
		default:
			fmt.Println("HAHA, VERY FUNNY, DO THAT AGAIN")
		}
	}
}
