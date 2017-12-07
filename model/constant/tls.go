package constant

// define tls global define constant

var (
	CLOSE_NOTIFY = byte(0)
	RESET_FULL_HANDSHAKE = byte(1)
	UNEXPECTED_MESSAGE = byte(10)
	BAD_RECORD_MAC = byte(20)
	DECRYPTION_FAILED = byte(21)
	RECORD_OVERFLOW = byte(22)
	DECOMPRESSION_FAILURE = byte(30)
	HANDSHAKE_FAILURE = byte(40)
	BAD_CERTIFICATE = byte(42)
	UNSUPPORTED_CERTIFICATE = byte(43)
	CERTIFICATE_REVOKED = byte(44)
	CERTIFICATE_EXPIRED = byte(45)
	CERTIFICATE_UNKNOWN = byte(46)
	ILLEGAL_PARAMETER = byte(47)
	UNKNOWN_CA = byte(48)
	ACCESS_DENIED = byte(49)
	DECODE_ERROR = byte(50)
	DECRYPT_ERROR = byte(51)
)

var (
	HELLO_REQUEST = byte(0)
	CLIENT_HELLO = byte(1)
	SERVER_HELLO = byte(2)
	CERTIFICATE = byte(4)
	SERVER_KEY_EXCHANGE = byte(5)
	CERTIFICATE_REQUEST = byte(6)
	CHANGE_CIPHER_SPEC = byte(7)
	CERTIFICATE_VERIFY = byte(8)
	CLIENT_KEY_EXCHANGE = byte(9)
	FINISHED = byte(10)
	INIT_CONTENT = byte(11)
	RESOURCE_HASH = byte(12)
	RESOURCE_HASH_OPT = byte(13)
	RESOURCE_HASH_RES = byte(14)
	RESOURCE_H5_HASH = byte(15)
)