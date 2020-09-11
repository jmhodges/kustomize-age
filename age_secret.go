// Copyright 2020 Jeffrey M Hodges
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"filippo.io/age"
	"sigs.k8s.io/kustomize/api/kv"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

type plugin struct {
	h                *resmap.PluginHelpers
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	types.SecretArgs

	AgeKeyEnvVar string `json:"ageKeyEnvVar,omitempty yaml:"ageKeyEnvVar,omitempty" protobuf:"bytes,1,opt,name=ageKeyEnvVar`
}

//noinspection GoUnusedGlobalVariable
//nolint: golint
var KustomizePlugin plugin

func (p *plugin) Config(h *resmap.PluginHelpers, c []byte) error {
	p.h = h
	p.AgeKeyEnvVar = strings.TrimSpace(p.AgeKeyEnvVar)
	if p.AgeKeyEnvVar == "" {
		p.AgeKeyEnvVar = "KUBE_AGE_KEY"
	}
	return yaml.Unmarshal(c, p)
}

func (p *plugin) Generate() (resmap.ResMap, error) {
	args := p.SecretArgs
	args.Name = p.ObjectMeta.Name
	args.Namespace = p.ObjectMeta.Namespace

	key := strings.TrimSpace(os.Getenv(p.AgeKeyEnvVar))
	if key == "" {
		return nil, fmt.Errorf("The env var for the age key, %#v, is empty", p.AgeKeyEnvVar)
	}
	ident, err := age.ParseX25519Identity(key)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse the age key env var %#v as an age key: %w", p.AgeKeyEnvVar, err)
	}
	for _, file := range args.FileSources {
		encFilePath := file + ".age"
		secretFile := filepath.Join(p.h.Loader().Root(), encFilePath)
		f, err := os.Open(secretFile)
		if err != nil {
			return nil, fmt.Errorf("unable to open the file %s to match the requested OndiskEncryptedSecret file %s: %w", encFilePath, file, err)
		}
		r, err := age.Decrypt(f, ident)
		if err != nil {
			log.Fatalf("error: cannot decode file %s :: %v", secretFile, err)
		}
		out := &bytes.Buffer{}
		if _, err := io.Copy(out, r); err != nil {
			return nil, fmt.Errorf("Failed to read encrypted file: %v", err)
		}
		args.LiteralSources = append(args.LiteralSources, file+"="+out.String())
	}
	args.FileSources = nil

	return p.h.ResmapFactory().FromSecretArgs(
		kv.NewLoader(p.h.Loader(), p.h.Validator()), args)
}
