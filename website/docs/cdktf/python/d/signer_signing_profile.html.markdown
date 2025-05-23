---
subcategory: "Signer"
layout: "aws"
page_title: "AWS: aws_signer_signing_profile"
description: |-
  Provides a Signer Signing Profile data source.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_signer_signing_profile

Provides information about a Signer Signing Profile.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_signer_signing_profile import DataAwsSignerSigningProfile
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSignerSigningProfile(self, "production_signing_profile",
            name="prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK"
        )
```

## Argument Reference

This data source supports the following arguments:

* `name` - (Required) Name of the target signing profile.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN for the signing profile.
* `platform_display_name` - A human-readable name for the signing platform associated with the signing profile.
* `platform_id` - ID of the platform that is used by the target signing profile.
* `revocation_record` - Revocation information for a signing profile.
* `signature_validity_period` - The validity period for a signing job.
* `status` - Status of the target signing profile.
* `tags` - List of tags associated with the signing profile.
* `version` - Current version of the signing profile.
* `version_arn` - Signing profile ARN, including the profile version.

<!-- cache-key: cdktf-0.20.8 input-aba77f8dc958b5e0757ea2581dd2d80a47faca88c585f688dbc1436b3712ace6 -->