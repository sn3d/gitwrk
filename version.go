package gitwrk

var version string = "devel"

// GetVersion returns you version of Gitwk. If you get the 'devel'
// version, that means you're running directly from source code. The
// actual version is getting here via tag and ldflag:
//     go build -ldflags="-X 'github.com/unravela/gitwrk.version=v1.0.0'"
func GetVersion() string {
	return version
}
