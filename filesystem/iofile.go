package filesystem

import "io/fs"

type IOFile struct {
	Fser
}

func NewIOFile(fs Fser) IOFile {
	return IOFile{Fser: fs}
}

func (iof IOFile) Open(name string) (fs.File, error) {
	const op = "open"

	if !fs.ValidPath(name) {
		return nil, iof.wrapError(op, name, fs.ErrInvalid)
	}

	file, err := iof.Fser.Open(name)
	if err != nil {
		return nil, iof.wrapError(op, name, err)
	}

	if _, ok := file.(fs.ReadDirFile); !ok {
		return nil, iof.wrapError(op, name, fs.ErrInvalid)
	}

	return file, nil
}

func (iof IOFile) ReadFile(name string) ([]byte, error) {
	const op = "readfile"

	if !fs.ValidPath(name) {
		return nil, iof.wrapError(op, name, fs.ErrInvalid)
	}

	cxt, err := ReadFile(iof.Fser, name)
	if err != nil {
		return nil, iof.wrapError(op, name, err)
	}

	return ctx, nil
}

func (IOFile) wrapError(op, path string, err error) error {

	if _, ok := err.(*fs.PathError); ok {
		return err
	}

	return &fs.PathError{
		Op:   op,
		Path: path,
		Err:  err,
	}
}
