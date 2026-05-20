// Package generation2radialhiggsselfcouplingandnormalizationaudit implements
// Gate 773: Radial Higgs Self-Coupling Boundary Audit.
//
// Gate 772 wrote the locally normalized sealed Higgs potential as
// V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2. Gate 773 chooses a supplied
// vacuum representative, expands the radial/unitary-gauge mode x=(v+h)u_rad,
// computes the h^2, h^3, and h^4 coefficients, records the Feynman-rule
// normalization, and preserves the VEV, EWSB, pole-mass, Yukawa, scalar-runtime,
// and HistoryLoop firewalls.
package generation2radialhiggsselfcouplingandnormalizationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE773-RADIAL-HIGGS-SELF-COUPLING-BOUNDARY-AUDIT"

	StatusGate772CompletedSquarePotentialInherited = "PASS_GATE772_COMPLETED_SQUARE_POTENTIAL_INHERITED"
	StatusRadialFieldExpansionDefined              = "PASS_RADIAL_FIELD_EXPANSION_DEFINED"
	StatusLocalPotentialExpanded                   = "PASS_LOCAL_POTENTIAL_EXPANDED"
	StatusTreeRadialMassReconfirmed                = "PASS_TREE_RADIAL_MASS_RECONFIRMED"
	StatusSelfCouplingConventionSeparationAudited  = "PASS_SELF_COUPLING_CONVENTION_SEPARATION_AUDITED"
	StatusNumericalSelfCouplingLedgerComputed      = "PASS_NUMERICAL_SELF_COUPLING_LEDGER_COMPUTED"
	StatusPhysicalFirewallsEnforced                = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusSealedPotentialDeterminesTreeSelfCouplings = "CONDITIONAL_SUPPORT_SEALED_COMPLETED_SQUARE_POTENTIAL_DETERMINES_TREE_RADIAL_SELF_COUPLINGS"
	StatusTreeMassAndSelfCouplingsSameExpansion      = "CONDITIONAL_SUPPORT_TREE_MASS_AND_SELF_COUPLINGS_FOLLOW_FROM_SAME_RADIAL_EXPANSION"

	StatusRadialExpansionNotNativeHiggsTheorem   = "FAILED_ROUTE_RADIAL_EXPANSION_NOT_NATIVE_HIGGS_THEOREM"
	StatusTreeSelfCouplingsNotPhysicalMeasured   = "FAILED_ROUTE_TREE_SELF_COUPLINGS_NOT_PHYSICAL_MEASURED_COUPLINGS"
	StatusTreeProxyNotPoleMass                   = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoNativeVEVTheorem                     = "FAILED_ROUTE_NO_NATIVE_VEV_THEOREM"
	StatusNoNativeEWSBTheorem                    = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoYukawaOperatorOrEigenvalue           = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem      = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem         = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusGate773RadialHiggsSelfCouplingBoundary = "FIREWALL_PRESERVED_GATE773_RADIAL_HIGGS_SELF_COUPLING_BOUNDARY"
)

const (
	lambdaRuntimeEff = 0.12965256505060754
	vevConventionGeV = 246.2196508
)

type Gate772Inheritance struct {
	Inherited              bool
	CompletedSquareForm    string
	RealFourCoordinateForm string
	QuarticAirlock         string
	VEVSeal                string
	LambdaRuntimeEff       float64
	VEVGeV                 float64
	NativeHiggsTheorem     bool
	Verdict                string
}

type RadialFieldExpansion struct {
	VacuumRepresentative string
	UnitRadialCondition  string
	GaugeChoice          string
	NormExpression       string
	ExpansionExpression  string
	RadialGaugeNative    bool
	NativeEWSBTheorem    bool
	Verdict              string
}

type LocalPotentialExpansion struct {
	StartingPotential  string
	ExpandedPotential  string
	A2Formula          string
	A3Formula          string
	A4Formula          string
	A2GeV2             float64
	A3GeV              float64
	A4                 float64
	AlgebraicExpansion bool
	NativeHiggsTheorem bool
	Verdict            string
}

type TreeRadialMass struct {
	CanonicalMassConvention string
	MassSquaredFormula      string
	LambdaRuntimeEff        float64
	VEVGeV                  float64
	MassSquaredGeV2         float64
	MassGeV                 float64
	PoleMassTheorem         bool
	Verdict                 string
}

