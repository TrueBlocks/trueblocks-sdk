package sdk

import "github.com/TrueBlocks/trueblocks-chifra/v6/pkg/version"

func Version() string {
	version := version.NewVersion(version.LibraryVersion)
	return version.String()
}
