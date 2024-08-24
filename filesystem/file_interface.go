package filesystem

import (
	"io"
	"os"
)

// Filer is a interface, represent a file in filesystem
type Filer interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Stat() (os.FileInfo, error)
	WriteString(s string) (ret int, err error)
}

// Fser is the filesystem interface
type Fser interface {
	// Create create a file, returning the filer and an error, if any happens
	Create(name string) (Filer, error)

	// Mkdir create a directory, return error is any happens
	Mkdir(name string, perm os.FileMode) error

	// Open opens a file, returning the filer and an error, if any happens
	Open(name string) (Filer, error)

	// OpenFile opens a file using the given flags and the given filemode, returning
	// the filer and an error, if any happens
	OpenFile(name string, flag int, perm os.FileMode) (Filer, error)

	// Remove removes a file identified name, returning an error if any happens
	// It doest not fail if the path does not exist(return nil)
	Remove(name string) error

	// RemoveAll removes a directory path and any children it contains. It doest
	// not fail if the path does not exist(return nil)
	RemoveAll(path string) error

	// Rename renames a file name, returning an error if any happens
	Rename(oldname, newname string) error

	// Stat returns an FileInfo describing the identified name, or an error if any happen
	Stat(name string) (os.FileInfo, error)

	// Name returning the name of the file
	Name() string

	// Chmod changes the mode of the named file to perm(FileMode), returning an error
	// if any happens
	Chmod(name string, perm os.FileMode) error

	// Chown changes the uid and gid of the identified file, returning an error if
	// any happens
	Chown(name string, uid, gid int) error
}
