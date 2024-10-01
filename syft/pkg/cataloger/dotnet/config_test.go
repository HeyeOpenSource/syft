package dotnet

import (
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func Test_Config(t *testing.T) {
	type opts struct {
		remote    bool
		providers string
	}

	homedirCacheDisabled := homedir.DisableCache
	homedir.DisableCache = true
	t.Cleanup(func() {
		homedir.DisableCache = homedirCacheDisabled
	})

	allEnv := map[string]string{
		"HOME":                         "/usr/home",
		"NUGET_PACKAGE_PROVIDERS":      "",
		"NUGET_SEARCH_REMOTE_LICENSES": "",
	}

	tests := []struct {
		name     string
		env      map[string]string
		opts     opts
		expected CatalogerConfig
	}{
		{
			name: "absolute defaults",
			env:  map[string]string{},
			opts: opts{},
			expected: CatalogerConfig{
				SearchRemoteLicenses: false,
				Providers:            []string{"https://www.nuget.org/api/v2/package"},
			},
		},
		{
			name: "set via env defaults",
			env: map[string]string{
				"NUGET_PACKAGE_PROVIDERS":      "https://my.proxy",
				"NUGET_SEARCH_REMOTE_LICENSES": "false",
			},
			opts: opts{},
			expected: CatalogerConfig{
				SearchRemoteLicenses: false,
				Providers:            []string{"https://my.proxy"},
			},
		},
		{
			name: "set via configuration",
			env: map[string]string{
				"NUGET_PACKAGE_PROVIDERS":      "https://my.proxy",
				"NUGET_SEARCH_REMOTE_LICENSES": "true",
			},
			opts: opts{
				remote:    true,
				providers: "https://www.nuget.org/api/v2/package",
			},
			expected: CatalogerConfig{
				SearchRemoteLicenses: true,
				Providers:            []string{"https://www.nuget.org/api/v2/package"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for k, v := range allEnv {
				t.Setenv(k, v)
			}
			for k, v := range test.env {
				t.Setenv(k, v)
			}
			got := DefaultCatalogerConfig().
				WithSearchRemoteLicenses(test.opts.remote).
				WithProviders(test.opts.providers)

			assert.Equal(t, test.expected, got)
		})
	}
}
