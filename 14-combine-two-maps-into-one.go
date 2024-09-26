// Given two maps, merge them into a single map.
// If the same key exists in both map, sum the values from both maps
package main
import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int{
  result := make(map[string]int)
  for k, v := range map1 {
    result[k] = v
  }
  for k,v := range map2 {
    result[k] += v
  }

  return result
}

func main(){
  map1 := map[string]int {"apple":5, "banana":2}
  fmt.Println("map1:",map1)
  map2 := map[string]int {"banana":3, "orange":8}
  fmt.Println("map2:",map2)
  mergedMap := mergeMaps(map1, map2)
  fmt.Println("Merged Map:",mergedMap)
}

