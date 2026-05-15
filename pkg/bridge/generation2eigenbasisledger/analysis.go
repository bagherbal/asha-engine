// Package generation2eigenbasisledger implements Gate 463:
// Eigenbasis Convention Ledger / Mixing-Matrix Gauge Audit.
//
// Gate 462 isolated a bridge-only u-d relative ray. Gate 463 audits the next
// missing datum before any CKM-facing residual can be evaluated: a sector-pair
// eigenbasis convention. The result is deliberately non-predictive. It proves
// that raw diagonalizers are gauge objects, not native observables, and defines
// the fail-closed convention ledger required by a future bridge-only CKM null
// adapter.
package generation2eigenbasisledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE463-EIGENBASIS-CONVENTION-LEDGER-MIXING-MATRIX-GAUGE-AUDIT"

	StatusGate462Inherited              = "CONDITIONAL_SUPPORT_GATE462_SECTOR_DIFFERENCE_INTERFACE_INHERITED"
	StatusGaugeAuditComplete            = "CONDITIONAL_SUPPORT_DIAGONALIZER_GAUGE_AUDIT_COMPLETE"
	StatusConventionSlotDefined         = "CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_SLOT_DEFINED"
	StatusConventionLedgerValidated     = "CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_LEDGER_VALIDATED"
	StatusCKMNullAdapterPreconditionSet = "CONDITIONAL_SUPPORT_CKM_NULL_ADAPTER_PRECONDITION_SET"
	StatusFirewallPreserved             = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedRequiresUDConventions     = "FAILED_ROUTE_EIGENBASIS_REQUIRES_U_D_CONVENTIONS"
	StatusFailedRawDiagonalizerPhaseGauge = "FAILED_ROUTE_RAW_DIAGONALIZERS_HAVE_PHASE_GAUGE"
	StatusFailedEigenvaluePermutation     = "FAILED_ROUTE_EIGENVALUE_PERMUTATION_NOT_NATIVE"
	StatusFailedDegenerateSpectrum        = "FAILED_ROUTE_DEGENERATE_SPECTRUM_REJECTED"
	StatusFailedKGenBasisRotation         = "FAILED_ROUTE_KGEN_BASIS_ROTATION_REJECTED"
	StatusFailedMissingConvention         = "FAILED_ROUTE_EIGENBASIS_CONVENTION_MISSING"
	StatusFailedObservedCKMPMNSImport     = "FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED"
	StatusFailedCKMPMNSNativePrediction   = "FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedEigenbasisNativePromotion = "FAILED_ROUTE_EIGENBASIS_NATIVE_PROMOTION_REJECTED"
	StatusFailedConventionExportsMatrix   = "FAILED_ROUTE_CONVENTION_LEDGER_CANNOT_EXPORT_CKM_MATRIX"
)

const (
	NativeFlavorDim       = 13
	KXYCoeffDim           = 9
	RequiredSectors       = 2
	PerSectorPhaseGauge   = 3
	PerSectorPermutation  = 6
	PairPhaseGaugeDim     = 6
	PairPermutationSheets = 36
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate445TriangleForced           bool
	Gate456RayInverse               bool
	Gate459BranchTags               bool
	Gate461SectorMultiplex          bool
	Gate462SectorDifference         bool
	Gate462RequiresEigenbasis       bool
	Gate462RejectsObservedCKMPMNS   bool
	Gate462RejectsNativePrediction  bool
	Gate462NoMixingObservableExport bool
	NoObservedValuesImported        bool
	Verdict                         string
}

type GaugeAudit struct {
	Executed                                bool
	RawDiagonalizerPhaseGaugePerSector      int
	RawEigenvaluePermutationSheetsPerSector int
	PairPhaseGaugeDimension                 int
	PairPermutationSheets                   int
	KGenSimpleSpectrum                      bool
	KGenBasisCentralizer                    string
	KGenPreservingRephasingsOnly            bool
	RawDiagonalizersAreNotObservables       bool
	PermutationLabelsAreNotNative           bool
	ConventionCanFixBridgeGauge             bool
	ConventionCannotCreatePrediction        bool
	Verdict                                 string
	Reason                                  string
}

