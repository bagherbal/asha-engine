// Package generation2empiricalimportswitch implements Gate 465:
// Quark-Sector Empirical Import Switch / CKM Data Firewall.
//
// Gate 464 deliberately stopped at a synthetic CKM-null residual. Gate 465 adds
// the explicit empirical airlock required before quark-sector mass or CKM data
// may enter any future residual comparator. The airlock is fail-closed: the
// switch must be on, source/scale/scheme/uncertainty metadata must be present,
// bridge-only quarantine must be declared, and every native-promotion attempt is
// rejected before the record can touch the native theorem registry.
package generation2empiricalimportswitch

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE465-QUARK-SECTOR-EMPIRICAL-IMPORT-SWITCH-CKM-DATA-FIREWALL"

	StatusGate464Inherited           = "CONDITIONAL_SUPPORT_GATE464_CKM_NULL_RESIDUAL_INHERITED"
	StatusAirlockDefined             = "CONDITIONAL_SUPPORT_EMPIRICAL_IMPORT_AIRLOCK_DEFINED"
	StatusQuarantinedImportAccepted  = "CONDITIONAL_SUPPORT_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_COMPARATOR_LEDGER"
	StatusEmpiricalImportSwitchValid = "CONDITIONAL_SUPPORT_EMPIRICAL_IMPORT_SWITCH_VALIDATED"
	StatusFirewallPreserved          = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_AIRLOCK_OPEN"

	StatusFailedSwitchDisabled          = "FAILED_ROUTE_EMPIRICAL_IMPORT_SWITCH_DISABLED"
	StatusFailedMissingMetadata         = "FAILED_ROUTE_EMPIRICAL_METADATA_INCOMPLETE"
	StatusFailedMissingUncertainty      = "FAILED_ROUTE_EMPIRICAL_UNCERTAINTY_MISSING"
	StatusFailedMissingBridgeOnly       = "FAILED_ROUTE_EMPIRICAL_DATA_MISSING_BRIDGE_ONLY_QUARANTINE"
	StatusFailedUnsupportedLedger       = "FAILED_ROUTE_EMPIRICAL_DATA_UNSUPPORTED_LEDGER_REJECTED"
	StatusFailedNativePromotion         = "FAILED_ROUTE_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite     = "FAILED_ROUTE_EMPIRICAL_DATA_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedCKMPMNSNativePrediction = "FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedObservedDataAsTheorem   = "FAILED_ROUTE_OBSERVED_DATA_AS_THEOREM_INPUT_REJECTED"
)

const (
	NativeFlavorDim       = 13
	KXYCoeffDim           = 9
	RequiredMetadataCount = 4
	ComparatorLedger      = "quark-sector-comparator-ledger"
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate457ProvenanceContract        bool
	Gate463EigenbasisLedger          bool
	Gate464CKMNullResidualAdapter    bool
	Gate464RejectsObservedByDefault  bool
	Gate464RejectsNativePrediction   bool
	Gate464RejectsMatrixExport       bool
	Gate464DiagnosticOnly            bool
	NativeRegistryCleanBeforeAirlock bool
	Verdict                          string
}

type AirlockPolicy struct {
	Executed                     bool
	StateVariableName            string
	DefaultEmpiricalImport       bool
	RequiresExplicitTrue         bool
	RequiresSource               bool
	RequiresScale                bool
	RequiresScheme               bool
	RequiresUncertainty          bool
	RequiresBridgeOnlyQuarantine bool
	AllowedLedger                string
	RejectsNativePredictionLabel bool
	RejectsNativeLawLabel        bool
	RejectsNativeRegistryWrite   bool
	RejectsObservedDataAsTheorem bool
	RequiredMetadataCount        int
	Verdict                      string
	Reason                       string
}

type EmpiricalRecord struct {
	Name            string
	Observable      string
	Sector          string
	Source          string
	Scale           string
	Scheme          string
	Uncertainty     string
	ValueKind       string
	Unit            string
	NumericRedacted bool
	BridgeOnly      bool
	TargetLedger    string
}

type ImportRequest struct {
	Name                         string
	EmpiricalImport              bool
	Record                       EmpiricalRecord
	NativePredictionClaim        bool
	NativeLawClaim               bool
	NativeRegistryWriteRequested bool
	CKMPMNSNativePredictionClaim bool
	ObservedDataAsTheoremInput   bool
}

