package chat

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	nats "github.com/nats-io/go-nats"
)

type PrivMsg struct {
	badges      []string
	color       string
	displayName string
	emotes      string
	id          string
	mod         bool
	roomid      string
	subscriber  bool
	tmisentts   string
	turbo       bool
	userid      string
	usertype    string
	roomname    string
	message     string
}

func string2Prv(nc *nats.EncodedConn, msg string) {

	chatspl := strings.Split(msg, "PRIVMSG")

	message := strings.Split(chatspl[1], ":")
	var stream = strings.Replace(message[0], " ", "", 2)
	var chatMsg = message[1]
	meta := strings.Split(chatspl[0], ":")
	var chatmeta = meta[0]

	splitMsg := strings.Split(chatmeta, ";")

	var m map[string]string

	m = make(map[string]string)

	for i := range splitMsg {
		key, value := splitAfterQuals(splitMsg[i])
		m[key] = value
	}

	privmsg := &PrivMsg{
		message:     chatMsg,
		roomname:    stream,
		color:       m["color"],
		badges:      strings.Split(m["@badges"], ","),
		displayName: m["display-name"],
		emotes:      m["emotes"], //strings.Split(m["emotes"], ","),
		subscriber:  intToBool(m["subscriber"]),
		mod:         intToBool(m["mod"]),
		turbo:       intToBool(m["turbo"]),
		id:          m["id"],
		roomid:      m["room-id"],
		tmisentts:   m["tmi-sent-ts"],
		userid:      m["user-id"],
	}
	// log.Println(privmsg)
	_ = privmsg

	ch := &Channel{
		Name: stream,
		ID:   m["room-id"],
	}

	uMeta := &UserMeta{
		isMod:   intToBool(m["mod"]),
		isSub:   intToBool(m["subscriber"]),
		isTurbo: intToBool(m["turbo"]),
		color:   m["color"],
		badges:  strings.Split(m["@badges"], ","),
	}

	se := &Sender{
		Name:     m["display-name"],
		ID:       m["user-id"],
		UserMeta: *uMeta,
	}

	tw := &NanPlugin{
		Platform:  "Twitch",
		Channel:   *ch,
		Timestamp: strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
		Sender:    *se,
		Message:   chatMsg,
		SplitMsg:  strings.Split(chatMsg, " "),
		Plugin:    nil,
	}

	// _ = tw
	// log.Println(chatmeta)
	log.Printf("Platform: %s, Channel: %s, Sender: %t", tw.Platform, tw.Channel.Name, tw.Sender.UserMeta.isMod)
	log.Printf("%#v", tw)
	// log.Println(tw.Platform)

	bytes, err := json.Marshal(tw)
	if err != nil {
		log.Println(err)
	}
	nc.Publish("chatserv", bytes)

}

func intToBool(bstr string) bool {
	if bstr == "1" {
		return true
	}
	return false
}

func tokenMesgTokenIzer(msg, token string) string {

	splitMsg := strings.Split(msg, ";")

	var m map[string]string

	m = make(map[string]string)

	for i := range splitMsg {
		key, value := splitAfterQuals(splitMsg[i])
		m[key] = value
	}
	return m[msg]
}

func splitAfterQuals(msg string) (string, string) {

	split := strings.Split(msg, "=")

	return split[0], split[1]
}
