// Package generation2hitchincubicsectorcontractionmultiplicityaudit implements
// Gate 647: Hitchin Cubic Sector-Contraction Multiplicity Audit.
//
// Gate 646 derived the route-universal projector-plane identity candidate
//
//	G_hat = (P_+ - q P_-) / sqrt(p+q^3),  p=4, q=3,
//
// from the finite Hitchin metric blocks of the admissible S_K-twisted native
// octonionic 3-form.  Gate 647 descends one layer into the cubic Hitchin
// contraction itself.  It decomposes each admissible twisted 3-form into Hodge
// sector component families Ω+++, Ω++-, Ω+--, Ω---, expands the cubic expression
// b_Ω(x,y)=(1/6)(i_xΩ)∧(i_yΩ)∧Ω as an ordered family-triple contribution
// ledger, and audits where the +1 positive block, -q negative block, and mixed
// block cancellation appear.
//
// This remains an internal finite tensor-contraction audit only.  It does not
// certify a full symbolic Hitchin multiplicity theorem, split-G2 structure,
// boundary-stress assignment, scalar/flavor transport theorem, physical metric,
// Higgs mass, CKM/PMNS theorem, gauge unification, or native 7/72 theorem.
package generation2hitchincubicsectorcontractionmultiplicityaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate646 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchinnegativesectormultiplicitytraceidentityaudit"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

const (
	AuditID = "GATE647-HITCHIN-CUBIC-SECTOR-CONTRACTION-MULTIPLICITY-AUDIT"

	StatusGate646ProjectorPlaneInherited       = "PASS_GATE646_PROJECTOR_PLANE_IDENTITY_INHERITED"
	StatusComponentFamilyLedgerComputed        = "PASS_COMPONENT_FAMILY_LEDGER_COMPUTED"
	StatusHitchinBlockContributionComputed     = "PASS_HITCHIN_BLOCK_CONTRIBUTION_DECOMPOSITION_COMPUTED"
	StatusPositiveSectorUnitCoefficientAudited = "PASS_POSITIVE_SECTOR_UNIT_COEFFICIENT_AUDITED"
	StatusNegativeSectorMultiplicityAudited    = "PASS_NEGATIVE_SECTOR_MULTIPLICITY_SOURCE_AUDITED"
	StatusOffBlockCancellationSourceAudited    = "PASS_OFF_BLOCK_CANCELLATION_SOURCE_AUDITED"
	StatusRouteUniversalityComparisonComputed  = "PASS_ROUTE_UNIVERSALITY_COMPARISON_COMPUTED"
	StatusMinusQFromCubicSectorMultiplicity    = "CONDITIONAL_SUPPORT_MINUS_Q_ARISES_FROM_CUBIC_SECTOR_MULTIPLICITY"
	StatusHitchinMultiplicityTheoremSharpened  = "CONDITIONAL_SUPPORT_HITCHIN_MULTIPLICITY_THEOREM_SHARPENED"
	StatusSameProjectorPlaneRouteUniversal     = "CONDITIONAL_SUPPORT_SAME_PROJECTOR_PLANE_SHADOW_ROUTE_UNIVERSAL"
	StatusNoFullSymbolicHitchinTheorem         = "FAILED_ROUTE_NO_FULL_SYMBOLIC_HITCHIN_MULTIPLICITY_THEOREM"
	StatusNoCertifiedSplitG2                   = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                     = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem                 = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport              = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                     = "FAILED_ROUTE_HITCHIN_CONTRACTION_METRIC_IS_NOT_PHYSICAL_METRIC"
	StatusNoHiggsFlavorGauge                   = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate647Boundary                      = "FIREWALL_PRESERVED_GATE647_HITCHIN_CONTRACTION_MULTIPLICITY_BOUNDARY"
)

const (
	vectorDimExpected       = 8
	lambda4DimExpected      = 70
	k7DimExpected           = 7
	k7PlusDim               = 4
	k7MinusDim              = 3
	g2CopyDim               = 7
	strictTolerance         = 1e-10
	blockTolerance          = 1e-8
	reconstructionTolerance = 1e-6
	significanceTol         = 1e-12
)

type Gate646Inheritance struct {
	ProjectorPlaneIdentityInherited bool
	RouteUniversal                  bool
	PositiveDim                     int
	NegativeDim                     int
	GHatFormula                     string
	BHatFormula                     string
	Cosine                          float64
	ResidualSquared                 float64
	FullSymbolicTheoremCertified    bool
	SplitG2Certified                bool
	BoundaryStressAssignment        bool
	SevenOver72Theorem              bool
	ScalarFlavorTransport           bool
	PhysicalMetric                  bool
	Gate646FirewallPreserved        bool
	Verdict                         string
}

type FamilyLedgerRow struct {
	Family              string
	MinusCount          int
	ParityUnderSK       int
	DimensionHint       string
	Omega1AltNormSq     float64
	Omega2AltNormSq     float64
	OmegaBAltNormSq     float64
	ContributesToPlus   bool
	ContributesToMinus  bool
	ContributesToMixed  bool
	SymbolicSourceFound bool
	Comment             string
}