type ImportResult struct {
	Imported                   bool
	Quarantined                bool
	ComparatorLedgerWritten    bool
	NativeRegistryWritten      bool
	NativePredictionLogged     bool
	NativeLawLogged            bool
	ObservedDataAsTheoremInput bool
	MetadataComplete           bool
	SwitchOpen                 bool
	BridgeOnly                 bool
	Verdict                    string
	Reason                     string
}

type Case struct {
	Name     string
	Request  ImportRequest
	Accepted bool
	Result   ImportResult
	Verdict  string
	Reason   string
}

type Sieve struct {
	Executed                           bool
	Cases                              []Case
	AcceptedCaseCount                  int
	RejectedCaseCount                  int
	ClosedSwitchRejected               bool
	MissingMetadataRejected            bool
	MissingUncertaintyRejected         bool
	MissingBridgeOnlyRejected          bool
	UnsupportedLedgerRejected          bool
	NativePromotionRejected            bool
	NativeRegistryWriteRejected        bool
	CKMPMNSNativePredictionRejected    bool
	ObservedDataAsTheoremRejected      bool
	QuarantinedQuarkMassImportAccepted bool
	QuarantinedCKMImportAccepted       bool
	AllAcceptedQuarantined             bool
	NoAcceptedNativeRegistryWrite      bool
	Verdict                            string
	Reason                             string
}

type Firewall struct {
	Executed                       bool
	AirlockCanOpen                 bool
	EmpiricalRowsImported          int
	AllImportedRowsQuarantined     bool
	EmpiricalDataInNativeRegistry  bool
	NativePredictionFromEmpirical  bool
	NativeLawFromEmpirical         bool
	ObservedDataUsedAsTheoremInput bool
	CKMMatrixNativePrediction      bool
	PMNSMatrixNativePrediction     bool
	CKMMatrixConstructed           bool
	CKMEntryComputed               bool
	QuarkMassNativePrediction      bool
	YukawaNativePrediction         bool
	KGenStillForced                bool
	XTriangleStillForced           bool
	YPhaseStillQuarantined         bool
	SectorCoefficientsStillSealed  bool
	NativeFlavorDimAfter           int
	KXYCoeffDimAfter               int
	Verdict                        string
	Reason                         string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Policy      AirlockPolicy
	Sieve       Sieve
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Policy = buildPolicy()
	a.Sieve = buildSieve()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate444KGenForced:                true,
		Gate445TriangleForced:            true,
		Gate457ProvenanceContract:        true,
		Gate463EigenbasisLedger:          true,
		Gate464CKMNullResidualAdapter:    true,
		Gate464RejectsObservedByDefault:  true,
		Gate464RejectsNativePrediction:   true,
		Gate464RejectsMatrixExport:       true,
		Gate464DiagnosticOnly:            true,
		NativeRegistryCleanBeforeAirlock: true,
		Verdict:                          StatusGate464Inherited,
	}
}

func buildPolicy() AirlockPolicy {
	return AirlockPolicy{
		Executed:                     true,
		StateVariableName:            "empirical_import",
		DefaultEmpiricalImport:       false,
		RequiresExplicitTrue:         true,
		RequiresSource:               true,
		RequiresScale:                true,
		RequiresScheme:               true,
		RequiresUncertainty:          true,
		RequiresBridgeOnlyQuarantine: true,
		AllowedLedger:                ComparatorLedger,
		RejectsNativePredictionLabel: true,
		RejectsNativeLawLabel:        true,
		RejectsNativeRegistryWrite:   true,
		RejectsObservedDataAsTheorem: true,
		RequiredMetadataCount:        RequiredMetadataCount,
		Verdict:                      StatusAirlockDefined,
		Reason:                       "external quark-sector data may pass only through empirical_import=true into a bridge-only comparator ledger with source, scale, scheme, and uncertainty metadata",
	}
}

