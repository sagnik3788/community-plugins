
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

# Imported from pipe-cd/pipecd 
.PHONY: build/plugins
build/plugins: PLUGINS_BIN_DIR ?= ~/.piped/plugins
build/plugins: PLUGINS_SRC_DIR ?= ./plugins
build/plugins: PLUGINS_OUT_DIR ?= ${PWD}/.artifacts/plugins
build/plugins: PLUGINS ?= $(shell find $(PLUGINS_SRC_DIR) -mindepth 1 -maxdepth 1 -type d | while read -r dir; do basename "$$dir"; done | paste -sd, -) # comma separated list of plugins. eg: PLUGINS=kubernetes,ecs,lambda
build/plugins:
	@echo "PLUGINS: $(PLUGINS)"
	mkdir -p $(PLUGINS_BIN_DIR)
	@echo "Building plugins..."
	@for plugin in $(shell echo $(PLUGINS) | tr ',' ' '); do \
		if [ ! -f $(PLUGINS_SRC_DIR)/$$plugin/go.mod ]; then \
			echo "Skipped plugin: $$plugin (no go.mod found)"; \
			continue; \
		fi; \
		echo "Building plugin: $$plugin"; \
		go -C $(PLUGINS_SRC_DIR)/$$plugin build -o $(PLUGINS_OUT_DIR)/$$plugin . \
			&& cp $(PLUGINS_OUT_DIR)/$$plugin $(PLUGINS_BIN_DIR)/$$plugin; \
		if [ $$? -ne 0 ]; then \
			echo "Failed to build plugin: $$plugin"; \
		fi; \
	done
	@echo "Plugins are built to $(PLUGINS_OUT_DIR) and copied to $(PLUGINS_BIN_DIR)"

# .PHONY: test/go

# .PHONY: push/plugins

# .PHONY: precommit