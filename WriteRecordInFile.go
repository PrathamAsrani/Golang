// Create a text file
// Read and write text file
// Copy a text file into new text file
// Create a zip file
// Write a structure of binary data into a file

package main

import (
	"archive/zip"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Student struct {
	age  uint8
	cgpa float64
}

func WriteData() {
	file, err := os.Create("Sample.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("Unable to create file: %s\n", err)
		panic(err)
	}

	len, err := file.WriteString("Hello, world!")

	if err != nil {
		fmt.Printf("Unable to write data: %s\n", err)
		panic(err)
	}

	fmt.Printf("%d letters written in the file\n", len)
}

func ReadData() {
	textData, err := ioutil.ReadFile("Sample.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nData in file is %s\n", textData)
}
func main() {
	WriteData()
	ReadData()

	MyFile, err := os.Stat("Sample.txt")
	if err != nil {
		fmt.Println("File not found")
	}

	fmt.Println("File Permissions: ", MyFile.Mode())
	fmt.Println("File Permissions: ", MyFile.ModTime())
	fmt.Printf("Size of file in bytes: %d\n", MyFile.Size())

	fmt.Println()
	fmt.Printf("Copying sample to copy\n")
	CopyFile()
	fmt.Println()

	ZipCreate()
	fmt.Println()

	BinaryCreate()
	fmt.Println()

}

func CopyFile() {
	existingFile, err := os.Open("Sample.txt")
	defer existingFile.Close()

	if err != nil {
		fmt.Println("Unable to open file: ", err)
	}
	copyFile, err := os.Create("Copy.txt")
	if err != nil {
		fmt.Println("Unable to create copy file: ", err)
	}
	defer copyFile.Close()

	len, err := io.Copy(copyFile, existingFile)
	if err != nil {
		fmt.Println("Unable to copy file: ", err)
	}

	fmt.Printf("%d bytes copied successfully\n", len)
}

func ZipCreate() {
	filePtr, err := os.Create("Empty.zip")
	if err != nil {
		fmt.Println("Unable to create: ", err)
	}

	// Create a zip writter object using file pointer
	MyZipWriter := zip.NewWriter(filePtr)

	err = MyZipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}

	filePtr.Close()
	fmt.Println("Empty.zip created successfully")
}

func BinaryCreate() {
	file, err := os.Create("data.bin")
	if err != nil {
		fmt.Println("Unable to create")
	}

	var st = Student{20, 7.5}

	err = binary.Write(file, binary.LittleEndian, st)
	if err != nil {
		fmt.Println("Write failed: ", err)
	}
	fmt.Println("Structure written into file successfuly")
	file.Close()
}
