
package main
import "fmt"


func substr(input string, start int, length int) string {
    asRunes := []rune(input)
    
    if start >= len(asRunes) {
        return ""
    }
    
    if start+length > len(asRunes) {
        length = len(asRunes) - start
    }
    
    return string(asRunes[start : start+length])
}

func main() {
    
  strng := "Gaurav"
  var subStr string
  
  for i:=0;i< len(strng);i++{
      
  
  subStr = substr(strng,i,6)
  fmt.Println("Generated Substring is",subStr)
  }
  
}
