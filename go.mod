module github.com/jmhodges/kustomize-age

go 1.18

require (
	filippo.io/age v1.0.0-beta4
	sigs.k8s.io/kustomize/api v0.6.0
	sigs.k8s.io/yaml v1.2.0
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/sys v0.0.0-20191002063906-3421d5a6bb1c // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

exclude (
	github.com/pkg/errors v0.8.0
	github.com/pkg/errors v0.8.1
	github.com/pkg/errors v0.9.0
	github.com/russross/blackfriday v2.0.0+incompatible
	sigs.k8s.io/kustomize/api v0.2.0
	sigs.k8s.io/kustomize/cmd/config v0.2.0
)
