package packets_test

import (
	"encoding/json"
	"testing"

	"github.com/VirtuConnect/packets"
)

func TestParsing(t *testing.T) {
	packet := packets.StatusPacket("", packets.StatusSuccess)
	data, _ := json.Marshal(packet)
	var up packets.UnparsedPacket
	json.Unmarshal(data, &up)
	p, err := packets.ParsePacket(&up)

	if err != nil {
		t.Error(err)
	}

	result, err := packets.IsStatusSuccess(p)
	if err != nil {
		t.Error(err)
	}

	if !result {
		sp, _ := packets.GetStatusPacket(p)
		t.Error("Invalid ", sp)
	}
}
