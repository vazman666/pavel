package pkg

import (
	"encoding/json"
	"net/http"
	"time"
	"www/models"

	_ "github.com/go-sql-driver/mysql"
)

func ApiAdd(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("ApiAdd\n")
	status := "success"
	message := "строка добавлена"
	var item models.Product
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.Date = time.Now().Format("02 January 2006")
	item.Color = "#ffffff"
	//log.Printf("Received JSON: %+v\n", item)
	if len(item.Partnum) == 0 || item.Qty == 0 || item.Price == 0 || len(item.Provider) == 0 || len(item.Name) == 0 || len(item.Provider) == 0 || len(item.Remark) == 0 {
		status = "fail"
		message = "Не хватило параметра"
	} else {
		_, err = models.Database.Exec("insert into mybase.mybase (Partnum, Qty, Price,Provider,Name,Remark,Date,Color) values (?, ?, ?, ?, ?, ?, ?, ?)",
			item.Partnum, item.Qty, item.Price, item.Provider, item.Name, item.Remark, item.Date, item.Color)

		if err != nil {
			status = "fail"
			message = "ошибка доступа к базе"
		}
	}
	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}
