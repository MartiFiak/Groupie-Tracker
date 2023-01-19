package groupietrackers

import (
	"strings"
)

func GetArtistWithStr(shearchFilter string, artistLoad []Artist) []Artist {
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		if strings.Contains(TurnStringToShearch(artist.Name), TurnStringToShearch(shearchFilter)) {
			artistFiltered = append(artistFiltered, artist)
		}
	}
	return artistFiltered
}

func FiltredByMembersNumber(artistLoad []Artist, n []string) []Artist {
	artistFiltered := []Artist{}
	for _, artist := range artistLoad {
		for _, nmembre := range n {
			if (len(artist.Member) == AtoiWithoutErr(nmembre) && nmembre != "7") || (nmembre == "7" && len(artist.Member) >= AtoiWithoutErr(nmembre)){
				artistFiltered = append(artistFiltered, artist)
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

func CheckNumberSelect(n []string)[]string{
	newn := []string{}
	for _,nc := range n{
		if nc != ""{
			newn = append(newn, nc)
		}
	}
	return newn
}

func TurnStringToShearch(str string) string {
	/*       Turn :  fdsfKJHJUGKHLJ dsf ezrtf _è-'4941 into : fdsfkjhjugkhljdsfezrtf_è-'4941*/
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
