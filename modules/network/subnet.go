package network

import (
	"fmt"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
)

type SubnetArgs struct {
	Name                  string
	Description           string
	Network               pulumi.Output
	ProjectId             pulumi.StringOutput
	Region                string
	PrivateIpGoogleAccess bool
	IpCidrRange           string
	Purpose               string
	Role                  string
}

type Subnet struct {
	Args SubnetArgs
	Name string
}

func (subnet *Subnet) Create(ctx *pulumi.Context) (subnetwork *compute.Subnetwork, err error) {

	args := &compute.SubnetworkArgs{}
	args.Name = pulumi.String(subnet.Args.Name)
	args.Description = pulumi.String(subnet.Args.Description)
	args.Project = subnet.Args.ProjectId
	args.Network = subnet.Args.Network
	args.IpCidrRange = pulumi.String(subnet.Args.IpCidrRange)
	args.PrivateIpGoogleAccess = pulumi.Bool(true)

	subnet.Args.ProjectId.ApplyT(func(pid string) error {
		subnetwork, err = compute.NewSubnetwork(ctx, fmt.Sprintf("%s-%s", pid, args.Name), args)
		return err
	})
	if err != nil {
		log.Println(err)
	}
	ctx.Export("subnet", subnetwork)
	return subnetwork, err
}
