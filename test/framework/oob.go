package framework

const (
	tinkerbellBMCConsumerURL            = "T_TINKERBELL_BMC_CONSUMER_URL"
	tinkerbellBMCHMACSecret             = "T_TINKERBELL_BMC_HMAC_SECRETS"
	tinkerbellBMCTimestampHeader        = "T_TINKERBELL_BMC_TIMESTAMP_HEADER"
	tinkerbellBMCIncludedPayloadHeaders = "T_TINKERBELL_BMC_INCLUDED_PAYLOAD_HEADERS"
)

var requiredOOBEnvVars = []string{
	tinkerbellBMCConsumerURL,
	tinkerbellBMCHMACSecret,
	tinkerbellBMCTimestampHeader,
	tinkerbellBMCIncludedPayloadHeaders,
}

// RequiredOOBEnvVars returns the environment variables required to run OOB related e2e tests.
func RequiredOOBEnvVars() []string { _ = "STUB: not implemented"; return nil }

// WithOOBConfiguration sets up the required environment to run OOB e2e tests.
func WithOOBConfiguration() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}
