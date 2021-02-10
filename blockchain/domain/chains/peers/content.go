package peers

import "fmt"

type content struct {
	normal Server `hydro:"Normal, Normal"`
	tor    Server `hydro:"Tor, Tor"`
}

func createContentWithNormal(
	normal Server,
) Content {
	return createContentInternally(normal, nil)
}

func createContentWithTor(
	tor Server,
) Content {
	return createContentInternally(nil, tor)
}

func createContentInternally(
	normal Server,
	tor Server,
) Content {
	out := content{
		normal: normal,
		tor:    tor,
	}

	return &out
}

// String returns the content as string
func (obj *content) String() string {
	if obj.IsNormal() {
		return fmt.Sprintf(protocolPattern, NormalProtocol, protocolSeparator, obj.Normal().String())
	}

	return fmt.Sprintf(protocolPattern, TorProtocol, protocolSeparator, obj.Tor().String())
}

// IsNormal returns true if there is a normal server, false otherwise
func (obj *content) IsNormal() bool {
	return obj.normal != nil
}

// Normal returns the normal server, if any
func (obj *content) Normal() Server {
	return obj.normal
}

// IsTor returns true if there is a tor server, false otherwise
func (obj *content) IsTor() bool {
	return obj.tor != nil
}

// Tor returns the tor server, if any
func (obj *content) Tor() Server {
	return obj.tor
}
