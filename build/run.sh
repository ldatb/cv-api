# Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
#
# SPDX-License-Identifier: MIT
#!/usr/bin/bash

ROOT_DIR=$(dirname $(readlink -f "$0"))
RUN_FILE="cv-api.go"

go run ${ROOT_DIR}/../cmd/${RUN_FILE}
