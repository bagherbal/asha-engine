// Package pevobservabilityaudit implements Gate 220: PeV-threshold
// indirect-signature / experimental observability audit.
//
// Gate 219 established a sealed, single-scale PeV threshold for the
// Dirac (1,3,Y=1) + Dirac (8,2,Y=1/2) spectrum under the
// ThresholdSpectrumSeal and MatchingCorrectionSeal. Gate 220 does not add a
// new finite derivation. It asks whether that sealed PeV spectrum has obvious
// indirect phenomenological failures: electroweak precision parameters,
// Higgs-loop imprints, direct-reach separation, and cosmological relic safety.
package pevobservabilityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/inputsensitivityaudit"
)

const (
	StatusConditionalPhenomenology = "CONDITIONAL_PHENOMENOLOGY_PEV_INDIRECT_OBSERVABILITY_AUDIT"
	StatusWarningCosmology         = "CONDITIONAL_PHENOMENOLOGY_WITH_STABLE_RELIC_WARNING"

	ObservabilityAuditID = "GATE220-PEV-THRESHOLD-INDIRECT-SIGNATURE-AUDIT"
	mzGeV                = 91.1876
	vevGeV               = 246.21965
	mhGeV                = 125.20
	alphaEMInvMZ         = 127.955
	sin2ThetaW           = 0.23122
	fccProxyGeV          = 100000.0
)

type Gate219Snapshot struct {
	Gate219Inherited             bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	MatchingPlausible            bool
	BottomTauComplete            bool
	MBGeV                        float64
	MBMinGeV                     float64
	MBMaxGeV                     float64
	MStarGeV                     float64
	MStarMinGeV                  float64
	MStarMaxGeV                  float64
	WorstResidualOverEpsilon     float64
	TruthStatement               string
}

type SealedSpectrum struct {
	Name                  string
	Row1Name              string
	Row2Name              string
	Row1Rep               string
	Row2Rep               string
	Row1Dimension         int
	Row2Dimension         int
	Row1Charges           []int
	Row2ContainsColor     bool
	Row2ContainsWeak      bool
	SingleScaleMBGeV      float64
	ScaleRangeMinGeV      float64
	ScaleRangeMaxGeV      float64
	ConditionalOnly       bool
	DecayOperatorsDerived bool
	MassSplittingsDerived bool
	Verdict               string
}

type DirectReachAudit struct {
	DirectProductionReachGeV float64
	CentralMassGeV           float64
	MassOverReach            float64
	DirectProductionSafe     bool
	FutureHundredTeVProxy    bool
	Verdict                  string
}

type ElectroweakPrecisionAudit struct {
	MassScaleGeV               float64
	VEVOverM                   float64
	VEVOverMSquared            float64
	ParametricDeltaMGeV        float64
	TProxy                     float64
	TProxyFormula              string
	TreeLevelViolationDerived  bool
	HeavyYukawaCouplingDerived bool
	ObliqueSafe                bool
	Verdict                    string
}

type HiggsLoopAudit struct {
	ColoredOctetPresent          bool
	ChargedTripletPresent        bool
	HeavyYukawaCouplingDerived   bool
	NonDecouplingMassFromHiggs   bool
	DecouplingParameter          float64
	DiphotonAmplitudeProxy       float64
	GluonFusionAmplitudeProxy    float64
	HiggsLoopSafeUnderDecoupling bool
	Verdict                      string
}

type CosmologyAudit struct {
	TripletHasNeutralComponent      bool
	TripletChargedComponents        []int
	ColorOctetContainsColoredStates bool
	DecayOperatorDerived            bool
	MassSplittingDerived            bool
	StableNeutralRelicWarning       bool
	StableChargedRelicWarning       bool
	StableColoredRelicWarning       bool
	DarkMatterCandidateClaimed      bool
	OverclosureComputed             bool
	RecommendedSeal                 string
	Verdict                         string
}

type ObservabilitySummary struct {
	EWPOSafe               bool
	HiggsLoopSafe          bool
	DirectReachSafe        bool
	CosmologyWarning       bool
	FatalObservableFailure bool
	FutureTargets          []string
	Status                 string
	Comment                string
}

