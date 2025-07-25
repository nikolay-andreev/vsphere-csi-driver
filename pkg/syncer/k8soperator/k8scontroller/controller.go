/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8scontroller

import (
	cnstypes "github.com/vmware/govmomi/cns/types"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager, cnstypes.CnsClusterFlavor) error

// AddToManager adds all Controllers to the Manager
func AddToManager(manager manager.Manager, clusterFlavor cnstypes.CnsClusterFlavor) error {
	for _, f := range AddToManagerFuncs {
		if err := f(manager, clusterFlavor); err != nil {
			return err
		}
	}
	return nil
}
