package stddatabase

type Config struct {
	// ScanEmptyAsError for
	ScanEmptyAsError bool `toml:"empty_error"`
	ScanPrint        bool `toml:"scan"`
	ValuePrint       bool `toml:"value"`
}
