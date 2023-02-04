let sidebar = document.querySelector(".sidebar");
let sidebarBtn = document.querySelector(".bx-search");
let hright = document.getElementById("hright");
let filter = document.querySelector(".filter");
let form = document.getElementById("sidebarform");

var sidebarOpen = false;
var filterVisibility = false;

function _openSidebar(){
    hright.style.opacity = "0";
    sidebar.style.width = "250px";
    sidebar.style.padding = "10px";
    sidebarOpen = true;
}

function _closeSidebar(){
    hright.style.opacity = "1";
    sidebar.style.width = "0px";
    sidebar.style.margin = "0px";
    sidebar.style.padding = "0px";
    sidebarOpen = false;
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

function _rangeManager(slider, slider2, text){
    text.textContent = slider2.value + " - " + slider.value;
    if ('min' in slider) {
        slider.min = slider2.value;
    } 
    else {
        slider.setAttribute ("min", slider2.value);
    }
    if ('max' in slider) {
        slider2.max = slider.value;
    } 
    else {
        slider2.setAttribute ("max", slider.value);
    }
}

function _getData(data){
    console.log(data[0])
}

window.addEventListener('click', function(e){   
    if (sidebar.contains(e.target)){

    } else if (sidebarBtn.contains(e.target)){
        if (!sidebarOpen) {
            _openSidebar();
        } else {
            _closeSidebar();
        } 
    } else {
        _closeSidebar();
    }
});