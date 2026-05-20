// Package generation2scalarcanonicalnormalizationspectralquarticairlockaudit implements
// Gate 617: Scalar Canonical Normalization and Spectral Quartic Airlock Audit.
//
// Gate 616 identified the incomplete scalar canonical-normalization ledger as
// the next blocker for interpreting the GaugeScalarBoundaryStressSeal. Gate 617
// audits the scalar coefficient airlock from pre-canonical spectral-action
// scalar data through K_phi normalization into the runtime Standard Model
// quartic lambda(mu). It is a type-normalization audit only: no Higgs mass,
// stability, lambda=0, or native stress theorem is promoted.
package generation2scalarcanonicalnormalizationspectralquarticairlockaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2coefficientjacobianrankoneboundarystressaudit"
)

const (
	AuditID = "GATE617-SCALAR-CANONICAL-NORMALIZATION-SPECTRAL-QUARTIC-AIRLOCK-AUDIT"

	StatusGate616Inherited             = "PASS_GATE616_CANONICAL_NORMALIZATION_BLOCKER_INHERITED"
	StatusScalarCoefficientTypes       = "PASS_SCALAR_COEFFICIENT_TYPES_CLASSIFIED"
	StatusRuntimeConventionAudited     = "PASS_RUNTIME_LAMBDA_CONVENTION_AUDITED"
	StatusCanonicalMapWritten          = "PASS_CANONICAL_NORMALIZATION_MAP_WRITTEN_SYMBOLICALLY"
	StatusABF0Audited                  = "PASS_SPECTRAL_ACTION_A_B_F0_LANE_AUDITED"
	StatusRuntimeLambdaCanonicalLedger = "CONDITIONAL_SUPPORT_RUNTIME_LAMBDA_IS_CANONICAL_SM_QUARTIC_LEDGER"
	StatusRuntimeLambdaBridgeOnly      = "CONDITIONAL_SUPPORT_RUNTIME_LAMBDA_IS_V1_TRANSPORTED_BRIDGE_LEDGER"
	StatusNoKPhiTheorem                = "FAILED_ROUTE_NO_NATIVE_K_PHI_NORMALIZATION_THEOREM"
	StatusNoABF0ToLambdaAirlock        = "FAILED_ROUTE_NO_NATIVE_A_B_F0_TO_LAMBDA_AIRLOCK"
	StatusStressScalarRuntimeShadow    = "FAILED_ROUTE_STRESS_SEAL_SCALAR_SIDE_REMAINS_RUNTIME_SHADOW"
	StatusNoHiggsVEVOrMatching         = "FAILED_ROUTE_NO_NATIVE_HIGGS_VEV_OR_MATCHING_THEOREM"
	StatusNoNativeScalarQuartic        = "FAILED_ROUTE_NO_NATIVE_SCALAR_QUARTIC_DERIVATION"
	StatusNoLambdaZero                 = "FAILED_ROUTE_NO_NATIVE_LAMBDA_ZERO_BOUNDARY_THEOREM"
	StatusGate617Boundary              = "FIREWALL_PRESERVED_GATE617_SCALAR_NORMALIZATION_AIRLOCK_BOUNDARY"
)

type Inherited struct {
	Lambda12GeV       float64
	LambdaLambda12    float64
	XiBoundary        float64
	R3MinusOne        float64
	DeltaLambdaBridge float64
	Verdict           string
}

type ScalarCoefficient struct {
	Symbol      string
	Meaning     string
	Layer       string
	Status      string
	Native      bool
	Bridge      bool
	Observed    bool
	MissingData string
	Verdict     string
}

type RuntimeLambdaConventionAudit struct {
	RuntimeFormula        string
	PotentialConvention   string
	CanonicalSMConvention bool
	BridgeRuntime         bool
	RequiresMatching      bool
	Statement             string
	Verdict               string
}

type CanonicalNormalizationMap struct {
	PreCanonicalKinetic string
	PreCanonicalQuartic string
	CanonicalField      string
	CanonicalQuartic    string
	ExactFormulaNative  bool
	Dependency          string
	Statement           string
	Verdict             string
}

type SpectralActionABF0Audit struct {
	CandidateShape      string
	AAvailableNative    bool
	BAvailableNative    bool
	F0AvailableNative   bool
	KPhiAvailableNative bool
	FormulaCertified    bool
	ConventionCertified bool
	Statement           string
	Verdict             string
}

type RuntimeToBoundaryAirlockStatus struct {
	LambdaRuntimeValue       float64
	IsCanonicalRuntimeLedger bool
	IsV1Transported          bool
	TopMassSensitive         bool
	AlphaSSensitive          bool
	ThresholdSensitive       bool
	LoopOrderSensitive       bool
	EquivalentToPreCanonical bool
	Statement                string
	Verdict                  string
}

type StressSealImpactAssessment struct {
	OriginalScalarShadow    float64
	XiBoundary              float64
	RuntimeStressResidual   float64
	CanReplaceByLambdaCanon bool
	CanReplaceByLambdaPhi   bool
	ScalarSideStatus        string
	Statement               string
	Verdict                 string
}

