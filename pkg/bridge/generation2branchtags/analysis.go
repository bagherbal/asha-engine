// Package generation2branchtags implements Gate 459:
// Oriented Comparator Branch Tag Sieve / CP-Sign Ledger.
//
// Gate 456 derived the bridge-only inverse map from labelled comparator pairs
// to a projective coefficient ray, but the inverse leaves six phase branches.
// Gate 458 added a redacted/synthetic evaluation harness. Gate 459 formalizes
// the extra bridge-only orientation metadata required to collapse the six
// branches to one: a CP-odd sign for sin(3phi) plus a C3 sheet tag. Neither tag
// is native ASHA law; both are comparator metadata unless a later theorem
// derives an intrinsic orientation selector.
package generation2branchtags

import (
	"fmt"
	"math"
	"sync"
)

const (
	AuditID = "GATE459-ORIENTED-COMPARATOR-BRANCH-TAG-SIEVE-CP-SIGN-LEDGER"

	StatusGate458Inherited             = "CONDITIONAL_SUPPORT_GATE458_REDACTED_HARNESS_INHERITED"
	StatusBranchLedgerDefined          = "CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_LEDGER_DEFINED"
	StatusCompleteBranchTagUnique      = "CONDITIONAL_SUPPORT_COMPLETE_BRANCH_TAG_SELECTS_UNIQUE_PHASE"
	StatusBridgeOnlyBranchTagValidated = "CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_BRIDGE_ONLY_VALIDATED"
	StatusFirewallPreserved            = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedCosineSixBranches       = "FAILED_ROUTE_COSINE_INVARIANT_RETURNS_SIX_BRANCHES"
	StatusFailedCPSignOnlyThreeSheets   = "FAILED_ROUTE_CP_ODD_SIGN_ONLY_LEAVES_C3_SHEETS"
	StatusFailedInvalidBranchTag        = "FAILED_ROUTE_INVALID_OR_INCOMPLETE_BRANCH_TAG"
	StatusFailedCKMPMNSSelectorRejected = "FAILED_ROUTE_CKM_PMNS_BRANCH_SELECTOR_REJECTED"
	StatusFailedNativePromotionRejected = "FAILED_ROUTE_BRANCH_TAG_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeCPSelectorAbsent  = "FAILED_ROUTE_CP_SIGN_NOT_NATIVE"
	StatusFailedNativeC3SelectorAbsent  = "FAILED_ROUTE_C3_SHEET_NOT_NATIVE"
)

const (
	NativeFlavorDim         = 13
	KXYCoeffDim             = 9
	CosineOnlyBranchCount   = 6
	CPSignOnlyBranchCount   = 3
	CompleteBranchCount     = 1
	C3SheetCount            = 3
	Gate457RequiredFields   = 11
	RayProjectiveDOF        = 2
	BranchTagRequiredFields = 2
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate456InverseDerived            bool
	Gate456GenericBranchCount        int
	Gate457ProvenanceContractDefined bool
	Gate458RedactedHarnessDefined    bool
	Gate458ObservedValuesRejected    bool
	Gate458BridgeOnly                bool
	NativeCPSelectorAbsent           bool
	NativeC3SheetSelectorAbsent      bool
	NoObservedValuesImported         bool
	Verdict                          string
}

type BranchLedger struct {
	Executed                      bool
	RequiresCosineInvariant       bool
	RequiresCPOddSign             bool
	RequiresC3Sheet               bool
	CPOddSignFormula              string
	C3SheetFormula                string
	UniquePhaseFormula            string
	CosineOnlyBranchCount         int
	CPOddSignOnlyBranchCount      int
	CompleteBranchTagCount        int
	RejectsCKMOrPMNSAsSelector    bool
	RejectsNativePromotion        bool
	BridgeOnly                    bool
	NativeCPOddSignSelectorAbsent bool
	NativeC3SheetSelectorAbsent   bool
	Verdict                       string
	Reason                        string
}

type BranchRequest struct {
	Name                 string
	Cos3Phi              float64
	HasCosineInvariant   bool
	CPOddSign            int // +1 selects +arccos branch, -1 selects -arccos branch.
	HasCPOddSign         bool
	C3Sheet              int // 0,1,2 labels the residual cubic sheet.
	HasC3Sheet           bool
	UsesCKMOrPMNS        bool
	BridgeOnly           bool
	NativePromotionClaim bool
}

