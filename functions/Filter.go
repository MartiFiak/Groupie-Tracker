package groupietrackers

import (
	"strings"
)

func GetArtistWithStr(shearchFilter string, artistLoad  []Artist) []Artist{
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		if strings.Contains(TurnStringToShearch(artist.Name), TurnStringToShearch(shearchFilter)) {
			artistFiltered = append(artistFiltered, artist)
		}
	}
	return artistFiltered
}

func TurnStringToShearch(str string) string {
	/*       Turn :  fdsfKJHJUGKHLJ dsf ezrtf _è-'4941 into : fdsfkjhjugkhljdsfezrtf_è-'4941*/
	var nstr string
	for _, car := range str {
		switch {
		case 65 <= car && car <= 90:
			nstr = nstr + string(car+32)
		case car == 32:
			continue
		default:
			nstr = nstr + string(car)
		}
	}
	return nstr
}