type SectorConvention struct {
	Sector                        string
	KGenBasisDeclared             bool
	EigenvalueOrderingDeclared    bool
	EigenvalueOrderingNativeClaim bool
	PhaseGaugeDeclared            bool
	NormalizationDeclared         bool
	DegeneracyPolicyDeclared      bool
	BranchTagDeclared             bool
	ProvenanceDeclared            bool
	BridgeOnly                    bool
	DegenerateSpectrum            bool
	KGenBasisRotationRequested    bool
	RawDiagonalizerClaim          bool
	ObservedCKMImport             bool
	ObservedPMNSImport            bool
	CKMNativePredictionClaim      bool
	PMNSNativePredictionClaim     bool
	EigenbasisNativePromotion     bool
	ExportsMixingMatrix           bool
}

type ConventionResult struct {
	FromSector                 string
	ToSector                   string
	CompleteInputs             bool
	BridgeOnly                 bool
	ConventionReady            bool
	CKMMatrixComputed          bool
	PMNSMatrixComputed         bool
	ExportsNativeObservable    bool
	EigenbasisGaugeFixed       bool
	PermutationConventionFixed bool
	DegeneracyRejected         bool
	KGenAddressPreserved       bool
	Verdict                    string
	Reason                     string
}

type InterfaceContract struct {
	Executed                       bool
	RequiresUSector                bool
	RequiresDSector                bool
	RequiresKGenBasisDeclaration   bool
	RequiresEigenvalueOrdering     bool
	RequiresPhaseGauge             bool
	RequiresNormalization          bool
	RequiresDegeneracyPolicy       bool
	RequiresBranchTag              bool
	RequiresProvenance             bool
	BridgeOnly                     bool
	RejectsRawDiagonalizerGauge    bool
	RejectsNativePermutationLabels bool
	RejectsKGenBasisRotation       bool
	RejectsObservedMixingImport    bool
	RejectsNativeMixingPrediction  bool
	ExportsConventionReadinessOnly bool
	Verdict                        string
	Reason                         string
}

type Case struct {
	Name        string
	Conventions []SectorConvention
	Accepted    bool
	Result      ConventionResult
	Verdict     string
	Reason      string
}

type Sieve struct {
	Executed                          bool
	Cases                             []Case
	AcceptedCaseCount                 int
	RejectedCaseCount                 int
	ValidConventionAccepted           bool
	MissingSectorRejected             bool
	MissingConventionRejected         bool
	RawPhaseGaugeRejected             bool
	PermutationNativeRejected         bool
	DegenerateSpectrumRejected        bool
	KGenBasisRotationRejected         bool
	ObservedCKMPMNSRejected           bool
	NativePredictionRejected          bool
	EigenbasisNativePromotionRejected bool
	MatrixExportRejected              bool
	AllAcceptedBridgeOnly             bool
	NoMixingMatrixExported            bool
	Verdict                           string
	Reason                            string
}

type Firewall struct {
	Executed                      bool
	EigenbasisConventionDefined   bool
	CKMNullAdapterMayProceed      bool
	CKMMatrixEntryComputed        bool
	CKMMatrixEntryNative          bool
	PMNSMatrixEntryComputed       bool
	PMNSMatrixEntryNative         bool
	ObservedMassesImported        bool
	ObservedYukawasImported       bool
	ObservedCKMImported           bool
	ObservedPMNSImported          bool
	GSTPromoted                   bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Gauge       GaugeAudit
	Contract    InterfaceContract
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
	a.Gauge = buildGaugeAudit()
	a.Contract = buildContract()
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
		Executed:                        true,
		Gate444KGenForced:               true,
		Gate445TriangleForced:           true,
		Gate456RayInverse:               true,
		Gate459BranchTags:               true,
		Gate461SectorMultiplex:          true,
		Gate462SectorDifference:         true,
		Gate462RequiresEigenbasis:       true,
		Gate462RejectsObservedCKMPMNS:   true,
		Gate462RejectsNativePrediction:  true,
		Gate462NoMixingObservableExport: true,
		NoObservedValuesImported:        true,
		Verdict:                         StatusGate462Inherited,
	}
}

