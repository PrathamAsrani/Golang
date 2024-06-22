package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func test(x int) {
	fmt.Println("hello", x)
}

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	} else {
		return 1
	}
}

func summation(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + summation(n-1)
	}
}

var (
	sum int
)

// global var
const global string = "I am a global string"

const (
	typed     int     = 100
	untyped_1         = 'A'
	untyped_2         = 1.234
	PI        float32 = 3.14
)

func Registration() {
	var roll_number int
	var student_name string
	var paper_one float32
	var paper_two float32
	var paper_three float32
	fmt.Println("\n-----------------------------------------------------------------------------------------------------------\n")
	fmt.Println("Enter your name: ")
	fmt.Scanln(&student_name)
	fmt.Println("Enter your roll number: ")
	fmt.Scan(&roll_number)
	fmt.Println("Enter your paper1 marks: ")
	fmt.Scan(&paper_one)
	fmt.Println("Enter your paper2 marks: ")
	fmt.Scan(&paper_two)
	fmt.Println("Enter your paper3 marks: ")
	fmt.Scan(&paper_three)

	fmt.Println("\n-----------------------------------------------------------------------------------------------------------\n")
	fmt.Println("Student Report: \nStudent name :", student_name, "\troll number: ", roll_number, "\nPaper1:", paper_one, "\nPaper2:", paper_two, "\nPaper3:", paper_three, "\nTotal:", (paper_one+paper_two+paper_three)/300*100, "%")
	fmt.Println("\n-----------------------------------------------------------------------------------------------------------\n")
}

func greet(name string) {
	var displayName = func() {
		fmt.Println(name)
	}
	displayName()
}

func add() {
	var a, b int
	// fmt.Scanln(&a, &b)
	a = 1
	b = 2
	fmt.Println("ans:", a+b)
}

func area_circle(radius float32) float32 {
	return PI * radius
}

func area_rectangle(a int, b int) int {
	return (a * b)
}

func fun(n1, n2 int) (int, float32) {
	// sum := 1
	// return n1 + n2, float32((n1+n2)/2);
	return n2, float32(n1)
}

func dispalyNumber() func() int {
	number := 10
	return func() int {
		number++
		return number
	}
}

func calculated() func() int {
	num := 1
	return func() int {
		num = num + 2
		return num
	}
}

func function(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("\n", total)
}

