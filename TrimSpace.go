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
	var str string = "  India  ";
	var str2 string = "!%India!%";
	fmt.Println(strings.TrimSpace(str));
	fmt.Println(strings.Trim(str2, "!%"));

	a:= "12";
	b:= "32";
	c:= "06";
	d:= "12";
	fmt.Println(strings.Compare(a, b));
	fmt.Println(strings.Compare(b, a));
	fmt.Println(strings.Compare(a, c));
	fmt.Println(strings.Compare(a, d));
	fmt.Println(strings.Compare(c, a));
}
