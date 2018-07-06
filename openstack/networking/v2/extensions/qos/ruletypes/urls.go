package ruletypes

import "github.com/chjlangzi/gophercloud"

func listRuleTypesURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("qos", "rule-types")
}
