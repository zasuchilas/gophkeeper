package app

import "log/slog"

type buildInfo struct {
	version string
	date    string
	commit  string
}

func NewBuildInfo(version, date, commit string) *buildInfo {
	return &buildInfo{version: version, date: date, commit: commit}
}

// buildInfo output build info.
func (a *app) buildInfoOutput() {
	slog.Info("Build info:")
	slog.Info("* app: gophkeeper server")
	slog.Info("* version: %s \n", a.build.version)
	slog.Info("* date: %s \n", a.build.date)
	slog.Info("* commit: %s \n", a.build.commit)
}
