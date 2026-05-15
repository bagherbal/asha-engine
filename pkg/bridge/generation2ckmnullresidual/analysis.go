// Package generation2ckmnullresidual implements Gate 464:
// CKM Null Residual Adapter / Convention-Ready Symbolic Map.
//
// Gates 462-463 produced the two ingredients required before any CKM-facing
// comparison is meaningful: a bridge-only u-d relative ray and a bridge-only
// eigenbasis convention ledger. Gate 464 composes those ingredients into a
// symbolic/null residual harness. It deliberately does not construct V_CKM,
// does not export any matrix entry, and does not import observed CKM/PMNS data.
// The only accepted output is a convention-fixed comparator diagnostic on
// synthetic bridge rays.
package generation2ckmnullresidual

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE464-CKM-NULL-RESIDUAL-ADAPTER-CONVENTION-READY-SYMBOLIC-MAP"

	StatusGate463Inherited                 = "CONDITIONAL_SUPPORT_GATE463_EIGENBASIS_LEDGER_INHERITED"
	StatusCKMNullMapDefined                = "CONDITIONAL_SUPPORT_CKM_NULL_RESIDUAL_SYMBOLIC_MAP_DEFINED"
	StatusCKMNullResidualComputed          = "CONDITIONAL_SUPPORT_CKM_NULL_RESIDUAL_BRIDGE_ONLY_COMPUTED"
	StatusCKMNullResidualFirewallValidated = "CONDITIONAL_SUPPORT_CKM_NULL_RESIDUAL_FIREWALL_VALIDATED"
	StatusFirewallPreserved                = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedRequiresRelativeRay          = "FAILED_ROUTE_CKM_NULL_RESIDUAL_REQUIRES_RELATIVE_RAY"
	StatusFailedRequiresEigenbasisConvention = "FAILED_ROUTE_CKM_NULL_RESIDUAL_REQUIRES_EIGENBASIS_CONVENTION"
	StatusFailedRequiresBranchAndProvenance  = "FAILED_ROUTE_CKM_NULL_RESIDUAL_REQUIRES_BRANCH_AND_PROVENANCE"
	StatusFailedObservedCKMPMNSImport        = "FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED"
	StatusFailedCKMPMNSNativePrediction      = "FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedCKMMatrixExport              = "FAILED_ROUTE_CKM_MATRIX_EXPORT_REJECTED"
	StatusFailedRawDiagonalizer              = "FAILED_ROUTE_RAW_DIAGONALIZER_REJECTED"
	StatusFailedDegenerateSpectrum           = "FAILED_ROUTE_DEGENERATE_SPECTRUM_REJECTED"
	StatusFailedKGenBasisRotation            = "FAILED_ROUTE_KGEN_BASIS_ROTATION_REJECTED"
	StatusFailedNativeResidualPromotion      = "FAILED_ROUTE_CKM_RESIDUAL_NATIVE_PROMOTION_REJECTED"
	StatusFailedGSTAsCKMSelector             = "FAILED_ROUTE_GST_FRITZSCH_AS_CKM_SELECTOR_REJECTED"
)

const (
	NativeFlavorDim  = 13
	KXYCoeffDim      = 9
	RelativeRayDOF   = 2
	ConventionFields = 7
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate445TriangleForced           bool
	Gate456RayInverse               bool
	Gate459BranchTags               bool
	Gate461SectorMultiplex          bool
	Gate462RelativeRay              bool
	Gate462RejectsObservedCKMPMNS   bool
	Gate462RejectsNativePrediction  bool
	Gate463EigenbasisLedger         bool
	Gate463RejectsRawGauge          bool
	Gate463RejectsPermutationNative bool
	Gate463ReadyForResidualAdapter  bool
	NoObservedValuesImported        bool
	Verdict                         string
}