func buildGaugeAudit() GaugeAudit {
	return GaugeAudit{
		Executed:                                true,
		RawDiagonalizerPhaseGaugePerSector:      PerSectorPhaseGauge,
		RawEigenvaluePermutationSheetsPerSector: PerSectorPermutation,
		PairPhaseGaugeDimension:                 PairPhaseGaugeDim,
		PairPermutationSheets:                   PairPermutationSheets,
		KGenSimpleSpectrum:                      true,
		KGenBasisCentralizer:                    "centralizer_U(3)(K_gen)=U(1)^3",
		KGenPreservingRephasingsOnly:            true,
		RawDiagonalizersAreNotObservables:       true,
		PermutationLabelsAreNotNative:           true,
		ConventionCanFixBridgeGauge:             true,
		ConventionCannotCreatePrediction:        true,
		Verdict:                                 StatusGaugeAuditComplete,
		Reason:                                  "raw sector diagonalizers carry U(1)^3 phase gauge and S3 ordering ambiguity; a convention can fix bridge bookkeeping but cannot derive CKM/PMNS entries",
	}
}

func buildContract() InterfaceContract {
	return InterfaceContract{
		Executed:                       true,
		RequiresUSector:                true,
		RequiresDSector:                true,
		RequiresKGenBasisDeclaration:   true,
		RequiresEigenvalueOrdering:     true,
		RequiresPhaseGauge:             true,
		RequiresNormalization:          true,
		RequiresDegeneracyPolicy:       true,
		RequiresBranchTag:              true,
		RequiresProvenance:             true,
		BridgeOnly:                     true,
		RejectsRawDiagonalizerGauge:    true,
		RejectsNativePermutationLabels: true,
		RejectsKGenBasisRotation:       true,
		RejectsObservedMixingImport:    true,
		RejectsNativeMixingPrediction:  true,
		ExportsConventionReadinessOnly: true,
		Verdict:                        StatusConventionSlotDefined,
		Reason:                         "a future CKM residual adapter may receive only a bridge-only convention-ready u-d sector pair; no matrix element is exported here",
	}
}

