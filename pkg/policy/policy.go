package policy

type BeePolicy struct {
	Name    string
	Request struct {
		Method  string `yaml:"method"`
		URL     string `yaml:"url"`
		Headers []struct {
			Key   string `yaml:"key"`
			Value string `yaml:"value"`
		} `yaml:"headers"`
	} `yaml:"request"`
}
