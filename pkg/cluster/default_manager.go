package cluster

var defaultManager *ConfigManager

func init() {
	var err error
	defaultManager, err = NewDefaultConfigManager()
	if err != nil {
		panic(err)
	}
}

func manager() *ConfigManager { _ = "STUB: not implemented"; return nil }

func NewDefaultConfigManager() (*ConfigManager, error) { _ = "STUB: not implemented"; return nil, nil }
