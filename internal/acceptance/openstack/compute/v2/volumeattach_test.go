//go:build acceptance || compute || volumeattach
// +build acceptance compute volumeattach

package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/internal/acceptance/clients"
	bs "github.com/gophercloud/gophercloud/internal/acceptance/openstack/blockstorage/v2"
	"github.com/gophercloud/gophercloud/internal/acceptance/tools"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestVolumeAttachAttachment(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	volume, err := bs.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	defer bs.DeleteVolume(t, blockClient, volume)

	client.Microversion = "2.79"
	volumeAttachment, err := CreateVolumeAttachment(t, client, blockClient, server, volume)
	th.AssertNoErr(t, err)
	defer DeleteVolumeAttachment(t, client, blockClient, server, volumeAttachment)

	tools.PrintResource(t, volumeAttachment)

	th.AssertEquals(t, volumeAttachment.ServerID, server.ID)
}
