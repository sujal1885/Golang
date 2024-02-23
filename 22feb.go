package main
import "fmt"

// defining a function
// func greet(){
// 	fmt.Println("Good Morning")
// }

// func addNumbers(){
// 	n1:=12
// 	n2:=8

// 	sum:=n1+n2
// 	fmt.Println("Sum: ", sum)
// }

// func addNumbers(n1 int,n2 int){
// 	sum:=n1+n2
// 	fmt.Println("Sum: ", sum)
// }

// func addNumbers(n1 int,n2 int) int {
// 	sum:=n1+n2
// 	return sum
// }

// func addNumbers1(n1 int,n2 int) int {
// 	sum:=n1+n2
// 	return sum
// 	fmt.Println("After the return statement")
// }

// func calculate(n1 int,n2 int)(int,int){
// 	sum:=n1+n2
// 	difference := n1-n2
// 	return sum , difference
// }

// func getSquare(num int){
// 	square := num*num
// 	fmt.Printf("Square of %d is %d\n", num, square)
// }


// func addNumbers2(){
// 	var sum int
// 	sum = 14
// }

func countDown(number int){
	if(number<=0){
		fmt.Println("Countdown Stops\n");
		return
	}
	fmt.Println(number)
	countDown(number-1)
}

func main(){
	// greet()
	// addNumbers()
	// var x ,y int
	// fmt.Println("Enter two numbers : ")
	// fmt.Scan(&x,&y)
	// addNumbers(x,y)

	// var x ,y int
	// fmt.Println("Enter two numbers : ")
	// fmt.Scan(&x,&y)
	// fmt.Println("The sum of  the two numbers is :",addNumbers(x,y))

	// addNumbers1(5,6)
	// var x ,y int
	// fmt.Println("Enter two numbers : ")
	// fmt.Scan(&x,&y)
	// sum,difference := calculate(x,y)
	// fmt.Println("Sum: ",sum,"\nDifference: ",difference)

	// getSquare(5)

	// var sum int
	// sum = 0
	// addNumbers2()
	// fmt.Println("Sum is  : ",sum)

	countDown(5)
}

