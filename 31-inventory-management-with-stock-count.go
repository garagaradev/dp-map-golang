//write a function that manages the inventory of products. Each product has a name (string)
//and a stock count (int). Implement a system to add new products, update stock, and print
//the current inventory
package main
import "fmt"

func addStock(inventory map[string]int, product string, count int){
  inventory[product] += count
}

func main(){
  inventory := map[string]int{
    "Laptop":10,
    "Phone":5,
  }
  fmt.Println("Initial inventory:", inventory)

  addStock(inventory, "Tablet", 7)
  addStock(inventory, "Phone", 3)

  fmt.Println("Current inventory:", inventory)


}
