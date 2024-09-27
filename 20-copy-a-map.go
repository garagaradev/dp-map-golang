// write a function that copies the contents of one map to another map.
// the map has string keys and integer values.
package main
import "fmt"

func copyMap(original map[string]int) map[string]int {
  copy := make(map[string]int)
  for key, value := range original {
    copy[key] = value
  }
  return copy
}

func main(){
  scores := map[string]int{"Jack":85,"Alice":90, "Bob":75}
  fmt.Println("Score Info:", scores)
  copiedScores := copyMap(scores)
  fmt.Println("Copied Scores:", copiedScores)
}
