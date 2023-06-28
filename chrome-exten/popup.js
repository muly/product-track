document.addEventListener('DOMContentLoaded', function() {
    event.preventDefault();

    // chrome.tabs.query({currentWindow:true, active : true},function(tabs){
    //    var Url=(tabs[0].url)
    //    console.log(Url)
    // });
    
    // async function getCurrentTab() {
    //   let queryOptions = { active: true, currentWindow: true };
     
    //    let [tab] = await chrome.tabs.query(queryOptions);
    //    localStorage.setItem('tabname' , tab);
    //     return tab;
    //    }
     
    //   getCurrentTab()
    //   .then((data) => { console.log('newdata',data)})
    //   .then(() => { console.log('error')});
    // function fetchActiveTabURL() {
    //   chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
    //     if (tabs && tabs.length > 0) {
    //       var url = tabs[0].url;
    //       // or perform any other action with the URL
    //     }
    //   });
    // }
    
    // // Call the function to fetch the URL
    //   var Url=fetchActiveTabURL();  
    //   console.log(Url)
    // Ensure the extension has the "tabs" permission in the manifest.json file.

// Function to fetch the URL of the current active tab and pass it to a callback function
    function fetchActiveTabURL(callback) {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
        if (tabs && tabs.length > 0) {
        var url = tabs[0].url;
        callback(url);
        }
      });
    }

// Example usage
    var activeTabURL;

    fetchActiveTabURL(function(url) {
    activeTabURL = url;
  // console.log(activeTabURL); // or perform any other action with the URL
   

   

    var submitBtn = document.getElementById('submitBtn');
    
    
    submitBtn.addEventListener('click', function() {
      var selectedOption = document.querySelector('input[name="option"]:checked').value;
      var minPriceThreshold = document.getElementById('minPrice').value;
      if (selectedOption === 'availability') {
        fetch(`https://smuly-test-ground.ue.r.appspot.com/track/availability`, {
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
            
            
            fetch(`https://smuly-test-ground.ue.r.appspot.com/track/price`, {
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
        });
      });
      });
      
    
    
    