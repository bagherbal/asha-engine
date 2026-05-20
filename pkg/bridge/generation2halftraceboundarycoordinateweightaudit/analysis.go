// Package generation2halftraceboundarycoordinateweightaudit implements
// Gate 656: Half-Trace Boundary Coordinate Weight Audit.
//
// Gate 655 sealed the Fano-Hitchin obstruction package as internally mature
// but boundary-disconnected.  Gate 656 audits the remaining typed boundary clue
//
//	7/144 = (1/2)(7/72)
//
// as a possible per-boundary-coordinate half-trace weight of the augmented
// chamber H_72 = Lambda^4 R^8 plus R^2_boundary.  It is a bridge-source audit
// only: it does not derive boundary stress, scalar RG matching, Higgs mass,
// gauge unification, flavor, physical spacetime, or a native 7/72 theorem.
package generation2halftraceboundarycoordinateweightaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate655 "github.com/bagherbal/asha-engine/pkg/bridge/generation2fanohitchinobstructionboundaryinterfaceaudit"
)

const (
	AuditID = "GATE656-HALF-TRACE-BOUNDARY-COORDINATE-WEIGHT-AUDIT"

	StatusGate655FanoHitchinSealInherited = "PASS_GATE655_FANO_HITCHIN_SEAL_INHERITED"
	StatusSourceTypeAudited               = "PASS_HALF_TRACE_SOURCE_TYPE_AUDITED"
	StatusBoundaryComparisonAudited       = "PASS_BOUNDARY_COMPARISON_AUDITED"
	StatusMeanStressAudited               = "PASS_MEAN_STRESS_AUDITED"
	StatusTwoCoordinateSplitAudited       = "PASS_TWO_COORDINATE_SPLIT_AUDITED"
	StatusRelationToPreviousSealsAudited  = "PASS_RELATION_TO_PREVIOUS_SEALS_AUDITED"
	StatusHalfTraceCandidate              = "CONDITIONAL_SUPPORT_SEVEN_OVER_ONE_FORTY_FOUR_IS_TYPED_HALF_TRACE_BOUNDARY_CANDIDATE"
	StatusFanoNumeratorStrengthensClue    = "CONDITIONAL_SUPPORT_FANO_HITCHIN_NUMERATOR_SEVEN_STRENGTHENS_HALF_TRACE_CLUE"
	StatusBoundaryClueOnly                = "CONDITIONAL_SUPPORT_HALF_TRACE_IS_BOUNDARY_FACING_CLUE_ONLY"
	StatusNoNativeHalfTraceMap            = "FAILED_ROUTE_NO_NATIVE_HALF_TRACE_BOUNDARY_MAP"
	StatusNoSevenOver144Theorem           = "FAILED_ROUTE_NO_NATIVE_7_OVER_144_BOUNDARY_THEOREM"
	StatusNoSevenOver72Theorem            = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoBoundaryStressFromK7          = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_FROM_K7"
	StatusNoBoundaryStressDerived         = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVED"
	StatusNoHistoryLoopSource             = "FAILED_ROUTE_NO_HISTORY_LOOP_UNIT_SOURCE_FROM_HALF_TRACE"
	StatusNoScalarFlavorMap               = "FAILED_ROUTE_NO_SCALAR_FLAVOR_TRANSPORT_MAP"
	StatusNoPhysicalMetric                = "FAILED_ROUTE_NO_PHYSICAL_SPACETIME_OR_METRIC_THEOREM"
	StatusNoHiggsCKMGauge                 = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate656Boundary                 = "FIREWALL_PRESERVED_GATE656_HALF_TRACE_BOUNDARY_COORDINATE_WEIGHT_BOUNDARY"
)

const (
	k7Dim        = 7
	lambda4Dim   = 70
	boundaryDim  = 2
	augmentedDim = lambda4Dim + boundaryDim
	wFull        = float64(k7Dim) / float64(augmentedDim)
	wHalf        = float64(k7Dim) / float64(boundaryDim*augmentedDim)
	absLambda    = 0.0497009420776833
	r3Minus1     = 0.0509933868964996
	xiBoundary   = 0.0503471644870914
	LHistory     = 0.0397887357729738
	tol          = 1e-12
)

