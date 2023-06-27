document.addEventListener('DOMContentLoaded', function() {
    event.preventDefault();

    console.log("submit button is activated")
    var submitBtn = document.getElementById('submitBtn');
    console.log(submitBtn)
    submitBtn.addEventListener('click', function() {
      var selectedOption = document.querySelector('input[name="option"]:checked').value;
      var minPriceThreshold = document.getElementById('minPrice').value;
      if (selectedOption === 'availability') {
        fetch(`https://smuly-test-ground.ue.r.appspot.com/track/availability`, {
            method: "POST",
            mode:"no-cors",
            headers : { "Content-Type" : "application/json" } ,
            body: JSON.stringify({
                url:"https://www.flipkart.com/samsung-galaxy-f13-nightsky-green-64-gb/p/itmeadfda1bd23fa?pid=MOBGENJWF4KJTPEN&lid=LSTMOBGENJWF4KJTPENS2XJXA&marketplace=FLIPKART&store=tyy%2F4io&srno=b_1_1&otracker=clp_banner_1_14.bannerX3.BANNER_mobile-phones-store_ARV72AV1ALWY&fm=neo%2Fmerchandising&iid=d2edff5a-41be-41c5-825a-89fbf189399b.MOBGENJWF4KJTPEN.SEARCH&ppt=clp&ppn=mobile-phones-store&ssid=m55kbf5z280000001686801244020"
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
                    url:"https://www.flipkart.com/samsung-galaxy-f13-nightsky-green-64-gb/p/itmeadfda1bd23fa?pid=MOBGENJWF4KJTPEN&lid=LSTMOBGENJWF4KJTPENS2XJXA&marketplace=FLIPKART&store=tyy%2F4io&srno=b_1_1&otracker=clp_banner_1_14.bannerX3.BANNER_mobile-phones-store_ARV72AV1ALWY&fm=neo%2Fmerchandising&iid=d2edff5a-41be-41c5-825a-89fbf189399b.MOBGENJWF4KJTPEN.SEARCH&ppt=clp&ppn=mobile-phones-store&ssid=m55kbf5z280000001686801244020",
                    min_threshold: minPriceThreshold.value
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
      
    
    
    