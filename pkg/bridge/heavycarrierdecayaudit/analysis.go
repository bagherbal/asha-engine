// Package heavycarrierdecayaudit implements Gate 221: heavy-carrier decay
// and mass-splitting / cosmological-relic safety audit.
//
// Gate 220 found that the sealed PeV spectrum is safe against direct-reach,
// electroweak-precision, and Higgs-loop probes by decoupling, but it logged a
// serious stable-relic warning. Gate 221 asks the next physical question:
// can the engine legally make the heavy carriers decay, or split their charged
// and neutral components, without inventing new couplings?
//
// The gate deliberately does not compute relic abundance from an absent
// Lagrangian. It audits the operator basis, mass-splitting semantics, and BBN
// lifetime threshold. If no decay operator exists, the lifetime is classified
// as unbounded/infinite for safety purposes and the cosmology route fails until
// a future finite theorem or explicit RelicDecaySeal supplies the missing data.
package heavycarrierdecayaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/pevobservabilityaudit"
)

const (
	StatusFailedCosmologicalPathology = "FAILED_ROUTE_COSMOLOGICAL_PATHOLOGY"
	StatusRelicDecaySealRequired      = "RELIC_DECAY_SEAL_REQUIRED_NOT_GRANTED"

	DecayAuditID = "GATE221-HEAVY-CARRIER-DECAY-RELIC-SAFETY-AUDIT"
	hbarGeVS     = 6.582119569e-25
	bbnSeconds   = 1.0
	vevGeV       = 246.21965
	pionMassGeV  = 0.13957039
)

type Gate220Snapshot struct {
	Gate220Inherited             bool
	DirectReachSafe              bool
	EWPOSafe                     bool
	HiggsLoopSafe                bool
	CosmologyWarning             bool
	StableNeutralRelicWarning    bool
	StableChargedRelicWarning    bool
	StableColoredRelicWarning    bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	DecayOperatorsDerived        bool
	MassSplittingsDerived        bool
	MBGeV                        float64
	MMinGeV                      float64
	MMaxGeV                      float64
	TruthStatement               string
}

type Carrier struct {
	Name                string
	Representation      string
	Dimension           int
	ColorCharged        bool
	WeakCharged         bool
	Hypercharge         string
	ElectricCharges     []int
	HasNeutralComponent bool
	HasChargedComponent bool
	HasColoredComponent bool
	ConditionalOnly     bool
}

type OperatorCandidate struct {
	Name                    string
	Dimension               int
	TargetCarrier           string
	SymbolicForm            string
	GaugeInvariant          bool
	LorentzStructureDerived bool
	FiniteAlgebraSupported  bool
	RequiresNewCoupling     bool
	RequiresNewLightField   bool
	RespectsLeptoquarkSeal  bool
	DecayWidthComputable    bool
	BBNLifetimeComputable   bool
	Verdict                 string
}

type OperatorBasisAudit struct {
	Candidates                []OperatorCandidate
	CandidatesAudited         int
	GaugeInvariantCandidates  int
	FiniteSupportedCandidates int
	DecayOperatorsDerived     int
	AnyDecayWidthComputable   bool
	PortalOperatorFound       bool
	NativePortalSearchFailed  bool
	Verdict                   string
}

type MassSplittingAudit struct {
	TreeDegenerateBySeal            bool
	ElectroweakLoopSplittingDerived bool
	VEVCouplingSplittingDerived     bool
	ProxyVEVSquaredOverMGeV         float64
	ProxyAbovePionThreshold         bool
	ChargedToNeutralCascadeDerived  bool
	ColorHadronizationDerived       bool
	StableChargedRisk               bool
	StableColoredRisk               bool
	Verdict                         string
}

type LifetimeAudit struct {
	BBNThresholdSeconds       float64
	RequiredWidthGeV          float64
	DecayWidthDerived         bool
	DecayWidthGeV             float64
	LifetimeSeconds           float64
	LifetimeFinite            bool
	PassesBBN                 bool
	FailsBBNByOperatorAbsence bool
	Dimension5ToyWidthNotUsed bool
	Dimension6ToyWidthNotUsed bool
	Verdict                   string
}

