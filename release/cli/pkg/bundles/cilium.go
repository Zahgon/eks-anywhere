// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
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

package bundles

import (
	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

const (
	ciliumImageName         = "cilium"
	ciliumOperatorImageName = "operator-generic"
	ciliumHelmChartName     = "cilium-chart"
	ciliumHelmChart         = "cilium"
	ciliumImage             = "cilium"
	ciliumOperatorImage     = "operator-generic"
)

func GetCiliumBundle(r *releasetypes.ReleaseConfig) (anywherev1alpha1.CiliumBundle, error) {
	_ = "STUB: not implemented"
	return *new(anywherev1alpha1.CiliumBundle), nil
}

// Helm charts are in the same repository and have the same
// sem version as the corresponding container image but omiting the initial "v"

func getCiliumImageDigest(gitRootPath, imageName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type imageDefinition struct {
	name, image, registry, tag string
	builder                    imageBuilder
}

type imageBuilder func(digest string) anywherev1alpha1.Image

func containerImage(name, image, registry, tag string) imageDefinition {
	_ = "STUB: not implemented"
	return *new(imageDefinition)
}

func chart(name, image, registry, tag string) imageDefinition {
	_ = "STUB: not implemented"
	return *new(imageDefinition)
}
