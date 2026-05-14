// Package flavoralignmentdmabsence implements Gate 224: Flavor alignment safety
// audit / Dark Matter absence theorem.
//
// Gate 223 rescued the sealed PeV spectrum from immediate relic falsification by
// granting a RelicDecaySeal conditional on explicit EFT portals. Gate 224 asks
// the next necessary phenomenological question: do these portals create a flavor
// disaster, and what do their fast decays imply for heavy-sector dark matter?
//
// The gate does not compute precision rare-decay rates. That would require a
// flavor basis, CKM/PMNS embedding, Wilson matrices, and hadronic matrix
// elements, none of which are finite-core theorems. Instead it formally seals the
// minimal requirement: the portal tensors must be aligned predominantly with the
// third generation. Under that FlavorAlignmentSeal and the already granted
// RelicDecaySeal, the sealed PeV carriers have zero present-day stable relic
// abundance; dark matter must be sought outside this heavy threshold sector.
package flavoralignmentdmabsence

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/coloredoctetportal"
)

const (
	AuditID = "GATE224-FLAVOR-ALIGNMENT-DARK-MATTER-ABSENCE-AUDIT"

	StatusConditionalFlavorAligned = "CONDITIONAL_PHENOMENOLOGY_FLAVOR_ALIGNMENT_SEAL_GRANTED"
	StatusHeavyDMAbsent            = "HEAVY_SECTOR_DARK_MATTER_ABSENCE_THEOREM"
	StatusRelicSealPreserved       = "RELIC_DECAY_SEAL_PRESERVED_UNDER_FLAVOR_ALIGNMENT"
)

type Gate223Snapshot struct {
	Gate223Inherited             bool
	RelicDecaySealGranted        bool
	TripletPortalActive          bool
	OctetPortalActive            bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	MBGeV                        float64
	LambdaEFTMaxGeV              float64
	RequiredWidthGeV             float64
	TruthStatement               string
}

type PortalFlavorTensor struct {
	Name                             string
	SymbolicForm                     string
	Carrier                          string
	FlavorIndices                    []string
	GenericTensorComponents          int
	PreferredAlignedEntry            string
	EFTScaleGeV                      float64
	GeneratesFlavorRisk              bool
	RiskClasses                      []string
	ThirdGenerationAlignmentRequired bool
	NativeAlignmentDerived           bool
	SealEligible                     bool
	Verdict                          string
}

type FlavorFCNCAudit struct {
	PortalsAudited              []PortalFlavorTensor
	GenericFlavorSafe           bool
	ArbitraryFirstSecondAllowed bool
	FlavorChangingRiskLogged    bool
	ExactRareDecayRatesComputed bool
	HadronicMatrixElementsKnown bool
	WilsonMatricesKnown         bool
	CKMPMNSBasisDerived         bool
	AlignmentCanProtect         bool
	Verdict                     string
}

type FlavorAlignmentSealAudit struct {
	SealName                   string
	SealGranted                bool
	AlignmentRule              string
	QuarantinedInputs          []string
	ForbiddenWithoutFutureSeal []string
	NativeFlavorTheoremDerived bool
	StillPhenomenological      bool
	OperationalStatus          string
	Verdict                    string
}

type DarkMatterAbsenceAudit struct {
	RelicDecaySealActive      bool
	FlavorAlignmentSealActive bool
	TripletDecaysBeforeBBN    bool
	OctetDecaysBeforeBBN      bool
	PresentDayStableFraction  float64
	OmegaHeavySectorH2        float64
	ThermalFreezeoutComputed  bool
	BoltzmannHistoryComputed  bool
	HeavySectorDMCandidate    bool
	DarkMatterDeferredTo      []string
	TheoremName               string
	Verdict                   string
}

type FutureRouteAudit struct {
	RequiredNextObjects []string
	OpenDMInventory     []string
	OpenFlavorInventory []string
	Verdict             string
}

type FirewallAudit struct {
	Gate223Inherited             bool
	RelicDecaySealActive         bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	NativeFlavorClaimed          bool
	ExactFCNCRatesClaimed        bool
	FlavorInputsTuned            bool
	WilsonCoefficientsDerived    bool
	RelicAbundanceThermalClaimed bool
	HeavySectorDMClaimed         bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	FlavorAlignmentSealGranted bool
	GenericFlavorObstructed    bool
	HeavySectorDMAbsent        bool
	RelicDecaySealStillValid   bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	Gate223         Gate223Snapshot
	Gate223Analysis coloredoctetportal.Analysis
	Flavor          FlavorFCNCAudit
	Seal            FlavorAlignmentSealAudit
	DarkMatter      DarkMatterAbsenceAudit
	Future          FutureRouteAudit
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
	defaultOnce.Do(func() {
		g223, err := coloredoctetportal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 223 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g223)
	})
	return defaultA, defaultErr
}

