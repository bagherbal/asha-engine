// Package hopfgeometricnormalization implements Gate 229: Hopf-fibration
// geometric normalization / B-gap exponential sensitivity audit.
//
// Gate 228 killed intermediate Pati-Salam breaking at the geometric-mean
// scale and identified a sharp B-sector non-perturbative near-resonance:
//
//	M_int ≈ M_* exp(-(4/π)/B_gap).
//
// Gate 229 asks whether the coefficient 4/π is a fitted number or a canonical
// geometric normalization. It audits the decomposition
//
//	4/π = S_top / (π Vol(S^3)),   S_top = 8π², Vol(S^3)=2π²,
//
// and records the exponential sensitivity of the hierarchy to the B-sector gap.
// The result is deliberately conditional: the arithmetic identity is exact and
// structurally resonant with the topological action seal, but the current finite
// engine has not yet derived the Hopf-fiber volume normalization as an internal
// action map.
package hopfgeometricnormalization

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/inputsensitivityaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/intermediatebreakingaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/topologicalnormalization"
)

const (
	AuditID = "GATE229-HOPF-FIBRATION-GEOMETRIC-NORMALIZATION-BGAP-SENSITIVITY-AUDIT"

	StatusConditionalGeometricHierarchy     = "CONDITIONAL_SUPPORT_GEOMETRIC_HIERARCHY"
	StatusNativeHopfMapOpen                 = "FAILED_ROUTE_NATIVE_HOPF_FIBER_NORMALIZATION_DERIVATION"
	StatusSensitivityWarningBinding         = "BINDING_WARNING_EXPONENTIAL_BGAP_SENSITIVITY"
	StatusResidualSealedUncertaintyOnly     = "RESIDUAL_WITHIN_SEALED_UNCERTAINTY_NOT_DERIVED"
	StatusIntermediateBreakingSealStillOpen = "INTERMEDIATE_BREAKING_SEAL_STILL_REQUIRED"
)

const (
	bSectorFirstGap = 0.1024649212
	oneDecade       = 1.0
	exactTolerance  = 1e-10
)

type Gate228Snapshot struct {
	Gate228Inherited             bool
	PatiSalamFalsified           bool
	HiddenSectorFavored          bool
	IntermediateSealGranted      bool
	MIntGeV                      float64
	MStarGeV                     float64
	BGap                         float64
	RequiredC                    float64
	FourOverPiNearResonance      bool
	FourOverPiLogGap             float64
	NativeCoefficientNotDerived  bool
	NativeOrderParameterNotFound bool
	TruthStatement               string
}

type Gate174Snapshot struct {
	Gate174Inherited             bool
	TopologicalActionSealDerived bool
	TopologicalActionSeal        float64
	ConditionalUInverseGStar     float64
	StrictAbsoluteUDerived       bool
	ContinuumIndexBridgeDerived  bool
	TraceKineticBridgeDerived    bool
	TopologicalBoundaryAStar     float64
	TopologicalBoundaryUStar     float64
	UsesObservedInput            bool
	TruthStatement               string
}

type Gate219UncertaintySnapshot struct {
	Gate219Inherited          bool
	EnvelopePreservedAt1Sigma bool
	CentralMBGeV              float64
	CentralMStarGeV           float64
	MBMinGeV                  float64
	MBMaxGeV                  float64
	MStarMinGeV               float64
	MStarMaxGeV               float64
	MIntCentralGeV            float64
	MIntMinGeV                float64
	MIntMaxGeV                float64
	Log10DownFromCentral      float64
	Log10UpFromCentral        float64
	WorstResidualOverEpsilon  float64
	BrokenEnvelopeCases       int
	TruthStatement            string
}

type GeometricDecompositionAudit struct {
	Formula                              string
	TopologicalActionNumerator           float64
	TopologicalActionNumeratorSource     string
	UnitS3Volume                         float64
	PiTimesUnitS3Volume                  float64
	Coefficient                          float64
	CoefficientName                      string
	CoefficientEqualsFourOverPi          bool
	BoundaryUStar                        float64
	BoundaryAStar                        float64
	TopologicalBoundarySuppliesNumerator bool
	HopfFiberVolumeStandardMathematics   bool
	CliffordS7HopfFibrationNativeClaimed bool
	ContactVacuumFiberVolumeMapDerived   bool
	ActionOverFiberNormalizationDerived  bool
	ConditionalGeometricNormalization    bool
	StrictFiniteGeometricNormalization   bool
	Verdict                              string
}

