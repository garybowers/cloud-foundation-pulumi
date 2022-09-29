package main

import (
	"cloud-foundation-pulumi/project"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	resourcemanagerv3 "github.com/pulumi/pulumi-google-native/sdk/go/google/cloudresourcemanager/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
)

var (
	// replace these values
	org_id          = ""
	billing_account = ""
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {

	postfix, _ := randomHex(4)

	pulumi.Run(func(ctx *pulumi.Context) error {
		//conf := config.New(ctx, "google-native")
		stack := ctx.Stack()
		fmt.Println(stack)

		// Create a top level folder with the stack name
		folder, err := resourcemanagerv3.NewFolder(ctx, stack, &resourcemanagerv3.FolderArgs{
			Parent:      pulumi.String("organizations/" + org_id),
			DisplayName: pulumi.String(stack),
		})
		if err != nil {
			log.Println(err)
		}

		err = project.ProjectFactory(ctx, stack, "gke-test-1-"+postfix, "gke-test-1-"+postfix, folder.Name, billing_account)
		fmt.Println(err)

		return nil
	})
}
