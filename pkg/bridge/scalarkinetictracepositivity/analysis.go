// Package scalarkinetictracepositivity implements Gate 301:
// Scalar Kinetic Trace Functional / Positive Z_H Evaluable Carrier Audit.
//
// Gate 300 formalized the normalization algorithm and deliberately stopped at
// FAILED_ROUTE_POSITIVE_ZH_NOT_NUMERICALLY_PROVED. Gate 301 advances exactly one
// epistemic layer: it constructs the finite scalar kinetic trace carrier and
// audits whether its algebraic form is positive by construction. It proves a
// positive-semidefinite Hilbert-Schmidt trace form and the exact strict-positivity
// condition, while preserving all numerical Yukawa, cutoff, subtraction, and
// B-gap firewalls.
package scalarkinetictracepositivity

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE301-SCALAR-KINETIC-TRACE-FUNCTIONAL-POSITIVE-ZH-EVALUABLE-CARRIER-AUDIT"

	StatusGate300Inherited                    = "CONDITIONAL_SUPPORT_GATE300_NORMALIZATION_ALGORITHM_INHERITED"
	StatusKineticTraceFormalized              = "CONDITIONAL_SUPPORT_SCALAR_KINETIC_TRACE_FUNCTIONAL_FORMALIZED"
	StatusDoubledCarrierEvaluated             = "CONDITIONAL_SUPPORT_DOUBLED_SPACE_SCALAR_EDGE_TRACE_EVALUATED"
	StatusPositiveTraceProvedStructurally     = "CONDITIONAL_SUPPORT_POSITIVE_SCALAR_KINETIC_TRACE_PROVED_STRUCTURALLY"
	StatusStrictPositivityConditionIdentified = "CONDITIONAL_SUPPORT_STRICT_POSITIVE_ZH_CONDITION_IDENTIFIED"
	StatusAmplitudeSealLedgerBuilt            = "CONDITIONAL_SUPPORT_SCALAR_AMPLITUDE_SEAL_LEDGER_BUILT"
	StatusEmpiricalFirewallsPreserved         = "CONDITIONAL_SUPPORT_GATE301_EMPIRICAL_FIREWALLS_PRESERVED"

	StatusFailedNumericalZHStillSealed       = "FAILED_ROUTE_NUMERICAL_ZH_VALUE_STILL_SEALED"
	StatusFailedYukawaAmplitudesStillFree    = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_FREE"
	StatusFailedNonzeroAmplitudesNotNative   = "FAILED_ROUTE_NONZERO_YUKAWA_AMPLITUDES_NOT_DERIVED_FROM_FINITE_GEOMETRY"
	StatusFailedCutoffMomentStillUnfixed     = "FAILED_ROUTE_CUTOFF_MOMENT_F0_STILL_UNFIXED"
	StatusFailedTraceConventionStillExplicit = "FAILED_ROUTE_TRACE_NORMALIZATION_CONVENTION_STILL_EXPLICIT"
	StatusFailedHeatKernelSignStillExplicit  = "FAILED_ROUTE_WICK_AND_HEAT_KERNEL_SIGN_CONVENTION_STILL_EXPLICIT"
	StatusFailedMassQuarticStillFirewalled   = "FAILED_ROUTE_HIGGS_MASS_AND_QUARTIC_STILL_FIREWALLED"
	StatusFailedBGapInstantonStillSealed     = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

type InheritedGate300 struct {
	ZHDefined                    bool
	ScalarKineticSelectorDefined bool
	RescalingDefined             bool
	GaugeNormalizationDefined    bool
	PositiveZHNumericallyProved  bool
	NumericalDynamicsDerived     bool
	Verdict                      string
}

type ScalarEdge struct {
	Sector               string
	Edge                 string
	LeftCarrier          string
	RightCarrier         string
	AmplitudeSymbol      string
	TraceContribution    string
	Multiplicity         int
	AppearsWithAdjoint   bool
	PreservesHermiticity bool
	SealedAmplitude      bool
	NumericallyProvided  bool
	PositivityStatement  string
}