func buildSieve() Sieve {
	cases := []Case{
		{Name: "valid quarantined quark-mass import", Request: validQuarkMassImport()},
		{Name: "valid quarantined CKM import", Request: validCKMImport()},
		{Name: "switch disabled", Request: switchDisabledImport()},
		{Name: "missing source metadata", Request: missingSourceImport()},
		{Name: "missing scale metadata", Request: missingScaleImport()},
		{Name: "missing scheme metadata", Request: missingSchemeImport()},
		{Name: "missing uncertainty metadata", Request: missingUncertaintyImport()},
		{Name: "missing bridge-only quarantine", Request: missingBridgeOnlyImport()},
		{Name: "unsupported target ledger", Request: unsupportedLedgerImport()},
		{Name: "native prediction promotion", Request: nativePredictionImport()},
		{Name: "native law promotion", Request: nativeLawImport()},
		{Name: "native registry write", Request: nativeRegistryWriteImport()},
		{Name: "CKM/PMNS native prediction", Request: ckmNativePredictionImport()},
		{Name: "observed data as theorem input", Request: observedDataAsTheoremImport()},
	}

	s := Sieve{Executed: true, Cases: make([]Case, 0, len(cases))}
	for _, c := range cases {
		res, accepted, verdict, reason := EvaluateImport(c.Request)
		c.Result = res
		c.Accepted = accepted
		c.Verdict = verdict
		c.Reason = reason
		if accepted {
			s.AcceptedCaseCount++
		} else {
			s.RejectedCaseCount++
		}
		s.ClosedSwitchRejected = s.ClosedSwitchRejected || verdict == StatusFailedSwitchDisabled
		s.MissingMetadataRejected = s.MissingMetadataRejected || verdict == StatusFailedMissingMetadata
		s.MissingUncertaintyRejected = s.MissingUncertaintyRejected || verdict == StatusFailedMissingUncertainty
		s.MissingBridgeOnlyRejected = s.MissingBridgeOnlyRejected || verdict == StatusFailedMissingBridgeOnly
		s.UnsupportedLedgerRejected = s.UnsupportedLedgerRejected || verdict == StatusFailedUnsupportedLedger
		s.NativePromotionRejected = s.NativePromotionRejected || verdict == StatusFailedNativePromotion
		s.NativeRegistryWriteRejected = s.NativeRegistryWriteRejected || verdict == StatusFailedNativeRegistryWrite
		s.CKMPMNSNativePredictionRejected = s.CKMPMNSNativePredictionRejected || verdict == StatusFailedCKMPMNSNativePrediction
		s.ObservedDataAsTheoremRejected = s.ObservedDataAsTheoremRejected || verdict == StatusFailedObservedDataAsTheorem
		s.QuarantinedQuarkMassImportAccepted = s.QuarantinedQuarkMassImportAccepted || (c.Name == "valid quarantined quark-mass import" && accepted && verdict == StatusQuarantinedImportAccepted)
		s.QuarantinedCKMImportAccepted = s.QuarantinedCKMImportAccepted || (c.Name == "valid quarantined CKM import" && accepted && verdict == StatusQuarantinedImportAccepted)
		s.Cases = append(s.Cases, c)
	}
	s.AllAcceptedQuarantined = true
	s.NoAcceptedNativeRegistryWrite = true
	for _, c := range s.Cases {
		if c.Accepted && (!c.Result.Quarantined || !c.Result.ComparatorLedgerWritten || c.Result.NativeRegistryWritten || c.Result.NativePredictionLogged || c.Result.NativeLawLogged) {
			s.AllAcceptedQuarantined = false
			s.NoAcceptedNativeRegistryWrite = false
		}
	}
	s.Verdict = StatusEmpiricalImportSwitchValid
	s.Reason = "the airlock accepts only explicitly switched-on, fully metadated, bridge-only rows into the comparator ledger and rejects every native registry or theorem-promotion route"
	return s
}

func baseRecord() EmpiricalRecord {
	return EmpiricalRecord{
		Name:            "redacted empirical row",
		Observable:      "redacted_quark_sector_comparator",
		Sector:          "u-d",
		Source:          "PDG-or-explicit-external-source",
		Scale:           "declared-renormalization-scale",
		Scheme:          "declared-renormalization-scheme",
		Uncertainty:     "declared-uncertainty",
		ValueKind:       "observed-redacted",
		Unit:            "declared-unit-or-dimensionless",
		NumericRedacted: true,
		BridgeOnly:      true,
		TargetLedger:    ComparatorLedger,
	}
}

func validQuarkMassImport() ImportRequest {
	r := baseRecord()
	r.Name = "redacted quark mass comparator row"
	r.Observable = "quark_mass_ratio_or_running_mass"
	r.Sector = "u/d quark sector"
	r.Unit = "declared mass unit"
	return ImportRequest{Name: "valid quarantined quark-mass import", EmpiricalImport: true, Record: r}
}