type NativeStatus struct {
	NativeKPhi          bool
	NativeScalarMetric  bool
	NativeLambdaPhi     bool
	NativeABF0ToLambda  bool
	NativeVEV           bool
	NativeMatching      bool
	NativeLambdaZero    bool
	NativeStressTheorem bool
	Statement           string
	Verdict             string
}

type Firewalls struct {
	ClaimsLambdaZero       bool
	ClaimsHiggsStability   bool
	ClaimsHiggsPoleMass    bool
	ClaimsNativeQuartic    bool
	ClaimsNativeStressSeal bool
	ClaimsGaugeUnification bool
	Verdict                string
}

type Analysis struct {
	Inherited          Inherited
	ScalarCoefficients []ScalarCoefficient
	RuntimeConvention  RuntimeLambdaConventionAudit
	CanonicalMap       CanonicalNormalizationMap
	ABF0Audit          SpectralActionABF0Audit
	RuntimeAirlock     RuntimeToBoundaryAirlockStatus
	StressSealImpact   StressSealImpactAssessment
	NativeStatus       NativeStatus
	Firewalls          Firewalls
	Truth              string
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
	g616, err := generation2coefficientjacobianrankoneboundarystressaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate616 predecessor: %w", err)
	}
	inherited := inherit(g616)
	coeffs := buildScalarCoefficientTable()
	convention := buildRuntimeConvention()
	canonical := buildCanonicalMap()
	abf0 := buildABF0Audit()
	airlock := buildRuntimeAirlock(inherited)
	impact := buildStressImpact(inherited)
	native := buildNativeStatus()
	firewalls := auditFirewalls()
	truth := "Gate 617 audits the scalar canonical-normalization airlock. Runtime lambda is a canonical Standard Model quartic ledger transported by v1 RG, but ASHA does not yet supply a native K_phi/scalar metric normalization theorem, a pre-canonical Lambda_phi ledger, a certified a,b,f0-to-lambda formula, Higgs VEV derivation, or matching theorem. Therefore the scalar side of the GaugeScalarBoundaryStressSeal remains a runtime shadow until the canonical scalar airlock closes."
	return Analysis{inherited, coeffs, convention, canonical, abf0, airlock, impact, native, firewalls, truth}, nil
}

func inherit(a generation2coefficientjacobianrankoneboundarystressaudit.Analysis) Inherited {
	return Inherited{
		Lambda12GeV:       a.Inherited.Lambda12GeV,
		LambdaLambda12:    a.Inherited.LambdaLambda12,
		XiBoundary:        a.Inherited.XiBoundary,
		R3MinusOne:        a.Inherited.R3MinusOne,
		DeltaLambdaBridge: a.Inherited.DeltaLambdaBoundary,
		Verdict:           StatusGate616Inherited,
	}
}

func buildScalarCoefficientTable() []ScalarCoefficient {
	return []ScalarCoefficient{
		{"K_phi", "scalar kinetic normalization multiplying |D_phi phi|^2", "pre-canonical scalar metric", "bridge/missing native ledger", false, true, false, "native scalar metric/K_phi normalization theorem", StatusNoKPhiTheorem},
		{"Lambda_phi", "pre-canonical scalar quartic coefficient multiplying |phi|^4", "pre-canonical spectral-action scalar potential", "symbolic slot", false, true, false, "certified spectral-action coefficient and convention", StatusNoNativeScalarQuartic},
		{"lambda_canon", "post-canonical scalar quartic Lambda_phi/K_phi^2 up to convention", "canonical scalar coefficient", "bridge map only", false, true, false, "K_phi and Lambda_phi airlock", StatusNoABF0ToLambdaAirlock},
		{"lambda_runtime", "SM endpoint quartic m_H^2/(2v^2) transported to Lambda_12", "runtime canonical SM ledger", "observed/bridge", false, true, true, "higher-loop/matching/top-mass precision", StatusRuntimeLambdaCanonicalLedger},
		{"a", "finite Yukawa quadratic trace coefficient", "finite spectral-action trace lane", "native trace object, not scalar airlock theorem", true, true, false, "certified map into canonical lambda", StatusABF0Audited},
		{"b", "finite Yukawa quartic trace coefficient", "finite spectral-action trace lane", "native trace object, not scalar airlock theorem", true, true, false, "certified b/a^2 normalization and convention", StatusABF0Audited},
		{"f0", "spectral-action cutoff moment controlling dimension-four coefficient lanes", "cutoff/coefficient lane", "native symbol/bridge value", false, true, false, "sector normalization and scalar map", StatusABF0Audited},
		{"scalar metric normalization", "field-space metric used before canonical scalar rescaling", "normalization airlock", "missing", false, true, false, "native scalar metric theorem", StatusNoKPhiTheorem},
		{"v", "Higgs VEV inferred from G_F in runtime", "observed endpoint ledger", "observed/bridge", false, true, true, "native VEV and matching theorem", StatusNoHiggsVEVOrMatching},
	}
}

