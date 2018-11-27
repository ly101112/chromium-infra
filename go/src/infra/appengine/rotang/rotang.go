// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package rotang handles the scheduling of oncall rotations.
package rotang

import (
	"time"

	"context"
	"go.chromium.org/gae/service/mail"
	"go.chromium.org/luci/server/router"
	"golang.org/x/oauth2"
)

// Rota represents a named rotation and it's shift entries.
type Rota struct {
	// Name of the Rota, should match a Configuration.
	Name string
	// Entries contain the shift entries for this rota.
	Entries []*ShiftEntry
}

// Configuration represents a rota configuration.
type Configuration struct {
	Config  Config
	Members []ShiftMember
}

// Config contains the rota configuration.
type Config struct {
	Name             string
	Description      string
	Calendar         string
	TokenID          string
	Owners           []string // TODO(olakar): Change from email to groups.
	Email            Email
	ShiftsToSchedule int
	Shifts           ShiftConfig
	Expiration       int
	Enabled          bool
}

// ShiftConfig holds the Shift configuration.
type ShiftConfig struct {
	// StartTime represents the start-time of the first shift.
	// Only the Time of day is considered.
	// Defaults to 00:00 PDT.
	StartTime time.Time
	// Length sets the number of days a shift lasts.
	Length int
	// Skip defines a number of days with no oncalls.
	Skip int
	// Shifts represents the shifts over a 24hour period.
	Shifts []Shift
	// ShiftMembers specifides number of members per shift.
	ShiftMembers int
	// Generator used to schedule new shifts.
	Generator string
}

// Shift represents a shift in a 24h rotation.
type Shift struct {
	// Name of the shift - Eg. "MTV Shift"
	Name string
	// Duration is the duration of the shift.
	Duration time.Duration
}

// ShiftEntry represents one scheduled shift.
type ShiftEntry struct {
	// Name of the Shift this entry belongs to.
	Name string
	// OnCall are the members on-call for this shift.
	OnCall    []ShiftMember
	StartTime time.Time
	EndTime   time.Time
	// Comment is an optional comment where the rota algo
	// can add some extra information.
	Comment string
	// EvtID can be used to match an event to a calendar event.
	EvtID string
}

// Email contains the Subject and Body templates for email generated by the rota.
type Email struct {
	// Subject is a string used as a template run against the Member structure to generate the e-mail Subject
	//  text.
	Subject string
	// Body is a string used as a template run againste the Member structure to generate the e-mail Body.
	Body string
	// DaysBeforeWarn sets the number of days before an on-call shift the notification e-mail is sent.
	DaysBeforeNotify int
}

// ShiftMember holds the information needed for a member of a shift.
type ShiftMember struct {
	Email     string
	ShiftName string
}

// Member represents one member of a rotation.
type Member struct {
	Name        string `json:"full_name"`
	Email       string `json:"email_address"`
	TZ          time.Location
	OOO         []OOO
	Preferences []Preference
}

// OOO contains one Out-of-Office event.
type OOO struct {
	Start    time.Time
	Duration time.Duration
	Comment  string
}

// Preference is used for Members to signal shift preferences.
type Preference int

// Possible preferences for a Member.
const (
	NoPreferences Preference = iota
	NoWeekends
	NoMonday
	NoTuesday
	NoWednesday
	NoThursday
	NoFriday
	NoSaturday
	NoSunday
	NoOncall
)

// ConfigStorer defines the Store interface.
type ConfigStorer interface {
	// CreateRotaConfig stores a Configuration in backend storage.
	CreateRotaConfig(ctx context.Context, cfg *Configuration) error
	// UpdateRotaConfig updates the specified configuration.
	UpdateRotaConfig(ctx context.Context, cfg *Configuration) error
	// RotaConfig fetches the specified rota Configuration from backend storage.
	// Leaving `name` empty will return a slice of all stored rota configurations.
	RotaConfig(ctx context.Context, name string) ([]*Configuration, error)
	// DeleteRotaConfig removes the specified rota from backend storage.
	DeleteRotaConfig(ctx context.Context, name string) error
	// AddRotaMember add a member to the backend store.
	AddRotaMember(ctx context.Context, rota string, member *ShiftMember) error
	// DeleteRotaMember deletes the specified member from backend storage.
	DeleteRotaMember(ctx context.Context, rota, email string) error
	// MemberOf returns the rotations the specified email is a member of.
	MemberOf(ctx context.Context, email string) ([]string, error)
	// EnableRota enables jobs to consider rotation.
	EnableRota(ctx context.Context, rota string) error
	// DisableRota disables jobs for rotation.
	DisableRota(ctx context.Context, rota string) error
	// RotaEnabled returns the Enabled state of a rota.
	RotaEnabled(ctx context.Context, rota string) (bool, error)
}

// MemberStorer defines the member store interface.
type MemberStorer interface {
	Member(ctx context.Context, email string) (*Member, error)
	AllMembers(ctx context.Context) ([]Member, error)
	CreateMember(ctx context.Context, member *Member) error
	UpdateMember(cxt context.Context, in *Member) error
	DeleteMember(ctx context.Context, email string) error
}

// TokenStorer is used to store OAuth2 tokens.
type TokenStorer interface {
	CreateToken(ctx context.Context, email string, token *oauth2.Token) error
	Token(ctx context.Context, email string) (*oauth2.Token, error)
	DeleteToken(ctx context.Context, email string) error
}

// ShiftStorer is used to store Shift entries.
type ShiftStorer interface {
	AddShifts(ctx context.Context, rota string, entries []ShiftEntry) error
	AllShifts(cxt context.Context, rota string) ([]ShiftEntry, error)
	ShiftsFromTo(ctx context.Context, rota string, from, to time.Time) ([]ShiftEntry, error)
	Shift(cxt context.Context, rota string, start time.Time) (*ShiftEntry, error)
	DeleteAllShifts(ctx context.Context, rota string) error
	DeleteShift(ctx context.Context, rota string, start time.Time) error
	UpdateShift(ctx context.Context, rota string, shift *ShiftEntry) error
	Oncall(ctx context.Context, at time.Time, rota string) (*ShiftEntry, error)
}

// RotaGenerator is used to generate oncall rotations.
type RotaGenerator interface {
	Name() string
	Generate(sc *Configuration, start time.Time, previous []ShiftEntry, members []Member, shiftsToSchedule int) ([]ShiftEntry, error)
}

// MailSender is used to send E-mails.
type MailSender interface {
	Send(context.Context, *mail.Message) error
}

// Calenderer is used to handle the shared rota calendars.
type Calenderer interface {
	CreateEvent(ctx *router.Context, cfg *Configuration, shifts []ShiftEntry) ([]ShiftEntry, error)
	Event(ctx *router.Context, cfg *Configuration, shift *ShiftEntry) (*ShiftEntry, error)
	Events(ctx *router.Context, cfg *Configuration, from, to time.Time) ([]ShiftEntry, error)
	UpdateEvent(ctx *router.Context, cfg *Configuration, updated *ShiftEntry) (*ShiftEntry, error)
	DeleteEvent(ctx *router.Context, cfg *Configuration, shift *ShiftEntry) error
	TrooperOncall(ctx *router.Context, calid, match string, t time.Time) ([]string, error)
	TrooperShifts(ctx *router.Context, calid, match string, from, to time.Time) ([]ShiftEntry, error)
}

// Info contains information used to fill in Email Subject/Body and Calendar descriptions.
type Info struct {
	RotaName    string
	ShiftConfig ShiftConfig
	ShiftEntry  ShiftEntry
	Member      Member
	MemberURL   string
}
