package main

import (
	"rt/colors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"image"
	"image/png"
	"os"
)

func main() {
	port := ":3000"
	fs := http.FileServer(http.Dir("./assets"))
	
	log.Printf("Start Config Scene...\n")
	// Parser the scene
	log.Printf("Start Render Scene...\n")
	render()
	// Render the rt
	log.Printf("Set Scene in Web...\n")
	// Functions for set the scene in example.html
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", serveTemplate)
	log.Printf("Wath Scene on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func render() {
	var c rtcolors.RtColorRGB
	c.RtColorRGB(0.5, 0.0, 0.7)
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			img.Set(x, y, c.GetPixel())
		}
	}
	arq, _ := os.Create("./assets/image.png")
	defer arq.Close()
	png.Encode(arq, img)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmp := "assets/templates"
	lp := filepath.Join(tmp, "layout.html")
	fp := filepath.Join(tmp, filepath.Clean(r.URL.Path))

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
