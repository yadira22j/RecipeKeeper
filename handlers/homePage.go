package handlers
 //
import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) { // standard template for a handler , response and request 2 way communication website to server
	fmt.Fprintf(w, "hey")
}
