// write a function that takes a map of string to int and a function as a parameter.
// the function should apply the passed-in function to each value in the map
package main
import "fmt"

func applyFunc(m map[string]int, f func(int) int) {
  for key, value := range m {
    m[key] = f(value)
  }
}

func main(){
  numbers := map[string]int{"a":1, "b":2, "c":3}
  fmt.Println("Initial map:",numbers)
  double := func(n int) int {return 2 * n}
  applyFunc(numbers, double)
  fmt.Println("Current map:", numbers)

}
