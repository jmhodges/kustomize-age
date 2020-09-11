kustomize-age is a plugin for `kustomize` that allows the user to keep local
files encrypted with `age` that are decrypted before being used as `files`
inside a kubernetes `Secret`.

Unlike kustomize-sops it does not encrypt the whole `Secret` yaml file, but just
the files mentioned in the `Secret`'s `files` attribute.

Use
===

First, run `make install` and it will build this as a Go plugin and install it
into your `$XDG_CONFIG_HOME`. (The Makefile is small, so you can read it to see
exactly what it's doing.)

Then, you'll create your `Secret` yaml file normally but with a few changes:

* `apiVersion` should be set to `apiVersion: kustomize-age/v1`
* `kind` should be set to `OndiskEncryptedSecret` instead of `Secret` (though,
  the final object that's added to your k8s cluster will be a `Secret`).
* `files` should be the file paths you want decrypted and stored in the `Secret`
  without the `.age` suffix on them. For example, a `mycredentials.key` in
  `files` of the `OndiskEncryptedSecret` will correspond to the age encrypted
  file on disk at `mycredentials.key.age`)

Those `.age` file paths will be the file paths you wanted included in the
`Secret` but encrypted with the `age` key specified.

By default, the `age` key used to decrypt these files is expected to be in the
`KUBE_AGE_KEY` environment variable. The environment variable used can be
overrided by setting the `ageKeyEnvVar` in the `OndiskEncryptedSecret`.

Example
=======

Supposing you have a file path you want in your secret at
`mycredentials.key.age`

```
$ ls
mycredentials.key.age mycredentials.yaml

$ cat mycredentials.yaml
apiVersion: kustomize-age/v1
kind: OndiskEncryptedSecret
metadata:
    name: some-svc-credentials
files:
    - mycredentials.key

# likely, you'll want to delete this next line and use the
# default KUBE_AGE_KEY env var.
ageKeyEnvVar: ANOTHER_ENV_VAR
```

Caveat
======
As with many Go plugins, you may have to fork this repo and adjust its `go.mod`
in order to correct package mismatches with your `kustomize` binary.