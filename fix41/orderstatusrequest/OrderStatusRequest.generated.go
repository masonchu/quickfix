package orderstatusrequest

import (
	"github.com/shopspring/decimal"

	"github.com/masonchu/quickfix/enum"
	"github.com/masonchu/quickfix/field"
	"github.com/masonchu/quickfix/fix41"
	"github.com/masonchu/quickfix/tag"
	"github.com/quickfixgo/quickfix"
)

// OrderStatusRequest is the fix41 OrderStatusRequest type, MsgType = H.
type OrderStatusRequest struct {
	fix41.Header
	*quickfix.Body
	fix41.Trailer
	Message *quickfix.Message
}

// FromMessage creates a OrderStatusRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) OrderStatusRequest {
	return OrderStatusRequest{
		Header:  fix41.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix41.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m OrderStatusRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a OrderStatusRequest initialized with the required fields for OrderStatusRequest.
func New(clordid field.ClOrdIDField, symbol field.SymbolField, side field.SideField) (m OrderStatusRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix41.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("H"))
	m.Set(clordid)
	m.Set(symbol)
	m.Set(side)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg OrderStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "H", r
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m OrderStatusRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetIDSource sets IDSource, Tag 22.
func (m OrderStatusRequest) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m OrderStatusRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m OrderStatusRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSide sets Side, Tag 54.
func (m OrderStatusRequest) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m OrderStatusRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m OrderStatusRequest) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetExecBroker sets ExecBroker, Tag 76.
func (m OrderStatusRequest) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m OrderStatusRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m OrderStatusRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetClientID sets ClientID, Tag 109.
func (m OrderStatusRequest) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m OrderStatusRequest) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m OrderStatusRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetPutOrCall sets PutOrCall, Tag 201.
func (m OrderStatusRequest) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m OrderStatusRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetMaturityDay sets MaturityDay, Tag 205.
func (m OrderStatusRequest) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m OrderStatusRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m OrderStatusRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m OrderStatusRequest) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIDSource gets IDSource, Tag 22.
func (m OrderStatusRequest) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m OrderStatusRequest) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m OrderStatusRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m OrderStatusRequest) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m OrderStatusRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m OrderStatusRequest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExecBroker gets ExecBroker, Tag 76.
func (m OrderStatusRequest) GetExecBroker() (v string, err quickfix.MessageRejectError) {
	var f field.ExecBrokerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m OrderStatusRequest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m OrderStatusRequest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClientID gets ClientID, Tag 109.
func (m OrderStatusRequest) GetClientID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m OrderStatusRequest) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m OrderStatusRequest) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201.
func (m OrderStatusRequest) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m OrderStatusRequest) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDay gets MaturityDay, Tag 205.
func (m OrderStatusRequest) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m OrderStatusRequest) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m OrderStatusRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m OrderStatusRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasIDSource returns true if IDSource is present, Tag 22.
func (m OrderStatusRequest) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m OrderStatusRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m OrderStatusRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSide returns true if Side is present, Tag 54.
func (m OrderStatusRequest) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m OrderStatusRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m OrderStatusRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasExecBroker returns true if ExecBroker is present, Tag 76.
func (m OrderStatusRequest) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m OrderStatusRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m OrderStatusRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasClientID returns true if ClientID is present, Tag 109.
func (m OrderStatusRequest) HasClientID() bool {
	return m.Has(tag.ClientID)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m OrderStatusRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m OrderStatusRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201.
func (m OrderStatusRequest) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m OrderStatusRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasMaturityDay returns true if MaturityDay is present, Tag 205.
func (m OrderStatusRequest) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m OrderStatusRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m OrderStatusRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}
