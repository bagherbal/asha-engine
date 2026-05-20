// Package generation2higgssocketmissingsealindependenceandsourcecandidateaudit implements
// Gate 720: Higgs Socket Missing-Seal Independence and Source-Candidate Audit.
//
// Gate 719 assembled the conditional internal electroweak Higgs socket
// g_int(n,q)=C ⊕ span(qJ_H(n)) acting on K7+_J(n) ~= C^2.  Gate 720 audits
// the two missing representation choices left by that assembly: n, a twistor
// point / complex-structure selector in S^2(K7-), and q, a scalar phase-line /
// hypercharge normalization.  It checks their source candidates, proves their
// type distinction at the current bridge layer, blocks scalar/event-probability
// shortcuts, and classifies them as separate seals.  It deliberately does not
// derive physical SU(2)_L×U(1)_Y, hypercharge, Higgs mass, scalar runtime,
// Yukawa operators/eigenvalues, CKM/PMNS, flavor hierarchy, or a native 7/72
// theorem.
package generation2higgssocketmissingsealindependenceandsourcecandidateaudit

import (
	"fmt"
	"strings"
	"sync"

	gate719 "github.com/bagherbal/asha-engine/pkg/bridge/generation2conditionalelectroweakhiggssocketassemblyandmissingsealaudit"
)

const (
	AuditID = "GATE720-HIGGS-SOCKET-MISSING-SEAL-INDEPENDENCE-AND-SOURCE-CANDIDATE-AUDIT"

	StatusGate719ConditionalHiggsSocketInherited  = "PASS_GATE719_CONDITIONAL_HIGGS_SOCKET_INHERITED"
	StatusNSelectorSourceCandidatesAudited        = "PASS_N_SELECTOR_SOURCE_CANDIDATES_AUDITED"
	StatusQNormalizationSourceCandidatesAudited   = "PASS_Q_NORMALIZATION_SOURCE_CANDIDATES_AUDITED"
	StatusNAndQTypeDistinctionAudited             = "PASS_N_AND_Q_TYPE_DISTINCTION_AUDITED"
	StatusForbiddenShortcutsAudited               = "PASS_FORBIDDEN_SHORTCUTS_AUDITED"
	StatusMissingSealClassificationDefined        = "PASS_MISSING_SEAL_CLASSIFICATION_DEFINED"
	StatusPhysicalFirewallEnforced                = "PASS_PHYSICAL_FIREWALL_ENFORCED"
	StatusNAndQTypeDistinctMissingSeals           = "CONDITIONAL_SUPPORT_N_AND_Q_ARE_TYPE_DISTINCT_MISSING_SEALS"
	StatusNRequiresTwistorOrVacuumSelectorSeal    = "CONDITIONAL_SUPPORT_N_REQUIRES_TWISTOR_OR_VACUUM_SELECTOR_SEAL"
	StatusQRequiresHyperchargeNormalizationSeal   = "CONDITIONAL_SUPPORT_Q_REQUIRES_HYPERCHARGE_NORMALIZATION_SEAL"
	StatusConditionalHiggsSocketReadyButNotNative = "CONDITIONAL_SUPPORT_CONDITIONAL_HIGGS_SOCKET_IS_STRUCTURALLY_READY_BUT_NOT_NATIVE"
	StatusNoNativeTwistorSelectorN                = "FAILED_ROUTE_NO_NATIVE_TWISTOR_SELECTOR_N"
	StatusNoNativeHyperchargeNormalizationQ       = "FAILED_ROUTE_NO_NATIVE_HYPERCHARGE_NORMALIZATION_Q"
	StatusScalarBridgeDataDoNotSelectN            = "FAILED_ROUTE_SCALAR_BRIDGE_DATA_DO_NOT_SELECT_N"
	StatusK7EventProbabilityDoesNotFixQ           = "FAILED_ROUTE_K7_EVENT_PROBABILITY_DOES_NOT_FIX_Q"
	StatusNoFullPhysicalHiggsDoubletTheorem       = "FAILED_ROUTE_NO_FULL_PHYSICAL_HIGGS_DOUBLET_THEOREM"
	StatusNoHiggsMassOrScalarRuntimeTheorem       = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate720MissingSealIndependenceBoundary  = "FIREWALL_PRESERVED_GATE720_MISSING_SEAL_INDEPENDENCE_BOUNDARY"
)

