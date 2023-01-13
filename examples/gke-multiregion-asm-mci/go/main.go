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

package main

import (
	"fmt"
	"log"

	"github.com/garybowers/cloud-foundation-pulumi/modules/network"
	"github.com/garybowers/cloud-foundation-pulumi/modules/project"
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
