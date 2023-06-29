var apiurl=`https://smuly-test-ground.ue.r.appspot.com`
document.addEventListener('DOMContentLoaded', function() {
    event.preventDefault()
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
    submitBtn.addEventListener('click', function() {  
      var selectedOption = document.querySelector('input[name="option"]:checked').value;
      var minPriceThreshold = document.getElementById('minPrice').value;
      if (selectedOption === 'availability') {
        fetch(apiurl+`/track/availability`, {
            method: "POST",
            mode:"no-cors",
            headers : { "Content-Type" : "application/json" } ,
            body: JSON.stringify({
              url:activeTabURL
            })
        })
            .then((req) => {
                alert('Successful') ;
            })
            .catch((err) => {
                console.log(err. Message)
            })
        console.log('Tracking availability...');
      } else if (selectedOption === 'price') {
          fetch(apiurl+`/track/price`, {
              method: "POST",
              mode:"no-cors",
              headers : { "Content-Type" : "application/json" } ,
              body: JSON.stringify({
                  url:activeTabURL,
                  min_threshold:parseFloat(minPriceThreshold)
              })
          })
              .then((req) => {
                  alert('Successful') ;
              })
              .catch((err) => {
                console.log(err. Message)
              })
         console.log('Tracking price with min threshold:', minPriceThreshold);
      }
      window.close();
    });
  });
});
      
    
    
    