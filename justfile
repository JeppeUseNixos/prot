version := env_var_or_default("APP_VERSION", "dev")
flags := "-s -w -X 'github.com/z3co/prot/cmd.Version={{version}}' -X 'github.com/z3co/prot/cmd.Commit=$(git rev-parse --short HEAD)'"
release:
	#!/usr/bin/env bash
	mkdir -p dist
	for os in linux darwin; do
		for arch in amd64 arm64; do
			echo "Building $os-$arch..."
			CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -ldflags="{{flags}}" -o prot
			tar -czf "dist/prot-{{version}}-$os-$arch.tar.gz" prot
			rm prot
		done
	done
	echo "Building windows-amd64..."
	GOOS=windows GOARCH=amd64 go build -ldflags="{{flags}}" -o prot.exe
	zip "dist/prot-{{version}}-windows-amd64.zip" prot.exe
	rm prot.exe

	echo "Generating checksum..."
	cd dist && sha256sum *.tar.gz *.zip > checksums.txt
	echo "Done!"

