## Installation
Nix 
```sh
nix profile add github:z3co/protv2
```
Other linux 
```sh
curl -sSL https://raw.githubusercontent.com/z3co/protv2/refs/heads/main/install.sh | sh
```
Windows
Download the zip and place the exe in your path

## Building

Clone the repo and run the following commands

Nix:
```sh
nix build .#dev
```
or to build from github
```sh
nix build
```

Other:
```sh
go build -o prot
```

## Make a new release

Update the version in package.nix, refresh the nix shell
update the version in install.sh and run ```just release```
