// Package matrixinvariantkoideaudit implements Gate 351:
// Matrix Invariant / Koide-Type Trace Polynomial Audit.
//
// Gate 351 follows the Gate-350 invariant-program pointer.  It audits whether
// the finite ASHA generation data natively imposes Koide-like root-trace or
// characteristic-polynomial constraints on Yukawa singular values.  Empirical
// charged-lepton and quark comparisons are quarantined as comparison data only;
// no observed mass is used to promote a theorem.
package matrixinvariantkoideaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE351-MATRIX-INVARIANT-KOIDE-TYPE-TRACE-POLYNOMIAL-AUDIT"

	StatusKoideInvariantFormalized         = "CONDITIONAL_SUPPORT_KOIDE_TYPE_INVARIANT_FORMALIZED"
	StatusTrialityInvariantSieveExecuted   = "CONDITIONAL_SUPPORT_TRIALITY_INVARIANT_SIEVE_EXECUTED"
	StatusEmpiricalKoideAlignmentCataloged = "CONDITIONAL_SUPPORT_EMPIRICAL_KOIDE_ALIGNMENT_CATALOGED"
	StatusCharacteristicPolynomialAudited  = "CONDITIONAL_SUPPORT_CHARACTERISTIC_POLYNOMIAL_INVARIANTS_AUDITED"
	StatusParameterReductionAssessed       = "CONDITIONAL_SUPPORT_PARAMETER_REDUCTION_ASSESSED"

	StatusTensionLeptonKoideEmpiricalNotNative = "CONDITIONAL_TENSION_CHARGED_LEPTON_KOIDE_MATCH_IS_EMPIRICAL_NOT_NATIVE"
	StatusTensionTauEtaDoesNotMandateKoide     = "CONDITIONAL_TENSION_TAU_ETA_DOES_NOT_MANDATE_KOIDE_TWO_THIRDS"
	StatusTensionBGapNoRootTraceOperator       = "CONDITIONAL_TENSION_BGAP_NO_ROOT_TRACE_OPERATOR_DERIVED"
	StatusTensionQuarkKoideNotUniversal        = "CONDITIONAL_TENSION_QUARK_KOIDE_VARIANTS_NOT_UNIVERSAL"

	StatusFailedMatrixTraceInvariantNotDerived = "FAILED_ROUTE_MATRIX_TRACE_INVARIANT_NOT_DERIVED"
	StatusFailedKoideConstraintNotNative       = "FAILED_ROUTE_KOIDE_CONSTRAINT_NOT_DERIVED_FROM_FINITE_GEOMETRY"
	StatusFailedYukawaCharacteristicNotLocked  = "FAILED_ROUTE_YUKAWA_CHARACTERISTIC_POLYNOMIAL_NOT_LOCKED"
	StatusFailedNoVacuumReductionProved        = "FAILED_ROUTE_NO_ADDITIONAL_VACUUM_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenCoordinatesNotReached     = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 350
	startingVacuumInputs = 15

	bGap                 = 0.102464921191
	topologicalResonance = 4.0 / math.Pi
)

type Span struct {
	AuditID          string
	InheritedGate    int
	AddsEmpiricalFit bool
	Purpose          string
	Verdict          string
}

type KoideInvariant struct {
	Formalized       bool
	Formula          string
	Target           float64
	RootAngleDegrees float64
	Interpretation   string
	Verdict          string
}

type Spectrum struct {
	Name                   string
	Values                 []float64
	Units                  string
	Quarantined            bool
	KoideK                 float64
	DeviationFromTwoThirds float64
	RelativeDeviation      float64
	Verdict                string
}

type TrialitySieve struct {
	Executed                bool
	TauEta                  []float64
	MagnitudeSquaredWeights []float64
	MagnitudeSquaredKoide   float64
	AbsoluteTauKoide        float64
	SignedTraceSum          float64
	SignedTraceNormSquared  float64
	BGap                    float64
	Resonance               float64
	BGapCandidateValues     []float64
	NativeTwoThirdsMandated bool
	Verdict                 string
}

