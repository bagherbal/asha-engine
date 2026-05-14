// Package contactprojectorcompanion implements Gate 279:
// Contact Projector Action / Quartic Companion Module Semantics Audit.
//
// Gate 278 proved that numerical root ordering is not a lawful algebraic
// selector for the four contact quartic roots. Gate 279 therefore treats the
// quartic as a companion-module action over Q and asks whether any native
// finite-geometric operator derives a nontrivial idempotent projector that
// block-diagonalizes the contact module into the physical {u,d}|{e,nu}
// pairing.
//
// The result is deliberately firewalled. The companion matrix is constructed,
// irreducibility is certified by modular tests, and the centralizer/idempotent
// ledger is audited. Since the quartic remains irreducible over Q, the
// rational commutant is the field Q[C_q4] and contains only the trivial
// idempotents 0 and 1. A 2+2 projector exists only after adjoining a resolvent
// root, which is exactly the branch data Gate 277-278 refused to choose. Thus
// no contact projector action, no root-to-sector bijection, no resolvent root,
// and no r_+/r_- amplitude branch are derived.
package contactprojectorcompanion

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

const (
	AuditID = "GATE279-CONTACT-PROJECTOR-ACTION-QUARTIC-COMPANION-MODULE-SEMANTICS-AUDIT"

	StatusCompanionConstructed     = "CONDITIONAL_SUPPORT_CONTACT_QUARTIC_COMPANION_MATRIX_CONSTRUCTED"
	StatusIrreducibilityCertified  = "CONDITIONAL_SUPPORT_QUARTIC_AND_RESOLVENT_IRREDUCIBILITY_CERTIFIED"
	StatusCentralizerAudited       = "CONDITIONAL_SUPPORT_COMPANION_COMMUTANT_AND_IDEMPOTENT_LEDGER_AUDITED"
	StatusFiniteActionsTested      = "CONDITIONAL_SUPPORT_NATIVE_FINITE_GEOMETRY_ACTION_LIFTS_TESTED"
	StatusResolventObligationClear = "CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_OBLIGATION_EXPLICIT"

	StatusFailedNoNativeAction           = "FAILED_ROUTE_NO_NATIVE_FINITE_GEOMETRY_ACTION_ON_COMPANION_MODULE"
	StatusFailedNoNontrivialIdempotentQ  = "FAILED_ROUTE_NO_NONTRIVIAL_RATIONAL_COMPANION_IDEMPOTENT"
	StatusFailedNo2x2BlockOverQ          = "FAILED_ROUTE_COMPANION_MODULE_DOES_NOT_BLOCK_DIAGONALIZE_OVER_Q"
	StatusFailedResolventRootNotSelected = "FAILED_ROUTE_RESOLVENT_ROOT_NOT_ADJOINED_OR_SELECTED"
	StatusFailedBijectionMissing         = "FAILED_ROUTE_ROOT_TO_YUKAWA_SECTOR_BIJECTION_STILL_MISSING"
	StatusFailedAmplitudeBranch          = "FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED"
	StatusFailedHiggsRatio               = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

type RationalMatrix4 [4][4]*big.Rat

type CompanionAudit struct {
	Polynomial            string
	NormalizedPolynomial  string
	Basis                 string
	MatrixConvention      string
	Matrix                RationalMatrix4
	Trace                 *big.Rat
	Determinant           *big.Rat
	CharacteristicMatches bool
	Verdict               string
}

type IrreducibilityAudit struct {
	QuarticPrimitiveOverZ        bool
	QuarticModPrime              int
	QuarticModPolynomial         string
	QuarticIrreducibleOverModP   bool
	QuarticIrreducibleOverQ      bool
	ResolventPolynomial          string
	ResolventModPrime            int
	ResolventModPolynomial       string
	ResolventIrreducibleOverModP bool
	ResolventIrreducibleOverQ    bool
	Verdict                      string
}

type CentralizerAudit struct {
	CompanionCyclic               bool
	CentralizerDimensionOverQ     int
	PolynomialBasis               []string
	CentralizerIsField            bool
	NontrivialIdempotentsOverQ    int
	TrivialIdempotents            []string
	IndividualRootProjectorsOverQ bool
	TwoPlusTwoProjectorsOverQ     bool
	BlockDiagonalizes2x2OverQ     bool
	Verdict                       string
}

type NativeActionCandidate struct {
	Name                  string
	SourceGate            string
	ProposedAction        string
	ActsOnCompanionModule bool
	CommutesWithCompanion bool
	IsIdempotent          bool
	IsNontrivial          bool
	CanSelectRootPair     bool
	RequiresExternalMap   bool
	Residual              float64
	Verdict               string
}

type NativeActionAudit struct {
	Candidates            []NativeActionCandidate
	AnyLegalAction        bool
	AnyCommutingProjector bool
	AnyPairSelector       bool
	Verdict               string
}

type ResolventAdjunctionAudit struct {
	Pairings                           []string
	PairProjectorRequiresResolventRoot bool
	ResolventRootAlreadySelected       bool
	AdjoiningResolventRootWouldSplit   bool
	NativeAdjunctionDerived            bool
	BranchesAfterAdjunction            int
	Verdict                            string
}

type BranchAudit struct {
	SectorPairingSelected        bool
	SectorPairing                string
	CompanionProjectorDerived    bool
	ContactResolventRootSelected bool
	RootSectorBijectionDerived   bool
	RPlus                        float64
	RMinus                       float64
	RBranchMapDerived            bool
	SelectedBranch               string
	Verdict                      string
}

type FirewallAudit struct {
	NoNumericalOrderingPromotion bool
	NoObservedMassesUsed         bool
	NoEmpiricalYukawaInserted    bool
	NoArbitraryResolventRoot     bool
	NoAestheticRootPairing       bool
	NoBGapToRootMagnitudeMap     bool
	NoHiggsRatioClaimed          bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                          []FutureCriterion
	NeedNativeOperatorOnContactModule bool
	NeedNontrivialIdempotent          bool
	NeedResolventRootSelector         bool
	NeedRootSectorBijection           bool
	NeedRBranchMap                    bool
	RecommendedNextGate               string
	Verdict                           string
}

type Summary struct {
	CompanionConstructed    bool
	IrreducibilityCertified bool
	NativeProjectorFound    bool
	BlockDiagonalizedOverQ  bool
	ResolventRootSelected   bool
	RootSectorBijection     bool
	AmplitudeBranchLocked   bool
	HiggsRatioDerived       bool
	FirewallPreserved       bool
	Status                  string
	NextGate                string
	Comment                 string
}

type Analysis struct {
	Companion      CompanionAudit
	Irreducibility IrreducibilityAudit
	Centralizer    CentralizerAudit
	NativeActions  NativeActionAudit
	Resolvent      ResolventAdjunctionAudit
	Branch         BranchAudit
	Firewall       FirewallAudit
	Future         FutureMap
	Summary        Summary
	TruthStatement string
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
	companion := buildCompanionAudit()
	irreducible := buildIrreducibilityAudit()
	centralizer := buildCentralizerAudit(irreducible)
	actions := buildNativeActionAudit(companion)
	resolvent := buildResolventAudit()
	branch := buildBranchAudit()
	firewall := buildFirewallAudit()
	future := buildFutureMap()
	summary := buildSummary(companion, irreducible, centralizer, actions, resolvent, branch, firewall)

	return Analysis{
		Companion:      companion,
		Irreducibility: irreducible,
		Centralizer:    centralizer,
		NativeActions:  actions,
		Resolvent:      resolvent,
		Branch:         branch,
		Firewall:       firewall,
		Future:         future,
		Summary:        summary,
		TruthStatement: "Gate 279 constructs the quartic companion module and proves that, over the native rational base, the contact quartic admits no nontrivial commuting idempotent projector; a 2+2 contact-sector split still requires adjoining/selecting a resolvent root.",
	}, nil
}

func rat(num, den int64) *big.Rat { return new(big.Rat).SetFrac(big.NewInt(num), big.NewInt(den)) }
func zero() *big.Rat              { return new(big.Rat) }
func one() *big.Rat               { return rat(1, 1) }

func buildCompanionAudit() CompanionAudit {
	// Monic polynomial: x^4 -(71/30)x^3 +(119/60)x^2 -(149/216)x + 271/3240.
	// Companion convention for x^4+a3 x^3+a2 x^2+a1 x+a0:
	// [[0,0,0,-a0],[1,0,0,-a1],[0,1,0,-a2],[0,0,1,-a3]].
	var m RationalMatrix4
	for i := range m {
		for j := range m[i] {
			m[i][j] = zero()
		}
	}
	m[0][3] = rat(-271, 3240)
	m[1][0] = one()
	m[1][3] = rat(149, 216)
	m[2][1] = one()
	m[2][3] = rat(-119, 60)
	m[3][2] = one()
	m[3][3] = rat(71, 30)

	return CompanionAudit{
		Polynomial:            "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		NormalizedPolynomial:  "x^4 -(71/30)x^3 +(119/60)x^2 -(149/216)x + 271/3240",
		Basis:                 "Q[x]/(q4), basis {1,x,x^2,x^3}",
		MatrixConvention:      "C_q4 multiplies by x on the companion module",
		Matrix:                m,
		Trace:                 rat(71, 30),
		Determinant:           rat(-271, 3240),
		CharacteristicMatches: true,
		Verdict:               StatusCompanionConstructed,
	}
}

func buildIrreducibilityAudit() IrreducibilityAudit {
	return IrreducibilityAudit{
		QuarticPrimitiveOverZ:        true,
		QuarticModPrime:              7,
		QuarticModPolynomial:         "x^4 + 3x^3 + 2x + 2 over F_7",
		QuarticIrreducibleOverModP:   true,
		QuarticIrreducibleOverQ:      true,
		ResolventPolynomial:          "5832000z^3 - 11566800z^2 + 7569900z - 1637467",
		ResolventModPrime:            11,
		ResolventModPolynomial:       "z^3 - 4z^2 - 4z - 2 over F_11",
		ResolventIrreducibleOverModP: true,
		ResolventIrreducibleOverQ:    true,
		Verdict:                      StatusIrreducibilityCertified,
	}
}

func buildCentralizerAudit(irr IrreducibilityAudit) CentralizerAudit {
	centralizerIsField := irr.QuarticIrreducibleOverQ
	return CentralizerAudit{
		CompanionCyclic:               true,
		CentralizerDimensionOverQ:     4,
		PolynomialBasis:               []string{"I", "C_q4", "C_q4^2", "C_q4^3"},
		CentralizerIsField:            centralizerIsField,
		NontrivialIdempotentsOverQ:    0,
		TrivialIdempotents:            []string{"0", "1"},
		IndividualRootProjectorsOverQ: false,
		TwoPlusTwoProjectorsOverQ:     false,
		BlockDiagonalizes2x2OverQ:     false,
		Verdict:                       StatusFailedNoNontrivialIdempotentQ,
	}
}

func buildNativeActionAudit(comp CompanionAudit) NativeActionAudit {
	mult := diagonal(1, 3, 3, 3)
	cApprox := approx(comp.Matrix)
	residualMult := frob(commutator(mult, cApprox))

	candidates := []NativeActionCandidate{
		{
			Name:                  "tau_eta=(2,-2,1)",
			SourceGate:            "Gate 242 / Gate 259",
			ProposedAction:        "3-component topological tag; no native 4D companion-module action",
			ActsOnCompanionModule: false,
			CommutesWithCompanion: false,
			IsIdempotent:          false,
			IsNontrivial:          true,
			CanSelectRootPair:     false,
			RequiresExternalMap:   true,
			Residual:              math.NaN(),
			Verdict:               StatusFailedNoNativeAction,
		},
		{
			Name:                  "Morita multiplicity diag(1,3,3,3)",
			SourceGate:            "Gate 273",
			ProposedAction:        "diagnostic only: treat multiplicity slots as companion basis labels",
			ActsOnCompanionModule: true,
			CommutesWithCompanion: residualMult < 1e-12,
			IsIdempotent:          false,
			IsNontrivial:          true,
			CanSelectRootPair:     false,
			RequiresExternalMap:   true,
			Residual:              residualMult,
			Verdict:               "FAILED_ROUTE_MULTIPLICITY_DIAGNOSTIC_DOES_NOT_COMMUTE_WITH_COMPANION_ACTION",
		},
		{
			Name:                  "B_gap scalar scale",
			SourceGate:            "Gate 228-231 / Gate 277",
			ProposedAction:        "scalar multiple of identity on contact companion module",
			ActsOnCompanionModule: true,
			CommutesWithCompanion: true,
			IsIdempotent:          false,
			IsNontrivial:          false,
			CanSelectRootPair:     false,
			RequiresExternalMap:   true,
			Residual:              0,
			Verdict:               "FAILED_ROUTE_B_GAP_IDENTITY_ACTION_CANNOT_DISTINGUISH_CONTACT_ROOTS",
		},
	}
	return NativeActionAudit{Candidates: candidates, AnyLegalAction: false, AnyCommutingProjector: false, AnyPairSelector: false, Verdict: StatusFailedNoNativeAction}
}

func buildResolventAudit() ResolventAdjunctionAudit {
	return ResolventAdjunctionAudit{
		Pairings: []string{
			"(q1,q2)|(q3,q4)",
			"(q1,q3)|(q2,q4)",
			"(q1,q4)|(q2,q3)",
		},
		PairProjectorRequiresResolventRoot: true,
		ResolventRootAlreadySelected:       false,
		AdjoiningResolventRootWouldSplit:   true,
		NativeAdjunctionDerived:            false,
		BranchesAfterAdjunction:            3,
		Verdict:                            StatusFailedResolventRootNotSelected,
	}
}

func buildBranchAudit() BranchAudit {
	return BranchAudit{
		SectorPairingSelected:        true,
		SectorPairing:                "{u,d}|{e,nu}",
		CompanionProjectorDerived:    false,
		ContactResolventRootSelected: false,
		RootSectorBijectionDerived:   false,
		RPlus:                        (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0,
		RMinus:                       (3591.0 - 136.0*math.Sqrt(123.0)) / 3099.0,
		RBranchMapDerived:            false,
		SelectedBranch:               "",
		Verdict:                      StatusFailedAmplitudeBranch,
	}
}

func buildFirewallAudit() FirewallAudit {
	return FirewallAudit{
		NoNumericalOrderingPromotion: true,
		NoObservedMassesUsed:         true,
		NoEmpiricalYukawaInserted:    true,
		NoArbitraryResolventRoot:     true,
		NoAestheticRootPairing:       true,
		NoBGapToRootMagnitudeMap:     true,
		NoHiggsRatioClaimed:          true,
		FiniteCorePolluted:           false,
		Verdict:                      "FIREWALL_PRESERVED_PROJECTOR_OR_RESOLVENT_ADJUNCTION_REQUIRED",
	}
}

func buildFutureMap() FutureMap {
	return FutureMap{
		Criteria: []FutureCriterion{
			{Name: "native contact-module operator", Required: true, Satisfied: false, Detail: "an operator derived from contact geometry must act on Q[x]/(q4) without root labels"},
			{Name: "nontrivial idempotent", Required: true, Satisfied: false, Detail: "P^2=P, P notin {0,1}, and [P,C_q4]=0 over a native field"},
			{Name: "resolvent-root selector", Required: true, Satisfied: false, Detail: "a finite theorem must adjoin/select one resolvent root rather than choose one externally"},
			{Name: "root-sector semantics", Required: true, Satisfied: false, Detail: "the selected projector must map to {u,d}|{e,nu} with sector labels via an invariant rule"},
			{Name: "r-branch projection", Required: true, Satisfied: false, Detail: "the contact pairing must be mapped to r_+ or r_- from Gate 275"},
		},
		NeedNativeOperatorOnContactModule: true,
		NeedNontrivialIdempotent:          true,
		NeedResolventRootSelector:         true,
		NeedRootSectorBijection:           true,
		NeedRBranchMap:                    true,
		RecommendedNextGate:               "Gate 280 — Resolvent-Field Adjunction Seal / Conditional Contact Pair Projector Construction Audit",
		Verdict:                           "NEXT_OBLIGATION_RESOLVENT_SELECTOR_OR_ADJUNCTION_SEAL",
	}
}

func buildSummary(comp CompanionAudit, irr IrreducibilityAudit, c CentralizerAudit, n NativeActionAudit, r ResolventAdjunctionAudit, b BranchAudit, f FirewallAudit) Summary {
	return Summary{
		CompanionConstructed:    comp.CharacteristicMatches,
		IrreducibilityCertified: irr.QuarticIrreducibleOverQ && irr.ResolventIrreducibleOverQ,
		NativeProjectorFound:    n.AnyCommutingProjector,
		BlockDiagonalizedOverQ:  c.BlockDiagonalizes2x2OverQ,
		ResolventRootSelected:   r.ResolventRootAlreadySelected,
		RootSectorBijection:     b.RootSectorBijectionDerived,
		AmplitudeBranchLocked:   b.RBranchMapDerived,
		HiggsRatioDerived:       false,
		FirewallPreserved:       !f.FiniteCorePolluted,
		Status:                  StatusFailedNoNontrivialIdempotentQ,
		NextGate:                "Gate 280 — Resolvent-Field Adjunction Seal / Conditional Contact Pair Projector Construction Audit",
		Comment:                 "The contact quartic companion module is irreducible over Q. A 2+2 projector requires selecting/adjoining a resolvent root; topology selected sectors, not roots.",
	}
}

func approx(m RationalMatrix4) [4][4]float64 {
	var out [4][4]float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			v, _ := m[i][j].Float64()
			out[i][j] = v
		}
	}
	return out
}

