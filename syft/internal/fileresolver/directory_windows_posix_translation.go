package fileresolver

import (
	"path"
	"path/filepath"
	"strings"
)

func windowsToPosix(windowsPath string) (posixPath string) {
	// volume should be encoded at the start (e.g /c/<path>) where c is the volume
	volumeName := filepath.VolumeName(windowsPath)
	pathWithoutVolume := strings.TrimPrefix(windowsPath, volumeName)
	volumeLetter := strings.ToLower(strings.TrimSuffix(volumeName, ":"))

	if volumeLetter == "" {
		// We have a relative path
		return path.Clean(strings.ReplaceAll(windowsPath, "\\", "/"))
	}

	// translate non-escaped backslash to forwardslash
	translatedPath := strings.ReplaceAll(pathWithoutVolume, "\\", "/")

	// always have `/` as the root... join all components, e.g.:
	// convert: C:\\some\windows\Place
	// into: /c/some/windows/Place
	return path.Clean("/" + strings.Join([]string{volumeLetter, translatedPath}, "/"))
}

func posixToWindows(posixPath string) (windowsPath string) {
	// decode the volume (e.g. /c/<path> --> C:\\) - There should always be a volume name.
	pathFields := strings.Split(posixPath, "/")
	if len(pathFields[0]) == 0 &&
		len(pathFields[1]) == 1 &&
		strings.Contains("abcdefghijklmnopqrstuvwxyz", pathFields[1]) {
		// We have an absolute path
		volumeName := strings.ToUpper(pathFields[1]) + `:\\`

		// translate non-escaped forward slashes into backslashes
		remainingTranslatedPath := strings.Join(pathFields[2:], "\\")

		// combine volume name and backslash components
		return filepath.Clean(volumeName + remainingTranslatedPath)
	}

	return filepath.Clean(strings.Join(pathFields, "\\"))
}
