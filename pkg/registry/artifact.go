package registry

// Artifact to head release dependency.
type Artifact struct {
	Registry   string
	Repository string
	Tag        string
	Digest     string
}

// NewArtifact creates a new artifact object.
func NewArtifact(registry, repository, tag, digest string) Artifact {
	_ = "STUB: not implemented"
	return *new(Artifact)
}

// NewArtifactFromURI creates a new artifact object from a URI.
func NewArtifactFromURI(uri string) Artifact { _ = "STUB: not implemented"; return *new(Artifact) }

// Version returns tag or digest.
func (art *Artifact) Version() string { _ = "STUB: not implemented"; return "" }

// VersionedImage returns full URI for image.
func (art *Artifact) VersionedImage() string { _ = "STUB: not implemented"; return "" }