func validCKMImport() ImportRequest {
	r := baseRecord()
	r.Name = "redacted CKM comparator row"
	r.Observable = "CKM_angle_or_rephasing_invariant"
	r.Sector = "u-d mixing sector"
	r.Unit = "dimensionless"
	return ImportRequest{Name: "valid quarantined CKM import", EmpiricalImport: true, Record: r}
}

func switchDisabledImport() ImportRequest {
	x := validQuarkMassImport()
	x.EmpiricalImport = false
	return x
}

func missingSourceImport() ImportRequest {
	x := validQuarkMassImport()
	x.Record.Source = ""
	return x
}

func missingScaleImport() ImportRequest {
	x := validQuarkMassImport()
	x.Record.Scale = ""
	return x
}

func missingSchemeImport() ImportRequest {
	x := validQuarkMassImport()
	x.Record.Scheme = ""
	return x
}

func missingUncertaintyImport() ImportRequest {
	x := validQuarkMassImport()
	x.Record.Uncertainty = ""
	return x
}

func missingBridgeOnlyImport() ImportRequest {
	x := validQuarkMassImport()
	x.Record.BridgeOnly = false
	return x
}

func unsupportedLedgerImport() ImportRequest {
	x := validQuarkMassImport()
	x.Record.TargetLedger = "native-theorem-registry"
	return x
}

func nativePredictionImport() ImportRequest {
	x := validQuarkMassImport()
	x.NativePredictionClaim = true
	return x
}

func nativeLawImport() ImportRequest {
	x := validQuarkMassImport()
	x.NativeLawClaim = true
	return x
}

func nativeRegistryWriteImport() ImportRequest {
	x := validQuarkMassImport()
	x.NativeRegistryWriteRequested = true
	return x
}

func ckmNativePredictionImport() ImportRequest {
	x := validCKMImport()
	x.CKMPMNSNativePredictionClaim = true
	return x
}

func observedDataAsTheoremImport() ImportRequest {
	x := validCKMImport()
	x.ObservedDataAsTheoremInput = true
	return x
}

func EvaluateImport(req ImportRequest) (ImportResult, bool, string, string) {
	res := ImportResult{SwitchOpen: req.EmpiricalImport, BridgeOnly: req.Record.BridgeOnly}
	if !req.EmpiricalImport {
		res.Verdict = StatusFailedSwitchDisabled
		res.Reason = "empirical_import must be explicitly true before any external quark-sector data can enter the comparator ledger"
		return res, false, res.Verdict, res.Reason
	}
	if req.NativePredictionClaim || req.NativeLawClaim {
		res.Verdict = StatusFailedNativePromotion
		res.Reason = "empirical records cannot be logged as native_prediction or native_law"
		return res, false, res.Verdict, res.Reason
	}
	if req.NativeRegistryWriteRequested {
		res.Verdict = StatusFailedNativeRegistryWrite
		res.Reason = "the empirical airlock has no write path into the native theorem registry"
		return res, false, res.Verdict, res.Reason
	}
	if req.CKMPMNSNativePredictionClaim {
		res.Verdict = StatusFailedCKMPMNSNativePrediction
		res.Reason = "CKM/PMNS values may be bridge comparators only and cannot become native predictions"
		return res, false, res.Verdict, res.Reason
	}
	if req.ObservedDataAsTheoremInput {
		res.Verdict = StatusFailedObservedDataAsTheorem
		res.Reason = "observed data cannot be used as an input premise for native theorem verification"
		return res, false, res.Verdict, res.Reason
	}
	if req.Record.TargetLedger != ComparatorLedger {
		res.Verdict = StatusFailedUnsupportedLedger
		res.Reason = "empirical data may target only the quarantined quark-sector comparator ledger"
		return res, false, res.Verdict, res.Reason
	}
	if !req.Record.BridgeOnly {
		res.Verdict = StatusFailedMissingBridgeOnly
		res.Reason = "empirical records must carry bridge_only=true quarantine metadata"
		return res, false, res.Verdict, res.Reason
	}
	if strings.TrimSpace(req.Record.Uncertainty) == "" {
		res.Verdict = StatusFailedMissingUncertainty
		res.Reason = "empirical records must declare uncertainty before import"
		return res, false, res.Verdict, res.Reason
	}
	if !metadataComplete(req.Record) {
		res.Verdict = StatusFailedMissingMetadata
		res.Reason = "empirical records require source, scale, scheme, and uncertainty metadata"
		return res, false, res.Verdict, res.Reason
	}
	res = ImportResult{
		Imported:                   true,
		Quarantined:                true,
		ComparatorLedgerWritten:    true,
		NativeRegistryWritten:      false,
		NativePredictionLogged:     false,
		NativeLawLogged:            false,
		ObservedDataAsTheoremInput: false,
		MetadataComplete:           true,
		SwitchOpen:                 true,
		BridgeOnly:                 true,
		Verdict:                    StatusQuarantinedImportAccepted,
		Reason:                     "record imported into the quarantined comparator ledger only; native theorem registry remains untouched",
	}
	return res, true, res.Verdict, res.Reason
}

