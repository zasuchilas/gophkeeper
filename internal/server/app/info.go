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
	slog.Info("gophkeeper server",
		slog.String("version", a.build.version),
		slog.String("date", a.build.date),
		slog.String("commit", a.build.commit),
	)
}
