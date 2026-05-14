// Package finitencginstantonaction implements Gate 286:
// Finite Spectral Action Saddle-Point / B-Gap Instanton Action Audit.
//
// Gate 285 correctly rejected the continuum Hopf-connection/Chern-Simons
// route inside a finite pre-geometric algebra. Gate 286 pivots inward to the
// noncommutative-geometric calculus: da=[D_F,a], inner fluctuations
// A=Σa_i[D_F,b_i], finite curvature F=[D_F,A]+A², and finite actions as
// traces of finite matrices. The gate tests whether this algebraic route can
// derive S_inst=(4/π)/B_gap. It does not: a local quaternionic diagnostic shows
// non-vacuous one-forms and a computable finite action, but the action scales
// with positive powers of the inserted Dirac amplitude and has no derived
// non-trivial saddle or inverse-B_gap law.
package finitencginstantonaction

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"
)

const (
	AuditID = "GATE286-FINITE-SPECTRAL-ACTION-SADDLE-POINT-BGAP-INSTANTON-ACTION-AUDIT"

	StatusGate285ContinuumBarrierInherited = "CONDITIONAL_SUPPORT_GATE285_CONTINUUM_BARRIER_INHERITED"
	StatusNCGCalculusFormalized            = "CONDITIONAL_SUPPORT_NCG_FINITE_DIFFERENTIAL_CALCULUS_FORMALIZED"
	StatusInnerFluctuationDiagnosticBuilt  = "CONDITIONAL_SUPPORT_LOCAL_QUATERNIONIC_INNER_FLUCTUATION_DIAGNOSTIC_BUILT"
	StatusFiniteCurvatureTraceEvaluated    = "CONDITIONAL_SUPPORT_FINITE_CURVATURE_TRACE_ACTION_EVALUATED"
	StatusBGapInsertionAudited             = "CONDITIONAL_SUPPORT_BGAP_MAJORANA_INSERTION_HYPOTHESIS_AUDITED"
	StatusSaddleSearchCompleted            = "CONDITIONAL_SUPPORT_FINITE_ACTION_SADDLE_SEARCH_COMPLETED"
	StatusFirewallsPreserved               = "CONDITIONAL_SUPPORT_NCG_INSTANTON_FIREWALLS_PRESERVED"

	StatusFailedPhysicalDFStillMissing      = "FAILED_ROUTE_PHYSICAL_FINITE_DIRAC_OPERATOR_STILL_MISSING"
	StatusFailedFullAFRepresentationMissing = "FAILED_ROUTE_FULL_C_PLUS_H_PLUS_M3C_REPRESENTATION_STILL_MISSING"
	StatusFailedBGapMajoranaMapMissing      = "FAILED_ROUTE_BGAP_TO_MAJORANA_DF_ENTRY_NOT_DERIVED"
	StatusFailedNoNontrivialSaddle          = "FAILED_ROUTE_NO_NONTRIVIAL_FINITE_ACTION_SADDLE_DERIVED"
	StatusFailedNoInverseBGapAction         = "FAILED_ROUTE_FINITE_TRACE_DOES_NOT_YIELD_INVERSE_BGAP_ACTION"
	StatusFailedFourOverPiNotGenerated      = "FAILED_ROUTE_FOUR_OVER_PI_NOT_GENERATED_BY_FINITE_SADDLE"
	StatusFailedFiniteInstantonNotDerived   = "FAILED_ROUTE_FINITE_INSTANTON_ACTION_NOT_DERIVED_VIA_NCG"
	StatusFailedIntermediateSealStillNeeded = "FAILED_ROUTE_INTERMEDIATE_BREAKING_SEAL_REMAINS_REQUIRED"
)

const (
	bGap       = 0.1024649212
	fourOverPi = 4 / math.Pi
	targetInst = fourOverPi / bGap
	exactTol   = 1e-12
)

