package main

import (
	"log"
	"strings"
	"time"

	"bufio"
	"context"
	"github.com/TheBookPeople/avaya"
	"github.com/TheBookPeople/avaya/soap"
	"os"
)

func main() {
	ctx := context.Background()

	host := os.Getenv("AACC_HOST")
	if host == "" {
		log.Fatal("Env AACC_HOST not set")
	}

	// setup client
	client := avaya.NewClient("http", host)

	// check if chat is online
	online, err := client.IsSkillsetInService(ctx, "WC_Default_Skillset")
	if err != nil {
		log.Fatal(err)
	}
	if !online {
		log.Fatal("Offline.")
	}

	// start conversation
	conv, err := avaya.NewDirectConversation(ctx, client, "Test", "test@test2.co.uk", "WC_Default_Skillset")
	if err != nil {
		log.Fatal(err)
	}

	for answered := false; ; {

		messages, advisorTyping, err := conv.ReadMessages(ctx)

		// error soap.SOAPFault provides more detail
		if sf, ok := err.(*soap.SOAPFault); ok {
			if strings.Contains(sf.String, "errorcode: 2831") {
				log.Print("Waiting for an answer...")
			} else if strings.Contains(sf.String, "errorcode: 2833") {
				log.Fatal("Too many concurrent connections for this customer: ", sf.String)
			} else {
				log.Printf("Soapfault: %v", sf)
			}
		} else if err != nil {
			log.Fatal(err)
		}

		if advisorTyping {
			log.Println("Advisor typing...")
		}

		for _, m := range messages {
			if !answered {
				answered = true
				log.Println("answered!")
			}
			log.Println(m.Time, m.Type, m.Text, m.Hidden)
		}

		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				t := scanner.Text()
				if t == "quit" {
					conv.Close(ctx)
					return
				}
				if err := conv.WriteMessage(ctx, scanner.Text()); err != nil {
					log.Println("write error: %v", err)
				}
			}
		}()

		time.Sleep(1 * time.Second) // ReadMessages is non-blocking
	}
}
