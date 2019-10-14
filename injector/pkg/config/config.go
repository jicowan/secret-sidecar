package config

type Config struct {
	// HTTP Server settings
	Port    int
	TlsCert string
	TlsKey  string

	// Sidecar settings
	SidecarImage  string
	SidecarCpu    string
	SidecarMemory string
	SecretName    string
	Region        string
	LogLevel      string
	EcrSecret     bool
}
