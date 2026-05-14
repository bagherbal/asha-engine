// Package doubledspacerepresentation implements Gate 294:
// Doubled-Space Representation / Opposite Algebra Action Assembly Audit.
//
// Gate 293 showed that the doubled-space swap J_swap gives the desired KO6
// grading sign without selecting an arbitrary one-mode odd twist.  Gate 294
// asks the next question: can the full finite algebra C⊕H⊕M3(C) be represented
// on the doubled physical finite Hilbert space so that J_swap constructs a
// lawful opposite action and the zero/order-one conditions can be verified?
//
// The audit deliberately distinguishes three notions that are often conflated:
//  1. a KO-sign-correct doubled swap operator;
//  2. an associative representation of the direct-sum algebra;
//  3. the physical SM bimodule where weak SU(2) and color act on the same
//     quark doublet through left/right module structure rather than through a
//     naive direct-sum left action.
//
// Result: J_swap is sign-correct and a formal opposite-action formula can be
// written for any already-valid representation.  However, the naive full left
// action that tries to place H and M3(C) simultaneously on Q_L as
// q⊗I_3 + I_2⊗B is not a representation of the direct-sum algebra: cross terms
// appear although (0,q,0)(0,0,B)=0 in C⊕H⊕M3(C).  Block-separated direct-sum
// representations are associative but not the physical quark weak/color
// bimodule.  Therefore the physical opposite action and full order-one theorem
// remain unconstructed.
package doubledspacerepresentation

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE294-DOUBLED-SPACE-REPRESENTATION-OPPOSITE-ALGEBRA-ACTION-ASSEMBLY-AUDIT"

	StatusGate293Inherited       = "CONDITIONAL_SUPPORT_GATE293_JSWAP_KO6_CANDIDATE_INHERITED"
	StatusDoubledKOSignsVerified = "CONDITIONAL_SUPPORT_DOUBLED_JSWAP_KO_SIGNS_VERIFIED"
	StatusRepresentationAudited  = "CONDITIONAL_SUPPORT_C_PLUS_H_PLUS_M3C_REPRESENTATION_CANDIDATES_AUDITED"
	StatusNaiveFailureCertified  = "CONDITIONAL_SUPPORT_NAIVE_WEAK_COLOR_DIRECT_SUM_ACTION_FAILURE_CERTIFIED"
	StatusFormalOppositeDefined  = "CONDITIONAL_SUPPORT_FORMAL_JSWAP_OPPOSITE_ACTION_FORMULA_DEFINED"
	StatusOrderOnePreflight      = "CONDITIONAL_SUPPORT_ORDER_ONE_REQUIREMENTS_RESTATED_ON_DOUBLED_SPACE"

	StatusFailedPhysicalRepMissing     = "FAILED_ROUTE_PHYSICAL_HF_REPRESENTATION_OF_C_PLUS_H_PLUS_M3C_NOT_DERIVED"
	StatusFailedNaiveNotRepresentation = "FAILED_ROUTE_NAIVE_QLEFT_H_AND_COLOR_ACTION_IS_NOT_DIRECT_SUM_REPRESENTATION"
	StatusFailedBlockNotSMBimodule     = "FAILED_ROUTE_BLOCK_SEPARATED_ACTION_IS_ASSOCIATIVE_BUT_NOT_PHYSICAL_SM_BIMODULE"
	StatusFailedOppositeNotConstructed = "FAILED_ROUTE_PHYSICAL_OPPOSITE_ACTION_NOT_CONSTRUCTED"
	StatusFailedOrderOneNotVerified    = "FAILED_ROUTE_FULL_ORDER_ONE_CONDITION_NOT_VERIFIED"
	StatusFailedDFStillUnselected      = "FAILED_ROUTE_CANONICAL_DF_STILL_UNSELECTED"
	StatusFailedDynamicsFirewalled     = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const (
	fiberDim    = 4
	doubledDim  = 2 * fiberDim
	qLeftToyDim = 6 // weak doublet ⊗ color triplet diagnostic
)