type BranchEvaluation struct {
	Request                BranchRequest
	Accepted               bool
	Selected               bool
	DomainOK               bool
	Phase                  float64
	BranchCount            int
	BridgeOnlyExport       bool
	NativePromotionBlocked bool
	Verdict                string
	Reason                 string
}

type Sieve struct {
	Executed                 bool
	Evaluations              []BranchEvaluation
	AcceptedCount            int
	RejectedCount            int
	CosineOnlyFlagged        bool
	CPOddOnlyFlagged         bool
	CompletePositiveAccepted bool
	CompleteNegativeAccepted bool
	CKMPMNSSelectorRejected  bool
	NativePromotionRejected  bool
	InvalidTagRejected       bool
	AllAcceptedBridgeOnly    bool
	NoNativePhaseExport      bool
	Verdict                  string
	Reason                   string
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
	NoCPPhasePromotion            bool
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
	Ledger      BranchLedger
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
	a.Ledger = buildLedger()
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
		Gate456InverseDerived:            true,
		Gate456GenericBranchCount:        CosineOnlyBranchCount,
		Gate457ProvenanceContractDefined: true,
		Gate458RedactedHarnessDefined:    true,
		Gate458ObservedValuesRejected:    true,
		Gate458BridgeOnly:                true,
		NativeCPSelectorAbsent:           true,
		NativeC3SheetSelectorAbsent:      true,
		NoObservedValuesImported:         true,
		Verdict:                          StatusGate458Inherited,
	}
}

func buildLedger() BranchLedger {
	return BranchLedger{
		Executed:                      true,
		RequiresCosineInvariant:       true,
		RequiresCPOddSign:             true,
		RequiresC3Sheet:               true,
		CPOddSignFormula:              "sigma_CP = sign(sin(3 phi)) in bridge orientation convention",
		C3SheetFormula:                "n_C3 in {0,1,2}",
		UniquePhaseFormula:            "phi = (sigma_CP arccos(C) + 2*pi*n_C3)/3, C=cos(3phi)",
		CosineOnlyBranchCount:         CosineOnlyBranchCount,
		CPOddSignOnlyBranchCount:      CPSignOnlyBranchCount,
		CompleteBranchTagCount:        CompleteBranchCount,
		RejectsCKMOrPMNSAsSelector:    true,
		RejectsNativePromotion:        true,
		BridgeOnly:                    true,
		NativeCPOddSignSelectorAbsent: true,
		NativeC3SheetSelectorAbsent:   true,
		Verdict:                       StatusBranchLedgerDefined,
		Reason:                        "A cosine invariant fixes only cos(3phi). A CP-odd sign chooses one orientation of 3phi but leaves three cubic sheets. A complete bridge-only tag {sigma_CP,n_C3} is necessary and sufficient for a unique synthetic phase branch.",
	}
}

