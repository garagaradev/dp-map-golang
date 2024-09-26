// create a map of string to int, where the string represents a product
// and the int represents its stock quantity.
// Print the total number of products in the map
package main
import "fmt"

func main(){
  products := map[string]int{
    "Laptop":5,
    "Mouse":10,
    "Keyboard":7,
  }
  fmt.Println("Products info:", products)
  fmt.Println("Number of products:", len(products))
}
