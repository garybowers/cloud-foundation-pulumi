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

type ClusterArgs struct {
	Name                  string
	Description           string
	Project               string
	Location              string
	ReleaseChannel        string //The release channel of this cluster. Accepted values are `UNSPECIFIED`, `RAPID`, `REGULAR` and `STABLE`. Defaults to `UNSPECIFIED`.
	Network               string
	Subnetwork            string
	NetworkProject        string //projectId for the project that contains the network.
	CreateServiceAccount  bool   //Creates a service account per nodepool.
	ServiceAccount        string //Only used if the CreateServiceAccount is false.
	RemoveDefaultNodePool bool   //Deletes the default nodepool, defaults to false.

}

type Cluster struct {
	Args ClusterArgs
	Name string
}

func (cluster *Cluster) Create(ctx *pulumi.Context) (gkeCluster *container.Cluster, err error) {
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
