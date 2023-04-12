package conf

import (
	"github.com/charliego93/flagx"
	"github.com/transerver/pkg1/configs"
)

var (
	// check type
	_ configs.IBootstrap = (*bootstrap)(nil)

	// Bootstrap instance
	Bootstrap = &bootstrap{}
)

type bootstrap struct {
	configs.Base `json:",inline" yaml:",inline"`
	I18nPath     string `json:"i18n,omitempty" yaml:"i18n"`
}

func init() {
	i18n := flagx.String("i18n,i", "i18n", flagx.WithDescription("i18n folder/file path"))
	p := flagx.String("config,c", "config.yaml", flagx.WithDescription("config path"))
	flagx.MustParse()

	configs.RegisterBootstrap(Bootstrap)
	configs.Parse(configs.NewFileLoader(*p))
	Bootstrap.I18nPath = *i18n
}
