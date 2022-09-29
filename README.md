# Google Cloud - Foundation Example in Pulumi Native

## Development

A dockerfile has been provided which provides all the necessary tooling to get started including:

* golang
* pulumi
* google-cloud-sdk
* git
* make
* code-server

To get started simply run `docker build . -t pulumi:dev` then run the container mounting your local google credentials and source code repository e.g. `docker `