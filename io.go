//go:build !js
// +build !js

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func DataDir() (string, error) {
	var xdg string
	if runtime.GOOS == "windows" {
		xdg = os.Getenv("LOCALAPPDATA")
	} else {
		xdg = os.Getenv("XDG_DATA_HOME")
	}

	if xdg == "" {
		xdg = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}

	dataDir := filepath.Join(xdg, "rrouge")
	_, err := os.Stat(dataDir)
	if err != nil {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return "", fmt.Errorf("building data directory: %v\n", err)
		}
	}

	return dataDir, nil
}

func SaveFile(filename string, data []byte) error {
	dataDir, err := DataDir()
	if err != nil {
		return err
	}
	tempSaveFile := filepath.Join(dataDir, "temp-"+filename)
	f, err := os.OpenFile(tempSaveFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	saveFile := filepath.Join(dataDir, filename)
	if err := os.Rename(f.Name(), saveFile); err != nil {
		return err
	}
	return err
}

func SaveConfig() error {
	data, err := GameConfig.ConfigSave()
	if err != nil {
		return err
	}

	return SaveFile("config.gob", data)
}

func LoadConfig() (bool, error) {
	dataDir, err := DataDir()
	if err != nil {
		return false, err
	}

	saveFile := filepath.Join(dataDir, "config.gob")
	_, err = os.Stat(saveFile)
	if err != nil {
		// no save file, new game
		return false, nil
	}

	data, err := ioutil.ReadFile(saveFile)
	if err != nil {
		return false, err
	}

	c, err := DecodeConfigSave(data)
	if err != nil {
		return false, err
	}

	if c.Version != GameConfig.Version {
		return false, nil
	}

	GameConfig = *c
	return true, nil
}
