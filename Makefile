BINARY_NAME=simplecli

release:
	@echo "Creating release..."
	@read -p "Version: " version; \
	git tag $$version; \
	git push origin $$version