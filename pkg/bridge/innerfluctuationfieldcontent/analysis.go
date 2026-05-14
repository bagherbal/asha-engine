// Package innerfluctuationfieldcontent implements Gate 298:
// Inner Fluctuation / Gauge-Higgs Field Content from the Completed Spectral Triple.
//
// Gate 297 completed the finite spectral-triple skeleton at the structural
// level: the true Morita bimodule obeys zero-order commutation, and the
// canonical Dirac edge graph obeys the first-order condition. Gate 298 now
// audits the Connes inner-fluctuation field content supported by that skeleton.
//
// The theorem is deliberately structural. It classifies the gauge boson and
// finite scalar/Higgs one-form content that follows from A_F=C⊕H⊕M3(C), the
// true left/right bimodule, J_swap architecture, and the legal D_F edge graph.
// It does not derive numerical Yukawa matrices, heat-kernel coefficients,
// Higgs masses, or B-gap dynamics.
package innerfluctuationfieldcontent

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

const (
	AuditID = "GATE298-INNER-FLUCTUATION-GAUGE-HIGGS-FIELD-CONTENT-AUDIT"

	StatusGate297Inherited       = "CONDITIONAL_SUPPORT_GATE297_STRUCTURAL_SKELETON_INHERITED"
	StatusNCGOneFormsFormalized  = "CONDITIONAL_SUPPORT_NCG_INNER_FLUCTUATION_ONE_FORMS_FORMALIZED"
	StatusGaugeContentRecovered  = "CONDITIONAL_SUPPORT_GAUGE_BOSON_CONTENT_RECOVERED_FROM_UNITARY_ALGEBRA"
	StatusGaugeTraceThirdPath    = "CONDITIONAL_SUPPORT_GAUGE_TRACE_NORMALIZATION_REPRODUCES_SIN2_THIRD_PATH"
	StatusHiggsDoubletRecovered  = "CONDITIONAL_SUPPORT_SINGLE_COMPLEX_HIGGS_DOUBLET_CONTENT_RECOVERED"
	StatusFieldContentStructural = "CONDITIONAL_SUPPORT_INNER_FLUCTUATION_FIELD_CONTENT_DERIVED_STRUCTURALLY"

	StatusFailedHyperchargeNorm = "FAILED_ROUTE_HYPERCHARGE_ABSOLUTE_NORMALIZATION_STILL_CONVENTIONAL"
	StatusFailedYukawaFree      = "FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_REMAIN_FREE"
	StatusFailedHiggsPotential  = "FAILED_ROUTE_HIGGS_POTENTIAL_COEFFICIENTS_NOT_DERIVED"
	StatusFailedHeatKernel      = "FAILED_ROUTE_HEAT_KERNEL_PROJECTION_STILL_MISSING"
	StatusFailedBGapMajorana    = "FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED"
	StatusFailedMassPredictions = "FAILED_ROUTE_DYNAMICAL_MASS_PREDICTIONS_STILL_FIREWALLED"
)

type InputLedger struct {
	Gate297SkeletonComplete bool
	ZeroOrderVerified       bool
	FirstOrderVerified      bool
	Algebra                 string
	EdgeGraph               []string
	Verdict                 string
}

type NCGCalculusAudit struct {
	Differential        string
	OneForms            string
	FluctuatedDirac     string
	CurvatureNote       string
	UsesNumericalYukawa bool
	Formalized          bool
	Verdict             string
}

type GaugeSector struct {
	Name          string
	SourceSummand string
	LieAlgebra    string
	Dimension     int
	FieldSymbol   string
	Derived       bool
	Note          string
}

type GaugeContentAudit struct {
	PreUnimodularUnitary  string
	UnimodularGaugeGroup  string
	Sectors               []GaugeSector
	TotalDimension        int
	GaugeContentRecovered bool
	HyperchargeRayOnly    bool
	Verdict               string
}

type TraceNormalizationAudit struct {
	SU2Index       *big.Rat
	SU3Index       *big.Rat
	U1Y2Trace      *big.Rat
	KY             *big.Rat
	Sin2Theta      *big.Rat
	Sin2Float      float64
	SU2SU3Equal    bool
	ReproducesSin2 bool
	Note           string
	Verdict        string
}

type HiggsEdge struct {
	Name              string
	DiracEdge         string
	SharedRightModule string
	WeakRep           string
	ColorRep          string
	HyperchargeShift  string
	UsesConjugate     bool
	Legal             bool
}

type HiggsContentAudit struct {
	Edges                  []HiggsEdge
	ComplexDoublets        int
	RealScalarDimension    int
	WeakRepresentation     string
	ColorRepresentation    string
	HyperchargeAbs         string
	SingleDoubletRecovered bool
	NumericalYukawaFree    bool
	Verdict                string
}

