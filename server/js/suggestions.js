let artist_boxs = document.getElementsByClassName("artist_box");

class Artist{
    constructor(id,image,name,creationdate){
        this.id;
        this.image;
        this.name;
        this.creationdate;
    }
}

let allArtist = {}
let currentArtist = []

function _addArtist(id, image, name, creationDate){
    allArtist[id] = {id:id,image:image,name:name,creationDate:creationDate};
    displayArtist(id, image,name, creationDate, "artist_band");
}

document.addEventListener("keyup", function(event) {
    if (sidebarOpen) {
        let artistbandIsDisplay = false;
        let searchBar = document.querySelector(".shearch");
        currentsearch = String(searchBar.value);
        /*for(let i=0; i < artist_boxs.length; i++){
            artistId = artist_boxs[i].id;
            if(_formatString(allArtist[artistId].name).includes(_formatString(currentsearch))){
                artist_boxs[i].style.display = "list-item";
                artistbandIsDisplay = true;
            }else {
                artist_boxs[i].style.display = "none";
            }
            
        }
        let artist_band_sect = document.getElementById("artist_band");
        if(artistbandIsDisplay){
            artist_band_sect.style.display = "block";
        }else { 
            artist_band_sect.style.display = "none";
        }*/
        document.getElementById("artist_band").querySelector(".artists").innerHTML = "";
        for(const [key, artist] of Object.entries(allArtist)){
            if(_formatString(artist.name).includes(_formatString(currentsearch))){
                console.log(key, artist);
                displayArtist(artist.id, artist.image,artist.name, artist.creationDate, "artist_band");
                artistbandIsDisplay = true;
            }
        }
        let artist_band_sect = document.getElementById("artist_band");
        if(artistbandIsDisplay){
            artist_band_sect.style.display = "block";
        }else { 
            artist_band_sect.style.display = "none";
        }
    }

});

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