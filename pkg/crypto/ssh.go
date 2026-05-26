package crypto

import (
	"io"

	"github.com/aws/eks-anywhere/pkg/filewriter"
)

// SshKeysize is the key size used when calling NewSshKeyPair().
const SshKeySize = 4096

// NewSshKeyPair creates an RSA public key pair and writes each part to privateOut and publicOut. The output
// written to privateOut and pulicKeyOut is formatted as ssh-keygen would format keys.
// The private key part is PEM encoded with the key data formatted in PKCS1, ASN.1 DER as typically done by
// the ssh-keygen GNU tool. The public key is formatted as an SSH Authorized Key suitable for storing on servers.
func NewSshKeyPair(privateOut, publicOut io.Writer) error { _ = "STUB: not implemented"; return nil }

// NewSshKeyPairUsingFileWriter provides a mechanism for generating SSH key pairs and writing them to the writer
// direcftory context. It exists to create compatibility with filewriter.FileWriter and compliment older code.
// The string returned is a path to the private key written to disk using writer.
// The bytes returned are the public key formated as specified in NewSshKeyPair().
func NewSshKeyPairUsingFileWriter(writer filewriter.FileWriter, privateKeyFilename, publicKeyFilename string) (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
