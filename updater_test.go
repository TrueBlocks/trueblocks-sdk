package sdk

import (
	"os"
	"testing"
	"time"
)

func TestNeedsUpdateWithFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write some data to the file to set its modification time
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempFile.Name(), Type: File},
	}
	u, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTimeStamp to the current modification time of the file
	u.LastTimeStamp = fileModTime(tempFile.Name())

	// Modify the file to update its modification time
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Explicitly update the file's modification time
	newModTime := time.Now().Add(time.Second)
	if err := os.Chtimes(tempFile.Name(), newModTime, newModTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := u.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTimeStamp <= u.LastTimeStamp {
		t.Errorf("expected updated LastTimeStamp to be greater than original")
	}
}

func TestNeedsUpdateWithFolder(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	tempFile, err := os.CreateTemp(tempDir, "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Write some data to the file to set its modification time
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempDir, Type: Folder},
	}
	u, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTimeStamp to the current modification time of the file
	u.LastTimeStamp = fileModTime(tempFile.Name())

	// Modify the file to update its modification time
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Explicitly update the file's modification time
	newModTime := time.Now().Add(time.Second)
	if err := os.Chtimes(tempFile.Name(), newModTime, newModTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := u.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTimeStamp <= u.LastTimeStamp {
		t.Errorf("expected updated LastTimeStamp to be greater than original")
	}
}

func fileModTime(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.ModTime().Unix()
}
