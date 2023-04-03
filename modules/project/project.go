package project

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectArgs struct {
	ProjectId         string
	Name              string
	AutoCreateNetwork bool
	BillingAccount    string
	FolderId          string
	Labels            map[string]string
	OrgId             string
	Services          []string
}

type Project struct {
	Args ProjectArgs
	pulumi.ResourceState
	Name string
}

func (project *Project) Create(ctx *pulumi.Context) (gcpProject *organizations.Project, err error) {
	args := &organizations.ProjectArgs{}
	args.Name = pulumi.String(project.Args.Name)

	if project.Args.ProjectId == "" {
		postfix, err := random.NewRandomString(ctx, fmt.Sprintf("postfix-%s", args.Name), &random.RandomStringArgs{
			Length:  pulumi.Int(6),
			Special: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
			Lower:   pulumi.Bool(false),
		})
		if err != nil {
			log.Println(err)
		}
		args.ProjectId = pulumi.Sprintf("%s-%s", project.Args.Name, postfix.Result)
	} else {
		args.ProjectId = pulumi.String(project.Args.ProjectId)
	}

	args.BillingAccount = pulumi.String(project.Args.BillingAccount)
	args.AutoCreateNetwork = pulumi.Bool(project.Args.AutoCreateNetwork)

	if project.Args.FolderId == "" {
		args.OrgId = pulumi.String(project.Args.OrgId)
	} else {
		args.FolderId = pulumi.String(project.Args.FolderId)
	}

	//args.Labels = project.Args.Labels

	gcpProject, err = organizations.NewProject(ctx, project.Args.Name, args)
	err = project.EnableServices(ctx, gcpProject.ProjectId, project.Args.Services)
	if err != nil {
		log.Println(err)
	}

	ctx.Export("projectId", gcpProject.ProjectId)
	return gcpProject, err
}

func (project *Project) EnableServices(ctx *pulumi.Context, projectId pulumi.StringOutput, Services []string) (err error) {
	projectId.ApplyT(func(pid string) error {
		for i, s := range Services {
			fmt.Println(s)
			_, err := projects.NewService(ctx, fmt.Sprintf("service-%s-%d-%s", pid, i, s), &projects.ServiceArgs{
				DisableDependentServices: pulumi.Bool(false),
				DisableOnDestroy:         pulumi.Bool(false),
				Project:                  projectId,
				Service:                  pulumi.String(s),
			})
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})
	return nil
}
