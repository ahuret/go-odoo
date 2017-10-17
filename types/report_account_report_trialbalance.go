package types

import (
	"time"
)

type ReportAccountReportTrialbalance struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportAccountReportTrialbalanceNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportAccountReportTrialbalanceModel string = "report.account.report_trialbalance"

type ReportAccountReportTrialbalances []ReportAccountReportTrialbalance

type ReportAccountReportTrialbalancesNil []ReportAccountReportTrialbalanceNil

func (s *ReportAccountReportTrialbalance) NilableType_() interface{} {
	return &ReportAccountReportTrialbalanceNil{}
}

func (ns *ReportAccountReportTrialbalanceNil) Type_() interface{} {
	s := &ReportAccountReportTrialbalance{}
	return load(ns, s)
}

func (s *ReportAccountReportTrialbalances) NilableType_() interface{} {
	return &ReportAccountReportTrialbalancesNil{}
}

func (ns *ReportAccountReportTrialbalancesNil) Type_() interface{} {
	s := &ReportAccountReportTrialbalances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportTrialbalance))
	}
	return s
}
