# eks-ami-finder

[![GitHub Actions](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/eks-ami-finder?status.svg)](https://godoc.org/github.com/guessi/eks-ami-finder)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/eks-ami-finder)](https://goreportcard.com/report/github.com/guessi/eks-ami-finder)
[![GitHub release](https://img.shields.io/github/release/guessi/eks-ami-finder.svg)](https://github.com/guessi/eks-ami-finder/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/eks-ami-finder)](https://github.com/guessi/eks-ami-finder/blob/main/go.mod)

A command-line tool to find Amazon EKS optimized AMI IDs across different versions and regions.

## 📋 Overview

Users often need to pin their AMI to a specific version of Amazon EKS Optimized AMI for security or compliance reasons. However, AWS documentation only provides methods to retrieve the latest AMI versions:

- [Retrieve recommended Amazon Linux AMI IDs](https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id.html)
- [Retrieve recommended Microsoft Windows AMI IDs](https://docs.aws.amazon.com/eks/latest/userguide/retrieve-windows-ami-id.html)
- [Retrieve recommended Bottlerocket AMI IDs](https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id-bottlerocket.html)

`eks-ami-finder` fills this gap by providing access to historical AMI information for Amazon Linux, Windows, and Bottlerocket-based EKS optimized AMIs.

## 🔢 Prerequisites

* An IAM Role/User with [ec2:DescribeImages](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImages.html) permission.

## 🚀 Quick start

### Basic Usage

```bash
eks-ami-finder help
```

### Find the latest AMIs

```bash
eks-ami-finder
```

### Find AMIs by Release Date

```bash
# Find all AMIs released in September 2025 (prefix match), with no region specify
eks-ami-finder --release-date 202509

# Find AMIs released on a specific date
eks-ami-finder --release-date 20250920 --region us-east-1
```

### Find AMIs by Kubernetes Version

```bash
# Find AMIs for Kubernetes 1.34
eks-ami-finder --kubernetes-version 1.34 --region us-east-1

# Combine Kubernetes version with specific release date
eks-ami-finder --kubernetes-version 1.34 --release-date 20250920 --region us-east-1
```

### Filter by AMI Type

```bash
# Find Amazon Linux 2023 AMIs
eks-ami-finder --ami-type AL2023_x86_64_STANDARD --region us-east-1

# Find Windows AMIs
eks-ami-finder --ami-type WINDOWS_CORE_2022_x86_64 --region us-east-1

# Find Bottlerocket AMIs
eks-ami-finder --ami-type BOTTLEROCKET_x86_64 --region us-east-1
```

### Example Output

```bash
eks-ami-finder --kubernetes-version 1.34 --release-date 20250920 --region us-east-1

+-----------+-----------------------+-------------------------------------------------------+---------------------------------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                                                  | Description                                                                                                   | DeprecationTime          | Architecture |
+-----------+-----------------------+-------------------------------------------------------+---------------------------------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0093e29064b926113 | amazon-eks-node-al2023-x86_64-standard-1.34-v20250920 | EKS-optimized Kubernetes node based on Amazon Linux 2023, (k8s: 1.34.1, containerd: 2.1.4-1.eks.amzn2023.0.1) | 2027-09-24T00:36:26.000Z | x86_64       |
+-----------+-----------------------+-------------------------------------------------------+---------------------------------------------------------------------------------------------------------------+--------------------------+--------------+
```

### Key Capabilities

- **Historical AMI Search**: Find specific versions of EKS-optimized AMIs, not just the latest.
- **Multi-OS Support**: Search Amazon Linux, Windows, and Bottlerocket AMIs.
- **Flexible Filtering**: Filter by Kubernetes version, release date, AMI type, region, etc.

## ❓ FAQ

### Q: How does `eks-ami-finder` look up AMI IDs?

`eks-ami-finder` first identifies the Owner IDs of the AMIs ([source](hack/ami-owner-info-check.sh)), then filters AMI IDs released by these Owner IDs ([source](cmd/search.go)) using pattern matching. It's that simple!

### Q: Where can I find the definition for the `--ami-type` flag value?

See the [amiType](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) definition in the AWS documentation.

### Q: Does an AMI description guarantee it's an official build?

Not necessarily. AMI descriptions like `EKS-optimized Kubernetes node based on Amazon Linux 2023, (k8s: 1.34.1, containerd: 2.1.4-1.eks.amzn2023.0.1)` can be defined by anyone. You still need to verify that it comes from the Amazon EKS team by checking the Owner ID.

## 👷 Install

### For macOS/Linux users (Recommended)

Brand new install

```bash
brew tap guessi/tap && brew update && brew install eks-ami-finder
```

To upgrade version

```bash
brew update && brew upgrade eks-ami-finder
```

### Manually setup (Linux, Windows, macOS)

<details><!-- markdownlint-disable-line -->
<summary>Click to expand!</summary><!-- markdownlint-disable-line -->

#### For Linux users

```bash
curl -fsSL https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Linux-$(uname -m).tar.gz -o - | tar zxvf -
mv -vf ./eks-ami-finder /usr/local/bin/eks-ami-finder
```

#### For macOS users

```bash
curl -fsSL https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Darwin-$(uname -m).tar.gz -o - | tar zxvf -
mv -vf ./eks-ami-finder /usr/local/bin/eks-ami-finder
```

#### For Windows users

```powershell
$SRC = 'https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Windows-x86_64.tar.gz'
$DST = 'C:\Temp\eks-ami-finder-Windows-x86_64.tar.gz'
Invoke-RestMethod -Uri $SRC -OutFile $DST
```

</details>

## ⚖️ License

[Apache-2.0](LICENSE)
