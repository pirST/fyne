//go:build !ios && !android

package mobile

import (
	"io"

	"github.com/pirST/fyne/v2"
	intRepo "github.com/pirST/fyne/v2/internal/repository"
	"github.com/pirST/fyne/v2/storage/repository"
)

func deleteURI(u fyne.URI) error {
	// no-op as we use the internal FileRepository
	return nil
}

func existsURI(fyne.URI) (bool, error) {
	// no-op as we use the internal FileRepository
	return false, nil
}

func nativeFileOpen(*fileOpen) (io.ReadCloser, error) {
	// no-op as we use the internal FileRepository
	return nil, nil
}

func nativeFileSave(*fileSave) (io.WriteCloser, error) {
	// no-op as we use the internal FileRepository
	return nil, nil
}

func registerRepository(d *driver) {
	repository.Register("file", intRepo.NewFileRepository())
}