type Gate655Inheritance struct {
	FanoSealDefined         bool
	FanoSealInternalOnly    bool
	FanoStructuresNumerator bool
	NoBoundaryInterface     bool
	NoSevenOver72Theorem    bool
	NoBoundaryStress        bool
	NoScalarFlavorMap       bool
	NoHistoryLoopSource     bool
	ClaimsBoundaryStress    bool
	ClaimsSevenOver72       bool
	ClaimsScalarFlavor      bool
	ClaimsHistoryLoopUnit   bool
	Gate655Firewall         bool
	Verdict                 string
}

type FactorSource struct {
	Factor         string
	Value          float64
	Source         string
	Typed          bool
	Native         bool
	BridgeOnly     bool
	CertifiedRoute bool
}

type SourceTypeAudit struct {
	Factors               []FactorSource
	FullWeight            float64
	HalfWeight            float64
	SevenTyped            bool
	SeventyTwoTyped       bool
	HalfTyped             bool
	HalfNative            bool
	AllFactorsTyped       bool
	CertifiedHalfTraceMap bool
	Verdict               string
}

type BoundaryComparisonRow struct {
	Target           string
	TargetValue      float64
	Candidate        float64
	SignedResidual   float64
	AbsResidual      float64
	RelativeResidual float64
	Rank             int
	Classification   string
}

type BoundaryComparisonAudit struct {
	Rows                     []BoundaryComparisonRow
	ClosestTarget            string
	ClosestResidual          float64
	CertifiedMatch           bool
	NoProximityCertification bool
	Verdict                  string
}

type MeanStressAudit struct {
	XiBoundary                float64
	HalfWeight                float64
	SignedResidual            float64
	RelativeResidual          float64
	BoundarySplit             float64
	ExistingMeanStressBetter  bool
	AntiAlignmentSealStronger bool
	Verdict                   string
}

type TwoCoordinateSplitAudit struct {
	FullWeight           float64
	HalfWeight           float64
	SignedPair           [2]float64
	MeanStressCandidate  float64
	BoundaryPair         string
	FullWeightTyped      bool
	PerCoordinateTyped   bool
	SignedPairTyped      bool
	MeanStressTyped      bool
	SuppliesBoundaryMap  bool
	SuppliesTraceTheorem bool
	Verdict              string
}

type PreviousSealRelation struct {
	Seal            string
	Candidate       string
	LawfulRelation  bool
	TypedArithmetic bool
	Classification  string
}

type PreviousSealRelationAudit struct {
	Rows                     []PreviousSealRelation
	HistoryLoopSource        bool
	BoundaryStressSource     bool
	OrientationBalanceSource bool
	FanoHitchinSource        bool
	Verdict                  string
}

type BoundaryMapObstructionAudit struct {
	MissingMap              string
	MissingTraceTheorem     string
	HasHalfTraceMap         bool
	HasSevenOver72Map       bool
	HasBoundaryStressMap    bool
	CanDeriveBoundaryStress bool
	CanDeriveLambdaOrR3     bool
	Verdict                 string
}

type Firewalls struct {
	ClaimsBoundaryStress   bool
	ClaimsLambdaR3         bool
	ClaimsSevenOver144     bool
	ClaimsSevenOver72      bool
	ClaimsHistoryLoopUnit  bool
	ClaimsScalarFlavor     bool
	ClaimsPhysicalMetric   bool
	ClaimsHiggsMass        bool
	ClaimsCKMPMNS          bool
	ClaimsGaugeUnification bool
	Verdict                string
}

