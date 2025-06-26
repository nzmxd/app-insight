package config

type Server struct {
	Downloader     Downloader `mapstructure:"downloader" json:"downloader" yaml:"downloader"`
	StaticAnalyzer Analyzer   `mapstructure:"static-analyzer" json:"static-analyzer" yaml:"static-analyzer"`
}

type Downloader struct {
	Proxy        string   `mapstructure:"proxy" json:"proxy" yaml:"proxy"`
	BackProxies  []string `mapstructure:"back-proxies" json:"back-proxies" yaml:"back-proxies"`
	Worker       int      `mapstructure:"worker" json:"worker" yaml:"worker"`
	Source       string   `mapstructure:"source" json:"source" yaml:"source"`
	Timeout      int      `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	MaxRetry     int      `mapstructure:"max-retry" json:"max-retry" yaml:"max-retry"`
	SavePath     string   `mapstructure:"save-path" json:"save-path" yaml:"save-path"`
	LimitSize    int      `mapstructure:"limit-size" json:"limit-size" yaml:"limit-size"`
	UploadEnable bool     `mapstructure:"upload-enabled" json:"upload-enabled" yaml:"upload-enabled"`
}

type Analyzer struct {
	Use     string `mapstructure:"use" json:"use" yaml:"use"`
	Exec    string `mapstructure:"exec" json:"exec" yaml:"exec"`
	Worker  int    `mapstructure:"worker" json:"worker" yaml:"worker"`
	Workdir string `mapstructure:"workdir" json:"workdir" yaml:"workdir"`
}
