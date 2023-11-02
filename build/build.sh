# Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
#
# SPDX-License-Identifier: MIT
#!/usr/bin/bash

# Define all vars for the script
ROOT_DIR=$(dirname $(readlink -f "$0"))
OUTPUT_FILE="API"
RUN_FILE="cv-api.go"
VERBOSE=false
BUILD_FLAGS=""

# Get opts
usage() {
    cat >&2 <<EOF
Usage:
$(basename $0)

Options:
    -o:  Specify the name of the output file (Defaults to API)
    -f:  File to run inside the cmd folder (Defaults to cv-api.go)
    -v:  Verbose moded
EOF
}

while getopts 'o:f:vh' arg; do
    case "${arg}" in
        o) OUTPUT_FILE="${OPTARG}" ;;
        f) RUN_FILE="${OPTARG}" ;;
        v) VERBOSE=true ;;
        h) usage
            exit 0 ;;
        *) usage
            exit 1 ;;
    esac
done

# Construct build flags
if [ $VERBOSE = true ]; then
    BUILD_FLAGS="${BUILD_FLAGS} -x"
fi

go build -o $OUTPUT_FILE $BUILD_FLAGS ${ROOT_DIR}/../cmd/${RUN_FILE}
