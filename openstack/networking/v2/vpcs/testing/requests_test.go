package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/chjlangzi/gophercloud/openstack/networking/v2/common"
	"github.com/chjlangzi/gophercloud/openstack/networking/v2/extensions/portsecurity"
	"github.com/chjlangzi/gophercloud/openstack/networking/v2/vpcs"
	"github.com/chjlangzi/gophercloud/pagination"
	th "github.com/chjlangzi/gophercloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	err := vpcs.List(client, vpcs.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := vpcs.ExtractVpcs(page)
		if err != nil {
			t.Errorf("Failed to extract networks: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedNetworkSlice, actual)

		return true, nil
	})

	if err != nil {
		fmt.Println("error>>%v",err)
	}
	fmt.Println("%d",count)
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListWithExtensions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	type networkWithExt struct {
		vpcs.Vpc
		portsecurity.PortSecurityExt
	}

	var allNetworks []networkWithExt

	allPages, err := vpcs.List(client, vpcs.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	err = vpcs.ExtractVpcsInto(allPages, &allNetworks)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, allNetworks[0].Status, "ACTIVE")
	th.AssertEquals(t, allNetworks[0].PortSecurityEnabled, true)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := vpcs.Get(fake.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &Network1, n)
}

func TestGetWithExtensions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	var networkWithExtensions struct {
		vpcs.Vpc
		portsecurity.PortSecurityExt
	}

	err := vpcs.Get(fake.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").ExtractInto(&networkWithExtensions)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, networkWithExtensions.Status, "ACTIVE")
	th.AssertEquals(t, networkWithExtensions.PortSecurityEnabled, true)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	iTrue := true
	options := vpcs.CreateOpts{Name: "private", IsDefault: &iTrue}
	n, err := vpcs.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Status, "ACTIVE")
	th.AssertDeepEquals(t, &Network2, n)
}

func TestCreateWithOptionalFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateOptionalFieldsRequest)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{}`)
	})

	iTrue := true
	options := vpcs.CreateOpts{
		Name:                  "public",
		IsDefault:              &iTrue,
		TenantID:              "12345",
	}
	_, err := vpcs.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs/4e8e5957-649f-477b-9e5b-f1f75b21c03c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	d := "description abcd"
	options := vpcs.UpdateOpts{Name: "new_network_name", Description: d}
	n, err := vpcs.Update(fake.ServiceClient(), "4e8e5957-649f-477b-9e5b-f1f75b21c03c", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Name, "new_network_name")
	th.AssertEquals(t, n.Description, d)
	th.AssertEquals(t, n.ID, "4e8e5957-649f-477b-9e5b-f1f75b21c03c")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs/4e8e5957-649f-477b-9e5b-f1f75b21c03c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := vpcs.Delete(fake.ServiceClient(), "4e8e5957-649f-477b-9e5b-f1f75b21c03c")
	th.AssertNoErr(t, res.Err)
}

func TestCreatePortSecurity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreatePortSecurityRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreatePortSecurityResponse)
	})

	var networkWithExtensions struct {
		vpcs.Vpc
		portsecurity.PortSecurityExt
	}

	iTrue := true
	iFalse := false
	networkCreateOpts := vpcs.CreateOpts{Name: "private", IsDefault: &iTrue}
	createOpts := portsecurity.VpcCreateOptsExt{
		CreateOptsBuilder:   networkCreateOpts,
		PortSecurityEnabled: &iFalse,
	}

	err := vpcs.Create(fake.ServiceClient(), createOpts).ExtractInto(&networkWithExtensions)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, networkWithExtensions.Status, "ACTIVE")
	th.AssertEquals(t, networkWithExtensions.PortSecurityEnabled, false)
}

func TestUpdatePortSecurity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/vpcs/4e8e5957-649f-477b-9e5b-f1f75b21c03c", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdatePortSecurityRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdatePortSecurityResponse)
	})

	var networkWithExtensions struct {
		vpcs.Vpc
		portsecurity.PortSecurityExt
	}

	iFalse := false
	vpcUpdateOpts := vpcs.UpdateOpts{}
	updateOpts := portsecurity.VpcUpdateOptsExt{
		UpdateOptsBuilder:   vpcUpdateOpts,
		PortSecurityEnabled: &iFalse,
	}

	err := vpcs.Update(fake.ServiceClient(), "4e8e5957-649f-477b-9e5b-f1f75b21c03c", updateOpts).ExtractInto(&networkWithExtensions)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, networkWithExtensions.Name, "private")
	th.AssertEquals(t, networkWithExtensions.IsDefault, true)
	th.AssertEquals(t, networkWithExtensions.ID, "4e8e5957-649f-477b-9e5b-f1f75b21c03c")
	th.AssertEquals(t, networkWithExtensions.PortSecurityEnabled, false)
}
