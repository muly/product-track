document.addEventListener('DOMContentLoaded', function() {
    event.preventDefault();

    chrome.tabs.query({currentWindow:true, active : true},function(tabs){
       var Url=(tabs[0].url)
       console.log(Url)
    });
    
    // async function getCurrentTab() {
    //   let queryOptions = { active: true, currentWindow: true };
     
    //    let [tab] = await chrome.tabs.query(queryOptions);
    //    localStorage.setItem('tabname' , tab);
    //     return tab;
    //    }
     
    //   getCurrentTab()
    //   .then((data) => { console.log('newdata',data)})
    //   .then(() => { console.log('error')});
      
   

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
              //url:objectname.url
                
               
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
                    //url:objectname.url,
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
      
    
    
    