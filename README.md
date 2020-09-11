kustomize-age is a plugin for kubernetes to decrypt files with age before using
them as secret contents. It's a placeholder until sops, and kustomize-sops can
support age.

You may have to fork this repo and adjust its `go.mod` in order to correct
package mismatches with your `kustomize` binary.

To install, run `make install` and it will build this as a Go plugin and install it into your `$XDG_CONFIG_HOME`