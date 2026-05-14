// Package fullflavorledgerclosure implements Gate 267:
// Full Flavor Ledger Closure / Quark-Lepton Empirical Firewall Summary Audit.
//
// Gate 267 is a manifest gate. It does not attempt a new flavor derivation.
// It consolidates Gates 261-266 into an auditable boundary ledger: what the
// finite Cℓ(1,7) core has supplied as kinematic structure, what has been
// reconstructed algebraically from sealed empirical textures, and what remains
// unavailable until a future finite spectral/action theorem exists.
package fullflavorledgerclosure

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/empiricalflavorledger"
)

const (
	AuditID = "GATE267-FULL-FLAVOR-LEDGER-CLOSURE-QUARK-LEPTON-EMPIRICAL-FIREWALL-SUMMARY-AUDIT"

	StatusGate266Inherited             = "CONDITIONAL_SUPPORT_GATE266_QUARK_LEPTON_RECONSTRUCTION_INHERITED"
	StatusGeometricLedgerCompiled      = "CONDITIONAL_SUPPORT_GEOMETRIC_FLAVOR_DERIVATION_LEDGER_COMPILED"
	StatusEmpiricalLedgerCompiled      = "CONDITIONAL_SUPPORT_EMPIRICAL_FLAVOR_INPUT_LEDGER_COMPILED"
	StatusReconstructionManifestClosed = "CONDITIONAL_SUPPORT_QUARK_LEPTON_RECONSTRUCTION_MANIFEST_CLOSED"
	StatusFutureTheoremCriteriaDefined = "CONDITIONAL_SUPPORT_FUTURE_FLAVOR_SEAL_LIFT_CRITERIA_DEFINED"
	StatusFullFlavorLedgerClosed       = "CONDITIONAL_SUPPORT_FULL_FLAVOR_LEDGER_CLOSED_AND_SEALED"
	StatusNoNativeFlavorAmplitude      = "FAILED_ROUTE_NO_NATIVE_FLAVOR_AMPLITUDE_DERIVATION"
	StatusNoNativeCKMPMNS              = "FAILED_ROUTE_CKM_PMNS_NUMERICS_REMAIN_EMPIRICAL"
	StatusNoNativeFermionMasses        = "FAILED_ROUTE_FERMION_MASSES_REMAIN_EMPIRICAL"
	StatusSpectralActionMissing        = "FAILED_ROUTE_FINITE_SPECTRAL_ACTION_FOR_YUKAWA_AMPLITUDES_MISSING"
	StatusMajoranaNatureNotDerived     = "FAILED_ROUTE_MAJORANA_OR_DIRAC_NEUTRINO_NATURE_NOT_FINITE_DERIVED"
	StatusSpontaneousOrientationSealed = "FAILED_ROUTE_SPONTANEOUS_FLAVOR_ORIENTATION_REMAINS_SEALED"
)

type Gate266Inheritance struct {
	EmpiricalYukawaSealActive     bool
	QuarkSVDCKMReconstructed      bool
	ChargedLeptonSVDReconstructed bool
	NeutrinoTakagiReconstructed   bool
	PMNSReconstructed             bool
	LargeAngleLeptonStructure     bool
	QuarkNativeDerivation         bool
	LeptonNativeDerivation        bool
	EmpiricalBoundaryPreserved    bool
	MajoranaNatureFiniteDerived   bool
	RepresentativeDataOnly        bool
	Verdict                       string
}

type LedgerItem struct {
	Name       string
	SourceGate string
	Kind       string
	Status     string
	FiniteCore bool
	Sealed     bool
	Statement  string
}

type GeometricDerivationLedger struct {
	Items                          []LedgerItem
	SCarrierSpaceRecorded          bool
	FiniteAlgebraRecorded          bool
	GaugeMatterChargeRecorded      bool
	ThreeGenerationCapacity        bool
	TauEtaSourceMapRecorded        bool
	AdTauMixingComplementRecorded  bool
	TrialityHermitianBasisRecorded bool
	YukawaAmplitudeDerived         bool
	CKMPMNSDerived                 bool
	FermionMassesDerived           bool
	Verdict                        string
}

