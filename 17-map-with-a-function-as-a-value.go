// create a map where the key is a string representing a mathematical operation
// ("add","substract","multiply","divide") and the value is a function that performs that operation on two integers.
package main
import "fmt"

func main(){
  operations := map[string] func(int, int) int {
    "add": func(a, b int) int {return a + b},
    "substract": func(a,b int) int {return a - b},
    "multiply": func(a, b int) int {return a * b},
    "divide": func(a, b int) int {return a / b},
  }

  fmt.Println("10 + 5 = ", operations["add"](10,5))
  fmt.Println("10 - 5 = ", operations["substract"](10,5))
  fmt.Println("10 * 5 = ", operations["multiply"](10,5))
  fmt.Println("10 / 5 = ", operations["divide"](10,5))
}