const (
	nSourceCandidateCount  = 7
	qSourceCandidateCount  = 4
	forbiddenShortcutCount = 5
	sealCount              = 2
)

type Gate719Inheritance struct {
	Inherited                   bool
	SocketAssembled             bool
	RequiresN                   bool
	RequiresQ                   bool
	RepresentationCompatible    bool
	NativeTwistorSelector       bool
	NativeHyperchargeNorm       bool
	CanonicalThetaH             bool
	PhysicalHiggsDoubletTheorem bool
	HiggsMassOrRuntime          bool
	YukawaOperatorOrEigenvalue  bool
	Verdict                     string
}

type SourceCandidate struct {
	Name          string
	SourceType    string
	CandidateRole string
	TypedMapFound bool
	NativeSource  bool
	Rejected      bool
	Reason        string
}

type NSelectorSourceAudit struct {
	LivesInS2K7Minus       bool
	SelectsJH              bool
	SelectsPhaseLine       bool
	SelectsComplexCarrier  bool
	Candidates             []SourceCandidate
	NativeSelectorFound    bool
	RequiresSelectorSeal   bool
	ExternalCandidateFound bool
	Verdict                string
}

type QNormalizationSourceAudit struct {
	LivesInRNonzero           bool
	NormalizesPhaseGenerator  bool
	Candidates                []SourceCandidate
	CanMatchTargetConvention  bool
	NativeQDerived            bool
	RequiresNormalizationSeal bool
	Verdict                   string
}

type TypeDistinctionAudit struct {
	NType                 string
	QType                 string
	TypeDistinct          bool
	ChangingNChangesLine  bool
	ChangingQRescalesLine bool
	NCanDetermineQ        bool
	QCanDetermineN        bool
	IndependentAtLevel    bool
	Verdict               string
}

type ForbiddenShortcut struct {
	Name       string
	Invalid    bool
	Reason     string
	Firewalled bool
}

type ForbiddenShortcutsAudit struct {
	Shortcuts               []ForbiddenShortcut
	ScalarQuantitiesSelectN bool
	EventProbabilityFixesQ  bool
	AllShortcutsRejected    bool
	Verdict                 string
}

type MissingSealClassification struct {
	TwistorSelectorSeal          string
	HyperchargeNormalizationSeal string
	Seals                        []string
	ConditionalSocketRemains     bool
	DerivedNative                bool
	Verdict                      string
}

type PhysicalFirewall struct {
	ConditionalSocketPhysicalHiggsTheorem bool
	MatchedQDerivedHypercharge            bool
	ChosenNDerivedVacuumOrientation       bool
	K7MinusSelectorFlavorHierarchy        bool
	K7PlusPhysicalHiggsMassTheorem        bool
	ScalarPotential                       bool
	QuarticRuntimeLambda                  bool
	HiggsPoleMass                         bool
	YukawaOperators                       bool
	FlavorHierarchy                       bool
	CKMPMNS                               bool
	Verdict                               string
}

