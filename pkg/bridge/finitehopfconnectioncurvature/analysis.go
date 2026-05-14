// Package finitehopfconnectioncurvature implements Gate 285:
// Finite Hopf Connection & Curvature / Chern-Simons Boundary Winding Audit.
//
// Gate 284 formalized the candidate instanton action
//
//	S_inst = (4/π)/B_gap,
//
// but refused to promote it because no finite Hopf connection A, curvature F,
// Chern-Simons boundary functional, or B_gap-as-coupling theorem was derived.
// Gate 285 audits exactly that missing gauge-theoretic machinery. The result is
// intentionally strict: the standard continuum Hopf/BPST connection template is
// recorded as a mathematical target, while the ASHA finite core still lacks a
// native finite connection and Chern-Simons evaluator.
package finitehopfconnectioncurvature

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactvacuumhopfaction"
)

const (
	AuditID = "GATE285-FINITE-HOPF-CONNECTION-CURVATURE-CHERN-SIMONS-BOUNDARY-WINDING-AUDIT"

	StatusGate284Inherited              = "CONDITIONAL_SUPPORT_GATE284_CONTACT_VACUUM_ACTION_REQUIREMENTS_INHERITED"
	StatusConnectionTargetsFormalized   = "CONDITIONAL_SUPPORT_HOPF_CONNECTION_TARGETS_FORMALIZED"
	StatusCurvatureRequirementsAudited  = "CONDITIONAL_SUPPORT_CURVATURE_TWO_FORM_REQUIREMENTS_AUDITED"
	StatusCSBoundaryRequirementsAudited = "CONDITIONAL_SUPPORT_CHERN_SIMONS_BOUNDARY_WINDING_REQUIREMENTS_AUDITED"
	StatusActionFunctionalReevaluated   = "CONDITIONAL_SUPPORT_INSTANTON_ACTION_FUNCTIONAL_REEVALUATED"
	StatusFirewallsPreserved            = "CONDITIONAL_SUPPORT_FINITE_HOPF_CONNECTION_FIREWALLS_PRESERVED"

	StatusFailedFiniteConnectionMissing     = "FAILED_ROUTE_FINITE_HOPF_CONNECTION_NOT_DERIVED"
	StatusFailedFiniteCurvatureMissing      = "FAILED_ROUTE_FINITE_CURVATURE_TWO_FORM_NOT_DERIVED"
	StatusFailedCSFunctionalMissing         = "FAILED_ROUTE_CHERN_SIMONS_BOUNDARY_FUNCTIONAL_NOT_DERIVED"
	StatusFailedIntegerWindingMissing       = "FAILED_ROUTE_INTEGER_BOUNDARY_WINDING_NOT_DERIVED"
	StatusFailedBGapCouplingMissing         = "FAILED_ROUTE_BGAP_AS_INSTANTON_COUPLING_NOT_DERIVED"
	StatusFailedOrderParameterMissing       = "FAILED_ROUTE_HIDDEN_SECTOR_ORDER_PARAMETER_STILL_NOT_DERIVED"
	StatusFailedActionEvaluationMissing     = "FAILED_ROUTE_FINITE_HOPF_CONNECTION_AND_ACTION_NOT_EVALUATED"
	StatusFailedIntermediateSealStillNeeded = "FAILED_ROUTE_INTERMEDIATE_BREAKING_SEAL_REMAINS_REQUIRED"
)

// Gate284Snapshot is a lightweight audited snapshot of the immediately previous
// gate. It intentionally avoids importing older dynamic chains beyond Gate 284.
type Gate284Snapshot struct {
	Gate284Inherited              bool
	InstantonFunctionalFormalized bool
	ContactVacuumMapDerived       bool
	HiddenOrderParameterDerived   bool
	ResidualCorrectionDerived     bool
	IntermediateTheoremUpgraded   bool
	IntermediateSealGranted       bool
	BGap                          float64
	Coefficient                   float64
	CoefficientExact              string
	CandidateExponent             float64
	PredictedMIntGeV              float64
	TargetMIntGeV                 float64
	Log10Gap                      float64
	TruthStatement                string
}

