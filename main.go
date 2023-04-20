package main

import (
	"cityapplication/Controller"
	"cityapplication/Database"
	"net/http"
)

func init() {
	Database.Getconnection()
	Database.Migratedatabase()
}
}

func main() {
	//st := time.Now()
	//tmp := template.Must(template.ParseFiles("/views/home.html"))
	mux := Controller.Serving_func()
	http.ListenAndServe("0.0.0.0:2000", mux)

	//Database.Db.Create(&city_info_structs{City: "dudhani", Population: 400000})

	//fmt.Println(time.Since(st))
}