type Analysis struct {
	Gate719   Gate719Inheritance
	NSelector NSelectorSourceAudit
	QNorm     QNormalizationSourceAudit
	Types     TypeDistinctionAudit
	Shortcuts ForbiddenShortcutsAudit
	Seals     MissingSealClassification
	Physical  PhysicalFirewall
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
	g719, err := gate719.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate719 inheritance unavailable: %w", err)
	}
	inherited := buildGate719Inheritance(g719)
	nAudit := buildNSelectorSourceAudit()
	qAudit := buildQNormalizationSourceAudit()
	typeAudit := buildTypeDistinctionAudit(nAudit, qAudit)
	shortcuts := buildForbiddenShortcutsAudit()
	seals := buildMissingSealClassification(nAudit, qAudit, typeAudit)
	physical := buildPhysicalFirewall()
	truth := "Gate 720 audits the two missing choices left by the conditional Higgs socket g_int(n,q)=C ⊕ span(qJ_H(n)).  The twistor selector n is a direction in S^2(K7-) that changes the complex structure, phase line, and complex carrier; the normalization q is a scalar in R^× that rescales the chosen phase generator.  Existing native Hodge/Fano data, scalar-wall data, history defects, and K7 event probability do not source these choices.  At the current theorem level n and q are type-distinct missing seals: a TwistorSelectorSeal and a HyperchargeNormalizationSeal.  The conditional socket is structurally ready but remains non-native and non-physical."
	return Analysis{Gate719: inherited, NSelector: nAudit, QNorm: qAudit, Types: typeAudit, Shortcuts: shortcuts, Seals: seals, Physical: physical, Truth: truth}, nil
}

func buildGate719Inheritance(g gate719.Analysis) Gate719Inheritance {
	return Gate719Inheritance{
		Inherited:                   g.Socket.Assembled && g.Intertwiner.RepresentationCompatible,
		SocketAssembled:             g.Socket.Assembled,
		RequiresN:                   g.Socket.RequiresN && g.Intertwiner.RequiresN && g.Choices.TwistorPointN,
		RequiresQ:                   g.Socket.RequiresQ && g.Intertwiner.RequiresQ && g.Choices.PhaseNormalizationQ,
		RepresentationCompatible:    g.Intertwiner.RepresentationCompatible,
		NativeTwistorSelector:       g.U1Inherited.NativeTwistorSelector || g.Physical.NDerivedVacuumSelector,
		NativeHyperchargeNorm:       g.U1Inherited.NativeThetaYNormalization || g.Hypercharge.QDerivedNatively || g.Physical.QDerivedHypercharge,
		CanonicalThetaH:             g.Choices.CanonicalThetaH,
		PhysicalHiggsDoubletTheorem: g.Physical.K7PlusPhysicalHiggsDoublet,
		HiggsMassOrRuntime:          g.Physical.HiggsPoleMass || g.Physical.QuarticRuntimeLambda || g.Physical.ScalarPotential,
		YukawaOperatorOrEigenvalue:  g.Physical.YukawaOperator,
		Verdict:                     StatusGate719ConditionalHiggsSocketInherited,
	}
}

func buildNSelectorSourceAudit() NSelectorSourceAudit {
	candidates := []SourceCandidate{
		{Name: "Hodge polarity", SourceType: "native split", CandidateRole: "separates K7+ from K7-", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "gives K7+ ⊕ K7- but no vector inside K7-"},
		{Name: "Fano volume eta_123", SourceType: "native orientation", CandidateRole: "orients K7-", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "orientation/volume does not select a unit axis"},
		{Name: "Fano frame eta_a", SourceType: "SO(3)-covariant frame", CandidateRole: "three-channel frame", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "frame is defined up to SO(3), not a canonical point on S^2"},
		{Name: "Boundary scalar S_split", SourceType: "scalar bridge coordinate", CandidateRole: "boundary split amplitude", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "a scalar carries no K7- direction"},
		{Name: "Scalar-wall airlock lambda", SourceType: "scalar normalization anchor", CandidateRole: "shared signed wall unit", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "scalar-wall unit carries no axis in K7-"},
		{Name: "History defects kappa_lambda,kappa_e", SourceType: "scalar history coordinates", CandidateRole: "history readout data", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "scalar deficits do not define a K7- vector"},
		{Name: "OrientationBalanceSeal / flavor wall", SourceType: "external bridge candidate", CandidateRole: "possible environmental selector source", TypedMapFound: false, NativeSource: false, Rejected: false, Reason: "candidate only; no typed map into K7- exists yet"},
	}
	return NSelectorSourceAudit{
		LivesInS2K7Minus:       true,
		SelectsJH:              true,
		SelectsPhaseLine:       true,
		SelectsComplexCarrier:  true,
		Candidates:             candidates,
		NativeSelectorFound:    false,
		RequiresSelectorSeal:   true,
		ExternalCandidateFound: true,
		Verdict: strings.Join([]string{
			StatusNSelectorSourceCandidatesAudited,
			StatusNRequiresTwistorOrVacuumSelectorSeal,
			StatusNoNativeTwistorSelectorN,
		}, "; "),
	}
}

