package template

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

var layout = template.Must(
	template.New("layout.html").Funcs(layoutFuncs).ParseFiles("templates/layout.html"),
)

var templates = template.Must(template.New("t").ParseGlob("./templates/**/*.html"))

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


func LastRenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}){
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error( w, "Page not found!", http.StatusInternalServerError, )
	}
}
