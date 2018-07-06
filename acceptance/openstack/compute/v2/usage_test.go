// +build acceptance compute usage

package v2

import (
	"strings"
	"testing"
	"time"

	"github.com/chjlangzi/gophercloud/acceptance/clients"
	"github.com/chjlangzi/gophercloud/acceptance/tools"
	"github.com/chjlangzi/gophercloud/openstack/compute/v2/extensions/usage"
	th "github.com/chjlangzi/gophercloud/testhelper"
)

func TestUsageSingleTenant(t *testing.T) {
	t.Skip("This is not passing in OpenLab. Works locally")

	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	DeleteServer(t, client, server)

	endpointParts := strings.Split(client.Endpoint, "/")
	tenantID := endpointParts[4]

	end := time.Now()
	start := end.AddDate(0, -1, 0)
	opts := usage.SingleTenantOpts{
		Start: &start,
		End:   &end,
	}

	page, err := usage.SingleTenant(client, tenantID, opts).AllPages()
	th.AssertNoErr(t, err)

	tenantUsage, err := usage.ExtractSingleTenant(page)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, tenantUsage)

	if tenantUsage.TotalHours == 0 {
		t.Fatalf("TotalHours should not be 0")
	}
}
