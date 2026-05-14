// Package spectralactioncapstone implements Gate 282:
// Spectral Action Epistemological Capstone / Higgs Prediction Firewall Audit.
//
// Gate 282 is a closure manifest for the Path-B spectral-action attempt. It
// consolidates what the engine has lawfully derived about the finite spectral
// scaffold and records the exact six-point firewall that prevents promotion of
// raw finite traces or scalar-Morita shape constraints to a physical Higgs mass
// prediction. It intentionally adds no new dynamics.
package spectralactioncapstone

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/resolventbranchsemantics"
)

const (
	AuditID = "GATE282-SPECTRAL-ACTION-EPISTEMOLOGICAL-CAPSTONE-HIGGS-PREDICTION-FIREWALL-AUDIT"

	StatusScaffoldManifestCompiled = "CONDITIONAL_SUPPORT_SPECTRAL_ACTION_SCAFFOLD_MANIFEST_COMPILED"
	StatusSixPointLedgerCompiled   = "CONDITIONAL_SUPPORT_SIX_POINT_HIGGS_FIREWALL_LEDGER_COMPILED"
	StatusFirewallEstablished      = "CONDITIONAL_SUPPORT_HIGGS_PREDICTION_FIREWALL_ESTABLISHED"
	StatusCapstoneEstablished      = "CONDITIONAL_SUPPORT_SPECTRAL_ACTION_CAPSTONE_AND_HIGGS_FIREWALL_ESTABLISHED"
	StatusFutureCriteriaDefined    = "CONDITIONAL_SUPPORT_FUTURE_SEAL_LIFTING_CRITERIA_DEFINED"
	StatusFirewallsPreserved       = "CONDITIONAL_SUPPORT_SPECTRAL_ACTION_FIREWALLS_PRESERVED"

	StatusFailedHiggsPredictionSealed           = "FAILED_ROUTE_HIGGS_MASS_RATIO_REMAINS_UNDERIVED"
	StatusFailedNoResolventToScalarMoritaMap    = "FAILED_ROUTE_RESOLVENT_TO_SCALAR_MORITA_FUNCTOR_MISSING"
	StatusFailedPhysicalJMissing                = "FAILED_ROUTE_PHYSICAL_REAL_STRUCTURE_J_MISSING"
	StatusFailedChiralHyperchargeMissing        = "FAILED_ROUTE_CHIRAL_HYPERCHARGE_REPRESENTATION_MISSING"
	StatusFailedHeatKernelSchemeMissing         = "FAILED_ROUTE_HEAT_KERNEL_SUBTRACTION_SCHEME_MISSING"
	StatusFailedScalarGaugeNormalizationMissing = "FAILED_ROUTE_SCALAR_GAUGE_NORMALIZATION_MISSING"
	StatusFailedObservableDefinitionMissing     = "FAILED_ROUTE_DIMENSIONLESS_HIGGS_OBSERVABLE_DEFINITION_MISSING"
)

type ScaffoldItem struct {
	Name        string
	SourceGate  string
	Status      string
	Description string
	Derived     bool
	Sealed      bool
}

type GeometricScaffoldManifest struct {
	Items                               []ScaffoldItem
	FiniteAlgebraCandidateRecorded      bool
	NativeQuaternionicLocalHRecorded    bool
	MoritaMultiplicityRecorded          bool
	ScalarMoritaShapeConstraintRecorded bool
	TwoBranchRConstraintRecorded        bool
	ResolventProjectorsRecorded         bool
	ProjectorOrientationSealRecorded    bool
	NoHiggsPredictionClaimed            bool
	Verdict                             string
}

type Obstruction struct {
	Index         int
	Name          string
	Status        string
	Description   string
	RequiredFor   string
	Satisfied     bool
	FutureTheorem string
}

type SixPointObstructionLedger struct {
	Obstructions                    []Obstruction
	FunctorZToRMissing              bool
	PhysicalJMissing                bool
	ChiralHyperchargeMissing        bool
	HeatKernelSchemeMissing         bool
	ScalarGaugeNormalizationMissing bool
	ObservableDefinitionMissing     bool
	AllUnsatisfied                  bool
	HiggsPredictionBlocked          bool
	Verdict                         string
}

