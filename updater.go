package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

type UpdateType string

const (
	Timer      UpdateType = "timer"
	File       UpdateType = "file"
	Folder     UpdateType = "folder"
	FileSize   UpdateType = "fileSize"
	FolderSize UpdateType = "folderSize"
)

func (u *UpdateType) String() string {
	return map[UpdateType]string{
		Timer:      "Timer",
		File:       "File",
		Folder:     "Folder",
		FileSize:   "FileSize",
		FolderSize: "FolderSize",
	}[*u]
}

type Updater struct {
	Name          string        `json:"name"`
	LastTimeStamp int64         `json:"lastTimeStamp"`
	LastTotalSize int64         `json:"lastTotalSize"`
	Items         []UpdaterItem `json:"items"`
}

type UpdaterItem struct {
	Path     string        `json:"path"`
	Duration time.Duration `json:"duration"`
	Type     UpdateType    `json:"type"`
}

func (u *Updater) String() string {
	bytes, _ := json.MarshalIndent(u, "", "  ")
	return string(bytes)
}

// NewUpdater creates a new Updater instance. It logs a fatal error if invalid parameters are provided.
func NewUpdater(name string, items []UpdaterItem) (Updater, error) {
	if len(items) == 0 {
		logger.Fatal("must provide at least one UpdaterItem")
	}

	now := time.Now()
	ret := Updater{
		Name:          name,
		Items:         items,
		LastTimeStamp: now.Unix(),
		LastTotalSize: 0,
	}

	ret.debugV("NewUpdater created with LastTimeStamp:", ret.LastTimeStamp, "LastTotalSize:", ret.LastTotalSize)
	return ret, nil
}

// NeedsUpdate checks if the updater needs to be updated based on the paths or duration.
//
// Summary:
// - If the Timer type has exceeded its duration, it will trigger an update.
// - If the File or Folder types find any file whose modification time is later than LastTimeStamp, it will trigger an update.
// - If the FileSize or FolderSize items notice an increase in the total file size of all FileSize or FolderSize items, it will trigger an update.
//
// Returns:
// - Updater: A potentially updated Updater instance with updated LastTimeStamp and/or LastTotalSize.
// - bool: A boolean indicating whether an update is needed (true) or not (false).
// - error: An error object containing any errors encountered during the update check.
func (u *Updater) NeedsUpdate() (Updater, bool, error) {
	var needsUpdate bool
	var errs []error
	var mostRecentTime int64
	var totalSize int64

	u.debugV("NeedsUpdate called with LastTimeStamp:", u.LastTimeStamp, "LastTotalSize:", u.LastTotalSize)

	for _, item := range u.Items {
		switch item.Type {
		case Timer:
			updated, needed, err := u.needsUpdateTime(item.Duration)
			if err != nil {
				errs = append(errs, err)
			}
			if needed {
				needsUpdate = true
				if updated.LastTimeStamp > mostRecentTime {
					mostRecentTime = updated.LastTimeStamp
				}
			}
		case File:
			updated, needed, err := u.needsUpdateFiles(item.Path)
			if err != nil {
				errs = append(errs, err)
			}
			if needed {
				needsUpdate = true
				if updated.LastTimeStamp > mostRecentTime {
					mostRecentTime = updated.LastTimeStamp
				}
			}
		case Folder:
			updated, needed, err := u.needsUpdateFolder(item.Path)
			if err != nil {
				errs = append(errs, err)
			}
			if needed {
				needsUpdate = true
				if updated.LastTimeStamp > mostRecentTime {
					mostRecentTime = updated.LastTimeStamp
				}
			}
		case FileSize:
			// Skip FileSize type items in this loop
			continue
		case FolderSize:
			// Skip FolderSize type items in this loop
			continue
		default:
			logger.Fatal("unknown path type" + fmt.Sprintf(" %s", item.Type))
			return Updater{}, false, errors.New("unknown path type")
		}
	}

	// Handle FileSize and FolderSize type items in a separate loop
	for _, item := range u.Items {
		switch item.Type {
		case FileSize:
			fileSize := file.FileSize(item.Path)
			totalSize += fileSize
		case FolderSize:
			if err := walk.ForEveryFileInFolder(item.Path, func(filePath string, _ any) (bool, error) {
				fileSize := file.FileSize(filePath)
				totalSize += fileSize
				return true, nil
			}, nil); err != nil {
				errs = append(errs, fmt.Errorf("failed to process folder %s: %v", item.Path, err))
			}
		}
	}
	u.debugV("Total size calculated:", totalSize, "LastTotalSize:", u.LastTotalSize)

	if totalSize != u.LastTotalSize {
		u.debug(mark("File size condition met", u.LastTotalSize, totalSize))
		needsUpdate = true
	}

	if needsUpdate {
		newUpdater := *u
		if mostRecentTime == 0 {
			mostRecentTime = time.Now().Unix()
		}
		newUpdater.LastTimeStamp = mostRecentTime
		newUpdater.LastTotalSize = totalSize
		u.debugV("Updating LastTimeStamp to:", newUpdater.LastTimeStamp, "and LastTotalSize to:", newUpdater.LastTotalSize)
		return newUpdater, true, combineErrors(errs)
	}

	return *u, false, combineErrors(errs)
}

