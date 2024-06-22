package main

import(
	"fmt"
	"io"
	"os"
)

func main(){
	f, err:= os.Open("Sample.txt");
	if err != nil {
		fmt.Println(err);
	}
	c, err := os.Create("Copy.txt")
	if err != nil {
		fmt.Println(err);
	}
	len, err := io.Copy(c, f);
	if err != nil {
		fmt.Println(err);
	}
	fmt.Printf("Copied %d bytes\n", len);
	c.Close();
	f.Close();
}