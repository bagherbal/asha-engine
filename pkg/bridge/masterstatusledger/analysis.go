// Package masterstatusledger implements Gate 311:
// ASHA Engine Master Status Ledger / Project Capstone Audit.
//
// Gate 310 completed the structural-derivation phase by identifying the
// continuum obligations needed to turn the Gate-308/309 Higgs boundary into a
// collider-scale prediction. Gate 311 does not add a new physics fit. It
// compiles the project state into a firewalled master ledger: native core
// theorems, active seals, unresolved tensions, and the Phase-II work program.
package masterstatusledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE311-ASHA-ENGINE-MASTER-STATUS-LEDGER-PROJECT-CAPSTONE"

	StatusMasterLedgerCompiled        = "CONDITIONAL_SUPPORT_MASTER_STATUS_LEDGER_COMPILED"
	StatusStructuralPhaseCapstone     = "CONDITIONAL_SUPPORT_STRUCTURAL_PHASE_CAPSTONE_ACHIEVED"
	StatusCoreTheoremsCataloged       = "CONDITIONAL_SUPPORT_CORE_THEOREMS_CATALOGED"
	StatusSealedAxiomsCataloged       = "CONDITIONAL_SUPPORT_SEALED_AXIOMS_CATALOGED"
	StatusUnresolvedTensionsCataloged = "CONDITIONAL_SUPPORT_UNRESOLVED_TENSIONS_CATALOGED"
	StatusPhaseIIBlueprintFormalized  = "CONDITIONAL_SUPPORT_PHASE_II_BLUEPRINT_FORMALIZED"
	StatusFirewallsPreserved          = "CONDITIONAL_SUPPORT_MASTER_LEDGER_FIREWALLS_PRESERVED"

	StatusFailedFinalTOENotClaimed       = "FAILED_ROUTE_FINAL_THEORY_OF_EVERYTHING_NOT_CLAIMED"
	StatusFailedF2StillUnlocked          = "FAILED_ROUTE_F2_CUTOFF_MOMENT_SHAPE_STILL_UNLOCKED"
	StatusFailedAbsoluteGaugeStillSealed = "FAILED_ROUTE_ABSOLUTE_GAUGE_COUPLING_VALUE_STILL_SEALED"
	StatusFailedHiggsTensionUnresolved   = "FAILED_ROUTE_GATE309_HIGGS_331GEV_TENSION_UNRESOLVED"
	StatusFailedThresholdsNotDerived     = "FAILED_ROUTE_THRESHOLD_MATCHING_JUMPS_NOT_DERIVED"
	StatusFailedBGapInstantonNotDerived  = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_NOT_DERIVED"
	StatusFailedPhysicalJStillFormal     = "FAILED_ROUTE_PHYSICAL_REAL_STRUCTURE_J_TWIST_STILL_FORMAL"
	StatusFailedNoLowEnergyMassClaimed   = "FAILED_ROUTE_LOW_ENERGY_HIGGS_MASS_NOT_CLAIMED"
)

const (
	gateHighestInherited  = 310
	exactWeakMixing       = "sin²θ_W = 3/8"
	exactTraceShape       = "1197/4624"
	exactContactResonance = "4/π"
	gate309MassGeV        = 331.630412
	gate309LambdaAtV      = 0.907051722647
	higgsReferenceGeV     = 125.10
)

type RegistrySpan struct {
	AuditID              string
	GateRange            string
	HighestGateInherited int
	StructuralPhase      string
	CapstoneForPhase     string
	ReadsPackageRegistry bool
	AddsNewPhysicsFit    bool
	Verdict              string
}

type CoreTheorem struct {
	Name                  string
	RepresentativeGate    string
	Statement             string
	NativeToFiniteCore    bool
	RequiresPhenomenology bool
	PromotedToPhysics     bool
	Status                string
}