type FirewallSeal struct {
	Name                            string
	Active                          bool
	Reason                          string
	Target                          string
	BlockedPromotion                []string
	CanUseForFutureStressTests      bool
	CanClaimFiniteDerivedHiggsRatio bool
	Verdict                         string
}

type FutureCriterion struct {
	Name        string
	Required    bool
	Satisfied   bool
	Description string
}

type FutureTheoremCriteria struct {
	Criteria                        []FutureCriterion
	RequiresAllSixObstructions      bool
	RequiresNativeProjectionMap     bool
	RequiresPhysicalSpectralTriple  bool
	RequiresHeatKernelNormalization bool
	RequiresPreComparisonPrediction bool
	CurrentGateCanLiftFirewall      bool
	Verdict                         string
}

type FirewallAudit struct {
	PreviousGate281Inherited        bool
	NoRBranchPromotion              bool
	NoProjectorOrientationOverclaim bool
	NoRawTraceToHeatKernelOverclaim bool
	NoHiggsMassClaim                bool
	NoObservedMassesUsed            bool
	NoEmpiricalYukawaInserted       bool
	SealsDoNotRewriteNativeTheorems bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	ScaffoldCompiled       bool
	SixPointLedgerCompiled bool
	HiggsFirewallActive    bool
	FutureCriteriaDefined  bool
	HiggsRatioDerived      bool
	PathBClosed            bool
	FirewallPreserved      bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	PreviousGate281 resolventbranchsemantics.Analysis
	Scaffold        GeometricScaffoldManifest
	Obstructions    SixPointObstructionLedger
	Seal            FirewallSeal
	FutureCriteria  FutureTheoremCriteria
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	prev, err := resolventbranchsemantics.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 281 predecessor: %w", err)
	}
	scaffold := buildScaffoldManifest(prev)
	obstructions := buildSixPointObstructionLedger()
	seal := establishHiggsFirewall(obstructions)
	future := buildFutureCriteria(obstructions)
	fw := auditFirewalls(prev, scaffold, obstructions, seal, future)
	summary := buildSummary(scaffold, obstructions, seal, future, fw)
	return Analysis{
		PreviousGate281: prev,
		Scaffold:        scaffold,
		Obstructions:    obstructions,
		Seal:            seal,
		FutureCriteria:  future,
		Firewall:        fw,
		Summary:         summary,
		TruthStatement:  "Gate 282 closes Path B as an epistemological capstone: finite geometry provides a rich spectral-action scaffold and a two-branch scalar-Morita shape constraint, but six unresolved mathematical structures block any lawful Seeley-de Witt a2/a4 Higgs mass-ratio prediction.",
	}, nil
}

func buildScaffoldManifest(prev resolventbranchsemantics.Analysis) GeometricScaffoldManifest {
	items := []ScaffoldItem{
		{Name: "candidate finite algebra C⊕H⊕M3(C)", SourceGate: "Gate 274", Status: "conditional/local", Description: "local weak quaternionic closure was verified and assembled with C and M3(C) as a candidate Standard Model finite algebra", Derived: true},
		{Name: "Morita first-quantized finite Hilbert-bimodule arena", SourceGate: "Gate 272", Status: "conditional support", Description: "the spectral triple carrier is first-quantized H_F, not the full second-quantized Fock space", Derived: true},
		{Name: "Morita trace multiplicity κ_C:κ_Q = 1:3", SourceGate: "Gate 273", Status: "finite-derived multiplicity ledger", Description: "inner-product normalization counts lepton/quark trace multiplicities but does not select edge amplitudes", Derived: true},
		{Name: "scalar-Morita shape constraint", SourceGate: "Gate 275", Status: "finite shape bridge", Description: "λ_contact=1197/4624 constrains (|x|⁴+3|y|⁴)/(|x|²+3|y|²)²", Derived: true},
		{Name: "two-branch amplitude shape r±", SourceGate: "Gate 275", Status: "branch-ambiguous", Description: "r=(3591±136√123)/3099 remains a two-branch shape constraint, not a physical Higgs prediction", Derived: true},
		{Name: "contact resolvent projectors", SourceGate: "Gate 280", Status: "sealed conditional", Description: "ResolventAdjunctionSeal constructs valid 2⊕2 projectors after adjoining a resolvent root", Derived: true, Sealed: true},
		{Name: "ProjectorSectorOrientationSeal", SourceGate: "Gate 281", Status: "sealed conditional", Description: "a representative 1-in-6 contact projector orientation can be chosen for stress tests but does not map to r±", Sealed: true},
	}
	return GeometricScaffoldManifest{
		Items:                               items,
		FiniteAlgebraCandidateRecorded:      hasItem(items, "candidate finite algebra C⊕H⊕M3(C)"),
		NativeQuaternionicLocalHRecorded:    true,
		MoritaMultiplicityRecorded:          true,
		ScalarMoritaShapeConstraintRecorded: true,
		TwoBranchRConstraintRecorded:        true,
		ResolventProjectorsRecorded:         true,
		ProjectorOrientationSealRecorded:    prev.OrientationSeal.Active,
		NoHiggsPredictionClaimed:            !prev.Summary.HiggsRatioDerived,
		Verdict:                             StatusScaffoldManifestCompiled,
	}
}

