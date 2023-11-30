# product-track
online product price and availability tracking

## deploying app engine api

TODO: need to add notes here

## deploying chrome extension

### first time deploying chrome extension

TODO: need to add the detailed steps associated with the chrome webstore review and publish process. 

### re-deploying chrome extension to chrome store
step 1. generate the zip file:
    make prepare_chrome_ext_for_dev_deployment (this command might change, see makefile)
step 2: use the generated chrome-exten.zip file to update to chrome webstore:
    -> https://chrome.google.com/webstore/devconsole/
    -> product-track 
    -> build 
    -> package 
    -> upload new package 
    -> browse and pick the zip file


### deploying chrome extension to local chrome store
temp notes required to update the script to deploy local chromestore
    - for local chrome deployment: 
step 1: gcp console -> APIs & Services -> credentials -> OAuth 2.0 Client IDs -> Chrome client (for chrome store deployment) -> copy client id
step 2: update oauth2->client_id value in chrome-exten/manifest.json file.


## known errors:


### OAuth2
scenario: error recorded in chrome extension's console when trying to Google login.

error: 
    signin.js:5 OAuth2 request failed: Service responded with error: 'bad client id: ******************************.apps.googleusercontent.com'

root cause: unknown.

workaround: generate new client id in gcp console and use the new client id in manifest.json file

