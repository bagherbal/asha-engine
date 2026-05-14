// Package finitehopfaction implements Gate 230: Octonionic Instanton /
// finite Hopf-action map and hidden order-parameter audit.
//
// Gate 229 found a sharp conditional resonance
//
//	M_Hopf = M_* exp(-(4/π)/B_gap),
//
// where 4/π = S_top/(π Vol(S^3)). Gate 230 asks the harder question:
// whether the finite Clifford/G2/contact core actually derives the dynamical
// instanton equation, Hopf-fiber action localization, and hidden order
// parameter needed to promote that resonance into an IntermediateBreakingSeal.
//
// The answer is intentionally strict: the geometric ingredients are real, but
// the dynamic map is not yet present. The route is logged as an obstruction,
// not repaired by hand.
package finitehopfaction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fourcyclechernweil"
	"github.com/bagherbal/asha-engine/pkg/bridge/hopfgeometricnormalization"
)

const (
	AuditID = "GATE230-OCTONIONIC-INSTANTON-FINITE-HOPF-ACTION-MAP-AUDIT"

	StatusFiniteInstantonFailed       = "FAILED_ROUTE_FINITE_INSTANTON_DERIVATION"
	StatusHopfActionMapFailed         = "FAILED_ROUTE_HOPF_ACTION_LOCALIZATION_MAP"
	StatusHiddenOrderParameterFailed  = "FAILED_ROUTE_HIDDEN_ORDER_PARAMETER_DERIVATION"
	StatusGeometricResonanceInherited = "CONDITIONAL_SUPPORT_HOPF_GEOMETRIC_RESONANCE_INHERITED"
	StatusIntermediateSealNotGranted  = "INTERMEDIATE_BREAKING_SEAL_NOT_GRANTED"
)

type Gate229Snapshot struct {
	Gate229Inherited            bool
	GeometricHierarchySupported bool
	NativeHopfMapDerived        bool
	IntermediateSealGranted     bool
	SensitivityBinding          bool
	ResidualPlausiblyCovered    bool
	MIntTargetGeV               float64
	MIntHopfGeV                 float64
	MStarGeV                    float64
	BGap                        float64
	HopfCoefficient             float64
	RequiredCoefficient         float64
	TopologicalActionNumerator  float64
	HopfFiberVolume             float64
	CoefficientEqualsFourOverPi bool
	TruthStatement              string
}

type ChernWeilSnapshot struct {
	Gate181Inherited               bool
	GaugeAlgebraClosed             bool
	RepresentationTraceRatioClosed bool
	TopologicalSealAvailable       bool
	PrincipalBundleDerived         bool
	ConnectionOnFourCarrierDerived bool
	CurvatureTwoFormDerived        bool
	TracePairingDerived            bool
	IntegralTrFedgeFDerived        bool
	IntegerInstantonNumberDerived  bool
	ContinuumNormalizationPromoted bool
	InstantonBridgePromoted        bool
	TruthStatement                 string
}

type OctonionicInstantonAudit struct {
	CandidateEquation                  string
	GeometryCarrier                    string
	G2ContactPredataAvailable          bool
	CliffordOctonionicPredataAvailable bool
	TopologicalActionSealAvailable     bool
	BGapAvailable                      bool
	PrincipalBundleDerived             bool
	GaugeConnectionDerived             bool
	CurvatureTwoFormDerived            bool
	G2SelfDualityProjectorDerived      bool
	FiniteYangMillsActionDerived       bool
	BPSOrCriticalPointEquationDerived  bool
	NontrivialFiniteSolutionDerived    bool
	IntegerTopologicalChargeDerived    bool
	OctonionicInstantonDerived         bool
	Verdict                            string
}

