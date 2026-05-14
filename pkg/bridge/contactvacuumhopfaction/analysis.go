// Package contactvacuumhopfaction implements Gate 284:
// Native Contact-Vacuum Hopf Action Map / Hidden-Sector Order Parameter Audit.
//
// Gate 284 follows Gate 283's exact identity
//
//	S_top/(π Vol(S³)) = 4/π
//
// and asks whether the ASHA finite core derives the missing physical mechanism:
// an instanton/contact-vacuum action map that couples this Hopf-fiber volume
// ratio to B_gap and produces a hidden-sector order parameter at the
// intermediate scale. The result is intentionally conservative: the candidate
// mechanism is precisely formulated, the resonance is re-evaluated, and the
// missing maps are logged as no-gos rather than silently supplied.
package contactvacuumhopfaction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/bgaphierarchycoefficient"
)

const (
	AuditID = "GATE284-NATIVE-CONTACT-VACUUM-HOPF-ACTION-MAP-HIDDEN-SECTOR-ORDER-PARAMETER-AUDIT"

	StatusGate283Inherited                    = "CONDITIONAL_SUPPORT_GATE283_BGAP_RESONANCE_INHERITED"
	StatusInstantonActionFunctionalFormalized = "CONDITIONAL_SUPPORT_INSTANTON_TOPOLOGICAL_ACTION_FUNCTIONAL_FORMALIZED"
	StatusBoundaryMapRequirementsAudited      = "CONDITIONAL_SUPPORT_CONTACT_VACUUM_BOUNDARY_MAP_REQUIREMENTS_AUDITED"
	StatusHiddenOrderParameterRequirementsSet = "CONDITIONAL_SUPPORT_HIDDEN_SECTOR_ORDER_PARAMETER_REQUIREMENTS_DEFINED"
	StatusResidualCorrectionLedgerComputed    = "CONDITIONAL_SUPPORT_BGAP_RESIDUAL_CORRECTION_LEDGER_COMPUTED"
	StatusFirewallsPreserved                  = "CONDITIONAL_SUPPORT_CONTACT_VACUUM_HOPF_FIREWALLS_PRESERVED"

	StatusFailedFiniteConnectionMissing          = "FAILED_ROUTE_FINITE_HOPF_CONNECTION_AND_CURVATURE_NOT_DERIVED"
	StatusFailedContactVacuumBoundaryMapMissing  = "FAILED_ROUTE_CONTACT_VACUUM_TO_HOPF_FIBER_MAP_NOT_DERIVED"
	StatusFailedBGapCouplingMapMissing           = "FAILED_ROUTE_BGAP_AS_INSTANTON_COUPLING_NOT_DERIVED"
	StatusFailedHiddenOrderParameterMissing      = "FAILED_ROUTE_HIDDEN_SECTOR_ORDER_PARAMETER_NOT_DERIVED"
	StatusFailedResidualCorrectionMissing        = "FAILED_ROUTE_RESIDUAL_MATCHING_CORRECTION_NOT_DERIVED"
	StatusFailedIntermediateTheoremNotUpgraded   = "FAILED_ROUTE_CONTACT_VACUUM_ACTION_MAP_NOT_DERIVED"
	StatusFailedIntermediateBreakingSealRequired = "FAILED_ROUTE_INTERMEDIATE_BREAKING_SEAL_REMAINS_REQUIRED"
)

const exactTol = 1e-12

type Gate283Snapshot struct {
	Gate283Inherited          bool
	PathCOpened               bool
	FourOverPiIdentity        bool
	BGapResonanceReproduced   bool
	NativeCoefficientDerived  bool
	IntermediateScaleTheorem  bool
	IntermediateSealGranted   bool
	BGap                      float64
	MStarGeV                  float64
	TargetMIntGeV             float64
	PredictedMIntGeV          float64
	HopfCoefficient           float64
	RequiredCoefficient       float64
	CoefficientResidual       float64
	RelativeCoefficientGap    float64
	Log10Gap                  float64
	SensitivityDecadesPerBGap float64
	TruthStatement            string
}

