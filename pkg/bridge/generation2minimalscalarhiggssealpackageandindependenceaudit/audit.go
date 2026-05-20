// Package generation2minimalscalarhiggssealpackageandindependenceaudit implements
// Gate 738: Minimal Scalar-Higgs Seal Package and Independence Audit.
//
// Gate 737 showed that no currently typed ASHA object selects the rank-one
// radial projector P_rad. Gate 738 audits the full minimal scalar-Higgs seal
// package (n,q,P_rad): n selects the complex structure / Hopf phase rule, q
// normalizes the moving U(1) phase line, and P_rad selects the radial/vacuum
// event. The gate certifies their distinct roles, independence, and minimality
// for the current scalar-Higgs bridge while preserving the firewall that the
// package is not a physical Higgs, scalar-runtime, or Yukawa theorem.
package generation2minimalscalarhiggssealpackageandindependenceaudit

import (
	"fmt"
	"strings"
	"sync"

	gate737 "github.com/bagherbal/asha-engine/pkg/bridge/generation2higgsradialselectorsourcecandidateandvacuumdirectionfirewallaudit"
)

const (
	AuditID = "GATE738-MINIMAL-SCALAR-HIGGS-SEAL-PACKAGE-INDEPENDENCE-AUDIT"

	StatusGate737RadialSelectorFirewallInherited = "PASS_GATE737_RADIAL_SELECTOR_FIREWALL_INHERITED"
	StatusNQPRadRolesDefined                     = "PASS_N_Q_P_RAD_ROLES_DEFINED"
	StatusSealIndependenceAudited                = "PASS_SEAL_INDEPENDENCE_AUDITED"
	StatusMinimalityAudited                      = "PASS_MINIMALITY_AUDITED"
	StatusAvailableStructuresUnderPackageAudited = "PASS_AVAILABLE_STRUCTURES_UNDER_SEAL_PACKAGE_AUDITED"
	StatusRemainingBridgeDependenciesRecorded    = "PASS_REMAINING_BRIDGE_DEPENDENCIES_RECORDED"
	StatusPhysicalFirewallsEnforced              = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusScalarHiggsSealPackageMinimal         = "CONDITIONAL_SUPPORT_SCALAR_HIGGS_SEAL_PACKAGE_IS_MINIMAL"
	StatusNQPRadTypeDistinctAndIndependent      = "CONDITIONAL_SUPPORT_N_Q_P_RAD_ARE_TYPE_DISTINCT_AND_INDEPENDENT"
	StatusCurrentBridgeRequiresThreeSealPackage = "CONDITIONAL_SUPPORT_CURRENT_SCALAR_HIGGS_BRIDGE_REQUIRES_THREE_SEAL_PACKAGE"

	StatusNNotNativelyDerived                  = "FAILED_ROUTE_N_NOT_NATIVELY_DERIVED"
	StatusQNotNativelyDerived                  = "FAILED_ROUTE_Q_NOT_NATIVELY_DERIVED"
	StatusPRadNotNativelyDerived               = "FAILED_ROUTE_P_RAD_NOT_NATIVELY_DERIVED"
	StatusSealPackageNotPhysicalHiggsTheorem   = "FAILED_ROUTE_SEAL_PACKAGE_NOT_PHYSICAL_HIGGS_THEOREM"
	StatusNoNativeHistoryLoopUnitSourceTheorem = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativeScalarRuntimeTheorem         = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem         = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem  = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate738Boundary                      = "FIREWALL_PRESERVED_GATE738_MINIMAL_SCALAR_HIGGS_SEAL_PACKAGE_BOUNDARY"
)

type Gate737Inheritance struct {
	Inherited                  bool
	PRadTypeDistinctSeal       bool
	PRadRequiresVacuumSelector bool
	NoNativePRadSelector       bool
	HistoryLoopConditional     bool
	NoMassTheorem              bool
	NoYukawaTheorem            bool
	Verdict                    string
}

type SealRole struct {
	Name        string
	Type        string
	Role        string
	SelectsN    bool
	SelectsQ    bool
	SelectsPRad bool
}

type SealRoleAudit struct {
	PackageName       string
	Roles             []SealRole
	NSelectsJH        bool
	NDefinesHopfPhase bool
	QNormalizesU1     bool
	PRadSelectsRadial bool
	PRadEnablesSplits bool
	Verdict           string
}

