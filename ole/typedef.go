package ole

type Handle uintptr
type HWND uintptr

type COAUTHIDENTITY struct {
	User           *uint16
	UserLength     uint32
	Domain         *uint16
	DomainLength   uint32
	Password       *uint16
	PasswordLength uint32
	Flags          uint32
}

type COAUTHINFO struct {
	AuthnSvc           uint32
	AuthzSvc           uint32
	ServerPrincName    *uint16
	AuthnLevel         uint32
	ImpersonationLevel uint32
	AuthIdentityData   *COAUTHIDENTITY
	Capabilities       uint32
}

type COSERVERINFO struct {
	Reserved1 uint32
	Aame      *uint16
	AuthInfo  *COAUTHINFO
	Reserved2 uint32
}

type BIND_OPTS3 struct {
	CbStruct          uint32
	Flags             uint32
	Mode              uint32
	TickCountDeadline uint32
	TrackFlags        uint32
	ClassContext      uint32
	Locale            uint32
	ServerInfo        *COSERVERINFO
	Hwnd              HWND
}
