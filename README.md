# eks-ami-finder

[![GitHub Actions](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/eks-ami-finder?status.svg)](https://godoc.org/github.com/guessi/eks-ami-finder)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/eks-ami-finder)](https://goreportcard.com/report/github.com/guessi/eks-ami-finder)
[![GitHub release](https://img.shields.io/github/release/guessi/eks-ami-finder.svg)](https://github.com/guessi/eks-ami-finder/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/eks-ami-finder)](https://github.com/guessi/eks-ami-finder/blob/master/go.mod)

retrieve Amazon EKS AMI with filters

## 🚀 Quick start

```bash
$ eks-ami-finder --help
```

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.32 --release-date 202501 # for all 1.32 AMIs released with specific month (prefix match)

+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0c8b3670978968623 | amazon-eks-node-1.32-v20250116 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.32.0, containerd: 1.7.*) | 2027-01-16T22:46:55.000Z | x86_64       |
| us-east-1 | ami-0226f5d1e7c6fc56e | amazon-eks-node-1.32-v20250103 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.32.0, containerd: 1.7.*) | 2027-01-04T00:22:02.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
```

</details>

## :accessibility: FAQ

Q: Where can I find he definition for the value of `--ami-type` flag?

A: See [amiType](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) definition here.

## 👷 Install

### For macOS users (Recommended)

```bash
brew tap guessi/tap && brew update && brew install eks-ami-finder
```

### Manually setup (Linux, Windows, macOS)

<details><!-- markdownlint-disable-line -->
<summary>Click to expand!</summary><!-- markdownlint-disable-line -->

### For Linux users

```bash
curl -fsSL https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Linux-$(uname -m).tar.gz -o - | tar zxvf -
mv -vf ./eks-ami-finder /usr/local/bin/eks-ami-finder
```

### For macOS users

```bash
curl -fsSL https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Darwin-$(uname -m).tar.gz -o - | tar zxvf -
mv -vf ./eks-ami-finder /usr/local/bin/eks-ami-finder
```

### For Windows users

```powershell
$SRC = 'https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Windows-x86_64.tar.gz'
$DST = 'C:\Temp\eks-ami-finder-Windows-x86_64.tar.gz'
Invoke-RestMethod -Uri $SRC -OutFile $DST
```

</details>

## ⚖️ License

[Apache-2.0](LICENSE)
