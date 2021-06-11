package network

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
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
		fmt.Printf("Unable to convert network config to []byte: %+v\n", r)
		return false, err
	}

	if err = ioutil.WriteFile(filePath, config, 0644); err != nil {
		fmt.Printf("Unable to write network config to file %s: %s\n", filePath, string(config))
		return false, err
	}
	return true, nil
}
