package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Hello Template Data Struct",
		Name:  "Ardi",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	})
}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Ardi",
		"Address": map[string]interface{}{
			"Street": "Jl Raya",
			"City":   "Jakarta",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/layout.gohtml", "./templates/footer.gohtml"))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Ardi",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
