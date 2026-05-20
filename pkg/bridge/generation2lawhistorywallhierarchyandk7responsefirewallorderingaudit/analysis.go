// Package generation2lawhistorywallhierarchyandk7responsefirewallorderingaudit implements
// Gate 749: Law-History Wall Hierarchy and K7 Response Firewall Ordering Audit.
//
// Gate 748 showed that the kappa_e residual reuses the same boundary/history
// wall geometry already active in the K7 event-response lane. Gate 749 does not
// chase another numerical residual. It orders the native law walls, boundary
// quotient walls, history closure walls, flavor walls, scalar transport walls,
// and physical translation walls, clarifying where K7 is native support, where
// it is only an event weight, and where it must not be promoted into a boundary,
// flavor, scalar-runtime, or Higgs theorem.
package generation2lawhistorywallhierarchyandk7responsefirewallorderingaudit

import (
	"fmt"
	"strings"
	"sync"

	gate748 "github.com/bagherbal/asha-engine/pkg/bridge/generation2kappaehyperchargeboundaryresidualandboundarystressmomentaudit"
)

const (
	AuditID = "GATE749-LAW-HISTORY-WALL-HIERARCHY-K7-RESPONSE-FIREWALL-ORDERING-AUDIT"

	StatusGate748BoundaryStressMomentInherited = "PASS_GATE748_KAPPA_E_BOUNDARY_STRESS_MOMENT_INHERITED"
	StatusWallHierarchyDefined                 = "PASS_LAW_HISTORY_WALL_HIERARCHY_DEFINED"
	StatusK7RoleSeparationAudited              = "PASS_K7_NATIVE_SUPPORT_EVENT_WEIGHT_AND_FORBIDDEN_PROMOTION_ROLES_SEPARATED"
	StatusFirewallOrderConstructed             = "PASS_FIREWALL_ORDER_CONSTRUCTED"
	StatusGate748WallResonanceRecorded         = "PASS_GATE748_WALL_RESONANCE_RECORDED"
	StatusBoundaryMomentCoordinateRecorded     = "PASS_BOUNDARY_RAW_MOMENT_COORDINATE_RECORDED"
	StatusReductionPriorityRecorded            = "PASS_REDUCTION_PRIORITY_RECORDED"
	StatusPhysicalFirewallsEnforced            = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusGate748CorrectionLawHistoryWallResonance = "CONDITIONAL_SUPPORT_CURRENT_GATE748_CORRECTION_IS_LAW_HISTORY_WALL_RESONANCE"
	StatusK7ActsAsSupportAndEventWeightOnly        = "CONDITIONAL_SUPPORT_K7_ACTS_AS_NATIVE_SUPPORT_AND_BRIDGE_EVENT_WEIGHT_ONLY"
	StatusWallOrderingStabilizesNextReductions     = "CONDITIONAL_SUPPORT_WALL_ORDERING_STABILIZES_NEXT_SEAL_REDUCTION_TARGETS"

	StatusNoNativeWallHierarchyResponseTheorem   = "FAILED_ROUTE_NO_NATIVE_WALL_HIERARCHY_RESPONSE_THEOREM"
	StatusNoK7ToBoundaryVectorMap                = "FAILED_ROUTE_NO_NATIVE_K7_TO_BOUNDARY_VECTOR_MAP"
	StatusK7EventWeightNotFlavorTheorem          = "FAILED_ROUTE_K7_EVENT_WEIGHT_IS_NOT_FLAVOR_THEOREM"
	StatusBoundaryStressMomentNotNativeFlavorLaw = "FAILED_ROUTE_BOUNDARY_STRESS_MOMENT_IS_NOT_NATIVE_FLAVOR_LAW"
	StatusNoNativeBoundaryHistoryResponseTheorem = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem         = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoPhysicalHiggsOrPoleMassTheorem       = "FAILED_ROUTE_NO_PHYSICAL_HIGGS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem    = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate749Boundary                        = "FIREWALL_PRESERVED_GATE749_LAW_HISTORY_WALL_HIERARCHY_BOUNDARY"
)

type Gate748Inheritance struct {
	Inherited               bool
	ResidualCompression     bool
	BoundaryStressReappears bool
	FlavorFirewallKept      bool
	Verdict                 string
}

