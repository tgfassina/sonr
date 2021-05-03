package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	olc "github.com/google/open-location-code/go"
	"github.com/h2non/filetype"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ***************************** //
// ** ConnectionRequest Mgnmt ** //
// ***************************** //
func (req *ConnectionRequest) AttachGeoToRequest(geo *GeoIP) *ConnectionRequest {
	if req.Latitude != 0 && req.Longitude != 0 {
		return req
	}
	req.Latitude = geo.Latitude
	req.Longitude = geo.Longitude
	return req
}

func (g *GeoIP) GetLocation() *Location {
	return &Location{
		Name:        g.State,
		Continent:   g.Continent,
		CountryCode: g.Country,
		Latitude:    g.Latitude,
		Longitude:   g.Longitude,
		MinorOlc:    olc.Encode(float64(g.Latitude), float64(g.Longitude), 6),
		MajorOlc:    olc.Encode(float64(g.Latitude), float64(g.Longitude), 4),
	}
}

func (req *ConnectionRequest) GetLocation() *Location {
	return &Location{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		MinorOlc:  olc.Encode(float64(req.Latitude), float64(req.Longitude), 6),
		MajorOlc:  olc.Encode(float64(req.Latitude), float64(req.Longitude), 4),
	}
}

func (req *ConnectionRequest) IsDesktop() bool {
	return req.Device.Platform == Platform_MacOS || req.Device.Platform == Platform_Linux || req.Device.Platform == Platform_Windows
}

func (req *ConnectionRequest) IsMobile() bool {
	return req.Device.Platform == Platform_IOS || req.Device.Platform == Platform_Android
}

// **************************** //
// ** Remote Info Management ** //
// **************************** //

// ^ Get Remote Point Info ^
func GetRemoteInfo(list []string) RemoteInfo {
	return RemoteInfo{
		Display: fmt.Sprintf("%s %s %s", list[0], list[1], list[2]),
		Topic:   fmt.Sprintf("%s-%s-%s", list[0], list[1], list[2]),
		Count:   int32(len(list)),
		IsJoin:  false,
		Words:   list,
	}
}

// *********************************** //
// ** Outgoing File Info Management ** //
// *********************************** //
// ^ Method Returns File Mime at Path ^ //
func GetFileMime(m *Metadata) (*MIME, error) {
	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(m.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		return nil, err
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		return nil, err
	}

	// @ 1. Set Mime for Metadata
	m.Mime = &MIME{
		Type:    MIME_Type(MIME_Type_value[kind.MIME.Type]),
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 2. Create Mime Protobuf
	return m.Mime, nil
}

// ^ Method Returns File Payload at Path ^ //
func GetFilePayload(m *Metadata) Payload {
	// @ 3. Find Payload
	if m.Mime.Type == MIME_image || m.Mime.Type == MIME_video || m.Mime.Type == MIME_audio {
		return Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(m.Path)

		// Cross Check Extension
		if ext == ".pdf" {
			return Payload_PDF
		} else if ext == ".ppt" || ext == ".pptx" {
			return Payload_PRESENTATION
		} else if ext == ".xls" || ext == ".xlsm" || ext == ".xlsx" || ext == ".csv" {
			return Payload_SPREADSHEET
		} else if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".ttf" {
			return Payload_TEXT
		} else {
			return Payload_OTHER
		}
	}
}

// ^ Method Returns File Size at Path ^ //
func GetFileSize(m *Metadata) (int32, error) {
	// Open File at Path
	file, err := os.Open(m.Path)
	if err != nil {
		return 0, err
	}

	// Find Info
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Set Size
	m.Size = int32(info.Size())
	return m.Size, nil
}

// *********************************** //
// ** Incoming File Info Management ** //
// *********************************** //
type InFile struct {
	Payload       Payload
	Metadata      *Metadata
	ChunkBaseChan chan Chunk64
	ChunkBufChan  chan ChunkBuffer
}

// ********************** //
// ** Lobby Management ** //
// ********************** //
// ^ Returns as Lobby Buffer ^
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

// ^ Add/Update Peer in Lobby ^
func (l *Lobby) Add(peer *Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Remove Peer from Lobby ^
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Sync Between Remote Peers Lobby ^
func (l *Lobby) Sync(ref *Lobby, remotePeer *Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			if l.User.IsNotPeerIDString(id) {
				l.Add(peer)
			}
		}
	}
	l.Add(remotePeer)
}
