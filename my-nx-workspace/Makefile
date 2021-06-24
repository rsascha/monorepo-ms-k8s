PROJECT = api
PATH_TO_PROJECT = apps
PROJECT_PATH = $(PATH_TO_PROJECT)/$(PROJECT)
VERSION_FILE = $(PROJECT_PATH)/k8s/version
OLD_VERSION = $(shell cat $(VERSION_FILE))
VERSION = $(shell echo $$(($(OLD_VERSION) + 1)))
FULL_VERSION = $(shell echo $(shell npm version --json | jq -r '.["my-nx-workspace"]').$(VERSION))
DOCKER_IMAGE = localhost:32000/$(PROJECT):$(FULL_VERSION)
DEPLOYMENT_TEMPLATE = $(PROJECT_PATH)/k8s/templates/deployment.yaml
DEPLOYMENT_TARGET = $(PROJECT_PATH)/k8s/overlays/dev/deployment.yaml

api-release:
	nx run api:build
	npx ncc build dist/apps/api/main.js -o dist/apps/api-ncc/
	docker build -f apps/api/Dockerfile -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	export DOCKER_IMAGE=$(DOCKER_IMAGE) \
		&& envsubst < $(DEPLOYMENT_TEMPLATE) > $(DEPLOYMENT_TARGET)
	@sed -i '1s/^/# ----------------------\n/' $(DEPLOYMENT_TARGET)
	@sed -i '2s/^/# --- GENERATED FILE ---\n/' $(DEPLOYMENT_TARGET)
	@sed -i '3s/^/# ----------------------\n/' $(DEPLOYMENT_TARGET)
	kustomize build $(PROJECT_PATH)/k8s/overlays/dev/ > $(PROJECT_PATH)/k8s/dev.yaml
	kubectl apply -f $(PROJECT_PATH)/k8s/dev.yaml
	@echo $(VERSION) > $(VERSION_FILE)
