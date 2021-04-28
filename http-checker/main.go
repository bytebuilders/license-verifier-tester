package main

import (
	"encoding/json"
	"fmt"
	"go.bytebuilders.dev/license-verifier/apis/licenses/v1alpha1"
	"gomodules.xyz/sets"
	"strings"
)

const resp = `{
  "apiVersion": "licenses.appscode.com/v1alpha1",
  "clusters": [
    "44728b88-bafc-4bfe-ae08-bf95103082c8"
  ],
  "features": [
    "kubedb-community",
    "kubedb-ext-stash",
    "kubedb-autoscaler",
    "kubedb-enterprise"
  ],
  "id": "3196503369123089194",
  "issuer": "byte.builders",
  "kind": "License",
  "notAfter": "2021-05-28T16:49:14Z",
  "notBefore": "2021-04-28T16:49:14Z",
  "planName": "kubedb-enterprise",
  "reason": "",
  "status": "active",
  "user": {
    "email": "tamal@appscode.com",
    "name": "Tamal Saha"
  }
}`

var SupportedProducts = []string{"stash-enterprise", "kubedb-ext-stash"}

func main() {
	err := check(SupportedProducts)
	if err != nil {
		panic(err)
	}
}

func check(features []string) error {
	var license v1alpha1.License
	err := json.Unmarshal([]byte(resp), &license)
	if err != nil {
		return err
	}

	if license.Status != v1alpha1.LicenseActive {
		return fmt.Errorf("license %s is not active, status: %s, reason: %s", license.ID, license.Status, license.Reason)
	}

	if !sets.NewString(license.Features...).HasAny(features...) {
		return fmt.Errorf("license %s is not valid for products %q", license.ID, strings.Join(features, ","))
	}
	return nil
}
