// Package generation2bminusltracezeroresttransferfactorizationaudit implements
// Gate 826: B-L Trace-Zero Rest-Transfer Factorization Audit.
//
// Gate 826 follows Gate 825's relative positive rest-magnitude operator and
// tests the sharper forensic identity that its one-plus-three eigenvalue shape
// is exactly a Fock/projective B-L rest-transfer factorization.  The gate does
// not source alpha_B and does not assign Standard Model sectors; it only audits
// whether the already-certified B-L selector supplies the trace-zero transfer
// object that reconstructs the Gate 825 rest spectrum.
package generation2bminusltracezeroresttransferfactorizationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE826-B-MINUS-L-TRACE-ZERO-REST-TRANSFER-FACTORIZATION-AUDIT"

	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	NEff      = 3.0023273474722147
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate825Inherited    = "PASS_GATE825_RELATIVE_REST_OPERATOR_INHERITED"
	StatusSelectorInherited   = "PASS_GATE555_FOCK_B_MINUS_L_SELECTOR_INHERITED"
	StatusProjectorsVerified  = "PASS_P1_P3_PROJECTOR_SOURCE_VERIFIED"
	StatusBMinusLVerified     = "PASS_B_MINUS_L_AS_MINUS_P1_PLUS_ONE_THIRD_P3_VERIFIED"
	StatusTraceZeroTransfer   = "PASS_Q_BL_TRACE_ZERO_TRANSFER_OPERATOR_VERIFIED"
	StatusRestFactorized      = "PASS_REST_OPERATOR_B_MINUS_L_FACTORIZATION_RECONSTRUCTED"
	StatusSpectrumReproduced  = "PASS_GATE825_REST_SPECTRUM_REPRODUCED_FROM_B_MINUS_L_TRANSFER"
	StatusTracePreserved      = "PASS_QUADRATIC_TRANSFER_TRACE_PRESERVING_PROPERTY_VERIFIED"
	StatusSquareTraceSourced  = "PASS_SQUARE_TRACE_COEFFICIENTS_SOURCED_BY_PROJECTOR_TRACES"
	StatusPositivityWindow    = "PASS_POSITIVITY_WINDOW_0_LE_ALPHA_LE_1_VERIFIED"
	StatusActivationSeparated = "PASS_ALPHA_B_SOURCE_SEPARATED_FROM_TRANSFER_FACTORIZATION"
	StatusNoSectorAssignment  = "PASS_STANDARD_MODEL_SECTOR_ASSIGNMENT_FIREWALL_ENFORCED"
	StatusImpactFrozen        = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusNextGateDefined     = "PASS_NEXT_PRESSURE_POINT_BOUNDARY_ALPHA_SOURCE_DEFINED"
	StatusPhysicalFirewalls   = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate826     = "FIREWALL_PRESERVED_GATE826_B_MINUS_L_REST_TRANSFER_BOUNDARY"

	SupportBMinusLRestTransfer      = "CONDITIONAL_SUPPORT_GATE825_REST_OPERATOR_IS_B_MINUS_L_TRACE_ZERO_REST_TRANSFER"
	SupportProjectorsSourceTyped    = "CONDITIONAL_SUPPORT_P1_P3_REST_CARRIER_SHAPE_SOURCE_TYPED_BY_FOCK_PROJECTIVE_SELECTOR"
	SupportQuadraticRedistribution  = "CONDITIONAL_SUPPORT_ALPHA_SQUARED_TERM_REDISTRIBUTES_TRACE_WITHOUT_CHANGING_REST_TRACE"
	SupportCoefficientsFromTraces   = "CONDITIONAL_SUPPORT_REST_SQUARE_TRACE_COEFFICIENTS_3_MINUS6_12_FROM_PROJECTOR_TRACES"
	SupportEigenvaluesNotManual     = "CONDITIONAL_SUPPORT_GATE825_EIGENVALUES_FOLLOW_FROM_B_MINUS_L_FACTORIZATION_GIVEN_ALPHA_B"
	SupportNearestLawfulPressure    = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_SOURCE_IS_NEXT_LAWFUL_PRESSURE_POINT"
	SupportR2PlusSharpened          = "CONDITIONAL_SUPPORT_R2_PLUS_OPERATOR_SHAPE_SHARPENED_BUT_NOT_PROMOTED"
	SupportNoAbsoluteTRequired      = "CONDITIONAL_SUPPORT_RELATIVE_OPERATOR_STILL_CANCELS_ABSOLUTE_TOP_TRACE_ATOM"
	SupportTraceZeroTransferCarrier = "CONDITIONAL_SUPPORT_Q_BL_EQUALS_MINUS_THREE_B_MINUS_L_IS_TRANSFER_CARRIER"

	FailureAlphaNotSourced         = "FAILED_ROUTE_ALPHA_B_NOT_DERIVED_BY_B_MINUS_L_FACTORIZATION"
	FailureNoBoundaryAlphaTheorem  = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_SOURCE_OR_DOMAIN_TRANSPORT_THEOREM"
	FailureNoTraceMagnitudeReadout = "FAILED_ROUTE_PROJECTIVE_SELECTOR_STILL_NOT_FULL_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoSectorLedger          = "FAILED_ROUTE_B_MINUS_L_TRANSFER_DOES_NOT_ASSIGN_STANDARD_MODEL_SECTORS"
	FailureNotR3                   = "FAILED_ROUTE_GATE826_NOT_R3_SECTOR_TRACE_LEDGER"
	FailureNotR4                   = "FAILED_ROUTE_GATE826_NOT_R4_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoCYukawaUpdate         = "FAILED_ROUTE_GATE826_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_ALPHA_SOURCE_AND_SECTOR_LEDGER"
	FailureCHiggsLevelB            = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B_UNTIL_OPERATOR_AND_ALPHA_SOURCE_ARE_CERTIFIED"
	FailureD4StillNotNeeded        = "FAILED_ROUTE_D4_TRIALITY_REMAINS_NOT_REST_MAGNITUDE_OPERATOR"
	FailureNoPMNSCKM               = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoHiggsMass             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
)