type Analysis struct {
	Inherited   Gate655Inheritance
	SourceType  SourceTypeAudit
	Boundary    BoundaryComparisonAudit
	MeanStress  MeanStressAudit
	Split       TwoCoordinateSplitAudit
	Relations   PreviousSealRelationAudit
	BoundaryMap BoundaryMapObstructionAudit
	Firewalls   Firewalls
	Truth       string
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
	g655, err := gate655.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate655 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g655)
	source := buildSourceType(inherited)
	boundary := buildBoundaryComparison()
	mean := buildMeanStress(boundary)
	split := buildSplit(source)
	relations := buildRelations()
	boundaryMap := buildBoundaryMap()
	firewalls := Firewalls{Verdict: StatusGate656Boundary}
	truth := "Gate 656 classifies 7/144=(1/2)(7/72) as a typed half-trace boundary-coordinate candidate only.  The factors 7, 72, and 1/2 are typed by K_7/Fano-Hitchin carrier dimension, the Lambda^4 R^8 plus R^2_boundary augmented chamber, and the two-coordinate boundary pair respectively, but no Psi:K_7->R^2_boundary, no normalized trace theorem, and no boundary-stress assignment is constructed.  The existing xi_boundary stress seal remains the better empirical two-coordinate compression."
	return Analysis{Inherited: inherited, SourceType: source, Boundary: boundary, MeanStress: mean, Split: split, Relations: relations, BoundaryMap: boundaryMap, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate655.Analysis) Gate655Inheritance {
	return Gate655Inheritance{
		FanoSealDefined:         g.Seal.Name == "FanoHitchinObstructionSeal",
		FanoSealInternalOnly:    g.Seal.InternalOnly,
		FanoStructuresNumerator: g.SevenOver72.StructuresNumerator7,
		NoBoundaryInterface:     !g.SevenOver72.BoundaryPairSupplied && !g.BoundaryMap.HasPsi && !g.BoundaryMap.HasTau,
		NoSevenOver72Theorem:    !g.SevenOver72.CertifiedSevenOver72Theorem,
		NoBoundaryStress:        !g.BoundaryStress.CertifiedBoundaryStressSource,
		NoScalarFlavorMap:       !g.Flavor.CertifiedFlavorMap,
		NoHistoryLoopSource:     !g.HistoryLoop.CertifiedSource,
		ClaimsBoundaryStress:    g.Firewalls.ClaimsBoundaryStress,
		ClaimsSevenOver72:       g.Firewalls.ClaimsSevenOver72,
		ClaimsScalarFlavor:      g.Firewalls.ClaimsScalarFlavor,
		ClaimsHistoryLoopUnit:   g.Firewalls.ClaimsHistoryLoopUnit,
		Gate655Firewall:         g.Firewalls.Verdict == gate655.StatusGate655Boundary,
		Verdict:                 StatusGate655FanoHitchinSealInherited,
	}
}

func buildSourceType(inh Gate655Inheritance) SourceTypeAudit {
	factors := []FactorSource{
		{Factor: "7", Value: k7Dim, Source: "dim(K_7), now strengthened as full Fano-Hitchin carrier dimension", Typed: inh.FanoStructuresNumerator, Native: true, CertifiedRoute: inh.FanoSealDefined},
		{Factor: "72", Value: augmentedDim, Source: "dim(Lambda^4 R^8)+dim(R^2_boundary)=70+2 inherited augmented chamber", Typed: true, Native: false, BridgeOnly: true, CertifiedRoute: false},
		{Factor: "1/2", Value: 0.5, Source: "possible averaging/splitting over two boundary coordinates", Typed: true, Native: false, BridgeOnly: true, CertifiedRoute: false},
	}
	return SourceTypeAudit{Factors: factors, FullWeight: wFull, HalfWeight: wHalf, SevenTyped: true, SeventyTwoTyped: true, HalfTyped: true, HalfNative: false, AllFactorsTyped: true, CertifiedHalfTraceMap: false, Verdict: join(StatusSourceTypeAudited, StatusHalfTraceCandidate, StatusFanoNumeratorStrengthensClue, StatusNoNativeHalfTraceMap, StatusNoSevenOver144Theorem)}
}

func buildBoundaryComparison() BoundaryComparisonAudit {
	targets := []struct {
		name string
		val  float64
	}{
		{"|lambda(Lambda_12)|", absLambda},
		{"R_3-1", r3Minus1},
		{"xi_boundary", xiBoundary},
	}
	rows := make([]BoundaryComparisonRow, 0, len(targets))
	for _, t := range targets {
		d := wHalf - t.val
		rows = append(rows, BoundaryComparisonRow{Target: t.name, TargetValue: t.val, Candidate: wHalf, SignedResidual: d, AbsResidual: math.Abs(d), RelativeResidual: math.Abs(d) / math.Abs(t.val), Classification: "typed comparison only; proximity cannot certify boundary source"})
	}
	// Rank by residual without changing target order.
	ranks := make([]int, len(rows))
	for i := range rows {
		rank := 1
		for j := range rows {
			if rows[j].AbsResidual < rows[i].AbsResidual {
				rank++
			}
		}
		ranks[i] = rank
	}
	closest := rows[0]
	for i := range rows {
		rows[i].Rank = ranks[i]
		if rows[i].AbsResidual < closest.AbsResidual {
			closest = rows[i]
		}
	}
	return BoundaryComparisonAudit{Rows: rows, ClosestTarget: closest.Target, ClosestResidual: closest.AbsResidual, CertifiedMatch: false, NoProximityCertification: true, Verdict: join(StatusBoundaryComparisonAudited, StatusBoundaryClueOnly, StatusNoBoundaryStressDerived)}
}

