package servers

type host struct {
	clear string
	onion string
}

func createHostWithClear(
	clear string,
) Host {
	return createHostInternally(clear, "")
}

func createHostWithOnion(
	onion string,
) Host {
	return createHostInternally("", onion)
}

func createHostInternally(
	clear string,
	onion string,
) Host {
	out := host{
		clear: clear,
		onion: onion,
	}

	return &out
}

// String returns the string representation of the host
func (obj *host) String() string {
	if obj.IsClear() {
		return obj.clear
	}

	return obj.onion
}

// IsClear returns true if there is a clear host, false otherwise
func (obj *host) IsClear() bool {
	return obj.clear != ""
}

// Clear returns the clear host, if any
func (obj *host) Clear() string {
	return obj.clear
}

// IsOnion returns true if there is an onion host, false otherwise
func (obj *host) IsOnion() bool {
	return obj.onion != ""
}

// Onion returns the onion host, if any
func (obj *host) Onion() string {
	return obj.onion
}
