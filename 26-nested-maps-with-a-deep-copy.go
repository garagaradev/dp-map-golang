//create a map where the value is another map.
//write a function that performs a deep copy of such a nested map.
package main
import "fmt"

func deepCopyMap(original map[string]map[string]int) map[string]map[string]int {
  copy := make(map[string]map[string]int)
  for k, v := range original {
    innerMap := make(map[string]int)
    for innerK, innerV := range v {
        innerMap[innerK] = innerV
    }
    copy[k] = innerMap
  }
  return copy
}

func main(){
  nestedMap := map[string]map[string]int {
    "group1": {"Jack":85, "Alice":90},
    "group2": {"Bob":75, "Eve": 95},
  }
  fmt.Println("Nested map:", nestedMap)
  copiedMap := deepCopyMap(nestedMap)
  fmt.Println("Copied map:", copiedMap)
}
