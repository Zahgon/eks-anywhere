package config

import (
	_ "embed"
)

const (
	EksavSphereUsernameKey = "EKSA_VSPHERE_USERNAME"
	EksavSpherePasswordKey = "EKSA_VSPHERE_PASSWORD"
	// EksavSphereCPUsernameKey holds Username for cloud provider.
	EksavSphereCPUsernameKey = "EKSA_VSPHERE_CP_USERNAME"
	// EksavSphereCPPasswordKey holds Password for cloud provider.
	EksavSphereCPPasswordKey = "EKSA_VSPHERE_CP_PASSWORD"
)

type VSphereUserConfig struct {
	EksaVsphereUsername   string
	EksaVspherePassword   string
	EksaVsphereCPUsername string
	EksaVsphereCPPassword string
}

//go:embed static/globalPrivs.json
var VSphereGlobalPrivsFile string

//go:embed static/eksUserPrivs.json
var VSphereUserPrivsFile string

//go:embed static/adminPrivs.json
var VSphereAdminPrivsFile string

//go:embed static/readOnlyPrivs.json
var VSphereReadOnlyPrivs string

func NewVsphereUserConfig() *VSphereUserConfig { _ = "STUB: not implemented"; return nil }

// Cloud provider credentials
