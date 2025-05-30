# eks-ami-finder

[![GitHub Actions](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/eks-ami-finder?status.svg)](https://godoc.org/github.com/guessi/eks-ami-finder)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/eks-ami-finder)](https://goreportcard.com/report/github.com/guessi/eks-ami-finder)
[![GitHub release](https://img.shields.io/github/release/guessi/eks-ami-finder.svg)](https://github.com/guessi/eks-ami-finder/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/eks-ami-finder)](https://github.com/guessi/eks-ami-finder/blob/main/go.mod)

retrieve Amazon EKS AMI with filters

## 🔢 Prerequisites

* An IAM Role/User with [ec2:DescribeImages](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImages.html) permission.

## 🚀 Quick start

```bash
$ eks-ami-finder --help
```

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.33 --release-date 202505 # for all 1.33 AMIs released with specific month (prefix match)

+-----------+-----------------------+-------------------------------------------------------+--------------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                                                  | Description                                                                                | DeprecationTime          | Architecture |
+-----------+-----------------------+-------------------------------------------------------+--------------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-014f71cc7221992de | amazon-eks-node-al2023-x86_64-standard-1.33-v20250519 | EKS-optimized Kubernetes node based on Amazon Linux 2023, (k8s: 1.33.0, containerd: 1.7.*) | 2027-05-20T17:09:14.000Z | x86_64       |
+-----------+-----------------------+-------------------------------------------------------+--------------------------------------------------------------------------------------------+--------------------------+--------------+
```

## :accessibility: FAQ

Q: How does `eks-ami-finder` lookup the AMI IDs? what's the magic behind the scene?

A: `eks-ami-finder` will first find out the Owner IDs of the AMI [HERE](hack/ami-owner-info-check.sh), then filter out AMI IDs that released by these Owner IDs [HERE](cmd/search.go) with patterns, just that simple!

Q: Where can I find he definition for the value of `--ami-type` flag?

A: See [amiType](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) definition here.

## 👷 Install

### For macOS/Linux users (Recommended)

```bash
brew tap guessi/tap && brew update && brew install eks-ami-finder
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
