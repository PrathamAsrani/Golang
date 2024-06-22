package main

import(
	"fmt"
	"strings"
)

func main(){
	var str string = "ABC PQR LMN PQR";
	var atr string = "XYZ LMN LMN PQR";
	var ptr string = "hello";
	fmt.Println(strings.ReplaceAll(str, "PQR", "TUV"));
	fmt.Println(strings.ReplaceAll(atr, "LMN", "WXY"));

	fmt.Println(strings.Title(ptr));	
	for i:= 0; i < len(str); i++ {
		fmt.Printf("%c\t%d\n", str[i], str[i]);
	}

	var array = [5] int {1, 5, 8, 0, 3}
	fmt.Println(array)
}