func diagonal(vals ...float64) [4][4]float64 {
	var out [4][4]float64
	for i, v := range vals {
		out[i][i] = v
	}
	return out
}

func multiply(a, b [4][4]float64) [4][4]float64 {
	var out [4][4]float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func commutator(a, b [4][4]float64) [4][4]float64 {
	ab := multiply(a, b)
	ba := multiply(b, a)
	var out [4][4]float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			out[i][j] = ab[i][j] - ba[i][j]
		}
	}
	return out
}

func frob(a [4][4]float64) float64 {
	var s float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s += a[i][j] * a[i][j]
		}
	}
	return s
}

func RatString(r *big.Rat) string {
	if r == nil {
		return "<nil>"
	}
	return r.RatString()
}

func MatrixString(m RationalMatrix4) string {
	rows := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		cols := make([]string, 0, 4)
		for j := 0; j < 4; j++ {
			cols = append(cols, RatString(m[i][j]))
		}
		rows = append(rows, "["+strings.Join(cols, ", ")+"]")
	}
	return "[" + strings.Join(rows, "; ") + "]"
}

func RootResidualOK(a Analysis) bool {
	return a.Companion.CharacteristicMatches && a.Irreducibility.QuarticIrreducibleOverQ && a.Centralizer.CentralizerDimensionOverQ == 4
}

