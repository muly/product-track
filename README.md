# product-track
online product price and availability tracking

## deploying app engine api

TODO: need to add notes here

## deploying chrome extension


### chrome extension: manifest file

Env variables: (saved in `secrets` folder)
**key:**
- for deployment to gcp: 
retrieve the correct key from Chrome Developer Dashboard (https://chrome.google.com) -> open the extension page -> Package tab -> click View public key.

**client_id:**
- for deployment to gcp: 
    1. gcp console: https://console.cloud.google.com/
    1. APIs & Services 
    1. Credentials
    1. OAuth 2.0 Client IDs
    1. copy the `Client ID` under `Chrome client (for chrome store deployment)`

- for deployment to local: 
    1. same steps as above except the last step
    1. copy the `Client ID` under `Chrome client (for local deployment)`


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

### verifying the draft/published code (before it is approved by google)
1. go to chrome web store developer dashboard: https://chrome.google.com/webstore/devconsole/
1. go to product-track: https://chrome.google.com/webstore/devconsole/92b09e82-ea96-4718-9dda-f14771a34b3c/ichhakcbialminoadfkhalilmdhkmifn/edit
1. go to Package page
1. under the Draft or the Published section, click the .crx file link under CRX file column
1. this should download and attempt to install.
1. Note: Draft might fail to download, so confirm by clicking Download suspicious file
1. now you should see the actual error (if any) on the top of the page.
1. if successful you will see a message Apps. extensions, and user scripts cannot be added from this website

### deploying chrome extension to local chrome store
temp notes required to update the script to deploy local chrome store
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

### Download error: Invalid manifest
scenario: after the chrome extension latest version submitted in chrome web store is approved by google, when trying to install (i.e. using `Add to Chrome` button) the extension, this error is seen.

error: 
    Download error: Invalid manifest

root cause: 
    this error message is very generic. the main error can be found using the steps indicated in section "verifying the draft/published code (before it is approved by google)" above.


## privacy policy

dev url: https://smuly-test-ground.ue.r.appspot.com/privacy-policy.html

