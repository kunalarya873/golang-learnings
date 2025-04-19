package main
import "fmt"


func main(){
	age := 20
	if age < 10 {
		fmt.Println("You are a kid")
	} else if age < 20 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
	}
}