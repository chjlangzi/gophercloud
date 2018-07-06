package services

import (
	"github.com/chjlangzi/gophercloud"
	"github.com/chjlangzi/gophercloud/pagination"
)

// List makes a request against the API to list services.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return ServicePage{pagination.SinglePageBase(r)}
	})
}
