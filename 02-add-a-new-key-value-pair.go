package main
import "fmt"

func main(){
  fruits := map[string]int{
    "apple":5,
    "banana":2,
    "orange":3,
  }
  fmt.Println("Initial fruits:",fruits)
  fruits["grape"] = 4
  fmt.Println("Current fruits:",fruits)
}
