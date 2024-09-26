//iterate over all kv pairs
package main
import "fmt"

func main(){
  fruits:=map[string]int{
    "apple":5,
    "banana":8,
    "orange":13,
    "grapes":7,
  }

  for k,v := range fruits {
    fmt.Println(k,"->",v)
  }
}
