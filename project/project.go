package project

import (
		"log"
		"fmt"
		"strings"
		"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
		resourcemanagerv3 "github.com/pulumi/pulumi-google-native/sdk/go/google/cloudresourcemanager/v3"
		cloudbilling "google.golang.org/api/cloudbilling/v1"
		serviceusage "google.golang.org/api/serviceusage/v1"
		"context"
		"time"
		"google.golang.org/api/googleapi"
		"net/http"
)

type Project struct {
}

func NewProject(ctx *pulumi.Context, stack string, pid string, pname string, parentFolder pulumi.StringOutput, billingAccount string) error {
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

	var p Project


	projectid := pulumi.StringOutput(project.ProjectId).Apply()

	err = p.enableAPIs(ctx, stack, projectid)
	if err != nil {
		log.Println(err)
	}

	err = p.updateProjectBillingAccount(ctx, stack, projectid, billingAccount)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(project)

	return nil
}

func (p *Project) enableAPIs(ctx *pulumi.Context, stack, pid string) error {

	log.Println("[DEBUG]: Enabling basic API's on project %q", pid)
	project := fmt.Sprintf("projects/%s", pid)

	err := p.enableAPI(ctx, stack, project, "serviceusage.googleapis.com")
	if err != nil{
		log.Println(err)
		return err
	}

	err = p.enableAPI(ctx, stack, project, "cloudbilling.googleapis.com")
	if err != nil{
		log.Println(err)
		return err
	}

	
	return nil
}

func (p *Project) enableAPI(ctx *pulumi.Context, stack, project string, api string) error {
	const retries int = 3

	context := context.TODO()
	serviceUsageClient, err := serviceusage.NewService(context)
	if err != nil{
		log.Println(err)
		return err
	}

	apiName := fmt.Sprintf("%s/services/%s", project, api)
	req := serviceUsageClient.Services.Enable(apiName, &serviceusage.EnableServiceRequest{})
	var retry int
	for {
		retry++
		time.Sleep(time.Second)
		_, err := req.Do()
		if err != nil {
			ae, ok := err.(*googleapi.Error)
			if ok && ae.Code == http.StatusForbidden && retry <= retries {
				log.Printf("[DEBUG]")
			}
			return err
		}
		return nil
	}


	return nil
}

func (p *Project) updateProjectBillingAccount(ctx *pulumi.Context, stack string, pid string, bid string) error {
	log.Printf("[DEBUG]: Attaching project %q to billing account %q", pid, bid)

	project := fmt.Sprintf("projects/%s", pid)
	billingAccount := fmt.Sprintf("billingAccounts/%S", strings.TrimSuffix(bid, "\n"))

	context := context.TODO()

	cloudBillingClient, err := cloudbilling.NewService(context)
	if err != nil {
		log.Println(err)
		return err
	}
	
	info, err := cloudBillingClient.Projects.GetBillingInfo(project).Do()
	if err != nil {
		log.Println(err)
		return err
	}

	info.BillingAccountName = billingAccount
	info.BillingEnabled = true

	_, err = cloudBillingClient.Projects.UpdateBillingInfo(project, info).Do()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}