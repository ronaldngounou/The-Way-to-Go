package main 
import "fmt"
func main(){
	// var fName string 
	// fmt.Println("Enter first name: ")
	// fmt.Scanln(&fName)

	// fmt.Print(fName + "\n")

	// Pointers
	var i1 = 5
	var ptr1 *int 
	ptr1 = &i1
	fmt.Print("The address of i1: %p", ptr1)
	var i2 = 10
	var ptr2 *int 
	ptr2 = &i2
	fmt.Print("The address of i2: %p \n", ptr2)

}