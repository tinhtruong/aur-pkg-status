package aur

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

// PackageUpdate represent a package update information
type PackageUpdate struct {
	Name, InstalledVersion, LatestVersion string
}

// GetPackageUpdates return a list of package update
func GetPackageUpdates(ctx context.Context) ([]PackageUpdate, error) {
	pkgNames, err := getInstalledPackageNames(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get installed AUR package names")
	}

	installedPkgs, err := getInstalledPackages(ctx, pkgNames)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get installed AUR package info")
	}

	latestPkgs, err := getLatestPackages(ctx, pkgNames)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get latest AUR package info")
	}
	result := make([]PackageUpdate, 0)
	for _, name := range pkgNames {
		update := PackageUpdate{Name: name}
		if installedPkg, found := filterByName(installedPkgs, name); found {
			update.InstalledVersion = installedPkg.Version
		}

		if latestPkg, found := filterByName(latestPkgs, name); found {
			update.LatestVersion = latestPkg.Version
		}
		result = append(result, update)

	}
	return result, nil
}

func filterByName(pkgs []aurPackage, name string) (aurPackage, bool) {
	for _, pkg := range pkgs {
		if pkg.Name == name {
			return pkg, true
		}
	}
	return aurPackage{}, false
}

type aurPackage struct {
	Name, Version string
}

func getLatestPackages(ctx context.Context, pkgNames []string) ([]aurPackage, error) {
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
	jsonResp := &webRPCResult{}
	if err = json.Unmarshal(data, jsonResp); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}
	if jsonResp.Type == "error" {
		return nil, errors.Errorf("AUR Web RPC return error %s", jsonResp.Error)
	}
	result := make([]aurPackage, len(jsonResp.Results))
	for _, r := range jsonResp.Results {
		if r.Name != "" {
			result = append(result, aurPackage{Name: r.Name, Version: r.Version})
		}
	}
	return result, nil
}

type webRPCResult struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	ResultCount int    `json:"resultcount"`
	Results     []struct {
		Name    string `json:"Name"`
		Version string `json:"Version"`
	} `json:"results"`
	Error string `json:"error"`
}

func getInstalledPackages(ctx context.Context, pkgNames []string) ([]aurPackage, error) {
	result := make([]aurPackage, len(pkgNames))
	g, ctx := errgroup.WithContext(ctx)
	for i, pkg := range pkgNames {
		i, pkg := i, pkg
		g.Go(func() error {
			cmd := exec.CommandContext(ctx, "pacman", "-Q", pkg)
			data, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrapf(err, "failed ot get output of pacman -Q %s", pkg)
			}
			result[i] = aurPackage{Name: pkg, Version: strings.TrimSpace(strings.Replace(string(data), pkg, "", -1))}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}

func getInstalledPackageNames(ctx context.Context) ([]string, error) {
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
