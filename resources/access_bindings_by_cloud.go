// Code generated by yandex cloud generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
)

func AccessBindingsByCloud() *schema.Table {
	return &schema.Table{
		Name:        "yandex_access_bindings_by_cloud",
		Resolver:    fetchAccessBindingsByCloud,
		Multiplex:   client.MultiplexBy(client.Clouds),
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceId"),
			},
			{
				Name:     "role_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleId"),
			},
			{
				Name:     "subject_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject.Id"),
			},
			{
				Name:     "subject_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subject.Type"),
			},
		},
	}
}

type CloudAccessBinding struct {
	*access.AccessBinding
	ResourceId string
}

func fetchAccessBindingsByCloud(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	resourceClient := c.Services.ResourceManager.Cloud()

	resp, err := resourceClient.ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
		ResourceId: c.MultiplexedResourceId,
	})
	if err != nil {
		return err
	}

	for {
		for _, value := range resp.GetAccessBindings() {
			res <- CloudAccessBinding{
				AccessBinding: value,
				ResourceId:    c.MultiplexedResourceId,
			}
		}

		if resp.GetNextPageToken() == "" {
			break
		}

		resp, err = resourceClient.ListAccessBindings(ctx, &access.ListAccessBindingsRequest{
			ResourceId: c.MultiplexedResourceId,
			PageToken:  resp.GetNextPageToken(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
