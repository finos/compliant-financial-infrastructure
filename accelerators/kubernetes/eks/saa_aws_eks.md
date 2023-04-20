# Service Approval Accelerator (SAA) v0.2 - AWS EKS - v2.0

## 1. Baseline

The policies included in the following document(s) are required for compliance,
unless otherwise noted in section 2: [Amendments to the Baseline](#2-amendments-to-the-baseline)

| Benchmark/Framework | Purpose | Name | Link | Comments |
|---|---|---|---|---|
| CIS      | Security | CIS Amazon Elastic Kubernetes Service (EKS) Benchmark v1.2.0 | [downloads.cisecurity.org](https://downloads.cisecurity.org/#/) | Cover recommendations for both profiles; see ammendments.  |
| PCI      | Compliance | Payment Card Industry Data Security Standard (PCI DSS) v3.2.1 | [pcisecuritystandards.org/document_library/](https://www.pcisecuritystandards.org/document_library/) | |

## 2. Amendments to the Baseline

1.1 Benchmark/Framework: CIS Amazon EKS Benchmark v1.2.0
1.2 Recommendation: 3.3.1 Prefer using a container-optimized OS when possible (Manual) - Profile Applicability: Level 2
1.3 Action: User's discretion.

2.1 Benchmark/Framework: CIS Amazon EKS Benchmark v1.2.0
2.2 Recommendation: 4.4.1 Prefer using secrets as files over secrets as environment variables (Manual) - Profile Applicability: Level 2
2.3 Action: User's discretion.

3.1 Benchmark/Framework: CIS Amazon EKS Benchmark v1.2.0
3.2 Recommendation: 4.4.2 Consider external secret storage (Manual) - Profile Applicability: Level 2
3.3 Action: User's discretion.

4.1 Benchmark/Framework: CIS Amazon EKS Benchmark v1.2.0
4.2 Recommendation: 5.1.4 Minimize Container Registries to only those approved (Manual) - Profile Applicability: Level 2
4.3 Action: User's discretion.

5.1 Benchmark/Framework: CIS Amazon EKS Benchmark v1.2.0
5.2 Recommendation: 5.4.5 Encrypt traffic to HTTPS load balancers with TLS certificates (Manual) - Profile Applicability: Level 2
5.3 Action: User's discretion.

6.1 Benchmark/Framework: CIS Amazon EKS Benchmark v1.2.0
6.2 Recommendation: 5.6.1 Consider Fargate for running untrusted workloads (Manual) - Profile Applicability: Level 1
6.3 Action: User's discretion.

## 3. Extensions to the Baseline

All guidelines provided by the aforementioned baseline are to be followed without modification or exception.

## 4. Security & Compliance Mapping

| Security Benchmark/Framework Control | PCI DSS v3.2.1 Requirement(s) | Comments |
|---|---|---|
| CIS Amazon EKS Benchmark v1.2.0 - 2.1.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.1.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.1.2 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.1.3 | Requirement 2| NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.1.4 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.1 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.2 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.3 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.4 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.5 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.6 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.7 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.8 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.9 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.10 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 3.2.11 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.1 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.2 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.3 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.4 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.5 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.6 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.7 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.1.8 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.2 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.3 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.4 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.5 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.6 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.7 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.2.8 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.3.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.3.2 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.6.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.6.2 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 4.6.3 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.1.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.1.2 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.1.3 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.2.1 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.3.1 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.4.1 | Requirements 2 & 8 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.4.2 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.4.3 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.4.4 | Requirement 2 | NA |
| CIS Amazon EKS Benchmark v1.2.0 - 5.5.1 | Requirements 2 & 8 | NA |
