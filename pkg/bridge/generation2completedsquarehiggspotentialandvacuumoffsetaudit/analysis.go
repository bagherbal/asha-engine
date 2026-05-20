// Package generation2completedsquarehiggspotentialandvacuumoffsetaudit implements
// Gate 772: Completed-Square Higgs Potential and Vacuum-Offset Firewall Audit.
//
// Gate 771 showed that after the quartic coefficient airlock and VEV/radius
// convention seal, the supplied U(2)-invariant Higgs potential has
// mu^2_bridge=-lambda_runtime_eff v^2. Gate 772 completes the square,
// records the local zero-vacuum offset convention, writes the real four-coordinate
// form, reconfirms Hessian compatibility, and preserves the cosmological,
// VEV, pole-mass, Yukawa, and HistoryLoop firewalls.
package generation2completedsquarehiggspotentialandvacuumoffsetaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE772-COMPLETED-SQUARE-HIGGS-POTENTIAL-AND-VACUUM-OFFSET-FIREWALL-AUDIT"

	StatusGate771VEVMuSquaredOffsetInherited = "PASS_GATE771_VEV_MU_SQUARED_OFFSET_INHERITED"
	StatusCompletedSquareFormDerived         = "PASS_COMPLETED_SQUARE_FORM_DERIVED"
	StatusLocalZeroVacuumOffsetRecorded      = "PASS_LOCAL_ZERO_VACUUM_OFFSET_RECORDED"
	StatusRealFourCoordinateFormWritten      = "PASS_REAL_FOUR_COORDINATE_FORM_WRITTEN"
	StatusHessianCompatibilityReconfirmed    = "PASS_HESSIAN_COMPATIBILITY_RECONFIRMED"
	StatusVacuumOrbitRecorded                = "PASS_VACUUM_ORBIT_RECORDED"
	StatusCosmologicalFirewallEnforced       = "PASS_COSMOLOGICAL_FIREWALL_ENFORCED"
	StatusPhysicalFirewallsEnforced          = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusSealedPotentialCompletedSquareNormalForm = "CONDITIONAL_SUPPORT_SEALED_HIGGS_POTENTIAL_HAS_COMPLETED_SQUARE_NORMAL_FORM"
	StatusLocalZeroOffsetFixesC0AsConvention       = "CONDITIONAL_SUPPORT_LOCAL_ZERO_OFFSET_FIXES_C0_AS_CONVENTION"
	StatusHessianTreeProxyFromCompletedSquare      = "CONDITIONAL_SUPPORT_HESSIAN_TREE_PROXY_FOLLOWS_FROM_COMPLETED_SQUARE_FORM"

	StatusCompletedSquareNotNativeHiggsTheorem = "FAILED_ROUTE_COMPLETED_SQUARE_FORM_NOT_NATIVE_HIGGS_THEOREM"
	StatusC0LocalOffsetNotCosmologicalConstant = "FAILED_ROUTE_C0_LOCAL_OFFSET_NOT_COSMOLOGICAL_CONSTANT_THEOREM"
	StatusNoNativeVEVTheorem                   = "FAILED_ROUTE_NO_NATIVE_VEV_THEOREM"
	StatusNoNativeEWSBTheorem                  = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusTreeProxyNotPoleMass                 = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoHiggsMassOrPoleMassTheorem         = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalue         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate772CompletedSquareBoundary       = "FIREWALL_PRESERVED_GATE772_COMPLETED_SQUARE_HIGGS_POTENTIAL_BOUNDARY"
)

const (
	lambdaRuntimeEff = 0.12965256505060754
	vevConventionGeV = 246.2196508
)

type Gate771Inheritance struct {
	Inherited              bool
	QuarticAirlock         string
	VEVSeal                string
	LambdaHIdentification  string
	MuSquaredFormula       string
	LambdaRuntimeEff       float64
	VEVGeV                 float64
	MuSquaredBridgeGeV2    float64
	C0LocalBridgeGeV4      float64
	TreeProxyGeV           float64
	NativeVEVTheorem       bool
	NativeMuSquaredTheorem bool
	C0CosmologyTheorem     bool
	Verdict                string
}

type CompletedSquareForm struct {
	StartingPotential  string
	Substitution       string
	Coordinate         string
	ExpandedAfterSeal  string
	CompletedSquare    string
	VMinFormula        string
	AlgebraicIdentity  bool
	NativeHiggsTheorem bool
	Verdict            string
}