type EmpiricalInputLedger struct {
	Items                        []LedgerItem
	SpontaneousCarrierSealActive bool
	EmpiricalYukawaSealActive    bool
	WeakFrameOrientationSealed   bool
	ScalarVEVAlignmentSealed     bool
	QuarkTexturesSealed          bool
	LeptonTexturesSealed         bool
	CKMEntriesSealed             bool
	PMNSEntriesSealed            bool
	MajoranaChoiceSealed         bool
	DoesNotRewriteFiniteCore     bool
	Verdict                      string
}

type ReconstructionVerification struct {
	QuarkSVDCKMVerified             bool
	ChargedLeptonSVDVerified        bool
	MajoranaTakagiVerified          bool
	PMNSVerified                    bool
	SVDIsAlgebraicReconstruction    bool
	TakagiIsAlgebraicReconstruction bool
	ObservablePipelineWorksOnData   bool
	ObservablePipelinePredictsData  bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Reason    string
}

type FutureTheoremCriteria struct {
	Criteria                        []FutureCriterion
	RequiresFiniteSpectralAction    bool
	RequiresCanonicalFiniteDirac    bool
	RequiresHeatKernelCoefficients  bool
	RequiresYukawaAmplitudeMap      bool
	RequiresHopfOrBGapProjection    bool
	RequiresNormalizationScheme     bool
	RequiresMassAndMixingPrediction bool
	CurrentGateCanLiftSeal          bool
	RecommendedNextGate             string
	Verdict                         string
}

