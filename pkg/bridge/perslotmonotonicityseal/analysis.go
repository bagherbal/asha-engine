// Package perslotmonotonicityseal implements Gate 291:
// Per-Slot Monotonicity Seal / Final Spectral Synthesis Audit.
//
// Gate 290 proved that Morita multiplicity gives a total trace-capacity
// diagnostic but cannot natively impose the stronger per-slot amplitude
// ordering that would select r_+.  Gate 291 therefore introduces an explicit
// phenomenological seal: PerSlotMonotonicitySeal.  Under that seal, the lower
// branch r_- is vetoed, the upper branch r_+ is selected, and the final reduced
// trace identity is recomputed.  The gate deliberately does not promote this
// raw trace proxy to a Seeley-de Witt Higgs prediction; the six-point Higgs
// firewall remains active.
package perslotmonotonicityseal

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE291-PER-SLOT-MONOTONICITY-SEAL-FINAL-SPECTRAL-SYNTHESIS-AUDIT"

	StatusGate290Inherited       = "CONDITIONAL_SUPPORT_GATE290_TRACE_CAPACITY_BARRIER_INHERITED"
	StatusSealActivated          = "CONDITIONAL_SUPPORT_PER_SLOT_MONOTONICITY_SEAL_ACTIVATED"
	StatusVacuumLocked           = "CONDITIONAL_SUPPORT_R_PLUS_VACUUM_BRANCH_LOCKED_UNDER_SEAL"
	StatusTraceMomentsComputed   = "CONDITIONAL_SUPPORT_LOCKED_REDUCED_TRACE_MOMENTS_COMPUTED"
	StatusFinalSynthesisAchieved = "CONDITIONAL_SUPPORT_VACUUM_LOCKED_AND_FINAL_SPECTRAL_SYNTHESIS_ACHIEVED"
	StatusHiggsFirewallPreserved = "CONDITIONAL_SUPPORT_SIX_POINT_HIGGS_FIREWALL_REMAINS_ACTIVE"

	StatusFailedSealIsPhenomenological     = "FAILED_ROUTE_PER_SLOT_MONOTONICITY_NOT_NATIVE_GEOMETRIC_THEOREM"
	StatusFailedRMinusVetoSealedNotDerived = "FAILED_ROUTE_R_MINUS_VETO_IS_SEALED_NOT_DERIVED"
	StatusFailedHeatKernelMissing          = "FAILED_ROUTE_HEAT_KERNEL_PROJECTION_STILL_MISSING"
	StatusFailedScalarGaugeMissing         = "FAILED_ROUTE_SCALAR_GAUGE_NORMALIZATION_STILL_MISSING"
	StatusFailedHiggsMassNotClaimed        = "FAILED_ROUTE_RAW_TRACE_PROXY_NOT_PHYSICAL_HIGGS_RATIO"
)

const (
	kappaC = 1.0
	kappaQ = 3.0

	contactZeta0 = 7.0
	contactZeta2 = 61.0 / 25.0
	contactZeta4 = 257629.0 / 202500.0
	reducedA0    = kappaC + kappaQ
	sTop         = 8 * math.Pi * math.Pi

	lambdaNum = 1197.0
	lambdaDen = 4624.0
	lambda    = lambdaNum / lambdaDen
)

type Gate290Inheritance struct {
	TotalCapacityBothBranchesPass bool
	PerSlotDiagnosticSelectsRPlus bool
	PerSlotRuleNativeTheorem      bool
	BranchPreviouslySelected      bool
	HiggsPredictionDerived        bool
	Verdict                       string
}

type PerSlotMonotonicitySeal struct {
	Name               string
	Active             bool
	Rule               string
	Provenance         string
	Phenomenological   bool
	NativeTheorem      bool
	VetoedBranch       string
	SelectedBranch     string
	OperationalMeaning string
	Verdict            string
}

type Branch struct {
	Name             string
	Exact            string
	R                float64
	AbsYOverX        float64
	X                float64
	LeptonD2         float64
	QuarkD2          float64
	LeptonD4         float64
	QuarkD4          float64
	D2               float64
	D4               float64
	D4OverD2         float64
	D4OverD2Squared  float64
	PerSlotLeptonD2  float64
	PerSlotQuarkD2   float64
	PerSlotLeptonD4  float64
	PerSlotQuarkD4   float64
	PerSlotPassD2    bool
	PerSlotPassD4    bool
	SurvivesSeal     bool
	SelectionVerdict string
}

type LockedVacuumState struct {
	SelectedBranch   Branch
	VetoedBranches   []string
	UniqueUnderSeal  bool
	NativeUnique     bool
	XSource          string
	RSource          string
	ExactREquation   string
	ExactXConstraint string
	Verdict          string
}

