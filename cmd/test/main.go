package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/masonchu/quickfix"
	"github.com/masonchu/quickfix/field"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,    // SIGINT
		syscall.SIGTERM, // SIGTERM
	)

	if err := createOrder(c); err != nil {
		fmt.Printf("%+v", err)
	}
}

func createOrder(c chan os.Signal) error {

	cfg, err := os.Open("fix.cfg")
	if err != nil {
		return fmt.Errorf("Error opening %v, %v\n", "fix.cfg", err)
	}
	defer cfg.Close()

	stringData, readErr := io.ReadAll(cfg)
	if readErr != nil {
		return fmt.Errorf("Error reading cfg: %s,", readErr)
	}

	appSettings, err := quickfix.ParseSettings(bytes.NewReader(stringData))
	if err != nil {
		return fmt.Errorf("Error reading cfg: %s,", err)
	}
	appSettings.SessionSettings()
	appSettings.GlobalSettings().Set("InitiateLogon", "N")
	fmt.Printf("%+v\n", appSettings)
	app := FtxApplication{}
	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)

	if err != nil {
		return fmt.Errorf("Error creating file log factory: %s,", err)
	}

	initiator, err := quickfix.NewInitiator(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
	if err != nil {
		return fmt.Errorf("Unable to create Initiator: %s\n", err)
	}

	err = initiator.Start()
	if err != nil {
		return fmt.Errorf("Unable to start Initiator: %s\n", err)
	}

	if err != nil {
		return fmt.Errorf("Unable to send order: %s\n", err)
	}

	<-c
	initiator.Stop()
	return nil
}

func hmac256hex(payload string, secret string) string {
	sig := hmac.New(sha256.New, []byte(secret))
	sig.Write([]byte(payload))
	return hex.EncodeToString(sig.Sum(nil)[:])
}

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func queryHeader(h header) {
	h.Set(field.NewSenderCompID("iw5ZmF7MpLrh6bTxtPPf7HYkZH3KYz4HAsGiCiMR"))
	h.Set(field.NewTargetCompID("FTX"))
}

type FtxApplication struct {
}

//Notification of a session begin created.
func (f FtxApplication) OnCreate(sessionID quickfix.SessionID) {
	fmt.Printf("FtxApplication OnCreate %v\n", sessionID)
}

//Notification of a session successfully logging on.
func (f FtxApplication) OnLogon(sessionID quickfix.SessionID) {
	fmt.Printf("FtxApplication OnLogon %v\n", sessionID)
}

//Notification of a session logging off or disconnecting.
func (f FtxApplication) OnLogout(sessionID quickfix.SessionID) {
	fmt.Printf("FtxApplication OnLogout %v\n", sessionID)
}

//Notification of admin message being sent to target.
func (f FtxApplication) ToAdmin(message *quickfix.Message, sessionID quickfix.SessionID) {
	fmt.Printf("FtxApplication ToAdmin %v %v\n", sessionID, message)
}

//Notification of app message being sent to target.
func (f FtxApplication) ToApp(message *quickfix.Message, sessionID quickfix.SessionID) error {
	fmt.Printf("FtxApplication ToApp %v %v\n", sessionID, message)
	return nil
}

//Notification of admin message being received from target.
func (f FtxApplication) FromAdmin(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Printf("FtxApplication FromAdmin %v %v\n", sessionID, message)
	return nil
}

//Notification of app message being received from target.
func (f FtxApplication) FromApp(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Printf("FtxApplication FromApp %v %v\n", sessionID, message)
	return nil
}
