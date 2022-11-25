package config

type Log struct {
	ConsoleLoggingEnabled bool   `yaml:"ConsoleLoggingEnabled"`
	EncodeLogsAsJson      bool   `yaml:"EncodeLogsAsJson"`
	FileLoggingEnabled    bool   `yaml:"FileLoggingEnabled"`
	Directory             string `yaml:"Directory"`
	Filename              string `yaml:"Filename"`
	MaxSize               int    `yaml:"MaxSize"`
	MaxBackups            int    `yaml:"MaxBackups"`
	MaxAge                int    `yaml:"MaxAge"`
}
