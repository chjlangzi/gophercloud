package vpcs

import (
	"github.com/chjlangzi/gophercloud"
	"github.com/chjlangzi/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a vpc resource.
func (r commonResult) Extract() (*Vpc, error) {
	var s Vpc
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "vpc")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a vpc.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a vpc.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a vpc.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

type Router struct{
	Id string `json:"id"`
}

// vpc represents, well, a vpc.
type Vpc struct {
	// UUID for the vpc
	ID string `json:"id"`

	// Human-readable name for the vpc. Might not be unique.
	Name string `json:"name"`

	// Indicates whether vpc is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional
	// values.
	Status string `json:"status"`

	// network_ids associated with this vpc.
	NetworkIds []string `json:"network_ids"`

	// routers associated with this vpc.
	Routers []Router `json:"routers"`

	// TenantID is the project owner of the vpc.
	TenantID string `json:"tenant_id"`

	// description for the vpc
	Description string `json:"description"`

	// createTime for the vpc
	CreatedAt string `json:"created_at"`

	// updateTime for the vpc
	UpdatedAt string `json:"updated_at"`

	// is_default for the vpc
	IsDefault bool `json:"is_default"`

	// cidr for the vpc
	Cidr string `json:"cidr"`
}

// vpcPage is the page returned by a pager when traversing over a
// collection of vpcs.
type VpcPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of vpcs has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r VpcPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"vpcs_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a vpcPage struct is empty.
func (r VpcPage) IsEmpty() (bool, error) {
	is, err := ExtractVpcs(r)
	return len(is) == 0, err
}

// ExtractVpcs accepts a Page struct, specifically a VpcPage struct,
// and extracts the elements into a slice of vpc structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractVpcs(r pagination.Page) ([]Vpc, error) {
	var s []Vpc
	err := ExtractVpcsInto(r, &s)
	return s, err
}

func ExtractVpcsInto(r pagination.Page, v interface{}) error {
	return r.(VpcPage).Result.ExtractIntoSlicePtr(v, "vpcs")
}