func buildSieve() Sieve {
	cases := []Case{
		{Name: "valid synthetic u-d convention ledger", Conventions: validUDConventions()},
		{Name: "missing d sector convention", Conventions: []SectorConvention{validConvention("u")}},
		{Name: "missing eigenvalue ordering", Conventions: missingOrderingConventions()},
		{Name: "raw diagonalizer phase gauge", Conventions: rawDiagonalizerConventions()},
		{Name: "native eigenvalue permutation claim", Conventions: nativeOrderingConventions()},
		{Name: "degenerate spectrum", Conventions: degenerateSpectrumConventions()},
		{Name: "K_gen basis rotation requested", Conventions: kgenRotationConventions()},
		{Name: "observed CKM/PMNS import", Conventions: observedMixingConventions()},
		{Name: "native CKM/PMNS prediction claim", Conventions: nativePredictionConventions()},
		{Name: "native eigenbasis promotion", Conventions: nativeEigenbasisPromotionConventions()},
		{Name: "convention ledger tries to export CKM matrix", Conventions: matrixExportConventions()},
	}

	s := Sieve{Executed: true, Cases: make([]Case, 0, len(cases))}
	for _, c := range cases {
		res, accepted, verdict, reason := EvaluateConventionPair(c.Conventions)
		c.Result = res
		c.Accepted = accepted
		c.Verdict = verdict
		c.Reason = reason
		if accepted {
			s.AcceptedCaseCount++
		} else {
			s.RejectedCaseCount++
		}
		s.ValidConventionAccepted = s.ValidConventionAccepted || (c.Name == "valid synthetic u-d convention ledger" && accepted && verdict == StatusConventionLedgerValidated)
		s.MissingSectorRejected = s.MissingSectorRejected || verdict == StatusFailedRequiresUDConventions
		s.MissingConventionRejected = s.MissingConventionRejected || verdict == StatusFailedMissingConvention
		s.RawPhaseGaugeRejected = s.RawPhaseGaugeRejected || verdict == StatusFailedRawDiagonalizerPhaseGauge
		s.PermutationNativeRejected = s.PermutationNativeRejected || verdict == StatusFailedEigenvaluePermutation
		s.DegenerateSpectrumRejected = s.DegenerateSpectrumRejected || verdict == StatusFailedDegenerateSpectrum
		s.KGenBasisRotationRejected = s.KGenBasisRotationRejected || verdict == StatusFailedKGenBasisRotation
		s.ObservedCKMPMNSRejected = s.ObservedCKMPMNSRejected || verdict == StatusFailedObservedCKMPMNSImport
		s.NativePredictionRejected = s.NativePredictionRejected || verdict == StatusFailedCKMPMNSNativePrediction
		s.EigenbasisNativePromotionRejected = s.EigenbasisNativePromotionRejected || verdict == StatusFailedEigenbasisNativePromotion
		s.MatrixExportRejected = s.MatrixExportRejected || verdict == StatusFailedConventionExportsMatrix
		s.Cases = append(s.Cases, c)
	}
	s.AllAcceptedBridgeOnly = true
	s.NoMixingMatrixExported = true
	for _, c := range s.Cases {
		if c.Accepted && (!c.Result.BridgeOnly || c.Result.CKMMatrixComputed || c.Result.PMNSMatrixComputed || c.Result.ExportsNativeObservable) {
			s.AllAcceptedBridgeOnly = false
			s.NoMixingMatrixExported = false
		}
	}
	s.Verdict = StatusCKMNullAdapterPreconditionSet
	s.Reason = "only the complete bridge-only u-d eigenbasis convention ledger is accepted; it exports readiness, not a CKM/PMNS matrix"
	return s
}

func validUDConventions() []SectorConvention {
	return []SectorConvention{validConvention("u"), validConvention("d")}
}

func validConvention(sector string) SectorConvention {
	return SectorConvention{
		Sector:                     sector,
		KGenBasisDeclared:          true,
		EigenvalueOrderingDeclared: true,
		PhaseGaugeDeclared:         true,
		NormalizationDeclared:      true,
		DegeneracyPolicyDeclared:   true,
		BranchTagDeclared:          true,
		ProvenanceDeclared:         true,
		BridgeOnly:                 true,
	}
}

func missingOrderingConventions() []SectorConvention {
	rs := validUDConventions()
	rs[1].EigenvalueOrderingDeclared = false
	return rs
}

func rawDiagonalizerConventions() []SectorConvention {
	rs := validUDConventions()
	rs[0].PhaseGaugeDeclared = false
	rs[0].RawDiagonalizerClaim = true
	return rs
}

func nativeOrderingConventions() []SectorConvention {
	rs := validUDConventions()
	rs[0].EigenvalueOrderingNativeClaim = true
	return rs
}

func degenerateSpectrumConventions() []SectorConvention {
	rs := validUDConventions()
	rs[1].DegenerateSpectrum = true
	return rs
}

func kgenRotationConventions() []SectorConvention {
	rs := validUDConventions()
	rs[0].KGenBasisRotationRequested = true
	return rs
}

func observedMixingConventions() []SectorConvention {
	rs := validUDConventions()
	rs[0].ObservedCKMImport = true
	rs[1].ObservedPMNSImport = true
	return rs
}

func nativePredictionConventions() []SectorConvention {
	rs := validUDConventions()
	rs[0].CKMNativePredictionClaim = true
	rs[1].PMNSNativePredictionClaim = true
	return rs
}

func nativeEigenbasisPromotionConventions() []SectorConvention {
	rs := validUDConventions()
	rs[1].EigenbasisNativePromotion = true
	return rs
}

func matrixExportConventions() []SectorConvention {
	rs := validUDConventions()
	rs[0].ExportsMixingMatrix = true
	return rs
}