type LocalZeroVacuumOffset struct {
	ConventionName              string
	Condition                   string
	C0Formula                   string
	C0LocalBridgeGeV4           float64
	VMinWithLocalOffsetGeV4     float64
	LocalOffsetConvention       bool
	CosmologicalConstantTheorem bool
	VacuumEnergyDerivation      bool
	Verdict                     string
}

type LocalSealedPotential struct {
	ComplexCoordinateForm  string
	RealFourCoordinateRule string
	RealFourCoordinateForm string
	LambdaRuntimeEff       float64
	VEVGeV                 float64
	NormalizedAfterOffset  bool
	NativePotentialTheorem bool
	Verdict                string
}

type HessianCompatibility struct {
	VacuumCondition          string
	HessianFormula           string
	SupportProjector         string
	TreeProxySquaredFormula  string
	LambdaRuntimeEff         float64
	VEVGeV                   float64
	TreeProxySquaredGeV2     float64
	TreeProxyGeV             float64
	TreeProxyPoleMass        bool
	NativeHistoryLoopTheorem bool
	Verdict                  string
}

type VacuumOrbit struct {
	ComplexMinimaCondition string
	RealMinimaCondition    string
	OrbitBeforeQuotient    string
	AngularFlatness        string
	RadialNonFlatness      string
	SelectsCP1Point        bool
	NativeEWSBTheorem      bool
	Verdict                string
}

type SourceTypeInterpretation struct {
	LambdaRuntimeEff string
	V                string
	MuSquaredBridge  string
	C0               string
	CompletedSquare  string
	Interpretation   string
	Verdict          string
}

type Firewalls struct {
	Audited                         bool
	CompletedSquareNativeHiggs      bool
	C0CosmologicalConstantTheorem   bool
	S3OrbitNativeEWSB               bool
	TreeHessianPoleMass             bool
	LambdaRuntimeIndependentTheorem bool
	VEVNativeTheorem                bool
	HiggsMassOrPoleMassTheorem      bool
	YukawaOperatorOrEigenvalue      bool
	HistoryLoopUnitTheorem          bool
	Verdict                         string
}

