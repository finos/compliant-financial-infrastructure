# EKS - Terraform

## Requirements

Before running the Terraform scripts, make sure to set the [IAM required permissions](./iam-required-permissions.md) first.

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