type CharacteristicPolynomialAudit struct {
	Audited                        bool
	Invariants                     []string
	KoideAsPolynomial              string
	OneConstraintCapacity          bool
	RequiresRootTraceOperator      bool
	CharacteristicPolynomialLocked bool
	Verdict                        string
}

type ParameterReduction struct {
	StartingVacuumInputs             int
	ChargedLeptonReductionIfPromoted int
	ReductionProved                  int
	RemainingVacuumInputs            int
	SevenSealTargetReached           bool
	Verdict                          string
}

type Summary struct {
	Executed              bool
	AnyInvariantPromoted  bool
	RemainingVacuumInputs int
	OneLine               string
	Status                string
}

type Analysis struct {
	Span             Span
	Koide            KoideInvariant
	EmpiricalSpectra []Spectrum
	Triality         TrialitySieve
	Characteristic   CharacteristicPolynomialAudit
	Reduction        ParameterReduction
	Summary          Summary
	Truth            string
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
	span := compileSpan()
	koide := formalizeKoide()
	spectra := empiricalComparisons()
	triality := auditTrialityInvariants()
	characteristic := auditCharacteristicPolynomial()
	reduction := assessReduction(triality, characteristic)
	summary := compileSummary(reduction)
	truth := "Gate 351 confirms the right mathematical search space: Koide-type relations are trace/root-trace constraints on full matrices, not single-eigenvalue power laws.  The charged-lepton Koide alignment is cataloged, but ASHA's installed τ_eta, B-gap, and 4/π invariants do not natively derive K=2/3 or lock a Yukawa characteristic polynomial."
	return Analysis{Span: span, Koide: koide, EmpiricalSpectra: spectra, Triality: triality, Characteristic: characteristic, Reduction: reduction, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsEmpiricalFit: false, Purpose: "audit whether ASHA finite generation invariants natively impose Koide-type root-trace/characteristic-polynomial constraints", Verdict: StatusKoideInvariantFormalized}
}

func formalizeKoide() KoideInvariant {
	return KoideInvariant{
		Formalized:       true,
		Formula:          "K(m1,m2,m3)=Tr(M)/(Tr(sqrt(M)))^2 = (m1+m2+m3)/(sqrt(m1)+sqrt(m2)+sqrt(m3))^2",
		Target:           2.0 / 3.0,
		RootAngleDegrees: 45.0,
		Interpretation:   "K=2/3 means the positive root-mass vector has angle 45 degrees with the democratic vector (1,1,1).",
		Verdict:          StatusKoideInvariantFormalized,
	}
}

func empiricalComparisons() []Spectrum {
	// Quarantined comparison values.  Because K is homogeneous, units cancel.
	spectra := []Spectrum{
		{Name: "charged_leptons_pole_proxy", Values: []float64{0.51099895, 105.6583755, 1776.86}, Units: "MeV", Quarantined: true},
		{Name: "up_type_quark_rough_running_proxy", Values: []float64{2.16, 1270.0, 172760.0}, Units: "MeV", Quarantined: true},
		{Name: "down_type_quark_rough_running_proxy", Values: []float64{4.67, 93.4, 4180.0}, Units: "MeV", Quarantined: true},
	}
	for i := range spectra {
		k := KoideK(spectra[i].Values)
		spectra[i].KoideK = k
		spectra[i].DeviationFromTwoThirds = k - 2.0/3.0
		spectra[i].RelativeDeviation = spectra[i].DeviationFromTwoThirds / (2.0 / 3.0)
		if spectra[i].Name == "charged_leptons_pole_proxy" {
			spectra[i].Verdict = strings.Join([]string{StatusEmpiricalKoideAlignmentCataloged, StatusTensionLeptonKoideEmpiricalNotNative}, ";")
		} else {
			spectra[i].Verdict = strings.Join([]string{StatusEmpiricalKoideAlignmentCataloged, StatusTensionQuarkKoideNotUniversal}, ";")
		}
	}
	return spectra
}

