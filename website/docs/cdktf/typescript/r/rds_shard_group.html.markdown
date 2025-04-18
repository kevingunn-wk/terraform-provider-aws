---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_rds_shard_group"
description: |-
  Terraform resource for managing an Amazon Aurora Limitless Database DB shard group.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_rds_shard_group

Terraform resource for managing an Amazon Aurora Limitless Database DB shard group

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RdsShardGroup } from "./.gen/providers/aws/";
import { RdsCluster } from "./.gen/providers/aws/rds-cluster";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new RdsCluster(this, "example", {
      clusterIdentifier: "example-limitless-cluster",
      cluster_scalability_type: "limitless",
      enabledCloudwatchLogsExports: ["postgresql"],
      engine: "aurora-postgresql",
      engineMode: "",
      engineVersion: "16.6-limitless",
      masterPassword: "must_be_eight_characters",
      masterUsername: "foo",
      monitoringInterval: 5,
      monitoringRoleArn: Token.asString(awsIamRoleExample.arn),
      performanceInsightsEnabled: true,
      performanceInsightsRetentionPeriod: 31,
      storageType: "aurora-iopt1",
    });
    const awsRdsShardGroupExample = new RdsShardGroup(this, "example_1", {
      db_cluster_identifier: example.id,
      db_shard_group_identifier: "example-shard-group",
      max_acu: 1200,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRdsShardGroupExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

For more detailed documentation about each argument, refer to the [AWS official documentation](https://docs.aws.amazon.com/cli/latest/reference/rds/create-integration.html).

This resource supports the following arguments:

* `compute_redundancy` - (Optional) Specifies whether to create standby DB shard groups for the DB shard group. Valid values are:
    * `0` - Creates a DB shard group without a standby DB shard group. This is the default value.
    * `1` - Creates a DB shard group with a standby DB shard group in a different Availability Zone (AZ).
    * `2` - Creates a DB shard group with two standby DB shard groups in two different AZs.
* `dbClusterIdentifier` - (Required) The name of the primary DB cluster for the DB shard group.
* `db_shard_group_identifier` - (Required) The name of the DB shard group.
* `max_acu` - (Required) The maximum capacity of the DB shard group in Aurora capacity units (ACUs).
* `min_acu` - (Optional) The minimum capacity of the DB shard group in Aurora capacity units (ACUs).
* `publiclyAccessible` - (Optional) Indicates whether the DB shard group is publicly accessible.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the shard group.
* `db_shard_group_resource_id` - The AWS Region-unique, immutable identifier for the DB shard group.
* `endpoint` - The connection endpoint for the DB shard group.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `45m`)
* `update` - (Default `45m`)
* `delete` - (Default `45m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import shard group using the `db_shard_group_identifier`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RdsShardGroup } from "./.gen/providers/aws/";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RdsShardGroup.generateConfigForImport(
      this,
      "example",
      "example-shard-group"
    );
  }
}

```

Using `terraform import`, import shard group using the `db_shard_group_identifier`. For example:

```console
% terraform import aws_rds_shard_group.example example-shard-group
```

<!-- cache-key: cdktf-0.20.8 input-d1de69af0e1ebc3ee02a4625d081d7f69a3d90a58ef24b466911dbb74ddeb57c -->