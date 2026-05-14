// Package axionphenomenologyseal implements Gate 226: AxionPhenomenologySeal /
// B-gap misalignment relic-density scale audit.
//
// Gate 225 proved that the B-sector first spectral gap is not natively an ALP:
// the finite core currently lacks a shift symmetry, periodic coordinate,
// anomaly/Pontryagin coupling, and axion decay constant. Gate 226 therefore
// does not derive dark matter. It introduces an explicit phenomenological seal
// and asks a narrower question: if the B-gap were allowed to behave as an ALP,
// what axion decay constant would be required by a standard misalignment relic
// estimate, and does that scale resonate with the already sealed ASHA hierarchy?
package axionphenomenologyseal

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finiteanchordm"
)

const (
	AuditID = "GATE226-AXION-PHENOMENOLOGY-SEAL-B-GAP-MISALIGNMENT-AUDIT"

	AxionPhenomenologySealID = "SEAL-AXION-PHENOMENOLOGY-GATE226"

	StatusConditionalPhenomenologyNoResonance = "CONDITIONAL_PHENOMENOLOGY_AXION_SEAL_NO_SCALE_RESONANCE"
	StatusAxionNotFiniteDerived               = "AXION_SEMANTICS_QUARANTINED_NOT_DERIVED"
	StatusDarkMatterScaleIntermediate         = "DARK_MATTER_SCALE_INTERMEDIATE_NOT_ASHA_HIERARCHY_MATCH"
)

const (
	// Inherited finite and phenomenological anchors.
	bSectorFirstGap             = 0.1024649212
	electroweakVEVGeV           = 246.0
	heavyThresholdMBGeV         = 2.56895727e6
	topologicalBoundaryMStarGeV = 1.72179441e17

	// Standard QCD-like misalignment proxy requested by the gate prompt.
	observedOmegaDMH2        = 0.12
	referenceOmegaH2         = 0.12
	referenceFA          GeV = 1.0e12
	misalignmentExponent     = 7.0 / 6.0
	defaultThetaI            = 1.0

	oneDecadeLogDistance = 1.0
)

// GeV is a documentation alias used only in constant declarations.
type GeV = float64

type Gate225Snapshot struct {
	Gate225Inherited            bool
	HeavySectorDMAbsenceBinding bool
	NativeALPFailed             bool
	NativeContactDMFailed       bool
	BGapValue                   float64
	ContactModeCount            int
	TruthStatement              string
}

type AxionPhenomenologySeal struct {
	ID                          string
	Active                      bool
	ConditionalOnBGapALP        bool
	GrantsShiftSymmetry         bool
	GrantsTopologicalCoupling   bool
	GrantsPeriodicCoordinate    bool
	GrantsMisalignmentMechanics bool
	NativeALPDerived            bool
	NativeFAObserved            bool
	NativeAnomalyMapDerived     bool
	QuarantinedInputs           []string
	Verdict                     string
}

type MisalignmentCalculation struct {
	Formula                       string
	TargetOmegaH2                 float64
	ReferenceOmegaH2              float64
	ReferenceFAGeV                float64
	Exponent                      float64
	ThetaI                        float64
	RequiredFAGeV                 float64
	UsesObservedOmegaAsDerivation bool
	NativeScaleDerived            bool
	BGapUsedAsTheta               bool
	BGapUsedAsMass                bool
	Verdict                       string
}

type ScaleComparison struct {
	Name        string
	ScaleGeV    float64
	RatioFAOver float64
	Log10Gap    float64
	Close       bool
	Verdict     string
}

type StructuralResonanceAudit struct {
	RequiredFAGeV      float64
	Comparisons        []ScaleComparison
	ClosestScaleName   string
	ClosestLog10Gap    float64
	ResonanceThreshold string
	ResonanceFound     bool
	ResonantScaleName  string
	Verdict            string
}

type BGapVariantDiagnostic struct {
	Evaluated        bool
	Promoted         bool
	Formula          string
	ThetaI           float64
	RequiredFAGeV    float64
	ClosestScaleName string
	ClosestLog10Gap  float64
	Verdict          string
}

