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

```bash
eks-ami-finder help
```

```bash
eks-ami-finder --release-date 202509 # for all AMIs released with specific month (prefix match)

+-----------+-----------------------+-------------------------------------------------------+---------------------------------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                                                  | Description                                                                                                   | DeprecationTime          | Architecture |
+-----------+-----------------------+-------------------------------------------------------+---------------------------------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0093e29064b926113 | amazon-eks-node-al2023-x86_64-standard-1.34-v20250920 | EKS-optimized Kubernetes node based on Amazon Linux 2023, (k8s: 1.34.1, containerd: 2.1.4-1.eks.amzn2023.0.1) | 2027-09-24T00:36:26.000Z | x86_64       |
| us-east-1 | ami-031271ba36c0b8711 | amazon-eks-node-al2023-x86_64-standard-1.34-v20250915 | EKS-optimized Kubernetes node based on Amazon Linux 2023, (k8s: 1.34.0, containerd: 2.1.4-1.eks.amzn2023.0.1) | 2027-09-16T18:17:54.000Z | x86_64       |
+-----------+-----------------------+-------------------------------------------------------+---------------------------------------------------------------------------------------------------------------+--------------------------+--------------+
```

## :accessibility: FAQ

Q: How does `eks-ami-finder` lookup the AMI IDs? what's the magic behind the scene?

A: `eks-ami-finder` will first find out the Owner IDs of the AMI [HERE](hack/ami-owner-info-check.sh), then filter out AMI IDs that released by these Owner IDs [HERE](cmd/search.go) with patterns, just that simple!

Q: Where can I find he definition for the value of `--ami-type` flag?

A: See [amiType](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) definition here.

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
