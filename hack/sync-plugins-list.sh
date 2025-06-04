#!/bin/bash
# This script updates the plugins list in the issue templates.

plugins=$(ls plugins | sed 's/\// /g')
issue_templates=".github/ISSUE_TEMPLATE/bug-report.yaml .github/ISSUE_TEMPLATE/enhancement.yaml"

# 1. Issue Templates
plugins_list=""
for plugin in $plugins; do
    plugins_list="${plugins_list}        - $plugin\\
"
done

for f in $issue_templates; do
    sed -i '' "/# --- Start plugins list ---/,/# --- End plugins list ---/c\\
        # --- Start plugins list ---\\
$plugins_list\\        # --- End plugins list ---\\
" "$f"
done

# 2. labeler
labeler_config_file=".github/labeler.yaml"
plugin_labels=""
for plugin in $plugins; do
    plugin_labels="${plugin_labels}\\
plugin/$plugin: \\
  - changed-files: \\
      - any-glob-to-any-file: "plugins/$plugin/**"\\
"
done
sed -i '' "/# --- Start plugins list ---/,/# --- End plugins list ---/c\\
# --- Start plugins list ---\\
$plugin_labels\\# --- End plugins list ---\\
" "$labeler_config_file"