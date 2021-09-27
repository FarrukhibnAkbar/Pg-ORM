package main

import(
	// "fmt"
	"net/http"
	"app/modules/user"
	"app/modules/product"
	"github.com/gorilla/mux"
)

func main() {
	
	r := mux.NewRouter()

	r.HandleFunc("/{Id}", user.GET_ID)
	r.HandleFunc("/products", product.GetMany)

	// fmt.Println(config.DB_HOST)
	// fmt.Println(config.SMTP_HOST)

	http.ListenAndServe(":8080", r)
} 