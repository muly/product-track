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
          fetch(apiurl + `/store-email`, {                            
            method: 'POST',
            mode: 'no-cors',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email: data.email }),
          })
          .then((req) =>  {
            window.location.href="./popup.html"
          })
          .catch((err) => {
            console.log(err. Message)
          })
        })
        .catch(function(error) {
          console.error('Error retrieving user email:', error);
        });
    });
  });  