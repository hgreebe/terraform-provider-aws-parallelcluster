---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pcluster_cluster Data Source - pcluster"
subcategory: ""
description: |-
  Get detailed information about an existing cluster.
---

# pcluster_cluster (Data Source)

Get detailed information about an existing cluster.

## Example Usage

```terraform
data "pcluster_cluster" "example" {
  cluster_name = var.cluster_name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) The name of the cluster.

### Optional

- `filters` (List of Map of String) Filter the log streams.
- `region` (String) The AWS Region that the cluster is in.

### Read-Only

- `cluster` (Object) the cluster. (see [below for nested schema](#nestedatt--cluster))
- `log_streams` (List of Map of String) List of logstreams.
- `stack_events` (List of Map of String) Events that are associated with the stack for an image build.

<a id="nestedatt--cluster"></a>
### Nested Schema for `cluster`

Read-Only:

- `cloudFormationStackStatus` (String)
- `cloudformationStackArn` (String)
- `clusterConfiguration` (String)
- `clusterName` (String)
- `clusterStatus` (String)
- `computeFleetStatus` (String)
- `creationTime` (String)
- `failures` (List of Map of String)
- `headNode` (Map of String)
- `lastUpdatedTime` (String)
- `loginNodes` (Object) (see [below for nested schema](#nestedobjatt--cluster--loginNodes))
- `region` (String)
- `scheduler` (Object) (see [below for nested schema](#nestedobjatt--cluster--scheduler))
- `tags` (List of Map of String)
- `version` (String)

<a id="nestedobjatt--cluster--loginNodes"></a>
### Nested Schema for `cluster.loginNodes`

Read-Only:

- `address` (String)
- `healthyNodes` (String)
- `scheme` (String)
- `status` (String)
- `unhealthyNodes` (String)


<a id="nestedobjatt--cluster--scheduler"></a>
### Nested Schema for `cluster.scheduler`

Read-Only:

- `metadata` (Object) (see [below for nested schema](#nestedobjatt--cluster--scheduler--metadata))
- `type` (String)

<a id="nestedobjatt--cluster--scheduler--metadata"></a>
### Nested Schema for `cluster.scheduler.metadata`

Read-Only:

- `name` (String)
- `version` (String)