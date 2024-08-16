package web

import (
	"cyoa/helpers"
	"cyoa/types"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartServer(gameConfig types.CYOAGameConfig, defaultStory string) {
	router := mux.NewRouter()

	router.HandleFunc("/{chapter}", chapterHandler(gameConfig))
	router.HandleFunc("/", defaultHandler(gameConfig[defaultStory]))

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}

func chapterHandler(c types.CYOAGameConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := os.ReadFile("./index.html")
		if err != nil {
			w.Write([]byte("An error has occured"))
			log.Fatal(err)
		}

		t, err := template.New("webpage").Parse(string(html))
		helpers.CheckError(err)

		chapter := r.URL.Path[1:]
		data := c[chapter]

		err = t.Execute(w, data)
		helpers.CheckError(err)
	}
}

func defaultHandler(introStory types.Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := os.ReadFile("./index.html")
		if err != nil {
			w.Write([]byte("An error has occured"))
			log.Fatal(err)
		}

		t, err := template.New("webpage").Parse(string(html))
		helpers.CheckError(err)

		data := introStory

		err = t.Execute(w, data)
		helpers.CheckError(err)
	}
}