type RelicDecaySealAudit struct {
	SealName          string
	SealGranted       bool
	SealRequired      bool
	ReasonNotGranted  string
	WouldNeedToSupply []string
	OperationalStatus string
	Verdict           string
}

type CosmologicalSafetySummary struct {
	NeutralRelicSafe bool
	ChargedRelicSafe bool
	ColoredRelicSafe bool
	BBNSafe          bool
	CosmologyCleared bool
	FatalPathology   bool
	Status           string
	NextGate         string
	Comment          string
}

type FirewallAudit struct {
	Gate220Inherited             bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	DecayOperatorInvented        bool
	MassSplittingInvented        bool
	LifetimeComputedFromAbsentOp bool
	RelicAbundanceComputed       bool
	DarkMatterClaimed            bool
	ArbitraryCouplingIntroduced  bool
	BBNUsedAsFilterOnly          bool
	PeVMassFiniteDerived         bool
	RecommendedNextGate          string
	RemainingUnknowns            []string
	Verdict                      string
}

type Analysis struct {
	Gate220         Gate220Snapshot
	Gate220Analysis pevobservabilityaudit.Analysis
	Carriers        []Carrier
	Operators       OperatorBasisAudit
	MassSplitting   MassSplittingAudit
	Lifetime        LifetimeAudit
	RelicSeal       RelicDecaySealAudit
	Summary         CosmologicalSafetySummary
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
		g220, err := pevobservabilityaudit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 220 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g220)
	})
	return defaultA, defaultErr
}

func Build(g220 pevobservabilityaudit.Analysis) (Analysis, error) {
	snap := snapshotFromGate220(g220)
	if !snap.Gate220Inherited || !snap.CosmologyWarning {
		return Analysis{}, fmt.Errorf("Gate 221 requires Gate 220 stable-relic warning")
	}
	carriers := buildCarriers()
	opAudit := auditOperatorBasis(carriers)
	split := auditMassSplitting(snap, carriers)
	life := auditLifetime(opAudit)
	seal := auditRelicDecaySeal(opAudit, split, life)
	summary := summarize(opAudit, split, life, seal)
	firewall := auditFirewall(snap)
	truth := buildTruth(snap, opAudit, split, life, seal, summary)
	return Analysis{Gate220: snap, Gate220Analysis: g220, Carriers: carriers, Operators: opAudit, MassSplitting: split, Lifetime: life, RelicSeal: seal, Summary: summary, Firewall: firewall, TruthStatement: truth}, nil
}

func snapshotFromGate220(a pevobservabilityaudit.Analysis) Gate220Snapshot {
	return Gate220Snapshot{
		Gate220Inherited:             a.Summary.Status == pevobservabilityaudit.StatusWarningCosmology,
		DirectReachSafe:              a.Summary.DirectReachSafe,
		EWPOSafe:                     a.Summary.EWPOSafe,
		HiggsLoopSafe:                a.Summary.HiggsLoopSafe,
		CosmologyWarning:             a.Summary.CosmologyWarning,
		StableNeutralRelicWarning:    a.Cosmology.StableNeutralRelicWarning,
		StableChargedRelicWarning:    a.Cosmology.StableChargedRelicWarning,
		StableColoredRelicWarning:    a.Cosmology.StableColoredRelicWarning,
		ThresholdSpectrumSealActive:  a.Firewall.ThresholdSpectrumSealActive,
		MatchingCorrectionSealActive: a.Firewall.MatchingCorrectionSealActive,
		EmpiricalCarrierSealActive:   a.Firewall.EmpiricalCarrierSealInherited,
		LeptoquarkDynamicsSealActive: a.Firewall.LeptoquarkDynamicsSealInherited,
		DecayOperatorsDerived:        a.Spectrum.DecayOperatorsDerived,
		MassSplittingsDerived:        a.Spectrum.MassSplittingsDerived,
		MBGeV:                        a.Spectrum.SingleScaleMBGeV,
		MMinGeV:                      a.Spectrum.ScaleRangeMinGeV,
		MMaxGeV:                      a.Spectrum.ScaleRangeMaxGeV,
		TruthStatement:               a.TruthStatement,
	}
}