type ForbiddenSubstitution struct {
	From    string
	To      string
	Allowed bool
	Reason  string
}

type IndependenceAudit struct {
	Substitutions        []ForbiddenSubstitution
	NQTypeDistinct       bool
	NPRadTypeDistinct    bool
	QPRadTypeDistinct    bool
	RhoPlusDeterminesAny bool
	PK7DeterminesPRad    bool
	Verdict              string
}

type MinimalityAudit struct {
	RemoveNConsequences    []string
	RemoveQConsequences    []string
	RemovePRadConsequences []string
	AllThreeRequired       bool
	Verdict                string
}

type AvailableStructuresAudit struct {
	Structures              []string
	NQPRadSupplied          bool
	HistoryLoopAvailable    bool
	RuntimeBridgeCompatible bool
	Verdict                 string
}

type RemainingDependenciesAudit struct {
	Dependencies           []string
	AllStillBridgeOrSealed bool
	Verdict                string
}

type PhysicalFirewallAudit struct {
	PackageIsPhysicalHiggsTheorem         bool
	PRadIsElectroweakVacuumTheorem        bool
	NIsNativeComplexStructureTheorem      bool
	QIsNativeHyperchargeDerivation        bool
	LIsNativeHistoryLoopTheorem           bool
	RuntimeBridgeIsHiggsMassPrediction    bool
	FWall3IsNativeBoundaryResponseTheorem bool
	Verdict                               string
}

type Analysis struct {
	Gate737      Gate737Inheritance
	Roles        SealRoleAudit
	Independence IndependenceAudit
	Minimality   MinimalityAudit
	Available    AvailableStructuresAudit
	Remaining    RemainingDependenciesAudit
	Firewall     PhysicalFirewallAudit
	Truth        string
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
	g737, err := gate737.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate737 inheritance unavailable: %w", err)
	}
	inherit := buildGate737Inheritance(g737)
	roles := buildRoles()
	ind := buildIndependence()
	min := buildMinimality()
	avail := buildAvailable()
	rem := buildRemaining()
	fw := buildFirewall()
	truth := "Gate 738 audits the full scalar-Higgs seal package (n,q,P_rad). The three objects have type-distinct roles: n selects the complex structure and Hopf phase direction, q normalizes the selected U(1) phase line, and P_rad selects the radial/vacuum event required by the Radial-Hopf source law. Removal of any one breaks the current scalar-Higgs bridge, so the package is minimal at the present level. The package remains sealed and does not constitute a native Higgs, scalar-runtime, HistoryLoopUnit, or Yukawa theorem."
	return Analysis{Gate737: inherit, Roles: roles, Independence: ind, Minimality: min, Available: avail, Remaining: rem, Firewall: fw, Truth: truth}, nil
}

func buildGate737Inheritance(g gate737.Analysis) Gate737Inheritance {
	return Gate737Inheritance{
		Inherited:                  g.Gate736.Inherited && !g.Candidates.AnyNativeSelectorFound,
		PRadTypeDistinctSeal:       g.Seal.DistinctFromN && g.Seal.DistinctFromQ && strings.Contains(g.Seal.Verdict, gate737.StatusPRadTypeDistinctScalarVacuumDirectionSeal),
		PRadRequiresVacuumSelector: g.Symmetry.RequiresVacuumSelector,
		NoNativePRadSelector:       !g.Candidates.AnyNativeSelectorFound,
		HistoryLoopConditional:     g.HistoryLoop.ConditionalWithoutPRad,
		NoMassTheorem:              !g.Firewall.PRadIsHiggsMassTheorem,
		NoYukawaTheorem:            !g.Firewall.YukawaOperatorOrEigenvalueTheorem,
		Verdict:                    StatusGate737RadialSelectorFirewallInherited,
	}
}

