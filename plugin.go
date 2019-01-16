package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Plugin struct {
	Home     string
	SSHKey   string
	Hosts    []string
	Commands []string
}

func (p Plugin) Exec() error {
	sshDir := fmt.Sprintf("%s/.ssh", p.Home)
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		return errors.Wrapf(err, "Exec cmd: mkdir -p %s", sshDir)
	}
	log.Infof("Exec cmd: mkdir -p %s", sshDir)

	if err := ioutil.WriteFile(fmt.Sprintf("%s/id_rsa", sshDir), []byte(p.SSHKey), 0600); err != nil {
		return errors.Wrapf(err, "Write file: %s/id_rsa", sshDir)
	}
	log.Infof("Write file: %s/id_rsa", sshDir)

	if err := p.writeSSHConfigFile(sshDir); err != nil {
		return err
	}
	log.Infof("Write file: %s/config", sshDir)

	for _, command := range p.Commands {
		out, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
		if err != nil {
			return errors.Wrapf(err, "Exec cmd: %s: %s", command, out)
		}
		log.Infof("Exec cmd: %s:  %s", command, string(out))
	}

	return nil
}

func (p Plugin) writeSSHConfigFile(sshDir string) (err error) {
	sshConfigFile, err := os.OpenFile(fmt.Sprintf("%s/config", sshDir), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return errors.Wrapf(err, "Open file: %s/config", sshDir)
	}

	defer func() {
		if closeErr := sshConfigFile.Close(); closeErr != nil {
			err = errors.Wrapf(closeErr, "Close file: %s/config", sshDir)
		}
	}()

	for _, host := range p.Hosts {
		if _, err := fmt.Fprintf(sshConfigFile, "Host %s\n\tStrictHostKeyChecking no\n\n", host); err != nil {
			return errors.Wrapf(err, "Write file: %s/config (host: %s)", sshDir, host)
		}
	}
	return nil
}