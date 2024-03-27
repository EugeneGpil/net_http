package pathParam

import "net/http"

func Run() {
	router := createRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /item/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte(id + "\n"))
	})

	router.HandleFunc("/item/{id}/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("default\n"))
	})

	router.HandleFunc("test.com/item/{id}/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test.com\n"))
	})

	router.HandleFunc("POST test.com/item/{id}/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST test.com\n"))
	})

	// router.HandleFunc("/item/test/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("TEST\n"))
	// })

	// router.HandleFunc("/{entity}/test/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("entity test\n"))
	// })

	return router
}
