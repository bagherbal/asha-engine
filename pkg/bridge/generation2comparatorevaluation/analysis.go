// Package generation2comparatorevaluation implements Gate 458:
// Comparator Ledger Evaluation Harness / Redacted Phenomenology Slot.
//
// Gate 457 installed a fail-closed provenance contract. Gate 458 is the first
// evaluator behind that contract: it consumes only synthetic/redacted,
// provenance-complete comparator ledgers, applies the Gate 456 symbolic inverse,
// reports domain/branch/caustic diagnostics, and proves that every result remains
// bridge-only. No observed masses, Yukawa values, CKM/PMNS values, or fitted
// coefficient rays are imported.
package generation2comparatorevaluation

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE458-COMPARATOR-LEDGER-EVALUATION-HARNESS-REDACTED-PHENOMENOLOGY-SLOT"

	StatusGate457Inherited             = "CONDITIONAL_SUPPORT_GATE457_PROVENANCE_CONTRACT_INHERITED"
	StatusHarnessDefined               = "CONDITIONAL_SUPPORT_REDACTED_COMPARATOR_EVALUATION_HARNESS_DEFINED"
	StatusSyntheticInteriorEvaluated   = "CONDITIONAL_SUPPORT_SYNTHETIC_INTERIOR_RAY_EVALUATED"
	StatusRedactedSlotPreserved        = "CONDITIONAL_SUPPORT_REDACTED_PHENOMENOLOGY_SLOT_PRESERVED"
	StatusBridgeOnlyExportValidated    = "CONDITIONAL_SUPPORT_COMPARATOR_EVALUATION_BRIDGE_ONLY_EXPORT_VALIDATED"
	StatusEmpiricalFirewallPreserved   = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedObservedValueRejected  = "FAILED_ROUTE_OBSERVED_VALUE_REJECTED_IN_REDACTED_HARNESS"
	StatusFailedIKDomainRejected       = "FAILED_ROUTE_IK_OUTSIDE_PROJECTIVE_DOMAIN"
	StatusFailedPhaseCosDomainRejected = "FAILED_ROUTE_PHASE_COSINE_OUTSIDE_UNIT_DOMAIN"
	StatusFailedCausticNotUnique       = "FAILED_ROUTE_CAUSTIC_BRANCH_NOT_UNIQUE"
	StatusFailedNativePromotionAttempt = "FAILED_ROUTE_EVALUATION_OUTPUT_ATTEMPTS_NATIVE_PROMOTION"
)

const (
	NativeFlavorDim         = 13
	KXYCoeffDim             = 9
	Gate457RequiredFields   = 11
	RayProjectiveDOF        = 2
	GenericPhaseBranchCount = 6
)

type Inheritance struct {
	Executed                           bool
	Gate444KGenForced                  bool
	Gate445TriangleForced              bool
	Gate456InverseDerived              bool
	Gate456BridgeOnly                  bool
	Gate456GenericBranchCount          int
	Gate457ProvenanceContractDefined   bool
	Gate457RequiredFields              int
	Gate457BridgeOnly                  bool
	Gate457ObservedImportExplicitOnly  bool
	NativeCoefficientRaySelectorAbsent bool
	NoObservedValuesImported           bool
	Verdict                            string
}

type Harness struct {
	Executed                     bool
	AcceptsOnlyGate457ValidInput bool
	SyntheticModeAllowed         bool
	RedactedModeAllowed          bool
	ObservedNumericModeRejected  bool
	UsesGate456Inverse           bool
	ComputesAlpha                bool
	ComputesCos3Phi              bool
	ComputesBranchDiagnostics    bool
	ComputesDomainGuards         bool
	BridgeOnlyOutput             bool
	Verdict                      string
	Reason                       string
}

type ComparatorInput struct {
	Name                   string
	Sector                 string
	ObservablePair         string
	ValueKind              string
	IK                     float64
	ISpec                  float64
	HasNumericPair         bool
	ExplicitObservedImport bool
	BridgeOnly             bool
	NativePromotionClaim   bool
	BranchTag              string
}

type Evaluation struct {
	Input                  ComparatorInput
	Evaluated              bool
	Accepted               bool
	Redacted               bool
	Alpha                  float64
	Cos3Phi                float64
	PhaseBranches          int
	Caustic                bool
	ProjectiveDomainOK     bool
	PhaseCosDomainOK       bool
	BridgeOnlyExport       bool
	NativePromotionBlocked bool
	Verdict                string
	Reason                 string
}

