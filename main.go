package main
import (
	"fmt"
	"strings"
)

func main(){
	// 3 different ways to declare variables in Go
	// var name = "Willow"
	// name2 := "Dan"
	// var name3 string = "Charles"

	// Multiple variable declaration
	var one, two, three, four, five int = 1, 2, 3, 4, 5 

	var(
		num int
		number int = 1
		greetings string = "hello"
	)
	// age := 2
	// isRaining := true
	fmt.Println(one, two, three, four, five)
	fmt.Println(num)
	fmt.Println(number)
	fmt.Println(greetings)
	
	// fmt.Println(age)
	// fmt.Println(isRaining)

	msg := "Golang Programming Language"
	msg2 := "Golang"	
	//contains(value used as the library, the var you are checking exists in the library)
	fmt.Println(strings.Contains(msg, msg2))

	stringOne := "CAPITAL LETTERS"
	fmt.Println(strings.ToLower(stringOne))

	//conditional statememts
	password := "1234"
	// var foods 
	// food := ["pizza", "meatpie", "jollof"]
	if len(password) > 7{
		fmt.Println("Password is secure")
	}else{
		fmt.Println("Password is not secure")
	}

	// format specifiers accessing a specific char in a string
	message := "One Piece"
	fmt.Printf("%c", message[0])

	// length of a string
	fmt.Println(len(message)) 


	// switch statements
	num2 := 1000

	switch{
		case(num2 > 5):
			fmt.Println("Number is greater than 5")
		case(num2 < 5):
			fmt.Println("Number is less than 5")
	}

	 
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i, day := range days{
		fmt.Println(i, day)
	}
}