package options

import (
	"strings"

	"github.com/anchore/syft/syft/pkg/cataloger/dotnet"
)

type dotnetConfig struct {
	SearchNuGetLicenses  bool   `yaml:"search-nuget-licenses" json:"search-nuget-licenses" mapstructure:"search-nuget-licenses"`
	SearchRemoteLicenses bool   `yaml:"search-remote-licenses" json:"search-remote-licenses" mapstructure:"search-remote-licenses"`
	Providers            string `yaml:"package-providers,omitempty" json:"package-providers,omitempty" mapstructure:"package-providers"`
}

func defaultDotnetConfig() dotnetConfig {
	def := dotnet.DefaultCatalogerConfig()
	return dotnetConfig{
		SearchNuGetLicenses:  def.SearchNuGetLicenses,
		SearchRemoteLicenses: def.SearchRemoteLicenses,
		Providers:            strings.Join(def.Providers, ","),
	}
}
