// Package generation2leptonempiricalimportswitch implements Gate 477:
// Lepton-Sector Empirical Import Switch / PMNS Data Firewall.
//
// Gate 476 constructed only a synthetic PMNS-null residual socket. Gate 477
// adds the explicit observed-data airlock for charged-lepton, neutrino, and
// PMNS residual-target rows. It intentionally does not evaluate observed PMNS
// data, does not infer I_K, and does not construct U_PMNS. External lepton rows
// can enter only a quarantined bridge comparator ledger when empirical_import is
// explicitly true and the full Gate 475 metadata/policy envelope is present.
package generation2leptonempiricalimportswitch

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE477-LEPTON-SECTOR-EMPIRICAL-IMPORT-SWITCH-PMNS-DATA-FIREWALL"

	StatusGate476Inherited          = "CONDITIONAL_SUPPORT_GATE476_PMNS_NULL_SOCKET_INHERITED"
	StatusLeptonAirlockDefined      = "CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_AIRLOCK_DEFINED"
	StatusQuarantinedLeptonAccepted = "CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_LEDGER"
	StatusLeptonImportSwitchValid   = "CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_SWITCH_VALIDATED"
	StatusFirewallPreserved         = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_LEPTON_AIRLOCK_OPEN"

	StatusFailedSwitchDisabled             = "FAILED_ROUTE_LEPTON_EMPIRICAL_IMPORT_SWITCH_DISABLED"
	StatusFailedMissingMetadata            = "FAILED_ROUTE_LEPTON_EMPIRICAL_METADATA_INCOMPLETE"
	StatusFailedMissingUncertainty         = "FAILED_ROUTE_LEPTON_EMPIRICAL_UNCERTAINTY_MISSING"
	StatusFailedMissingBridgeOnly          = "FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_MISSING_BRIDGE_ONLY_QUARANTINE"
	StatusFailedUnsupportedLedger          = "FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_UNSUPPORTED_LEDGER_REJECTED"
	StatusFailedMissingLeptonPolicies      = "FAILED_ROUTE_LEPTON_EMPIRICAL_NEUTRINO_POLICIES_MISSING"
	StatusFailedPMNSAsRayInput             = "FAILED_ROUTE_PMNS_USED_AS_EMPIRICAL_LEPTON_RAY_INPUT_REJECTED"
	StatusFailedNativePromotion            = "FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite        = "FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedPMNSNativePrediction       = "FAILED_ROUTE_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedObservedDataAsTheoremInput = "FAILED_ROUTE_LEPTON_OBSERVED_DATA_AS_THEOREM_INPUT_REJECTED"
)

const (
	NativeFlavorDim       = 13
	KXYCoeffDim           = 9
	RequiredMetadataCount = 4
	ComparatorLedger      = "lepton-sector-comparator-ledger"
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate475LeptonPreflight           bool
	Gate476PMNSNullResidualAdapter   bool
	Gate476SyntheticOnly             bool
	Gate476RejectsObservedPMNS       bool
	Gate476RejectsPMNSAsRayInput     bool
	Gate476RejectsNativePrediction   bool
	Gate476RejectsMatrixExport       bool
	Gate476DiagnosticOnly            bool
	NativeRegistryCleanBeforeAirlock bool
	Verdict                          string
}

type AirlockPolicy struct {
	Executed                            bool
	StateVariableName                   string
	DefaultEmpiricalImport              bool
	RequiresExplicitTrue                bool
	RequiresSource                      bool
	RequiresScale                       bool
	RequiresScheme                      bool
	RequiresUncertainty                 bool
	RequiresBridgeOnlyQuarantine        bool
	RequiresNeutrinoOrderingPolicy      bool
	RequiresAbsoluteNeutrinoScalePolicy bool
	RequiresMajoranaDiracPhasePolicy    bool
	AllowsPMNSResidualTarget            bool
	AllowsPMNSAsRayInput                bool
	AllowedLedger                       string
	RejectsNativePredictionLabel        bool
	RejectsNativeLawLabel               bool
	RejectsNativeRegistryWrite          bool
	RejectsObservedDataAsTheorem        bool
	RequiredMetadataCount               int
	Verdict                             string
	Reason                              string
}

