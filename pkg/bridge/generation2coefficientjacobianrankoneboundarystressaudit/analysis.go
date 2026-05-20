// Package generation2coefficientjacobianrankoneboundarystressaudit implements
// Gate 616: Spectral-Action Coefficient Jacobian and Rank-One Boundary Stress Audit.
//
// Gate 615 showed that the GaugeScalarBoundaryStressSeal is type-safe as a
// bridge coefficient deformation after normalization, but that the native
// spectral-action grammar does not supply an SU(3)-only deformation, a sector-
// split f0, a scalar quartic boundary theorem, or a C3-lambda relation. Gate
// 616 audits the symbolic coefficient Jacobian: can a single coefficient
// direction produce the anti-aligned stress shadow (+xi,-xi), or does the
// bridge seal require independent gauge/scalar history slots?
package generation2coefficientjacobianrankoneboundarystressaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2spectralactioncoefficientgrammaraudit"
)

const (
	AuditID = "GATE616-COEFFICIENT-JACOBIAN-RANK-ONE-BOUNDARY-STRESS-AUDIT"

	StatusGate615Inherited              = "PASS_GATE615_COEFFICIENT_GRAMMAR_INHERITED"
	StatusDependencyGraphBuilt          = "PASS_COEFFICIENT_DEPENDENCY_GRAPH_BUILT"
	StatusNormalizedShadowMapDefined    = "PASS_NORMALIZED_SHADOW_MAP_DEFINED"
	StatusSymbolicJacobianAudited       = "PASS_SYMBOLIC_JACOBIAN_AUDITED"
	StatusRankOneSourceTested           = "PASS_RANK_ONE_SOURCE_CANDIDATES_TESTED"
	StatusAntiAlignmentAudited          = "PASS_ANTI_ALIGNMENT_TEST_AUDITED"
	StatusCanonicalNormalizationAudited = "PASS_CANONICAL_NORMALIZATION_AUDITED"
	StatusBridgeQStressDefinable        = "CONDITIONAL_SUPPORT_BOUNDARY_Q_STRESS_BRIDGE_SLOT_DEFINABLE"
	StatusOnlyRankTwoSlots              = "CONDITIONAL_SUPPORT_ONLY_RANK_TWO_INDEPENDENT_SLOTS_AVAILABLE_IN_NATIVE_GRAMMAR"
	StatusNoNativeRankOne               = "FAILED_ROUTE_NO_NATIVE_RANK_ONE_COEFFICIENT_SOURCE"
	StatusC3LambdaRankTwo               = "FAILED_ROUTE_C3_ONLY_AND_LAMBDA_ONLY_ARE_RANK_TWO_INDEPENDENT_SLOTS"
	StatusNoCoefficientAntiAlign        = "FAILED_ROUTE_NO_COEFFICIENT_SOURCE_FOR_ANTI_ALIGNED_STRESS"
	StatusNoF0SectorSplit               = "FAILED_ROUTE_NO_NATIVE_SECTOR_SPLIT_F0"
	StatusNoC3LambdaRelation            = "FAILED_ROUTE_NO_NATIVE_C3_LAMBDA_RELATION"
	StatusCanonicalScalarIncomplete     = "FAILED_ROUTE_CANONICAL_SCALAR_NORMALIZATION_LEDGER_INCOMPLETE"
	StatusNoThresholdMatching           = "FAILED_ROUTE_NO_NATIVE_THRESHOLD_OR_MATCHING_THEOREM"
	StatusNoNativeXi                    = "FAILED_ROUTE_NO_NATIVE_XI_BOUNDARY_THEOREM"
	StatusGate616Boundary               = "FIREWALL_PRESERVED_GATE616_COEFFICIENT_JACOBIAN_BOUNDARY"
)

type Inherited struct {
	Lambda12GeV         float64
	R3MinusOne          float64
	LambdaLambda12      float64
	XiBoundary          float64
	Delta3ColorBoundary float64
	DeltaLambdaBoundary float64
	Eta3                float64
	BoundaryResidual    float64
	ResidualOverXi      float64
	Verdict             string
}

type CoefficientDependency struct {
	Source        string
	AffectsGauge  string // +, -, 0, unknown
	AffectsScalar string // +, -, 0, unknown
	Dependency    string
	Certified     bool
	Bridge        bool
	Native        bool
	Obstruction   string
	Verdict       string
}

