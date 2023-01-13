// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
