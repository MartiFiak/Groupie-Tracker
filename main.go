package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
	"net/http"
	"text/template"
)

var currentband groupietrackers.CurrentBand

func main() {

	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/index.html"))

	data := groupietrackers.SetGlobalData(groupietrackers.GetAPIData("https://groupietrackers.herokuapp.com/api"))
	dataartist := groupietrackers.SetArtistData(groupietrackers.GetAPIData(data.Artist + "/15"))
	datalocation := groupietrackers.SetLocationData(groupietrackers.GetAPIData(dataartist.Locations))
	datadate := groupietrackers.SetDateData(groupietrackers.GetAPIData(dataartist.ConcertDate))
	datadatelocation := groupietrackers.SetRelationData(groupietrackers.GetAPIData(dataartist.Relations))
	fmt.Println("Logo :", dataartist.Image, "\nLe groupe : ", dataartist.Name, "\nCréé en : ", dataartist.CreationDate, "\nA pour membre :", dataartist.Member, "\nLeur premier album est paru en :", dataartist.FirstAlbum)
	fmt.Println("Location : ", datalocation.Locations, "\nDate de Concert : ", datadate.Dates, "\nDateRelation : ", datadatelocation.DatesLocations)

	currentband.Image = dataartist.Image
	currentband.Name = dataartist.Name
	currentband.Member = dataartist.Member
	currentband.CreationDate = dataartist.CreationDate
	currentband.Relations = datadatelocation.DatesLocations

	tmpl.Execute(w, currentband)
}
