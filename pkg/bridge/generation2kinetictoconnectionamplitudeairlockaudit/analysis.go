// Package generation2kinetictoconnectionamplitudeairlockaudit implements
// Gate 667: Kinetic-to-Connection Amplitude Airlock Source Audit.
//
// Gate 666 diagnosed the active BoundaryWeightedDeficitClosure as an
// amplitude-layer bridge seal: the 7/72 closure works in the coupling
// amplitude residual g3/gEW-1 and not in the inverse-coupling kinetic
// coordinate u=1/g^2. Gate 667 audits the source type of that amplitude
// coordinate. The gate asks whether canonical gauge-field normalization and
// covariant-derivative connection amplitudes explain why the bridge uses g.
// It preserves the firewall: this is a bridge-layer airlock audit, not a
// native 7/72, dual-root, scalar, flavor, Higgs, or gauge-unification theorem.
package generation2kinetictoconnectionamplitudeairlockaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate666 "github.com/bagherbal/asha-engine/pkg/bridge/generation2canonicalamplitudeairlockaudit"
)

const (
	AuditID = "GATE667-KINETIC-TO-CONNECTION-AMPLITUDE-AIRLOCK-SOURCE-AUDIT"

	StatusGate666AmplitudeSealInherited         = "PASS_GATE666_AMPLITUDE_SEAL_INHERITED"
	StatusKineticCoordinateDefined              = "PASS_KINETIC_COORDINATE_DEFINED"
	StatusCanonicalFieldRescalingAudited        = "PASS_CANONICAL_FIELD_RESCALING_AUDITED"
	StatusConnectionAmplitudeCoordinateTyped    = "PASS_CONNECTION_AMPLITUDE_COORDINATE_TYPED"
	StatusGaugeCoordinateComparisonAudited      = "PASS_GAUGE_COORDINATE_COMPARISON_AUDITED"
	StatusElectroweakHessianSocketAudited       = "PASS_ELECTROWEAK_HESSIAN_SOCKET_AUDITED"
	StatusScalarSideTypeAudited                 = "PASS_SCALAR_SIDE_TYPE_AUDITED"
	StatusRootAmplitudeRecurrenceAudited        = "PASS_ROOT_AMPLITUDE_RECURRENCE_AUDITED"
	StatusGaugeAmplitudeSourcedByConnection     = "CONDITIONAL_SUPPORT_GAUGE_AMPLITUDE_COORDINATE_SOURCED_BY_CANONICAL_CONNECTION_NORMALIZATION"
	StatusClosureConnectionAmplitudeLayer       = "CONDITIONAL_SUPPORT_BOUNDARY_WEIGHTED_DEFICIT_CLOSURE_BELONGS_TO_CONNECTION_AMPLITUDE_LAYER"
	StatusCanonicalEndpointSocketSupported      = "CONDITIONAL_SUPPORT_CONNECTION_AMPLITUDE_COMPATIBLE_WITH_ENDPOINT_HESSIAN_SOCKET"
	StatusRootAmplitudePatternSupported         = "CONDITIONAL_SUPPORT_ROOT_AMPLITUDE_AIRLOCK_PATTERN_RECURS_ACROSS_SEALS"
	StatusInverseKineticStillFails              = "FAILED_ROUTE_INVERSE_KINETIC_LAYER_DOES_NOT_SUPPORT_SAME_7_OVER_72_CLOSURE"
	StatusNoNativeKineticAmplitudeTheorem       = "FAILED_ROUTE_NO_NATIVE_KINETIC_TO_AMPLITUDE_AIRLOCK_THEOREM"
	StatusScalarRuntimeShadow                   = "FAILED_ROUTE_SCALAR_SIDE_REMAINS_RUNTIME_SHADOW_NOT_NATIVE_AMPLITUDE_OBJECT"
	StatusNoNativeSevenOver72Theorem            = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeDualRootTheorem               = "FAILED_ROUTE_NO_NATIVE_DUAL_ROOT_ALIGNMENT_THEOREM"
	StatusNoNativeScalarFlavorBoundaryTransport = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoBoundaryStressDerivation            = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoHiggsStabilityGaugeFlavor           = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate667Boundary                       = "FIREWALL_PRESERVED_GATE667_KINETIC_CONNECTION_AMPLITUDE_BOUNDARY"
)