type SelfCouplingConventions struct {
	PotentialCoefficientConvention string
	FeynmanRuleConvention          string
	Lambda3Formula                 string
	Lambda4Formula                 string
	Lambda3Alternative             string
	Lambda4Alternative             string
	ConventionSeparated            bool
	PhysicalMeasuredCouplings      bool
	Verdict                        string
}

type NumericalLedger struct {
	LambdaRuntimeEff float64
	VEVGeV           float64
	A2GeV2           float64
	A3GeV            float64
	A4               float64
	MassGeV          float64
	Lambda3GeV       float64
	Lambda4          float64
	LedgerComputed   bool
	Verdict          string
}

type SourceTypeInterpretation struct {
	LambdaRuntimeEff string
	V                string
	H                string
	MassProxy        string
	Lambda3Lambda4   string
	Interpretation   string
	Verdict          string
}

type Firewalls struct {
	Audited                         bool
	RadialExpansionNativeHiggs      bool
	TreeSelfCouplingsMeasured       bool
	TreeProxyPoleMass               bool
	RadialGaugeNativeEWSB           bool
	LambdaRuntimeIndependentTheorem bool
	VEVNativeTheorem                bool
	YukawaOperatorOrEigenvalue      bool
	HistoryLoopUnitTheorem          bool
	Verdict                         string
}