type TraceMomentSynthesis struct {
	D2                  float64
	D4                  float64
	D4OverD2            float64
	D4OverD2Squared     float64
	ContactLambda       float64
	ContactLambdaExact  string
	LambdaResidualAbs   float64
	D2Formula           string
	D4Formula           string
	ShapeFormula        string
	ShapeMatchesContact bool
	Verdict             string
}

type HiggsFirewall struct {
	BranchLockedUnderSeal          bool
	RawProxyComputed               bool
	HeatKernelProjectionDerived    bool
	ScalarGaugeNormalization       bool
	DimensionlessObservableFinal   bool
	PhysicalHiggsPredictionClaimed bool
	ActiveObstructions             []string
	Verdict                        string
}

type Firewalls struct {
	SealDoesNotRewriteGate290       bool
	RMinusVetoMarkedSealed          bool
	RawProxyNotPromotedToA2A4       bool
	HeatKernelObstructionPreserved  bool
	ScalarGaugeObstructionPreserved bool
	PhysicalHiggsMassNotClaimed     bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	Gate290Inherited       bool
	SealActivated          bool
	BranchLockedUnderSeal  bool
	BranchLockedNatively   bool
	TraceMomentsComputed   bool
	ContactShapeReproduced bool
	RawProxyComputed       bool
	HiggsPredictionDerived bool
	FirewallPreserved      bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	Inheritance Gate290Inheritance
	Seal        PerSlotMonotonicitySeal
	Branches    []Branch
	Locked      LockedVacuumState
	Trace       TraceMomentSynthesis
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
	inh := inheritGate290()
	seal := activateSeal()
	branches, err := buildBranches(seal)
	if err != nil {
		return Analysis{}, err
	}
	locked, err := lockVacuum(branches)
	if err != nil {
		return Analysis{}, err
	}
	trace := synthesizeTrace(locked.SelectedBranch)
	higgs := auditHiggsFirewall(seal, trace)
	fw := auditFirewalls(seal, locked, higgs)
	summary := buildSummary(inh, seal, locked, trace, higgs, fw)
	truth := "Gate 291 activates the PerSlotMonotonicitySeal to choose r_+ as a representative physical vacuum branch. This is an explicit phenomenological orientation rule, not a finite-core theorem. Under the seal, the reduced trace moments are computed and their dimensionless shape exactly reproduces the Gate-169 contact scalar shape 1197/4624. The result is a final raw spectral synthesis, not a Seeley-de Witt Higgs mass prediction; the heat-kernel and scalar/gauge normalization firewalls remain active."
	return Analysis{Inheritance: inh, Seal: seal, Branches: branches, Locked: locked, Trace: trace, Higgs: higgs, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate290() Gate290Inheritance {
	return Gate290Inheritance{TotalCapacityBothBranchesPass: true, PerSlotDiagnosticSelectsRPlus: true, PerSlotRuleNativeTheorem: false, BranchPreviouslySelected: false, HiggsPredictionDerived: false, Verdict: StatusGate290Inherited}
}

func activateSeal() PerSlotMonotonicitySeal {
	return PerSlotMonotonicitySeal{
		Name:               "PerSlotMonotonicitySeal",
		Active:             true,
		Rule:               "Tr(P_Q D_F^{2n})/κ_Q > Tr(P_C D_F^{2n})/κ_C for n=1,2 in the reduced 1⊕3 Morita edge ledger",
		Provenance:         "phenomenological sector-orientation axiom isolated by Gate 290; not derived from κ_C:κ_Q multiplicity alone",
		Phenomenological:   true,
		NativeTheorem:      false,
		VetoedBranch:       "r_minus",
		SelectedBranch:     "r_plus",
		OperationalMeaning: "quark edge amplitude per slot exceeds lepton edge amplitude per slot; equivalently r=|y/x|²>1 in the reduced scalar-Morita proxy",
		Verdict:            StatusSealActivated,
	}
}

func buildBranches(seal PerSlotMonotonicitySeal) ([]Branch, error) {
	defs := []struct {
		name, exact string
		r           float64
	}{
		{"r_plus", "(3591 + 136√123)/3099", (3591 + 136*math.Sqrt(123)) / 3099},
		{"r_minus", "(3591 - 136√123)/3099", (3591 - 136*math.Sqrt(123)) / 3099},
	}
	branches := make([]Branch, 0, len(defs))
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
		perL2 := l2 / kappaC
		perQ2 := q2 / kappaQ
		perL4 := l4 / kappaC
		perQ4 := q4 / kappaQ
		passD2 := perQ2 > perL2
		passD4 := perQ4 > perL4
		survives := seal.Active && d.name == seal.SelectedBranch && passD2 && passD4
		verdict := StatusFailedRMinusVetoSealedNotDerived
		if survives {
			verdict = StatusVacuumLocked
		}
		branches = append(branches, Branch{Name: d.name, Exact: d.exact, R: d.r, AbsYOverX: math.Sqrt(d.r), X: x, LeptonD2: l2, QuarkD2: q2, LeptonD4: l4, QuarkD4: q4, D2: d2, D4: d4, D4OverD2: d4 / d2, D4OverD2Squared: d4 / (d2 * d2), PerSlotLeptonD2: perL2, PerSlotQuarkD2: perQ2, PerSlotLeptonD4: perL4, PerSlotQuarkD4: perQ4, PerSlotPassD2: passD2, PerSlotPassD4: passD4, SurvivesSeal: survives, SelectionVerdict: verdict})
	}
	return branches, nil
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
		if root > 0 && !math.IsNaN(root) && !math.IsInf(root, 0) {
			return root, nil
		}
	}
	return math.NaN(), fmt.Errorf("no positive X root for r=%g", r)
}

