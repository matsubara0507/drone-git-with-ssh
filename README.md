# drone-git-with-ssh
Drone plugin: git with ssh

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-with-ssh
docker build --rm -t matsubara0507/git-with-ssh .
```

## Usage

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