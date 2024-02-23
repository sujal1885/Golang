package main
import "fmt"
// import "math"
// import "io"
// import "os"

func main(){
	// var dd int = 17
	// var mm int = 04
	// var yy int = 2021
	// var str string
	// str = fmt.Sprintf("%02d-%02d-%04d",dd,mm,yy)
	// io.WriteString(os.Stdout,str)

	// var str string
	// fmt.Print("Enter a string: ")
	// fmt.Scan(str)
	// fmt.Printf("Result:%s\n",str)

	// var in int
	// fmt.Print("Enter a int: ")
	// fmt.Scan(&in)
	// fmt.Printf("Result:%d\n",in)

	// var str1 string
	// var str2 string
	// var in int
	// fmt.Print("Enter two strings and int ")
	// fmt.Scan(&str1,&str2,&in)
	// fmt.Printf("Result:%s , Result:%s , Result:%d",str1,str2,in)

	
	// var a int = 123
	// var b uint = 0
	// // Assign the value of a to b
	// b = (uint)(a)
	// fmt.Printf("a= %d, b=%d\n",a,b)

	// var a int = 64
	// var b uint = 0

	// b = uint(math.Sqrt(float64(a)))
	// fmt.Printf("b=%d\n",b)

	// var x int = 225
	// var r float32
	// r = float32(math.Sqrt(float64(x)))
	// fmt.Println("The square root of ",x," is ",r)


	// explicit type conversion
	var x int = 42
	var y float64 = float64(x)
	var z uint = uint(y)
	fmt.Printf("Value of x is %d and type is %T \n",x,x)
	fmt.Printf("Value of y is %.2f and type is %T \n",y,y)
	fmt.Printf("Value of z is %d and type is %T \n",z,z)
}