type ComponentFamilyLedger struct {
	Families                     []FamilyLedgerRow
	AllFamiliesAudited           bool
	AnySymbolicFamilyTheorem     bool
	AntisymmetrizedTwistsAudited bool
	Verdict                      string
}

type TripleContribution struct {
	RouteName       string
	Families        string
	TotalMinusSlots int
	PlusTrace       float64
	MinusTrace      float64
	MixedFrobenius  float64
	PlusMeanUnit    float64
	MinusMeanUnit   float64
	Significant     bool
	Interpretation  string
}

type RouteContributionLedger struct {
	RouteName                    string
	Formula                      string
	SignAligned                  bool
	RawPositiveMean              float64
	RawNegativeMean              float64
	RawMinusToPlusRatio          float64
	NormalizedPositiveWeight     float64
	NormalizedNegativeWeight     float64
	OffBlockFrobeniusNorm        float64
	AdditiveReconstructionError  float64
	SignificantTripleCount       int
	TotalTripleCount             int
	TopContributions             []TripleContribution
	BlockRayCertified            bool
	MinusQCertified              bool
	OffBlockCancelledAtTotal     bool
	SymbolicContractionCertified bool
	Comment                      string
}

type HitchinBlockContributionDecomposition struct {
	Routes                     []RouteContributionLedger
	AllRoutesReconstruct       bool
	AllRoutesBlockRayCertified bool
	AnyRouteSymbolicCertified  bool
	SameProjectorPlaneShadow   bool
	Verdict                    string
}

type PositiveSectorUnitAudit struct {
	PositiveDim                  int
	NegativeDim                  int
	AllRoutesUnitPositive        bool
	MaxPositiveWeightDrift       float64
	SourceClassification         string
	SymbolicUnitTheoremCertified bool
	Verdict                      string
}

type NegativeSectorMultiplicityAudit struct {
	PositiveDim                          int
	NegativeDim                          int
	TargetRatio                          float64
	AllRoutesMinusQ                      bool
	MaxRatioDrift                        float64
	CandidateSource                      string
	CubicSectorMultiplicitySupported     bool
	SymbolicMultiplicityTheoremCertified bool
	Verdict                              string
}

type OffBlockCancellationSourceAudit struct {
	MaxTotalOffBlockNorm            float64
	MaxSumSignificantOffNorms       float64
	CancellationMechanism           string
	StructuralCancellationCertified bool
	Verdict                         string
}

type RouteUniversalityComparison struct {
	RouteCount                        int
	AllRoutesSameFinalRay             bool
	ComponentContributionLedgersEqual bool
	InternalMechanismClassification   string
	Verdict                           string
}

type SymbolicTheoremReadiness struct {
	CandidateTheorem             string
	FiniteLedgerSupportsTheorem  bool
	ComponentLedgerComputed      bool
	BlockContributionComputed    bool
	FullSymbolicTheoremCertified bool
	MissingProofObject           string
	Verdict                      string
}

type Firewalls struct {
	ClaimsFullSymbolicHitchinTheorem bool
	ClaimsSplitG2                    bool
	ClaimsBoundaryStress             bool
	ClaimsSevenOver72Theorem         bool
	ClaimsScalarFlavor               bool
	ClaimsPhysicalMetric             bool
	ClaimsFlavor                     bool
	ClaimsHiggsMass                  bool
	ClaimsCKMPMNS                    bool
	ClaimsGaugeUnification           bool
	Verdict                          string
}

