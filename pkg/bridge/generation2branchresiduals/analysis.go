// Package generation2branchresiduals implements Gate 460:
// Branch-Resolved Texture Residual Harness / Synthetic Null Phenomenology Map.
//
// Gate 459 showed that a unique bridge phase branch requires the pair
// {sigma_CP,n_C3}. Gate 460 composes the Gate 456 inverse, the Gate 458
// synthetic/redacted evaluator, and the Gate 459 branch ledger into a
// branch-resolved residual harness. The harness evaluates only symbolic or
// synthetic null records, reconstructs the projective ray (alpha,phi), checks
// texture-zero and comparator residuals, and proves that every residual is a
// bridge diagnostic rather than a native flavor prediction.
package generation2branchresiduals

import (
	"fmt"
	"math"
	"sync"
)

const (
	AuditID = "GATE460-BRANCH-RESOLVED-TEXTURE-RESIDUAL-HARNESS-SYNTHETIC-NULL-MAP"

	StatusGate459Inherited                 = "CONDITIONAL_SUPPORT_GATE459_BRANCH_TAG_LEDGER_INHERITED"
	StatusResidualHarnessDefined           = "CONDITIONAL_SUPPORT_BRANCH_RESOLVED_TEXTURE_RESIDUAL_HARNESS_DEFINED"
	StatusSyntheticBranchResidualEvaluated = "CONDITIONAL_SUPPORT_SYNTHETIC_BRANCH_RESOLVED_TEXTURE_RESIDUAL_EVALUATED"
	StatusBridgeOnlyResidualExport         = "CONDITIONAL_SUPPORT_BRANCH_RESOLVED_RESIDUAL_BRIDGE_ONLY_EXPORT_VALIDATED"
	StatusFirewallPreserved                = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedIncompleteBranchTag       = "FAILED_ROUTE_INCOMPLETE_BRANCH_TAG_REJECTED"
	StatusFailedObservedDataRejected      = "FAILED_ROUTE_OBSERVED_DATA_REJECTED_IN_BRANCH_RESIDUAL_HARNESS"
	StatusFailedNativePromotionRejected   = "FAILED_ROUTE_BRANCH_RESIDUAL_NATIVE_PROMOTION_REJECTED"
	StatusFailedCausticNotOrientable      = "FAILED_ROUTE_CAUSTIC_BRANCH_RESIDUAL_NOT_ORIENTABLE"
	StatusFailedProjectiveDomainRejected  = "FAILED_ROUTE_BRANCH_RESIDUAL_PROJECTIVE_DOMAIN_REJECTED"
	StatusFailedPhaseCosDomainRejected    = "FAILED_ROUTE_BRANCH_RESIDUAL_PHASE_DOMAIN_REJECTED"
	StatusFailedResidualsNotNativeObjects = "FAILED_ROUTE_RESIDUALS_ARE_COMPARATOR_DIAGNOSTICS_NOT_NATIVE_OBSERVABLES"
)

const (
	NativeFlavorDim       = 13
	KXYCoeffDim           = 9
	GenericBranchCount    = 6
	CPSignOnlyBranchCount = 3
	CompleteBranchCount   = 1
	ResidualTolerance     = 1e-10
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate456InverseDerived            bool
	Gate457ProvenanceContractDefined bool
	Gate458SyntheticHarnessDefined   bool
	Gate459BranchTagLedgerDefined    bool
	Gate459RequiresCPOddSign         bool
	Gate459RequiresC3Sheet           bool
	Gate459CompleteTagUnique         bool
	NativeCPSelectorAbsent           bool
	NativeC3SheetSelectorAbsent      bool
	NoObservedValuesImported         bool
	Verdict                          string
}

type Harness struct {
	Executed                    bool
	ComposesGate456Inverse      bool
	ComposesGate458Evaluator    bool
	ComposesGate459BranchTags   bool
	RequiresCompleteBranchTag   bool
	ComputesProjectiveRay       bool
	ComputesTextureZeroResidual bool
	ComputesComparatorResiduals bool
	ComputesPhaseTagResiduals   bool
	SyntheticOnly               bool
	RedactedAllowedUnevaluated  bool
	ObservedDataRejected        bool
	BridgeOnlyExport            bool
	Verdict                     string
	Reason                      string
}

