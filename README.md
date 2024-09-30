# eks-ami-finder

[![GitHub Actions](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/eks-ami-finder?status.svg)](https://godoc.org/github.com/guessi/eks-ami-finder)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/eks-ami-finder)](https://goreportcard.com/report/github.com/guessi/eks-ami-finder)
[![GitHub release](https://img.shields.io/github/release/guessi/eks-ami-finder.svg)](https://github.com/guessi/eks-ami-finder/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/eks-ami-finder)](https://github.com/guessi/eks-ami-finder/blob/master/go.mod)

retrieve Amazon EKS AMI with filters

## Usage

```bash
$ eks-ami-finder --help

NAME:
   eks-ami-finder - retrieve Amazon EKS AMI with filters

USAGE:
   eks-ami-finder [global options] command [command options]

COMMANDS:
   version, v  Print version number
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value, -r value              Region for the AMI (default: "us-east-1")
   --owner-id value, -o value            Owner ID of the AMI
   --ami-type value, -t value            AMI Type for the AMI (default: "AL2_x86_64")
   --kubernetes-version value, -V value  Kubernetes version for AMI (default: "1.31")
   --release-date value, -d value        Release date with [yyyy], [yyyymm] or [yyyymmdd] format
   --include-deprecated                  (default: false)
   --max-results value, -n value         (default: "20")
   --debug                               (default: false)
   --help, -h                            show help
   --version, -v                         print the version
```

## Sample Output

<details><!-- markdownlint-disable-line -->
<summary>Click to expand!</summary><!-- markdownlint-disable-line -->

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.31 --release-date 2024 # for all 1.31 AMIs released in 2024

+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                            | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-03a66e914971f8646 | amazon-eks-node-1.31-v20240924 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.31.0, containerd: 1.7.11-*) | 2026-09-25T02:01:32.000Z | x86_64       |
| us-east-1 | ami-00ec84c1189958713 | amazon-eks-node-1.31-v20240917 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.31.0, containerd: 1.7.11-*) | 2026-09-18T06:26:49.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
```

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.31 --release-date 202409 # for all 1.31 AMIs released with specific month

+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                            | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-03a66e914971f8646 | amazon-eks-node-1.31-v20240924 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.31.0, containerd: 1.7.11-*) | 2026-09-25T02:01:32.000Z | x86_64       |
| us-east-1 | ami-00ec84c1189958713 | amazon-eks-node-1.31-v20240917 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.31.0, containerd: 1.7.11-*) | 2026-09-18T06:26:49.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
```

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.31 --release-date 20240924 # for all 1.31 AMIs released with specific date

+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                            | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-03a66e914971f8646 | amazon-eks-node-1.31-v20240924 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.31.0, containerd: 1.7.11-*) | 2026-09-25T02:01:32.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+----------------------------------------------------------------------------------------+--------------------------+--------------+
```

</details>

## FAQ

Q: Where can I find he definition for the value of `--ami-type` flag?

A: See [amiType](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) definition here.

## Install

### Homebrew

```bash
brew tap guessi/tap && brew install eks-ami-finder
```

### For non-Homebrew users

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

## License

[Apache-2.0](LICENSE)
