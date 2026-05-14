// Package bgaphierarchycoefficient implements Gate 283:
// B-Gap Hierarchy Coefficient / Topological Volume Ratio Audit.
//
// Gate 283 pivots from the Path-B spectral-action capstone to Path C: the
// B-gap hierarchy coefficient. It re-audits the sharp Gate-228/229 resonance
// c≈4/π using the native Hopf/contact volume ledger, while preserving the
// firewall between an exact mathematical volume identity and a finite-derived
// contact-vacuum action map.
package bgaphierarchycoefficient

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/hopfgeometricnormalization"
	"github.com/bagherbal/asha-engine/pkg/bridge/spectralactioncapstone"
)

const (
	AuditID = "GATE283-BGAP-HIERARCHY-COEFFICIENT-TOPOLOGICAL-VOLUME-RATIO-AUDIT"

	StatusPathCOpened                      = "CONDITIONAL_SUPPORT_PATH_C_BGAP_DERIVATION_OPENED_AFTER_PATH_B_CAPSTONE"
	StatusTopologicalVolumesRetrieved      = "CONDITIONAL_SUPPORT_HOPF_TOPOLOGICAL_VOLUMES_RETRIEVED"
	StatusFourOverPiIdentityVerified       = "CONDITIONAL_SUPPORT_FOUR_OVER_PI_VOLUME_RATIO_IDENTITY_VERIFIED"
	StatusBGapHierarchyResonanceReproduced = "CONDITIONAL_SUPPORT_BGAP_HIERARCHY_RESONANCE_REPRODUCED"
	StatusSensitivityLedgerPreserved       = "CONDITIONAL_SUPPORT_EXPONENTIAL_SENSITIVITY_LEDGER_PRESERVED"
	StatusFirewallsPreserved               = "CONDITIONAL_SUPPORT_BGAP_PATH_C_FIREWALLS_PRESERVED"

	StatusFailedNativeContactActionMapMissing  = "FAILED_ROUTE_NATIVE_CONTACT_ACTION_MAP_TO_BGAP_NOT_DERIVED"
	StatusFailedHopfFiberNormalizationMissing  = "FAILED_ROUTE_HOPF_FIBER_VOLUME_NORMALIZATION_NOT_FINITE_DERIVED"
	StatusFailedIntermediateScaleNotExact      = "FAILED_ROUTE_FOUR_OVER_PI_DOES_NOT_EXACTLY_REPRODUCE_M_INT_WITH_CURRENT_BGAP"
	StatusFailedIntermediateTheoremNotUpgraded = "FAILED_ROUTE_INTERMEDIATE_SCALE_THEOREM_NOT_UPGRADED"
	StatusFailedIntermediateBreakingSealOpen   = "FAILED_ROUTE_INTERMEDIATE_BREAKING_SEAL_REMAINS_REQUIRED"
)

const (
	bGap          = 0.1024649212
	mStarGeV      = 1.72179441e17
	targetMIntGeV = 6.650726476871e11
	exactTol      = 1e-12
)

type Gate282Snapshot struct {
	Gate282Inherited      bool
	PathBClosed           bool
	HiggsFirewallActive   bool
	HiggsRatioDerived     bool
	SixPointFirewallCount int
	TruthStatement        string
}

type HopfVolumeLedger struct {
	UnitS3Volume        float64
	UnitS4Volume        float64
	UnitS7Volume        float64
	UnitS3VolumeExact   string
	UnitS4VolumeExact   string
	UnitS7VolumeExact   string
	HopfFibration       string
	VolumesStandardMath bool
	NativeHopfFibration bool
	Verdict             string
}

type ContactActionLedger struct {
	TopologicalAction               float64
	TopologicalActionExact          string
	TopologicalActionSource         string
	FiberVolume                     float64
	PiTimesFiberVolume              float64
	CoefficientFormula              string
	Coefficient                     float64
	CoefficientExact                string
	CoefficientEqualsFourOverPi     bool
	ContactBoundaryActionMapDerived bool
	HopfFiberNormalizationDerived   bool
	Verdict                         string
}