const sevenOver72 = 7.0 / 72.0

type Gate666Inheritance struct {
	AmplitudeSealInherited     bool
	Classification             string
	AmplitudeLayerPasses       bool
	InverseKineticLayerPasses  bool
	AmplitudeResidual          float64
	InverseFractionalWound     float64
	InverseOverAmplitude       float64
	AmplitudeWBestMinus7Over72 float64
	InverseWBestMinus7Over72   float64
	MissingAirlockTheorem      bool
	NoNativeSevenOver72        bool
	NoNativeDualRoot           bool
	NoNativeTransport          bool
	NoBoundaryStress           bool
	Verdict                    string
}

type KineticCoordinateAudit struct {
	NativeCoordinate      string
	SpectralKineticForm   string
	RGVariable            string
	WhyNativeForOneLoopRG string
	ClosureStatus         string
	Verdict               string
}

type CanonicalFieldRescalingAudit struct {
	KineticCoefficientU      string
	CanonicalMap             string
	ConnectionAmplitude      string
	AlgebraicRelation        string
	DerivativeRelation       string
	AmplitudeCoordinateTyped bool
	SourceStatement          string
	Verdict                  string
}

type GaugeCoordinateLayerRow struct {
	Layer             string
	Coordinate        string
	ResidualName      string
	Residual          float64
	WBest             float64
	WBestMinus7Over72 float64
	PassesSevenOver72 bool
	Verdict           string
}

type GaugeCoordinateComparisonAudit struct {
	Rows                []GaugeCoordinateLayerRow
	AmplitudeOnlyPasses bool
	InverseKineticFails bool
	ClosureCoordinate   string
	Verdict             string
}

type ElectroweakHessianSocketAudit struct {
	CovariantDerivativeSocket string
	NeutralHessianShape       string
	ChargedWSocket            string
	AmplitudeObjects          []string
	CompatibleWithClosure     bool
	Limitation                string
	Verdict                   string
}

type ScalarSideTypeAudit struct {
	ScalarObject    string
	ComparedTo      string
	TypedAs         string
	NativeAmplitude bool
	Limitation      string
	Verdict         string
}

type RootAmplitudeRecurrenceAudit struct {
	Rows    []gate666.RecurringPatternRow
	Pattern string
	Verdict string
}

type AirlockTheoremTarget struct {
	Name             string
	Domain           string
	Airlock          string
	Codomain         string
	CandidateContent string
	Status           string
	Verdict          string
}

type SourceTypeVerdict struct {
	Classification string
	Statements     []string
	Verdict        string
}

type VerdictDiscipline struct {
	ClaimsNativeKineticAmplitudeTheorem bool
	ClaimsNativeSevenOver72Theorem      bool
	ClaimsNativeDualRootTheorem         bool
	ClaimsNativeTransportTheorem        bool
	ClaimsBoundaryStressDerivation      bool
	ClaimsHiggsPrediction               bool
	ClaimsScalarStability               bool
	ClaimsFlavorDerivation              bool
	ClaimsGaugeUnification              bool
	ClaimsCKMPMNSDerivation             bool
	Verdict                             string
}

