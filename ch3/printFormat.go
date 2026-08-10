package main
import ( 
	"fmt"
)
func main(){
// decalare variables
// string
str:= "mirror in x.0"
ing := 33
isTrue := false
// put a value with formating
fmt.Printf("this is string: %v\n",str)
// get datatypes
fmt.Printf("bunch of datatypes : %T,%T,%T",str,ing,isTrue)
// converte to float from printf
fmt.Printf("\nconverte to float %.2f\n",float64(ing))

}
