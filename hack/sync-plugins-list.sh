#!/bin/bash
# This script updates the plugins list in the issue templates.

plugins=$(ls plugins | sed 's/\// /g')
files=".github/ISSUE_TEMPLATE/bug-report.yaml .github/ISSUE_TEMPLATE/enhancement.yaml"

plugins_list=""

for plugin in $plugins; do
    plugins_list="${plugins_list}        - $plugin\\
"
done

for f in $files; do
    sed -i '' "/# --- Start plugins list ---/,/# --- End plugins list ---/c\\
        # --- Start plugins list ---\\
$plugins_list\\        # --- End plugins list ---\\
" "$f"
done