type FirewallAudit struct {
	Gate219Inherited                bool
	ThresholdSpectrumSealActive     bool
	MatchingCorrectionSealActive    bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	PeVMassFiniteDerived            bool
	DecayOperatorInvented           bool
	HeavyHiggsYukawaInvented        bool
	MassSplittingInvented           bool
	DarkMatterClaimed               bool
	OverclosureComputed             bool
	PhysicalObservationClaimed      bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Analysis struct {
	Gate219         Gate219Snapshot
	Gate219Analysis inputsensitivityaudit.Analysis
	Spectrum        SealedSpectrum
	DirectReach     DirectReachAudit
	EWPO            ElectroweakPrecisionAudit
	HiggsLoops      HiggsLoopAudit
	Cosmology       CosmologyAudit
	Summary         ObservabilitySummary
	Firewall        FirewallAudit
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g219, err := inputsensitivityaudit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 219 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g219)
	})
	return defaultA, defaultErr
}

func Build(g219 inputsensitivityaudit.Analysis) (Analysis, error) {
	snap := snapshotFromGate219(g219)
	if !snap.Gate219Inherited || !snap.ThresholdSpectrumSealActive || !snap.MatchingCorrectionSealActive || !snap.MatchingPlausible {
		return Analysis{}, fmt.Errorf("Gate 220 requires the Gate 219 sealed PeV-scale matching-plausible input-sensitivity audit")
	}
	spectrum := buildSpectrum(snap)
	direct := auditDirectReach(spectrum)
	ewpo := auditEWPO(spectrum)
	higgs := auditHiggsLoops(spectrum)
	cosmo := auditCosmology(spectrum)
	summary := summarize(direct, ewpo, higgs, cosmo)
	firewall := auditFirewall(g219, snap)
	truth := buildTruth(spectrum, direct, ewpo, higgs, cosmo, summary)
	return Analysis{Gate219: snap, Gate219Analysis: g219, Spectrum: spectrum, DirectReach: direct, EWPO: ewpo, HiggsLoops: higgs, Cosmology: cosmo, Summary: summary, Firewall: firewall, TruthStatement: truth}, nil
}

func snapshotFromGate219(a inputsensitivityaudit.Analysis) Gate219Snapshot {
	return Gate219Snapshot{
		Gate219Inherited:             a.Summary.Status == inputsensitivityaudit.StatusConditionalPhenomenology,
		ThresholdSpectrumSealActive:  a.Firewall.ThresholdSpectrumSealInherited,
		MatchingCorrectionSealActive: a.Firewall.MatchingCorrectionSealActive,
		MatchingPlausible:            a.CentralFit.MatchingPlausible && a.Sensitivity.BrokenEnvelopeCases == 0,
		BottomTauComplete:            a.Completeness.BottomYukawaIncluded && a.Completeness.TauYukawaIncluded,
		MBGeV:                        a.CentralFit.MBGeV,
		MBMinGeV:                     a.Sensitivity.MBMinGeV,
		MBMaxGeV:                     a.Sensitivity.MBMaxGeV,
		MStarGeV:                     a.CentralFit.MStarGeV,
		MStarMinGeV:                  a.Sensitivity.MStarMinGeV,
		MStarMaxGeV:                  a.Sensitivity.MStarMaxGeV,
		WorstResidualOverEpsilon:     a.Sensitivity.WorstResidualOverEpsilon,
		TruthStatement:               a.TruthStatement,
	}
}

func buildSpectrum(g Gate219Snapshot) SealedSpectrum {
	return SealedSpectrum{
		Name:                  "sealed PeV single-scale spectrum",
		Row1Name:              "Dirac electroweak triplet",
		Row2Name:              "Dirac color-octet weak doublet",
		Row1Rep:               "(1,3,Y=1)",
		Row2Rep:               "(8,2,Y=1/2)",
		Row1Dimension:         3,
		Row2Dimension:         16,
		Row1Charges:           []int{0, 1, 2},
		Row2ContainsColor:     true,
		Row2ContainsWeak:      true,
		SingleScaleMBGeV:      g.MBGeV,
		ScaleRangeMinGeV:      g.MBMinGeV,
		ScaleRangeMaxGeV:      g.MBMaxGeV,
		ConditionalOnly:       true,
		DecayOperatorsDerived: false,
		MassSplittingsDerived: false,
		Verdict:               "the spectrum is inherited only under ThresholdSpectrumSeal and MatchingCorrectionSeal; carrier decays and splittings remain un-derived",
	}
}