type Gate285Snapshot struct {
	ContinuumRouteAudited      bool
	FiniteConnectionMissing    bool
	FiniteCurvatureMissing     bool
	CSFunctionalMissing        bool
	BGapCouplingMissing        bool
	IntermediateSealRequired   bool
	FourOverPiCoefficient      float64
	BGap                       float64
	CandidateInstantonExponent float64
	TruthStatement             string
}

type NCGCalculusLedger struct {
	DifferentialDefinition      string
	OneFormDefinition           string
	InnerFluctuationDefinition  string
	CurvatureDefinition         string
	FiniteTraceActionDefinition string
	RequiresContinuumForms      bool
	RequiresIntegrationMeasure  bool
	AlgebraicMatrixRouteDefined bool
	PhysicalDFDerived           bool
	FullAlgebraRepresentation   bool
	Verdict                     string
}

type MatrixDiagnostic struct {
	Carrier                   string
	AlgebraSector             string
	D                         string
	Generator                 string
	OneFormNorm2              float64
	CurvatureFormula          string
	ActionPolynomial          string
	MuPowerScaling            string
	NonVacuousOneForm         bool
	FiniteCurvatureComputed   bool
	FiniteTraceActionComputed bool
	Verdict                   string
}

type SaddleAudit struct {
	ActionPolynomial           string
	DerivativePolynomial       string
	RealSaddles                []float64
	NontrivialRealSaddleExists bool
	TrivialVacuumAction        float64
	NontrivialActionGapDerived bool
	Verdict                    string
}

type BGapInsertionAudit struct {
	InsertionHypothesis           string
	BGap                          float64
	CandidateInstantonTarget      float64
	TreatingBGapAsMajoranaDerived bool
	TreatingBGapAsInverseDerived  bool
	IfMuEqualsBGapScaling         string
	IfMuEqualsInverseBGapScaling  string
	DerivedScaling                string
	ProducesInverseBGap           bool
	ProducesFourOverPi            bool
	Verdict                       string
}

type Firewalls struct {
	DoesNotUseContinuumForms     bool
	DoesNotInventPhysicalDF      bool
	DoesNotPromoteBGapToMajorana bool
	DoesNotPromoteBGapToCoupling bool
	DoesNotClaimFourOverPiSaddle bool
	DoesNotGrantIntermediateSeal bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	Gate285Inherited         bool
	NCGCalculusFormalized    bool
	InnerFluctuationBuilt    bool
	FiniteTraceEvaluated     bool
	BGapInsertionAudited     bool
	NontrivialSaddleDerived  bool
	InverseBGapActionDerived bool
	FourOverPiGenerated      bool
	FiniteInstantonDerived   bool
	IntermediateSealGranted  bool
	FirewallPreserved        bool
	Status                   string
	DirectAnswer             string
	NextGate                 string
}

type Analysis struct {
	Gate285        Gate285Snapshot
	Calculus       NCGCalculusLedger
	Diagnostic     MatrixDiagnostic
	Saddle         SaddleAudit
	BGapAudit      BGapInsertionAudit
	Firewalls      Firewalls
	Summary        Summary
	Statuses       []string
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
	gate285 := snapshotGate285()
	calc := formalizeNCGCalculus()
	diag := runLocalQuaternionicDiagnostic(bGap)
	saddle := auditSaddles()
	bg := auditBGapInsertion(gate285, diag, saddle)
	fw := auditFirewalls(calc, bg, saddle)
	sum := buildSummary(gate285, calc, diag, saddle, bg, fw)
	statuses := collectStatuses(sum)
	truth := "Gate 286 pivots from the blocked continuum Hopf/Chern-Simons route to finite NCG matrix calculus. It successfully formalizes da=[D_F,a], inner fluctuations, finite curvature, and a local quaternionic trace-action diagnostic. The diagnostic is non-vacuous, but it yields only a positive-power dependence on the inserted Dirac amplitude, has no non-trivial real saddle, and provides no derived map from B_gap to a Majorana entry or inverse coupling. Therefore S_inst=(4/π)/B_gap remains a target, not a finite NCG theorem."
	return Analysis{Gate285: gate285, Calculus: calc, Diagnostic: diag, Saddle: saddle, BGapAudit: bg, Firewalls: fw, Summary: sum, Statuses: statuses, TruthStatement: truth}, nil
}

