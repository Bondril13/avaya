package CIWebCommsWs

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

func TestLastReadTime(t *testing.T) {

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "    ")
	enc.Encode(&ReadChatMessage{
		ContactID:  12345678,
		IsWriting:  false,
		SessionKey: "session key",
		LastReadTime: &CIDateTime{
			Milliseconds: 9876543,
		},
	})
	fmt.Println("moo")
}
