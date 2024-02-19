package main
import "fmt"
// import "math"

func main(){
	// integer:=23
	// fmt.Printf("%T %T\n",integer,&integer)

	// number , floatingNumber := 238,1234.575883939
	// fmt.Printf("Default: %f \n",floatingNumber)
	// fmt.Printf(".2f: %15.2f\n",floatingNumber)
	// fmt.Printf(".4f: %.4f\n",floatingNumber)
	// fmt.Printf("Scientific: %e \n",floatingNumber)
	// fmt.Printf("Decimal: %d \n",number)
	// fmt.Printf("Binary: %b \n",number)
	// fmt.Printf("Octal: %o \n",number)
	// fmt.Printf("HexaDecimal: %X \n",number)

	// fmt.Printf("%*s\n",40,"text")
	// fmt.Printf("%040s\n","text")
	// text := "Rajesh"
	// fmt.Printf("%*s\n",40,text)

	// var ch byte = 'A'
	// bch := "B"
	// fmt.Printf("%T %T\n",ch,bch)

	// fmt.Printf("%d %c\n",ch,ch)
	
	// abs accepts only float
	// var val float64 = -19.25
	// fmt.Printf("Absolute value of %f is %f",val,math.Abs(val))

	//maximum value between two numbers using math.max()
	// var num1 float64 = 10.25
	// var num2 float64 = 20.14
	// var large float64 = 0

	// large = math.Max(num1,num2)
	// fmt.Printf("Largest number is : %f",large)

	//minimum value between two numbers using math.min()
	// var num1 float64 = 10.25
	// var num2 float64 = 20.14
	// var small float64 = 0

	// small = math.Min(num1,num2)
	// fmt.Printf("Smallest number is : %f",small)

	// var num float64 = 102
	// var p float64 = 3.0
	// var result float64 = 0

	// result = math.Pow(num,p)
	// fmt.Printf("Result: %f",result)

	//ceil

	// var num1 float64 = 10.25
	// var num2 float64 = 20.52

	// var result float64 = 0

	// result = math.Ceil(num1)
	// fmt.Printf("Result is : %f",result)

	// result = math.Ceil(num2)
	// fmt.Printf("\nResult is : %f",result)

	//floor
	// var num1 float64 = 10.25
	// var num2 float64 = 20.52

	// var result float64 = 0

	// result = math.Floor(num1)
	// fmt.Printf("Result is : %f",result)

	// result = math.Floor(num2)
	// fmt.Printf("\nResult is : %f",result)

	// use without importing fmt
	// println("Using println instead of fmt.Println")
	// print("Using print instead of fmt.Print")

	// x:=10
	// y:=2
	// fmt.Println(x+y)

	// x := 10
	// x++
	// fmt.Println(x)

	// var x = 10
	// x+=5
	// fmt.Println(x)

	// scanning input
	var in int
	fmt.Println("Enter an integer")
	fmt.Scan(&in)
	fmt.Println("Entered input is ",in," and its modulo 2 is ",in%2)




}