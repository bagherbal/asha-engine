// Package resolventfieldadjunction implements Gate 280:
// Resolvent Field Adjunction / Contact Projector Construction Audit.
//
// Gate 279 proved that the contact quartic companion module is irreducible
// over Q and therefore has no nontrivial rational commuting idempotent. Gate
// 280 activates a sealed resolvent-adjunction schema: after adjoining a root of
// the irreducible resolvent cubic, the quartic can be conditionally split into
// a 2+2 pair of quadratics, and numerical/auditable projectors can be
// constructed for each possible branch. The result is intentionally bounded:
// the adjunction seal supplies a conditional branch space, not a finite-core
// theorem selecting one branch as physical. Thus the contact projectors are
// constructed branch-by-branch, but the root-to-sector bijection and Gate-275
// r_+/r_- amplitude branch remain unselected unless a future theorem supplies
// a native resolvent-root selector.
package resolventfieldadjunction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactprojectorcompanion"
)

const (
	AuditID = "GATE280-RESOLVENT-FIELD-ADJUNCTION-CONTACT-PROJECTOR-CONSTRUCTION-AUDIT"

	StatusSealActivated         = "CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_SEAL_ACTIVATED"
	StatusBranchesConstructed   = "CONDITIONAL_SUPPORT_RESOLVENT_FIELD_BRANCHES_CONSTRUCTED"
	StatusQuadraticFactors      = "CONDITIONAL_SUPPORT_CONDITIONAL_QUADRATIC_FACTORIZATIONS_CONSTRUCTED"
	StatusProjectorsConstructed = "CONDITIONAL_SUPPORT_CONDITIONAL_CONTACT_PROJECTORS_CONSTRUCTED"
	StatusOrthogonalityVerified = "CONDITIONAL_SUPPORT_PROJECTOR_ORTHOGONALITY_VERIFIED_PER_BRANCH"
	StatusFirewallPreserved     = "CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_FIREWALLS_PRESERVED"

	StatusFailedNoNativeSelector  = "FAILED_ROUTE_NO_NATIVE_RESOLVENT_ROOT_SELECTOR_DERIVED"
	StatusFailedNoSectorBijection = "FAILED_ROUTE_PROJECTORS_NOT_MAPPED_TO_PHYSICAL_SECTORS"
	StatusFailedNoRBranchMap      = "FAILED_ROUTE_RESOLVENT_TO_RPLUS_RMINUS_BRANCH_MAP_MISSING"
	StatusFailedAmplitudeBranch   = "FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED"
	StatusFailedHiggsRatio        = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

// Matrix4 is a small auditable real matrix used only for residual checks of
// algebraic identities. The exact theorem statement is about the field
// extension; floating residuals are diagnostics, not source of truth.
type Matrix4 [4][4]float64

type ResolventAdjunctionSeal struct {
	Name                        string
	Active                      bool
	FieldBefore                 string
	FieldAfterSchema            string
	SelectionStatus             string
	IsSpontaneousBoundary       bool
	GrantsConditionalProjectors bool
	GrantsNativeBranchSelection bool
	Verdict                     string
}

type QuadraticFactor struct {
	Name       string
	Pair       string
	Sum        float64
	Product    float64
	Polynomial string
}

type ProjectorAudit struct {
	Name                          string
	PolynomialCoefficients        []float64 // ascending p0+p1 x+p2 x^2+p3 x^3
	Matrix                        Matrix4
	IdempotentResidual            float64
	CommutesWithCompanionResidual float64
	Trace                         float64
	RankApprox                    float64
	Verdict                       string
}

type BranchAudit struct {
	Name                     string
	Pairing                  string
	ResolventRootZ           float64
	ResolventResidualAbs     float64
	FactorA                  QuadraticFactor
	FactorB                  QuadraticFactor
	FactorizationResidualAbs float64
	ProjectorA               ProjectorAudit
	ProjectorB               ProjectorAudit
	SumToIdentityResidual    float64
	OrthogonalityResidual    float64
	ProjectorsValid          bool
	SectorAssigned           bool
	AssignedSectorA          string
	AssignedSectorB          string
	RBranchSelected          bool
	SelectedRBranch          string
	Verdict                  string
}

type BranchSpaceAudit struct {
	Branches                   []BranchAudit
	BranchCount                int
	AllBranchesProjectorsValid bool
	AnyNativeBranchSelected    bool
	NativeSelectedBranch       string
	ConditionalBranchCount     int
	Verdict                    string
}

type SectorBijectionAudit struct {
	Gate277SectorPairing             string
	SectorPairingSelected            bool
	ConditionalProjectorsExist       bool
	RequiresMappingProjectorToSector bool
	MappingDerived                   bool
	PossibleConditionalMaps          int
	UsesNumericalRootOrdering        bool
	Verdict                          string
}

type RBranchAudit struct {
	RPlus                  float64
	RMinus                 float64
	ResolventToRMapDerived bool
	UniqueAmplitudeBranch  bool
	SelectedBranch         string
	Verdict                string
}

type FirewallAudit struct {
	NoArbitraryResolventRootPromoted      bool
	NoNumericalOrderingPromotion          bool
	NoEmpiricalYukawaInserted             bool
	NoObservedMassesUsed                  bool
	ConditionalAdjunctionNotNativeTheorem bool
	NoProjectorSectorOverclaim            bool
	NoHiggsRatioClaimed                   bool
	FiniteCorePolluted                    bool
	Verdict                               string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                     []FutureCriterion
	NeedNativeResolventSelector  bool
	NeedProjectorSectorSemantics bool
	NeedResolventToRBranchMap    bool
	NeedHeatKernelProjection     bool
	RecommendedNextGate          string
	Verdict                      string
}

type Summary struct {
	SealActivated                    bool
	ConditionalProjectorsConstructed bool
	AllBranchProjectorsValid         bool
	NativeResolventRootSelected      bool
	SectorBijectionDerived           bool
	AmplitudeBranchLocked            bool
	HiggsRatioDerived                bool
	FirewallPreserved                bool
	Status                           string
	NextGate                         string
	Comment                          string
}

type Analysis struct {
	PreviousGate279 contactprojectorcompanion.Analysis
	Seal            ResolventAdjunctionSeal
	BranchSpace     BranchSpaceAudit
	SectorBijection SectorBijectionAudit
	RBranch         RBranchAudit
	Firewall        FirewallAudit
	Future          FutureMap
	Summary         Summary
	TruthStatement  string
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
	prev, err := contactprojectorcompanion.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 279 predecessor: %w", err)
	}
	seal := buildSeal()
	branches := buildBranchSpace()
	sector := auditSectorBijection(branches)
	rbranch := auditRBranch()
	firewall := auditFirewall(seal, branches, sector, rbranch)
	future := buildFuture(sector, rbranch)
	summary := buildSummary(seal, branches, sector, rbranch, firewall)
	return Analysis{
		PreviousGate279: prev,
		Seal:            seal,
		BranchSpace:     branches,
		SectorBijection: sector,
		RBranch:         rbranch,
		Firewall:        firewall,
		Future:          future,
		Summary:         summary,
		TruthStatement:  "Gate 280 activates a ResolventAdjunctionSeal and constructs valid conditional 2+2 contact projectors on every resolvent branch, but no finite-core theorem selects which branch is physical or maps it to the Gate-275 r_+/r_- amplitude branch.",
	}, nil
}

