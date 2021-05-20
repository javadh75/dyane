package link

import (
	"log"

	"github.com/vishvananda/netlink"
)

// AddVlan adds a VLan with its parent link name and it configs
func AddVlan(name string, parentName string, vlanID int, vlanProtocol netlink.VlanProtocol) (bool, error) {
	parentLink, err := netlink.LinkByName(parentName)
	if err != nil {
		log.Printf("Error in getting parent link in AddVlan(%s, %s, %d, %v): %s", name, parentName, vlanID, vlanProtocol, err)
		return false, err
	}

	vlanAttrs := netlink.NewLinkAttrs()
	vlanAttrs.Name = name
	vlanAttrs.ParentIndex = parentLink.Attrs().Index

	vlan := &netlink.Vlan{LinkAttrs: vlanAttrs, VlanId: vlanID, VlanProtocol: vlanProtocol}
	if err := netlink.LinkAdd(vlan); err != nil {
		log.Printf("Error in adding link in AddVlan(%s, %s, %d, %v): %s", name, parentName, vlanID, vlanProtocol, err)
		return false, err
	}
	return true, nil
}