type Analysis struct {
	Gate772     Gate772Inheritance
	Radial      RadialFieldExpansion
	Expansion   LocalPotentialExpansion
	Mass        TreeRadialMass
	Conventions SelfCouplingConventions
	Numerical   NumericalLedger
	SourceTypes SourceTypeInterpretation
	Firewalls   Firewalls
	Truth       string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		clone := *cache
		return &clone, nil
	}

	if math.IsNaN(lambdaRuntimeEff) || math.IsInf(lambdaRuntimeEff, 0) || lambdaRuntimeEff <= 0 {
		return nil, fmt.Errorf("invalid lambda_runtime_eff ledger: %.17g", lambdaRuntimeEff)
	}
	if math.IsNaN(vevConventionGeV) || math.IsInf(vevConventionGeV, 0) || vevConventionGeV <= 0 {
		return nil, fmt.Errorf("invalid VEV convention ledger: %.17g", vevConventionGeV)
	}

	v2 := vevConventionGeV * vevConventionGeV
	a2 := lambdaRuntimeEff * v2
	a3 := lambdaRuntimeEff * vevConventionGeV
	a4 := lambdaRuntimeEff / 4
	massSquared := 2 * a2
	mass := math.Sqrt(massSquared)
	lambda3 := 6 * lambdaRuntimeEff * vevConventionGeV
	lambda4 := 6 * lambdaRuntimeEff

	a := &Analysis{
		Gate772: Gate772Inheritance{
			Inherited:              true,
			CompletedSquareForm:    "V_local(phi)=lambda_runtime_eff(phi^dagger phi-v^2/2)^2",
			RealFourCoordinateForm: "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2",
			QuarticAirlock:         "HiggsQuarticRuntimeCoefficientSeal",
			VEVSeal:                "VEVConventionSeal",
			LambdaRuntimeEff:       lambdaRuntimeEff,
			VEVGeV:                 vevConventionGeV,
			NativeHiggsTheorem:     false,
			Verdict:                StatusGate772CompletedSquarePotentialInherited,
		},
		Radial: RadialFieldExpansion{
			VacuumRepresentative: "x_0=v u_rad",
			UnitRadialCondition:  "||u_rad||=1",
			GaugeChoice:          "x=(v+h)u_rad",
			NormExpression:       "||x||^2=(v+h)^2",
			ExpansionExpression:  "V_local(h)=(lambda_runtime_eff/4)((v+h)^2-v^2)^2",
			RadialGaugeNative:    false,
			NativeEWSBTheorem:    false,
			Verdict:              StatusRadialFieldExpansionDefined,
		},
		Expansion: LocalPotentialExpansion{
			StartingPotential:  "V_local(h)=(lambda_runtime_eff/4)[(v+h)^2-v^2]^2",
			ExpandedPotential:  "V_local(h)=lambda_runtime_eff v^2 h^2+lambda_runtime_eff v h^3+(lambda_runtime_eff/4)h^4",
			A2Formula:          "A_2=lambda_runtime_eff v^2",
			A3Formula:          "A_3=lambda_runtime_eff v",
			A4Formula:          "A_4=lambda_runtime_eff/4",
			A2GeV2:             a2,
			A3GeV:              a3,
			A4:                 a4,
			AlgebraicExpansion: true,
			NativeHiggsTheorem: false,
			Verdict:            StatusLocalPotentialExpanded,
		},
		Mass: TreeRadialMass{
			CanonicalMassConvention: "V(h) contains (1/2)m_h^2 h^2",
			MassSquaredFormula:      "m_H_tree_proxy^2=2lambda_runtime_eff v^2",
			LambdaRuntimeEff:        lambdaRuntimeEff,
			VEVGeV:                  vevConventionGeV,
			MassSquaredGeV2:         massSquared,
			MassGeV:                 mass,
			PoleMassTheorem:         false,
			Verdict:                 StatusTreeRadialMassReconfirmed,
		},
		Conventions: SelfCouplingConventions{
			PotentialCoefficientConvention: "V(h)=A_2h^2+A_3h^3+A_4h^4",
			FeynmanRuleConvention:          "V(h)=(1/2)m_h^2h^2+(1/3!)lambda_3 h^3+(1/4!)lambda_4 h^4",
			Lambda3Formula:                 "lambda_3=6lambda_runtime_eff v",
			Lambda4Formula:                 "lambda_4=6lambda_runtime_eff",
			Lambda3Alternative:             "lambda_3=3m_h^2/v",
			Lambda4Alternative:             "lambda_4=3m_h^2/v^2",
			ConventionSeparated:            true,
			PhysicalMeasuredCouplings:      false,
			Verdict:                        StatusSelfCouplingConventionSeparationAudited,
		},
		Numerical: NumericalLedger{
			LambdaRuntimeEff: lambdaRuntimeEff,
			VEVGeV:           vevConventionGeV,
			A2GeV2:           a2,
			A3GeV:            a3,
			A4:               a4,
			MassGeV:          mass,
			Lambda3GeV:       lambda3,
			Lambda4:          lambda4,
			LedgerComputed:   true,
			Verdict:          StatusNumericalSelfCouplingLedgerComputed,
		},
		SourceTypes: SourceTypeInterpretation{
			LambdaRuntimeEff: "sealed bridge quartic after Gate770 airlock",
			V:                "supplied VEV convention, not native VEV theorem",
			H:                "radial fluctuation coordinate after choosing a vacuum representative and radial gauge",
			MassProxy:        "radial Hessian tree proxy, not pole mass",
			Lambda3Lambda4:   "tree-level self-coupling proxies under the chosen convention",
			Interpretation:   "Gate773 expands only the locally sealed potential; it normalizes tree radial coefficients but does not promote them to measured pole or self-coupling observables.",
			Verdict:          "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED",
		},
		Firewalls: Firewalls{
			Audited:                         true,
			RadialExpansionNativeHiggs:      false,
			TreeSelfCouplingsMeasured:       false,
			TreeProxyPoleMass:               false,
			RadialGaugeNativeEWSB:           false,
			LambdaRuntimeIndependentTheorem: false,
			VEVNativeTheorem:                false,
			YukawaOperatorOrEigenvalue:      false,
			HistoryLoopUnitTheorem:          false,
			Verdict:                         StatusGate773RadialHiggsSelfCouplingBoundary,
		},
		Truth: "Gate773 expands the sealed completed-square Higgs potential around a supplied radial representative and derives only tree-level convention-normalized mass, cubic, and quartic self-coupling proxies; it does not derive the VEV, EWSB, pole mass, measured self-couplings, Yukawa ledger, or HistoryLoopUnit.",
	}

	cache = a
	clone := *a
	return &clone, nil
}