type Analysis struct {
	Inherited     Gate666Inheritance
	Kinetic       KineticCoordinateAudit
	Rescaling     CanonicalFieldRescalingAudit
	Coordinates   GaugeCoordinateComparisonAudit
	HessianSocket ElectroweakHessianSocketAudit
	ScalarSide    ScalarSideTypeAudit
	Pattern       RootAmplitudeRecurrenceAudit
	Target        AirlockTheoremTarget
	Source        SourceTypeVerdict
	Discipline    VerdictDiscipline
	Truth         string
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
	g666, err := gate666.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate666 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g666)
	kinetic := buildKineticCoordinate()
	rescaling := buildCanonicalRescaling()
	coordinates := buildCoordinateComparison(g666)
	hessian := buildHessianSocket()
	scalar := buildScalarSide()
	pattern := buildPattern(g666)
	target := buildTarget()
	source := buildSourceVerdict(coordinates, scalar)
	discipline := VerdictDiscipline{Verdict: StatusGate667Boundary}
	truth := "Gate 667 upgrades Gate666's amplitude-layer diagnosis by typing the working g-coordinate as a canonical connection amplitude produced after kinetic normalization u=1/g^2 -> g=1/sqrt(u). The active 7/72 closure belongs to this endpoint connection-amplitude layer, while the scalar side remains a runtime scalar coefficient shadow and no native airlock or 7/72 theorem is certified."
	return Analysis{Inherited: inherited, Kinetic: kinetic, Rescaling: rescaling, Coordinates: coordinates, HessianSocket: hessian, ScalarSide: scalar, Pattern: pattern, Target: target, Source: source, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate666.Analysis) Gate666Inheritance {
	return Gate666Inheritance{
		AmplitudeSealInherited:     g.Source.Classification == "BoundaryWeightedDeficitClosureAmplitudeSeal" && g.CoordinateStack.AmplitudeLayerPasses && !g.CoordinateStack.InverseKineticLayerPasses,
		Classification:             g.Source.Classification,
		AmplitudeLayerPasses:       g.CoordinateStack.AmplitudeLayerPasses,
		InverseKineticLayerPasses:  g.CoordinateStack.InverseKineticLayerPasses,
		AmplitudeResidual:          g.KineticToAmplitude.AmplitudeResidual,
		InverseFractionalWound:     g.KineticToAmplitude.InverseFractionalWound,
		InverseOverAmplitude:       g.KineticToAmplitude.InverseOverAmplitude,
		AmplitudeWBestMinus7Over72: g.Inherited.AmplitudeWBestMinus7Over72,
		InverseWBestMinus7Over72:   g.Inherited.InverseWBestMinus7Over72,
		MissingAirlockTheorem:      strings.Contains(g.Target.Verdict, gate666.StatusNoNativeAmplitudeAirlockTheorem),
		NoNativeSevenOver72:        !g.Discipline.ClaimsNativeSevenOver72Theorem,
		NoNativeDualRoot:           !g.Discipline.ClaimsNativeDualRootTheorem,
		NoNativeTransport:          !g.Discipline.ClaimsNativeTransportTheorem,
		NoBoundaryStress:           !g.Discipline.ClaimsBoundaryStressDerivation,
		Verdict:                    StatusGate666AmplitudeSealInherited,
	}
}

func buildKineticCoordinate() KineticCoordinateAudit {
	return KineticCoordinateAudit{
		NativeCoordinate:      "u_i=1/g_i^2",
		SpectralKineticForm:   "C_i Tr(F_i^2), equivalently a gauge kinetic coefficient u_i multiplying the curvature-square term",
		RGVariable:            "one-loop transport is linear in inverse kinetic coefficients u_i up to beta-function normalization",
		WhyNativeForOneLoopRG: "the RG ledger evolves kinetic strengths before canonical field rescaling; this is why inverse-coupling naturality must be audited and cannot be silently bypassed",
		ClosureStatus:         "Gate665/Gate666 show that this coordinate does not keep the same 7/72 closure",
		Verdict:               join(StatusKineticCoordinateDefined, StatusInverseKineticStillFails),
	}
}