func main() {
	println("hii I am println")
	print("hii I am print\n")
	fmt.Printf("HelloWorld")
	fmt.Print("\tHelloWorld")
	fmt.Println("\tHelloWorld")

	// static
	var a int = 100

	// inferred
	x := 10
	b := 25.5

	fmt.Println(a, x, b)

	// pascal case
	FirstName := "Pratham"

	// Camel Case
	lastName := "Asrani"

	// snake case
	my_full_name := "Pratham Ashok Asrani"

	fmt.Println(FirstName, "\n", lastName, "\n", my_full_name)

	fmt.Println("PI:", PI)

	// typed constant
	const A int = 1

	// untyped constant
	const B = 1

	// fun(3)
	var num int
	print("\nEnter the number: ")
	fmt.Scan(&num)
	fmt.Println("Factorial of 5 is", factorial(num))
	fmt.Println("Summation of 5 is", summation(num))

	// A = 2

	fmt.Println(A, B, global)
	fmt.Println(typed, untyped_1, untyped_2)

	fmt.Printf("The answer is %v\n", 42) // %v prints the value of an untyped constant
	fmt.Printf("float %f\n", untyped_2)
	fmt.Printf("int %d\n", typed)
	fmt.Printf("String %s\n", my_full_name)

	add()

	fmt.Println("area of rec:", area_rectangle(2, 3))
	fmt.Println("area of circle:", area_circle(2))

	// Registration()

	num1 := 5
	num2 := 10
	sum, avg := fun(num1, num2)
	fmt.Println("sum:", sum, "avg:", avg)

	// Password to be hashed
	password := "MySecurePassword123"
	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Output the hashed password
	fmt.Println("Original Password:", password)
	fmt.Println("Hashed Password: ", string(hashedPassword))
	// Comparing a password attempt with the hashed password passwordAttempt := "MySecurePassword123"
	passwordAttempt := "MySecurePassword123"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(passwordAttempt))
	if err == nil {
		fmt.Println("Password match!")
	} else {
		fmt.Println("Password does not match!")
	}

	test(3)
	_test := func(x int) {
		fmt.Println("Hello", x)
	}
	__test := func(x int) int {
		return x * 2
	}(8)
	_test(4)
	fmt.Println(__test)

	m1 := md5.Sum([]byte("Hello World"))
	m2 := sha256.Sum256([]byte("hii"))
	fmt.Println(m1)
	fmt.Println(m2)

	summation := func(n1 int, n2 int) int {
		return n1 + n2
	}
	findSquare := func(n1 int) int {
		return n1 * n1
	}
	fmt.Println(findSquare(summation(7, 9)))

	_a := dispalyNumber()
	fmt.Println(_a())

	greet("John")

	cal := calculated()
	fmt.Println(cal())
	fmt.Println(cal())
	fmt.Println(cal())

	YYYY, MM, DD := time.Now().Date()
	Hour, Minute, Second := time.Now().Clock()
	fmt.Println(YYYY, MM, DD)
	fmt.Println(Hour, Minute, Second)
	i := 1
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}

	// var n1, n2, tmp uint32 = 0, 0, 0
	// if _, err := fmt.Scan(&n1, &n2); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// for n2 != 0 {
	// 	tmp = n1 % n2
	// 	n1 = n2
	// 	n2 = tmp
	// }
	// fmt.Println("GCD", n1)

	n3, n4 := 1, 2
	if n3 > n4 {
		fmt.Println("n3 is greater")
	} else {
		fmt.Println("n4 is greater or equal to n3")
	}

	// var n5 int
	// fmt.Scan(&n5)
	// if n5&1 == 1 {
	// 	fmt.Println("odd")
	// } else {
	// 	fmt.Println("even")
	// }

	n6, n7 := 8, 10
	if n6 != n7 {
		fmt.Printf("%d is not equal to %d\n", n6, n7)
	} else {
		fmt.Printf("%d is equal to %d\n", n6, n7)
	}

	// // Learning to call an api
	// apiUrl := "https://api.example.com/data"

	// response, err := http.Get(apiUrl)
	// if err != nil {
	// 	fmt.Printf("Error making Get request : %s \n", err)
	// 	return
	// }
	// defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Printf("Error reading resp %s \n", err)
	// 	return
	// }
	// fmt.Println(string(body))

	switch 5 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}

	function(10, 20)
	function(10, 20, 30)
	function(10, 20, 30, 40)
	nums := []int{10, 20, 30, 40}
	function(nums...)

	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean Your House")
	case today.Day() <= 10:
		fmt.Println("Buy Some Flowers")
	case today.Day() > 15:
		fmt.Println("Go for a Run")
	case today.Day() == 25:
		fmt.Println("Buy some food")
	default:
		fmt.Println(today)
	}

	switch x1 := 1; x1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	var x2 interface{} = "RKNEC"
	switch itr := x2.(type) {
	case int:
		fmt.Println("int", itr)
	case float64, float32:
		fmt.Println("float")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	}

	itr := 0
loop:
	fmt.Println(itr)
	itr++
	if itr < 5 {
		goto loop
	}
	fmt.Println("lopp ends here")

	var1 := 0
while_loop:
	if var1 < 5 {
		fmt.Println(var1)
		var1++
		goto while_loop
	}

	var2 := 0
	for var2 < 5 {
		fmt.Println(var2)
		var2++
	}

	var2 = 0
	for {
		fmt.Println(var2)
		var2++
		if var2 > 4 {
			break
		}
	}

	/*
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		// Using the Config value, create the DynamoDB client
		svc := dynamodb.NewFromConfig(cfg)

		// Build the request with its input parameters
		resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
			Limit: aws.Int32(5),
		})
		if err != nil {
			log.Fatalf("failed to list tables, %v", err)
		}
		fmt.Println("Tables:")
		for _, tableName := range resp.TableNames {
			fmt.Println(tableName)
		}
	*/

	var str string = "Hello world"
	var sub_str = "world"
	if strings.Contains(str, sub_str) == true {
		fmt.Println("String", str, " contains substr:", sub_str)
	} else {
		fmt.Println("String", str, " does not contain substr:", sub_str)
	}

	fmt.Println("string in uppercase:", strings.ToUpper(str))
	fmt.Println("string in lowercase:", strings.ToLower(str))

	str = "India-is-a-great-country"
	var array []string = strings.Split(str, "-")
	fmt.Println("Splited Strings:")
	for i := 0; i < len(array); i++ {
		if i < len(array)-1 {
			fmt.Print(array[i], ", ")
		} else {
			fmt.Print(array[i], "\n")
		}
	}

	final := strings.Repeat(str, 4)
	fmt.Println(final, "\t")

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	println()

	final = strings.Replace(str, "i", "I", -1)
	fmt.Println(final, "\t")
	final = strings.Replace(str, "i", "I", 1)
	fmt.Println(final, "\t")
	final = strings.Replace(str, "i", "I", 2)
	fmt.Println(final, "\t")

	str_arr := []string{"India", "is", "a", "great", "country"}
	final = strings.Join(str_arr, "-")
	fmt.Println(final)
}