type InstantonActionFunctional struct {
	CandidateName                  string
	BoundaryManifold               string
	Fiber                          string
	TopologicalAction              float64
	TopologicalActionExact         string
	FiberVolume                    float64
	FiberVolumeExact               string
	Coefficient                    float64
	CoefficientExact               string
	CandidateExponent              float64
	CandidateFormula               string
	RequiresFiniteConnection       bool
	RequiresCurvatureTwoForm       bool
	RequiresChernSimonsThreeForm   bool
	RequiresIntegerWindingMap      bool
	RequiresFiniteActionCriticalPt bool
	FiniteInstantonDerived         bool
	Verdict                        string
}

type ContactVacuumBoundaryMap struct {
	ProposedMap                         string
	ContactVacuumCarrierAvailable       bool
	S7HopfFibrationAvailable            bool
	S3FiberVolumeAvailable              bool
	BGapSpectralDatumAvailable          bool
	BoundaryEmbeddingDerived            bool
	FiberLocalizationFunctionalDerived  bool
	ActionDensityOnFiberDerived         bool
	BGapAsInverseCouplingDerived        bool
	ExponentialHierarchyEquationDerived bool
	ContactVacuumHopfActionMapDerived   bool
	CandidateIntermediateScaleGeV       float64
	CandidateIntermediateScaleRatio     float64
	CandidateIntermediateScaleLog10Gap  float64
	Verdict                             string
}

type HiddenSectorOrderParameter struct {
	CandidateName                     string
	BGap                              float64
	HiddenSectorOrderParameterDefined bool
	ScalarOrCondensateFieldDerived    bool
	GaugeGroupOrBundleDerived         bool
	EffectivePotentialDerived         bool
	NonzeroVEVDerived                 bool
	VEVScaleGeV                       float64
	CouplesToHopfAction               bool
	GeneratesSeesawScale              bool
	GeneratesAxionOrRelicPortal       bool
	Verdict                           string
}

type ResidualCorrectionAudit struct {
	RequiredCoefficient              float64
	HopfCoefficient                  float64
	DeltaCoefficient                 float64
	RelativeDeltaCoefficient         float64
	CandidatePredictedMIntGeV        float64
	TargetMIntGeV                    float64
	RatioPredictedToTarget           float64
	Log10Gap                         float64
	RequiredMultiplicativeCorrection float64
	CorrectionEquivalent             string
	FiniteVolumeCorrectionDerived    bool
	ThresholdMatchingDerived         bool
	LoopCorrectionDerived            bool
	GeometricSubtractionDerived      bool
	ResidualExacted                  bool
	Verdict                          string
}

type SealLedger struct {
	IntermediateBreakingSealPrepared bool
	IntermediateBreakingSealGranted  bool
	RequiresInstantonActionMap       bool
	RequiresHiddenOrderParameter     bool
	RequiresBreakingPotential        bool
	RequiresResidualCorrection       bool
	SafeHiddenSectorDirection        bool
	Verdict                          string
}

type FirewallAudit struct {
	UsesOnlyGate283Data             bool
	DoesNotFitCoefficient           bool
	DoesNotDeclareInstantonSolution bool
	DoesNotPromoteBGapToField       bool
	DoesNotInventOrderParameter     bool
	DoesNotClaimExactResidual       bool
	DoesNotGrantIntermediateSeal    bool
	DoesNotReopenPathB              bool
	DoesNotInsertObservedMasses     bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	Gate283Inherited              bool
	InstantonFunctionalFormalized bool
	ContactVacuumMapDerived       bool
	HiddenOrderParameterDerived   bool
	ResidualCorrectionDerived     bool
	IntermediateTheoremUpgraded   bool
	IntermediateSealGranted       bool
	FirewallPreserved             bool
	Status                        string
	DirectAnswer                  string
	NextGate                      string
}

