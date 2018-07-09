package vpcsubnets

import "github.com/chjlangzi/gophercloud"

func listURL(c *gophercloud.ServiceClient, vpcId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets.json")
}

func createURL(c *gophercloud.ServiceClient, vpcId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets.json")
}

func deleteURL(c *gophercloud.ServiceClient, vpcId string, subnetId string) string {
	return c.ServiceURL("vpcs", vpcId, "subnets", subnetId)
}