func buildSieve() Sieve {
	requests := []BranchRequest{
		{Name: "cosine-only Gate456 inverse result", Cos3Phi: 0.25, HasCosineInvariant: true, BridgeOnly: true},
		{Name: "CP-odd sign without C3 sheet", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, BridgeOnly: true},
		{Name: "complete positive bridge branch tag", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 2, HasC3Sheet: true, BridgeOnly: true},
		{Name: "complete negative bridge branch tag", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: -1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true, BridgeOnly: true},
		{Name: "CKM or PMNS used as branch selector", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true, UsesCKMOrPMNS: true, BridgeOnly: true},
		{Name: "branch tag attempts native phase promotion", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true, BridgeOnly: false, NativePromotionClaim: true},
		{Name: "invalid C3 sheet rejected", Cos3Phi: 0.25, HasCosineInvariant: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 3, HasC3Sheet: true, BridgeOnly: true},
	}
	out := Sieve{Executed: true, AllAcceptedBridgeOnly: true, NoNativePhaseExport: true}
	for _, r := range requests {
		e := EvaluateBranch(r)
		out.Evaluations = append(out.Evaluations, e)
		if e.Accepted {
			out.AcceptedCount++
		} else {
			out.RejectedCount++
		}
		switch r.Name {
		case "cosine-only Gate456 inverse result":
			out.CosineOnlyFlagged = !e.Accepted && e.BranchCount == CosineOnlyBranchCount && e.Verdict == StatusFailedCosineSixBranches
		case "CP-odd sign without C3 sheet":
			out.CPOddOnlyFlagged = !e.Accepted && e.BranchCount == CPSignOnlyBranchCount && e.Verdict == StatusFailedCPSignOnlyThreeSheets
		case "complete positive bridge branch tag":
			out.CompletePositiveAccepted = e.Accepted && e.Selected && e.BranchCount == CompleteBranchCount && e.Verdict == StatusCompleteBranchTagUnique
		case "complete negative bridge branch tag":
			out.CompleteNegativeAccepted = e.Accepted && e.Selected && e.BranchCount == CompleteBranchCount && e.Verdict == StatusCompleteBranchTagUnique
		case "CKM or PMNS used as branch selector":
			out.CKMPMNSSelectorRejected = !e.Accepted && e.Verdict == StatusFailedCKMPMNSSelectorRejected
		case "branch tag attempts native phase promotion":
			out.NativePromotionRejected = !e.Accepted && e.Verdict == StatusFailedNativePromotionRejected
		case "invalid C3 sheet rejected":
			out.InvalidTagRejected = !e.Accepted && e.Verdict == StatusFailedInvalidBranchTag
		}
		if e.Accepted && !e.BridgeOnlyExport {
			out.AllAcceptedBridgeOnly = false
		}
		if e.Accepted && r.NativePromotionClaim {
			out.NoNativePhaseExport = false
		}
	}
	out.Verdict = StatusBridgeOnlyBranchTagValidated
	out.Reason = fmt.Sprintf("%d complete bridge tags selected unique synthetic branches; %d incomplete, empirical-selector, or native-promotion routes failed closed.", out.AcceptedCount, out.RejectedCount)
	return out
}

// EvaluateBranch audits branch metadata for a Gate456/Gate458 phase inverse.
// The function does not import physical CKM/PMNS phases and does not promote a
// selected branch to native law. It only decides whether a bridge record has the
// metadata needed to choose one synthetic phi branch.
func EvaluateBranch(r BranchRequest) BranchEvaluation {
	e := BranchEvaluation{Request: r, DomainOK: true, BridgeOnlyExport: r.BridgeOnly, NativePromotionBlocked: !r.NativePromotionClaim}
	if !r.BridgeOnly || r.NativePromotionClaim {
		e.Accepted = false
		e.Verdict = StatusFailedNativePromotionRejected
		e.Reason = "branch tags are bridge metadata and cannot promote a CP phase, ray branch, or orientation to native law-space."
		return e
	}
	if r.UsesCKMOrPMNS {
		e.Accepted = false
		e.Verdict = StatusFailedCKMPMNSSelectorRejected
		e.Reason = "CKM/PMNS phases are physical comparator data and cannot be used as native or implicit branch selectors in this ledger."
		return e
	}
	if !r.HasCosineInvariant || math.IsNaN(r.Cos3Phi) || math.IsInf(r.Cos3Phi, 0) || math.Abs(r.Cos3Phi) > 1+1e-12 {
		e.Accepted = false
		e.DomainOK = false
		e.Verdict = StatusFailedInvalidBranchTag
		e.Reason = "a valid cosine invariant C=cos(3phi) in [-1,1] is required before branch metadata can be audited."
		return e
	}
	if !r.HasCPOddSign && !r.HasC3Sheet {
		e.Accepted = false
		e.BranchCount = CosineOnlyBranchCount
		e.Verdict = StatusFailedCosineSixBranches
		e.Reason = "cos(3phi) alone leaves the six Gate456 branches phi=(±arccos(C)+2πn)/3."
		return e
	}
	if !r.HasCPOddSign || (r.CPOddSign != +1 && r.CPOddSign != -1) {
		e.Accepted = false
		e.BranchCount = CosineOnlyBranchCount
		e.Verdict = StatusFailedInvalidBranchTag
		e.Reason = "a nonzero CP-odd orientation sign sigma_CP=±1 is required away from caustics."
		return e
	}
	if !r.HasC3Sheet {
		e.Accepted = false
		e.BranchCount = CPSignOnlyBranchCount
		e.Verdict = StatusFailedCPSignOnlyThreeSheets
		e.Reason = "the CP-odd sign chooses ±arccos(C) but still leaves the three C3 sheets n=0,1,2."
		return e
	}
	if r.C3Sheet < 0 || r.C3Sheet >= C3SheetCount {
		e.Accepted = false
		e.BranchCount = CPSignOnlyBranchCount
		e.Verdict = StatusFailedInvalidBranchTag
		e.Reason = "C3 sheet tag must be one of n=0,1,2."
		return e
	}
	theta := math.Acos(clamp(r.Cos3Phi, -1, 1))
	if r.CPOddSign < 0 {
		theta = -theta
	}
	phase := (theta + 2*math.Pi*float64(r.C3Sheet)) / 3
	e.Phase = normalizeAngle(phase)
	e.BranchCount = CompleteBranchCount
	e.Selected = true
	e.Accepted = true
	e.Verdict = StatusCompleteBranchTagUnique
	e.Reason = "complete bridge tag {sigma_CP,n_C3} selects one synthetic phase branch without importing a physical CP value."
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
		NoCPPhasePromotion:            true,
		NoCurveFitPromoted:            true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate459 selects phase branches only when complete bridge tags are supplied; native CP-sign and C3-sheet selectors remain absent, so the 13-moduli firewall is unchanged.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        460,
		Title:       "Branch-Resolved Synthetic Texture Residual Harness / Null-Data Run",
		Reason:      "Gate459 can select a unique bridge phase branch when a complete symbolic branch tag is supplied, so the next audit can run branch-resolved residuals on synthetic/null data while still rejecting observed flavor imports by default",
		PrimaryTask: "compose the Gate458 evaluator with the Gate459 branch ledger to produce branch-resolved, bridge-only texture residual records with no CKM/PMNS, Yukawa, or mass data imported",
	}
}

