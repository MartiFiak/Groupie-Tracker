package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var currentband groupietrackers.CurrentBand
var pageData groupietrackers.PageData
var artistLoad []groupietrackers.Artist
var artistFiltered []groupietrackers.Artist
var currentID int
var data groupietrackers.ApiData

var FakeCurrentYear int
var FakeCurrentMonth time.Month
var FakeCurrentDay int

func main() {

	FakeCurrentYear, FakeCurrentMonth, FakeCurrentDay = time.Now().Date()
	FakeCurrentYear -= 3
	fmt.Println("Date Simulated :", FakeCurrentDay, FakeCurrentMonth, FakeCurrentYear)

	data = groupietrackers.SetGlobalData(groupietrackers.GetAPIData("https://groupietrackers.herokuapp.com/api"))

	go RealtimeData() /*       Permet de récuperer les artists en arrière plan sans ralentir l'affichage de la page         */
	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/allArtist", AllArtistHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/signup", SignUpHandler)
	http.ListenAndServe(":8080", nil)
}
func AllArtistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/allArtist.html", "./server/component/sidebar.html"))

	tmpl.Execute(w, nil)
}


func RealtimeData() {
	for { /*       Regenere les données des artistes toutes les minutes        */
		artistLoad = groupietrackers.SetArtist(groupietrackers.GetAPIData(data.Artist))
		time.Sleep(60 * time.Second)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./Signup.html"))

	switch r.Method {
	case "GET":
	case "POST":
	}

	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./Login.html"))

	switch r.Method {
	case "GET":
	case "POST":
	}
	tmpl.Execute(w, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/index.html", "./server/component/sidebar.html"))

	switch r.Method {
	case "GET":
		if r.URL.Query().Get("id") != "" {
			currentID, _ = strconv.Atoi(r.URL.Query().Get("id"))
		}
		pageData.Artists = artistLoad
	case "POST":
		shearchFilter := r.FormValue("shearch")
		creationdateFilter := r.FormValue("creationdate")
		nmemberFilter := []string{r.FormValue("one_members"), r.FormValue("tow_members"), r.FormValue("tree_members"), r.FormValue("four_members"), r.FormValue("five_members"), r.FormValue("six_members"), r.FormValue("more_members")}
		if shearchFilter != "" {
			if artistFiltered == nil {
				artistFiltered = groupietrackers.GetArtistWithStr(shearchFilter, artistLoad)
			} else {
				artistFiltered = groupietrackers.GetArtistWithStr(shearchFilter, artistFiltered)
			}
		}
		if creationdateFilter != "" {
			if artistFiltered == nil {
				artistFiltered = groupietrackers.FiltredByCreationDate(artistLoad, "1800", creationdateFilter)
			} else {
				artistFiltered = groupietrackers.FiltredByCreationDate(artistFiltered, "1800", creationdateFilter)
			}
		}
		nmemberFilter = groupietrackers.CheckNumberSelect(nmemberFilter)
		if len(nmemberFilter) != 0 {
			if artistFiltered == nil {
				artistFiltered = groupietrackers.FiltredByMembersNumber(artistLoad, nmemberFilter)
			} else {
				artistFiltered = groupietrackers.FiltredByMembersNumber(artistFiltered, nmemberFilter)
			}
		}
		pageData.Artists = artistFiltered
		artistFiltered = nil
	}
	rand.Seed(time.Now().UnixNano())
	pageData.MPageRArtist = []groupietrackers.Artist{}
	mi := 3
	if len(artistLoad) < 3 {
		mi = len(artistLoad)
	}
	for i := 0; i < mi; i++ {
		alreadyin := false
		artistrandom := artistLoad[rand.Intn(len(artistLoad))]
		for _, selectartiste := range pageData.MPageRArtist {
			if selectartiste.Id == artistrandom.Id {
				i--
				alreadyin = true
			}
		}
		if !alreadyin {
			pageData.MPageRArtist = append(pageData.MPageRArtist, artistrandom)
		}
	}

	tmpl.Execute(w, pageData)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/artist.html", "./server/component/sidebar.html"))

	switch r.Method {
	case "GET":
		if r.URL.Query().Get("id") != "" {
			currentID, _ = strconv.Atoi(r.URL.Query().Get("id"))
		}
		pageData.Artists = artistLoad
	case "POST":
		shearchFilter := r.FormValue("shearch")
		creationdateFilter := r.FormValue("creationdate")
		fmt.Println(shearchFilter)
		fmt.Println(creationdateFilter)
		if shearchFilter != "" {
			if artistFiltered == nil {
				artistFiltered = groupietrackers.GetArtistWithStr(shearchFilter, artistLoad)
			} else {
				artistFiltered = groupietrackers.GetArtistWithStr(shearchFilter, artistFiltered)
			}
		}
		if creationdateFilter != "" {
			if artistFiltered == nil {
				artistFiltered = groupietrackers.FiltredByCreationDate(artistLoad, "1800", creationdateFilter)
			} else {
				artistFiltered = groupietrackers.FiltredByCreationDate(artistFiltered, "1800", creationdateFilter)
			}
		}
		pageData.Artists = artistFiltered
		artistFiltered = nil
	}

	if currentID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	currentband = groupietrackers.UpdateCurrentBand(data.Artist + "/" + strconv.Itoa(currentID))
	pageData.Currentband = currentband
	tmpl.Execute(w, pageData)
}

func GetArtistXtoY(x, y int, apiArtist string) {
	for i := x; i <= y; i++ {
		if i > len(artistLoad) {
			artist := groupietrackers.SetArtistInfoData(groupietrackers.GetAPIData(apiArtist + "/" + strconv.Itoa(i)))
			artistLoad = append(artistLoad, artist)
		}
	}
}
