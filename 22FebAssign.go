package main
import "fmt"
                                                                                                                                                
func Fibonacci(num int) int {                                                       
    if(num<=1){
		return num
	}
	return Fibonacci(num-1)+Fibonacci(num-2)                                                    
}    

func factorial(n int) int {
	if (n == 1) { 
		return 1
	}
	return n * factorial(n-1)
}

func main() {                                                                   
	
	//fibonacci
    // fmt.Print("Enter number: ")                                                 
    // fmt.Scanf("%d", &num)                                                       
    // fmt.Println("Fibonacci is ",fibonacci(num));
	fmt.Println("Fibonacci sequence:")
    for i := 0; i < 10; i++ {
        fmt.Print(Fibonacci(i), " ")
    }
    fmt.Println()
	
	//factorial
    // fmt.Print("Enter number: ")                                                 
    // fmt.Scanf("%d", &num)                                                       
    // fmt.Println("factorial is ",factorial(num))
}