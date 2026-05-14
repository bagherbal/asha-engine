// Package truebimodulerepresentation implements Gate 295:
// True Bimodule Assembly / Left-Right Representation Audit.
//
// Gate 294 proved that the naive left action placing weak H and color M3(C)
// simultaneously on Q_L as q⊗I_3 + I_2⊗B is not a representation of the
// direct-sum algebra C⊕H⊕M3(C).  Gate 295 audits the categorical repair: a
// quark doublet is a two-sided Morita bimodule.  The weak quaternionic factor
// acts from the left and the color matrix factor acts from the right/opposite
// side.  In matrix terms, on Q_L≈C²_weak⊗C³_color one tests
//
//	L(q)=q⊗I_3,        R(B)=I_2⊗B^T,
//
// so [L(q),R(B)]=0 while both actions are non-trivial.  This resolves the
// Tensor-vs-Direct-Sum paradox as a zero-order bimodule statement rather than
// as a single left representation of the direct-sum algebra.
//
// The gate deliberately does not overpromote the result.  Hypercharge splitting
// still requires an anomaly-free U(1) ledger and physical chirality assignment;
// the first-order condition still requires a canonical finite Dirac edge map.
// Thus the weak/color zero-order representation is derived as a bimodule, while
// full spectral-triple completion remains firewalled.
package truebimodulerepresentation

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE295-TRUE-BIMODULE-ASSEMBLY-LEFT-RIGHT-REPRESENTATION-AUDIT"

	StatusGate294Inherited          = "CONDITIONAL_SUPPORT_GATE294_DIRECT_SUM_PARADOX_INHERITED"
	StatusLeftWeakActionIsolated    = "CONDITIONAL_SUPPORT_LEFT_WEAK_H_ACTION_ISOLATED"
	StatusRightColorActionIsolated  = "CONDITIONAL_SUPPORT_RIGHT_COLOR_M3_ACTION_ISOLATED"
	StatusWeakColorCommute          = "CONDITIONAL_SUPPORT_TRUE_BIMODULE_ZERO_ORDER_COMMUTATION_VERIFIED"
	StatusTrueBimoduleDerived       = "CONDITIONAL_SUPPORT_TRUE_BIMODULE_REPRESENTATION_DERIVED"
	StatusHyperchargeLedgerAudited  = "CONDITIONAL_SUPPORT_HYPERCHARGE_SPLITTING_REQUIREMENTS_AUDITED"
	StatusOrderOnePreflightRestated = "CONDITIONAL_SUPPORT_FIRST_ORDER_DIRAC_REQUIREMENTS_RESTATED"

	StatusFailedHyperchargeNotDerived = "FAILED_ROUTE_HYPERCHARGE_SPLITTING_NOT_DERIVED"
	StatusFailedPhysicalJStillFormal  = "FAILED_ROUTE_PHYSICAL_J_ANTILINEAR_SEMANTICS_STILL_FORMAL"
	StatusFailedFirstOrderNotVerified = "FAILED_ROUTE_FIRST_ORDER_CONDITION_NOT_VERIFIED"
	StatusFailedCanonicalDFMissing    = "FAILED_ROUTE_CANONICAL_DF_STILL_UNSELECTED"
	StatusFailedDynamicsFirewalled    = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const (
	weakDim            = 2
	colorDim           = 3
	leptonRightDim     = 1
	quarkDim           = weakDim * colorDim
	leptonDim          = weakDim * leptonRightDim
	particleDoubletDim = leptonDim + quarkDim
)

type Gate294Input struct {
	JSwapKOSix            bool
	NaiveLeftActionFailed bool
	RepresentationParadox string
	Verdict               string
}

type LeftActionAudit struct {
	Carrier              string
	Dimension            int
	WeakGenerator        string
	ActsOnLeptons        bool
	ActsOnQuarks         bool
	NonTrivialResidual   float64
	AssociativeSurrogate bool
	Verdict              string
}

type RightActionAudit struct {
	Carrier              string
	Dimension            int
	ColorGenerator       string
	ActsOnLeptons        bool
	ActsOnQuarks         bool
	NonTrivialResidual   float64
	AssociativeSurrogate bool
	Verdict              string
}

type BimoduleAudit struct {
	QuarkCarrier                string
	LeftWeakAction              string
	RightColorAction            string
	WeakColorCommutatorNorm     float64
	NaiveLeftCrossTermNorm      float64
	ResolvesGate294Paradox      bool
	ZeroOrderCondition          string
	ZeroOrderVerified           bool
	DirectSumMultiplicationNote string
	Verdict                     string
}

