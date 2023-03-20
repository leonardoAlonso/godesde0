package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the raring for our Pizza: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("tahnks for rating, ", input)
	fmt.Printf("type of this rating is %T ", input)
}
