// +build acceptance clientconfig

package clientconfig

import (
	"testing"

	"github.com/cro8ox/gophercloud/openstack/compute/v2/servers"

	acc_compute "github.com/cro8ox/gophercloud/acceptance/openstack/compute/v2"
	acc_tools "github.com/cro8ox/gophercloud/acceptance/tools"

	cc "github.com/cro8ox/utils/openstack/clientconfig"
)

func TestServerCreateDestroy(t *testing.T) {
	// This will be populated by environment variables.
	clientOpts := &cc.ClientOpts{}

	client, err := cc.NewServiceClient("compute", clientOpts)
	if err != nil {
		t.Fatalf("Unable to create client: %v", err)
	}

	server, err := acc_compute.CreateServer(t, client)
	if err != nil {
		t.Fatalf("Unable to create server: %v", err)
	}
	defer acc_compute.DeleteServer(t, client, server)

	newServer, err := servers.Get(client, server.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get server %s: %v", server.ID, err)
	}

	acc_tools.PrintResource(t, newServer)
}
