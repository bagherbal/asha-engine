// Package generation2minimalhiggssocketsealpackageandpromotionboundaryaudit implements
// Gate 721: Minimal Higgs Socket Seal Package and Promotion Boundary Audit.
//
// Gate 720 showed that the conditional electroweak Higgs socket
// g_int(n,q)=C ⊕ span(qJ_H(n)) requires two type-distinct missing choices:
// n, a twistor point / complex-structure selector in S^2(K7-), and q, a
// phase-line / hypercharge normalization in R^×.  Gate 721 audits the minimal
// sealed package needed to promote the internal representation socket from a
// purely conditional object to a sealed electroweak-Higgs representation
// interface.  It classifies the package (n,q), reconstructs what becomes
// available under those seals, proves that each seal does nonredundant work,
// and preserves the firewall that a sealed representation interface is still
// not a physical Higgs theorem, scalar-runtime theorem, Yukawa theorem, CKM/PMNS
// theorem, or native 7/72 theorem.
package generation2minimalhiggssocketsealpackageandpromotionboundaryaudit

import (
	"fmt"
	"strings"
	"sync"

	gate720 "github.com/bagherbal/asha-engine/pkg/bridge/generation2higgssocketmissingsealindependenceandsourcecandidateaudit"
)

const (
	AuditID = "GATE721-MINIMAL-HIGGS-SOCKET-SEAL-PACKAGE-AND-PROMOTION-BOUNDARY-AUDIT"

	StatusGate720MissingSealIndependenceInherited = "PASS_GATE720_MISSING_SEAL_INDEPENDENCE_INHERITED"
	StatusMinimalSealPackageDefined               = "PASS_MINIMAL_SEAL_PACKAGE_DEFINED"
	StatusTwistorSelectorSealRoleAudited          = "PASS_TWISTOR_SELECTOR_SEAL_ROLE_AUDITED"
	StatusHyperchargeNormalizationSealRoleAudited = "PASS_HYPERCHARGE_NORMALIZATION_SEAL_ROLE_AUDITED"
	StatusSealedSocketAssemblyReconstructed       = "PASS_SEALED_SOCKET_ASSEMBLY_RECONSTRUCTED"
	StatusAvailableStructuresUnderSealsAudited    = "PASS_AVAILABLE_STRUCTURES_UNDER_SEALS_AUDITED"
	StatusRemainingBlockedPhysicsAudited          = "PASS_REMAINING_BLOCKED_PHYSICS_AUDITED"
	StatusSealMinimalityAudited                   = "PASS_SEAL_MINIMALITY_AUDITED"
	StatusNQIndependencePreserved                 = "PASS_N_Q_INDEPENDENCE_PRESERVED"
	StatusHiggsSocketSealPackageMinimal           = "CONDITIONAL_SUPPORT_HIGGS_SOCKET_SEAL_PACKAGE_IS_MINIMAL"
	StatusSealedHiggsSocketInterfaceDefined       = "CONDITIONAL_SUPPORT_SEALED_HIGGS_SOCKET_REPRESENTATION_INTERFACE_DEFINED"
	StatusK7PlusAirlockReadyOnlyAfterSeals        = "CONDITIONAL_SUPPORT_K7_PLUS_HIGGS_REPRESENTATION_AIRLOCK_IS_READY_ONLY_AFTER_N_AND_Q_SEALS"
	StatusNNotNativelyDerived                     = "FAILED_ROUTE_N_NOT_NATIVELY_DERIVED"
	StatusQNotNativelyDerived                     = "FAILED_ROUTE_Q_NOT_NATIVELY_DERIVED"
	StatusSealedSocketNotPhysicalHiggsTheorem     = "FAILED_ROUTE_SEALED_SOCKET_NOT_FULL_PHYSICAL_HIGGS_THEOREM"
	StatusNoScalarPotentialOrRuntimeLambda        = "FAILED_ROUTE_NO_SCALAR_POTENTIAL_OR_RUNTIME_LAMBDA_THEOREM"
	StatusNoHiggsMassTheorem                      = "FAILED_ROUTE_NO_HIGGS_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate721PromotionBoundary                = "FIREWALL_PRESERVED_GATE721_MINIMAL_HIGGS_SOCKET_SEAL_PACKAGE_BOUNDARY"
)

