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
	if res.Body != nil {
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

func SetArtistInfoData(body []byte) Artist {
	artistdata := Artist{}
	jsonErr := json.Unmarshal(body, &artistdata)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return artistdata
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
	var cb CurrentBand
	dataartist := SetArtistData(GetAPIData(apiArtist))
	datadatelocation := SetRelationData(GetAPIData(dataartist.Relations))
	cb.Id = dataartist.Id
	cb.Image = dataartist.Image
	cb.Name = dataartist.Name
	cb.Member = dataartist.Member
	cb.CreationDate = dataartist.CreationDate
	cb.Relations = ChangeDateFormat(datadatelocation.DatesLocations)
	cb.FuturRelation, cb.PassRelation = CheckRelationTime(cb.Relations)
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

func CheckRelationTime(date map[string][][][]string) (map[string][][][]string, map[string][][][]string) {
	var fRelation map[string][][][]string
	var pRelation map[string][][][]string

	for _, pays := range date {
		fmt.Println(pays)
	}

	return fRelation, pRelation
}
