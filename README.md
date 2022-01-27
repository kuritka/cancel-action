# CANCEL-ACTION

# Execution
- If action is executed against branch (e.g: push to feature-branch), the full dev process is triggered `go mod tidy,`,`go build`,..
- If action is executed against with @tag the [release is used](https://github.com/kuritka/cancel-action-test/blob/main/.github/workflows/build.yaml)).
