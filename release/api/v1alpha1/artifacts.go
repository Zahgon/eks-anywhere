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

package v1alpha1

// Manifests returns a map of manifests for different components in a VersionsBundle.
func (vb *VersionsBundle) Manifests() map[string][]*string { _ = "STUB: not implemented"; return nil }

// Ovas returns a list of OVA archives in a VersionsBundle.
func (vb *VersionsBundle) Ovas() []Archive { _ = "STUB: not implemented"; return nil }

// CloudStackImages returns images needed for the CloudStack provider in a VersionsBundle.
func (vb *VersionsBundle) CloudStackImages() []Image { _ = "STUB: not implemented"; return nil }

// VsphereImages returns images needed for the vSphere provider in a VersionsBundle.
func (vb *VersionsBundle) VsphereImages() []Image { _ = "STUB: not implemented"; return nil }

// DockerImages returns images needed for the Docker provider in a VersionsBundle.
func (vb *VersionsBundle) DockerImages() []Image { _ = "STUB: not implemented"; return nil }

// SnowImages returns images needed for the Snow provider in a VersionsBundle.
func (vb *VersionsBundle) SnowImages() []Image { _ = "STUB: not implemented"; return nil }

// TinkerbellImages returns images needed for the Tinkerbell provider in a VersionsBundle.
func (vb *VersionsBundle) TinkerbellImages() []Image { _ = "STUB: not implemented"; return nil }

// NutanixImages returns images needed for the Nutanix provider in a VersionsBundle.
func (vb *VersionsBundle) NutanixImages() []Image { _ = "STUB: not implemented"; return nil }

// SharedImages returns images that are shared across different providers in a VersionsBundle.
func (vb *VersionsBundle) SharedImages() []Image { _ = "STUB: not implemented"; return nil }

// Images returns all images from the VersionsBundle by aggregating those from different providers.
func (vb *VersionsBundle) Images() []Image { _ = "STUB: not implemented"; return nil }

// Charts returns a map of Helm chart images used by different components in a VersionsBundle.
func (vb *VersionsBundle) Charts() map[string]*Image { _ = "STUB: not implemented"; return nil }