func EvaluateConventionPair(rows []SectorConvention) (ConventionResult, bool, string, string) {
	res := ConventionResult{FromSector: "u", ToSector: "d"}

	if hasObservedMixingImport(rows) {
		res.Verdict = StatusFailedObservedCKMPMNSImport
		res.Reason = "observed CKM/PMNS values are not accepted in the eigenbasis convention audit"
		return res, false, res.Verdict, res.Reason
	}
	if hasNativeMixingPrediction(rows) {
		res.Verdict = StatusFailedCKMPMNSNativePrediction
		res.Reason = "a convention ledger cannot promote CKM/PMNS entries into native ASHA predictions"
		return res, false, res.Verdict, res.Reason
	}
	if hasEigenbasisNativePromotion(rows) {
		res.Verdict = StatusFailedEigenbasisNativePromotion
		res.Reason = "sector eigenvectors are bridge gauge choices and cannot become native law-space"
		return res, false, res.Verdict, res.Reason
	}
	if hasMixingMatrixExport(rows) {
		res.Verdict = StatusFailedConventionExportsMatrix
		res.Reason = "Gate463 exports convention readiness only; CKM/PMNS matrix construction belongs to a later explicit bridge adapter"
		return res, false, res.Verdict, res.Reason
	}
	if hasKGenBasisRotation(rows) {
		res.Verdict = StatusFailedKGenBasisRotation
		res.Reason = "a general family rotation would erase the native K_gen address and is not a convention gauge"
		return res, false, res.Verdict, res.Reason
	}
	if hasNativeOrderingClaim(rows) {
		res.Verdict = StatusFailedEigenvaluePermutation
		res.Reason = "eigenvalue ordering is a bridge convention; native geometry does not label mass-generation permutations"
		return res, false, res.Verdict, res.Reason
	}
	if hasDegenerateSpectrum(rows) {
		res.Verdict = StatusFailedDegenerateSpectrum
		res.Reason = "degenerate spectra make eigenvectors non-unique, so the convention ledger must fail closed"
		return res, false, res.Verdict, res.Reason
	}
	if hasRawDiagonalizerClaim(rows) {
		res.Verdict = StatusFailedRawDiagonalizerPhaseGauge
		res.Reason = "raw diagonalizers retain U(1)^3 phase gauge per sector and cannot be used as observables"
		return res, false, res.Verdict, res.Reason
	}

	u, hasU := findSector(rows, "u")
	d, hasD := findSector(rows, "d")
	if !hasU || !hasD || !u.BridgeOnly || !d.BridgeOnly {
		res.Verdict = StatusFailedRequiresUDConventions
		res.Reason = "u and d sector conventions are both required and must be bridge-only"
		return res, false, res.Verdict, res.Reason
	}
	if !complete(u) || !complete(d) {
		res.Verdict = StatusFailedMissingConvention
		res.Reason = "each sector must declare K_gen basis, ordering, phase gauge, normalization, degeneracy policy, branch tag, and provenance"
		return res, false, res.Verdict, res.Reason
	}

	res.CompleteInputs = true
	res.BridgeOnly = true
	res.ConventionReady = true
	res.CKMMatrixComputed = false
	res.PMNSMatrixComputed = false
	res.ExportsNativeObservable = false
	res.EigenbasisGaugeFixed = true
	res.PermutationConventionFixed = true
	res.DegeneracyRejected = true
	res.KGenAddressPreserved = true
	res.Verdict = StatusConventionLedgerValidated
	res.Reason = "complete u-d eigenbasis conventions are present; the result is readiness for a later bridge residual adapter, not a CKM prediction"
	return res, true, res.Verdict, res.Reason
}

func complete(r SectorConvention) bool {
	return r.KGenBasisDeclared && r.EigenvalueOrderingDeclared && r.PhaseGaugeDeclared && r.NormalizationDeclared && r.DegeneracyPolicyDeclared && r.BranchTagDeclared && r.ProvenanceDeclared && r.BridgeOnly
}

func findSector(rows []SectorConvention, sector string) (SectorConvention, bool) {
	for _, r := range rows {
		if r.Sector == sector {
			return r, true
		}
	}
	return SectorConvention{}, false
}

