
.PHONY: gen/sync-plugins-list
gen/sync-plugins-list:
	sh hack/sync-plugins-list.sh

.PHONY: gen/sync-codeowners
gen/sync-codeowners:
	sh hack/sync-codeowners.sh

.PHONY: init-plugin
init-plugin:
	sh hack/init-plugin.sh $(PLUGIN_DIR_NAME) $(CODEOWNERS)

# Imported from pipe-cd/pipecd 
.PHONY: lint/go
lint/go: FIX ?= false
lint/go: VERSION ?= sha256:c2f5e6aaa7f89e7ab49f6bd45d8ce4ee5a030b132a5fbcac68b7959914a5a890 # golangci/golangci-lint:v1.64.7
lint/go: FLAGS ?= --rm -e GOCACHE=/repo/.cache/go-build -e GOLANGCI_LINT_CACHE=/repo/.cache/golangci-lint -v ${PWD}:/repo -it
lint/go: MODULES ?= $(shell find . -name go.mod | while read -r dir; do dirname "$$dir"; done | paste -sd, -) # comma separated list of modules. e.g. MODULES=.,plugins/xxx-plugin
lint/go:
	@echo "Linting go modules..."
	@for module in $(shell echo $(MODULES) | tr ',' ' '); do \
		echo "Linting module: $$module"; \
		docker run ${FLAGS} -w /repo/$$module golangci/golangci-lint@${VERSION} golangci-lint run -v --config /repo/.golangci.yml --fix=$(FIX); \
	done
