package main

import (
	"city_population/Controller"
	"city_population/Database"
	"net/http"
)

func init() {
	Database.Getconnection()
}

func main() {
	//st := time.Now()
	//tmp := template.Must(template.ParseFiles("/views/home.html"))
	mux := Controller.Serving_func()
	http.ListenAndServe("localhost:8880", mux)

	//Database.Db.Create(&city_info_structs{City: "dudhani", Population: 400000})

	//fmt.Println(time.Since(st))
}
