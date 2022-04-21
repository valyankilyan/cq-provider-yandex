// Code generated by cq-provider-yandex resource generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/thoas/go-funk"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
)

func ServerlessApiGateways() *schema.Table {
	return &schema.Table{
		Name:         "yandex_serverless_api_gateways",
		Resolver:     fetchServerlessApiGateways,
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
				Description: "Name of the API gateway. The name is unique within the folder.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the API gateway.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Status of the API gateway.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "domain",
				Type:        schema.TypeString,
				Description: "Default domain for the API gateway. Generated at creation time.",
				Resolver:    schema.PathResolver("Domain"),
			},
			{
				Name:        "log_group_id",
				Type:        schema.TypeString,
				Description: "ID of the log group for the API gateway.",
				Resolver:    schema.PathResolver("LogGroupId"),
			},
			{
				Name:        "connectivity_network_id",
				Type:        schema.TypeString,
				Description: "Network the gateway will have access to.\n It's essential to specify network with subnets in all availability zones.",
				Resolver:    schema.PathResolver("Connectivity.NetworkId"),
			},
			{
				Name:        "connectivity_subnet_id",
				Type:        schema.TypeStringArray,
				Description: "Complete list of subnets (from the same network) the gateway can be attached to.\n It's essential to specify at least one subnet for each availability zones.",
				Resolver:    schema.PathResolver("Connectivity.SubnetId"),
			},
		},

		Relations: []*schema.Table{
			{
				Name:        "yandex_apigateway_api_gateway_attached_domains",
				Resolver:    fetchApiGatewayApiGatewayAttachedDomains,
				Multiplex:   client.EmptyMultiplex,
				IgnoreError: client.IgnoreErrorHandler,
				Columns: []schema.Column{
					{
						Name:        "api_gateway_cq_id",
						Type:        schema.TypeUUID,
						Description: "cq_id of parent api_gateway",
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_gateway_id",
						Type:        schema.TypeString,
						Description: "id of parent api_gateway",
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "domain_id",
						Type:        schema.TypeString,
						Description: "ID of the domain.",
						Resolver:    schema.PathResolver("DomainId"),
					},
					{
						Name:        "certificate_id",
						Type:        schema.TypeString,
						Description: "ID of the domain certificate.",
						Resolver:    schema.PathResolver("CertificateId"),
					},
					{
						Name:        "enabled",
						Type:        schema.TypeBool,
						Description: "Enabling flag.",
						Resolver:    schema.PathResolver("Enabled"),
					},
					{
						Name:        "domain",
						Type:        schema.TypeString,
						Description: "Name of the domain.",
						Resolver:    schema.PathResolver("Domain"),
					},
				},
			},
		},
	}

}

func fetchServerlessApiGateways(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &apigateway.ListApiGatewayRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.ApiGateway.ApiGateway().ApiGatewayIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}

func fetchApiGatewayApiGatewayAttachedDomains(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	values := funk.Get(parent.Item, "AttachedDomains")

	if funk.IsIteratee(values) {
		funk.ForEach(values, func(value interface{}) {
			res <- value
		})
	}

	return nil
}
