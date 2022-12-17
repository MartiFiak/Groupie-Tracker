package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
	"net/http"
	"strconv"
	"text/template"
)

var currentband groupietrackers.CurrentBand
var pageData groupietrackers.PageData
var currentID int

func main() {

	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/index.html"))

	if currentID == 0 {
		currentID = 1
	}

	data := groupietrackers.SetGlobalData(groupietrackers.GetAPIData("https://groupietrackers.herokuapp.com/api"))

	switch r.Method {
	case "GET":
		if r.URL.Query().Get("id") != "" {
			currentID, _ = strconv.Atoi(r.URL.Query().Get("id"))
			fmt.Println()
		}
	}
	/*dataartist := groupietrackers.SetArtistData(groupietrackers.GetAPIData(data.Artist + "/" + strconv.Itoa(currentID)))
	datalocation := groupietrackers.SetLocationData(groupietrackers.GetAPIData(dataartist.Locations))
	datadate := groupietrackers.SetDateData(groupietrackers.GetAPIData(dataartist.ConcertDate))
	datadatelocation := groupietrackers.SetRelationData(groupietrackers.GetAPIData(dataartist.Relations))
	fmt.Println("Logo :", dataartist.Image, "\nLe groupe : ", dataartist.Name, "\nCréé en : ", dataartist.CreationDate, "\nA pour membre :", dataartist.Member, "\nLeur premier album est paru en :", dataartist.FirstAlbum)
	fmt.Println("Location : ", datalocation.Locations, "\nDate de Concert : ", datadate.Dates, "\nDateRelation : ", datadatelocation.DatesLocations)*/
	GetArtistXtoY(1, 10, data.Artist)
	currentband = groupietrackers.UpdateCurrentBand(data.Artist + "/" + strconv.Itoa(currentID))
	pageData.Currentband = currentband

	tmpl.Execute(w, pageData)
}

func GetArtistXtoY(x, y int, apiArtist string) {
	pageData.Artists = []groupietrackers.Artist{}
	for i := x; i <= y; i++ {
		artist := groupietrackers.SetArtistInfoData(groupietrackers.GetAPIData(apiArtist + "/" + strconv.Itoa(i)))
		pageData.Artists = append(pageData.Artists, artist)
	}
}
