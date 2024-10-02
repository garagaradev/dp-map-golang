//write a function that takes two maps of map[string]int, and merges them into one map.
//if a key exist in both maps, their values should be summed.
package main
import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int{
  result := make(map[string]int)

  for k, v := range m1 {
    result[k]  = v
  }

  for k, v := range m2 {
    result[k] += v
  }

  return result

}

func main(){
  m1:= map[string]int{"a":1, "b":2, "c":3}
  m2:= map[string]int{"b":3, "c":4, "d":5}

  merged := mergeMaps(m1, m2)
  fmt.Println("merged:", merged)
}