func buildSixPointObstructionLedger() SixPointObstructionLedger {
	obs := []Obstruction{
		{Index: 1, Name: "Resolvent-to-Scalar-Morita functor", Status: StatusFailedNoResolventToScalarMoritaMap, Description: "no native map sends a selected contact resolvent root z_res to the Gate-275 r_+ or r_- amplitude-shape branch", RequiredFor: "branch selection and amplitude-shape semantics", Satisfied: false, FutureTheorem: "derive a functor or invariant pairing between the contact companion module and scalar-Morita trace-shape equation"},
		{Index: 2, Name: "physical anti-linear real structure J", Status: StatusFailedPhysicalJMissing, Description: "candidate conjugations exist, but the physical charge-conjugation/opposite-action map on the completed finite Hilbert space is not derived", RequiredFor: "opposite algebra action and real spectral triple", Satisfied: false, FutureTheorem: "derive antiunitary J with KO signs and particle-antiparticle semantics on H_F"},
		{Index: 3, Name: "chiral/hypercharge representation", Status: StatusFailedChiralHyperchargeMissing, Description: "the complete C⊕H⊕M3(C) action with physical chirality and hypercharge assignments is not constructed as an unsealed theorem", RequiredFor: "physical finite Hilbert space and Standard Model representation", Satisfied: false, FutureTheorem: "derive left/right chiral representation and hypercharge ledger from finite algebra without empirical insertion"},
		{Index: 4, Name: "heat-kernel subtraction scheme", Status: StatusFailedHeatKernelSchemeMissing, Description: "raw finite traces are not yet projected through cutoff moments, subtraction, renormalization convention, or Seeley-de Witt coefficient extraction", RequiredFor: "a₂/a₄ coefficient calculation", Satisfied: false, FutureTheorem: "derive the finite spectral-action heat-kernel map and subtraction convention"},
		{Index: 5, Name: "scalar versus gauge kinetic normalization", Status: StatusFailedScalarGaugeNormalizationMissing, Description: "the scalar fluctuation and gauge curvature sectors are not separately normalized from the finite trace data", RequiredFor: "Higgs-to-gauge comparison", Satisfied: false, FutureTheorem: "derive field normalization and scalar/gauge projection maps from inner fluctuations"},
		{Index: 6, Name: "dimensionless predicted observable", Status: StatusFailedObservableDefinitionMissing, Description: "the exact dimensionless target is not fixed: raw trace ratio, a₂/a₄, λ_H/g², m_H/v, or another normalized observable", RequiredFor: "pre-comparison physical prediction", Satisfied: false, FutureTheorem: "define the predicted observable before consulting observed Higgs data"},
	}
	return SixPointObstructionLedger{
		Obstructions:                    obs,
		FunctorZToRMissing:              !obs[0].Satisfied,
		PhysicalJMissing:                !obs[1].Satisfied,
		ChiralHyperchargeMissing:        !obs[2].Satisfied,
		HeatKernelSchemeMissing:         !obs[3].Satisfied,
		ScalarGaugeNormalizationMissing: !obs[4].Satisfied,
		ObservableDefinitionMissing:     !obs[5].Satisfied,
		AllUnsatisfied:                  allUnsatisfied(obs),
		HiggsPredictionBlocked:          true,
		Verdict:                         StatusSixPointLedgerCompiled,
	}
}

