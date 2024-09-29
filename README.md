# Request Random Delay

This HTTP server exposes APIs to simulate random delays. Useful to test the Sidecar proxy behaviours (such as in Istio) when these delays are added.

## Build

- Use `make build` to build the processor and the test services.
- To build and push the Docker images use `PUSH_MULTIARCH=true make docker`. By default it only builds `linux/amd64` & `linux/arm64`.
  - The images get pushed to `australia-southeast1-docker.pkg.dev/field-engineering-apac/public-repo` but you can override this with the env var ``
- Run make help for all the build directives.

## Possible Configuration

| Env        | Description                                                                                  | Default |
|:-----------|:---------------------------------------------------------------------------------------------|:---|
| `BASE_DELAY`| Base value to generate the random delay, final delay is base value + random value,           | 0 |
| `SERVER_ID`| A friendly server id                                                                         | |
| `LOG_LEVEL`| To set the default log level ||