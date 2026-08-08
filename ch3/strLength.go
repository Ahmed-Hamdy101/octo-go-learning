package main
import ( 
	"fmt"
)
func main(){
// decalare variables
str1 :="string one "
str2 := "string two"
str3 :="string three"
// the length of output after printed
stringLength , err := fmt.Println(str1 , str2, str3)
/*
if i have a varaible and i don't use it 
stringLength , _ := fmt.Println(str1 , str2, str3)
*/
// if there is no error
if err == nil {
// print length
fmt.Println(stringLength)
}
}
