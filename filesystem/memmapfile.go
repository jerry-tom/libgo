import (
	"fmt"
)

type MemMapFile struct {
	mu   sync.RWMutex
	data map[string]*mem.FileData
	init sync.Once
}