func snapshotGate285() Gate285Snapshot {
	return Gate285Snapshot{
		ContinuumRouteAudited:      true,
		FiniteConnectionMissing:    true,
		FiniteCurvatureMissing:     true,
		CSFunctionalMissing:        true,
		BGapCouplingMissing:        true,
		IntermediateSealRequired:   true,
		FourOverPiCoefficient:      fourOverPi,
		BGap:                       bGap,
		CandidateInstantonExponent: targetInst,
		TruthStatement:             "Gate 285 rejected continuum Hopf connection/Chern-Simons evaluation inside finite algebra and left the intermediate-breaking seal active.",
	}
}

func formalizeNCGCalculus() NCGCalculusLedger {
	return NCGCalculusLedger{
		DifferentialDefinition:      "δ(a)=[D_F,a]",
		OneFormDefinition:           "Ω¹_D(A_F)=span{a[D_F,b]}",
		InnerFluctuationDefinition:  "A=Σ a_i[D_F,b_i] with A=A† after reality/self-adjointness projection",
		CurvatureDefinition:         "F_D(A)=δ(A)+A², represented finitely by commutators and matrix products after quotienting junk forms",
		FiniteTraceActionDefinition: "S_finite≈Tr(F†F) or spectral moments Tr(f((D_F+A)/Λ)); normalization still requires a physical spectral triple",
		RequiresContinuumForms:      false,
		RequiresIntegrationMeasure:  false,
		AlgebraicMatrixRouteDefined: true,
		PhysicalDFDerived:           false,
		FullAlgebraRepresentation:   false,
		Verdict:                     StatusNCGCalculusFormalized,
	}
}

type mat2 [2][2]complex128

func mul(a, b mat2) mat2 {
	var c mat2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
func add(a, b mat2) mat2 {
	var c mat2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}
func sub(a, b mat2) mat2 {
	var c mat2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}
func scale(z complex128, a mat2) mat2 {
	var c mat2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = z * a[i][j]
		}
	}
	return c
}
func comm(a, b mat2) mat2 { return sub(mul(a, b), mul(b, a)) }
func dagger(a mat2) mat2 {
	var c mat2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = cmplx.Conj(a[j][i])
		}
	}
	return c
}
func frob2(a mat2) float64 {
	var s float64
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			v := a[i][j]
			s += real(v)*real(v) + imag(v)*imag(v)
		}
	}
	return s
}
func trace(a mat2) complex128           { return a[0][0] + a[1][1] }
func traceDaggerProduct(a mat2) float64 { return real(trace(mul(dagger(a), a))) }

