package groupietrackers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var FakeCurrentYear int
var FakeCurrentMonth time.Month
var FakeCurrentDay int

func GetAPIData(apiUrl string) []byte {

	apiClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User", "groupie-tracker")
	res, getErr := apiClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}
	if res.Body != nil { // ?  Get "": unsupported protocol scheme "" 2023/01/22 23:48:26 http: panic serving [::1]:60987: runtime error: invalid memory address or nil pointer dereference
		// ? Get "https://groupietrackers.herokuapp.com/api": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	return body

}

func SetGlobalData(body []byte) ApiData {
	apidata := ApiData{}
	jsonErr := json.Unmarshal(body, &apidata)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return apidata
}

func GetGeoCodeData(body []byte) map[string]interface{} {
	ap := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &ap)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return ap
}

func GetCoord(geoCodeData map[string]interface{}) GeoCoord {
	coord := GeoCoord{}
	test := geoCodeData["results"].([]interface{})
	test2 := test[0].(map[string]interface{})
	test3 := test2["geometry"].(map[string]interface{})
	test4 := test3["location"].(map[string]interface{})
	coord.Lat = test4["lat"].(float64)
	coord.Long = test4["lng"].(float64)
	return coord
}

func SetArtistData(body []byte) artistStruct {
	artistdata := artistStruct{}
	jsonErr := json.Unmarshal(body, &artistdata)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return artistdata
}

func SetLocationData(body []byte) locationsStruct {
	locationdata := locationsStruct{}
	jsonErr := json.Unmarshal(body, &locationdata)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return locationdata
}

func SetDateData(body []byte) datesStruct {
	datedata := datesStruct{}
	jsonErr := json.Unmarshal(body, &datedata)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return datedata
}

func SetRelationData(body []byte) relationStruct {
	relationdata := relationStruct{}
	jsonErr := json.Unmarshal(body, &relationdata)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return relationdata
}

func SetArtist(body []byte) []Artist {
	artistd := []Artist{}
	jsonErr := json.Unmarshal(body, &artistd)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return artistd
}

func UpdateCurrentBand(apiArtist string) CurrentBand {
	//var cb CurrentBand
	dataartist := SetArtistData(GetAPIData(apiArtist))
	datadatelocation := SetRelationData(GetAPIData(dataartist.Relations)) // ? Erreur au lancement   Get "": unsupported protocol scheme "" 2023/01/22 23:48:26 http: panic serving [::1]:60989: runtime error: invalid memory address or nil pointer dereference

	cb := CurrentBand{
		Id:           dataartist.Id,
		Name:         dataartist.Name,
		Image:        dataartist.Image,
		Member:       dataartist.Member,
		CreationDate: dataartist.CreationDate,
		Relations:    ChangeDateFormat(datadatelocation.DatesLocations),
	}
	cb.FuturRelation, cb.PassRelation = CheckRelationTime(cb.Relations)
	/*cb.Id = dataartist.Id
	cb.Image = dataartist.Image
	cb.Name = dataartist.Name
	cb.Member = dataartist.Member
	cb.CreationDate = dataartist.CreationDate
	cb.Relations = ChangeDateFormat(datadatelocation.DatesLocations)
	cb.FuturRelation, cb.PassRelation = CheckRelationTime(cb.Relations)*/
	return cb
}

func ChangeDateFormat(date map[string][]string) map[string][][][]string {
	nDate := make(map[string][][][]string)

	for location, ldate := range date {
		var ville string
		var pays string

		if len(ldate) > 1 {
			lDate := [][]string{}
			var allDate [][]string
			for _, d := range ldate {
				allDate = append(allDate, strings.Split(d, "-"))
			}
			var day string
			var mday string
			var lday int
			var month string
			var year string

			pays = strings.Split(location, "-")[1]
			ville = strings.Split(location, "-")[0]

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
						if i == len(allDate)-1 {
							imonth, _ := strconv.Atoi(month)
							lDate = append(lDate, []string{ville, day, []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}[imonth-1], year})
						}
					} else {
						imonth, _ := strconv.Atoi(month)
						lDate = append(lDate, []string{ville, day, []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}[imonth-1], year})
						day = d[0]
						mday = d[0]
						lday, _ = strconv.Atoi(day)
						month = d[1]
						year = d[2]
					}
				}
			}

			nDate[pays] = append(nDate[pays], lDate)
		} else {
			pays = strings.Split(location, "-")[1]
			ville = strings.Split(location, "-")[0]

			lDate := [][]string{}
			imonth, _ := strconv.Atoi(strings.Split(ldate[0], "-")[1])
			lDate = append(lDate, []string{ville, strings.Split(ldate[0], "-")[0], []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}[imonth-1], strings.Split(ldate[0], "-")[2]})
			nDate[pays] = append(nDate[pays], lDate)
		}
	}
	return nDate
}

func SetCoordToEvent(event map[string][]Event) map[string][]Event {
	for pays := range event {
		for i, e := range event[pays] {
			e.Coord = GetCoord(GetGeoCodeData(GetAPIData("https://maps.googleapis.com/maps/api/geocode/json?address=" + e.City + "+" + e.Country + "&key=AIzaSyBq9H9P3Jazc6tUoqQ8fwBdMbgLhm0QSe4")))
			event[pays][i] = FormatFLocation(e)
		}
	}
	return event
}

func FormatFLocation(event Event) Event {
	city := strings.Split(event.City, "_")
	for i, value := range city {
		ncity := ""
		for j, lettre := range value {
			if j == 0 {
				ncity += string(lettre - 32)
			} else {
				ncity += string(lettre)
			}
		}
		city[i] = ncity
	}
	event.City = strings.Join(city, " ")

	country := strings.Split(event.Country, "_")
	for i, value := range country {
		ncountry := ""
		for j, lettre := range value {
			if j == 0 {
				ncountry += string(lettre - 32)
			} else {
				ncountry += string(lettre)
			}
		}
		country[i] = ncountry
	}
	event.Country = strings.Join(country, " ")

	return event
}

func CheckRelationTime(date map[string][][][]string) (map[string][]Event, [][]string) {

	FakeCurrentYear, FakeCurrentMonth, FakeCurrentDay = time.Now().Date()
	FakeCurrentYear -= 3
	fRelation := make(map[string][]Event)
	pRelation := [][]string{}

	for pays := range date {
		for _, location := range date[pays] {
			for _, rlocation := range location {
				checkEvent := Event{}
				switch {
				case AtoiWithoutErr(rlocation[3]) >= FakeCurrentYear:
					checkEvent.Country = pays
					checkEvent.City = rlocation[0]
					checkEvent.Date = rlocation[1:]
					fRelation[pays] = append(fRelation[pays], checkEvent)
				case AtoiWithoutErr(rlocation[3]) < FakeCurrentYear:
					rlocation = append(rlocation, pays)
					pRelation = append(pRelation, rlocation)
				default:
				}
			}
		}
	}

	pRelation = sortByNIndex(3, sortByNIndex(1, pRelation))
	if len(pRelation) >= 3 {
		pRelation = pRelation[:3]
	}
	/*fRelation = FormatFLocation(fRelation)*/
	fRelation = SetCoordToEvent(fRelation)
	return fRelation, pRelation
}