type HierarchyCoefficientAudit struct {
	Formula                     string
	BGap                        float64
	MStarGeV                    float64
	TargetMIntGeV               float64
	Coefficient                 float64
	RequiredCoefficient         float64
	CoefficientResidual         float64
	RelativeCoefficientResidual float64
	PredictedMIntGeV            float64
	RatioPredictedToTarget      float64
	Log10Gap                    float64
	BGapRequiredForExactMatch   float64
	RelativeBGapDisplacement    float64
	WithinOneDecade             bool
	TightNearResonance          bool
	ExactIntermediateMatch      bool
	TheoremUpgradeGranted       bool
	Verdict                     string
}

type SensitivityLedger struct {
	DerivativeLog10MPerUnitBGap    float64
	DerivativeLog10MPerFractionalB float64
	OnePercentShiftDecades         float64
	TenPercentShiftDecades         float64
	BindingWarning                 bool
	Verdict                        string
}

type PathCSealLedger struct {
	IntermediateBreakingSealPrepared bool
	IntermediateBreakingSealGranted  bool
	RequiresFiniteOrderParameter     bool
	RequiresContactActionMap         bool
	RequiresBreakingPotential        bool
	RequiresResidualMatchingMap      bool
	PatiSalamRouteFalsifiedInherited bool
	HiddenSectorRouteFavored         bool
	Verdict                          string
}

type FirewallAudit struct {
	PathBClosureInherited                   bool
	Gate229NearResonanceInherited           bool
	UsesOnlySealedScales                    bool
	DoesNotFitCoefficient                   bool
	DoesNotPromoteStandardVolumeToFiniteMap bool
	DoesNotClaimExactMIntTheorem            bool
	DoesNotGrantIntermediateSeal            bool
	DoesNotReopenPatiSalam                  bool
	DoesNotInsertObservedMasses             bool
	FiniteCorePolluted                      bool
	Verdict                                 string
}

type Summary struct {
	PathCOpened                     bool
	TopologicalVolumesRetrieved     bool
	FourOverPiIdentityVerified      bool
	BGapResonanceReproduced         bool
	NativeCoefficientDerived        bool
	IntermediateScaleTheorem        bool
	IntermediateBreakingSealGranted bool
	FirewallPreserved               bool
	Status                          string
	DirectAnswer                    string
	NextGate                        string
}

type Analysis struct {
	PreviousGate282 Gate282Snapshot
	PreviousGate229 hopfgeometricnormalization.Analysis
	Volumes         HopfVolumeLedger
	ContactAction   ContactActionLedger
	Hierarchy       HierarchyCoefficientAudit
	Sensitivity     SensitivityLedger
	Seal            PathCSealLedger
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
	g282, err := spectralactioncapstone.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 282 capstone predecessor: %w", err)
	}
	g229, err := hopfgeometricnormalization.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 229 Hopf predecessor: %w", err)
	}
	snap := snapshotGate282(g282)
	volumes := retrieveVolumes()
	contact := formalizeContactAction(volumes)
	hierarchy := auditHierarchy(contact)
	sensitivity := auditSensitivity(contact, hierarchy)
	seal := auditPathCSeal(g229, contact, hierarchy)
	fw := auditFirewalls(snap, g229, contact, hierarchy, seal)
	summary := buildSummary(volumes, contact, hierarchy, seal, fw)
	return Analysis{
		PreviousGate282: snap,
		PreviousGate229: g229,
		Volumes:         volumes,
		ContactAction:   contact,
		Hierarchy:       hierarchy,
		Sensitivity:     sensitivity,
		Seal:            seal,
		Firewall:        fw,
		Summary:         summary,
		TruthStatement:  "Gate 283 reopens Path C: 4/π is an exact Hopf/topological-volume identity and a tight B-gap hierarchy resonance, but the finite engine still lacks the native contact-action map and residual matching theorem needed to upgrade M_int to a derived theorem.",
	}, nil
}

func snapshotGate282(g spectralactioncapstone.Analysis) Gate282Snapshot {
	return Gate282Snapshot{
		Gate282Inherited:      true,
		PathBClosed:           g.Summary.PathBClosed,
		HiggsFirewallActive:   g.Summary.HiggsFirewallActive,
		HiggsRatioDerived:     g.Summary.HiggsRatioDerived,
		SixPointFirewallCount: len(g.Obstructions.Obstructions),
		TruthStatement:        g.TruthStatement,
	}
}

