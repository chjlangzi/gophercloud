package evacuate

import (
	"github.com/chjlangzi/gophercloud"
)

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