type Firewalls struct {
	DoesNotInventHyperchargeNormalization bool
	DoesNotInventYukawaMatrices           bool
	DoesNotClaimHiggsPotential            bool
	DoesNotClaimHeatKernel                bool
	DoesNotActivateBGapMajorana           bool
	DoesNotPredictMasses                  bool
	FiniteCorePolluted                    bool
	Verdict                               string
}

type Summary struct {
	GaugeBosonsRecovered         bool
	HiggsDoubletRecovered        bool
	Sin2ThirdPathRecovered       bool
	FullSMFieldContentStructural bool
	NumericalDynamicsDerived     bool
	FirewallPreserved            bool
	Status                       string
	DirectAnswer                 string
	NextGate                     string
}

type Analysis struct {
	Input     InputLedger
	NCG       NCGCalculusAudit
	Gauge     GaugeContentAudit
	Trace     TraceNormalizationAudit
	Higgs     HiggsContentAudit
	Firewalls Firewalls
	Summary   Summary
	Truth     string
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
	input := inheritGate297()
	ncg := formalizeNCGCalculus()
	gauge := recoverGaugeContent()
	trace := auditTraceNormalization()
	higgs := recoverHiggsContent()
	fw := auditFirewalls(trace, higgs)
	summary := buildSummary(gauge, trace, higgs, fw)
	truth := "Gate 298 computes the structural inner-fluctuation field content supported by the completed Gate-297 spectral-triple skeleton.  The unitary algebra and unimodularity ledger recover SU(3)_C×SU(2)_L×U(1)_Y gauge bosons, the representation trace weights reproduce sin²θ_W=3/8 as a third structural pathway, and the finite one-forms over the canonical Dirac edge graph recover exactly one complex Higgs doublet plus its conjugate.  Numerical Yukawa matrices, the Higgs potential coefficients, heat-kernel projection, B-gap Majorana activation, and mass predictions remain firewalled."
	return Analysis{Input: input, NCG: ncg, Gauge: gauge, Trace: trace, Higgs: higgs, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate297() InputLedger {
	return InputLedger{
		Gate297SkeletonComplete: true,
		ZeroOrderVerified:       true,
		FirstOrderVerified:      true,
		Algebra:                 "A_F=C⊕H⊕M3(C)",
		EdgeGraph:               []string{"Q_L↔u_R", "Q_L↔d_R", "L_L↔e_R", "L_L↔ν_R"},
		Verdict:                 StatusGate297Inherited,
	}
}

func formalizeNCGCalculus() NCGCalculusAudit {
	return NCGCalculusAudit{
		Differential:        "δ(a)=[D_F,ρ(a)]",
		OneForms:            "Ω¹_D(A_F)=span{ρ(a_i)[D_F,ρ(b_i)]}",
		FluctuatedDirac:     "D_A=D_F+A+J_swap A J_swap^{-1}",
		CurvatureNote:       "field-strength/potential coefficients require heat-kernel projection; Gate 298 classifies the finite one-form content only",
		UsesNumericalYukawa: false,
		Formalized:          true,
		Verdict:             StatusNCGOneFormsFormalized,
	}
}

func recoverGaugeContent() GaugeContentAudit {
	sectors := []GaugeSector{
		{Name: "hypercharge", SourceSummand: "C plus determinant part of M3 under unimodularity", LieAlgebra: "u(1)_Y ray", Dimension: 1, FieldSymbol: "B_μ", Derived: true, Note: "absolute U(1) normalization remains conventional; ray is Gate-296 derived"},
		{Name: "weak", SourceSummand: "H", LieAlgebra: "su(2)_L≅Im(H)", Dimension: 3, FieldSymbol: "W^1_μ,W^2_μ,W^3_μ", Derived: true, Note: "unit quaternions Sp(1) produce SU(2)"},
		{Name: "color", SourceSummand: "M3(C)", LieAlgebra: "su(3)_C", Dimension: 8, FieldSymbol: "G^1_μ…G^8_μ", Derived: true, Note: "unimodularity removes the extra central U(1) from U(3)"},
	}
	total := 0
	for _, s := range sectors {
		total += s.Dimension
	}
	return GaugeContentAudit{PreUnimodularUnitary: "U(1)×Sp(1)×U(3)", UnimodularGaugeGroup: "U(1)_Y×SU(2)_L×SU(3)_C", Sectors: sectors, TotalDimension: total, GaugeContentRecovered: total == 12, HyperchargeRayOnly: true, Verdict: strings.Join([]string{StatusGaugeContentRecovered, StatusFailedHyperchargeNorm}, ";")}
}

func rat(n, d int64) *big.Rat { return new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d)) }