func buildCarriers() []Carrier {
	return []Carrier{
		{Name: "Dirac electroweak triplet", Representation: "(1,3,Y=1)", Dimension: 3, ColorCharged: false, WeakCharged: true, Hypercharge: "1", ElectricCharges: []int{0, 1, 2}, HasNeutralComponent: true, HasChargedComponent: true, HasColoredComponent: false, ConditionalOnly: true},
		{Name: "Dirac color-octet weak doublet", Representation: "(8,2,Y=1/2)", Dimension: 16, ColorCharged: true, WeakCharged: true, Hypercharge: "1/2", ElectricCharges: []int{0, 1}, HasNeutralComponent: true, HasChargedComponent: true, HasColoredComponent: true, ConditionalOnly: true},
	}
}

func auditOperatorBasis(carriers []Carrier) OperatorBasisAudit {
	candidates := []OperatorCandidate{
		{Name: "renormalizable triplet-lepton-Higgs portal", Dimension: 4, TargetCarrier: "(1,3,Y=1)", SymbolicForm: "Psi_3 L H or Psi_3 L H†", GaugeInvariant: false, LorentzStructureDerived: false, FiniteAlgebraSupported: false, RequiresNewCoupling: true, RequiresNewLightField: false, RespectsLeptoquarkSeal: true, Verdict: "hypercharge mismatch for the sealed Y=1 fermion triplet; no finite Yukawa portal is derived"},
		{Name: "dimension-five triplet Higgs-lepton portal", Dimension: 5, TargetCarrier: "(1,3,Y=1)", SymbolicForm: "Psi_3 L H H / Λ variants", GaugeInvariant: false, LorentzStructureDerived: false, FiniteAlgebraSupported: false, RequiresNewCoupling: true, RequiresNewLightField: false, RespectsLeptoquarkSeal: true, Verdict: "no canonical gauge-invariant finite operator with coefficient and Lorentz contraction is supplied by the engine"},
		{Name: "octet-doublet quark-Higgs portal", Dimension: 4, TargetCarrier: "(8,2,Y=1/2)", SymbolicForm: "Psi_8 Q H plus color-octet contraction", GaugeInvariant: false, LorentzStructureDerived: false, FiniteAlgebraSupported: false, RequiresNewCoupling: true, RequiresNewLightField: true, RespectsLeptoquarkSeal: true, Verdict: "SM quarks are color triplets; a color-octet fermion cannot mix with them without an additional colored operator not derived by ASHA"},
		{Name: "dimension-six neutral-current decay portal", Dimension: 6, TargetCarrier: "both carriers", SymbolicForm: "Psi_heavy ψ_SM ψ_SM H / Λ² templates", GaugeInvariant: false, LorentzStructureDerived: false, FiniteAlgebraSupported: false, RequiresNewCoupling: true, RequiresNewLightField: false, RespectsLeptoquarkSeal: true, Verdict: "would be an external EFT coefficient; no finite current action, local field map, or suppression scale is derived"},
		{Name: "leptoquark-mediated colored-carrier decay", Dimension: 6, TargetCarrier: "(8,2,Y=1/2)", SymbolicForm: "Psi_8 q l q / Λ² through dormant u(4) slots", GaugeInvariant: true, LorentzStructureDerived: false, FiniteAlgebraSupported: false, RequiresNewCoupling: true, RequiresNewLightField: false, RespectsLeptoquarkSeal: false, Verdict: "blocked by the LeptoquarkDynamicsSeal; dormant u(4) slots cannot be used as propagators or coefficients"},
	}
	gaugeInv := 0
	finite := 0
	derived := 0
	computable := false
	for i := range candidates {
		if candidates[i].GaugeInvariant {
			gaugeInv++
		}
		if candidates[i].FiniteAlgebraSupported {
			finite++
		}
		if candidates[i].FiniteAlgebraSupported && candidates[i].LorentzStructureDerived && !candidates[i].RequiresNewCoupling && candidates[i].RespectsLeptoquarkSeal {
			candidates[i].DecayWidthComputable = true
			candidates[i].BBNLifetimeComputable = true
			derived++
			computable = true
		}
	}
	return OperatorBasisAudit{Candidates: candidates, CandidatesAudited: len(candidates), GaugeInvariantCandidates: gaugeInv, FiniteSupportedCandidates: finite, DecayOperatorsDerived: derived, AnyDecayWidthComputable: computable, PortalOperatorFound: derived > 0, NativePortalSearchFailed: derived == 0, Verdict: "no native decay portal is derived; all apparent portals either fail gauge/field semantics, require an external coupling/suppression scale, or violate the LeptoquarkDynamicsSeal"}
}

