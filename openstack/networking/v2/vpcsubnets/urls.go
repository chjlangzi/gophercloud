package vpcsubnets

import "github.com/chjlangzi/gophercloud"

func resourceURL(c *gophercloud.ServiceClient, vpcId string, subnetId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets", subnetId)
}

func rootURL(c *gophercloud.ServiceClient, vpcId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets")
}

func listURL(c *gophercloud.ServiceClient, vpcId string) string {
	return rootURL(c,vpcId)
}

func createURL(c *gophercloud.ServiceClient, vpcId string) string {
	return rootURL(c,vpcId)
}

func deleteURL(c *gophercloud.ServiceClient, vpcId string, subnetId string) string {
	return resourceURL(c,vpcId,subnetId)
}
