let artist_boxs = document.getElementsByClassName("artist_box");

class Artist{
    constructor(id,image,name,creationdate,firstalbum, members){
        this.id;
        this.image;
        this.name;
        this.creationdate;
        this.firstalbum;
        this.members;
    }
}

let allArtist = {}
let currentArtist = []

function _addArtist(id, image, name, creationDate, firstAlbum, members){
    allArtist[id] = {id:id,image:image,name:name,creationDate:creationDate,firstalbum:firstAlbum, members:members};
    displayArtist(id, image,name, creationDate, "artist_band");
}

document.addEventListener("keyup", function(event) {
    _executeFilter()
});

function _creationDateSlider(){
    _rangeManager(document.getElementById("myRange"), document.getElementById("cdTxt"));
    return document.getElementById("myRange").value;
}

function _FirstAlbumSlider(){
    _rangeManager(document.getElementById("myRangeFA"), document.getElementById("faTxt"));
    return document.getElementById("myRangeFA").value;
}

function _searchDate(date, currentsearch){
    return date == currentsearch;
}

function _searchArtistBand(name, currentsearch){
    return _formatString(name).includes(_formatString(currentsearch));
}

function _searchMembers(members, currentsearch){
    for(const name of members){
        if(_formatString(name).includes(_formatString(currentsearch))){
            return name;
        }
    }
    return "";
}

function _executeFilter(){
    if (sidebarOpen) {
        let artistbandIsDisplay = false;
        let creationdateIsDisplay = false;
        let firstalbumIsDisplay = false;
        let memberIsDisplay = false;
        let searchBar = document.querySelector(".shearch");
        currentsearch = String(searchBar.value);
        document.getElementById("artist_band").querySelector(".artists").innerHTML = "";
        document.getElementById("creation_date").querySelector(".artists").innerHTML = "";
        document.getElementById("firstalbumdate_date").querySelector(".artists").innerHTML = "";
        document.getElementById("members").querySelector(".artists").innerHTML="";
        for(const [key, artist] of Object.entries(allArtist)){
            if(_creationDateSlider() >= artist.creationDate && _FirstAlbumSlider() >= artist.firstalbum.split("-")[2]){
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