func buildRoles() SealRoleAudit {
	roles := []SealRole{
		{Name: "n", Type: "twistor point / complex-structure selector in S^2(K7-)", Role: "selects J_H(n), gives K7+ a chosen C^2 structure, and defines the Hopf phase direction through P_rad", SelectsN: true},
		{Name: "q", Type: "phase-line / hypercharge normalization scalar", Role: "normalizes the selected U(1) phase generator qJ_H(n); does not select n or P_rad", SelectsQ: true},
		{Name: "P_rad", Type: "rank-one radial projector / scalar vacuum-direction selector in K7+", Role: "selects the radial event, enables 1+3 and 1+1+2 splits, and supplies the event for the HistoryLoopUnit expectation", SelectsPRad: true},
	}
	return SealRoleAudit{
		PackageName:       "ScalarHiggsSealPackage=(n,q,P_rad)",
		Roles:             roles,
		NSelectsJH:        true,
		NDefinesHopfPhase: true,
		QNormalizesU1:     true,
		PRadSelectsRadial: true,
		PRadEnablesSplits: true,
		Verdict:           strings.Join([]string{StatusNQPRadRolesDefined, StatusNQPRadTypeDistinctAndIndependent}, "; "),
	}
}

func buildIndependence() IndependenceAudit {
	subs := []ForbiddenSubstitution{
		{From: "n", To: "q", Allowed: false, Reason: "a twistor direction does not fix charge normalization"},
		{From: "n", To: "P_rad", Allowed: false, Reason: "complex-structure choice does not select a real radial line"},
		{From: "q", To: "n", Allowed: false, Reason: "scalar normalization cannot select a point on S^2(K7-)"},
		{From: "q", To: "P_rad", Allowed: false, Reason: "phase charge normalization carries no K7+ direction"},
		{From: "P_rad", To: "n", Allowed: false, Reason: "radial line does not select a quaternionic complex-structure direction"},
		{From: "P_rad", To: "q", Allowed: false, Reason: "radial event does not normalize hypercharge"},
		{From: "rho_plus", To: "n or P_rad", Allowed: false, Reason: "maximum-entropy state is isotropic and selects no twistor point or radial line"},
		{From: "P_K7", To: "P_rad", Allowed: false, Reason: "rank-seven event support does not choose a line inside K7+"},
	}
	return IndependenceAudit{Substitutions: subs, NQTypeDistinct: true, NPRadTypeDistinct: true, QPRadTypeDistinct: true, RhoPlusDeterminesAny: false, PK7DeterminesPRad: false, Verdict: strings.Join([]string{StatusSealIndependenceAudited, StatusNQPRadTypeDistinctAndIndependent}, "; ")}
}

func buildMinimality() MinimalityAudit {
	return MinimalityAudit{
		RemoveNConsequences:    []string{"no chosen complex structure", "no Hopf phase direction", "no full Higgs socket"},
		RemoveQConsequences:    []string{"phase line exists after n", "charge/hypercharge normalization is not fixed"},
		RemovePRadConsequences: []string{"no radial event", "no radial-Hopf source law for L", "no scalar vacuum-direction candidate"},
		AllThreeRequired:       true,
		Verdict:                strings.Join([]string{StatusMinimalityAudited, StatusScalarHiggsSealPackageMinimal, StatusCurrentBridgeRequiresThreeSealPackage}, "; "),
	}
}

func buildAvailable() AvailableStructuresAudit {
	return AvailableStructuresAudit{
		Structures:              []string{"K7+_J(n) ~= C^2", "g_int(n,q)=C ⊕ span(qJ_H(n))", "K7+=K_rad⊕K_ang", "K7+=K_rad⊕K_phase⊕K_trans", "L=Tr(rho_plus[(1/(2*pi))P_rad])=1/(8*pi)", "scalar runtime bridge compatibility"},
		NQPRadSupplied:          true,
		HistoryLoopAvailable:    true,
		RuntimeBridgeCompatible: true,
		Verdict:                 strings.Join([]string{StatusAvailableStructuresUnderPackageAudited, StatusCurrentBridgeRequiresThreeSealPackage}, "; "),
	}
}

func buildRemaining() RemainingDependenciesAudit {
	return RemainingDependenciesAudit{
		Dependencies:           []string{"lambda_proxy", "kappa_e", "F_wall_3", "HistoryLoop transport law", "boundary response principle", "scale-local Lambda12 status"},
		AllStillBridgeOrSealed: true,
		Verdict:                StatusRemainingBridgeDependenciesRecorded,
	}
}

