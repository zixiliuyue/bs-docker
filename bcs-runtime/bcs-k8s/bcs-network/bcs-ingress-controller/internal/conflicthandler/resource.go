

package conflicthandler

// portSegment [start,end)
type portSegment struct {
	Start     int
	End       int
	Protocols []string
}

// return true if conflict
func (ps *portSegment) isConflictWithPort(isTCPUDPReuse bool, p portStruct) bool {
	if p.val >= ps.Start && p.val < ps.End {
		if isProtocolConflict(isTCPUDPReuse, ps.Protocols, p.Protocols) {
			return true
		}
	}

	return false
}

// return true if conflict
func (ps *portSegment) isConflictWithPortSeg(isTCPUDPReuse bool, newPs portSegment) bool {
	if max(ps.Start, newPs.Start) < min(ps.End, newPs.End) {
		if isProtocolConflict(isTCPUDPReuse, ps.Protocols, newPs.Protocols) {
			return true
		}
	}

	return false
}

type portStruct struct {
	val       int
	Protocols []string
}

// return true if conflict
func (p *portStruct) isConflictWithPort(isTCPUDPReuse bool, newPort portStruct) bool {
	if p.val == newPort.val {
		if isProtocolConflict(isTCPUDPReuse, p.Protocols, newPort.Protocols) {
			return true
		}
	}

	return false
}

// return true if conflict
func (p *portStruct) isConflictWithPortSeg(isTCPUDPReuse bool, ps portSegment) bool {
	return ps.isConflictWithPort(isTCPUDPReuse, *p)
}

type resource struct {
	usedPort        map[int]portStruct
	usedPortSegment []portSegment
}

func newResource() *resource {
	return &resource{
		usedPort:        make(map[int]portStruct),
		usedPortSegment: make([]portSegment, 0),
	}
}

// IsConflict return true if conflict with otherRes
func (r *resource) IsConflict(isTCPUDPReuse bool, otherRes *resource) bool {
	if otherRes == nil {
		return false
	}

	for portVal, port := range r.usedPort {
		if otherPort, ok := otherRes.usedPort[portVal]; ok {
			if port.isConflictWithPort(isTCPUDPReuse, otherPort) {
				return true
			}
		}

		for _, otherPortSeg := range otherRes.usedPortSegment {
			if port.isConflictWithPortSeg(isTCPUDPReuse, otherPortSeg) {
				return true
			}
		}
	}

	for _, portSeg := range r.usedPortSegment {
		for _, port := range otherRes.usedPort {
			if portSeg.isConflictWithPort(isTCPUDPReuse, port) {
				return true
			}
		}

		for _, otherPortSeg := range otherRes.usedPortSegment {
			if portSeg.isConflictWithPortSeg(isTCPUDPReuse, otherPortSeg) {
				return true
			}
		}
	}

	return false
}
