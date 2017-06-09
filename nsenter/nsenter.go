package nsenter

// Config is the nsenter configuration used to generate
// nsenter command
type Config struct {
	Cgroup              bool   // Enter cgroup namespace
	FollowContext       bool   // Set SELinux security context
	GID                 uint   // GID to use to execute given program
	IPC                 bool   // Enter IPC namespace
	Mount               bool   // Enter mount namespace
	Net                 bool   // Enter network namespace
	NoFork              bool   // Do not fork before executing the specified program
	PID                 bool   // Enter PID namespace
	PreserveCredentials bool   // Preserve current UID/GID when entering namespaces
	RootDirectory       string // Set the root directory, default to target process root directory
	Target              uint   // Target PID
	UID                 uint   // UID to use to execute given program
	User                bool   // Enter user namespace
	UTS                 bool   // Enter UTS namespace
	WorkingDirectory    string // Set the working directory, default to target process working directory
}
