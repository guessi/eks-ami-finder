#!/usr/bin/env bash

set -uo pipefail

readonly TARGET_VERSION="1.33"
readonly TARGET_VARIENT="${1:-}"

error_message() {
    echo "Error: $1" >&2
}

usage() {
    echo "Usage: $0 <Bottlerocket|AmazonLinux|Windows>" >&2
}

if [[ $# -ne 1 ]]; then
    error_message "Exactly one argument is required"
    usage
    exit 1
fi

if [[ -z ${TARGET_VARIENT} ]]; then
    error_message "TARGET_VARIENT must be specified"
    usage
    exit 1
fi

case ${TARGET_VARIENT} in
    Bottlerocket|AmazonLinux|Windows)
        ;;
    *)
        error_message "Invalid TARGET_VARIENT. Must be one of: Bottlerocket, AmazonLinux, Windows"
        usage
        exit 1
        ;;
esac

readonly REGIONS=(
    "af-south-1"
    "ap-east-1"
    "ap-northeast-1"
    "ap-northeast-2"
    "ap-northeast-3"
    "ap-south-1"
    "ap-south-2"
    "ap-southeast-1"
    "ap-southeast-2"
    "ap-southeast-3"
    "ap-southeast-4"
    "ap-southeast-5"
    "ap-southeast-7"
    "ca-central-1"
    "ca-west-1"
    "eu-central-1"
    "eu-central-2"
    "eu-north-1"
    "eu-south-1"
    "eu-south-2"
    "eu-west-1"
    "eu-west-2"
    "eu-west-3"
    "il-central-1"
    "me-central-1"
    "me-south-1"
    "mx-central-1"
    "sa-east-1"
    "us-east-1"
    "us-east-2"
    "us-west-1"
    "us-west-2"
)

get_ami_id() {
    local VARIANT="$1"
    local REGION="$2"
    local AMI_ID=""

    if [[ -z ${VARIANT} ]] || [[ -z ${REGION} ]]; then
        error_message "Both VARIANT and REGION must be specified"
        return 1
    fi

    case ${VARIANT} in
        "Bottlerocket")
            if ! AMI_ID=$(aws ssm get-parameter \
                --name "/aws/service/bottlerocket/aws-k8s-${TARGET_VERSION}/x86_64/latest/image_id" \
                --region "${REGION}" \
                --query "Parameter.Value" \
                --output text); then
                error_message "AWS SSM command failed for ${REGION}"
                return 1
            fi
            ;;
        "AmazonLinux")
            if ! AMI_ID=$(aws ssm get-parameter \
                --name "/aws/service/eks/optimized-ami/${TARGET_VERSION}/amazon-linux-2023/x86_64/standard/recommended/image_id" \
                --region "${REGION}" \
                --query "Parameter.Value" \
                --output text); then
                error_message "AWS SSM command failed for ${REGION}"
                return 1
            fi
            ;;
        "Windows")
            if ! AMI_ID=$(aws ssm get-parameter \
                --name "/aws/service/ami-windows-latest/Windows_Server-2022-English-Core-EKS_Optimized-${TARGET_VERSION}/image_id" \
                --region "${REGION}" \
                --query "Parameter.Value" \
                --output text); then
                error_message "AWS SSM command failed for ${REGION}"
                return 1
            fi
            ;;
        *)
            error_message "Unsupported VARIANT '${VARIANT}'"
            return 1
            ;;
    esac

    if [[ -z ${AMI_ID} ]] || [[ ${AMI_ID} == "None" ]]; then
        error_message "Could not find AMI ID for VARIANT ${VARIANT} in REGION ${REGION}"
        return 1
    fi

    echo "${AMI_ID}"
    return 0
}

get_ami_owner() {
    local AMI_ID="$1"
    local REGION="$2"
    local OWNER_ID=""

    if ! OWNER_ID=$(aws ec2 describe-images \
        --image-ids "${AMI_ID}" \
        --query 'Images[*].OwnerId' \
        --output text \
        --region "${REGION}"); then
        return 1
    fi

    echo "${OWNER_ID}"
    return 0
}

process_region() {
    local REGION="$1"
    local AMI_ID=""
    local OWNER_ID=""

    printf "%17s => " "${REGION}"

    if ! AMI_ID=$(get_ami_id "${TARGET_VARIENT}" "${REGION}"); then
        echo "AMI not found"
        return 1
    fi

    if ! OWNER_ID=$(get_ami_owner "${AMI_ID}" "${REGION}"); then
        echo "Failed to retrieve owner"
        return 1
    fi

    echo "${OWNER_ID}"
    return 0
}

main() {
    local RESULT_COUNT=0
    local FAILED_COUNT=0

    echo "----------------------------------------------------"
    echo "Checking AMI owners for ${TARGET_VARIENT} AMIs (${TARGET_VERSION})"
    echo "----------------------------------------------------"

    for TARGET_REGION in "${REGIONS[@]}"; do
        if process_region "${TARGET_REGION}"; then
            ((RESULT_COUNT++))
        else
            ((FAILED_COUNT++))
        fi
    done

    echo "----------------------------------------------------"
    echo "Summary: Found ${RESULT_COUNT} AMIs, ${FAILED_COUNT} region(s) failed"
    echo "----------------------------------------------------"

    [[ ${FAILED_COUNT} -eq 0 ]]
    return $?
}

main
exit $?
