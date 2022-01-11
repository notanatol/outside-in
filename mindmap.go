package mindmap

type Chunk byte

type File []Chunk

func (f *File) Equals(f2 *File) bool {
	return false
}

type (
	network interface {
		Sync(Chunk) error
	}
	localstore interface {
		Put(Chunk) error
		Has(*ID) (bool, error)
		Get(*ID) (Chunk, error)
	}
)

type App struct {
	network    network
	localstore localstore
}

type ID byte

func (a *App) Upload(f File) *ID {
	for _, chunk := range f {
		if mustStay(chunk) {
			_ = a.localstore.Put(chunk)
			continue
		}

		a.network.Sync(chunk)
	}

	return new(ID)
}

func mustStay(Chunk) bool {
	return false
}

func (a *App) Download(addr *ID) (*File, error) {
	if found, _ := a.localstore.Has(addr); found {
		chunk, _ := a.localstore.Get(addr)
		f := File([]Chunk{chunk})
		return &f, nil
	}

	// sync from network...

	return new(File), nil
}

func (a *App) Pin(addr *ID) bool {
	return false
}