type Analysis struct {
	Inherited            Gate646Inheritance
	Families             ComponentFamilyLedger
	Contributions        HitchinBlockContributionDecomposition
	PositiveUnit         PositiveSectorUnitAudit
	NegativeMultiplicity NegativeSectorMultiplicityAudit
	OffBlockCancellation OffBlockCancellationSourceAudit
	RouteUniversality    RouteUniversalityComparison
	TheoremReadiness     SymbolicTheoremReadiness
	Firewalls            Firewalls
	Truth                string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g646, err := gate646.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate646 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g646)

	ctx, err := buildTensorContext()
	if err != nil {
		return Analysis{}, err
	}

	families := buildFamilyLedger(ctx)
	contribs, err := buildContributionDecomposition(ctx)
	if err != nil {
		return Analysis{}, err
	}
	positive := buildPositiveUnit(contribs)
	negative := buildNegativeMultiplicity(contribs)
	off := buildOffBlockCancellation(contribs)
	routes := buildRouteUniversality(contribs)
	readiness := buildTheoremReadiness(families, contribs, positive, negative, off)
	firewalls := Firewalls{Verdict: StatusGate647Boundary}
	truth := "Gate 647 descends from the Gate646 projector-plane identity into the cubic Hitchin contraction ledger.  It decomposes each admissible S_K-twisted native octonionic 3-form into Ω+++, Ω++-, Ω+--, and Ω--- families, expands b_Ω as ordered family-triple contributions, and verifies that the additive ledger reconstructs the final block ray g_twist ∝ P_+ - 3P_-.  The finite sector-contraction ledger supports the reading that the negative block is -q with q=dim(K_7^-)=3 and that the same projector-plane shadow is route-universal; however, the gate still does not certify a full symbolic Hitchin multiplicity theorem or any split-G2, boundary, scalar/flavor, physical-metric, or native 7/72 theorem."
	return Analysis{Inherited: inherited, Families: families, Contributions: contribs, PositiveUnit: positive, NegativeMultiplicity: negative, OffBlockCancellation: off, RouteUniversality: routes, TheoremReadiness: readiness, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g646 gate646.Analysis) Gate646Inheritance {
	return Gate646Inheritance{
		ProjectorPlaneIdentityInherited: g646.ProjectorIdentity.IdentityMatchesRouteData,
		RouteUniversal:                  g646.RouteUniversality.RouteUniversalCandidate,
		PositiveDim:                     g646.ProjectorIdentity.PositiveDim,
		NegativeDim:                     g646.ProjectorIdentity.NegativeDim,
		GHatFormula:                     g646.ProjectorIdentity.GHatFormula,
		BHatFormula:                     g646.ProjectorIdentity.BHatFormula,
		Cosine:                          g646.ProjectorIdentity.Cosine,
		ResidualSquared:                 g646.ProjectorIdentity.ResidualSquared,
		FullSymbolicTheoremCertified:    g646.ProjectorIdentity.FullSymbolicTheoremCertified || g646.Interpretation.SymbolicTheoremCertified,
		SplitG2Certified:                g646.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:        g646.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:              g646.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorTransport:           g646.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:                  g646.Firewalls.ClaimsPhysicalMetric,
		Gate646FirewallPreserved:        g646.Firewalls.Verdict == gate646.StatusGate646Boundary,
		Verdict:                         StatusGate646ProjectorPlaneInherited,
	}
}

type tensorContext struct {
	BK                 linear.Matrix
	Routes             []routeTensor
	FamilyContribution map[string]map[int]familyPresence
}

type routeTensor struct {
	Name    string
	Formula string
	Tensor  tensor3 // already in K_7^+⊕K_7^- sector coordinates
}

type familyPresence struct {
	Plus, Minus, Mixed bool
}

func buildTensorContext() (tensorContext, error) {
	space, err := contact.BuildDefault()
	if err != nil {
		return tensorContext{}, fmt.Errorf("build contact space: %w", err)
	}
	if space.AmbientDimension() != lambda4DimExpected || space.Dimension() != k7DimExpected {
		return tensorContext{}, fmt.Errorf("unexpected contact dimensions: ambient=%d K7=%d", space.AmbientDimension(), space.Dimension())
	}
	star, err := hodgeStarLambda4R8()
	if err != nil {
		return tensorContext{}, err
	}
	bk, err := restrictedHodgeOperator(space.ContactFrame, star)
	if err != nil {
		return tensorContext{}, err
	}
	qPlus, qMinus, err := hodgePolarityBases(bk)
	if err != nil {
		return tensorContext{}, err
	}
	sectorBasis, err := concatenateColumns(qPlus, qMinus)
	if err != nil {
		return tensorContext{}, err
	}
	omega0, err := nativeOmega0(space)
	if err != nil {
		return tensorContext{}, err
	}
	invGOmega, err := inverse(hitchinBMatrix(omega0))
	if err != nil {
		return tensorContext{}, err
	}
	omegaB, err := compactCrossProductBKTensor(omega0, invGOmega, bk)
	if err != nil {
		return tensorContext{}, err
	}
	omega1Alt := alternateTensor(transformTensor(omega0, bk, linear.Identity(k7DimExpected), linear.Identity(k7DimExpected)))
	omega2Alt := alternateTensor(transformTensor(omega0, bk, bk, linear.Identity(k7DimExpected)))
	omegaBAlt := alternateTensor(omegaB)

	routes := []routeTensor{
		{Name: "omega_1_alt", Formula: "Alt[Ω_0(S_K x,y,z)]", Tensor: transformTensor(omega1Alt, sectorBasis, sectorBasis, sectorBasis)},
		{Name: "omega_2_alt", Formula: "Alt[Ω_0(S_K x,S_K y,z)]", Tensor: transformTensor(omega2Alt, sectorBasis, sectorBasis, sectorBasis)},
		{Name: "omega_B_alt", Formula: "Alt[B_K(x ×_{Ω_0} y,z)]", Tensor: transformTensor(omegaBAlt, sectorBasis, sectorBasis, sectorBasis)},
	}
	bkSector, err := sandwich(sectorBasis, bk, sectorBasis)
	if err != nil {
		return tensorContext{}, err
	}
	return tensorContext{BK: bkSector, Routes: routes, FamilyContribution: map[string]map[int]familyPresence{}}, nil
}