type CKMNullInput struct {
	Name                         string
	AlphaU                       float64
	PhiU                         float64
	AlphaD                       float64
	PhiD                         float64
	HasRelativeRay               bool
	HasEigenbasisConvention      bool
	HasOrderingConvention        bool
	HasPhaseGaugeConvention      bool
	HasNormalizationConvention   bool
	HasDegeneracyPolicy          bool
	HasBranchTags                bool
	HasProvenance                bool
	BridgeOnly                   bool
	SyntheticOnly                bool
	ObservedCKMImport            bool
	ObservedPMNSImport           bool
	CKMNativePredictionClaim     bool
	PMNSNativePredictionClaim    bool
	CKMMatrixExportRequested     bool
	RawDiagonalizerClaim         bool
	DegenerateSpectrum           bool
	KGenBasisRotationRequested   bool
	ResidualNativePromotionClaim bool
	GSTFritzschSelectorClaim     bool
}

type CKMNullResidual struct {
	DeltaAlpha                float64
	DeltaPhi                  float64
	PhaseChord                float64
	ProjectiveRayDistance     float64
	NullAlignmentResidual     float64
	CompleteInputs            bool
	ConventionFixed           bool
	BridgeOnly                bool
	SyntheticOnly             bool
	ExportsRelativeDiagnostic bool
	CKMMatrixConstructed      bool
	CKMEntryComputed          bool
	PMNSEntryComputed         bool
	ExportsNativeObservable   bool
	ObservedDataImported      bool
	Verdict                   string
	Reason                    string
}

type SymbolicMap struct {
	Executed                        bool
	RequiresRelativeRay             bool
	RequiresEigenbasisConvention    bool
	RequiresOrderingConvention      bool
	RequiresPhaseGaugeConvention    bool
	RequiresBranchTags              bool
	RequiresProvenance              bool
	RequiresBridgeOnlySyntheticMode bool
	RelativeRayDimension            int
	ConventionFieldCount            int
	ComputesDeltaAlpha              bool
	ComputesDeltaPhi                bool
	ComputesProjectiveDistance      bool
	ExportsResidualDiagnosticsOnly  bool
	CKMMatrixElementExported        bool
	CKMMatrixConstructed            bool
	Verdict                         string
	Reason                          string
}

type Case struct {
	Name     string
	Input    CKMNullInput
	Accepted bool
	Residual CKMNullResidual
	Verdict  string
	Reason   string
}

type Sieve struct {
	Executed                        bool
	Cases                           []Case
	AcceptedCaseCount               int
	RejectedCaseCount               int
	ValidSyntheticResidualAccepted  bool
	MissingRelativeRayRejected      bool
	MissingEigenbasisRejected       bool
	MissingBranchProvenanceRejected bool
	ObservedCKMPMNSRejected         bool
	NativePredictionRejected        bool
	MatrixExportRejected            bool
	RawDiagonalizerRejected         bool
	DegenerateSpectrumRejected      bool
	KGenBasisRotationRejected       bool
	NativeResidualPromotionRejected bool
	GSTSelectorRejected             bool
	AllAcceptedBridgeOnly           bool
	NoCKMMatrixConstructed          bool
	NoNativeObservableExport        bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                        bool
	CKMNullResidualAdapterDefined   bool
	CKMNullResidualMayRunBridgeOnly bool
	CKMMatrixConstructed            bool
	CKMMatrixEntryComputed          bool
	CKMMatrixEntryNative            bool
	PMNSMatrixEntryComputed         bool
	PMNSMatrixEntryNative           bool
	ObservedMassesImported          bool
	ObservedYukawasImported         bool
	ObservedCKMImported             bool
	ObservedPMNSImported            bool
	GSTFritzschPromoted             bool
	RelativeRayPromotedNative       bool
	EigenbasisPromotedNative        bool
	KGenStillForced                 bool
	XTriangleStillForced            bool
	YPhaseStillQuarantined          bool
	SectorCoefficientsStillSealed   bool
	NativeFlavorDimAfter            int
	KXYCoeffDimAfter                int
	Verdict                         string
	Reason                          string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Map         SymbolicMap
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
	a.Map = buildSymbolicMap()
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
		Gate462RelativeRay:              true,
		Gate462RejectsObservedCKMPMNS:   true,
		Gate462RejectsNativePrediction:  true,
		Gate463EigenbasisLedger:         true,
		Gate463RejectsRawGauge:          true,
		Gate463RejectsPermutationNative: true,
		Gate463ReadyForResidualAdapter:  true,
		NoObservedValuesImported:        true,
		Verdict:                         StatusGate463Inherited,
	}
}