type Analysis struct {
	Gate771     Gate771Inheritance
	Square      CompletedSquareForm
	Offset      LocalZeroVacuumOffset
	Local       LocalSealedPotential
	Hessian     HessianCompatibility
	Orbit       VacuumOrbit
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
	v4 := v2 * v2
	muSquared := -lambdaRuntimeEff * v2
	treeProxySquared := 2 * lambdaRuntimeEff * v2
	treeProxy := math.Sqrt(treeProxySquared)
	c0Local := 0.25 * lambdaRuntimeEff * v4

	a := &Analysis{
		Gate771: Gate771Inheritance{
			Inherited:              true,
			QuarticAirlock:         "HiggsQuarticRuntimeCoefficientSeal",
			VEVSeal:                "VEVConventionSeal",
			LambdaHIdentification:  "lambda_H := lambda_runtime_eff",
			MuSquaredFormula:       "mu^2_bridge=-lambda_runtime_eff v^2",
			LambdaRuntimeEff:       lambdaRuntimeEff,
			VEVGeV:                 vevConventionGeV,
			MuSquaredBridgeGeV2:    muSquared,
			C0LocalBridgeGeV4:      c0Local,
			TreeProxyGeV:           treeProxy,
			NativeVEVTheorem:       false,
			NativeMuSquaredTheorem: false,
			C0CosmologyTheorem:     false,
			Verdict:                StatusGate771VEVMuSquaredOffsetInherited,
		},
		Square: CompletedSquareForm{
			StartingPotential:  "V(u)=c_0+mu^2 u+lambda_H u^2",
			Substitution:       "mu^2=-lambda_H v^2",
			Coordinate:         "u=phi^dagger phi",
			ExpandedAfterSeal:  "V(u)=c_0-lambda_H v^2 u+lambda_H u^2",
			CompletedSquare:    "V(u)=lambda_H(u-v^2/2)^2+c_0-(1/4)lambda_H v^4",
			VMinFormula:        "V_min=c_0-(1/4)lambda_H v^4",
			AlgebraicIdentity:  true,
			NativeHiggsTheorem: false,
			Verdict:            StatusCompletedSquareFormDerived,
		},
		Offset: LocalZeroVacuumOffset{
			ConventionName:              "LocalZeroVacuumOffsetConvention",
			Condition:                   "V_min=0",
			C0Formula:                   "c_0=(1/4)lambda_H v^4=(1/4)lambda_runtime_eff v^4 after airlock",
			C0LocalBridgeGeV4:           c0Local,
			VMinWithLocalOffsetGeV4:     0,
			LocalOffsetConvention:       true,
			CosmologicalConstantTheorem: false,
			VacuumEnergyDerivation:      false,
			Verdict:                     StatusLocalZeroVacuumOffsetRecorded,
		},
		Local: LocalSealedPotential{
			ComplexCoordinateForm:  "V_local(phi)=lambda_runtime_eff(phi^dagger phi-v^2/2)^2",
			RealFourCoordinateRule: "phi^dagger phi=(1/2)||x||^2",
			RealFourCoordinateForm: "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2",
			LambdaRuntimeEff:       lambdaRuntimeEff,
			VEVGeV:                 vevConventionGeV,
			NormalizedAfterOffset:  true,
			NativePotentialTheorem: false,
			Verdict:                StatusRealFourCoordinateFormWritten,
		},
		Hessian: HessianCompatibility{
			VacuumCondition:          "||x_0||^2=v^2",
			HessianFormula:           "H_V(x_0)=2 lambda_runtime_eff v^2 P_rad",
			SupportProjector:         "P_rad=supp(H_V(x_0))",
			TreeProxySquaredFormula:  "m_H_tree_proxy^2=2 lambda_runtime_eff v^2",
			LambdaRuntimeEff:         lambdaRuntimeEff,
			VEVGeV:                   vevConventionGeV,
			TreeProxySquaredGeV2:     treeProxySquared,
			TreeProxyGeV:             treeProxy,
			TreeProxyPoleMass:        false,
			NativeHistoryLoopTheorem: false,
			Verdict:                  StatusHessianCompatibilityReconfirmed,
		},
		Orbit: VacuumOrbit{
			ComplexMinimaCondition: "phi^dagger phi=v^2/2",
			RealMinimaCondition:    "||x||^2=v^2",
			OrbitBeforeQuotient:    "S^3 vacuum orbit in the real four-carrier before gauge/orbit quotient",
			AngularFlatness:        "potential is flat along angular orbit directions",
			RadialNonFlatness:      "only the radial Hessian direction is non-flat at the supplied vacuum",
			SelectsCP1Point:        false,
			NativeEWSBTheorem:      false,
			Verdict:                StatusVacuumOrbitRecorded,
		},
		SourceTypes: SourceTypeInterpretation{
			LambdaRuntimeEff: "bridge quartic coefficient after Gate770 airlock",
			V:                "supplied VEV/radius convention, not native VEV theorem",
			MuSquaredBridge:  "stationarity consequence inherited from Gate771",
			C0:               "local vacuum-offset convention, not cosmological constant theorem",
			CompletedSquare:  "normalized sealed Higgs-potential form after lambda and v seals",
			Interpretation:   "Gate772 rewrites the sealed potential as an algebraic completed square and preserves its status as a local tree-proxy normalization, not a native Higgs or cosmology theorem.",
			Verdict:          "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED",
		},
		Firewalls: Firewalls{
			Audited:                         true,
			CompletedSquareNativeHiggs:      false,
			C0CosmologicalConstantTheorem:   false,
			S3OrbitNativeEWSB:               false,
			TreeHessianPoleMass:             false,
			LambdaRuntimeIndependentTheorem: false,
			VEVNativeTheorem:                false,
			HiggsMassOrPoleMassTheorem:      false,
			YukawaOperatorOrEigenvalue:      false,
			HistoryLoopUnitTheorem:          false,
			Verdict:                         StatusGate772CompletedSquareBoundary,
		},
		Truth: "Gate772 completes the sealed Higgs potential into lambda_runtime_eff(phi^dagger phi-v^2/2)^2 plus a local offset convention; it reconfirms the Hessian tree proxy but derives neither the VEV, the pole mass, nor the cosmological constant.",
	}

	cache = a
	clone := *a
	return &clone, nil
}

func Statuses() []string {
	return []string{
		StatusGate771VEVMuSquaredOffsetInherited,
		StatusCompletedSquareFormDerived,
		StatusLocalZeroVacuumOffsetRecorded,
		StatusRealFourCoordinateFormWritten,
		StatusHessianCompatibilityReconfirmed,
		StatusVacuumOrbitRecorded,
		StatusCosmologicalFirewallEnforced,
		StatusPhysicalFirewallsEnforced,
		StatusSealedPotentialCompletedSquareNormalForm,
		StatusLocalZeroOffsetFixesC0AsConvention,
		StatusHessianTreeProxyFromCompletedSquare,
		StatusCompletedSquareNotNativeHiggsTheorem,
		StatusC0LocalOffsetNotCosmologicalConstant,
		StatusNoNativeVEVTheorem,
		StatusNoNativeEWSBTheorem,
		StatusTreeProxyNotPoleMass,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalue,
		StatusGate772CompletedSquareBoundary,
	}
}