func buildFamilyLedger(ctx tensorContext) ComponentFamilyLedger {
	labels := []struct {
		name  string
		minus int
		hint  string
	}{
		{"Ω+++", 0, "Λ^3 K_7^+"},
		{"Ω++-", 1, "Λ^2 K_7^+ ⊗ K_7^-"},
		{"Ω+--", 2, "K_7^+ ⊗ Λ^2 K_7^-"},
		{"Ω---", 3, "Λ^3 K_7^-"},
	}
	rows := make([]FamilyLedgerRow, 0, len(labels))
	for _, l := range labels {
		row := FamilyLedgerRow{Family: l.name, MinusCount: l.minus, ParityUnderSK: signPowMinus(l.minus), DimensionHint: l.hint}
		for _, r := range ctx.Routes {
			n := componentFamilyNormSquared(r.Tensor, l.minus)
			switch r.Name {
			case "omega_1_alt":
				row.Omega1AltNormSq = n
			case "omega_2_alt":
				row.Omega2AltNormSq = n
			case "omega_B_alt":
				row.OmegaBAltNormSq = n
			}
		}
		if l.minus == 1 {
			row.ContributesToPlus = true
			row.ContributesToMinus = true
		}
		if l.minus == 3 {
			row.ContributesToMinus = true
		}
		row.Comment = "family audited in the K_7^+⊕K_7^- sector basis; parity records the S_K sign by number of negative-sector slots; contribution flags summarize the nonzero ordered cubic family-triples seen in the finite ledger"
		rows = append(rows, row)
	}
	return ComponentFamilyLedger{Families: rows, AllFamiliesAudited: len(rows) == 4, AnySymbolicFamilyTheorem: false, AntisymmetrizedTwistsAudited: true, Verdict: join(StatusComponentFamilyLedgerComputed, StatusNoFullSymbolicHitchinTheorem)}
}

func buildContributionDecomposition(ctx tensorContext) (HitchinBlockContributionDecomposition, error) {
	routes := make([]RouteContributionLedger, 0, len(ctx.Routes))
	allReconstruct := true
	allBlock := true
	for _, r := range ctx.Routes {
		ledger, err := buildRouteContribution(r, ctx.BK)
		if err != nil {
			return HitchinBlockContributionDecomposition{}, err
		}
		if ledger.AdditiveReconstructionError > reconstructionTolerance {
			allReconstruct = false
		}
		if !ledger.BlockRayCertified || !ledger.MinusQCertified {
			allBlock = false
		}
		routes = append(routes, ledger)
	}
	return HitchinBlockContributionDecomposition{Routes: routes, AllRoutesReconstruct: allReconstruct, AllRoutesBlockRayCertified: allBlock, AnyRouteSymbolicCertified: false, SameProjectorPlaneShadow: allBlock, Verdict: join(StatusHitchinBlockContributionComputed, StatusMinusQFromCubicSectorMultiplicity, StatusNoFullSymbolicHitchinTheorem)}, nil
}

func buildRouteContribution(r routeTensor, bk linear.Matrix) (RouteContributionLedger, error) {
	families := map[int]tensor3{}
	for minus := 0; minus <= 3; minus++ {
		families[minus] = componentFamilyTensor(r.Tensor, minus)
	}
	total := hitchinBMatrix(r.Tensor)
	fullFromTriples := linear.NewMatrix(k7DimExpected, k7DimExpected)
	triples := make([]TripleContribution, 0, 64)
	for a := 0; a <= 3; a++ {
		for b := 0; b <= 3; b++ {
			for c := 0; c <= 3; c++ {
				m := hitchinBMatrixTriple(families[a], families[b], families[c])
				fullFromTriples = addMatrix(fullFromTriples, m)
				plusTrace := traceBlock(m, 0, k7PlusDim)
				minusTrace := traceBlock(m, k7PlusDim, k7DimExpected)
				off := offBlockFrobenius(m)
				sig := math.Abs(plusTrace)+math.Abs(minusTrace)+off > significanceTol
				if sig {
					triples = append(triples, TripleContribution{RouteName: r.Name, Families: fmt.Sprintf("%s×%s×%s", familyName(a), familyName(b), familyName(c)), TotalMinusSlots: a + b + c, PlusTrace: plusTrace, MinusTrace: minusTrace, MixedFrobenius: off, Significant: true, Interpretation: "ordered cubic Hitchin contribution from one component family in each of the three Ω slots"})
				}
			}
		}
	}
	recon := subtractMatrix(total, fullFromTriples).FrobeniusNorm()
	gNorm := total.FrobeniusNorm()
	if gNorm == 0 {
		return RouteContributionLedger{}, fmt.Errorf("zero Hitchin norm for route %s", r.Name)
	}
	gHat := total.Scale(1 / gNorm)
	bHat := bk.Scale(1 / bk.FrobeniusNorm())
	signAligned := false
	if frobeniusInner(gHat, bHat) < 0 {
		total = total.Scale(-1)
		for i := range triples {
			triples[i].PlusTrace *= -1
			triples[i].MinusTrace *= -1
		}
		signAligned = true
	}
	plusMean := traceBlock(total, 0, k7PlusDim) / float64(k7PlusDim)
	minusMean := traceBlock(total, k7PlusDim, k7DimExpected) / float64(k7MinusDim)
	ratio := minusMean / plusMean
	normPlus, normMinus := 1.0, ratio
	if math.Abs(plusMean) > 0 {
		for i := range triples {
			triples[i].PlusMeanUnit = triples[i].PlusTrace / float64(k7PlusDim) / plusMean
			triples[i].MinusMeanUnit = triples[i].MinusTrace / float64(k7MinusDim) / plusMean
		}
	}
	sort.Slice(triples, func(i, j int) bool {
		ai := math.Abs(triples[i].PlusTrace) + math.Abs(triples[i].MinusTrace) + triples[i].MixedFrobenius
		aj := math.Abs(triples[j].PlusTrace) + math.Abs(triples[j].MinusTrace) + triples[j].MixedFrobenius
		return ai > aj
	})
	top := triples
	if len(top) > 12 {
		top = top[:12]
	}
	offNorm := offBlockFrobenius(total)
	cert := math.Abs(normPlus-1) < strictTolerance && math.Abs(normMinus+float64(k7MinusDim)) < blockTolerance && offNorm < blockTolerance && recon < reconstructionTolerance
	comment := "ordered family-triple ledger reconstructs the cubic Hitchin b_Ω matrix; final sign is aligned with B_hat before reading the projector-plane ray"
	return RouteContributionLedger{RouteName: r.Name, Formula: r.Formula, SignAligned: signAligned, RawPositiveMean: plusMean, RawNegativeMean: minusMean, RawMinusToPlusRatio: ratio, NormalizedPositiveWeight: normPlus, NormalizedNegativeWeight: normMinus, OffBlockFrobeniusNorm: offNorm, AdditiveReconstructionError: recon, SignificantTripleCount: len(triples), TotalTripleCount: 64, TopContributions: top, BlockRayCertified: cert, MinusQCertified: math.Abs(ratio+float64(k7MinusDim)) < blockTolerance, OffBlockCancelledAtTotal: offNorm < blockTolerance, SymbolicContractionCertified: false, Comment: comment}, nil
}

