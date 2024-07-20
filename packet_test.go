package packets_test

import (
	"encoding/json"
	"testing"

	"github.com/VirtuConnect/packets"
)

func TestParsing(t *testing.T) {
	packet := packets.StatusPacket("", packets.StatusNotAllowed.ToString())
	data, _ := json.Marshal(packet)
	var up packets.UnparsedPacket
	json.Unmarshal(data, &up)
	_, err := packets.ParsePacket(&up)

	if err != nil {
		t.Error(err)
	}
}