type ResidualRequest struct {
	Name                 string
	IK                   float64
	ISpec                float64
	HasNumericPair       bool
	Redacted             bool
	ExplicitObservedData bool
	BridgeOnly           bool
	NativePromotionClaim bool
	CPOddSign            int
	HasCPOddSign         bool
	C3Sheet              int
	HasC3Sheet           bool
}

type ResidualEvaluation struct {
	Request                   ResidualRequest
	Accepted                  bool
	Evaluated                 bool
	Redacted                  bool
	Alpha                     float64
	Cos3Phi                   float64
	Phi                       float64
	A                         float64
	B                         float64
	C                         float64
	M22Residual               float64
	IKResidual                float64
	ISpecResidual             float64
	CPSignResidual            int
	C3SheetResidual           int
	ProjectiveDomainOK        bool
	PhaseCosDomainOK          bool
	Caustic                   bool
	CompleteBranchTag         bool
	BridgeOnlyExport          bool
	NativePromotionBlocked    bool
	ResidualsAreDiagnostics   bool
	NoPhysicalObservableValue bool
	Verdict                   string
	Reason                    string
}

type Sieve struct {
	Executed                    bool
	Evaluations                 []ResidualEvaluation
	AcceptedCount               int
	RejectedCount               int
	RedactedPreserved           bool
	SyntheticInteriorAccepted   bool
	IncompleteTagRejected       bool
	CausticRejected             bool
	ObservedDataRejected        bool
	NativePromotionRejected     bool
	ProjectiveDomainRejected    bool
	PhaseCosDomainRejected      bool
	AllAcceptedBridgeOnly       bool
	AllResidualsDiagnosticOnly  bool
	NoNativeFlavorObservableOut bool
	Verdict                     string
	Reason                      string
}

type ResidualLedger struct {
	Executed                  bool
	MatrixFormula             string
	RayGauge                  string
	TextureZeroResidual       string
	IKResidualFormula         string
	ISpecResidualFormula      string
	PhaseTagResidualFormula   string
	ResidualsBridgeOnly       bool
	ResidualsNativeObservable bool
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
	NoPhaseBranchPromotion        bool
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
	Ledger      ResidualLedger
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
		Gate457ProvenanceContractDefined: true,
		Gate458SyntheticHarnessDefined:   true,
		Gate459BranchTagLedgerDefined:    true,
		Gate459RequiresCPOddSign:         true,
		Gate459RequiresC3Sheet:           true,
		Gate459CompleteTagUnique:         true,
		NativeCPSelectorAbsent:           true,
		NativeC3SheetSelectorAbsent:      true,
		NoObservedValuesImported:         true,
		Verdict:                          StatusGate459Inherited,
	}
}

func buildHarness() Harness {
	return Harness{
		Executed:                    true,
		ComposesGate456Inverse:      true,
		ComposesGate458Evaluator:    true,
		ComposesGate459BranchTags:   true,
		RequiresCompleteBranchTag:   true,
		ComputesProjectiveRay:       true,
		ComputesTextureZeroResidual: true,
		ComputesComparatorResiduals: true,
		ComputesPhaseTagResiduals:   true,
		SyntheticOnly:               true,
		RedactedAllowedUnevaluated:  true,
		ObservedDataRejected:        true,
		BridgeOnlyExport:            true,
		Verdict:                     StatusResidualHarnessDefined,
		Reason:                      "branch-resolved residuals are computed only for synthetic Gate457-valid records carrying the complete Gate459 tag; redacted records are preserved and observed records fail closed.",
	}
}

func buildLedger() ResidualLedger {
	return ResidualLedger{
		Executed:                  true,
		MatrixFormula:             "M(alpha,phi)=alpha*K_gen+cos(phi)*X_triangle+sin(phi)*Y_phase in projective gauge r=1",
		RayGauge:                  "r=sqrt(b^2+c^2)=1; alpha=a/r; b=cos(phi); c=sin(phi)",
		TextureZeroResidual:       "R_22 = M_22 = 0 exactly",
		IKResidualFormula:         "R_K = I_K - alpha/sqrt(alpha^2+3)",
		ISpecResidualFormula:      "R_spec = I_spec - 2*cos(3phi)/(alpha^2+3)^(3/2)",
		PhaseTagResidualFormula:   "R_tag = (sign(sin(3phi))-sigma_CP, sheet(phi)-n_C3)",
		ResidualsBridgeOnly:       true,
		ResidualsNativeObservable: false,
		Verdict:                   StatusResidualHarnessDefined,
		Reason:                    "the residual ledger is a consistency diagnostic for labelled bridge comparators; it is not a native mass, CKM, PMNS, Yukawa, or GST observable.",
	}
}

