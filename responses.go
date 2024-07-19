package packets

import "fmt"

type ResponsePaket struct {
	RequestId string      `json:"requestId"`
	Type      string      `json:"type"`
	Body      interface{} `json:"body"`
}

type CommandResponsePacket struct {
	Content string `json:"content"`
}

type StatusResponsePacket struct {
	Status string `json:"status"`
}

type TaskLaunchPacket struct {
	TaskId string `json:"TaskId"`
}

func StatusPacket(request string, status string) *Packet {
	return &Packet{
		PacketType: TypeStatusResponsePacket,
		Body: ResponsePaket{
			RequestId: request,
			Body: StatusResponsePacket{
				Status: status,
			},
		},
	}
}

func GetStatusPacket(packet *Packet) (Status, error) {
	if packet.PacketType != TypeStatusResponsePacket {
		return StatusFailure, fmt.Errorf("packed %p is not status packet", packet)
	}
	return ParseStatus(packet.Body.(StatusResponsePacket).Status)
}

func IsStatusSuccess(packet *Packet) (bool, error) {
	status, err := GetStatusPacket(packet)
	if err != nil {
		return false, err
	}
	isSuccess := status == StatusSuccess
	return isSuccess, nil
}

func GetResponsePacket(packet *Packet) (*CommandResponsePacket, error) {
	if packet.PacketType != TypeCommandResponsePacket {
		return nil, fmt.Errorf("provided packet is not a commandresponse packet")
	}
	commandRef, _ := packet.Body.(CommandResponsePacket)
	return &commandRef, nil
}