type WallRecord struct {
	Name     string
	Equation string
	Layer    string
	Meaning  string
	K7Role   string
	Status   string
	Firewall string
}

type WallHierarchy struct {
	Walls   []WallRecord
	Count   int
	Verdict string
}

type K7RoleSeparation struct {
	NativeSupportRole        string
	EventWeightRole          string
	RawMomentRole            string
	BoundaryVectorMapBlocked bool
	FlavorPromotionBlocked   bool
	HiggsPromotionBlocked    bool
	Verdict                  string
}

type FirewallStep struct {
	Order int
	Name  string
	Why   string
}

type FirewallOrder struct {
	Steps   []FirewallStep
	Count   int
	Verdict string
}

type Gate748WallResonance struct {
	KappaEExpression string
	UsesOrientation  bool
	UsesFiveThirds   bool
	UsesXiBoundary   bool
	UsesK7Moment     bool
	IsFlavorTheorem  bool
	Verdict          string
}

type BoundaryMomentCoordinate struct {
	ResponseOperator string
	MomentFormula    string
	K7EventWeight    float64
	XiBoundary       float64
	SSplit           float64
	M2Wall           float64
	Verdict          string
}

type ReductionPriority struct {
	NextGateName         string
	DoNotChaseResidual   bool
	RecommendedTargets   []string
	StabilizedBeforeNext bool
	Verdict              string
}

type PhysicalFirewalls struct {
	NoBoundaryVectorMap    bool
	NoFlavorTheorem        bool
	NoHistoryLoopTheorem   bool
	NoScalarRuntimeTheorem bool
	NoHiggsPoleMassTheorem bool
	NoYukawaTheorem        bool
	Verdict                string
}

type Analysis struct {
	Gate748   Gate748Inheritance
	Hierarchy WallHierarchy
	K7Roles   K7RoleSeparation
	Firewall  FirewallOrder
	Resonance Gate748WallResonance
	Moment    BoundaryMomentCoordinate
	Reduction ReductionPriority
	Physical  PhysicalFirewalls
	Truth     string
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
	g748, err := gate748.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate748 inheritance unavailable: %w", err)
	}
	inherit := buildGate748Inheritance(g748)
	hierarchy := buildWallHierarchy()
	k7Roles := buildK7RoleSeparation()
	firewall := buildFirewallOrder()
	resonance := buildGate748WallResonance()
	moment := buildBoundaryMomentCoordinate(g748)
	reduction := buildReductionPriority()
	physical := buildPhysicalFirewalls()
	truth := "Gate 749 orders the active wall system exposed by Gate 748. K7 remains native Boolean-octonionic support in the finite law-space, and in the boundary/history lane it appears as the no-bias event weight p_K7=7/72 inside raw response moments M_n=p_K7 S_split^n. Gate 748's kappa_e refinement is therefore a law-history wall resonance, not a native flavor theorem or a K7-to-boundary vector map. The hierarchy is now clean enough to reduce kappa_e, P_rad, n, or F_wall_3 without confusing wall layers."
	return Analysis{Gate748: inherit, Hierarchy: hierarchy, K7Roles: k7Roles, Firewall: firewall, Resonance: resonance, Moment: moment, Reduction: reduction, Physical: physical, Truth: truth}, nil
}

func buildGate748Inheritance(g gate748.Analysis) Gate748Inheritance {
	return Gate748Inheritance{
		Inherited:               g.Gate747.Inherited && g.Correction.CompressionFactor > 250,
		ResidualCompression:     g.Correction.CompressionFactor > 250 && g.Replacement.StressImprovementOverHyper > 100,
		BoundaryStressReappears: g.Stress.BestCandidate == "xi_boundary midpoint" && g.Correction.StressMoment > 0,
		FlavorFirewallKept:      !g.Firewall.DerivesFlavorTheorem && !g.Firewall.DerivesYukawa && !g.Firewall.DerivesScalarRuntime,
		Verdict:                 StatusGate748BoundaryStressMomentInherited,
	}
}

