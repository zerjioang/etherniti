package protocol

import (
	"bytes"
	"strconv"
	"unsafe"
)

type ServerStatusResponse struct {
	Cpus struct {
		Cores int `json:"cores"`
	} `json:"cpus"`
	Runtime struct {
		Compiler string `json:"compiler"`
	} `json:"runtime"`
	Version struct {
		Etherniti string `json:"etherniti"`
		HTTP      string `json:"http"`
		Go        string `json:"go"`
	} `json:"version"`
	Disk struct {
		All  uint64 `json:"all"`
		Used uint64 `json:"used"`
		Free uint64 `json:"free"`
	} `json:"disk"`
	Memory struct {
		Frees     uint64 `json:"frees"`
		Heapalloc uint64 `json:"heapalloc"`
		Alloc     uint64 `json:"alloc"`
		Total     uint64 `json:"total"`
		Sys       uint64 `json:"sys"`
		Mallocs   uint64 `json:"mallocs"`
	} `json:"memory"`
	Gc struct {
		Numgc       uint32 `json:"numgc"`
		NumForcedGC uint32 `json:"numForcedGC"`
	} `json:"gc"`
}

func (r ServerStatusResponse) Bytes() []byte {
	buffer.WriteString("x")
	data := `{"cpus":{"cores":`+itoa(r.Cpus.Cores)+`},"runtime":{"compiler":"`+r.Runtime.Compiler+`"},"version":{"etherniti":"`+r.Version.Etherniti+`","http":"`+r.Version.HTTP+`","go":"`+r.Version.Go+`"},"disk":{"all":`+itoau64(r.Disk.All)+`,"used":`+itoau64(r.Disk.Used)+`,"free":`+itoau64(r.Disk.Free)+`},"memory":{"frees":`+itoau64(r.Memory.Frees)+`,"heapalloc":`+itoau64(r.Memory.Heapalloc)+`,"alloc":`+itoau64(r.Memory.Alloc)+`,"total":`+itoau64(r.Memory.Total)+`,"sys":`+itoau64(r.Memory.Sys)+`,"mallocs":`+itoau64(r.Memory.Mallocs)+`},"gc":{"numgc":`+itoau32(r.Gc.Numgc)+`,"numForcedGC":`+itoau32(r.Gc.NumForcedGC)+`}}`
	return *(*[]byte)(unsafe.Pointer(&data))
}

func itoa(v int) string {
	return strconv.Itoa(v)
}

func itoau32(v uint32) string {
	return strconv.Itoa(int(v))
}

func itoau64(v uint64) string {
	return strconv.Itoa(int(v))
}

func tofloat64(v float64) string {
	//return fmt.Sprintf("%.6f", v)
	return strconv.FormatFloat(v, 'f', 6, 64)
}