func metadataComplete(r EmpiricalRecord) bool {
	return strings.TrimSpace(r.Source) != "" && strings.TrimSpace(r.Scale) != "" && strings.TrimSpace(r.Scheme) != "" && strings.TrimSpace(r.Uncertainty) != ""
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                       true,
		AirlockCanOpen:                 a.Sieve.QuarantinedQuarkMassImportAccepted && a.Sieve.QuarantinedCKMImportAccepted,
		EmpiricalRowsImported:          a.Sieve.AcceptedCaseCount,
		AllImportedRowsQuarantined:     a.Sieve.AllAcceptedQuarantined,
		EmpiricalDataInNativeRegistry:  false,
		NativePredictionFromEmpirical:  false,
		NativeLawFromEmpirical:         false,
		ObservedDataUsedAsTheoremInput: false,
		CKMMatrixNativePrediction:      false,
		PMNSMatrixNativePrediction:     false,
		CKMMatrixConstructed:           false,
		CKMEntryComputed:               false,
		QuarkMassNativePrediction:      false,
		YukawaNativePrediction:         false,
		KGenStillForced:                a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:           a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:         true,
		SectorCoefficientsStillSealed:  true,
		NativeFlavorDimAfter:           NativeFlavorDim,
		KXYCoeffDimAfter:               KXYCoeffDim,
		Verdict:                        StatusFirewallPreserved,
		Reason:                         "even with empirical_import=true, external rows are quarantined comparator inputs only; native K/X structural laws stay fixed and the 13-moduli flavor firewall remains unchanged",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        466,
		Title:       "Quark-Sector Observed Comparator Adapter / Redacted PDG Slot",
		Reason:      "Gate465 defines the empirical airlock; the next safe step is an optional bridge-only adapter that can read metadated quark/CKM records and compute labelled residuals without writing native predictions.",
		PrimaryTask: "compose imported comparator rows with Gate464 residual diagnostics under explicit provenance, scale, scheme, uncertainty, branch, and bridge-only controls",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate464CKMNullResidualAdapter || !a.Inheritance.Gate464RejectsObservedByDefault || !a.Inheritance.Gate464RejectsNativePrediction || !a.Inheritance.Gate464DiagnosticOnly || !a.Inheritance.NativeRegistryCleanBeforeAirlock {
		return fmt.Errorf("Gate464 inheritance incomplete")
	}
	if !a.Policy.Executed || a.Policy.StateVariableName != "empirical_import" || a.Policy.DefaultEmpiricalImport || !a.Policy.RequiresExplicitTrue || !a.Policy.RequiresSource || !a.Policy.RequiresScale || !a.Policy.RequiresScheme || !a.Policy.RequiresUncertainty || !a.Policy.RequiresBridgeOnlyQuarantine || a.Policy.AllowedLedger != ComparatorLedger || !a.Policy.RejectsNativePredictionLabel || !a.Policy.RejectsNativeLawLabel || !a.Policy.RejectsNativeRegistryWrite || !a.Policy.RejectsObservedDataAsTheorem || a.Policy.RequiredMetadataCount != RequiredMetadataCount {
		return fmt.Errorf("empirical airlock policy incomplete")
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCaseCount != 2 || a.Sieve.RejectedCaseCount != 12 || !a.Sieve.QuarantinedQuarkMassImportAccepted || !a.Sieve.QuarantinedCKMImportAccepted || !a.Sieve.AllAcceptedQuarantined || !a.Sieve.NoAcceptedNativeRegistryWrite {
		return fmt.Errorf("empirical import sieve did not fail closed")
	}
	if !(a.Sieve.ClosedSwitchRejected && a.Sieve.MissingMetadataRejected && a.Sieve.MissingUncertaintyRejected && a.Sieve.MissingBridgeOnlyRejected && a.Sieve.UnsupportedLedgerRejected && a.Sieve.NativePromotionRejected && a.Sieve.NativeRegistryWriteRejected && a.Sieve.CKMPMNSNativePredictionRejected && a.Sieve.ObservedDataAsTheoremRejected) {
		return fmt.Errorf("not all unsafe empirical import routes were rejected")
	}
	if !a.Firewall.Executed || !a.Firewall.AirlockCanOpen || a.Firewall.EmpiricalRowsImported != 2 || !a.Firewall.AllImportedRowsQuarantined || a.Firewall.EmpiricalDataInNativeRegistry || a.Firewall.NativePredictionFromEmpirical || a.Firewall.NativeLawFromEmpirical || a.Firewall.ObservedDataUsedAsTheoremInput || a.Firewall.CKMMatrixNativePrediction || a.Firewall.PMNSMatrixNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.QuarkMassNativePrediction || a.Firewall.YukawaNativePrediction || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("13-moduli firewall violated by empirical import switch")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.QuarantinedQuarkMassImportAccepted && a.Sieve.QuarantinedCKMImportAccepted && a.Sieve.NativePromotionRejected && a.Firewall.AllImportedRowsQuarantined && !a.Firewall.EmpiricalDataInNativeRegistry {
		return "Gate 465 validates the quark-sector empirical airlock: empirical_import=true can admit fully metadated quark-mass or CKM comparator rows into a quarantined bridge ledger, while every native_prediction, native_law, theorem-input, and native-registry route fails closed."
	}
	return "Gate 465 failed to validate the empirical import airlock."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t provenance=%t eigenbasis=%t gate464=%t observed_default_reject=%t native_prediction_reject=%t matrix_export_reject=%t diagnostic_only=%t native_registry_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate457ProvenanceContract, x.Gate463EigenbasisLedger, x.Gate464CKMNullResidualAdapter, x.Gate464RejectsObservedByDefault, x.Gate464RejectsNativePrediction, x.Gate464RejectsMatrixExport, x.Gate464DiagnosticOnly, x.NativeRegistryCleanBeforeAirlock, x.Verdict)
}

