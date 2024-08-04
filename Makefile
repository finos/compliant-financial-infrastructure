# Define the default target
.PHONY: ccc-wireframe

ccc-wireframe:
	@echo "Please provide the YAML component definition file:"
	@read -p "Component Definition Path (URL or local): " YAML_FILE; \
	read -p "CSP shortname (max 8 characters):" CSP_NAME; \
	read -p "Service shortname (max 8 characters): " SERVICE_NAME; \
	read -p "Version (max 8 characters): " VERSION; \
	cd generators/ccc && go run ccc-cfi-generator.go $$YAML_FILE $$CSP_NAME $$SERVICE_NAME $$VERSION && \
	mv $$CSP_NAME ../../; \
	echo "To re-run this command: go run ccc-cfi-generator.go $$YAML_FILE $$CSP_NAME $$SERVICE_NAME $$VERSION"

# Optional: A help target to display usage information
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run - Run the Go program with user-provided arguments"