func buildSeal() ResolventAdjunctionSeal {
	return ResolventAdjunctionSeal{
		Name:                        "ResolventAdjunctionSeal",
		Active:                      true,
		FieldBefore:                 "Q",
		FieldAfterSchema:            "Q(z_res), where z_res is one root of 5832000z^3-11566800z^2+7569900z-1637467",
		SelectionStatus:             "seal defines three conditional adjunction branches; no native branch value supplied",
		IsSpontaneousBoundary:       true,
		GrantsConditionalProjectors: true,
		GrantsNativeBranchSelection: false,
		Verdict:                     StatusSealActivated,
	}
}

func buildBranchSpace() BranchSpaceAudit {
	roots := []float64{0.28391219259200619017, 0.44112275728436631761, 0.74409663798084089953, 0.89753507880945325936}
	branchDefs := []struct {
		name string
		a    [2]int
		b    [2]int
	}{
		{"z_high_pairing_q1q2_q3q4", [2]int{0, 1}, [2]int{2, 3}},
		{"z_mid_pairing_q1q3_q2q4", [2]int{0, 2}, [2]int{1, 3}},
		{"z_low_pairing_q1q4_q2q3", [2]int{0, 3}, [2]int{1, 2}},
	}
	branches := make([]BranchAudit, 0, len(branchDefs))
	for _, d := range branchDefs {
		br := constructBranch(d.name, roots, d.a, d.b)
		branches = append(branches, br)
	}
	allValid := true
	for _, b := range branches {
		allValid = allValid && b.ProjectorsValid
	}
	return BranchSpaceAudit{
		Branches:                   branches,
		BranchCount:                len(branches),
		AllBranchesProjectorsValid: allValid,
		AnyNativeBranchSelected:    false,
		NativeSelectedBranch:       "",
		ConditionalBranchCount:     len(branches),
		Verdict:                    StatusProjectorsConstructed,
	}
}

