let popup = document.querySelector(".popup")

function _aboutUsPopUp(name, member, creationdate, firstalbum){
    popup.style.height = "100vh";
    var text = `${name} is ${member.length == 1 ? "an artist who appeared":"a group created"} in ${creationdate}. ${member.length == 1 ? "He saw his":`This band is made up of ${member} and saw their`} debut album release the ${firstalbum}.`
    popup.innerHTML = `<div class='aboutUsPP'><p>${text}</p></div>`
    popup.style.display = "flex";
}

popup.addEventListener('click', function(e){   
    if (document.querySelector('.aboutUsPP') != null && document.querySelector('.aboutUsPP').contains(e.target)){
    } else{
        popup.style.height = "0px";
        popup.style.display = "none";
        popup.innerHTML = "";
    }
    if (document.querySelector('.payementCard') != null && document.querySelector('.payementCard').contains(e.target)){
    } else{
        popup.style.height = "0px";
        popup.style.display = "none";
        popup.innerHTML = "";
    }
  });

function _payementPopUp(){
    popup.style.height = "100vh";
    popup.innerHTML = `<div class="payementCard">
    <div class="topSect">
      <div class="creditCard">
        <div class="top">
        
        </div>
        <div class="mid">
          <p class="cardNumber" id="cardNumber">
          ####  ####  ####  ####
          </p>
        </div>
        <div class="bot">
        
        </div>
        
      </div>
    </div>
    <form>
      <label>Card Number</label>
      <input type="text" oninput="_changeText(this)">
      <label>Card Holder</label>
      <input type="text">
      <div>
      <div class="grp">
        <label>Expiration Date</label>
        <input type="month">
      </div>
      <div class="grp">
        <label>CVV</label>
        <input type="text">
      </div>
      </div>
      <input type="submit">
    </form>
  </div>`
    popup.style.display = "flex";

}