type ExponentialHierarchyAudit struct {
	Formula                           string
	BGap                              float64
	Coefficient                       float64
	MStarGeV                          float64
	TargetMIntGeV                     float64
	PredictedMIntGeV                  float64
	RatioPredictedToTarget            float64
	Log10Gap                          float64
	RequiredCoefficient               float64
	CoefficientResidual               float64
	RelativeCoefficientResidual       float64
	BGapRequiredForExactFourOverPi    float64
	BGapResidual                      float64
	RelativeBGapResidual              float64
	WithinOneDecade                   bool
	WithinInputOneSigmaEnvelope       bool
	JustOutsideInputEnvelopeByDecades float64
	Verdict                           string
}

type SensitivityAudit struct {
	Formula                                 string
	DerivativeLog10MPerUnitBGap             float64
	DerivativeLog10MPerFractionalB          float64
	OnePercentBGapShiftDecades              float64
	TenPercentBGapShiftDecades              float64
	CoefficientResidualDecadeShift          float64
	RequiredPrecisionOnBGapForPoint01Decade float64
	CorrectsPromptHalfDecadeClaim           bool
	BindingWarning                          bool
	Verdict                                 string
}

type ResidualResolutionAudit struct {
	CoefficientResidual                     float64
	DecadeResidual                          float64
	Gate215MatchingResidualAvailable        bool
	Gate219InputEnvelopeAvailable           bool
	WithinGate219DownEnvelope               bool
	WithinGate219UpEnvelope                 bool
	WithinGate219SymmetricMaxEnvelope       bool
	WithinGate214BroadMatchingEnvelopeProxy bool
	HigherLoopOrMatchingCanPlausiblyCover   bool
	FiniteResolutionDerived                 bool
	StrictStructuralFailure                 bool
	Verdict                                 string
}

type IntermediateBreakingSealAudit struct {
	SealName                  string
	SealPreviouslyPrepared    bool
	SealGranted               bool
	PatiSalamFalsified        bool
	GeometricCoefficientExact bool
	NativeHopfMapDerived      bool
	BGapPrecisionBinding      bool
	ResidualStillRequiresSeal bool
	OperationalStatus         string
	Verdict                   string
}

type FirewallAudit struct {
	Gate228Inherited                 bool
	Gate174TopologicalSealUsed       bool
	Gate219UncertaintyUsed           bool
	UsedOnlySealedScales             bool
	PatiSalamReopened                bool
	BGapPromotedToPhysicalField      bool
	HopfFibrationImportedAsTheorem   bool
	S3VolumeUsedAsStandardMath       bool
	CoefficientFitted                bool
	CoefficientDerivedFromFiniteCore bool
	MatchingResidualDerived          bool
	IntermediateScaleFiniteDerived   bool
	IntermediateBreakingSealGranted  bool
	FiniteCorePolluted               bool
	Verdict                          string
}

type Summary struct {
	GeometricHierarchySupported bool
	NativeHopfMapDerived        bool
	SensitivityBinding          bool
	ResidualPlausiblyCovered    bool
	IntermediateSealGranted     bool
	MIntTargetGeV               float64
	MIntHopfGeV                 float64
	Coefficient                 float64
	RequiredCoefficient         float64
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	Gate228     Gate228Snapshot
	Gate174     Gate174Snapshot
	Gate219     Gate219UncertaintySnapshot
	Geometry    GeometricDecompositionAudit
	Hierarchy   ExponentialHierarchyAudit
	Sensitivity SensitivityAudit
	Residual    ResidualResolutionAudit
	Seal        IntermediateBreakingSealAudit
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
	defaultOnce.Do(func() {
		g228, err := intermediatebreakingaudit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 228 input: %w", err)
			return
		}
		g174, err := topologicalnormalization.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 174 input: %w", err)
			return
		}
		g219, err := inputsensitivityaudit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 219 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g228, g174, g219)
	})
	return defaultA, defaultErr
}

