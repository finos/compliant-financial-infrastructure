# EKS - Terraform

## Requirements

Before running the Terraform scripts, make sure to set the [IAM required permissions](./iam-required-permissions.md) first. You will have to remove the comments in the code if you decide to copy/paste them. Remember it is a best practice to use a role and assign this permissions as a managed policy rather than inline.

## Supported Versions

1. Terraform
- required_version = "~> 0.14"

2. Terraform AWS
- source  = "hashicorp/aws"
- version = ">= 3.20.0"

3. Terraform Kubernetes
- source  = "hashicorp/kubernetes"
- version = ">= 2.0.1"

4. AWS EKS
- cluster_version = "1.18"

## Instructions

1. Make sure you have the following installed and configured:
- AWS CLI
- AWS IAM Authenticator
- kubectl
- wget 
2. Set your working directory to: **eks-terraform-scripts**.
3. Execute:
```shell
terraform init
```
4. Execute:
```shell
terraform apply
```
5. Confirm the apply with a:
```shell
yes
```
6. Execute:
```shell
terraform apply
```
7. At this point you can configure **kubectl**:
```shell
aws eks --region $(terraform output -raw region) update-kubeconfig --name $(terraform output -raw cluster_name)
```
8. What you decide to do next is up to you; the cluster is ready for you to work with it.
9. Clean up:
```shell
terraform destroy
```