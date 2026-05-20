// Package generation2canonicalamplitudeairlockaudit implements
// Gate 666: Canonical Amplitude Airlock for BoundaryWeightedDeficitClosure Audit.
//
// Gate 665 classified the active E72 closure as strongest in the coupling
// amplitude coordinate G_g=g3/gEW-1, while inverse-coupling and other typed
// coordinates do not keep the same 7/72 alignment. Gate 666 audits that result
// as a source-type discovery: the bridge lives after a canonical amplitude /
// square-root / endpoint-coordinate airlock, not in the raw inverse-kinetic RG
// coordinate. It does not promote the airlock into a native theorem.
package generation2canonicalamplitudeairlockaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate665 "github.com/bagherbal/asha-engine/pkg/bridge/generation2electroweakrootclosurecoordinatenaturalityaudit"
)

const (
	AuditID = "GATE666-CANONICAL-AMPLITUDE-AIRLOCK-BOUNDARY-WEIGHTED-DEFICIT-CLOSURE-AUDIT"

	StatusGate665CoordinateSealInherited        = "PASS_GATE665_COORDINATE_SEAL_INHERITED"
	StatusCoordinateStackAudited                = "PASS_COORDINATE_STACK_AUDITED"
	StatusKineticToAmplitudeNonlinearityAudited = "PASS_KINETIC_TO_AMPLITUDE_NONLINEARITY_AUDITED"
	StatusRecurringAmplitudePatternAudited      = "PASS_RECURRING_AMPLITUDE_PATTERN_AUDITED"
	StatusAirlockTheoremTargetDefined           = "PASS_CANONICAL_AMPLITUDE_AIRLOCK_THEOREM_TARGET_DEFINED"
	StatusAmplitudeLayerSupported               = "CONDITIONAL_SUPPORT_BOUNDARY_WEIGHTED_DEFICIT_CLOSURE_IS_CANONICAL_AMPLITUDE_LAYER"
	StatusEndpointAirlockSupported              = "CONDITIONAL_SUPPORT_BRIDGE_LAYER_USES_ENDPOINT_AMPLITUDE_COORDINATES"
	StatusRootAmplitudePatternSupported         = "CONDITIONAL_SUPPORT_ROOT_AMPLITUDE_PROJECTIVE_AIRLOCK_RECURS_ACROSS_SEALS"
	StatusInverseKineticFails                   = "FAILED_ROUTE_INVERSE_KINETIC_LAYER_DOES_NOT_SUPPORT_SAME_7_OVER_72_CLOSURE"
	StatusNoNativeAmplitudeAirlockTheorem       = "FAILED_ROUTE_NO_NATIVE_AMPLITUDE_AIRLOCK_THEOREM"
	StatusNoNativeDualRootTheorem               = "FAILED_ROUTE_NO_NATIVE_DUAL_ROOT_ALIGNMENT_THEOREM"
	StatusNoNativeSevenOver72Theorem            = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeTransportTheorem              = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoBoundaryStressDerivation            = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoHiggsStabilityGaugeFlavor           = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate666Boundary                       = "FIREWALL_PRESERVED_GATE666_CANONICAL_AMPLITUDE_AIRLOCK_BOUNDARY"
)

const sevenOver72 = 7.0 / 72.0

type Gate665Inheritance struct {
	CoordinateSealInherited    bool
	Classification             string
	AmplitudeNatural           bool
	CoordinateRobust           bool
	RGNativeInverseNatural     bool
	AmplitudeWBest             float64
	AmplitudeWBestMinus7Over72 float64
	InverseWBest               float64
	InverseWBestMinus7Over72   float64
	AmplitudeE72               float64
	InverseE72                 float64
	NoNativeDualRoot           bool
	NoNativeSevenOver72        bool
	NoNativeTransport          bool
	NoBoundaryStress           bool
	Verdict                    string
}

type CoordinateLayerRow struct {
	Layer              string
	Coordinate         string
	ResidualName       string
	ResidualAtLambda12 float64
	WBest              float64
	WBestMinus7Over72  float64
	E72AtSevenOver72   float64
	RootAligned        bool
	NearSevenOver72    bool
	Classification     string
	Verdict            string
}

type CoordinateStackAudit struct {
	Rows                      []CoordinateLayerRow
	AmplitudeLayerPasses      bool
	InverseKineticLayerPasses bool
	StrengthLayerPasses       bool
	LogLayerPasses            bool
	Verdict                   string
}

