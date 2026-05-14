// Package finiteanchordm implements Gate 225: finite anchor Dark Matter
// viability / Axion-like particle (ALP) and Dark Sector audit.
//
// Gate 224 proved that the PeV threshold carriers decay before BBN and
// therefore contribute zero present-day heavy-sector dark matter. Gate 225
// returns to the finite anchors that were deliberately left unassigned: the
// B-sector first spectral gap and the seven contact partial-overlap modes. It
// asks whether these anchors already provide the minimal structures required
// for axion-like or sequestered dark-sector dark matter.
//
// The answer is currently no. The finite core supplies dimensionless spectral
// anchors, but not a continuous shift symmetry, not a Pontryagin/anomaly
// coupling, not a stable dark-sector action, and not a dimensionful scale such
// as f_a. The gate therefore logs a strict obstruction rather than inventing an
// axion decay constant or relic abundance.
package finiteanchordm

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/flavoralignmentdmabsence"
)

const (
	AuditID = "GATE225-FINITE-ANCHOR-DARK-MATTER-VIABILITY-AUDIT"

	StatusFailedFiniteAnchorDM       = "FAILED_ROUTE_FINITE_ANCHOR_DARK_MATTER_DERIVATION"
	StatusALPObstruction             = "FAILED_ROUTE_ALP_SHIFT_ANOMALY_SCALE_OBSTRUCTION"
	StatusDarkSectorObstruction      = "FAILED_ROUTE_CONTACT_DARK_SECTOR_STABILITY_OBSTRUCTION"
	StatusHeavyDMAbsenceStillBinding = "HEAVY_SECTOR_DM_ABSENCE_REMAINS_BINDING"
)

const (
	bSectorFirstGap       = 0.1024649212
	contactPartialModeN   = 7
	loopFactorInv16Pi2    = 1.0 / (16.0 * math.Pi * math.Pi)
	observedOmegaDMH2Hint = 0.12
)

type Gate224Snapshot struct {
	Gate224Inherited             bool
	HeavySectorDMAbsent          bool
	RelicDecaySealActive         bool
	FlavorAlignmentSealActive    bool
	ThresholdSpectrumSealActive  bool
	MatchingCorrectionSealActive bool
	EmpiricalCarrierSealActive   bool
	LeptoquarkDynamicsSealActive bool
	OmegaHeavySectorH2           float64
	TruthStatement               string
}

type FiniteAnchorInventory struct {
	BGapValue                   float64
	BGapDimensionless           bool
	BGapPhysicalMassDerived     bool
	BGapPeriodicCoordinate      bool
	ContactPartialModeCount     int
	ContactModesDimensionless   bool
	ContactModesPositiveAnchors bool
	ContactModesSMChargeDerived bool
	ContactModesSingletProved   bool
	ContactModesStableProved    bool
	ContactModesActionDerived   bool
	FiniteDarkMatterCandidate   bool
	InventoryVerdict            string
}

type ALPTopologicalAudit struct {
	AnchorName                   string
	CandidateScalar              float64
	GlobalShiftSymmetryDerived   bool
	CompactPeriodicFieldDerived  bool
	AxionDecayConstantDerived    bool
	InstantonPotentialDerived    bool
	PontryaginCouplingDerived    bool
	GaugeAnomalyCoefficientRows  int
	AnomalyVector                []float64
	QCDThetaRelaxationDerived    bool
	GenericALPStructureSupported bool
	QCDAxionStructureSupported   bool
	CanonicalNormalizationFound  bool
	ArbitraryCoefficientInserted bool
	Verdict                      string
}

type ContactDarkSectorAudit struct {
	ModeCount                  int
	SequesteredByNonPromotion  bool
	GaugeSingletTheoremDerived bool
	StabilitySymmetryDerived   bool
	DarkActionDerived          bool
	SelfInteractionDerived     bool
	MassScaleDerived           bool
	ThermalHistoryDerived      bool
	RelicAbundanceComputed     bool
	StrictDarkSectorSupported  bool
	CompatibleFutureRoute      bool
	Verdict                    string
}

