package fs

import (
	"fmt"

	"github.com/cyverse/go-irodsclient/irods/connection"
	irods_fs "github.com/cyverse/go-irodsclient/irods/fs"
	"github.com/cyverse/go-irodsclient/irods/types"
	"github.com/cyverse/go-irodsclient/irods/util"
)

// FileHandle ...
type FileHandle struct {
	FileSystem  *FileSystem
	Connection  *connection.IRODSConnection
	IRODSHandle *types.IRODSFileHandle
	Entry       *FSEntry
	Offset      int64
	OpenMode    types.FileOpenMode
}

// GetOffset returns current offset
func (handle *FileHandle) GetOffset() int64 {
	return handle.Offset
}

// IsReadMode returns true if file is opened with read mode
func (handle *FileHandle) IsReadMode() bool {
	return types.IsFileOpenFlagRead(handle.OpenMode)
}

// IsWriteMode returns true if file is opened with write mode
func (handle *FileHandle) IsWriteMode() bool {
	return types.IsFileOpenFlagWrite(handle.OpenMode)
}

// Close closes the file
func (handle *FileHandle) Close() error {
	defer handle.FileSystem.Session.ReturnConnection(handle.Connection)

	if handle.IsWriteMode() {
		handle.FileSystem.invalidateCachePath(handle.Entry.Path)
		handle.FileSystem.invalidateCachePath(util.GetIRODSPathDirname(handle.Entry.Path))
	}

	return irods_fs.CloseDataObject(handle.Connection, handle.IRODSHandle)
}

// Seek moves file pointer
func (handle *FileHandle) Seek(offset int64, whence types.Whence) (int64, error) {
	newOffset, err := irods_fs.SeekDataObject(handle.Connection, handle.IRODSHandle, offset, whence)
	if err != nil {
		return newOffset, err
	}

	handle.Offset = newOffset
	return newOffset, nil
}

// Read reads the file
func (handle *FileHandle) Read(length int) ([]byte, error) {
	if !handle.IsReadMode() {
		return nil, fmt.Errorf("File is opened with %s mode", handle.OpenMode)
	}

	bytes, err := irods_fs.ReadDataObject(handle.Connection, handle.IRODSHandle, length)
	if err != nil {
		return nil, err
	}

	handle.Offset += int64(len(bytes))
	return bytes, nil
}

// Write writes the file
func (handle *FileHandle) Write(data []byte) error {
	if !handle.IsWriteMode() {
		return fmt.Errorf("File is opened with %s mode", handle.OpenMode)
	}

	err := irods_fs.WriteDataObject(handle.Connection, handle.IRODSHandle, data)
	if err != nil {
		return err
	}

	handle.Offset += int64(len(data))

	// update
	if handle.Entry.Size < handle.Offset+int64(len(data)) {
		handle.Entry.Size = handle.Offset + int64(len(data))
	}

	return nil
}

// ToString stringifies the object
func (handle *FileHandle) ToString() string {
	return fmt.Sprintf("<FileHandle %d %s %s %s>", handle.Entry.ID, handle.Entry.Type, handle.Entry.Name, handle.OpenMode)
}
