package nsenter

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

// Config is the nsenter configuration used to generate
// nsenter command
type Config struct {
	Cgroup              bool   // Enter cgroup namespace
	FollowContext       bool   // Set SELinux security context
	GID                 int    // GID to use to execute given program
	IPC                 bool   // Enter IPC namespace
	Mount               bool   // Enter mount namespace
	Net                 bool   // Enter network namespace
	NoFork              bool   // Do not fork before executing the specified program
	PID                 bool   // Enter PID namespace
	PreserveCredentials bool   // Preserve current UID/GID when entering namespaces
	RootDirectory       string // Set the root directory, default to target process root directory
	Target              int    // Target PID (required)
	UID                 int    // UID to use to execute given program
	User                bool   // Enter user namespace
	UTS                 bool   // Enter UTS namespace
	WorkingDirectory    string // Set the working directory, default to target process working directory
}

// Execute the given program using the given nsenter configuration
// and return stdout/stderr or an error if command has failed
func (c *Config) Execute(program string, args ...string) (string, string, error) {
	cmd, err := c.buildCommand()
	if err != nil {
		return "", "", fmt.Errorf("Error while building command: %v", err)
	}

	// Prepare command
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Args = append(cmd.Args, program)
	cmd.Args = append(cmd.Args, args...)

	err = cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("Error while executing command: %v", err)
	}

	return stdout.String(), stderr.String(), nil
}

func (c *Config) buildCommand() (*exec.Cmd, error) {
	if c.Target == 0 {
		return nil, fmt.Errorf("Target must be specified")
	}

	var args []string
	args = append(args, "--target", strconv.Itoa(c.Target))

	if c.Cgroup {
		args = append(args, "--cgroup")
	}

	if c.FollowContext {
		args = append(args, "--follow-context")
	}

	if c.GID != 0 {
		args = append(args, "--setgid", strconv.Itoa(c.GID))
	}

	if c.IPC {
		args = append(args, "--ipc")
	}

	if c.Mount {
		args = append(args, "--mount")
	}

	if c.Net {
		args = append(args, "--net")
	}

	if c.NoFork {
		args = append(args, "--no-fork")
	}

	if c.PID {
		args = append(args, "--pid")
	}

	if c.PreserveCredentials {
		args = append(args, "--preserve-credentials")
	}

	if c.RootDirectory != "" {
		args = append(args, "--root", c.RootDirectory)
	}

	if c.UID != 0 {
		args = append(args, "--setuid", strconv.Itoa(c.UID))
	}

	if c.User {
		args = append(args, "--user")
	}

	if c.UTS {
		args = append(args, "--uts")
	}

	if c.WorkingDirectory != "" {
		args = append(args, "--wd", c.WorkingDirectory)
	}

	cmd := exec.Command("nsenter", args...)

	return cmd, nil
}
