package lobby

import (
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleEvents() {
	// @ Create Topic Handler
	topicHandler, err := lob.topic.EventHandler()
	if err != nil {
		log.Println(err)
		return
	}

	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := topicHandler.NextPeerEvent(lob.ctx)
		if err != nil {
			topicHandler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			lob.Exchange(lobEvent.Peer)
		}

		if lobEvent.Type == pubsub.PeerLeave {
			lob.removePeer(lobEvent.Peer)
		}

		md.GetState().NeedsWait()
	}
}

// ^ 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.messages)
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom == lob.host.ID() {
			continue
		}

		// Construct message
		notif := md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, &notif)
		if err != nil {
			continue
		}

		// Send valid messages onto the Messages channel
		lob.messages <- &notif
		md.GetState().NeedsWait()
	}

}

// ^ 1a. processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Data)
			}

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
		md.GetState().NeedsWait()
	}
}
