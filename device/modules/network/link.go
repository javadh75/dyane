package network

import (
	"github.com/vishvananda/netlink"
)

type DyaneLink struct {
	Valid        bool   `json:"valid"`
	Type         string `json:"type"`
	Index        int    `json:"index"`
	Name         string `json:"name"`
	HardwareAddr string `json:"mac"`
	Flags        string `json:"flags"`
	RawFlags     uint32 `json:"raw flags"`
	MTU          int    `json:"mtu"`
	ParentIndex  int    `json:"parent index"`
	MasterIndex  int    `json:"master index"`
	Alias        string `json:"alias"`
}

func (dl *DyaneLink) GetData(nll netlink.Link) {
	dl.Type = nll.Type()

	nll_attrs := nll.Attrs()
	dl.Index = nll_attrs.Index
	dl.Name = nll_attrs.Name
	dl.HardwareAddr = nll_attrs.HardwareAddr.String()
	dl.Flags = nll_attrs.Flags.String()
	dl.RawFlags = uint32(nll_attrs.Flags)
	dl.MTU = nll_attrs.MTU
	dl.ParentIndex = nll_attrs.ParentIndex
	dl.MasterIndex = nll_attrs.MasterIndex
	dl.Alias = nll_attrs.Alias

	dl.Valid = true
}

func GetLink(name string) (DyaneLink, error) {
	var dyaneLink DyaneLink
	dyaneLink.Valid = false
	link, err := netlink.LinkByName(name)
	if err != nil {
		return dyaneLink, err
	}
	dyaneLink.GetData(link)

	return dyaneLink, nil
}

func GetAllLinks() ([]DyaneLink, error) {
	var dyaneLinks []DyaneLink
	links, err := netlink.LinkList()
	if err != nil {
		return dyaneLinks, err
	}

	for _, link := range links {
		var dyaneLink DyaneLink
		dyaneLink.GetData(link)
		dyaneLinks = append(dyaneLinks, dyaneLink)
	}

	return dyaneLinks, nil
}
