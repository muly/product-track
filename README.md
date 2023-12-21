# product-track
online product price and availability tracking

## deploying app engine api

TODO: need to add notes here

## deploying chrome extension

### first time deploying chrome extension

steps: (TODO: to be reviewed )
1. zip content of "chrome-exten" folder
2. go to chrome webstore : https://chrome.google.com/webstore/category/extensions 
3. click on setting gear icon on top right , adjacent to the profile
4. click on "Developer Dashboard"
5. go to the product track extension: https://chrome.google.com/webstore/devconsole/92b09e82-ea96-4718-9dda-f14771a34b3c/ichhakcbialminoadfkhalilmdhkmifn/edit
6. go to Package page: https://chrome.google.com/webstore/devconsole/92b09e82-ea96-4718-9dda-f14771a34b3c/ichhakcbialminoadfkhalilmdhkmifn/edit/package
7. click on "Update New Package"
8. TODO: need to add more
TODO: need to add the detailed steps associated with the chrome webstore review and publish process. 

### re-deploying chrome extension to chrome store
step 1. generate the zip file:

`make prepare_chrome_ext_for_dev_deployment` (this command might change, see makefile)

step 2: use the generated chrome-exten.zip file to update to chrome webstore:
1. -> https://chrome.google.com/webstore/devconsole/
1. -> product-track 
1. -> build 
1. -> package 
1. -> upload new package 
1. -> browse and pick the zip file


### deploying chrome extension to local chrome store
temp notes required to update the script to deploy local chromestore
    - for local chrome deployment: 

step 1: gcp console 
1. -> APIs & Services 
1. -> credentials 
1. -> OAuth 2.0 Client IDs 
1. -> Chrome client (for chrome store deployment) 
1. -> copy client id

step 2: 

1. update oauth2
1. -> client_id value in chrome-exten/manifest.json file.


### webstore link

https://chromewebstore.google.com/detail/ichhakcbialminoadfkhalilmdhkmifn


## known errors:


### OAuth2
scenario: error recorded in chrome extension's console when trying to Google login.

error: 
    signin.js:5 OAuth2 request failed: Service responded with error: 'bad client id: ******************************.apps.googleusercontent.com'

root cause: unknown.

workaround: generate new client id in gcp console and use the new client id in manifest.json file

## privacy policy

dev url: https://smuly-test-ground.ue.r.appspot.com/privacy-policy.html