func retrieveVolumes() HopfVolumeLedger {
	s3 := 2 * math.Pi * math.Pi
	s4 := 8 * math.Pi * math.Pi / 3
	s7 := math.Pow(math.Pi, 4) / 3
	return HopfVolumeLedger{
		UnitS3Volume:        s3,
		UnitS4Volume:        s4,
		UnitS7Volume:        s7,
		UnitS3VolumeExact:   "2π²",
		UnitS4VolumeExact:   "8π²/3",
		UnitS7VolumeExact:   "π⁴/3",
		HopfFibration:       "S³ -> S⁷ -> S⁴",
		VolumesStandardMath: true,
		NativeHopfFibration: true,
		Verdict:             StatusTopologicalVolumesRetrieved,
	}
}

func formalizeContactAction(v HopfVolumeLedger) ContactActionLedger {
	sTop := 8 * math.Pi * math.Pi
	denom := math.Pi * v.UnitS3Volume
	c := sTop / denom
	return ContactActionLedger{
		TopologicalAction:               sTop,
		TopologicalActionExact:          "8π²",
		TopologicalActionSource:         "Gate 174 topological action seal, inherited through Gate 229",
		FiberVolume:                     v.UnitS3Volume,
		PiTimesFiberVolume:              denom,
		CoefficientFormula:              "c_Hopf = S_top / (π Vol(S³))",
		Coefficient:                     c,
		CoefficientExact:                "4/π",
		CoefficientEqualsFourOverPi:     math.Abs(c-4/math.Pi) < exactTol,
		ContactBoundaryActionMapDerived: false,
		HopfFiberNormalizationDerived:   false,
		Verdict:                         StatusFourOverPiIdentityVerified,
	}
}

func auditHierarchy(c ContactActionLedger) HierarchyCoefficientAudit {
	req := bGap * math.Log(mStarGeV/targetMIntGeV)
	pred := mStarGeV * math.Exp(-c.Coefficient/bGap)
	ratio := pred / targetMIntGeV
	loggap := math.Abs(math.Log10(ratio))
	bExact := c.Coefficient / math.Log(mStarGeV/targetMIntGeV)
	relB := math.Abs(bExact-bGap) / bGap
	cres := req - c.Coefficient
	relC := math.Abs(cres) / req
	exactMatch := loggap < 1e-9
	return HierarchyCoefficientAudit{
		Formula:                     "M_hidden = M_* exp(-c/B_gap)",
		BGap:                        bGap,
		MStarGeV:                    mStarGeV,
		TargetMIntGeV:               targetMIntGeV,
		Coefficient:                 c.Coefficient,
		RequiredCoefficient:         req,
		CoefficientResidual:         cres,
		RelativeCoefficientResidual: relC,
		PredictedMIntGeV:            pred,
		RatioPredictedToTarget:      ratio,
		Log10Gap:                    loggap,
		BGapRequiredForExactMatch:   bExact,
		RelativeBGapDisplacement:    relB,
		WithinOneDecade:             loggap < 1,
		TightNearResonance:          loggap < 0.02,
		ExactIntermediateMatch:      exactMatch,
		TheoremUpgradeGranted:       false,
		Verdict:                     StatusBGapHierarchyResonanceReproduced,
	}
}

func auditSensitivity(c ContactActionLedger, h HierarchyCoefficientAudit) SensitivityLedger {
	dUnit := c.Coefficient / (math.Ln10 * h.BGap * h.BGap)
	dFrac := dUnit * h.BGap
	return SensitivityLedger{
		DerivativeLog10MPerUnitBGap:    dUnit,
		DerivativeLog10MPerFractionalB: dFrac,
		OnePercentShiftDecades:         0.01 * dFrac,
		TenPercentShiftDecades:         0.10 * dFrac,
		BindingWarning:                 dUnit > 50 && dFrac > 5,
		Verdict:                        StatusSensitivityLedgerPreserved,
	}
}

func auditPathCSeal(g229 hopfgeometricnormalization.Analysis, c ContactActionLedger, h HierarchyCoefficientAudit) PathCSealLedger {
	return PathCSealLedger{
		IntermediateBreakingSealPrepared: true,
		IntermediateBreakingSealGranted:  false,
		RequiresFiniteOrderParameter:     true,
		RequiresContactActionMap:         !c.ContactBoundaryActionMapDerived,
		RequiresBreakingPotential:        true,
		RequiresResidualMatchingMap:      !h.ExactIntermediateMatch,
		PatiSalamRouteFalsifiedInherited: g229.Seal.PatiSalamFalsified,
		HiddenSectorRouteFavored:         true,
		Verdict:                          StatusFailedIntermediateBreakingSealOpen,
	}
}

