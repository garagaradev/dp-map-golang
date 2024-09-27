// write a function that converts a JSON string into a map[string]interface{}.
// Assume JSON contains arbitrary nested structures.
package main
import (
  "encoding/json"
  "fmt"
)

func jsonToMap(jsonStr string) (map[string]interface{}, error){
  var result map[string]interface{}
  err := json.Unmarshal([]byte(jsonStr), &result)
  return result, err
}

func main(){
  jsonStr := `{"name":"Jack", "age":30, "skills":{"Go":true, "Python":false}}`
  fmt.Println(jsonStr)
  m, err := jsonToMap(jsonStr)
  if err != nil {
    fmt.Println("Error:",err)
  }
  fmt.Println(m)
}
