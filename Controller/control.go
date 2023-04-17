package Controller

import (
	"cityapplication/Database"
	"cityapplication/Model"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var Tin *template.Template
var Tout *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		population, err := strconv.Atoi(r.PostFormValue("population"))
		if err != nil {
			fmt.Println("Population must be integer value")
		}
		data := &Model.City_info_structs{
			City:       r.PostFormValue("city"),
			Population: population,
		}
		Database.Db.Create(data)
		Tout, _ := template.ParseFiles("./view/output.html")
		Tout.Execute(w, "successfully created")
	} else {
		Tin, _ := template.ParseFiles("./view/home.html")
		Tin.Execute(w, nil)
	}

}

func viewall(w http.ResponseWriter, r *http.Request) {
	var all []Model.City_info_structs
	Database.Db.Find(&all)
	fmt.Println(all)
	Tout, _ := template.ParseFiles("./view/output.html")
	Tout.Execute(w, all)

}

func Serving_func() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/viewall", viewall)
	return mux
}
