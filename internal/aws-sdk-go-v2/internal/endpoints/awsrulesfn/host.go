package awsrulesfn

// IsVirtualHostableS3Bucket returns if the input is a DNS compatible bucket
// name and can be used with Amazon S3 virtual hosted style addressing. Similar
// to [rulesfn.IsValidHostLabel] with the added restriction that the length of label
// must be [3:63] characters long, all lowercase, and not formatted as an IP
// address.
func IsVirtualHostableS3Bucket(input string, allowSubDomains bool) bool {
	_ = "STUB: not implemented"
	// input should not be formatted as an IP address
	// NOTE: this will technically trip up on IPv6 hosts with zone IDs, but
	// validation further down will catch that anyway (it's guaranteed to have
	// unfriendly characters % and : if that's the case)
	return false
}

// validate special length constraints

// Validate no capital letters

// Validate valid host label
