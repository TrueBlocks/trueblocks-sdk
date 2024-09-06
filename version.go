package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/version"

func Version() string {
	version := version.NewVersion(version.LibraryVersion)
	return version.String()
}
