let sidebar = document.querySelector(".sidebar");
let sidebarBtn = document.querySelector(".bx-search");
let hright = document.getElementById("hright");
let filter = document.querySelector(".filter");
let form = document.getElementById("sidebarform");

var sidebarPosition = false;
var filterVisibility = false;
console.log(sidebarBtn);

function _openSidebar(){
    hright.style.opacity = "0";
    sidebar.style.width = "250px";
    sidebar.style.padding = "10px";
    console.log("Click", sidebar.style.width);
    sidebarPosition = true;
}

function _closeSidebar(){
    hright.style.opacity = "1";
    sidebar.style.width = "0px";
    sidebar.style.margin = "0px";
    sidebar.style.padding = "0px";
    console.log("Click", sidebar.style.width);
    sidebarPosition = false;
}

function _showFilter(){
    if (filterVisibility){
        filter.style.height = "0px";
        filterVisibility = false;
    } else {
        filter.style.height = "auto";
        filterVisibility = true;
    }

}

function _rangeManager(slider, text){
    text.textContent = "1900 - " + slider.value;
}

function _FirstAlbumSlider(){
    console.log(document.getElementById("myRangeFA").value);
    _rangeManager(document.getElementById("myRangeFA"), document.getElementById("faTxt"));
}

function _creationDateSlider(){
    console.log(document.getElementById("myRange").value);
    _rangeManager(document.getElementById("myRange"), document.getElementById("cdTxt"));
}

function _getData(data){
    console.log(data[0])
}

window.addEventListener('click', function(e){   
    if (sidebar.contains(e.target)){

    } else if (sidebarBtn.contains(e.target)){
        if (!sidebarPosition) {
            _openSidebar();
        } else {
            _closeSidebar();
        } 
    } else {
        _closeSidebar();
    }
});

document.addEventListener("keyup", function(event) {
    if (event.keyCode === 13 && sidebarPosition) {
        form.submit();
    }
});