func auditDirectReach(s SealedSpectrum) DirectReachAudit {
	ratio := s.SingleScaleMBGeV / fccProxyGeV
	return DirectReachAudit{DirectProductionReachGeV: fccProxyGeV, CentralMassGeV: s.SingleScaleMBGeV, MassOverReach: ratio, DirectProductionSafe: ratio > 10, FutureHundredTeVProxy: true, Verdict: fmt.Sprintf("central PeV threshold is %.3g times above a 100 TeV direct-production proxy; direct production is not a near-term observable", ratio)}
}

func auditEWPO(s SealedSpectrum) ElectroweakPrecisionAudit {
	m := s.SingleScaleMBGeV
	vOverM := vevGeV / m
	deltaM := vevGeV * vevGeV / m
	c2 := 1 - sin2ThetaW
	tProxy := (deltaM * deltaM) / (12 * math.Pi * sin2ThetaW * c2 * mzGeV * mzGeV)
	return ElectroweakPrecisionAudit{
		MassScaleGeV: m, VEVOverM: vOverM, VEVOverMSquared: vOverM * vOverM, ParametricDeltaMGeV: deltaM,
		TProxy: tProxy, TProxyFormula: "proxy T ≈ (ΔM)^2/(12π s_W² c_W² M_Z²), with ΔM≈v²/M_B; this is an upper-order parametric EFT audit, not a derived finite splitting",
		TreeLevelViolationDerived: false, HeavyYukawaCouplingDerived: false, ObliqueSafe: tProxy < 1e-4,
		Verdict: "PeV decoupling makes VEV-induced oblique effects parametrically tiny unless future seals introduce large non-decoupling Higgs couplings",
	}
}

func auditHiggsLoops(s SealedSpectrum) HiggsLoopAudit {
	dec := math.Pow(vevGeV/s.SingleScaleMBGeV, 2)
	// These proxies deliberately include representation multiplicities but no
	// invented heavy Yukawa. They are order-of-magnitude decoupling diagnostics.
	diphoton := dec * float64(s.Row1Dimension+2*s.Row2Dimension)
	gluon := dec * float64(s.Row2Dimension)
	return HiggsLoopAudit{ColoredOctetPresent: s.Row2ContainsColor, ChargedTripletPresent: true, HeavyYukawaCouplingDerived: false, NonDecouplingMassFromHiggs: false, DecouplingParameter: dec, DiphotonAmplitudeProxy: diphoton, GluonFusionAmplitudeProxy: gluon, HiggsLoopSafeUnderDecoupling: diphoton < 1e-5 && gluon < 1e-5, Verdict: "standard decoupling suppresses Higgs-loop imprints by v²/M_B² because no Higgs-generated heavy mass or heavy Yukawa coupling is derived"}
}

func auditCosmology(s SealedSpectrum) CosmologyAudit {
	neutral := false
	for _, q := range s.Row1Charges {
		if q == 0 {
			neutral = true
		}
	}
	return CosmologyAudit{
		TripletHasNeutralComponent: neutral, TripletChargedComponents: []int{1, 2}, ColorOctetContainsColoredStates: s.Row2ContainsColor,
		DecayOperatorDerived: s.DecayOperatorsDerived, MassSplittingDerived: s.MassSplittingsDerived,
		StableNeutralRelicWarning: !s.DecayOperatorsDerived && neutral, StableChargedRelicWarning: !s.DecayOperatorsDerived, StableColoredRelicWarning: !s.DecayOperatorsDerived && s.Row2ContainsColor,
		DarkMatterCandidateClaimed: false, OverclosureComputed: false, RecommendedSeal: "FutureDecayOperatorOrRelicSeal",
		Verdict: "cosmology is the active warning: without a derived decay operator or splitting theorem, the neutral triplet, charged partners, and colored octet states must not be declared cosmologically safe",
	}
}