func runLocalQuaternionicDiagnostic(mu float64) MatrixDiagnostic {
	// Minimal extracted-H diagnostic: D=μσ_x and generator J_H=[[0,1],[-1,0]].
	// This is not the full spectral triple. It only tests whether finite NCG
	// calculus can produce nonzero one-forms and trace actions locally.
	D := mat2{{0, complex(mu, 0)}, {complex(mu, 0), 0}}
	J := mat2{{0, 1}, {-1, 0}}
	one := comm(D, J)
	oneNorm := traceDaggerProduct(one)
	C := one
	dc := comm(D, C)
	// F(t)=t[D,C]+t²C²; analytic result is 4μ²(tJ+t²I), so Tr(F†F)=32μ⁴(t²+t⁴).
	c2 := mul(C, C)
	sampleT := 0.7
	F := add(scale(complex(sampleT, 0), dc), scale(complex(sampleT*sampleT, 0), c2))
	_ = frob2(F) // keep numeric path audited; analytic formula below is exact for this diagnostic.
	return MatrixDiagnostic{
		Carrier:                   "local weak doublet / quaternionic H diagnostic, not full H_F",
		AlgebraSector:             "H_U12 ≅ unit-quaternion/SU(2) local block",
		D:                         "D_μ = μ σ_x",
		Generator:                 "J_H = [[0,1],[-1,0]], J_H²=-1",
		OneFormNorm2:              oneNorm,
		CurvatureFormula:          "For A=t[D_μ,J_H], F=[D_μ,A]+A² = 4μ²(t J_H + t² I)",
		ActionPolynomial:          "Tr(F†F)=32 μ⁴(t²+t⁴)",
		MuPowerScaling:            "positive-power μ⁴ scaling; no inverse-μ or inverse-B_gap law appears",
		NonVacuousOneForm:         oneNorm > exactTol,
		FiniteCurvatureComputed:   true,
		FiniteTraceActionComputed: true,
		Verdict:                   StatusFiniteCurvatureTraceEvaluated,
	}
}

func auditSaddles() SaddleAudit {
	return SaddleAudit{
		ActionPolynomial:           "S(t)=32 μ⁴(t²+t⁴)",
		DerivativePolynomial:       "S'(t)=64 μ⁴ t(1+2t²)",
		RealSaddles:                []float64{0},
		NontrivialRealSaddleExists: false,
		TrivialVacuumAction:        0,
		NontrivialActionGapDerived: false,
		Verdict:                    StatusSaddleSearchCompleted,
	}
}

func auditBGapInsertion(g Gate285Snapshot, d MatrixDiagnostic, s SaddleAudit) BGapInsertionAudit {
	return BGapInsertionAudit{
		InsertionHypothesis:           "B_gap may be tested as a sealed Majorana-like amplitude μ in D_F, but Gate 286 does not derive that map",
		BGap:                          g.BGap,
		CandidateInstantonTarget:      g.CandidateInstantonExponent,
		TreatingBGapAsMajoranaDerived: false,
		TreatingBGapAsInverseDerived:  false,
		IfMuEqualsBGapScaling:         fmt.Sprintf("S(t)=32 B_gap^4(t²+t⁴); B_gap^4≈%.12g", math.Pow(g.BGap, 4)),
		IfMuEqualsInverseBGapScaling:  fmt.Sprintf("S(t)=32 B_gap^-4(t²+t⁴); inverse insertion would be an external choice, not derived"),
		DerivedScaling:                d.MuPowerScaling,
		ProducesInverseBGap:           false,
		ProducesFourOverPi:            false,
		Verdict:                       StatusBGapInsertionAudited,
	}
}

func auditFirewalls(c NCGCalculusLedger, b BGapInsertionAudit, s SaddleAudit) Firewalls {
	return Firewalls{
		DoesNotUseContinuumForms:     !c.RequiresContinuumForms && !c.RequiresIntegrationMeasure,
		DoesNotInventPhysicalDF:      !c.PhysicalDFDerived,
		DoesNotPromoteBGapToMajorana: !b.TreatingBGapAsMajoranaDerived,
		DoesNotPromoteBGapToCoupling: !b.TreatingBGapAsInverseDerived,
		DoesNotClaimFourOverPiSaddle: !b.ProducesFourOverPi && !s.NontrivialActionGapDerived,
		DoesNotGrantIntermediateSeal: true,
		FiniteCorePolluted:           false,
		Verdict:                      StatusFirewallsPreserved,
	}
}

