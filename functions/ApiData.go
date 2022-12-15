package groupietrackers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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


func SetGlobalData(body []byte) apiData {
	apidata := apiData{}
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