type EmpiricalRecord struct {
	Name                        string
	Observable                  string
	Sector                      string
	Source                      string
	SourceVersion               string
	Scale                       string
	Scheme                      string
	Uncertainty                 string
	ValueKind                   string
	Unit                        string
	NumericRedacted             bool
	BridgeOnly                  bool
	TargetLedger                string
	NeutrinoOrderingPolicy      string
	AbsoluteNeutrinoScalePolicy string
	MajoranaDiracPhasePolicy    string
	EigenbasisConvention        string
	PMNSResidualTarget          bool
	PMNSAsRayInput              bool
}

type ImportRequest struct {
	Name                         string
	EmpiricalImport              bool
	Record                       EmpiricalRecord
	NativePredictionClaim        bool
	NativeLawClaim               bool
	NativeRegistryWriteRequested bool
	PMNSNativePredictionClaim    bool
	ObservedDataAsTheoremInput   bool
}

type ImportResult struct {
	Imported                   bool
	Quarantined                bool
	ComparatorLedgerWritten    bool
	NativeRegistryWritten      bool
	NativePredictionLogged     bool
	NativeLawLogged            bool
	PMNSResidualTargetAccepted bool
	PMNSAsRayInput             bool
	ObservedDataAsTheoremInput bool
	MetadataComplete           bool
	LeptonPoliciesComplete     bool
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
	Executed                               bool
	Cases                                  []Case
	AcceptedCaseCount                      int
	RejectedCaseCount                      int
	ClosedSwitchRejected                   bool
	MissingMetadataRejected                bool
	MissingUncertaintyRejected             bool
	MissingBridgeOnlyRejected              bool
	UnsupportedLedgerRejected              bool
	MissingLeptonPoliciesRejected          bool
	PMNSAsRayInputRejected                 bool
	NativePromotionRejected                bool
	NativeRegistryWriteRejected            bool
	PMNSNativePredictionRejected           bool
	ObservedDataAsTheoremRejected          bool
	QuarantinedChargedLeptonImportAccepted bool
	QuarantinedNeutrinoImportAccepted      bool
	QuarantinedPMNSResidualTargetAccepted  bool
	AllAcceptedQuarantined                 bool
	NoAcceptedNativeRegistryWrite          bool
	NoAcceptedPMNSAsRayInput               bool
	Verdict                                string
	Reason                                 string
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
	PMNSMatrixNativePrediction     bool
	PMNSMatrixConstructed          bool
	PMNSEntryComputed              bool
	LeptonMassNativePrediction     bool
	NeutrinoMassNativePrediction   bool
	IKNativeSelectorFound          bool
	DENuComputedFromObserved       bool
	DENuNativePrediction           bool
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
		Gate475LeptonPreflight:           true,
		Gate476PMNSNullResidualAdapter:   true,
		Gate476SyntheticOnly:             true,
		Gate476RejectsObservedPMNS:       true,
		Gate476RejectsPMNSAsRayInput:     true,
		Gate476RejectsNativePrediction:   true,
		Gate476RejectsMatrixExport:       true,
		Gate476DiagnosticOnly:            true,
		NativeRegistryCleanBeforeAirlock: true,
		Verdict:                          StatusGate476Inherited,
	}
}

func buildPolicy() AirlockPolicy {
	return AirlockPolicy{
		Executed:                            true,
		StateVariableName:                   "empirical_import",
		DefaultEmpiricalImport:              false,
		RequiresExplicitTrue:                true,
		RequiresSource:                      true,
		RequiresScale:                       true,
		RequiresScheme:                      true,
		RequiresUncertainty:                 true,
		RequiresBridgeOnlyQuarantine:        true,
		RequiresNeutrinoOrderingPolicy:      true,
		RequiresAbsoluteNeutrinoScalePolicy: true,
		RequiresMajoranaDiracPhasePolicy:    true,
		AllowsPMNSResidualTarget:            true,
		AllowsPMNSAsRayInput:                false,
		AllowedLedger:                       ComparatorLedger,
		RejectsNativePredictionLabel:        true,
		RejectsNativeLawLabel:               true,
		RejectsNativeRegistryWrite:          true,
		RejectsObservedDataAsTheorem:        true,
		RequiredMetadataCount:               RequiredMetadataCount,
		Verdict:                             StatusLeptonAirlockDefined,
		Reason:                              "external lepton/PMNS rows may pass only through empirical_import=true into a bridge-only comparator ledger with source, scale, scheme, uncertainty, and neutrino-policy metadata",
	}
}