type ConnectionCandidate struct {
	CandidateName                  string
	ContinuumTemplate              string
	Bundle                         string
	Boundary                       string
	GaugeAlgebra                   string
	HopfFibrationAvailable         bool
	S3FiberAvailable               bool
	LocalQuaternionicAlgebraHint   bool
	PrincipalBundleDerived         bool
	FiniteConnectionOneFormDerived bool
	ConnectionCoefficientsDerived  bool
	GaugeTransformationLawDerived  bool
	GlobalPatchDataDerived         bool
	NativeFiniteConnectionDerived  bool
	Verdict                        string
}

type CurvatureAudit struct {
	Formula                    string
	RequiresFiniteExteriorD    bool
	RequiresWedgeProduct       bool
	RequiresLieBracket         bool
	RequiresTracePairing       bool
	FiniteExteriorDDerived     bool
	WedgeProductDerived        bool
	LieBracketClosureAvailable bool
	TracePairingDerived        bool
	CurvatureTwoFormDerived    bool
	YangMillsDensityDerived    bool
	Verdict                    string
}

type ChernSimonsAudit struct {
	Formula                       string
	BoundaryManifold              string
	RequiresConnection            bool
	RequiresCurvature             bool
	RequiresOrientationAndMeasure bool
	RequiresIntegerWindingMap     bool
	S3BoundaryVolumeAvailable     bool
	OrientationAndMeasureDerived  bool
	BoundaryEmbeddingDerived      bool
	ChernSimonsThreeFormDerived   bool
	IntegralEvaluatorDerived      bool
	IntegerWindingNumberDerived   bool
	BoundaryWindingEvaluated      bool
	Verdict                       string
}

type ActionFunctionalAudit struct {
	CandidateFormula             string
	Coefficient                  float64
	CoefficientExact             string
	CandidateExponent            float64
	TopologicalRatioAvailable    bool
	FiniteConnectionDerived      bool
	CurvatureDerived             bool
	ChernSimonsWindingDerived    bool
	BGapAvailable                bool
	BGapAsInverseCouplingDerived bool
	HiddenOrderParameterDerived  bool
	ActionEvaluationDerived      bool
	IntermediateScaleTheorem     bool
	Verdict                      string
}

type CouplingLedger struct {
	BGap                          float64
	CandidateCouplingStatement    string
	BGapSpectralDatumAvailable    bool
	CouplingNormalizationDerived  bool
	InverseCouplingMapDerived     bool
	GaugeKineticNormalizationOpen bool
	ContactVacuumBoundaryOpen     bool
	Verdict                       string
}

type FirewallAudit struct {
	UsesOnlyGate284Data          bool
	DoesNotInventConnection      bool
	DoesNotInventCurvature       bool
	DoesNotInventCSFunctional    bool
	DoesNotPromoteBGapToCoupling bool
	DoesNotClaimIntegerWinding   bool
	DoesNotDeclareOrderParameter bool
	DoesNotGrantIntermediateSeal bool
	DoesNotFitResidual           bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	Gate284Inherited        bool
	ConnectionTargetAudited bool
	FiniteConnectionDerived bool
	CurvatureDerived        bool
	CSWindingDerived        bool
	BGapCouplingDerived     bool
	ActionEvaluated         bool
	IntermediateTheorem     bool
	FirewallPreserved       bool
	Status                  string
	DirectAnswer            string
	NextGate                string
}

