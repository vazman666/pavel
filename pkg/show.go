package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"www/models"

	_ "github.com/go-sql-driver/mysql"
)

func Show(w http.ResponseWriter, r *http.Request) {
	rows, err := models.Database.Query("select * from mybase.mybase")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	Products := []models.Product{}
	var sum float32 = 0

	for rows.Next() {

		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Partnum, &p.Qty, &p.Price, &p.Provider, &p.Name, &p.Remark, &p.Date, &p.Color)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p.Summ = float32(p.Qty) * p.Price
		Products = append(Products, p)
		sum += float32(p.Qty) * p.Price

	}
	//fmt.Printf("%v\n", models.Products)
	//w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	jsonResp, err := json.Marshal(Products)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

