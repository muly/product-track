var apiurl=`https://smuly-test-ground.ue.r.appspot.com`
document.getElementById('signInBtn').addEventListener('click', function() {
    chrome.identity.getAuthToken({ interactive: true, scopes: ['email'] }, function(token) {
      if (chrome.runtime.lastError) {
        console.error(chrome.runtime.lastError.message);
        return;
      }
      fetch('https://www.googleapis.com/oauth2/v2/userinfo?access_token=' + token)
        .then(function(response) {
          if (!response.ok) {
            throw new Error('Error: ' + response.status);
          }
          return response.json();
        })
        .then(function(data) {
           //need to write error handling
           console.log('User email1:', data.email);
          //chrome.tabs.update({ url: 'popup.html' });
          window.location.href="./popup.html"
          
          fetch(apiurl + `/store-email`, {
            method: 'POST',
            mode: 'no-cors',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email: data.email }),
          })
          .then((req) => {
            alert("succesfull") ;
            window.close();
          })
          .catch((err) => {
            console.log(err. Message)
          })
          console.log('User email2:', data.email);
          
        })
        .catch(function(error) {
          console.error('Error retrieving user email:', error);
        });
    });
  });  