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
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.32 --release-date 202505 # for all 1.32 AMIs released with specific month (prefix match)

+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0f2e4735b924be9d0 | amazon-eks-node-1.32-v20250519 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.32.3, containerd: 1.7.*) | 2027-05-20T17:09:13.000Z | x86_64       |
| us-east-1 | ami-0bda772ad7684f8d5 | amazon-eks-node-1.32-v20250514 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.32.3, containerd: 1.7.*) | 2027-05-14T04:13:57.000Z | x86_64       |
| us-east-1 | ami-0f0e6b8d1eb5bf2cf | amazon-eks-node-1.32-v20250505 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.32.3, containerd: 1.7.*) | 2027-05-06T18:29:33.000Z | x86_64       |
| us-east-1 | ami-08075f9ccf102cac9 | amazon-eks-node-1.32-v20250501 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.32.3, containerd: 1.7.*) | 2027-05-01T05:32:25.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
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