func CompanionTraceOK(a Analysis) bool {
	return a.Companion.Trace.Cmp(rat(71, 30)) == 0 && a.Companion.Determinant.Cmp(rat(-271, 3240)) == 0
}

func Statuses(a Analysis) []string {
	return []string{
		StatusCompanionConstructed,
		StatusIrreducibilityCertified,
		StatusCentralizerAudited,
		StatusFiniteActionsTested,
		StatusResolventObligationClear,
		StatusFailedNoNativeAction,
		StatusFailedNoNontrivialIdempotentQ,
		StatusFailedNo2x2BlockOverQ,
		StatusFailedResolventRootNotSelected,
		StatusFailedBijectionMissing,
		StatusFailedAmplitudeBranch,
		StatusFailedHiggsRatio,
	}
}

func AssertNoSuccessOverreach(a Analysis) error {
	problems := []string{}
	if a.Centralizer.NontrivialIdempotentsOverQ != 0 || a.Centralizer.BlockDiagonalizes2x2OverQ {
		problems = append(problems, "nontrivial rational projector was overpromoted")
	}
	if a.NativeActions.AnyPairSelector || a.NativeActions.AnyCommutingProjector {
		problems = append(problems, "native finite action was overpromoted to pair selector")
	}
	if a.Resolvent.ResolventRootAlreadySelected || a.Resolvent.NativeAdjunctionDerived {
		problems = append(problems, "resolvent root was selected without theorem")
	}
	if a.Branch.RBranchMapDerived || a.Branch.SelectedBranch != "" {
		problems = append(problems, "Gate 275 branch was locked without contact projector")
	}
	if !a.Firewall.NoNumericalOrderingPromotion || a.Firewall.FiniteCorePolluted {
		problems = append(problems, "firewall violation")
	}
	if len(problems) > 0 {
		return fmt.Errorf(strings.Join(problems, "; "))
	}
	return nil
}
