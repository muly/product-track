document.addEventListener('DOMContentLoaded', function() {
    var availabilityCheckbox = document.getElementById('availabilityCheckbox');
    var priceCheckbox = document.getElementById('priceCheckbox');
    var minPriceContainer = document.getElementById('minPriceContainer');
    var submitButton = document.getElementById('submitButton');
  
    priceCheckbox.addEventListener('change', function() {
      if (priceCheckbox.checked) {
        minPriceContainer.style.display = 'block';
      } else {
        minPriceContainer.style.display = 'none';
      }
    });
  
    submitButton.addEventListener('click', function() {
      var trackAvailability = availabilityCheckbox.checked;
      var trackPrice = priceCheckbox.checked;
      var minPrice = document.getElementById('minPrice').value;
  
      // Perform API calls based on selected options
      if (trackAvailability) {
        trackAvailabilityAPI();
      }
      if (trackPrice) {
        trackPriceAPI(minPrice);
      }
  
    });
  
    function trackAvailabilityAPI() {
        fetch(`https://smuly-test-ground.ue.r.appspot.com/track/availability`, {
            method: "POST",
            mode:"no-cors",
            headers : { "Content-Type" : "application/json" } ,
            body: JSON.stringify({
                url:"https://www.flipkart.com/samsung-galaxy-f13-nightsky-green-64-gb/p/itmeadfda1bd23fa?pid=MOBGENJWF4KJTPEN&lid=LSTMOBGENJWF4KJTPENS2XJXA&marketplace=FLIPKART&store=tyy%2F4io&srno=b_1_1&otracker=clp_banner_1_14.bannerX3.BANNER_mobile-phones-store_ARV72AV1ALWY&fm=neo%2Fmerchandising&iid=d2edff5a-41be-41c5-825a-89fbf189399b.MOBGENJWF4KJTPEN.SEARCH&ppt=clp&ppn=mobile-phones-store&ssid=m55kbf5z280000001686801244020",
            })
        })
            .then((req) => {
                alert('Successful') ;
                })
            .catch((err) => {
                console.log(err. Message)
            })
                        
        
      console.log('Calling availability tracking API...');
    }
  
    function trackPriceAPI(minPrice) {
        fetch(`https://smuly-test-ground.ue.r.appspot.com/track/price`, {
            method: "POST",
            mode:"no-cors",
            headers : { "Content-Type" : "application/json" } ,
            body: JSON.stringify({
                url:"https://www.flipkart.com/samsung-galaxy-f13-nightsky-green-64-gb/p/itmeadfda1bd23fa?pid=MOBGENJWF4KJTPEN&lid=LSTMOBGENJWF4KJTPENS2XJXA&marketplace=FLIPKART&store=tyy%2F4io&srno=b_1_1&otracker=clp_banner_1_14.bannerX3.BANNER_mobile-phones-store_ARV72AV1ALWY&fm=neo%2Fmerchandising&iid=d2edff5a-41be-41c5-825a-89fbf189399b.MOBGENJWF4KJTPEN.SEARCH&ppt=clp&ppn=mobile-phones-store&ssid=m55kbf5z280000001686801244020",
            })
        })
            .then((req) => {
                alert('Successful') ;
                })
            .catch((err) => {
                console.log(err. Message)
        })
      console.log('Calling price tracking API with min price threshold:', minPrice);
    }
  });
  