package link

import (
	"log"

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