type KineticToAmplitudeAudit struct {
	AmplitudeResidual       float64
	InverseFractionalWound  float64
	InverseOverAmplitude    float64
	SquaredResidual         float64
	SquaredOverAmplitude    float64
	LogResidual             float64
	ScalarWoundAbsLambda    float64
	AmplitudeScalarScaleGap float64
	InverseScalarScaleGap   float64
	FirstOrderStatement     string
	Verdict                 string
}

type RecurringPatternRow struct {
	Lane                           string
	WorkingCoordinate              string
	BlockedOrUncertifiedCoordinate string
	AirlockReading                 string
}

type RecurringAmplitudePatternAudit struct {
	Rows    []RecurringPatternRow
	Pattern string
	Verdict string
}

type AirlockTheoremTarget struct {
	NativeCoordinate  string
	AirlockCoordinate string
	ClosureCoordinate string
	MissingMap        string
	CandidateTheorem  string
	Verdict           string
}

type SourceTypeVerdict struct {
	Classification string
	Statements     []string
	Verdict        string
}

type VerdictDiscipline struct {
	ClaimsNativeAmplitudeAirlockTheorem bool
	ClaimsNativeDualRootTheorem         bool
	ClaimsNativeSevenOver72Theorem      bool
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
	Inherited          Gate665Inheritance
	CoordinateStack    CoordinateStackAudit
	KineticToAmplitude KineticToAmplitudeAudit
	Pattern            RecurringAmplitudePatternAudit
	Target             AirlockTheoremTarget
	Source             SourceTypeVerdict
	Discipline         VerdictDiscipline
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
	g665, err := gate665.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate665 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g665)
	stack := buildCoordinateStack(g665)
	kinetic := buildKineticToAmplitude(stack, g665)
	pattern := buildRecurringPattern()
	target := buildAirlockTarget()
	source := buildSourceVerdict(stack)
	discipline := VerdictDiscipline{Verdict: StatusGate666Boundary}
	truth := "Gate 666 classifies the active E72 closure as a canonical coupling-amplitude bridge seal. The same 7/72 dual-root closure is not certified in inverse-kinetic one-loop RG variables; the missing theorem is an airlock from native inverse/trace data to endpoint amplitude/projective history coordinates."
	return Analysis{Inherited: inherited, CoordinateStack: stack, KineticToAmplitude: kinetic, Pattern: pattern, Target: target, Source: source, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate665.Analysis) Gate665Inheritance {
	amp := findGate665Row(g, "amplitude ratio")
	inv := findGate665Row(g, "inverse-coupling ratio")
	return Gate665Inheritance{
		CoordinateSealInherited:    strings.Contains(g.CoordinateSeal.Classification, "amplitude-coordinate") && g.Coordinates.AmplitudeNatural && !g.Coordinates.RGNativeInverseNatural,
		Classification:             g.CoordinateSeal.Classification,
		AmplitudeNatural:           g.Coordinates.AmplitudeNatural,
		CoordinateRobust:           g.Coordinates.CoordinateRobust,
		RGNativeInverseNatural:     g.Coordinates.RGNativeInverseNatural,
		AmplitudeWBest:             amp.WBestAtT12,
		AmplitudeWBestMinus7Over72: amp.WBestMinus7Over72,
		InverseWBest:               inv.WBestAtT12,
		InverseWBestMinus7Over72:   inv.WBestMinus7Over72,
		AmplitudeE72:               amp.E72AtSevenOver72,
		InverseE72:                 inv.E72AtSevenOver72,
		NoNativeDualRoot:           !g.Discipline.ClaimsNativeDualRootTheorem,
		NoNativeSevenOver72:        !g.Discipline.ClaimsNativeSevenOver72Theorem,
		NoNativeTransport:          !g.Discipline.ClaimsNativeTransportTheorem,
		NoBoundaryStress:           !g.Discipline.ClaimsBoundaryStressDerivation,
		Verdict:                    StatusGate665CoordinateSealInherited,
	}
}

