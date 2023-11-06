var apiurl=`https://smuly-test-ground.ue.r.appspot.com`
//var apiurl=`http:localhost:8006`
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
 
  document.addEventListener('DOMContentLoaded', function() {
    // Function to fetch active tab URL
    function fetchActiveTabURL(callback) {
        chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
            if (tabs && tabs.length > 0) {
                var url = tabs[0].url;
                callback(url);
            }
        });
    }

    // DOM elements
    const availabilityFields = document.getElementById('availabilityFields');
    const priceFields = document.getElementById('priceFields');
    const availabilityOption = document.querySelector('input[value="availability"]');
    const priceOption = document.querySelector('input[value="price"]');
    const submitBtn = document.getElementById('submitBtn');

    // Fetch active tab URL
    var activeTabURL;
    fetchActiveTabURL(function(url) {
        activeTabURL = url;
    });

    // Event listener for radio button changes
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
        var zipCode = document.getElementById('zipCode').value;
            if (selectedOption === 'availability') {
            fetch(apiurl+'/track/availability', {
                method: "POST",
                mode: "no-cors",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    url: activeTabURL,
                    emailid: emailid,
                    zipcode: zipCode
                })
            })
            .then((response) => {
                alert("Successful");
                window.close();
            })
            .catch((err) => {
                console.log(err.message);
            });
        } else if (selectedOption === 'price') {
            fetch(apiurl+'/track/price', {
                method: "POST",
                mode: "no-cors",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    url: activeTabURL,
                    min_threshold: parseFloat(minPriceThreshold),
                    emailid: emailid
                })
            })
            .then((response) => {
                alert("Successful");
                window.close();
            })
            .catch((err) => {
                console.log(err.message);
            });
        }
      });
    });
});

      
    
    
    