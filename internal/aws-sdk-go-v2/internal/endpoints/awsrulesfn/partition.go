package awsrulesfn

// Partition provides the metadata describing an AWS partition.
type Partition struct {
	ID            string                     `json:"id"`
	Regions       map[string]RegionOverrides `json:"regions"`
	RegionRegex   string                     `json:"regionRegex"`
	DefaultConfig PartitionConfig            `json:"outputs"`
}

// PartitionConfig provides the endpoint metadata for an AWS region or partition.
type PartitionConfig struct {
	Name               string `json:"name"`
	DnsSuffix          string `json:"dnsSuffix"`
	DualStackDnsSuffix string `json:"dualStackDnsSuffix"`
	SupportsFIPS       bool   `json:"supportsFIPS"`
	SupportsDualStack  bool   `json:"supportsDualStack"`
}

type RegionOverrides struct {
	Name               *string `json:"name"`
	DnsSuffix          *string `json:"dnsSuffix"`
	DualStackDnsSuffix *string `json:"dualStackDnsSuffix"`
	SupportsFIPS       *bool   `json:"supportsFIPS"`
	SupportsDualStack  *bool   `json:"supportsDualStack"`
}

const defaultPartition = "aws"

func getPartition(partitions []Partition, region string) *PartitionConfig {
	_ = "STUB: not implemented"
	return nil
}

func mergeOverrides(into PartitionConfig, from RegionOverrides) PartitionConfig {
	_ = "STUB: not implemented"
	return *new(PartitionConfig)
}
