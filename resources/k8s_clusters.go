// Code generated by cq-provider-yandex resource generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func K8SClusters() *schema.Table {
	return &schema.Table{
		Name:         "yandex_k8s_clusters",
		Resolver:     fetchK8SClusters,
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
				Description: "Name of the Kubernetes cluster.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the Kubernetes cluster. 0-256 characters long.",
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
				Description: "Status of the Kubernetes cluster.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "health",
				Type:        schema.TypeString,
				Description: "Health of the Kubernetes cluster.",
				Resolver:    client.EnumPathResolver("Health"),
			},
			{
				Name:        "network_id",
				Type:        schema.TypeString,
				Description: "ID of the network the Kubernetes cluster belongs to.",
				Resolver:    schema.PathResolver("NetworkId"),
			},
			{
				Name:        "master_master_type_zonal_master_zone_id",
				Type:        schema.TypeString,
				Description: "ID of the availability zone where the master resides.",
				Resolver:    schema.PathResolver("Master.MasterType.ZonalMaster.ZoneId"),
			},
			{
				Name:        "master_master_type_zonal_master_internal_v_4_address",
				Type:        schema.TypeString,
				Description: "IPv4 internal network address that is assigned to the master.",
				Resolver:    schema.PathResolver("Master.MasterType.ZonalMaster.InternalV4Address"),
			},
			{
				Name:        "master_master_type_zonal_master_external_v_4_address",
				Type:        schema.TypeString,
				Description: "IPv4 external network address that is assigned to the master.",
				Resolver:    schema.PathResolver("Master.MasterType.ZonalMaster.ExternalV4Address"),
			},
			{
				Name:        "master_master_type_regional_master_region_id",
				Type:        schema.TypeString,
				Description: "ID of the region where the master resides.",
				Resolver:    schema.PathResolver("Master.MasterType.RegionalMaster.RegionId"),
			},
			{
				Name:        "master_master_type_regional_master_internal_v_4_address",
				Type:        schema.TypeString,
				Description: "IPv4 internal network address that is assigned to the master.",
				Resolver:    schema.PathResolver("Master.MasterType.RegionalMaster.InternalV4Address"),
			},
			{
				Name:        "master_master_type_regional_master_external_v_4_address",
				Type:        schema.TypeString,
				Description: "IPv4 external network address that is assigned to the master.",
				Resolver:    schema.PathResolver("Master.MasterType.RegionalMaster.ExternalV4Address"),
			},
			{
				Name:        "master_version",
				Type:        schema.TypeString,
				Description: "Version of Kubernetes components that runs on the master.",
				Resolver:    schema.PathResolver("Master.Version"),
			},
			{
				Name:        "master_endpoints_internal_v_4_endpoint",
				Type:        schema.TypeString,
				Description: "Internal endpoint that can be used to connect to the master from cloud networks.",
				Resolver:    schema.PathResolver("Master.Endpoints.InternalV4Endpoint"),
			},
			{
				Name:        "master_endpoints_external_v_4_endpoint",
				Type:        schema.TypeString,
				Description: "External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).",
				Resolver:    schema.PathResolver("Master.Endpoints.ExternalV4Endpoint"),
			},
			{
				Name:        "master_master_auth_cluster_ca_certificate",
				Type:        schema.TypeString,
				Description: "PEM-encoded public certificate that is the root of trust for the Kubernetes cluster.",
				Resolver:    schema.PathResolver("Master.MasterAuth.ClusterCaCertificate"),
			},
			{
				Name:        "master_version_info_current_version",
				Type:        schema.TypeString,
				Description: "Current Kubernetes version, format: major.minor (e.g. 1.15).",
				Resolver:    schema.PathResolver("Master.VersionInfo.CurrentVersion"),
			},
			{
				Name:        "master_version_info_new_revision_available",
				Type:        schema.TypeBool,
				Description: "Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well\n as some internal component updates - new features or bug fixes in Yandex specific\n components either on the master or nodes.",
				Resolver:    schema.PathResolver("Master.VersionInfo.NewRevisionAvailable"),
			},
			{
				Name:        "master_version_info_new_revision_summary",
				Type:        schema.TypeString,
				Description: "Description of the changes to be applied when updating to the latest\n revision. Empty if new_revision_available is false.",
				Resolver:    schema.PathResolver("Master.VersionInfo.NewRevisionSummary"),
			},
			{
				Name:        "master_version_info_version_deprecated",
				Type:        schema.TypeBool,
				Description: "The current version is on the deprecation schedule, component (master or node group)\n should be upgraded.",
				Resolver:    schema.PathResolver("Master.VersionInfo.VersionDeprecated"),
			},
			{
				Name:        "master_security_group_ids",
				Type:        schema.TypeStringArray,
				Description: "Master security groups.",
				Resolver:    schema.PathResolver("Master.SecurityGroupIds"),
			},
			{
				Name:        "ip_allocation_policy_cluster_ipv_4_cidr_block",
				Type:        schema.TypeString,
				Description: "CIDR block. IP range for allocating pod addresses.\n\n It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be\n set up for this CIDR blocks in node subnets.",
				Resolver:    schema.PathResolver("IpAllocationPolicy.ClusterIpv4CidrBlock"),
			},
			{
				Name:        "ip_allocation_policy_node_ipv_4_cidr_mask_size",
				Type:        schema.TypeBigInt,
				Description: "Size of the masks that are assigned for each node in the cluster.\n\n If not specified, 24 is used.",
				Resolver:    schema.PathResolver("IpAllocationPolicy.NodeIpv4CidrMaskSize"),
			},
			{
				Name:        "ip_allocation_policy_service_ipv_4_cidr_block",
				Type:        schema.TypeString,
				Description: "CIDR block. IP range Kubernetes service Kubernetes cluster IP addresses will be allocated from.\n\n It should not overlap with any subnet in the network the Kubernetes cluster located in.",
				Resolver:    schema.PathResolver("IpAllocationPolicy.ServiceIpv4CidrBlock"),
			},
			{
				Name:        "ip_allocation_policy_cluster_ipv_6_cidr_block",
				Type:        schema.TypeString,
				Description: "IPv6 range for allocating pod IP addresses.",
				Resolver:    schema.PathResolver("IpAllocationPolicy.ClusterIpv6CidrBlock"),
			},
			{
				Name:        "ip_allocation_policy_service_ipv_6_cidr_block",
				Type:        schema.TypeString,
				Description: "IPv6 range for allocating Kubernetes service IP addresses",
				Resolver:    schema.PathResolver("IpAllocationPolicy.ServiceIpv6CidrBlock"),
			},
			{
				Name:        "internet_gateway_gateway_ipv_4_address",
				Type:        schema.TypeString,
				Description: "Gateway IPv4 address.",
				Resolver:    schema.PathResolver("InternetGateway.GatewayIpv4Address"),
			},
			{
				Name:        "service_account_id",
				Type:        schema.TypeString,
				Description: "Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster.",
				Resolver:    schema.PathResolver("ServiceAccountId"),
			},
			{
				Name:        "node_service_account_id",
				Type:        schema.TypeString,
				Description: "Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.",
				Resolver:    schema.PathResolver("NodeServiceAccountId"),
			},
			{
				Name:        "release_channel",
				Type:        schema.TypeString,
				Description: "When creating a Kubernetes cluster, you should specify one of three release channels. The release channel contains several Kubernetes versions.\n Channels differ in the set of available versions, the management of auto-updates, and the updates received.\n You can't change the channel once the Kubernetes cluster is created, you can only recreate the Kubernetes cluster and specify a new release channel.\n For more details see [documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/release-channels-and-updates).",
				Resolver:    client.EnumPathResolver("ReleaseChannel"),
			},
			{
				Name:        "network_policy_provider",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.EnumPathResolver("NetworkPolicy.Provider"),
			},
			{
				Name:        "kms_provider_key_id",
				Type:        schema.TypeString,
				Description: "KMS key ID for secrets encryption.\n To obtain a KMS key ID use a [yandex.cloud.kms.v1.SymmetricKeyService.List] request.",
				Resolver:    schema.PathResolver("KmsProvider.KeyId"),
			},
			{
				Name:        "log_group_id",
				Type:        schema.TypeString,
				Description: "Log group where cluster stores cluster system logs, like audit, events, or controlplane logs.",
				Resolver:    schema.PathResolver("LogGroupId"),
			},
			{
				Name:        "network_implementation_cilium_routing_mode",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.EnumPathResolver("NetworkImplementation.Cilium.RoutingMode"),
			},
		},
	}

}

func fetchK8SClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &k8s.ListClustersRequest{FolderId: c.MultiplexedResourceId}
	it := c.Services.K8S.Cluster().ClusterIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
