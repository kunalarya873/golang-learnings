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
	var role = "admin"
	var hasPermissions = false
	if role == "admin" && hasPermissions {
		fmt.Println("You are an admin with permissions")
	} else{
		fmt.Println("You are not an admin with permissions")
	}
	if age:=15 ; age < 10 {
		fmt.Println("You are a kid")
	}
}