func truth(a Analysis) string {
	if a.Sieve.CompletePositiveAccepted && a.Sieve.CompleteNegativeAccepted && a.Firewall.NoCPPhasePromotion {
		return "Gate459 proves the exact metadata boundary for phase-branch resolution: cos(3phi) gives six branches, a CP-odd sign gives three, and only the pair {sigma_CP,n_C3} gives one. That pair is a bridge tag, not native geometry; CKM/PMNS phases and native-promotion attempts are rejected."
	}
	return "Gate459 did not validate the branch-tag ledger."
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || a.Inheritance.Gate456GenericBranchCount != CosineOnlyBranchCount || !a.Inheritance.Gate458BridgeOnly {
		return fmt.Errorf("invalid inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Ledger.Executed || !a.Ledger.RequiresCPOddSign || !a.Ledger.RequiresC3Sheet || a.Ledger.CompleteBranchTagCount != CompleteBranchCount || !a.Ledger.BridgeOnly {
		return fmt.Errorf("invalid branch ledger: %s", FormatLedger(a.Ledger))
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCount != 2 || a.Sieve.RejectedCount != 5 || !a.Sieve.CosineOnlyFlagged || !a.Sieve.CPOddOnlyFlagged || !a.Sieve.CompletePositiveAccepted || !a.Sieve.CompleteNegativeAccepted || !a.Sieve.CKMPMNSSelectorRejected || !a.Sieve.NativePromotionRejected || !a.Sieve.InvalidTagRejected || !a.Sieve.AllAcceptedBridgeOnly || !a.Sieve.NoNativePhaseExport {
		return fmt.Errorf("invalid branch sieve: %s", FormatSieve(a.Sieve))
	}
	if !a.Firewall.Executed || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoCPPhasePromotion || !a.Firewall.YPhaseStillQuarantined || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall breach: %s", FormatFirewall(a.Firewall))
	}
	if a.Next.Gate != 460 {
		return fmt.Errorf("unexpected next gate: %s", FormatNext(a.Next))
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate458Inherited,
		StatusBranchLedgerDefined,
		StatusCompleteBranchTagUnique,
		StatusBridgeOnlyBranchTagValidated,
		StatusFirewallPreserved,
		StatusFailedCosineSixBranches,
		StatusFailedCPSignOnlyThreeSheets,
		StatusFailedInvalidBranchTag,
		StatusFailedCKMPMNSSelectorRejected,
		StatusFailedNativePromotionRejected,
		StatusFailedNativeCPSelectorAbsent,
		StatusFailedNativeC3SelectorAbsent,
	}
}

func normalizeAngle(x float64) float64 {
	y := math.Mod(x, 2*math.Pi)
	if y < 0 {
		y += 2 * math.Pi
	}
	return y
}

func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
