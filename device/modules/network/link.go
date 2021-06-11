package network

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	sdc "github.com/javadh75/systemd-config"
	"github.com/vishvananda/netlink"
)

func DelLink(name string) (bool, error) {
	link, err := GetLink(name)
	if err != nil {
		return false, nil
	}
	if err = netlink.LinkDel(link); err != nil {
		log.Printf("Error in deleting link in DelLink(%s): %s", name, err)
		return false, nil
	}
	return true, nil
}

func GetLink(name string) (netlink.Link, error) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		log.Printf("Error in getting link in GetLink(%s): %s", name, err)
	}
	return link, err
}

func GetAllLinks() ([]netlink.Link, error) {
	links, err := netlink.LinkList()
	if err != nil {
		log.Printf("Error in getting links in GetAllLinks(): %s", err)
	}
	return links, err
}

func InitCurrentConfigs() error {
	links, err := GetAllLinks()
	if err != nil {
		log.Printf("Error in getting links for initializing current configs")
		return err
	}
	for _, link := range links {
		if link.Attrs().Name == "lo" {
			continue
		}
		existance, err := CheckLinkConfigExistense(link.Attrs().Name)
		if err != nil {
			return err
		}
		if !existance {
			linkConfig := CreateLinkConfig(link)
			NetworkConfigWriter(linkConfig, link.Attrs().Name)
		}
	}
	return nil
}

func CheckLinkConfigExistense(name string) (bool, error) {
	linkFilePath := filepath.Join(BASE_SYSTEMD_NETWORKD_CONFIG_PATH, name+NETWORKD_LINK_CONFIG_SUFFIX)
	_, err := os.Stat(linkFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("Error in checking existanse of file %s: %v\n", linkFilePath, err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func CreateLinkConfig(link netlink.Link) io.Reader {
	linkUnit := sdc.NewUnit()

	matchSection := CreateMatchSection(link.Attrs().Name, link.Attrs().HardwareAddr, link.Type())
	linkUnit.Sections = append(linkUnit.Sections, matchSection)

	linkSection := CreateLinkSection(link.Attrs().HardwareAddr, link.Attrs().MTU, link.Attrs().Flags&net.FlagMulticast != 0)
	linkUnit.Sections = append(linkUnit.Sections, linkSection)

	addresses, _ := netlink.AddrList(link, 0)
	for _, address := range addresses {
		addressSection := CreateAddressSection(*address.IPNet)
		linkUnit.Sections = append(linkUnit.Sections, addressSection)
	}
	return sdc.Serialize(linkUnit)
}
