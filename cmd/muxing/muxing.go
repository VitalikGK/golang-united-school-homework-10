package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", NameParam)
	router.HandleFunc("/bad", Bad)
	// router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "mess.html")
	// 	DataParam(w, r)
	// })
	router.HandleFunc("/data/{PARAM}", DataParam).Methods(http.MethodPost)
	router.HandleFunc("/", YourHandler)

	http.Handle("/", router)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Page")
}

func Bad(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Status: %d", http.StatusInternalServerError), http.StatusInternalServerError)
}

func NameParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["PARAM"]
	response := fmt.Sprintf("Hello, %s!", param)
	fmt.Fprint(w, response)
}

func DataParam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//mes := r.FormValue("PARAM")
		vars := mux.Vars(r)
		mes := vars["PARAM"]

		response := fmt.Sprintf("I got message:\n%s", mes)
		fmt.Fprint(w, response)
		fmt.Println("resp ", response)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("PORT", "8080")
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