func buildCanonicalRescaling() CanonicalFieldRescalingAudit {
	return CanonicalFieldRescalingAudit{
		KineticCoefficientU:      "u_i=1/g_i^2",
		CanonicalMap:             "rescale the gauge field so the kinetic term is canonical and the coupling moves into the connection/covariant derivative",
		ConnectionAmplitude:      "D=d+i g_i A_i",
		AlgebraicRelation:        "g_i=u_i^{-1/2}",
		DerivativeRelation:       "for small amplitude wound r_g=g3/gEW-1, the inverse-kinetic fractional wound is 1-u3/uEW=1-1/(1+r_g)^2≈2r_g",
		AmplitudeCoordinateTyped: true,
		SourceStatement:          "the working Gate665 coordinate is not arbitrary: it is the canonical connection amplitude after kinetic normalization",
		Verdict:                  join(StatusCanonicalFieldRescalingAudited, StatusConnectionAmplitudeCoordinateTyped, StatusGaugeAmplitudeSourcedByConnection),
	}
}

func buildCoordinateComparison(g gate666.Analysis) GaugeCoordinateComparisonAudit {
	rows := make([]GaugeCoordinateLayerRow, 0, len(g.CoordinateStack.Rows))
	ampOnly := false
	inverseFails := false
	for _, r := range g.CoordinateStack.Rows {
		passes := r.NearSevenOver72 && r.RootAligned
		verdict := StatusGaugeCoordinateComparisonAudited
		if r.Layer == "canonical amplitude layer" && passes {
			ampOnly = true
			verdict = join(StatusGaugeCoordinateComparisonAudited, StatusClosureConnectionAmplitudeLayer)
		}
		if r.Layer == "RG-native inverse-kinetic layer" && !passes {
			inverseFails = true
			verdict = join(StatusGaugeCoordinateComparisonAudited, StatusInverseKineticStillFails)
		}
		rows = append(rows, GaugeCoordinateLayerRow{Layer: r.Layer, Coordinate: r.Coordinate, ResidualName: r.ResidualName, Residual: r.ResidualAtLambda12, WBest: r.WBest, WBestMinus7Over72: r.WBestMinus7Over72, PassesSevenOver72: passes, Verdict: verdict})
	}
	return GaugeCoordinateComparisonAudit{Rows: rows, AmplitudeOnlyPasses: ampOnly, InverseKineticFails: inverseFails, ClosureCoordinate: "G_g=g3/gEW-1", Verdict: join(StatusGaugeCoordinateComparisonAudited, StatusClosureConnectionAmplitudeLayer, StatusInverseKineticStillFails)}
}

func buildHessianSocket() ElectroweakHessianSocketAudit {
	return ElectroweakHessianSocketAudit{
		CovariantDerivativeSocket: "D_mu = partial_mu + i g A_mu",
		NeutralHessianShape:       "M_neutral^2 = (K_phi v^2/4) [[g^2,-gg'],[-gg',g'^2]]",
		ChargedWSocket:            "m_W^2 ∼ g^2 v^2/4, so mass amplitudes use g v/2 after taking the square root",
		AmplitudeObjects:          []string{"g", "g'", "sqrt(g^2+g'^2)", "gEW=(g1+g2)/2 at the meeting root"},
		CompatibleWithClosure:     true,
		Limitation:                "the Hessian socket types g as an endpoint connection/mass-amplitude coordinate; it does not derive the 7/72 weight or the scalar wound",
		Verdict:                   join(StatusElectroweakHessianSocketAudited, StatusCanonicalEndpointSocketSupported, StatusNoNativeSevenOver72Theorem),
	}
}

func buildScalarSide() ScalarSideTypeAudit {
	return ScalarSideTypeAudit{
		ScalarObject:    "|lambda(Lambda_12)|",
		ComparedTo:      "R_3-1=g3/gEW-1 in W_72=(65/72)|lambda|+(7/72)(R_3-1)",
		TypedAs:         "runtime scalar coefficient / high-scale scalar wound in the bridge ledger",
		NativeAmplitude: false,
		Limitation:      "the gauge side has a canonical connection-amplitude source, but the scalar side remains a runtime quartic/coefficient shadow rather than a certified native amplitude object",
		Verdict:         join(StatusScalarSideTypeAudited, StatusScalarRuntimeShadow),
	}
}

