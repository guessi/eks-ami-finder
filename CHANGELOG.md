<!-- markdownlint-disable -->

# v1.3.0 / 2025-06-20

* Support for Auto Mode AMI search

# v1.2.2 / 2025-06-10

* Fix incorrect AMI returns for `BOTTLEROCKET_ARM_64_FIPS`
* Allow passing `--timeout` to avoid command running too long
* Early exit if EC2 endpoint not reachable for target region
* Fix multiple logic handling issues

# v1.2.1 / 2025-06-08

* Fix missing `me-central-1` setup for AL2/AL2023-based AMI

# v1.2.0 / 2025-06-06

* Update region support for `ap-east-2`

# v1.1.0 / 2025-05-30

* Set default version to 1.33
* Set default AMI_TYPE as `AL2023_x86_64_STANDARD`
* Allow user to search deprecated Windows Server 2016 based AMIs
* Make error messages understandable to fools

# v1.0.22 / 2025-05-28

* Update region support for `ap-southeast-7` and `mx-central-1`

# v1.0.21 / 2025-05-22

* Add support for `AL2023_ARM_64_NVIDIA`
* Fix incorrect ami filter for `AL2023_x86_64_NVIDIA`
* Migrate to `github.com/urfave/cli/v3`
* Bump dependencies

# v1.0.20 / 2025-04-04

* Extend support for `BOTTLEROCKET_ARM_64_FIPS` and `BOTTLEROCKET_x86_64_FIPS`
* Build with golang 1.24
* Dependencies update

# v1.0.19 / 2025-01-24

* Set default version to 1.32

# v1.0.18 / 2024-12-27

* Biuld with golang 1.23
* Minor code refactoring
* Dependencies update

# v1.0.17 / 2024-12-16

* Dependencies update

# v1.0.16 / 2024-09-30

* Set default version to 1.31
* Dependencies update
* Extend support for AL2023_x86_64_NEURON and AL2023_x86_64_NVIDIA

# v1.0.15 / 2024-09-09

* Dependencies update
* Extend support for all AMI TYPEs (AL2/AL2023/BOTTLEROCKET/WINDOWS)

# v1.0.14 / 2024-08-02

* Add support for new region - Kuala Lumpur (ap-southeast-5)
* Dependencies update

# v1.0.13 / 2024-06-24

* Add support for AL2023 based AMI searching

# v1.0.12 / 2024-06-15

* Set default version to 1.30
* Dependencies update

# v1.0.11 / 2024-03-24

* Biuld with golang 1.22
* Dependencies update

# v1.0.10 / 2024-02-18

* Implement sub-command "version"

# v1.0.9 / 2024-02-16

* Include "Architecture" info
* Dependencies update

# v1.0.8 / 2024-02-02

* Loosen check for `--release-date` input
* Dependencies update

# v1.0.7 / 2024-01-27

* Support new region - Calgary (ca-west-1)
* Set default version to 1.29
* Dependencies update

# v1.0.6 / 2024-01-07

* Dependencies update

# v1.0.5 / 2023-11-20

* Introduce `--max-results` flag
* Add paginator support for DescribeImages()
* Include `DeprecationTime` info in the output

# v1.0.4 / 2023-10-12

* Set default version to 1.28

# v1.0.3 / 2023-10-12

* Biuld with golang 1.21
* Dependencies update

# v1.0.2 / 2023-08-08

* Identical with v1.0.1

# v1.0.1 / 2023-08-08

* Dependencies update

# v1.0.0 / 2023-05-13

* Initial Release