type NormalizedShadowMap struct {
	PreferredPair     string
	AlternatePair     string
	RawPair           string
	RawPairTypeSafe   bool
	PreferredTypeSafe bool
	ColorShadow       float64
	ScalarShadow      float64
	XiBoundary        float64
	Statement         string
	Verdict           string
}

type JacobianEntry struct {
	Source    string
	DColor    string
	DScalar   string
	Exact     bool
	Certified bool
	Comment   string
	Verdict   string
}

type RankOneSourceCandidate struct {
	Source                 string
	ProducesColorPositive  bool
	ProducesScalarNegative bool
	RankOneBridgeDefinable bool
	Native                 bool
	RequiresExtraSeal      bool
	Classification         string
	Obstruction            string
	Verdict                string
}

type RankClassification struct {
	NativeRankOneFound      bool
	BridgeRankOneDefinable  bool
	RankTwoIndependentSlots bool
	GrammarInsufficient     bool
	BestClassification      string
	Statement               string
	Verdict                 string
}

type AntiAlignmentTest struct {
	Candidate             string
	CanForceAntiAlignment bool
	Native                bool
	StressResidual        float64
	ResidualOverXi        float64
	Statement             string
	Verdict               string
}

type CanonicalNormalizationAudit struct {
	RuntimeLambdaConvention    string
	KPhiKnown                  bool
	CanonicalScalarLedgerKnown bool
	CanAuditLambdaBeforeAfterK bool
	Statement                  string
	Verdict                    string
}

type NativeStatus struct {
	SectorSplitF0       bool
	NativeQStress       bool
	C3LambdaRelation    bool
	ScalarNormalization bool
	ThresholdMatching   bool
	NativeXi            bool
	Statement           string
	Verdict             string
}

type Firewalls struct {
	ClaimsXiNative           bool
	ClaimsLambdaZero         bool
	ClaimsHiggsMass          bool
	ClaimsHiggsStability     bool
	ClaimsGaugeUnification   bool
	ClaimsThresholdExistence bool
	ClaimsNativeCorrection   bool
	Verdict                  string
}

type Analysis struct {
	Inherited              Inherited
	DependencyGraph        []CoefficientDependency
	ShadowMap              NormalizedShadowMap
	Jacobian               []JacobianEntry
	RankOneCandidates      []RankOneSourceCandidate
	RankClassification     RankClassification
	AntiAlignment          AntiAlignmentTest
	CanonicalNormalization CanonicalNormalizationAudit
	NativeStatus           NativeStatus
	Firewalls              Firewalls
	Truth                  string
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
	g615, err := generation2spectralactioncoefficientgrammaraudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate615 predecessor: %w", err)
	}
	inherited := inherit(g615)
	deps := buildDependencyGraph()
	shadow := buildShadowMap(inherited)
	jac := buildJacobian()
	candidates := buildRankOneCandidates()
	rank := classifyRank(candidates)
	anti := buildAntiAlignment(inherited)
	canon := buildCanonicalNormalization()
	native := buildNativeStatus()
	firewalls := auditFirewalls()
	truth := "Gate 616 audits the coefficient Jacobian behind the GaugeScalarBoundaryStressSeal. The spectral-action grammar can define a bridge q_stress slot whose normalized shadow is (+xi,-xi), but no native single coefficient direction is found. C3-only and lambda-only are independent rank-two slots, common f0 and K_phi lanes are normalization-sensitive or insufficient, and no sector-split f0, C3-lambda law, scalar normalization theorem, or threshold matching theorem is present. xi_boundary remains a bridge stress seal, not native ASHA law."
	return Analysis{inherited, deps, shadow, jac, candidates, rank, anti, canon, native, firewalls, truth}, nil
}

func inherit(a generation2spectralactioncoefficientgrammaraudit.Analysis) Inherited {
	return Inherited{
		Lambda12GeV:         a.Inherited.Lambda12GeV,
		R3MinusOne:          a.Inherited.R3MinusOne,
		LambdaLambda12:      a.Inherited.LambdaLambda12,
		XiBoundary:          a.Inherited.XiBoundary,
		Delta3ColorBoundary: a.Inherited.Delta3ColorBoundary,
		DeltaLambdaBoundary: a.Inherited.DeltaLambdaBoundary,
		Eta3:                a.Inherited.Eta3,
		BoundaryResidual:    a.Inherited.BoundaryResidual,
		ResidualOverXi:      a.Inherited.ResidualOverXi,
		Verdict:             StatusGate615Inherited,
	}
}

