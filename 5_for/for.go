package main
import "fmt"
func main(){
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i:= 0; i<3; i++ {
		fmt.Println(i)
	}
	//Shorthand
	for i:= range(3){
		fmt.Println(i)
	}
}