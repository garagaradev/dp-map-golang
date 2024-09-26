// create a map where the key is a struct representing a point with X and Y coordinates,
// and the value is a string representing the point's name
package main
import "fmt"

type Point struct {
  X, Y int
}

func main(){
  points := map[Point]string {
    {X:1, Y:2}:"A",
    {X:3, Y:4}:"B",
    {X:5, Y:6}:"C",
  }
  fmt.Println(points)
}

