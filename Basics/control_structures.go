package main 
import "fmt"
func ageCondition(){
	age := 10

	if age < 18 {
		fmt.Print("You are a young person.")
	} else {
		fmt.Print("You are an adult.")
	}
}

