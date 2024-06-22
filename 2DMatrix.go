package main

import(
	"fmt"
)

func main(){
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("\nThe entered Matrix is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j]);
		}
		fmt.Println();
	}
	fmt.Println("\nThe entered Matrix left diagonal is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if(i == j) {
				fmt.Printf("%d ", matrix[i][j]);
			}else{
				fmt.Printf(" ");
			}
		}
		fmt.Println();
	}
	fmt.Println("\nThe entered Matrix right diagonal is: \n");
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if(i + j == 2) {
				fmt.Printf("%d ", matrix[i][j]);
			}else{
				fmt.Printf(" ");
			}
		}
		fmt.Println();
	}

}