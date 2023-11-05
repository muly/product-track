// var apiurl=`https://smuly-test-ground.ue.r.appspot.com`
var apiurl=`http://localhost:8006`
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
        var url = tabs[0].url;
        callback(url);
        }
      });
    }
    var activeTabURL;
    fetchActiveTabURL(function(url) {
    activeTabURL = url;
    var submitBtn = document.getElementById('submitBtn');
    submitBtn.addEventListener('click', function(event) {  
      event.preventDefault();
      var selectedOption = document.querySelector('input[name="option"]:checked').value;
      var minPriceThreshold = document.getElementById('minPrice').value;
      if (selectedOption === 'availability') {
        fetch(apiurl+`/track/availability`, {
            method: "POST",
            mode:"no-cors",
            headers : { "Content-Type" : "application/json" } ,
            body: JSON.stringify({
              url:activeTabURL,
              emailid:emailid
            })
        })
        .then((response) => {
          if (response.ok) {
              alert("Successful"+ response.status);
          } else {
              alert("Error - Status Code: " + response.status);
          }
          window.close();
      })
            .catch((err) => {
                console.log(err. Message)
            })
      }
      else if (selectedOption === 'price') {
          fetch(apiurl+`/track/price`, {
              method: "POST",
              mode:"no-cors",
              headers : { "Content-Type" : "application/json" } ,
              body: JSON.stringify({
                  url:activeTabURL,
                  min_threshold:parseFloat(minPriceThreshold),
                  emailid:emailid
              })
          })
          .then((response) => {
            if (response.ok) {
                alert("Successful- Status Code:"+ response.status);
            } else {
                alert("Error - Status Code: " + response.status);
            }
            window.close();
        })
              .catch((err) => {
                console.log(err. Message)
              })
        }  
    });
  });
});
      
    
    
    