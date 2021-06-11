package network

import (
	"fmt"
	"net"

	sdc "github.com/javadh75/systemd-config"
)

func CreateMatchSection(name string, mac net.HardwareAddr, linkType string) *sdc.Section {
	matchSection := sdc.NewSection("Match")

	nameOption := sdc.NewUnitOption("Name", name)
	matchSection.Options = append(matchSection.Options, nameOption)

	macOption := sdc.NewUnitOption("PermanentMACAddress", mac.String())
	matchSection.Options = append(matchSection.Options, macOption)

	sdcLinkType := "ether"
	switch linkType {
	case "vlan":
		sdcLinkType = "vlan"
	}
	linkTypeOption := sdc.NewUnitOption("Type", sdcLinkType)
	matchSection.Options = append(matchSection.Options, linkTypeOption)

	return matchSection
}

func CreateLinkSection(mac net.HardwareAddr, mtu int, multicast bool) *sdc.Section {
	linkSection := sdc.NewSection("Link")

	macOption := sdc.NewUnitOption("MACAddress", mac.String())
	linkSection.Options = append(linkSection.Options, macOption)

	mtuOption := sdc.NewUnitOption("MTUBytes", fmt.Sprintf("%d", mtu))
	linkSection.Options = append(linkSection.Options, mtuOption)

	multicastOption := sdc.NewUnitOption("Multicast", fmt.Sprintf("%t", multicast))
	linkSection.Options = append(linkSection.Options, multicastOption)

	return linkSection
}

func CreateAddressSection(ip net.IPNet) *sdc.Section {
	addressSection := sdc.NewSection("Address")

	addressOption := sdc.NewUnitOption("Address", ip.String())
	addressSection.Options = append(addressSection.Options, addressOption)

	return addressSection
}
