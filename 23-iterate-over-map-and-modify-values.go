//write a function that iterates over a map of string to int and multiplies each value by 2.
package main
import "fmt"

func multiplyValues(m map[string]int){
  for key, value := range m {
    m[key] = 2 * value
  }
}

func main(){
  ages := map[string]int{"Jack":28, "Alice":34, "Bob":23}
  fmt.Println("Initial data:", ages)
  multiplyValues(ages)
  fmt.Println("Data after multiplied:", ages)
}