type Gate293Input struct {
	CandidateName string
	DoubledDim    int
	J2Sign        int
	JGammaSign    int
	KOSixLike     bool
	Verdict       string
}

type JSwapAudit struct {
	DoubledDimension int
	J2Sign           int
	JGammaSign       int
	KOSixLike        bool
	ResidualJ2       float64
	ResidualGamma    float64
	Verdict          string
}

type RepresentationCandidate struct {
	Name                  string
	CarrierDescription    string
	Associative           bool
	Unital                bool
	FaithfulToAllSummands bool
	PhysicalSMBimodule    bool
	FailureMechanism      string
	Residual              float64
	Verdict               string
}

type NaiveWeakColorDiagnostic struct {
	AlgebraProductShouldVanish bool
	ImageProductNorm           float64
	MultiplicativityResidual   float64
	Explanation                string
	Verdict                    string
}

type OppositeActionAudit struct {
	JSwapAvailable              bool
	RequiresValidRepresentation bool
	Formula                     string
	ConstructedForPhysicalHF    bool
	ZeroOrderVerified           bool
	Missing                     []string
	Verdict                     string
}

type OrderOneAudit struct {
	ZeroOrderCondition              string
	FirstOrderCondition             string
	PhysicalDFAvailable             bool
	PhysicalRepresentationAvailable bool
	OrderOneVerified                bool
	DiracConstraintsDerived         bool
	Verdict                         string
}

type Firewalls struct {
	DoesNotPromoteNaiveAction bool
	DoesNotInventHF           bool
	DoesNotInventOpposite     bool
	DoesNotSelectDF           bool
	DoesNotUnlockHiggs        bool
	DoesNotUnlockBGap         bool
	FiniteCorePolluted        bool
	Verdict                   string
}

type Summary struct {
	JSwapKOSix              bool
	AnyAssociativeCandidate bool
	PhysicalSMRepDerived    bool
	OppositeConstructed     bool
	OrderOneVerified        bool
	CanonicalDFSelected     bool
	DynamicsUnblocked       bool
	FirewallPreserved       bool
	Status                  string
	DirectAnswer            string
	NextGate                string
}

