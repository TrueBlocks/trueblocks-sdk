package sdk

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v5/pkg/version"

func Version() string {
	version := version.NewVersion(version.LibraryVersion)
	return version.String()
}
