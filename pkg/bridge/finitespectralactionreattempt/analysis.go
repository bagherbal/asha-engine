// Package finitespectralactionreattempt implements Gate 268:
// Finite Spectral Action Re-Attempt / Seeley-de Witt Coefficient Audit.
//
// Gate 267 closed the full flavor ledger and named the finite spectral action
// as the only lawful path to reopen amplitude dynamics. Gate 268 performs the
// next disciplined re-attempt. It audits whether the currently available
// scaffold—S_C, gamma, candidate J, C ⊕ M3(C), and the formal odd self-adjoint
// D_F family—is already enough to compute genuine Seeley-de Witt coefficients
// and a Higgs-to-gauge mass ratio independent of sealed Yukawa amplitudes.
//
// The answer is deliberately split. Raw finite spectral moments can be computed
// for any chosen dimensionless D_F ansatz, and the unit-incidence diagnostic is
// exactly evaluable. But those moments are not yet physical Seeley-de Witt
// coefficients: the canonical finite Dirac block, non-vacuous order-one
// calculus, gauge fluctuation/projection, cutoff-moment normalization, scalar
// fluctuation map, and subtraction scheme remain missing. Therefore no Higgs
// mass ratio is derived; the gate closes as a precise obstruction and a clean
// target for the next spectral theorem.
package finitespectralactionreattempt

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fullflavorledgerclosure"
)