func buildSymbolicMap() SymbolicMap {
	return SymbolicMap{
		Executed:                        true,
		RequiresRelativeRay:             true,
		RequiresEigenbasisConvention:    true,
		RequiresOrderingConvention:      true,
		RequiresPhaseGaugeConvention:    true,
		RequiresBranchTags:              true,
		RequiresProvenance:              true,
		RequiresBridgeOnlySyntheticMode: true,
		RelativeRayDimension:            RelativeRayDOF,
		ConventionFieldCount:            ConventionFields,
		ComputesDeltaAlpha:              true,
		ComputesDeltaPhi:                true,
		ComputesProjectiveDistance:      true,
		ExportsResidualDiagnosticsOnly:  true,
		CKMMatrixElementExported:        false,
		CKMMatrixConstructed:            false,
		Verdict:                         StatusCKMNullMapDefined,
		Reason:                          "the CKM-null map computes only Delta_alpha, Delta_phi, and projective relative-ray distance after the Gate463 convention ledger is complete",
	}
}

func buildSieve() Sieve {
	cases := []Case{
		{Name: "valid synthetic convention-fixed u-d null residual", Input: validInput()},
		{Name: "missing relative u-d ray", Input: missingRelativeRayInput()},
		{Name: "missing eigenbasis convention", Input: missingEigenbasisInput()},
		{Name: "missing branch tag and provenance", Input: missingBranchProvenanceInput()},
		{Name: "observed CKM/PMNS import", Input: observedMixingInput()},
		{Name: "native CKM/PMNS prediction claim", Input: nativePredictionInput()},
		{Name: "CKM matrix export request", Input: matrixExportInput()},
		{Name: "raw diagonalizer claim", Input: rawDiagonalizerInput()},
		{Name: "degenerate spectrum", Input: degenerateSpectrumInput()},
		{Name: "K_gen basis rotation requested", Input: kgenRotationInput()},
		{Name: "native residual promotion", Input: nativeResidualPromotionInput()},
		{Name: "GST/Fritzsch used as CKM selector", Input: gstSelectorInput()},
	}

	s := Sieve{Executed: true, Cases: make([]Case, 0, len(cases))}
	for _, c := range cases {
		res, accepted, verdict, reason := EvaluateCKMNullResidual(c.Input)
		c.Residual = res
		c.Accepted = accepted
		c.Verdict = verdict
		c.Reason = reason
		if accepted {
			s.AcceptedCaseCount++
		} else {
			s.RejectedCaseCount++
		}
		s.ValidSyntheticResidualAccepted = s.ValidSyntheticResidualAccepted || (c.Name == "valid synthetic convention-fixed u-d null residual" && accepted && verdict == StatusCKMNullResidualComputed)
		s.MissingRelativeRayRejected = s.MissingRelativeRayRejected || verdict == StatusFailedRequiresRelativeRay
		s.MissingEigenbasisRejected = s.MissingEigenbasisRejected || verdict == StatusFailedRequiresEigenbasisConvention
		s.MissingBranchProvenanceRejected = s.MissingBranchProvenanceRejected || verdict == StatusFailedRequiresBranchAndProvenance
		s.ObservedCKMPMNSRejected = s.ObservedCKMPMNSRejected || verdict == StatusFailedObservedCKMPMNSImport
		s.NativePredictionRejected = s.NativePredictionRejected || verdict == StatusFailedCKMPMNSNativePrediction
		s.MatrixExportRejected = s.MatrixExportRejected || verdict == StatusFailedCKMMatrixExport
		s.RawDiagonalizerRejected = s.RawDiagonalizerRejected || verdict == StatusFailedRawDiagonalizer
		s.DegenerateSpectrumRejected = s.DegenerateSpectrumRejected || verdict == StatusFailedDegenerateSpectrum
		s.KGenBasisRotationRejected = s.KGenBasisRotationRejected || verdict == StatusFailedKGenBasisRotation
		s.NativeResidualPromotionRejected = s.NativeResidualPromotionRejected || verdict == StatusFailedNativeResidualPromotion
		s.GSTSelectorRejected = s.GSTSelectorRejected || verdict == StatusFailedGSTAsCKMSelector
		s.Cases = append(s.Cases, c)
	}
	s.AllAcceptedBridgeOnly = true
	s.NoCKMMatrixConstructed = true
	s.NoNativeObservableExport = true
	for _, c := range s.Cases {
		if c.Accepted && (!c.Residual.BridgeOnly || c.Residual.CKMMatrixConstructed || c.Residual.CKMEntryComputed || c.Residual.PMNSEntryComputed || c.Residual.ExportsNativeObservable) {
			s.AllAcceptedBridgeOnly = false
			s.NoCKMMatrixConstructed = false
			s.NoNativeObservableExport = false
		}
	}
	s.Verdict = StatusCKMNullResidualFirewallValidated
	s.Reason = "only a complete synthetic bridge row with Gate462 relative ray and Gate463 eigenbasis convention is accepted; the output is a null residual diagnostic, not V_CKM"
	return s
}