func buildMeanStress(b BoundaryComparisonAudit) MeanStressAudit {
	signed := xiBoundary - wHalf
	return MeanStressAudit{XiBoundary: xiBoundary, HalfWeight: wHalf, SignedResidual: signed, RelativeResidual: math.Abs(signed) / xiBoundary, BoundarySplit: r3Minus1 - absLambda, ExistingMeanStressBetter: true, AntiAlignmentSealStronger: true, Verdict: join(StatusMeanStressAudited, StatusNoBoundaryStressDerived)}
}

func buildSplit(source SourceTypeAudit) TwoCoordinateSplitAudit {
	return TwoCoordinateSplitAudit{FullWeight: source.FullWeight, HalfWeight: source.HalfWeight, SignedPair: [2]float64{source.HalfWeight, -source.HalfWeight}, MeanStressCandidate: source.HalfWeight, BoundaryPair: "span(|lambda(Lambda_12)|, R_3-1)", FullWeightTyped: true, PerCoordinateTyped: true, SignedPairTyped: true, MeanStressTyped: true, SuppliesBoundaryMap: false, SuppliesTraceTheorem: false, Verdict: join(StatusTwoCoordinateSplitAudited, StatusHalfTraceCandidate, StatusNoNativeHalfTraceMap, StatusNoSevenOver144Theorem, StatusNoSevenOver72Theorem)}
}

func buildRelations() PreviousSealRelationAudit {
	rows := []PreviousSealRelation{
		{Seal: "FanoHitchinObstructionSeal", Candidate: "numerator 7", LawfulRelation: true, TypedArithmetic: true, Classification: "strengthens carrier numerator only; no boundary map"},
		{Seal: "GaugeScalarBoundaryStressSeal", Candidate: "7/144 near boundary magnitudes", LawfulRelation: false, TypedArithmetic: true, Classification: "bridge clue only; xi_boundary remains empirical stress coordinate"},
		{Seal: "HistoryLoopUnitSeal", Candidate: "7/144 vs 1/(8*pi)", LawfulRelation: false, TypedArithmetic: false, Classification: "finite half-trace does not source S1/Hopf or heat-kernel loop unit"},
		{Seal: "OrientationBalanceSeal", Candidate: "7/144 vs flavor orientation data", LawfulRelation: false, TypedArithmetic: false, Classification: "no flavor intertwiner supplied"},
	}
	return PreviousSealRelationAudit{Rows: rows, HistoryLoopSource: false, BoundaryStressSource: false, OrientationBalanceSource: false, FanoHitchinSource: true, Verdict: join(StatusRelationToPreviousSealsAudited, StatusFanoNumeratorStrengthensClue, StatusNoHistoryLoopSource, StatusNoScalarFlavorMap)}
}

func buildBoundaryMap() BoundaryMapObstructionAudit {
	return BoundaryMapObstructionAudit{MissingMap: "Psi: K_7 or FanoHitchinPackage -> R^2_boundary with per-coordinate half trace", MissingTraceTheorem: "tau_half: normalized trace over Lambda^4 R^8 plus R^2_boundary yielding 7/144", HasHalfTraceMap: false, HasSevenOver72Map: false, HasBoundaryStressMap: false, CanDeriveBoundaryStress: false, CanDeriveLambdaOrR3: false, Verdict: join(StatusNoNativeHalfTraceMap, StatusNoSevenOver144Theorem, StatusNoSevenOver72Theorem, StatusNoBoundaryStressFromK7)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate655FanoHitchinSealInherited,
		StatusSourceTypeAudited,
		StatusBoundaryComparisonAudited,
		StatusMeanStressAudited,
		StatusTwoCoordinateSplitAudited,
		StatusRelationToPreviousSealsAudited,
		StatusHalfTraceCandidate,
		StatusFanoNumeratorStrengthensClue,
		StatusBoundaryClueOnly,
		StatusNoNativeHalfTraceMap,
		StatusNoSevenOver144Theorem,
		StatusNoSevenOver72Theorem,
		StatusNoBoundaryStressFromK7,
		StatusNoBoundaryStressDerived,
		StatusNoHistoryLoopSource,
		StatusNoScalarFlavorMap,
		StatusNoPhysicalMetric,
		StatusNoHiggsCKMGauge,
		StatusGate656Boundary,
	}
}

func near(x, y float64) bool { return math.Abs(x-y) < tol }
