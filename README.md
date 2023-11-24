# product-track
online product price and availability tracking

## known errors:


### OAuth2
scenario: error recorded in chrome extension's console when trying to Google login.

error: 
    signin.js:5 OAuth2 request failed: Service responded with error: 'bad client id: ******************************.apps.googleusercontent.com'

root cause: unknown.

workaround: generate new client id in gcp console and use the new client id in manifest.json file