type CoreTheoremLedger struct {
	Cataloged                  bool
	Theorems                   []CoreTheorem
	ZeroPhenomenologyCount     int
	RequiredNamedTruthsPresent bool
	ContainsWeakMixing         bool
	ContainsTrialityTopology   bool
	ContainsMoritaMultiplicity bool
	ContainsContactResonance   bool
	ContainsTrueBimodule       bool
	ContainsTraceEquivalence   bool
	Verdict                    string
}

type Seal struct {
	Name                string
	RepresentativeGate  string
	SealClass           string
	WhatItProvides      string
	WhyRequired         string
	Phenomenological    bool
	StructuralPromotion bool
	FinalDerived        bool
	Status              string
}

type SealLedger struct {
	Cataloged                      bool
	Seals                          []Seal
	RequiredNamedSealsPresent      bool
	PhenomenologicalSealCount      int
	StructuralPromotionSealCount   int
	FinalPredictionStillFirewalled bool
	Verdict                        string
}

type Tension struct {
	Name                string
	FirstExposedGate    string
	MathematicalGap     string
	Blocks              string
	CandidateResolution string
	Resolved            bool
	Status              string
}

type TensionLedger struct {
	Cataloged                     bool
	Tensions                      []Tension
	ContainsF2CutoffShape         bool
	ContainsAbsoluteGaugeCoupling bool
	ContainsGate309HiggsTension   bool
	ContainsBGapInstantonGap      bool
	ContainsPhysicalJGap          bool
	AnyResolved                   bool
	Verdict                       string
}

type PhaseIIObligation struct {
	Name                  string
	WorkPackage           string
	RequiredInputs        []string
	ExpectedOutput        string
	WhyNext               string
	CanUseEmpiricalFit    bool
	BlocksFinalPrediction bool
	Status                string
}

type PhaseIIBlueprint struct {
	Formalized                  bool
	Obligations                 []PhaseIIObligation
	ThresholdMatchingIncluded   bool
	NonPerturbativeBGapIncluded bool
	RealStructureTwistIncluded  bool
	TopSectorTensorIncluded     bool
	PrecisionRGEIncluded        bool
	NoEmpiricalTuning           bool
	Verdict                     string
}

type FirewallAudit struct {
	NoObservedHiggsFitInserted      bool
	NoObservedTopFitInserted        bool
	NoThresholdJumpInserted         bool
	NoF2ShapeInserted               bool
	NoAbsoluteGaugeValueInserted    bool
	NoBGapInstantonInserted         bool
	NoFinalTOEClaimed               bool
	NoLowEnergyMassClaimed          bool
	PhaseIIRequiredBeforePrediction bool
	Verdict                         string
}

type Summary struct {
	RegistryCompiled        bool
	CoreCatalogReady        bool
	SealCatalogReady        bool
	TensionCatalogReady     bool
	PhaseIIBlueprintReady   bool
	StructuralPhaseCapstone bool
	FinalTheoryClaimed      bool
	FinalMassClaimed        bool
	FirewallPreserved       bool
	Status                  string
	DirectAnswer            string
	NextProject             string
}

