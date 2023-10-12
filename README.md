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
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.24 --release-date 2023
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+---------------------------------------------------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                                                                     |
+-----------+-----------------------+--------------------------------+---------------------------------------------------------------------------------------------------------------------------------+
| us-east-1 | ami-0cbbb5a8c6f670bb6 | amazon-eks-node-1.24-v20230816 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.15, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-06f28e19c3ba73ef7 | amazon-eks-node-1.24-v20230728 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.15, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-08a2e6a8e82737230 | amazon-eks-node-1.24-v20230711 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.15, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-09e3c1f18094d099d | amazon-eks-node-1.24-v20230703 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-08c95f33fc51670df | amazon-eks-node-1.24-v20230607 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-0ea1e9d29dabc105b | amazon-eks-node-1.24-v20230526 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-08f0666b83a9a3cb9 | amazon-eks-node-1.24-v20230513 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-01b58531bffcd77d0 | amazon-eks-node-1.24-v20230509 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-0805526053854501b | amazon-eks-node-1.24-v20230501 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-0ce0bc9be2a044a29 | amazon-eks-node-1.24-v20230411 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-02f5ecb082b74cd86 | amazon-eks-node-1.24-v20230406 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.19-1.amzn2.0.1) |
| us-east-1 | ami-0733d88d7fb98418c | amazon-eks-node-1.24-v20230322 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.6-1.amzn2.0.2)  |
| us-east-1 | ami-0b4795e99297c2650 | amazon-eks-node-1.24-v20230304 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.10, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.6-1.amzn2.0.2)  |
| us-east-1 | ami-01ced323515f177b0 | amazon-eks-node-1.24-v20230217 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.10, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.6-1.amzn2.0.2)  |
| us-east-1 | ami-06bf8e441ff8de6c6 | amazon-eks-node-1.24-v20230203 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.9, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.6-1.amzn2.0.2)   |
| us-east-1 | ami-08794756cd16cd445 | amazon-eks-node-1.24-v20230127 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.9, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.6-1.amzn2.0.2)   |
| us-east-1 | ami-06c9b6a12f5bd0a96 | amazon-eks-node-1.24-v20230105 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.7, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.6-1.amzn2.0.2)   |
+-----------+-----------------------+--------------------------------+---------------------------------------------------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.24 --release-date 202304
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+---------------------------------------------------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                                                                     |
+-----------+-----------------------+--------------------------------+---------------------------------------------------------------------------------------------------------------------------------+
| us-east-1 | ami-0ce0bc9be2a044a29 | amazon-eks-node-1.24-v20230411 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.*)              |
| us-east-1 | ami-02f5ecb082b74cd86 | amazon-eks-node-1.24-v20230406 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.19-1.amzn2.0.1) |
+-----------+-----------------------+--------------------------------+---------------------------------------------------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.24 --release-date 202305
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                                                        |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| us-east-1 | ami-0ea1e9d29dabc105b | amazon-eks-node-1.24-v20230526 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
| us-east-1 | ami-08f0666b83a9a3cb9 | amazon-eks-node-1.24-v20230513 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
| us-east-1 | ami-01b58531bffcd77d0 | amazon-eks-node-1.24-v20230509 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
| us-east-1 | ami-0805526053854501b | amazon-eks-node-1.24-v20230501 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.11, docker: 20.10.17-1.amzn2.0.1, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.24 --release-date 202306
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                                                        |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| us-east-1 | ami-08c95f33fc51670df | amazon-eks-node-1.24-v20230607 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.24 --release-date 202307
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                                                        |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| us-east-1 | ami-06f28e19c3ba73ef7 | amazon-eks-node-1.24-v20230728 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.15, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
| us-east-1 | ami-08a2e6a8e82737230 | amazon-eks-node-1.24-v20230711 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.15, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
| us-east-1 | ami-09e3c1f18094d099d | amazon-eks-node-1.24-v20230703 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.13, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
```
</details>

```bash
$ eks-ami-finder --region us-east-1 --kubernetes-version 1.24 --release-date 202308
```

<details>
<summary>Click to expand!</summary>

```
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| Region    | AMI ID                | Name                           | Description                                                                                                        |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
| us-east-1 | ami-0cbbb5a8c6f670bb6 | amazon-eks-node-1.24-v20230816 | EKS Kubernetes Worker AMI with AmazonLinux2 image, (k8s: 1.24.15, docker: 20.10.23-1.amzn2.0.1, containerd: 1.6.*) |
+-----------+-----------------------+--------------------------------+--------------------------------------------------------------------------------------------------------------------+
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
