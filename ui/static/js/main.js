var longurl = document.querySelector('.longurl');
var shortname = document.querySelector('.shortname');
var shorturl = document.querySelector('.shorturl');
var button = document.querySelector('.button');

button.onclick = buttonOnClick
button.disabled = true;

longurl.addEventListener('input', changeBackground);
shortname.addEventListener('input', changeBackground);

function changeBackground() {
  if (longurl.value !== '' && shortname.value !== '') {
    button.style.background = 'green';
    button.disabled = false;
  } else {
    button.style.background = 'grey';
    button.disabled = true; 
  }
}

async function buttonOnClick() {
    let UrlMapping = {
        longurl: longurl.value,
        shortname: shortname.value
      };
      
    let response = await fetch('/create', {
    method: 'POST',
    body: JSON.stringify(UrlMapping)
    });

    if (response.status == 400) {
        let result = await response.json();
        alert(result.StatusMessage)
    }
    else {
        let result = await response.json();
        shorturl.value = result.StatusMessage
    }
    
}