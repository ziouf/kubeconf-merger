package v1

type Config struct {
	ApiVersion     string            `yaml:"apiVersion"`
	Kind           string            `yaml:"kind"`
	CurrentContext string            `yaml:"current-context"`
	Contexts       []Context         `yaml:"contexts"`
	Clusters       []Cluster         `yaml:"clusters"`
	Users          []User            `yaml:"users"`
	Preferences    map[string]string `yaml:"preferences"`
}

type Context struct {
	Name    string `yaml:"name"`
	Context struct {
		Cluster string `yaml:"cluster"`
		User    string `yaml:"user"`
	} `yaml:"context"`
}

type Cluster struct {
	Name    string `yaml:"name"`
	Cluster struct {
		CertificateAuthorityData string `yaml:"certificate-authority-data"`
		Server                   string `yaml:"server"`
	} `yaml:"cluster"`
}

type User struct {
	Name string `yaml:"name"`
	User struct {
		Token string `yaml:"token"`
	} `yaml:"user"`
}

func (c1 *Config) Merge(c2 Config) Config {
	return Config{
		ApiVersion:     c1.ApiVersion,
		Kind:           c1.Kind,
		CurrentContext: c1.CurrentContext,
		Contexts:       append(c1.Contexts, c2.Contexts...),
		Clusters:       append(c1.Clusters, c2.Clusters...),
		Users:          append(c1.Users, c2.Users...),
		Preferences:    c1.Preferences,
	}
}
