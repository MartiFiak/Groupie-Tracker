package groupietrackers

type apiData struct {
	Artist    string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type artistStruct struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Member       []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type locationsStruct struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string      `json:"dates"`
}

type datesStruct struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type relationStruct struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}