func buildQNormalizationSourceAudit() QNormalizationSourceAudit {
	candidates := []SourceCandidate{
		{Name: "Target Higgs hypercharge convention Y_H=1/2", SourceType: "target convention", CandidateRole: "matches q to physical lane convention", TypedMapFound: true, NativeSource: false, Rejected: false, Reason: "can fix q relative to the target, but this is matching not derivation"},
		{Name: "Spectral-triple hypercharge normalization", SourceType: "finite electroweak target lane", CandidateRole: "supplies target normalization", TypedMapFound: true, NativeSource: false, Rejected: false, Reason: "target lane may normalize U(1)_Y but does not derive why internal L_n chooses that q"},
		{Name: "Generator norm convention", SourceType: "basis/norm convention", CandidateRole: "sets ||qJ_H||", TypedMapFound: true, NativeSource: false, Rejected: true, Reason: "norm convention alone is not physical hypercharge"},
		{Name: "Gauge kinetic normalization", SourceType: "coupling normalization candidate", CandidateRole: "may constrain coupling units", TypedMapFound: false, NativeSource: false, Rejected: true, Reason: "does not assign the internal phase charge without a typed intertwiner"},
	}
	return QNormalizationSourceAudit{
		LivesInRNonzero:           true,
		NormalizesPhaseGenerator:  true,
		Candidates:                candidates,
		CanMatchTargetConvention:  true,
		NativeQDerived:            false,
		RequiresNormalizationSeal: true,
		Verdict: strings.Join([]string{
			StatusQNormalizationSourceCandidatesAudited,
			StatusQRequiresHyperchargeNormalizationSeal,
			StatusNoNativeHyperchargeNormalizationQ,
		}, "; "),
	}
}

func buildTypeDistinctionAudit(n NSelectorSourceAudit, q QNormalizationSourceAudit) TypeDistinctionAudit {
	return TypeDistinctionAudit{
		NType:                 "n ∈ S^2(K7-) : twistor point / complex-structure selector",
		QType:                 "q ∈ R^× : scalar normalization on L_n / charge convention",
		TypeDistinct:          n.LivesInS2K7Minus && q.LivesInRNonzero,
		ChangingNChangesLine:  true,
		ChangingQRescalesLine: true,
		NCanDetermineQ:        false,
		QCanDetermineN:        false,
		IndependentAtLevel:    true,
		Verdict: strings.Join([]string{
			StatusNAndQTypeDistinctionAudited,
			StatusNAndQTypeDistinctMissingSeals,
		}, "; "),
	}
}

