package main
import "fmt"

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

	
	
}