func summarize(d DirectReachAudit, e ElectroweakPrecisionAudit, h HiggsLoopAudit, c CosmologyAudit) ObservabilitySummary {
	cosmoWarning := c.StableNeutralRelicWarning || c.StableChargedRelicWarning || c.StableColoredRelicWarning
	status := StatusConditionalPhenomenology
	if cosmoWarning {
		status = StatusWarningCosmology
	}
	return ObservabilitySummary{EWPOSafe: e.ObliqueSafe, HiggsLoopSafe: h.HiggsLoopSafeUnderDecoupling, DirectReachSafe: d.DirectProductionSafe, CosmologyWarning: cosmoWarning, FatalObservableFailure: false, FutureTargets: []string{"PeV-threshold threshold effects in precision unification", "indirect Higgs/EW deviations only if future heavy-Higgs couplings are sealed", "cosmological relic/decay-channel audit for neutral, charged, and colored heavy carriers"}, Status: status, Comment: "Gate 220 finds indirect precision safety by decoupling, but logs a stable-relic warning because decay and mass-splitting semantics are not derived."}
}

func auditFirewall(g219 inputsensitivityaudit.Analysis, snap Gate219Snapshot) FirewallAudit {
	return FirewallAudit{Gate219Inherited: snap.Gate219Inherited, ThresholdSpectrumSealActive: snap.ThresholdSpectrumSealActive, MatchingCorrectionSealActive: snap.MatchingCorrectionSealActive, EmpiricalCarrierSealInherited: g219.Firewall.EmpiricalCarrierSealInherited, LeptoquarkDynamicsSealInherited: g219.Firewall.LeptoquarkDynamicsSealInherited, PeVMassFiniteDerived: false, DecayOperatorInvented: false, HeavyHiggsYukawaInvented: false, MassSplittingInvented: false, DarkMatterClaimed: false, OverclosureComputed: false, PhysicalObservationClaimed: false, RecommendedNextGate: "Gate 221 — sealed decay-operator and cosmological relic safety audit", OpenRequirements: []string{"derive or seal heavy carrier decay operators", "derive or seal charged/neutral mass splittings", "compute relic abundance only after dynamics and decay operators exist", "audit higher-dimensional operators connecting the PeV sector to SM fields"}, Verdict: "observability is conditional phenomenology; no finite-derived mass, decay, Higgs coupling, or dark-matter claim is introduced"}
}

func buildTruth(s SealedSpectrum, d DirectReachAudit, e ElectroweakPrecisionAudit, h HiggsLoopAudit, c CosmologyAudit, sum ObservabilitySummary) string {
	return fmt.Sprintf("Gate 220 audits the sealed PeV spectrum at M_B=%.9g GeV. Direct production is separated by a factor %.4g from a 100 TeV proxy. EWPO proxy T=%.6g from ΔM≈v²/M_B=%.6g GeV, and Higgs-loop decoupling proxy v²/M_B²=%.6g, so precision/Higgs loop effects are parametrically safe under current seals. Cosmology remains warning=%t because no decay operator or mass-splitting theorem is derived for the neutral, charged, and colored heavy states. Status=%s.", s.SingleScaleMBGeV, d.MassOverReach, e.TProxy, e.ParametricDeltaMGeV, h.DecouplingParameter, sum.CosmologyWarning, sum.Status)
}

func FormatGate219(g Gate219Snapshot) string {
	return fmt.Sprintf("gate219=%t spectrumSeal=%t matchSeal=%t plausible=%t bottomTau=%t MB=%.9g [%.9g,%.9g] M*=%.9g [%.9g,%.9g] worstδ/ε=%.6g", g.Gate219Inherited, g.ThresholdSpectrumSealActive, g.MatchingCorrectionSealActive, g.MatchingPlausible, g.BottomTauComplete, g.MBGeV, g.MBMinGeV, g.MBMaxGeV, g.MStarGeV, g.MStarMinGeV, g.MStarMaxGeV, g.WorstResidualOverEpsilon)
}