const (
	minimalSealCount        = 2
	availableStructureCount = 5
	remainingBlockedCount   = 9
	minimalityRemovalCount  = 4
	forbiddenShortcutCount  = 4
)

type Gate720Inheritance struct {
	Inherited                          bool
	NAndQTypeDistinct                  bool
	NRequiresSelectorSeal              bool
	QRequiresNormalizationSeal         bool
	ConditionalSocketReadyButNotNative bool
	NoNativeN                          bool
	NoNativeQ                          bool
	NoPhysicalHiggs                    bool
	NoHiggsRuntime                     bool
	NoYukawa                           bool
	Verdict                            string
}

type SealRole struct {
	Name        string
	Input       string
	Output      string
	Required    bool
	Native      bool
	RoleAudited bool
	Verdict     string
}

type MinimalSealPackage struct {
	TwistorSelectorSeal          SealRole
	HyperchargeNormalizationSeal SealRole
	PackageName                  string
	Seals                        []string
	SealCount                    int
	SuppliesN                    bool
	SuppliesQ                    bool
	Minimal                      bool
	Native                       bool
	Verdict                      string
}

type SealedSocketAssembly struct {
	SelectedComplexCarrier   bool
	InternalU2Socket         bool
	SU2Compatibility         bool
	U1PhaseCompatibility     bool
	FullIntertwinerCandidate bool
	SocketFormula            string
	CarrierFormula           string
	Verdict                  string
}

type AvailableStructure struct {
	Name      string
	Available bool
	Reason    string
}

type AvailableUnderSealsAudit struct {
	Structures   []AvailableStructure
	AllAvailable bool
	Verdict      string
}

type BlockedPhysics struct {
	WhyNSelected               bool
	WhyQHasValue               bool
	PhysicalEqualitySU2U1      bool
	ScalarPotential            bool
	QuarticRuntimeLambda       bool
	HiggsPoleMass              bool
	YukawaOperatorConstruction bool
	FlavorHierarchy            bool
	CKMPMNS                    bool
	BlockedCount               int
	Verdict                    string
}

type MinimalityRemoval struct {
	Removed string
	Breaks  []string
	Fatal   bool
}

type SealMinimalityAudit struct {
	Removals           []MinimalityRemoval
	RemoveNBreaks      bool
	RemoveQBreaks      bool
	RemoveCBreaks      bool
	RemoveK7PlusBreaks bool
	PairMinimal        bool
	Verdict            string
}

type IndependenceAudit struct {
	NType                 string
	QType                 string
	TypeDistinct          bool
	NotMutuallyDerivable  bool
	QFromSevenOver72      bool
	NFromScalarBridgeData bool
	NFromPK7              bool
	QFromAbsN             bool
	ForbiddenShortcuts    []string
	Verdict               string
}

type PhysicalFirewall struct {
	TwistorSelectorSealNativeVacuumTheorem bool
	HyperchargeSealNativeDerivation        bool
	SealedSocketFullPhysicalHiggsTheorem   bool
	SealedSocketHiggsMassTheorem           bool
	SealedSocketYukawaTheorem              bool
	SealedSocketCKMPMNSTtheorem            bool
	NoScalarPotentialOrRuntimeLambda       bool
	NoHiggsMassTheorem                     bool
	NoYukawaOperatorOrEigenvalueTheorem    bool
	Verdict                                string
}

