package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CreateMultTable(limit int) {
	welcome := "Welcome to user create mult table"
	fmt.Println(welcome)

	reader := bufio.NewScanner(os.Stdin)

	// comma ok || err ok
	var number int
	var err error
	for {
		fmt.Println("Enter the number to create the table ")
		if reader.Scan() {
			number, err = strconv.Atoi(reader.Text())
			if err != nil {
				continue
			}
			break

		}

	}

	fmt.Printf("Table from 1 to %d is:\n", limit)
	for i := 1; i <= limit; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
