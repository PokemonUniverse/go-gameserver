package netmsg

import (
	pnet "github.com/PokemonUniverse/nonamelib/network"
)

const (
	LOGINSTATUS_IDLE            = 0
	LOGINSTATUS_WRONGACCOUNT    = 1
	LOGINSTATUS_SERVERERROR     = 2
	LOGINSTATUS_DATABASEERROR   = 3
	LOGINSTATUS_ALREADYLOGGEDIN = 4
	LOGINSTATUS_READY           = 5
	LOGINSTATUS_CHARBANNED      = 6
	LOGINSTATUS_SERVERCLOSED    = 7
	LOGINSTATUS_WRONGVERSION    = 8
	LOGINSTATUS_FAILPROFILELOAD = 9
)

type LoginMessage struct {
	// Receive
	Username      string
	Password      string
	ClientVersion int // uint16

	// Send
	Status int // uint32
}

// GetHeader returns the header value of this message
func (m *LoginMessage) GetHeader() uint8 {
	return pnet.HEADER_LOGIN
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *LoginMessage) ReadPacket(_packet pnet.IPacket) error {
	var err error
	if m.Username, err = _packet.ReadString(); err != nil {
		return err
	}
	if m.Password, err = _packet.ReadString(); err != nil {
		return err
	}
	clientVersion, err := _packet.ReadUint16()
	if err != nil {
		return err
	}
	m.ClientVersion = int(clientVersion)

	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *LoginMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint32(uint32(m.Status))

	return packet
}
