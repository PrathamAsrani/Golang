package main

import(
	"fmt"
)

func show(){
	fmt.Println("in show func");
}

func main(){
	defer show();
	fmt.Println("line 1");
	fmt.Println("line 2");
	defer display();
}

func display(){
	fmt.Println("in display func");
}