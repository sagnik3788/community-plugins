# Example Stage Plugin

| Status        |           |
| ------------- |-----------|
| Stability     | alpha   |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/pipe-cd/community-plugins?query=is%3Aissue%20is%3Aopen%20label%3Aplugin%2Fexample-stage%20&label=open&color=orange)](https://github.com/pipe-cd/community-plugins/issues?q=is%3Aopen+is%3Aissue+label%3Aplugin%2Fexample-stage) [![Closed issues](https://img.shields.io/github/issues-search/pipe-cd/community-plugins?query=is%3Aissue%20is%3Aclosed%20label%3Aplugin%2Fexample-stage%20&label=closed&color=blue)](https://github.com/pipe-cd/community-plugins/issues?q=is%3Aclosed+is%3Aissue+label%3Aplugin%2Fexample-stage) |
| [Code Owners](https://github.com/pipe-cd/community-plugins/blob/main/CONTRIBUTING.md#becoming-a-code-owner)   | [@your-account](https://www.github.com/) - Seeking more code owners!  |

<!-- | Feature Status |  |
|---------|--------|
| QuickSync | alpha |
| PipelineSync | beta |
| LiveState View | beta |
| DriftDetection | - |
| PlanPreview | - | -->

This plugin is an example plugin of stages. Each stage just shows a message on UI.

## Stages

- **EXAMPLE_HELLO**: shows a message on UI
    ```
    Hello <name> from the example PLAN stage!
    CommonMessage: <commonMessage>
    ```

- **EXAMPLE_GOODBYE**: shows a message on UI
    ```
    Goodbye from example GOODBYE stage!
    Message: <message>
    CommonMessage: <commonMessage>
    ```

## Configuration

### Piped Plugin Config

| Field | Type | Description | Required |
|-|-|-|-|
| commonMessage | string | The common message to be shown in all stages. | No |

<!-- DeployTarget Config -->
<!-- Application Config -->

### Stage Config

EXAMPLE_HELLO Stage:

| Field | Type | Description | Required |
|-|-|-|-|
| name | string | The name to be shown in the stage. | Yes |

EXAMPLE_GOODBYE Stage:

| Field | Type | Description | Required |
|-|-|-|-|
| message | string | The message to be shown in the stage. | No |


<!-- ## Notes -->
<!-- If there're some notable points, describe them here -->