func buildRuntimeConvention() RuntimeLambdaConventionAudit {
	return RuntimeLambdaConventionAudit{
		RuntimeFormula:        "lambda(M_Z)=m_H^2/(2v^2); lambda(mu) transported by v1 one-loop/top-dominant SM RG",
		PotentialConvention:   "V(H)=-mu^2 |H|^2 + lambda |H|^4 in the Standard Model convention used by the runtime ledger",
		CanonicalSMConvention: true,
		BridgeRuntime:         true,
		RequiresMatching:      true,
		Statement:             "Runtime lambda is a canonical SM quartic ledger, not a pre-canonical spectral-action coefficient. It can be compared to spectral-action scalar data only through a K_phi/Lambda_phi/matching airlock.",
		Verdict:               StatusRuntimeConventionAudited,
	}
}

func buildCanonicalMap() CanonicalNormalizationMap {
	return CanonicalNormalizationMap{
		PreCanonicalKinetic: "K_phi |D_phi phi|^2",
		PreCanonicalQuartic: "Lambda_phi |phi|^4",
		CanonicalField:      "phi_c = sqrt(K_phi) phi",
		CanonicalQuartic:    "lambda_canon = Lambda_phi / K_phi^2 up to normalization and Higgs-doublet convention",
		ExactFormulaNative:  false,
		Dependency:          "lambda_runtime(mu) ?= RG+matching+Normalize[Lambda_phi,K_phi,scalar metric,a,b,f0]",
		Statement:           "The formal canonical map is type-clear, but ASHA lacks the native K_phi/Lambda_phi/scalar metric ledger needed to identify runtime lambda(Lambda_12) with a spectral-action quartic coefficient.",
		Verdict:             StatusCanonicalMapWritten,
	}
}

func buildABF0Audit() SpectralActionABF0Audit {
	return SpectralActionABF0Audit{
		CandidateShape:      "lambda_canon may have symbolic dependence on b/a^2, f0, and scalar normalization in spectral-action conventions",
		AAvailableNative:    true,
		BAvailableNative:    true,
		F0AvailableNative:   false,
		KPhiAvailableNative: false,
		FormulaCertified:    false,
		ConventionCertified: false,
		Statement:           "The a,b trace lanes are native polynomial trace objects, but the project currently has no certified native a,b,f0,K_phi-to-canonical-lambda airlock or convention-complete scalar quartic formula.",
		Verdict:             StatusNoABF0ToLambdaAirlock,
	}
}

func buildRuntimeAirlock(h Inherited) RuntimeToBoundaryAirlockStatus {
	return RuntimeToBoundaryAirlockStatus{
		LambdaRuntimeValue:       h.LambdaLambda12,
		IsCanonicalRuntimeLedger: true,
		IsV1Transported:          true,
		TopMassSensitive:         true,
		AlphaSSensitive:          true,
		ThresholdSensitive:       true,
		LoopOrderSensitive:       true,
		EquivalentToPreCanonical: false,
		Statement:                "lambda(Lambda_12) is a v1 transported canonical SM runtime quartic. It is not equivalent to a pre-canonical spectral-action quartic without K_phi, scalar metric, matching, and higher-loop ledgers.",
		Verdict:                  StatusRuntimeLambdaBridgeOnly,
	}
}

func buildStressImpact(h Inherited) StressSealImpactAssessment {
	return StressSealImpactAssessment{
		OriginalScalarShadow:    h.LambdaLambda12,
		XiBoundary:              h.XiBoundary,
		RuntimeStressResidual:   h.R3MinusOne + h.LambdaLambda12,
		CanReplaceByLambdaCanon: false,
		CanReplaceByLambdaPhi:   false,
		ScalarSideStatus:        "runtime canonical SM shadow only",
		Statement:               "Replacing lambda_runtime with a spectral-action lambda_canon or Lambda_phi/K_phi^2 could change the GaugeScalarBoundaryStressSeal interpretation. Until the scalar airlock closes, the scalar side remains a runtime shadow.",
		Verdict:                 StatusStressScalarRuntimeShadow,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		NativeKPhi:          false,
		NativeScalarMetric:  false,
		NativeLambdaPhi:     false,
		NativeABF0ToLambda:  false,
		NativeVEV:           false,
		NativeMatching:      false,
		NativeLambdaZero:    false,
		NativeStressTheorem: false,
		Statement:           "Current ASHA does not provide native K_phi, scalar metric, Lambda_phi, a,b,f0-to-lambda theorem, Higgs VEV derivation, scalar matching theorem, lambda=0 boundary theorem, or native gauge-scalar stress theorem.",
		Verdict:             StatusNoKPhiTheorem,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, StatusGate617Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate616Inherited,
		StatusScalarCoefficientTypes,
		StatusRuntimeConventionAudited,
		StatusCanonicalMapWritten,
		StatusABF0Audited,
		StatusRuntimeLambdaCanonicalLedger,
		StatusRuntimeLambdaBridgeOnly,
		StatusNoKPhiTheorem,
		StatusNoABF0ToLambdaAirlock,
		StatusStressScalarRuntimeShadow,
		StatusNoHiggsVEVOrMatching,
		StatusNoNativeScalarQuartic,
		StatusNoLambdaZero,
		StatusGate617Boundary,
	}
}

func almost(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