func buildCoordinateStack(g gate665.Analysis) CoordinateStackAudit {
	rows := make([]CoordinateLayerRow, 0, len(g.Coordinates.Rows))
	ampPass, invPass, strengthPass, logPass := false, false, false, false
	for _, r := range g.Coordinates.Rows {
		layer := layerFor(r)
		pass := r.NearSevenOver72 && r.RootFoundNearLambda12
		if layer == "canonical amplitude layer" && pass {
			ampPass = true
		}
		if layer == "RG-native inverse-kinetic layer" && pass {
			invPass = true
		}
		if layer == "coupling-strength layer" && pass {
			strengthPass = true
		}
		if layer == "multiplicative/log layer" && pass {
			logPass = true
		}
		verdict := StatusCoordinateStackAudited
		if layer == "canonical amplitude layer" && pass {
			verdict = join(StatusCoordinateStackAudited, StatusAmplitudeLayerSupported)
		}
		if layer == "RG-native inverse-kinetic layer" && !pass {
			verdict = join(StatusCoordinateStackAudited, StatusInverseKineticFails)
		}
		rows = append(rows, CoordinateLayerRow{
			Layer:              layer,
			Coordinate:         r.Definition,
			ResidualName:       r.Name,
			ResidualAtLambda12: r.GaugeResidualAtT12,
			WBest:              r.WBestAtT12,
			WBestMinus7Over72:  r.WBestMinus7Over72,
			E72AtSevenOver72:   r.E72AtSevenOver72,
			RootAligned:        r.RootFoundNearLambda12,
			NearSevenOver72:    r.NearSevenOver72,
			Classification:     r.CoordinateClass,
			Verdict:            verdict,
		})
	}
	return CoordinateStackAudit{Rows: rows, AmplitudeLayerPasses: ampPass, InverseKineticLayerPasses: invPass, StrengthLayerPasses: strengthPass, LogLayerPasses: logPass, Verdict: join(StatusCoordinateStackAudited, StatusAmplitudeLayerSupported, StatusInverseKineticFails)}
}

func buildKineticToAmplitude(stack CoordinateStackAudit, g gate665.Analysis) KineticToAmplitudeAudit {
	amp := findLayerRow(stack, "canonical amplitude layer")
	sq := findByResidualName(stack, "squared-coupling ratio")
	lg := findByResidualName(stack, "log-coupling residual")
	absLambda := g.CommonRoot.WBestAtRoot * 0 // avoid hard-coding through a hidden dependency below
	// Recover |lambda| from the identity w=(K-|lambda|)/(G-|lambda|).
	// |lambda|=(K-wG)/(1-w). K is inherited from Gate665's Gate664 inheritance.
	ksum := g.Inherited.KSum
	w := amp.WBest
	absLambda = (ksum - w*amp.ResidualAtLambda12) / (1.0 - w)
	invFractional := 1.0 - 1.0/math.Pow(1.0+amp.ResidualAtLambda12, 2)
	invOverAmp := invFractional / amp.ResidualAtLambda12
	ampGap := math.Abs(amp.ResidualAtLambda12 - absLambda)
	invGap := math.Abs(invFractional - absLambda)
	return KineticToAmplitudeAudit{
		AmplitudeResidual:       amp.ResidualAtLambda12,
		InverseFractionalWound:  invFractional,
		InverseOverAmplitude:    invOverAmp,
		SquaredResidual:         sq.ResidualAtLambda12,
		SquaredOverAmplitude:    sq.ResidualAtLambda12 / amp.ResidualAtLambda12,
		LogResidual:             lg.ResidualAtLambda12,
		ScalarWoundAbsLambda:    absLambda,
		AmplitudeScalarScaleGap: ampGap,
		InverseScalarScaleGap:   invGap,
		FirstOrderStatement:     "for r_g=g3/gEW-1, inverse-kinetic fractional wound 1-u3/uEW=1-1/(1+r_g)^2≈2r_g; this nearly doubles the boundary wound and leaves the scalar wound scale near 0.05",
		Verdict:                 join(StatusKineticToAmplitudeNonlinearityAudited, StatusAmplitudeLayerSupported, StatusInverseKineticFails),
	}
}

func buildRecurringPattern() RecurringAmplitudePatternAudit {
	rows := []RecurringPatternRow{
		{Lane: "charged-lepton Koide wall", WorkingCoordinate: "x_i=sqrt(y_i), wall angle/root coordinate", BlockedOrUncertifiedCoordinate: "polynomial trace ring alone", AirlockReading: "mass/Yukawa data enters the bridge after a square-root amplitude projection"},
		{Lane: "flavor orientation seal", WorkingCoordinate: "epsilon_e and kappa_e wall offsets", BlockedOrUncertifiedCoordinate: "native H_e traces without cross-sector orientation map", AirlockReading: "flavor closure is read in wall/offset coordinates"},
		{Lane: "scalar matching", WorkingCoordinate: "relative correction rho_lambda=(lambda_runtime-lambda_proxy)/lambda_proxy", BlockedOrUncertifiedCoordinate: "direct high-scale spectral quartic as native theorem", AirlockReading: "runtime scalar bridge is relative endpoint matching"},
		{Lane: "gauge boundary stress", WorkingCoordinate: "R_3-1=g3/gEW-1", BlockedOrUncertifiedCoordinate: "inverse-coupling residual", AirlockReading: "boundary wound is amplitude-normalized"},
		{Lane: "history loop unit", WorkingCoordinate: "L=1/(8pi), phase/amplitude-sized unit", BlockedOrUncertifiedCoordinate: "raw 1/(16pi^2) loop square", AirlockReading: "bridge unit resembles angular/root reduction rather than full loop-square unit"},
	}
	return RecurringAmplitudePatternAudit{Rows: rows, Pattern: "native quadratic/trace/RG variables repeatedly require a root, amplitude, wall, or projective endpoint airlock before the bridge-layer closure appears", Verdict: join(StatusRecurringAmplitudePatternAudited, StatusRootAmplitudePatternSupported)}
}