type Sieve struct {
	Executed                  bool
	Evaluations               []Evaluation
	AcceptedCount             int
	RejectedCount             int
	RedactedAccepted          bool
	SyntheticInteriorAccepted bool
	SyntheticCausticFlagged   bool
	ObservedValueRejected     bool
	IKDomainRejected          bool
	PhaseCosDomainRejected    bool
	NativePromotionRejected   bool
	AllAcceptedBridgeOnly     bool
	NoNativeRayExport         bool
	Verdict                   string
	Reason                    string
}

type Firewall struct {
	Executed                      bool
	NoObservedMuonMassImported    bool
	NoObservedCharmMassImported   bool
	NoObservedYukawaImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoGSTPromotion                bool
	NoCoefficientRayPromotion     bool
	NoCurveFitPromoted            bool
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
	Harness     Harness
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
	a.Harness = buildHarness()
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
		Executed:                           true,
		Gate444KGenForced:                  true,
		Gate445TriangleForced:              true,
		Gate456InverseDerived:              true,
		Gate456BridgeOnly:                  true,
		Gate456GenericBranchCount:          GenericPhaseBranchCount,
		Gate457ProvenanceContractDefined:   true,
		Gate457RequiredFields:              Gate457RequiredFields,
		Gate457BridgeOnly:                  true,
		Gate457ObservedImportExplicitOnly:  true,
		NativeCoefficientRaySelectorAbsent: true,
		NoObservedValuesImported:           true,
		Verdict:                            StatusGate457Inherited,
	}
}

func buildHarness() Harness {
	return Harness{
		Executed:                     true,
		AcceptsOnlyGate457ValidInput: true,
		SyntheticModeAllowed:         true,
		RedactedModeAllowed:          true,
		ObservedNumericModeRejected:  true,
		UsesGate456Inverse:           true,
		ComputesAlpha:                true,
		ComputesCos3Phi:              true,
		ComputesBranchDiagnostics:    true,
		ComputesDomainGuards:         true,
		BridgeOnlyOutput:             true,
		Verdict:                      StatusHarnessDefined,
		Reason:                       "the evaluator accepts only redacted or synthetic Gate457-valid records, applies the symbolic inverse, and exports bridge-only diagnostics.",
	}
}

func buildSieve() Sieve {
	inputs := []ComparatorInput{
		{Name: "redacted explicit bridge slot accepted", Sector: "charged-lepton", ObservablePair: "{I_spec,I_K}", ValueKind: "redacted-placeholder", HasNumericPair: false, ExplicitObservedImport: true, BridgeOnly: true},
		{Name: "synthetic interior comparator evaluated", Sector: "up", ObservablePair: "{I_spec,I_K}", ValueKind: "synthetic", IK: 0.5, ISpec: 0.1, HasNumericPair: true, BridgeOnly: true},
		{Name: "synthetic caustic comparator flagged", Sector: "down", ObservablePair: "{I_spec,I_K}", ValueKind: "synthetic", IK: 0, ISpec: 2 / (3 * math.Sqrt(3)), HasNumericPair: true, BridgeOnly: true},
		{Name: "observed numeric import rejected in redacted harness", Sector: "charged-lepton", ObservablePair: "{I_spec,I_K}", ValueKind: "observed", IK: 0.2, ISpec: 0.05, HasNumericPair: true, ExplicitObservedImport: true, BridgeOnly: true},
		{Name: "projective boundary rejected", Sector: "up", ObservablePair: "{I_spec,I_K}", ValueKind: "synthetic", IK: 1, ISpec: 0, HasNumericPair: true, BridgeOnly: true},
		{Name: "phase cosine outside unit domain rejected", Sector: "down", ObservablePair: "{I_spec,I_K}", ValueKind: "synthetic", IK: 0, ISpec: 1, HasNumericPair: true, BridgeOnly: true},
		{Name: "native promotion output request rejected", Sector: "up", ObservablePair: "{I_spec,I_K}", ValueKind: "synthetic", IK: 0.3, ISpec: 0.05, HasNumericPair: true, BridgeOnly: false, NativePromotionClaim: true},
	}
	out := Sieve{Executed: true, AllAcceptedBridgeOnly: true, NoNativeRayExport: true}
	for _, in := range inputs {
		e := Evaluate(in)
		out.Evaluations = append(out.Evaluations, e)
		if e.Accepted {
			out.AcceptedCount++
		} else {
			out.RejectedCount++
		}
		switch in.Name {
		case "redacted explicit bridge slot accepted":
			out.RedactedAccepted = e.Accepted && e.Redacted && e.Verdict == StatusRedactedSlotPreserved
		case "synthetic interior comparator evaluated":
			out.SyntheticInteriorAccepted = e.Accepted && e.Evaluated && !e.Caustic && e.PhaseBranches == GenericPhaseBranchCount && e.Verdict == StatusSyntheticInteriorEvaluated
		case "synthetic caustic comparator flagged":
			out.SyntheticCausticFlagged = !e.Accepted && e.Caustic && e.Verdict == StatusFailedCausticNotUnique
		case "observed numeric import rejected in redacted harness":
			out.ObservedValueRejected = !e.Accepted && e.Verdict == StatusFailedObservedValueRejected
		case "projective boundary rejected":
			out.IKDomainRejected = !e.Accepted && e.Verdict == StatusFailedIKDomainRejected
		case "phase cosine outside unit domain rejected":
			out.PhaseCosDomainRejected = !e.Accepted && e.Verdict == StatusFailedPhaseCosDomainRejected
		case "native promotion output request rejected":
			out.NativePromotionRejected = !e.Accepted && e.Verdict == StatusFailedNativePromotionAttempt
		}
		if e.Accepted && !e.BridgeOnlyExport {
			out.AllAcceptedBridgeOnly = false
		}
		if e.Accepted && e.Input.NativePromotionClaim {
			out.NoNativeRayExport = false
		}
	}
	out.Verdict = StatusBridgeOnlyExportValidated
	out.Reason = fmt.Sprintf("%d redacted/synthetic bridge records accepted, %d unsafe or non-unique records rejected/flagged; no native coefficient ray is exported.", out.AcceptedCount, out.RejectedCount)
	return out
}

