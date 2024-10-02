//create a map where the value is another map. For example, store information about 
//cities within a country. The outer map has countries as keys, 
//and the inner map has cities and their populations.
package main
import "fmt"

func main(){
  countries := map[string]map[string]int {
    "USA":{
      "New York":8500000,
      "Los Angeles":4000000,
    },
    "Canada":{
      "Toronto": 3000000,
      "Vancouver": 675000,
    },
  }
  countries["Canada"]["Montreal"] = 1780000
  fmt.Println(countries)
}