func validInput() CKMNullInput {
	return CKMNullInput{
		Name:                       "valid synthetic convention-fixed u-d null residual",
		AlphaU:                     0.25,
		PhiU:                       0.40,
		AlphaD:                     -0.10,
		PhiD:                       0.95,
		HasRelativeRay:             true,
		HasEigenbasisConvention:    true,
		HasOrderingConvention:      true,
		HasPhaseGaugeConvention:    true,
		HasNormalizationConvention: true,
		HasDegeneracyPolicy:        true,
		HasBranchTags:              true,
		HasProvenance:              true,
		BridgeOnly:                 true,
		SyntheticOnly:              true,
	}
}

func missingRelativeRayInput() CKMNullInput {
	x := validInput()
	x.HasRelativeRay = false
	return x
}

func missingEigenbasisInput() CKMNullInput {
	x := validInput()
	x.HasEigenbasisConvention = false
	return x
}

func missingBranchProvenanceInput() CKMNullInput {
	x := validInput()
	x.HasBranchTags = false
	x.HasProvenance = false
	return x
}

func observedMixingInput() CKMNullInput {
	x := validInput()
	x.ObservedCKMImport = true
	x.ObservedPMNSImport = true
	return x
}

func nativePredictionInput() CKMNullInput {
	x := validInput()
	x.CKMNativePredictionClaim = true
	x.PMNSNativePredictionClaim = true
	return x
}

func matrixExportInput() CKMNullInput {
	x := validInput()
	x.CKMMatrixExportRequested = true
	return x
}

func rawDiagonalizerInput() CKMNullInput {
	x := validInput()
	x.RawDiagonalizerClaim = true
	x.HasPhaseGaugeConvention = false
	return x
}

func degenerateSpectrumInput() CKMNullInput {
	x := validInput()
	x.DegenerateSpectrum = true
	return x
}

func kgenRotationInput() CKMNullInput {
	x := validInput()
	x.KGenBasisRotationRequested = true
	return x
}

func nativeResidualPromotionInput() CKMNullInput {
	x := validInput()
	x.ResidualNativePromotionClaim = true
	return x
}

func gstSelectorInput() CKMNullInput {
	x := validInput()
	x.GSTFritzschSelectorClaim = true
	return x
}