func buildForbiddenShortcutsAudit() ForbiddenShortcutsAudit {
	shortcuts := []ForbiddenShortcut{
		{Name: "q from |n|", Invalid: true, Reason: "|n|=1 by construction and carries no charge normalization", Firewalled: true},
		{Name: "n from q", Invalid: true, Reason: "a scalar normalization cannot select a direction on S^2(K7-)", Firewalled: true},
		{Name: "n from lambda or S_split", Invalid: true, Reason: "a scalar bridge coordinate cannot select a K7- direction without a typed scalar->vector map", Firewalled: true},
		{Name: "q from 7/72", Invalid: true, Reason: "7/72 is K7 event probability, not hypercharge normalization", Firewalled: true},
		{Name: "n from K7 event support", Invalid: true, Reason: "P_K7 selects the whole 7D carrier, not an axis inside K7-", Firewalled: true},
	}
	return ForbiddenShortcutsAudit{
		Shortcuts:               shortcuts,
		ScalarQuantitiesSelectN: false,
		EventProbabilityFixesQ:  false,
		AllShortcutsRejected:    allShortcutsRejected(shortcuts),
		Verdict: strings.Join([]string{
			StatusForbiddenShortcutsAudited,
			StatusScalarBridgeDataDoNotSelectN,
			StatusK7EventProbabilityDoesNotFixQ,
		}, "; "),
	}
}

func allShortcutsRejected(xs []ForbiddenShortcut) bool {
	for _, x := range xs {
		if !x.Invalid || !x.Firewalled {
			return false
		}
	}
	return true
}

