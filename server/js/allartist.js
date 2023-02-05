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