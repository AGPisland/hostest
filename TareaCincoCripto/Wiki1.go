package gowiki

import (
	"io/ioutil"
	"strings"
)

func FusionCssHtml(html string, css string) string {
	a := strings.Replace(html, "STYLECSS", css, 10)
	//fmt.Println(a)
	return a
}

type Page struct {
	Title string
	Body  []byte
}

//Save Guarda un texto con el body d eP
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
LoadPage esto lee un archivo en el directorio y retorna una interfaz
con el archivo en el body.
*/
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		body = nil
	}
	return &Page{Title: title, Body: body}, nil //crea la pagina pero no la guarda
}

/*Prueba de funciones*/
func Main2() {
	p1 := &Page{Title: "TestPage", Body: []byte(Pi)}
	p1.Save()
	_, _ = LoadPage("TestPage")
	//fmt.Println(string(p2.Body))
}