func auditMassSplitting(snap Gate220Snapshot, carriers []Carrier) MassSplittingAudit {
	proxy := vevGeV * vevGeV / snap.MBGeV
	return MassSplittingAudit{TreeDegenerateBySeal: !snap.MassSplittingsDerived, ElectroweakLoopSplittingDerived: false, VEVCouplingSplittingDerived: false, ProxyVEVSquaredOverMGeV: proxy, ProxyAbovePionThreshold: proxy > pionMassGeV, ChargedToNeutralCascadeDerived: false, ColorHadronizationDerived: false, StableChargedRisk: true, StableColoredRisk: true, Verdict: fmt.Sprintf("no charged/neutral or colored-hadron splitting theorem is derived; the v²/M proxy is %.6g GeV and is diagnostic only, not a decay theorem", proxy)}
}

func auditLifetime(op OperatorBasisAudit) LifetimeAudit {
	required := hbarGeVS / bbnSeconds
	if op.AnyDecayWidthComputable {
		// No current branch reaches this. Kept for theorem completeness.
		width := required * 10
		life := hbarGeVS / width
		return LifetimeAudit{BBNThresholdSeconds: bbnSeconds, RequiredWidthGeV: required, DecayWidthDerived: true, DecayWidthGeV: width, LifetimeSeconds: life, LifetimeFinite: true, PassesBBN: life < bbnSeconds, Verdict: "derived decay width passes BBN"}
	}
	return LifetimeAudit{BBNThresholdSeconds: bbnSeconds, RequiredWidthGeV: required, DecayWidthDerived: false, DecayWidthGeV: 0, LifetimeSeconds: math.Inf(1), LifetimeFinite: false, PassesBBN: false, FailsBBNByOperatorAbsence: true, Dimension5ToyWidthNotUsed: true, Dimension6ToyWidthNotUsed: true, Verdict: "no decay width is legal because no decay operator is derived; for cosmological safety this is treated as an infinite-lifetime BBN failure route"}
}

func auditRelicDecaySeal(op OperatorBasisAudit, split MassSplittingAudit, life LifetimeAudit) RelicDecaySealAudit {
	granted := op.PortalOperatorFound && life.PassesBBN && (split.ChargedToNeutralCascadeDerived || !split.StableChargedRisk) && !split.StableColoredRisk
	return RelicDecaySealAudit{SealName: "RelicDecaySeal", SealGranted: granted, SealRequired: !granted, ReasonNotGranted: "no finite decay operator, no mass-splitting/cascade theorem, and no computable width below the BBN lifetime threshold", WouldNeedToSupply: []string{"gauge-invariant heavy-to-SM operator basis", "Lorentz/local-field contraction", "finite or sealed coupling coefficient", "suppression scale", "charged/neutral mass splitting", "colored-state hadronization/decay channel", "width Γ with τ < 1 second"}, OperationalStatus: StatusRelicDecaySealRequired, Verdict: "RelicDecaySeal is required for future phenomenology but is not granted by Gate 221"}
}

func summarize(op OperatorBasisAudit, split MassSplittingAudit, life LifetimeAudit, seal RelicDecaySealAudit) CosmologicalSafetySummary {
	fatal := op.NativePortalSearchFailed && life.FailsBBNByOperatorAbsence && !seal.SealGranted
	return CosmologicalSafetySummary{NeutralRelicSafe: false, ChargedRelicSafe: false, ColoredRelicSafe: false, BBNSafe: life.PassesBBN, CosmologyCleared: false, FatalPathology: fatal, Status: StatusFailedCosmologicalPathology, NextGate: "Gate 222 — finite or sealed relic-decay portal construction / cosmological rescue audit", Comment: "Gate 221 fails cosmological safety: PeV carriers are precision-safe but cannot be declared relic-safe without a decay portal, splitting theorem, and BBN-safe lifetime."}
}

