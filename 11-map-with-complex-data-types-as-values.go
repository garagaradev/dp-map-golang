// Create a map where the key is a string (representing a country)
// and the value is a struct that holds information like population (int) and capital city (string).
package main
import "fmt"

type CountryInfo struct {
  Population int
  Capital    string
}

func main(){
  countries := map[string] CountryInfo {
    "USA": {Population: 331002651, Capital: "Washington, D.C."},
    "France": {Population:65273511, Capital: "Paris"},
    "Japan": {Population:126476461, Capital: "Tokyo"},
  }
  fmt.Println(countries)
}