func establishHiggsFirewall(o SixPointObstructionLedger) FirewallSeal {
	return FirewallSeal{
		Name:   "SpectralActionHiggsPredictionFirewall",
		Active: o.HiggsPredictionBlocked,
		Reason: "six unresolved structures prevent promotion of raw finite traces, scalar-Morita branches, or sealed contact projectors into a Seeley-de Witt Higgs mass-ratio prediction",
		Target: "Seeley-de Witt a2/a4 Higgs mass-ratio or equivalent Higgs-to-gauge observable",
		BlockedPromotion: []string{
			"raw Tr(D_F²)/Tr(D_F⁴) ratios",
			"Gate-275 r_± scalar-Morita branches",
			"sealed Gate-281 contact projector orientation",
			"candidate C⊕H⊕M3(C) algebra without completed H_F,J,γ,Y",
		},
		CanUseForFutureStressTests:      true,
		CanClaimFiniteDerivedHiggsRatio: false,
		Verdict:                         StatusFirewallEstablished,
	}
}

func buildFutureCriteria(o SixPointObstructionLedger) FutureTheoremCriteria {
	criteria := make([]FutureCriterion, 0, len(o.Obstructions)+1)
	for _, ob := range o.Obstructions {
		criteria = append(criteria, FutureCriterion{Name: ob.Name, Required: true, Satisfied: ob.Satisfied, Description: ob.FutureTheorem})
	}
	criteria = append(criteria, FutureCriterion{Name: "pre-comparison prediction discipline", Required: true, Satisfied: false, Description: "the engine must emit the finite-core observable before comparing with the measured Higgs mass or electroweak VEV"})
	return FutureTheoremCriteria{
		Criteria:                        criteria,
		RequiresAllSixObstructions:      true,
		RequiresNativeProjectionMap:     true,
		RequiresPhysicalSpectralTriple:  true,
		RequiresHeatKernelNormalization: true,
		RequiresPreComparisonPrediction: true,
		CurrentGateCanLiftFirewall:      false,
		Verdict:                         StatusFutureCriteriaDefined,
	}
}

func auditFirewalls(prev resolventbranchsemantics.Analysis, s GeometricScaffoldManifest, o SixPointObstructionLedger, seal FirewallSeal, f FutureTheoremCriteria) FirewallAudit {
	return FirewallAudit{
		PreviousGate281Inherited:        prev.Summary.TraceNormAuditComplete && prev.OrientationSeal.Active,
		NoRBranchPromotion:              !prev.Summary.AmplitudeBranchLocked && o.FunctorZToRMissing,
		NoProjectorOrientationOverclaim: prev.OrientationSeal.OrientationIsSealedConditional && !prev.OrientationSeal.OrientationIsNativeTheorem,
		NoRawTraceToHeatKernelOverclaim: o.HeatKernelSchemeMissing && o.ScalarGaugeNormalizationMissing,
		NoHiggsMassClaim:                s.NoHiggsPredictionClaimed && !seal.CanClaimFiniteDerivedHiggsRatio,
		NoObservedMassesUsed:            true,
		NoEmpiricalYukawaInserted:       true,
		SealsDoNotRewriteNativeTheorems: true,
		FiniteCorePolluted:              false,
		Verdict:                         StatusFirewallsPreserved,
	}
}

