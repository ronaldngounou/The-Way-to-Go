package main 
import "fmt"
func main(){
	var fName string 
	fmt.Println("Enter first name: ")
	fmt.Scanln(&fName)

	fmt.Print(fName + "\n")
}