func Statuses() []string {
	return []string{
		StatusGate772CompletedSquarePotentialInherited,
		StatusRadialFieldExpansionDefined,
		StatusLocalPotentialExpanded,
		StatusTreeRadialMassReconfirmed,
		StatusSelfCouplingConventionSeparationAudited,
		StatusNumericalSelfCouplingLedgerComputed,
		StatusPhysicalFirewallsEnforced,
		StatusSealedPotentialDeterminesTreeSelfCouplings,
		StatusTreeMassAndSelfCouplingsSameExpansion,
		StatusRadialExpansionNotNativeHiggsTheorem,
		StatusTreeSelfCouplingsNotPhysicalMeasured,
		StatusTreeProxyNotPoleMass,
		StatusNoNativeVEVTheorem,
		StatusNoNativeEWSBTheorem,
		StatusNoYukawaOperatorOrEigenvalue,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusGate773RadialHiggsSelfCouplingBoundary,
	}
}

func FormatGate772(g Gate772Inheritance) string {
	return fmt.Sprintf("inherited=%t square=%s real=%s airlock=%s vevSeal=%s lambda=%.17g v=%.10f nativeHiggs=%t verdict=%s", g.Inherited, g.CompletedSquareForm, g.RealFourCoordinateForm, g.QuarticAirlock, g.VEVSeal, g.LambdaRuntimeEff, g.VEVGeV, g.NativeHiggsTheorem, g.Verdict)
}

func FormatRadial(r RadialFieldExpansion) string {
	return fmt.Sprintf("vacuum=%s unit=%s gauge=%s norm=%s expansion=%s nativeGauge=%t nativeEWSB=%t verdict=%s", r.VacuumRepresentative, r.UnitRadialCondition, r.GaugeChoice, r.NormExpression, r.ExpansionExpression, r.RadialGaugeNative, r.NativeEWSBTheorem, r.Verdict)
}

func FormatExpansion(e LocalPotentialExpansion) string {
	return fmt.Sprintf("start=%s expanded=%s A2=%s A3=%s A4=%s A2val=%.12f A3val=%.12f A4val=%.17g algebraic=%t nativeHiggs=%t verdict=%s", e.StartingPotential, e.ExpandedPotential, e.A2Formula, e.A3Formula, e.A4Formula, e.A2GeV2, e.A3GeV, e.A4, e.AlgebraicExpansion, e.NativeHiggsTheorem, e.Verdict)
}

func FormatMass(m TreeRadialMass) string {
	return fmt.Sprintf("convention=%s formula=%s lambda=%.17g v=%.10f m2=%.12f m=%.12f poleMass=%t verdict=%s", m.CanonicalMassConvention, m.MassSquaredFormula, m.LambdaRuntimeEff, m.VEVGeV, m.MassSquaredGeV2, m.MassGeV, m.PoleMassTheorem, m.Verdict)
}

func FormatConventions(c SelfCouplingConventions) string {
	return fmt.Sprintf("potential=%s feynman=%s lambda3=%s lambda4=%s alt3=%s alt4=%s separated=%t measured=%t verdict=%s", c.PotentialCoefficientConvention, c.FeynmanRuleConvention, c.Lambda3Formula, c.Lambda4Formula, c.Lambda3Alternative, c.Lambda4Alternative, c.ConventionSeparated, c.PhysicalMeasuredCouplings, c.Verdict)
}

func FormatNumerical(n NumericalLedger) string {
	return fmt.Sprintf("lambda=%.17g v=%.10f A2=%.12f A3=%.12f A4=%.17g m=%.12f lambda3=%.12f lambda4=%.17g computed=%t verdict=%s", n.LambdaRuntimeEff, n.VEVGeV, n.A2GeV2, n.A3GeV, n.A4, n.MassGeV, n.Lambda3GeV, n.Lambda4, n.LedgerComputed, n.Verdict)
}

func FormatSourceTypes(s SourceTypeInterpretation) string {
	return strings.Join([]string{s.LambdaRuntimeEff, s.V, s.H, s.MassProxy, s.Lambda3Lambda4, s.Interpretation, s.Verdict}, " | ")
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("audited=%t radialNative=%t measuredSelfCouplings=%t poleMass=%t gaugeEWSB=%t independentRuntime=%t nativeVEV=%t yukawa=%t historyLoop=%t verdict=%s", f.Audited, f.RadialExpansionNativeHiggs, f.TreeSelfCouplingsMeasured, f.TreeProxyPoleMass, f.RadialGaugeNativeEWSB, f.LambdaRuntimeIndependentTheorem, f.VEVNativeTheorem, f.YukawaOperatorOrEigenvalue, f.HistoryLoopUnitTheorem, f.Verdict)
}
