//write a function that takes  a map of string to int and returns all keys
//that have the maximum value.
package main
import "fmt"

func findMaxKeys(m map[string]int) []string {
  var max int
  var keys []string
  for _, value := range m {
    if value > max {
      max = value
    }
  }

  for key, value := range m {
    if value == max {
      keys = append(keys, key)
    }
  }
  return keys
}

func main(){
  scores := map[string]int{"Jack":85, "Alice":90, "Bob":90, "Eve":75, "Carol":60, "Matt":70}
  fmt.Println("Initial scores:", scores)
  maxScores := findMaxKeys(scores)
  fmt.Println("Max scores:", maxScores)
}
