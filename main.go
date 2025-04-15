package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Results struct {
	Results int
}

type Operands struct {
	Operand1 int
	Operand2 int
}

var templ *template.Template

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)
	// Full path that identifies the resource:
	full := filepath.Join(dir, "views")
	full = full + "/*.html"
	templ = template.Must(template.ParseGlob(full))
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
	err := templ.ExecuteTemplate(w, "index", nil)
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
	err := templ.ExecuteTemplate(w, "result", results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 - operands.Operand2

	results := Results{
		Results: result,
	}
	err := templ.ExecuteTemplate(w, "result", results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 * operands.Operand2

	results := Results{
		Results: result,
	}
	err := templ.ExecuteTemplate(w, "result", results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDivide(w http.ResponseWriter, r *http.Request) {

	operands := parseOperands(r)

	result := operands.Operand1 / operands.Operand2

	results := Results{
		Results: result,
	}
	err := templ.ExecuteTemplate(w, "result", results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