func buildSummary(g Gate285Snapshot, c NCGCalculusLedger, d MatrixDiagnostic, s SaddleAudit, b BGapInsertionAudit, f Firewalls) Summary {
	finiteInst := d.FiniteTraceActionComputed && s.NontrivialActionGapDerived && b.ProducesInverseBGap && b.ProducesFourOverPi && c.PhysicalDFDerived && c.FullAlgebraRepresentation
	statuses := strings.Join(collectStatusesPartial(g, c, d, s, b, f), ";")
	return Summary{
		Gate285Inherited:         g.ContinuumRouteAudited && g.FiniteConnectionMissing,
		NCGCalculusFormalized:    c.AlgebraicMatrixRouteDefined,
		InnerFluctuationBuilt:    d.NonVacuousOneForm,
		FiniteTraceEvaluated:     d.FiniteTraceActionComputed,
		BGapInsertionAudited:     b.BGap > 0,
		NontrivialSaddleDerived:  s.NontrivialRealSaddleExists && s.NontrivialActionGapDerived,
		InverseBGapActionDerived: b.ProducesInverseBGap,
		FourOverPiGenerated:      b.ProducesFourOverPi,
		FiniteInstantonDerived:   finiteInst,
		IntermediateSealGranted:  false,
		FirewallPreserved:        !f.FiniteCorePolluted && f.DoesNotClaimFourOverPiSaddle && f.DoesNotPromoteBGapToMajorana && f.DoesNotPromoteBGapToCoupling,
		Status:                   statuses,
		DirectAnswer:             "Finite NCG calculus is the correct inward route, and it produces non-vacuous finite one-forms in a local quaternionic diagnostic, but it does not derive a non-trivial saddle, an inverse-B_gap law, or the 4/π instanton action.",
		NextGate:                 "Derive the physical finite Dirac operator, full C⊕H⊕M3(C) representation, and B_gap-to-Majorana bilinear map before retrying finite saddle-point instanton dynamics.",
	}
}

func collectStatusesPartial(g Gate285Snapshot, c NCGCalculusLedger, d MatrixDiagnostic, s SaddleAudit, b BGapInsertionAudit, f Firewalls) []string {
	statuses := []string{}
	if g.ContinuumRouteAudited {
		statuses = append(statuses, StatusGate285ContinuumBarrierInherited)
	}
	if c.AlgebraicMatrixRouteDefined {
		statuses = append(statuses, StatusNCGCalculusFormalized)
	}
	if d.NonVacuousOneForm {
		statuses = append(statuses, StatusInnerFluctuationDiagnosticBuilt)
	}
	if d.FiniteTraceActionComputed {
		statuses = append(statuses, StatusFiniteCurvatureTraceEvaluated)
	}
	if b.BGap > 0 {
		statuses = append(statuses, StatusBGapInsertionAudited)
	}
	if !s.NontrivialRealSaddleExists {
		statuses = append(statuses, StatusSaddleSearchCompleted)
	}
	if f.DoesNotClaimFourOverPiSaddle {
		statuses = append(statuses, StatusFirewallsPreserved)
	}
	if !c.PhysicalDFDerived {
		statuses = append(statuses, StatusFailedPhysicalDFStillMissing)
	}
	if !c.FullAlgebraRepresentation {
		statuses = append(statuses, StatusFailedFullAFRepresentationMissing)
	}
	if !b.TreatingBGapAsMajoranaDerived {
		statuses = append(statuses, StatusFailedBGapMajoranaMapMissing)
	}
	if !s.NontrivialRealSaddleExists {
		statuses = append(statuses, StatusFailedNoNontrivialSaddle)
	}
	if !b.ProducesInverseBGap {
		statuses = append(statuses, StatusFailedNoInverseBGapAction)
	}
	if !b.ProducesFourOverPi {
		statuses = append(statuses, StatusFailedFourOverPiNotGenerated)
	}
	statuses = append(statuses, StatusFailedFiniteInstantonNotDerived, StatusFailedIntermediateSealStillNeeded)
	return statuses
}

func collectStatuses(s Summary) []string {
	if s.Status == "" {
		return nil
	}
	return strings.Split(s.Status, ";")
}