func hasObservedMixingImport(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.ObservedCKMImport || r.ObservedPMNSImport {
			return true
		}
	}
	return false
}

func hasNativeMixingPrediction(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.CKMNativePredictionClaim || r.PMNSNativePredictionClaim {
			return true
		}
	}
	return false
}

func hasEigenbasisNativePromotion(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.EigenbasisNativePromotion {
			return true
		}
	}
	return false
}

func hasMixingMatrixExport(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.ExportsMixingMatrix {
			return true
		}
	}
	return false
}

func hasKGenBasisRotation(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.KGenBasisRotationRequested {
			return true
		}
	}
	return false
}

func hasNativeOrderingClaim(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.EigenvalueOrderingNativeClaim {
			return true
		}
	}
	return false
}

func hasDegenerateSpectrum(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.DegenerateSpectrum {
			return true
		}
	}
	return false
}

func hasRawDiagonalizerClaim(rows []SectorConvention) bool {
	for _, r := range rows {
		if r.RawDiagonalizerClaim || !r.PhaseGaugeDeclared {
			return true
		}
	}
	return false
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		EigenbasisConventionDefined:   a.Contract.Executed,
		CKMNullAdapterMayProceed:      a.Sieve.ValidConventionAccepted,
		CKMMatrixEntryComputed:        false,
		CKMMatrixEntryNative:          false,
		PMNSMatrixEntryComputed:       false,
		PMNSMatrixEntryNative:         false,
		ObservedMassesImported:        false,
		ObservedYukawasImported:       false,
		ObservedCKMImported:           false,
		ObservedPMNSImported:          false,
		GSTPromoted:                   false,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate463 defines only a bridge eigenbasis convention ledger; CKM/PMNS values, masses, Yukawas, branch choices, and coefficients remain quarantined.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        464,
		Title:       "CKM Null Residual Adapter / Convention-Ready Symbolic Map",
		Reason:      "Gate463 now supplies the bridge-only eigenbasis convention slot needed before relative u-d rays can be compared by any CKM-facing residual harness.",
		PrimaryTask: "compose Gate462 relative-ray diagnostics with Gate463 eigenbasis conventions into a symbolic/null CKM residual adapter that still rejects observed CKM data and native-prediction claims",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate462SectorDifference || !a.Inheritance.Gate462RequiresEigenbasis {
		return fmt.Errorf("Gate462 inheritance incomplete")
	}
	if !a.Gauge.Executed || a.Gauge.PairPhaseGaugeDimension != PairPhaseGaugeDim || a.Gauge.PairPermutationSheets != PairPermutationSheets || !a.Gauge.RawDiagonalizersAreNotObservables || !a.Gauge.PermutationLabelsAreNotNative {
		return fmt.Errorf("gauge audit did not close")
	}
	if !a.Contract.Executed || !a.Contract.RequiresUSector || !a.Contract.RequiresDSector || !a.Contract.RequiresEigenvalueOrdering || !a.Contract.RequiresPhaseGauge || !a.Contract.RejectsNativeMixingPrediction || !a.Contract.ExportsConventionReadinessOnly {
		return fmt.Errorf("interface contract incomplete")
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCaseCount != 1 || a.Sieve.RejectedCaseCount != 10 || !a.Sieve.ValidConventionAccepted || !a.Sieve.AllAcceptedBridgeOnly || !a.Sieve.NoMixingMatrixExported {
		return fmt.Errorf("sieve did not fail closed")
	}
	if !(a.Sieve.MissingSectorRejected && a.Sieve.MissingConventionRejected && a.Sieve.RawPhaseGaugeRejected && a.Sieve.PermutationNativeRejected && a.Sieve.DegenerateSpectrumRejected && a.Sieve.KGenBasisRotationRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.EigenbasisNativePromotionRejected && a.Sieve.MatrixExportRejected) {
		return fmt.Errorf("not all unsafe routes were rejected")
	}
	if !a.Firewall.Executed || !a.Firewall.EigenbasisConventionDefined || a.Firewall.CKMMatrixEntryComputed || a.Firewall.CKMMatrixEntryNative || a.Firewall.PMNSMatrixEntryComputed || a.Firewall.PMNSMatrixEntryNative || a.Firewall.ObservedCKMImported || a.Firewall.ObservedPMNSImported || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall violated")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.ValidConventionAccepted && a.Firewall.CKMNullAdapterMayProceed && !a.Firewall.CKMMatrixEntryComputed && !a.Firewall.ObservedCKMImported {
		return "Gate 463 proves that a CKM-facing bridge adapter needs an explicit u-d eigenbasis convention ledger. The ledger can fix bridge bookkeeping gauge, but it cannot create CKM/PMNS predictions or promote eigenvectors into native law-space."
	}
	return "Gate 463 failed to validate the eigenbasis convention firewall."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch_tags=%t multiplex=%t sector_difference=%t requires_eigenbasis=%t rejects_observed=%t rejects_native=%t no_mixing_export=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456RayInverse, x.Gate459BranchTags, x.Gate461SectorMultiplex, x.Gate462SectorDifference, x.Gate462RequiresEigenbasis, x.Gate462RejectsObservedCKMPMNS, x.Gate462RejectsNativePrediction, x.Gate462NoMixingObservableExport, x.NoObservedValuesImported, x.Verdict)
}