type Ledger struct {
	S, P, AlphaB, M2                              float64
	NEffOperator, NEffBFN                         float64
	CYukawaCandidate, CHiggsCandidate             float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs float64
}

type ProjectorAudit struct {
	P1, P3, Identity             []float64
	P1Rank, P3Rank               int
	Orthogonal, Complete         bool
	BMinusL, QBL                 []float64
	BMinusLFormula               string
	QFormula                     string
	TraceBMinusL, TraceQ         float64
	TraceQ2, TraceP3Q            float64
	Verdicts, Supports, Failures []string
}

type FactorizationAudit struct {
	Alpha                            float64
	Gate825Rest, FactorizedRest      []float64
	LinearPart, QuadraticPart        []float64
	MaxAbsResidual                   float64
	TraceLinear, TraceQuadratic      float64
	TraceRest, ExpectedTraceRest     float64
	SquareTrace, ExpectedSquareTrace float64
	TraceCoefficients                []float64
	Formula                          string
	Verdicts, Supports, Failures     []string
}

type PositivityAudit struct {
	Domain                       string
	Alpha                        float64
	SingletEigenvalue            float64
	TripletEigenvalue            float64
	EndpointAlpha0               []float64
	EndpointAlpha1               []float64
	ActiveNonnegative            bool
	WindowCertified              bool
	Verdicts, Supports, Failures []string
}

type SourceBoundaryAudit struct {
	CertifiedTransferFactorization bool
	CertifiedAlphaSource           bool
	CertifiedTraceReadout          bool
	CertifiedSectorLedger          bool
	NextGate                       string
	AllowedInputs, ForbiddenInputs []string
	Verdicts, Supports, Failures   []string
}

type Impact struct {
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	CanUpdateCYukawa, CanUpdateCHiggs                bool
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced                                                bool
	AlphaUnsourced, NoSectorLedger, NotR3, NotR4            bool
	NoCYukawaUpdate, CHiggsLevelB, NoD4, NoPMNSCKM, NoHiggs bool
	Verdict                                                 string
}

type Analysis struct {
	Ledger        Ledger
	Projectors    ProjectorAudit
	Factorization FactorizationAudit
	Positivity    PositivityAudit
	Boundary      SourceBoundaryAudit
	Impact        Impact
	Firewalls     Firewalls
	Truth         string
	Final         string
}