func buildDependencyGraph() []CoefficientDependency {
	return []CoefficientDependency{
		{"f0", "+ common gauge normalization", "unknown after scalar canonical normalization", "shared heat-kernel/coefficient lane", false, true, false, "common f0 is not sector-specific and does not force anti-aligned C3-lambda stress", StatusDependencyGraphBuilt},
		{"f2", "0", "mass/quadratic scalar potential lane", "cutoff moment / relevant scalar mass lane", false, true, false, "does not supply color kinetic correction", StatusDependencyGraphBuilt},
		{"f4", "0", "cosmological/volume lane", "cutoff moment / leading volume term", false, true, false, "not a gauge-scalar stress source", StatusDependencyGraphBuilt},
		{"C_3", "+", "0", "color inverse kinetic coefficient", true, true, false, "affects color slot only; needs independent scalar slot for stress pair", StatusC3LambdaRankTwo},
		{"lambda", "0", "+", "scalar quartic coefficient", true, true, false, "affects scalar slot only; needs independent color slot for stress pair", StatusC3LambdaRankTwo},
		{"K_phi", "0", "unknown / rescaling-sensitive", "scalar canonical normalization", false, true, false, "canonical scalar normalization ledger is incomplete", StatusCanonicalScalarIncomplete},
		{"a", "0", "unknown through scalar normalization", "Yukawa trace power sum", true, true, false, "can enter scalar formulas but not color kinetic slot", StatusDependencyGraphBuilt},
		{"b", "0", "+ through scalar quartic formula if supplied", "Yukawa quartic power sum", true, true, false, "no native relation to C3", StatusNoC3LambdaRelation},
		{"q_boundary", "+ by definition", "- by definition", "explicit bridge matching/stress slot", false, true, false, "bridge definable but tautological without native theorem", StatusBridgeQStressDefinable},
	}
}

func buildShadowMap(h Inherited) NormalizedShadowMap {
	return NormalizedShadowMap{
		PreferredPair:     "(G_color,S_scalar)=(R_3-1, lambda(Lambda_12))",
		AlternatePair:     "(eta_3, 2lambda(Lambda_12))",
		RawPair:           "(delta_3^color_boundary, delta_lambda_boundary)",
		RawPairTypeSafe:   false,
		PreferredTypeSafe: true,
		ColorShadow:       h.R3MinusOne,
		ScalarShadow:      h.LambdaLambda12,
		XiBoundary:        h.XiBoundary,
		Statement:         "The rank audit is type-safe only after mapping raw coefficient corrections to normalized dimensionless shadows such as (R_3-1, lambda) or (eta_3, 2lambda).",
		Verdict:           StatusNormalizedShadowMapDefined,
	}
}

func buildJacobian() []JacobianEntry {
	return []JacobianEntry{
		{"f0", "+ common", "unknown", false, false, "common normalization may affect several lanes but is not SU(3)-specific and can cancel after canonical normalization", StatusSymbolicJacobianAudited},
		{"sector-split f0_3", "+", "0", false, false, "would move color only; missing natively", StatusNoF0SectorSplit},
		{"C_3", "+", "0", true, false, "color-only slot", StatusC3LambdaRankTwo},
		{"lambda", "0", "+", true, false, "scalar-only slot", StatusC3LambdaRankTwo},
		{"K_phi", "0", "unknown", false, false, "field rescaling changes canonical lambda if ledger is supplied", StatusCanonicalScalarIncomplete},
		{"b/a^2", "0", "+", false, false, "scalar quartic proxy only", StatusNoC3LambdaRelation},
		{"finite Yukawa trace deformation", "0", "unknown/+", false, false, "scalar/Yukawa lane but no color kinetic source", StatusSymbolicJacobianAudited},
		{"q_boundary stress", "+", "-", false, false, "bridge-defined source can be assigned the anti-aligned shadow but is not native", StatusBridgeQStressDefinable},
	}
}

