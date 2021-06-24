
# https://kubectl.docs.kubernetes.io/installation/kustomize/binaries/
install-kustomize: 
	curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
	sudo mv ./kustomize /usr/bin/

# https://kubectl.docs.kubernetes.io/installation/kubectl/binaries/
install-kubectl:
	curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(shell curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/bin/
	