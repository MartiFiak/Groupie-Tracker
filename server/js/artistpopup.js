let popup = document.querySelector(".popup")

function _aboutUsPopUp(name, member, creationdate, firstalbum){
    var text = `${name} is ${member.length == 1 ? "an artist who appeared":"a group created"} in ${creationdate}. ${member.length == 1 ? "He saw his":`This band is made up of ${member} and saw their`} debut album release the ${firstalbum}.`
    popup.innerHTML = `<div class='aboutUsPP'><p>${text}</p></div>`
    popup.style.display = "flex";
}

popup.addEventListener('click', function(e){   
    if (document.querySelector('.aboutUsPP').contains(e.target)){
    } else{
        popup.style.display = "none";
    }
  });