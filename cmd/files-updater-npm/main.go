package main

import (
	npmUpdater "github.com/ted-vo/files-updater-npm/pkg/updater"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
	"github.com/ted-vo/semantic-release/v3/pkg/updater"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		FilesUpdater: func() updater.FilesUpdater {
			return &npmUpdater.Updater{}
		},
	})
}