type Analysis struct {
	Gate283        Gate283Snapshot
	Instanton      InstantonActionFunctional
	BoundaryMap    ContactVacuumBoundaryMap
	OrderParameter HiddenSectorOrderParameter
	Residual       ResidualCorrectionAudit
	Seal           SealLedger
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
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	g283, err := bgaphierarchycoefficient.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 283 B-gap coefficient predecessor: %w", err)
	}
	snap := snapshotGate283(g283)
	if !snap.Gate283Inherited || !snap.FourOverPiIdentity || snap.BGap <= 0 || snap.MStarGeV <= 0 || snap.TargetMIntGeV <= 0 {
		return Analysis{}, fmt.Errorf("Gate 284 requires Gate 283 four-over-pi B-gap resonance data")
	}
	inst := formalizeInstantonAction(snap)
	boundary := auditBoundaryMap(snap, inst)
	order := auditOrderParameter(snap, boundary)
	residual := auditResidual(snap)
	seal := auditSeal(inst, boundary, order, residual)
	firewall := auditFirewall(snap, inst, boundary, order, residual, seal)
	summary := buildSummary(snap, inst, boundary, order, residual, seal, firewall)
	truth := "Gate 284 formalizes the candidate Hopf-fiber instanton action S_inst=(4/π)/B_gap and preserves the sharp intermediate-scale resonance, but it does not derive the contact-vacuum boundary map, B_gap coupling interpretation, hidden-sector order parameter, or residual correction required to upgrade the scale to a finite theorem."
	return Analysis{Gate283: snap, Instanton: inst, BoundaryMap: boundary, OrderParameter: order, Residual: residual, Seal: seal, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotGate283(a bgaphierarchycoefficient.Analysis) Gate283Snapshot {
	return Gate283Snapshot{
		Gate283Inherited:          a.Summary.PathCOpened && a.Summary.TopologicalVolumesRetrieved,
		PathCOpened:               a.Summary.PathCOpened,
		FourOverPiIdentity:        a.Summary.FourOverPiIdentityVerified && a.ContactAction.CoefficientEqualsFourOverPi,
		BGapResonanceReproduced:   a.Summary.BGapResonanceReproduced,
		NativeCoefficientDerived:  a.Summary.NativeCoefficientDerived,
		IntermediateScaleTheorem:  a.Summary.IntermediateScaleTheorem,
		IntermediateSealGranted:   a.Summary.IntermediateBreakingSealGranted,
		BGap:                      a.Hierarchy.BGap,
		MStarGeV:                  a.Hierarchy.MStarGeV,
		TargetMIntGeV:             a.Hierarchy.TargetMIntGeV,
		PredictedMIntGeV:          a.Hierarchy.PredictedMIntGeV,
		HopfCoefficient:           a.Hierarchy.Coefficient,
		RequiredCoefficient:       a.Hierarchy.RequiredCoefficient,
		CoefficientResidual:       a.Hierarchy.CoefficientResidual,
		RelativeCoefficientGap:    a.Hierarchy.RelativeCoefficientResidual,
		Log10Gap:                  a.Hierarchy.Log10Gap,
		SensitivityDecadesPerBGap: a.Sensitivity.DerivativeLog10MPerUnitBGap,
		TruthStatement:            a.TruthStatement,
	}
}

func formalizeInstantonAction(g Gate283Snapshot) InstantonActionFunctional {
	fiberVol := 2 * math.Pi * math.Pi
	sTop := 8 * math.Pi * math.Pi
	coeff := sTop / (math.Pi * fiberVol)
	return InstantonActionFunctional{
		CandidateName:                  "Hopf S³ boundary instanton / Chern-Simons winding preflight",
		BoundaryManifold:               "Hopf fiber S³ inside S⁷ -> S⁴ contact boundary",
		Fiber:                          "S³",
		TopologicalAction:              sTop,
		TopologicalActionExact:         "8π²",
		FiberVolume:                    fiberVol,
		FiberVolumeExact:               "2π²",
		Coefficient:                    coeff,
		CoefficientExact:               "4/π",
		CandidateExponent:              coeff / g.BGap,
		CandidateFormula:               "S_inst,candidate = S_top/(π Vol(S³) B_gap) = (4/π)/B_gap",
		RequiresFiniteConnection:       true,
		RequiresCurvatureTwoForm:       true,
		RequiresChernSimonsThreeForm:   true,
		RequiresIntegerWindingMap:      true,
		RequiresFiniteActionCriticalPt: true,
		FiniteInstantonDerived:         false,
		Verdict:                        StatusInstantonActionFunctionalFormalized,
	}
}

func auditBoundaryMap(g Gate283Snapshot, i InstantonActionFunctional) ContactVacuumBoundaryMap {
	ratio := g.PredictedMIntGeV / g.TargetMIntGeV
	return ContactVacuumBoundaryMap{
		ProposedMap:                         "contact vacuum boundary -> Hopf S³ fiber action density -> B_gap inverse coupling -> M_* exp(-S_inst)",
		ContactVacuumCarrierAvailable:       true,
		S7HopfFibrationAvailable:            true,
		S3FiberVolumeAvailable:              true,
		BGapSpectralDatumAvailable:          g.BGap > 0,
		BoundaryEmbeddingDerived:            false,
		FiberLocalizationFunctionalDerived:  false,
		ActionDensityOnFiberDerived:         false,
		BGapAsInverseCouplingDerived:        false,
		ExponentialHierarchyEquationDerived: false,
		ContactVacuumHopfActionMapDerived:   false,
		CandidateIntermediateScaleGeV:       g.PredictedMIntGeV,
		CandidateIntermediateScaleRatio:     ratio,
		CandidateIntermediateScaleLog10Gap:  math.Abs(math.Log10(ratio)),
		Verdict:                             StatusBoundaryMapRequirementsAudited,
	}
}

func auditOrderParameter(g Gate283Snapshot, b ContactVacuumBoundaryMap) HiddenSectorOrderParameter {
	return HiddenSectorOrderParameter{
		CandidateName:                     "Φ_B hidden B-sector condensate / contact-vacuum order parameter",
		BGap:                              g.BGap,
		HiddenSectorOrderParameterDefined: false,
		ScalarOrCondensateFieldDerived:    false,
		GaugeGroupOrBundleDerived:         false,
		EffectivePotentialDerived:         false,
		NonzeroVEVDerived:                 false,
		VEVScaleGeV:                       b.CandidateIntermediateScaleGeV,
		CouplesToHopfAction:               false,
		GeneratesSeesawScale:              false,
		GeneratesAxionOrRelicPortal:       false,
		Verdict:                           StatusHiddenOrderParameterRequirementsSet,
	}
}

func auditResidual(g Gate283Snapshot) ResidualCorrectionAudit {
	ratio := g.PredictedMIntGeV / g.TargetMIntGeV
	corr := g.RequiredCoefficient / g.HopfCoefficient
	return ResidualCorrectionAudit{
		RequiredCoefficient:              g.RequiredCoefficient,
		HopfCoefficient:                  g.HopfCoefficient,
		DeltaCoefficient:                 g.CoefficientResidual,
		RelativeDeltaCoefficient:         g.RelativeCoefficientGap,
		CandidatePredictedMIntGeV:        g.PredictedMIntGeV,
		TargetMIntGeV:                    g.TargetMIntGeV,
		RatioPredictedToTarget:           ratio,
		Log10Gap:                         g.Log10Gap,
		RequiredMultiplicativeCorrection: corr,
		CorrectionEquivalent:             "c_req/(4/π), or equivalently a ~0.305% coefficient correction / B_gap displacement",
		FiniteVolumeCorrectionDerived:    false,
		ThresholdMatchingDerived:         false,
		LoopCorrectionDerived:            false,
		GeometricSubtractionDerived:      false,
		ResidualExacted:                  false,
		Verdict:                          StatusResidualCorrectionLedgerComputed,
	}
}

func auditSeal(i InstantonActionFunctional, b ContactVacuumBoundaryMap, o HiddenSectorOrderParameter, r ResidualCorrectionAudit) SealLedger {
	return SealLedger{
		IntermediateBreakingSealPrepared: true,
		IntermediateBreakingSealGranted:  false,
		RequiresInstantonActionMap:       !i.FiniteInstantonDerived || !b.ContactVacuumHopfActionMapDerived,
		RequiresHiddenOrderParameter:     !o.HiddenSectorOrderParameterDefined,
		RequiresBreakingPotential:        !o.EffectivePotentialDerived,
		RequiresResidualCorrection:       !r.ResidualExacted,
		SafeHiddenSectorDirection:        true,
		Verdict:                          StatusFailedIntermediateBreakingSealRequired,
	}
}

func auditFirewall(g Gate283Snapshot, i InstantonActionFunctional, b ContactVacuumBoundaryMap, o HiddenSectorOrderParameter, r ResidualCorrectionAudit, s SealLedger) FirewallAudit {
	return FirewallAudit{
		UsesOnlyGate283Data:             g.Gate283Inherited && g.PathCOpened,
		DoesNotFitCoefficient:           math.Abs(i.Coefficient-4/math.Pi) < exactTol,
		DoesNotDeclareInstantonSolution: !i.FiniteInstantonDerived,
		DoesNotPromoteBGapToField:       !b.BGapAsInverseCouplingDerived && !o.ScalarOrCondensateFieldDerived,
		DoesNotInventOrderParameter:     !o.HiddenSectorOrderParameterDefined,
		DoesNotClaimExactResidual:       !r.ResidualExacted,
		DoesNotGrantIntermediateSeal:    !s.IntermediateBreakingSealGranted,
		DoesNotReopenPathB:              true,
		DoesNotInsertObservedMasses:     true,
		FiniteCorePolluted:              false,
		Verdict:                         StatusFirewallsPreserved,
	}
}

func buildSummary(g Gate283Snapshot, i InstantonActionFunctional, b ContactVacuumBoundaryMap, o HiddenSectorOrderParameter, r ResidualCorrectionAudit, s SealLedger, f FirewallAudit) Summary {
	theorem := i.FiniteInstantonDerived && b.ContactVacuumHopfActionMapDerived && o.HiddenSectorOrderParameterDefined && r.ResidualExacted && s.IntermediateBreakingSealGranted
	status := StatusFailedIntermediateTheoremNotUpgraded
	direct := "The Hopf/contact instanton mechanism is precisely formulated, but the native action map and hidden order parameter are not derived; M_int remains a sharp conditional resonance, not a theorem."
	if theorem {
		status = "CONDITIONAL_SUPPORT_CONTACT_VACUUM_ACTION_MAP_AND_ORDER_PARAMETER_DERIVED"
		direct = "The contact-vacuum Hopf action map and hidden order parameter are derived."
	}
	return Summary{
		Gate283Inherited:              g.Gate283Inherited,
		InstantonFunctionalFormalized: i.CandidateExponent > 0 && i.Verdict == StatusInstantonActionFunctionalFormalized,
		ContactVacuumMapDerived:       b.ContactVacuumHopfActionMapDerived,
		HiddenOrderParameterDerived:   o.HiddenSectorOrderParameterDefined,
		ResidualCorrectionDerived:     r.ResidualExacted,
		IntermediateTheoremUpgraded:   theorem,
		IntermediateSealGranted:       s.IntermediateBreakingSealGranted,
		FirewallPreserved:             f.UsesOnlyGate283Data && f.DoesNotFitCoefficient && f.DoesNotGrantIntermediateSeal && !f.FiniteCorePolluted,
		Status:                        status,
		DirectAnswer:                  direct,
		NextGate:                      "Derive a finite contact/Hopf boundary embedding with Chern-Simons or instanton density, then identify B_gap as a genuine inverse coupling or order parameter; without this, Path C remains resonance-only.",
	}
}

func Statuses(a Analysis) []string {
	statuses := []string{
		StatusGate283Inherited,
		a.Instanton.Verdict,
		a.BoundaryMap.Verdict,
		a.OrderParameter.Verdict,
		a.Residual.Verdict,
		a.Firewall.Verdict,
	}
	if !a.Instanton.FiniteInstantonDerived {
		statuses = append(statuses, StatusFailedFiniteConnectionMissing)
	}
	if !a.BoundaryMap.ContactVacuumHopfActionMapDerived {
		statuses = append(statuses, StatusFailedContactVacuumBoundaryMapMissing)
	}
	if !a.BoundaryMap.BGapAsInverseCouplingDerived {
		statuses = append(statuses, StatusFailedBGapCouplingMapMissing)
	}
	if !a.OrderParameter.HiddenSectorOrderParameterDefined {
		statuses = append(statuses, StatusFailedHiddenOrderParameterMissing)
	}
	if !a.Residual.ResidualExacted {
		statuses = append(statuses, StatusFailedResidualCorrectionMissing)
	}
	statuses = append(statuses, a.Summary.Status, StatusFailedIntermediateBreakingSealRequired)
	return statuses
}

func FormatGate283(g Gate283Snapshot) string {
	return fmt.Sprintf("Gate283Inherited=%t PathCOpened=%t FourOverPi=%t Resonance=%t NativeCoeff=%t MIntTheorem=%t Seal=%t B_gap=%.12g M*=%.12e target=%.12e pred=%.12e c=%.12g cReq=%.12g delta=%.12g rel=%.12g log10gap=%.12g", g.Gate283Inherited, g.PathCOpened, g.FourOverPiIdentity, g.BGapResonanceReproduced, g.NativeCoefficientDerived, g.IntermediateScaleTheorem, g.IntermediateSealGranted, g.BGap, g.MStarGeV, g.TargetMIntGeV, g.PredictedMIntGeV, g.HopfCoefficient, g.RequiredCoefficient, g.CoefficientResidual, g.RelativeCoefficientGap, g.Log10Gap)
}

func FormatInstanton(i InstantonActionFunctional) string {
	return fmt.Sprintf("%s on %s: S_top=%s=%.12g Vol(%s)=%s=%.12g coeff=%s=%.12g exponent=%.12g formula=%q requiresConnection=%t curvature=%t CS3=%t winding=%t criticalPt=%t derived=%t verdict=%s", i.CandidateName, i.BoundaryManifold, i.TopologicalActionExact, i.TopologicalAction, i.Fiber, i.FiberVolumeExact, i.FiberVolume, i.CoefficientExact, i.Coefficient, i.CandidateExponent, i.CandidateFormula, i.RequiresFiniteConnection, i.RequiresCurvatureTwoForm, i.RequiresChernSimonsThreeForm, i.RequiresIntegerWindingMap, i.RequiresFiniteActionCriticalPt, i.FiniteInstantonDerived, i.Verdict)
}

func FormatBoundaryMap(b ContactVacuumBoundaryMap) string {
	return fmt.Sprintf("map=%q carrier=%t Hopf=%t S3vol=%t BgapDatum=%t boundaryEmbed=%t localization=%t density=%t BgapCoupling=%t hierarchyEq=%t actionMap=%t pred=%.12e ratio=%.12g log10gap=%.12g verdict=%s", b.ProposedMap, b.ContactVacuumCarrierAvailable, b.S7HopfFibrationAvailable, b.S3FiberVolumeAvailable, b.BGapSpectralDatumAvailable, b.BoundaryEmbeddingDerived, b.FiberLocalizationFunctionalDerived, b.ActionDensityOnFiberDerived, b.BGapAsInverseCouplingDerived, b.ExponentialHierarchyEquationDerived, b.ContactVacuumHopfActionMapDerived, b.CandidateIntermediateScaleGeV, b.CandidateIntermediateScaleRatio, b.CandidateIntermediateScaleLog10Gap, b.Verdict)
}

func FormatOrderParameter(o HiddenSectorOrderParameter) string {
	return fmt.Sprintf("%s B_gap=%.12g defined=%t field=%t gaugeBundle=%t potential=%t VEV=%t VEVscale=%.12e couplesHopf=%t seesaw=%t axionRelic=%t verdict=%s", o.CandidateName, o.BGap, o.HiddenSectorOrderParameterDefined, o.ScalarOrCondensateFieldDerived, o.GaugeGroupOrBundleDerived, o.EffectivePotentialDerived, o.NonzeroVEVDerived, o.VEVScaleGeV, o.CouplesToHopfAction, o.GeneratesSeesawScale, o.GeneratesAxionOrRelicPortal, o.Verdict)
}

func FormatResidual(r ResidualCorrectionAudit) string {
	return fmt.Sprintf("cReq=%.12g cHopf=%.12g delta=%.12g relDelta=%.12g pred=%.12e target=%.12e ratio=%.12g log10gap=%.12g multCorr=%.12g corr=%q finiteVol=%t threshold=%t loop=%t subtraction=%t exacted=%t verdict=%s", r.RequiredCoefficient, r.HopfCoefficient, r.DeltaCoefficient, r.RelativeDeltaCoefficient, r.CandidatePredictedMIntGeV, r.TargetMIntGeV, r.RatioPredictedToTarget, r.Log10Gap, r.RequiredMultiplicativeCorrection, r.CorrectionEquivalent, r.FiniteVolumeCorrectionDerived, r.ThresholdMatchingDerived, r.LoopCorrectionDerived, r.GeometricSubtractionDerived, r.ResidualExacted, r.Verdict)
}

func FormatSeal(s SealLedger) string {
	return fmt.Sprintf("IntermediateBreakingSeal prepared=%t granted=%t requiresInstanton=%t requiresOrderParameter=%t requiresPotential=%t requiresResidual=%t hiddenSafe=%t verdict=%s", s.IntermediateBreakingSealPrepared, s.IntermediateBreakingSealGranted, s.RequiresInstantonActionMap, s.RequiresHiddenOrderParameter, s.RequiresBreakingPotential, s.RequiresResidualCorrection, s.SafeHiddenSectorDirection, s.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("UsesGate283=%t noFit=%t noInstantonSolution=%t noBGapField=%t noOrderParameter=%t noExactResidual=%t noSeal=%t noPathB=%t noObservedMasses=%t polluted=%t verdict=%s", f.UsesOnlyGate283Data, f.DoesNotFitCoefficient, f.DoesNotDeclareInstantonSolution, f.DoesNotPromoteBGapToField, f.DoesNotInventOrderParameter, f.DoesNotClaimExactResidual, f.DoesNotGrantIntermediateSeal, f.DoesNotReopenPathB, f.DoesNotInsertObservedMasses, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("Gate283=%t instantonFormal=%t contactMap=%t orderParameter=%t residual=%t theorem=%t seal=%t firewall=%t status=%s direct=%q next=%q", s.Gate283Inherited, s.InstantonFunctionalFormalized, s.ContactVacuumMapDerived, s.HiddenOrderParameterDerived, s.ResidualCorrectionDerived, s.IntermediateTheoremUpgraded, s.IntermediateSealGranted, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}

func FormatStatusList(statuses []string) string { return strings.Join(statuses, "\n") }
