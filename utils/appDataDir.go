package utils

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"runtime"
)

func GetAppDataDir(appName string) string {
	homeDir, err := homedir.Dir()

	dotName := fmt.Sprintf(".%s", appName)
	if err != nil {
		return dotName
	}

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("LOCALAPPDATA")
		if appData == "" {
			appData = os.Getenv("APPDATA")
		}

		if appData != "" {
			return filepath.Join(appData, appName)
		}
	case "darwin":
		if homeDir != "" {
			return filepath.Join(homeDir, "Library", "Application Support", appName)
		}
	case "linux":
		xdgDataHome := os.Getenv("XDG_DATA_HOME")

		if xdgDataHome == "" {
			if homeDir == "" {
				return filepath.Join(homeDir, dotName)
			}

			xdgDataHome = filepath.Join(homeDir, ".local", "share")
		}

		return filepath.Join(xdgDataHome, appName)
	default:
		if homeDir != "" {
			return filepath.Join(homeDir, dotName)
		}
	}

	return dotName
}
