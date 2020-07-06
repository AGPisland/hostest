package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"bitbucket.com/local/gowiki"
	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := gowiki.LoadPage(title) //->  si no puede abrir un titulo no puede entrar a la pagina y sera redirigido
	if p == nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, string(p.Body), p, "view")
}

//editHandler : solo para editar una pagina que no tiene direccion. esto es se guardara en un dato
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	//p, err := gowiki.LoadPage(title)
	//if err != nil {
	//	p = &gowiki.Page{Title: title}
	//}
	p := &gowiki.Page{Title: title}
	renderTemplate(w, gowiki.Edit, p, "edit")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body") //?post?
	p := &gowiki.Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	//p := &gowiki.Page{Title: "main"}
	//aca debo fucionar el main con el json de estadosl
	mystringResponseParticleApi := gowiki.Getonlinedevices() // recupero el jsn
	formatjson, _ := json.MarshalIndent(mystringResponseParticleApi, "", "	")
	//fmt.Printf("%+v\n", string(formatjson))
	//gowiki.FusionJsonHtmlEstatus(v) //
	p := gowiki.Getwebmain() //retorna la estructura web y podemos ver que debe imprimir
	//p.Side = mystringResponseParticleApi
	p.Side, _ = strconv.Unquote(string(formatjson))
	//fmt.Println(p.Side)
	//gowiki.Main es el template guardado en go
	//p es un interfaz. y

	mystring1 := gowiki.FusionCssHtml(gowiki.Main, gowiki.StyleCss)

	renderTemplate(w, mystring1, p, "main")
}

func othrHandler(w http.ResponseWriter, r *http.Request) {
	p := &gowiki.Page{Title: "othr"}
	//gowiki.Main es el template guardado en go
	renderTemplate(w, gowiki.FusionCssHtml(gowiki.Main, gowiki.StyleCss), p, "othr")
}

func renderTemplate(w http.ResponseWriter, stringHtmlMemoriaLocalVariable string, InterfacesPa interface{}, Name string) {
	t, err := gowiki.ParseFiles(stringHtmlMemoriaLocalVariable, Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, InterfacesPa) //esto genera el render de las interfacez son el templete y lo
	//incorpora en el response writer.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	gowiki.Recuperadordeviceindb()

	apoyo := gowiki.GetEstrucApoyo()
	gowiki.InitApi(apoyo.Token, apoyo.Deviceparticle)

	//gowiki.Main2()

	//domains := []string{"example.com", "www.example.com"}
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/view/", viewHandler)
	r.HandleFunc("/edit/", editHandler)
	r.HandleFunc("/save/", saveHandler)
	r.HandleFunc("/main/", mainHandler)
	r.HandleFunc("/otro/", othrHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8181",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//Esta funcion bloquea el programa.
	gowiki.Open("http://localhost:8181/main/")

	log.Fatal(srv.ListenAndServe())
	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	//	log.Fatal(http.ListenAndServe(":8181", nil))
}

func saveHandle2r(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body") //boyd es un string!/
	p := &gowiki.Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
