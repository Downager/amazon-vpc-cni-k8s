module github.com/aws/amazon-vpc-cni-k8s/test

go 1.14

require (
	github.com/apparentlymart/go-cidr v1.0.1
	github.com/aws/amazon-vpc-cni-k8s v1.7.10
	github.com/aws/amazon-vpc-cni-k8s/test/agent v0.0.0-20211209222755-86ece934e91a
	github.com/aws/amazon-vpc-resource-controller-k8s v1.0.7
	github.com/aws/aws-sdk-go v1.43.16
	github.com/fatih/color v1.13.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/onsi/ginkgo/v2 v2.4.0
	github.com/onsi/gomega v1.23.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.4.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
	helm.sh/helm/v3 v3.11.1
	k8s.io/api v0.26.0
	k8s.io/apimachinery v0.26.0
	k8s.io/cli-runtime v0.26.0
	k8s.io/client-go v0.26.0
	sigs.k8s.io/controller-runtime v0.11.2
)
