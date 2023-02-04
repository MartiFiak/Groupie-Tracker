package groupietrackers

import (
	"net/http"
	"strconv"
	"strings"
)

func GetFilterUse(r *http.Request, artistFiltered []Artist, artistLoad []Artist) []Artist {
	// ! Checks and applies the filters used for the search
	shearchFilter := r.FormValue("shearch")
	creationdateFilter := r.FormValue("creationdate")
	firstalbumdateFilte := r.FormValue("firstalbumdate")
	locationFilter := r.FormValue("locationfilter")
	nmemberFilter := []string{ // * Contains the value of all check-buttons
		r.FormValue("one_members"),
		r.FormValue("tow_members"),
		r.FormValue("tree_members"),
		r.FormValue("four_members"),
		r.FormValue("five_members"),
		r.FormValue("six_members"),
		r.FormValue("more_members"),
	}
	if shearchFilter != "" {
		if artistFiltered == nil {
			artistFiltered = WhichContainsString(shearchFilter, artistLoad)
		} else {
			artistFiltered = WhichContainsString(shearchFilter, artistFiltered)
		}
	}
	if creationdateFilter != "" {
		if artistFiltered == nil {
			artistFiltered = FiltredByCreationDate(artistLoad, "1800", creationdateFilter)
		} else {
			artistFiltered = FiltredByCreationDate(artistFiltered, "1800", creationdateFilter)
		}
	}
	if firstalbumdateFilte != "" {
		if artistFiltered == nil {
			artistFiltered = FiltredByFirstAlbum(artistLoad, "1800", firstalbumdateFilte)
		} else {
			artistFiltered = FiltredByFirstAlbum(artistFiltered, "1800", firstalbumdateFilte)
		}
	}
	nmemberFilter = CheckNumberSelect(nmemberFilter)
	if len(nmemberFilter) != 0 {
		if artistFiltered == nil {
			artistFiltered = FiltredByMembersNumber(artistLoad, nmemberFilter)
		} else {
			artistFiltered = FiltredByMembersNumber(artistFiltered, nmemberFilter)
		}
	}
	if locationFilter != "" {
		if artistFiltered == nil {
			artistFiltered = FiltredByLocations(artistLoad, locationFilter)
		} else {
			artistFiltered = FiltredByLocations(artistFiltered, locationFilter)
		}
	}

	return artistFiltered
}

func WhichContainsString(content string, listToCheck []Artist) []Artist { // ? Here we search in the names of the artists/groups
	// ! Check if a string contains a similarity
	containsString := []Artist{}
	for _, element := range listToCheck {
		if strings.Contains(TurnStringToShearch(element.Name), TurnStringToShearch(content)) {
			containsString = append(containsString, element)
		} else if strconv.Itoa(element.CreationDate) == content || strings.Split(element.FirstAlbum, "-")[2] == content {
			containsString = append(containsString, element)
		}else {
			if len(FiltredByLocations([]Artist{element}, content)) != 0 {
				containsString = append(containsString, element)
			} else if SearchMembers(element, content){
				containsString = append(containsString, element)
			}
		}

	}
	return containsString
}

func SearchMembers(artist Artist, search string) bool{
	for _, member := range artist.Member {
		if strings.Contains(TurnStringToShearch(member), TurnStringToShearch(search)){
			return true
		}
	}
	return false
}

func FiltredByMembersNumber(artistLoad []Artist, n []string) []Artist {
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		for _, nmembre := range n {
			if (len(artist.Member) == AtoiWithoutErr(nmembre) && nmembre != "7") || (nmembre == "7" && len(artist.Member) >= AtoiWithoutErr(nmembre)) {
				artistFiltered = append(artistFiltered, artist)
			}
		}
	}
	return artistFiltered
}

func FiltredByLocations(artistLoad []Artist, locationSearch string) []Artist {
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		locations := artist.FormatLocations
		for _, location := range locations {
			if strings.Contains(TurnStringToShearch(location), TurnStringToShearch(locationSearch)) {
				artistFiltered = append(artistFiltered, artist)
				break
			}
		}
	}
	return artistFiltered
}

func FiltredByCreationDate(artistLoad []Artist, mindate, maxdate string) []Artist {
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		if AtoiWithoutErr(mindate) <= artist.CreationDate && artist.CreationDate <= AtoiWithoutErr(maxdate) {
			artistFiltered = append(artistFiltered, artist)
		}

	}
	return artistFiltered
}

func FiltredByFirstAlbum(artistLoad []Artist, mindate, maxdate string) []Artist {
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		if (AtoiWithoutErr(mindate) <= AtoiWithoutErr(strings.Split(artist.FirstAlbum, "-")[2])) && (AtoiWithoutErr(strings.Split(artist.FirstAlbum, "-")[2]) <= AtoiWithoutErr(maxdate)) {
			artistFiltered = append(artistFiltered, artist)
		}

	}
	return artistFiltered
}

func CheckNumberSelect(n []string) []string {
	newn := []string{}
	for _, nc := range n {
		if nc != "" {
			newn = append(newn, nc)
		}
	}
	return newn
}

func TurnStringToShearch(str string) string {
	// ! Changes the "special" character string (which may contain " " and capital letters) into a "simple" character string (no spaces or capital letters)
	var nstr string
	for _, car := range str {
		switch {
		case 65 <= car && car <= 90:
			nstr = nstr + string(car+32)
		case car == 32:
		default:
			nstr = nstr + string(car)
		}
	}
	return nstr
}
