# eks-ami-finder

[![GitHub Actions](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/guessi/eks-ami-finder/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/eks-ami-finder?status.svg)](https://godoc.org/github.com/guessi/eks-ami-finder)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/eks-ami-finder)](https://goreportcard.com/report/github.com/guessi/eks-ami-finder)
[![GitHub release](https://img.shields.io/github/release/guessi/eks-ami-finder.svg)](https://github.com/guessi/eks-ami-finder/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/eks-ami-finder)](https://github.com/guessi/eks-ami-finder/blob/master/go.mod)

retrieve Amazon EKS AMI with filters

# Usage

```bash
$ eks-ami-finder --version

eks-ami-finder version 1.0.5
```

```bash
$ eks-ami-finder --help

NAME:
   eks-ami-finder - retrieve Amazon EKS AMI with filters

USAGE:
   eks-ami-finder [global options] command [command options] [arguments...]

VERSION:
   1.0.5

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value, -r value              Region for the AMI (default: "us-east-1")
   --owner-id value, -o value            Owner ID of the AMI
   --ami-type value, -t value            x86_64, x86_64-gpu, arm64 (default: "x86_64")
   --kubernetes-version value, -V value  Kubernetes version for AMI (default: "1.28")
   --release-date value, -d value        Release date with [yyyymmdd] date string format
   --include-deprecated                  (default: false)
   --max-results value, -n value         (default: "20")
   --debug                               (default: false)
   --help, -h                            show help
   --version, -v                         print the version
```

# Sample Output

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.28 \
      --release-date 2023 # for all 1.28 AMIs released in 2023
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| us-east-1 | ami-0d881c8e9d4844a86 | amazon-eks-node-1.28-v20231230 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-30T08:27:35.000Z |
| us-east-1 | ami-0df88a6d3d96762e8 | amazon-eks-node-1.28-v20231201 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-04T00:06:54.000Z |
| us-east-1 | ami-0e0b0f2cb811d16b0 | amazon-eks-node-1.28-v20231116 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.6.*) | 2025-11-16T08:14:03.000Z |
| us-east-1 | ami-02872df47199586cc | amazon-eks-node-1.28-v20231106 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.6.*) | 2025-11-07T19:18:57.000Z |
| us-east-1 | ami-0c97930d0d19e564a | amazon-eks-node-1.28-v20231027 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.2, containerd: 1.6.*) | 2025-10-27T05:45:38.000Z |
| us-east-1 | ami-0dd7006cb3a28d563 | amazon-eks-node-1.28-v20231002 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.1, containerd: 1.6.*) | 2025-10-03T04:21:57.000Z |
| us-east-1 | ami-0164b8ae1906d3372 | amazon-eks-node-1.28-v20230919 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.1, containerd: 1.6.*) | 2025-09-20T19:16:35.000Z |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
```
</details>

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.28 \
      --release-date 2023 --max-results 3 # for all 1.28 AMIs released in 2023 and show only most recent 3 releases.
```

<details>
<summary>Click to expand!</summary>
```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| us-east-1 | ami-0d881c8e9d4844a86 | amazon-eks-node-1.28-v20231230 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-30T08:27:35.000Z |
| us-east-1 | ami-0df88a6d3d96762e8 | amazon-eks-node-1.28-v20231201 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-04T00:06:54.000Z |
| us-east-1 | ami-0e0b0f2cb811d16b0 | amazon-eks-node-1.28-v20231116 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.6.*) | 2025-11-16T08:14:03.000Z |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
```
</details>

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.28 \
      --release-date 202312 # for specific month
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| us-east-1 | ami-0d881c8e9d4844a86 | amazon-eks-node-1.28-v20231230 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-30T08:27:35.000Z |
| us-east-1 | ami-0df88a6d3d96762e8 | amazon-eks-node-1.28-v20231201 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-04T00:06:54.000Z |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
```
</details>

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.28 \
      --release-date 20231201 # for specific date
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         | DeprecationTime          |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
| us-east-1 | ami-0df88a6d3d96762e8 | amazon-eks-node-1.28-v20231201 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.28.3, containerd: 1.7.*) | 2025-12-04T00:06:54.000Z |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+--------------------------+
```
</details>

# Install

### Homebrew

```bash
$ brew tap guessi/tap && brew install eks-ami-finder
```

### For non-Homebrew users, click `Details` to view more methods.

<details>

### For Linux users

```bash
$ curl -fsSL https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Linux-$(uname -m).tar.gz -o - | tar zxvf -
$ mv ./eks-ami-finder /usr/local/bin/eks-ami-finder
```

### For macOS users

```bash
$ curl -fsSL https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Darwin-$(uname -m).tar.gz -o - | tar zxvf -
$ mv ./eks-ami-finder /usr/local/bin/eks-ami-finder
```

### For Windows users

```powershell
PS> $SRC = 'https://github.com/guessi/eks-ami-finder/releases/latest/download/eks-ami-finder-Windows-x86_64.tar.gz'
PS> $DST = 'C:\Temp\eks-ami-finder-Windows-x86_64.tar.gz'
PS> Invoke-RestMethod -Uri $SRC -OutFile $DST
```
</details>

# License

[Apache-2.0](LICENSE)