func buildMissingSealClassification(n NSelectorSourceAudit, q QNormalizationSourceAudit, types TypeDistinctionAudit) MissingSealClassification {
	seals := []string{
		"TwistorSelectorSeal: supplies n ∈ S^2(K7-)",
		"HyperchargeNormalizationSeal: supplies q ∈ R^× for qJ_H(n)",
	}
	return MissingSealClassification{
		TwistorSelectorSeal:          seals[0],
		HyperchargeNormalizationSeal: seals[1],
		Seals:                        seals,
		ConditionalSocketRemains:     n.RequiresSelectorSeal && q.RequiresNormalizationSeal && types.IndependentAtLevel,
		DerivedNative:                false,
		Verdict: strings.Join([]string{
			StatusMissingSealClassificationDefined,
			StatusNRequiresTwistorOrVacuumSelectorSeal,
			StatusQRequiresHyperchargeNormalizationSeal,
			StatusConditionalHiggsSocketReadyButNotNative,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalFirewall {
	return PhysicalFirewall{
		ConditionalSocketPhysicalHiggsTheorem: false,
		MatchedQDerivedHypercharge:            false,
		ChosenNDerivedVacuumOrientation:       false,
		K7MinusSelectorFlavorHierarchy:        false,
		K7PlusPhysicalHiggsMassTheorem:        false,
		ScalarPotential:                       false,
		QuarticRuntimeLambda:                  false,
		HiggsPoleMass:                         false,
		YukawaOperators:                       false,
		FlavorHierarchy:                       false,
		CKMPMNS:                               false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallEnforced,
			StatusNoFullPhysicalHiggsDoubletTheorem,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate720MissingSealIndependenceBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate719ConditionalHiggsSocketInherited,
		StatusNSelectorSourceCandidatesAudited,
		StatusQNormalizationSourceCandidatesAudited,
		StatusNAndQTypeDistinctionAudited,
		StatusForbiddenShortcutsAudited,
		StatusMissingSealClassificationDefined,
		StatusPhysicalFirewallEnforced,
		StatusNAndQTypeDistinctMissingSeals,
		StatusNRequiresTwistorOrVacuumSelectorSeal,
		StatusQRequiresHyperchargeNormalizationSeal,
		StatusConditionalHiggsSocketReadyButNotNative,
		StatusNoNativeTwistorSelectorN,
		StatusNoNativeHyperchargeNormalizationQ,
		StatusScalarBridgeDataDoNotSelectN,
		StatusK7EventProbabilityDoesNotFixQ,
		StatusNoFullPhysicalHiggsDoubletTheorem,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate720MissingSealIndependenceBoundary,
	}
}

func FormatGate719(x Gate719Inheritance) string {
	return fmt.Sprintf("inherited=%t assembled=%t n=%t q=%t compatible=%t nativeN=%t nativeQ=%t thetaH=%t higgs=%t mass=%t yukawa=%t verdict=%q", x.Inherited, x.SocketAssembled, x.RequiresN, x.RequiresQ, x.RepresentationCompatible, x.NativeTwistorSelector, x.NativeHyperchargeNorm, x.CanonicalThetaH, x.PhysicalHiggsDoubletTheorem, x.HiggsMassOrRuntime, x.YukawaOperatorOrEigenvalue, x.Verdict)
}

func FormatCandidates(xs []SourceCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s[type=%s role=%s map=%t native=%t rejected=%t reason=%s]", x.Name, x.SourceType, x.CandidateRole, x.TypedMapFound, x.NativeSource, x.Rejected, x.Reason))
	}
	return strings.Join(parts, " | ")
}

func FormatNSelector(x NSelectorSourceAudit) string {
	return fmt.Sprintf("S2=%t JH=%t line=%t carrier=%t candidates=%d native=%t seal=%t externalCandidate=%t verdict=%q :: %s", x.LivesInS2K7Minus, x.SelectsJH, x.SelectsPhaseLine, x.SelectsComplexCarrier, len(x.Candidates), x.NativeSelectorFound, x.RequiresSelectorSeal, x.ExternalCandidateFound, x.Verdict, FormatCandidates(x.Candidates))
}

func FormatQNorm(x QNormalizationSourceAudit) string {
	return fmt.Sprintf("R*=%t phase=%t candidates=%d match=%t native=%t seal=%t verdict=%q :: %s", x.LivesInRNonzero, x.NormalizesPhaseGenerator, len(x.Candidates), x.CanMatchTargetConvention, x.NativeQDerived, x.RequiresNormalizationSeal, x.Verdict, FormatCandidates(x.Candidates))
}

func FormatTypes(x TypeDistinctionAudit) string {
	return fmt.Sprintf("nType=%q qType=%q distinct=%t changeN=%t changeQ=%t nToQ=%t qToN=%t independent=%t verdict=%q", x.NType, x.QType, x.TypeDistinct, x.ChangingNChangesLine, x.ChangingQRescalesLine, x.NCanDetermineQ, x.QCanDetermineN, x.IndependentAtLevel, x.Verdict)
}

func FormatShortcuts(x ForbiddenShortcutsAudit) string {
	parts := make([]string, 0, len(x.Shortcuts))
	for _, s := range x.Shortcuts {
		parts = append(parts, fmt.Sprintf("%s invalid=%t firewalled=%t reason=%s", s.Name, s.Invalid, s.Firewalled, s.Reason))
	}
	return fmt.Sprintf("count=%d scalarsSelectN=%t pFixesQ=%t allRejected=%t verdict=%q :: %s", len(x.Shortcuts), x.ScalarQuantitiesSelectN, x.EventProbabilityFixesQ, x.AllShortcutsRejected, x.Verdict, strings.Join(parts, " | "))
}

func FormatSeals(x MissingSealClassification) string {
	return fmt.Sprintf("twistor=%q hypercharge=%q seals=%d conditional=%t derived=%t verdict=%q", x.TwistorSelectorSeal, x.HyperchargeNormalizationSeal, len(x.Seals), x.ConditionalSocketRemains, x.DerivedNative, x.Verdict)
}

func FormatPhysical(x PhysicalFirewall) string {
	return fmt.Sprintf("socketPhysical=%t qDerived=%t nDerived=%t flavorAxis=%t higgsMass=%t potential=%t quartic=%t pole=%t yukawa=%t flavor=%t ckm=%t verdict=%q", x.ConditionalSocketPhysicalHiggsTheorem, x.MatchedQDerivedHypercharge, x.ChosenNDerivedVacuumOrientation, x.K7MinusSelectorFlavorHierarchy, x.K7PlusPhysicalHiggsMassTheorem, x.ScalarPotential, x.QuarticRuntimeLambda, x.HiggsPoleMass, x.YukawaOperators, x.FlavorHierarchy, x.CKMPMNS, x.Verdict)
}