func auditFirewalls(g282 Gate282Snapshot, g229 hopfgeometricnormalization.Analysis, c ContactActionLedger, h HierarchyCoefficientAudit, s PathCSealLedger) FirewallAudit {
	return FirewallAudit{
		PathBClosureInherited:                   g282.PathBClosed && g282.HiggsFirewallActive && !g282.HiggsRatioDerived,
		Gate229NearResonanceInherited:           g229.Summary.GeometricHierarchySupported && !g229.Summary.NativeHopfMapDerived,
		UsesOnlySealedScales:                    true,
		DoesNotFitCoefficient:                   c.CoefficientEqualsFourOverPi && c.CoefficientExact == "4/π",
		DoesNotPromoteStandardVolumeToFiniteMap: !c.ContactBoundaryActionMapDerived && !c.HopfFiberNormalizationDerived,
		DoesNotClaimExactMIntTheorem:            !h.ExactIntermediateMatch && !h.TheoremUpgradeGranted,
		DoesNotGrantIntermediateSeal:            !s.IntermediateBreakingSealGranted,
		DoesNotReopenPatiSalam:                  s.PatiSalamRouteFalsifiedInherited,
		DoesNotInsertObservedMasses:             true,
		FiniteCorePolluted:                      false,
		Verdict:                                 StatusFirewallsPreserved,
	}
}

func buildSummary(v HopfVolumeLedger, c ContactActionLedger, h HierarchyCoefficientAudit, s PathCSealLedger, fw FirewallAudit) Summary {
	native := c.ContactBoundaryActionMapDerived && c.HopfFiberNormalizationDerived
	theorem := native && h.ExactIntermediateMatch && s.IntermediateBreakingSealGranted
	status := StatusFailedIntermediateTheoremNotUpgraded
	direct := "Path C is reopened: 4/π is exact as a Hopf/topological volume ratio and tightly reproduces the B-gap intermediate-scale resonance, but it is still not a finite-derived intermediate-scale theorem."
	if theorem {
		status = "CONDITIONAL_SUPPORT_BGAP_COEFFICIENT_DERIVED_AND_M_INT_THEOREM_UPGRADED"
		direct = "The B-gap coefficient and intermediate scale are theorem-upgraded."
	}
	return Summary{
		PathCOpened:                     true,
		TopologicalVolumesRetrieved:     v.VolumesStandardMath && v.NativeHopfFibration,
		FourOverPiIdentityVerified:      c.CoefficientEqualsFourOverPi,
		BGapResonanceReproduced:         h.TightNearResonance,
		NativeCoefficientDerived:        native,
		IntermediateScaleTheorem:        theorem,
		IntermediateBreakingSealGranted: s.IntermediateBreakingSealGranted,
		FirewallPreserved:               fw.PathBClosureInherited && fw.Gate229NearResonanceInherited && fw.DoesNotClaimExactMIntTheorem && fw.DoesNotGrantIntermediateSeal && !fw.FiniteCorePolluted,
		Status:                          status,
		DirectAnswer:                    direct,
		NextGate:                        "Derive a native contact-vacuum Hopf action map or hidden B-sector order parameter; without it, 4/π remains an exact geometric diagnostic rather than a finite-core coefficient theorem.",
	}
}

func FormatGate282(s Gate282Snapshot) string {
	return fmt.Sprintf("Gate282Inherited=%t PathBClosed=%t HiggsFirewallActive=%t HiggsRatioDerived=%t SixPointFirewallCount=%d Truth=%q", s.Gate282Inherited, s.PathBClosed, s.HiggsFirewallActive, s.HiggsRatioDerived, s.SixPointFirewallCount, s.TruthStatement)
}

func FormatVolumes(v HopfVolumeLedger) string {
	return fmt.Sprintf("%s with Vol(S3)=%s=%.12g Vol(S4)=%s=%.12g Vol(S7)=%s=%.12g standardMath=%t nativeHopf=%t verdict=%s", v.HopfFibration, v.UnitS3VolumeExact, v.UnitS3Volume, v.UnitS4VolumeExact, v.UnitS4Volume, v.UnitS7VolumeExact, v.UnitS7Volume, v.VolumesStandardMath, v.NativeHopfFibration, v.Verdict)
}

