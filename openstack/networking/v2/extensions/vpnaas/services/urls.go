package services

import "github.com/chjlangzi/gophercloud"

const (
	rootPath     = "vpn"
	resourcePath = "vpnservices"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
