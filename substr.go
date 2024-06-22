package main

import(
	"fmt"
	"strings"
)

func count(s, substr string) int {
    count := 0
    for strings.Index(s, substr) != -1 {
        count++
        s = s[strings.Index(s, substr)+len(substr):]
    }
    return count
}

func main(){
	var str string = "India";
	var substr string = "In";
	if strings.HasPrefix(str, substr) == true {
		fmt.Println("The specified substring", substr, "is prefix of", str);
	}else{
		fmt.Println("The given string does not starts with the specified substring")
	}
	if strings.HasSuffix(str, substr) == true {
		fmt.Println("The specified substring", substr, "is suffix of", str);
	}else{
		fmt.Println("The given string does not starts with the specified substring")
	}

	cnt := count(str, "i");
	fmt.Println(cnt)
}