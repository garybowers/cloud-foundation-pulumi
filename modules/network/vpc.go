package network

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpcArgs struct {
	Name                        string
	Description                 string
	ProjectId                   string
	RoutingMode                 string
	AutoCreateSubnetworks       bool
	DeleteDefaultRoutesOnCreate bool
	EnableUIaInternalIpv6       bool
	InternalIpv6Range           bool
}

type Vpc struct {
	Args VpcArgs
	Name string
}

func (vpc *Vpc) Create(ctx *pulumi.Context) (vpcNetwork pulumi.Output, err error) {
	args := &compute.NetworkArgs{}
	args.Name = pulumi.String(vpc.Args.Name)
	args.Description = pulumi.String(vpc.Args.Description)
	args.Project = vpc.Args.ProjectId
	args.AutoCreateSubnetworks = pulumi.Bool(vpc.Args.AutoCreateSubnetworks)
	args.DeleteDefaultRoutesOnCreate = pulumi.Bool(vpc.Args.DeleteDefaultRoutesOnCreate)

	vpcSL := vpc.Args.ProjectId.ApplyT(func(pid string) pulumi.StringOutput {
		vpcnetwork, err := compute.NewNetwork(ctx, fmt.Sprintf("%s-%s", pid, args.Name), args)
		if err != nil {
			log.Println(err)
		}
		return vpcnetwork.SelfLink
	})

	//ctx.Export("vpc", vpct.ID())
	//return vpct, err
	return vpcSL, err
}