func EvaluateCKMNullResidual(in CKMNullInput) (CKMNullResidual, bool, string, string) {
	res := CKMNullResidual{}
	if in.ObservedCKMImport || in.ObservedPMNSImport {
		res.Verdict = StatusFailedObservedCKMPMNSImport
		res.Reason = "observed CKM/PMNS data are not accepted in the null residual harness"
		return res, false, res.Verdict, res.Reason
	}
	if in.CKMNativePredictionClaim || in.PMNSNativePredictionClaim {
		res.Verdict = StatusFailedCKMPMNSNativePrediction
		res.Reason = "a CKM/PMNS prediction claim is outside the native Gate464 contract"
		return res, false, res.Verdict, res.Reason
	}
	if in.CKMMatrixExportRequested {
		res.Verdict = StatusFailedCKMMatrixExport
		res.Reason = "Gate464 exports residual diagnostics only and cannot construct or export V_CKM"
		return res, false, res.Verdict, res.Reason
	}
	if in.ResidualNativePromotionClaim {
		res.Verdict = StatusFailedNativeResidualPromotion
		res.Reason = "CKM-null residuals are bridge diagnostics and cannot be promoted to native observables"
		return res, false, res.Verdict, res.Reason
	}
	if in.GSTFritzschSelectorClaim {
		res.Verdict = StatusFailedGSTAsCKMSelector
		res.Reason = "GST/Fritzsch relations are external texture assumptions and cannot be used as CKM branch selectors"
		return res, false, res.Verdict, res.Reason
	}
	if in.KGenBasisRotationRequested {
		res.Verdict = StatusFailedKGenBasisRotation
		res.Reason = "rotating away from the native K_gen address is not an allowed CKM residual convention"
		return res, false, res.Verdict, res.Reason
	}
	if in.RawDiagonalizerClaim || !in.HasPhaseGaugeConvention {
		res.Verdict = StatusFailedRawDiagonalizer
		res.Reason = "raw diagonalizers retain phase gauge and must not be used without the Gate463 convention ledger"
		return res, false, res.Verdict, res.Reason
	}
	if in.DegenerateSpectrum {
		res.Verdict = StatusFailedDegenerateSpectrum
		res.Reason = "degenerate spectra make eigenbasis residuals non-unique, so the adapter fails closed"
		return res, false, res.Verdict, res.Reason
	}
	if !in.HasRelativeRay {
		res.Verdict = StatusFailedRequiresRelativeRay
		res.Reason = "a Gate462 u-d relative ray is required before the null residual map can run"
		return res, false, res.Verdict, res.Reason
	}
	if !in.HasEigenbasisConvention || !in.HasOrderingConvention || !in.HasNormalizationConvention || !in.HasDegeneracyPolicy {
		res.Verdict = StatusFailedRequiresEigenbasisConvention
		res.Reason = "the Gate463 eigenbasis convention ledger must declare ordering, phase gauge, normalization, and degeneracy policy"
		return res, false, res.Verdict, res.Reason
	}
	if !in.HasBranchTags || !in.HasProvenance || !in.BridgeOnly || !in.SyntheticOnly {
		res.Verdict = StatusFailedRequiresBranchAndProvenance
		res.Reason = "complete branch tags, provenance, bridge_only=true, and synthetic/null mode are required"
		return res, false, res.Verdict, res.Reason
	}

	dalpha := in.AlphaD - in.AlphaU
	dphi := wrapPi(in.PhiD - in.PhiU)
	chord := 2 * math.Sin(dphi/2)
	dist := math.Sqrt(dalpha*dalpha + chord*chord)
	res = CKMNullResidual{
		DeltaAlpha:                dalpha,
		DeltaPhi:                  dphi,
		PhaseChord:                chord,
		ProjectiveRayDistance:     dist,
		NullAlignmentResidual:     dist,
		CompleteInputs:            true,
		ConventionFixed:           true,
		BridgeOnly:                true,
		SyntheticOnly:             true,
		ExportsRelativeDiagnostic: true,
		CKMMatrixConstructed:      false,
		CKMEntryComputed:          false,
		PMNSEntryComputed:         false,
		ExportsNativeObservable:   false,
		ObservedDataImported:      false,
		Verdict:                   StatusCKMNullResidualComputed,
		Reason:                    "computed only the convention-fixed relative-ray residual d_ud; no CKM/PMNS matrix element was constructed",
	}
	return res, true, res.Verdict, res.Reason
}

