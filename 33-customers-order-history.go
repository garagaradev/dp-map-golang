package main
import "fmt"

func addOrder(orders map[string][]string, customer string, order string){
  orders[customer] = append(orders[customer], order)
}

func main(){
  orders := map[string][]string{
    "Alice": {"Laptop","Phone"},
    "Bob":{"Tablet"},
  }
  fmt.Println("Initial orders:", orders)

  addOrder(orders, "Alice", "Headphones")
  addOrder(orders, "Bob", "Smartwatch")

  fmt.Println("Customer order history:", orders)
}
