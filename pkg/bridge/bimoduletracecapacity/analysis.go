// Package bimoduletracecapacity implements Gate 290:
// Bimodule Trace Capacity Sieve / Sector Hierarchy Audit.
//
// Gate 289 exposed branch sensitivity in lepton/quark sector-projected traces
// but found no native rule selecting r_+ or r_-.  Gate 290 stress-tests the
// proposed Morita trace-capacity principle: since κ_C:κ_Q=1:3, can the finite
// Hilbert bimodule demand that quark-sector spectral weight dominate lepton
// spectral weight and thereby veto the lower branch?  The audit separates a
// weak total-capacity inequality, which both branches pass, from a stronger
// per-slot amplitude monotonicity inequality, which would select r_+ but is not
// derived from multiplicity alone.
package bimoduletracecapacity

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE290-BIMODULE-TRACE-CAPACITY-SIEVE-SECTOR-HIERARCHY-AUDIT"

	StatusGate289Inherited      = "CONDITIONAL_SUPPORT_GATE289_SECTOR_TRACE_DIAGNOSTIC_INHERITED"
	StatusCapacityFormalized    = "CONDITIONAL_SUPPORT_TRACE_CAPACITY_CANDIDATES_FORMALIZED"
	StatusBranchStressCompleted = "CONDITIONAL_SUPPORT_BRANCH_STRESS_TEST_COMPLETED"
	StatusWeakBoundLedger       = "CONDITIONAL_SUPPORT_TOTAL_CAPACITY_BOUND_AUDITED"
	StatusStrongBoundDiagnostic = "CONDITIONAL_SUPPORT_PER_SLOT_MONOTONIC_BOUND_DIAGNOSTIC_EXPOSED"
	StatusFirewallPreserved     = "CONDITIONAL_SUPPORT_TRACE_CAPACITY_FIREWALLS_PRESERVED"

	StatusFailedNoNativeCapacityLaw  = "FAILED_ROUTE_NO_NATIVE_TRACE_CAPACITY_BOUND_DERIVED"
	StatusFailedWeakBoundBothPass    = "FAILED_ROUTE_TOTAL_CAPACITY_BOUND_DOES_NOT_SELECT_BRANCH"
	StatusFailedStrongBoundUnsealed  = "FAILED_ROUTE_PER_SLOT_MONOTONIC_BOUND_IS_EXTRA_SELECTION_AXIOM"
	StatusFailedCannotVetoRMinus     = "FAILED_ROUTE_GEOMETRY_CANNOT_STRICTLY_VETO_R_MINUS_DISTRIBUTION"
	StatusFailedBranchNotSelected    = "FAILED_ROUTE_BRANCH_NOT_SELECTED_BY_TRACE_CAPACITY_BOUND"
	StatusFailedHiggsRatioNotDerived = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

const (
	// Gate-273 Morita multiplicities.
	kappaC = 1.0
	kappaQ = 3.0

	// Gate-162 contact cutoff and topological action snapshots used by Gate 288.
	sTop         = 8 * math.Pi * math.Pi
	contactZeta0 = 7.0
	contactZeta2 = 61.0 / 25.0
	contactZeta4 = 257629.0 / 202500.0
	reducedA0    = kappaC + kappaQ

	// Gate-169 scalar/contact shape.
	lambdaNum = 1197
	lambdaDen = 4624
)

type Gate289Inheritance struct {
	SectorProjectedTracesBranchSensitive bool
	GammaTracesBranchBlind               bool
	PhysicalJDerived                     bool
	SelectionPrincipleDerived            bool
	BranchesSurvived                     []string
	Verdict                              string
}

type Branch struct {
	Name      string
	Exact     string
	R         float64
	X         float64
	AbsYOverX float64
	D2Total   float64
	D4Total   float64
	Shape     float64
	LeptonD2  float64
	QuarkD2   float64
	LeptonD4  float64
	QuarkD4   float64
	PerC_D2   float64
	PerQ_D2   float64
	PerC_D4   float64
	PerQ_D4   float64
}

type CapacityCandidate struct {
	Name          string
	Formula       string
	DerivedNative bool
	Description   string
	Verdict       string
}

type CapacityFormalization struct {
	KappaC             float64
	KappaQ             float64
	MultiplicityRatio  string
	CandidateBounds    []CapacityCandidate
	NativeBoundDerived bool
	Verdict            string
}

