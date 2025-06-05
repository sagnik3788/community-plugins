#!/bin/bash
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