const (
	AuditID = "GATE268-FINITE-SPECTRAL-ACTION-SEELEY-DE-WITT-COEFFICIENT-AUDIT"

	StatusGate267Inherited             = "CONDITIONAL_SUPPORT_GATE267_FLAVOR_LEDGER_CLOSURE_INHERITED"
	StatusScaffoldRetrieved            = "CONDITIONAL_SUPPORT_SPECTRAL_SCAFFOLD_RETRIEVED"
	StatusFormalDFFamilyAvailable      = "CONDITIONAL_SUPPORT_FORMAL_ODD_SELF_ADJOINT_DF_FAMILY_AVAILABLE"
	StatusRawMomentsEvaluated          = "CONDITIONAL_SUPPORT_RAW_FINITE_SPECTRAL_MOMENTS_EVALUATED"
	StatusAmplitudeDependenceExposed   = "CONDITIONAL_SUPPORT_DF_MOMENT_AMPLITUDE_DEPENDENCE_EXPOSED"
	StatusFailedCanonicalDF            = "FAILED_ROUTE_CANONICAL_FINITE_DIRAC_OPERATOR_NOT_DERIVED"
	StatusFailedSDW                    = "FAILED_ROUTE_SEELEY_DE_WITT_COEFFICIENTS_NOT_DERIVED"
	StatusFailedHiggsRatio             = "FAILED_ROUTE_HIGGS_MASS_RATIO_NOT_DERIVED"
	StatusFailedGaugeProjection        = "FAILED_ROUTE_GAUGE_KINETIC_PROJECTION_MISSING"
	StatusFailedScalarFluctuation      = "FAILED_ROUTE_SCALAR_FLUCTUATION_MAP_MISSING"
	StatusFailedCutoffNormalization    = "FAILED_ROUTE_CUTOFF_MOMENTS_AND_SUBTRACTION_SCHEME_MISSING"
	StatusEmpiricalYukawaSealPreserved = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate267Inheritance struct {
	FlavorLedgerClosed     bool
	KinematicsDerived      bool
	DynamicsSealed         bool
	EmpiricalYukawaSeal    bool
	FutureSpectralRequired bool
	NativeFlavorDynamics   bool
	FiniteCorePolluted     bool
	Verdict                string
}

type SpectralScaffoldAudit struct {
	CarrierName                 string
	ComplexFockDimension        int
	DoubledCarrierDimension     int
	GammaAvailable              bool
	GammaTraceZero              bool
	CandidateJAvailable         bool
	CandidateJIsPreflightOnly   bool
	FiniteAlgebraName           string
	NativeFiniteAlgebraRecorded bool
	OrderOneVerified            bool
	GaugeFluctuationMapDerived  bool
	ScalarFluctuationMapDerived bool
	Verdict                     string
}

type FiniteDiracAssemblyAudit struct {
	FamilyFormula               string
	Dimension                   int
	LeftDimension               int
	RightDimension              int
	FreeRealParameters          int
	OddWithGammaByConstruction  bool
	SelfAdjointByConstruction   bool
	UnitIncidenceRepresentative bool
	CanonicalBlockSelected      bool
	PhysicalChiralityDerived    bool
	JDCompatibilityDerived      bool
	OrderOneCalculusVerified    bool
	UsesObservedMasses          bool
	UsesYukawaAmplitudes        bool
	PromotablePhysicalDF        bool
	Verdict                     string
}

type MomentRow struct {
	Name                   string
	SingularValues         []float64
	Dimension              int
	TraceD0                float64
	TraceD2                float64
	TraceD4                float64
	RawA2OverA4            float64
	EffectiveParticipation float64
	UsesCanonicalDF        bool
	UsesObservedInput      bool
	Comment                string
}

type SpectralTraceAudit struct {
	Rows                           []MomentRow
	MomentsComputed                bool
	UnitRepresentativeComputed     bool
	DeformedRepresentativeComputed bool
	RawMomentRatioInvariant        bool
	DependsOnDFSingularValues      bool
	SeeleyDeWittMapDerived         bool
	CutoffMomentsDerived           bool
	NormalizationSchemeDerived     bool
	Verdict                        string
}

type HiggsMassRatioAudit struct {
	CandidateFormula               string
	RequiresA2Coefficient          bool
	RequiresA4GaugeCoefficient     bool
	RequiresScalarHessianMap       bool
	RequiresGaugeKineticProjection bool
	RequiresCutoffNormalization    bool
	RequiresCanonicalDF            bool
	IndependentOfYukawaAmplitudes  bool
	UnitDiagnosticRatio            float64
	DeformedDiagnosticRatio        float64
	DiagnosticStable               bool
	HiggsRatioDerived              bool
	HiggsMassPredicted             bool
	Verdict                        string
}

type FirewallAudit struct {
	EmpiricalYukawaSealPreserved    bool
	SpontaneousCarrierSealPreserved bool
	NoObservedMassInserted          bool
	NoVEVInserted                   bool
	NoCutoffScaleInserted           bool
	RawMomentsNotPromoted           bool
	NoHiggsPredictionClaim          bool
	NoGaugeCouplingPredictionClaim  bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type FutureObligation struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureTheoremMap struct {
	Obligations               []FutureObligation
	CanonicalDFRequired       bool
	HeatKernelRequired        bool
	GaugeProjectionRequired   bool
	ScalarFluctuationRequired bool
	SubtractionSchemeRequired bool
	ActionFunctionalRequired  bool
	CanDeriveHiggsRatioNow    bool
	RecommendedNextGate       string
	Verdict                   string
}

type Summary struct {
	Gate267Inherited        bool
	ScaffoldRetrieved       bool
	FormalDFFamilyAvailable bool
	RawMomentsEvaluated     bool
	MomentDependenceExposed bool
	SeeleyDeWittDerived     bool
	HiggsRatioDerived       bool
	FirewallPreserved       bool
	Status                  string
	NextGate                string
	Comment                 string
}

type Analysis struct {
	PreviousGate267 fullflavorledgerclosure.Analysis
	Inheritance     Gate267Inheritance
	Scaffold        SpectralScaffoldAudit
	Dirac           FiniteDiracAssemblyAudit
	Trace           SpectralTraceAudit
	Higgs           HiggsMassRatioAudit
	Firewall        FirewallAudit
	Future          FutureTheoremMap
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := fullflavorledgerclosure.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 267 predecessor: %w", err)
			return
		}
		inh := inheritGate267(prev)
		scaffold := retrieveScaffold()
		dirac := assembleFormalDirac(scaffold)
		trace := evaluateSpectralMoments(dirac)
		higgs := auditHiggsMassRatio(dirac, trace)
		firewall := buildFirewall(inh, dirac, trace, higgs)
		future := defineFutureMap(scaffold, dirac, trace, higgs)
		summary := summarize(inh, scaffold, dirac, trace, higgs, firewall, future)
		truth := buildTruth(inh, scaffold, dirac, trace, higgs, firewall, future)
		defaultA = Analysis{PreviousGate267: prev, Inheritance: inh, Scaffold: scaffold, Dirac: dirac, Trace: trace, Higgs: higgs, Firewall: firewall, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate267(prev fullflavorledgerclosure.Analysis) Gate267Inheritance {
	return Gate267Inheritance{
		FlavorLedgerClosed:     prev.Summary.FullFlavorLedgerClosed,
		KinematicsDerived:      prev.Firewall.KinematicsDerived,
		DynamicsSealed:         prev.Firewall.DynamicsSealed,
		EmpiricalYukawaSeal:    prev.Firewall.EmpiricalYukawaSealPreserved,
		FutureSpectralRequired: prev.FutureCriteria.RequiresFiniteSpectralAction && prev.FutureCriteria.RequiresCanonicalFiniteDirac && prev.FutureCriteria.RequiresHeatKernelCoefficients,
		NativeFlavorDynamics:   prev.Summary.NativeFlavorDynamicsDerived,
		FiniteCorePolluted:     prev.Firewall.FiniteCorePolluted,
		Verdict:                StatusGate267Inherited + "; Gate 267 named the spectral-action route as an obligation, not as an already-satisfied theorem",
	}
}

func retrieveScaffold() SpectralScaffoldAudit {
	return SpectralScaffoldAudit{
		CarrierName:                 "S_C = Λ*(C^4), represented by the 16-state complex Fock bookkeeping carrier",
		ComplexFockDimension:        16,
		DoubledCarrierDimension:     16,
		GammaAvailable:              true,
		GammaTraceZero:              true,
		CandidateJAvailable:         true,
		CandidateJIsPreflightOnly:   true,
		FiniteAlgebraName:           "C ⊕ M3(C)",
		NativeFiniteAlgebraRecorded: true,
		OrderOneVerified:            false,
		GaugeFluctuationMapDerived:  false,
		ScalarFluctuationMapDerived: false,
		Verdict:                     StatusScaffoldRetrieved + "; scaffold exists, but the order-one calculus and fluctuation maps are not yet physical spectral-action data",
	}
}

func assembleFormalDirac(scaffold SpectralScaffoldAudit) FiniteDiracAssemblyAudit {
	dim := scaffold.DoubledCarrierDimension
	half := dim / 2
	return FiniteDiracAssemblyAudit{
		FamilyFormula:               "D_F(M) = [[0,M],[M†,0]] on the gamma-even/gamma-odd split",
		Dimension:                   dim,
		LeftDimension:               half,
		RightDimension:              half,
		FreeRealParameters:          2 * half * half,
		OddWithGammaByConstruction:  scaffold.GammaAvailable,
		SelfAdjointByConstruction:   true,
		UnitIncidenceRepresentative: true,
		CanonicalBlockSelected:      false,
		PhysicalChiralityDerived:    false,
		JDCompatibilityDerived:      false,
		OrderOneCalculusVerified:    false,
		UsesObservedMasses:          false,
		UsesYukawaAmplitudes:        false,
		PromotablePhysicalDF:        false,
		Verdict:                     StatusFormalDFFamilyAvailable + "; only a dimensionless odd self-adjoint family is assembled, not a canonical physical finite Dirac operator",
	}
}

func evaluateSpectralMoments(dirac FiniteDiracAssemblyAudit) SpectralTraceAudit {
	unit := momentRow("unit-incidence diagnostic D_F", repeat(1, dirac.LeftDimension), false, false, "All singular values are 1. This is a legal diagnostic representative, not a canonical theorem-selected D_F.")
	deformed := momentRow("deformed diagnostic D_F", []float64{1, 1, 1, 1, 1, 1, 1, 2}, false, false, "One singular value is changed to expose amplitude dependence of Tr(D²)/Tr(D⁴).")
	rows := []MomentRow{unit, deformed}
	stable := approx(unit.RawA2OverA4, deformed.RawA2OverA4, 1e-12)
	return SpectralTraceAudit{
		Rows:                           rows,
		MomentsComputed:                true,
		UnitRepresentativeComputed:     true,
		DeformedRepresentativeComputed: true,
		RawMomentRatioInvariant:        stable,
		DependsOnDFSingularValues:      !stable,
		SeeleyDeWittMapDerived:         false,
		CutoffMomentsDerived:           false,
		NormalizationSchemeDerived:     false,
		Verdict:                        StatusRawMomentsEvaluated + "; " + StatusAmplitudeDependenceExposed + "; raw moments exist but change under legal D_F deformations and are not Seeley-de Witt coefficients without a heat-kernel/action map",
	}
}

func momentRow(name string, sigmas []float64, canonical bool, observed bool, comment string) MomentRow {
	var s2, s4 float64
	for _, s := range sigmas {
		s2 += s * s
		s4 += s * s * s * s
	}
	tr2 := 2 * s2
	tr4 := 2 * s4
	ratio := math.Inf(1)
	eff := 0.0
	if tr4 != 0 {
		ratio = tr2 / tr4
	}
	if s4 != 0 {
		eff = (s2 * s2) / s4
	}
	return MomentRow{Name: name, SingularValues: append([]float64(nil), sigmas...), Dimension: 2 * len(sigmas), TraceD0: float64(2 * len(sigmas)), TraceD2: tr2, TraceD4: tr4, RawA2OverA4: ratio, EffectiveParticipation: eff, UsesCanonicalDF: canonical, UsesObservedInput: observed, Comment: comment}
}

func auditHiggsMassRatio(dirac FiniteDiracAssemblyAudit, trace SpectralTraceAudit) HiggsMassRatioAudit {
	unit, deform := math.NaN(), math.NaN()
	if len(trace.Rows) >= 2 {
		unit = trace.Rows[0].RawA2OverA4
		deform = trace.Rows[1].RawA2OverA4
	}
	stable := approx(unit, deform, 1e-12)
	return HiggsMassRatioAudit{
		CandidateFormula:               "formal diagnostic only: m_H²/g² ∝ a₂/a₄ after gauge projection, scalar fluctuation map, cutoff moments, and subtraction scheme",
		RequiresA2Coefficient:          true,
		RequiresA4GaugeCoefficient:     true,
		RequiresScalarHessianMap:       true,
		RequiresGaugeKineticProjection: true,
		RequiresCutoffNormalization:    true,
		RequiresCanonicalDF:            true,
		IndependentOfYukawaAmplitudes:  false,
		UnitDiagnosticRatio:            unit,
		DeformedDiagnosticRatio:        deform,
		DiagnosticStable:               stable,
		HiggsRatioDerived:              false,
		HiggsMassPredicted:             false,
		Verdict:                        StatusFailedHiggsRatio + "; the computed ratio is a raw moment diagnostic, depends on the unselected D_F singular spectrum, and lacks the scalar/gauge projection needed for a Higgs mass theorem",
	}
}

func buildFirewall(inh Gate267Inheritance, dirac FiniteDiracAssemblyAudit, trace SpectralTraceAudit, higgs HiggsMassRatioAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalYukawaSealPreserved:    inh.EmpiricalYukawaSeal,
		SpontaneousCarrierSealPreserved: true,
		NoObservedMassInserted:          !dirac.UsesObservedMasses,
		NoVEVInserted:                   true,
		NoCutoffScaleInserted:           true,
		RawMomentsNotPromoted:           trace.MomentsComputed && !trace.SeeleyDeWittMapDerived && !higgs.HiggsRatioDerived,
		NoHiggsPredictionClaim:          !higgs.HiggsMassPredicted,
		NoGaugeCouplingPredictionClaim:  !higgs.HiggsRatioDerived,
		FiniteCorePolluted:              false,
		Verdict:                         StatusEmpiricalYukawaSealPreserved + "; Gate 268 computes diagnostics but refuses to promote them into physical mass or coupling predictions",
	}
}

func defineFutureMap(scaffold SpectralScaffoldAudit, dirac FiniteDiracAssemblyAudit, trace SpectralTraceAudit, higgs HiggsMassRatioAudit) FutureTheoremMap {
	obligations := []FutureObligation{
		{Name: "canonical finite D_F selector", Required: true, Satisfied: dirac.CanonicalBlockSelected, Detail: "Select M in D_F(M) from finite algebra, order-one calculus, or an action extremum; representative unit incidence is insufficient."},
		{Name: "physical chirality and J reality", Required: true, Satisfied: dirac.PhysicalChiralityDerived && dirac.JDCompatibilityDerived, Detail: "Candidate gamma/J preflight must become a physical KO/order-one spectral triple, not bookkeeping."},
		{Name: "non-vacuous order-one calculus", Required: true, Satisfied: scaffold.OrderOneVerified && dirac.OrderOneCalculusVerified, Detail: "The algebra representation must produce real one-forms and constrain D_F non-vacuously."},
		{Name: "finite heat-kernel / cutoff-moment map", Required: true, Satisfied: trace.SeeleyDeWittMapDerived && trace.CutoffMomentsDerived, Detail: "Raw Tr(D^0), Tr(D^2), Tr(D^4) must be related to a0,a2,a4 with cutoff-function moments."},
		{Name: "gauge kinetic projection for a4", Required: true, Satisfied: scaffold.GaugeFluctuationMapDerived && higgs.RequiresA4GaugeCoefficient && false, Detail: "A4 must project onto the derived gauge kinetic Hessian with normalization and subtraction scheme."},
		{Name: "scalar fluctuation/Higgs Hessian map for a2", Required: true, Satisfied: scaffold.ScalarFluctuationMapDerived && false, Detail: "A2 must couple to the Higgs/scalar fluctuation with a finite Hessian normalization independent of sealed Yukawa fits."},
		{Name: "scale and subtraction convention", Required: true, Satisfied: trace.NormalizationSchemeDerived, Detail: "A dimensionless ratio can be physical only after common normalization, scheme, and subtraction conventions are fixed."},
		{Name: "non-empirical prediction before comparison", Required: true, Satisfied: higgs.HiggsMassPredicted, Detail: "The engine must emit a mass/coupling ratio before observing or fitting m_H, v, or Yukawa data."},
	}
	return FutureTheoremMap{
		Obligations:               obligations,
		CanonicalDFRequired:       true,
		HeatKernelRequired:        true,
		GaugeProjectionRequired:   true,
		ScalarFluctuationRequired: true,
		SubtractionSchemeRequired: true,
		ActionFunctionalRequired:  true,
		CanDeriveHiggsRatioNow:    false,
		RecommendedNextGate:       "Gate 269 — Canonical Finite Dirac Selector / Order-One Spectral Triple Completion Audit",
		Verdict:                   "Gate 268 defines the exact missing theorem stack: canonical D_F first, then heat-kernel/gauge/scalar projection before any Higgs ratio claim",
	}
}

func summarize(inh Gate267Inheritance, scaffold SpectralScaffoldAudit, dirac FiniteDiracAssemblyAudit, trace SpectralTraceAudit, higgs HiggsMassRatioAudit, firewall FirewallAudit, future FutureTheoremMap) Summary {
	status := strings.Join([]string{
		StatusScaffoldRetrieved,
		StatusFormalDFFamilyAvailable,
		StatusRawMomentsEvaluated,
		StatusAmplitudeDependenceExposed,
		StatusFailedCanonicalDF,
		StatusFailedSDW,
		StatusFailedGaugeProjection,
		StatusFailedScalarFluctuation,
		StatusFailedCutoffNormalization,
		StatusFailedHiggsRatio,
	}, "; ")
	return Summary{
		Gate267Inherited:        inh.FlavorLedgerClosed && inh.FutureSpectralRequired,
		ScaffoldRetrieved:       scaffold.GammaAvailable && scaffold.CandidateJAvailable && scaffold.NativeFiniteAlgebraRecorded,
		FormalDFFamilyAvailable: dirac.SelfAdjointByConstruction && dirac.OddWithGammaByConstruction,
		RawMomentsEvaluated:     trace.MomentsComputed,
		MomentDependenceExposed: trace.DependsOnDFSingularValues,
		SeeleyDeWittDerived:     trace.SeeleyDeWittMapDerived,
		HiggsRatioDerived:       higgs.HiggsRatioDerived,
		FirewallPreserved:       firewall.EmpiricalYukawaSealPreserved && firewall.RawMomentsNotPromoted && !firewall.FiniteCorePolluted,
		Status:                  status,
		NextGate:                future.RecommendedNextGate,
		Comment:                 "Gate 268 re-attempts the finite spectral action and computes raw D_F moments, but the Higgs mass ratio remains blocked until a canonical physical D_F and spectral projection theorem exist.",
	}
}

func buildTruth(inh Gate267Inheritance, scaffold SpectralScaffoldAudit, dirac FiniteDiracAssemblyAudit, trace SpectralTraceAudit, higgs HiggsMassRatioAudit, firewall FirewallAudit, future FutureTheoremMap) string {
	return fmt.Sprintf("Gate 268 truth: scaffold=%t formalDF=%t rawMoments=%t momentDependence=%t seeleyDeWitt=%t higgsRatio=%t firewall=%t next=%s", scaffold.NativeFiniteAlgebraRecorded && scaffold.GammaAvailable && scaffold.CandidateJAvailable, dirac.SelfAdjointByConstruction && dirac.OddWithGammaByConstruction, trace.MomentsComputed, trace.DependsOnDFSingularValues, trace.SeeleyDeWittMapDerived, higgs.HiggsRatioDerived, firewall.RawMomentsNotPromoted && !firewall.FiniteCorePolluted, future.RecommendedNextGate)
}

func repeat(v float64, n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = v
	}
	return out
}

func approx(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func FormatInheritance(a Gate267Inheritance) string {
	return fmt.Sprintf("closed=%t kinematics=%t dynamicsSealed=%t empiricalSeal=%t spectralRequired=%t nativeDynamics=%t polluted=%t verdict=%q", a.FlavorLedgerClosed, a.KinematicsDerived, a.DynamicsSealed, a.EmpiricalYukawaSeal, a.FutureSpectralRequired, a.NativeFlavorDynamics, a.FiniteCorePolluted, a.Verdict)
}

func FormatScaffold(a SpectralScaffoldAudit) string {
	return fmt.Sprintf("carrier=%q dim=%d doubled=%d gamma=%t gammaTraceZero=%t J=%t Jpreflight=%t algebra=%q nativeAlgebra=%t orderOne=%t gaugeMap=%t scalarMap=%t verdict=%q", a.CarrierName, a.ComplexFockDimension, a.DoubledCarrierDimension, a.GammaAvailable, a.GammaTraceZero, a.CandidateJAvailable, a.CandidateJIsPreflightOnly, a.FiniteAlgebraName, a.NativeFiniteAlgebraRecorded, a.OrderOneVerified, a.GaugeFluctuationMapDerived, a.ScalarFluctuationMapDerived, a.Verdict)
}

func FormatDirac(a FiniteDiracAssemblyAudit) string {
	return fmt.Sprintf("formula=%q dim=%d left=%d right=%d params=%d odd=%t selfAdjoint=%t unitRep=%t canonical=%t chirality=%t JD=%t orderOne=%t obsMass=%t yukawa=%t promotable=%t verdict=%q", a.FamilyFormula, a.Dimension, a.LeftDimension, a.RightDimension, a.FreeRealParameters, a.OddWithGammaByConstruction, a.SelfAdjointByConstruction, a.UnitIncidenceRepresentative, a.CanonicalBlockSelected, a.PhysicalChiralityDerived, a.JDCompatibilityDerived, a.OrderOneCalculusVerified, a.UsesObservedMasses, a.UsesYukawaAmplitudes, a.PromotablePhysicalDF, a.Verdict)
}

func FormatMomentRow(r MomentRow) string {
	return fmt.Sprintf("%s dim=%d sigmas=%s TrD0=%.6g TrD2=%.6g TrD4=%.6g rawA2OverA4=%.6g eff=%.6g canonical=%t observed=%t comment=%q", r.Name, r.Dimension, formatFloats(r.SingularValues), r.TraceD0, r.TraceD2, r.TraceD4, r.RawA2OverA4, r.EffectiveParticipation, r.UsesCanonicalDF, r.UsesObservedInput, r.Comment)
}

func FormatTrace(a SpectralTraceAudit) string {
	rows := make([]string, 0, len(a.Rows))
	for _, r := range a.Rows {
		rows = append(rows, FormatMomentRow(r))
	}
	return fmt.Sprintf("moments=%t unit=%t deformed=%t invariant=%t depends=%t sdw=%t cutoff=%t norm=%t rows=[%s] verdict=%q", a.MomentsComputed, a.UnitRepresentativeComputed, a.DeformedRepresentativeComputed, a.RawMomentRatioInvariant, a.DependsOnDFSingularValues, a.SeeleyDeWittMapDerived, a.CutoffMomentsDerived, a.NormalizationSchemeDerived, strings.Join(rows, " | "), a.Verdict)
}

func FormatHiggs(a HiggsMassRatioAudit) string {
	return fmt.Sprintf("formula=%q requiresA2=%t requiresA4=%t scalar=%t gauge=%t cutoff=%t DF=%t independent=%t unit=%.6g deformed=%.6g stable=%t ratioDerived=%t massPredicted=%t verdict=%q", a.CandidateFormula, a.RequiresA2Coefficient, a.RequiresA4GaugeCoefficient, a.RequiresScalarHessianMap, a.RequiresGaugeKineticProjection, a.RequiresCutoffNormalization, a.RequiresCanonicalDF, a.IndependentOfYukawaAmplitudes, a.UnitDiagnosticRatio, a.DeformedDiagnosticRatio, a.DiagnosticStable, a.HiggsRatioDerived, a.HiggsMassPredicted, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("empiricalSeal=%t ssbSeal=%t noMass=%t noVEV=%t noCutoff=%t rawNotPromoted=%t noHiggs=%t noGauge=%t polluted=%t verdict=%q", a.EmpiricalYukawaSealPreserved, a.SpontaneousCarrierSealPreserved, a.NoObservedMassInserted, a.NoVEVInserted, a.NoCutoffScaleInserted, a.RawMomentsNotPromoted, a.NoHiggsPredictionClaim, a.NoGaugeCouplingPredictionClaim, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureTheoremMap) string {
	missing := make([]string, 0, len(a.Obligations))
	for _, o := range a.Obligations {
		if o.Required && !o.Satisfied {
			missing = append(missing, o.Name)
		}
	}
	return fmt.Sprintf("obligations=%d missing=[%s] canonicalDF=%t heat=%t gauge=%t scalar=%t subtraction=%t action=%t canDerive=%t next=%q verdict=%q", len(a.Obligations), strings.Join(missing, "; "), a.CanonicalDFRequired, a.HeatKernelRequired, a.GaugeProjectionRequired, a.ScalarFluctuationRequired, a.SubtractionSchemeRequired, a.ActionFunctionalRequired, a.CanDeriveHiggsRatioNow, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate267=%t scaffold=%t formalDF=%t raw=%t dependence=%t sdw=%t higgs=%t firewall=%t status=%q next=%q comment=%q", a.Gate267Inherited, a.ScaffoldRetrieved, a.FormalDFFamilyAvailable, a.RawMomentsEvaluated, a.MomentDependenceExposed, a.SeeleyDeWittDerived, a.HiggsRatioDerived, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}

func formatFloats(v []float64) string {
	cp := append([]float64(nil), v...)
	sort.Float64s(cp)
	parts := make([]string, len(cp))
	for i, x := range cp {
		parts[i] = fmt.Sprintf("%.6g", x)
	}
	return "[" + strings.Join(parts, ",") + "]"
}