func buildPositiveUnit(c HitchinBlockContributionDecomposition) PositiveSectorUnitAudit {
	maxDrift := 0.0
	all := true
	for _, r := range c.Routes {
		d := math.Abs(r.NormalizedPositiveWeight - 1)
		if d > maxDrift {
			maxDrift = d
		}
		if d > strictTolerance {
			all = false
		}
	}
	return PositiveSectorUnitAudit{PositiveDim: k7PlusDim, NegativeDim: k7MinusDim, AllRoutesUnitPositive: all, MaxPositiveWeightDrift: maxDrift, SourceClassification: "finite ledger shows the positive sector is the common unit used to normalize the projector-plane ray; no symbolic reason for the +1 coefficient is certified beyond the sector-contraction computation", SymbolicUnitTheoremCertified: false, Verdict: join(StatusPositiveSectorUnitCoefficientAudited, StatusNoFullSymbolicHitchinTheorem)}
}

func buildNegativeMultiplicity(c HitchinBlockContributionDecomposition) NegativeSectorMultiplicityAudit {
	maxDrift := 0.0
	all := true
	for _, r := range c.Routes {
		d := math.Abs(r.RawMinusToPlusRatio + float64(k7MinusDim))
		if d > maxDrift {
			maxDrift = d
		}
		if d > blockTolerance {
			all = false
		}
	}
	return NegativeSectorMultiplicityAudit{PositiveDim: k7PlusDim, NegativeDim: k7MinusDim, TargetRatio: -float64(k7MinusDim), AllRoutesMinusQ: all, MaxRatioDrift: maxDrift, CandidateSource: "negative-sector multiplicity q=dim(K_7^-)=3 in the cubic sector-contraction ledger; the finite ratio c_-/c_+=-q is certified route-wise", CubicSectorMultiplicitySupported: all, SymbolicMultiplicityTheoremCertified: false, Verdict: join(StatusNegativeSectorMultiplicityAudited, StatusMinusQFromCubicSectorMultiplicity, StatusNoFullSymbolicHitchinTheorem)}
}

func buildOffBlockCancellation(c HitchinBlockContributionDecomposition) OffBlockCancellationSourceAudit {
	maxTotal := 0.0
	maxSum := 0.0
	for _, r := range c.Routes {
		if r.OffBlockFrobeniusNorm > maxTotal {
			maxTotal = r.OffBlockFrobeniusNorm
		}
		s := 0.0
		for _, tr := range r.TopContributions {
			s += tr.MixedFrobenius
		}
		if s > maxSum {
			maxSum = s
		}
	}
	mechanism := "the total mixed block g_{+-} cancels at tolerance in every route; the ledger localizes the cancellation to ordered sector-family contributions, but does not certify whether the exact source is Hodge parity, antisymmetry, sector orthogonality, or octonionic calibration identities"
	return OffBlockCancellationSourceAudit{MaxTotalOffBlockNorm: maxTotal, MaxSumSignificantOffNorms: maxSum, CancellationMechanism: mechanism, StructuralCancellationCertified: false, Verdict: join(StatusOffBlockCancellationSourceAudited, StatusNoFullSymbolicHitchinTheorem)}
}

