package protocol

//snoflake based composited id
// show flake generates 8 bytes
type DatumId [8 + 1 + 8]byte

type ItemKey struct {
	Left  []byte
	Sep   []byte
	Right []byte
}