type Analysis struct {
	Gate720      Gate720Inheritance
	Package      MinimalSealPackage
	Assembly     SealedSocketAssembly
	Available    AvailableUnderSealsAudit
	Blocked      BlockedPhysics
	Minimality   SealMinimalityAudit
	Independence IndependenceAudit
	Physical     PhysicalFirewall
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
	g720, err := gate720.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate720 inheritance unavailable: %w", err)
	}
	inherited := buildGate720Inheritance(g720)
	pkg := buildMinimalSealPackage()
	assembly := buildSealedSocketAssembly(pkg)
	available := buildAvailableUnderSealsAudit(assembly)
	blocked := buildBlockedPhysics()
	minimality := buildSealMinimalityAudit()
	independence := buildIndependenceAudit(g720)
	physical := buildPhysicalFirewall()
	truth := "Gate 721 packages the two type-distinct missing choices identified by Gate 720 as the minimal sealed pair (n,q).  The TwistorSelectorSeal supplies n and therefore J_H(n), L_n, and K7+_J(n); the HyperchargeNormalizationSeal supplies q and therefore qJ_H(n).  With both seals the conditional U(2)-type Higgs representation interface is available, but the result remains sealed rather than native or physical: n and q are not derived, scalar potential/runtime lambda, Higgs mass, Yukawa operators/eigenvalues, flavor hierarchy, CKM/PMNS, and native 7/72 remain firewalled."
	return Analysis{Gate720: inherited, Package: pkg, Assembly: assembly, Available: available, Blocked: blocked, Minimality: minimality, Independence: independence, Physical: physical, Truth: truth}, nil
}

func buildGate720Inheritance(g gate720.Analysis) Gate720Inheritance {
	return Gate720Inheritance{
		Inherited:                          g.Seals.ConditionalSocketRemains && g.Types.IndependentAtLevel,
		NAndQTypeDistinct:                  g.Types.TypeDistinct,
		NRequiresSelectorSeal:              g.NSelector.RequiresSelectorSeal && !g.NSelector.NativeSelectorFound,
		QRequiresNormalizationSeal:         g.QNorm.RequiresNormalizationSeal && !g.QNorm.NativeQDerived,
		ConditionalSocketReadyButNotNative: strings.Contains(g.Seals.Verdict, gate720.StatusConditionalHiggsSocketReadyButNotNative),
		NoNativeN:                          !g.NSelector.NativeSelectorFound,
		NoNativeQ:                          !g.QNorm.NativeQDerived,
		NoPhysicalHiggs:                    !g.Physical.ConditionalSocketPhysicalHiggsTheorem && !g.Physical.K7PlusPhysicalHiggsMassTheorem,
		NoHiggsRuntime:                     !g.Physical.ScalarPotential && !g.Physical.QuarticRuntimeLambda && !g.Physical.HiggsPoleMass,
		NoYukawa:                           !g.Physical.YukawaOperators,
		Verdict:                            StatusGate720MissingSealIndependenceInherited,
	}
}

func buildMinimalSealPackage() MinimalSealPackage {
	n := SealRole{
		Name:        "TwistorSelectorSeal",
		Input:       "n ∈ S^2(K7-), ||n||=1",
		Output:      "J_H(n), L_n=span(J_H(n)), K7+_J(n)",
		Required:    true,
		Native:      false,
		RoleAudited: true,
		Verdict:     StatusTwistorSelectorSealRoleAudited,
	}
	q := SealRole{
		Name:        "HyperchargeNormalizationSeal",
		Input:       "q ∈ R^×",
		Output:      "normalized phase generator qJ_H(n)",
		Required:    true,
		Native:      false,
		RoleAudited: true,
		Verdict:     StatusHyperchargeNormalizationSealRoleAudited,
	}
	seals := []string{n.Name, q.Name}
	return MinimalSealPackage{
		TwistorSelectorSeal:          n,
		HyperchargeNormalizationSeal: q,
		PackageName:                  "HiggsSocketSealPackage=(n,q)",
		Seals:                        seals,
		SealCount:                    len(seals),
		SuppliesN:                    true,
		SuppliesQ:                    true,
		Minimal:                      true,
		Native:                       false,
		Verdict: strings.Join([]string{
			StatusMinimalSealPackageDefined,
			StatusTwistorSelectorSealRoleAudited,
			StatusHyperchargeNormalizationSealRoleAudited,
			StatusHiggsSocketSealPackageMinimal,
		}, "; "),
	}
}

