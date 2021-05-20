package address

import (
	"log"

	link "github.com/javadh75/dyane/device/modules/network/link"
	"github.com/vishvananda/netlink"
)

func AddIPAddress(linkName string, IP string) error {
	nllink, err := link.GetLink(linkName)
	if err != nil {
		return err
	}
	address, _ := netlink.ParseAddr(IP)
	netlink.AddrAdd(nllink, address)
	return nil
}

func ParseIPAddress(IP string) (*netlink.Addr, error) {
	address, err := netlink.ParseAddr(IP)
	if err != nil {
		log.Printf("Error in parsing IP address in ParseIPAddress(%s): %s", IP, err)
	}
	return address, err
}