func buildSieve() Sieve {
	cases := []Case{
		{Name: "valid charged-lepton comparator import", Request: validChargedLeptonImport()},
		{Name: "valid neutrino comparator import", Request: validNeutrinoImport()},
		{Name: "valid PMNS residual-target import", Request: validPMNSResidualTargetImport()},
		{Name: "switch disabled", Request: switchDisabledImport()},
		{Name: "missing source metadata", Request: missingSourceImport()},
		{Name: "missing scale metadata", Request: missingScaleImport()},
		{Name: "missing scheme metadata", Request: missingSchemeImport()},
		{Name: "missing uncertainty metadata", Request: missingUncertaintyImport()},
		{Name: "missing bridge-only quarantine", Request: missingBridgeOnlyImport()},
		{Name: "unsupported target ledger", Request: unsupportedLedgerImport()},
		{Name: "missing neutrino ordering policy", Request: missingNeutrinoOrderingImport()},
		{Name: "PMNS as lepton ray input", Request: pmnsAsRayInputImport()},
		{Name: "native prediction promotion", Request: nativePredictionImport()},
		{Name: "native law promotion", Request: nativeLawImport()},
		{Name: "native registry write", Request: nativeRegistryWriteImport()},
		{Name: "PMNS native prediction", Request: pmnsNativePredictionImport()},
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
		s.MissingLeptonPoliciesRejected = s.MissingLeptonPoliciesRejected || verdict == StatusFailedMissingLeptonPolicies
		s.PMNSAsRayInputRejected = s.PMNSAsRayInputRejected || verdict == StatusFailedPMNSAsRayInput
		s.NativePromotionRejected = s.NativePromotionRejected || verdict == StatusFailedNativePromotion
		s.NativeRegistryWriteRejected = s.NativeRegistryWriteRejected || verdict == StatusFailedNativeRegistryWrite
		s.PMNSNativePredictionRejected = s.PMNSNativePredictionRejected || verdict == StatusFailedPMNSNativePrediction
		s.ObservedDataAsTheoremRejected = s.ObservedDataAsTheoremRejected || verdict == StatusFailedObservedDataAsTheoremInput
		s.QuarantinedChargedLeptonImportAccepted = s.QuarantinedChargedLeptonImportAccepted || (c.Name == "valid charged-lepton comparator import" && accepted)
		s.QuarantinedNeutrinoImportAccepted = s.QuarantinedNeutrinoImportAccepted || (c.Name == "valid neutrino comparator import" && accepted)
		s.QuarantinedPMNSResidualTargetAccepted = s.QuarantinedPMNSResidualTargetAccepted || (c.Name == "valid PMNS residual-target import" && accepted && res.PMNSResidualTargetAccepted)
		s.Cases = append(s.Cases, c)
	}
	s.AllAcceptedQuarantined = true
	s.NoAcceptedNativeRegistryWrite = true
	s.NoAcceptedPMNSAsRayInput = true
	for _, c := range s.Cases {
		if c.Accepted && (!c.Result.Quarantined || !c.Result.ComparatorLedgerWritten || c.Result.NativeRegistryWritten || c.Result.NativePredictionLogged || c.Result.NativeLawLogged) {
			s.AllAcceptedQuarantined = false
			s.NoAcceptedNativeRegistryWrite = false
		}
		if c.Accepted && c.Result.PMNSAsRayInput {
			s.NoAcceptedPMNSAsRayInput = false
		}
	}
	s.Verdict = StatusLeptonImportSwitchValid
	s.Reason = "the lepton airlock accepts only explicitly switched-on, fully metadated, bridge-only lepton/PMNS-residual rows into the comparator ledger and rejects every coordinate-smuggling or native-registry route"
	return s
}

func baseRecord() EmpiricalRecord {
	return EmpiricalRecord{
		Name:                        "redacted lepton empirical row",
		Observable:                  "redacted_lepton_sector_comparator",
		Sector:                      "e/nu",
		Source:                      "PDG-or-explicit-neutrino-global-fit-source",
		SourceVersion:               "declared-source-version",
		Scale:                       "declared-common-scale",
		Scheme:                      "declared-renormalization-or-fit-scheme",
		Uncertainty:                 "declared-uncertainty",
		ValueKind:                   "observed-redacted",
		Unit:                        "declared-unit-or-dimensionless",
		NumericRedacted:             true,
		BridgeOnly:                  true,
		TargetLedger:                ComparatorLedger,
		NeutrinoOrderingPolicy:      "declared-normal-or-inverted-ordering-policy",
		AbsoluteNeutrinoScalePolicy: "declared-absolute-scale-policy",
		MajoranaDiracPhasePolicy:    "declared-majorana-dirac-phase-policy",
		EigenbasisConvention:        "declared-lepton-eigenbasis-convention",
	}
}