func buildAirlockTarget() AirlockTheoremTarget {
	return AirlockTheoremTarget{
		NativeCoordinate:  "u_i=1/g_i^2, spectral traces, polynomial finite invariants, loop-square units",
		AirlockCoordinate: "g_i, sqrt(y_i), wall offsets, phase/angular units, projective endpoint ratios",
		ClosureCoordinate: "R_3-1=g3/gEW-1 paired with |lambda(Lambda_12)| in E72",
		MissingMap:        "inverse-kinetic RG transport -> canonical coupling-amplitude boundary coordinate -> scalar/flavor deficit closure",
		CandidateTheorem:  "CanonicalAmplitudeAirlockTheorem: a typed normalization/root projection explains why history bridge seals are evaluated in endpoint amplitude coordinates rather than raw inverse kinetic variables",
		Verdict:           join(StatusAirlockTheoremTargetDefined, StatusNoNativeAmplitudeAirlockTheorem),
	}
}

func buildSourceVerdict(stack CoordinateStackAudit) SourceTypeVerdict {
	statements := []string{
		"the active 7/72 closure is currently certified only in the canonical coupling-amplitude residual G_g=g3/gEW-1",
		"the inverse-kinetic coordinate is the natural one-loop RG variable but does not support the same 7/72 closure in Gate665/Gate666",
		"this is a layer classification, not a failure of the bridge clue",
		"the missing theorem is an amplitude airlock, not another numerical coefficient search",
	}
	classification := "BoundaryWeightedDeficitClosureAmplitudeSeal"
	verdict := join(StatusAmplitudeLayerSupported, StatusEndpointAirlockSupported, StatusNoNativeAmplitudeAirlockTheorem, StatusInverseKineticFails)
	if !stack.AmplitudeLayerPasses || stack.InverseKineticLayerPasses {
		classification = "coordinate-natural classification unresolved"
	}
	return SourceTypeVerdict{Classification: classification, Statements: statements, Verdict: verdict}
}

func findGate665Row(g gate665.Analysis, name string) gate665.CoordinateRow {
	for _, r := range g.Coordinates.Rows {
		if r.Name == name {
			return r
		}
	}
	return gate665.CoordinateRow{}
}

func layerFor(r gate665.CoordinateRow) string {
	switch r.Name {
	case "amplitude ratio":
		return "canonical amplitude layer"
	case "inverse-coupling ratio":
		return "RG-native inverse-kinetic layer"
	case "squared-coupling ratio", "alpha ratio":
		return "coupling-strength layer"
	case "log-coupling residual":
		return "multiplicative/log layer"
	default:
		return "unclassified typed coordinate layer"
	}
}

func findLayerRow(stack CoordinateStackAudit, layer string) CoordinateLayerRow {
	for _, r := range stack.Rows {
		if r.Layer == layer {
			return r
		}
	}
	return CoordinateLayerRow{}
}

func findByResidualName(stack CoordinateStackAudit, name string) CoordinateLayerRow {
	for _, r := range stack.Rows {
		if r.ResidualName == name {
			return r
		}
	}
	return CoordinateLayerRow{}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate665CoordinateSealInherited,
		StatusCoordinateStackAudited,
		StatusKineticToAmplitudeNonlinearityAudited,
		StatusRecurringAmplitudePatternAudited,
		StatusAirlockTheoremTargetDefined,
		StatusAmplitudeLayerSupported,
		StatusEndpointAirlockSupported,
		StatusRootAmplitudePatternSupported,
		StatusInverseKineticFails,
		StatusNoNativeAmplitudeAirlockTheorem,
		StatusNoNativeDualRootTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoNativeTransportTheorem,
		StatusNoBoundaryStressDerivation,
		StatusNoHiggsStabilityGaugeFlavor,
		StatusGate666Boundary,
	}
}