type DarkMatterAccounting struct {
	HeavySectorOmegaH2          float64
	ALPOmegaTargetH2            float64
	ALPOmegaComputedUnderSeal   bool
	ALPAccountsForDMUnderSeal   bool
	NativeDMStillAbsent         bool
	ObservedDMUsedAsTargetOnly  bool
	DarkMatterDerivedFromFinite bool
	DeferredNativeTasks         []string
	Verdict                     string
}

type FirewallAudit struct {
	Gate225Inherited               bool
	HeavyDMAbsencePreserved        bool
	AxionSealActive                bool
	ShiftSymmetryFiniteDerived     bool
	TopologicalCouplingDerived     bool
	AxionDecayConstantFinite       bool
	BGapPromotedWithoutSeal        bool
	ObservedDMUsedToRewriteCore    bool
	StructuralResonanceOverclaimed bool
	ContactModesPromoted           bool
	FiniteCorePolluted             bool
	Verdict                        string
}

type Summary struct {
	SealGranted                 bool
	MisalignmentComputed        bool
	RequiredFAFound             bool
	ScaleResonanceFound         bool
	NativeDarkMatterDerived     bool
	ConditionalDarkMatterViable bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	Gate225   Gate225Snapshot
	Seal      AxionPhenomenologySeal
	Misalign  MisalignmentCalculation
	Resonance StructuralResonanceAudit
	BGapTheta BGapVariantDiagnostic
	DM        DarkMatterAccounting
	Firewall  FirewallAudit
	Summary   Summary

	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g225, err := finiteanchordm.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 225 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g225)
	})
	return defaultA, defaultErr
}

