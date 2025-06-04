# CONTRIBUTING.md

## Adding a New Plugin

Please follow the steps below. **You MUST NOT initialize a new plugin by opening a PR directly to avoid confusion.**

1. Please [open a 'New plugin proposal' issue](https://github.com/pipe-cd/community-plugins/issues/new?template=new-plugin.yaml).
2. A maintainer will create and merge a PR to initialize the plugin directory, CODEOWNERS, and so on by [the 'New Plugin' workflow](https://github.com/pipe-cd/community-plugins/actions/workflows/new-plugin.yaml).
3. The maintainer will close the issue and you can start working on the plugin.

## Becoming a Code Owner

A code owner is responsible for maintaining the plugin, including triaging issues, reviewing PRs, and submitting bug fixes.

There are two ways to become a code owner:
- When [initializing a new plugin](#adding-a-new-plugin)
- When the current code owner(s) invites a new one. Then, follow these steps:
    1. Update the Code Owners in the README.md of the plugin.
    2. Execute `make sync/codeowners` to update the CODEOWNERS file.
    3. Create a PR and get it merged.

_What to do when a code owner became inactive?: TBD_

## Issues

- Opening a new issue
- Finding 
- Discuss
- Good first issues: TBD

## Pull Request

### Reviewing

- Code Owners must lead reviewing for their plugin. `@pipe-cd/pipecd-approvers` is just a helper and does not assure the quality of the plugin.
- After CI passes and the code owner approves the PR, the code owner can merge it.

- _Note: When you are the only one code owner of a plugin and you want to submit a PR, `@pipe-cd/pipecd-approvers` can review it instantly. However, they are **not fully responsible** for the plugin and you should find a new code owner._

## Development, Testing

TBD

We're preparing a guide of the SDK/API.

See [example-stage plugin](examples/example-stage) for the example.

## Release Procedure

TBD
