// write a function that returns the intersection of two maps.
// i.e. the keys that exists in both maps along with their values.
package main
import "fmt"

func intersectMaps(map1, map2 map[string]int) map[string]int {
  intersection := make(map[string]int)
  for k, v := range map1 {
    if v2, ok := map2[k]; ok && v == v2 {
      intersection[k] = v
    }
  }
  return intersection
}

func main(){
  map1 := map[string]int {"apple":5, "banana":2, "orange":8}
  fmt.Println("Map1:", map1)
  map2 := map[string]int {"banana":2, "orange":10, "grape":3}
  fmt.Println("Map2:", map2)
  fmt.Println(intersectMaps(map1, map2))
}
