package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	pkgNames, err := getAURPackages(ctx)
	if err != nil {
		panic(err)
	}

	localAurPkgs, err := getLocalAURPackageVersions(ctx, pkgNames)
	if err != nil {
		panic(err)
	}

	remoteAurPkgs, err := getRemoteAURPackageVersions(ctx, pkgNames)
	if err != nil {
		panic(err)
	}

	fmt.Printf("local pkg version: %+v\n", localAurPkgs)
	fmt.Printf("remote pkg version: %+v\n", remoteAurPkgs)
}

func getRemoteAURPackageVersions(ctx context.Context, pkgNames []string) ([]aurPkg, error) {
	url := "https://aur.archlinux.org/rpc/?v=5&type=info"
	for _, name := range pkgNames {
		url = fmt.Sprintf("%s&arg[]=%s", url, name)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create http request for url %s", url)
	}
	req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get package info for url %s", url)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	jsonResp := &aurWebRPCResult{}
	if err = json.Unmarshal(data, jsonResp); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}
	if jsonResp.Type == "error" {
		return nil, errors.Errorf("AUR Web RPC return error %s", jsonResp.Error)
	}
	result := make([]aurPkg, len(jsonResp.Results))
	for _, r := range jsonResp.Results {
		if r.Name != "" {
			result = append(result, aurPkg{Name: r.Name, Version: r.Version})
		}
	}
	return result, nil
}

type aurPkg struct {
	Name, Version string
}

type aurWebRPCResult struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	ResultCount int    `json:"resultcount"`
	Results     []struct {
		Name    string `json:"Name"`
		Version string `json:"Version"`
	} `json:"results"`
	Error string `json:"error"`
}

func getLocalAURPackageVersions(ctx context.Context, pkgNames []string) ([]aurPkg, error) {
	result := make([]aurPkg, len(pkgNames))
	g, ctx := errgroup.WithContext(ctx)
	for i, pkg := range pkgNames {
		i, pkg := i, pkg
		g.Go(func() error {
			cmd := exec.CommandContext(ctx, "pacman", "-Q", pkg)
			data, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrapf(err, "failed ot get output of pacman -Q %s", pkg)
			}
			result[i] = aurPkg{Name: pkg, Version: strings.TrimSpace(strings.Replace(string(data), pkg, "", -1))}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}

func getAURPackages(ctx context.Context) ([]string, error) {
	cmd := exec.CommandContext(ctx, "pacman", "-Qqm")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrapf(err, "failed ot get output of pacman -Qqm")
	}
	result := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}
