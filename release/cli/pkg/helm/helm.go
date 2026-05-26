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

package helm

import (
	"github.com/go-logr/logr"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	ctrl "sigs.k8s.io/controller-runtime"

	anywherev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	releasetypes "github.com/aws/eks-anywhere/release/cli/pkg/types"
)

var HelmLog = ctrl.Log.WithName("HelmLog")

// helmDriver implements PackageDriver to install packages from Helm charts.
type helmDriver struct {
	cfg      *action.Configuration
	log      logr.Logger
	settings *cli.EnvSettings
}

func NewHelm() (*helmDriver, error) { _ = "STUB: not implemented"; return nil, nil }

func GetHelmDest(d *helmDriver, r *releasetypes.ReleaseConfig, sourceImageURI, assetName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetChartImageTags(d *helmDriver, helmDest string) (*Requires, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ModifyAndPushChartYaml(i releasetypes.ImageArtifact, r *releasetypes.ReleaseConfig, d *helmDriver, helmDest string, eksaArtifacts map[string][]releasetypes.Artifact, shaMap map[string]anywherev1alpha1.Image) error {
	_ = "STUB: not implemented"
	return nil
}

// Overwrite Chart.yaml

// If the chart is packages, we find the image tag values and overide them in the values.yaml.

func (d *helmDriver) HelmRegistryLogin(r *releasetypes.ReleaseConfig, remoteType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *helmDriver) HelmRegistryLogout(r *releasetypes.ReleaseConfig, remoteType string) error {
	_ = "STUB: not implemented"
	return nil
}

// PullHelmChart will take in a a remote Helm URI and attempt to pull down the chart if it exists.
func (d *helmDriver) PullHelmChart(name, version string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PushHelmChart will take in packaged helm chart and push to a remote URI.
func PushHelmChart(packaged, URI string) error { _ = "STUB: not implemented"; return nil }

// PackageHelmChart will package a dir into a helm chart.
func PackageHelmChart(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// helmLog wraps logr.Logger to make it compatible with helm's DebugLog.
func helmLog(log logr.Logger) action.DebugLog {
	_ = "STUB: not implemented"
	return *new(action.DebugLog)
}

// UnTarHelmChart will attempt to move the helm chart out of the helm cache, by untaring it to the pwd and creating the filesystem to unpack it into.
func UnTarHelmChart(chartRef, chartPath, dest string) error { _ = "STUB: not implemented"; return nil }

// Checks directory check errors such as permission issues to read

// Untar the files, and create the directory structure

// HasRequires checks for the existance of the requires.yaml within the helm directory.
func HasRequires(helmdir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ValidateHelmRequires runs the parse file into struct function, and validations.
func ValidateHelmRequires(fileName string) (*Requires, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateHelmRequiresContent loops over the validation tests.
func validateHelmRequiresContent(helmrequires *Requires) error {
	_ = "STUB: not implemented"
	return nil
}

var helmRequiresValidations = []func(*Requires) error{
	validateHelmRequiresName,
}

func validateHelmRequiresName(helmrequires *Requires) error { _ = "STUB: not implemented"; return nil }

// validateHelmRequiresNotEmpty checks that it has at least one image in the spec.
func (helmrequires *Requires) validateHelmRequiresNotEmpty() error {
	_ = "STUB: not implemented"
	// Check if Projects are listed
	return nil
}

// parseHelmRequires will attempt to unpack the requires.yaml into the Go struct `Requires`.
func parseHelmRequires(fileName string, helmrequires *Requires) error {
	_ = "STUB: not implemented"
	return nil
}

// Chart yaml functions

// HasChart checks for the existance of the Chart.yaml within the helm directory.
func HasChart(helmdir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ValidateHelmChart runs the parse file into struct function, and validations.
func ValidateHelmChart(fileName string) (*chart.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseHelmChart will attempt to unpack the Chart.yaml into the Go struct `Chart`.
func parseHelmChart(fileName string, helmChart *chart.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func OverwriteChartYaml(filename string, helmChart *chart.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func OverWriteChartValuesImageTag(filename string, tagMap map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func OverWriteChartValuesImageSha(filename string, shaMap map[string]anywherev1alpha1.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func GetPackagesImageTags(packagesArtifacts map[string][]releasetypes.Artifact) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
