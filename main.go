package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
	"strings"
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

	go RealtimeData() /*       Permet de récuperer les artists en parallèle de la gestion de nos pages        */
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

	switch r.Method {
	case "GET":
	case "POST":
		pageData.Artists = groupietrackers.GetFilterUse(r, artistFiltered, artistLoad)
		artistFiltered = nil
	}

	tmpl.Execute(w, pageData)
}

func RealtimeData() {
	// !       Regenere les données des artistes toutes les minutes
	for {
		artistLoad = groupietrackers.SetArtist(groupietrackers.GetAPIData(data.Artist))
		for index, artist := range artistLoad {
			locations := groupietrackers.SetLocationData(groupietrackers.GetAPIData(artist.Locations))
			for _, location := range locations.Locations {
				ville_pays := strings.Split(location, "-")
				ville_pays[0] = strings.Replace(ville_pays[0], "_", " ", -1)
				ville_pays[1] = strings.Replace(ville_pays[1], "_", " ", -1)
				if !groupietrackers.SContains(pageData.Locations, ville_pays[0]){
					pageData.Locations = append(pageData.Locations, ville_pays[0])
				}
				if !groupietrackers.SContains(pageData.Locations, ville_pays[1]){
					pageData.Locations = append(pageData.Locations, ville_pays[1])
				}
				artistLoad[index].FormatLocations = append(artistLoad[index].FormatLocations, ville_pays[0] + " " + ville_pays[1])
			}
		}
		time.Sleep(60 * time.Second)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/signup.html"))

	switch r.Method {
	case "GET":
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		confpassword := r.FormValue("confirmpassword")

		if password == confpassword {
			if username == "" || password == "" {
				fmt.Println("Please all field nead to be write")
			} else {
				if groupietrackers.GetUserData(username).Username != "" {
					fmt.Println("Error User Already exist")
				} else {
					newUser := groupietrackers.User{
						Username: username,
						Password: password,
					}
					groupietrackers.AddUser(newUser)
				}
			}
		} else {
			fmt.Println("Password don't match")
		}

	}

	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/login.html"))

	switch r.Method {
	case "GET":
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		if password != "" {
			if username != ""{
				if groupietrackers.GetUserData(username).Username != "" && groupietrackers.GetUserData(username).Password == password {
					pageData.CurrentUser = groupietrackers.GetUserData(username)
				}
			}
		}
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
		pageData.Artists = groupietrackers.GetFilterUse(r, artistFiltered, artistLoad)
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
		pageData.Artists = groupietrackers.GetFilterUse(r, artistFiltered, artistLoad)
		artistFiltered = nil
	}

	if currentID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	currentband = groupietrackers.UpdateCurrentBand(data.Artist + "/" + strconv.Itoa(currentID)) // ? Erreur au lancement
	pageData.Currentband = currentband
	tmpl.Execute(w, pageData)
}
