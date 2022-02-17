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

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()

	rows, err := models.Database.Query("select * from mybase.mybase")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []models.Product{}
	var sum float32 = 0

	for rows.Next() {

		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Partnum, &p.Qty, &p.Price, &p.Provider, &p.Name, &p.Remark, &p.Date, &p.Color)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p.Summ = float32(p.Qty) * p.Price
		products = append(products, p)
		sum += float32(p.Qty) * p.Price

	}

	files := []string{
		"./ui/html/home.page.tmpl",
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
	err = ts.Execute(w, products)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	fmt.Fprintf(w, "Сумма итого: %v\n", sum)

}
func Count(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()

	vars := mux.Vars(r)
	Id := vars["Id"]
	fmt.Println("Id=", Id)
	row := models.Database.QueryRow("select * from mybase.mybase where Id = ?", Id)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Partnum, &prod.Qty, &prod.Price, &prod.Provider, &prod.Name, &prod.Remark, &prod.Date, &prod.Color)
	if err != nil {
		fmt.Println(err)

	}
	if prod.Color == "#ffffff" {

		prod.Color = "#fa8e47"
	} else {

		prod.Color = "#ffffff"
	}

	_, err = models.Database.Exec("upDate mybase.mybase set Color=? where Id=?", prod.Color, Id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/pavel", 301)
}
func EditPage(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	vars := mux.Vars(r)
	Id := vars["Id"]

	row := models.Database.QueryRow("select * from mybase.mybase where Id = ?", Id)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Partnum, &prod.Qty, &prod.Price, &prod.Provider, &prod.Name, &prod.Remark, &prod.Date, &prod.Color)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		files := []string{
			"./ui/html/edit.tmpl",
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
		err = ts.Execute(w, prod)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	Id := r.FormValue("Id")
	Partnum := r.FormValue("Partnum")
	Qty, _ := strconv.Atoi(r.FormValue("Qty"))
	Price, _ := strconv.ParseFloat(r.FormValue("Price"), 32)
	Provider := r.FormValue("Provider")
	Name := r.FormValue("Name")
	Remark := r.FormValue("Remark")
	Date := time.Now().Format("02 January 2006")
	Color := "#ffffff"

	_, err = models.Database.Exec("upDate mybase.mybase set Partnum=?, Qty=?, Price=?,Provider=?,Name=?,Remark=?,Date=? , Color=? where Id=?",
		Partnum, Qty, Price, Provider, Name, Remark, Date, Color, Id)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/pavel", 301)
}
func Del(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	_, err := models.Database.Query("DELETE FROM mybase.mybase WHERE color='#fa8e47';")

	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/pavel", 301)

}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	vars := mux.Vars(r)
	Id := vars["Id"]

	_, err := models.Database.Exec("delete from mybase.mybase where Id = ?", Id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/pavel", 301)
}
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
func Checkout(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	rows, err := models.Database.Query("select * from mybase.mybase")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []models.Product{}

	for rows.Next() {

		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Partnum, &p.Qty, &p.Price, &p.Provider, &p.Name, &p.Remark, &p.Date, &p.Color)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p.Color == "#fa8e47" {
			products = append(products, p)
			fmt.Println(products)
		}
	}
	sum := float32(0.0)
	for _, j := range products {
		fmt.Println(j.Price, j.Qty)
		sum = sum + j.Price*(float32)(j.Qty)
	}
	tmpl, _ := template.ParseFiles("templates/show.html")
	tmpl.Execute(w, products)
	fmt.Fprintf(w, "\nОбщая сумма: %6.2f\n", sum)
	http.Redirect(w, r, "/pavel", 301)
}