type FirewallManifest struct {
	KinematicsDerived               bool
	DynamicsSealed                  bool
	SpontaneousCarrierSealPreserved bool
	EmpiricalYukawaSealPreserved    bool
	NoMassPredictionClaim           bool
	NoCKMPMNSPredictionClaim        bool
	NoMajoranaNatureClaim           bool
	NoSpectralActionClaim           bool
	ClosureDoesNotAddNewPhysics     bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	Gate266Inherited            bool
	GeometricLedgerClosed       bool
	EmpiricalLedgerClosed       bool
	ReconstructionsVerified     bool
	FutureCriteriaDefined       bool
	FullFlavorLedgerClosed      bool
	NativeFlavorDynamicsDerived bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	PreviousGate266 empiricalflavorledger.Analysis
	Inheritance     Gate266Inheritance
	Geometric       GeometricDerivationLedger
	Empirical       EmpiricalInputLedger
	Reconstruction  ReconstructionVerification
	FutureCriteria  FutureTheoremCriteria
	Firewall        FirewallManifest
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := empiricalflavorledger.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 266 predecessor: %w", err)
			return
		}
		inh := inheritGate266(prev)
		geo := compileGeometricLedger(inh)
		emp := compileEmpiricalLedger(inh)
		recon := verifyReconstructions(inh)
		criteria := defineFutureCriteria()
		firewall := buildFirewall(geo, emp, recon, criteria)
		summary := summarize(inh, geo, emp, recon, criteria, firewall)
		truth := buildTruth(inh, geo, emp, recon, criteria, firewall)
		defaultA = Analysis{PreviousGate266: prev, Inheritance: inh, Geometric: geo, Empirical: emp, Reconstruction: recon, FutureCriteria: criteria, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate266(prev empiricalflavorledger.Analysis) Gate266Inheritance {
	return Gate266Inheritance{
		EmpiricalYukawaSealActive:     prev.Seal.Activated && prev.Seal.ExplicitlyQuarantined,
		QuarkSVDCKMReconstructed:      prev.Inheritance.QuarkSVDCKMVerified,
		ChargedLeptonSVDReconstructed: prev.Summary.ChargedLeptonSVDCompleted && prev.ChargedSVD.Passed,
		NeutrinoTakagiReconstructed:   prev.Summary.NeutrinoTakagiCompleted && prev.NeutrinoTakagi.Passed,
		PMNSReconstructed:             prev.Summary.PMNSReconstructed && prev.PMNS.Verified,
		LargeAngleLeptonStructure:     prev.Summary.LargeAnglesAudited && prev.LargeAngles.LargeAngleStructure,
		QuarkNativeDerivation:         prev.Inheritance.QuarkNativeDerivation,
		LeptonNativeDerivation:        prev.Summary.NativeDerivation,
		EmpiricalBoundaryPreserved:    prev.Summary.EmpiricalBoundaryPreserved && !prev.Firewall.FiniteCorePolluted,
		MajoranaNatureFiniteDerived:   prev.NeutrinoTakagi.DerivedNeutrinoNature,
		RepresentativeDataOnly:        prev.Data.RepresentativeNotPrecision || prev.Inheritance.RepresentativeQuarkDataWarning,
		Verdict:                       StatusGate266Inherited + "; quark and lepton observable reconstructions are inherited as sealed algebraic reconstructions, not finite-core predictions",
	}
}

func compileGeometricLedger(inh Gate266Inheritance) GeometricDerivationLedger {
	items := []LedgerItem{
		{Name: "complexified Fock carrier S_C = Λ*(C^4)", SourceGate: "Gates 14, 233-235", Kind: "carrier", Status: "finite/bridge structural support", FiniteCore: true, Statement: "The flavor module has a lawful generation/carrier arena, but carrier existence is not a mass-amplitude theorem."},
		{Name: "native finite algebra C ⊕ M3(C)", SourceGate: "Gate 238", Kind: "finite algebra", Status: "conditional finite support", FiniteCore: true, Statement: "The M3(C) block supplies the natural bilinear matrix arena used later by the flavor closure."},
		{Name: "SM gauge/matter/charge kinematic ledger", SourceGate: "Gates 14-26, 166-167", Kind: "kinematics", Status: "derived kinematic support", FiniteCore: true, Statement: "Gauge representations and charge tables constrain the stage of flavor physics but do not assign Yukawa amplitudes."},
		{Name: "τ_eta = diag(2,-2,1) generation-breaking source", SourceGate: "Gates 242, 261", Kind: "generation topology", Status: "conditional support", FiniteCore: true, Statement: "τ_eta provides 1⊕1⊕1 generation capacity and a diagonal source map on M3(C)."},
		{Name: "ad_τ texture decomposition", SourceGate: "Gate 261", Kind: "texture algebra", Status: "conditional support", FiniteCore: true, Statement: "The commutator [τ,E_ij]=(λ_i-λ_j)E_ij splits M3(C) into a 3D diagonal commutant and 6D off-diagonal complement."},
		{Name: "Hermitian triality real/phase basis C+C^T and i(C-C^T)", SourceGate: "Gate 262", Kind: "mixing basis", Status: "conditional support", FiniteCore: true, Statement: "Triality supplies noncommuting off-diagonal directions, but Gate 262 forbids promoting symmetry bases into amplitudes."},
		{Name: "finite trace/action diagnostics on triality basis", SourceGate: "Gate 263", Kind: "action preflight", Status: "failed route for dynamics", FiniteCore: true, Statement: "Trace norms verify a basis metric but remain degenerate and do not select α, β, γ."},
	}
	return GeometricDerivationLedger{
		Items:                          items,
		SCarrierSpaceRecorded:          containsItem(items, "S_C"),
		FiniteAlgebraRecorded:          containsItem(items, "M3"),
		GaugeMatterChargeRecorded:      containsItem(items, "charge"),
		ThreeGenerationCapacity:        containsItem(items, "τ_eta"),
		TauEtaSourceMapRecorded:        containsItem(items, "source"),
		AdTauMixingComplementRecorded:  containsItem(items, "ad_τ"),
		TrialityHermitianBasisRecorded: containsItem(items, "triality"),
		YukawaAmplitudeDerived:         false,
		CKMPMNSDerived:                 inh.QuarkNativeDerivation || inh.LeptonNativeDerivation,
		FermionMassesDerived:           false,
		Verdict:                        StatusGeometricLedgerCompiled + "; finite geometry derives the kinematic flavor arena and structural texture bases, not numerical flavor dynamics",
	}
}

func compileEmpiricalLedger(inh Gate266Inheritance) EmpiricalInputLedger {
	items := []LedgerItem{
		{Name: "weak-frame embedding orientation", SourceGate: "Gates 256-259", Kind: "SpontaneousCarrierSeal", Status: StatusSpontaneousOrientationSealed, Sealed: true, Statement: "The scalar/spinor carrier alignment is a sealed vacuum-orientation datum."},
		{Name: "scalar VEV and Higgs alignment", SourceGate: "Gates 187, 256", Kind: "SpontaneousCarrierSeal", Status: StatusSpontaneousOrientationSealed, Sealed: true, Statement: "VEV orientation and absolute electroweak scale are boundary data unless a vacuum functional is derived."},
		{Name: "Dirac Yukawa amplitudes", SourceGate: "Gates 263-265", Kind: "EmpiricalYukawaSeal", Status: StatusNoNativeFlavorAmplitude, Sealed: true, Statement: "The coefficients that set quark and charged-lepton masses are empirical texture data."},
		{Name: "quark full textures, masses, and CKM entries", SourceGate: "Gate 265", Kind: "EmpiricalYukawaSeal", Status: StatusNoNativeCKMPMNS, Sealed: true, Statement: "SVD reconstructs observables from full textures, but the entries are not finite-core outputs."},
		{Name: "lepton charged texture, neutrino texture, and PMNS entries", SourceGate: "Gate 266", Kind: "EmpiricalYukawaSeal", Status: StatusNoNativeCKMPMNS, Sealed: true, Statement: "SVD/Takagi reconstructs PMNS from sealed lepton data."},
		{Name: "Majorana/Dirac neutrino nature and neutrino ordering", SourceGate: "Gate 266", Kind: "EmpiricalYukawaSeal", Status: StatusMajoranaNatureNotDerived, Sealed: true, Statement: "The Majorana branch is a representative sealed witness, not a finite theorem."},
	}
	return EmpiricalInputLedger{
		Items:                        items,
		SpontaneousCarrierSealActive: true,
		EmpiricalYukawaSealActive:    inh.EmpiricalYukawaSealActive,
		WeakFrameOrientationSealed:   true,
		ScalarVEVAlignmentSealed:     true,
		QuarkTexturesSealed:          inh.QuarkSVDCKMReconstructed,
		LeptonTexturesSealed:         inh.ChargedLeptonSVDReconstructed && inh.NeutrinoTakagiReconstructed,
		CKMEntriesSealed:             inh.QuarkSVDCKMReconstructed,
		PMNSEntriesSealed:            inh.PMNSReconstructed,
		MajoranaChoiceSealed:         !inh.MajoranaNatureFiniteDerived,
		DoesNotRewriteFiniteCore:     !inh.QuarkNativeDerivation && !inh.LeptonNativeDerivation,
		Verdict:                      StatusEmpiricalLedgerCompiled + "; all flavor amplitudes, full textures, mixing entries, and neutrino-nature choices remain quarantined boundary data",
	}
}

func verifyReconstructions(inh Gate266Inheritance) ReconstructionVerification {
	return ReconstructionVerification{
		QuarkSVDCKMVerified:             inh.QuarkSVDCKMReconstructed,
		ChargedLeptonSVDVerified:        inh.ChargedLeptonSVDReconstructed,
		MajoranaTakagiVerified:          inh.NeutrinoTakagiReconstructed,
		PMNSVerified:                    inh.PMNSReconstructed,
		SVDIsAlgebraicReconstruction:    true,
		TakagiIsAlgebraicReconstruction: true,
		ObservablePipelineWorksOnData:   inh.QuarkSVDCKMReconstructed && inh.ChargedLeptonSVDReconstructed && inh.NeutrinoTakagiReconstructed && inh.PMNSReconstructed,
		ObservablePipelinePredictsData:  false,
		FiniteCorePolluted:              inh.QuarkNativeDerivation || inh.LeptonNativeDerivation || !inh.EmpiricalBoundaryPreserved,
		Verdict:                         StatusReconstructionManifestClosed + "; SVD/Takagi reconstruct flavor observables from sealed textures, while prediction remains forbidden",
	}
}

func defineFutureCriteria() FutureTheoremCriteria {
	criteria := []FutureCriterion{
		{Name: "canonical finite Dirac operator D_F on doubled S_C", Required: true, Satisfied: false, Reason: "Gates 233-235 opened legal matrix families and real-structure sieves, but no canonical physical D_F selector is derived."},
		{Name: "finite spectral action / heat-kernel map", Required: true, Satisfied: false, Reason: "A lawful Tr(f(D_F/Λ)) or finite Seeley-de Witt analogue must be available before amplitudes can be dynamics."},
		{Name: "computed a0, a2, a4 coefficients with normalization scheme", Required: true, Satisfied: false, Reason: "Trace diagnostics alone are not enough; coefficient extraction requires cutoff, subtraction, gauge projection, and normalization conventions."},
		{Name: "action map from spectral coefficients to M3(C) Yukawa weights", Required: true, Satisfied: false, Reason: "The theorem must output sector-dependent weights for τ_eta, C+C^T, i(C-C^T), or a larger justified texture basis."},
		{Name: "Hopf/B_gap/topological phase projection into generation endomorphisms", Required: true, Satisfied: false, Reason: "Gate 263 found no functor from scalar scales or Hopf residuals into off-diagonal M3(C) amplitudes."},
		{Name: "non-empirical mass and mixing predictions with residual audit", Required: true, Satisfied: false, Reason: "A seal-lifting theorem must produce masses/CKM/PMNS numbers before seeing the empirical ledger, then compare residuals."},
	}
	return FutureTheoremCriteria{
		Criteria:                        criteria,
		RequiresFiniteSpectralAction:    true,
		RequiresCanonicalFiniteDirac:    true,
		RequiresHeatKernelCoefficients:  true,
		RequiresYukawaAmplitudeMap:      true,
		RequiresHopfOrBGapProjection:    true,
		RequiresNormalizationScheme:     true,
		RequiresMassAndMixingPrediction: true,
		CurrentGateCanLiftSeal:          false,
		RecommendedNextGate:             "Gate 268 — Finite Spectral Action Re-Attempt / Seeley-de Witt a0-a2-a4 Coefficient Audit on doubled S_C",
		Verdict:                         StatusFutureTheoremCriteriaDefined + "; any future flavor seal lift requires a native finite spectral/action theorem, not another empirical fit",
	}
}

func buildFirewall(geo GeometricDerivationLedger, emp EmpiricalInputLedger, recon ReconstructionVerification, criteria FutureTheoremCriteria) FirewallManifest {
	return FirewallManifest{
		KinematicsDerived:               geo.SCarrierSpaceRecorded && geo.ThreeGenerationCapacity && geo.AdTauMixingComplementRecorded && geo.TrialityHermitianBasisRecorded,
		DynamicsSealed:                  emp.EmpiricalYukawaSealActive && emp.QuarkTexturesSealed && emp.LeptonTexturesSealed,
		SpontaneousCarrierSealPreserved: emp.SpontaneousCarrierSealActive && emp.WeakFrameOrientationSealed && emp.ScalarVEVAlignmentSealed,
		EmpiricalYukawaSealPreserved:    emp.EmpiricalYukawaSealActive && emp.DoesNotRewriteFiniteCore,
		NoMassPredictionClaim:           !geo.FermionMassesDerived && !recon.ObservablePipelinePredictsData,
		NoCKMPMNSPredictionClaim:        !geo.CKMPMNSDerived && !recon.ObservablePipelinePredictsData,
		NoMajoranaNatureClaim:           emp.MajoranaChoiceSealed,
		NoSpectralActionClaim:           !criteria.CurrentGateCanLiftSeal,
		ClosureDoesNotAddNewPhysics:     true,
		FiniteCorePolluted:              recon.FiniteCorePolluted,
		Verdict:                         StatusFullFlavorLedgerClosed + "; the engine derives flavor kinematics, reconstructs sealed phenomenology, and closes numerical flavor dynamics behind explicit seals",
	}
}

func summarize(inh Gate266Inheritance, geo GeometricDerivationLedger, emp EmpiricalInputLedger, recon ReconstructionVerification, criteria FutureTheoremCriteria, firewall FirewallManifest) Summary {
	return Summary{
		Gate266Inherited:            inh.EmpiricalYukawaSealActive && inh.EmpiricalBoundaryPreserved,
		GeometricLedgerClosed:       geo.SCarrierSpaceRecorded && geo.TauEtaSourceMapRecorded && geo.AdTauMixingComplementRecorded && !geo.YukawaAmplitudeDerived,
		EmpiricalLedgerClosed:       emp.EmpiricalYukawaSealActive && emp.QuarkTexturesSealed && emp.LeptonTexturesSealed && emp.CKMEntriesSealed && emp.PMNSEntriesSealed,
		ReconstructionsVerified:     recon.ObservablePipelineWorksOnData && !recon.ObservablePipelinePredictsData,
		FutureCriteriaDefined:       len(criteria.Criteria) >= 6 && !criteria.CurrentGateCanLiftSeal,
		FullFlavorLedgerClosed:      firewall.KinematicsDerived && firewall.DynamicsSealed && firewall.EmpiricalYukawaSealPreserved && !firewall.FiniteCorePolluted,
		NativeFlavorDynamicsDerived: false,
		Status:                      StatusFullFlavorLedgerClosed,
		NextGate:                    criteria.RecommendedNextGate,
		Comment:                     "Gate 267 is a closure manifest: it does not predict new flavor numbers; it records the exact boundary and the theorem obligations required to reopen it.",
	}
}

func buildTruth(inh Gate266Inheritance, geo GeometricDerivationLedger, emp EmpiricalInputLedger, recon ReconstructionVerification, criteria FutureTheoremCriteria, firewall FirewallManifest) string {
	parts := []string{
		"Gate 267 closes the ASHA flavor ledger: finite Cℓ(1,7) geometry supplies the carrier, generation, and texture-basis kinematics, while all numerical Yukawa amplitudes, CKM/PMNS entries, fermion masses, and neutrino-nature choices remain sealed empirical boundary data.",
		fmt.Sprintf("Inherited reconstructions: quarkSVDCKM=%t chargedLeptonSVD=%t neutrinoTakagi=%t PMNS=%t boundaryPreserved=%t.", inh.QuarkSVDCKMReconstructed, inh.ChargedLeptonSVDReconstructed, inh.NeutrinoTakagiReconstructed, inh.PMNSReconstructed, inh.EmpiricalBoundaryPreserved),
		fmt.Sprintf("Derived ledger items=%d sealed ledger items=%d future criteria=%d.", len(geo.Items), len(emp.Items), len(criteria.Criteria)),
		fmt.Sprintf("Firewall: kinematicsDerived=%t dynamicsSealed=%t massPrediction=%t ckmPmnsPrediction=%t finitePollution=%t.", firewall.KinematicsDerived, firewall.DynamicsSealed, !firewall.NoMassPredictionClaim, !firewall.NoCKMPMNSPredictionClaim, firewall.FiniteCorePolluted),
	}
	return strings.Join(parts, " ")
}

func containsItem(items []LedgerItem, needle string) bool {
	needle = strings.ToLower(needle)
	for _, it := range items {
		if strings.Contains(strings.ToLower(it.Name), needle) || strings.Contains(strings.ToLower(it.Statement), needle) {
			return true
		}
	}
	return false
}
