package mailbox

import (
	"errors"
	"os"
	"time"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
)

// Textile API definitions
const (
	// Textile Client API URL
	TextileClientURL = "https://api.textile.io"

	// Textile Miner Index Target
	TextileMinerIdx = "api.minerindex.hub.textile.io:443"

	// Textile Mailbox Directory
	TextileMailboxDirName = "mailbox"
)

var (
	logger              = golog.Default.Child("protocols/mailbox")
	ErrMailboxDisabled  = errors.New("Mailbox not enabled, cannot perform request.")
	ErrMissingAPIKey    = errors.New("Missing Textile API Key in env")
	ErrMissingAPISecret = errors.New("Missing Textile API Secret in env")
	ErrFailedEntry      = errors.New("Failed to get Topmost entry from Queue")
	ErrFailedAuth       = errors.New("Failed to Authenticate message")
	ErrEmptyRequests    = errors.New("Empty Request list provided")
	ErrRequestNotFound  = errors.New("Request not found in list")
)

// fetchApiKeys fetches the Textile Api/Secrect keys from the environment
func fetchApiKeys() (string, string, error) {
	// Get API Key
	key, ok := os.LookupEnv("TEXTILE_HUB_KEY")
	if !ok {
		return "", "", ErrMissingAPIKey
	}

	// Get API Secret
	secret, ok := os.LookupEnv("TEXTILE_HUB_SECRET")
	if !ok {
		return "", "", ErrMissingAPISecret
	}
	return key, secret, nil
}

// getMailboxPath returns the mailbox path from the device
func (mb *MailboxProtocol) getMailboxPath() (string, error) {
	// Get Mailbox Path
	path, err := fs.ThirdParty.GenPath(TextileMailboxDirName)
	if err != nil {
		logger.Errorf("%s - Failed to Find Existing Mailbox at Path", err)
		return "", err
	}
	return path, nil
}

// loadMailbox loads an existing mailbox instance
func (mb *MailboxProtocol) loadMailbox() error {
	logger.Debug("Loading Mailbox...")

	// Get Mailbox Path
	path, err := mb.getMailboxPath()
	if err != nil {
		logger.Errorf("%s - Failed to Create New Mailbox at Path", err)
		return err
	}

	// // Return Existing Mailbox
	// mailbox, err := mb.mail.GetLocalMailbox(mb.ctx, path)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Load Existing Mailbox", err)
	// 	return err
	// }

	// // Set mailbox
	// mb.mailbox = mailbox
	logger.Debug("Existing Mailbox has been loaded.", golog.Fields{"path": path})
	return nil
}

// newMailbox creates a new mailbox instance
func (mb *MailboxProtocol) newMailbox() error {
	logger.Debug("Creating new Mailbox...")

	// Get Mailbox Path
	path, err := mb.getMailboxPath()
	if err != nil {
		logger.Errorf("%s - Failed to Create New Mailbox at Path", err)
		return err
	}

	// Fetch API Keys
	// key, secret, err := fetchApiKeys()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new mailbox", err)
	// 	return err
	// }

	// // Create a new mailbox with config
	// mailbox, err := mb.mail.NewMailbox(mb.ctx, local.Config{
	// 	Path: path,
	// 	// Identity:  privKey.ThreadIdentity(),
	// 	APIKey:    key,
	// 	APISecret: secret,
	// })

	// // Check if Err is for ErrMailboxExists
	// if err == local.ErrMailboxExists {
	// 	logger.Debug("Mailbox already exists no need to create a new one")
	// 	// Load Existing Mailbox
	// 	return mb.loadMailbox()
	// }

	// // Check for errors
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create mailbox", err)
	// 	return err
	// }

	// // Set mailbox
	// mb.mailbox = mailbox
	logger.Debug("New Mailbox has been created.", golog.Fields{"path": path})
	return nil
}

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *api.DecisionEvent {
	return &api.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *api.InviteEvent {
	return &api.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}
