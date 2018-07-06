// +build acceptance networking lbaas_v2 listeners

package lbaas_v2

import (
	"testing"

	"github.com/chjlangzi/gophercloud/acceptance/clients"
	"github.com/chjlangzi/gophercloud/acceptance/tools"
	"github.com/chjlangzi/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
)

func TestListenersList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	allPages, err := listeners.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list listeners: %v", err)
	}

	allListeners, err := listeners.ExtractListeners(allPages)
	if err != nil {
		t.Fatalf("Unable to extract listeners: %v", err)
	}

	for _, listener := range allListeners {
		tools.PrintResource(t, listener)
	}
}
