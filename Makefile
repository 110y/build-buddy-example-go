BAZELISK_VERSION := 1.8.1

.PHONY: bazelisk
bazelisk:
	curl -sSL \
		"https://github.com/bazelbuild/bazelisk/releases/download/v$(BAZELISK_VERSION)/bazelisk-$(shell go env GOOS)-$(shell go env GOARCH)" \
		-o ./bin/bazelisk && \
		chmod +x ./bin/bazelisk