func lockVacuum(branches []Branch) (LockedVacuumState, error) {
	var selected *Branch
	vetoed := []string{}
	for i := range branches {
		b := branches[i]
		if b.SurvivesSeal {
			selected = &branches[i]
		} else {
			vetoed = append(vetoed, b.Name)
		}
	}
	if selected == nil {
		return LockedVacuumState{}, fmt.Errorf("seal did not select a surviving branch")
	}
	return LockedVacuumState{SelectedBranch: *selected, VetoedBranches: vetoed, UniqueUnderSeal: true, NativeUnique: false, XSource: "Gate 288 contact-spectral cutoff quadratic with reduced a0=4", RSource: "Gate 275 scalar-Morita shape branch plus Gate 291 PerSlotMonotonicitySeal", ExactREquation: "3099r² - 7182r + 3427 = 0", ExactXConstraint: "7X²(1+3r²)+(61/25)X(1+3r)+(257629/202500)·4=8π²", Verdict: StatusVacuumLocked}, nil
}

func synthesizeTrace(b Branch) TraceMomentSynthesis {
	residual := math.Abs(b.D4OverD2Squared - lambda)
	return TraceMomentSynthesis{D2: b.D2, D4: b.D4, D4OverD2: b.D4OverD2, D4OverD2Squared: b.D4OverD2Squared, ContactLambda: lambda, ContactLambdaExact: "1197/4624", LambdaResidualAbs: residual, D2Formula: "Tr(D_F²)=X(1+3r)", D4Formula: "Tr(D_F⁴)=X²(1+3r²)", ShapeFormula: "Tr(D_F⁴)/(Tr(D_F²))²", ShapeMatchesContact: residual < 1e-12, Verdict: StatusTraceMomentsComputed}
}

func auditHiggsFirewall(seal PerSlotMonotonicitySeal, trace TraceMomentSynthesis) HiggsFirewall {
	obs := []string{StatusFailedHeatKernelMissing, StatusFailedScalarGaugeMissing, StatusFailedHiggsMassNotClaimed}
	return HiggsFirewall{BranchLockedUnderSeal: seal.Active && seal.SelectedBranch == "r_plus", RawProxyComputed: trace.ShapeMatchesContact, HeatKernelProjectionDerived: false, ScalarGaugeNormalization: false, DimensionlessObservableFinal: false, PhysicalHiggsPredictionClaimed: false, ActiveObstructions: obs, Verdict: StatusHiggsFirewallPreserved}
}

func auditFirewalls(seal PerSlotMonotonicitySeal, locked LockedVacuumState, h HiggsFirewall) Firewalls {
	return Firewalls{SealDoesNotRewriteGate290: seal.Phenomenological && !seal.NativeTheorem && !locked.NativeUnique, RMinusVetoMarkedSealed: len(locked.VetoedBranches) == 1 && locked.VetoedBranches[0] == "r_minus", RawProxyNotPromotedToA2A4: !h.DimensionlessObservableFinal, HeatKernelObstructionPreserved: !h.HeatKernelProjectionDerived, ScalarGaugeObstructionPreserved: !h.ScalarGaugeNormalization, PhysicalHiggsMassNotClaimed: !h.PhysicalHiggsPredictionClaimed, FiniteCorePolluted: false, Verdict: StatusHiggsFirewallPreserved}
}

