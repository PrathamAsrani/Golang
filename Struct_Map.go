package main

import(
	"fmt"
)

type Vertex struct{
	Lat, Long float64
}

func main(){
	var m = map[string]Vertex {
		"Bell Labs": {
			40.68, -74.39,
		},
		"Google": {
			37.422, -122.084,
		},
	}
	fmt.Println("Map: ",m);

	for _, val := range m {
		fmt.Printf("%v: %v\n", val.Lat, val.Long);
	}

	flowerColor := map[string]string {"Sunflower": "Yellow"};

	delete(flowerColor, "Sunflower");
	val, flag := flowerColor["Sunflower"];
	fmt.Println(val, flag);
}