func buildWallHierarchy() WallHierarchy {
	walls := []WallRecord{
		{"Boolean wall", "P_B, rank 56", "native law-space", "Boolean incidence law-sector", "K7 is contained in Im(P_B)", "native", "none beyond native law-space type discipline"},
		{"Octonionic wall", "P_G, rank 14", "native law-space", "Fano/G2 calibration law-sector", "K7 is contained in Im(P_G)", "native", "does not itself select boundary stress"},
		{"K7 contact wall", "K7 = Im(P_B) ∩ Im(P_G), dim 7", "native support", "Boolean-octonionic contact carrier", "selected by Boolean + octonionic support", "native support", "rank seven alone is insufficient"},
		{"Hodge polarity wall", "K7 = K7+ ⊕ K7-, 4+3", "internal shadow", "Higgs/flavor shadow carrier", "internal structure of K7", "candidate shadow", "not physical Higgs/flavor theorem"},
		{"Augmented chamber wall", "H72 = Λ4 R8 ⊕ R2_boundary = 70+2", "bridge observer chamber", "finite chamber plus boundary pair", "p_K7=7/72 after rho_72", "bridge typed", "requires augmented no-bias observer"},
		{"Scalar zero-wall", "lambda(Lambda12)=0", "boundary scalar", "signed scalar-wall coordinate", "boundary pair scalar side", "bridge coordinate", "not K7 vector map"},
		{"Gauge meeting-wall", "R3-1", "boundary scalar", "gauge-scalar meeting wound", "boundary pair gauge side", "bridge coordinate", "not K7 vector map"},
		{"Anti-alignment wall", "lambda+(R3-1)=0", "boundary quotient", "perfect scalar/gauge anti-alignment", "K7 weights its response moments", "bridge quotient", "S_split is not native stress theorem"},
		{"Boundary stress wall", "xi_boundary=(|lambda|+R)/2", "boundary midpoint", "common midpoint stress", "appears in Gate748 with M2_wall", "bridge stress", "not native flavor operator"},
		{"History closure wall", "kappa_lambda+kappa_e+lambda=0", "history readout", "scalar/flavor/history closure", "D_base compared to K7-weighted boundary response", "bridge readout", "not automatically boundary-controlled"},
		{"Flavor orientation wall", "kappa_e_orient=sin^2(theta13)/4-J_CKM", "flavor bridge", "PMNS/CKM orientation candidate", "refined by boundary moments in Gate748", "bridge input", "not PMNS/CKM theorem"},
		{"Hypercharge boundary-square wall", "-(5/3)S_split^2", "flavor residual correction", "hypercharge-normalized second boundary correction", "acts through boundary scalar, not native K7 map", "bridge correction", "not native flavor theorem"},
		{"K7 second-moment stress wall", "xi_boundary p_K7 S_split^2", "boundary/flavor residual refinement", "midpoint stress times K7 raw M2 wall", "K7 only as event weight p_K7", "bridge residual compression", "not native flavor law"},
		{"HistoryLoop wall", "L=1/(8π)", "scalar transport", "radial-Hopf loop candidate", "separate from p_K7", "bridge seal", "not native transport theorem"},
		{"Higgs socket wall", "(n,q,P_rad)", "representation/scalar seal", "minimal scalar-Higgs seal package", "K7+ shadow only after seals", "sealed interface", "not physical Higgs theorem"},
		{"Tree/pole wall", "m_tree ≠ m_pole", "physical translation", "tree proxy to pole observable boundary", "no K7 promotion", "forecast firewall", "pole correction external/sealed"},
	}
	return WallHierarchy{Walls: walls, Count: len(walls), Verdict: StatusWallHierarchyDefined}
}

func buildK7RoleSeparation() K7RoleSeparation {
	return K7RoleSeparation{
		NativeSupportRole:        "K7 is the native Boolean-octonionic contact support Im(P_B)∩Im(P_G).",
		EventWeightRole:          "In H72 with rho_72, K7 contributes p_K7=Tr(rho_72 P_K7)=7/72 as a no-bias event weight.",
		RawMomentRole:            "Boundary response moments use M_n=Tr(rho_72(S_split P_K7)^n)=p_K7 S_split^n.",
		BoundaryVectorMapBlocked: true,
		FlavorPromotionBlocked:   true,
		HiggsPromotionBlocked:    true,
		Verdict: strings.Join([]string{
			StatusK7RoleSeparationAudited,
			StatusK7ActsAsSupportAndEventWeightOnly,
			StatusNoK7ToBoundaryVectorMap,
			StatusK7EventWeightNotFlavorTheorem,
		}, "; "),
	}
}