func validChargedLeptonImport() ImportRequest {
	r := baseRecord()
	r.Name = "redacted charged-lepton comparator row"
	r.Observable = "charged_lepton_mass_or_rank_complete_comparator"
	r.Sector = "e charged-lepton sector"
	return ImportRequest{Name: "valid charged-lepton comparator import", EmpiricalImport: true, Record: r}
}

func validNeutrinoImport() ImportRequest {
	r := baseRecord()
	r.Name = "redacted neutrino comparator row"
	r.Observable = "neutrino_mass_or_rank_complete_comparator"
	r.Sector = "nu neutrino sector"
	return ImportRequest{Name: "valid neutrino comparator import", EmpiricalImport: true, Record: r}
}

func validPMNSResidualTargetImport() ImportRequest {
	r := baseRecord()
	r.Name = "redacted PMNS residual target row"
	r.Observable = "PMNS_residual_target_angle_or_invariant"
	r.Sector = "e-nu residual target"
	r.Unit = "dimensionless"
	r.PMNSResidualTarget = true
	return ImportRequest{Name: "valid PMNS residual-target import", EmpiricalImport: true, Record: r}
}

func switchDisabledImport() ImportRequest {
	x := validChargedLeptonImport()
	x.EmpiricalImport = false
	return x
}
func missingSourceImport() ImportRequest {
	x := validChargedLeptonImport()
	x.Record.Source = ""
	return x
}
func missingScaleImport() ImportRequest {
	x := validChargedLeptonImport()
	x.Record.Scale = ""
	return x
}
func missingSchemeImport() ImportRequest {
	x := validChargedLeptonImport()
	x.Record.Scheme = ""
	return x
}
func missingUncertaintyImport() ImportRequest {
	x := validChargedLeptonImport()
	x.Record.Uncertainty = ""
	return x
}
func missingBridgeOnlyImport() ImportRequest {
	x := validChargedLeptonImport()
	x.Record.BridgeOnly = false
	return x
}
func unsupportedLedgerImport() ImportRequest {
	x := validChargedLeptonImport()
	x.Record.TargetLedger = "native-theorem-registry"
	return x
}
func missingNeutrinoOrderingImport() ImportRequest {
	x := validNeutrinoImport()
	x.Record.NeutrinoOrderingPolicy = ""
	return x
}
func pmnsAsRayInputImport() ImportRequest {
	x := validPMNSResidualTargetImport()
	x.Record.PMNSAsRayInput = true
	return x
}
func nativePredictionImport() ImportRequest {
	x := validNeutrinoImport()
	x.NativePredictionClaim = true
	return x
}
func nativeLawImport() ImportRequest { x := validNeutrinoImport(); x.NativeLawClaim = true; return x }
func nativeRegistryWriteImport() ImportRequest {
	x := validNeutrinoImport()
	x.NativeRegistryWriteRequested = true
	return x
}
func pmnsNativePredictionImport() ImportRequest {
	x := validPMNSResidualTargetImport()
	x.PMNSNativePredictionClaim = true
	return x
}
func observedDataAsTheoremImport() ImportRequest {
	x := validPMNSResidualTargetImport()
	x.ObservedDataAsTheoremInput = true
	return x
}