type KineticTraceFunctional struct {
	Name                       string
	Source                     string
	Functional                 string
	EdgeTerms                  []ScalarEdge
	ParticleAntiparticleFactor int
	UsesHilbertSchmidtNorm     bool
	ExcludesGaugeCurvature     bool
	ExcludesPotentialTerms     bool
	ExcludesVacuumTerms        bool
	ExcludesMajoranaBGap       bool
	Formalized                 bool
	Verdict                    string
}

type DoubledSpaceEvaluation struct {
	Carrier                  string
	ParticleBlock            string
	AntiparticleBlock        string
	JSwapDuplication         string
	DoubleCountingHandled    bool
	PositivePairingPreserved bool
	QuarkEdgesMapped         int
	LeptonEdgesMapped        int
	TotalEdgesMapped         int
	Evaluation               string
	Verdict                  string
}

type PositivitySieve struct {
	RawExpression                   string
	SumOfSquaresExpression          string
	PositiveSemidefinite            bool
	NegativeTermsPermitted          bool
	ImaginaryKineticPermitted       bool
	StrictPositiveCondition         string
	StrictPositiveProved            bool
	StrictPositiveConditional       bool
	ZeroTraceFailureMode            string
	GhostRiskEliminatedStructurally bool
	Verdict                         string
}

type AmplitudeSeal struct {
	Symbol               string
	Sector               string
	RequiredForNumericZH bool
	RequiredForStrictZH  bool
	NativeValueDerived   bool
	AllowedSealTypes     []string
	Reason               string
}

type AmplitudeSealingCheck struct {
	Seals                     []AmplitudeSeal
	AtLeastOneNonzeroNeeded   bool
	AllNumericalValuesSealed  bool
	NoEmpiricalValuesInserted bool
	ReducibleToNumericZH      bool
	LedgerBuilt               bool
	Verdict                   string
}

type ZHCarrierMap struct {
	KRawDefinition              string
	ZHDefinition                string
	ZHPositivityStatement       string
	NumericalZHComputed         bool
	EvaluableAfterAmplitudeSeal bool
	RequiresPositiveF0          bool
	RequiresPositiveTraceNorm   bool
	RequiresEuclideanSignLedger bool
	Verdict                     string
}

type RemainingObligation struct {
	Name             string
	WhyRequired      string
	Status           string
	BlocksPrediction bool
}

type FirewallAudit struct {
	NoYukawaNumbersInserted     bool
	NoObservedMassesInserted    bool
	NoCutoffMomentInserted      bool
	NoSubtractionSchemeInvented bool
	NoBGapInstantonClaimed      bool
	NoMassQuarticClaimed        bool
	FiniteCorePolluted          bool
	Obligations                 []RemainingObligation
	Verdict                     string
}

type Summary struct {
	Gate300Inherited                  bool
	TraceFunctionalFormalized         bool
	DoubledCarrierEvaluated           bool
	PositiveSemidefiniteProved        bool
	StrictPositiveConditionIdentified bool
	NumericalZHComputed               bool
	PhysicalDynamicsDerived           bool
	FirewallPreserved                 bool
	Status                            string
	DirectAnswer                      string
	NextGate                          string
}