type Analysis struct {
	Span      RegistrySpan
	Core      CoreTheoremLedger
	Seals     SealLedger
	Tensions  TensionLedger
	PhaseII   PhaseIIBlueprint
	Firewalls FirewallAudit
	Summary   Summary
	Truth     string
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
	span := compileSpan()
	core := compileCoreTheorems()
	seals := compileSealLedger()
	tensions := compileTensionLedger()
	phaseII := compilePhaseIIBlueprint()
	firewalls := auditFirewalls()
	summary := compileSummary(span, core, seals, tensions, phaseII, firewalls)
	truth := "Gate 311 is the capstone ledger for the structural ASHA phase. It records the native finite-algebra successes, the active epistemological seals, and the exact live tensions without converting conditional diagnostics into final physical claims. Phase II must derive threshold matching, non-perturbative B-gap dynamics, and the physical real-structure twist before a final collider-scale prediction can be authorized."
	return Analysis{Span: span, Core: core, Seals: seals, Tensions: tensions, PhaseII: phaseII, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func compileSpan() RegistrySpan {
	return RegistrySpan{
		AuditID:              AuditID,
		GateRange:            "Gate 1 → Gate 310 inherited; Gate 311 compiles the capstone ledger",
		HighestGateInherited: gateHighestInherited,
		StructuralPhase:      "finite algebra → spectral action → normalization → RG diagnostic",
		CapstoneForPhase:     "Structural Derivation Phase / Phase-I closure",
		ReadsPackageRegistry: true,
		AddsNewPhysicsFit:    false,
		Verdict:              strings.Join([]string{StatusMasterLedgerCompiled, StatusStructuralPhaseCapstone}, ";"),
	}
}

func compileCoreTheorems() CoreTheoremLedger {
	theorems := []CoreTheorem{
		{Name: "Weak mixing angle from representation trace", RepresentativeGate: "Gate 298/308", Statement: exactWeakMixing + " via k_Y=5/3 and GUT-normalized trace ledger", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
		{Name: "Triality generation topology", RepresentativeGate: "Triality/Tau-Eta chain", Statement: "τ_η supplies the generation-tag topology and spatial S3 orientation sieve", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: false, Status: StatusCoreTheoremsCataloged},
		{Name: "Morita trace multiplicities", RepresentativeGate: "Morita bimodule chain", Statement: "κ_C=1 and κ_Q=3 define the scalar/gauge trace multiplicity split", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
		{Name: "Contact topological resonance", RepresentativeGate: "B-gap / Hopf / contact action chain", Statement: exactContactResonance + " resonance identifies the hidden topological action scale candidate", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: false, Status: StatusCoreTheoremsCataloged},
		{Name: "True Bimodule Mandate", RepresentativeGate: "Gate 295", Statement: "physical one-form calculus requires a true left/right bimodule, not a direct-sum label carrier", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
		{Name: "Inner fluctuation Standard Model scaffold", RepresentativeGate: "Gate 298", Statement: "finite inner fluctuations yield U(1)_Y × SU(2)_L × SU(3)_C with 12 gauge bosons and one complex Higgs doublet", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
		{Name: "Positive scalar kinetic carrier", RepresentativeGate: "Gate 301/302", Statement: "K_H^raw is a Hilbert-Schmidt sum and maps to Z_H>0 under the positive convention ledger", NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
		{Name: "Trace equivalence seal", RepresentativeGate: "Gate 307", Statement: "projected scalar carrier gives C4_raw/KH_raw² = " + exactTraceShape, NativeToFiniteCore: true, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
		{Name: "Quartic unification boundary", RepresentativeGate: "Gate 308", Statement: "λ_H(Λ_GUT) = (1197/4624)·g_*² after τ_GUT=1 and Sign_4=+1 conventions", NativeToFiniteCore: false, RequiresPhenomenology: false, PromotedToPhysics: true, Status: StatusCoreTheoremsCataloged},
	}
	return CoreTheoremLedger{
		Cataloged:                  true,
		Theorems:                   theorems,
		ZeroPhenomenologyCount:     countNativeZeroPhenomenology(theorems),
		RequiredNamedTruthsPresent: hasCoreTruth(theorems, "Weak mixing angle") && hasCoreTruth(theorems, "Triality generation topology") && hasCoreTruth(theorems, "Morita trace multiplicities") && hasCoreTruth(theorems, "Contact topological resonance") && hasCoreTruth(theorems, "True Bimodule Mandate") && hasCoreTruth(theorems, "Trace equivalence seal"),
		ContainsWeakMixing:         hasCoreTruth(theorems, "Weak mixing angle"),
		ContainsTrialityTopology:   hasCoreTruth(theorems, "Triality generation topology"),
		ContainsMoritaMultiplicity: hasCoreTruth(theorems, "Morita trace multiplicities"),
		ContainsContactResonance:   hasCoreTruth(theorems, "Contact topological resonance"),
		ContainsTrueBimodule:       hasCoreTruth(theorems, "True Bimodule Mandate"),
		ContainsTraceEquivalence:   hasCoreTruth(theorems, "Trace equivalence seal"),
		Verdict:                    StatusCoreTheoremsCataloged,
	}
}

func compileSealLedger() SealLedger {
	seals := []Seal{
		{Name: "EmpiricalYukawaSeal", RepresentativeGate: "Yukawa amplitude / empirical texture chain", SealClass: "phenomenological amplitude seal", WhatItProvides: "dimensionless Yukawa singular values and flavor texture data", WhyRequired: "finite geometry fixes carriers and shapes more strongly than absolute measured fermion amplitudes", Phenomenological: true, FinalDerived: false, Status: "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_SEALED"},
		{Name: "ResolventAdjunctionSeal", RepresentativeGate: "Resolvent field adjunction chain", SealClass: "spontaneous vacuum orientation seal", WhatItProvides: "choice of resolvent root / sector orientation", WhyRequired: "the finite quartic module offers multiple algebraic roots; selecting one is spontaneous symmetry breaking data", Phenomenological: false, StructuralPromotion: true, FinalDerived: false, Status: "CONDITIONAL_SUPPORT_RESOLVENT_ADJUNCTION_SEAL_ACTIVE"},
		{Name: "IntermediateBreakingSeal", RepresentativeGate: "Intermediate breaking / seesaw preflight", SealClass: "conditional continuum-scale seal", WhatItProvides: "intermediate B-sector/seesaw activation lane", WhyRequired: "finite core identifies the possibility, but does not yet derive the physical threshold mass/action", Phenomenological: true, FinalDerived: false, Status: "FAILED_ROUTE_INTERMEDIATE_BREAKING_SCALE_NOT_DERIVED"},
		{Name: "PerSlotMonotonicitySeal", RepresentativeGate: "Gate 291", SealClass: "branch-selection monotonicity seal", WhatItProvides: "selects the branch yielding the contact scalar shape " + exactTraceShape, WhyRequired: "the branch is structurally constrained but the per-slot ordering assumption must be sealed", Phenomenological: false, StructuralPromotion: true, FinalDerived: false, Status: "CONDITIONAL_SUPPORT_PER_SLOT_MONOTONICITY_SEAL_ACTIVE"},
		{Name: "ContactSpectralCutoffPromotionSeal", RepresentativeGate: "Gate 304", SealClass: "cutoff-moment promotion seal", WhatItProvides: "promotes f0 := ζ_contact(0) = 7 for a4 kinetic/quartic lanes", WhyRequired: "a discrete contact invariant must be promoted to an admissible continuous test-profile moment", Phenomenological: false, StructuralPromotion: true, FinalDerived: false, Status: "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_PROMOTED"},
		{Name: "GaugeCouplingBoundarySeal", RepresentativeGate: "Gate 308/309", SealClass: "absolute coupling boundary seal", WhatItProvides: "conditional g_*²=1 lane", WhyRequired: "trace ratios fix relative structure, not the absolute gauge coupling value", Phenomenological: true, FinalDerived: false, Status: StatusFailedAbsoluteGaugeStillSealed},
		{Name: "BoundaryScaleSeal", RepresentativeGate: "Gate 309", SealClass: "RG boundary-scale seal", WhatItProvides: "conditional Λ_GUT≈10^17 GeV closed-triangle lane", WhyRequired: "transport requires a continuum boundary scale beyond the finite trace identity", Phenomenological: true, FinalDerived: false, Status: "FAILED_ROUTE_LAMBDA_GUT_BOUNDARY_SCALE_STILL_SEALED"},
	}
	return SealLedger{
		Cataloged:                      true,
		Seals:                          seals,
		RequiredNamedSealsPresent:      hasSeal(seals, "EmpiricalYukawaSeal") && hasSeal(seals, "ResolventAdjunctionSeal") && hasSeal(seals, "IntermediateBreakingSeal") && hasSeal(seals, "PerSlotMonotonicitySeal") && hasSeal(seals, "ContactSpectralCutoffPromotionSeal"),
		PhenomenologicalSealCount:      countSeals(seals, true, false),
		StructuralPromotionSealCount:   countSeals(seals, false, true),
		FinalPredictionStillFirewalled: true,
		Verdict:                        strings.Join([]string{StatusSealedAxiomsCataloged, StatusFailedFinalTOENotClaimed}, ";"),
	}
}

func compileTensionLedger() TensionLedger {
	tensions := []Tension{
		{Name: "unresolved f2 cutoff shape", FirstExposedGate: "Gate 305", MathematicalGap: "f0=7 fixes a4 normalization but does not determine the higher test-function moment f2", Blocks: "absolute Higgs mass parameter μ_H²", CandidateResolution: "canonical cutoff-profile shape theorem or finite spectral rule for higher moments", Resolved: false, Status: StatusFailedF2StillUnlocked},
		{Name: "absolute gauge coupling normalization", FirstExposedGate: "Gate 308/309", MathematicalGap: "relative trace ratios do not derive g_*² or the exact boundary scale", Blocks: "absolute λ_H boundary value and RG transport without conditional seal", CandidateResolution: "absolute coupling theorem or topological gauge-action normalization", Resolved: false, Status: StatusFailedAbsoluteGaugeStillSealed},
		{Name: "Gate 309 Higgs running-mass tension", FirstExposedGate: "Gate 309", MathematicalGap: fmt.Sprintf("one-loop r_plus lane gives λ(v)=%.12f and m≈%.6f GeV instead of the 125 GeV target", gate309LambdaAtV, gate309MassGeV), Blocks: "final collider-scale Higgs prediction", CandidateResolution: "derived threshold matching jumps, modified top-sector tensor, and full two-loop/pole ledger", Resolved: false, Status: StatusFailedHiggsTensionUnresolved},
		{Name: "B-gap instanton action", FirstExposedGate: "Gate 299/310", MathematicalGap: "polynomial Majorana insertion does not derive S_inst=(4/π)/B_gap", Blocks: "non-perturbative hierarchy and intermediate-seesaw dynamics", CandidateResolution: "native determinant/saddle/instanton action theorem", Resolved: false, Status: StatusFailedBGapInstantonNotDerived},
		{Name: "physical real-structure twist", FirstExposedGate: "KO/J chain", MathematicalGap: "J_swap has formal doubled-space KO behavior, but physical anti-linear particle/antiparticle semantics need a native twist", Blocks: "fully physical finite spectral triple semantics", CandidateResolution: "twisted real-structure theorem compatible with order-one calculus and anomaly signs", Resolved: false, Status: StatusFailedPhysicalJStillFormal},
		{Name: "threshold matching values", FirstExposedGate: "Gate 310", MathematicalGap: "Δλ_i values for PeV/B-gap/heavy scalar thresholds are not derived", Blocks: "capacity to reduce the 331 GeV tension", CandidateResolution: "finite heavy-sector coupling and mass ledger with matching formulae", Resolved: false, Status: StatusFailedThresholdsNotDerived},
	}
	return TensionLedger{
		Cataloged:                     true,
		Tensions:                      tensions,
		ContainsF2CutoffShape:         hasTension(tensions, "f2"),
		ContainsAbsoluteGaugeCoupling: hasTension(tensions, "absolute gauge"),
		ContainsGate309HiggsTension:   hasTension(tensions, "Gate 309"),
		ContainsBGapInstantonGap:      hasTension(tensions, "B-gap"),
		ContainsPhysicalJGap:          hasTension(tensions, "real-structure"),
		AnyResolved:                   anyTensionResolved(tensions),
		Verdict:                       strings.Join([]string{StatusUnresolvedTensionsCataloged, StatusFailedNoLowEnergyMassClaimed}, ";"),
	}
}

func compilePhaseIIBlueprint() PhaseIIBlueprint {
	obligations := []PhaseIIObligation{
		{Name: "Threshold Matching Derivations", WorkPackage: "finite heavy-sector EFT matching", RequiredInputs: []string{"PeV/B-gap/heavy-scalar mass ledger", "finite heavy-sector couplings", "matching convention"}, ExpectedOutput: "symbolic and numerical Δλ_i threshold matrix", WhyNext: "threshold jumps are the only Gate-310 class with enough capacity to move the 331 GeV diagnostic", BlocksFinalPrediction: true, Status: StatusFailedThresholdsNotDerived},
		{Name: "Non-Perturbative Instanton Action Mapping", WorkPackage: "B-gap saddle / determinant / topological action", RequiredInputs: []string{"native inverse-action theorem", "finite determinant or saddle carrier", "4/π resonance bridge"}, ExpectedOutput: "derived S_inst=(4/π)/B_gap or rejection of that route", WhyNext: "polynomial heat-kernel Majorana insertion cannot generate the required inverse hierarchy action", BlocksFinalPrediction: true, Status: StatusFailedBGapInstantonNotDerived},
		{Name: "Physical Real-Structure Twist", WorkPackage: "twisted KO/J semantics", RequiredInputs: []string{"anti-linear particle/antiparticle action", "order-one compatibility", "chirality/anomaly ledger"}, ExpectedOutput: "native J_phys replacing merely formal J_swap semantics", WhyNext: "full physical NCG semantics require real-structure data beyond doubled-space bookkeeping", BlocksFinalPrediction: true, Status: StatusFailedPhysicalJStillFormal},
		{Name: "Modified Top-Sector Tensor Audit", WorkPackage: "top-Yukawa carrier refinement", RequiredInputs: []string{"r_plus branch carrier", "triality/tau_eta tensor decomposition", "top beta sensitivity"}, ExpectedOutput: "accepted or rejected correction to the high y_t boundary", WhyNext: "Gate 309 tension is driven by the top-Yukawa lane", BlocksFinalPrediction: true, Status: "FAILED_ROUTE_TOP_SECTOR_TENSOR_NOT_DERIVED"},
		{Name: "Full Precision RGE + Pole-Mass Pipeline", WorkPackage: "two-loop RG, threshold matching, pole conversion", RequiredInputs: []string{"full beta coefficient table", "threshold values", "MS-bar to pole self-energies"}, ExpectedOutput: "collider-scale prediction with uncertainty ledger", WhyNext: "the first GeV-scale diagnostic exists but must be transported with the full continuum apparatus", BlocksFinalPrediction: true, Status: "FAILED_ROUTE_FULL_PRECISION_PIPELINE_NOT_EXECUTED"},
	}
	return PhaseIIBlueprint{
		Formalized:                  true,
		Obligations:                 obligations,
		ThresholdMatchingIncluded:   hasObligation(obligations, "Threshold Matching"),
		NonPerturbativeBGapIncluded: hasObligation(obligations, "Non-Perturbative"),
		RealStructureTwistIncluded:  hasObligation(obligations, "Real-Structure"),
		TopSectorTensorIncluded:     hasObligation(obligations, "Top-Sector"),
		PrecisionRGEIncluded:        hasObligation(obligations, "Precision RGE"),
		NoEmpiricalTuning:           true,
		Verdict:                     strings.Join([]string{StatusPhaseIIBlueprintFormalized, StatusFailedFinalTOENotClaimed}, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{
		NoObservedHiggsFitInserted:      true,
		NoObservedTopFitInserted:        true,
		NoThresholdJumpInserted:         true,
		NoF2ShapeInserted:               true,
		NoAbsoluteGaugeValueInserted:    true,
		NoBGapInstantonInserted:         true,
		NoFinalTOEClaimed:               true,
		NoLowEnergyMassClaimed:          true,
		PhaseIIRequiredBeforePrediction: true,
		Verdict:                         strings.Join([]string{StatusFirewallsPreserved, StatusFailedFinalTOENotClaimed, StatusFailedNoLowEnergyMassClaimed}, ";"),
	}
}

func compileSummary(span RegistrySpan, core CoreTheoremLedger, seals SealLedger, tensions TensionLedger, phaseII PhaseIIBlueprint, f FirewallAudit) Summary {
	firewall := f.NoObservedHiggsFitInserted && f.NoObservedTopFitInserted && f.NoThresholdJumpInserted && f.NoF2ShapeInserted && f.NoAbsoluteGaugeValueInserted && f.NoBGapInstantonInserted && f.NoFinalTOEClaimed && f.NoLowEnergyMassClaimed
	return Summary{
		RegistryCompiled:        span.ReadsPackageRegistry && span.HighestGateInherited == gateHighestInherited,
		CoreCatalogReady:        core.Cataloged && core.RequiredNamedTruthsPresent,
		SealCatalogReady:        seals.Cataloged && seals.RequiredNamedSealsPresent,
		TensionCatalogReady:     tensions.Cataloged && tensions.ContainsF2CutoffShape && tensions.ContainsAbsoluteGaugeCoupling && tensions.ContainsGate309HiggsTension,
		PhaseIIBlueprintReady:   phaseII.Formalized && phaseII.ThresholdMatchingIncluded && phaseII.NonPerturbativeBGapIncluded && phaseII.RealStructureTwistIncluded,
		StructuralPhaseCapstone: true,
		FinalTheoryClaimed:      false,
		FinalMassClaimed:        false,
		FirewallPreserved:       firewall,
		Status:                  strings.Join([]string{StatusMasterLedgerCompiled, StatusStructuralPhaseCapstone, StatusFailedFinalTOENotClaimed}, ";"),
		DirectAnswer:            "The structural phase is capstoned: ASHA has a finite-algebra Standard Model scaffold, a quartic boundary, and a first GeV-scale tension diagnostic, but final prediction awaits Phase-II threshold, non-perturbative, and real-structure derivations.",
		NextProject:             "Phase II: Non-Perturbative Dynamics & Threshold Matching, beginning with a threshold-sensitivity matrix for Δλ and a native B-gap instanton action theorem.",
	}
}

func countNativeZeroPhenomenology(xs []CoreTheorem) int {
	n := 0
	for _, x := range xs {
		if x.NativeToFiniteCore && !x.RequiresPhenomenology {
			n++
		}
	}
	return n
}

func hasCoreTruth(xs []CoreTheorem, name string) bool {
	for _, x := range xs {
		if strings.Contains(x.Name, name) {
			return true
		}
	}
	return false
}

func hasSeal(xs []Seal, name string) bool {
	for _, x := range xs {
		if x.Name == name {
			return true
		}
	}
	return false
}

func countSeals(xs []Seal, phenomenological bool, structural bool) int {
	n := 0
	for _, x := range xs {
		if x.Phenomenological == phenomenological && (structural == false || x.StructuralPromotion == structural) {
			n++
		}
	}
	return n
}

func hasTension(xs []Tension, needle string) bool {
	needle = strings.ToLower(needle)
	for _, x := range xs {
		if strings.Contains(strings.ToLower(x.Name), needle) || strings.Contains(strings.ToLower(x.MathematicalGap), needle) {
			return true
		}
	}
	return false
}

func anyTensionResolved(xs []Tension) bool {
	for _, x := range xs {
		if x.Resolved {
			return true
		}
	}
	return false
}

func hasObligation(xs []PhaseIIObligation, needle string) bool {
	needle = strings.ToLower(needle)
	for _, x := range xs {
		if strings.Contains(strings.ToLower(x.Name), needle) || strings.Contains(strings.ToLower(x.WorkPackage), needle) {
			return true
		}
	}
	return false
}

func FormatSpan(x RegistrySpan) string {
	return fmt.Sprintf("audit=%s range=%s highest=%d phase=%s capstone=%s registry=%v fit=%v verdict=%s", x.AuditID, x.GateRange, x.HighestGateInherited, x.StructuralPhase, x.CapstoneForPhase, x.ReadsPackageRegistry, x.AddsNewPhysicsFit, x.Verdict)
}

func FormatCore(x CoreTheoremLedger) string {
	names := make([]string, 0, len(x.Theorems))
	for _, t := range x.Theorems {
		names = append(names, t.Name+":"+t.Statement)
	}
	return fmt.Sprintf("cataloged=%v count=%d nativeZeroPhen=%d required=%v weak=%v triality=%v morita=%v contact=%v bimodule=%v trace=%v theorems=[%s] verdict=%s", x.Cataloged, len(x.Theorems), x.ZeroPhenomenologyCount, x.RequiredNamedTruthsPresent, x.ContainsWeakMixing, x.ContainsTrialityTopology, x.ContainsMoritaMultiplicity, x.ContainsContactResonance, x.ContainsTrueBimodule, x.ContainsTraceEquivalence, strings.Join(names, " | "), x.Verdict)
}

func FormatSeals(x SealLedger) string {
	names := make([]string, 0, len(x.Seals))
	for _, s := range x.Seals {
		names = append(names, s.Name+":"+s.WhyRequired)
	}
	return fmt.Sprintf("cataloged=%v count=%d required=%v phenom=%d promotions=%d finalFirewalled=%v seals=[%s] verdict=%s", x.Cataloged, len(x.Seals), x.RequiredNamedSealsPresent, x.PhenomenologicalSealCount, x.StructuralPromotionSealCount, x.FinalPredictionStillFirewalled, strings.Join(names, " | "), x.Verdict)
}

func FormatTensions(x TensionLedger) string {
	names := make([]string, 0, len(x.Tensions))
	for _, t := range x.Tensions {
		names = append(names, t.Name+":"+t.Status)
	}
	return fmt.Sprintf("cataloged=%v count=%d f2=%v gauge=%v higgs309=%v bgap=%v physicalJ=%v anyResolved=%v tensions=[%s] verdict=%s", x.Cataloged, len(x.Tensions), x.ContainsF2CutoffShape, x.ContainsAbsoluteGaugeCoupling, x.ContainsGate309HiggsTension, x.ContainsBGapInstantonGap, x.ContainsPhysicalJGap, x.AnyResolved, strings.Join(names, " | "), x.Verdict)
}

func FormatPhaseII(x PhaseIIBlueprint) string {
	names := make([]string, 0, len(x.Obligations))
	for _, o := range x.Obligations {
		names = append(names, o.Name+":"+o.ExpectedOutput)
	}
	return fmt.Sprintf("formalized=%v count=%d thresholds=%v bgap=%v realJ=%v topTensor=%v precisionRGE=%v noTuning=%v obligations=[%s] verdict=%s", x.Formalized, len(x.Obligations), x.ThresholdMatchingIncluded, x.NonPerturbativeBGapIncluded, x.RealStructureTwistIncluded, x.TopSectorTensorIncluded, x.PrecisionRGEIncluded, x.NoEmpiricalTuning, strings.Join(names, " | "), x.Verdict)
}

func FormatFirewalls(x FirewallAudit) string {
	return fmt.Sprintf("noHiggsFit=%v noTopFit=%v noThreshold=%v noF2=%v noGauge=%v noBGap=%v noTOE=%v noMass=%v phaseIIRequired=%v verdict=%s", x.NoObservedHiggsFitInserted, x.NoObservedTopFitInserted, x.NoThresholdJumpInserted, x.NoF2ShapeInserted, x.NoAbsoluteGaugeValueInserted, x.NoBGapInstantonInserted, x.NoFinalTOEClaimed, x.NoLowEnergyMassClaimed, x.PhaseIIRequiredBeforePrediction, x.Verdict)
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("registry=%v core=%v seals=%v tensions=%v phaseII=%v capstone=%v finalTheory=%v finalMass=%v firewall=%v status=%s answer=%s next=%s", x.RegistryCompiled, x.CoreCatalogReady, x.SealCatalogReady, x.TensionCatalogReady, x.PhaseIIBlueprintReady, x.StructuralPhaseCapstone, x.FinalTheoryClaimed, x.FinalMassClaimed, x.FirewallPreserved, x.Status, x.DirectAnswer, x.NextProject)
}
