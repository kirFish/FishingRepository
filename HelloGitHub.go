package main

import(

	"fmt"

)


func main(){

	var myGreetings string = "Hello"
	var adress string = "GitHub"

	fmt.Println(myGreetings + " " + adress)

	sum := 1
	for( sum < 10 ){

		sum = sum + 1
		fmt.Println(myGreetings + "," + adress)

	}


}