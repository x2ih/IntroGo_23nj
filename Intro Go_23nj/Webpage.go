package main

// bron https://go.dev/doc/articles/wiki/
// Introduction
import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// Data Structures
type PaginaVar struct {
	Datum string
	Tijd  string
}

func main() {
	fs := http.FileServer(http.Dir("./assest"))
	http.Handle("/assest/", http.StripPrefix("/assest/", fs))
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/"))))
	http.HandleFunc("/", IndexHtml)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// functie voor Datum en Tijd
func IndexHtml(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now
	IndexPaVar := PaginaVar{ //store the date and time in a struct
		Datum: now.Format("02-01-2006"),
		Tijd:  now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("index.html") //parse the html file index.html
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, IndexPaVar) //execute the template and pass it the IndexPaVar- struct to fill in the gaps
	if err != nil {                // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
