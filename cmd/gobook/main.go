package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

/*
docker build -t go-app .
docker run -v $(pwd):/app -w /app -p 8080:8080 -it go-app sh
*/

//função de entrada do programa
func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if(err != nil) {
		panic(err)
	}

	//espera daqui pra baixo tudo ser executado, qd terminar de executar, aí dá o close
	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandlers := web.NewBookHandler(bookService)

	//NewServerMux é o servidor web
	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
} 
