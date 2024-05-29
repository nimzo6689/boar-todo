package config

type ApplicationConfig struct {
	Colors    Colors
	Shortcuts Shortcuts
}

var Configuration *ApplicationConfig = &ApplicationConfig{}

// Default configuration which config file overwrites
func DefaultConfig() *ApplicationConfig {
	conf := &ApplicationConfig{
		Colors:    defaultColors(),
		Shortcuts: defaultShortcuts(),
	}
	return conf
}