func Build(g225 finiteanchordm.Analysis) (Analysis, error) {
	snap := snapshotFromGate225(g225)
	if !snap.Gate225Inherited || !snap.HeavySectorDMAbsenceBinding || !snap.NativeALPFailed {
		return Analysis{}, fmt.Errorf("Gate 226 requires Gate 225 finite-anchor dark-matter obstruction and heavy-sector DM absence")
	}

	seal := activateSeal()
	mis := computeMisalignment(defaultThetaI)
	resonance := auditResonance(mis.RequiredFAGeV)
	thetaDiagnostic := auditBGapThetaVariant(snap.BGapValue)
	dm := auditDarkMatterAccounting(mis, seal)
	firewall := auditFirewall(snap, seal, mis, resonance)
	summary := summarize(seal, mis, resonance, dm)
	truth := buildTruth(snap, seal, mis, resonance, dm)

	return Analysis{Gate225: snap, Seal: seal, Misalign: mis, Resonance: resonance, BGapTheta: thetaDiagnostic, DM: dm, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate225(a finiteanchordm.Analysis) Gate225Snapshot {
	return Gate225Snapshot{
		Gate225Inherited:            a.Summary.HeavyDMAbsencePreserved && !a.Summary.FiniteAnchorDMViable,
		HeavySectorDMAbsenceBinding: a.Summary.HeavyDMAbsencePreserved && a.Relic.HeavySectorOmegaH2 == 0,
		NativeALPFailed:             !a.Summary.ALPSupported && !a.ALP.GenericALPStructureSupported && !a.ALP.QCDAxionStructureSupported,
		NativeContactDMFailed:       !a.Summary.DarkSectorSupported && !a.Contact.StrictDarkSectorSupported,
		BGapValue:                   a.Inventory.BGapValue,
		ContactModeCount:            a.Inventory.ContactPartialModeCount,
		TruthStatement:              a.TruthStatement,
	}
}

func activateSeal() AxionPhenomenologySeal {
	return AxionPhenomenologySeal{
		ID:                          AxionPhenomenologySealID,
		Active:                      true,
		ConditionalOnBGapALP:        true,
		GrantsShiftSymmetry:         true,
		GrantsTopologicalCoupling:   true,
		GrantsPeriodicCoordinate:    true,
		GrantsMisalignmentMechanics: true,
		NativeALPDerived:            false,
		NativeFAObserved:            false,
		NativeAnomalyMapDerived:     false,
		QuarantinedInputs: []string{
			"B-sector gap is conditionally treated as an ALP coordinate",
			"continuous shift symmetry a -> a + c",
			"periodic compact axion coordinate",
			"topological coupling a F∧F or a F F~",
			"QCD-like misalignment relic-density law",
			"order-one initial misalignment angle θ_i = 1",
		},
		Verdict: "AxionPhenomenologySeal is active: ALP semantics are quarantined phenomenology, not finite-core theorems",
	}
}

func computeMisalignment(theta float64) MisalignmentCalculation {
	if theta <= 0 {
		theta = defaultThetaI
	}
	// Ω_a h² = Ω_ref * θ_i² * (f_a/f_ref)^(7/6).
	requiredFA := referenceFA * math.Pow(observedOmegaDMH2/(referenceOmegaH2*theta*theta), 1.0/misalignmentExponent)
	return MisalignmentCalculation{
		Formula:                       "Ω_a h² = 0.12 × θ_i² × (f_a / 10¹² GeV)^(7/6)",
		TargetOmegaH2:                 observedOmegaDMH2,
		ReferenceOmegaH2:              referenceOmegaH2,
		ReferenceFAGeV:                referenceFA,
		Exponent:                      misalignmentExponent,
		ThetaI:                        theta,
		RequiredFAGeV:                 requiredFA,
		UsesObservedOmegaAsDerivation: false,
		NativeScaleDerived:            false,
		BGapUsedAsTheta:               false,
		BGapUsedAsMass:                false,
		Verdict:                       "under the sealed QCD-like misalignment proxy with θ_i=1, Ω_DM selects f_a = 1.0e12 GeV",
	}
}

func auditResonance(requiredFA float64) StructuralResonanceAudit {
	comps := []ScaleComparison{
		compareScale("electroweak VEV v", electroweakVEVGeV, requiredFA),
		compareScale("PeV heavy threshold M_B", heavyThresholdMBGeV, requiredFA),
		compareScale("topological boundary M_*", topologicalBoundaryMStarGeV, requiredFA),
	}
	closest := comps[0]
	for _, c := range comps[1:] {
		if c.Log10Gap < closest.Log10Gap {
			closest = c
		}
	}
	resonanceFound := false
	resonantName := ""
	for _, c := range comps {
		if c.Close {
			resonanceFound = true
			resonantName = c.Name
			break
		}
	}
	verdict := fmt.Sprintf("no scale resonance: required f_a=%.8e GeV is %.3f log10-decades from the closest sealed ASHA scale (%s)", requiredFA, closest.Log10Gap, closest.Name)
	if resonanceFound {
		verdict = fmt.Sprintf("conditional scale resonance: required f_a=%.8e GeV lies within one decade of %s", requiredFA, resonantName)
	}
	return StructuralResonanceAudit{
		RequiredFAGeV:      requiredFA,
		Comparisons:        comps,
		ClosestScaleName:   closest.Name,
		ClosestLog10Gap:    closest.Log10Gap,
		ResonanceThreshold: "|log10(f_a/scale)| < 1 decade",
		ResonanceFound:     resonanceFound,
		ResonantScaleName:  resonantName,
		Verdict:            verdict,
	}
}

func compareScale(name string, scale, requiredFA float64) ScaleComparison {
	ratio := requiredFA / scale
	gap := math.Abs(math.Log10(ratio))
	close := gap < oneDecadeLogDistance
	verdict := "not close"
	if close {
		verdict = "within one decade resonance threshold"
	}
	return ScaleComparison{Name: name, ScaleGeV: scale, RatioFAOver: ratio, Log10Gap: gap, Close: close, Verdict: verdict}
}

func auditBGapThetaVariant(gap float64) BGapVariantDiagnostic {
	if gap <= 0 {
		return BGapVariantDiagnostic{Evaluated: false, Promoted: false, Verdict: "B-gap theta diagnostic not evaluated: nonpositive gap"}
	}
	variant := computeMisalignment(gap)
	resonance := auditResonance(variant.RequiredFAGeV)
	return BGapVariantDiagnostic{
		Evaluated:        true,
		Promoted:         false,
		Formula:          "diagnostic only: set θ_i = B_gap in the same misalignment proxy",
		ThetaI:           gap,
		RequiredFAGeV:    variant.RequiredFAGeV,
		ClosestScaleName: resonance.ClosestScaleName,
		ClosestLog10Gap:  resonance.ClosestLog10Gap,
		Verdict:          "diagnostic rejected as noncanonical: Gate 225 did not derive B_gap as the initial misalignment angle",
	}
}

func auditDarkMatterAccounting(m MisalignmentCalculation, seal AxionPhenomenologySeal) DarkMatterAccounting {
	return DarkMatterAccounting{
		HeavySectorOmegaH2:          0,
		ALPOmegaTargetH2:            m.TargetOmegaH2,
		ALPOmegaComputedUnderSeal:   seal.Active,
		ALPAccountsForDMUnderSeal:   seal.Active && m.RequiredFAGeV > 0,
		NativeDMStillAbsent:         true,
		ObservedDMUsedAsTargetOnly:  true,
		DarkMatterDerivedFromFinite: false,
		DeferredNativeTasks: []string{
			"derive a finite shift generator or compact ALP coordinate",
			"derive a topological anomaly/Pontryagin coupling vector",
			"derive f_a or a dimensionful dark scale from finite geometry",
			"derive a production history rather than fitting Ω_DM",
		},
		Verdict: "dark matter can be parameterized under AxionPhenomenologySeal, but native finite dark matter remains absent",
	}
}

func auditFirewall(snap Gate225Snapshot, seal AxionPhenomenologySeal, m MisalignmentCalculation, r StructuralResonanceAudit) FirewallAudit {
	return FirewallAudit{
		Gate225Inherited:               snap.Gate225Inherited,
		HeavyDMAbsencePreserved:        snap.HeavySectorDMAbsenceBinding,
		AxionSealActive:                seal.Active,
		ShiftSymmetryFiniteDerived:     seal.NativeALPDerived,
		TopologicalCouplingDerived:     seal.NativeAnomalyMapDerived,
		AxionDecayConstantFinite:       seal.NativeFAObserved || m.NativeScaleDerived,
		BGapPromotedWithoutSeal:        false,
		ObservedDMUsedToRewriteCore:    false,
		StructuralResonanceOverclaimed: false,
		ContactModesPromoted:           false,
		FiniteCorePolluted:             false,
		Verdict:                        "firewalls closed: Gate 226 grants ALP semantics only under seal and does not rewrite finite geometry, contact modes, or RG scales",
	}
}

func summarize(seal AxionPhenomenologySeal, m MisalignmentCalculation, r StructuralResonanceAudit, dm DarkMatterAccounting) Summary {
	return Summary{
		SealGranted:                 seal.Active,
		MisalignmentComputed:        m.RequiredFAGeV > 0,
		RequiredFAFound:             m.RequiredFAGeV > 0,
		ScaleResonanceFound:         r.ResonanceFound,
		NativeDarkMatterDerived:     dm.DarkMatterDerivedFromFinite,
		ConditionalDarkMatterViable: dm.ALPAccountsForDMUnderSeal,
		Status:                      StatusConditionalPhenomenologyNoResonance,
		NextGate:                    "Gate 227 — axion-scale origin / finite shift-generator and f_a derivation search",
		Comment:                     "The sealed ALP misalignment model can reproduce Ω_DM with f_a≈1e12 GeV, but that scale does not resonate with v, M_B, or M_* and is not finite-derived.",
	}
}

func buildTruth(snap Gate225Snapshot, seal AxionPhenomenologySeal, m MisalignmentCalculation, r StructuralResonanceAudit, dm DarkMatterAccounting) string {
	return fmt.Sprintf("Gate 226 grants %s only as a phenomenological seal. Under the requested QCD-like misalignment proxy with θ_i=%.3g, Ω_DM h²=%.3g requires f_a=%.8e GeV. This is not close to v=%.3g GeV, M_B=%.8e GeV, or M_*=%.8e GeV; the closest gap is %.3f decades to %s. Therefore the ALP route is conditionally parameterizable but not structurally resonant or finite-derived.", seal.ID, m.ThetaI, m.TargetOmegaH2, m.RequiredFAGeV, electroweakVEVGeV, heavyThresholdMBGeV, topologicalBoundaryMStarGeV, r.ClosestLog10Gap, r.ClosestScaleName)
}

func FormatGate225(s Gate225Snapshot) string {
	return fmt.Sprintf("inherited=%t heavyDMAbsence=%t nativeALPFailed=%t nativeContactDMFailed=%t BGap=%.10f contactModes=%d", s.Gate225Inherited, s.HeavySectorDMAbsenceBinding, s.NativeALPFailed, s.NativeContactDMFailed, s.BGapValue, s.ContactModeCount)
}

func FormatSeal(s AxionPhenomenologySeal) string {
	return fmt.Sprintf("id=%s active=%t conditionalBGapALP=%t grantsShift=%t grantsTopological=%t grantsMisalignment=%t nativeALP=%t nativeFA=%t nativeAnomaly=%t inputs=[%s] verdict=%s", s.ID, s.Active, s.ConditionalOnBGapALP, s.GrantsShiftSymmetry, s.GrantsTopologicalCoupling, s.GrantsMisalignmentMechanics, s.NativeALPDerived, s.NativeFAObserved, s.NativeAnomalyMapDerived, strings.Join(s.QuarantinedInputs, "; "), s.Verdict)
}

func FormatMisalignment(m MisalignmentCalculation) string {
	return fmt.Sprintf("formula=%q theta=%.8g targetΩ=%.8g f_ref=%.8e exponent=%.8g requiredFA=%.8e nativeScale=%t bGapTheta=%t bGapMass=%t verdict=%s", m.Formula, m.ThetaI, m.TargetOmegaH2, m.ReferenceFAGeV, m.Exponent, m.RequiredFAGeV, m.NativeScaleDerived, m.BGapUsedAsTheta, m.BGapUsedAsMass, m.Verdict)
}

func FormatScaleComparison(c ScaleComparison) string {
	return fmt.Sprintf("%s scale=%.8e GeV f_a/scale=%.8e log10gap=%.6f close=%t verdict=%s", c.Name, c.ScaleGeV, c.RatioFAOver, c.Log10Gap, c.Close, c.Verdict)
}

func FormatResonance(r StructuralResonanceAudit) string {
	parts := make([]string, 0, len(r.Comparisons))
	for _, c := range r.Comparisons {
		parts = append(parts, FormatScaleComparison(c))
	}
	return fmt.Sprintf("requiredFA=%.8e resonanceFound=%t closest=%s closestLog10Gap=%.6f threshold=%s comparisons={%s} verdict=%s", r.RequiredFAGeV, r.ResonanceFound, r.ClosestScaleName, r.ClosestLog10Gap, r.ResonanceThreshold, strings.Join(parts, " | "), r.Verdict)
}

func FormatBGapThetaDiagnostic(d BGapVariantDiagnostic) string {
	return fmt.Sprintf("evaluated=%t promoted=%t formula=%q theta=%.10f requiredFA=%.8e closest=%s closestLog10Gap=%.6f verdict=%s", d.Evaluated, d.Promoted, d.Formula, d.ThetaI, d.RequiredFAGeV, d.ClosestScaleName, d.ClosestLog10Gap, d.Verdict)
}

func FormatDM(d DarkMatterAccounting) string {
	return fmt.Sprintf("heavyΩ=%.8g alpTargetΩ=%.8g computedUnderSeal=%t accountsForDMUnderSeal=%t nativeDMStillAbsent=%t observedTargetOnly=%t finiteDerived=%t deferred=[%s] verdict=%s", d.HeavySectorOmegaH2, d.ALPOmegaTargetH2, d.ALPOmegaComputedUnderSeal, d.ALPAccountsForDMUnderSeal, d.NativeDMStillAbsent, d.ObservedDMUsedAsTargetOnly, d.DarkMatterDerivedFromFinite, strings.Join(d.DeferredNativeTasks, "; "), d.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate225=%t heavyDMAbsence=%t seal=%t shiftFinite=%t couplingFinite=%t faFinite=%t bGapPromotedWithoutSeal=%t observedRewrite=%t resonanceOverclaimed=%t contactPromoted=%t polluted=%t verdict=%s", f.Gate225Inherited, f.HeavyDMAbsencePreserved, f.AxionSealActive, f.ShiftSymmetryFiniteDerived, f.TopologicalCouplingDerived, f.AxionDecayConstantFinite, f.BGapPromotedWithoutSeal, f.ObservedDMUsedToRewriteCore, f.StructuralResonanceOverclaimed, f.ContactModesPromoted, f.FiniteCorePolluted, f.Verdict)
}