func FormatSpectrum(s SealedSpectrum) string {
	return fmt.Sprintf("%s rows=[%s %s dim=%d Q=%v; %s %s dim=%d color=%t weak=%t] MB=%.9g range=[%.9g,%.9g] conditional=%t decayDerived=%t splitDerived=%t", s.Name, s.Row1Name, s.Row1Rep, s.Row1Dimension, s.Row1Charges, s.Row2Name, s.Row2Rep, s.Row2Dimension, s.Row2ContainsColor, s.Row2ContainsWeak, s.SingleScaleMBGeV, s.ScaleRangeMinGeV, s.ScaleRangeMaxGeV, s.ConditionalOnly, s.DecayOperatorsDerived, s.MassSplittingsDerived)
}

func FormatDirectReach(d DirectReachAudit) string {
	return fmt.Sprintf("reach=%.9g GeV mass=%.9g GeV mass/reach=%.6g safe=%t :: %s", d.DirectProductionReachGeV, d.CentralMassGeV, d.MassOverReach, d.DirectProductionSafe, d.Verdict)
}

func FormatEWPO(e ElectroweakPrecisionAudit) string {
	return fmt.Sprintf("M=%.9g v/M=%.9g (v/M)^2=%.9g ΔMproxy=%.9g GeV Tproxy=%.9g safe=%t treeDerived=%t heavyYukawa=%t formula=%q", e.MassScaleGeV, e.VEVOverM, e.VEVOverMSquared, e.ParametricDeltaMGeV, e.TProxy, e.ObliqueSafe, e.TreeLevelViolationDerived, e.HeavyYukawaCouplingDerived, e.TProxyFormula)
}

func FormatHiggs(h HiggsLoopAudit) string {
	return fmt.Sprintf("colored=%t chargedTriplet=%t heavyYukawa=%t nonDecoupling=%t v2/M2=%.9g hγγproxy=%.9g ggHproxy=%.9g safe=%t :: %s", h.ColoredOctetPresent, h.ChargedTripletPresent, h.HeavyYukawaCouplingDerived, h.NonDecouplingMassFromHiggs, h.DecouplingParameter, h.DiphotonAmplitudeProxy, h.GluonFusionAmplitudeProxy, h.HiggsLoopSafeUnderDecoupling, h.Verdict)
}

func FormatCosmology(c CosmologyAudit) string {
	return fmt.Sprintf("neutral=%t charged=%v colored=%t decayDerived=%t splitDerived=%t warnings=(neutral %t,charged %t,colored %t) DMclaim=%t overclosure=%t next=%s :: %s", c.TripletHasNeutralComponent, c.TripletChargedComponents, c.ColorOctetContainsColoredStates, c.DecayOperatorDerived, c.MassSplittingDerived, c.StableNeutralRelicWarning, c.StableChargedRelicWarning, c.StableColoredRelicWarning, c.DarkMatterCandidateClaimed, c.OverclosureComputed, c.RecommendedSeal, c.Verdict)
}

func FormatSummary(s ObservabilitySummary) string {
	return fmt.Sprintf("directSafe=%t EWPOsafe=%t HiggsSafe=%t cosmoWarning=%t fatal=%t status=%s targets=[%s]", s.DirectReachSafe, s.EWPOSafe, s.HiggsLoopSafe, s.CosmologyWarning, s.FatalObservableFailure, s.Status, strings.Join(s.FutureTargets, "; "))
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate219=%t spectrumSeal=%t matchSeal=%t carrier=%t lq=%t massFinite=%t decayInvented=%t heavyYukawaInvented=%t splitInvented=%t DM=%t overclosure=%t observation=%t next=%q", f.Gate219Inherited, f.ThresholdSpectrumSealActive, f.MatchingCorrectionSealActive, f.EmpiricalCarrierSealInherited, f.LeptoquarkDynamicsSealInherited, f.PeVMassFiniteDerived, f.DecayOperatorInvented, f.HeavyHiggsYukawaInvented, f.MassSplittingInvented, f.DarkMatterClaimed, f.OverclosureComputed, f.PhysicalObservationClaimed, f.RecommendedNextGate)
}
