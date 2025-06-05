#!/bin/bash

# Copyright 2025 The PipeCD Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script updates the CODEOWNERS file by syncing from README files of each plugin.

COMMON_APPROVERS="@pipe-cd/pipecd-approvers"

# 1. List plugins
plugins=$(ls plugins | gsed 's/\// /g')
codeowners=""

for plugin in $plugins; do
    accounts=$(grep '\[Code Owners\]' plugins/$plugin/README.md | grep -o '@[^ )]*]' | gsed 's|]||'  | tr '\n' ' ')
    codeowners="${codeowners}plugins/$plugin/ $COMMON_APPROVERS $accounts\\
"
done

echo "$codeowners"

# 2. Update CODEOWNERS
gsed -i "/# --- Start plugins ---/,/# --- End plugins ---/c\\
# --- Start plugins ---\\
$codeowners\\# --- End plugins ---" CODEOWNERS
