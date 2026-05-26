/*
Copyright 2022 The Tinkerbell Authors.

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

package v1beta1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	osUbuntu             = "ubuntu"
	defaultUbuntuVersion = "20.04"
)

var (
	_ webhook.CustomValidator = &TinkerbellCluster{}
	_ webhook.CustomDefaulter = &TinkerbellCluster{}
)

// SetupWebhookWithManager sets up and registers the webhook with the manager.
func (c *TinkerbellCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:wrapcheck

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (c *TinkerbellCluster) ValidateCreate(_ context.Context, _ runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"

	// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
	return *new(admission.Warnings), nil
}

func (c *TinkerbellCluster) ValidateUpdate(_ context.Context, _, _ runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"

	// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
	return *new(admission.Warnings), nil
}

func (c *TinkerbellCluster) ValidateDelete(_ context.Context, _ runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func defaultVersionForOSDistro(distro string) string { _ = "STUB: not implemented"; return "" }

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type.
func (c *TinkerbellCluster) Default(_ context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}
