package main

import (
	"context"
	"fmt"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func main() {
	cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID: "f8cdef31-a31e-4b4a-93e4-5f571e91255a",
		ClientID: "8206ae2c-8e14-493e-88f7-19ebf7f2c3ca",
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(message.Message)
			return nil
		},
	})

	if err != nil {
		fmt.Printf("Error creating credentials: %v\n", err)
	}

	auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Files.Read"})
	if err != nil {
		fmt.Printf("Error authentication provider: %v\n", err)
		return
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		fmt.Printf("Error creating adapter: %v\n", err)
		return
	}
	client := msgraphsdk.NewGraphServiceClient(adapter)

	result, err := client.Users().Get(nil)
	if err != nil {
		fmt.Printf("Error getting users: %v\n", err)
		return err
	}

	result, err = client.Me().Drive().Get()
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
	}
	fmt.Printf("Found Drive : %v\n", *result.GetId())

	// Use PageIterator to iterate through all users
	/*	pageIterator, err := msgraphsdk.NewPageIterator(result, adapter, models.CreateUserCollectionResponseFromDiscriminatorValue)

		err = pageIterator.Iterate(func(pageItem interface{}) bool {
			user := pageItem.(models.Userable)
			fmt.Printf("%s\n", *user.GetDisplayName())
			// Return true to continue the iteration
			return true
		})
	*/
}