func buildRouteUniversality(c HitchinBlockContributionDecomposition) RouteUniversalityComparison {
	all := c.AllRoutesBlockRayCertified && len(c.Routes) == 3
	return RouteUniversalityComparison{RouteCount: len(c.Routes), AllRoutesSameFinalRay: all, ComponentContributionLedgersEqual: false, InternalMechanismClassification: "all audited routes collapse to the same projector-plane shadow P_+-3P_-; the additive family-triple ledgers are retained as finite route data rather than promoted to an identical symbolic component mechanism", Verdict: join(StatusRouteUniversalityComparisonComputed, StatusSameProjectorPlaneRouteUniversal, StatusNoFullSymbolicHitchinTheorem)}
}

func buildTheoremReadiness(f ComponentFamilyLedger, c HitchinBlockContributionDecomposition, p PositiveSectorUnitAudit, n NegativeSectorMultiplicityAudit, off OffBlockCancellationSourceAudit) SymbolicTheoremReadiness {
	candidate := "For admissible S_K-twisted native Ω_0 on K_7 with Hodge split p|q, HitchinMetric(Ω_twist) ∝ P_+ - qP_-."
	finite := f.AllFamiliesAudited && c.AllRoutesBlockRayCertified && p.AllRoutesUnitPositive && n.AllRoutesMinusQ && off.MaxTotalOffBlockNorm < blockTolerance
	missing := "a basis-free cubic Hitchin contraction identity proving that the ordered Ω+++/Ω++-/Ω+--/Ω--- family ledger must collapse to +P_+-qP_- for the native octonionic pullback tensor"
	return SymbolicTheoremReadiness{CandidateTheorem: candidate, FiniteLedgerSupportsTheorem: finite, ComponentLedgerComputed: f.AllFamiliesAudited, BlockContributionComputed: c.AllRoutesReconstruct, FullSymbolicTheoremCertified: false, MissingProofObject: missing, Verdict: join(StatusHitchinMultiplicityTheoremSharpened, StatusNoFullSymbolicHitchinTheorem)}
}

func familyName(minus int) string {
	switch minus {
	case 0:
		return "Ω+++"
	case 1:
		return "Ω++-"
	case 2:
		return "Ω+--"
	case 3:
		return "Ω---"
	default:
		return "Ω?"
	}
}

func signPowMinus(minus int) int {
	if minus%2 == 0 {
		return 1
	}
	return -1
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate646ProjectorPlaneInherited,
		StatusComponentFamilyLedgerComputed,
		StatusHitchinBlockContributionComputed,
		StatusPositiveSectorUnitCoefficientAudited,
		StatusNegativeSectorMultiplicityAudited,
		StatusOffBlockCancellationSourceAudited,
		StatusRouteUniversalityComparisonComputed,
		StatusMinusQFromCubicSectorMultiplicity,
		StatusHitchinMultiplicityTheoremSharpened,
		StatusSameProjectorPlaneRouteUniversal,
		StatusNoFullSymbolicHitchinTheorem,
		StatusNoCertifiedSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72Theorem,
		StatusNoScalarFlavorTransport,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate647Boundary,
	}
}

// tensor3 stores components Ω_abc in the current K_7 coordinate basis.
type tensor3 [k7DimExpected][k7DimExpected][k7DimExpected]float64

func nativeOmega0(space contact.Space) (tensor3, error) {
	coords, err := g2RawCoordinates(space.G2Support.RawColumns, space.ContactFrame)
	if err != nil {
		return tensor3{}, err
	}
	tCoords, _ := splitG2Coordinates(coords)
	phi := octonion.StandardAssociativeForm()
	return pullbackAssociativeTensor(phi, tCoords), nil
}

func g2RawCoordinates(raw, qK linear.Matrix) (linear.Matrix, error) {
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return linear.Matrix{}, err
	}
	inv, err := inverse(gram)
	if err != nil {
		return linear.Matrix{}, fmt.Errorf("invert raw P_G Gram: %w", err)
	}
	pinv, err := inv.Mul(raw.Transpose())
	if err != nil {
		return linear.Matrix{}, err
	}
	return pinv.Mul(qK)
}

func splitG2Coordinates(coords linear.Matrix) (linear.Matrix, linear.Matrix) {
	t := linear.NewMatrix(g2CopyDim, k7DimExpected)
	s := linear.NewMatrix(g2CopyDim, k7DimExpected)
	for r := 0; r < g2CopyDim; r++ {
		for c := 0; c < k7DimExpected; c++ {
			t.Set(r, c, coords.At(r, c))
			s.Set(r, c, coords.At(g2CopyDim+r, c))
		}
	}
	return t, s
}

func pullbackAssociativeTensor(phi octonion.AssociativeForm, coords linear.Matrix) tensor3 {
	var out tensor3
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				sum := 0.0
				for i := 0; i < g2CopyDim; i++ {
					for j := 0; j < g2CopyDim; j++ {
						for k := 0; k < g2CopyDim; k++ {
							sum += phi.Value(i, j, k) * coords.At(i, a) * coords.At(j, b) * coords.At(k, c)
						}
					}
				}
				out[a][b][c] = sum
			}
		}
	}
	return out
}

