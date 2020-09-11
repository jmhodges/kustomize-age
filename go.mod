module github.com/jmhodges/kustomize-age

go 1.14

require (
	filippo.io/age v1.0.0-beta4
	sigs.k8s.io/kustomize/api v0.6.0
	sigs.k8s.io/yaml v1.2.0
)

exclude (
	github.com/pkg/errors v0.8.0
	github.com/pkg/errors v0.8.1
	github.com/pkg/errors v0.9.0
	github.com/russross/blackfriday v2.0.0+incompatible
	sigs.k8s.io/kustomize/api v0.2.0
	sigs.k8s.io/kustomize/cmd/config v0.2.0
)
