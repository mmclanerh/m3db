// DefaultMmapConfiguration is the default mmap configuration.
func DefaultMmapConfiguration() MmapConfiguration {
	return MmapConfiguration{
		HugeTLB: MmapHugeTLBConfiguration{
			Enabled:   true,    // Enable when on a platform that supports
			Threshold: 2 << 14, // 32kb and above mappings use huge pages
		},
	}
}


	// Mmap is the mmap options which features are primarily platform dependent
	Mmap *MmapConfiguration `yaml:"mmap"`
}

// MmapConfiguration is the mmap configuration.
type MmapConfiguration struct {
	// HugeTLB is the huge pages configuration which will only take affect
	// on platforms that support it, currently just linux
	HugeTLB MmapHugeTLBConfiguration `yaml:"hugeTLB"`
}

// MmapHugeTLBConfiguration is the mmap huge TLB configuration.
type MmapHugeTLBConfiguration struct {
	// Enabled if true or disabled if false
	Enabled bool `yaml:"enabled"`

	// Threshold is the threshold on which to use the huge TLB flag if enabled
	Threshold int64 `yaml:"threshold"`

// MmapConfiguration returns the effective mmap configuration.
func (p FilesystemConfiguration) MmapConfiguration() MmapConfiguration {
	if p.Mmap == nil {
		return DefaultMmapConfiguration()
	}
	return *p.Mmap
}