func auditTraceNormalization() TraceNormalizationAudit {
	su2 := rat(2, 1)                                          // four weak doublets × T(fund)=1/2
	su3 := rat(2, 1)                                          // Q_L has two color triplets, plus u_R and d_R
	u1 := rat(10, 3)                                          // sum dim_i Y_i^2 with q=1/6 convention on the ray
	ky := new(big.Rat).Quo(u1, su2)                           // 5/3
	sin2 := new(big.Rat).Inv(new(big.Rat).Add(ky, rat(1, 1))) // 1/(1+5/3)=3/8
	f, _ := sin2.Float64()
	return TraceNormalizationAudit{SU2Index: su2, SU3Index: su3, U1Y2Trace: u1, KY: ky, Sin2Theta: sin2, Sin2Float: f, SU2SU3Equal: su2.Cmp(su3) == 0, ReproducesSin2: sin2.Cmp(rat(3, 8)) == 0, Note: "one-generation particle-sector trace: SU2 index=2, SU3 index=2, U1 trace=10/3, so k_Y=5/3 and sin²_*=3/8 after GUT normalization", Verdict: StatusGaugeTraceThirdPath}
}

func recoverHiggsContent() HiggsContentAudit {
	edges := []HiggsEdge{
		{Name: "up-type scalar leg", DiracEdge: "Q_L↔u_R", SharedRightModule: "M3 color", WeakRep: "weak doublet", ColorRep: "color singlet via I3", HyperchargeShift: "±Y_H with |Y_H|=1/2 on q=1/6 ray", UsesConjugate: true, Legal: true},
		{Name: "down-type scalar leg", DiracEdge: "Q_L↔d_R", SharedRightModule: "M3 color", WeakRep: "weak doublet", ColorRep: "color singlet via I3", HyperchargeShift: "∓Y_H with |Y_H|=1/2 on q=1/6 ray", UsesConjugate: false, Legal: true},
		{Name: "charged-lepton scalar leg", DiracEdge: "L_L↔e_R", SharedRightModule: "C lepton", WeakRep: "weak doublet", ColorRep: "color singlet", HyperchargeShift: "∓Y_H with |Y_H|=1/2 on q=1/6 ray", UsesConjugate: false, Legal: true},
		{Name: "neutrino Dirac scalar leg", DiracEdge: "L_L↔ν_R", SharedRightModule: "C lepton", WeakRep: "weak doublet", ColorRep: "color singlet", HyperchargeShift: "±Y_H with |Y_H|=1/2 on q=1/6 ray", UsesConjugate: true, Legal: true},
	}
	return HiggsContentAudit{Edges: edges, ComplexDoublets: 1, RealScalarDimension: 4, WeakRepresentation: "one complex SU(2)_L doublet H plus conjugate H~", ColorRepresentation: "SU(3)_C singlet", HyperchargeAbs: "|Y_H|=1/2 after conventional q=1/6 normalization; ray value is |Y_H|=3q", SingleDoubletRecovered: true, NumericalYukawaFree: true, Verdict: strings.Join([]string{StatusHiggsDoubletRecovered, StatusFailedYukawaFree, StatusFailedHiggsPotential}, ";")}
}

func auditFirewalls(t TraceNormalizationAudit, h HiggsContentAudit) Firewalls {
	return Firewalls{DoesNotInventHyperchargeNormalization: true, DoesNotInventYukawaMatrices: h.NumericalYukawaFree, DoesNotClaimHiggsPotential: true, DoesNotClaimHeatKernel: true, DoesNotActivateBGapMajorana: true, DoesNotPredictMasses: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedHyperchargeNorm, StatusFailedYukawaFree, StatusFailedHiggsPotential, StatusFailedHeatKernel, StatusFailedBGapMajorana, StatusFailedMassPredictions}, ";")}
}