func auditTrialityInvariants() TrialitySieve {
	tau := []float64{2, -2, 1}
	mag2 := normalize([]float64{4, 4, 1})
	mag2Koide := KoideK(mag2)
	absTauKoide := KoideK([]float64{2, 2, 1})
	signedSum := 1.0
	signedNorm := 9.0
	// Candidate dimensionless ASHA numbers near Koide-like scales.  None is
	// an installed root-trace operator acting on Yukawa eigenvalues.
	candidates := []float64{
		bGap,
		1.0 / bGap,
		topologicalResonance,
		topologicalResonance * bGap,
		3 * topologicalResonance * bGap,
		2.0 / 3.0,
	}
	sort.Float64s(candidates)
	return TrialitySieve{
		Executed:                true,
		TauEta:                  tau,
		MagnitudeSquaredWeights: mag2,
		MagnitudeSquaredKoide:   mag2Koide,
		AbsoluteTauKoide:        absTauKoide,
		SignedTraceSum:          signedSum,
		SignedTraceNormSquared:  signedNorm,
		BGap:                    bGap,
		Resonance:               topologicalResonance,
		BGapCandidateValues:     candidates,
		NativeTwoThirdsMandated: false,
		Verdict: strings.Join([]string{
			StatusTrialityInvariantSieveExecuted,
			StatusTensionTauEtaDoesNotMandateKoide,
			StatusTensionBGapNoRootTraceOperator,
			StatusFailedKoideConstraintNotNative,
		}, ";"),
	}
}

func auditCharacteristicPolynomial() CharacteristicPolynomialAudit {
	return CharacteristicPolynomialAudit{
		Audited: true,
		Invariants: []string{
			"s1 = Tr(M) = m1+m2+m3",
			"s2 = 1/2[(Tr M)^2 - Tr(M^2)] = m1m2+m1m3+m2m3",
			"s3 = det(M) = m1m2m3",
			"r1 = Tr(sqrt(M)) = sqrt(m1)+sqrt(m2)+sqrt(m3)",
		},
		KoideAsPolynomial:              "3 Tr(M) - 2[Tr(sqrt(M))]^2 = 0",
		OneConstraintCapacity:          true,
		RequiresRootTraceOperator:      true,
		CharacteristicPolynomialLocked: false,
		Verdict: strings.Join([]string{
			StatusCharacteristicPolynomialAudited,
			StatusFailedYukawaCharacteristicNotLocked,
			StatusFailedMatrixTraceInvariantNotDerived,
		}, ";"),
	}
}

func assessReduction(t TrialitySieve, c CharacteristicPolynomialAudit) ParameterReduction {
	reduction := 0
	if t.NativeTwoThirdsMandated && c.CharacteristicPolynomialLocked {
		reduction = 1
	}
	return ParameterReduction{
		StartingVacuumInputs:             startingVacuumInputs,
		ChargedLeptonReductionIfPromoted: 1,
		ReductionProved:                  reduction,
		RemainingVacuumInputs:            startingVacuumInputs - reduction,
		SevenSealTargetReached:           startingVacuumInputs-reduction == 7,
		Verdict: strings.Join([]string{
			StatusParameterReductionAssessed,
			StatusFailedNoVacuumReductionProved,
			StatusFailedSevenCoordinatesNotReached,
		}, ";"),
	}
}

func compileSummary(r ParameterReduction) Summary {
	promoted := r.ReductionProved > 0
	status := StatusFailedMatrixTraceInvariantNotDerived
	if promoted {
		status = StatusParameterReductionAssessed
	}
	return Summary{Executed: true, AnyInvariantPromoted: promoted, RemainingVacuumInputs: r.RemainingVacuumInputs, OneLine: fmt.Sprintf("Gate 351 audits Koide/root-trace invariants but promotes no native Yukawa constraint: vacuum inputs remain %d.", r.RemainingVacuumInputs), Status: status}
}

