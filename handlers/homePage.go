package handlers
 //
import (
	//"fmt"
	"net/http"
	"html/template"
	"log"
	"recipeKeeper/data"
)

func HomePage(w http.ResponseWriter, r *http.Request) { // standard template for a handler , response and request 2 way communication website to server
	//fmt.Fprintf(w, "hey")
	tmpl, err := template.ParseFiles("frontend/homePage.html")
	if err != nil {
		log.Fatal(err, "ERROR: problem with home file path")
	}
	tmpl.Execute(w, data.AllRecipes) // execute this template on to response write with the data from all recipes, put the data into the template and post to homepage
}
