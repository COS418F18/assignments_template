package chandy_lamport

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

const debug = false

// ====================================
//  Messages exchanged between servers
// ====================================

// An event that represents the sending of a message.
// This is expected to be queued in `link.events`.
type SendMessageEvent struct {
	src     string
	dest    string
	message interface{}
	// The message will be received by the server at or after this time step
	receiveTime int
}

// A message sent from one server to another for token passing.
// This is expected to be encapsulated within a `sendMessageEvent`.
type TokenMessage struct {
	numTokens int
}

func (m TokenMessage) String() string {
	return fmt.Sprintf("token(%v)", m.numTokens)
}

// A message sent from one server to another during the chandy-lamport algorithm.
// This is expected to be encapsulated within a `sendMessageEvent`.
type MarkerMessage struct {
	snapshotId int
}

func (m MarkerMessage) String() string {
	return fmt.Sprintf("marker(%v)", m.snapshotId)
}

// =======================
//  Events used by logger
// =======================

// A message that signifies receiving of a message on a particular server
// This is used only for debugging that is not sent between servers
type ReceivedMessageEvent struct {
	src     string
	dest    string
	message interface{}
}

func (m ReceivedMessageEvent) String() string {
	switch msg := m.message.(type) {
	case TokenMessage:
		return fmt.Sprintf("%v received %v tokens from %v", m.dest, msg.numTokens, m.src)
	case MarkerMessage:
		return fmt.Sprintf("%v received marker(%v) from %v", m.dest, msg.snapshotId, m.src)
	}
	return fmt.Sprintf("Unrecognized message: %v", m.message)
}

// A message that signifies sending of a message on a particular server
// This is used only for debugging that is not sent between servers
type SentMessageEvent struct {
	src     string
	dest    string
	message interface{}
}

func (m SentMessageEvent) String() string {
	switch msg := m.message.(type) {
	case TokenMessage:
		return fmt.Sprintf("%v sent %v tokens to %v", m.src, msg.numTokens, m.dest)
	case MarkerMessage:
		return fmt.Sprintf("%v sent marker(%v) to %v", m.src, msg.snapshotId, m.dest)
	}
	return fmt.Sprintf("Unrecognized message: %v", m.message)
}

// A message that signifies the beginning of the snapshot process on a particular server.
// This is used only for debugging that is not sent between servers.
type StartSnapshot struct {
	serverId string
	snapshotId int
}

func (m StartSnapshot) String() string {
	return fmt.Sprintf("%v startSnapshot(%v)", m.serverId, m.snapshotId)
}

// A message that signifies the end of the snapshot process on a particular server.
// This is used only for debugging that is not sent between servers.
type EndSnapshot struct {
	serverId string
	snapshotId int
}

func (m EndSnapshot) String() string {
	return fmt.Sprintf("%v endSnapshot(%v)", m.serverId, m.snapshotId)
}

// ================================================
//  Events injected to the system by the simulator
// ================================================

// An event parsed from the .event files that represent the passing of tokens
// from one server to another
type PassTokenEvent struct {
	src    string
	dest   string
	tokens int
}

// An event parsed from the .event files that represent the initiation of the
// chandy-lamport snapshot algorithm
type SnapshotEvent struct {
	serverId string
}

// A message recorded during the snapshot process
type SnapshotMessage struct {
	src     string
	dest    string
	message interface{}
}

// State recorded during the snapshot process
type SnapshotState struct {
	id       int
	tokens   map[string]int // key = server ID, value = num tokens
	messages []*SnapshotMessage
}

// =====================
//  Misc helper methods
// =====================

// If the error is not nil, terminate
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Return the keys of the given map in sorted order.
// Note: The argument passed in MUST be a map, otherwise an error will be thrown.
func getSortedKeys(m interface{}) []string {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		log.Fatal("Attempted to access sorted keys of a non-map: ", m)
	}
	keys := make([]string, 0)
	for _, k := range v.MapKeys() {
		keys = append(keys, k.String())
	}
	sort.Strings(keys)
	return keys
}
