package main

import(
	"fmt"
)

func main(){
	matrix1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var ans[3][3] int;
	fmt.Println("\nThe entered Matrix is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ans[i][j] = matrix1[i][j] + matrix2[i][j];
		}
	}
	fmt.Println("\nThe Matrix1 is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix1[i][j]);
		}
		fmt.Println();
	}
	fmt.Println("\nThe Matrix2 is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix2[i][j]);
		}
		fmt.Println();
	}
	fmt.Println("\nThe ans is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", ans[i][j]);
		}
		fmt.Println();
	}
}