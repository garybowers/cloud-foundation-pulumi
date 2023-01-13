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