type MisalignmentPreflight struct {
	Formula                     string
	RequiresAxionMass           bool
	RequiresDecayConstant       bool
	RequiresInitialAngle        bool
	RequiresCosmologicalHistory bool
	BGapCanDimensionalizeMass   bool
	NativeFAFound               bool
	NativeMassFound             bool
	OmegaComputed               bool
	OmegaValueH2                float64
	TargetOmegaH2               float64
	MisalignmentViable          bool
	Verdict                     string
}

type RelicDensityAccounting struct {
	HeavySectorOmegaH2          float64
	FiniteAnchorOmegaH2Computed bool
	FiniteAnchorOmegaH2         float64
	TotalModelOmegaComputed     bool
	ObservedOmegaDMH2Reference  float64
	DarkMatterStillOpen         bool
	DMDeferredTo                []string
	Verdict                     string
}

type FirewallAudit struct {
	Gate224Inherited                bool
	HeavySectorDMAbsencePreserved   bool
	RelicDecaySealPreserved         bool
	FlavorAlignmentSealPreserved    bool
	BGapUsedAsPhysicalMass          bool
	BGapUsedAsAxionScale            bool
	ContactModesPromotedToParticles bool
	ShiftSymmetryInvented           bool
	PontryaginCouplingInvented      bool
	AxionDecayConstantInvented      bool
	RelicDensityInvented            bool
	ObservedOmegaUsedForDerivation  bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	ALPSupported            bool
	DarkSectorSupported     bool
	MisalignmentComputed    bool
	HeavyDMAbsencePreserved bool
	FiniteAnchorDMViable    bool
	Status                  string
	NextGate                string
	Comment                 string
}

type Analysis struct {
	Gate224        Gate224Snapshot
	Inventory      FiniteAnchorInventory
	ALP            ALPTopologicalAudit
	Contact        ContactDarkSectorAudit
	Misalign       MisalignmentPreflight
	Relic          RelicDensityAccounting
	Firewall       FirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g224, err := flavoralignmentdmabsence.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 224 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g224)
	})
	return defaultA, defaultErr
}

