package gopractice

import("fmt"
		"net/http"
)

func ExampleJson() {
	
	b := NewBook("harry potter", "1234-A")

	fmt.Printf("got book %s", b)

	//Output: got book X
}

func ExampleServer() {

    http.HandleFunc("/book/", handler("silent hill"))
    http.ListenAndServe(":8080", nil)

    //Output: none
}