type Analysis struct {
	Input           Gate293Input
	JSwap           JSwapAudit
	Representations []RepresentationCandidate
	NaiveDiagnostic NaiveWeakColorDiagnostic
	Opposite        OppositeActionAudit
	OrderOne        OrderOneAudit
	Firewalls       Firewalls
	Summary         Summary
	Truth           string
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
	input := inheritGate293()
	j := auditJSwap()
	naive := auditNaiveWeakColorAction()
	reps := auditRepresentationCandidates(naive)
	opp := auditOppositeAction(j, reps)
	order := auditOrderOne(reps, opp)
	fw := auditFirewalls(reps, opp, order)
	summary := buildSummary(j, reps, opp, order, fw)
	truth := "Gate 294 verifies that J_swap on H_F⊕H_F* has the required KO6 grading sign and can formally define an opposite-action formula for an already-valid representation.  The hard obstruction is the representation itself: the naive Q_L action q⊗I_3 + I_2⊗B is not a representation of the direct-sum algebra C⊕H⊕M3(C), while block-separated representations are associative but not the physical weak-doublet/color bimodule.  Therefore the physical opposite algebra action and full order-one theorem remain blocked."
	return Analysis{Input: input, JSwap: j, Representations: reps, NaiveDiagnostic: naive, Opposite: opp, OrderOne: order, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate293() Gate293Input {
	return Gate293Input{CandidateName: "J_swap on H_F⊕H_F* with grading diag(γ,-γ)", DoubledDim: doubledDim, J2Sign: +1, JGammaSign: -1, KOSixLike: true, Verdict: StatusGate293Inherited}
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
func blockDiag(blocks ...matrix) matrix {
	n := 0
	for _, b := range blocks {
		n += len(b)
	}
	out := zero(n)
	off := 0
	for _, b := range blocks {
		for i := range b {
			for j := range b[i] {
				out[off+i][off+j] = b[i][j]
			}
		}
		off += len(b)
	}
	return out
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
func frobDiff(a, b matrix) float64 { return frob(matAdd(a, b, -1)) }
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
func swap(n int) matrix {
	out := zero(2 * n)
	for i := 0; i < n; i++ {
		out[i][n+i] = 1
		out[n+i][i] = 1
	}
	return out
}

func auditJSwap() JSwapAudit {
	gammaF := diag(1, -1, -1, 1)
	gammaD := blockDiag(gammaF, scalarMul(gammaF, -1))
	j := swap(fiberDim)
	j2 := frobDiff(matMul(j, j), eye(doubledDim))
	jg := matMul(j, gammaD)
	gj := matMul(gammaD, j)
	anti := frob(matAdd(jg, gj, +1))
	return JSwapAudit{DoubledDimension: doubledDim, J2Sign: +1, JGammaSign: -1, KOSixLike: j2 < 1e-12 && anti < 1e-12, ResidualJ2: j2, ResidualGamma: anti, Verdict: StatusDoubledKOSignsVerified}
}
func scalarMul(a matrix, s float64) matrix {
	out := zero(len(a))
	for i := range a {
		for j := range a[i] {
			out[i][j] = s * a[i][j]
		}
	}
	return out
}

func auditNaiveWeakColorAction() NaiveWeakColorDiagnostic {
	// In C⊕H⊕M3(C), a=(0,q,0) and b=(0,0,B) multiply to zero.
	// A tempting Q_L carrier map sends a -> q⊗I_3 and b -> I_2⊗B.
	// Then rho(a)rho(b)=q⊗B, generally nonzero, proving this is a tensor-product
	// algebra action, not a representation of the direct-sum algebra.
	q := diag(1, -1)      // simple quaternionic/weak generator surrogate
	b := diag(2, -1, 0.5) // simple color generator surrogate
	rhoA := kron(q, eye(3))
	rhoB := kron(eye(2), b)
	imageProduct := matMul(rhoA, rhoB)
	res := frob(imageProduct) // since rho(ab)=rho(0)=0
	return NaiveWeakColorDiagnostic{AlgebraProductShouldVanish: true, ImageProductNorm: res, MultiplicativityResidual: res, Explanation: "For direct-sum algebra elements (0,q,0)(0,0,B)=0, but the naive quark-doublet image gives (q⊗I3)(I2⊗B)=q⊗B≠0.  This is an H⊗M3 tensor-product action, not a left representation of C⊕H⊕M3(C).", Verdict: StatusNaiveFailureCertified + ";" + StatusFailedNaiveNotRepresentation}
}

func auditRepresentationCandidates(n NaiveWeakColorDiagnostic) []RepresentationCandidate {
	return []RepresentationCandidate{
		{Name: "naive_Q_L_tensor_action", CarrierDescription: "Q_L≈C²_weak⊗C³_color with ρ(q,B)=q⊗I3 + I2⊗B", Associative: false, Unital: false, FaithfulToAllSummands: true, PhysicalSMBimodule: false, FailureMechanism: n.Explanation, Residual: n.MultiplicativityResidual, Verdict: StatusFailedNaiveNotRepresentation},
		{Name: "block_separated_direct_sum_action", CarrierDescription: "C, H, and M3(C) act on disjoint summand blocks; this is a direct-sum representation but does not put weak and color structure on the same quark doublet", Associative: true, Unital: true, FaithfulToAllSummands: true, PhysicalSMBimodule: false, FailureMechanism: "Associativity is recovered only by separating the summands; the physical Q_L weak-doublet/color-triplet bimodule is lost.", Residual: 0, Verdict: StatusRepresentationAudited + ";" + StatusFailedBlockNotSMBimodule},
		{Name: "Morita_two_sided_bimodule_target", CarrierDescription: "H and M3(C) must act through left/right module structure with J-defined opposite action, not as one naive left direct-sum action", Associative: true, Unital: true, FaithfulToAllSummands: false, PhysicalSMBimodule: false, FailureMechanism: "This is the correct categorical target, but the physical H_F sub-bimodule, hypercharge/chirality attachment, and J-defined opposite action are not derived in this gate.", Residual: 0, Verdict: StatusRepresentationAudited + ";" + StatusFailedPhysicalRepMissing},
	}
}

func auditOppositeAction(j JSwapAudit, reps []RepresentationCandidate) OppositeActionAudit {
	physical := false
	for _, r := range reps {
		if r.PhysicalSMBimodule && r.Associative {
			physical = true
		}
	}
	missing := []string{"physical H_F sub-bimodule for C⊕H⊕M3(C)", "hypercharge/chirality assignment", "anti-linear complex/quaternionic conjugation semantics beyond real J_swap matrix", "left/right Morita module placement of H and M3(C)", "canonical D_F edge map"}
	return OppositeActionAudit{JSwapAvailable: j.KOSixLike, RequiresValidRepresentation: true, Formula: "ρ°(a)=J_swap ρ(a*) J_swap^{-1}", ConstructedForPhysicalHF: physical, ZeroOrderVerified: false, Missing: missing, Verdict: StatusFormalOppositeDefined + ";" + StatusFailedOppositeNotConstructed}
}

func auditOrderOne(reps []RepresentationCandidate, opp OppositeActionAudit) OrderOneAudit {
	return OrderOneAudit{ZeroOrderCondition: "[ρ(a),ρ°(b)]=0", FirstOrderCondition: "[[D_F,ρ(a)],ρ°(b)]=0", PhysicalDFAvailable: false, PhysicalRepresentationAvailable: opp.ConstructedForPhysicalHF, OrderOneVerified: false, DiracConstraintsDerived: false, Verdict: StatusOrderOnePreflight + ";" + StatusFailedOrderOneNotVerified + ";" + StatusFailedDFStillUnselected}
}

func auditFirewalls(reps []RepresentationCandidate, opp OppositeActionAudit, order OrderOneAudit) Firewalls {
	return Firewalls{DoesNotPromoteNaiveAction: true, DoesNotInventHF: true, DoesNotInventOpposite: !opp.ConstructedForPhysicalHF, DoesNotSelectDF: !order.DiracConstraintsDerived, DoesNotUnlockHiggs: true, DoesNotUnlockBGap: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedPhysicalRepMissing, StatusFailedNaiveNotRepresentation, StatusFailedOppositeNotConstructed, StatusFailedOrderOneNotVerified, StatusFailedDFStillUnselected, StatusFailedDynamicsFirewalled}, ";")}
}

func buildSummary(j JSwapAudit, reps []RepresentationCandidate, opp OppositeActionAudit, order OrderOneAudit, fw Firewalls) Summary {
	anyAssoc := false
	for _, r := range reps {
		if r.Associative {
			anyAssoc = true
		}
	}
	statuses := []string{StatusGate293Inherited, StatusDoubledKOSignsVerified, StatusRepresentationAudited, StatusNaiveFailureCertified, StatusFormalOppositeDefined, StatusOrderOnePreflight, StatusFailedPhysicalRepMissing, StatusFailedNaiveNotRepresentation, StatusFailedBlockNotSMBimodule, StatusFailedOppositeNotConstructed, StatusFailedOrderOneNotVerified, StatusFailedDFStillUnselected, StatusFailedDynamicsFirewalled}
	return Summary{JSwapKOSix: j.KOSixLike, AnyAssociativeCandidate: anyAssoc, PhysicalSMRepDerived: false, OppositeConstructed: opp.ConstructedForPhysicalHF, OrderOneVerified: order.OrderOneVerified, CanonicalDFSelected: order.DiracConstraintsDerived, DynamicsUnblocked: false, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotPromoteNaiveAction, Status: strings.Join(statuses, ";"), DirectAnswer: "J_swap solves the KO-sign problem, but it does not by itself supply a physical representation.  The naive weak/color Q_L action fails direct-sum multiplicativity, while block-separated actions lose the physical SM bimodule.  The physical opposite algebra action and order-one theorem remain blocked.", NextGate: "Derive the physical H_F sub-bimodule as a two-sided Morita representation of C⊕H⊕M3(C), including chirality, hypercharge, and anti-linear conjugation semantics, before re-running order-one."}
}

func FormatInput(g Gate293Input) string {
	return fmt.Sprintf("candidate=%q dim=%d J2=%+d Jgamma=%+d KO6=%t verdict=%s", g.CandidateName, g.DoubledDim, g.J2Sign, g.JGammaSign, g.KOSixLike, g.Verdict)
}
func FormatJSwap(j JSwapAudit) string {
	return fmt.Sprintf("dim=%d J2=%+d Jgamma=%+d KO6=%t residuals=(%.3g,%.3g) verdict=%s", j.DoubledDimension, j.J2Sign, j.JGammaSign, j.KOSixLike, j.ResidualJ2, j.ResidualGamma, j.Verdict)
}
func FormatNaive(n NaiveWeakColorDiagnostic) string {
	return fmt.Sprintf("abZero=%t imageProductNorm=%.6g residual=%.6g explanation=%q verdict=%s", n.AlgebraProductShouldVanish, n.ImageProductNorm, n.MultiplicativityResidual, n.Explanation, n.Verdict)
}
func FormatRepresentation(r RepresentationCandidate) string {
	return fmt.Sprintf("name=%s assoc=%t unital=%t faithful=%t physicalSM=%t residual=%.6g carrier=%q failure=%q verdict=%s", r.Name, r.Associative, r.Unital, r.FaithfulToAllSummands, r.PhysicalSMBimodule, r.Residual, r.CarrierDescription, r.FailureMechanism, r.Verdict)
}
func FormatRepresentations(rs []RepresentationCandidate) string {
	parts := []string{}
	for _, r := range rs {
		parts = append(parts, FormatRepresentation(r))
	}
	return strings.Join(parts, " || ")
}
func FormatOpposite(o OppositeActionAudit) string {
	return fmt.Sprintf("Jswap=%t requiresRep=%t formula=%q physical=%t zeroOrder=%t missing=%s verdict=%s", o.JSwapAvailable, o.RequiresValidRepresentation, o.Formula, o.ConstructedForPhysicalHF, o.ZeroOrderVerified, strings.Join(o.Missing, " | "), o.Verdict)
}
func FormatOrderOne(o OrderOneAudit) string {
	return fmt.Sprintf("zero=%q first=%q physicalDF=%t physicalRep=%t verified=%t constraints=%t verdict=%s", o.ZeroOrderCondition, o.FirstOrderCondition, o.PhysicalDFAvailable, o.PhysicalRepresentationAvailable, o.OrderOneVerified, o.DiracConstraintsDerived, o.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noNaive=%t noHF=%t noOpposite=%t noDF=%t noHiggs=%t noBGap=%t polluted=%t verdict=%s", f.DoesNotPromoteNaiveAction, f.DoesNotInventHF, f.DoesNotInventOpposite, f.DoesNotSelectDF, f.DoesNotUnlockHiggs, f.DoesNotUnlockBGap, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("JswapKO6=%t assocCandidate=%t physicalSM=%t opposite=%t orderOne=%t canonicalDF=%t dynamics=%t firewall=%t next=%q status=%s", s.JSwapKOSix, s.AnyAssociativeCandidate, s.PhysicalSMRepDerived, s.OppositeConstructed, s.OrderOneVerified, s.CanonicalDFSelected, s.DynamicsUnblocked, s.FirewallPreserved, s.NextGate, s.Status)
}