func auditFirewall(snap Gate220Snapshot) FirewallAudit {
	return FirewallAudit{Gate220Inherited: snap.Gate220Inherited, ThresholdSpectrumSealActive: snap.ThresholdSpectrumSealActive, MatchingCorrectionSealActive: snap.MatchingCorrectionSealActive, EmpiricalCarrierSealActive: snap.EmpiricalCarrierSealActive, LeptoquarkDynamicsSealActive: snap.LeptoquarkDynamicsSealActive, DecayOperatorInvented: false, MassSplittingInvented: false, LifetimeComputedFromAbsentOp: false, RelicAbundanceComputed: false, DarkMatterClaimed: false, ArbitraryCouplingIntroduced: false, BBNUsedAsFilterOnly: true, PeVMassFiniteDerived: false, RecommendedNextGate: "Gate 222 — finite or sealed relic-decay portal construction / cosmological rescue audit", RemainingUnknowns: []string{"heavy-carrier decay operator", "heavy-light coupling coefficient", "suppression scale", "charged-neutral splitting", "colored-state decay or confinement history", "Boltzmann/relic abundance after a legal decay channel exists"}, Verdict: "Gate 221 uses BBN only as a safety filter and refuses to compute a lifetime, relic density, or dark-matter claim from absent dynamics"}
}

func buildTruth(snap Gate220Snapshot, op OperatorBasisAudit, split MassSplittingAudit, life LifetimeAudit, seal RelicDecaySealAudit, sum CosmologicalSafetySummary) string {
	return fmt.Sprintf("Gate 221 inherits the Gate-220 PeV spectrum at M_B=%.9g GeV and audits whether its neutral, charged, and colored carriers can decay. It finds %d candidate portal classes, but zero finite-supported decay operators and zero computable widths. The diagnostic splitting proxy v²/M_B=%.6g GeV is not a theorem. Because no Γ is legal, the BBN τ<%.3g s requirement fails by operator absence. RelicDecaySeal is required but not granted. Status=%s.", snap.MBGeV, op.CandidatesAudited, split.ProxyVEVSquaredOverMGeV, life.BBNThresholdSeconds, sum.Status)
}

func FormatGate220(g Gate220Snapshot) string {
	return fmt.Sprintf("gate220=%t direct=%t EWPO=%t Higgs=%t cosmoWarning=%t warnings=(N %t,C %t,color %t) seals=(spectrum %t,matching %t,carrier %t,lq %t) decayDerived=%t splitDerived=%t MB=%.9g [%.9g,%.9g]", g.Gate220Inherited, g.DirectReachSafe, g.EWPOSafe, g.HiggsLoopSafe, g.CosmologyWarning, g.StableNeutralRelicWarning, g.StableChargedRelicWarning, g.StableColoredRelicWarning, g.ThresholdSpectrumSealActive, g.MatchingCorrectionSealActive, g.EmpiricalCarrierSealActive, g.LeptoquarkDynamicsSealActive, g.DecayOperatorsDerived, g.MassSplittingsDerived, g.MBGeV, g.MMinGeV, g.MMaxGeV)
}

func FormatCarrier(c Carrier) string {
	return fmt.Sprintf("%s %s dim=%d color=%t weak=%t Y=%s Q=%v neutral=%t charged=%t colored=%t conditional=%t", c.Name, c.Representation, c.Dimension, c.ColorCharged, c.WeakCharged, c.Hypercharge, c.ElectricCharges, c.HasNeutralComponent, c.HasChargedComponent, c.HasColoredComponent, c.ConditionalOnly)
}