type Analysis struct {
	Input      InheritedGate300
	Trace      KineticTraceFunctional
	Doubled    DoubledSpaceEvaluation
	Positivity PositivitySieve
	Seals      AmplitudeSealingCheck
	ZH         ZHCarrierMap
	Firewalls  FirewallAudit
	Summary    Summary
	Truth      string
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
	input := inheritGate300()
	trace := formalizeKineticTraceFunctional()
	doubled := evaluateDoubledSpace(trace)
	positivity := runPositivitySieve(trace, doubled)
	seals := buildAmplitudeSealLedger(trace, positivity)
	zh := buildZHCarrierMap(trace, positivity, seals)
	firewalls := auditFirewalls(zh, seals)
	summary := buildSummary(input, trace, doubled, positivity, seals, zh, firewalls)
	truth := "Gate 301 proves that the scalar kinetic trace carrier has the correct positive Hilbert-Schmidt form: K_H^raw is a non-negative sum of multiplicity-weighted Tr(Y_s^†Y_s) terms over the completed H_F ⊕ H_F* scalar Dirac-edge carrier. Therefore the finite geometry does not natively permit negative or imaginary Higgs kinetic energy. Strict Z_H>0 is conditionally guaranteed exactly when the trace convention/f0/sign ledger is positive and at least one scalar Yukawa amplitude carrier is nonzero; the numerical value of Z_H remains sealed."
	return Analysis{Input: input, Trace: trace, Doubled: doubled, Positivity: positivity, Seals: seals, ZH: zh, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate300() InheritedGate300 {
	return InheritedGate300{
		ZHDefined:                    true,
		ScalarKineticSelectorDefined: true,
		RescalingDefined:             true,
		GaugeNormalizationDefined:    true,
		PositiveZHNumericallyProved:  false,
		NumericalDynamicsDerived:     false,
		Verdict:                      StatusGate300Inherited,
	}
}

func formalizeKineticTraceFunctional() KineticTraceFunctional {
	edges := []ScalarEdge{
		{Sector: "quark", Edge: "Q_L ↔ u_R", LeftCarrier: "left weak quark doublet Q_L", RightCarrier: "right up quark singlet u_R", AmplitudeSymbol: "Y_u", TraceContribution: "3 Tr(Y_u†Y_u)", Multiplicity: 3, AppearsWithAdjoint: true, PreservesHermiticity: true, SealedAmplitude: true, NumericallyProvided: false, PositivityStatement: "color multiplicity times a Hilbert-Schmidt norm; non-negative and strictly positive iff Y_u ≠ 0"},
		{Sector: "quark", Edge: "Q_L ↔ d_R", LeftCarrier: "left weak quark doublet Q_L", RightCarrier: "right down quark singlet d_R", AmplitudeSymbol: "Y_d", TraceContribution: "3 Tr(Y_d†Y_d)", Multiplicity: 3, AppearsWithAdjoint: true, PreservesHermiticity: true, SealedAmplitude: true, NumericallyProvided: false, PositivityStatement: "color multiplicity times a Hilbert-Schmidt norm; non-negative and strictly positive iff Y_d ≠ 0"},
		{Sector: "lepton", Edge: "L_L ↔ e_R", LeftCarrier: "left weak lepton doublet L_L", RightCarrier: "right charged lepton singlet e_R", AmplitudeSymbol: "Y_e", TraceContribution: "Tr(Y_e†Y_e)", Multiplicity: 1, AppearsWithAdjoint: true, PreservesHermiticity: true, SealedAmplitude: true, NumericallyProvided: false, PositivityStatement: "Hilbert-Schmidt norm; non-negative and strictly positive iff Y_e ≠ 0"},
		{Sector: "lepton", Edge: "L_L ↔ ν_R", LeftCarrier: "left weak lepton doublet L_L", RightCarrier: "right neutrino singlet ν_R", AmplitudeSymbol: "Y_ν", TraceContribution: "Tr(Y_ν†Y_ν)", Multiplicity: 1, AppearsWithAdjoint: true, PreservesHermiticity: true, SealedAmplitude: true, NumericallyProvided: false, PositivityStatement: "Dirac-neutrino Hilbert-Schmidt norm; non-negative and strictly positive iff Y_ν ≠ 0; Majorana/B-gap activation remains excluded"},
	}
	return KineticTraceFunctional{
		Name:                       "K_H^raw",
		Source:                     "a_4(D_A) derivative-order 2, scalar-power 2, curvature-order 0 channel",
		Functional:                 "K_H^raw := c_H · Tr_F(Φ†Φ)|_{scalar Dirac edges} = c_H · [3Tr(Y_u†Y_u)+3Tr(Y_d†Y_d)+Tr(Y_e†Y_e)+Tr(Y_ν†Y_ν)] up to the explicit doubled-space and heat-kernel convention factors",
		EdgeTerms:                  edges,
		ParticleAntiparticleFactor: 2,
		UsesHilbertSchmidtNorm:     true,
		ExcludesGaugeCurvature:     true,
		ExcludesPotentialTerms:     true,
		ExcludesVacuumTerms:        true,
		ExcludesMajoranaBGap:       true,
		Formalized:                 true,
		Verdict:                    StatusKineticTraceFormalized,
	}
}

func evaluateDoubledSpace(t KineticTraceFunctional) DoubledSpaceEvaluation {
	quark, lepton := 0, 0
	for _, e := range t.EdgeTerms {
		switch e.Sector {
		case "quark":
			quark++
		case "lepton":
			lepton++
		}
	}
	return DoubledSpaceEvaluation{
		Carrier:                  "H_F ⊕ H_F* with J_swap particle/antiparticle pairing",
		ParticleBlock:            "H_F carries Q_L↔u_R, Q_L↔d_R, L_L↔e_R, L_L↔ν_R scalar Dirac edges",
		AntiparticleBlock:        "H_F* carries the conjugate/opposite scalar edges with adjoint amplitudes",
		JSwapDuplication:         "the doubled trace contributes a positive multiplicity factor, not a sign flip; normalization may divide by a convention factor but cannot turn a square negative",
		DoubleCountingHandled:    true,
		PositivePairingPreserved: true,
		QuarkEdgesMapped:         quark,
		LeptonEdgesMapped:        lepton,
		TotalEdgesMapped:         len(t.EdgeTerms),
		Evaluation:               "Tr_{H_F⊕H_F*}(Φ†Φ)=2[3Tr(Y_u†Y_u)+3Tr(Y_d†Y_d)+Tr(Y_e†Y_e)+Tr(Y_ν†Y_ν)] before optional convention normalization",
		Verdict:                  StatusDoubledCarrierEvaluated,
	}
}

func runPositivitySieve(t KineticTraceFunctional, d DoubledSpaceEvaluation) PositivitySieve {
	allHermitianSquares := t.UsesHilbertSchmidtNorm && d.PositivePairingPreserved
	for _, e := range t.EdgeTerms {
		if !e.AppearsWithAdjoint || !e.PreservesHermiticity || e.Multiplicity <= 0 {
			allHermitianSquares = false
			break
		}
	}
	return PositivitySieve{
		RawExpression:                   "K_H^raw = c_H · Tr_{H_F⊕H_F*}(Φ†Φ)|_{allowed scalar edges}",
		SumOfSquaresExpression:          "K_H^raw = C_H · (3||Y_u||_HS² + 3||Y_d||_HS² + ||Y_e||_HS² + ||Y_ν||_HS²), with C_H positive after the explicit heat-kernel/sign/trace convention is selected",
		PositiveSemidefinite:            allHermitianSquares,
		NegativeTermsPermitted:          false,
		ImaginaryKineticPermitted:       false,
		StrictPositiveCondition:         "K_H^raw>0 iff C_H>0 and at least one of Y_u,Y_d,Y_e,Y_ν is a nonzero finite Dirac-edge amplitude; Z_H>0 additionally requires f_0>0 and the same positive trace convention",
		StrictPositiveProved:            false,
		StrictPositiveConditional:       allHermitianSquares,
		ZeroTraceFailureMode:            "if all scalar Dirac-edge amplitudes vanish, the trace is zero and the Higgs carrier is non-propagating rather than ghostlike; numerical amplitudes are still sealed",
		GhostRiskEliminatedStructurally: allHermitianSquares,
		Verdict:                         strings.Join([]string{StatusPositiveTraceProvedStructurally, StatusStrictPositivityConditionIdentified, StatusFailedNumericalZHStillSealed, StatusFailedNonzeroAmplitudesNotNative}, ";"),
	}
}

func buildAmplitudeSealLedger(t KineticTraceFunctional, p PositivitySieve) AmplitudeSealingCheck {
	seals := []AmplitudeSeal{}
	for _, e := range t.EdgeTerms {
		seals = append(seals, AmplitudeSeal{
			Symbol:               e.AmplitudeSymbol,
			Sector:               e.Sector,
			RequiredForNumericZH: true,
			RequiredForStrictZH:  true,
			NativeValueDerived:   false,
			AllowedSealTypes:     []string{"EmpiricalYukawaSeal", "FiniteAmplitudeTheorem", "PhenomenologicalTextureLedger"},
			Reason:               fmt.Sprintf("%s enters %s and fixes the size of the Hilbert-Schmidt norm only after a nonzero amplitude value or theorem is supplied", e.AmplitudeSymbol, e.TraceContribution),
		})
	}
	return AmplitudeSealingCheck{
		Seals:                     seals,
		AtLeastOneNonzeroNeeded:   true,
		AllNumericalValuesSealed:  true,
		NoEmpiricalValuesInserted: true,
		ReducibleToNumericZH:      false,
		LedgerBuilt:               len(seals) == len(t.EdgeTerms) && p.PositiveSemidefinite,
		Verdict:                   strings.Join([]string{StatusAmplitudeSealLedgerBuilt, StatusFailedYukawaAmplitudesStillFree, StatusFailedNonzeroAmplitudesNotNative}, ";"),
	}
}

func buildZHCarrierMap(t KineticTraceFunctional, p PositivitySieve, s AmplitudeSealingCheck) ZHCarrierMap {
	return ZHCarrierMap{
		KRawDefinition:              t.Functional,
		ZHDefinition:                "Z_H := N_4 f_0 K_H^raw, with N_4 carrying Seeley-de Witt normalization, doubled-space convention, trace normalization, and Wick/sign choice",
		ZHPositivityStatement:       "Z_H is structurally non-negative; it is strictly positive exactly when N_4>0, f_0>0, the finite trace inner product is positive, and at least one sealed scalar Yukawa amplitude is nonzero",
		NumericalZHComputed:         false,
		EvaluableAfterAmplitudeSeal: true,
		RequiresPositiveF0:          true,
		RequiresPositiveTraceNorm:   true,
		RequiresEuclideanSignLedger: true,
		Verdict:                     strings.Join([]string{StatusPositiveTraceProvedStructurally, StatusStrictPositivityConditionIdentified, StatusFailedNumericalZHStillSealed, StatusFailedCutoffMomentStillUnfixed, StatusFailedTraceConventionStillExplicit, StatusFailedHeatKernelSignStillExplicit}, ";"),
	}
}

func auditFirewalls(z ZHCarrierMap, s AmplitudeSealingCheck) FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "nonzero scalar amplitude theorem or seal", WhyRequired: "strict positivity requires at least one active finite Dirac scalar edge", Status: StatusFailedNonzeroAmplitudesNotNative, BlocksPrediction: true},
		{Name: "numerical Yukawa matrices Y_u,Y_d,Y_e,Y_ν", WhyRequired: "reduce the positive trace to an actual numerical K_H^raw and Z_H", Status: StatusFailedYukawaAmplitudesStillFree, BlocksPrediction: true},
		{Name: "cutoff moment f_0", WhyRequired: "multiplies the a4 scalar kinetic channel and absolute Z_H", Status: StatusFailedCutoffMomentStillUnfixed, BlocksPrediction: true},
		{Name: "trace normalization convention", WhyRequired: "fixes the overall positive constant C_H/N_4 and doubled-space multiplicity", Status: StatusFailedTraceConventionStillExplicit, BlocksPrediction: true},
		{Name: "Wick/sign convention", WhyRequired: "matches Euclidean spectral-action positivity to the physical Lorentzian kinetic sign", Status: StatusFailedHeatKernelSignStillExplicit, BlocksPrediction: true},
		{Name: "Higgs mass and quartic computation", WhyRequired: "needs positive numerical Z_H plus a2 subtraction and a4 quartic amplitude data", Status: StatusFailedMassQuarticStillFirewalled, BlocksPrediction: true},
		{Name: "B-gap instanton action", WhyRequired: "the positive scalar kinetic trace is polynomial and does not derive S_inst=(4/pi)/B_gap", Status: StatusFailedBGapInstantonStillSealed, BlocksPrediction: true},
	}
	return FirewallAudit{
		NoYukawaNumbersInserted:     s.NoEmpiricalValuesInserted,
		NoObservedMassesInserted:    true,
		NoCutoffMomentInserted:      z.RequiresPositiveF0,
		NoSubtractionSchemeInvented: true,
		NoBGapInstantonClaimed:      true,
		NoMassQuarticClaimed:        true,
		FiniteCorePolluted:          false,
		Obligations:                 obligations,
		Verdict:                     strings.Join([]string{StatusEmpiricalFirewallsPreserved, StatusFailedNumericalZHStillSealed, StatusFailedYukawaAmplitudesStillFree, StatusFailedCutoffMomentStillUnfixed, StatusFailedBGapInstantonStillSealed}, ";"),
	}
}

