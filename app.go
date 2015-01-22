package main

import "fmt"
import m "github.com/yundream/simpleweb/lib/database"

import "io/ioutil"
import "net/http"

func loadPage(filename string) []byte {
	contents, _ := ioutil.ReadFile(filename)
	return contents
}

func htmlhandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("html" + r.URL.String())
	w.Write(body)
}

type Handler struct {
	dir   string
	ctype string
}

func (h Handler) Run(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile(h.dir + r.URL.String())
	w.Header().Add("content-type", h.ctype)
	w.Write(body)
}

func main() {
	fmt.Println("Simple Web Server")
	db := new(m.Database)
	db.Open("/data")
	//db.FindVcard("yundream.json")
	//handle := Handler{}
	http.HandleFunc("/", Handler{dir: "html", ctype: "text/html"}.Run)
	http.HandleFunc("/css/", Handler{dir: "public", ctype: "text/css"}.Run)
	http.HandleFunc("/js/", Handler{dir: "public", ctype: "application/javascript"}.Run)

	http.ListenAndServe(":8080", nil)
}
