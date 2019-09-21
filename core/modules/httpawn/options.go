package httpawn

var (
	// default options that are loaded if
	// no options are defined
	DefaultOptions = PawnOptions{
		EnablePipelining: false,
	}
)

// options allowed by httpawn
type PawnOptions struct {
	EnablePipelining bool `json:"pipelining"`
}
