package conf

import "flag"

const (
	FlagConfig = "config"
	FlagI18n   = "i18n"
)

var (
	ConfigPath string
	I18nPath   string
)

func init() {
	flag.StringVar(&ConfigPath, FlagConfig, "config.yaml", "config file path")
	flag.StringVar(&I18nPath, FlagI18n, "i18n", "i18n folder/file path")
	flag.Parse()
}
