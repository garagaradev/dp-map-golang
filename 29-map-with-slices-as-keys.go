//write a function that creates a map where the keys are slices of strings
//and the value are integers. How would you manage comparison of slice keys
package main
import "fmt"

func addEntry(m map[string]int, key []string, value int){
  keyStr := fmt.Sprint(key)
  m[keyStr] = value
}

func main(){
  m := make(map[string]int)
  key1 := []string{"apple","banana"}
  key2 := []string{"orange", "grape"}
  fmt.Println("key1:", key1)
  fmt.Println("key2:", key2)
  addEntry(m, key1, 5)
  addEntry(m, key2, 10)

  fmt.Println(m)
}
