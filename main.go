package main

import (
	"net/http"
	//"fmt"
	"recipeKeeper/data"
	"recipeKeeper/handlers"
)

func main() {
	data.FetchAllRecipes()
	//fmt.Println (data.AllRecipes)
	http.HandleFunc("/", handlers.HomePage) // if i have this request (/) do this - print response from homepage func
	http.ListenAndServe(":8000", nil) // to create a default server

}

//http for computer to comunicate -> to request data from the server (web server)
// url endpoint - page -> link that to a page/function (url path) - HANDLER (several pages several handlers), a way of organising the data
// ctrl + c to cancel -> to make changes 
// restful api - 