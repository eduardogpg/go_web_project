package utils

import(
	"html/template"
	"net/http"
	"fmt"
	"bytes"
	"log"
)

var initLayoutFuncs = template.FuncMap{
	"yield": func() (string, error){
		return "", fmt.Errorf("Yield called innappropriatley")
	},
}

var layout = template.Must( template.New("layout.html").Funcs(initLayoutFuncs).ParseFiles("views/templates/layout.html"), )
var templates = template.Must(template.New("t").ParseGlob("./views/templates/**/*.html"))

func TestRenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}){
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error){
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()) , err
		},
	}

	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)

	if err != nil{
		log.Println(err)
		http.Error( w, "Page not found!", http.StatusInternalServerError, )
	}
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}){
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error){
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()) , err
		},
	}
	
	var NewLayout = template.Must( template.New("layout.html").Funcs(funcs).ParseFiles("views/templates/layout.html"), )
	err := NewLayout.Execute(w, data)
	if err != nil{
		log.Println(err)
		http.Error( w, "Algo paso!", http.StatusInternalServerError, )	
	}
}





