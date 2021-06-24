// +build mage

/*

=== Build and deploy the applications to K8s ===

Sorry - it's mega bashy, but better than a Makefile!
I'm using bash, cat, jq, sed, ...
Feel free to remove all linux commands, so that this can be executed on Windows :)

Debugging:
	mage -v release:api
*/

package main

import (
	"strconv"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Release mg.Namespace
type Debug mg.Namespace

// Build and release 'api'.
func (Release) Api() {
	build("api")
	compress("api")
	dockerBuild("api")
	release("api")
}

// Build and release 'movie-name-generator'.
func (Release) MovieNameGenerator() {
	build("movie-name-generator")
	dockerBuild("movie-name-generator")
	release("movie-name-generator")
}

// Perform just the docker build, push, version bump.
func (Debug) MovieNameGeneratorDockerBuild() {
	dockerBuild("movie-name-generator")
}

// Release just k8s changes (kustomize and kubectl apply).
func (Debug) MovieNameGeneratorK8s() {
	release("movie-name-generator")
}

// Just perform the nx build.
func build(project string) {
	sh.Run("nx", "run", project+":build")
}

// Compress NodeJS code and include all related NodeJS Modules.
// This is important, because we have a single packages.json and a single node_modules folder for all applications!
// 'ncc' will include needed packages only.
func compress(project string) {
	sh.Run("npx", "ncc", "build", "dist/apps/"+project+"/main.js", "-o", "dist/apps/"+project+"-ncc/")
}

func dockerBuild(project string) {
	projectPath := "apps/" + project
	versionFile := projectPath + "/k8s/version"
	oldVersion, _ := sh.Output("cat", versionFile)
	oldVersionInt, _ := strconv.Atoi(oldVersion)
	buildVersion := strconv.Itoa(oldVersionInt + 1)
	npmVersion, _ := sh.Output("bash", "-c", "npm version --json | jq -r '.[\"my-nx-workspace\"]'")
	fullVersion := npmVersion + "." + buildVersion
	dockerImage := "localhost:32000/" + project + ":" + fullVersion
	deploymentTemplate := projectPath + "/k8s/templates/deployment.yaml"
	deploymentTarget := projectPath + "/k8s/overlays/dev/deployment.yaml"

	sh.Run("docker", "build", "-f", "apps/"+project+"/Dockerfile", "-t", dockerImage, ".")
	sh.Run("docker", "push", dockerImage)
	sh.Run("bash", "-c", "export DOCKER_IMAGE="+dockerImage+" && envsubst < "+deploymentTemplate+" > "+deploymentTarget)
	sh.Run("bash", "-c", "echo "+buildVersion+" > "+versionFile)
}

func release(project string) {
	projectPath := "apps/" + project

	sh.Run("bash", "-c", "kustomize build "+projectPath+"/k8s/overlays/dev/ > "+projectPath+"/k8s/dev.yaml")
	sh.Run("kubectl", "apply", "-f", projectPath+"/k8s/dev.yaml")
}
