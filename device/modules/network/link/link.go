package link

import (
	"log"

	"github.com/vishvananda/netlink"
)

func DelByName(name string) (bool, error) {
	link, err := GetLinkByName(name)
	if err != nil {
		return false, nil
	}
	if err = netlink.LinkDel(link); err != nil {
		log.Printf("Error in deleting link in DelByName(%s): %s", name, err)
		return false, nil
	}
	return true, nil
}

func DelByIndex(index int) (bool, error) {
	link, err := GetLinkByIndex(index)
	if err != nil {
		return false, nil
	}
	if err = netlink.LinkDel(link); err != nil {
		log.Printf("Error in deleting link in DelByIndex(%d): %s", index, err)
		return false, nil
	}
	return true, nil
}

func GetLinkByName(name string) (netlink.Link, error) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		log.Printf("Error in getting link in GetLinkByName(%s): %s", name, err)
		return nil, err
	}
	return link, nil
}

func GetLinkByIndex(index int) (netlink.Link, error) {
	link, err := netlink.LinkByIndex(index)
	if err != nil {
		log.Printf("Error in getting link in GetLinkByIndex(%d): %s", index, err)
		return nil, err
	}
	return link, nil
}

func GetAllLinks() ([]netlink.Link, error) {
	links, err := netlink.LinkList()
	if err != nil {
		log.Printf("Error in getting links in GetAllLinks(): %s", err)
		return nil, err
	}
	return links, nil
}
