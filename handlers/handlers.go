package asciiartweb

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	asciiartweb "asciiartweb/Functions"
)

type Error struct {
	Status string
	Type   string
}

type DataAscii struct {
	Value   string
	Result  string
	Banner1 string
	Banner2 string
	Banner3 string
}

func GatherBannerData() *DataAscii {
	Data := &DataAscii{}
	files, _ := os.ReadDir("Files/")
	for _, file := range files {
		if file.Name() == "standard.txt" {
			Data.Banner1 = file.Name()
		} else if file.Name() == "shadow.txt" {
			Data.Banner2 = file.Name()
		} else if file.Name() == "thinkertoy.txt" {
			Data.Banner3 = file.Name()
		}
	}
	return Data
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == http.MethodGet {
			Data := GatherBannerData()
			RenderTemplate(w, "./templates/index.html", Data, http.StatusOK)
		} else {
			a := Error{Status: "405", Type: "Method Not Allowed"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusMethodNotAllowed)
		}
	default:
		a := Error{Status: "404", Type: "Page Not Found"}
		RenderTemplate(w, "./templates/errorPage.html", a, http.StatusNotFound)
	}
}

func AsciiArtPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" || banner == "" {
			a := Error{Status: "400", Type: "Bad Request"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusBadRequest)
			return
		}

		if len([]rune(text)) > 200 {
			a := Error{Status: "400", Type: "Bad Request"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusBadRequest)
			return
		}

		result := asciiartweb.AsciiArt(text, banner)

		if result == "Banner not found" || result == "All caracters" {
			a := Error{Status: "500", Type: "Server Error"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusInternalServerError)
			return
		} else if result == "Special charactere is not allowed." {
			a := Error{Status: "400", Type: "Bad Request"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusBadRequest)
			return
		}

		Data := GatherBannerData()

		Data.Value = text
		Data.Result = "\n" + result + "\n"

		err := RenderTemplate(w, "./templates/index.html", Data, http.StatusOK)
		if err != nil {
			a := Error{Status: "500", Type: "Server Error"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusInternalServerError)
			return
		}
	} else {
		a := Error{Status: "405", Type: "Method Not Allowed"}
		RenderTemplate(w, "./templates/errorPage.html", a, http.StatusMethodNotAllowed)
		return
	}
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		files, err := os.ReadDir("assets/css")
		if err != nil {
			a := Error{Status: "500", Type: "Server Error"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusInternalServerError)
			return
		}
		for _, file := range files {
			if r.URL.Path == "/assets/css/"+file.Name() {
				fs := http.Dir("assets/css")
				http.StripPrefix("/assets/css", http.FileServer(fs)).ServeHTTP(w, r)
				return
			}
		}

		a := Error{Status: "404", Type: "Page Not Found"}
		RenderTemplate(w, "./templates/errorPage.html", a, http.StatusNotFound)
		return
	}

	a := Error{Status: "405", Type: "Method Not Allowed"}
	RenderTemplate(w, "./templates/errorPage.html", a, http.StatusMethodNotAllowed)
}

func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		files, err := os.ReadDir("assets/images")
		if err != nil {
			a := Error{Status: "500", Type: "Server Error"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusInternalServerError)
			return
		}
		for _, file := range files {
			if r.URL.Path == "/assets/images/"+file.Name() {
				fs := http.Dir("assets/images")
				http.StripPrefix("/assets/images", http.FileServer(fs)).ServeHTTP(w, r)
				return
			}
		}

		a := Error{Status: "404", Type: "Page Not Found"}
		RenderTemplate(w, "./templates/errorPage.html", a, http.StatusNotFound)
		return
	}

	a := Error{Status: "405", Type: "Method Not Allowed"}
	RenderTemplate(w, "./templates/errorPage.html", a, http.StatusMethodNotAllowed)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data any, status int) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}
	return nil
}