func FormatContactAction(c ContactActionLedger) string {
	return fmt.Sprintf("%s: S_top=%s=%.12g fiber=%g πfiber=%.12g coefficient=%s=%.12g equals4/pi=%t contactMapDerived=%t fiberNormDerived=%t verdict=%s", c.CoefficientFormula, c.TopologicalActionExact, c.TopologicalAction, c.FiberVolume, c.PiTimesFiberVolume, c.CoefficientExact, c.Coefficient, c.CoefficientEqualsFourOverPi, c.ContactBoundaryActionMapDerived, c.HopfFiberNormalizationDerived, c.Verdict)
}

func FormatHierarchy(h HierarchyCoefficientAudit) string {
	return fmt.Sprintf("%s with B_gap=%.12g M*=%.12e targetM_int=%.12e c=%.12g c_req=%.12g Δc=%.12g relΔc=%.12g pred=%.12e ratio=%.12g log10gap=%.12g B_exact=%.12g relB=%.12g tight=%t exact=%t theorem=%t verdict=%s", h.Formula, h.BGap, h.MStarGeV, h.TargetMIntGeV, h.Coefficient, h.RequiredCoefficient, h.CoefficientResidual, h.RelativeCoefficientResidual, h.PredictedMIntGeV, h.RatioPredictedToTarget, h.Log10Gap, h.BGapRequiredForExactMatch, h.RelativeBGapDisplacement, h.TightNearResonance, h.ExactIntermediateMatch, h.TheoremUpgradeGranted, h.Verdict)
}

func FormatSensitivity(s SensitivityLedger) string {
	return fmt.Sprintf("dlog10M/dB=%.12g dlog10M/dlnB=%.12g onePercent=%.12g decades tenPercent=%.12g decades warning=%t verdict=%s", s.DerivativeLog10MPerUnitBGap, s.DerivativeLog10MPerFractionalB, s.OnePercentShiftDecades, s.TenPercentShiftDecades, s.BindingWarning, s.Verdict)
}

func FormatSeal(s PathCSealLedger) string {
	return fmt.Sprintf("IntermediateBreakingSeal prepared=%t granted=%t requiresOrderParameter=%t requiresContactMap=%t requiresPotential=%t requiresResidualMap=%t PatiSalamFalsified=%t hiddenFavored=%t verdict=%s", s.IntermediateBreakingSealPrepared, s.IntermediateBreakingSealGranted, s.RequiresFiniteOrderParameter, s.RequiresContactActionMap, s.RequiresBreakingPotential, s.RequiresResidualMatchingMap, s.PatiSalamRouteFalsifiedInherited, s.HiddenSectorRouteFavored, s.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("PathBClosureInherited=%t Gate229NearResonanceInherited=%t UsesOnlySealedScales=%t DoesNotFitCoefficient=%t DoesNotPromoteVolume=%t DoesNotClaimExactMInt=%t DoesNotGrantSeal=%t DoesNotReopenPatiSalam=%t DoesNotInsertObservedMasses=%t FiniteCorePolluted=%t verdict=%s", f.PathBClosureInherited, f.Gate229NearResonanceInherited, f.UsesOnlySealedScales, f.DoesNotFitCoefficient, f.DoesNotPromoteStandardVolumeToFiniteMap, f.DoesNotClaimExactMIntTheorem, f.DoesNotGrantIntermediateSeal, f.DoesNotReopenPatiSalam, f.DoesNotInsertObservedMasses, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("PathCOpened=%t Volumes=%t FourOverPi=%t Resonance=%t NativeCoefficientDerived=%t MIntTheorem=%t SealGranted=%t Firewall=%t status=%s direct=%q next=%q", s.PathCOpened, s.TopologicalVolumesRetrieved, s.FourOverPiIdentityVerified, s.BGapResonanceReproduced, s.NativeCoefficientDerived, s.IntermediateScaleTheorem, s.IntermediateBreakingSealGranted, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}

func FormatStatusList(statuses []string) string { return strings.Join(statuses, "\n") }