type HopfFiberActionMapAudit struct {
	Formula                              string
	TopologicalActionNumerator           float64
	HopfFiberVolume                      float64
	HopfCoefficient                      float64
	TargetMIntGeV                        float64
	HopfMIntGeV                          float64
	Log10Gap                             float64
	S7HopfFibrationStandardMathAvailable bool
	S3FiberVolumeStandardMathAvailable   bool
	ContactVacuumToS7MapDerived          bool
	FiberLocalizationFunctionalDerived   bool
	ActionDensityOnFiberDerived          bool
	BGapAsInstantonCouplingDerived       bool
	HopfActionMapDerived                 bool
	ConditionalShapeSupported            bool
	Verdict                              string
}

type HiddenOrderParameterAudit struct {
	CandidateName                 string
	BGap                          float64
	ScalarSpectralAnchorAvailable bool
	ContinuousFieldDerived        bool
	LocalEffectiveActionDerived   bool
	PotentialDerived              bool
	VEVAtHopfScaleDerived         bool
	ShiftSymmetryBreakingDerived  bool
	AxionOrEFTScaleGenerated      bool
	HiddenOrderParameterDerived   bool
	Verdict                       string
}

type IntermediateBreakingSealAudit struct {
	SealName                     string
	PreviouslyPrepared           bool
	Granted                      bool
	RequiredInstantonDerived     bool
	RequiredHopfActionMapDerived bool
	RequiredOrderParameter       bool
	GeometricResonanceInherited  bool
	OperationalStatus            string
	Verdict                      string
}

type FirewallAudit struct {
	UsedOnlySealedInputs            bool
	ObservedInputsIntroduced        bool
	PatiSalamReopened               bool
	LeptoquarkDynamicsReopened      bool
	InstatonEquationInvented        bool
	DFOrConnectionInvented          bool
	HopfActionNormalizationFitted   bool
	BGapPromotedToPhysicalField     bool
	HiddenVEVInvented               bool
	IntermediateBreakingSealGranted bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type Summary struct {
	GeometricResonanceInherited bool
	FiniteInstantonDerived      bool
	HopfActionMapDerived        bool
	HiddenOrderParameterDerived bool
	IntermediateSealGranted     bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	Gate229        Gate229Snapshot
	ChernWeil      ChernWeilSnapshot
	Instanton      OctonionicInstantonAudit
	HopfAction     HopfFiberActionMapAudit
	OrderParameter HiddenOrderParameterAudit
	Seal           IntermediateBreakingSealAudit
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
		g229, err := hopfgeometricnormalization.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 229 input: %w", err)
			return
		}
		g181, err := fourcyclechernweil.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 181 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g229, g181)
	})
	return defaultA, defaultErr
}