func Build(g223 coloredoctetportal.Analysis) (Analysis, error) {
	snap := snapshotFromGate223(g223)
	if !snap.Gate223Inherited || !snap.RelicDecaySealGranted || !snap.TripletPortalActive || !snap.OctetPortalActive {
		return Analysis{}, fmt.Errorf("Gate 224 requires Gate 223 RelicDecaySeal with both triplet and octet portals active")
	}
	flavor := auditFlavor(snap)
	seal := auditFlavorAlignmentSeal(flavor)
	dm := auditDarkMatterAbsence(snap, seal)
	future := auditFutureRoutes()
	firewall := auditFirewall(snap, seal, dm)
	summary := summarize(flavor, seal, dm)
	truth := buildTruth(snap, flavor, seal, dm, summary)
	return Analysis{Gate223: snap, Gate223Analysis: g223, Flavor: flavor, Seal: seal, DarkMatter: dm, Future: future, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate223(a coloredoctetportal.Analysis) Gate223Snapshot {
	return Gate223Snapshot{
		Gate223Inherited:             a.Summary.FullRelicDecaySeal && !a.Summary.Rank1SpectrumFalsified && a.TensorSearch.OctetPortalFound,
		RelicDecaySealGranted:        a.RelicSeal.SealGranted,
		TripletPortalActive:          a.RelicSeal.TripletPortal != "",
		OctetPortalActive:            a.RelicSeal.OctetPortal != "" && a.TensorSearch.OctetPortalFound,
		ThresholdSpectrumSealActive:  a.Firewall.ThresholdSpectrumSealActive,
		MatchingCorrectionSealActive: a.Firewall.MatchingCorrectionSealActive,
		EmpiricalCarrierSealActive:   a.Firewall.EmpiricalCarrierSealActive,
		LeptoquarkDynamicsSealActive: a.Firewall.LeptoquarkDynamicsSealActive,
		MBGeV:                        a.Kinematics.MBGeV,
		LambdaEFTMaxGeV:              a.Kinematics.ConservativeLambdaMaxGeV,
		RequiredWidthGeV:             a.Kinematics.RequiredWidthGeV,
		TruthStatement:               a.TruthStatement,
	}
}

func auditFlavor(snap Gate223Snapshot) FlavorFCNCAudit {
	portals := []PortalFlavorTensor{
		{
			Name:                             "triplet lepton-Higgs portal",
			SymbolicForm:                     "y_T^i Ψ_3^a (L_i σ^a H†) + h.c.",
			Carrier:                          "Dirac (1,3,Y=1)",
			FlavorIndices:                    []string{"i ∈ {e, μ, τ}"},
			GenericTensorComponents:          3,
			PreferredAlignedEntry:            "i=3 / τ-family lepton doublet",
			EFTScaleGeV:                      snap.MBGeV,
			GeneratesFlavorRisk:              true,
			RiskClasses:                      []string{"charged-lepton flavor violation if multiple lepton flavors are active", "non-aligned heavy-lepton mixing", "radiative dipole operators after electroweak symmetry breaking"},
			ThirdGenerationAlignmentRequired: true,
			NativeAlignmentDerived:           false,
			SealEligible:                     true,
			Verdict:                          "generic lepton-flavor tensor is not certified safe; third-generation alignment is required as a seal",
		},
		{
			Name:                             "octet three-fermion portal",
			SymbolicForm:                     "(c_8^{ijk}/Λ²) bar(Ψ8)^a_i (Q_i u^c_j e^c_k)^a + h.c.",
			Carrier:                          "Dirac (8,2,Y=1/2)",
			FlavorIndices:                    []string{"i,j,k ∈ {1,2,3}"},
			GenericTensorComponents:          27,
			PreferredAlignedEntry:            "Q_3 u^c_3 τ^c",
			EFTScaleGeV:                      snap.LambdaEFTMaxGeV,
			GeneratesFlavorRisk:              true,
			RiskClasses:                      []string{"charged-lepton flavor violation", "ΔF=2 meson mixing through non-aligned quark tensors", "rare top/bottom/tau decay tails", "flavor off-diagonal four-fermion operators"},
			ThirdGenerationAlignmentRequired: true,
			NativeAlignmentDerived:           false,
			SealEligible:                     true,
			Verdict:                          "generic 27-component tensor is flavor-dangerous; only aligned third-generation dominance is admitted by the seal",
		},
		{
			Name:                             "octet chromomagnetic-Higgs-lepton portal",
			SymbolicForm:                     "(c'_8{}^k/Λ²) bar(Ψ8)^a_i σ^{μν} e^c_k H†_i G^a_{μν} + h.c.",
			Carrier:                          "Dirac (8,2,Y=1/2)",
			FlavorIndices:                    []string{"k ∈ {e, μ, τ}"},
			GenericTensorComponents:          3,
			PreferredAlignedEntry:            "τ^c",
			EFTScaleGeV:                      snap.LambdaEFTMaxGeV,
			GeneratesFlavorRisk:              true,
			RiskClasses:                      []string{"leptonic dipole flavor violation if k is not aligned", "electron/muon precision tails", "post-EWSB chromomagnetic heavy decay flavor choice"},
			ThirdGenerationAlignmentRequired: true,
			NativeAlignmentDerived:           false,
			SealEligible:                     true,
			Verdict:                          "dipole portal is viable only as a sealed aligned τ-family operator; no native flavor theorem is present",
		},
	}
	return FlavorFCNCAudit{
		PortalsAudited:              portals,
		GenericFlavorSafe:           false,
		ArbitraryFirstSecondAllowed: false,
		FlavorChangingRiskLogged:    true,
		ExactRareDecayRatesComputed: false,
		HadronicMatrixElementsKnown: false,
		WilsonMatricesKnown:         false,
		CKMPMNSBasisDerived:         false,
		AlignmentCanProtect:         allSealEligible(portals),
		Verdict:                     "generic flavor tensors are not accepted as safe; a FlavorAlignmentSeal is required before the RelicDecaySeal can be used phenomenologically",
	}
}

func allSealEligible(xs []PortalFlavorTensor) bool {
	for _, x := range xs {
		if !x.SealEligible || !x.ThirdGenerationAlignmentRequired || x.NativeAlignmentDerived {
			return false
		}
	}
	return true
}

func auditFlavorAlignmentSeal(f FlavorFCNCAudit) FlavorAlignmentSealAudit {
	grant := !f.GenericFlavorSafe && !f.ArbitraryFirstSecondAllowed && f.FlavorChangingRiskLogged && f.AlignmentCanProtect
	status := "FLAVOR_ALIGNMENT_SEAL_DENIED"
	verdict := "FlavorAlignmentSeal is denied"
	if grant {
		status = "FLAVOR_ALIGNMENT_SEAL_GRANTED_CONDITIONAL_ON_THIRD_GENERATION_DOMINANCE"
		verdict = "FlavorAlignmentSeal granted: decay-portal flavor tensors are quarantined to third-generation-dominant entries; this is phenomenological alignment, not a finite theorem"
	}
	return FlavorAlignmentSealAudit{
		SealName:                   "FlavorAlignmentSeal",
		SealGranted:                grant,
		AlignmentRule:              "c_8^{ijk}, c'_8{}^k, y_T^i are zero or negligibly small outside third-generation aligned entries unless a future finite flavor theorem says otherwise",
		QuarantinedInputs:          []string{"portal flavor tensors", "generation basis", "CKM/PMNS leakage model", "rare-decay Wilson matrices", "hadronic matrix elements", "experimental flavor likelihoods"},
		ForbiddenWithoutFutureSeal: []string{"arbitrary first-generation portal couplings", "arbitrary second-generation portal couplings", "generic flavor-anarchic c_8^{ijk}", "claiming FCNC safety from gauge invariance alone"},
		NativeFlavorTheoremDerived: false,
		StillPhenomenological:      true,
		OperationalStatus:          status,
		Verdict:                    verdict,
	}
}

func auditDarkMatterAbsence(snap Gate223Snapshot, seal FlavorAlignmentSealAudit) DarkMatterAbsenceAudit {
	decaySafe := snap.RelicDecaySealGranted && seal.SealGranted && snap.RequiredWidthGeV > 0
	omega := 0.0
	stable := 1.0
	if decaySafe {
		stable = 0.0
	}
	return DarkMatterAbsenceAudit{
		RelicDecaySealActive:      snap.RelicDecaySealGranted,
		FlavorAlignmentSealActive: seal.SealGranted,
		TripletDecaysBeforeBBN:    decaySafe,
		OctetDecaysBeforeBBN:      decaySafe,
		PresentDayStableFraction:  stable,
		OmegaHeavySectorH2:        omega,
		ThermalFreezeoutComputed:  false,
		BoltzmannHistoryComputed:  false,
		HeavySectorDMCandidate:    false,
		DarkMatterDeferredTo:      []string{"seven unassigned contact partial-overlap modes", "B-sector spectral gap / axion-like route", "future finite stable neutral sector", "non-heavy-threshold cosmological sector"},
		TheoremName:               "Heavy_Sector_Dark_Matter_Absence_Theorem",
		Verdict:                   "under RelicDecaySeal + FlavorAlignmentSeal, the PeV carriers decay before BBN and leave no present-day heavy-sector dark matter; Ω_heavy h²=0 for this sector only",
	}
}

func auditFutureRoutes() FutureRouteAudit {
	return FutureRouteAudit{
		RequiredNextObjects: []string{"finite or sealed flavor basis", "rare-decay Wilson-coefficient envelope", "Boltzmann relic calculation for the sealed decays", "independent dark-matter candidate audit"},
		OpenDMInventory:     []string{"contact modes", "B-sector gap", "sterile/Fock vacuum seed", "new finite neutral sector if later derived"},
		OpenFlavorInventory: []string{"generation-breaking source map", "Yukawa texture seal", "CKM/PMNS alignment bridge", "third-generation dominance mechanism"},
		Verdict:             "the RG/cosmology bridge is conditionally viable, but dark matter and flavor precision are now separate frontier problems",
	}
}

func auditFirewall(snap Gate223Snapshot, seal FlavorAlignmentSealAudit, dm DarkMatterAbsenceAudit) FirewallAudit {
	return FirewallAudit{
		Gate223Inherited:             snap.Gate223Inherited,
		RelicDecaySealActive:         snap.RelicDecaySealGranted,
		ThresholdSpectrumSealActive:  snap.ThresholdSpectrumSealActive,
		MatchingCorrectionSealActive: snap.MatchingCorrectionSealActive,
		EmpiricalCarrierSealActive:   snap.EmpiricalCarrierSealActive,
		LeptoquarkDynamicsSealActive: snap.LeptoquarkDynamicsSealActive,
		NativeFlavorClaimed:          seal.NativeFlavorTheoremDerived,
		ExactFCNCRatesClaimed:        false,
		FlavorInputsTuned:            false,
		WilsonCoefficientsDerived:    false,
		RelicAbundanceThermalClaimed: dm.ThermalFreezeoutComputed || dm.BoltzmannHistoryComputed,
		HeavySectorDMClaimed:         dm.HeavySectorDMCandidate,
		FiniteCorePolluted:           false,
		Verdict:                      "Gate 224 grants only a phenomenological flavor-alignment seal and a conditional heavy-sector DM absence statement; it does not derive flavor, Wilson coefficients, or a dark-matter model",
	}
}

func summarize(f FlavorFCNCAudit, seal FlavorAlignmentSealAudit, dm DarkMatterAbsenceAudit) Summary {
	status := StatusConditionalFlavorAligned + "; " + StatusHeavyDMAbsent + "; " + StatusRelicSealPreserved
	return Summary{
		FlavorAlignmentSealGranted: seal.SealGranted,
		GenericFlavorObstructed:    !f.GenericFlavorSafe && !f.ArbitraryFirstSecondAllowed,
		HeavySectorDMAbsent:        dm.OmegaHeavySectorH2 == 0 && !dm.HeavySectorDMCandidate && dm.PresentDayStableFraction == 0,
		RelicDecaySealStillValid:   dm.RelicDecaySealActive && seal.SealGranted,
		Status:                     status,
		NextGate:                   "Gate 225 — dark-matter candidate inventory / finite neutral-sector viability audit",
		Comment:                    "The PeV threshold sector is conditionally cosmology-safe only after RelicDecaySeal and FlavorAlignmentSeal; it cannot be the dark matter sector.",
	}
}

func buildTruth(snap Gate223Snapshot, f FlavorFCNCAudit, seal FlavorAlignmentSealAudit, dm DarkMatterAbsenceAudit, s Summary) string {
	return fmt.Sprintf("Gate 224 inherits Gate 223's RelicDecaySeal at M_B=%.9g GeV and Λ_EFT≲%.9g GeV. Generic flavor tensors are rejected as unsafe; %d portal flavor structures require third-generation alignment, so %s is granted only as phenomenology. Under RelicDecaySeal + FlavorAlignmentSeal the heavy PeV carriers decay before BBN and have Ω_heavy h²=%.1f today. Dark matter is deferred to the unassigned finite core. Status=%s.", snap.MBGeV, snap.LambdaEFTMaxGeV, len(f.PortalsAudited), seal.SealName, dm.OmegaHeavySectorH2, s.Status)
}

func FormatPortal(p PortalFlavorTensor) string {
	return fmt.Sprintf("%s carrier=%s form=%s indices=%s components=%d preferred=%s EFTscale=%.6g flavorRisk=%t risks=[%s] thirdGenSealRequired=%t nativeAlignment=%t sealEligible=%t :: %s", p.Name, p.Carrier, p.SymbolicForm, strings.Join(p.FlavorIndices, ","), p.GenericTensorComponents, p.PreferredAlignedEntry, p.EFTScaleGeV, p.GeneratesFlavorRisk, strings.Join(p.RiskClasses, "; "), p.ThirdGenerationAlignmentRequired, p.NativeAlignmentDerived, p.SealEligible, p.Verdict)
}

func FormatFlavor(f FlavorFCNCAudit) string {
	parts := make([]string, 0, len(f.PortalsAudited))
	for _, p := range f.PortalsAudited {
		parts = append(parts, FormatPortal(p))
	}
	return fmt.Sprintf("genericSafe=%t firstSecondAllowed=%t riskLogged=%t exactRates=%t hadronicKnown=%t WilsonKnown=%t CKM_PMNS=%t alignmentProtects=%t portals={%s} :: %s", f.GenericFlavorSafe, f.ArbitraryFirstSecondAllowed, f.FlavorChangingRiskLogged, f.ExactRareDecayRatesComputed, f.HadronicMatrixElementsKnown, f.WilsonMatricesKnown, f.CKMPMNSBasisDerived, f.AlignmentCanProtect, strings.Join(parts, " | "), f.Verdict)
}

func FormatSeal(s FlavorAlignmentSealAudit) string {
	return fmt.Sprintf("seal=%s granted=%t rule=%q quarantined=[%s] forbidden=[%s] nativeFlavor=%t phenomenological=%t status=%s :: %s", s.SealName, s.SealGranted, s.AlignmentRule, strings.Join(s.QuarantinedInputs, "; "), strings.Join(s.ForbiddenWithoutFutureSeal, "; "), s.NativeFlavorTheoremDerived, s.StillPhenomenological, s.OperationalStatus, s.Verdict)
}

func FormatDarkMatter(d DarkMatterAbsenceAudit) string {
	return fmt.Sprintf("relicSeal=%t flavorSeal=%t tripletBeforeBBN=%t octetBeforeBBN=%t stableFraction=%.1f OmegaHeavyH2=%.1f thermal=%t boltzmann=%t heavyDMCandidate=%t deferred=[%s] theorem=%s :: %s", d.RelicDecaySealActive, d.FlavorAlignmentSealActive, d.TripletDecaysBeforeBBN, d.OctetDecaysBeforeBBN, d.PresentDayStableFraction, d.OmegaHeavySectorH2, d.ThermalFreezeoutComputed, d.BoltzmannHistoryComputed, d.HeavySectorDMCandidate, strings.Join(d.DarkMatterDeferredTo, "; "), d.TheoremName, d.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate223=%t relicSeal=%t thresholdSeal=%t matchingSeal=%t carrierSeal=%t lqSeal=%t nativeFlavor=%t exactFCNC=%t flavorTuned=%t WilsonDerived=%t thermalRelicClaim=%t heavyDMClaim=%t finitePolluted=%t :: %s", f.Gate223Inherited, f.RelicDecaySealActive, f.ThresholdSpectrumSealActive, f.MatchingCorrectionSealActive, f.EmpiricalCarrierSealActive, f.LeptoquarkDynamicsSealActive, f.NativeFlavorClaimed, f.ExactFCNCRatesClaimed, f.FlavorInputsTuned, f.WilsonCoefficientsDerived, f.RelicAbundanceThermalClaimed, f.HeavySectorDMClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("flavorSeal=%t genericFlavorObstructed=%t heavyDMAbsent=%t relicSealValid=%t status=%s next=%s :: %s", s.FlavorAlignmentSealGranted, s.GenericFlavorObstructed, s.HeavySectorDMAbsent, s.RelicDecaySealStillValid, s.Status, s.NextGate, s.Comment)
}
