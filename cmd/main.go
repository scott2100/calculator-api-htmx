package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

type Results struct {
	Results int
}

type Operands struct {
	Operand1 int
	Operand2 int
}

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("../views/*.html"))
}

func main() {

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/subtract", handleSubtract)
	http.HandleFunc("/multiply", handleMultiply)
	http.HandleFunc("/divide", handleDivide)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 + operands.Operand2

	results := Results{
		Results: result,
	}

	templ = template.Must(template.ParseGlob("../views/results.html"))
	templ.ExecuteTemplate(w, "results.html", results)
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 - operands.Operand2

	results := Results{
		Results: result,
	}

	templ = template.Must(template.ParseGlob("../views/results.html"))
	templ.ExecuteTemplate(w, "results.html", results)
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 * operands.Operand2

	results := Results{
		Results: result,
	}

	templ = template.Must(template.ParseGlob("../views/results.html"))
	templ.ExecuteTemplate(w, "results.html", results)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 / operands.Operand2

	results := Results{
		Results: result,
	}

	templ = template.Must(template.ParseGlob("../views/results.html"))
	templ.ExecuteTemplate(w, "results.html", results)
}

func parseOperands(r *http.Request) Operands {
	a := r.FormValue("a")
	b := r.FormValue("b")

	num1, err := strconv.Atoi(a)
	if err != nil {
		log.Println(err)
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		log.Println(err)
	}

	return Operands{num1, num2}
}
