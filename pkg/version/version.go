package version

var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func GetVersion() string {
	return Version
}

func GetBuildInfo() string {
	return "Version: " + Version + "\nBuild Time: " + BuildTime + "\nGit Commit: " + GitCommit
}