func wrapPi(x float64) float64 {
	for x <= -math.Pi {
		x += 2 * math.Pi
	}
	for x > math.Pi {
		x -= 2 * math.Pi
	}
	return x
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                        true,
		CKMNullResidualAdapterDefined:   a.Map.Executed,
		CKMNullResidualMayRunBridgeOnly: a.Sieve.ValidSyntheticResidualAccepted,
		CKMMatrixConstructed:            false,
		CKMMatrixEntryComputed:          false,
		CKMMatrixEntryNative:            false,
		PMNSMatrixEntryComputed:         false,
		PMNSMatrixEntryNative:           false,
		ObservedMassesImported:          false,
		ObservedYukawasImported:         false,
		ObservedCKMImported:             false,
		ObservedPMNSImported:            false,
		GSTFritzschPromoted:             false,
		RelativeRayPromotedNative:       false,
		EigenbasisPromotedNative:        false,
		KGenStillForced:                 a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:            a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:          true,
		SectorCoefficientsStillSealed:   true,
		NativeFlavorDimAfter:            NativeFlavorDim,
		KXYCoeffDimAfter:                KXYCoeffDim,
		Verdict:                         StatusFirewallPreserved,
		Reason:                          "Gate464 composes relative-ray and convention ledgers into a synthetic bridge residual only; V_CKM, U_PMNS, masses, Yukawas, GST assumptions, and coefficient values remain sealed.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        465,
		Title:       "Quark-Sector Empirical Import Switch / CKM Data Firewall",
		Reason:      "Gate464 now has a safe null residual map; the next boundary is an explicit switch that distinguishes synthetic/null mode from any future observed CKM comparator import.",
		PrimaryTask: "define a fail-closed observed-data import path for CKM-facing bridge comparators with source, scheme, scale, uncertainty, and native-promotion rejection",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate462RelativeRay || !a.Inheritance.Gate463EigenbasisLedger || !a.Inheritance.Gate463ReadyForResidualAdapter || !a.Inheritance.NoObservedValuesImported {
		return fmt.Errorf("Gate463 inheritance incomplete")
	}
	if !a.Map.Executed || !a.Map.RequiresRelativeRay || !a.Map.RequiresEigenbasisConvention || !a.Map.RequiresOrderingConvention || !a.Map.RequiresPhaseGaugeConvention || !a.Map.RequiresBranchTags || !a.Map.RequiresProvenance || !a.Map.RequiresBridgeOnlySyntheticMode || !a.Map.ExportsResidualDiagnosticsOnly || a.Map.CKMMatrixConstructed || a.Map.CKMMatrixElementExported {
		return fmt.Errorf("symbolic CKM-null map contract incomplete")
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCaseCount != 1 || a.Sieve.RejectedCaseCount != 11 || !a.Sieve.ValidSyntheticResidualAccepted || !a.Sieve.AllAcceptedBridgeOnly || !a.Sieve.NoCKMMatrixConstructed || !a.Sieve.NoNativeObservableExport {
		return fmt.Errorf("sieve did not fail closed")
	}
	if !(a.Sieve.MissingRelativeRayRejected && a.Sieve.MissingEigenbasisRejected && a.Sieve.MissingBranchProvenanceRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.MatrixExportRejected && a.Sieve.RawDiagonalizerRejected && a.Sieve.DegenerateSpectrumRejected && a.Sieve.KGenBasisRotationRejected && a.Sieve.NativeResidualPromotionRejected && a.Sieve.GSTSelectorRejected) {
		return fmt.Errorf("not all unsafe CKM residual routes were rejected")
	}
	if !a.Firewall.Executed || !a.Firewall.CKMNullResidualAdapterDefined || !a.Firewall.CKMNullResidualMayRunBridgeOnly || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMMatrixEntryComputed || a.Firewall.CKMMatrixEntryNative || a.Firewall.PMNSMatrixEntryComputed || a.Firewall.PMNSMatrixEntryNative || a.Firewall.ObservedCKMImported || a.Firewall.ObservedPMNSImported || a.Firewall.GSTFritzschPromoted || a.Firewall.RelativeRayPromotedNative || a.Firewall.EigenbasisPromotedNative || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall violated")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.ValidSyntheticResidualAccepted && a.Firewall.CKMNullResidualMayRunBridgeOnly && !a.Firewall.CKMMatrixEntryComputed && !a.Firewall.ObservedCKMImported {
		return "Gate 464 validates a CKM-null residual adapter: with a Gate462 relative ray and Gate463 eigenbasis convention, the engine may compute a synthetic bridge-only d_ud residual. That residual is not V_CKM, not a CKM entry, and not a native prediction."
	}
	return "Gate 464 failed to validate the CKM-null residual firewall."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch_tags=%t multiplex=%t relative_ray=%t rel_rejects_observed=%t rel_rejects_native=%t eigenbasis=%t raw_rejected=%t permutation_rejected=%t ready=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456RayInverse, x.Gate459BranchTags, x.Gate461SectorMultiplex, x.Gate462RelativeRay, x.Gate462RejectsObservedCKMPMNS, x.Gate462RejectsNativePrediction, x.Gate463EigenbasisLedger, x.Gate463RejectsRawGauge, x.Gate463RejectsPermutationNative, x.Gate463ReadyForResidualAdapter, x.NoObservedValuesImported, x.Verdict)
}