func EvaluateImport(req ImportRequest) (ImportResult, bool, string, string) {
	res := ImportResult{SwitchOpen: req.EmpiricalImport, BridgeOnly: req.Record.BridgeOnly, PMNSAsRayInput: req.Record.PMNSAsRayInput}
	if !req.EmpiricalImport {
		res.Verdict = StatusFailedSwitchDisabled
		res.Reason = "empirical_import must be explicitly true before external lepton or PMNS data can enter the comparator ledger"
		return res, false, res.Verdict, res.Reason
	}
	if req.NativePredictionClaim || req.NativeLawClaim {
		res.Verdict = StatusFailedNativePromotion
		res.Reason = "lepton empirical records cannot be logged as native_prediction or native_law"
		return res, false, res.Verdict, res.Reason
	}
	if req.NativeRegistryWriteRequested {
		res.Verdict = StatusFailedNativeRegistryWrite
		res.Reason = "the lepton empirical airlock has no write path into the native theorem registry"
		return res, false, res.Verdict, res.Reason
	}
	if req.PMNSNativePredictionClaim {
		res.Verdict = StatusFailedPMNSNativePrediction
		res.Reason = "PMNS values may be bridge residual targets only and cannot become native predictions"
		return res, false, res.Verdict, res.Reason
	}
	if req.ObservedDataAsTheoremInput {
		res.Verdict = StatusFailedObservedDataAsTheoremInput
		res.Reason = "observed lepton/PMNS data cannot be used as an input premise for native theorem verification"
		return res, false, res.Verdict, res.Reason
	}
	if req.Record.PMNSAsRayInput {
		res.Verdict = StatusFailedPMNSAsRayInput
		res.Reason = "PMNS values may be residual targets, not alpha/phi/I_K ray-coordinate sources"
		return res, false, res.Verdict, res.Reason
	}
	if req.Record.TargetLedger != ComparatorLedger {
		res.Verdict = StatusFailedUnsupportedLedger
		res.Reason = "lepton empirical data may target only the quarantined lepton-sector comparator ledger"
		return res, false, res.Verdict, res.Reason
	}
	if !req.Record.BridgeOnly {
		res.Verdict = StatusFailedMissingBridgeOnly
		res.Reason = "lepton empirical records must carry bridge_only=true quarantine metadata"
		return res, false, res.Verdict, res.Reason
	}
	if strings.TrimSpace(req.Record.Uncertainty) == "" {
		res.Verdict = StatusFailedMissingUncertainty
		res.Reason = "lepton empirical records must declare uncertainty before import"
		return res, false, res.Verdict, res.Reason
	}
	if !metadataComplete(req.Record) {
		res.Verdict = StatusFailedMissingMetadata
		res.Reason = "lepton empirical records require source, scale, scheme, and uncertainty metadata"
		return res, false, res.Verdict, res.Reason
	}
	if !leptonPoliciesComplete(req.Record) {
		res.Verdict = StatusFailedMissingLeptonPolicies
		res.Reason = "lepton empirical records require neutrino ordering, absolute neutrino scale, Majorana/Dirac phase, and eigenbasis-convention policies"
		return res, false, res.Verdict, res.Reason
	}
	res = ImportResult{
		Imported:                   true,
		Quarantined:                true,
		ComparatorLedgerWritten:    true,
		NativeRegistryWritten:      false,
		NativePredictionLogged:     false,
		NativeLawLogged:            false,
		PMNSResidualTargetAccepted: req.Record.PMNSResidualTarget,
		PMNSAsRayInput:             false,
		ObservedDataAsTheoremInput: false,
		MetadataComplete:           true,
		LeptonPoliciesComplete:     true,
		SwitchOpen:                 true,
		BridgeOnly:                 true,
		Verdict:                    StatusQuarantinedLeptonAccepted,
		Reason:                     "lepton record imported into the quarantined comparator ledger only; native theorem registry remains untouched",
	}
	return res, true, res.Verdict, res.Reason
}

func metadataComplete(r EmpiricalRecord) bool {
	return strings.TrimSpace(r.Source) != "" && strings.TrimSpace(r.Scale) != "" && strings.TrimSpace(r.Scheme) != "" && strings.TrimSpace(r.Uncertainty) != ""
}

