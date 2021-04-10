module go.bytebuilders.dev/license-tester

go 1.15

require (
	github.com/spf13/pflag v1.0.5
	go.bytebuilders.dev/license-verifier v0.8.0
	go.bytebuilders.dev/license-verifier/kubernetes v0.8.0
	k8s.io/apiserver v0.18.9
	k8s.io/client-go v0.18.9
	k8s.io/klog v1.0.0
	kmodules.xyz/client-go v0.0.0-20210406074814-1b9f0e240c49
)

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.18.10-0.20200922195747-1bd1cc8f00d1
