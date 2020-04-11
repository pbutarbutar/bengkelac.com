package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"bengkelac.com/libs"
	"bengkelac.com/libs/tplmanager"
	_ "github.com/qodrorid/godaemon"
)

type IndexPage struct {
	Title string
}
type UserData struct {
	Name        string
	City        string
	Nationality string
}

type SkillSet struct {
	Language string
	Level    string
}

type SkillSets []*SkillSet

func loadConfiguration(fileName string) {
	dirpath, _ := libs.Loadconfig("config.json")
	log.Println("layout path: ", dirpath.LayoutPath)
	log.Println("include path: ", dirpath.IncludePath)
	tplmanager.SetTemplateConfig(dirpath.LayoutPath, dirpath.IncludePath)
}

func index(w http.ResponseWriter, r *http.Request) {

	/*if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}*/

	titleHome := &IndexPage{Title: "Service AC/KULKAS/MESIN CUCI DAERAH CIPUTATA, SAWANGAN, PONDOK CABE, PAMULANG, PONDOK INDAH"}

	err := tplmanager.RenderTemplate(w, "index.tmpl", titleHome)
	if err != nil {
		log.Println(err)
	}
}

func aboutMe(w http.ResponseWriter, r *http.Request) {

	userData := &UserData{Name: "Asit Dhal", City: "Bhubaneswar", Nationality: "Indian"}
	err := tplmanager.RenderTemplate(w, "aboutme.tmpl", userData)
	if err != nil {
		log.Println(err)
	}
}

func skillSet(w http.ResponseWriter, r *http.Request) {

	skillSets := SkillSets{&SkillSet{Language: "Golang", Level: "Beginner"},
		&SkillSet{Language: "C++", Level: "Advanced"},
		&SkillSet{Language: "Python", Level: "Advanced"}}
	err := tplmanager.RenderTemplate(w, "skillset.tmpl", skillSets)
	if err != nil {
		log.Println(err)
	}
}

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	template.Must(template.ParseGlob("views/*.tmpl"))
	loadConfiguration("config.json")
	tplmanager.LoadTemplates()

	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	fmt.Println("Running port 8081")
	http.HandleFunc("/", index)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/skillset", skillSet)
	server.ListenAndServe()
}