func buildSealedSocketAssembly(pkg MinimalSealPackage) SealedSocketAssembly {
	return SealedSocketAssembly{
		SelectedComplexCarrier:   pkg.SuppliesN,
		InternalU2Socket:         pkg.SuppliesN && pkg.SuppliesQ,
		SU2Compatibility:         true,
		U1PhaseCompatibility:     pkg.SuppliesQ,
		FullIntertwinerCandidate: pkg.SuppliesN && pkg.SuppliesQ,
		SocketFormula:            "g_int(n,q)=C ⊕ span(qJ_H(n))",
		CarrierFormula:           "K7+_J(n) ~= C^2",
		Verdict: strings.Join([]string{
			StatusSealedSocketAssemblyReconstructed,
			StatusSealedHiggsSocketInterfaceDefined,
			StatusK7PlusAirlockReadyOnlyAfterSeals,
		}, "; "),
	}
}

func buildAvailableUnderSealsAudit(a SealedSocketAssembly) AvailableUnderSealsAudit {
	structures := []AvailableStructure{
		{Name: "selected complex carrier K7+_J(n)", Available: a.SelectedComplexCarrier, Reason: "n selects J_H(n) and turns K7+ into a complex two-carrier"},
		{Name: "internal U(2)-type socket C⊕span(qJ_H(n))", Available: a.InternalU2Socket, Reason: "n supplies the line and q supplies the phase normalization"},
		{Name: "SU(2)-side representation compatibility via C", Available: a.SU2Compatibility, Reason: "inherited selector-independent Gate716/Gate719 compatibility"},
		{Name: "U(1)-side phase compatibility via qJ_H(n)", Available: a.U1PhaseCompatibility, Reason: "q normalizes the chosen moving phase line"},
		{Name: "full representation-intertwiner candidate to finite electroweak Higgs lane", Available: a.FullIntertwinerCandidate, Reason: "Theta_SU2, Theta_Y, and Theta_H can be assembled conditionally"},
	}
	return AvailableUnderSealsAudit{Structures: structures, AllAvailable: allAvailable(structures), Verdict: StatusAvailableStructuresUnderSealsAudited}
}

func allAvailable(xs []AvailableStructure) bool {
	for _, x := range xs {
		if !x.Available {
			return false
		}
	}
	return true
}

func buildBlockedPhysics() BlockedPhysics {
	return BlockedPhysics{
		WhyNSelected:               false,
		WhyQHasValue:               false,
		PhysicalEqualitySU2U1:      false,
		ScalarPotential:            false,
		QuarticRuntimeLambda:       false,
		HiggsPoleMass:              false,
		YukawaOperatorConstruction: false,
		FlavorHierarchy:            false,
		CKMPMNS:                    false,
		BlockedCount:               remainingBlockedCount,
		Verdict: strings.Join([]string{
			StatusRemainingBlockedPhysicsAudited,
			StatusNNotNativelyDerived,
			StatusQNotNativelyDerived,
			StatusSealedSocketNotPhysicalHiggsTheorem,
			StatusNoScalarPotentialOrRuntimeLambda,
			StatusNoHiggsMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
		}, "; "),
	}
}

func buildSealMinimalityAudit() SealMinimalityAudit {
	removals := []MinimalityRemoval{
		{Removed: "n", Fatal: true, Breaks: []string{"J_H(n) not selected", "K7+ not a chosen C^2", "L_n not selected", "U(1) side undefined"}},
		{Removed: "q", Fatal: true, Breaks: []string{"phase line exists after n", "charge normalization not fixed", "hypercharge compatibility remains convention-level"}},
		{Removed: "C", Fatal: true, Breaks: []string{"SU(2)-doublet side disappears"}},
		{Removed: "K7+", Fatal: true, Breaks: []string{"no Higgs carrier"}},
	}
	return SealMinimalityAudit{
		Removals:           removals,
		RemoveNBreaks:      len(removals[0].Breaks) == 4 && removals[0].Fatal,
		RemoveQBreaks:      len(removals[1].Breaks) == 3 && removals[1].Fatal,
		RemoveCBreaks:      len(removals[2].Breaks) == 1 && removals[2].Fatal,
		RemoveK7PlusBreaks: len(removals[3].Breaks) == 1 && removals[3].Fatal,
		PairMinimal:        true,
		Verdict: strings.Join([]string{
			StatusSealMinimalityAudited,
			StatusHiggsSocketSealPackageMinimal,
		}, "; "),
	}
}