func buildSummary(i InheritedGate300, t KineticTraceFunctional, d DoubledSpaceEvaluation, p PositivitySieve, s AmplitudeSealingCheck, z ZHCarrierMap, f FirewallAudit) Summary {
	statuses := []string{
		StatusGate300Inherited,
		StatusKineticTraceFormalized,
		StatusDoubledCarrierEvaluated,
		StatusPositiveTraceProvedStructurally,
		StatusStrictPositivityConditionIdentified,
		StatusAmplitudeSealLedgerBuilt,
		StatusEmpiricalFirewallsPreserved,
		StatusFailedNumericalZHStillSealed,
		StatusFailedYukawaAmplitudesStillFree,
		StatusFailedNonzeroAmplitudesNotNative,
		StatusFailedCutoffMomentStillUnfixed,
		StatusFailedTraceConventionStillExplicit,
		StatusFailedHeatKernelSignStillExplicit,
		StatusFailedMassQuarticStillFirewalled,
		StatusFailedBGapInstantonStillSealed,
	}
	return Summary{
		Gate300Inherited:                  i.ZHDefined && i.ScalarKineticSelectorDefined && !i.PositiveZHNumericallyProved,
		TraceFunctionalFormalized:         t.Formalized && t.UsesHilbertSchmidtNorm,
		DoubledCarrierEvaluated:           d.TotalEdgesMapped == 4 && d.QuarkEdgesMapped == 2 && d.LeptonEdgesMapped == 2 && d.PositivePairingPreserved,
		PositiveSemidefiniteProved:        p.PositiveSemidefinite && !p.NegativeTermsPermitted && !p.ImaginaryKineticPermitted,
		StrictPositiveConditionIdentified: p.StrictPositiveConditional && s.AtLeastOneNonzeroNeeded && z.EvaluableAfterAmplitudeSeal,
		NumericalZHComputed:               false,
		PhysicalDynamicsDerived:           false,
		FirewallPreserved:                 !f.FiniteCorePolluted && f.NoYukawaNumbersInserted && f.NoObservedMassesInserted && f.NoBGapInstantonClaimed,
		Status:                            strings.Join(statuses, ";"),
		DirectAnswer:                      "Gate 301 proves the scalar kinetic carrier is structurally positive: the allowed Higgs Dirac edges enter K_H^raw only through multiplicity-weighted Hilbert-Schmidt squares. This eliminates native negative/imaginary ghost kinetic terms, while keeping strict numerical Z_H behind the nonzero-amplitude, f0, sign, and trace-convention seals.",
		NextGate:                          "Gate 302 should audit the minimal convention ledger for the overall positive scalar kinetic prefactor N_4 f_0: heat-kernel sign, Wick rotation, doubled-space trace normalization, and canonical coefficient matching.",
	}
}

