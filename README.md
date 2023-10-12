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

eks-ami-finder version 1.0.4
```

```bash
$ eks-ami-finder --help

NAME:
   eks-ami-finder - retrieve Amazon EKS AMI with filters

USAGE:
   eks-ami-finder [global options] command [command options] [arguments...]

VERSION:
   1.0.4

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value, -r value              Region for the AMI (default: "us-east-1")
   --owner-id value, -o value            Owner ID of the AMI
   --ami-type value, -t value            x86_64, x86_64-gpu, arm64 (default: "x86_64")
   --kubernetes-version value, -V value  Kubernetes version for AMI (default: "1.28")
   --release-date value, -d value        Release date with [yyyymmdd] date string format
   --include-deprecated                  (default: false)
   --debug                               (default: false)
   --help, -h                            show help
   --version, -v                         print the version
```

# Sample Output

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.27 \
      --release-date 2023 # for all 1.27 AMIs released in 2023
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
| us-east-1 | ami-0474c5fe3b15d9685 | amazon-eks-node-1.27-v20231002 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.5, containerd: 1.6.*) |
| us-east-1 | ami-0c92ea9c7c0380b66 | amazon-eks-node-1.27-v20230919 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.5, containerd: 1.6.*) |
| us-east-1 | ami-013895b64fa9cbcba | amazon-eks-node-1.27-v20230825 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.4, containerd: 1.6.*) |
| us-east-1 | ami-080632d422a0f7cc5 | amazon-eks-node-1.27-v20230816 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.3, containerd: 1.6.*) |
| us-east-1 | ami-0bc4534a93057f9fb | amazon-eks-node-1.27-v20230728 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.3, containerd: 1.6.*) |
| us-east-1 | ami-0ae32cfe425c3643a | amazon-eks-node-1.27-v20230711 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.3, containerd: 1.6.*) |
| us-east-1 | ami-061112afff4339a5f | amazon-eks-node-1.27-v20230703 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.1, containerd: 1.6.*) |
| us-east-1 | ami-0fe06c902df3a937b | amazon-eks-node-1.27-v20230607 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.1, containerd: 1.6.*) |
| us-east-1 | ami-0b94943bd76cb55c2 | amazon-eks-node-1.27-v20230526 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.1, containerd: 1.6.*) |
| us-east-1 | ami-0e38f9978e7cac6dc | amazon-eks-node-1.27-v20230513 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.1, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.27 \
      --release-date 202308 # for specific month
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
| us-east-1 | ami-013895b64fa9cbcba | amazon-eks-node-1.27-v20230825 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.4, containerd: 1.6.*) |
| us-east-1 | ami-080632d422a0f7cc5 | amazon-eks-node-1.27-v20230816 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.3, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder \
      --region us-east-1 \
      --kubernetes-version 1.27 \
      --release-date 20230825 # for specific date
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                         |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
| us-east-1 | ami-013895b64fa9cbcba | amazon-eks-node-1.27-v20230825 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.27.4, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+-------------------------------------------------------------------------------------+
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
