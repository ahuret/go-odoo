package types

import (
	"time"
)

type MailShortcode struct {
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	Description   string    `xmlrpc:"description"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	ShortcodeType string    `xmlrpc:"shortcode_type"`
	Source        string    `xmlrpc:"source"`
	Substitution  string    `xmlrpc:"substitution"`
	UnicodeSource string    `xmlrpc:"unicode_source"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type MailShortcodeNil struct {
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	Description   interface{} `xmlrpc:"description"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	ShortcodeType interface{} `xmlrpc:"shortcode_type"`
	Source        interface{} `xmlrpc:"source"`
	Substitution  interface{} `xmlrpc:"substitution"`
	UnicodeSource interface{} `xmlrpc:"unicode_source"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var MailShortcodeModel string = "mail.shortcode"

type MailShortcodes []MailShortcode

type MailShortcodesNil []MailShortcodeNil

func (s *MailShortcode) NilableType_() interface{} {
	return &MailShortcodeNil{}
}

func (ns *MailShortcodeNil) Type_() interface{} {
	s := &MailShortcode{}
	return load(ns, s)
}

func (s *MailShortcodes) NilableType_() interface{} {
	return &MailShortcodesNil{}
}

func (ns *MailShortcodesNil) Type_() interface{} {
	s := &MailShortcodes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailShortcode))
	}
	return s
}
