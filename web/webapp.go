package web

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"primitive-webapp/primitive"
	"strconv"
)

type Data struct {
	Name string
}

func SetServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/image/", showImage)
	fs := http.FileServer(http.Dir("./web/img/"))
	mux.Handle("/img/", http.StripPrefix("/img/", fs))

	http.ListenAndServe(":8000", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("web/templates/home.html"))
	tpl.Execute(w, "")
}

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	numshapesstr := r.FormValue("numshapes")
	shapestr := r.FormValue(("shape"))

	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]

	f, err := ioutil.TempFile("./web/img/", fmt.Sprintf("in_img*.%s", ext))
	checkError(err)
	defer f.Close()

	_, err = io.Copy(f, file)
	checkError(err)

	fo, err := ioutil.TempFile("./web/img/", fmt.Sprintf("out_img*.%s", ext))
	checkError(err)
	defer fo.Close()

	numshapes, err := strconv.Atoi(numshapesstr)
	checkError(err)

	shape, err := strconv.Atoi(shapestr)
	checkError(err)

	_, err = primitive.Primitive(f.Name(), fo.Name(), numshapes, primitive.WithMode(primitive.Mode(shape)))
	checkError(err)

	http.Redirect(w, r, fmt.Sprintf("/image/%s", filepath.Base(fo.Name())), http.StatusFound)

}

func showImage(w http.ResponseWriter, r *http.Request) {
	
	data := []Data{
		{Name: filepath.Base(r.URL.Path)},
	}
	tpl := template.Must(template.ParseFiles("web/templates/showimage.html"))
	tpl.Execute(w, data)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// func makeTempFile(file io.Reader, prefix, suffix string) (string, error) {
// 	f, err := ioutil.TempFile("./web/img/", fmt.Sprintf("%s_img*.%s", prefix, suffix))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close()

// 	_, err = io.Copy(f, file)
// 	if err != nil {
// 		return "", err
// 	}
// 	return f.Name(), nil
// }
