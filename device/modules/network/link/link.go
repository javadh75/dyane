package vlan

import (
	"log"

	"github.com/vishvananda/netlink"
)

func DelByName(name string) (bool, error) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		log.Printf("Error in getting link in DelByName(%s): %s", name, err)
		return false, nil
	}
	if err = netlink.LinkDel(link); err != nil {
		log.Printf("Error in deleting link in DelByName(%s): %s", name, err)
		return false, nil
	}
	return true, nil
}

func DelByIndex(index int) (bool, error) {
	link, err := netlink.LinkByIndex(index)
	if err != nil {
		log.Printf("Error in getting link in DelByIndex(%d): %s", index, err)
		return false, nil
	}
	if err = netlink.LinkDel(link); err != nil {
		log.Printf("Error in deleting link in DelByIndex(%d): %s", index, err)
		return false, nil
	}
	return true, nil
}
