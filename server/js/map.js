let coords = document.getElementsByClassName("coord");
console.log(coords);


let usableCoord = [];
for (let i = 0; i < coords.length; i++) {

  coords[i].onmouseover = function(){
    console.log("hover");
  };
  coords[i].addEventListener("mouseover", function( event ) {
    console.log("hover")
  
  },false);
  coord = {lat: coords[i].textContent.split("|")[0], lng: coords[i].textContent.split("|")[1]}
  usableCoord.push(coord);
}


let pos = {lat: Number(usableCoord[0]["lat"]),lng: Number(usableCoord[0]["lng"]) };


function changeMapFocus(lat,lng){
  pos = {lat: lat, lng: lng}
  initMap()
}


function initMap() {
    const map = new google.maps.Map(document.getElementById("map"), {
      zoom: 10,
      center: pos,
    });
    const marker = new google.maps.Marker({
      position: pos,
      map: map,
    });
  }
  
  window.initMap = initMap;