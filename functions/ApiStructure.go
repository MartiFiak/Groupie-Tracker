package groupietrackers

type CurrentBand struct {
	Id            int
	Name          string
	Image         string
	Member        []string
	CreationDate  int
	Relations     map[string][][][]string
	FuturRelation map[string][]Event
	PassRelation  [][]string
}

type Event struct {
	Country string
	City    string
	Date    []string
	Coord   GeoCoord
}

type GeoCoord struct {
	Lat  float64
	Long float64
}

type PageData struct {
	Currentband  CurrentBand
	Artists      []Artist
	MPageRArtist []Artist
	Locations    []string
}

type Artist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	CreationDate    int      `json:"creationdate"`
	Member          []string `json:"members"`
	FirstAlbum      string   `json:"firstAlbum"`
	Locations       string   `json:"locations"`
	FormatLocations []string
}

type ApiData struct {
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
	Dates     string   `json:"dates"`
}

type datesStruct struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type relationStruct struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type User struct {
	Username  string
	Password  string
	Favorites []string
}
