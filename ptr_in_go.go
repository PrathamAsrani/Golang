package main

import(
	"fmt"
)

func main(){
	a := 10
	var ptr *int = &a
	fmt.Println(ptr);

	fmt.Println(a);
	*ptr = 30;
	fmt.Println(a);

	name := "name";
	var ptr_to_string *string = &name;

	fmt.Println("Value of the ptr_to_string: ", ptr_to_string);
	fmt.Println("Value of the ptr_to_string: ", *ptr_to_string);
}