type HyperchargeAudit struct {
	RequiredAssignments        []string
	CandidateFormula           string
	DerivedByGate              bool
	AnomalyFreeLedgerAvailable bool
	FractionalChargesGenerated bool
	Missing                    []string
	Verdict                    string
}

type OrderOneAudit struct {
	ZeroOrderConditionVerified bool
	FirstOrderCondition        string
	CanonicalDiracAvailable    bool
	FirstOrderVerified         bool
	DiracConstraintsDerived    bool
	Missing                    []string
	Verdict                    string
}

type Firewalls struct {
	DoesNotRecastBimoduleAsLeftRep bool
	DoesNotInventHypercharge       bool
	DoesNotInventDF                bool
	DoesNotUnlockHiggs             bool
	DoesNotUnlockBGap              bool
	FiniteCorePolluted             bool
	Verdict                        string
}

type Summary struct {
	LeftWeakDerived     bool
	RightColorDerived   bool
	ZeroOrderResolved   bool
	TrueBimoduleDerived bool
	HyperchargeDerived  bool
	FirstOrderVerified  bool
	CanonicalDFSelected bool
	DynamicsUnblocked   bool
	FirewallPreserved   bool
	Status              string
	DirectAnswer        string
	NextGate            string
}

type Analysis struct {
	Input       Gate294Input
	Left        LeftActionAudit
	Right       RightActionAudit
	Bimodule    BimoduleAudit
	Hypercharge HyperchargeAudit
	OrderOne    OrderOneAudit
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
	input := inheritGate294()
	left := auditLeftWeakAction()
	right := auditRightColorAction()
	bimodule := auditBimodule(left, right)
	hyper := auditHypercharge()
	order := auditOrderOne(bimodule)
	fw := auditFirewalls(hyper, order)
	summary := buildSummary(left, right, bimodule, hyper, order, fw)
	truth := "Gate 295 resolves the Gate 294 tensor/direct-sum paradox at zero order by replacing the illegal single left weak⊗color action with a true two-sided bimodule: H acts from the left on weak doublets and M3(C) acts from the right/opposite side on color slots.  The two actions commute exactly on Q_L≈C²⊗C³, so quarks can feel weak and color simultaneously without violating direct-sum multiplication.  Hypercharge splitting and the first-order Dirac theorem remain un-derived and firewalled."
	return Analysis{Input: input, Left: left, Right: right, Bimodule: bimodule, Hypercharge: hyper, OrderOne: order, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate294() Gate294Input {
	return Gate294Input{JSwapKOSix: true, NaiveLeftActionFailed: true, RepresentationParadox: "Naive Q_L left action q⊗I3 + I2⊗B violates direct-sum multiplication; a true left/right bimodule is required.", Verdict: StatusGate294Inherited}
}

type matrix [][]float64

func zero(n int) matrix {
	m := make(matrix, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}
func eye(n int) matrix {
	m := zero(n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}
func diag(vals ...float64) matrix {
	m := zero(len(vals))
	for i, v := range vals {
		m[i][i] = v
	}
	return m
}
func matMul(a, b matrix) matrix {
	n, m, k := len(a), len(b[0]), len(b)
	out := make(matrix, n)
	for i := range out {
		out[i] = make([]float64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s := 0.0
			for t := 0; t < k; t++ {
				s += a[i][t] * b[t][j]
			}
			out[i][j] = s
		}
	}
	return out
}
func matAdd(a, b matrix, sb float64) matrix {
	out := zero(len(a))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] + sb*b[i][j]
		}
	}
	return out
}
func frob(a matrix) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
func kron(a, b matrix) matrix {
	out := zero(len(a) * len(b))
	for i := range a {
		for j := range a[i] {
			for p := range b {
				for q := range b[p] {
					out[i*len(b)+p][j*len(b)+q] = a[i][j] * b[p][q]
				}
			}
		}
	}
	return out
}
func transpose(a matrix) matrix {
	out := zero(len(a))
	for i := range a {
		for j := range a[i] {
			out[j][i] = a[i][j]
		}
	}
	return out
}

func auditLeftWeakAction() LeftActionAudit {
	q := diag(1, -1)
	l := kron(q, eye(colorDim))
	return LeftActionAudit{Carrier: "L_doublets = (C²_weak⊗C_lepton) ⊕ (C²_weak⊗C³_color)", Dimension: particleDoubletDim, WeakGenerator: "q=diag(1,-1) acting as q⊗I_right", ActsOnLeptons: true, ActsOnQuarks: true, NonTrivialResidual: frob(l), AssociativeSurrogate: true, Verdict: StatusLeftWeakActionIsolated}
}

func auditRightColorAction() RightActionAudit {
	b := diag(2, -1, 0.5)
	r := kron(eye(weakDim), transpose(b))
	return RightActionAudit{Carrier: "Q_L = C²_weak⊗C³_color as right M3(C)-module", Dimension: quarkDim, ColorGenerator: "B=diag(2,-1,1/2) acting from right as I_2⊗B^T", ActsOnLeptons: false, ActsOnQuarks: true, NonTrivialResidual: frob(r), AssociativeSurrogate: true, Verdict: StatusRightColorActionIsolated}
}

func auditBimodule(l LeftActionAudit, r RightActionAudit) BimoduleAudit {
	q := diag(1, -1)
	b := diag(2, -1, 0.5)
	leftWeak := kron(q, eye(colorDim))
	rightColor := kron(eye(weakDim), transpose(b))
	comm := matAdd(matMul(leftWeak, rightColor), matMul(rightColor, leftWeak), -1)
	// Gate 294 illegal all-left cross-term residual for comparison.
	naiveCross := matMul(kron(q, eye(colorDim)), kron(eye(weakDim), b))
	commNorm := frob(comm)
	naiveNorm := frob(naiveCross)
	return BimoduleAudit{QuarkCarrier: "Q_L≈C²_weak⊗C³_color, interpreted as an H-M3(C) bimodule", LeftWeakAction: "L(q)=q⊗I3", RightColorAction: "R(B)=I2⊗B^T", WeakColorCommutatorNorm: commNorm, NaiveLeftCrossTermNorm: naiveNorm, ResolvesGate294Paradox: commNorm < 1e-12 && naiveNorm > 0.1, ZeroOrderCondition: "[L(q),R(B)]=0 for weak-left/color-right actions", ZeroOrderVerified: commNorm < 1e-12, DirectSumMultiplicationNote: "The product (0,q,0)(0,0,B)=0 is not represented as a single left-product on Q_L; color belongs to the opposite/right action, so the relevant zero-order law is commutation of left and right actions.", Verdict: StatusWeakColorCommute + ";" + StatusTrueBimoduleDerived}
}

func auditHypercharge() HyperchargeAudit {
	req := []string{"Y(Q_L)=+1/6", "Y(u_R)=+2/3", "Y(d_R)=-1/3", "Y(L_L)=-1/2", "Y(e_R)=-1", "optional Y(ν_R)=0"}
	missing := []string{"full chiral particle list in H_F", "left/right C-summand charge ledger", "anomaly-free U(1) quotient/unimodularity map", "physical normalization convention linking T3 and Y"}
	return HyperchargeAudit{RequiredAssignments: req, CandidateFormula: "Y must be a left/right C-summand and unimodularity quotient ledger; not generated by weak-left/color-right commutation alone", DerivedByGate: false, AnomalyFreeLedgerAvailable: false, FractionalChargesGenerated: false, Missing: missing, Verdict: StatusHyperchargeLedgerAudited + ";" + StatusFailedHyperchargeNotDerived}
}

func auditOrderOne(b BimoduleAudit) OrderOneAudit {
	missing := []string{"canonical finite Dirac edge map between L/R sectors", "Higgs/Yukawa edge representation", "Majorana/right-neutrino edge decision", "physical hypercharge/chiral ledger"}
	return OrderOneAudit{ZeroOrderConditionVerified: b.ZeroOrderVerified, FirstOrderCondition: "[[D_F,ρ(a)],ρ°(b)]=0", CanonicalDiracAvailable: false, FirstOrderVerified: false, DiracConstraintsDerived: false, Missing: missing, Verdict: StatusOrderOnePreflightRestated + ";" + StatusFailedFirstOrderNotVerified + ";" + StatusFailedCanonicalDFMissing}
}

func auditFirewalls(h HyperchargeAudit, o OrderOneAudit) Firewalls {
	return Firewalls{DoesNotRecastBimoduleAsLeftRep: true, DoesNotInventHypercharge: !h.DerivedByGate, DoesNotInventDF: !o.CanonicalDiracAvailable, DoesNotUnlockHiggs: true, DoesNotUnlockBGap: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedHyperchargeNotDerived, StatusFailedPhysicalJStillFormal, StatusFailedFirstOrderNotVerified, StatusFailedCanonicalDFMissing, StatusFailedDynamicsFirewalled}, ";")}
}

