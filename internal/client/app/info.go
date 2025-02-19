package app

import (
	"fmt"
	"strings"
)

type buildInfo struct {
	version string
	date    string
	commit  string
}

func NewBuildInfo(version, date, commit string) *buildInfo {
	return &buildInfo{version: version, date: date, commit: commit}
}

func (b *buildInfo) String() string {
	var version string
	if b.version != "" {
		version = "ver. " + b.version
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s %s", version, b.date, b.commit))
}
