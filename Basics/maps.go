package main 
import "fmt"

func MapsTutorial() {
	mapLit := map[string]int{"one": 1, "two": 2}
	var mapAssigned map[string]int 
	mapCreated := make(map[string]float32)

	mapAssigned = mapLit 
	mapCreated["Tom"] = 1238120
	mapCreated["John"] = 1540943
	mapCreated["Paul"] = 8093232
	fmt.Println(mapAssigned)
	fmt.Printf("%d\n", mapLit["two"])
 	
}