package vpcs

import (
	"github.com/chjlangzi/gophercloud"
	"github.com/chjlangzi/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNetworkListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the vpc attributes you want to see returned. SortKey allows you to sort
// by a particular vpc attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Status       string `q:"status"`
	Name         string `q:"name"`
	TenantID     string `q:"tenant_id"`
	ID           string `q:"id"`
	SortKey      string `q:"sort_key"`
	SortDir      string `q:"sort_dir"`
	Cidr         string `q:"cidr"`
	Description  string `q:"description"`
	IsDefault    *bool  `q:"is_default"`
	Fields       []string `q:"fields"`
}

// ToVpcListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNetworkListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// vpcs. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToNetworkListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return VpcPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific vpc based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVpcCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a vpc.
type CreateOpts struct {
	Vpc                   string   `json:"vpc,omitempty"`
	Name                  string   `json:"name,omitempty"`
	TenantID              string   `json:"tenant_id,omitempty"`
	Description           string   `json:"description,omitempty"`
	Cidr                  string   `json:"cidr,omitempty"`
	IsDefault             *bool    `json:"is_default,omitempty"`
}

// ToVpcCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToVpcCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "vpc")
}

// Create accepts a CreateOpts struct and creates a new vpc using the values
// provided. This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// vpc. An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVpcCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVpcUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a vpc.
type UpdateOpts struct {
	Name         string  `json:"name,omitempty"`
	Description  string  `json:"description,omitempty"`
	Cidr         string  `json:"cidr,omitempty"`
}

// ToVpcUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToVpcUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "vpc")
}

// Update accepts a UpdateOpts struct and updates an existing vpc using the
// values provided. For more information, see the Create function.
func Update(c *gophercloud.ServiceClient, vpcID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVpcUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, vpcID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

// Delete accepts a unique ID and deletes the vpc associated with it.
func Delete(c *gophercloud.ServiceClient, vpcID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, vpcID), nil)
	return
}

// IDFromName is a convenience function that returns a vpc's ID, given
// its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractVpcs(pages)
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
		return "", gophercloud.ErrResourceNotFound{Name: name, ResourceType: "vpc"}
	case 1:
		return id, nil
	default:
		return "", gophercloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "vpc"}
	}
}