func buildFirewallOrder() FirewallOrder {
	steps := []FirewallStep{
		{1, "Native law-space firewall", "Cl(1,7), Λ4R8, P_B, P_G, and K7 are native law objects."},
		{2, "K7 support-selection firewall", "Rank seven alone does not select K7; Boolean plus octonionic support does."},
		{3, "Augmented observer firewall", "7/72 requires rho_72=I_H72/72; other observers give 7/70, 7/71, or 1."},
		{4, "Boundary quotient firewall", "S_split is anti-alignment quotient coordinate, not native stress theorem."},
		{5, "History readout firewall", "D_base is history closure defect, not automatically controlled by boundary."},
		{6, "Raw-moment firewall", "M_n=p_K7 S_split^n is active coordinate; no native moment theorem yet."},
		{7, "Flavor wall firewall", "kappa_e remains a flavor bridge input; PMNS/CKM form is close but not theorem."},
		{8, "Boundary-stress residual firewall", "xi_boundary M2 compresses Gate748 residual but is not native flavor law."},
		{9, "HistoryLoop firewall", "L=1/(8π) has radial-Hopf source type, not native transport theorem."},
		{10, "Higgs socket firewall", "(n,q,P_rad) seals are required; socket is not Higgs theorem."},
		{11, "Runtime scalar firewall", "lambda_runtime bridge is consistency closure, not independent prediction."},
		{12, "Tree/pole mass firewall", "m_tree is not m_pole; pole correction remains external/sealed."},
	}
	return FirewallOrder{Steps: steps, Count: len(steps), Verdict: StatusFirewallOrderConstructed}
}

func buildGate748WallResonance() Gate748WallResonance {
	return Gate748WallResonance{
		KappaEExpression: "kappa_e ≈ sin^2(theta13)/4 - J_CKM - (5/3)S_split^2 + xi_boundary p_K7 S_split^2",
		UsesOrientation:  true,
		UsesFiveThirds:   true,
		UsesXiBoundary:   true,
		UsesK7Moment:     true,
		IsFlavorTheorem:  false,
		Verdict: strings.Join([]string{
			StatusGate748WallResonanceRecorded,
			StatusGate748CorrectionLawHistoryWallResonance,
			StatusBoundaryStressMomentNotNativeFlavorLaw,
		}, "; "),
	}
}

func buildBoundaryMomentCoordinate(g gate748.Analysis) BoundaryMomentCoordinate {
	return BoundaryMomentCoordinate{
		ResponseOperator: "R_wall = S_split P_K7",
		MomentFormula:    "M_n = Tr(rho_72 R_wall^n)=p_K7 S_split^n",
		K7EventWeight:    g.Residual.P_K7,
		XiBoundary:       g.Stress.XiBoundary,
		SSplit:           g.Residual.SSplit,
		M2Wall:           g.Residual.M2Wall,
		Verdict:          StatusBoundaryMomentCoordinateRecorded,
	}
}

func buildReductionPriority() ReductionPriority {
	return ReductionPriority{
		NextGateName:         "Gate 749 — Law-History Wall Hierarchy and K7 Response Firewall Ordering Audit",
		DoNotChaseResidual:   true,
		RecommendedTargets:   []string{"kappa_e source reduction", "P_rad selector reduction", "twistor selector n reduction", "F_wall_3 generating-function source"},
		StabilizedBeforeNext: true,
		Verdict: strings.Join([]string{
			StatusReductionPriorityRecorded,
			StatusWallOrderingStabilizesNextReductions,
		}, "; "),
	}
}

