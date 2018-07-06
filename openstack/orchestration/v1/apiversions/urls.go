package apiversions

import "github.com/chjlangzi/gophercloud"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
