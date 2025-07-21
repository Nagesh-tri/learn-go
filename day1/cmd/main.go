package main;

import (
	"fmt"
	"os"
	"hello"
)

func main()  {
	
	// fmt.Printf("Hello, world");
	// using args 
	// fmt.Printf("hello, "+ os.Args[1]);
	//
	/*
	if len(os.Args) > 1 {
		fmt.Printf("Hello, "+os.Args[1]);
	} else{
		fmt.Print("Hello, World !");
	}
	*/

	// do same with packages
	/*
	if len(os.Args) > 1 {
		fmt.Printf(hello.Greet(os.Args[1]));
	} else{
		fmt.Print(hello.Greet( "World !"));
	}
	*/

	//call many
	fmt.Print(hello.GreetMany(os.Args[1:]));



}