func KoideK(values []float64) float64 {
	var sum, roots float64
	for _, v := range values {
		if v < 0 {
			return math.NaN()
		}
		sum += v
		roots += math.Sqrt(v)
	}
	if roots == 0 {
		return math.NaN()
	}
	return sum / (roots * roots)
}

func normalize(values []float64) []float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	out := make([]float64, len(values))
	for i, v := range values {
		out[i] = v / total
	}
	return out
}

func Statuses(a Analysis) []string {
	return []string{
		StatusKoideInvariantFormalized,
		StatusTrialityInvariantSieveExecuted,
		StatusEmpiricalKoideAlignmentCataloged,
		StatusCharacteristicPolynomialAudited,
		StatusParameterReductionAssessed,
		StatusTensionLeptonKoideEmpiricalNotNative,
		StatusTensionTauEtaDoesNotMandateKoide,
		StatusTensionBGapNoRootTraceOperator,
		StatusTensionQuarkKoideNotUniversal,
		StatusFailedMatrixTraceInvariantNotDerived,
		StatusFailedKoideConstraintNotNative,
		StatusFailedYukawaCharacteristicNotLocked,
		StatusFailedNoVacuumReductionProved,
		StatusFailedSevenCoordinatesNotReached,
	}
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("%s | inherited gate=%d | adds_empirical_fit=%t | %s", s.AuditID, s.InheritedGate, s.AddsEmpiricalFit, s.Purpose)
}

func FormatKoide(k KoideInvariant) string {
	return fmt.Sprintf("formalized=%t; target=%.15f; root_angle=%.1f°; formula=%s", k.Formalized, k.Target, k.RootAngleDegrees, k.Formula)
}

func FormatSpectrum(s Spectrum) string {
	return fmt.Sprintf("%s values=%v %s; K=%.15f; ΔK=%.15e; rel=%.15e; quarantined=%t", s.Name, s.Values, s.Units, s.KoideK, s.DeviationFromTwoThirds, s.RelativeDeviation, s.Quarantined)
}

func FormatTriality(t TrialitySieve) string {
	return fmt.Sprintf("τ_eta=%v; normalized |τ|²=%v; K(|τ|²)=%.15f; K(|τ|)=%.15f; signed_sum=%.1f; signed_norm²=%.1f; B_gap=%.12f; 4/π=%.12f; mandates_2/3=%t", t.TauEta, t.MagnitudeSquaredWeights, t.MagnitudeSquaredKoide, t.AbsoluteTauKoide, t.SignedTraceSum, t.SignedTraceNormSquared, t.BGap, t.Resonance, t.NativeTwoThirdsMandated)
}

func FormatCharacteristic(c CharacteristicPolynomialAudit) string {
	return fmt.Sprintf("audited=%t; invariants=%s; koide_polynomial=%s; one_constraint_capacity=%t; root_trace_operator_required=%t; characteristic_locked=%t", c.Audited, strings.Join(c.Invariants, " | "), c.KoideAsPolynomial, c.OneConstraintCapacity, c.RequiresRootTraceOperator, c.CharacteristicPolynomialLocked)
}

func FormatReduction(r ParameterReduction) string {
	return fmt.Sprintf("starting=%d; charged_lepton_reduction_if_promoted=%d; reduction_proved=%d; remaining=%d; seven_target=%t", r.StartingVacuumInputs, r.ChargedLeptonReductionIfPromoted, r.ReductionProved, r.RemainingVacuumInputs, r.SevenSealTargetReached)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t; any_invariant_promoted=%t; remaining=%d; %s", s.Executed, s.AnyInvariantPromoted, s.RemainingVacuumInputs, s.OneLine)
}