func buildRankOneCandidates() []RankOneSourceCandidate {
	return []RankOneSourceCandidate{
		{"common f0 shift", true, false, false, false, true, "not a certified anti-aligned rank-one source", "scalar response is normalization-dependent and C3-specificity is absent", StatusNoCoefficientAntiAlign},
		{"sector-split f0_3 shift", true, false, false, false, true, "color-only bridge slot", "sector-split f0 is missing natively and does not move lambda", StatusNoF0SectorSplit},
		{"C_3-only shift", true, false, false, false, false, "rank-two independent color slot", "does not affect scalar quartic", StatusC3LambdaRankTwo},
		{"lambda-only shift", false, true, false, false, false, "rank-two independent scalar slot", "does not affect color kinetic coefficient", StatusC3LambdaRankTwo},
		{"K_phi shift", false, false, false, false, true, "scalar normalization-sensitive slot", "canonical scalar normalization ledger is incomplete", StatusCanonicalScalarIncomplete},
		{"b/a^2 scalar quartic shift", false, true, false, false, true, "scalar-only bridge/theory proxy", "no native C3 relation", StatusNoC3LambdaRelation},
		{"finite Yukawa trace deformation", false, false, false, false, true, "scalar/Yukawa-sensitive but not color source", "no color kinetic channel", StatusNoCoefficientAntiAlign},
		{"q_boundary stress", true, true, true, false, true, "bridge rank-one source definable by the seal", "tautological unless a threshold/coefficient theorem supplies it", StatusBridgeQStressDefinable},
	}
}

func classifyRank(rows []RankOneSourceCandidate) RankClassification {
	bridge := false
	native := false
	for _, r := range rows {
		if r.RankOneBridgeDefinable {
			bridge = true
		}
		if r.Native && r.ProducesColorPositive && r.ProducesScalarNegative {
			native = true
		}
	}
	return RankClassification{
		NativeRankOneFound:      native,
		BridgeRankOneDefinable:  bridge,
		RankTwoIndependentSlots: true,
		GrammarInsufficient:     false,
		BestClassification:      "B/C hybrid: bridge q_stress can be defined, but native grammar supplies only independent C3 and lambda slots",
		Statement:               "No native rank-one coefficient source is found. A bridge q_stress can be declared to have shadow (+xi,-xi), but without a sector-split f0, threshold matching theorem, or C3-lambda relation, the native grammar remains rank-two for the actual color/scalar corrections.",
		Verdict:                 StatusNoNativeRankOne,
	}
}

func buildAntiAlignment(h Inherited) AntiAlignmentTest {
	return AntiAlignmentTest{
		Candidate:             "q_stress -> (R_3-1, lambda)=(+xi,-xi)",
		CanForceAntiAlignment: false,
		Native:                false,
		StressResidual:        h.BoundaryResidual,
		ResidualOverXi:        math.Abs(h.BoundaryResidual) / h.XiBoundary,
		Statement:             "The bridge stress coordinate records the observed anti-alignment, but no coefficient source forces R_3-1+lambda=0. Anti-alignment remains a sealed boundary-history compression.",
		Verdict:               StatusNoCoefficientAntiAlign,
	}
}

func buildCanonicalNormalization() CanonicalNormalizationAudit {
	return CanonicalNormalizationAudit{
		RuntimeLambdaConvention:    "SM endpoint convention lambda=m_H^2/(2v^2), then v1 one-loop/top-dominant RG transport",
		KPhiKnown:                  false,
		CanonicalScalarLedgerKnown: false,
		CanAuditLambdaBeforeAfterK: false,
		Statement:                  "The runtime lambda is in canonical SM convention, but ASHA's internal K_phi/scalar metric normalization ledger is not sufficient to audit pre- versus post-canonical spectral-action lambda as a native coefficient.",
		Verdict:                    StatusCanonicalScalarIncomplete,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		SectorSplitF0:       false,
		NativeQStress:       false,
		C3LambdaRelation:    false,
		ScalarNormalization: false,
		ThresholdMatching:   false,
		NativeXi:            false,
		Statement:           "Current ASHA supplies no sector-split f0, native q_stress, C3-lambda coefficient relation, complete scalar canonical normalization theorem, or finite threshold matching theorem.",
		Verdict:             StatusNoNativeRankOne,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, false, StatusGate616Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate615Inherited,
		StatusDependencyGraphBuilt,
		StatusNormalizedShadowMapDefined,
		StatusSymbolicJacobianAudited,
		StatusRankOneSourceTested,
		StatusAntiAlignmentAudited,
		StatusCanonicalNormalizationAudited,
		StatusBridgeQStressDefinable,
		StatusOnlyRankTwoSlots,
		StatusNoNativeRankOne,
		StatusC3LambdaRankTwo,
		StatusNoCoefficientAntiAlign,
		StatusNoF0SectorSplit,
		StatusNoC3LambdaRelation,
		StatusCanonicalScalarIncomplete,
		StatusNoThresholdMatching,
		StatusNoNativeXi,
		StatusGate616Boundary,
	}
}
