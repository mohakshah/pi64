package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/bamarni/pi64/pkg/pi64"
	"github.com/cheggaaa/pb"
	"golang.org/x/crypto/openpgp"
)

var (
	tarPath = "/root/linux.tar.gz"
	pubKey  = `
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v2

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
=Euwm
-----END PGP PUBLIC KEY BLOCK-----
`
)

func main() {
	os.Exit(run())
}

func run() int {
	if os.Geteuid() != 0 {
		fmt.Fprintln(os.Stderr, "pi64-update must be run as root")
		return 1
	}

	keyring, err := openpgp.ReadArmoredKeyRing(strings.NewReader(pubKey))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't read keyring : "+err.Error())
		return 1
	}

	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	latestReleaseResp, err := client.Get("https://github.com/mohakshah/pi64-kernel/releases/latest")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't request for latest release : "+err.Error())
		return 1
	}

	latestRelease, err := latestReleaseResp.Location()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't parse latest release location : "+err.Error())
		return 1
	}

	metadata, err := pi64.ReadMetadata()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't read pi64 metadata.")
		return 1
	}

	currentVersion := metadata.KernelVersion
	if currentVersion == "" {
		currentVersion = metadata.Version
	}

	latestVersion := path.Base(latestRelease.String())

	if latestVersion <= currentVersion {
		fmt.Fprintln(os.Stderr, "You're already using the latest version.")
		return 0
	}

	releaseEndpoint := "https://github.com/mohakshah/pi64-kernel/releases/download/" + latestVersion

	fmt.Fprintf(os.Stderr, "Downloading '%s' release.\n", latestRelease)

	tarResp, err := http.Get(releaseEndpoint + "/linux.tar.gz")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get linux.tar.gz : "+err.Error())
		return 1
	}
	defer tarResp.Body.Close()

	tarFile, err := os.Create(tarPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create %s : %s\n", tarPath, err)
		return 1
	}
	defer tarFile.Close()
	defer os.Remove(tarPath)

	sig, err := http.Get(releaseEndpoint + "/linux.tar.gz.sig")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get linux.tar.gz.sig : "+err.Error())
		return 1
	}
	defer sig.Body.Close()

	bar := pb.New64(tarResp.ContentLength).SetUnits(pb.U_BYTES)
	bar.Start()

	// Wrap the response body to :
	// - a TeeReader to check against the PGP signature
	// - a proxy reader for the progress bar display
	reader := bar.NewProxyReader(io.TeeReader(tarResp.Body, tarFile))

	_, err = openpgp.CheckDetachedSignature(keyring, reader, sig.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't verify signature : "+err.Error())
		return 1
	}

	if err := exec.Command("tar", "-zxvf", tarPath, "-C", "/").Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't extract "+tarPath+" : "+err.Error())
		return 1
	}

	metadata.KernelVersion = latestVersion
	if err := pi64.WriteMetadata(metadata); err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't write metadata : "+err.Error())
	}

	fmt.Fprintln(os.Stderr, "Your kernel has been updated! You'll have to reboot for this to take effect.")
	return 0
}