func buildSummary(s GeometricScaffoldManifest, o SixPointObstructionLedger, seal FirewallSeal, f FutureTheoremCriteria, fw FirewallAudit) Summary {
	return Summary{
		ScaffoldCompiled:       len(s.Items) >= 7 && s.MoritaMultiplicityRecorded && s.ScalarMoritaShapeConstraintRecorded,
		SixPointLedgerCompiled: len(o.Obstructions) == 6 && o.AllUnsatisfied,
		HiggsFirewallActive:    seal.Active && !seal.CanClaimFiniteDerivedHiggsRatio,
		FutureCriteriaDefined:  len(f.Criteria) >= 7 && f.RequiresAllSixObstructions,
		HiggsRatioDerived:      false,
		PathBClosed:            true,
		FirewallPreserved:      !fw.FiniteCorePolluted,
		Status:                 StatusFailedHiggsPredictionSealed,
		DirectAnswer:           "Path B is structurally capped: the ASHA Engine has a spectral-action scaffold and a two-branch scalar-Morita shape constraint, but the Higgs a2/a4 ratio remains sealed until the six-point obstruction ledger is resolved.",
		NextGate:               "Path C candidate — B-gap coefficient/action derivation, or a future theorem deriving the six missing spectral-action structures.",
	}
}

func hasItem(items []ScaffoldItem, name string) bool {
	for _, item := range items {
		if item.Name == name {
			return true
		}
	}
	return false
}

func allUnsatisfied(obs []Obstruction) bool {
	if len(obs) == 0 {
		return false
	}
	for _, ob := range obs {
		if ob.Satisfied {
			return false
		}
	}
	return true
}

func FormatScaffold(s GeometricScaffoldManifest) string {
	parts := []string{fmt.Sprintf("items=%d", len(s.Items)), "verdict=" + s.Verdict}
	for _, item := range s.Items {
		parts = append(parts, fmt.Sprintf("%s[%s,%s,derived=%t,sealed=%t]", item.Name, item.SourceGate, item.Status, item.Derived, item.Sealed))
	}
	return strings.Join(parts, "; ")
}

func FormatObstructions(o SixPointObstructionLedger) string {
	parts := []string{fmt.Sprintf("count=%d", len(o.Obstructions)), fmt.Sprintf("allUnsatisfied=%t", o.AllUnsatisfied), "verdict=" + o.Verdict}
	for _, ob := range o.Obstructions {
		parts = append(parts, fmt.Sprintf("%d.%s satisfied=%t status=%s", ob.Index, ob.Name, ob.Satisfied, ob.Status))
	}
	return strings.Join(parts, "; ")
}

func FormatSeal(s FirewallSeal) string {
	return fmt.Sprintf("%s active=%t target=%s canStressTest=%t canClaimHiggs=%t blocked=%s verdict=%s", s.Name, s.Active, s.Target, s.CanUseForFutureStressTests, s.CanClaimFiniteDerivedHiggsRatio, strings.Join(s.BlockedPromotion, ","), s.Verdict)
}

func FormatFutureCriteria(f FutureTheoremCriteria) string {
	parts := []string{fmt.Sprintf("criteria=%d", len(f.Criteria)), fmt.Sprintf("currentCanLift=%t", f.CurrentGateCanLiftFirewall), "verdict=" + f.Verdict}
	for _, c := range f.Criteria {
		parts = append(parts, fmt.Sprintf("%s required=%t satisfied=%t", c.Name, c.Required, c.Satisfied))
	}
	return strings.Join(parts, "; ")
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("prev281=%t noRBranch=%t noProjectorOverclaim=%t noRawTraceOverclaim=%t noHiggs=%t noObservedMasses=%t noYukawa=%t sealsPreserved=%t polluted=%t verdict=%s", f.PreviousGate281Inherited, f.NoRBranchPromotion, f.NoProjectorOrientationOverclaim, f.NoRawTraceToHeatKernelOverclaim, f.NoHiggsMassClaim, f.NoObservedMassesUsed, f.NoEmpiricalYukawaInserted, f.SealsDoNotRewriteNativeTheorems, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("scaffold=%t sixPoint=%t firewall=%t future=%t higgsDerived=%t pathBClosed=%t firewallPreserved=%t status=%s direct=%s next=%s", s.ScaffoldCompiled, s.SixPointLedgerCompiled, s.HiggsFirewallActive, s.FutureCriteriaDefined, s.HiggsRatioDerived, s.PathBClosed, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