func leptonPoliciesComplete(r EmpiricalRecord) bool {
	return strings.TrimSpace(r.NeutrinoOrderingPolicy) != "" && strings.TrimSpace(r.AbsoluteNeutrinoScalePolicy) != "" && strings.TrimSpace(r.MajoranaDiracPhasePolicy) != "" && strings.TrimSpace(r.EigenbasisConvention) != ""
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                       true,
		AirlockCanOpen:                 a.Sieve.QuarantinedChargedLeptonImportAccepted && a.Sieve.QuarantinedNeutrinoImportAccepted && a.Sieve.QuarantinedPMNSResidualTargetAccepted,
		EmpiricalRowsImported:          a.Sieve.AcceptedCaseCount,
		AllImportedRowsQuarantined:     a.Sieve.AllAcceptedQuarantined,
		EmpiricalDataInNativeRegistry:  false,
		NativePredictionFromEmpirical:  false,
		NativeLawFromEmpirical:         false,
		ObservedDataUsedAsTheoremInput: false,
		PMNSMatrixNativePrediction:     false,
		PMNSMatrixConstructed:          false,
		PMNSEntryComputed:              false,
		LeptonMassNativePrediction:     false,
		NeutrinoMassNativePrediction:   false,
		IKNativeSelectorFound:          false,
		DENuComputedFromObserved:       false,
		DENuNativePrediction:           false,
		KGenStillForced:                a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:           a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:         true,
		SectorCoefficientsStillSealed:  true,
		NativeFlavorDimAfter:           NativeFlavorDim,
		KXYCoeffDimAfter:               KXYCoeffDim,
		Verdict:                        StatusFirewallPreserved,
		Reason:                         "even with empirical_import=true, lepton and PMNS residual rows are quarantined comparator inputs only; no observed lepton quantity can become native law or compute d_e_nu without a later rank-complete evaluator",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 478, Title: "Observed lepton rank-complete ledger / PMNS non-computation audit", Reason: "Gate477 opens the lepton empirical airlock but does not evaluate observed PMNS data or infer I_K.", PrimaryTask: "ingest an explicit observed lepton ledger and fail closed unless e and nu sectors supply rank-complete I_spec/I_K/branch tags, neutrino policies, and bridge-only provenance"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate476PMNSNullResidualAdapter || !a.Inheritance.Gate476SyntheticOnly || !a.Inheritance.Gate476RejectsObservedPMNS || !a.Inheritance.Gate476RejectsPMNSAsRayInput || !a.Inheritance.Gate476RejectsNativePrediction || !a.Inheritance.Gate476DiagnosticOnly || !a.Inheritance.NativeRegistryCleanBeforeAirlock {
		return fmt.Errorf("Gate476 inheritance incomplete")
	}
	if !a.Policy.Executed || a.Policy.StateVariableName != "empirical_import" || a.Policy.DefaultEmpiricalImport || !a.Policy.RequiresExplicitTrue || !a.Policy.RequiresSource || !a.Policy.RequiresScale || !a.Policy.RequiresScheme || !a.Policy.RequiresUncertainty || !a.Policy.RequiresBridgeOnlyQuarantine || !a.Policy.RequiresNeutrinoOrderingPolicy || !a.Policy.RequiresAbsoluteNeutrinoScalePolicy || !a.Policy.RequiresMajoranaDiracPhasePolicy || !a.Policy.AllowsPMNSResidualTarget || a.Policy.AllowsPMNSAsRayInput || a.Policy.AllowedLedger != ComparatorLedger || !a.Policy.RejectsNativePredictionLabel || !a.Policy.RejectsNativeLawLabel || !a.Policy.RejectsNativeRegistryWrite || !a.Policy.RejectsObservedDataAsTheorem || a.Policy.RequiredMetadataCount != RequiredMetadataCount {
		return fmt.Errorf("lepton empirical airlock policy incomplete")
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCaseCount != 3 || a.Sieve.RejectedCaseCount != 14 || !a.Sieve.QuarantinedChargedLeptonImportAccepted || !a.Sieve.QuarantinedNeutrinoImportAccepted || !a.Sieve.QuarantinedPMNSResidualTargetAccepted || !a.Sieve.AllAcceptedQuarantined || !a.Sieve.NoAcceptedNativeRegistryWrite || !a.Sieve.NoAcceptedPMNSAsRayInput {
		return fmt.Errorf("lepton empirical import sieve did not fail closed")
	}
	if !(a.Sieve.ClosedSwitchRejected && a.Sieve.MissingMetadataRejected && a.Sieve.MissingUncertaintyRejected && a.Sieve.MissingBridgeOnlyRejected && a.Sieve.UnsupportedLedgerRejected && a.Sieve.MissingLeptonPoliciesRejected && a.Sieve.PMNSAsRayInputRejected && a.Sieve.NativePromotionRejected && a.Sieve.NativeRegistryWriteRejected && a.Sieve.PMNSNativePredictionRejected && a.Sieve.ObservedDataAsTheoremRejected) {
		return fmt.Errorf("not all unsafe lepton empirical import routes were rejected")
	}
	if !a.Firewall.Executed || !a.Firewall.AirlockCanOpen || a.Firewall.EmpiricalRowsImported != 3 || !a.Firewall.AllImportedRowsQuarantined || a.Firewall.EmpiricalDataInNativeRegistry || a.Firewall.NativePredictionFromEmpirical || a.Firewall.NativeLawFromEmpirical || a.Firewall.ObservedDataUsedAsTheoremInput || a.Firewall.PMNSMatrixNativePrediction || a.Firewall.PMNSMatrixConstructed || a.Firewall.PMNSEntryComputed || a.Firewall.LeptonMassNativePrediction || a.Firewall.NeutrinoMassNativePrediction || a.Firewall.IKNativeSelectorFound || a.Firewall.DENuComputedFromObserved || a.Firewall.DENuNativePrediction || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("13-moduli firewall violated by lepton empirical import switch")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.QuarantinedChargedLeptonImportAccepted && a.Sieve.QuarantinedNeutrinoImportAccepted && a.Sieve.QuarantinedPMNSResidualTargetAccepted && a.Sieve.NativePromotionRejected && a.Firewall.AllImportedRowsQuarantined && !a.Firewall.EmpiricalDataInNativeRegistry {
		return "Gate 477 validates the lepton-sector empirical airlock: empirical_import=true can admit fully metadated charged-lepton, neutrino, and PMNS residual-target rows into a quarantined bridge ledger, while PMNS-as-ray, native_prediction, native_law, theorem-input, and native-registry routes fail closed."
	}
	return "Gate 477 failed to validate the lepton empirical import airlock."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t preflight=%t gate476=%t synthetic_only=%t observed_pmns_reject=%t pmns_ray_reject=%t native_prediction_reject=%t matrix_export_reject=%t diagnostic_only=%t native_registry_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate475LeptonPreflight, x.Gate476PMNSNullResidualAdapter, x.Gate476SyntheticOnly, x.Gate476RejectsObservedPMNS, x.Gate476RejectsPMNSAsRayInput, x.Gate476RejectsNativePrediction, x.Gate476RejectsMatrixExport, x.Gate476DiagnosticOnly, x.NativeRegistryCleanBeforeAirlock, x.Verdict)
}