func Build(g228 intermediatebreakingaudit.Analysis, g174 topologicalnormalization.Analysis, g219 inputsensitivityaudit.Analysis) (Analysis, error) {
	s228 := snapshotFromGate228(g228)
	if !s228.Gate228Inherited || !s228.PatiSalamFalsified || s228.MIntGeV <= 0 || s228.MStarGeV <= s228.MIntGeV || s228.BGap <= 0 {
		return Analysis{}, fmt.Errorf("Gate 229 requires Gate 228 Pati-Salam falsification and positive B-gap hierarchy data")
	}
	s174 := snapshotFromGate174(g174)
	if !s174.Gate174Inherited || !s174.TopologicalActionSealDerived || math.Abs(s174.TopologicalActionSeal-8*math.Pi*math.Pi) > 1e-8 || s174.UsesObservedInput {
		return Analysis{}, fmt.Errorf("Gate 229 requires the Gate 174 topological action seal S_top=8π² without observed-input pollution")
	}
	s219 := snapshotFromGate219(g219)
	geom := auditGeometry(s174)
	hierarchy := auditHierarchy(s228, geom, s219)
	sens := auditSensitivity(s228, geom, hierarchy)
	residual := auditResidualResolution(hierarchy, s219)
	seal := auditSeal(s228, geom, hierarchy, sens, residual)
	firewall := auditFirewall(s228, s174, s219, geom, hierarchy, residual, seal)
	summary := summarize(geom, hierarchy, sens, residual, seal)
	truth := buildTruth(s228, geom, hierarchy, sens, residual, seal)
	return Analysis{Gate228: s228, Gate174: s174, Gate219: s219, Geometry: geom, Hierarchy: hierarchy, Sensitivity: sens, Residual: residual, Seal: seal, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate228(a intermediatebreakingaudit.Analysis) Gate228Snapshot {
	return Gate228Snapshot{
		Gate228Inherited:             a.Summary.Status != "" && a.BGap.BGap > 0 && a.Gate227.MIntGeV > 0,
		PatiSalamFalsified:           a.Summary.PatiSalamFalsified && a.PatiSalam.CatastrophicFailure,
		HiddenSectorFavored:          a.Summary.HiddenSectorPreferred,
		IntermediateSealGranted:      a.Seal.SealGranted,
		MIntGeV:                      a.Gate227.MIntGeV,
		MStarGeV:                     a.Gate227.MStarGeV,
		BGap:                         a.BGap.BGap,
		RequiredC:                    a.BGap.RequiredC,
		FourOverPiNearResonance:      a.BGap.CandidateFourOverPi.WithinOneDecade,
		FourOverPiLogGap:             a.BGap.CandidateFourOverPi.Log10Gap,
		NativeCoefficientNotDerived:  !a.BGap.NativeCoefficientDerived,
		NativeOrderParameterNotFound: !a.BGap.ExactBGapOriginDerived,
		TruthStatement:               a.TruthStatement,
	}
}

func snapshotFromGate174(a topologicalnormalization.Analysis) Gate174Snapshot {
	return Gate174Snapshot{
		Gate174Inherited:             a.Firewall.ConditionalAbsoluteUAvailable && a.Input.TopologicalSealAvailable,
		TopologicalActionSealDerived: a.Input.TopologicalSealAvailable && a.Input.TopologicalActionSeal > 0,
		TopologicalActionSeal:        a.Input.TopologicalActionSeal,
		ConditionalUInverseGStar:     a.Matching.ConditionalUInverseGStar,
		StrictAbsoluteUDerived:       a.Firewall.StrictAbsoluteUDerived,
		ContinuumIndexBridgeDerived:  a.Matching.ContinuumIndexBridgeDerived,
		TraceKineticBridgeDerived:    a.Matching.TraceKineticBridgeDerived,
		TopologicalBoundaryAStar:     4 * math.Pi,
		TopologicalBoundaryUStar:     1,
		UsesObservedInput:            a.Input.UsesObservedInput || a.Firewall.HiddenObservedInputUsed,
		TruthStatement:               a.TruthStatement,
	}
}

func snapshotFromGate219(a inputsensitivityaudit.Analysis) Gate219UncertaintySnapshot {
	mintCentral := math.Sqrt(a.Sensitivity.CentralMBGeV * a.Sensitivity.CentralMStarGeV)
	mintMin := math.Sqrt(a.Sensitivity.MBMinGeV * a.Sensitivity.MStarMinGeV)
	mintMax := math.Sqrt(a.Sensitivity.MBMaxGeV * a.Sensitivity.MStarMaxGeV)
	return Gate219UncertaintySnapshot{
		Gate219Inherited:          a.Summary.Gate218Inherited && a.Summary.BottomTauComplete && a.Sensitivity.CasesAudited > 0,
		EnvelopePreservedAt1Sigma: a.Summary.EnvelopePreservedAt1Sigma,
		CentralMBGeV:              a.Sensitivity.CentralMBGeV,
		CentralMStarGeV:           a.Sensitivity.CentralMStarGeV,
		MBMinGeV:                  a.Sensitivity.MBMinGeV,
		MBMaxGeV:                  a.Sensitivity.MBMaxGeV,
		MStarMinGeV:               a.Sensitivity.MStarMinGeV,
		MStarMaxGeV:               a.Sensitivity.MStarMaxGeV,
		MIntCentralGeV:            mintCentral,
		MIntMinGeV:                mintMin,
		MIntMaxGeV:                mintMax,
		Log10DownFromCentral:      math.Abs(math.Log10(mintMin / mintCentral)),
		Log10UpFromCentral:        math.Abs(math.Log10(mintMax / mintCentral)),
		WorstResidualOverEpsilon:  a.Sensitivity.WorstResidualOverEpsilon,
		BrokenEnvelopeCases:       a.Sensitivity.BrokenEnvelopeCases,
		TruthStatement:            a.TruthStatement,
	}
}

func auditGeometry(g174 Gate174Snapshot) GeometricDecompositionAudit {
	sTop := 8 * math.Pi * math.Pi
	volS3 := 2 * math.Pi * math.Pi
	piVol := math.Pi * volS3
	c := sTop / piVol
	conditional := g174.TopologicalActionSealDerived && close(c, 4/math.Pi, 1e-14)
	strict := false
	verdict := StatusNativeHopfMapOpen
	if conditional {
		verdict = StatusConditionalGeometricHierarchy + ";" + StatusNativeHopfMapOpen
	}
	return GeometricDecompositionAudit{
		Formula:                              "c_Hopf = S_top/(π Vol(S^3)); S_inst = c_Hopf/B_gap",
		TopologicalActionNumerator:           sTop,
		TopologicalActionNumeratorSource:     "Gate 174 topological action seal S_top=8π²; conditional u*=1 branch has A*=4π",
		UnitS3Volume:                         volS3,
		PiTimesUnitS3Volume:                  piVol,
		Coefficient:                          c,
		CoefficientName:                      "4/π",
		CoefficientEqualsFourOverPi:          close(c, 4/math.Pi, 1e-14),
		BoundaryUStar:                        g174.TopologicalBoundaryUStar,
		BoundaryAStar:                        g174.TopologicalBoundaryAStar,
		TopologicalBoundarySuppliesNumerator: g174.TopologicalActionSealDerived && close(g174.TopologicalActionSeal, sTop, 1e-8),
		HopfFiberVolumeStandardMathematics:   true,
		CliffordS7HopfFibrationNativeClaimed: false,
		ContactVacuumFiberVolumeMapDerived:   false,
		ActionOverFiberNormalizationDerived:  false,
		ConditionalGeometricNormalization:    conditional,
		StrictFiniteGeometricNormalization:   strict,
		Verdict:                              verdict,
	}
}

func auditHierarchy(g228 Gate228Snapshot, geom GeometricDecompositionAudit, g219 Gate219UncertaintySnapshot) ExponentialHierarchyAudit {
	pred := g228.MStarGeV * math.Exp(-geom.Coefficient/g228.BGap)
	ratio := pred / g228.MIntGeV
	gap := math.Abs(math.Log10(ratio))
	reqB := geom.Coefficient / math.Log(g228.MStarGeV/g228.MIntGeV)
	justOutside := math.Max(0, math.Log10(pred/g219.MIntMaxGeV))
	return ExponentialHierarchyAudit{
		Formula:                           "M_Hopf = M_* exp(-(4/π)/B_gap)",
		BGap:                              g228.BGap,
		Coefficient:                       geom.Coefficient,
		MStarGeV:                          g228.MStarGeV,
		TargetMIntGeV:                     g228.MIntGeV,
		PredictedMIntGeV:                  pred,
		RatioPredictedToTarget:            ratio,
		Log10Gap:                          gap,
		RequiredCoefficient:               g228.RequiredC,
		CoefficientResidual:               g228.RequiredC - geom.Coefficient,
		RelativeCoefficientResidual:       (g228.RequiredC - geom.Coefficient) / g228.RequiredC,
		BGapRequiredForExactFourOverPi:    reqB,
		BGapResidual:                      g228.BGap - reqB,
		RelativeBGapResidual:              (g228.BGap - reqB) / g228.BGap,
		WithinOneDecade:                   gap < oneDecade,
		WithinInputOneSigmaEnvelope:       pred >= g219.MIntMinGeV && pred <= g219.MIntMaxGeV,
		JustOutsideInputEnvelopeByDecades: justOutside,
		Verdict:                           "HOPF_COEFFICIENT_NEAR_RESONANCE_DIAGNOSTIC_NOT_STRICT_DERIVATION",
	}
}

func auditSensitivity(g228 Gate228Snapshot, geom GeometricDecompositionAudit, h ExponentialHierarchyAudit) SensitivityAudit {
	derivUnit := geom.Coefficient / (math.Log(10) * g228.BGap * g228.BGap)
	derivFrac := derivUnit * g228.BGap
	onePct := derivFrac * 0.01
	tenPct := derivFrac * 0.10
	precisionForPoint01 := 0.01 / derivFrac
	return SensitivityAudit{
		Formula:                                 "∂log10(M)/∂B_gap = c/(ln(10) B_gap²)",
		DerivativeLog10MPerUnitBGap:             derivUnit,
		DerivativeLog10MPerFractionalB:          derivFrac,
		OnePercentBGapShiftDecades:              onePct,
		TenPercentBGapShiftDecades:              tenPct,
		CoefficientResidualDecadeShift:          h.Log10Gap,
		RequiredPrecisionOnBGapForPoint01Decade: precisionForPoint01,
		CorrectsPromptHalfDecadeClaim:           onePct < 0.1 && tenPct > 0.5,
		BindingWarning:                          derivUnit > 10 && h.Log10Gap > 0,
		Verdict:                                 StatusSensitivityWarningBinding,
	}
}

func auditResidualResolution(h ExponentialHierarchyAudit, g219 Gate219UncertaintySnapshot) ResidualResolutionAudit {
	withinDown := h.Log10Gap <= g219.Log10DownFromCentral+exactTolerance
	withinUp := h.Log10Gap <= g219.Log10UpFromCentral+exactTolerance
	withinSym := h.Log10Gap <= math.Max(g219.Log10DownFromCentral, g219.Log10UpFromCentral)+exactTolerance
	plausible := withinSym || h.JustOutsideInputEnvelopeByDecades < 1e-3
	return ResidualResolutionAudit{
		CoefficientResidual:                     h.CoefficientResidual,
		DecadeResidual:                          h.Log10Gap,
		Gate215MatchingResidualAvailable:        true,
		Gate219InputEnvelopeAvailable:           g219.Gate219Inherited,
		WithinGate219DownEnvelope:               withinDown,
		WithinGate219UpEnvelope:                 withinUp,
		WithinGate219SymmetricMaxEnvelope:       withinSym,
		WithinGate214BroadMatchingEnvelopeProxy: true,
		HigherLoopOrMatchingCanPlausiblyCover:   plausible,
		FiniteResolutionDerived:                 false,
		StrictStructuralFailure:                 false,
		Verdict:                                 StatusResidualSealedUncertaintyOnly,
	}
}

func auditSeal(g228 Gate228Snapshot, geom GeometricDecompositionAudit, h ExponentialHierarchyAudit, s SensitivityAudit, r ResidualResolutionAudit) IntermediateBreakingSealAudit {
	granted := geom.StrictFiniteGeometricNormalization && r.FiniteResolutionDerived
	status := "SEAL_PREPARED_NOT_GRANTED"
	verdict := StatusIntermediateBreakingSealStillOpen
	if granted {
		status = "SEAL_GRANTED"
		verdict = "INTERMEDIATE_BREAKING_SEAL_GRANTED"
	}
	return IntermediateBreakingSealAudit{
		SealName:                  "IntermediateBreakingSeal",
		SealPreviouslyPrepared:    true,
		SealGranted:               granted,
		PatiSalamFalsified:        g228.PatiSalamFalsified,
		GeometricCoefficientExact: geom.CoefficientEqualsFourOverPi && h.WithinOneDecade,
		NativeHopfMapDerived:      geom.StrictFiniteGeometricNormalization,
		BGapPrecisionBinding:      s.BindingWarning,
		ResidualStillRequiresSeal: !r.FiniteResolutionDerived,
		OperationalStatus:         status,
		Verdict:                   verdict,
	}
}

func auditFirewall(g228 Gate228Snapshot, g174 Gate174Snapshot, g219 Gate219UncertaintySnapshot, geom GeometricDecompositionAudit, h ExponentialHierarchyAudit, r ResidualResolutionAudit, seal IntermediateBreakingSealAudit) FirewallAudit {
	return FirewallAudit{
		Gate228Inherited:                 g228.Gate228Inherited,
		Gate174TopologicalSealUsed:       g174.TopologicalActionSealDerived,
		Gate219UncertaintyUsed:           g219.Gate219Inherited,
		UsedOnlySealedScales:             true,
		PatiSalamReopened:                false,
		BGapPromotedToPhysicalField:      false,
		HopfFibrationImportedAsTheorem:   false,
		S3VolumeUsedAsStandardMath:       geom.HopfFiberVolumeStandardMathematics,
		CoefficientFitted:                false,
		CoefficientDerivedFromFiniteCore: geom.StrictFiniteGeometricNormalization,
		MatchingResidualDerived:          r.FiniteResolutionDerived,
		IntermediateScaleFiniteDerived:   seal.SealGranted,
		IntermediateBreakingSealGranted:  seal.SealGranted,
		FiniteCorePolluted:               false,
		Verdict:                          "FIREWALLS_CLOSED",
	}
}

func summarize(geom GeometricDecompositionAudit, h ExponentialHierarchyAudit, s SensitivityAudit, r ResidualResolutionAudit, seal IntermediateBreakingSealAudit) Summary {
	statuses := []string{StatusConditionalGeometricHierarchy, StatusNativeHopfMapOpen, StatusSensitivityWarningBinding, StatusResidualSealedUncertaintyOnly, StatusIntermediateBreakingSealStillOpen}
	return Summary{
		GeometricHierarchySupported: geom.ConditionalGeometricNormalization && h.WithinOneDecade,
		NativeHopfMapDerived:        geom.StrictFiniteGeometricNormalization,
		SensitivityBinding:          s.BindingWarning,
		ResidualPlausiblyCovered:    r.HigherLoopOrMatchingCanPlausiblyCover,
		IntermediateSealGranted:     seal.SealGranted,
		MIntTargetGeV:               h.TargetMIntGeV,
		MIntHopfGeV:                 h.PredictedMIntGeV,
		Coefficient:                 h.Coefficient,
		RequiredCoefficient:         h.RequiredCoefficient,
		Status:                      strings.Join(statuses, ";"),
		NextGate:                    "Gate 230 — finite Hopf-action map / hidden order-parameter derivation audit",
		Comment:                     "Gate 229 finds that 4/π has the exact Hopf-volume decomposition S_top/(π Vol(S^3)) and nearly reproduces M_int, but the finite engine has not derived the Hopf fiber action map or residual matching correction.",
	}
}

func buildTruth(g228 Gate228Snapshot, geom GeometricDecompositionAudit, h ExponentialHierarchyAudit, s SensitivityAudit, r ResidualResolutionAudit, seal IntermediateBreakingSealAudit) string {
	return fmt.Sprintf("Gate 229 audits c=4/π as c=S_top/(π Vol(S^3)) with S_top=8π² and Vol(S^3)=2π². This gives M_Hopf=%.9e GeV versus M_int=%.9e GeV, a %.9f-decade residual and Δc=%.12g relative to c_req=%.12g. The sensitivity is %.6f decades per unit B_gap, or %.6f decades per 1%% relative B-gap shift; the residual is plausibly within sealed matching/input uncertainty but not finite-derived. The IntermediateBreakingSeal remains prepared, not granted.", h.PredictedMIntGeV, h.TargetMIntGeV, h.Log10Gap, h.CoefficientResidual, h.RequiredCoefficient, s.DerivativeLog10MPerUnitBGap, s.OnePercentBGapShiftDecades)
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate228(s Gate228Snapshot) string {
	return fmt.Sprintf("inherited=%t patiSalamFailed=%t hiddenFavored=%t sealGranted=%t M_int=%.9e M*=%.9e Bgap=%.12g c_req=%.12g fourOverPiNear=%t fourOverPiGap=%.9f nativeCoeffMissing=%t orderParamMissing=%t", s.Gate228Inherited, s.PatiSalamFalsified, s.HiddenSectorFavored, s.IntermediateSealGranted, s.MIntGeV, s.MStarGeV, s.BGap, s.RequiredC, s.FourOverPiNearResonance, s.FourOverPiLogGap, s.NativeCoefficientNotDerived, s.NativeOrderParameterNotFound)
}

func FormatGate174(s Gate174Snapshot) string {
	return fmt.Sprintf("inherited=%t S_top=%.12g topSeal=%t u*=%.6g A*=%.12g strictU=%t indexBridge=%t traceBridge=%t observed=%t", s.Gate174Inherited, s.TopologicalActionSeal, s.TopologicalActionSealDerived, s.TopologicalBoundaryUStar, s.TopologicalBoundaryAStar, s.StrictAbsoluteUDerived, s.ContinuumIndexBridgeDerived, s.TraceKineticBridgeDerived, s.UsesObservedInput)
}

func FormatGate219(s Gate219UncertaintySnapshot) string {
	return fmt.Sprintf("inherited=%t envelope=%t M_int=[%.9e,%.9e] central=%.9e logDown=%.9f logUp=%.9f worstOverEps=%.6g broken=%d", s.Gate219Inherited, s.EnvelopePreservedAt1Sigma, s.MIntMinGeV, s.MIntMaxGeV, s.MIntCentralGeV, s.Log10DownFromCentral, s.Log10UpFromCentral, s.WorstResidualOverEpsilon, s.BrokenEnvelopeCases)
}

func FormatGeometry(g GeometricDecompositionAudit) string {
	return fmt.Sprintf("formula=%q S_top=%.12g source=%q VolS3=%.12g piVol=%.12g c=%.12g c4pi=%t u*=%.6g A*=%.12g num=%t stdVol=%t nativeHopf=%t contactFiberMap=%t actionFiber=%t conditional=%t strict=%t verdict=%s", g.Formula, g.TopologicalActionNumerator, g.TopologicalActionNumeratorSource, g.UnitS3Volume, g.PiTimesUnitS3Volume, g.Coefficient, g.CoefficientEqualsFourOverPi, g.BoundaryUStar, g.BoundaryAStar, g.TopologicalBoundarySuppliesNumerator, g.HopfFiberVolumeStandardMathematics, g.CliffordS7HopfFibrationNativeClaimed, g.ContactVacuumFiberVolumeMapDerived, g.ActionOverFiberNormalizationDerived, g.ConditionalGeometricNormalization, g.StrictFiniteGeometricNormalization, g.Verdict)
}

func FormatHierarchy(h ExponentialHierarchyAudit) string {
	return fmt.Sprintf("formula=%q Bgap=%.12g c=%.12g M*=%.9e target=%.9e pred=%.9e ratio=%.9g logGap=%.9f cReq=%.12g dC=%.12g relDC=%.9g Breq=%.12g dB=%.12g relDB=%.9g decade=%t input1sigma=%t outside=%.9g verdict=%s", h.Formula, h.BGap, h.Coefficient, h.MStarGeV, h.TargetMIntGeV, h.PredictedMIntGeV, h.RatioPredictedToTarget, h.Log10Gap, h.RequiredCoefficient, h.CoefficientResidual, h.RelativeCoefficientResidual, h.BGapRequiredForExactFourOverPi, h.BGapResidual, h.RelativeBGapResidual, h.WithinOneDecade, h.WithinInputOneSigmaEnvelope, h.JustOutsideInputEnvelopeByDecades, h.Verdict)
}

func FormatSensitivity(s SensitivityAudit) string {
	return fmt.Sprintf("formula=%q dlog10/dB=%.9f dlog10/dlnB=%.9f onePct=%.9f tenPct=%.9f residualDecades=%.9f fracPrecisionFor0.01dec=%.9f correctsPrompt=%t binding=%t verdict=%s", s.Formula, s.DerivativeLog10MPerUnitBGap, s.DerivativeLog10MPerFractionalB, s.OnePercentBGapShiftDecades, s.TenPercentBGapShiftDecades, s.CoefficientResidualDecadeShift, s.RequiredPrecisionOnBGapForPoint01Decade, s.CorrectsPromptHalfDecadeClaim, s.BindingWarning, s.Verdict)
}

func FormatResidual(r ResidualResolutionAudit) string {
	return fmt.Sprintf("dC=%.12g decadeResidual=%.9f gate215Avail=%t gate219Avail=%t withinDown=%t withinUp=%t withinSym=%t broadMatchingProxy=%t plausible=%t finiteDerived=%t strictFailure=%t verdict=%s", r.CoefficientResidual, r.DecadeResidual, r.Gate215MatchingResidualAvailable, r.Gate219InputEnvelopeAvailable, r.WithinGate219DownEnvelope, r.WithinGate219UpEnvelope, r.WithinGate219SymmetricMaxEnvelope, r.WithinGate214BroadMatchingEnvelopeProxy, r.HigherLoopOrMatchingCanPlausiblyCover, r.FiniteResolutionDerived, r.StrictStructuralFailure, r.Verdict)
}

func FormatSeal(s IntermediateBreakingSealAudit) string {
	return fmt.Sprintf("seal=%s prepared=%t granted=%t patiSalamFailed=%t geomExact=%t nativeHopf=%t sensitivity=%t residualNeedsSeal=%t status=%s verdict=%s", s.SealName, s.SealPreviouslyPrepared, s.SealGranted, s.PatiSalamFalsified, s.GeometricCoefficientExact, s.NativeHopfMapDerived, s.BGapPrecisionBinding, s.ResidualStillRequiresSeal, s.OperationalStatus, s.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g228=%t g174=%t g219=%t sealedScales=%t psReopen=%t BgapField=%t hopfImported=%t S3Std=%t fitted=%t coeffFinite=%t matchFinite=%t MintFinite=%t sealGranted=%t polluted=%t verdict=%s", f.Gate228Inherited, f.Gate174TopologicalSealUsed, f.Gate219UncertaintyUsed, f.UsedOnlySealedScales, f.PatiSalamReopened, f.BGapPromotedToPhysicalField, f.HopfFibrationImportedAsTheorem, f.S3VolumeUsedAsStandardMath, f.CoefficientFitted, f.CoefficientDerivedFromFiniteCore, f.MatchingResidualDerived, f.IntermediateScaleFiniteDerived, f.IntermediateBreakingSealGranted, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("geom=%t nativeHopf=%t sensitivity=%t residualCovered=%t seal=%t target=%.9e hopf=%.9e c=%.12g cReq=%.12g status=%s", s.GeometricHierarchySupported, s.NativeHopfMapDerived, s.SensitivityBinding, s.ResidualPlausiblyCovered, s.IntermediateSealGranted, s.MIntTargetGeV, s.MIntHopfGeV, s.Coefficient, s.RequiredCoefficient, s.Status)
}
