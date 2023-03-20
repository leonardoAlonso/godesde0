package variables

import "fmt"

const LoginToken string = "token" // Public

func Get_variables() {
	var username string = "Leonardo"
	fmt.Println(username)
	fmt.Printf("Variable if of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable if of type: %T \n", isLoggedIn)

	var smallValue uint8 = 10
	fmt.Println(smallValue)
	fmt.Printf("Variable if of type: %T \n", smallValue)

	var floatValue float32 = 10.3
	fmt.Println(floatValue)
	fmt.Printf("Variable if of type: %T \n", floatValue)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable if of type: %T \n", anotherVariable)

	// implicit type

	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable if of type: %T \n", website)

	// no var style

	numberOfUser := 300
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable if of type: %T \n", LoginToken)
}
