package netipx

import (
	"net"
	"net/netip"
)

// FromStdIPRaw returns an IP from the standard library's IP type.
// If std is invalid, ok is false.
// Unlike FromStdIP, FromStdIPRaw does not do an implicit Unmap if
// len(std) == 16 and contains an IPv6-mapped IPv4 address.
//
// Deprecated: use netip.AddrFromSlice instead.
func FromStdIPRaw(std net.IP) (ip netip.Addr, ok bool) {
	return netip.AddrFromSlice(std)
}

// AddrNext returns the IP following ip.
// If there is none, it returns the IP zero value.
//
// Deprecated: use netip.Addr.Next instead.
func AddrNext(ip netip.Addr) netip.Addr {
	addr := u128From16(ip.As16()).addOne()
	if ip.Is4() {
		if uint32(addr.lo) == 0 {
			// Overflowed.
			return netip.Addr{}
		}
		return addr.IP4()
	} else {
		if addr.isZero() {
			// Overflowed
			return netip.Addr{}
		}
		return addr.IP6().WithZone(ip.Zone())
	}
}

// AddrPrior returns the IP before ip.
// If there is none, it returns the IP zero value.
//
// Deprecated: use netip.Addr.Prev instead.
func AddrPrior(ip netip.Addr) netip.Addr {
	addr := u128From16(ip.As16())
	if ip.Is4() {
		if uint32(addr.lo) == 0 {
			return netip.Addr{}
		}
		return addr.subOne().IP4()
	} else {
		if addr.isZero() {
			return netip.Addr{}
		}
		return addr.subOne().IP6().WithZone(ip.Zone())
	}
}

// Valid reports whether r.From() and r.To() are both non-zero and
// obey the documented requirements: address families match, and From
// is less than or equal to To.
//
// Deprecated: use the correctly named and identical IsValid method instead.
func (r IPRange) Valid() bool { return r.IsValid() }
