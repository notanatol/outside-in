package mindmap

import "testing"

func TestUploadFile(t *testing.T) {
	file := make(File, 2)

	app := new(App)

	addr := app.Upload(file)

	dFile, _ := app.Download(addr)

	if !file.Equals(dFile) {
		t.Fail()
	}
}

func TestUploadChunk(t *testing.T) {
	chunk := new(Chunk)

	file := File([]Chunk{*chunk})

	app := new(App)

	addr := app.Upload(file)

	dFile, _ := app.Download(addr)

	if !file.Equals(dFile) {
		t.Fail()
	}
}

func TestPinChunk(t *testing.T) {
	chunk := new(Chunk)

	file := File([]Chunk{*chunk})

	app := new(App)

	addr := app.Upload(file)

	ok := app.Pin(addr)

	// assert pinned content avoid eviction

	if !ok {
		t.Fail()
	}
}