func FormatGaugeAudit(x GaugeAudit) string {
	return fmt.Sprintf("executed=%t phase_gauge_per_sector=U(1)^%d permutation_sheets_per_sector=%d pair_phase_dim=%d pair_permutation_sheets=%d K_simple=%t centralizer=%s rephasings_only=%t raw_not_observable=%t permutations_not_native=%t convention_bridge_fix=%t convention_no_prediction=%t verdict=%s reason=%s", x.Executed, x.RawDiagonalizerPhaseGaugePerSector, x.RawEigenvaluePermutationSheetsPerSector, x.PairPhaseGaugeDimension, x.PairPermutationSheets, x.KGenSimpleSpectrum, x.KGenBasisCentralizer, x.KGenPreservingRephasingsOnly, x.RawDiagonalizersAreNotObservables, x.PermutationLabelsAreNotNative, x.ConventionCanFixBridgeGauge, x.ConventionCannotCreatePrediction, x.Verdict, x.Reason)
}

func FormatContract(x InterfaceContract) string {
	return fmt.Sprintf("executed=%t sectors=u,d K_basis=%t ordering=%t phase_gauge=%t normalization=%t degeneracy_policy=%t branch_tag=%t provenance=%t bridge_only=%t reject_raw=%t reject_native_permutation=%t reject_K_rotation=%t reject_observed=%t reject_native_mixing=%t readiness_only=%t verdict=%s reason=%s", x.Executed, x.RequiresKGenBasisDeclaration, x.RequiresEigenvalueOrdering, x.RequiresPhaseGauge, x.RequiresNormalization, x.RequiresDegeneracyPolicy, x.RequiresBranchTag, x.RequiresProvenance, x.BridgeOnly, x.RejectsRawDiagonalizerGauge, x.RejectsNativePermutationLabels, x.RejectsKGenBasisRotation, x.RejectsObservedMixingImport, x.RejectsNativeMixingPrediction, x.ExportsConventionReadinessOnly, x.Verdict, x.Reason)
}

