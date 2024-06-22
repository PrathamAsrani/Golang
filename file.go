package main

import(
	"fmt"
	"os"
	"io/ioutil"
)

func writeData(){
	file, err := os.Create("Sample.txt")
	if err != nil{
		fmt.Println("File not created");
	}
	fmt.Println("File created");
	len, err := file.WriteString("Hello World");
	if err != nil {
		fmt.Printf("Error writing to the file: %v\n",err);
	}
	fmt.Println("File created", len);
	file.Close();
}

func readData(){
	text, err := ioutil.ReadFile("Sample.txt");
	if err != nil {
		fmt.Printf("Error reading from the file: %v\n", err)
	}
	fmt.Printf("%s", text);
}

func main(){
	writeData();
	readData();
}