// needsUpdateTime checks if the specified duration has passed since the last update.
func (u *Updater) needsUpdateTime(duration time.Duration) (Updater, bool, error) {
	now := time.Now().Unix()
	u.debugV("Current time:", now, "Last update timestamp:", u.LastTimeStamp, "Duration (seconds):", duration.Seconds())
	if now-int64(duration.Seconds()) >= u.LastTimeStamp {
		u.debug(mark("Duration condition met", u.LastTimeStamp, now))
		newUpdater := *u
		newUpdater.LastTimeStamp = now
		return newUpdater, true, nil
	}
	u.debugV("Duration condition not met, no update needed")
	return *u, false, nil
}

// needsUpdateFiles checks if the file's modification time is more recent than the last update.
func (u *Updater) needsUpdateFiles(path string) (Updater, bool, error) {
	u.debugV("Checking files for updates")
	modTime, err := file.GetModTime(path)
	if err != nil {
		return *u, false, fmt.Errorf("failed to get modification time for file %s: %v", path, err)
	}
	u.debugV("File:", relativize(path), "Modification time:", modTime.Unix())
	if modTime.Unix() > u.LastTimeStamp {
		u.debug(mark("File time condition met", u.LastTimeStamp, modTime.Unix()))
		newUpdater := *u
		newUpdater.LastTimeStamp = modTime.Unix()
		return newUpdater, true, nil
	}
	u.debugV("File modification condition not met, no update needed")
	return *u, false, nil
}

// needsUpdateFolder checks if any file within the folder has a modification time more recent than the last update.
func (u *Updater) needsUpdateFolder(path string) (Updater, bool, error) {
	u.debugV("Checking folders for updates")
	var maxLastTs int64
	var errs []error

	if err := walk.ForEveryFileInFolder(path, func(filePath string, _ any) (bool, error) {
		modTime, err := file.GetModTime(filePath)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get modification time for file %s: %v", filePath, err))
			return true, nil
		}
		u.debugV("File in folder:", relativize(filePath), "Modification time:", modTime.Unix())
		if modTime.Unix() > maxLastTs {
			maxLastTs = modTime.Unix()
		}
		return true, nil
	}, nil); err != nil {
		errs = append(errs, fmt.Errorf("failed to process folder %s: %v", path, err))
	}

	if maxLastTs > u.LastTimeStamp {
		u.debug(mark("Folder modification condition met", u.LastTimeStamp, maxLastTs))
		newUpdater := *u
		newUpdater.LastTimeStamp = maxLastTs
		return newUpdater, true, combineErrors(errs)
	}

	u.debugV("Folder modification condition not met, no update needed")
	return *u, false, combineErrors(errs)
}

// combineErrors combines multiple errors into a single error.
func combineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	errStrings := make([]string, len(errs))
	for i, err := range errs {
		errStrings[i] = err.Error()
	}
	return errors.New(strings.Join(errStrings, "\n"))
}

// SetChain updates the paths by replacing an old chain segment with a new one.
func (u *Updater) SetChain(oldChain, newChain string) {
	newItems := []UpdaterItem{}
	for _, item := range u.Items {
		ps := string(os.PathSeparator)
		oc := ps + oldChain
		nc := ps + newChain
		newPath := strings.ReplaceAll(item.Path, oc, nc)
		newItems = append(newItems, UpdaterItem{Path: newPath, Duration: item.Duration, Type: item.Type})
	}
	u.Items = newItems
	u.Reset()
}

// Reset sets the LastTimeStamp and LastTotalSize to 0 which causes a reload on the next call to NeedsUpdate.
func (u *Updater) Reset() {
	u.debugV("Reset called, setting LastTimeStamp and LastTotalSize to 0")
	u.LastTimeStamp = 0
	u.LastTotalSize = 0
}

var debugging bool
var debugType = ""
var debugVerbose = false

func init() {
	debugType = os.Getenv("TB_DEBUG_UPDATE")
	debugging = len(debugType) > 0
	debugVerbose = os.Getenv("TB_VERBOSE") == "true"
}

func mark(msg string, t1, t2 int64) string {
	return fmt.Sprintf("%s%s: updating ts %d ----> %d%s", colors.BrightRed, msg, t1, t2, colors.Off)
}

func (u *Updater) debug(args ...interface{}) {
	if debugging && (debugType == "true" || strings.Contains(debugType, u.Name)) {
		head := colors.Green + fmt.Sprintf("%10.10s:", u.Name) + colors.BrightYellow
		modifiedArgs := append([]interface{}{head}, args...)
		modifiedArgs = append(modifiedArgs, colors.Off)
		logger.Info(modifiedArgs...)
	}
}

// debugV is a verbose debug function that calls u.debug.
func (u *Updater) debugV(args ...interface{}) {
	if debugVerbose {
		u.debug(args...)
	}
}

// relativize modifies the given path by relativizing it with the specified partial paths.
func relativize(path string) string {
	partialPaths := []string{
		"/Users/jrush/Data/trueblocks/v1.0.0/cache/",
		"/Users/jrush/Data/trueblocks/v1.0.0/unchained/",
		"/Users/jrush/Library/Application Support/TrueBlocks/",
	}

	for _, partialPath := range partialPaths {
		if after, ok := strings.CutPrefix(path, partialPath); ok {
			return "./" + after
		}
	}

	return path
}
