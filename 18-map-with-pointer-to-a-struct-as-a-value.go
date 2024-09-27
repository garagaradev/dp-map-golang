// create a map where the key is a string representing a book title
// and the value is a pointer to a struct representing the book details (author and number of pages)
package main
import "fmt"

type Book struct {
  Author string
  NumPages int
}

func main(){
  books := map[string]*Book{
    "The Go Programming Language": {Author: "Alan Donovan",NumPages: 400},
    "Gun, Germ, and Steel": {Author: "Jared Diamond", NumPages:528},
  }
  fmt.Println(books)
  fmt.Println("Author of The Go Programming Language:",books["The Go Programming Language"].Author)
  fmt.Println("Author of Gun, Germ, and Steel:", books["Gun, Germ, and Steel"].Author)
}
