package gowiki

import (
	"html/template"
	"os/exec"
	"runtime"
)

//jurarse experto
func ParseFiles(structdata string, name string) (*template.Template, error) {
	return ParseFiles2(name, nil, structdata)
}

/*
ParseFiles2 es un intento de diferente funcion con la libreria que abre o genera un
template html sin archivos. //intentemos hacer una pagina con 3 direcciones y de gran volumen html, compilar y probar
en otro computador
*/
func ParseFiles2(name2 string, t *template.Template, structdata string) (*template.Template, error) {
	s := structdata //tiene que ser un string
	name := name2

	var tmpl *template.Template

	if t == nil {
		t = template.New(name)
	}

	if name == t.Name() {
		tmpl = t
	} else {
		tmpl = t.New(name2)
	}
	_, err := tmpl.Parse(s)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// open opens the specified URL in the default browser of the user.
func Open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
