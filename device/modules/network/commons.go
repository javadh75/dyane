package network

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	BASE_SYSTEMD_NETWORKD_CONFIG_PATH = "/etc/systemd/network/"

	NETWORKD_LINK_CONFIG_SUFFIX = ".network"
)

func NetworkConfigWriter(r io.Reader, name string) (bool, error) {
	return ConfigWriter(r, filepath.Join(BASE_SYSTEMD_NETWORKD_CONFIG_PATH, name+NETWORKD_LINK_CONFIG_SUFFIX))
}

func ConfigWriter(r io.Reader, filePath string) (bool, error) {
	config, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Unable to convert systemd config to []byte: %+v\n", r)
		return false, err
	}

	if err = os.WriteFile(filePath, config, 0644); err != nil {
		fmt.Printf("Unable to write systemd config to file %s: %s\n", filePath, string(config))
		return false, err
	}
	return true, nil
}

func ConfigReader(filePath string) (io.Reader, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Unable to read config %s: %s\n", filePath, err)
		return nil, err
	}
	s := string(f)
	return strings.NewReader(s), nil
}