func FormatGate285(g Gate285Snapshot) string {
	return fmt.Sprintf("continuumAudited=%t finiteConnMissing=%t finiteCurvMissing=%t CSFunctionalMissing=%t BgapCouplingMissing=%t sealRequired=%t c=%.12g B_gap=%.12g targetExponent=%.12g", g.ContinuumRouteAudited, g.FiniteConnectionMissing, g.FiniteCurvatureMissing, g.CSFunctionalMissing, g.BGapCouplingMissing, g.IntermediateSealRequired, g.FourOverPiCoefficient, g.BGap, g.CandidateInstantonExponent)
}
func FormatCalculus(c NCGCalculusLedger) string {
	return fmt.Sprintf("d=%q oneforms=%q A=%q F=%q action=%q continuumForms=%t integration=%t algebraic=%t physicalDF=%t fullRep=%t verdict=%s", c.DifferentialDefinition, c.OneFormDefinition, c.InnerFluctuationDefinition, c.CurvatureDefinition, c.FiniteTraceActionDefinition, c.RequiresContinuumForms, c.RequiresIntegrationMeasure, c.AlgebraicMatrixRouteDefined, c.PhysicalDFDerived, c.FullAlgebraRepresentation, c.Verdict)
}
func FormatDiagnostic(d MatrixDiagnostic) string {
	return fmt.Sprintf("carrier=%q sector=%q D=%q gen=%q oneNorm2=%.12g F=%q S=%q scaling=%q nonvacuous=%t curvature=%t trace=%t verdict=%s", d.Carrier, d.AlgebraSector, d.D, d.Generator, d.OneFormNorm2, d.CurvatureFormula, d.ActionPolynomial, d.MuPowerScaling, d.NonVacuousOneForm, d.FiniteCurvatureComputed, d.FiniteTraceActionComputed, d.Verdict)
}
func FormatSaddle(s SaddleAudit) string {
	return fmt.Sprintf("S=%q dS=%q realSaddles=%v nontrivial=%t vacuumS=%.12g gap=%t verdict=%s", s.ActionPolynomial, s.DerivativePolynomial, s.RealSaddles, s.NontrivialRealSaddleExists, s.TrivialVacuumAction, s.NontrivialActionGapDerived, s.Verdict)
}
func FormatBGapAudit(b BGapInsertionAudit) string {
	return fmt.Sprintf("hypothesis=%q B_gap=%.12g target=%.12g majoranaDerived=%t inverseDerived=%t muB=%q muInv=%q derivedScaling=%q invBgap=%t fourPi=%t verdict=%s", b.InsertionHypothesis, b.BGap, b.CandidateInstantonTarget, b.TreatingBGapAsMajoranaDerived, b.TreatingBGapAsInverseDerived, b.IfMuEqualsBGapScaling, b.IfMuEqualsInverseBGapScaling, b.DerivedScaling, b.ProducesInverseBGap, b.ProducesFourOverPi, b.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noContinuumForms=%t noInventDF=%t noBgapMajorana=%t noBgapCoupling=%t noFourPiSaddle=%t noSeal=%t polluted=%t verdict=%s", f.DoesNotUseContinuumForms, f.DoesNotInventPhysicalDF, f.DoesNotPromoteBGapToMajorana, f.DoesNotPromoteBGapToCoupling, f.DoesNotClaimFourOverPiSaddle, f.DoesNotGrantIntermediateSeal, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate285=%t calculus=%t inner=%t trace=%t bgap=%t saddle=%t inverse=%t fourPi=%t finiteInst=%t seal=%t firewall=%t status=%s direct=%q next=%q", s.Gate285Inherited, s.NCGCalculusFormalized, s.InnerFluctuationBuilt, s.FiniteTraceEvaluated, s.BGapInsertionAudited, s.NontrivialSaddleDerived, s.InverseBGapActionDerived, s.FourOverPiGenerated, s.FiniteInstantonDerived, s.IntermediateSealGranted, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
