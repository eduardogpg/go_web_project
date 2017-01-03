package utils

import(
	"html/template"
	"net/http"
	"fmt"
	"bytes"
)

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error){
		return "", fmt.Errorf("Yield called innappropriatley")
	},
}

var layout = template.Must( template.New("layout.html").Funcs(layoutFuncs).ParseFiles("views/templates/layout.html"), )
var templates = template.Must(template.New("t").ParseGlob("./views/templates/**/*.html"))

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}){
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
		http.Error( w, "Page not found!", http.StatusInternalServerError, )
	}
}