func Build(g224 flavoralignmentdmabsence.Analysis) (Analysis, error) {
	snap := snapshotFromGate224(g224)
	if !snap.Gate224Inherited || !snap.HeavySectorDMAbsent {
		return Analysis{}, fmt.Errorf("Gate 225 requires Gate 224 heavy-sector dark-matter absence theorem")
	}
	inventory := auditInventory()
	alp := auditALP(inventory)
	contact := auditContactDarkSector(inventory)
	misalign := auditMisalignment(inventory, alp)
	relic := auditRelicAccounting(snap, alp, contact, misalign)
	firewall := auditFirewall(snap, inventory, alp, contact, misalign, relic)
	summary := summarize(snap, alp, contact, misalign)
	truth := buildTruth(snap, inventory, alp, contact, misalign, relic, summary)
	return Analysis{Gate224: snap, Inventory: inventory, ALP: alp, Contact: contact, Misalign: misalign, Relic: relic, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate224(a flavoralignmentdmabsence.Analysis) Gate224Snapshot {
	return Gate224Snapshot{
		Gate224Inherited:             a.Summary.FlavorAlignmentSealGranted && a.Summary.HeavySectorDMAbsent && a.Summary.RelicDecaySealStillValid,
		HeavySectorDMAbsent:          a.DarkMatter.OmegaHeavySectorH2 == 0 && !a.DarkMatter.HeavySectorDMCandidate,
		RelicDecaySealActive:         a.Firewall.RelicDecaySealActive,
		FlavorAlignmentSealActive:    a.DarkMatter.FlavorAlignmentSealActive,
		ThresholdSpectrumSealActive:  a.Firewall.ThresholdSpectrumSealActive,
		MatchingCorrectionSealActive: a.Firewall.MatchingCorrectionSealActive,
		EmpiricalCarrierSealActive:   a.Firewall.EmpiricalCarrierSealActive,
		LeptoquarkDynamicsSealActive: a.Firewall.LeptoquarkDynamicsSealActive,
		OmegaHeavySectorH2:           a.DarkMatter.OmegaHeavySectorH2,
		TruthStatement:               a.TruthStatement,
	}
}

func auditInventory() FiniteAnchorInventory {
	return FiniteAnchorInventory{
		BGapValue:                   bSectorFirstGap,
		BGapDimensionless:           true,
		BGapPhysicalMassDerived:     false,
		BGapPeriodicCoordinate:      false,
		ContactPartialModeCount:     contactPartialModeN,
		ContactModesDimensionless:   true,
		ContactModesPositiveAnchors: true,
		ContactModesSMChargeDerived: false,
		ContactModesSingletProved:   false,
		ContactModesStableProved:    false,
		ContactModesActionDerived:   false,
		FiniteDarkMatterCandidate:   false,
		InventoryVerdict:            "finite anchors exist, but none carries the full dark-matter semantic package: mass, stability, action, and cosmological production",
	}
}

func auditALP(inv FiniteAnchorInventory) ALPTopologicalAudit {
	return ALPTopologicalAudit{
		AnchorName:                   "B-sector first spectral gap",
		CandidateScalar:              inv.BGapValue,
		GlobalShiftSymmetryDerived:   false,
		CompactPeriodicFieldDerived:  false,
		AxionDecayConstantDerived:    false,
		InstantonPotentialDerived:    false,
		PontryaginCouplingDerived:    false,
		GaugeAnomalyCoefficientRows:  0,
		AnomalyVector:                []float64{},
		QCDThetaRelaxationDerived:    false,
		GenericALPStructureSupported: false,
		QCDAxionStructureSupported:   false,
		CanonicalNormalizationFound:  false,
		ArbitraryCoefficientInserted: false,
		Verdict:                      "B-gap is a dimensionless spectral scalar only; no continuous shift symmetry, periodic axion coordinate, f_a, instanton potential, or F∧F anomaly coupling is derived",
	}
}

func auditContactDarkSector(inv FiniteAnchorInventory) ContactDarkSectorAudit {
	return ContactDarkSectorAudit{
		ModeCount:                  inv.ContactPartialModeCount,
		SequesteredByNonPromotion:  true,
		GaugeSingletTheoremDerived: false,
		StabilitySymmetryDerived:   false,
		DarkActionDerived:          false,
		SelfInteractionDerived:     false,
		MassScaleDerived:           false,
		ThermalHistoryDerived:      false,
		RelicAbundanceComputed:     false,
		StrictDarkSectorSupported:  false,
		CompatibleFutureRoute:      true,
		Verdict:                    "contact modes are compatible future dark-sector anchors, but non-promotion to SM carriers is not a proof of gauge-singlet stable dark matter",
	}
}

func auditMisalignment(inv FiniteAnchorInventory, alp ALPTopologicalAudit) MisalignmentPreflight {
	return MisalignmentPreflight{
		Formula:                     "Ω_a h² ∝ θ_i² f_a² m_a^{1/2} after a dimensional axion mass m_a and decay constant f_a are supplied",
		RequiresAxionMass:           true,
		RequiresDecayConstant:       true,
		RequiresInitialAngle:        true,
		RequiresCosmologicalHistory: true,
		BGapCanDimensionalizeMass:   inv.BGapPhysicalMassDerived,
		NativeFAFound:               alp.AxionDecayConstantDerived,
		NativeMassFound:             false,
		OmegaComputed:               false,
		OmegaValueH2:                0,
		TargetOmegaH2:               observedOmegaDMH2Hint,
		MisalignmentViable:          false,
		Verdict:                     "misalignment cannot be evaluated: the B-gap is dimensionless and the engine has no native f_a, axion mass, initial angle, or cosmological history map",
	}
}

func auditRelicAccounting(snap Gate224Snapshot, alp ALPTopologicalAudit, contact ContactDarkSectorAudit, mis MisalignmentPreflight) RelicDensityAccounting {
	computed := alp.GenericALPStructureSupported || contact.StrictDarkSectorSupported || mis.OmegaComputed
	return RelicDensityAccounting{
		HeavySectorOmegaH2:          snap.OmegaHeavySectorH2,
		FiniteAnchorOmegaH2Computed: computed,
		FiniteAnchorOmegaH2:         0,
		TotalModelOmegaComputed:     computed,
		ObservedOmegaDMH2Reference:  observedOmegaDMH2Hint,
		DarkMatterStillOpen:         !computed,
		DMDeferredTo: []string{
			"derive a finite shift symmetry and anomaly map for the B-sector gap",
			"derive a stable gauge-singlet dark action for the seven contact partial-overlap modes",
			"derive a dimensionful scale f_dark or f_a without using observed Ω_DM as input",
			"derive a production mechanism and cosmological history before computing Ω h²",
		},
		Verdict: "heavy-sector dark matter is absent and finite-anchor dark matter remains open; Ω_DM is not computed by Gate 225",
	}
}

func auditFirewall(snap Gate224Snapshot, inv FiniteAnchorInventory, alp ALPTopologicalAudit, contact ContactDarkSectorAudit, mis MisalignmentPreflight, relic RelicDensityAccounting) FirewallAudit {
	return FirewallAudit{
		Gate224Inherited:                snap.Gate224Inherited,
		HeavySectorDMAbsencePreserved:   snap.HeavySectorDMAbsent && relic.HeavySectorOmegaH2 == 0,
		RelicDecaySealPreserved:         snap.RelicDecaySealActive,
		FlavorAlignmentSealPreserved:    snap.FlavorAlignmentSealActive,
		BGapUsedAsPhysicalMass:          inv.BGapPhysicalMassDerived,
		BGapUsedAsAxionScale:            false,
		ContactModesPromotedToParticles: false,
		ShiftSymmetryInvented:           alp.GlobalShiftSymmetryDerived && !alp.GenericALPStructureSupported,
		PontryaginCouplingInvented:      alp.PontryaginCouplingDerived && len(alp.AnomalyVector) == 0,
		AxionDecayConstantInvented:      alp.AxionDecayConstantDerived && !alp.CanonicalNormalizationFound,
		RelicDensityInvented:            relic.FiniteAnchorOmegaH2Computed && !mis.OmegaComputed,
		ObservedOmegaUsedForDerivation:  false,
		FiniteCorePolluted:              false,
		Verdict:                         "firewalls closed: Gate 225 does not invent f_a, a shift symmetry, F∧F coupling, contact singlet fields, or Ω_DM",
	}
}

func summarize(snap Gate224Snapshot, alp ALPTopologicalAudit, contact ContactDarkSectorAudit, mis MisalignmentPreflight) Summary {
	viable := alp.GenericALPStructureSupported || contact.StrictDarkSectorSupported || mis.MisalignmentViable
	return Summary{
		ALPSupported:            alp.GenericALPStructureSupported || alp.QCDAxionStructureSupported,
		DarkSectorSupported:     contact.StrictDarkSectorSupported,
		MisalignmentComputed:    mis.OmegaComputed,
		HeavyDMAbsencePreserved: snap.HeavySectorDMAbsent,
		FiniteAnchorDMViable:    viable,
		Status:                  StatusFailedFiniteAnchorDM,
		NextGate:                "Gate 226 — finite dark symmetry / shift-generator construction search",
		Comment:                 "The finite anchors are real inventory, but they are not yet dark matter. The next honest path is to search for a finite symmetry generator or stable dark action before any relic-density calculation.",
	}
}

func buildTruth(snap Gate224Snapshot, inv FiniteAnchorInventory, alp ALPTopologicalAudit, contact ContactDarkSectorAudit, mis MisalignmentPreflight, relic RelicDensityAccounting, s Summary) string {
	return fmt.Sprintf("Gate 225 preserves the Gate-224 heavy-sector dark-matter absence theorem and audits finite anchors for a replacement dark sector. B-gap=%.10f and contact modes=%d are real dimensionless finite data, but no shift symmetry, F∧F anomaly map, f_a, stable contact singlet action, or relic-production law is derived. Therefore finite-anchor dark matter remains an open route, not a theorem.", inv.BGapValue, inv.ContactPartialModeCount)
}

func FormatInventory(inv FiniteAnchorInventory) string {
	return fmt.Sprintf("BGap=%.12f dimensionless=%t physicalMass=%t periodic=%t contactModes=%d contactCharges=%t contactSinglet=%t contactStable=%t contactAction=%t verdict=%s", inv.BGapValue, inv.BGapDimensionless, inv.BGapPhysicalMassDerived, inv.BGapPeriodicCoordinate, inv.ContactPartialModeCount, inv.ContactModesSMChargeDerived, inv.ContactModesSingletProved, inv.ContactModesStableProved, inv.ContactModesActionDerived, inv.InventoryVerdict)
}

func FormatALP(a ALPTopologicalAudit) string {
	return fmt.Sprintf("anchor=%s scalar=%.12f shift=%t periodic=%t f_a=%t FwedgeF=%t anomalyRows=%d qcdAxion=%t genericALP=%t canonicalNorm=%t verdict=%s", a.AnchorName, a.CandidateScalar, a.GlobalShiftSymmetryDerived, a.CompactPeriodicFieldDerived, a.AxionDecayConstantDerived, a.PontryaginCouplingDerived, a.GaugeAnomalyCoefficientRows, a.QCDAxionStructureSupported, a.GenericALPStructureSupported, a.CanonicalNormalizationFound, a.Verdict)
}

func FormatContact(c ContactDarkSectorAudit) string {
	return fmt.Sprintf("modes=%d sequesteredByNonPromotion=%t singletTheorem=%t stability=%t darkAction=%t massScale=%t thermalHistory=%t omega=%t strictDarkSector=%t compatibleFuture=%t verdict=%s", c.ModeCount, c.SequesteredByNonPromotion, c.GaugeSingletTheoremDerived, c.StabilitySymmetryDerived, c.DarkActionDerived, c.MassScaleDerived, c.ThermalHistoryDerived, c.RelicAbundanceComputed, c.StrictDarkSectorSupported, c.CompatibleFutureRoute, c.Verdict)
}

func FormatMisalignment(m MisalignmentPreflight) string {
	return fmt.Sprintf("formula=%q needsMass=%t needsFA=%t needsTheta=%t nativeMass=%t nativeFA=%t omegaComputed=%t targetOmega≈%.3f verdict=%s", m.Formula, m.RequiresAxionMass, m.RequiresDecayConstant, m.RequiresInitialAngle, m.NativeMassFound, m.NativeFAFound, m.OmegaComputed, m.TargetOmegaH2, m.Verdict)
}

func FormatRelic(r RelicDensityAccounting) string {
	return fmt.Sprintf("OmegaHeavy=%.3g finiteOmegaComputed=%t totalOmegaComputed=%t observedRef≈%.3f stillOpen=%t deferred=[%s] verdict=%s", r.HeavySectorOmegaH2, r.FiniteAnchorOmegaH2Computed, r.TotalModelOmegaComputed, r.ObservedOmegaDMH2Reference, r.DarkMatterStillOpen, strings.Join(r.DMDeferredTo, "; "), r.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate224=%t heavyDMabsence=%t relicSeal=%t flavorSeal=%t BGapMass=%t BGapAxionScale=%t contactPromoted=%t shiftInvented=%t FwedgeFInvented=%t faInvented=%t omegaInvented=%t observedOmegaInput=%t finitePolluted=%t verdict=%s", f.Gate224Inherited, f.HeavySectorDMAbsencePreserved, f.RelicDecaySealPreserved, f.FlavorAlignmentSealPreserved, f.BGapUsedAsPhysicalMass, f.BGapUsedAsAxionScale, f.ContactModesPromotedToParticles, f.ShiftSymmetryInvented, f.PontryaginCouplingInvented, f.AxionDecayConstantInvented, f.RelicDensityInvented, f.ObservedOmegaUsedForDerivation, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("ALP=%t darkSector=%t misalignment=%t heavyDMAbsent=%t finiteAnchorDM=%t status=%s next=%s", s.ALPSupported, s.DarkSectorSupported, s.MisalignmentComputed, s.HeavyDMAbsencePreserved, s.FiniteAnchorDMViable, s.Status, s.NextGate)
}

func LoopScaledBGap() float64 { return bSectorFirstGap * loopFactorInv16Pi2 }
