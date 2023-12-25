package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test/mathoperations"
)

const (
	sum       = "sum"
	div       = "div"
	operation = "operation"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query(), "query")
	for k, v := range r.URL.Query() {
		fmt.Println(k, v)
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handleDiv(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is division!")
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is sum!")
}

func handlerMath(w http.ResponseWriter, r *http.Request) {
	opr := r.URL.Query().Get(operation)

	input1 := r.URL.Query().Get("input1")
	input1F, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "expected a number, got: %v", input1)
		return
	}

	input2 := r.URL.Query().Get("input2")
	input2F, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "expected a number, got: %v", input2)
		return
	}

	switch opr {
	case sum:
		fmt.Fprintf(w, "result for %s of %s and %s = %v", opr, input1, input2, mathoperations.Sum(input1F, input2F))
		return
	case div:
		fmt.Fprintf(w, "result for %s of %s and %s = %v", opr, input1, input2, mathoperations.Div(input1F, input2F))
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unkown opr: %s\n", opr)
		return
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/div", handleDiv)
	http.HandleFunc("/math", handlerMath)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
