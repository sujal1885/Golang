package main
import "fmt"

func sum(number int) int {
	if number<=0{
		return 0
	}else{
		return number+sum(number-1)
	}
}

// Anonymous function global
var  show = func()  {
	fmt.Println("I won't show");
}

// var fact = func(n int) int {
// 	if n == 0 {
// 		return 1
// 	}else{
// 		return n*fact(n-1)
// 	}
// }

func findSquare(num int) int {
	square := num*num
	return square
}



func main(){
	// var num int
	// fmt.Println("Enter a number : ")
	// fmt.Scan(&num)
	// if num<0{
	// 	fmt.Println("Number can't be negative");
	// }else{
	// 	fmt.Println("The sum of all natural numbers upto entered number is ",sum(num))
	// }

	// Anonymous function
	// var greet = func(){
	// 	fmt.Println("Hello World")
	// }

	// greet()
	// var watch = show
	// show()
	// watch()

	// anonymous function with parameters
	// var sum = func(n1 int,n2 int){
	// 	sum := n1+n2
	// 	fmt.Println("Sum is : ",sum)
	// }
	// sum(5,3)
	// fmt.Printf("The datatype of sum is %T",sum)


	//factorial anonymous function
	// fmt.Println("The factorial value is ",fact(5))

	// Anonymous function with return value
	// var sum = func(n1,n2 int) int {
	// 	sum := n1+n2
	// 	return sum
	// }

	// fmt.Println("Sum is ",sum(4,6))

	fmt.Println("The square is ",findSquare(sum(5)))
}