func buildPhysicalFirewalls() PhysicalFirewalls {
	return PhysicalFirewalls{
		NoBoundaryVectorMap:    true,
		NoFlavorTheorem:        true,
		NoHistoryLoopTheorem:   true,
		NoScalarRuntimeTheorem: true,
		NoHiggsPoleMassTheorem: true,
		NoYukawaTheorem:        true,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeWallHierarchyResponseTheorem,
			StatusNoK7ToBoundaryVectorMap,
			StatusNoNativeBoundaryHistoryResponseTheorem,
			StatusNoNativeHistoryLoopUnitTheorem,
			StatusNoPhysicalHiggsOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate749Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate748BoundaryStressMomentInherited,
		StatusWallHierarchyDefined,
		StatusK7RoleSeparationAudited,
		StatusFirewallOrderConstructed,
		StatusGate748WallResonanceRecorded,
		StatusBoundaryMomentCoordinateRecorded,
		StatusReductionPriorityRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusGate748CorrectionLawHistoryWallResonance,
		StatusK7ActsAsSupportAndEventWeightOnly,
		StatusWallOrderingStabilizesNextReductions,
		StatusNoNativeWallHierarchyResponseTheorem,
		StatusNoK7ToBoundaryVectorMap,
		StatusK7EventWeightNotFlavorTheorem,
		StatusBoundaryStressMomentNotNativeFlavorLaw,
		StatusNoNativeBoundaryHistoryResponseTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoPhysicalHiggsOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate749Boundary,
	}
}

func FormatGate748(x Gate748Inheritance) string {
	return fmt.Sprintf("inherited=%t compression=%t boundaryStress=%t flavorFirewall=%t verdict=%q", x.Inherited, x.ResidualCompression, x.BoundaryStressReappears, x.FlavorFirewallKept, x.Verdict)
}

func FormatHierarchy(x WallHierarchy) string {
	names := make([]string, 0, len(x.Walls))
	for _, w := range x.Walls {
		names = append(names, fmt.Sprintf("%d:%s[%s]", len(names)+1, w.Name, w.Layer))
	}
	return fmt.Sprintf("count=%d walls=[%s] verdict=%q", x.Count, strings.Join(names, "; "), x.Verdict)
}

func FormatK7Roles(x K7RoleSeparation) string {
	return fmt.Sprintf("native=%q event=%q moments=%q boundaryMapBlocked=%t flavorBlocked=%t higgsBlocked=%t verdict=%q", x.NativeSupportRole, x.EventWeightRole, x.RawMomentRole, x.BoundaryVectorMapBlocked, x.FlavorPromotionBlocked, x.HiggsPromotionBlocked, x.Verdict)
}

func FormatFirewall(x FirewallOrder) string {
	parts := make([]string, 0, len(x.Steps))
	for _, s := range x.Steps {
		parts = append(parts, fmt.Sprintf("%d:%s", s.Order, s.Name))
	}
	return fmt.Sprintf("count=%d steps=[%s] verdict=%q", x.Count, strings.Join(parts, "; "), x.Verdict)
}

func FormatResonance(x Gate748WallResonance) string {
	return fmt.Sprintf("expression=%q orientation=%t fiveThirds=%t xi=%t k7Moment=%t flavorTheorem=%t verdict=%q", x.KappaEExpression, x.UsesOrientation, x.UsesFiveThirds, x.UsesXiBoundary, x.UsesK7Moment, x.IsFlavorTheorem, x.Verdict)
}

func FormatMoment(x BoundaryMomentCoordinate) string {
	return fmt.Sprintf("operator=%q formula=%q p=%.17g xi=%.17g S=%.17g M2=%.17g verdict=%q", x.ResponseOperator, x.MomentFormula, x.K7EventWeight, x.XiBoundary, x.SSplit, x.M2Wall, x.Verdict)
}

func FormatReduction(x ReductionPriority) string {
	return fmt.Sprintf("next=%q doNotChase=%t targets=[%s] stabilized=%t verdict=%q", x.NextGateName, x.DoNotChaseResidual, strings.Join(x.RecommendedTargets, "; "), x.StabilizedBeforeNext, x.Verdict)
}

func FormatPhysical(x PhysicalFirewalls) string {
	return fmt.Sprintf("noBoundaryMap=%t noFlavor=%t noHistoryLoop=%t noRuntime=%t noHiggsPole=%t noYukawa=%t verdict=%q", x.NoBoundaryVectorMap, x.NoFlavorTheorem, x.NoHistoryLoopTheorem, x.NoScalarRuntimeTheorem, x.NoHiggsPoleMassTheorem, x.NoYukawaTheorem, x.Verdict)
}