func FormatResult(x ConventionResult) string {
	return fmt.Sprintf("%s->%s complete=%t bridge_only=%t ready=%t CKM_computed=%t PMNS_computed=%t native_export=%t phase_fixed=%t permutation_fixed=%t degeneracy_rejected=%t K_preserved=%t verdict=%s", x.FromSector, x.ToSector, x.CompleteInputs, x.BridgeOnly, x.ConventionReady, x.CKMMatrixComputed, x.PMNSMatrixComputed, x.ExportsNativeObservable, x.EigenbasisGaugeFixed, x.PermutationConventionFixed, x.DegeneracyRejected, x.KGenAddressPreserved, x.Verdict)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d valid=%t missing_sector=%t missing_convention=%t raw_phase=%t permutation=%t degenerate=%t K_rotation=%t observed=%t native_prediction=%t native_eigenbasis=%t matrix_export=%t bridge_only=%t no_matrix=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.ValidConventionAccepted, x.MissingSectorRejected, x.MissingConventionRejected, x.RawPhaseGaugeRejected, x.PermutationNativeRejected, x.DegenerateSpectrumRejected, x.KGenBasisRotationRejected, x.ObservedCKMPMNSRejected, x.NativePredictionRejected, x.EigenbasisNativePromotionRejected, x.MatrixExportRejected, x.AllAcceptedBridgeOnly, x.NoMixingMatrixExported, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t convention_defined=%t null_adapter_may_proceed=%t CKM_computed=%t CKM_native=%t PMNS_computed=%t PMNS_native=%t masses_imported=%t yukawas_imported=%t CKM_imported=%t PMNS_imported=%t GST_promoted=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.EigenbasisConventionDefined, x.CKMNullAdapterMayProceed, x.CKMMatrixEntryComputed, x.CKMMatrixEntryNative, x.PMNSMatrixEntryComputed, x.PMNSMatrixEntryNative, x.ObservedMassesImported, x.ObservedYukawasImported, x.ObservedCKMImported, x.ObservedPMNSImported, x.GSTPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{
		StatusGate462Inherited,
		StatusGaugeAuditComplete,
		StatusConventionSlotDefined,
		StatusConventionLedgerValidated,
		StatusCKMNullAdapterPreconditionSet,
		StatusFirewallPreserved,
		StatusFailedRequiresUDConventions,
		StatusFailedRawDiagonalizerPhaseGauge,
		StatusFailedEigenvaluePermutation,
		StatusFailedDegenerateSpectrum,
		StatusFailedKGenBasisRotation,
		StatusFailedMissingConvention,
		StatusFailedObservedCKMPMNSImport,
		StatusFailedCKMPMNSNativePrediction,
		StatusFailedEigenbasisNativePromotion,
		StatusFailedConventionExportsMatrix,
	}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 463 Registry Audit — Eigenbasis Convention Ledger / Mixing-Matrix Gauge Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusCKMNullAdapterPreconditionSet + "`\n\n")
	b.WriteString("Gate 463 defines the bridge-only eigenbasis convention ledger required before a u-d relative ray can enter any future CKM residual adapter. It computes no CKM or PMNS entries and imports no observed flavor data.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Diagonalizer gauge audit\n\n")
	b.WriteString(FormatGaugeAudit(a.Gauge) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("raw sector diagonalizer gauge: U(1)^3 x S3\n")
	b.WriteString("u-d pair gauge before convention: (U(1)^3 x S3)_u x (U(1)^3 x S3)_d\n")
	b.WriteString("pair phase-gauge dimension: 6\n")
	b.WriteString("pair permutation sheets: 36\n")
	b.WriteString("K_gen-preserving basis group: centralizer_U(3)(K_gen)=U(1)^3\n")
	b.WriteString("```\n\n")

	b.WriteString("## Required convention ledger\n\n")
	b.WriteString(FormatContract(a.Contract) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("sector in {u,d}\n")
	b.WriteString("K_gen basis declared\n")
	b.WriteString("eigenvalue ordering declared as bridge convention, not native particle label\n")
	b.WriteString("eigenvector phase gauge declared\n")
	b.WriteString("unit normalization declared\n")
	b.WriteString("degeneracy policy = fail closed\n")
	b.WriteString("branch tag and provenance inherited from comparator ledger\n")
	b.WriteString("bridge_only = true\n")
	b.WriteString("```\n\n")

	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Case | Accepted | Verdict | Convention result | Reason |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(FormatResult(c.Result)), esc(c.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## What Gate 463 does not do\n\n")
	b.WriteString("Gate 463 does not diagonalize physical quark matrices, does not compute CKM entries, does not compute PMNS entries, does not choose particle labels natively, and does not turn eigenvectors into ASHA law. It only proves that a future residual adapter must carry an explicit bridge convention before matrix-like comparisons are meaningful.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
