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

package kubernetesengine

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodepoolArgs struct {
	Name       string //Name of the nodepool.
	NamePrefix string //Creates a unique name for the nodepool, replaces Name.
	Project    string //Project for the nodepool, should be the same as the cluster.
	Cluster    string //Cluster name for the nodepool to be provisioned in.
	Location   string //Region or zone for the nodes
	NodeCount  int    //The number of nodes to provision is overridden
	NodeConfig struct {
		Preemptible    bool
		MachineType    string
		ServiceAccount string
	}
}

type Nodepool struct {
	Args NodepoolArgs
	Name string
}

func (nodepool *Nodepool) Create(ctx *pulumi.Context) (gkeNodepool *container.Cluster, err error) {
	args := &cluster.ClusterArgs{}

	args.Name = pulimi.String(cluster.Args.Name)
	args.Description = pulumi.String(cluster.Args.Description)
	args.Project = pulumi.String(cluster.Args.Project)
	args.RemoveDefaultNodePool = pulumi.Bool(cluster.Args.RemoveDefaultNodePool)

	cluster, err := container.NewCluster(ctx, fmt.Sprintf("gke-cluster-%s", args.Name), &args{})
	if err != nil {
		return nil, err
	}
	return cluster, err
}