func FormatPolicy(x AirlockPolicy) string {
	return fmt.Sprintf("executed=%t state=%s default=%t explicit_true=%t source=%t scale=%t scheme=%t uncertainty=%t bridge_quarantine=%t ledger=%s reject_native_prediction=%t reject_native_law=%t reject_native_registry=%t reject_theorem_input=%t metadata_count=%d verdict=%s reason=%s", x.Executed, x.StateVariableName, x.DefaultEmpiricalImport, x.RequiresExplicitTrue, x.RequiresSource, x.RequiresScale, x.RequiresScheme, x.RequiresUncertainty, x.RequiresBridgeOnlyQuarantine, x.AllowedLedger, x.RejectsNativePredictionLabel, x.RejectsNativeLawLabel, x.RejectsNativeRegistryWrite, x.RejectsObservedDataAsTheorem, x.RequiredMetadataCount, x.Verdict, x.Reason)
}

func FormatRecord(x EmpiricalRecord) string {
	return fmt.Sprintf("name=%s observable=%s sector=%s source=%s scale=%s scheme=%s uncertainty=%s kind=%s unit=%s redacted=%t bridge_only=%t ledger=%s", x.Name, x.Observable, x.Sector, empty(x.Source), empty(x.Scale), empty(x.Scheme), empty(x.Uncertainty), x.ValueKind, x.Unit, x.NumericRedacted, x.BridgeOnly, x.TargetLedger)
}

func FormatRequest(x ImportRequest) string {
	return fmt.Sprintf("empirical_import=%t record={%s} native_prediction=%t native_law=%t native_registry_write=%t ckm_pmns_native=%t observed_as_theorem=%t", x.EmpiricalImport, FormatRecord(x.Record), x.NativePredictionClaim, x.NativeLawClaim, x.NativeRegistryWriteRequested, x.CKMPMNSNativePredictionClaim, x.ObservedDataAsTheoremInput)
}