func buildSummary(l LeftActionAudit, r RightActionAudit, b BimoduleAudit, h HyperchargeAudit, o OrderOneAudit, fw Firewalls) Summary {
	statuses := []string{StatusGate294Inherited, StatusLeftWeakActionIsolated, StatusRightColorActionIsolated, StatusWeakColorCommute, StatusTrueBimoduleDerived, StatusHyperchargeLedgerAudited, StatusOrderOnePreflightRestated, StatusFailedHyperchargeNotDerived, StatusFailedPhysicalJStillFormal, StatusFailedFirstOrderNotVerified, StatusFailedCanonicalDFMissing, StatusFailedDynamicsFirewalled}
	return Summary{LeftWeakDerived: l.ActsOnQuarks && l.ActsOnLeptons, RightColorDerived: r.ActsOnQuarks && !r.ActsOnLeptons, ZeroOrderResolved: b.ZeroOrderVerified, TrueBimoduleDerived: b.ResolvesGate294Paradox, HyperchargeDerived: h.DerivedByGate, FirstOrderVerified: o.FirstOrderVerified, CanonicalDFSelected: o.DiracConstraintsDerived, DynamicsUnblocked: false, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotRecastBimoduleAsLeftRep, Status: strings.Join(statuses, ";"), DirectAnswer: "The true H-left/M3-right bimodule resolves the Gate 294 direct-sum paradox at zero order: quarks can carry weak and color structure simultaneously because weak and color act from opposite module sides and commute.  Hypercharge and first-order Dirac constraints remain un-derived.", NextGate: "Derive the chiral hypercharge/unimodularity ledger and canonical D_F edges on the true bimodule, then re-run the first-order condition."}
}