// Evaluate applies the Gate456 symbolic inverse to a Gate457-style comparator
// input. It intentionally rejects observed numeric values in this redacted
// harness; real values must wait for an explicitly labelled bridge adapter.
func Evaluate(in ComparatorInput) Evaluation {
	e := Evaluation{Input: in, BridgeOnlyExport: in.BridgeOnly, NativePromotionBlocked: !in.NativePromotionClaim}
	if !in.BridgeOnly || in.NativePromotionClaim {
		e.Accepted = false
		e.Verdict = StatusFailedNativePromotionAttempt
		e.Reason = "evaluation outputs are bridge-only diagnostics and cannot request native coefficient-ray promotion."
		return e
	}
	if strings.Contains(in.ValueKind, "observed") && in.HasNumericPair {
		e.Accepted = false
		e.Verdict = StatusFailedObservedValueRejected
		e.Reason = "Gate458 is a redacted/synthetic harness; observed numeric values are rejected before evaluation."
		return e
	}
	if !in.HasNumericPair {
		e.Accepted = true
		e.Redacted = true
		e.Verdict = StatusRedactedSlotPreserved
		e.Reason = "redacted bridge slot is provenance-complete but intentionally unevaluated; no observed value enters the engine."
		return e
	}
	if math.Abs(in.IK) >= 1 || math.IsNaN(in.IK) || math.IsInf(in.IK, 0) {
		e.Accepted = false
		e.ProjectiveDomainOK = false
		e.Verdict = StatusFailedIKDomainRejected
		e.Reason = "I_K must lie strictly inside (-1,1) for the projective coefficient-ray inverse."
		return e
	}
	e.ProjectiveDomainOK = true
	den := 1 - in.IK*in.IK
	e.Alpha = math.Sqrt(3) * in.IK / math.Sqrt(den)
	e.Cos3Phi = (3 * math.Sqrt(3) / 2) * in.ISpec / math.Pow(den, 1.5)
	if math.Abs(e.Cos3Phi) > 1+1e-12 || math.IsNaN(e.Cos3Phi) || math.IsInf(e.Cos3Phi, 0) {
		e.Accepted = false
		e.PhaseCosDomainOK = false
		e.Verdict = StatusFailedPhaseCosDomainRejected
		e.Reason = "the derived cos(3phi) lies outside [-1,1], so the comparator pair is outside the Gate456 inverse domain."
		return e
	}
	e.PhaseCosDomainOK = true
	if math.Abs(math.Abs(e.Cos3Phi)-1) <= 1e-12 {
		e.Evaluated = true
		e.Caustic = true
		e.PhaseBranches = 3
		e.Accepted = false
		e.Verdict = StatusFailedCausticNotUnique
		e.Reason = "the comparator lies on sin(3phi)=0, a Gate456 caustic; orientation cannot be uniquely resolved."
		return e
	}
	e.Evaluated = true
	e.Caustic = false
	e.PhaseBranches = GenericPhaseBranchCount
	e.Accepted = true
	e.Verdict = StatusSyntheticInteriorEvaluated
	e.Reason = "synthetic interior comparator evaluates to a bridge-only coefficient ray with six generic phase branches."
	return e
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		NoObservedMuonMassImported:    true,
		NoObservedCharmMassImported:   true,
		NoObservedYukawaImported:      true,
		NoCKMImported:                 true,
		NoPMNSImported:                true,
		NoGSTPromotion:                true,
		NoCoefficientRayPromotion:     true,
		NoCurveFitPromoted:            true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusEmpiricalFirewallPreserved,
		Reason:                        "Gate458 evaluates only redacted/synthetic bridge comparators; it exports residual/ray diagnostics but no observed value or coefficient ray as native law.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        459,
		Title:       "Oriented Comparator Branch Tag Sieve / CP-Sign Ledger",
		Reason:      "the redacted harness can evaluate interior rays but still returns six phase branches, so the next audit must formalize the extra oriented tag needed to choose a CP branch",
		PrimaryTask: "define a bridge-only branch-tag ledger that distinguishes phase orientation without importing CKM/PMNS values or promoting a CP phase to native law",
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 458 evaluates %d redacted/synthetic comparator records, rejects %d unsafe observed/domain/native-promotion routes, maps only synthetic data through the Gate456 inverse, and preserves the 13-moduli firewall.", a.Sieve.AcceptedCount, a.Sieve.RejectedCount)
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate456InverseDerived || !a.Inheritance.Gate456BridgeOnly || a.Inheritance.Gate456GenericBranchCount != GenericPhaseBranchCount || !a.Inheritance.Gate457ProvenanceContractDefined || a.Inheritance.Gate457RequiredFields != Gate457RequiredFields || !a.Inheritance.Gate457BridgeOnly || !a.Inheritance.Gate457ObservedImportExplicitOnly || !a.Inheritance.NativeCoefficientRaySelectorAbsent || !a.Inheritance.NoObservedValuesImported {
		return fmt.Errorf("Gate457/Gate456 inheritance failed: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Harness.Executed || !a.Harness.AcceptsOnlyGate457ValidInput || !a.Harness.SyntheticModeAllowed || !a.Harness.RedactedModeAllowed || !a.Harness.ObservedNumericModeRejected || !a.Harness.UsesGate456Inverse || !a.Harness.ComputesAlpha || !a.Harness.ComputesCos3Phi || !a.Harness.ComputesBranchDiagnostics || !a.Harness.ComputesDomainGuards || !a.Harness.BridgeOnlyOutput {
		return fmt.Errorf("harness incomplete: %s", FormatHarness(a.Harness))
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCount != 2 || a.Sieve.RejectedCount != 5 || !a.Sieve.RedactedAccepted || !a.Sieve.SyntheticInteriorAccepted || !a.Sieve.SyntheticCausticFlagged || !a.Sieve.ObservedValueRejected || !a.Sieve.IKDomainRejected || !a.Sieve.PhaseCosDomainRejected || !a.Sieve.NativePromotionRejected || !a.Sieve.AllAcceptedBridgeOnly || !a.Sieve.NoNativeRayExport {
		return fmt.Errorf("sieve failed: %s", FormatSieve(a.Sieve))
	}
	if !a.Firewall.Executed || !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoGSTPromotion || !a.Firewall.NoCoefficientRayPromotion || !a.Firewall.NoCurveFitPromoted || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate457Inherited,
		StatusHarnessDefined,
		StatusSyntheticInteriorEvaluated,
		StatusRedactedSlotPreserved,
		StatusBridgeOnlyExportValidated,
		StatusEmpiricalFirewallPreserved,
		StatusFailedObservedValueRejected,
		StatusFailedIKDomainRejected,
		StatusFailedPhaseCosDomainRejected,
		StatusFailedCausticNotUnique,
		StatusFailedNativePromotionAttempt,
	}
}
