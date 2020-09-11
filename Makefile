
XDG_CONFIG_HOME ?= $(HOME)/.config

.PHONY: install clean

all: OndiskEncryptedSecret.so

OndiskEncryptedSecret.so:
	go build -trimpath -buildmode plugin -o OndiskEncryptedSecret.so age_secret.go

install_dir = $(XDG_CONFIG_HOME)/kustomize/plugin/kustomize-age/v1/ondiskencryptedsecret/

install: OndiskEncryptedSecret.so
	mkdir -p $(install_dir) && cp OndiskEncryptedSecret.so $(install_dir)

clean:
	rm -f OndiskEncryptedSecret.so