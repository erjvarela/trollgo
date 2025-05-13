package sshAgent

import (
	"golang.org/x/crypto/ssh"

	"github.com/erjvarela/trollgo/internal/config"
)

type SshConfig struct{}

func InitSSHConfiguration(Cfg *config.Configuration) (*ssh.ClientConfig, error) {
	sshConfig := &ssh.ClientConfig{
		User: Cfg.SshDefaultUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(Cfg.SshDefaultPassword), // Only offer password authentication
		},
		// WARNING: This is insecure and bypasses server host key verification.
		// Only use for trusted local networks or testing.
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return sshConfig, nil
}