func constructBranch(name string, roots []float64, pairA, pairB [2]int) BranchAudit {
	rA := []float64{roots[pairA[0]], roots[pairA[1]]}
	rB := []float64{roots[pairB[0]], roots[pairB[1]]}
	sumA, prodA := rA[0]+rA[1], rA[0]*rA[1]
	sumB, prodB := rB[0]+rB[1], rB[0]*rB[1]
	z := prodA + prodB
	factorA := QuadraticFactor{Name: "factor_A", Pair: fmt.Sprintf("(q%d,q%d)", pairA[0]+1, pairA[1]+1), Sum: sumA, Product: prodA, Polynomial: fmt.Sprintf("x^2 - %.15gx + %.15g", sumA, prodA)}
	factorB := QuadraticFactor{Name: "factor_B", Pair: fmt.Sprintf("(q%d,q%d)", pairB[0]+1, pairB[1]+1), Sum: sumB, Product: prodB, Polynomial: fmt.Sprintf("x^2 - %.15gx + %.15g", sumB, prodB)}
	coeffA := interpolationProjector(roots, map[int]float64{pairA[0]: 1, pairA[1]: 1, pairB[0]: 0, pairB[1]: 0})
	coeffB := interpolationProjector(roots, map[int]float64{pairA[0]: 0, pairA[1]: 0, pairB[0]: 1, pairB[1]: 1})
	C := companionFloat()
	PA := polyMat(coeffA, C)
	PB := polyMat(coeffB, C)
	pA := auditProjector("P_"+factorA.Pair, coeffA, PA, C)
	pB := auditProjector("P_"+factorB.Pair, coeffB, PB, C)
	sumRes := frob(sub(add(PA, PB), identity4()))
	orthRes := frob(multiply(PA, PB))
	factRes := factorizationResidual(sumA, prodA, sumB, prodB)
	resRes := math.Abs(resolventPoly(z))
	valid := pA.IdempotentResidual < 1e-8 && pB.IdempotentResidual < 1e-8 && pA.CommutesWithCompanionResidual < 1e-8 && pB.CommutesWithCompanionResidual < 1e-8 && sumRes < 1e-8 && orthRes < 1e-8
	return BranchAudit{
		Name:                     name,
		Pairing:                  factorA.Pair + "|" + factorB.Pair,
		ResolventRootZ:           z,
		ResolventResidualAbs:     resRes,
		FactorA:                  factorA,
		FactorB:                  factorB,
		FactorizationResidualAbs: factRes,
		ProjectorA:               pA,
		ProjectorB:               pB,
		SumToIdentityResidual:    sumRes,
		OrthogonalityResidual:    orthRes,
		ProjectorsValid:          valid,
		SectorAssigned:           false,
		AssignedSectorA:          "",
		AssignedSectorB:          "",
		RBranchSelected:          false,
		SelectedRBranch:          "",
		Verdict:                  branchVerdict(valid),
	}
}

func branchVerdict(valid bool) string {
	if valid {
		return StatusOrthogonalityVerified
	}
	return "FAILED_ROUTE_CONDITIONAL_PROJECTOR_RESIDUAL_TOO_LARGE"
}

func auditProjector(name string, coeff []float64, P, C Matrix4) ProjectorAudit {
	idRes := frob(sub(multiply(P, P), P))
	commRes := frob(commutator(P, C))
	tr := trace(P)
	return ProjectorAudit{Name: name, PolynomialCoefficients: coeff, Matrix: P, IdempotentResidual: idRes, CommutesWithCompanionResidual: commRes, Trace: tr, RankApprox: tr, Verdict: branchVerdict(idRes < 1e-8 && commRes < 1e-8)}
}

