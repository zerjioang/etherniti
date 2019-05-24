// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

import (
	"bytes"
	"strconv"
)

type ServerStatusResponse struct {
	Architecture string  `json:"arch"`
	Os           string  `json:"os"`
	Cpus         Cpus    `json:"cpus"`
	Runtime      Runtime `json:"runtime"`
	Version      Version `json:"version"`
	Disk         Disk    `json:"disk"`
	Memory       Memory  `json:"memory"`
	Gc           Gc      `json:"gc"`
}

func (r *ServerStatusResponse) Reset() {
	r.Disk.Reset()
	r.Memory.Reset()
	r.Gc.Reset()
}

type Cpus struct {
	Cores int `json:"cores"`
}
type Runtime struct {
	Compiler string `json:"compiler"`
}
type Version struct {
	Etherniti string `json:"etherniti"`
	HTTP      string `json:"http"`
	Go        string `json:"go"`
}
type Disk struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func (disk *Disk) Reset() {
	disk.All = 0
	disk.Used = 0
	disk.Free = 0
}

type Memory struct {
	Frees     uint64 `json:"frees"`
	Heapalloc uint64 `json:"heapalloc"`
	Alloc     uint64 `json:"alloc"`
	Total     uint64 `json:"total"`
	Sys       uint64 `json:"sys"`
	Mallocs   uint64 `json:"mallocs"`
}

func (memory *Memory) Reset() {
	memory.Frees = 0
	memory.Heapalloc = 0
	memory.Alloc = 0
	memory.Total = 0
	memory.Sys = 0
	memory.Mallocs = 0
}

type Gc struct {
	Numgc       uint32 `json:"numgc"`
	NumForcedGC uint32 `json:"numForcedGC"`
}

func (gc *Gc) Reset() {
	gc.Numgc = 0
	gc.NumForcedGC = 0
}

func (r ServerStatusResponse) Bytes(buffer *bytes.Buffer) []byte {
	buffer.WriteString(`{"system":{"os":"`)
	buffer.WriteString(r.Os)
	buffer.WriteString(`","arch":"`)
	buffer.WriteString(r.Architecture)
	buffer.WriteString(`"},`)
	buffer.WriteString(`"cpus":{"cores":`)
	buffer.WriteString(itoa(r.Cpus.Cores))
	buffer.WriteString(`},"runtime":{"compiler":"`)
	buffer.WriteString(r.Runtime.Compiler)
	buffer.WriteString(`"},"version":{"etherniti":"`)
	buffer.WriteString(r.Version.Etherniti)
	buffer.WriteString(`","http":"`)
	buffer.WriteString(r.Version.HTTP)
	buffer.WriteString(`","go":"`)
	buffer.WriteString(r.Version.Go)
	buffer.WriteString(`"},"disk":{"all":`)
	buffer.WriteString(itoau64(r.Disk.All))
	buffer.WriteString(`,"used":`)
	buffer.WriteString(itoau64(r.Disk.Used))
	buffer.WriteString(`,"free":`)
	buffer.WriteString(itoau64(r.Disk.Free))
	buffer.WriteString(`},"memory":{"frees":`)
	buffer.WriteString(itoau64(r.Memory.Frees))
	buffer.WriteString(`,"heapalloc":`)
	buffer.WriteString(itoau64(r.Memory.Heapalloc))
	buffer.WriteString(`,"alloc":`)
	buffer.WriteString(itoau64(r.Memory.Alloc))
	buffer.WriteString(`,"total":`)
	buffer.WriteString(itoau64(r.Memory.Total))
	buffer.WriteString(`,"sys":`)
	buffer.WriteString(itoau64(r.Memory.Sys))
	buffer.WriteString(`,"mallocs":`)
	buffer.WriteString(itoau64(r.Memory.Mallocs))
	buffer.WriteString(`},"gc":{"numgc":`)
	buffer.WriteString(itoau32(r.Gc.Numgc))
	buffer.WriteString(`,"numForcedGC":`)
	buffer.WriteString(itoau32(r.Gc.NumForcedGC))
	buffer.WriteString(`}}`)
	return buffer.Bytes()
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
