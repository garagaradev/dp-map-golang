// write a function that detects duplicates in a slice of strings using a map
// and returns a new slice containing only the duplicates
package main
import "fmt"

func findDuplicates(slice []string)[]string{
  occurences := make(map[string]int)
  var duplicates []string

  for _, value := range slice {
    occurences[value]++
  }

  for key, count := range occurences {
    if count > 1 {
      duplicates = append(duplicates, key)
    }
  }

  return duplicates
}

func main(){
  names := []string{"Jack","Alice","Bob","Alice","Jack","Eve"}
  fmt.Println("Initial names:", names)
  duplicates := findDuplicates(names)
  fmt.Println("Duplicated names:",duplicates)


}