type Analysis struct {
	Gate284     Gate284Snapshot
	Connection  ConnectionCandidate
	Curvature   CurvatureAudit
	ChernSimons ChernSimonsAudit
	Action      ActionFunctionalAudit
	Coupling    CouplingLedger
	Firewall    FirewallAudit
	Summary     Summary

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
	g284, err := contactvacuumhopfaction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 284 contact-vacuum action predecessor: %w", err)
	}
	snap := snapshotGate284(g284)
	if !snap.Gate284Inherited || !snap.InstantonFunctionalFormalized || snap.BGap <= 0 || snap.Coefficient <= 0 {
		return Analysis{}, fmt.Errorf("Gate 285 requires Gate 284 formalized instanton candidate and positive B-gap data")
	}
	conn := auditConnection(snap)
	curv := auditCurvature(conn)
	cs := auditChernSimons(conn, curv)
	action := auditAction(snap, conn, curv, cs)
	coupling := auditCoupling(snap, action)
	firewall := auditFirewall(snap, conn, curv, cs, action, coupling)
	summary := buildSummary(snap, conn, curv, cs, action, coupling, firewall)
	truth := "Gate 285 audits the missing gauge-theoretic bridge beneath the 4/π B-gap resonance. It records the continuum Hopf/BPST/Chern-Simons connection template as the correct target, but the ASHA finite core still has not derived a finite connection one-form, curvature two-form, Chern-Simons boundary evaluator, integer winding map, or B_gap-as-instanton-coupling theorem. The intermediate scale therefore remains a resonance behind the IntermediateBreakingSeal."
	return Analysis{Gate284: snap, Connection: conn, Curvature: curv, ChernSimons: cs, Action: action, Coupling: coupling, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotGate284(a contactvacuumhopfaction.Analysis) Gate284Snapshot {
	return Gate284Snapshot{
		Gate284Inherited:              a.Summary.Gate283Inherited && a.Summary.InstantonFunctionalFormalized,
		InstantonFunctionalFormalized: a.Summary.InstantonFunctionalFormalized,
		ContactVacuumMapDerived:       a.Summary.ContactVacuumMapDerived,
		HiddenOrderParameterDerived:   a.Summary.HiddenOrderParameterDerived,
		ResidualCorrectionDerived:     a.Summary.ResidualCorrectionDerived,
		IntermediateTheoremUpgraded:   a.Summary.IntermediateTheoremUpgraded,
		IntermediateSealGranted:       a.Summary.IntermediateSealGranted,
		BGap:                          a.Gate283.BGap,
		Coefficient:                   a.Instanton.Coefficient,
		CoefficientExact:              a.Instanton.CoefficientExact,
		CandidateExponent:             a.Instanton.CandidateExponent,
		PredictedMIntGeV:              a.BoundaryMap.CandidateIntermediateScaleGeV,
		TargetMIntGeV:                 a.Gate283.TargetMIntGeV,
		Log10Gap:                      a.BoundaryMap.CandidateIntermediateScaleLog10Gap,
		TruthStatement:                a.TruthStatement,
	}
}

func auditConnection(g Gate284Snapshot) ConnectionCandidate {
	return ConnectionCandidate{
		CandidateName:                  "canonical Hopf/BPST-like S³-fiber connection target",
		ContinuumTemplate:              "A = Im(q†dq)/(1+|x|²) or equivalent SU(2) Hopf connection; finite version must be derived, not imported",
		Bundle:                         "S³ -> S⁷ -> S⁴",
		Boundary:                       "S³ Hopf fiber/contact boundary",
		GaugeAlgebra:                   "su(2) / Im(H)",
		HopfFibrationAvailable:         true,
		S3FiberAvailable:               true,
		LocalQuaternionicAlgebraHint:   true,
		PrincipalBundleDerived:         false,
		FiniteConnectionOneFormDerived: false,
		ConnectionCoefficientsDerived:  false,
		GaugeTransformationLawDerived:  false,
		GlobalPatchDataDerived:         false,
		NativeFiniteConnectionDerived:  false,
		Verdict:                        StatusFailedFiniteConnectionMissing,
	}
}

func auditCurvature(c ConnectionCandidate) CurvatureAudit {
	return CurvatureAudit{
		Formula:                    "F = dA + A∧A",
		RequiresFiniteExteriorD:    true,
		RequiresWedgeProduct:       true,
		RequiresLieBracket:         true,
		RequiresTracePairing:       true,
		FiniteExteriorDDerived:     false,
		WedgeProductDerived:        false,
		LieBracketClosureAvailable: c.LocalQuaternionicAlgebraHint,
		TracePairingDerived:        false,
		CurvatureTwoFormDerived:    false,
		YangMillsDensityDerived:    false,
		Verdict:                    StatusFailedFiniteCurvatureMissing,
	}
}

func auditChernSimons(c ConnectionCandidate, f CurvatureAudit) ChernSimonsAudit {
	return ChernSimonsAudit{
		Formula:                       "CS₃(A)=Tr(A∧dA + (2/3)A∧A∧A)",
		BoundaryManifold:              "S³ Hopf fiber",
		RequiresConnection:            true,
		RequiresCurvature:             true,
		RequiresOrientationAndMeasure: true,
		RequiresIntegerWindingMap:     true,
		S3BoundaryVolumeAvailable:     c.S3FiberAvailable,
		OrientationAndMeasureDerived:  false,
		BoundaryEmbeddingDerived:      false,
		ChernSimonsThreeFormDerived:   false,
		IntegralEvaluatorDerived:      false,
		IntegerWindingNumberDerived:   false,
		BoundaryWindingEvaluated:      false,
		Verdict:                       StatusFailedCSFunctionalMissing,
	}
}

func auditAction(g Gate284Snapshot, c ConnectionCandidate, f CurvatureAudit, cs ChernSimonsAudit) ActionFunctionalAudit {
	return ActionFunctionalAudit{
		CandidateFormula:             "S_inst,candidate = S_top/(π Vol(S³) B_gap) = (4/π)/B_gap",
		Coefficient:                  g.Coefficient,
		CoefficientExact:             g.CoefficientExact,
		CandidateExponent:            g.CandidateExponent,
		TopologicalRatioAvailable:    math.Abs(g.Coefficient-4/math.Pi) < 1e-12,
		FiniteConnectionDerived:      c.NativeFiniteConnectionDerived,
		CurvatureDerived:             f.CurvatureTwoFormDerived,
		ChernSimonsWindingDerived:    cs.BoundaryWindingEvaluated && cs.IntegerWindingNumberDerived,
		BGapAvailable:                g.BGap > 0,
		BGapAsInverseCouplingDerived: false,
		HiddenOrderParameterDerived:  false,
		ActionEvaluationDerived:      false,
		IntermediateScaleTheorem:     false,
		Verdict:                      StatusFailedActionEvaluationMissing,
	}
}

func auditCoupling(g Gate284Snapshot, a ActionFunctionalAudit) CouplingLedger {
	return CouplingLedger{
		BGap:                          g.BGap,
		CandidateCouplingStatement:    "candidate g_B² ∝ B_gap so S_inst ∝ 1/B_gap; currently a required map, not a theorem",
		BGapSpectralDatumAvailable:    g.BGap > 0,
		CouplingNormalizationDerived:  false,
		InverseCouplingMapDerived:     false,
		GaugeKineticNormalizationOpen: true,
		ContactVacuumBoundaryOpen:     true,
		Verdict:                       StatusFailedBGapCouplingMissing,
	}
}

func auditFirewall(g Gate284Snapshot, c ConnectionCandidate, f CurvatureAudit, cs ChernSimonsAudit, a ActionFunctionalAudit, coupling CouplingLedger) FirewallAudit {
	polluted := c.NativeFiniteConnectionDerived || f.CurvatureTwoFormDerived || cs.BoundaryWindingEvaluated || a.BGapAsInverseCouplingDerived || a.IntermediateScaleTheorem || coupling.InverseCouplingMapDerived
	verdict := StatusFirewallsPreserved
	if polluted {
		verdict = "FAILED_ROUTE_FIREWALL_POLLUTED_BY_UNDERIVED_CONNECTION_OR_COUPLING"
	}
	return FirewallAudit{
		UsesOnlyGate284Data:          g.Gate284Inherited,
		DoesNotInventConnection:      !c.NativeFiniteConnectionDerived,
		DoesNotInventCurvature:       !f.CurvatureTwoFormDerived,
		DoesNotInventCSFunctional:    !cs.ChernSimonsThreeFormDerived && !cs.BoundaryWindingEvaluated,
		DoesNotPromoteBGapToCoupling: !a.BGapAsInverseCouplingDerived && !coupling.InverseCouplingMapDerived,
		DoesNotClaimIntegerWinding:   !cs.IntegerWindingNumberDerived,
		DoesNotDeclareOrderParameter: !a.HiddenOrderParameterDerived,
		DoesNotGrantIntermediateSeal: !a.IntermediateScaleTheorem,
		DoesNotFitResidual:           true,
		FiniteCorePolluted:           polluted,
		Verdict:                      verdict,
	}
}

func buildSummary(g Gate284Snapshot, c ConnectionCandidate, f CurvatureAudit, cs ChernSimonsAudit, a ActionFunctionalAudit, coupling CouplingLedger, fw FirewallAudit) Summary {
	return Summary{
		Gate284Inherited:        g.Gate284Inherited,
		ConnectionTargetAudited: c.HopfFibrationAvailable && c.S3FiberAvailable && c.LocalQuaternionicAlgebraHint,
		FiniteConnectionDerived: c.NativeFiniteConnectionDerived,
		CurvatureDerived:        f.CurvatureTwoFormDerived,
		CSWindingDerived:        cs.BoundaryWindingEvaluated,
		BGapCouplingDerived:     coupling.InverseCouplingMapDerived,
		ActionEvaluated:         a.ActionEvaluationDerived,
		IntermediateTheorem:     a.IntermediateScaleTheorem,
		FirewallPreserved:       !fw.FiniteCorePolluted && fw.DoesNotInventConnection && fw.DoesNotPromoteBGapToCoupling,
		Status:                  StatusFailedActionEvaluationMissing,
		DirectAnswer:            "No. Gate 285 identifies the correct finite Hopf connection/Chern-Simons machinery needed for Path C, but it does not derive the connection, curvature, boundary winding, or B_gap coupling map.",
		NextGate:                "derive a finite Hopf connection or formally cap Path C behind a connection/action/coupling firewall",
	}
}

func StatusList(a Analysis) []string {
	statuses := []string{
		StatusGate284Inherited,
		StatusConnectionTargetsFormalized,
		StatusCurvatureRequirementsAudited,
		StatusCSBoundaryRequirementsAudited,
		StatusActionFunctionalReevaluated,
		a.Firewall.Verdict,
	}
	if !a.Connection.NativeFiniteConnectionDerived {
		statuses = append(statuses, StatusFailedFiniteConnectionMissing)
	}
	if !a.Curvature.CurvatureTwoFormDerived {
		statuses = append(statuses, StatusFailedFiniteCurvatureMissing)
	}
	if !a.ChernSimons.ChernSimonsThreeFormDerived {
		statuses = append(statuses, StatusFailedCSFunctionalMissing)
	}
	if !a.ChernSimons.IntegerWindingNumberDerived {
		statuses = append(statuses, StatusFailedIntegerWindingMissing)
	}
	if !a.Action.BGapAsInverseCouplingDerived {
		statuses = append(statuses, StatusFailedBGapCouplingMissing)
	}
	if !a.Action.HiddenOrderParameterDerived {
		statuses = append(statuses, StatusFailedOrderParameterMissing)
	}
	if !a.Action.ActionEvaluationDerived {
		statuses = append(statuses, StatusFailedActionEvaluationMissing)
	}
	statuses = append(statuses, StatusFailedIntermediateSealStillNeeded)
	return statuses
}

func FormatGate284(g Gate284Snapshot) string {
	return fmt.Sprintf("Gate284Inherited=%t instantonFormal=%t contactMap=%t orderParameter=%t residual=%t theorem=%t seal=%t B_gap=%.12g coeff=%s=%.12g exponent=%.12g pred=%.12e target=%.12e log10gap=%.12g", g.Gate284Inherited, g.InstantonFunctionalFormalized, g.ContactVacuumMapDerived, g.HiddenOrderParameterDerived, g.ResidualCorrectionDerived, g.IntermediateTheoremUpgraded, g.IntermediateSealGranted, g.BGap, g.CoefficientExact, g.Coefficient, g.CandidateExponent, g.PredictedMIntGeV, g.TargetMIntGeV, g.Log10Gap)
}

func FormatConnection(c ConnectionCandidate) string {
	return fmt.Sprintf("%s bundle=%s boundary=%s algebra=%s template=%q Hopf=%t S3=%t Hhint=%t principalBundle=%t A=%t coeffs=%t gaugeLaw=%t patches=%t native=%t verdict=%s", c.CandidateName, c.Bundle, c.Boundary, c.GaugeAlgebra, c.ContinuumTemplate, c.HopfFibrationAvailable, c.S3FiberAvailable, c.LocalQuaternionicAlgebraHint, c.PrincipalBundleDerived, c.FiniteConnectionOneFormDerived, c.ConnectionCoefficientsDerived, c.GaugeTransformationLawDerived, c.GlobalPatchDataDerived, c.NativeFiniteConnectionDerived, c.Verdict)
}

func FormatCurvature(c CurvatureAudit) string {
	return fmt.Sprintf("formula=%q requiresD=%t wedge=%t bracket=%t trace=%t finiteD=%t wedgeDerived=%t bracketAvail=%t traceDerived=%t F=%t YM=%t verdict=%s", c.Formula, c.RequiresFiniteExteriorD, c.RequiresWedgeProduct, c.RequiresLieBracket, c.RequiresTracePairing, c.FiniteExteriorDDerived, c.WedgeProductDerived, c.LieBracketClosureAvailable, c.TracePairingDerived, c.CurvatureTwoFormDerived, c.YangMillsDensityDerived, c.Verdict)
}

func FormatChernSimons(c ChernSimonsAudit) string {
	return fmt.Sprintf("formula=%q boundary=%s reqA=%t reqF=%t reqMeasure=%t reqWinding=%t S3vol=%t orient=%t embed=%t CS3=%t integral=%t integer=%t evaluated=%t verdict=%s", c.Formula, c.BoundaryManifold, c.RequiresConnection, c.RequiresCurvature, c.RequiresOrientationAndMeasure, c.RequiresIntegerWindingMap, c.S3BoundaryVolumeAvailable, c.OrientationAndMeasureDerived, c.BoundaryEmbeddingDerived, c.ChernSimonsThreeFormDerived, c.IntegralEvaluatorDerived, c.IntegerWindingNumberDerived, c.BoundaryWindingEvaluated, c.Verdict)
}

func FormatAction(a ActionFunctionalAudit) string {
	return fmt.Sprintf("formula=%q coeff=%s=%.12g exponent=%.12g topRatio=%t A=%t F=%t CS=%t Bgap=%t BgapCoupling=%t orderParameter=%t action=%t theorem=%t verdict=%s", a.CandidateFormula, a.CoefficientExact, a.Coefficient, a.CandidateExponent, a.TopologicalRatioAvailable, a.FiniteConnectionDerived, a.CurvatureDerived, a.ChernSimonsWindingDerived, a.BGapAvailable, a.BGapAsInverseCouplingDerived, a.HiddenOrderParameterDerived, a.ActionEvaluationDerived, a.IntermediateScaleTheorem, a.Verdict)
}

func FormatCoupling(c CouplingLedger) string {
	return fmt.Sprintf("B_gap=%.12g statement=%q datum=%t norm=%t inverseMap=%t kineticOpen=%t boundaryOpen=%t verdict=%s", c.BGap, c.CandidateCouplingStatement, c.BGapSpectralDatumAvailable, c.CouplingNormalizationDerived, c.InverseCouplingMapDerived, c.GaugeKineticNormalizationOpen, c.ContactVacuumBoundaryOpen, c.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("usesGate284=%t noA=%t noF=%t noCS=%t noBGapCoupling=%t noWinding=%t noOrderParam=%t noSeal=%t noResidualFit=%t polluted=%t verdict=%s", f.UsesOnlyGate284Data, f.DoesNotInventConnection, f.DoesNotInventCurvature, f.DoesNotInventCSFunctional, f.DoesNotPromoteBGapToCoupling, f.DoesNotClaimIntegerWinding, f.DoesNotDeclareOrderParameter, f.DoesNotGrantIntermediateSeal, f.DoesNotFitResidual, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("Gate284=%t targetAudited=%t A=%t F=%t CS=%t BgapCoupling=%t action=%t theorem=%t firewall=%t status=%s direct=%q next=%q", s.Gate284Inherited, s.ConnectionTargetAudited, s.FiniteConnectionDerived, s.CurvatureDerived, s.CSWindingDerived, s.BGapCouplingDerived, s.ActionEvaluated, s.IntermediateTheorem, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}

func FormatStatusList(statuses []string) string { return strings.Join(statuses, "\n") }
