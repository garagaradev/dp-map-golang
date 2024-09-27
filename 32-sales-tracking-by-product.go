// write a function that tracks sales of products. Each product has a name and a total number of units sold.
// udpate the sales when a product is sold and print the total sales for each product.
package main
import "fmt"

func recordSale(sales map[string]int, product string, quantity int){
  sales[product] += quantity
}

func main(){
  sales := map[string]int{
    "Laptop": 50,
    "Phone": 30,
  }
  fmt.Println("Initial sales:", sales)
  recordSale(sales,"Laptop", 5)
  recordSale(sales, "Phone", 10)

  fmt.Println("Current sales:", sales)

}
