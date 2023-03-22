package main

import (
	"database/sql"
	"log"
	"net/http"

	"www/models"
	"www/pkg"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func helloMuxGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<head>	<meta http-equiv=\"refresh\" content=\"0; URL=http://www.japautozap.ru/\" /> </head> "))
}

/*
	func Pavel(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Func Pavel\n"))
	}
*/
func main() {
	db, err := sql.Open("mysql", "vazman:rbhgbxb1@/mybase")

	if err != nil {
		log.Println(err)
	}
	models.Database = db
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/", helloMuxGoHandler)
	r.HandleFunc("/pavel", pkg.IndexHandler)
	r.HandleFunc("/pavel/add/{Id:[0-9]+}", pkg.Count)
	r.HandleFunc("/pavel/edit/{Id:[0-9]+}", pkg.EditPage).Methods("GET")
	r.HandleFunc("/pavel/edit/{Id:[0-9]+}", pkg.EditHandler).Methods("POST")
	r.HandleFunc("/pavel/delete/{Id:[0-9]+}", pkg.DeleteHandler)
	r.HandleFunc("/pavel/create", pkg.CreateHandler)
	r.HandleFunc("/pavel/Checkout", pkg.Checkout)
	r.HandleFunc("/pavel/del", pkg.Del)
	r.HandleFunc("/pavel/show", pkg.Show)
	r.HandleFunc("/pavel/api/add", pkg.ApiAdd)
	log.Fatal(http.ListenAndServe(":10011", r))
}