func auditSectorBijection(branches BranchSpaceAudit) SectorBijectionAudit {
	return SectorBijectionAudit{
		Gate277SectorPairing:             "{u,d}|{e,nu}",
		SectorPairingSelected:            true,
		ConditionalProjectorsExist:       branches.AllBranchesProjectorsValid,
		RequiresMappingProjectorToSector: true,
		MappingDerived:                   false,
		PossibleConditionalMaps:          6, // 3 resolvent branches × 2 projector-label orientations.
		UsesNumericalRootOrdering:        false,
		Verdict:                          StatusFailedNoSectorBijection,
	}
}

func auditRBranch() RBranchAudit {
	return RBranchAudit{
		RPlus:                  (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0,
		RMinus:                 (3591.0 - 136.0*math.Sqrt(123.0)) / 3099.0,
		ResolventToRMapDerived: false,
		UniqueAmplitudeBranch:  false,
		SelectedBranch:         "",
		Verdict:                StatusFailedNoRBranchMap,
	}
}

func auditFirewall(seal ResolventAdjunctionSeal, branches BranchSpaceAudit, sector SectorBijectionAudit, r RBranchAudit) FirewallAudit {
	return FirewallAudit{
		NoArbitraryResolventRootPromoted:      !branches.AnyNativeBranchSelected && !seal.GrantsNativeBranchSelection,
		NoNumericalOrderingPromotion:          !sector.UsesNumericalRootOrdering,
		NoEmpiricalYukawaInserted:             true,
		NoObservedMassesUsed:                  true,
		ConditionalAdjunctionNotNativeTheorem: true,
		NoProjectorSectorOverclaim:            !sector.MappingDerived,
		NoHiggsRatioClaimed:                   !r.UniqueAmplitudeBranch,
		FiniteCorePolluted:                    false,
		Verdict:                               StatusFirewallPreserved,
	}
}

func buildFuture(sector SectorBijectionAudit, r RBranchAudit) FutureMap {
	return FutureMap{
		Criteria: []FutureCriterion{
			{Name: "native resolvent-root selector", Required: true, Satisfied: false, Detail: "a finite theorem must select one of the three z_res roots rather than relying on the adjunction seal alone"},
			{Name: "projector-to-sector semantics", Required: true, Satisfied: sector.MappingDerived, Detail: "the selected projector pair must be mapped to {u,d}|{e,nu} without numerical root ordering"},
			{Name: "resolvent-to-r branch map", Required: true, Satisfied: r.ResolventToRMapDerived, Detail: "the contact resolvent branch must be tied to r_+ or r_- from the Gate-275 scalar-Morita bridge"},
			{Name: "physical J/hypercharge/heat-kernel projection", Required: true, Satisfied: false, Detail: "a physical spectral triple and Seeley-de Witt map are still needed before a Higgs mass ratio"},
		},
		NeedNativeResolventSelector:  true,
		NeedProjectorSectorSemantics: true,
		NeedResolventToRBranchMap:    true,
		NeedHeatKernelProjection:     true,
		RecommendedNextGate:          "Gate 281 — Resolvent Branch Semantics / Projector-to-Sector Orientation Seal Audit",
		Verdict:                      "NEXT_OBLIGATION_NATIVE_RESOLVENT_SELECTOR_OR_EXPLICIT_BRANCH_SEAL_VALUE",
	}
}

func buildSummary(seal ResolventAdjunctionSeal, branches BranchSpaceAudit, sector SectorBijectionAudit, r RBranchAudit, fw FirewallAudit) Summary {
	return Summary{
		SealActivated:                    seal.Active,
		ConditionalProjectorsConstructed: branches.BranchCount == 3 && branches.AllBranchesProjectorsValid,
		AllBranchProjectorsValid:         branches.AllBranchesProjectorsValid,
		NativeResolventRootSelected:      branches.AnyNativeBranchSelected,
		SectorBijectionDerived:           sector.MappingDerived,
		AmplitudeBranchLocked:            r.UniqueAmplitudeBranch,
		HiggsRatioDerived:                false,
		FirewallPreserved:                !fw.FiniteCorePolluted,
		Status:                           StatusFailedNoNativeSelector,
		NextGate:                         "Gate 281 — Resolvent Branch Semantics / Projector-to-Sector Orientation Seal Audit",
		Comment:                          "Resolvent adjunction constructs valid conditional 2+2 companion projectors for all three branches. The adjunction seal does not itself select the physical branch, root-sector bijection, or r_+/r_- amplitude branch.",
	}
}

func companionFloat() Matrix4 {
	return Matrix4{{0, 0, 0, -271.0 / 3240.0}, {1, 0, 0, 149.0 / 216.0}, {0, 1, 0, -119.0 / 60.0}, {0, 0, 1, 71.0 / 30.0}}
}

func resolventPoly(z float64) float64 { return 5832000*z*z*z - 11566800*z*z + 7569900*z - 1637467 }

func factorizationResidual(s1, p1, s2, p2 float64) float64 {
	// Compare coefficients of (x^2-s1 x+p1)(x^2-s2 x+p2) to normalized q4.
	a3 := -(s1 + s2)
	a2 := p1 + p2 + s1*s2
	a1 := -(s1*p2 + s2*p1)
	a0 := p1 * p2
	target := []float64{-71.0 / 30.0, 119.0 / 60.0, -149.0 / 216.0, 271.0 / 3240.0}
	got := []float64{a3, a2, a1, a0}
	var sum float64
	for i := range got {
		d := got[i] - target[i]
		sum += d * d
	}
	return math.Sqrt(sum)
}

func interpolationProjector(roots []float64, values map[int]float64) []float64 {
	// Solve Vandermonde system for degree <=3 polynomial p with p(root_i)=value_i.
	var A [4][5]float64
	for i, r := range roots {
		pow := 1.0
		for j := 0; j < 4; j++ {
			A[i][j] = pow
			pow *= r
		}
		A[i][4] = values[i]
	}
	// Gaussian elimination.
	for col := 0; col < 4; col++ {
		piv := col
		for r := col + 1; r < 4; r++ {
			if math.Abs(A[r][col]) > math.Abs(A[piv][col]) {
				piv = r
			}
		}
		if piv != col {
			A[piv], A[col] = A[col], A[piv]
		}
		div := A[col][col]
		for j := col; j < 5; j++ {
			A[col][j] /= div
		}
		for r := 0; r < 4; r++ {
			if r == col {
				continue
			}
			factor := A[r][col]
			for j := col; j < 5; j++ {
				A[r][j] -= factor * A[col][j]
			}
		}
	}
	return []float64{A[0][4], A[1][4], A[2][4], A[3][4]}
}

func polyMat(coeff []float64, C Matrix4) Matrix4 {
	I := identity4()
	C2 := multiply(C, C)
	C3 := multiply(C2, C)
	mats := []Matrix4{I, C, C2, C3}
	var out Matrix4
	for k, c := range coeff {
		out = add(out, scale(c, mats[k]))
	}
	return out
}

func identity4() Matrix4 {
	var m Matrix4
	for i := 0; i < 4; i++ {
		m[i][i] = 1
	}
	return m
}
func add(a, b Matrix4) Matrix4 {
	var o Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			o[i][j] = a[i][j] + b[i][j]
		}
	}
	return o
}
func sub(a, b Matrix4) Matrix4 {
	var o Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			o[i][j] = a[i][j] - b[i][j]
		}
	}
	return o
}
func scale(s float64, a Matrix4) Matrix4 {
	var o Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			o[i][j] = s * a[i][j]
		}
	}
	return o
}
func multiply(a, b Matrix4) Matrix4 {
	var o Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				o[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return o
}
func commutator(a, b Matrix4) Matrix4 { return sub(multiply(a, b), multiply(b, a)) }
func frob(a Matrix4) float64 {
	var s float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
func trace(a Matrix4) float64 {
	var s float64
	for i := 0; i < 4; i++ {
		s += a[i][i]
	}
	return s
}

func FormatSeal(s ResolventAdjunctionSeal) string {
	return fmt.Sprintf("name=%s active=%t before=%s after=%s selection=%q spontaneous=%t conditionalProjectors=%t nativeSelection=%t verdict=%s", s.Name, s.Active, s.FieldBefore, s.FieldAfterSchema, s.SelectionStatus, s.IsSpontaneousBoundary, s.GrantsConditionalProjectors, s.GrantsNativeBranchSelection, s.Verdict)
}

func FormatQuadratic(q QuadraticFactor) string {
	return fmt.Sprintf("%s pair=%s sum=%.15g product=%.15g poly=%q", q.Name, q.Pair, q.Sum, q.Product, q.Polynomial)
}

func FormatProjector(p ProjectorAudit) string {
	coeffs := make([]string, 0, len(p.PolynomialCoefficients))
	for _, c := range p.PolynomialCoefficients {
		coeffs = append(coeffs, fmt.Sprintf("%.12g", c))
	}
	return fmt.Sprintf("%s coeff=[%s] idRes=%.3g commRes=%.3g trace=%.12g rank≈%.12g verdict=%s", p.Name, strings.Join(coeffs, ","), p.IdempotentResidual, p.CommutesWithCompanionResidual, p.Trace, p.RankApprox, p.Verdict)
}

func FormatBranch(b BranchAudit) string {
	return fmt.Sprintf("%s pairing=%s z=%.15g resZ=%.3g factors={%s | %s} factRes=%.3g PA={%s} PB={%s} sumI=%.3g orth=%.3g valid=%t sectorAssigned=%t rSelected=%t verdict=%s", b.Name, b.Pairing, b.ResolventRootZ, b.ResolventResidualAbs, FormatQuadratic(b.FactorA), FormatQuadratic(b.FactorB), b.FactorizationResidualAbs, FormatProjector(b.ProjectorA), FormatProjector(b.ProjectorB), b.SumToIdentityResidual, b.OrthogonalityResidual, b.ProjectorsValid, b.SectorAssigned, b.RBranchSelected, b.Verdict)
}

func FormatBranchSpace(a BranchSpaceAudit) string {
	parts := make([]string, 0, len(a.Branches))
	for _, b := range a.Branches {
		parts = append(parts, FormatBranch(b))
	}
	return fmt.Sprintf("count=%d allValid=%t nativeSelected=%t selected=%q conditional=%d branches={%s} verdict=%s", a.BranchCount, a.AllBranchesProjectorsValid, a.AnyNativeBranchSelected, a.NativeSelectedBranch, a.ConditionalBranchCount, strings.Join(parts, "; "), a.Verdict)
}

func FormatSector(a SectorBijectionAudit) string {
	return fmt.Sprintf("gate277=%q selected=%t projectors=%t requiresMap=%t mapping=%t maps=%d ordering=%t verdict=%s", a.Gate277SectorPairing, a.SectorPairingSelected, a.ConditionalProjectorsExist, a.RequiresMappingProjectorToSector, a.MappingDerived, a.PossibleConditionalMaps, a.UsesNumericalRootOrdering, a.Verdict)
}

func FormatRBranch(a RBranchAudit) string {
	return fmt.Sprintf("rPlus=%.15g rMinus=%.15g resolventToR=%t unique=%t selected=%q verdict=%s", a.RPlus, a.RMinus, a.ResolventToRMapDerived, a.UniqueAmplitudeBranch, a.SelectedBranch, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("arbitraryRoot=%t numericalOrdering=%t empiricalY=%t masses=%t conditionalNotNative=%t noSectorOverclaim=%t noHiggs=%t polluted=%t verdict=%s", a.NoArbitraryResolventRootPromoted, a.NoNumericalOrderingPromotion, a.NoEmpiricalYukawaInserted, a.NoObservedMassesUsed, a.ConditionalAdjunctionNotNativeTheorem, a.NoProjectorSectorOverclaim, a.NoHiggsRatioClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	parts := make([]string, 0, len(a.Criteria))
	for _, c := range a.Criteria {
		parts = append(parts, fmt.Sprintf("%s[required=%t satisfied=%t detail=%s]", c.Name, c.Required, c.Satisfied, c.Detail))
	}
	return fmt.Sprintf("resolvent=%t semantics=%t rMap=%t heat=%t criteria={%s} next=%s verdict=%s", a.NeedNativeResolventSelector, a.NeedProjectorSectorSemantics, a.NeedResolventToRBranchMap, a.NeedHeatKernelProjection, strings.Join(parts, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("seal=%t projectors=%t allValid=%t nativeRoot=%t sector=%t r=%t higgs=%t firewall=%t next=%s status=%s comment=%q", a.SealActivated, a.ConditionalProjectorsConstructed, a.AllBranchProjectorsValid, a.NativeResolventRootSelected, a.SectorBijectionDerived, a.AmplitudeBranchLocked, a.HiggsRatioDerived, a.FirewallPreserved, a.NextGate, a.Status, a.Comment)
}

func AssertNoOverclaim(a Analysis) error {
	if a.BranchSpace.AnyNativeBranchSelected || a.SectorBijection.MappingDerived || a.RBranch.UniqueAmplitudeBranch || a.Summary.HiggsRatioDerived || a.Firewall.FiniteCorePolluted {
		return fmt.Errorf("Gate 280 overclaimed: summary=%+v", a.Summary)
	}
	return nil
}
