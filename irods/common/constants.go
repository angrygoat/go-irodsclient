package common

// constants
const (
	// VERSION
	IRODSVersionRelease string = "4.3.0"
	IRODSVersionAPI     string = "d"

	// Magic Numbers
	MaxQueryRows        int = 500
	MaxPasswordLength   int = 50
	ReadWriteBufferSize int = 64 * 1024

	/*
		MAX_SQL_ATTR               int = 50
		MAX_PATH_ALLOWED           int = 1024
		MAX_NAME_LEN               int = MAX_PATH_ALLOWED + 64
		MAX_SQL_ROWS               int = 256
		DEFAULT_CONNECTION_TIMEOUT int = 120
	*/

	// Magic Strings
	/*
		AUTH_SCHEME_KEY string = "a_scheme"
		AUTH_USER_KEY   string = "a_user"
		AUTH_PWD_KEY    string = "a_pw"
		AUTH_TTL_KEY    string = "a_ttl"

		GSI_AUTH_PLUGIN string = "GSI"
		GSI_OID         string = "1.3.6.1.4.1.3536.1.1" // taken from http://j.mp/2hDeczm

		PAM_AUTH_PLUGIN string = "PAM"
	*/

	// Environment variables
	TestingPropertiesPath string = "TESTING_PROPERTIES"
)
