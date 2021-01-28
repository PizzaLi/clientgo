module k8s/clientgo

go 1.15

replace k8s/clientgo/incluster => ./incluster

replace k8s/clientgo/outcluster => ./outcluster

require (
	github.com/Azure/go-autorest/autorest v0.11.12 // indirect
	github.com/gophercloud/gophercloud v0.14.0 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392 // indirect
	golang.org/x/net v0.0.0-20201201195509-5d6afe98e0b7 // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)
