let filter = document.querySelector(".filter");

var filterVisibility = false;

function changesize(self){
    self.style.width = "50vh";
    self.querySelector("img").style.filter = "saturate(1)"
}

function mouseleave(self){
    self.style.width = "15vh";
    self.querySelector("img").style.filter = "saturate(20%)"
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

function _rangeCD(){
    _rangeManager(document.getElementById("myRangeAMax"),document.getElementById("myRangeAMin"), document.getElementById("cdATxt"))
}

function _rangeFA(){
    _rangeManager(document.getElementById("myRangeFAAMax"),document.getElementById("myRangeFAAMin"), document.getElementById("faATxt"))
}

function _showFilterAA(){
    if (filterVisibility){
        filter.style.height = "0px";
        filterVisibility = false;
    } else {
        filter.style.height = "auto";
        filterVisibility = true;
    }
}

const scrollContainer = document.querySelector(".sect2AA");

scrollContainer.addEventListener("wheel", (evt) => {
    evt.preventDefault();
    scrollContainer.scrollLeft += evt.deltaY;
});