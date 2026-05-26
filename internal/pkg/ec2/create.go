package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var dockerLogsUserData = `
#!/bin/bash
cat <<'EOF' >> /etc/docker/daemon.json
{
  "log-driver": "journald",
  "log-level": "debug"
}
EOF
systemctl restart docker --no-block
`

func CreateInstance(session *session.Session, amiId, key, tag, instanceProfileName, subnetId, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// EC2 Request token bucket has a refill rate of 2 request tokens
// per second, so waiting between 5 and 10 seconds per retry with a backoff factor of 1.5 should be sufficient

func isThrottleError(err error) bool { _ = "STUB: not implemented"; return false }

func getRandomSubnetID(subnetIDsStr string) string { _ = "STUB: not implemented"; return "" }
