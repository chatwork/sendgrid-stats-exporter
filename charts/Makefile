KUBERNETES_VERSION="1.18.1"
RELEASE = $$(basename $$PWD)

.PHONY: install
install:
	helm upgrade -i --wait $(RELEASE) .

.PHONY: lint
lint: lint-default

.PHONY: lint-default
lint-default:
	@echo "=> Linting default values.yaml"
	helm lint --strict
	@echo "=> Validating default value.yaml"
	helm template . | kubeval --kubernetes-version $(KUBERNETES_VERSION)
	@echo ""

.PHONY: test
test:
	helm test --logs $(RELEASE)

.PHONY: uninstall
uninstall:
	helm uninstall $(RELEASE)