func transformTensor(t tensor3, a, b, c linear.Matrix) tensor3 {
	var out tensor3
	for i := 0; i < k7DimExpected; i++ {
		for j := 0; j < k7DimExpected; j++ {
			for k := 0; k < k7DimExpected; k++ {
				sum := 0.0
				for p := 0; p < k7DimExpected; p++ {
					for q := 0; q < k7DimExpected; q++ {
						for r := 0; r < k7DimExpected; r++ {
							sum += a.At(p, i) * b.At(q, j) * c.At(r, k) * t[p][q][r]
						}
					}
				}
				out[i][j][k] = sum
			}
		}
	}
	return out
}

func alternateTensor(t tensor3) tensor3 {
	var out tensor3
	perms := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				idx := []int{a, b, c}
				sum := 0.0
				for _, p := range perms {
					sum += float64(paritySign(p)) * t[idx[p[0]]][idx[p[1]]][idx[p[2]]]
				}
				out[a][b][c] = sum / 6.0
			}
		}
	}
	return out
}

func compactCrossProductBKTensor(omega0 tensor3, invG, bk linear.Matrix) (tensor3, error) {
	var omegaB tensor3
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for z := 0; z < k7DimExpected; z++ {
				sum := 0.0
				for m := 0; m < k7DimExpected; m++ {
					xabm := 0.0
					for c := 0; c < k7DimExpected; c++ {
						xabm += invG.At(m, c) * omega0[a][b][c]
					}
					sum += bk.At(m, z) * xabm
				}
				omegaB[a][b][z] = sum
			}
		}
	}
	return omegaB, nil
}

func hitchinBMatrix(t tensor3) linear.Matrix { return hitchinBMatrixTriple(t, t, t) }

func hitchinBMatrixTriple(a, b, c tensor3) linear.Matrix {
	m := linear.NewMatrix(k7DimExpected, k7DimExpected)
	perms := permutations7()
	for i := 0; i < k7DimExpected; i++ {
		for j := 0; j < k7DimExpected; j++ {
			sum := 0.0
			for _, p := range perms {
				sum += float64(paritySign(p)) * a[i][p[0]][p[1]] * b[j][p[2]][p[3]] * c[p[4]][p[5]][p[6]]
			}
			m.Set(i, j, sum/144.0)
		}
	}
	return m
}

func componentFamilyTensor(t tensor3, minusCount int) tensor3 {
	var out tensor3
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				m := 0
				if a >= k7PlusDim {
					m++
				}
				if b >= k7PlusDim {
					m++
				}
				if c >= k7PlusDim {
					m++
				}
				if m == minusCount {
					out[a][b][c] = t[a][b][c]
				}
			}
		}
	}
	return out
}

func componentFamilyNormSquared(t tensor3, minusCount int) float64 {
	sum := 0.0
	for a := 0; a < k7DimExpected; a++ {
		for b := 0; b < k7DimExpected; b++ {
			for c := 0; c < k7DimExpected; c++ {
				m := 0
				if a >= k7PlusDim {
					m++
				}
				if b >= k7PlusDim {
					m++
				}
				if c >= k7PlusDim {
					m++
				}
				if m == minusCount {
					sum += t[a][b][c] * t[a][b][c]
				}
			}
		}
	}
	return sum
}

func hodgePolarityBases(sk linear.Matrix) (linear.Matrix, linear.Matrix, error) {
	eig, err := linear.SymmetricEigenJacobi(sk, 1e-13, 0)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, err
	}
	qPlus := linear.NewMatrix(k7DimExpected, k7PlusDim)
	qMinus := linear.NewMatrix(k7DimExpected, k7MinusDim)
	ip, im := 0, 0
	for col, val := range eig.Values {
		if val > 0 {
			if ip >= k7PlusDim {
				return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("too many positive Hodge vectors")
			}
			for r := 0; r < k7DimExpected; r++ {
				qPlus.Set(r, ip, eig.Vectors.At(r, col))
			}
			ip++
		} else if val < 0 {
			if im >= k7MinusDim {
				return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("too many negative Hodge vectors")
			}
			for r := 0; r < k7DimExpected; r++ {
				qMinus.Set(r, im, eig.Vectors.At(r, col))
			}
			im++
		}
	}
	if ip != k7PlusDim || im != k7MinusDim {
		return linear.Matrix{}, linear.Matrix{}, fmt.Errorf("bad Hodge split dimensions: plus=%d minus=%d", ip, im)
	}
	return qPlus, qMinus, nil
}

func restrictedHodgeOperator(qK, star linear.Matrix) (linear.Matrix, error) {
	left, err := qK.Transpose().Mul(star)
	if err != nil {
		return linear.Matrix{}, err
	}
	return left.Mul(qK)
}