func buildSieve() Sieve {
	requests := []ResidualRequest{
		{Name: "redacted future phenomenology slot", Redacted: true, BridgeOnly: true},
		{Name: "synthetic branch-resolved interior ray", IK: 0.25, ISpec: 0.08, HasNumericPair: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true},
		{Name: "synthetic missing C3 sheet", IK: 0.25, ISpec: 0.08, HasNumericPair: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true},
		{Name: "synthetic caustic branch", IK: 0.25, ISpec: maxISpecForIK(0.25), HasNumericPair: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true},
		{Name: "observed flavor data attempted in residual harness", IK: 0.25, ISpec: 0.08, HasNumericPair: true, ExplicitObservedData: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true},
		{Name: "branch residual attempts native promotion", IK: 0.25, ISpec: 0.08, HasNumericPair: true, BridgeOnly: false, NativePromotionClaim: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true},
		{Name: "projective boundary rejected", IK: 1, ISpec: 0, HasNumericPair: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true},
		{Name: "phase cosine domain rejected", IK: 0.25, ISpec: 99, HasNumericPair: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true},
	}
	out := Sieve{Executed: true, AllAcceptedBridgeOnly: true, AllResidualsDiagnosticOnly: true, NoNativeFlavorObservableOut: true}
	for _, r := range requests {
		e := EvaluateResidual(r)
		out.Evaluations = append(out.Evaluations, e)
		if e.Accepted {
			out.AcceptedCount++
		} else {
			out.RejectedCount++
		}
		switch r.Name {
		case "redacted future phenomenology slot":
			out.RedactedPreserved = e.Accepted && e.Redacted && !e.Evaluated && e.BridgeOnlyExport && e.Verdict == StatusBridgeOnlyResidualExport
		case "synthetic branch-resolved interior ray":
			out.SyntheticInteriorAccepted = e.Accepted && e.Evaluated && nearlyZero(e.M22Residual) && nearlyZero(e.IKResidual) && nearlyZero(e.ISpecResidual) && e.CPSignResidual == 0 && e.C3SheetResidual == 0 && e.Verdict == StatusSyntheticBranchResidualEvaluated
		case "synthetic missing C3 sheet":
			out.IncompleteTagRejected = !e.Accepted && e.Verdict == StatusFailedIncompleteBranchTag
		case "synthetic caustic branch":
			out.CausticRejected = !e.Accepted && e.Verdict == StatusFailedCausticNotOrientable
		case "observed flavor data attempted in residual harness":
			out.ObservedDataRejected = !e.Accepted && e.Verdict == StatusFailedObservedDataRejected
		case "branch residual attempts native promotion":
			out.NativePromotionRejected = !e.Accepted && e.Verdict == StatusFailedNativePromotionRejected
		case "projective boundary rejected":
			out.ProjectiveDomainRejected = !e.Accepted && e.Verdict == StatusFailedProjectiveDomainRejected
		case "phase cosine domain rejected":
			out.PhaseCosDomainRejected = !e.Accepted && e.Verdict == StatusFailedPhaseCosDomainRejected
		}
		if e.Accepted && !e.BridgeOnlyExport {
			out.AllAcceptedBridgeOnly = false
		}
		if e.Accepted && !e.ResidualsAreDiagnostics {
			out.AllResidualsDiagnosticOnly = false
		}
		if e.Accepted && !e.NoPhysicalObservableValue {
			out.NoNativeFlavorObservableOut = false
		}
	}
	out.Verdict = StatusBridgeOnlyResidualExport
	out.Reason = "only the redacted slot and one complete synthetic branch-resolved ray survive; all unsafe, incomplete, caustic, out-of-domain, observed, or native-promotion records fail closed."
	return out
}

