package main 

const englishHelloPrefix = "Hello, "

func Hello(word string) string {
	if word == ""{
		word = "World"
	}
	return englishHelloPrefix + word
}


