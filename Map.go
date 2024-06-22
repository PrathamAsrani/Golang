package main

import(
	"fmt"
)

func main(){
	subjectsMarks := map[string]float64{"Golang": 100, "Cpp": 100, "Java": 100, "All": 100};
	fmt.Println(subjectsMarks["Golang"]);
	fmt.Println(subjectsMarks);

	flowerColor := map[string]string {"Sunflower": "Yellow"};
	fmt.Println(flowerColor["Sunflower"]);


	/*Maps are mutuable*/
	capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"};
	fmt.Println("Initial map: ", capital);

	fmt.Println("value before update: ", capital["US"]);
	capital["US"] = "Washington";
	fmt.Println("value after update: ", capital["US"]);

	delete(capital,"Nepal");
	fmt.Println(capital);

	squaredNumbers := map[int]int{2 : 4, 3 : 9, 4 : 16, 5 : 25};
	for num, sqr := range squaredNumbers {
		fmt.Printf("Square of %d is %d\n", num, sqr);
	}

	student := make(map[int]string)
	student[1] = "Pratham Asrani"
	student[2] = "John Doe"
	student[3] = "Honey Singh"
	fmt.Println(student)
	
}
