# Prot
Prot is a project based todolist the todolist is connected to the folder you are in and the git branch you are on

## Installation
Nix 
```sh
nix profile add github:JeppeUseNixos/prot
```
Other linux 
```sh
curl -sSL https://raw.githubusercontent.com/JeppeUseNixos/prot/refs/heads/main/install.sh | sh
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