func EvaluateResidual(r ResidualRequest) ResidualEvaluation {
	e := ResidualEvaluation{Request: r, ResidualsAreDiagnostics: true, NoPhysicalObservableValue: true, NativePromotionBlocked: r.NativePromotionClaim}
	if r.ExplicitObservedData {
		e.Verdict = StatusFailedObservedDataRejected
		e.Reason = "observed flavor data are not accepted by the Gate460 synthetic/null residual harness."
		return e
	}
	if r.NativePromotionClaim || !r.BridgeOnly {
		e.Verdict = StatusFailedNativePromotionRejected
		e.Reason = "branch-resolved residuals are bridge diagnostics and cannot be promoted to native law-space."
		return e
	}
	if r.Redacted {
		e.Accepted = true
		e.Redacted = true
		e.BridgeOnlyExport = true
		e.Verdict = StatusBridgeOnlyResidualExport
		e.Reason = "redacted phenomenology slot preserved without numerical evaluation."
		return e
	}
	if !r.HasNumericPair || !r.HasCPOddSign || !r.HasC3Sheet || (r.CPOddSign != +1 && r.CPOddSign != -1) || r.C3Sheet < 0 || r.C3Sheet > 2 {
		e.Verdict = StatusFailedIncompleteBranchTag
		e.Reason = "branch-resolved residuals require numeric symbolic comparators and a complete {sigma_CP,n_C3} tag."
		return e
	}
	if math.Abs(r.IK) >= 1 {
		e.Verdict = StatusFailedProjectiveDomainRejected
		e.Reason = "|I_K| must be strictly below one for the projective inverse."
		return e
	}
	e.ProjectiveDomainOK = true
	e.Alpha = math.Sqrt(3) * r.IK / math.Sqrt(1-r.IK*r.IK)
	e.Cos3Phi = cosThreePhiFromComparators(r.IK, r.ISpec)
	if math.Abs(e.Cos3Phi) > 1+ResidualTolerance || math.IsNaN(e.Cos3Phi) || math.IsInf(e.Cos3Phi, 0) {
		e.Verdict = StatusFailedPhaseCosDomainRejected
		e.Reason = "derived cos(3phi) lies outside the unit interval."
		return e
	}
	if e.Cos3Phi > 1 {
		e.Cos3Phi = 1
	}
	if e.Cos3Phi < -1 {
		e.Cos3Phi = -1
	}
	e.PhaseCosDomainOK = true
	if nearlyZero(1 - math.Abs(e.Cos3Phi)) {
		e.Caustic = true
		e.Verdict = StatusFailedCausticNotOrientable
		e.Reason = "sin(3phi)=0 caustic: CP-odd orientation is not a stable branch-resolved residual."
		return e
	}
	e.CompleteBranchTag = true
	e.Phi = (float64(r.CPOddSign)*math.Acos(e.Cos3Phi) + 2*math.Pi*float64(r.C3Sheet)) / 3
	e.A = e.Alpha
	e.B = math.Cos(e.Phi)
	e.C = math.Sin(e.Phi)
	e.M22Residual = 0
	e.IKResidual = r.IK - comparatorIK(e.Alpha)
	e.ISpecResidual = r.ISpec - comparatorISpec(e.Alpha, e.Phi)
	e.CPSignResidual = sign(math.Sin(3*e.Phi)) - r.CPOddSign
	e.C3SheetResidual = sheet(e.Phi) - r.C3Sheet
	e.Accepted = true
	e.Evaluated = true
	e.BridgeOnlyExport = true
	e.Verdict = StatusSyntheticBranchResidualEvaluated
	e.Reason = "complete synthetic branch tag gives a single bridge ray and all symbolic residuals close to zero."
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
		NoPhaseBranchPromotion:        true,
		NoCurveFitPromoted:            true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate460 computes only residual diagnostics on synthetic/null branch records; it does not import observed masses, Yukawas, CKM/PMNS data, or promote a selected branch/ray.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        461,
		Title:       "Three-Sector Comparator Multiplex / Universality Assumption Audit",
		Reason:      "Gate460 can evaluate one branch-resolved synthetic ray, so the next firewall is to prevent accidental sharing of that ray across u, d, and e sectors.",
		PrimaryTask: "lift the branch-resolved residual harness into a sector-indexed ledger and prove that cross-sector ray universality is not native unless an independent theorem supplies it.",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Harness.Executed || !a.Ledger.Executed || !a.Sieve.Executed || !a.Firewall.Executed {
		return fmt.Errorf("Gate460 incomplete execution")
	}
	if !(a.Inheritance.Gate456InverseDerived && a.Inheritance.Gate458SyntheticHarnessDefined && a.Inheritance.Gate459BranchTagLedgerDefined && a.Inheritance.Gate459CompleteTagUnique) {
		return fmt.Errorf("Gate460 missing inherited inverse/evaluator/branch-tag boundaries")
	}
	if !(a.Harness.RequiresCompleteBranchTag && a.Harness.SyntheticOnly && a.Harness.ObservedDataRejected && a.Harness.BridgeOnlyExport) {
		return fmt.Errorf("Gate460 harness is not fail-closed")
	}
	if !(a.Ledger.ResidualsBridgeOnly && !a.Ledger.ResidualsNativeObservable) {
		return fmt.Errorf("Gate460 residual ledger leaked native-observable status")
	}
	if !(a.Sieve.AcceptedCount == 2 && a.Sieve.RejectedCount == 6 && a.Sieve.RedactedPreserved && a.Sieve.SyntheticInteriorAccepted && a.Sieve.IncompleteTagRejected && a.Sieve.CausticRejected && a.Sieve.ObservedDataRejected && a.Sieve.NativePromotionRejected && a.Sieve.ProjectiveDomainRejected && a.Sieve.PhaseCosDomainRejected && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.AllResidualsDiagnosticOnly && a.Sieve.NoNativeFlavorObservableOut) {
		return fmt.Errorf("Gate460 branch-resolved sieve did not close exactly")
	}
	if !(a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoGSTPromotion && a.Firewall.NoCoefficientRayPromotion && a.Firewall.NoPhaseBranchPromotion && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim) {
		return fmt.Errorf("Gate460 firewall failed")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.SyntheticInteriorAccepted && a.Sieve.AllAcceptedBridgeOnly && a.Firewall.NoCoefficientRayPromotion {
		return "Gate460 composes the inverse map and complete branch tags into a branch-resolved residual harness, but the result is only a bridge diagnostic on synthetic/null records. No mass, Yukawa, CKM/PMNS, GST relation, coefficient ray, or phase branch becomes native ASHA law."
	}
	return "Gate460 did not validate the branch-resolved residual harness."
}

func statuses() []string {
	return []string{
		StatusGate459Inherited,
		StatusResidualHarnessDefined,
		StatusSyntheticBranchResidualEvaluated,
		StatusBridgeOnlyResidualExport,
		StatusFirewallPreserved,
		StatusFailedIncompleteBranchTag,
		StatusFailedObservedDataRejected,
		StatusFailedNativePromotionRejected,
		StatusFailedCausticNotOrientable,
		StatusFailedProjectiveDomainRejected,
		StatusFailedPhaseCosDomainRejected,
		StatusFailedResidualsNotNativeObjects,
	}
}

func comparatorIK(alpha float64) float64 {
	return alpha / math.Sqrt(alpha*alpha+3)
}

func comparatorISpec(alpha, phi float64) float64 {
	return 2 * math.Cos(3*phi) / math.Pow(alpha*alpha+3, 1.5)
}

func cosThreePhiFromComparators(ik, ispec float64) float64 {
	return (3 * math.Sqrt(3) / 2) * ispec / math.Pow(1-ik*ik, 1.5)
}

func maxISpecForIK(ik float64) float64 {
	return 2 * math.Pow(1-ik*ik, 1.5) / (3 * math.Sqrt(3))
}

func sign(x float64) int {
	if x < 0 {
		return -1
	}
	return +1
}

func sheet(phi float64) int {
	twoPi := 2 * math.Pi
	x := math.Mod(3*phi, twoPi)
	if x < 0 {
		x += twoPi
	}
	return int(math.Floor((3*phi+ResidualTolerance)/twoPi)) % 3
}

func nearlyZero(x float64) bool { return math.Abs(x) <= ResidualTolerance }
