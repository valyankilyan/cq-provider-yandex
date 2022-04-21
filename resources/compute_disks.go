// Code generated by cq-provider-yandex resource generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func ComputeDisks() *schema.Table {
	return &schema.Table{
		Name:         "yandex_compute_disks",
		Resolver:     fetchComputeDisks,
		Multiplex:    client.MultiplexBy(client.Folders),
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the resource.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{NotNull: !false, Unique: true},
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Description: "ID of the folder that the resource belongs to.",
				Resolver:    client.ResolveFolderID,
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the disk. 1-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the disk. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "type_id",
				Type:        schema.TypeString,
				Description: "ID of the disk type.",
				Resolver:    schema.PathResolver("TypeId"),
			},
			{
				Name:        "zone_id",
				Type:        schema.TypeString,
				Description: "ID of the availability zone where the disk resides.",
				Resolver:    schema.PathResolver("ZoneId"),
			},
			{
				Name:        "size",
				Type:        schema.TypeBigInt,
				Description: "Size of the disk, specified in bytes.",
				Resolver:    schema.PathResolver("Size"),
			},
			{
				Name:        "block_size",
				Type:        schema.TypeBigInt,
				Description: "Block size of the disk, specified in bytes.",
				Resolver:    schema.PathResolver("BlockSize"),
			},
			{
				Name:        "product_ids",
				Type:        schema.TypeStringArray,
				Description: "License IDs that indicate which licenses are attached to this resource.\n License IDs are used to calculate additional charges for the use of the virtual machine.\n\n The correct license ID is generated by Yandex Cloud. IDs are inherited by new resources created from this resource.\n\n If you know the license IDs, specify them when you create the image.\n For example, if you create a disk image using a third-party utility and load it into Yandex Object Storage, the license IDs will be lost.\n You can specify them in the [yandex.cloud.compute.v1.ImageService.Create] request.",
				Resolver:    schema.PathResolver("ProductIds"),
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Current status of the disk.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "source_source_image_id",
				Type:        schema.TypeString,
				Description: "ID of the image that was used for disk creation.",
				Resolver:    schema.PathResolver("Source.SourceImageId"),
			},
			{
				Name:        "source_source_snapshot_id",
				Type:        schema.TypeString,
				Description: "ID of the snapshot that was used for disk creation.",
				Resolver:    schema.PathResolver("Source.SourceSnapshotId"),
			},
			{
				Name:        "instance_ids",
				Type:        schema.TypeStringArray,
				Description: "Array of instances to which the disk is attached.",
				Resolver:    schema.PathResolver("InstanceIds"),
			},
			{
				Name:        "disk_placement_policy_placement_group_id",
				Type:        schema.TypeString,
				Description: "Placement group ID.",
				Resolver:    schema.PathResolver("DiskPlacementPolicy.PlacementGroupId"),
			},
		},
	}

}

func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &compute.ListDisksRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.Compute.Disk().DiskIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