func FormatPolicy(x AirlockPolicy) string {
	return fmt.Sprintf("executed=%t state=%s default=%t explicit_true=%t source=%t scale=%t scheme=%t uncertainty=%t bridge_quarantine=%t ordering_policy=%t absolute_scale_policy=%t majorana_dirac_policy=%t pmns_target=%t pmns_as_ray=%t ledger=%s reject_native_prediction=%t reject_native_law=%t reject_native_registry=%t reject_theorem_input=%t metadata_count=%d verdict=%s reason=%s", x.Executed, x.StateVariableName, x.DefaultEmpiricalImport, x.RequiresExplicitTrue, x.RequiresSource, x.RequiresScale, x.RequiresScheme, x.RequiresUncertainty, x.RequiresBridgeOnlyQuarantine, x.RequiresNeutrinoOrderingPolicy, x.RequiresAbsoluteNeutrinoScalePolicy, x.RequiresMajoranaDiracPhasePolicy, x.AllowsPMNSResidualTarget, x.AllowsPMNSAsRayInput, x.AllowedLedger, x.RejectsNativePredictionLabel, x.RejectsNativeLawLabel, x.RejectsNativeRegistryWrite, x.RejectsObservedDataAsTheorem, x.RequiredMetadataCount, x.Verdict, x.Reason)
}

func FormatRecord(x EmpiricalRecord) string {
	return fmt.Sprintf("name=%s observable=%s sector=%s source=%s version=%s scale=%s scheme=%s uncertainty=%s kind=%s unit=%s redacted=%t bridge_only=%t ledger=%s ordering=%s abs_scale=%s maj_dirac=%s eigenbasis=%s pmns_target=%t pmns_as_ray=%t", x.Name, x.Observable, x.Sector, empty(x.Source), empty(x.SourceVersion), empty(x.Scale), empty(x.Scheme), empty(x.Uncertainty), x.ValueKind, x.Unit, x.NumericRedacted, x.BridgeOnly, x.TargetLedger, empty(x.NeutrinoOrderingPolicy), empty(x.AbsoluteNeutrinoScalePolicy), empty(x.MajoranaDiracPhasePolicy), empty(x.EigenbasisConvention), x.PMNSResidualTarget, x.PMNSAsRayInput)
}

func FormatRequest(x ImportRequest) string {
	return fmt.Sprintf("empirical_import=%t record={%s} native_prediction=%t native_law=%t native_registry_write=%t pmns_native=%t observed_as_theorem=%t", x.EmpiricalImport, FormatRecord(x.Record), x.NativePredictionClaim, x.NativeLawClaim, x.NativeRegistryWriteRequested, x.PMNSNativePredictionClaim, x.ObservedDataAsTheoremInput)
}

