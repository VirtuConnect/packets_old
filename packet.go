package packets

import (
	"encoding/json"
	"fmt"
)

type Packet struct {
	PacketType string      `json:"packetType"`
	Body       interface{} `json:"body"`
}

type UnparsedPacket struct {
	PacketType string                 `json:"packetType"`
	Body       map[string]interface{} `json:"body"`
}

const (
	TypePacketReadingError    = "RedingError"
	TypePacketParsingError    = "ParsingError"
	TypePacketProcessingError = "ProcessingError"
	TypeCommandResponsePacket = "CommandResponse"
	TypeStatusResponsePacket  = "StatusResponse"
	TypeTaskLaunchPacket      = "TaskLaunchPacket"
	TypeErrorPacket           = "Error"
)

var PacketReadingErrorPacket = Packet{
	PacketType: TypePacketReadingError,
	Body:       "Failed to read packet",
}

var PacketParsingErrorPacket = Packet{
	PacketType: TypePacketParsingError,
	Body:       "Failed to parse the packets body (missing fields or invalid data)",
}

var HelloWorldPacket = Packet{
	PacketType: "HelloWorld",
	Body:       "Hello its me",
}

var PacketProcessingErrorPacket = Packet{
	PacketType: TypePacketParsingError,
	Body:       "Failed to process package. Unkown Failure",
}

func ParsePacket(packet *UnparsedPacket) (*Packet, error) {
	result := &Packet{PacketType: packet.PacketType}

	//convert the unparsed body into json
	jsonData, err := json.Marshal(packet.Body)
	if err != nil {
		return nil, err
	}

	//switch the different cases
	switch packet.PacketType {
	case TypeCommandResponsePacket:
		var commandResponse CommandResponsePacket
		err = json.Unmarshal(jsonData, &commandResponse)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}

		result.Body = commandResponse
	case TypeStatusResponsePacket:
		var statusResponse StatusResponsePacket
		err = json.Unmarshal(jsonData, &statusResponse)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}
		result.Body = statusResponse
	case TypeTaskLaunchPacket:
		var launchPacket TaskLaunchPacket
		err = json.Unmarshal(jsonData, &launchPacket)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}
		result.Body = launchPacket
	case TypePlayVideoRequestPacket:
		var packet PlayVideoRequestPacket
		err = json.Unmarshal(jsonData, &packet)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}
		result.Body = packet
	case TypePlayAudioRequestPacket:
		var packet PlayAudioRequestPacket
		err = json.Unmarshal(jsonData, &packet)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}
		result.Body = packet
	case TypePlayYoutubeRequestPacket:
		var packet PlayYoutubeRequestPacket
		err = json.Unmarshal(jsonData, &packet)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}
		result.Body = packet
	case TypeRunCommandRequestPacket:
		var packet RunCommandRequestPacket
		err = json.Unmarshal(jsonData, &packet)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to CommandResponsePacket: %w", err)
		}
		result.Body = packet
	case TypeErrorPacket:
		var errorResponse ErrorPacket
		err = json.Unmarshal(jsonData, &errorResponse)

		if err != nil {
			return nil, fmt.Errorf("error unmarshaling body to ErrorPacket: %w", err)
		}
		result.Body = errorResponse
	default:
		return nil, fmt.Errorf("invalid packet content type `%s`", packet.PacketType)
	}
	return result, nil
}
