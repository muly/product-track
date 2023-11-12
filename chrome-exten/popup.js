//var apiurl=`https://smuly-test-ground.ue.r.appspot.com`
var apiurl=`http:localhost:8006`
var activeUrl
let emailid 
chrome.identity.getAuthToken({ interactive: true, scopes: ['email'] }, function(token) {
  if (chrome.runtime.lastError) {
    console.error(chrome.runtime.lastError.message);
    return;
  }
  fetch('https://www.googleapis.com/oauth2/v2/userinfo?access_token=' +token)
      .then(function(response) {
        if (!response.ok) {
          throw new Error('Error: ' + response.status);
        }
        return response.json();
  })
  .then(function(data) {
      emailid=data.email
  })
  })
document.addEventListener('DOMContentLoaded', function() {
  function fetchActiveTabURL(callback) {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      if (tabs && tabs.length > 0) {
        activeUrl= tabs[0].url;
        console.log(tabs[0].url)
        callback(activeUrl);
      }
    });
  }
  var activeTabURL;
  fetchActiveTabURL(function(activeUrl) {
    activeTabURL = activeUrl;
    const availabilityFields = document.getElementById('zipcode');
    const priceFields = document.getElementById('priceThreshold');
    const availabilityOption = document.querySelector('input[value="availability"]');
    const priceOption = document.querySelector('input[value="price"]');
    var submitBtn = document.getElementById('submitBtn');
    availabilityOption.addEventListener('change', function() {
    if (availabilityOption.checked) {
        availabilityFields.style.display = 'block';
        priceFields.style.display = 'none';
      }
    });
    priceOption.addEventListener('change', function() {
    if (priceOption.checked) {
          availabilityFields.style.display = 'none';
          priceFields.style.display = 'block';
      }
    });
    submitBtn.addEventListener('click', function(event) {  
      event.preventDefault();
      var selectedOption = document.querySelector('input[name="option"]:checked').value;
      var minPriceThreshold = document.getElementById('minPrice').value;
      var zipCode = document.getElementById('zipcodeValue').value;
      var payload = {
        url: activeTabURL,
        emailid: emailid
      };
      if (selectedOption === 'availability') {
          payload.zipcode = parseInt(zipCode);
      } else if (selectedOption === 'price') {
          payload.min_threshold = parseFloat(minPriceThreshold);
      }
      fetch(apiurl + '/track/' + selectedOption, {
          method: "POST",
          //mode: "no-cors",
          //headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload)
      })
      .then(function(response) {
        if (!response.ok) {
          alert("our extension doesnot support the given url")
        }else{
          alert("successful");
          window.close();
        }
      })
      .catch((err) => {
         console.error('There has been a problem with your fetch operation:', err);
       });
    });
  });
});
      
    
    
    