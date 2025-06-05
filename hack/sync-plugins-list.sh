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

# This script updates the plugins list in the issue templates.

plugins=$(ls plugins | gsed 's/\// /g')
issue_templates=".github/ISSUE_TEMPLATE/bug-report.yaml .github/ISSUE_TEMPLATE/enhancement.yaml"

# 1. Issue Templates
plugins_list=""
for plugin in $plugins; do
    plugins_list="${plugins_list}        - $plugin\\
"
done

for f in $issue_templates; do
    gsed -i "/# --- Start plugins list ---/,/# --- End plugins list ---/c\\
        # --- Start plugins list ---\\
$plugins_list\\        # --- End plugins list ---" "$f"
done

# 2. labeler
labeler_config_file=".github/labeler.yaml"
plugin_labels=""
for plugin in $plugins; do
    plugin_labels="${plugin_labels}\\
plugin/$plugin:\\
  - changed-files:\\
      - any-glob-to-any-file: "plugins/$plugin/**"\\
"
done
gsed -i "/# --- Start plugins list ---/,/# --- End plugins list ---/c\\
# --- Start plugins list ---\\
$plugin_labels\\# --- End plugins list ---" "$labeler_config_file"
