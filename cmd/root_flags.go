package cmd

type RootFlags struct {
	configFile string
	apikey     string
	baseUrl    string
	verbose    bool
}

func NewRootFlags() *RootFlags {
	return &RootFlags{}
}