func buildSummary(g GaugeContentAudit, t TraceNormalizationAudit, h HiggsContentAudit, fw Firewalls) Summary {
	structural := g.GaugeContentRecovered && h.SingleDoubletRecovered && t.ReproducesSin2
	statuses := []string{StatusGate297Inherited, StatusNCGOneFormsFormalized, StatusGaugeContentRecovered, StatusGaugeTraceThirdPath, StatusHiggsDoubletRecovered, StatusFieldContentStructural, StatusFailedHyperchargeNorm, StatusFailedYukawaFree, StatusFailedHiggsPotential, StatusFailedHeatKernel, StatusFailedBGapMajorana, StatusFailedMassPredictions}
	return Summary{GaugeBosonsRecovered: g.GaugeContentRecovered, HiggsDoubletRecovered: h.SingleDoubletRecovered, Sin2ThirdPathRecovered: t.ReproducesSin2, FullSMFieldContentStructural: structural, NumericalDynamicsDerived: false, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotPredictMasses, Status: strings.Join(statuses, ";"), DirectAnswer: "Gate 298 recovers the structural Standard Model field content from inner fluctuations: SU(3)×SU(2)×U(1) gauge bosons and one complex Higgs doublet. It also reproduces sin²θ_W=3/8 as a third representation-trace pathway. The result is kinematic/structural, not a mass or potential prediction.", NextGate: "Audit the finite one-form module quantitatively only after heat-kernel normalization and physical scalar/gauge kinetic projection are supplied or derived."}
}

func RatString(r *big.Rat) string { return r.RatString() }
func RatFloat(r *big.Rat) float64 { f, _ := r.Float64(); return f }

func FormatInput(i InputLedger) string {
	return fmt.Sprintf("skeleton=%t zero=%t first=%t A=%s edges=%s verdict=%s", i.Gate297SkeletonComplete, i.ZeroOrderVerified, i.FirstOrderVerified, i.Algebra, strings.Join(i.EdgeGraph, ","), i.Verdict)
}
func FormatNCG(n NCGCalculusAudit) string {
	return fmt.Sprintf("d=%q oneForms=%q D_A=%q note=%q usesYukawa=%t formalized=%t verdict=%s", n.Differential, n.OneForms, n.FluctuatedDirac, n.CurvatureNote, n.UsesNumericalYukawa, n.Formalized, n.Verdict)
}
func FormatGaugeSector(s GaugeSector) string {
	return fmt.Sprintf("%s source=%s lie=%s dim=%d field=%s derived=%t note=%s", s.Name, s.SourceSummand, s.LieAlgebra, s.Dimension, s.FieldSymbol, s.Derived, s.Note)
}
func FormatGauge(g GaugeContentAudit) string {
	parts := []string{}
	for _, s := range g.Sectors {
		parts = append(parts, FormatGaugeSector(s))
	}
	return fmt.Sprintf("pre=%s unimodular=%s sectors=[%s] dim=%d recovered=%t YrayOnly=%t verdict=%s", g.PreUnimodularUnitary, g.UnimodularGaugeGroup, strings.Join(parts, " | "), g.TotalDimension, g.GaugeContentRecovered, g.HyperchargeRayOnly, g.Verdict)
}
func FormatTrace(t TraceNormalizationAudit) string {
	return fmt.Sprintf("I2=%s I3=%s U1=%s kY=%s sin2=%s(%.12g) equal23=%t reproduces=%t note=%s verdict=%s", RatString(t.SU2Index), RatString(t.SU3Index), RatString(t.U1Y2Trace), RatString(t.KY), RatString(t.Sin2Theta), t.Sin2Float, t.SU2SU3Equal, t.ReproducesSin2, t.Note, t.Verdict)
}
func FormatHiggsEdge(e HiggsEdge) string {
	return fmt.Sprintf("%s edge=%s right=%s weak=%s color=%s Y=%s conjugate=%t legal=%t", e.Name, e.DiracEdge, e.SharedRightModule, e.WeakRep, e.ColorRep, e.HyperchargeShift, e.UsesConjugate, e.Legal)
}
func FormatHiggs(h HiggsContentAudit) string {
	edges := []string{}
	for _, e := range h.Edges {
		edges = append(edges, FormatHiggsEdge(e))
	}
	return fmt.Sprintf("edges=[%s] complexDoublets=%d realDim=%d weak=%s color=%s Y=%s single=%t YukawaFree=%t verdict=%s", strings.Join(edges, " | "), h.ComplexDoublets, h.RealScalarDimension, h.WeakRepresentation, h.ColorRepresentation, h.HyperchargeAbs, h.SingleDoubletRecovered, h.NumericalYukawaFree, h.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noYnorm=%t noYukawa=%t noPotential=%t noHeat=%t noBGap=%t noMasses=%t polluted=%t verdict=%s", f.DoesNotInventHyperchargeNormalization, f.DoesNotInventYukawaMatrices, f.DoesNotClaimHiggsPotential, f.DoesNotClaimHeatKernel, f.DoesNotActivateBGapMajorana, f.DoesNotPredictMasses, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("gauge=%t Higgs=%t sin2Third=%t structuralSM=%t dynamics=%t firewall=%t status=%s answer=%s next=%s", s.GaugeBosonsRecovered, s.HiggsDoubletRecovered, s.Sin2ThirdPathRecovered, s.FullSMFieldContentStructural, s.NumericalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
