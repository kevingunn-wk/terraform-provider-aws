// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newLogicallyAirGappedVaultResource,
			TypeName: "aws_backup_logically_air_gapped_vault",
			Name:     "Logically Air Gapped Vault",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newRestoreTestingPlanResource,
			TypeName: "aws_backup_restore_testing_plan",
			Name:     "Restore Testing Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newRestoreTestingSelectionResource,
			TypeName: "aws_backup_restore_testing_selection",
			Name:     "Restore Testing Plan Selection",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceFramework,
			TypeName: "aws_backup_framework",
			Name:     "Framework",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourcePlan,
			TypeName: "aws_backup_plan",
			Name:     "Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceReportPlan,
			TypeName: "aws_backup_report_plan",
			Name:     "Report Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceSelection,
			TypeName: "aws_backup_selection",
			Name:     "Selection",
		},
		{
			Factory:  dataSourceVault,
			TypeName: "aws_backup_vault",
			Name:     "Vault",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceFramework,
			TypeName: "aws_backup_framework",
			Name:     "Framework",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGlobalSettings,
			TypeName: "aws_backup_global_settings",
			Name:     "Global Settings",
		},
		{
			Factory:  resourcePlan,
			TypeName: "aws_backup_plan",
			Name:     "Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRegionSettings,
			TypeName: "aws_backup_region_settings",
			Name:     "Region Settings",
		},
		{
			Factory:  resourceReportPlan,
			TypeName: "aws_backup_report_plan",
			Name:     "Report Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSelection,
			TypeName: "aws_backup_selection",
			Name:     "Selection",
		},
		{
			Factory:  resourceVault,
			TypeName: "aws_backup_vault",
			Name:     "Vault",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVaultLockConfiguration,
			TypeName: "aws_backup_vault_lock_configuration",
			Name:     "Vault Lock Configuration",
		},
		{
			Factory:  resourceVaultNotifications,
			TypeName: "aws_backup_vault_notifications",
			Name:     "Vault Notifications",
		},
		{
			Factory:  resourceVaultPolicy,
			TypeName: "aws_backup_vault_policy",
			Name:     "Vault Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Backup
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*backup.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*backup.Options){
		backup.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return backup.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*backup.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*backup.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *backup.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*backup.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