func FormatMap(x SymbolicMap) string {
	return fmt.Sprintf("executed=%t relative_ray=%t eigenbasis=%t ordering=%t phase_gauge=%t branch_tags=%t provenance=%t synthetic=%t relative_dof=%d convention_fields=%d delta_alpha=%t delta_phi=%t distance=%t diagnostics_only=%t CKM_element_exported=%t CKM_constructed=%t verdict=%s reason=%s", x.Executed, x.RequiresRelativeRay, x.RequiresEigenbasisConvention, x.RequiresOrderingConvention, x.RequiresPhaseGaugeConvention, x.RequiresBranchTags, x.RequiresProvenance, x.RequiresBridgeOnlySyntheticMode, x.RelativeRayDimension, x.ConventionFieldCount, x.ComputesDeltaAlpha, x.ComputesDeltaPhi, x.ComputesProjectiveDistance, x.ExportsResidualDiagnosticsOnly, x.CKMMatrixElementExported, x.CKMMatrixConstructed, x.Verdict, x.Reason)
}

func FormatInput(x CKMNullInput) string {
	return fmt.Sprintf("alpha_u=%.12g phi_u=%.12g alpha_d=%.12g phi_d=%.12g relative=%t eigenbasis=%t ordering=%t phase_gauge=%t norm=%t degeneracy_policy=%t tags=%t provenance=%t bridge=%t synthetic=%t observed_CKM=%t observed_PMNS=%t native_CKM=%t native_PMNS=%t export_matrix=%t raw=%t degenerate=%t K_rotation=%t native_residual=%t GST_selector=%t", x.AlphaU, x.PhiU, x.AlphaD, x.PhiD, x.HasRelativeRay, x.HasEigenbasisConvention, x.HasOrderingConvention, x.HasPhaseGaugeConvention, x.HasNormalizationConvention, x.HasDegeneracyPolicy, x.HasBranchTags, x.HasProvenance, x.BridgeOnly, x.SyntheticOnly, x.ObservedCKMImport, x.ObservedPMNSImport, x.CKMNativePredictionClaim, x.PMNSNativePredictionClaim, x.CKMMatrixExportRequested, x.RawDiagonalizerClaim, x.DegenerateSpectrum, x.KGenBasisRotationRequested, x.ResidualNativePromotionClaim, x.GSTFritzschSelectorClaim)
}