type BranchCapacityResult struct {
	BranchName             string
	R                      float64
	X                      float64
	LeptonD2               float64
	QuarkD2                float64
	LeptonD4               float64
	QuarkD4                float64
	PerSlotLeptonD2        float64
	PerSlotQuarkD2         float64
	PerSlotLeptonD4        float64
	PerSlotQuarkD4         float64
	TotalD2CapacityPass    bool
	TotalD4CapacityPass    bool
	PerSlotD2MonotonicPass bool
	PerSlotD4MonotonicPass bool
	LeptonFourthFraction   float64
	QuarkFourthFraction    float64
	Verdict                string
}

type BranchStressTest struct {
	Results               []BranchCapacityResult
	BothPassTotalCapacity bool
	ExactlyOnePassPerSlot bool
	PerSlotSelectedBranch string
	Verdict               string
}

type GeometricVetoAudit struct {
	TotalCapacityIsNativeDiagnostic bool
	TotalCapacityCanVetoRMinus      bool
	PerSlotCapacityWouldVetoRMinus  bool
	PerSlotCapacityIsNativeTheorem  bool
	RMinusViolatesDerivedGeometry   bool
	Reason                          string
	Verdict                         string
}

type BranchSelection struct {
	BranchesEvaluated    []string
	EliminatedByNative   []string
	EliminatedByUnsealed []string
	SurvivingBranches    []string
	UniqueBranchSelected bool
	SelectedBranch       string
	Verdict              string
}

type HiggsFirewall struct {
	BranchSelected                  bool
	HeatKernelProjectionDerived     bool
	ScalarGaugeNormalizationDerived bool
	HiggsPredictionClaimed          bool
	Verdict                         string
}

type Firewalls struct {
	DoesNotPromoteMultiplicityToAmplitude bool
	DoesNotUsePerSlotBoundAsTheorem       bool
	DoesNotVetoRMinusWithoutSelector      bool
	DoesNotClaimBranchSelection           bool
	DoesNotClaimHiggsPrediction           bool
	FiniteCorePolluted                    bool
	Verdict                               string
}

