/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"go.bytebuilders.dev/license-verifier/info"
	verifier "go.bytebuilders.dev/license-verifier/kubernetes"

	flag "github.com/spf13/pflag"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var (
	masterURL      = flag.String("master", "", "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	kubeconfigPath = flag.String("kubeconfig", "", "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	licenseFile    = flag.String("license-file", "", "Path to license file.")
)

func main() {
	klog.InitFlags(nil)
	_ = flag.Set("v", "3")
	defer klog.Flush()

	flag.Parse()

	PrintInfo()

	config, err := clientcmd.BuildConfigFromFlags(*masterURL, *kubeconfigPath)
	if err != nil {
		klog.Fatalf("could not get Kubernetes config: %s", err)
	}

	stopCh := genericapiserver.SetupSignalHandler()

	//nolint:errcheck
	go verifier.VerifyLicensePeriodically(config, *licenseFile, stopCh)

	select {}
}

func PrintInfo() {
	fmt.Println("EnforceLicense", info.EnforceLicense)
	fmt.Println("LicenseCA", info.LicenseCA)
	fmt.Println("ProductOwnerName", info.ProductOwnerName)
	fmt.Println("ProductOwnerUID", info.ProductOwnerUID)
	fmt.Println("ProductName", info.ProductName)
	fmt.Println("ProductUID", info.ProductUID)
}