func FormatResidual(x CKMNullResidual) string {
	return fmt.Sprintf("Delta_alpha=%.12g Delta_phi=%.12g phase_chord=%.12g d_ud=%.12g R_null=%.12g complete=%t convention=%t bridge=%t synthetic=%t diagnostic=%t CKM_constructed=%t CKM_entry=%t PMNS_entry=%t native_export=%t observed_import=%t verdict=%s reason=%s", x.DeltaAlpha, x.DeltaPhi, x.PhaseChord, x.ProjectiveRayDistance, x.NullAlignmentResidual, x.CompleteInputs, x.ConventionFixed, x.BridgeOnly, x.SyntheticOnly, x.ExportsRelativeDiagnostic, x.CKMMatrixConstructed, x.CKMEntryComputed, x.PMNSEntryComputed, x.ExportsNativeObservable, x.ObservedDataImported, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d valid=%t missing_ray=%t missing_eigenbasis=%t missing_branch_provenance=%t observed=%t native_prediction=%t matrix_export=%t raw=%t degenerate=%t K_rotation=%t native_residual=%t GST_selector=%t bridge_only=%t no_CKM_matrix=%t no_native=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.ValidSyntheticResidualAccepted, x.MissingRelativeRayRejected, x.MissingEigenbasisRejected, x.MissingBranchProvenanceRejected, x.ObservedCKMPMNSRejected, x.NativePredictionRejected, x.MatrixExportRejected, x.RawDiagonalizerRejected, x.DegenerateSpectrumRejected, x.KGenBasisRotationRejected, x.NativeResidualPromotionRejected, x.GSTSelectorRejected, x.AllAcceptedBridgeOnly, x.NoCKMMatrixConstructed, x.NoNativeObservableExport, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t adapter_defined=%t bridge_only_run=%t CKM_constructed=%t CKM_entry=%t CKM_native=%t PMNS_entry=%t PMNS_native=%t masses_imported=%t yukawas_imported=%t CKM_imported=%t PMNS_imported=%t GST_promoted=%t relative_native=%t eigenbasis_native=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.CKMNullResidualAdapterDefined, x.CKMNullResidualMayRunBridgeOnly, x.CKMMatrixConstructed, x.CKMMatrixEntryComputed, x.CKMMatrixEntryNative, x.PMNSMatrixEntryComputed, x.PMNSMatrixEntryNative, x.ObservedMassesImported, x.ObservedYukawasImported, x.ObservedCKMImported, x.ObservedPMNSImported, x.GSTFritzschPromoted, x.RelativeRayPromotedNative, x.EigenbasisPromotedNative, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{
		StatusGate463Inherited,
		StatusCKMNullMapDefined,
		StatusCKMNullResidualComputed,
		StatusCKMNullResidualFirewallValidated,
		StatusFirewallPreserved,
		StatusFailedRequiresRelativeRay,
		StatusFailedRequiresEigenbasisConvention,
		StatusFailedRequiresBranchAndProvenance,
		StatusFailedObservedCKMPMNSImport,
		StatusFailedCKMPMNSNativePrediction,
		StatusFailedCKMMatrixExport,
		StatusFailedRawDiagonalizer,
		StatusFailedDegenerateSpectrum,
		StatusFailedKGenBasisRotation,
		StatusFailedNativeResidualPromotion,
		StatusFailedGSTAsCKMSelector,
	}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 464 Registry Audit — CKM Null Residual Adapter / Convention-Ready Symbolic Map\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusCKMNullResidualFirewallValidated + "`\n\n")
	b.WriteString("Gate 464 composes the Gate 462 relative-ray diagnostic with the Gate 463 eigenbasis convention ledger. It computes only a synthetic bridge residual and does not construct or export CKM/PMNS matrix entries.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Symbolic/null map\n\n")
	b.WriteString(FormatMap(a.Map) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("Delta_alpha_ud = alpha_d - alpha_u\n")
	b.WriteString("Delta_phi_ud   = wrap_pi(phi_d - phi_u)\n")
	b.WriteString("d_ud           = sqrt(Delta_alpha_ud^2 + 4 sin^2(Delta_phi_ud/2))\n")
	b.WriteString("R_null         = d_ud\n")
	b.WriteString("forbidden      = V_CKM, CKM_ij, U_PMNS, PMNS_ij, observed masses, Yukawas, GST/Fritzsch selectors\n")
	b.WriteString("```\n\n")
	b.WriteString("The residual is a comparator diagnostic on bridge rays. It is not a CKM angle, not a CKM matrix element, and not a native flavor observable.\n\n")

	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Case | Accepted | Verdict | Input | Residual | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s | %s | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(FormatInput(c.Input)), esc(FormatResidual(c.Residual)), esc(c.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Boundary statement\n\n")
	b.WriteString("Gate 464 is the strongest safe CKM-facing computation currently permitted by the post-444 family board: it can measure relative bridge-ray separation after conventions are declared, but it cannot turn that separation into V_CKM without an explicit future observed-comparator import switch.\n\n")

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
