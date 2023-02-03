function changesize(self){
    self.style.width = "50vh";
    self.querySelector("img").style.filter = "saturate(1)"
}

function mouseleave(self){
    self.style.width = "15vh";
    self.querySelector("img").style.filter = "saturate(20%)"
}