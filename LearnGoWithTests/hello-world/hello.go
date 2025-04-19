package main 

<<<<<<< HEAD
const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(word string, language string) string {
	if word == ""{
		word = "World"
	}
	
	return greetingPrefix(language) + word
}


func greetingPrefix(language string) (prefix string){
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
=======
import "fmt"

func Hello(word string) string {
	return "Hello " + word
}

func main() {
	fmt.Println(Hello("Jean"))
>>>>>>> 1adb18a9f4bd138a596c4f16baba006504428d5a
}

