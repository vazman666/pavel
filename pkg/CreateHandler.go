package pkg

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"text/template"
	"time"
	"www/models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		Partnum := r.FormValue("Partnum")
		Qty, _ := strconv.Atoi(r.FormValue("Qty"))
		Price, _ := strconv.ParseFloat(r.FormValue("Price"), 32)
		fmt.Println(reflect.TypeOf(Price))
		Provider := r.FormValue("Provider")
		Name := r.FormValue("Name")
		Remark := r.FormValue("Remark")
		Date := time.Now().Format("02 January 2006")
		Color := "#ffffff"
		_, err = models.Database.Exec("insert into mybase.mybase (Partnum, Qty, Price,Provider,Name,Remark,Date,Color) values (?, ?, ?, ?, ?, ?, ?, ?)",
			Partnum, Qty, Price, Provider, Name, Remark, Date, Color)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/pavel", 301)
	} else {
		fmt.Printf("get")
		files := []string{
			"./ui/html/create.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",
		}

		// Используем функцию template.ParseFiles() для чтения файлов шаблона.
		// Если возникла ошибка, мы запишем детальное сообщение ошибки и
		// используя функцию http.Error() мы отправим пользователю
		// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// Затем мы используем метод Execute() для записи содержимого
		// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
		// возможность отправки динамических данных в шаблон.
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
