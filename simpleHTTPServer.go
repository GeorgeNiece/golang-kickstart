package main
import (
"fmt"
"net/http"
"strconv"
)
var counter int = 1
// define a type for the response
type Hello struct{ counter int}
// let that type implement the ServeHTTP method (defined in
// interface http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if(counter%2==0){
        fmt.Fprint(w, "Hello!" + strconv.Itoa(counter))
	   fmt.Println("Even")
	   counter++ 
	   return
    }else{
        fmt.Fprint(w, "Goodbye!" + strconv.Itoa(counter))
	   fmt.Println("Odd")
	   counter++ 
	   return
    }
}

func main() {
var h Hello
http.ListenAndServe("localhost:8686", h)
}
// Here's the method signature of http.ServeHTTP:
// type Handler interface {
// ServeHTTP(w http.ResponseWriter, r *http.Request)
// }