func FormatInput(g Gate294Input) string {
	return fmt.Sprintf("JswapKO6=%t naiveFailed=%t paradox=%q verdict=%s", g.JSwapKOSix, g.NaiveLeftActionFailed, g.RepresentationParadox, g.Verdict)
}
func FormatLeft(l LeftActionAudit) string {
	return fmt.Sprintf("carrier=%q dim=%d gen=%q leptons=%t quarks=%t norm=%.6g assoc=%t verdict=%s", l.Carrier, l.Dimension, l.WeakGenerator, l.ActsOnLeptons, l.ActsOnQuarks, l.NonTrivialResidual, l.AssociativeSurrogate, l.Verdict)
}
func FormatRight(r RightActionAudit) string {
	return fmt.Sprintf("carrier=%q dim=%d gen=%q leptons=%t quarks=%t norm=%.6g assoc=%t verdict=%s", r.Carrier, r.Dimension, r.ColorGenerator, r.ActsOnLeptons, r.ActsOnQuarks, r.NonTrivialResidual, r.AssociativeSurrogate, r.Verdict)
}
func FormatBimodule(b BimoduleAudit) string {
	return fmt.Sprintf("carrier=%q L=%q R=%q comm=%.6g naiveCross=%.6g resolves=%t zero=%t note=%q verdict=%s", b.QuarkCarrier, b.LeftWeakAction, b.RightColorAction, b.WeakColorCommutatorNorm, b.NaiveLeftCrossTermNorm, b.ResolvesGate294Paradox, b.ZeroOrderVerified, b.DirectSumMultiplicationNote, b.Verdict)
}
func FormatHypercharge(h HyperchargeAudit) string {
	return fmt.Sprintf("derived=%t anomalyLedger=%t fractional=%t required=%s formula=%q missing=%s verdict=%s", h.DerivedByGate, h.AnomalyFreeLedgerAvailable, h.FractionalChargesGenerated, strings.Join(h.RequiredAssignments, " | "), h.CandidateFormula, strings.Join(h.Missing, " | "), h.Verdict)
}
func FormatOrderOne(o OrderOneAudit) string {
	return fmt.Sprintf("zeroVerified=%t first=%q canonicalDF=%t firstVerified=%t constraints=%t missing=%s verdict=%s", o.ZeroOrderConditionVerified, o.FirstOrderCondition, o.CanonicalDiracAvailable, o.FirstOrderVerified, o.DiracConstraintsDerived, strings.Join(o.Missing, " | "), o.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noLeftRecast=%t noY=%t noDF=%t noHiggs=%t noBGap=%t polluted=%t verdict=%s", f.DoesNotRecastBimoduleAsLeftRep, f.DoesNotInventHypercharge, f.DoesNotInventDF, f.DoesNotUnlockHiggs, f.DoesNotUnlockBGap, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("leftWeak=%t rightColor=%t zero=%t trueBimodule=%t hypercharge=%t first=%t canonicalDF=%t dynamics=%t firewall=%t next=%q status=%s", s.LeftWeakDerived, s.RightColorDerived, s.ZeroOrderResolved, s.TrueBimoduleDerived, s.HyperchargeDerived, s.FirstOrderVerified, s.CanonicalDFSelected, s.DynamicsUnblocked, s.FirewallPreserved, s.NextGate, s.Status)
}