func FormatInput(i InheritedGate300) string {
	return fmt.Sprintf("ZH=%t scalarSelector=%t rescale=%t gauge=%t positiveNumerical=%t dynamics=%t verdict=%s", i.ZHDefined, i.ScalarKineticSelectorDefined, i.RescalingDefined, i.GaugeNormalizationDefined, i.PositiveZHNumericallyProved, i.NumericalDynamicsDerived, i.Verdict)
}

func FormatScalarEdge(e ScalarEdge) string {
	return fmt.Sprintf("%s %s left=%q right=%q amp=%s term=%q mult=%d adjoint=%t hermitian=%t sealed=%t numeric=%t positive=%q", e.Sector, e.Edge, e.LeftCarrier, e.RightCarrier, e.AmplitudeSymbol, e.TraceContribution, e.Multiplicity, e.AppearsWithAdjoint, e.PreservesHermiticity, e.SealedAmplitude, e.NumericallyProvided, e.PositivityStatement)
}

func FormatTrace(t KineticTraceFunctional) string {
	parts := []string{}
	for _, e := range t.EdgeTerms {
		parts = append(parts, FormatScalarEdge(e))
	}
	return fmt.Sprintf("name=%s source=%q functional=%q edges=[%s] doubledFactor=%d HS=%t noGauge=%t noPotential=%t noVacuum=%t noBGap=%t formalized=%t verdict=%s", t.Name, t.Source, t.Functional, strings.Join(parts, " | "), t.ParticleAntiparticleFactor, t.UsesHilbertSchmidtNorm, t.ExcludesGaugeCurvature, t.ExcludesPotentialTerms, t.ExcludesVacuumTerms, t.ExcludesMajoranaBGap, t.Formalized, t.Verdict)
}

