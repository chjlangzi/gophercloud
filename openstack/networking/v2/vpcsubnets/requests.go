package vpcsubnets

import (
	"github.com/chjlangzi/gophercloud"
	"github.com/chjlangzi/gophercloud/pagination"
	"fmt"
	"encoding/json"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSubnetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the subnet attributes you want to see returned. SortKey allows you to sort
// by a particular subnet attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {

}

// ToSubnetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSubnetListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// subnets. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those subnets that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *gophercloud.ServiceClient,vpcId string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c,vpcId)
	if opts != nil {
		query, err := opts.ToSubnetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return SubnetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// List request.
type CreateOptsBuilder interface {
	ToSubnetCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents the attributes used when creating a new subnet.
type CreateOpts struct {

	// CIDR is the address CIDR of the subnet.
	CIDR string `json:"cidr,omitempty"`

	// Name is a human-readable name of the subnet.
	Name string `json:"name,omitempty"`

	// GatewayIP sets gateway information for the subnet. Setting to nil will
	// cause a default gateway to automatically be created. Setting to an empty
	// string will cause the subnet to be created with no gateway. Setting to
	// an explicit address will set that address as the gateway.
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// IPVersion is the IP version for the subnet.
	IPVersion gophercloud.IPVersion `json:"ip_version,omitempty"`

	// DNSNameservers are the nameservers to be set via DHCP.
	DNSNameservers []string `json:"dns_nameservers,omitempty"`

	Vpc *bool `json:"vpc"`
}

// ToSubnetCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "subnet")
	if err != nil {
		return nil, err
	}

	if m := b["subnet"].(map[string]interface{}); m["gateway_ip"] == "" {
		m["gateway_ip"] = nil
	}

	return b, nil
}

// Create accepts a CreateOpts struct and creates a new subnet using the values
// provided. You must remember to provide a valid NetworkID, CIDR and IP
// version.
func Create(c *gophercloud.ServiceClient,vpcId string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSubnetCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	jsonBytes, _ := json.Marshal(b)
	fmt.Println("subnet>>>>%s:", string(jsonBytes))
	fmt.Println("url>>%s,",createURL(c,vpcId))
	_, r.Err = c.Post(createURL(c,vpcId), b, &r.Body, nil)
	return
}

// Delete accepts a unique ID and deletes the subnet associated with it.
func Delete(c *gophercloud.ServiceClient, vpcId string, subnetId string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, vpcId, subnetId), nil)
	return
}

// IDFromName is a convenience function that returns a subnet's ID,
// given its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, "",nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractSubnets(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		return "", gophercloud.ErrResourceNotFound{Name: name, ResourceType: "subnet"}
	case 1:
		return id, nil
	default:
		return "", gophercloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "subnet"}
	}
}
