package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
	"net/http"
	"strconv"
	"strings"
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
	dataartist := groupietrackers.SetArtistData(groupietrackers.GetAPIData(data.Artist + "/3"))
	datalocation := groupietrackers.SetLocationData(groupietrackers.GetAPIData(dataartist.Locations))
	datadate := groupietrackers.SetDateData(groupietrackers.GetAPIData(dataartist.ConcertDate))
	datadatelocation := groupietrackers.SetRelationData(groupietrackers.GetAPIData(dataartist.Relations))
	fmt.Println("Logo :", dataartist.Image, "\nLe groupe : ", dataartist.Name, "\nCréé en : ", dataartist.CreationDate, "\nA pour membre :", dataartist.Member, "\nLeur premier album est paru en :", dataartist.FirstAlbum)
	fmt.Println("Location : ", datalocation.Locations, "\nDate de Concert : ", datadate.Dates, "\nDateRelation : ", datadatelocation.DatesLocations)

	currentband.Image = dataartist.Image
	currentband.Name = dataartist.Name
	currentband.Member = dataartist.Member
	// dd/mm/yyyy
	currentband.CreationDate = dataartist.CreationDate
	currentband.Relations = ChangeDateFormat(datadatelocation.DatesLocations)

	tmpl.Execute(w, currentband)
}

func ChangeDateFormat(date map[string][]string) map[string][][]string {
	nDate := make(map[string][][]string)

	for location, ldate := range date {
		if len(ldate) > 1 {
			var lDate [][]string
			var allDate [][]string
			for _, d := range ldate {
				allDate = append(allDate, strings.Split(d, "-"))
			}
			var day string
			var mday string
			var lday int
			var month string
			var year string

			for i, d := range allDate {
				if i == 0 {
					day = d[0]
					mday = d[0]
					lday, _ = strconv.Atoi(day)
					month = d[1]
					year = d[2]
				} else {
					nlday, _ := strconv.Atoi(d[0])
					if month == d[1] && year == d[2] && lday-1 == nlday {
						day = d[0] + "-" + mday
						lday--
					} else {
						imonth, _ := strconv.Atoi(month)
						lDate = append(lDate, []string{day, []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}[imonth], year})
						day = d[0]
						mday = d[0]
						lday, _ = strconv.Atoi(day)
						month = d[1]
						year = d[2]
					}
				}
			}

			nDate[location] = lDate
		} else {
			var lDate [][]string
			imonth, _ := strconv.Atoi(strings.Split(ldate[0], "-")[1])
			lDate = append(lDate, []string{strings.Split(ldate[0], "-")[0], []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}[imonth],strings.Split(ldate[0], "-")[2]})
			nDate[location] = lDate
		}
	}

	return nDate
}
