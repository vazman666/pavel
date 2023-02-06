package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"www/models"
	"www/pkg"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func helloMuxGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<head>	<meta http-equiv=\"refresh\" content=\"0; URL=http://www.japautozap.ru/\" /> </head> "))
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		fmt.Printf(" Логгин %s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
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
	//r.HandleFunc("/pavel/edit/{Id:[0-9]+}", pkg.EditPage).Methods("GET")
	r.HandleFunc("/pavel/edit/{Id:[0-9]+}", pkg.EditPage).Methods("GET")
	r.HandleFunc("/pavel/edit/{Id:[0-9]+}", pkg.EditHandler).Methods("POST")
	r.HandleFunc("/pavel/delete/{Id:[0-9]+}", pkg.DeleteHandler)
	r.HandleFunc("/pavel/create", pkg.CreateHandler)
	r.HandleFunc("/pavel/Checkout", pkg.Checkout)
	r.HandleFunc("/pavel/del", pkg.Del)
	r.Use(Logging)
	log.Fatal(http.ListenAndServe(":10011", r))
}
