package packets

type ErrorPacket struct {
	Reason string `json:"reason"`
}

func Error(reason string) *Packet {
	return &Packet{
		PacketType: TypeErrorPacket,
		Body:       ErrorPacket{Reason: reason},
	}
}
