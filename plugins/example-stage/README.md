# example-stage plugin

| Metadata        |           |
| ------------- |-----------|
|[Stability](/README.md#stability-levels)     | In Development   |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/pipe-cd/community-plugins?query=is%3Aissue%20is%3Aopen%20label%3Aplugin%2Fexample-stage%20&label=open&color=orange)](https://github.com/pipe-cd/community-plugins/issues?q=is%3Aopen+is%3Aissue+label%3Aplugin%2Fexample-stage) |
| [Code Owners](/CONTRIBUTING.md#becoming-a-code-owner)   |  [@pipe-cd/pipecd-approvers](https://github.com/orgs/pipe-cd/teams/pipecd-approvers)  |

## Supported Features

- PipelineSync
<!-- 
- QuickSync
- Prune
- LiveState View
- DriftDetection
- PlanPreview
-->

<!-- You can add additional rows like 'PipelineSync by Istio', 'Analysis by <some-o11y-provider>', etc. -->

<!-- For stage plugins, only PipelineSync would be supported in most cases. -->

## Overview

This plugin is an example plugin of stages. Each stage just shows a message on the Deployment UI.

## Stages

### EXAMPLE_HELLO

It shows a message on the UI.
```
Hello <name> from the example PLAN stage!
CommonMessage: <commonMessage>
```

### EXAMPLE_GOODBYE

It shows a message on the UI.
```
Goodbye from example GOODBYE stage!
Message: <message>
CommonMessage: <commonMessage>
```


## Plugin Configuration

### Plugin scope config

| Field | Type | Description | Required | Default |
|-|-|-|-|-|
| commonMessage | string | The common message to be shown in all stages. | No | "" |

<!-- ### Deploy Target config -->

## Application Configuration

<!-- ### Application scope options -->

### Stage options

#### EXAMPLE_HELLO Stage

| Field | Type | Description | Required | Default |
|-|-|-|-|-|
| name | string | The name to be shown in the stage. | Yes | "" |

#### EXAMPLE_GOODBYE Stage

| Field | Type | Description | Required | Default |
|-|-|-|-|-|
| message | string | The message to be shown in the stage. | No | "" |
