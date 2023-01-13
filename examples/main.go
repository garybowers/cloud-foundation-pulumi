package main

import (
	"fmt"
	"log"
	"github.com/garybowers/b1-infra-new/modules/project"
	"github.com/garybowers/b1-infra-new/modules/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	var project project.Project
	var vpc network.Vpc
	var subnet network.Subnet

	pulumi.Run(func(ctx *pulumi.Context) error {
		project.Args.Name = "b1-services-new"
		project.Args.FolderId = "folders/415061719873"
		project.Args.BillingAccount = "01504C-A2522F-2110FA"
		project.Args.AutoCreateNetwork = false

		project.Args.Services = []string{"compute.googleapis.com", "container.googleapis.com"}

		prj, err := project.Create(ctx)
		if err != nil {
			log.Println(err)
		}

		vpc.Args.Name = "vpc-1"
		vpc.Args.ProjectId = prj.ProjectId
		//vpc.Args.Project = pulumi.Sprintf("%s", prj.ProjectId)


		vpcNetwork, err := vpc.Create(ctx)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("NETWORK:")

		subnet.Args.Name = "euw1"
		subnet.Args.Description = "europe-west1"
		subnet.Args.ProjectId = prj.ProjectId
		subnet.Args.Region = "europe-west1"
		subnet.Args.IpCidrRange = "10.76.34.0/22"
		subnet.Args.Network = vpcNetwork

		subnet, err := subnet.Create(ctx)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(subnet)

		return nil
	})
}