func buildSummary(inh Gate290Inheritance, seal PerSlotMonotonicitySeal, locked LockedVacuumState, trace TraceMomentSynthesis, h HiggsFirewall, fw Firewalls) Summary {
	statuses := []string{StatusGate290Inherited, StatusSealActivated, StatusVacuumLocked, StatusTraceMomentsComputed, StatusFinalSynthesisAchieved, StatusHiggsFirewallPreserved, StatusFailedSealIsPhenomenological, StatusFailedRMinusVetoSealedNotDerived, StatusFailedHeatKernelMissing, StatusFailedScalarGaugeMissing, StatusFailedHiggsMassNotClaimed}
	return Summary{Gate290Inherited: inh.PerSlotDiagnosticSelectsRPlus && !inh.PerSlotRuleNativeTheorem, SealActivated: seal.Active, BranchLockedUnderSeal: locked.UniqueUnderSeal, BranchLockedNatively: locked.NativeUnique, TraceMomentsComputed: trace.D2 > 0 && trace.D4 > 0, ContactShapeReproduced: trace.ShapeMatchesContact, RawProxyComputed: h.RawProxyComputed, HiggsPredictionDerived: h.PhysicalHiggsPredictionClaimed, FirewallPreserved: !fw.FiniteCorePolluted && fw.PhysicalHiggsMassNotClaimed, Status: strings.Join(statuses, ";"), DirectAnswer: "Under the explicit PerSlotMonotonicitySeal, r_+ is selected and r_- is vetoed. This is a sealed phenomenological branch lock, not a native theorem. The reduced raw trace proxy exactly reproduces 1197/4624, while the six-point Higgs firewall remains active.", NextGate: "Either stop with the sealed final synthesis, or derive the per-slot monotonicity rule, heat-kernel projection, scalar/gauge normalization, and physical J/hypercharge data before claiming a physical Higgs prediction."}
}

func FormatSeal(s PerSlotMonotonicitySeal) string {
	return fmt.Sprintf("%s active=%t native=%t phenom=%t rule=%q selected=%s vetoed=%s verdict=%s", s.Name, s.Active, s.NativeTheorem, s.Phenomenological, s.Rule, s.SelectedBranch, s.VetoedBranch, s.Verdict)
}

func FormatBranch(b Branch) string {
	return fmt.Sprintf("%s exact=%s r=%.15g |y/x|=%.15g X=%.15g D2(C,Q,total)=(%.15g,%.15g,%.15g) D4(C,Q,total)=(%.15g,%.15g,%.15g) perSlotD2(C,Q)=(%.15g,%.15g) perSlotD4(C,Q)=(%.15g,%.15g) survives=%t verdict=%s", b.Name, b.Exact, b.R, b.AbsYOverX, b.X, b.LeptonD2, b.QuarkD2, b.D2, b.LeptonD4, b.QuarkD4, b.D4, b.PerSlotLeptonD2, b.PerSlotQuarkD2, b.PerSlotLeptonD4, b.PerSlotQuarkD4, b.SurvivesSeal, b.SelectionVerdict)
}

func FormatLocked(l LockedVacuumState) string {
	return fmt.Sprintf("selected=%s vetoed=%v uniqueUnderSeal=%t nativeUnique=%t rSource=%q xSource=%q verdict=%s", l.SelectedBranch.Name, l.VetoedBranches, l.UniqueUnderSeal, l.NativeUnique, l.RSource, l.XSource, l.Verdict)
}

func FormatTrace(t TraceMomentSynthesis) string {
	return fmt.Sprintf("D2=%.15g D4=%.15g D4/D2=%.15g D4/D2^2=%.15g lambda=%s(%.15g) residual=%.3g match=%t verdict=%s", t.D2, t.D4, t.D4OverD2, t.D4OverD2Squared, t.ContactLambdaExact, t.ContactLambda, t.LambdaResidualAbs, t.ShapeMatchesContact, t.Verdict)
}

func FormatHiggs(h HiggsFirewall) string {
	return fmt.Sprintf("branchUnderSeal=%t rawProxy=%t heatKernel=%t scalarGauge=%t finalObs=%t physicalClaim=%t obstructions=%v verdict=%s", h.BranchLockedUnderSeal, h.RawProxyComputed, h.HeatKernelProjectionDerived, h.ScalarGaugeNormalization, h.DimensionlessObservableFinal, h.PhysicalHiggsPredictionClaimed, h.ActiveObstructions, h.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noRewrite=%t vetoSealed=%t proxyNotA2A4=%t heatKernel=%t scalarGauge=%t noPhysicalHiggs=%t polluted=%t verdict=%s", f.SealDoesNotRewriteGate290, f.RMinusVetoMarkedSealed, f.RawProxyNotPromotedToA2A4, f.HeatKernelObstructionPreserved, f.ScalarGaugeObstructionPreserved, f.PhysicalHiggsMassNotClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate290=%t seal=%t lockedUnderSeal=%t native=%t traces=%t shape=%t rawProxy=%t higgs=%t firewall=%t next=%q status=%s", s.Gate290Inherited, s.SealActivated, s.BranchLockedUnderSeal, s.BranchLockedNatively, s.TraceMomentsComputed, s.ContactShapeReproduced, s.RawProxyComputed, s.HiggsPredictionDerived, s.FirewallPreserved, s.NextGate, s.Status)
}
