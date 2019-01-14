FROM docker:git

ADD release/linux/amd64/drone-git-with-ssh /bin/
ENTRYPOINT ["/bin/drone-git-with-ssh"]