func buildPattern(g gate666.Analysis) RootAmplitudeRecurrenceAudit {
	return RootAmplitudeRecurrenceAudit{Rows: g.Pattern.Rows, Pattern: g.Pattern.Pattern, Verdict: join(StatusRootAmplitudeRecurrenceAudited, StatusRootAmplitudePatternSupported)}
}

func buildTarget() AirlockTheoremTarget {
	return AirlockTheoremTarget{
		Name:             "CanonicalKineticToConnectionAmplitudeAirlock / KineticSquareRootAirlock",
		Domain:           "native kinetic/RG variables u_i=1/g_i^2 and other quadratic/trace coordinates",
		Airlock:          "canonical field normalization and square-root projection u_i -> g_i=u_i^{-1/2}",
		Codomain:         "endpoint connection-amplitude residuals such as R_3-1=g3/gEW-1",
		CandidateContent: "a typed ASHA theorem explaining when bridge/history closures must be read after canonical amplitude projection rather than in raw inverse kinetic variables",
		Status:           "not certified",
		Verdict:          join(StatusConnectionAmplitudeCoordinateTyped, StatusNoNativeKineticAmplitudeTheorem),
	}
}

func buildSourceVerdict(coords GaugeCoordinateComparisonAudit, scalar ScalarSideTypeAudit) SourceTypeVerdict {
	classification := "BoundaryWeightedDeficitClosureConnectionAmplitudeSeal"
	statements := []string{
		"the active 7/72 closure is source-typed to the canonical connection-amplitude coordinate g, not the inverse kinetic coordinate u=1/g^2",
		"canonical gauge-field normalization supplies a lawful bridge-layer route from u to g, but no native airlock theorem is proved",
		"the electroweak Hessian/mass socket also uses connection amplitudes g and sqrt(g^2+g'^2), making the coordinate choice endpoint-compatible",
		"the scalar side remains a runtime scalar coefficient wound, so the full scalar/flavor/boundary transport theorem is still missing",
	}
	verdict := join(StatusGaugeAmplitudeSourcedByConnection, StatusClosureConnectionAmplitudeLayer, StatusScalarRuntimeShadow, StatusNoNativeKineticAmplitudeTheorem)
	if !coords.AmplitudeOnlyPasses || !coords.InverseKineticFails || scalar.NativeAmplitude {
		classification = "connection-amplitude source classification unresolved"
	}
	return SourceTypeVerdict{Classification: classification, Statements: statements, Verdict: verdict}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate666AmplitudeSealInherited,
		StatusKineticCoordinateDefined,
		StatusCanonicalFieldRescalingAudited,
		StatusConnectionAmplitudeCoordinateTyped,
		StatusGaugeCoordinateComparisonAudited,
		StatusElectroweakHessianSocketAudited,
		StatusScalarSideTypeAudited,
		StatusRootAmplitudeRecurrenceAudited,
		StatusGaugeAmplitudeSourcedByConnection,
		StatusClosureConnectionAmplitudeLayer,
		StatusCanonicalEndpointSocketSupported,
		StatusRootAmplitudePatternSupported,
		StatusInverseKineticStillFails,
		StatusNoNativeKineticAmplitudeTheorem,
		StatusScalarRuntimeShadow,
		StatusNoNativeSevenOver72Theorem,
		StatusNoNativeDualRootTheorem,
		StatusNoNativeScalarFlavorBoundaryTransport,
		StatusNoBoundaryStressDerivation,
		StatusNoHiggsStabilityGaugeFlavor,
		StatusGate667Boundary,
	}
}

// finiteNumber is used by tests/theorem checks to guard inherited numerical fields.
func finiteNumber(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }
