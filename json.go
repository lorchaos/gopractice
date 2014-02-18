package gopractice


import ("encoding/json"
		"fmt"
		"net/http"
)

type Book struct {
	
	Title string
	ISBN string
}


func NewBook(title, isbn string) string {

	b, _ := json.Marshal(Book{title, isbn})

	return string(b)
}

func handler(title string) (func(w http.ResponseWriter, r *http.Request)) {
    
	return func(w http.ResponseWriter, r *http.Request) {

		isbn := r.URL.Path[1:]

		b := NewBook(title, isbn)

    	fmt.Fprintf(w, b)
    }
}





