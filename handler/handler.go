package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("welcome to home"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "belajar web golang",
	// 	"content": "Belajar Web Golang bersama Lukman Agung Prakoso",
	// }

	//data := entity.Product{ID: 1, Name: "Mobilio", Price: 2000000, Stock: 4}

	data := []entity.Product{
		{ID: 1, Name: "Agiya", Price: 1000000, Stock: 12},
		{ID: 1, Name: "Avansa", Price: 2000000, Stock: 8},
		{ID: 1, Name: "Pajero", Price: 5000000, Stock: 3},
		{ID: 1, Name: "Expander", Price: 4000000, Stock: 15},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello word, saya sedang belajar golang web"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Product page ke : %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
			return
		}
	}
	http.Error(w, "Error is Happening, keep calm", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
		}

		name := r.Form.Get("name")
		massage := r.Form.Get(("massage"))

		data := map[string]interface{}{
			"name":    name,
			"massage": massage,
		}
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapening, keep calm", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Error is Happening, keep calm", http.StatusBadRequest)
	}

}
