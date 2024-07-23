package packets

import "fmt"

type ResponsePacket struct {
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

func StatusPacket(request string, status Status) *Packet {
	return &Packet{
		PacketType: TypeResponsePacket,
		Body: ResponsePacket{
			RequestId: request,
			Type:      TypeStatusResponsePacket,
			Body: StatusResponsePacket{
				Status: statusToString[status],
			},
		},
	}
}

func GetStatusPacket(packet *Packet) (Status, error) {
	if packet.PacketType != TypeResponsePacket {
		return -1, fmt.Errorf("packed %p is not status packet", packet)
	}
	res, ok := packet.Body.(ResponsePacket)
	if !ok {
		return -1, fmt.Errorf("invalid sturcture (expected response packet)")
	}
	pack, ok := res.Body.(StatusResponsePacket)
	if !ok {
		return -1, fmt.Errorf("invalid sturcture (expected statusresponse packet)")
	}
	return ParseStatus(pack.Status)
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
	if packet.PacketType != TypeResponsePacket {
		return nil, fmt.Errorf("provided packet is not a commandresponse packet")
	}
	ress, _ := packet.Body.(ResponsePacket)
	commandRef, _ := ress.Body.(CommandResponsePacket)
	return &commandRef, nil
}
