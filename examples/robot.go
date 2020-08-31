package examples

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunRobot() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name: ")

	// 读取数据直到碰到\n为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		// 异常退出
		os.Exit(1)
	} else {
		// 用切片操作删除最后的\n
		name := input[:len(input)-2]
		fmt.Printf("Hello, %s! What can i do for you?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}

		input = input[:len(input)-1]
		input = strings.ToLower(input)
		// fmt.Println(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, i didn't catch you.")
		}
	}
}
