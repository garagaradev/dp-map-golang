// create  a map where each key is a string (representing a country)
// and the value is another map with city names as keys
// and their population as values
package main
import "fmt"

func main(){
  countryCities := map[string]map[string]int{
    "USA":{
      "New York": 8175133,
      "Los Angeles": 3792621,
    },
    "France":{
      "Paris": 2140526,
      "Marseille": 861635,
    },
  }
  for country, city := range countryCities {
    fmt.Println("Country:",country)
    for city, population := range city {
      fmt.Println("    City:",city,", population:",population)
    }
  }
}
