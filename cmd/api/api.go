package main

import (
	"fmt"
	"net/http"
	/*
	"github.com/go-chi/chi"
	"github.com/ahmedaliv/goapi/internal/handlers"
 	log	"github.com/sirupsen/logrus"
	*/
)


func main() {
	/*
	log.SetReportCaller(true);

	var r *chi.Mux = chi.NewRouter();

	handlers.Handler(r)
	
	err := http.ListenAndServe(":8080", r);
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			}).Error("Failed to start the server");
		}
		*/
		/*
		mx := http.NewServeMux()
		mx.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
				return
			}
			if r.Method == "GET" {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed) {
				return
			} 
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("base route"))
			//fmt.Fprint(w, "Hello, world!")
		})
		fmt.Println("starting go api service ... ");
		http.ListenAndServe(":8080", mx)
		*/

		// a, err := Sum(5,6)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(a)
}


// func Sum(a int, b int) (int, error) {
// 	if a < 0 || b < 0 {
// 		return 0, error.New("a and b must be positive numbers")
// 	}
// 	return a + b, nil
// }