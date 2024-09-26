// a map with a slice as a value
package main
import "fmt"

func main(){
  cities:=map[string][]string{
    "Paris": {"Eiffel Tower","Loure Museum"},
    "New York":{"Statue of Liberty","Central Park"},
    "Tokyo":{"Shibuya Crossing","Tokyo Tower"},
    "Jakarta":{"Monas Tower","Dufan"},
  }

  for key,value := range cities {
    for _,places := range value {
      fmt.Println(key, "->",places)
    }
  }
}