func FormatResult(x ImportResult) string {
	return fmt.Sprintf("imported=%t quarantined=%t comparator_write=%t native_registry_write=%t native_prediction=%t native_law=%t theorem_input=%t metadata=%t switch_open=%t bridge_only=%t verdict=%s reason=%s", x.Imported, x.Quarantined, x.ComparatorLedgerWritten, x.NativeRegistryWritten, x.NativePredictionLogged, x.NativeLawLogged, x.ObservedDataAsTheoremInput, x.MetadataComplete, x.SwitchOpen, x.BridgeOnly, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d closed_switch=%t missing_metadata=%t missing_uncertainty=%t missing_bridge=%t unsupported_ledger=%t native_promotion=%t native_registry=%t ckm_pmns_native=%t theorem_input=%t quark_mass=%t ckm=%t quarantined=%t no_native_write=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.ClosedSwitchRejected, x.MissingMetadataRejected, x.MissingUncertaintyRejected, x.MissingBridgeOnlyRejected, x.UnsupportedLedgerRejected, x.NativePromotionRejected, x.NativeRegistryWriteRejected, x.CKMPMNSNativePredictionRejected, x.ObservedDataAsTheoremRejected, x.QuarantinedQuarkMassImportAccepted, x.QuarantinedCKMImportAccepted, x.AllAcceptedQuarantined, x.NoAcceptedNativeRegistryWrite, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t airlock_open=%t imported_rows=%d all_quarantined=%t empirical_in_native=%t native_prediction=%t native_law=%t theorem_input=%t CKM_native=%t PMNS_native=%t CKM_constructed=%t CKM_entry=%t quark_mass_native=%t yukawa_native=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.AirlockCanOpen, x.EmpiricalRowsImported, x.AllImportedRowsQuarantined, x.EmpiricalDataInNativeRegistry, x.NativePredictionFromEmpirical, x.NativeLawFromEmpirical, x.ObservedDataUsedAsTheoremInput, x.CKMMatrixNativePrediction, x.PMNSMatrixNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.QuarkMassNativePrediction, x.YukawaNativePrediction, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{
		StatusGate464Inherited,
		StatusAirlockDefined,
		StatusQuarantinedImportAccepted,
		StatusEmpiricalImportSwitchValid,
		StatusFirewallPreserved,
		StatusFailedSwitchDisabled,
		StatusFailedMissingMetadata,
		StatusFailedMissingUncertainty,
		StatusFailedMissingBridgeOnly,
		StatusFailedUnsupportedLedger,
		StatusFailedNativePromotion,
		StatusFailedNativeRegistryWrite,
		StatusFailedCKMPMNSNativePrediction,
		StatusFailedObservedDataAsTheorem,
	}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 465 Registry Audit — Quark-Sector Empirical Import Switch / CKM Data Firewall\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusEmpiricalImportSwitchValid + "`\n\n")
	b.WriteString("Gate 465 opens a controlled empirical airlock. It does not evaluate PDG data, does not construct CKM entries, and does not write observed values into the native theorem registry. It proves only that fully metadated external rows can enter a quarantined comparator ledger when `empirical_import=true`.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Airlock policy\n\n")
	b.WriteString(FormatPolicy(a.Policy) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("empirical_import default = false\n")
	b.WriteString("empirical_import must be true before external rows are accepted\n")
	b.WriteString("required metadata = {source, scale, scheme, uncertainty}\n")
	b.WriteString("allowed target     = quark-sector-comparator-ledger\n")
	b.WriteString("forbidden targets  = native theorem registry, native_prediction, native_law\n")
	b.WriteString("```\n\n")

	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Case | Accepted | Verdict | Request | Result | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s | %s | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(FormatRequest(c.Request)), esc(FormatResult(c.Result)), esc(c.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Native firewall proof\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("The airlock can be open while the native registry remains sealed: imported rows are bridge comparator inputs only. `K_gen`, the Generation-2 structural zero, and the full `X_triangle` topology remain native structural laws; phase rays, quark masses, Yukawa amplitudes, CKM/PMNS values, GST/Fritzsch assumptions, and all 9 charged K/X/Y coefficients remain non-native.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func empty(s string) string {
	if strings.TrimSpace(s) == "" {
		return "∅"
	}
	return s
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
