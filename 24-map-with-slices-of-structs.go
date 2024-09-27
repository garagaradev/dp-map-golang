// create a map where the key is a string representing a department and the value is a slice of structs.
// representing employees with their name and age.
package main
import "fmt"

type Employee struct {
  Name string
  Age int
}

func main(){
  departments := map[string][]Employee{
    "Engineering": {
      {Name: "Alice", Age: 30},
      {Name: "Bob", Age:25},
    },
    "HR": {
      {Name:"Carol",Age:35},
    },
  }
  fmt.Println(departments)
}
