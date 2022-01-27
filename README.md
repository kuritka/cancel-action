# CANCEL-ACTION

# Execution
- If action is executed against branch (e.g: push to feature-branch), the full dev process is triggered `go mod tidy,`,`go build`,..
- If action is executed against Tag (release, [usage](https://github.com/kuritka/cancel-action-test/blob/main/.github/workflows/build.yaml)), the binary is downloaded and executed (much faster)


## Release details and TODO
 - release.yaml releases binaries. 
 - Use tags `0.0.1`,`0.0.2`.... DO NOT use `v` prefix`v0.0.3`, because go-releaser needs to be set to provide 
name of binary containing `v` prefix.
 - run test workflow manually after release workflow finish