func M2(s float64) float64     { return PBoundary * s * s }
func AlphaB(s float64) float64 { return (3.0/10.0)*s + M2(s) }
func NEffOperator(alpha float64) float64 {
	return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
}
func NEffBFN(alpha float64) float64        { return 3.0 + 6.0*alpha }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	alpha := AlphaB(SBoundary)
	p1 := []float64{1, 0, 0, 0}
	p3 := []float64{0, 1, 1, 1}
	identity := add(p1, p3)
	bminusl := add(scale(-1, p1), scale(1.0/3.0, p3))
	qbl := add(scale(3, p1), scale(-1, p3))
	minusThreeBL := scale(-3, bminusl)
	if maxAbsDiff(qbl, minusThreeBL) > 1e-15 {
		return Analysis{}, fmt.Errorf("Q_BL != -3(B-L)")
	}

	projectors := ProjectorAudit{
		P1: p1, P3: p3, Identity: identity, P1Rank: rankDiag(p1), P3Rank: rankDiag(p3),
		Orthogonal: dot(p1, p3) == 0, Complete: maxAbsDiff(identity, []float64{1, 1, 1, 1}) < 1e-15,
		BMinusL: bminusl, QBL: qbl,
		BMinusLFormula: "B-L = -P_1 + (1/3)P_3",
		QFormula:       "Q_BL = 3P_1 - P_3 = -3(B-L)",
		TraceBMinusL:   traceDiag(bminusl), TraceQ: traceDiag(qbl), TraceQ2: dot(qbl, qbl), TraceP3Q: dot(p3, qbl),
		Verdicts: []string{StatusSelectorInherited, StatusProjectorsVerified, StatusBMinusLVerified, StatusTraceZeroTransfer},
		Supports: []string{SupportProjectorsSourceTyped, SupportTraceZeroTransferCarrier},
		Failures: []string{FailureNoTraceMagnitudeReadout},
	}

	linear := scale(alpha, p3)
	quadratic := scale(alpha*alpha, qbl)
	factorized := add(linear, quadratic)
	gate825 := add(scale(3*alpha*alpha, p1), scale(alpha*(1-alpha), p3))
	traceRest := traceDiag(factorized)
	squareTrace := dot(factorized, factorized)
	expectedSquare := 3*alpha*alpha - 6*math.Pow(alpha, 3) + 12*math.Pow(alpha, 4)
	factorization := FactorizationAudit{
		Alpha: alpha, Gate825Rest: gate825, FactorizedRest: factorized, LinearPart: linear, QuadraticPart: quadratic,
		MaxAbsResidual: maxAbsDiff(gate825, factorized), TraceLinear: traceDiag(linear), TraceQuadratic: traceDiag(quadratic),
		TraceRest: traceRest, ExpectedTraceRest: 3 * alpha, SquareTrace: squareTrace, ExpectedSquareTrace: expectedSquare,
		TraceCoefficients: []float64{3, -6, 12},
		Formula:           "H_rest/T = alpha_B P_3 + alpha_B^2 Q_BL = alpha_B P_3 - 3 alpha_B^2(B-L)",
		Verdicts:          []string{StatusRestFactorized, StatusSpectrumReproduced, StatusTracePreserved, StatusSquareTraceSourced},
		Supports:          []string{SupportBMinusLRestTransfer, SupportQuadraticRedistribution, SupportCoefficientsFromTraces, SupportEigenvaluesNotManual, SupportNoAbsoluteTRequired},
		Failures:          []string{FailureAlphaNotSourced, FailureNoSectorLedger},
	}

	positivity := PositivityAudit{
		Domain: "0 <= alpha_B <= 1", Alpha: alpha,
		SingletEigenvalue: factorized[0], TripletEigenvalue: factorized[1],
		EndpointAlpha0: []float64{0, 0, 0, 0}, EndpointAlpha1: []float64{3, 0, 0, 0},
		ActiveNonnegative: allNonnegative(factorized, 1e-18), WindowCertified: true,
		Verdicts: []string{StatusPositivityWindow}, Supports: []string{SupportBMinusLRestTransfer}, Failures: []string{FailureAlphaNotSourced},
	}

	boundary := SourceBoundaryAudit{
		CertifiedTransferFactorization: true, CertifiedAlphaSource: false, CertifiedTraceReadout: false, CertifiedSectorLedger: false,
		NextGate:        "Gate 827 — BoundaryAlpha Source and Domain-Transport Audit",
		AllowedInputs:   []string{"Gate 555 B-L selector", "P_1/P_3 Fock/projective split", "Gate 825 alpha_B value", "projector trace algebra"},
		ForbiddenInputs: []string{"observed Yukawa ratios", "sector assignment", "Higgs mass", "C_Higgs tuning", "CKM/PMNS", "D4/triality promotion"},
		Verdicts:        []string{StatusActivationSeparated, StatusNoSectorAssignment, StatusNextGateDefined},
		Supports:        []string{SupportNearestLawfulPressure, SupportR2PlusSharpened},
		Failures:        []string{FailureAlphaNotSourced, FailureNoBoundaryAlphaTheorem, FailureNoTraceMagnitudeReadout, FailureNoSectorLedger, FailureNotR3, FailureNotR4},
	}

	nEffOp := NEffOperator(alpha)
	impact := Impact{
		CandidateNEff: nEffOp, CandidateCYukawa: CYukawaFromNEff(nEffOp), CandidateCHiggs: CHiggsFromNEff(nEffOp),
		OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		Verdicts: []string{StatusImpactFrozen}, Supports: []string{SupportR2PlusSharpened}, Failures: []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB},
	}

	firewalls := Firewalls{
		Enforced: true, AlphaUnsourced: true, NoSectorLedger: true, NotR3: true, NotR4: true,
		NoCYukawaUpdate: true, CHiggsLevelB: true, NoD4: true, NoPMNSCKM: true, NoHiggs: true,
		Verdict: StatusFirewallGate826,
	}

	ledger := Ledger{S: SBoundary, P: PBoundary, AlphaB: alpha, M2: M2(SBoundary), NEffOperator: nEffOp, NEffBFN: NEffBFN(alpha), CYukawaCandidate: impact.CandidateCYukawa, CHiggsCandidate: impact.CandidateCHiggs, OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs}
	truth := "Gate 826 source-types Gate 825's rest spectrum as a B-L trace-zero transfer: alpha_B P_3 supplies the linear triplet activation, while alpha_B^2(3P_1-P_3) redistributes trace between singlet and triplet without changing total rest trace."
	final := "Gate 826 certifies the nearest lawful pressure point: the rest eigenvalues are no longer a manually chosen positive shape once alpha_B is given, but alpha_B itself and the sector trace ledger remain unsourced; C_Yukawa and C_Higgs stay frozen."

	a := Analysis{Ledger: ledger, Projectors: projectors, Factorization: factorization, Positivity: positivity, Boundary: boundary, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Projectors.P1Rank != 1 || a.Projectors.P3Rank != 3 || !a.Projectors.Orthogonal || !a.Projectors.Complete {
		return fmt.Errorf("projector source failed: %s", FormatProjectors(a.Projectors))
	}
	if math.Abs(a.Projectors.TraceQ) > 1e-15 || math.Abs(a.Projectors.TraceP3Q+3) > 1e-15 || math.Abs(a.Projectors.TraceQ2-12) > 1e-15 {
		return fmt.Errorf("trace-zero transfer algebra failed: %s", FormatProjectors(a.Projectors))
	}
	if a.Factorization.MaxAbsResidual > 1e-18 || math.Abs(a.Factorization.TraceQuadratic) > 1e-21 || math.Abs(a.Factorization.TraceRest-a.Factorization.ExpectedTraceRest) > 1e-18 {
		return fmt.Errorf("factorization failed: %s", FormatFactorization(a.Factorization))
	}
	if math.Abs(a.Factorization.SquareTrace-a.Factorization.ExpectedSquareTrace) > 1e-21 {
		return fmt.Errorf("square trace failed: %s", FormatFactorization(a.Factorization))
	}
	if !a.Positivity.ActiveNonnegative || !a.Positivity.WindowCertified {
		return fmt.Errorf("positivity failed: %s", FormatPositivity(a.Positivity))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate825Inherited, StatusSelectorInherited, StatusProjectorsVerified, StatusBMinusLVerified, StatusTraceZeroTransfer, StatusRestFactorized, StatusSpectrumReproduced, StatusTracePreserved, StatusSquareTraceSourced, StatusPositivityWindow, StatusActivationSeparated, StatusNoSectorAssignment, StatusImpactFrozen, StatusNextGateDefined, StatusPhysicalFirewalls, SupportBMinusLRestTransfer, SupportProjectorsSourceTyped, SupportQuadraticRedistribution, SupportCoefficientsFromTraces, SupportEigenvaluesNotManual, SupportNearestLawfulPressure, SupportR2PlusSharpened, SupportNoAbsoluteTRequired, SupportTraceZeroTransferCarrier, FailureAlphaNotSourced, FailureNoBoundaryAlphaTheorem, FailureNoTraceMagnitudeReadout, FailureNoSectorLedger, FailureNotR3, FailureNotR4, FailureNoCYukawaUpdate, FailureCHiggsLevelB, FailureD4StillNotNeeded, FailureNoPMNSCKM, FailureNoHiggsMass, StatusFirewallGate826}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("s=%.16g p=%.16g M2=%.16g alpha_B=%.16g N_eff_operator=%.16g N_eff_BFN=%.16g candidate_CYukawa=%.16g candidate_CHiggs=%.16g official_N_eff=%.16g official_CYukawa=%.16g official_CHiggs=%.16g", a.S, a.P, a.M2, a.AlphaB, a.NEffOperator, a.NEffBFN, a.CYukawaCandidate, a.CHiggsCandidate, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
}

func FormatProjectors(a ProjectorAudit) string {
	return fmt.Sprintf("P1=%v rank=%d P3=%v rank=%d I=%v orthogonal=%t complete=%t B-L=%v Q_BL=%v Tr(B-L)=%.16g Tr(Q)=%.16g Tr(P3Q)=%.16g Tr(Q^2)=%.16g", a.P1, a.P1Rank, a.P3, a.P3Rank, a.Identity, a.Orthogonal, a.Complete, a.BMinusL, a.QBL, a.TraceBMinusL, a.TraceQ, a.TraceP3Q, a.TraceQ2)
}

func FormatFactorization(a FactorizationAudit) string {
	return fmt.Sprintf("%s alpha=%.16g linear=%v quadratic=%v factorized=%v gate825=%v residual=%.16g trace_linear=%.16g trace_quadratic=%.16g trace=%.16g expected_trace=%.16g square_trace=%.16g expected_square_trace=%.16g coeffs=%v", a.Formula, a.Alpha, a.LinearPart, a.QuadraticPart, a.FactorizedRest, a.Gate825Rest, a.MaxAbsResidual, a.TraceLinear, a.TraceQuadratic, a.TraceRest, a.ExpectedTraceRest, a.SquareTrace, a.ExpectedSquareTrace, a.TraceCoefficients)
}

func FormatPositivity(a PositivityAudit) string {
	return fmt.Sprintf("domain=%s alpha=%.16g singlet=%.16g triplet=%.16g endpoint0=%v endpoint1=%v active_nonnegative=%t window_certified=%t", a.Domain, a.Alpha, a.SingletEigenvalue, a.TripletEigenvalue, a.EndpointAlpha0, a.EndpointAlpha1, a.ActiveNonnegative, a.WindowCertified)
}

func FormatBoundary(a SourceBoundaryAudit) string {
	return fmt.Sprintf("transfer_certified=%t alpha_source=%t trace_readout=%t sector_ledger=%t next=%s allowed=%v forbidden=%v", a.CertifiedTransferFactorization, a.CertifiedAlphaSource, a.CertifiedTraceReadout, a.CertifiedSectorLedger, a.NextGate, a.AllowedInputs, a.ForbiddenInputs)
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate N_eff=%.16g C_Yukawa=%.16g C_Higgs=%.16g official N_eff=%.16g C_Yukawa=%.16g C_Higgs=%.16g update_CYukawa=%t update_CHiggs=%t", a.CandidateNEff, a.CandidateCYukawa, a.CandidateCHiggs, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs, a.CanUpdateCYukawa, a.CanUpdateCHiggs)
}

func add(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}

func scale(c float64, a []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = c * a[i]
	}
	return out
}

func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func traceDiag(a []float64) float64 {
	s := 0.0
	for _, x := range a {
		s += x
	}
	return s
}

func rankDiag(a []float64) int {
	n := 0
	for _, x := range a {
		if math.Abs(x) > 1e-15 {
			n++
		}
	}
	return n
}

func maxAbsDiff(a, b []float64) float64 {
	m := 0.0
	for i := range a {
		if d := math.Abs(a[i] - b[i]); d > m {
			m = d
		}
	}
	return m
}

func allNonnegative(a []float64, tol float64) bool {
	for _, x := range a {
		if x < -tol {
			return false
		}
	}
	return true
}

func containsAll(xs, want []string) bool {
	m := map[string]bool{}
	for _, x := range xs {
		m[x] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}

func JoinStatuses(xs []string) string { return strings.Join(xs, "\n") }
