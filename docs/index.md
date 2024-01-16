---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pcluster Provider"
subcategory: ""
description: |-
  AWS ParallelCluster is an AWS supported open source cluster management tool that helps you to deploy and manage high performance computing (HPC) clusters in the AWS Cloud. It automatically sets up the required compute resources, scheduler, and shared filesystem. You can use AWS ParallelCluster with AWS Batch and Slurm schedulers.
  With AWS ParallelCluster, you can quickly build and deploy proof of concept and production HPC compute environments. You can also build and deploy a high level workflow on top of AWS ParallelCluster, such as a genomics portal that automates an entire DNA sequencing workflow.
---

# pcluster Provider

AWS ParallelCluster is an AWS supported open source cluster management tool that helps you to deploy and manage high performance computing (HPC) clusters in the AWS Cloud. It automatically sets up the required compute resources, scheduler, and shared filesystem. You can use AWS ParallelCluster with AWS Batch and Slurm schedulers.

With AWS ParallelCluster, you can quickly build and deploy proof of concept and production HPC compute environments. You can also build and deploy a high level workflow on top of AWS ParallelCluster, such as a genomics portal that automates an entire DNA sequencing workflow.

## Example Usage

```terraform
provider "pcluster" {
  role_arn = var.role_arn
  endpoint = var.endpoint
  profile  = var.profile
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_name` (String) The name of the ParallelCluster api. Used to retrieve the api endpoint if not given. Defaults to ParallelCluster.
- `aws_key` (String) The aws key used for deploying resources and query data sources.
- `aws_secret` (String, Sensitive) The aws secret used for deploying resources and query data sources.
- `endpoint` (String) The endpoint of the ParallelCluster API. If unset will be autodetected.
- `profile` (String) The aws profile used for deploying resources and query data sources.
- `region` (String) The region used for deploying resources and query data sources.
- `role_arn` (String) The role used for deploying resources and query data sources.