func Build(g229 hopfgeometricnormalization.Analysis, g181 fourcyclechernweil.Analysis) (Analysis, error) {
	s229 := snapshotFromGate229(g229)
	if !s229.Gate229Inherited || !s229.GeometricHierarchySupported || s229.MIntTargetGeV <= 0 || s229.MIntHopfGeV <= 0 || s229.BGap <= 0 {
		return Analysis{}, fmt.Errorf("Gate 230 requires Gate 229 Hopf-resonance data and positive B-gap hierarchy scales")
	}
	cw := snapshotFromChernWeil(g181)
	inst := auditOctonionicInstanton(s229, cw)
	hopf := auditHopfAction(s229, inst)
	order := auditOrderParameter(s229, hopf)
	seal := auditSeal(s229, inst, hopf, order)
	firewall := auditFirewall(s229, inst, hopf, order, seal)
	summary := summarize(s229, inst, hopf, order, seal)
	truth := buildTruth(s229, inst, hopf, order, seal)
	return Analysis{Gate229: s229, ChernWeil: cw, Instanton: inst, HopfAction: hopf, OrderParameter: order, Seal: seal, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate229(a hopfgeometricnormalization.Analysis) Gate229Snapshot {
	return Gate229Snapshot{
		Gate229Inherited:            a.Summary.Status != "" && a.Hierarchy.BGap > 0,
		GeometricHierarchySupported: a.Summary.GeometricHierarchySupported,
		NativeHopfMapDerived:        a.Summary.NativeHopfMapDerived,
		IntermediateSealGranted:     a.Summary.IntermediateSealGranted,
		SensitivityBinding:          a.Summary.SensitivityBinding,
		ResidualPlausiblyCovered:    a.Summary.ResidualPlausiblyCovered,
		MIntTargetGeV:               a.Summary.MIntTargetGeV,
		MIntHopfGeV:                 a.Summary.MIntHopfGeV,
		MStarGeV:                    a.Hierarchy.MStarGeV,
		BGap:                        a.Hierarchy.BGap,
		HopfCoefficient:             a.Summary.Coefficient,
		RequiredCoefficient:         a.Summary.RequiredCoefficient,
		TopologicalActionNumerator:  a.Geometry.TopologicalActionNumerator,
		HopfFiberVolume:             a.Geometry.UnitS3Volume,
		CoefficientEqualsFourOverPi: a.Geometry.CoefficientEqualsFourOverPi,
		TruthStatement:              a.TruthStatement,
	}
}

func snapshotFromChernWeil(a fourcyclechernweil.Analysis) ChernWeilSnapshot {
	return ChernWeilSnapshot{
		Gate181Inherited:               a.Firewall.RecommendedNextGate != "" || a.ChernWeil.TopologicalSealAvailable,
		GaugeAlgebraClosed:             a.ChernWeil.GaugeAlgebraClosed,
		RepresentationTraceRatioClosed: a.ChernWeil.RepresentationTraceRatioClosed,
		TopologicalSealAvailable:       a.ChernWeil.TopologicalSealAvailable,
		PrincipalBundleDerived:         a.ChernWeil.PrincipalBundleDerived,
		ConnectionOnFourCarrierDerived: a.ChernWeil.ConnectionOnFourCarrierDerived,
		CurvatureTwoFormDerived:        a.ChernWeil.CurvatureTwoFormDerived,
		TracePairingDerived:            a.ChernWeil.TracePairingDerived,
		IntegralTrFedgeFDerived:        a.ChernWeil.IntegralOfTrFedgeFDerived,
		IntegerInstantonNumberDerived:  a.ChernWeil.IntegerInstantonNumberDerived,
		ContinuumNormalizationPromoted: a.ChernWeil.ContinuumNormalizationPromoted,
		InstantonBridgePromoted:        a.Firewall.InstantonTraceBridgeDerived,
		TruthStatement:                 a.TruthStatement,
	}
}

func auditOctonionicInstanton(g Gate229Snapshot, cw ChernWeilSnapshot) OctonionicInstantonAudit {
	derived := cw.PrincipalBundleDerived && cw.ConnectionOnFourCarrierDerived && cw.CurvatureTwoFormDerived && cw.IntegerInstantonNumberDerived
	return OctonionicInstantonAudit{
		CandidateEquation:                  "G2 instanton preflight: F_A ∧ *φ = 0, equivalently F_A ∈ Λ²_14; Spin(7)/octonionic analogue would require a derived connection and curvature",
		GeometryCarrier:                    "Cℓ(1,7) / G2 contact vacuum with B-sector spectral gap",
		G2ContactPredataAvailable:          true,
		CliffordOctonionicPredataAvailable: true,
		TopologicalActionSealAvailable:     g.TopologicalActionNumerator > 0 && cw.TopologicalSealAvailable,
		BGapAvailable:                      g.BGap > 0,
		PrincipalBundleDerived:             cw.PrincipalBundleDerived,
		GaugeConnectionDerived:             cw.ConnectionOnFourCarrierDerived,
		CurvatureTwoFormDerived:            cw.CurvatureTwoFormDerived,
		G2SelfDualityProjectorDerived:      false,
		FiniteYangMillsActionDerived:       false,
		BPSOrCriticalPointEquationDerived:  false,
		NontrivialFiniteSolutionDerived:    false,
		IntegerTopologicalChargeDerived:    cw.IntegerInstantonNumberDerived,
		OctonionicInstantonDerived:         derived && false,
		Verdict:                            StatusFiniteInstantonFailed,
	}
}

func auditHopfAction(g Gate229Snapshot, inst OctonionicInstantonAudit) HopfFiberActionMapAudit {
	logGap := math.Abs(math.Log10(g.MIntHopfGeV / g.MIntTargetGeV))
	derived := inst.OctonionicInstantonDerived && false
	return HopfFiberActionMapAudit{
		Formula:                              "S_inst = S_top/(π Vol(S^3) B_gap); M = M_* exp(-S_inst)",
		TopologicalActionNumerator:           g.TopologicalActionNumerator,
		HopfFiberVolume:                      g.HopfFiberVolume,
		HopfCoefficient:                      g.HopfCoefficient,
		TargetMIntGeV:                        g.MIntTargetGeV,
		HopfMIntGeV:                          g.MIntHopfGeV,
		Log10Gap:                             logGap,
		S7HopfFibrationStandardMathAvailable: true,
		S3FiberVolumeStandardMathAvailable:   g.HopfFiberVolume > 0,
		ContactVacuumToS7MapDerived:          false,
		FiberLocalizationFunctionalDerived:   false,
		ActionDensityOnFiberDerived:          false,
		BGapAsInstantonCouplingDerived:       false,
		HopfActionMapDerived:                 derived,
		ConditionalShapeSupported:            g.GeometricHierarchySupported && g.CoefficientEqualsFourOverPi && logGap < 0.02,
		Verdict:                              StatusHopfActionMapFailed,
	}
}

func auditOrderParameter(g Gate229Snapshot, h HopfFiberActionMapAudit) HiddenOrderParameterAudit {
	derived := h.HopfActionMapDerived && false
	return HiddenOrderParameterAudit{
		CandidateName:                 "Φ_B / hidden B-sector order parameter",
		BGap:                          g.BGap,
		ScalarSpectralAnchorAvailable: g.BGap > 0,
		ContinuousFieldDerived:        false,
		LocalEffectiveActionDerived:   false,
		PotentialDerived:              false,
		VEVAtHopfScaleDerived:         false,
		ShiftSymmetryBreakingDerived:  false,
		AxionOrEFTScaleGenerated:      false,
		HiddenOrderParameterDerived:   derived,
		Verdict:                       StatusHiddenOrderParameterFailed,
	}
}

func auditSeal(g Gate229Snapshot, inst OctonionicInstantonAudit, h HopfFiberActionMapAudit, o HiddenOrderParameterAudit) IntermediateBreakingSealAudit {
	granted := inst.OctonionicInstantonDerived && h.HopfActionMapDerived && o.HiddenOrderParameterDerived
	status := "SEAL_PREPARED_NOT_GRANTED"
	if granted {
		status = "SEAL_GRANTED"
	}
	return IntermediateBreakingSealAudit{
		SealName:                     "IntermediateBreakingSeal",
		PreviouslyPrepared:           true,
		Granted:                      granted,
		RequiredInstantonDerived:     inst.OctonionicInstantonDerived,
		RequiredHopfActionMapDerived: h.HopfActionMapDerived,
		RequiredOrderParameter:       o.HiddenOrderParameterDerived,
		GeometricResonanceInherited:  g.GeometricHierarchySupported,
		OperationalStatus:            status,
		Verdict:                      StatusIntermediateSealNotGranted,
	}
}

func auditFirewall(g Gate229Snapshot, inst OctonionicInstantonAudit, h HopfFiberActionMapAudit, o HiddenOrderParameterAudit, s IntermediateBreakingSealAudit) FirewallAudit {
	return FirewallAudit{
		UsedOnlySealedInputs:            true,
		ObservedInputsIntroduced:        false,
		PatiSalamReopened:               false,
		LeptoquarkDynamicsReopened:      false,
		InstatonEquationInvented:        false,
		DFOrConnectionInvented:          false,
		HopfActionNormalizationFitted:   false,
		BGapPromotedToPhysicalField:     o.ContinuousFieldDerived,
		HiddenVEVInvented:               false,
		IntermediateBreakingSealGranted: s.Granted,
		FiniteCorePolluted:              false,
		Verdict:                         "FIREWALLS_CLOSED",
	}
}

func summarize(g Gate229Snapshot, inst OctonionicInstantonAudit, h HopfFiberActionMapAudit, o HiddenOrderParameterAudit, s IntermediateBreakingSealAudit) Summary {
	statuses := []string{StatusGeometricResonanceInherited, StatusFiniteInstantonFailed, StatusHopfActionMapFailed, StatusHiddenOrderParameterFailed, StatusIntermediateSealNotGranted}
	return Summary{
		GeometricResonanceInherited: g.GeometricHierarchySupported && h.ConditionalShapeSupported,
		FiniteInstantonDerived:      inst.OctonionicInstantonDerived,
		HopfActionMapDerived:        h.HopfActionMapDerived,
		HiddenOrderParameterDerived: o.HiddenOrderParameterDerived,
		IntermediateSealGranted:     s.Granted,
		Status:                      strings.Join(statuses, ";"),
		NextGate:                    "Gate 231 — hidden-sector order-parameter seal or alternate dark/intermediate-origin audit",
		Comment:                     "Gate 230 inherits the strong Hopf/B-gap hierarchy resonance but fails to derive the dynamical instanton, Hopf-fiber localization map, or hidden order parameter required to grant the IntermediateBreakingSeal.",
	}
}

func buildTruth(g Gate229Snapshot, inst OctonionicInstantonAudit, h HopfFiberActionMapAudit, o HiddenOrderParameterAudit, s IntermediateBreakingSealAudit) string {
	return fmt.Sprintf("Gate 230 audits whether the Gate-229 Hopf resonance is a dynamic theorem. The finite core contains real predata—G2/contact geometry, S_top=%.12g, B_gap=%.12g, and M_Hopf=%.9e GeV—but it does not derive a principal bundle/connection/curvature, a G2 or octonionic self-duality projector, a finite Yang-Mills action, a nontrivial instanton solution, a Hopf-fiber localization functional, or a hidden order parameter. Therefore the resonance remains conditional and the %s remains %s.", g.TopologicalActionNumerator, g.BGap, h.HopfMIntGeV, s.SealName, s.OperationalStatus)
}

func FormatGate229(s Gate229Snapshot) string {
	return fmt.Sprintf("inherited=%t geom=%t nativeHopf=%t seal=%t sensitivity=%t residualCovered=%t M_int=%.9e M_Hopf=%.9e M*=%.9e Bgap=%.12g c=%.12g cReq=%.12g S_top=%.12g VolS3=%.12g c4pi=%t", s.Gate229Inherited, s.GeometricHierarchySupported, s.NativeHopfMapDerived, s.IntermediateSealGranted, s.SensitivityBinding, s.ResidualPlausiblyCovered, s.MIntTargetGeV, s.MIntHopfGeV, s.MStarGeV, s.BGap, s.HopfCoefficient, s.RequiredCoefficient, s.TopologicalActionNumerator, s.HopfFiberVolume, s.CoefficientEqualsFourOverPi)
}

func FormatChernWeil(s ChernWeilSnapshot) string {
	return fmt.Sprintf("inherited=%t gauge=%t repTrace=%t S_top=%t bundle=%t conn=%t curv=%t trace=%t intTrFF=%t kZ=%t promoted=%t instanton=%t", s.Gate181Inherited, s.GaugeAlgebraClosed, s.RepresentationTraceRatioClosed, s.TopologicalSealAvailable, s.PrincipalBundleDerived, s.ConnectionOnFourCarrierDerived, s.CurvatureTwoFormDerived, s.TracePairingDerived, s.IntegralTrFedgeFDerived, s.IntegerInstantonNumberDerived, s.ContinuumNormalizationPromoted, s.InstantonBridgePromoted)
}

func FormatInstanton(a OctonionicInstantonAudit) string {
	return fmt.Sprintf("equation=%q carrier=%q G2=%t clifford=%t S_top=%t Bgap=%t bundle=%t conn=%t curv=%t g2Projector=%t YM=%t BPS=%t solution=%t kZ=%t derived=%t verdict=%s", a.CandidateEquation, a.GeometryCarrier, a.G2ContactPredataAvailable, a.CliffordOctonionicPredataAvailable, a.TopologicalActionSealAvailable, a.BGapAvailable, a.PrincipalBundleDerived, a.GaugeConnectionDerived, a.CurvatureTwoFormDerived, a.G2SelfDualityProjectorDerived, a.FiniteYangMillsActionDerived, a.BPSOrCriticalPointEquationDerived, a.NontrivialFiniteSolutionDerived, a.IntegerTopologicalChargeDerived, a.OctonionicInstantonDerived, a.Verdict)
}

func FormatHopfAction(a HopfFiberActionMapAudit) string {
	return fmt.Sprintf("formula=%q S_top=%.12g VolS3=%.12g c=%.12g target=%.9e Hopf=%.9e logGap=%.9f S7std=%t S3std=%t contactToS7=%t fiberFunctional=%t density=%t BgapCoupling=%t derived=%t conditional=%t verdict=%s", a.Formula, a.TopologicalActionNumerator, a.HopfFiberVolume, a.HopfCoefficient, a.TargetMIntGeV, a.HopfMIntGeV, a.Log10Gap, a.S7HopfFibrationStandardMathAvailable, a.S3FiberVolumeStandardMathAvailable, a.ContactVacuumToS7MapDerived, a.FiberLocalizationFunctionalDerived, a.ActionDensityOnFiberDerived, a.BGapAsInstantonCouplingDerived, a.HopfActionMapDerived, a.ConditionalShapeSupported, a.Verdict)
}

func FormatOrderParameter(a HiddenOrderParameterAudit) string {
	return fmt.Sprintf("candidate=%s Bgap=%.12g scalarAnchor=%t field=%t action=%t potential=%t VEV=%t shiftBreak=%t axionEFTScale=%t derived=%t verdict=%s", a.CandidateName, a.BGap, a.ScalarSpectralAnchorAvailable, a.ContinuousFieldDerived, a.LocalEffectiveActionDerived, a.PotentialDerived, a.VEVAtHopfScaleDerived, a.ShiftSymmetryBreakingDerived, a.AxionOrEFTScaleGenerated, a.HiddenOrderParameterDerived, a.Verdict)
}

func FormatSeal(a IntermediateBreakingSealAudit) string {
	return fmt.Sprintf("seal=%s prepared=%t granted=%t instanton=%t hopfMap=%t orderParam=%t geomInherited=%t status=%s verdict=%s", a.SealName, a.PreviouslyPrepared, a.Granted, a.RequiredInstantonDerived, a.RequiredHopfActionMapDerived, a.RequiredOrderParameter, a.GeometricResonanceInherited, a.OperationalStatus, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("sealedInputs=%t observed=%t patiSalam=%t leptoquark=%t instantonInvented=%t connInvented=%t fitted=%t BgapField=%t hiddenVEV=%t seal=%t polluted=%t verdict=%s", a.UsedOnlySealedInputs, a.ObservedInputsIntroduced, a.PatiSalamReopened, a.LeptoquarkDynamicsReopened, a.InstatonEquationInvented, a.DFOrConnectionInvented, a.HopfActionNormalizationFitted, a.BGapPromotedToPhysicalField, a.HiddenVEVInvented, a.IntermediateBreakingSealGranted, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("geomInherited=%t instanton=%t hopfMap=%t orderParam=%t seal=%t status=%s", s.GeometricResonanceInherited, s.FiniteInstantonDerived, s.HopfActionMapDerived, s.HiddenOrderParameterDerived, s.IntermediateSealGranted, s.Status)
}