func FormatResult(x ImportResult) string {
	return fmt.Sprintf("imported=%t quarantined=%t comparator_write=%t native_registry_write=%t native_prediction=%t native_law=%t pmns_target=%t pmns_as_ray=%t theorem_input=%t metadata=%t lepton_policies=%t switch_open=%t bridge_only=%t verdict=%s reason=%s", x.Imported, x.Quarantined, x.ComparatorLedgerWritten, x.NativeRegistryWritten, x.NativePredictionLogged, x.NativeLawLogged, x.PMNSResidualTargetAccepted, x.PMNSAsRayInput, x.ObservedDataAsTheoremInput, x.MetadataComplete, x.LeptonPoliciesComplete, x.SwitchOpen, x.BridgeOnly, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d closed_switch=%t missing_metadata=%t missing_uncertainty=%t missing_bridge=%t unsupported_ledger=%t missing_policies=%t pmns_as_ray=%t native_promotion=%t native_registry=%t pmns_native=%t theorem_input=%t charged_lepton=%t neutrino=%t pmns_target=%t quarantined=%t no_native_write=%t no_pmns_ray=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.ClosedSwitchRejected, x.MissingMetadataRejected, x.MissingUncertaintyRejected, x.MissingBridgeOnlyRejected, x.UnsupportedLedgerRejected, x.MissingLeptonPoliciesRejected, x.PMNSAsRayInputRejected, x.NativePromotionRejected, x.NativeRegistryWriteRejected, x.PMNSNativePredictionRejected, x.ObservedDataAsTheoremRejected, x.QuarantinedChargedLeptonImportAccepted, x.QuarantinedNeutrinoImportAccepted, x.QuarantinedPMNSResidualTargetAccepted, x.AllAcceptedQuarantined, x.NoAcceptedNativeRegistryWrite, x.NoAcceptedPMNSAsRayInput, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t airlock_open=%t imported_rows=%d all_quarantined=%t empirical_in_native=%t native_prediction=%t native_law=%t theorem_input=%t PMNS_native=%t PMNS_constructed=%t PMNS_entry=%t lepton_mass_native=%t neutrino_mass_native=%t IK_native=%t d_e_nu_observed=%t d_e_nu_native=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.AirlockCanOpen, x.EmpiricalRowsImported, x.AllImportedRowsQuarantined, x.EmpiricalDataInNativeRegistry, x.NativePredictionFromEmpirical, x.NativeLawFromEmpirical, x.ObservedDataUsedAsTheoremInput, x.PMNSMatrixNativePrediction, x.PMNSMatrixConstructed, x.PMNSEntryComputed, x.LeptonMassNativePrediction, x.NeutrinoMassNativePrediction, x.IKNativeSelectorFound, x.DENuComputedFromObserved, x.DENuNativePrediction, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate476Inherited, StatusLeptonAirlockDefined, StatusQuarantinedLeptonAccepted, StatusLeptonImportSwitchValid, StatusFirewallPreserved, StatusFailedSwitchDisabled, StatusFailedMissingMetadata, StatusFailedMissingUncertainty, StatusFailedMissingBridgeOnly, StatusFailedUnsupportedLedger, StatusFailedMissingLeptonPolicies, StatusFailedPMNSAsRayInput, StatusFailedNativePromotion, StatusFailedNativeRegistryWrite, StatusFailedPMNSNativePrediction, StatusFailedObservedDataAsTheoremInput}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 477 Registry Audit — Lepton-Sector Empirical Import Switch / PMNS Data Firewall\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusLeptonImportSwitchValid + "`\n\n")
	b.WriteString("Gate 477 opens the lepton observed-data airlock. It does not evaluate observed PMNS data, does not construct `U_PMNS`, does not infer `I_K`, and does not write observed values into the native theorem registry. It proves only that fully metadated lepton rows can enter a quarantined comparator ledger when `empirical_import=true`.\n\n")
	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Airlock policy\n\n")
	b.WriteString(FormatPolicy(a.Policy) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("empirical_import default = false\n")
	b.WriteString("empirical_import must be true before external lepton rows are accepted\n")
	b.WriteString("required metadata = {source, scale, scheme, uncertainty}\n")
	b.WriteString("required lepton policies = {neutrino_ordering, absolute_neutrino_scale, majorana_dirac_phase, eigenbasis_convention}\n")
	b.WriteString("allowed target = lepton-sector-comparator-ledger\n")
	b.WriteString("PMNS allowed role = residual target only\n")
	b.WriteString("forbidden roles = alpha/phi/I_K source, native theorem registry, native_prediction, native_law\n")
	b.WriteString("```\n\n")
	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| case | accepted | verdict | reason |\n|---|---:|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s |\n", c.Name, c.Accepted, c.Verdict, c.Reason))
	}
	b.WriteString("\n## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("## Status ledger\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func empty(s string) string {
	if strings.TrimSpace(s) == "" {
		return "<missing>"
	}
	return s
}