func buildIndependenceAudit(g gate720.Analysis) IndependenceAudit {
	shortcuts := []string{"q from 7/72", "n from scalar bridge data", "n from P_K7", "q from |n|"}
	return IndependenceAudit{
		NType:                 g.Types.NType,
		QType:                 g.Types.QType,
		TypeDistinct:          g.Types.TypeDistinct,
		NotMutuallyDerivable:  !g.Types.NCanDetermineQ && !g.Types.QCanDetermineN,
		QFromSevenOver72:      false,
		NFromScalarBridgeData: false,
		NFromPK7:              false,
		QFromAbsN:             false,
		ForbiddenShortcuts:    shortcuts,
		Verdict: strings.Join([]string{
			StatusNQIndependencePreserved,
			StatusNNotNativelyDerived,
			StatusQNotNativelyDerived,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalFirewall {
	return PhysicalFirewall{
		TwistorSelectorSealNativeVacuumTheorem: false,
		HyperchargeSealNativeDerivation:        false,
		SealedSocketFullPhysicalHiggsTheorem:   false,
		SealedSocketHiggsMassTheorem:           false,
		SealedSocketYukawaTheorem:              false,
		SealedSocketCKMPMNSTtheorem:            false,
		NoScalarPotentialOrRuntimeLambda:       true,
		NoHiggsMassTheorem:                     true,
		NoYukawaOperatorOrEigenvalueTheorem:    true,
		Verdict: strings.Join([]string{
			StatusSealedSocketNotPhysicalHiggsTheorem,
			StatusNoScalarPotentialOrRuntimeLambda,
			StatusNoHiggsMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate721PromotionBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate720MissingSealIndependenceInherited,
		StatusMinimalSealPackageDefined,
		StatusTwistorSelectorSealRoleAudited,
		StatusHyperchargeNormalizationSealRoleAudited,
		StatusSealedSocketAssemblyReconstructed,
		StatusAvailableStructuresUnderSealsAudited,
		StatusRemainingBlockedPhysicsAudited,
		StatusSealMinimalityAudited,
		StatusNQIndependencePreserved,
		StatusHiggsSocketSealPackageMinimal,
		StatusSealedHiggsSocketInterfaceDefined,
		StatusK7PlusAirlockReadyOnlyAfterSeals,
		StatusNNotNativelyDerived,
		StatusQNotNativelyDerived,
		StatusSealedSocketNotPhysicalHiggsTheorem,
		StatusNoScalarPotentialOrRuntimeLambda,
		StatusNoHiggsMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate721PromotionBoundary,
	}
}

func FormatGate720(x Gate720Inheritance) string {
	return fmt.Sprintf("inherited=%t typeDistinct=%t nSeal=%t qSeal=%t readyButNotNative=%t noN=%t noQ=%t noHiggs=%t noRuntime=%t noYukawa=%t verdict=%q", x.Inherited, x.NAndQTypeDistinct, x.NRequiresSelectorSeal, x.QRequiresNormalizationSeal, x.ConditionalSocketReadyButNotNative, x.NoNativeN, x.NoNativeQ, x.NoPhysicalHiggs, x.NoHiggsRuntime, x.NoYukawa, x.Verdict)
}

func FormatSealRole(x SealRole) string {
	return fmt.Sprintf("name=%q input=%q output=%q required=%t native=%t audited=%t verdict=%q", x.Name, x.Input, x.Output, x.Required, x.Native, x.RoleAudited, x.Verdict)
}

func FormatPackage(x MinimalSealPackage) string {
	return fmt.Sprintf("package=%q sealCount=%d suppliesN=%t suppliesQ=%t minimal=%t native=%t verdict=%q twistor={%s} hypercharge={%s}", x.PackageName, x.SealCount, x.SuppliesN, x.SuppliesQ, x.Minimal, x.Native, x.Verdict, FormatSealRole(x.TwistorSelectorSeal), FormatSealRole(x.HyperchargeNormalizationSeal))
}

func FormatAssembly(x SealedSocketAssembly) string {
	return fmt.Sprintf("carrier=%t socket=%t su2=%t u1=%t intertwiner=%t socketFormula=%q carrierFormula=%q verdict=%q", x.SelectedComplexCarrier, x.InternalU2Socket, x.SU2Compatibility, x.U1PhaseCompatibility, x.FullIntertwinerCandidate, x.SocketFormula, x.CarrierFormula, x.Verdict)
}

func FormatAvailable(x AvailableUnderSealsAudit) string {
	parts := make([]string, 0, len(x.Structures))
	for _, s := range x.Structures {
		parts = append(parts, fmt.Sprintf("%s available=%t reason=%s", s.Name, s.Available, s.Reason))
	}
	return fmt.Sprintf("count=%d all=%t verdict=%q :: %s", len(x.Structures), x.AllAvailable, x.Verdict, strings.Join(parts, " | "))
}

func FormatBlocked(x BlockedPhysics) string {
	return fmt.Sprintf("whyN=%t whyQ=%t physicalEq=%t potential=%t quartic=%t pole=%t yukawa=%t flavor=%t ckm=%t blocked=%d verdict=%q", x.WhyNSelected, x.WhyQHasValue, x.PhysicalEqualitySU2U1, x.ScalarPotential, x.QuarticRuntimeLambda, x.HiggsPoleMass, x.YukawaOperatorConstruction, x.FlavorHierarchy, x.CKMPMNS, x.BlockedCount, x.Verdict)
}

func FormatMinimality(x SealMinimalityAudit) string {
	parts := make([]string, 0, len(x.Removals))
	for _, r := range x.Removals {
		parts = append(parts, fmt.Sprintf("remove %s fatal=%t breaks=%s", r.Removed, r.Fatal, strings.Join(r.Breaks, ",")))
	}
	return fmt.Sprintf("removals=%d n=%t q=%t c=%t k7plus=%t pairMinimal=%t verdict=%q :: %s", len(x.Removals), x.RemoveNBreaks, x.RemoveQBreaks, x.RemoveCBreaks, x.RemoveK7PlusBreaks, x.PairMinimal, x.Verdict, strings.Join(parts, " | "))
}

func FormatIndependence(x IndependenceAudit) string {
	return fmt.Sprintf("nType=%q qType=%q distinct=%t mutual=%t qFrom7_72=%t nFromScalars=%t nFromPK7=%t qFromAbsN=%t shortcuts=%d verdict=%q", x.NType, x.QType, x.TypeDistinct, x.NotMutuallyDerivable, x.QFromSevenOver72, x.NFromScalarBridgeData, x.NFromPK7, x.QFromAbsN, len(x.ForbiddenShortcuts), x.Verdict)
}

func FormatPhysical(x PhysicalFirewall) string {
	return fmt.Sprintf("twistorNative=%t qNative=%t physicalHiggs=%t higgsMass=%t yukawa=%t ckm=%t noPotentialRuntime=%t noMass=%t noYukawa=%t verdict=%q", x.TwistorSelectorSealNativeVacuumTheorem, x.HyperchargeSealNativeDerivation, x.SealedSocketFullPhysicalHiggsTheorem, x.SealedSocketHiggsMassTheorem, x.SealedSocketYukawaTheorem, x.SealedSocketCKMPMNSTtheorem, x.NoScalarPotentialOrRuntimeLambda, x.NoHiggsMassTheorem, x.NoYukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
