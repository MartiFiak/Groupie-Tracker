class Artist{
    constructor(id,image,name,creationdate,firstalbum, members, locations){
        this.id;
        this.image;
        this.name;
        this.creationdate;
        this.firstalbum;
        this.members;
        this.locations;
    }
}


document.addEventListener("keyup", function(event) {
    _executeFilter()
});


let allArtist = {}
let currentArtist = []

function _addArtist(id, image, name, creationDate, firstAlbum, members, locations){
    allArtist[id] = {id:id,image:image,name:name,creationDate:creationDate,firstalbum:firstAlbum, members:members, locations:locations};
    displayArtist(id, image,name, creationDate, "artist_band");
}

function _creationDateSlider(){
    _rangeManager(document.getElementById("myRangeMax"),document.getElementById("myRangeMin"), document.getElementById("cdTxt"));
    return [document.getElementById("myRangeMin").value, document.getElementById("myRangeMax").value];
}

function _FirstAlbumSlider(){
    _rangeManager(document.getElementById("myRangeFAMax"),document.getElementById("myRangeFAMin"), document.getElementById("faTxt"));
    return [document.getElementById("myRangeFAMin").value, document.getElementById("myRangeFAMax").value];
}

function _searchDate(date, currentsearch){
    return date == currentsearch;
}

function _searchArtistBand(name, currentsearch){
    return _formatString(name).includes(_formatString(currentsearch));
}

function _filterLocations(locations, loc){
    for(const location of locations){
        if(_formatString(location).includes(_formatString(loc))){
            return true;
        }
    }
    return false;
}

function _searchLocations(locations, currentsearch){
    for(const location of locations){
        if(_formatString(location).includes(_formatString(currentsearch))){
            return location;
        }
    }
    return "";
}

function _searchMembers(members, currentsearch){
    for(const name of members){
        if(_formatString(name).includes(_formatString(currentsearch))){
            return name;
        }
    }
    return "";
}

function _filterByNumberOfMember(members){
    let o_Member = document.getElementById("one_members");
    if(o_Member.checked && members.length == 1){
        return true
    }
    let to_Member = document.getElementById("tow_members");
    if(to_Member.checked && members.length == 2){
        return true
    }
    let tr_Member = document.getElementById("tree_members");
    if(tr_Member.checked && members.length == 3){
        return true
    }
    let fo_Member = document.getElementById("four_members");
    if(fo_Member.checked && members.length == 4){
        return true
    }
    let fi_Member = document.getElementById("five_members");
    if(fi_Member.checked && members.length == 5){
        return true
    }
    let s_Member = document.getElementById("six_members");
    if(s_Member.checked && members.length == 6){
        return true
    }
    let m_Member = document.getElementById("more_members");
    if(m_Member.checked && members.length >= 7){
        return true
    }

    if(!o_Member.checked && !to_Member.checked && !tr_Member.checked && !fo_Member.checked && !fi_Member.checked && !s_Member.checked && !m_Member.checked){
        return true
    }

    return false;
}

function _executeFilter(){
    if (sidebarOpen) {
        let artistbandIsDisplay = false;
        let creationdateIsDisplay = false;
        let firstalbumIsDisplay = false;
        let memberIsDisplay = false;
        let locationsIsDisplay = false;
        let searchBar = document.querySelector(".shearch");
        currentsearch = String(searchBar.value);
        document.getElementById("artist_band").querySelector(".artists").innerHTML = "";
        document.getElementById("creation_date").querySelector(".artists").innerHTML = "";
        document.getElementById("firstalbumdate_date").querySelector(".artists").innerHTML = "";
        document.getElementById("members").querySelector(".artists").innerHTML="";
        document.getElementById("locations").querySelector(".artists").innerHTML="";
        for(const [key, artist] of Object.entries(allArtist)){
            if((_creationDateSlider()[0] <= artist.creationDate 
            && _creationDateSlider()[1] >= artist.creationDate) 
            && (_FirstAlbumSlider()[0] <= artist.firstalbum.split("-")[2] 
            && _FirstAlbumSlider()[1] >= artist.firstalbum.split("-")[2]) 
            && _filterByNumberOfMember(artist.members) 
            && _filterLocations(artist.locations, document.getElementById("locselect").value)){

                if(_searchArtistBand(artist.name, currentsearch)){
                    displayArtist(artist.id, artist.image,artist.name, artist.creationDate, "artist_band");
                    artistbandIsDisplay = true;
                }
                if(_searchDate(artist.creationDate, currentsearch)){
                    displayArtist(artist.id, artist.image,artist.name, artist.creationDate, "creation_date");
                    creationdateIsDisplay = true;
                }
                if(_searchDate(artist.firstalbum.split("-")[2], currentsearch)){
                    displayArtist(artist.id, artist.image,artist.name, "First album : " + artist.firstalbum.split("-")[2], "firstalbumdate_date");
                    firstalbumIsDisplay = true;
                }
                var m = _searchMembers(artist.members, currentsearch);
                if(m != "" && currentsearch != ""){
                    displayArtist(artist.id, artist.image,artist.name, m, "members");
                    memberIsDisplay = true;
                }
                var l = _searchLocations(artist.locations, currentsearch);
                if(l != "" && currentsearch != ""){
                    displayArtist(artist.id, artist.image,artist.name, l, "locations");
                    locationsIsDisplay = true;
                }
            }
        }
        let artist_band_sect = document.getElementById("artist_band");
        if(artistbandIsDisplay){
            artist_band_sect.style.display = "block";
        }else { 
            artist_band_sect.style.display = "none";
        }

        let creation_date_sect = document.getElementById("creation_date");
        if(creationdateIsDisplay){
            creation_date_sect.style.display = "block";
        }else { 
            creation_date_sect.style.display = "none";
        }

        let firstalbum_date_sect = document.getElementById("firstalbumdate_date");
        if(firstalbumIsDisplay){
            firstalbum_date_sect.style.display = "block";
        }else { 
            firstalbum_date_sect.style.display = "none";
        }

        let members_sect = document.getElementById("members");
        if(memberIsDisplay){
            members_sect.style.display = "block";
        }else { 
            members_sect.style.display = "none";
        }

        let locations_sect = document.getElementById("locations");
        if(locationsIsDisplay){
            locations_sect.style.display = "block";
        }else { 
            locations_sect.style.display = "none";
        }
    }

}


function _formatString(str){
    return String(String(String(str).toLowerCase()).split(" ").join(""));
}

function displayArtist(id, img, name, creationDate, where){
    let template = `<li class="artist_box" id="${id}"><a href="/artist?id=${id}">
                        <div class="artist"><img src="${img}" alt="">
                            <div>
                                <h3>${name}</h3>
                                <p>${creationDate}</p>
                            </div>
                        </div>
                    </a></li>`;
    document.getElementById(where).querySelector(".artists").innerHTML += template;
}