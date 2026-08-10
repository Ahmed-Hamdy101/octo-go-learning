package main
import ( 
	"fmt"
)
func main(){
	// initial value static type
	var str string = "Hello World!"
	// define variable which will work with `any`
	NumberTwo :=2
	// define integer
	var age  int = 30
	// define float
	var pi float32 = 3.14
	// define boolean
	var isTrue bool = true
	
	fmt.Println(NumberTwo,str,age,pi,isTrue)

}