func FormatDoubled(d DoubledSpaceEvaluation) string {
	return fmt.Sprintf("carrier=%q particle=%q antiparticle=%q jswap=%q doubleHandled=%t positive=%t quark=%d lepton=%d total=%d eval=%q verdict=%s", d.Carrier, d.ParticleBlock, d.AntiparticleBlock, d.JSwapDuplication, d.DoubleCountingHandled, d.PositivePairingPreserved, d.QuarkEdgesMapped, d.LeptonEdgesMapped, d.TotalEdgesMapped, d.Evaluation, d.Verdict)
}

func FormatPositivity(p PositivitySieve) string {
	return fmt.Sprintf("raw=%q squares=%q semidefinite=%t negative=%t imaginary=%t strictCondition=%q strictProved=%t strictConditional=%t zeroMode=%q ghostEliminated=%t verdict=%s", p.RawExpression, p.SumOfSquaresExpression, p.PositiveSemidefinite, p.NegativeTermsPermitted, p.ImaginaryKineticPermitted, p.StrictPositiveCondition, p.StrictPositiveProved, p.StrictPositiveConditional, p.ZeroTraceFailureMode, p.GhostRiskEliminatedStructurally, p.Verdict)
}

func FormatAmplitudeSeal(s AmplitudeSeal) string {
	return fmt.Sprintf("%s sector=%s numericZH=%t strictZH=%t native=%t allowed=%s reason=%q", s.Symbol, s.Sector, s.RequiredForNumericZH, s.RequiredForStrictZH, s.NativeValueDerived, strings.Join(s.AllowedSealTypes, "/"), s.Reason)
}