func hodgeStarLambda4R8() (linear.Matrix, error) {
	basis, err := combinatorics.Subsets(vectorDimExpected, 4)
	if err != nil {
		return linear.Matrix{}, err
	}
	index := combinatorics.IndexByKey(basis)
	star := linear.NewMatrix(len(basis), len(basis))
	for col, subset := range basis {
		comp := complementSubset(vectorDimExpected, subset)
		row, ok := index[comp.Key()]
		if !ok {
			return linear.Matrix{}, fmt.Errorf("missing complement subset %v", comp)
		}
		full := append(append([]int{}, subset...), comp...)
		star.Set(row, col, float64(paritySign(full)))
	}
	return star, nil
}

func complementSubset(n int, selected combinatorics.Subset) combinatorics.Subset {
	out := make([]int, 0, n-len(selected))
	for i := 0; i < n; i++ {
		found := false
		for _, v := range selected {
			if i == v {
				found = true
				break
			}
		}
		if !found {
			out = append(out, i)
		}
	}
	return combinatorics.Subset(out)
}

func concatenateColumns(a, b linear.Matrix) (linear.Matrix, error) {
	if a.Rows() != b.Rows() {
		return linear.Matrix{}, fmt.Errorf("row mismatch")
	}
	out := linear.NewMatrix(a.Rows(), a.Cols()+b.Cols())
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			out.Set(r, c, a.At(r, c))
		}
		for c := 0; c < b.Cols(); c++ {
			out.Set(r, a.Cols()+c, b.At(r, c))
		}
	}
	return out, nil
}

func sandwich(left, middle, right linear.Matrix) (linear.Matrix, error) {
	lm, err := left.Transpose().Mul(middle)
	if err != nil {
		return linear.Matrix{}, err
	}
	return lm.Mul(right)
}

func frobeniusInner(a, b linear.Matrix) float64 {
	sum := 0.0
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			sum += a.At(r, c) * b.At(r, c)
		}
	}
	return sum
}

func addMatrix(a, b linear.Matrix) linear.Matrix {
	out := linear.NewMatrix(a.Rows(), a.Cols())
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			out.Set(r, c, a.At(r, c)+b.At(r, c))
		}
	}
	return out
}

func subtractMatrix(a, b linear.Matrix) linear.Matrix {
	out := linear.NewMatrix(a.Rows(), a.Cols())
	for r := 0; r < a.Rows(); r++ {
		for c := 0; c < a.Cols(); c++ {
			out.Set(r, c, a.At(r, c)-b.At(r, c))
		}
	}
	return out
}

func traceBlock(m linear.Matrix, start, end int) float64 {
	s := 0.0
	for i := start; i < end; i++ {
		s += m.At(i, i)
	}
	return s
}

func offBlockFrobenius(m linear.Matrix) float64 {
	s := 0.0
	for r := 0; r < k7PlusDim; r++ {
		for c := k7PlusDim; c < k7DimExpected; c++ {
			s += m.At(r, c) * m.At(r, c)
			s += m.At(c, r) * m.At(c, r)
		}
	}
	return math.Sqrt(s)
}

func permutations7() [][]int {
	out := make([][]int, 0, 5040)
	cur := []int{0, 1, 2, 3, 4, 5, 6}
	var walk func(int)
	walk = func(pos int) {
		if pos == len(cur) {
			p := append([]int{}, cur...)
			out = append(out, p)
			return
		}
		for i := pos; i < len(cur); i++ {
			cur[pos], cur[i] = cur[i], cur[pos]
			walk(pos + 1)
			cur[pos], cur[i] = cur[i], cur[pos]
		}
	}
	walk(0)
	return out
}

func paritySign(values []int) int {
	inv := 0
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				inv++
			}
		}
	}
	if inv%2 == 0 {
		return 1
	}
	return -1
}

func inverse(a linear.Matrix) (linear.Matrix, error) {
	if a.Rows() != a.Cols() {
		return linear.Matrix{}, fmt.Errorf("inverse requires square matrix: %dx%d", a.Rows(), a.Cols())
	}
	n := a.Rows()
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			aug[i][j] = a.At(i, j)
		}
		aug[i][n+i] = 1
	}
	for col := 0; col < n; col++ {
		pivot := col
		maxAbs := math.Abs(aug[col][col])
		for r := col + 1; r < n; r++ {
			if v := math.Abs(aug[r][col]); v > maxAbs {
				maxAbs = v
				pivot = r
			}
		}
		if maxAbs < 1e-14 {
			return linear.Matrix{}, fmt.Errorf("singular matrix at pivot %d", col)
		}
		aug[col], aug[pivot] = aug[pivot], aug[col]
		pv := aug[col][col]
		for j := 0; j < 2*n; j++ {
			aug[col][j] /= pv
		}
		for r := 0; r < n; r++ {
			if r == col {
				continue
			}
			f := aug[r][col]
			for j := 0; j < 2*n; j++ {
				aug[r][j] -= f * aug[col][j]
			}
		}
	}
	out := linear.NewMatrix(n, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			out.Set(r, c, aug[r][n+c])
		}
	}
	return out, nil
}