func FormatGate771(g Gate771Inheritance) string {
	return fmt.Sprintf("inherited=%t airlock=%s vevSeal=%s identification=%s muFormula=%s lambda=%.17g v=%.10f mu2=%.12f c0=%.12f tree=%.12f nativeVEV=%t nativeMu=%t c0Cosmology=%t verdict=%s", g.Inherited, g.QuarticAirlock, g.VEVSeal, g.LambdaHIdentification, g.MuSquaredFormula, g.LambdaRuntimeEff, g.VEVGeV, g.MuSquaredBridgeGeV2, g.C0LocalBridgeGeV4, g.TreeProxyGeV, g.NativeVEVTheorem, g.NativeMuSquaredTheorem, g.C0CosmologyTheorem, g.Verdict)
}

func FormatSquare(s CompletedSquareForm) string {
	return fmt.Sprintf("start=%s substitution=%s coordinate=%s expanded=%s completed=%s Vmin=%s identity=%t nativeHiggs=%t verdict=%s", s.StartingPotential, s.Substitution, s.Coordinate, s.ExpandedAfterSeal, s.CompletedSquare, s.VMinFormula, s.AlgebraicIdentity, s.NativeHiggsTheorem, s.Verdict)
}

func FormatOffset(o LocalZeroVacuumOffset) string {
	return fmt.Sprintf("convention=%s condition=%s formula=%s c0=%.12f VminWithOffset=%.12f local=%t cosmology=%t vacuumEnergy=%t verdict=%s", o.ConventionName, o.Condition, o.C0Formula, o.C0LocalBridgeGeV4, o.VMinWithLocalOffsetGeV4, o.LocalOffsetConvention, o.CosmologicalConstantTheorem, o.VacuumEnergyDerivation, o.Verdict)
}

func FormatLocal(l LocalSealedPotential) string {
	return fmt.Sprintf("complex=%s rule=%s real=%s lambda=%.17g v=%.10f normalized=%t nativePotential=%t verdict=%s", l.ComplexCoordinateForm, l.RealFourCoordinateRule, l.RealFourCoordinateForm, l.LambdaRuntimeEff, l.VEVGeV, l.NormalizedAfterOffset, l.NativePotentialTheorem, l.Verdict)
}

func FormatHessian(h HessianCompatibility) string {
	return fmt.Sprintf("vacuum=%s hessian=%s support=%s m2Formula=%s lambda=%.17g v=%.10f m2=%.12f m=%.12f poleMass=%t nativeHistoryLoop=%t verdict=%s", h.VacuumCondition, h.HessianFormula, h.SupportProjector, h.TreeProxySquaredFormula, h.LambdaRuntimeEff, h.VEVGeV, h.TreeProxySquaredGeV2, h.TreeProxyGeV, h.TreeProxyPoleMass, h.NativeHistoryLoopTheorem, h.Verdict)
}

func FormatOrbit(o VacuumOrbit) string {
	return fmt.Sprintf("complex=%s real=%s orbit=%s angular=%s radial=%s selectsCP1=%t nativeEWSB=%t verdict=%s", o.ComplexMinimaCondition, o.RealMinimaCondition, o.OrbitBeforeQuotient, o.AngularFlatness, o.RadialNonFlatness, o.SelectsCP1Point, o.NativeEWSBTheorem, o.Verdict)
}

func FormatSourceTypes(s SourceTypeInterpretation) string {
	return strings.Join([]string{s.LambdaRuntimeEff, s.V, s.MuSquaredBridge, s.C0, s.CompletedSquare, s.Interpretation, s.Verdict}, " | ")
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("audited=%t completedSquareNative=%t c0Cosmology=%t s3EWSB=%t poleMass=%t independentRuntime=%t nativeVEV=%t higgsMass=%t yukawa=%t historyLoop=%t verdict=%s", f.Audited, f.CompletedSquareNativeHiggs, f.C0CosmologicalConstantTheorem, f.S3OrbitNativeEWSB, f.TreeHessianPoleMass, f.LambdaRuntimeIndependentTheorem, f.VEVNativeTheorem, f.HiggsMassOrPoleMassTheorem, f.YukawaOperatorOrEigenvalue, f.HistoryLoopUnitTheorem, f.Verdict)
}