type Summary struct {
	Gate289Inherited       bool
	CapacityFormalized     bool
	BranchStressCompleted  bool
	TotalCapacitySelects   bool
	PerSlotDiagnostic      bool
	NativeVetoDerived      bool
	UniqueBranchSelected   bool
	HiggsPredictionDerived bool
	FirewallPreserved      bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	Inheritance Gate289Inheritance
	Branches    []Branch
	Capacity    CapacityFormalization
	Stress      BranchStressTest
	Veto        GeometricVetoAudit
	Selection   BranchSelection
	Higgs       HiggsFirewall
	Firewalls   Firewalls
	Summary     Summary
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inh := inheritGate289()
	branches, err := buildBranches()
	if err != nil {
		return Analysis{}, err
	}
	cap := formalizeCapacity()
	stress := stressTestBranches(branches)
	veto := auditVeto(stress)
	sel := selectBranch(branches, veto)
	higgs := auditHiggs(sel)
	fw := auditFirewalls(cap, veto, sel, higgs)
	summary := buildSummary(inh, cap, stress, veto, sel, higgs, fw)
	truth := "Gate 290 shows that Morita multiplicity provides a trace-capacity diagnostic, not an amplitude-selection theorem. The weak total-capacity bound Tr(P_QD^n)>=Tr(P_CD^n) is satisfied by both r branches. A stronger per-slot monotonicity rule would select r_+, but that rule is not derived from κ_C:κ_Q=1:3 and would be an extra dynamical axiom. Therefore r_- cannot be strictly vetoed and the Higgs branch remains unselected."
	return Analysis{Inheritance: inh, Branches: branches, Capacity: cap, Stress: stress, Veto: veto, Selection: sel, Higgs: higgs, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate289() Gate289Inheritance {
	return Gate289Inheritance{SectorProjectedTracesBranchSensitive: true, GammaTracesBranchBlind: true, PhysicalJDerived: false, SelectionPrincipleDerived: false, BranchesSurvived: []string{"r_plus", "r_minus"}, Verdict: StatusGate289Inherited}
}

func buildBranches() ([]Branch, error) {
	defs := []struct {
		name, exact string
		r           float64
	}{
		{"r_plus", "(3591 + 136√123)/3099", (3591 + 136*math.Sqrt(123)) / 3099},
		{"r_minus", "(3591 - 136√123)/3099", (3591 - 136*math.Sqrt(123)) / 3099},
	}
	out := make([]Branch, 0, len(defs))
	for _, d := range defs {
		x, err := solvePositiveX(d.r)
		if err != nil {
			return nil, err
		}
		l2 := x
		q2 := kappaQ * x * d.r
		l4 := x * x
		q4 := kappaQ * x * x * d.r * d.r
		d2 := l2 + q2
		d4 := l4 + q4
		out = append(out, Branch{Name: d.name, Exact: d.exact, R: d.r, X: x, AbsYOverX: math.Sqrt(d.r), D2Total: d2, D4Total: d4, Shape: d4 / (d2 * d2), LeptonD2: l2, QuarkD2: q2, LeptonD4: l4, QuarkD4: q4, PerC_D2: l2 / kappaC, PerQ_D2: q2 / kappaQ, PerC_D4: l4 / kappaC, PerQ_D4: q4 / kappaQ})
	}
	return out, nil
}

func solvePositiveX(r float64) (float64, error) {
	A := contactZeta0 * (1 + 3*r*r)
	B := contactZeta2 * (1 + 3*r)
	C := contactZeta4*reducedA0 - sTop
	disc := B*B - 4*A*C
	if disc < 0 {
		return math.NaN(), fmt.Errorf("negative discriminant for r=%g", r)
	}
	roots := []float64{(-B + math.Sqrt(disc)) / (2 * A), (-B - math.Sqrt(disc)) / (2 * A)}
	for _, root := range roots {
		if root > 0 && !math.IsInf(root, 0) && !math.IsNaN(root) {
			return root, nil
		}
	}
	return math.NaN(), fmt.Errorf("no positive root for r=%g", r)
}

func formalizeCapacity() CapacityFormalization {
	candidates := []CapacityCandidate{
		{Name: "weak_total_sector_capacity", Formula: "Tr(P_QD_F^{2n}) >= Tr(P_CD_F^{2n})", DerivedNative: false, Description: "A diagnostic suggested by κ_Q>κ_C. It compares total sector spectral weight, not per-state edge amplitude.", Verdict: StatusWeakBoundLedger},
		{Name: "strong_per_slot_monotonicity", Formula: "Tr(P_QD_F^{2n})/κ_Q >= Tr(P_CD_F^{2n})/κ_C", DerivedNative: false, Description: "Equivalent in this reduced ledger to r^n>=1. It would select r_+, but it is an additional amplitude-ordering axiom, not a consequence of multiplicity.", Verdict: StatusStrongBoundDiagnostic},
	}
	return CapacityFormalization{KappaC: kappaC, KappaQ: kappaQ, MultiplicityRatio: "1:3", CandidateBounds: candidates, NativeBoundDerived: false, Verdict: StatusCapacityFormalized}
}

func stressTestBranches(branches []Branch) BranchStressTest {
	results := make([]BranchCapacityResult, 0, len(branches))
	bothTotal := true
	perSlotPass := []string{}
	for _, b := range branches {
		totalD2Pass := b.QuarkD2+1e-12 >= b.LeptonD2
		totalD4Pass := b.QuarkD4+1e-12 >= b.LeptonD4
		perD2Pass := b.PerQ_D2+1e-12 >= b.PerC_D2
		perD4Pass := b.PerQ_D4+1e-12 >= b.PerC_D4
		if !totalD2Pass || !totalD4Pass {
			bothTotal = false
		}
		if perD2Pass && perD4Pass {
			perSlotPass = append(perSlotPass, b.Name)
		}
		total4 := b.LeptonD4 + b.QuarkD4
		results = append(results, BranchCapacityResult{BranchName: b.Name, R: b.R, X: b.X, LeptonD2: b.LeptonD2, QuarkD2: b.QuarkD2, LeptonD4: b.LeptonD4, QuarkD4: b.QuarkD4, PerSlotLeptonD2: b.PerC_D2, PerSlotQuarkD2: b.PerQ_D2, PerSlotLeptonD4: b.PerC_D4, PerSlotQuarkD4: b.PerQ_D4, TotalD2CapacityPass: totalD2Pass, TotalD4CapacityPass: totalD4Pass, PerSlotD2MonotonicPass: perD2Pass, PerSlotD4MonotonicPass: perD4Pass, LeptonFourthFraction: b.LeptonD4 / total4, QuarkFourthFraction: b.QuarkD4 / total4, Verdict: StatusBranchStressCompleted})
	}
	selected := ""
	if len(perSlotPass) == 1 {
		selected = perSlotPass[0]
	}
	return BranchStressTest{Results: results, BothPassTotalCapacity: bothTotal, ExactlyOnePassPerSlot: len(perSlotPass) == 1, PerSlotSelectedBranch: selected, Verdict: StatusBranchStressCompleted}
}

func auditVeto(stress BranchStressTest) GeometricVetoAudit {
	reason := "The derived Morita data κ_C:κ_Q=1:3 count trace multiplicities. They license total sector diagnostics but do not impose a theorem that per-slot quark edge norms must exceed lepton edge norms. The weak total-capacity inequality is too weak because both branches satisfy it; the strong per-slot inequality would veto r_- but is precisely an extra amplitude-selection law."
	return GeometricVetoAudit{TotalCapacityIsNativeDiagnostic: true, TotalCapacityCanVetoRMinus: false, PerSlotCapacityWouldVetoRMinus: stress.ExactlyOnePassPerSlot && stress.PerSlotSelectedBranch == "r_plus", PerSlotCapacityIsNativeTheorem: false, RMinusViolatesDerivedGeometry: false, Reason: reason, Verdict: StatusFailedCannotVetoRMinus}
}

func selectBranch(branches []Branch, veto GeometricVetoAudit) BranchSelection {
	evaluated := make([]string, 0, len(branches))
	for _, b := range branches {
		evaluated = append(evaluated, b.Name)
	}
	survivors := append([]string{}, evaluated...)
	eliminatedNative := []string{}
	eliminatedUnsealed := []string{}
	if veto.PerSlotCapacityWouldVetoRMinus && !veto.PerSlotCapacityIsNativeTheorem {
		eliminatedUnsealed = append(eliminatedUnsealed, "r_minus_if_strong_per_slot_axiom_were_sealed")
	}
	return BranchSelection{BranchesEvaluated: evaluated, EliminatedByNative: eliminatedNative, EliminatedByUnsealed: eliminatedUnsealed, SurvivingBranches: survivors, UniqueBranchSelected: false, SelectedBranch: "", Verdict: StatusFailedBranchNotSelected}
}

func auditHiggs(sel BranchSelection) HiggsFirewall {
	return HiggsFirewall{BranchSelected: sel.UniqueBranchSelected, HeatKernelProjectionDerived: false, ScalarGaugeNormalizationDerived: false, HiggsPredictionClaimed: false, Verdict: StatusFailedHiggsRatioNotDerived}
}

func auditFirewalls(cap CapacityFormalization, veto GeometricVetoAudit, sel BranchSelection, h HiggsFirewall) Firewalls {
	return Firewalls{DoesNotPromoteMultiplicityToAmplitude: !cap.NativeBoundDerived, DoesNotUsePerSlotBoundAsTheorem: !veto.PerSlotCapacityIsNativeTheorem, DoesNotVetoRMinusWithoutSelector: !veto.RMinusViolatesDerivedGeometry, DoesNotClaimBranchSelection: !sel.UniqueBranchSelected, DoesNotClaimHiggsPrediction: !h.HiggsPredictionClaimed, FiniteCorePolluted: false, Verdict: StatusFirewallPreserved}
}

func buildSummary(inh Gate289Inheritance, cap CapacityFormalization, stress BranchStressTest, veto GeometricVetoAudit, sel BranchSelection, h HiggsFirewall, fw Firewalls) Summary {
	statuses := []string{StatusGate289Inherited, StatusCapacityFormalized, StatusBranchStressCompleted, StatusWeakBoundLedger, StatusStrongBoundDiagnostic, StatusFirewallPreserved, StatusFailedNoNativeCapacityLaw, StatusFailedWeakBoundBothPass, StatusFailedStrongBoundUnsealed, StatusFailedCannotVetoRMinus, StatusFailedBranchNotSelected, StatusFailedHiggsRatioNotDerived}
	return Summary{Gate289Inherited: inh.SectorProjectedTracesBranchSensitive && len(inh.BranchesSurvived) == 2, CapacityFormalized: len(cap.CandidateBounds) == 2, BranchStressCompleted: len(stress.Results) == 2, TotalCapacitySelects: !stress.BothPassTotalCapacity, PerSlotDiagnostic: stress.ExactlyOnePassPerSlot, NativeVetoDerived: veto.RMinusViolatesDerivedGeometry, UniqueBranchSelected: sel.UniqueBranchSelected, HiggsPredictionDerived: h.HiggsPredictionClaimed, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotClaimBranchSelection && fw.DoesNotClaimHiggsPrediction, Status: strings.Join(statuses, ";"), DirectAnswer: "No. The native 1⊕3 Morita multiplicities do not mathematically prefer one branch. A weak total-capacity bound is satisfied by both r_+ and r_-. A stronger per-slot monotonic bound would select r_+, but it is not derived from multiplicity and would need a new theorem or seal.", NextGate: "A future gate must derive a genuine sector-capacity functional, physical J/hypercharge representation, or an explicit sealed amplitude-ordering principle before discarding r_-."}
}

func FormatBranch(b Branch) string {
	return fmt.Sprintf("%s r=%.12g X=%.12g D2(C,Q,total)=(%.12g,%.12g,%.12g) D4(C,Q,total)=(%.12g,%.12g,%.12g) perSlotD4(C,Q)=(%.12g,%.12g)", b.Name, b.R, b.X, b.LeptonD2, b.QuarkD2, b.D2Total, b.LeptonD4, b.QuarkD4, b.D4Total, b.PerC_D4, b.PerQ_D4)
}

func FormatCapacity(c CapacityFormalization) string {
	parts := []string{}
	for _, b := range c.CandidateBounds {
		parts = append(parts, fmt.Sprintf("%s formula=%q native=%t verdict=%s", b.Name, b.Formula, b.DerivedNative, b.Verdict))
	}
	return fmt.Sprintf("κC=%.0f κQ=%.0f ratio=%s nativeBound=%t candidates=[%s] verdict=%s", c.KappaC, c.KappaQ, c.MultiplicityRatio, c.NativeBoundDerived, strings.Join(parts, "; "), c.Verdict)
}

func FormatStress(s BranchStressTest) string {
	parts := []string{}
	for _, r := range s.Results {
		parts = append(parts, fmt.Sprintf("%s totalPass(D2,D4)=(%t,%t) perSlotPass(D2,D4)=(%t,%t) C4frac=%.12g Q4frac=%.12g", r.BranchName, r.TotalD2CapacityPass, r.TotalD4CapacityPass, r.PerSlotD2MonotonicPass, r.PerSlotD4MonotonicPass, r.LeptonFourthFraction, r.QuarkFourthFraction))
	}
	return fmt.Sprintf("bothTotal=%t exactlyOnePerSlot=%t perSlotSelected=%q results=[%s] verdict=%s", s.BothPassTotalCapacity, s.ExactlyOnePassPerSlot, s.PerSlotSelectedBranch, strings.Join(parts, "; "), s.Verdict)
}

func FormatVeto(v GeometricVetoAudit) string {
	return fmt.Sprintf("totalNativeDiag=%t totalVetoRMinus=%t perSlotWouldVeto=%t perSlotNative=%t rMinusViolates=%t reason=%q verdict=%s", v.TotalCapacityIsNativeDiagnostic, v.TotalCapacityCanVetoRMinus, v.PerSlotCapacityWouldVetoRMinus, v.PerSlotCapacityIsNativeTheorem, v.RMinusViolatesDerivedGeometry, v.Reason, v.Verdict)
}

func FormatSelection(s BranchSelection) string {
	return fmt.Sprintf("evaluated=%v nativeElim=%v unsealedElim=%v survivors=%v unique=%t selected=%q verdict=%s", s.BranchesEvaluated, s.EliminatedByNative, s.EliminatedByUnsealed, s.SurvivingBranches, s.UniqueBranchSelected, s.SelectedBranch, s.Verdict)
}

func FormatHiggs(h HiggsFirewall) string {
	return fmt.Sprintf("branch=%t heatKernel=%t scalarGauge=%t claimed=%t verdict=%s", h.BranchSelected, h.HeatKernelProjectionDerived, h.ScalarGaugeNormalizationDerived, h.HiggsPredictionClaimed, h.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noMultToAmp=%t noPerSlotTheorem=%t noVeto=%t noBranch=%t noHiggs=%t polluted=%t verdict=%s", f.DoesNotPromoteMultiplicityToAmplitude, f.DoesNotUsePerSlotBoundAsTheorem, f.DoesNotVetoRMinusWithoutSelector, f.DoesNotClaimBranchSelection, f.DoesNotClaimHiggsPrediction, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate289=%t capacity=%t stress=%t totalSelects=%t perSlotDiagnostic=%t nativeVeto=%t unique=%t higgs=%t firewall=%t next=%q status=%s", s.Gate289Inherited, s.CapacityFormalized, s.BranchStressCompleted, s.TotalCapacitySelects, s.PerSlotDiagnostic, s.NativeVetoDerived, s.UniqueBranchSelected, s.HiggsPredictionDerived, s.FirewallPreserved, s.NextGate, s.Status)
}
