package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

var secretManagerClient *secretmanager.Client
var secretVersion *secretmanagerpb.AccessSecretVersionResponse

func main() {
	log.Println("main function started")
	initFirestore(context.Background())
	initEmailClient()

	testSecret()

	handleRequest()
}

func testSecret() {
	projectID := "149500152182" // project id in number format, not alpha string

	// Create the client.
	ctx := context.Background()
	var err error
	secretManagerClient, err = secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf("failed to setup client: %v", err)
		return
	}
	defer secretManagerClient.Close() 
	
	req := &secretmanagerpb.CreateSecretRequest{
		Parent:"smuly-test-ground" ,
		SecretId: projectID,
		Secret: &secretmanagerpb.Secret{
				Replication: &secretmanagerpb.Replication{
						Replication: &secretmanagerpb.Replication_Automatic_{
								Automatic: &secretmanagerpb.Replication_Automatic{},
						},
				},
		},
	}
	result, err := secretManagerClient.CreateSecret(ctx, req)
        if err != nil {
                return 
        }
    log.Printf( "Created secret: %s\n", result.Name)
	secretID := "TEST_SECRET"
	secretVersion, err := secretManagerClient.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/1", projectID, secretID),
	})
	if err != nil {
		log.Printf("failed to access secret version: %v", err)
		// ERROR: rpc error: code = PermissionDenied desc = Permission 'secretmanager.versions.access' denied for resource 'projects/149500152182/secrets/TEST_SECRET/versions/1' (or it may not exist).
		// Note: secretmanager.versions.access permission is part of "Secret Manager Secret Accessor" role. 
		// Fix: add "Secret Manager Secret Accessor" role to the Principal which is used as service account for the app engine. in this case it is the App Engine default service account (smuly-test-ground@appspot.gserviceaccount.com). 
		// steps: go to IAM -> IAM (https://console.cloud.google.com/iam-admin/iam?project=smuly-test-ground)  
		// -> edit the "App Engine default service account" principal
		// -> Add Another Role -> search and add "Secret Manager Secret Accessor" role -> Save
		return
	}
	log.Printf("secret %s is %s, %s", secretID, string(secretVersion.Payload.Data), secretVersion.Payload.String)
}

// function for processing a url according the url provided
func process(rawURL string) (product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return product{}, err
	}
	switch u.Hostname() {
	case "scrapeme.live":
		return scrapeme(rawURL)
	case "www.flipkart.com":
		return flipkart(rawURL)
	case "www.amazon.in":
		return amazon(rawURL)
	default:
		log.Printf("%s is not supported\n", u.Hostname())
		return product{}, err
	}
}
