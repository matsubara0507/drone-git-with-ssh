# drone-git-with-ssh

[![Build Status](https://cloud.drone.io/api/badges/matsubara0507/drone-git-with-ssh/status.svg)](https://cloud.drone.io/matsubara0507/drone-git-with-ssh)
[![GoDoc](https://godoc.org/github.com/matsubara0507/drone-git-with-ssh?status.svg)](https://godoc.org/github.com/matsubara0507/drone-git-with-ssh)
[![Go Report Card](https://goreportcard.com/badge/github.com/matsubara0507/drone-git-with-ssh)](https://goreportcard.com/report/github.com/matsubara0507/drone-git-with-ssh)
[![](https://images.microbadger.com/badges/image/matsubara0507/git-with-ssh.svg)](https://microbadger.com/images/matsubara0507/git-with-ssh "Get your own image badge on microbadger.com")

Drone plugin: git with ssh

## Build

Build the binary with the following commands:

```
go build
```

## Test

use drone cli

```
drone exec --event pull_request
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-with-ssh
docker build --rm -t matsubara0507/git-with-ssh .
```

## Usage

```
$ ./drone-git-with-ssh --help
NAME:
   git-with-ssh plugin - git-with-ssh plugin

USAGE:
   drone-git-with-ssh [global options] command [command options] [arguments...]

VERSION:
   1.0.0+0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --ssh-private-key value  SSH private key for git [$PLUGIN_SSH_PRIVATE_KEY]
   --hosts value            hosts by git with ssh (disable StrictHostKeyChecking) [$PLUGIN_SSH_HOSTS]
   --commands value         shell commands [$PLUGIN_COMMANDS]
   --home value             home directory for ssh (default: "/root") [$PLUGIN_HOME]
   --env-file value         source env file
   --help, -h               show help
   --version, -v            print the version
```

Execute from the working directory:

```sh
docker run --rm \
  -e PLUGIN_SSH_PRIVATE_KEY=ssh_key \
  -e PLUGIN_SSH_HOSTS=github.com,bitbucket.org \
  -e PLUGIN_COMMANDS='git clone git@github.com:matsubara0507/drone-git-with-ssh.git' \
  -v $(pwd):/root/work \
  -w /root/work \
  matsubara0507/git-with-ssh
```
