// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	acmpca_sdkv2 "github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCertificate,
			TypeName: "aws_acmpca_certificate",
			Name:     "Certificate",
		},
		{
			Factory:  dataSourceCertificateAuthority,
			TypeName: "aws_acmpca_certificate_authority",
			Name:     "Certificate Authority",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCertificate,
			TypeName: "aws_acmpca_certificate",
			Name:     "Certificate",
		},
		{
			Factory:  resourceCertificateAuthority,
			TypeName: "aws_acmpca_certificate_authority",
			Name:     "Certificate Authority",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  resourceCertificateAuthorityCertificate,
			TypeName: "aws_acmpca_certificate_authority_certificate",
			Name:     "Certificate Authority Certificate",
		},
		{
			Factory:  resourcePermission,
			TypeName: "aws_acmpca_permission",
			Name:     "Permission",
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_acmpca_policy",
			Name:     "Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ACMPCA
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*acmpca_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return acmpca_sdkv2.NewFromConfig(cfg, func(o *acmpca_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
