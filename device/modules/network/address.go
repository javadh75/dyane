package network

import (
	"log"

	"github.com/vishvananda/netlink"
)

func AddIPAddress(linkName string, IP string) error {
	nllink, err := GetLink(linkName)
	if err != nil {
		return err
	}
	address, err := netlink.ParseAddr(IP)
	if err != nil {
		return err
	}
	if err := netlink.AddrAdd(nllink, address); err != nil {
		log.Printf("Error in adding IP address in AddIPAddress(%s, %s): %s", linkName, IP, err)
		return err
	}
	return nil
}

func ParseIPAddress(IP string) (*netlink.Addr, error) {
	address, err := netlink.ParseAddr(IP)
	if err != nil {
		log.Printf("Error in parsing IP address in ParseIPAddress(%s): %s", IP, err)
	}
	return address, err
}

func DelIPAddress(linkName string, IP string) error {
	nllink, err := GetLink(linkName)
	if err != nil {
		return err
	}
	address, err := netlink.ParseAddr(IP)
	if err != nil {
		return err
	}
	if err := netlink.AddrDel(nllink, address); err != nil {
		log.Printf("Error in deleting IP address in DelIPAddress(%s, %s): %s", linkName, IP, err)
		return err
	}
	return nil
}
