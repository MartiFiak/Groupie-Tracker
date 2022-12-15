package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
)

func main() {
	data := groupietrackers.SetGlobalData(groupietrackers.GetAPIData("https://groupietrackers.herokuapp.com/api"))
	dataartist := groupietrackers.SetArtistData(groupietrackers.GetAPIData(data.Artist + "/15"))
	datalocation := groupietrackers.SetLocationData(groupietrackers.GetAPIData(dataartist.Locations))
	datadate := groupietrackers.SetDateData(groupietrackers.GetAPIData(dataartist.ConcertDate))
	datadatelocation := groupietrackers.SetRelationData(groupietrackers.GetAPIData(dataartist.Relations))
	fmt.Println("Logo :", dataartist.Image, "\nLe groupe : ", dataartist.Name, "\nCréé en : ", dataartist.CreationDate, "\nA pour membre :", dataartist.Member, "\nLeur premier album est paru en :", dataartist.FirstAlbum)
	fmt.Println("Location : ", datalocation.Locations, "\nDate de Concert : ", datadate.Dates, "\nDateRelation : ", datadatelocation.DatesLocations)
}