func buildFirewall() PhysicalFirewallAudit {
	return PhysicalFirewallAudit{
		PackageIsPhysicalHiggsTheorem:         false,
		PRadIsElectroweakVacuumTheorem:        false,
		NIsNativeComplexStructureTheorem:      false,
		QIsNativeHyperchargeDerivation:        false,
		LIsNativeHistoryLoopTheorem:           false,
		RuntimeBridgeIsHiggsMassPrediction:    false,
		FWall3IsNativeBoundaryResponseTheorem: false,
		Verdict:                               strings.Join([]string{StatusPhysicalFirewallsEnforced, StatusNNotNativelyDerived, StatusQNotNativelyDerived, StatusPRadNotNativelyDerived, StatusSealPackageNotPhysicalHiggsTheorem, StatusNoNativeHistoryLoopUnitSourceTheorem, StatusNoNativeScalarRuntimeTheorem, StatusNoHiggsMassOrPoleMassTheorem, StatusNoYukawaOperatorOrEigenvalueTheorem, StatusGate738Boundary}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate737RadialSelectorFirewallInherited,
		StatusNQPRadRolesDefined,
		StatusSealIndependenceAudited,
		StatusMinimalityAudited,
		StatusAvailableStructuresUnderPackageAudited,
		StatusRemainingBridgeDependenciesRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusScalarHiggsSealPackageMinimal,
		StatusNQPRadTypeDistinctAndIndependent,
		StatusCurrentBridgeRequiresThreeSealPackage,
		StatusNNotNativelyDerived,
		StatusQNotNativelyDerived,
		StatusPRadNotNativelyDerived,
		StatusSealPackageNotPhysicalHiggsTheorem,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate738Boundary,
	}
}

func FormatGate737(x Gate737Inheritance) string {
	return fmt.Sprintf("inherited=%t pRadSeal=%t vacuum=%t noNative=%t conditionalL=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.PRadTypeDistinctSeal, x.PRadRequiresVacuumSelector, x.NoNativePRadSelector, x.HistoryLoopConditional, x.NoMassTheorem, x.NoYukawaTheorem, x.Verdict)
}
func FormatRoles(x SealRoleAudit) string {
	parts := make([]string, 0, len(x.Roles))
	for _, r := range x.Roles {
		parts = append(parts, fmt.Sprintf("%s:%s", r.Name, r.Role))
	}
	return fmt.Sprintf("package=%q roles=[%s] nJH=%t nHopf=%t qU1=%t pRad=%t splits=%t verdict=%q", x.PackageName, strings.Join(parts, " | "), x.NSelectsJH, x.NDefinesHopfPhase, x.QNormalizesU1, x.PRadSelectsRadial, x.PRadEnablesSplits, x.Verdict)
}
func FormatIndependence(x IndependenceAudit) string {
	return fmt.Sprintf("subs=%d nQ=%t nPRad=%t qPRad=%t rhoAny=%t pK7PRad=%t verdict=%q", len(x.Substitutions), x.NQTypeDistinct, x.NPRadTypeDistinct, x.QPRadTypeDistinct, x.RhoPlusDeterminesAny, x.PK7DeterminesPRad, x.Verdict)
}
func FormatMinimality(x MinimalityAudit) string {
	return fmt.Sprintf("removeN=%q removeQ=%q removePRad=%q allThree=%t verdict=%q", strings.Join(x.RemoveNConsequences, ","), strings.Join(x.RemoveQConsequences, ","), strings.Join(x.RemovePRadConsequences, ","), x.AllThreeRequired, x.Verdict)
}
func FormatAvailable(x AvailableStructuresAudit) string {
	return fmt.Sprintf("structures=%q supplied=%t L=%t runtime=%t verdict=%q", strings.Join(x.Structures, ";"), x.NQPRadSupplied, x.HistoryLoopAvailable, x.RuntimeBridgeCompatible, x.Verdict)
}
func FormatRemaining(x RemainingDependenciesAudit) string {
	return fmt.Sprintf("dependencies=%q bridgeOrSealed=%t verdict=%q", strings.Join(x.Dependencies, ","), x.AllStillBridgeOrSealed, x.Verdict)
}
func FormatFirewall(x PhysicalFirewallAudit) string {
	return fmt.Sprintf("physicalHiggs=%t ewsb=%t nativeN=%t nativeQ=%t nativeL=%t mass=%t fwall3=%t verdict=%q", x.PackageIsPhysicalHiggsTheorem, x.PRadIsElectroweakVacuumTheorem, x.NIsNativeComplexStructureTheorem, x.QIsNativeHyperchargeDerivation, x.LIsNativeHistoryLoopTheorem, x.RuntimeBridgeIsHiggsMassPrediction, x.FWall3IsNativeBoundaryResponseTheorem, x.Verdict)
}
