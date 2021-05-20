package vlan

import (
	"log"

	"github.com/vishvananda/netlink"
)

// AddVlanByParentIndex adds a VLan with its parent link index and it configs
func AddVlanByParentIndex(name string, parentIndex int, vlanID int, vlanProtocol netlink.VlanProtocol) (bool, error) {
	vlanAttrs := netlink.NewLinkAttrs()
	vlanAttrs.Name = name
	vlanAttrs.ParentIndex = parentIndex

	vlan := &netlink.Vlan{LinkAttrs: vlanAttrs, VlanId: vlanID, VlanProtocol: vlanProtocol}
	if err := netlink.LinkAdd(vlan); err != nil {
		log.Printf("Error in adding link in AddVlanByParentIndex(%s, %d, %d, %v): %s", name, parentIndex, vlanID, vlanProtocol, err)
		return false, err
	}
	return true, nil
}

// AddVlanByParentName adds a VLan with its parent link name and it configs
func AddVlanByParentName(name string, parentName string, vlanID int, vlanProtocol netlink.VlanProtocol) (bool, error) {
	parentLink, err := netlink.LinkByName(parentName)
	if err != nil {
		log.Printf("Error in getting parent link in AddVlanByParentName(%s, %s, %d, %v): %s", name, parentName, vlanID, vlanProtocol, err)
		return false, err
	}

	vlanAttrs := netlink.NewLinkAttrs()
	vlanAttrs.Name = name
	vlanAttrs.ParentIndex = parentLink.Attrs().Index

	vlan := &netlink.Vlan{LinkAttrs: vlanAttrs, VlanId: vlanID, VlanProtocol: vlanProtocol}
	if err := netlink.LinkAdd(vlan); err != nil {
		log.Printf("Error in adding link in AddVlanByParentName(%s, %s, %d, %v): %s", name, parentName, vlanID, vlanProtocol, err)
		return false, err
	}
	return true, nil
}