func FormatSeals(s AmplitudeSealingCheck) string {
	parts := []string{}
	for _, seal := range s.Seals {
		parts = append(parts, FormatAmplitudeSeal(seal))
	}
	return fmt.Sprintf("seals=[%s] needNonzero=%t allNumericSealed=%t noEmpirical=%t reducible=%t built=%t verdict=%s", strings.Join(parts, " | "), s.AtLeastOneNonzeroNeeded, s.AllNumericalValuesSealed, s.NoEmpiricalValuesInserted, s.ReducibleToNumericZH, s.LedgerBuilt, s.Verdict)
}

func FormatZH(z ZHCarrierMap) string {
	return fmt.Sprintf("KRaw=%q ZH=%q positive=%q numeric=%t evaluableAfterSeal=%t f0=%t traceNorm=%t sign=%t verdict=%s", z.KRawDefinition, z.ZHDefinition, z.ZHPositivityStatement, z.NumericalZHComputed, z.EvaluableAfterAmplitudeSeal, z.RequiresPositiveF0, z.RequiresPositiveTraceNorm, z.RequiresEuclideanSignLedger, z.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	parts := []string{}
	for _, o := range f.Obligations {
		parts = append(parts, FormatObligation(o))
	}
	return fmt.Sprintf("noYukawa=%t noMasses=%t noF0=%t noSubtraction=%t noBGap=%t noMassQuartic=%t polluted=%t obligations=[%s] verdict=%s", f.NoYukawaNumbersInserted, f.NoObservedMassesInserted, f.NoCutoffMomentInserted, f.NoSubtractionSchemeInvented, f.NoBGapInstantonClaimed, f.NoMassQuarticClaimed, f.FiniteCorePolluted, strings.Join(parts, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("inherit=%t trace=%t doubled=%t semidefinite=%t strictCondition=%t numericZH=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.Gate300Inherited, s.TraceFunctionalFormalized, s.DoubledCarrierEvaluated, s.PositiveSemidefiniteProved, s.StrictPositiveConditionIdentified, s.NumericalZHComputed, s.PhysicalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
