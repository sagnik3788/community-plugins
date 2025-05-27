
.PHONY: gen/sync-plugins-list
gen/sync-plugins-list:
	sh hack/sync-plugins-list.sh

.PHONY: gen/sync-codeowners
gen/sync-codeowners:
	sh hack/sync-codeowners.sh

.PHONY: init-plugin
init-plugin:
	sh hack/init-plugin.sh $(PLUGIN_DIR_NAME) $(CODEOWNERS)