func FormatOperatorCandidate(c OperatorCandidate) string {
	return fmt.Sprintf("%s dim=%d target=%s form=%q gauge=%t lorentz=%t finite=%t newCoupling=%t newLight=%t lqSealOK=%t width=%t :: %s", c.Name, c.Dimension, c.TargetCarrier, c.SymbolicForm, c.GaugeInvariant, c.LorentzStructureDerived, c.FiniteAlgebraSupported, c.RequiresNewCoupling, c.RequiresNewLightField, c.RespectsLeptoquarkSeal, c.DecayWidthComputable, c.Verdict)
}

func FormatOperatorAudit(a OperatorBasisAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, c.Name+":"+c.Verdict)
	}
	return fmt.Sprintf("audited=%d gaugeInvariant=%d finiteSupported=%d derived=%d widthComputable=%t portal=%t failed=%t :: %s", a.CandidatesAudited, a.GaugeInvariantCandidates, a.FiniteSupportedCandidates, a.DecayOperatorsDerived, a.AnyDecayWidthComputable, a.PortalOperatorFound, a.NativePortalSearchFailed, strings.Join(parts, " | "))
}

func FormatMassSplitting(a MassSplittingAudit) string {
	return fmt.Sprintf("treeDegenerate=%t ewLoopDerived=%t vevDerived=%t v2/M=%.9g GeV pionThreshold=%.9g abovePion=%t chargedCascade=%t colorDecay=%t risks=(charged %t,colored %t) :: %s", a.TreeDegenerateBySeal, a.ElectroweakLoopSplittingDerived, a.VEVCouplingSplittingDerived, a.ProxyVEVSquaredOverMGeV, pionMassGeV, a.ProxyAbovePionThreshold, a.ChargedToNeutralCascadeDerived, a.ColorHadronizationDerived, a.StableChargedRisk, a.StableColoredRisk, a.Verdict)
}

func FormatLifetime(a LifetimeAudit) string {
	life := "Inf"
	if a.LifetimeFinite {
		life = fmt.Sprintf("%.9g", a.LifetimeSeconds)
	}
	return fmt.Sprintf("BBN<%.9g s requires Γ>%.9g GeV derivedWidth=%t Γ=%.9g τ=%s finite=%t pass=%t failByAbsence=%t toyD5unused=%t toyD6unused=%t :: %s", a.BBNThresholdSeconds, a.RequiredWidthGeV, a.DecayWidthDerived, a.DecayWidthGeV, life, a.LifetimeFinite, a.PassesBBN, a.FailsBBNByOperatorAbsence, a.Dimension5ToyWidthNotUsed, a.Dimension6ToyWidthNotUsed, a.Verdict)
}

func FormatRelicSeal(a RelicDecaySealAudit) string {
	return fmt.Sprintf("seal=%s granted=%t required=%t status=%s reason=%q needs=[%s]", a.SealName, a.SealGranted, a.SealRequired, a.OperationalStatus, a.ReasonNotGranted, strings.Join(a.WouldNeedToSupply, "; "))
}

func FormatSummary(a CosmologicalSafetySummary) string {
	return fmt.Sprintf("safe=(neutral %t,charged %t,colored %t,BBN %t) cleared=%t fatal=%t status=%s next=%q :: %s", a.NeutralRelicSafe, a.ChargedRelicSafe, a.ColoredRelicSafe, a.BBNSafe, a.CosmologyCleared, a.FatalPathology, a.Status, a.NextGate, a.Comment)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate220=%t seals=(spectrum %t,matching %t,carrier %t,lq %t) invented=(decay %t,split %t,coupling %t) lifetimeAbsent=%t relic=%t DM=%t BBNfilter=%t PeVFinite=%t next=%q", f.Gate220Inherited, f.ThresholdSpectrumSealActive, f.MatchingCorrectionSealActive, f.EmpiricalCarrierSealActive, f.LeptoquarkDynamicsSealActive, f.DecayOperatorInvented, f.MassSplittingInvented, f.ArbitraryCouplingIntroduced, f.LifetimeComputedFromAbsentOp, f.RelicAbundanceComputed, f.DarkMatterClaimed, f.BBNUsedAsFilterOnly, f.PeVMassFiniteDerived, f.RecommendedNextGate)
}
