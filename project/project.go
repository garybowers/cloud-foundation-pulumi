package project

import (
		"log"
		"fmt"
		"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
		resourcemanagerv3 "github.com/pulumi/pulumi-google-native/sdk/go/google/cloudresourcemanager/v3"
)

func ProjectFactory(ctx *pulumi.Context, stack string, pid string, pname string, parentFolder pulumi.StringOutput, billingAccount string) error {
	log.Printf("[DEBUG]: Creating new project %q", pid)

	project, err := resourcemanagerv3.NewProject(ctx, stack, &resourcemanagerv3.ProjectArgs{
		Parent: parentFolder,
		ProjectId: pulumi.String(pid),
		DisplayName: pulumi.String(pname),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	err = updateProjectBillingAccount(ctx, stack, pid)

	fmt.Println(project)

	return nil
}

func updateProjectBillingAccount(ctx *pulumi.Context, stack string, pid string) error {
	log.Printf("[DEBUG]: Attaching project %q to billing account %q", pid, "nul")

	return nil
}