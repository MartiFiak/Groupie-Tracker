let sidebar = document.querySelector(".sidebar");
let sidebarBtn = document.querySelector(".bx-search");

var sidebarPosition = false;
console.log(sidebarBtn);

function _openSidebar(){
    sidebar.style.width = "250px";
    sidebar.style.padding = "10px";
    console.log("Click", sidebar.style.width);
    sidebarPosition = true;
}

function _closeSidebar(){
    sidebar.style.width = "0px";
    sidebar.style.margin = "0px";
    sidebar.style.padding = "0px";
    console.log("Click", sidebar.style.width);
    sidebarPosition = false;
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