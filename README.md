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
   --ami-type value, -t value            x86_64, x86_64-gpu, arm64 (default: "x86_64")
   --ami-family value, -f value          AL2, AL2023 (default: "AL2")
   --kubernetes-version value, -V value  Kubernetes version for AMI (default: "1.30")
   --release-date value, -d value        Release date with [yyyy], [yyyymm] or [yyyymmdd] format
   --include-deprecated                  (default: false)
   --max-results value, -n value         (default: "20")
   --debug                               (default: false)
   --help, -h                            show help
```

## Sample Output

<details><!-- markdownlint-disable-line -->
<summary>Click to expand!</summary><!-- markdownlint-disable-line -->

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.30 --release-date 2024 # for all 1.30 AMIs released in 2024
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0b9e545aa4c20aac6 | amazon-eks-node-1.30-v20240605 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-06-06T02:21:36.000Z | x86_64       |
| us-east-1 | ami-0469fcb2e219afb31 | amazon-eks-node-1.30-v20240531 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-05-31T16:01:09.000Z | x86_64       |
| us-east-1 | ami-0d06f30e8d6e02990 | amazon-eks-node-1.30-v20240522 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-05-23T05:59:52.000Z | x86_64       |
| us-east-1 | ami-01d267b705a0521d1 | amazon-eks-node-1.30-v20240514 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-05-14T17:03:12.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+

```

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.30 --release-date 202405 # for all 1.30 AMIs released with specific month

+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0469fcb2e219afb31 | amazon-eks-node-1.30-v20240531 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-05-31T16:01:09.000Z | x86_64       |
| us-east-1 | ami-0d06f30e8d6e02990 | amazon-eks-node-1.30-v20240522 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-05-23T05:59:52.000Z | x86_64       |
| us-east-1 | ami-01d267b705a0521d1 | amazon-eks-node-1.30-v20240514 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-05-14T17:03:12.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
```

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.30 --release-date 20240605 # for all 1.30 AMIs released with specific date

+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          | Architecture |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
| us-east-1 | ami-0b9e545aa4c20aac6 | amazon-eks-node-1.30-v20240605 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.30.0, containerd: 1.7.*) | 2026-06-06T02:21:36.000Z | x86_64       |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+--------------+
```

</details>

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
