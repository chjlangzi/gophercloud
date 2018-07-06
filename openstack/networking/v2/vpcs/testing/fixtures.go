package testing

import (
	"github.com/chjlangzi/gophercloud/openstack/networking/v2/vpcs"
)

const ListResponse = `
{
    "vpcs": [
        {
            "status": "ACTIVE",
            "network_ids": [
                "54d6f61d-db07-451c-9ab3-b9609b6b6f0b"
            ],
            "name": "public",
            "tenant_id": "4fd44f30292945e481c7b8a0c8908869",
            "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
            "description": "Network1 description",
            "created_at": "2018-04-23",
            "is_default": false,
            "cidr": "172.17.0.0/16",
			"port_security_enabled": true
        },
        {
            "status": "ACTIVE",
            "network_ids": [
                "08eae331-0402-425a-923c-34f7cfe39c1b"
            ],
            "name": "private",
            "tenant_id": "26a7980765d0414dbc1fc1f88cdb7e6e",
            "id": "db193ab3-96e3-4cb3-8fc5-05f4296d0324",
            "description": "Network2 description",
            "created_at": "2018-05-23",
            "is_default": true,
            "cidr": "172.18.0.0/16",
			"port_security_enabled": false
        }
    ]
}`

const GetResponse = `
{
    "vpc": {
		"status": "ACTIVE",
		"network_ids": [
			"54d6f61d-db07-451c-9ab3-b9609b6b6f0b"
		],
		"name": "public",
		"tenant_id": "4fd44f30292945e481c7b8a0c8908869",
		"id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
		"description": "Network1 description",
		"created_at": "2018-04-23",
		"is_default": false,
		"cidr": "172.17.0.0/16",
		"port_security_enabled": true
    }
}`

const CreateRequest = `
{
    "vpc": {
        "name": "private",
        "is_default": true
    }
}`

const CreateResponse = `
{
    "vpc": {
		"status": "ACTIVE",
        "network_ids": [
            "08eae331-0402-425a-923c-34f7cfe39c1b"
        ],
        "name": "private",
        "tenant_id": "26a7980765d0414dbc1fc1f88cdb7e6e",
        "id": "db193ab3-96e3-4cb3-8fc5-05f4296d0324",
        "description": "Network2 description",
        "created_at": "2018-05-23",
        "is_default": true,
        "cidr": "172.18.0.0/16",
		"port_security_enabled": false
    }
}`

const CreatePortSecurityRequest = `
{
    "vpc": {
        "name": "private",
        "is_default": true,
        "port_security_enabled": false
    }
}`

const CreatePortSecurityResponse = `
{
    "vpc": {
        "status": "ACTIVE",
        "subnets": ["08eae331-0402-425a-923c-34f7cfe39c1b"],
        "name": "private",
        "admin_state_up": true,
        "tenant_id": "26a7980765d0414dbc1fc1f88cdb7e6e",
        "shared": false,
        "id": "db193ab3-96e3-4cb3-8fc5-05f4296d0324",
        "provider:segmentation_id": 9876543210,
        "provider:physical_network": null,
        "provider:network_type": "local",
        "port_security_enabled": false
    }
}`

const CreateOptionalFieldsRequest = `
{
  "vpc": {
      "name": "public",
      "is_default": true,
      "tenant_id": "12345"
  }
}`

const UpdateRequest = `
{
    "vpc": {
        "name": "new_network_name",
        "description": "description abcd"
    }
}`

const UpdateResponse = `
{
    "vpc": {
		"status": "ACTIVE",
		"network_ids": [
			"08eae331-0402-425a-923c-34f7cfe39c1b"
		],
		"name": "new_network_name",
		"tenant_id": "26a7980765d0414dbc1fc1f88cdb7e6e",
		"id": "4e8e5957-649f-477b-9e5b-f1f75b21c03c",
		"description": "description abcd",
		"created_at": "2018-05-23",
		"is_default": true,
		"cidr": "172.18.0.0/16",
		"port_security_enabled": false
    }
}`

const UpdatePortSecurityRequest = `
{
    "vpc": {
        "port_security_enabled": false
    }
}`

const UpdatePortSecurityResponse = `
{
    "vpc": {
    	"status": "ACTIVE",
		"network_ids": [
			"08eae331-0402-425a-923c-34f7cfe39c1b"
		],
		"name": "private",
		"tenant_id": "26a7980765d0414dbc1fc1f88cdb7e6e",
		"id": "4e8e5957-649f-477b-9e5b-f1f75b21c03c",
		"description": "description abcd",
		"created_at": "2018-05-23",
		"is_default": true,
		"cidr": "172.18.0.0/16",
		"port_security_enabled": false
    }
}`

var Network1 = vpcs.Vpc{
	Status:       "ACTIVE",
	NetworkIds:   []string{"54d6f61d-db07-451c-9ab3-b9609b6b6f0b"},
	Name:         "public",
	TenantID:     "4fd44f30292945e481c7b8a0c8908869",
	ID:           "d32019d3-bc6e-4319-9c1d-6722fc136a22",
	Description:  "Network1 description",
	CreatedAt:    "2018-04-23",
	IsDefault:    false,
	Cidr:         "172.17.0.0/16",
}

var Network2 = vpcs.Vpc{
	Status:       "ACTIVE",
	NetworkIds:   []string{"08eae331-0402-425a-923c-34f7cfe39c1b"},
	Name:         "private",
	TenantID:     "26a7980765d0414dbc1fc1f88cdb7e6e",
	ID:           "db193ab3-96e3-4cb3-8fc5-05f4296d0324",
	Description:  "Network2 description",
	CreatedAt:    "2018-05-23",
	IsDefault:    true,
	Cidr:         "172.18.0.0/16",
}

var ExpectedNetworkSlice = []vpcs.Vpc{Network1, Network2}
