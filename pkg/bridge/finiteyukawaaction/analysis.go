// Package finiteyukawaaction implements Gate 263:
// Finite Yukawa Action Functional / Triality-Hopf Amplitude Qualification Audit.
//
// Gate 262 exposed exact Hermitian triality real/phase matrices that populate
// the six-dimensional off-diagonal complement of ad_tau on M3(C). Gate 263 asks
// the dynamical question: does the finite core already contain an action
// functional that assigns canonical Yukawa amplitudes to those matrices, or are
// they still only a lawful basis awaiting an action/coefficient rule?
package finiteyukawaaction

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/tauetamixingpartner"
)

const (
	AuditID = "GATE263-FINITE-YUKAWA-ACTION-FUNCTIONAL-TRIALITY-HOPF-AMPLITUDE-QUALIFICATION-AUDIT"

	StatusGate262Inherited             = "CONDITIONAL_SUPPORT_GATE262_HERMITIAN_TRIALITY_BASIS_INHERITED"
	StatusTraceFunctionalsEvaluated    = "CONDITIONAL_SUPPORT_M3_TRACE_FUNCTIONALS_EVALUATED"
	StatusTraceMetricDegenerate        = "FAILED_ROUTE_TRACE_FUNCTIONALS_DO_NOT_SELECT_AMPLITUDES"
	StatusCanonicalActionNoMixing      = "FAILED_ROUTE_CANONICAL_ACTION_HAS_NO_NONCOMMUTING_TEXTURE_TERM"
	StatusSpectralActionNotReady       = "FAILED_ROUTE_FINITE_SPECTRAL_ACTION_NOT_READY_FOR_YUKAWA_AMPLITUDES"
	StatusBGapNoActionMap              = "FAILED_ROUTE_B_GAP_ACTION_MAP_TO_M3_OFFDIAGONAL_MISSING"
	StatusHopfNoProjection             = "FAILED_ROUTE_HOPF_PHASE_TO_TRIALITY_PHASE_PROJECTION_MISSING"
	StatusNoFiniteYukawaAction         = "FAILED_ROUTE_FINITE_YUKAWA_ACTION_FUNCTIONAL_NOT_DERIVED"
	StatusEmpiricalYukawaSealPreserved = "CONDITIONAL_SUPPORT_EMPIRICAL_YUKAWA_SEAL_PRESERVED"
	StatusPhysicalTextureStillBlocked  = "FAILED_ROUTE_PHYSICAL_YUKAWA_TEXTURE_STILL_BLOCKED"
	StatusCKMPMNSMassesStillBlocked    = "FAILED_ROUTE_CKM_PMNS_AND_FERMION_MASSES_STILL_BLOCKED"
)

type Gate262Inheritance struct {
	BilinearCarrierDefined        bool
	TauEtaDiagonalSourceOpened    bool
	OffDiagonalComplementExposed  bool
	HermitianTrialityBasisExposed bool
	RawNonCommutingPartnerExists  bool
	PreviousQualifiedPartnerFound bool
	PreviousPhysicalYukawaTexture bool
	PreviousCKMPMNSDerived        bool
	PreviousFermionMassesDerived  bool
	TauEtaEigenvalues             []int
	RealBasisName                 string
	PhaseBasisName                string
	RealBasisMatrix               tauetamixingpartner.Matrix3
	PhaseBasisMatrix              tauetamixingpartner.Matrix3
	RealCommutatorNormSquared     int
	PhaseCommutatorNormSquared    int
	Verdict                       string
}

type TraceFunctionalAudit struct {
	FunctionalName              string
	Formula                     string
	AppliesToM3                 bool
	ExactEvaluation             bool
	RealBasisValue              int
	PhaseBasisValue             int
	CrossValue                  int
	TauRealCrossValue           int
	TauPhaseCrossValue          int
	NonZeroOnBasis              bool
	DistinguishesRealAndPhase   bool
	SelectsAmplitudeCoefficient bool
	PromotableToYukawaAction    bool
	Verdict                     string
}

type NativeActionCandidate struct {
	Name                         string
	Source                       string
	Available                    bool
	Canonical                    bool
	ActsOnM3BilinearCarrier      bool
	HasFiniteTraceOrVariation    bool
	EvaluatesTrialityBasis       bool
	AssignsNonzeroCoefficients   bool
	SelectsRelativeAmplitude     bool
	MapsBGapToOffDiagonal        bool
	MapsHopfPhaseToPhaseBasis    bool
	RequiresMissingIngredient    string
	PhysicalYukawaTextureDerived bool
	Verdict                      string
}

type ScalarPhaseIntegrationAudit struct {
	BGapAvailableAsScale              bool
	BGapActionCoefficientDerived      bool
	BGapGenerationEndomorphismDerived bool
	BGapCanWeightTrialityBasis        bool
	HopfPhaseLedgerAvailable          bool
	HopfProjectionToKTrialityDerived  bool
	HopfCanFixCPPhase                 bool
	ScalarPhaseIntegrationDerived     bool
	Verdict                           string
}

type TextureConstructionAudit struct {
	DiagonalTauSourceAvailable      bool
	HermitianOffDiagonalBasisExists bool
	TraceMetricAvailable            bool
	FiniteActionCoefficientRule     bool
	RelativeRealPhaseWeightSelected bool
	OverallYukawaScaleSelected      bool
	FermionKindDependenceSelected   bool
	CandidateFormula                string
	FreeParameters                  []string
	PhysicalTextureConstructed      bool
	EmpiricalYukawaSealRequired     bool
	Verdict                         string
}

type FirewallAudit struct {
	Gate262RawBasisPreserved            bool
	DoesNotPromoteTraceMetricToDynamics bool
	DoesNotPromoteSymmetryToAmplitude   bool
	DoesNotUseObservedMasses            bool
	DoesNotUseObservedMixingAngles      bool
	DoesNotUseBGapWithoutMap            bool
	DoesNotUseHopfWithoutProjection     bool
	DoesNotClaimSpectralTripleComplete  bool
	EmpiricalYukawaSealPreserved        bool
	FiniteCorePolluted                  bool
	Verdict                             string
}

type Summary struct {
	Gate262Inherited              bool
	TraceFunctionalsEvaluated     bool
	TraceMetricDegenerate         bool
	NativeActionCandidateCount    int
	ActionCandidateQualified      bool
	ScalarPhaseIntegrationDerived bool
	FiniteYukawaActionDerived     bool
	PhysicalYukawaTextureDerived  bool
	CKMPMNSDerived                bool
	FermionMassesDerived          bool
	Status                        string
	NextGate                      string
	Comment                       string
}

type Analysis struct {
	PreviousGate262  tauetamixingpartner.Analysis
	Inheritance      Gate262Inheritance
	TraceAudits      []TraceFunctionalAudit
	ActionCandidates []NativeActionCandidate
	ScalarPhase      ScalarPhaseIntegrationAudit
	Texture          TextureConstructionAudit
	Firewall         FirewallAudit
	Summary          Summary
	TruthStatement   string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := tauetamixingpartner.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 262 predecessor: %w", err)
			return
		}
		inh := inheritGate262(prev)
		traceAudits := auditTraceFunctionals(inh)
		actionCandidates := auditNativeActionCandidates(inh, traceAudits)
		scalarPhase := auditScalarPhaseIntegration(actionCandidates)
		texture := auditTextureConstruction(inh, traceAudits, actionCandidates, scalarPhase)
		firewall := auditFirewall(inh, traceAudits, actionCandidates, scalarPhase, texture)
		summary := summarize(inh, traceAudits, actionCandidates, scalarPhase, texture)
		truth := buildTruth(inh, traceAudits, actionCandidates, scalarPhase, texture)
		defaultA = Analysis{PreviousGate262: prev, Inheritance: inh, TraceAudits: traceAudits, ActionCandidates: actionCandidates, ScalarPhase: scalarPhase, Texture: texture, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate262(a tauetamixingpartner.Analysis) Gate262Inheritance {
	byName := map[string]tauetamixingpartner.CandidateAudit{}
	for _, c := range a.Candidates {
		byName[c.Name] = c
	}
	realBasis := byName["A_triality_real=C+C^T"]
	phaseBasis := byName["K_triality_phase=i(C-C^T)"]
	return Gate262Inheritance{
		BilinearCarrierDefined:        a.Inheritance.BilinearCarrierDefined,
		TauEtaDiagonalSourceOpened:    a.Inheritance.TauEtaActionDerived,
		OffDiagonalComplementExposed:  a.Inheritance.OffDiagonalComplementDimension == 6,
		HermitianTrialityBasisExposed: a.TrialityPartner.HermitianRealPartNonCommuting && a.TrialityPartner.HermitianImaginaryPartNonCommuting,
		RawNonCommutingPartnerExists:  a.PartnerVerdict.RawNonCommutingPartnerExists,
		PreviousQualifiedPartnerFound: a.PartnerVerdict.QualifiedFiniteMixingPartnerFound,
		PreviousPhysicalYukawaTexture: a.PartnerVerdict.PhysicalYukawaTextureDerived,
		PreviousCKMPMNSDerived:        a.PartnerVerdict.CKMPMNSDerived,
		PreviousFermionMassesDerived:  a.PartnerVerdict.FermionMassesDerived,
		TauEtaEigenvalues:             append([]int(nil), a.Inheritance.TauEtaEigenvalues...),
		RealBasisName:                 realBasis.Name,
		PhaseBasisName:                phaseBasis.Name,
		RealBasisMatrix:               realBasis.Matrix,
		PhaseBasisMatrix:              phaseBasis.Matrix,
		RealCommutatorNormSquared:     realBasis.CommutatorFrobeniusNormSquared,
		PhaseCommutatorNormSquared:    phaseBasis.CommutatorFrobeniusNormSquared,
		Verdict:                       StatusGate262Inherited + "; inherited exact Hermitian triality real/phase basis from Gate 262 without rerunning broader historical chains",
	}
}

func auditTraceFunctionals(inh Gate262Inheritance) []TraceFunctionalAudit {
	tau := diagonalTau(inh.TauEtaEigenvalues)
	A := inh.RealBasisMatrix
	K := inh.PhaseBasisMatrix
	return []TraceFunctionalAudit{
		{
			FunctionalName:              "linear trace",
			Formula:                     "Tr(X)",
			AppliesToM3:                 true,
			ExactEvaluation:             true,
			RealBasisValue:              trace(A).Re,
			PhaseBasisValue:             trace(K).Re,
			CrossValue:                  0,
			TauRealCrossValue:           trace(mul(tau, A)).Re,
			TauPhaseCrossValue:          trace(mul(tau, K)).Re,
			NonZeroOnBasis:              trace(A).Re != 0 || trace(A).Im != 0 || trace(K).Re != 0 || trace(K).Im != 0,
			DistinguishesRealAndPhase:   false,
			SelectsAmplitudeCoefficient: false,
			PromotableToYukawaAction:    false,
			Verdict:                     "linear trace vanishes on the off-diagonal triality basis and cannot weight flavor mixing",
		},
		{
			FunctionalName:              "Hilbert-Schmidt trace metric",
			Formula:                     "Tr(X†Y), especially Tr(A²), Tr(K²), Tr(AK)",
			AppliesToM3:                 true,
			ExactEvaluation:             true,
			RealBasisValue:              hilbertSchmidt(A, A).Re,
			PhaseBasisValue:             hilbertSchmidt(K, K).Re,
			CrossValue:                  hilbertSchmidt(A, K).Re,
			TauRealCrossValue:           hilbertSchmidt(tau, A).Re,
			TauPhaseCrossValue:          hilbertSchmidt(tau, K).Re,
			NonZeroOnBasis:              hilbertSchmidt(A, A).Re != 0 && hilbertSchmidt(K, K).Re != 0,
			DistinguishesRealAndPhase:   hilbertSchmidt(A, A).Re != hilbertSchmidt(K, K).Re || hilbertSchmidt(A, K).Re != 0 || hilbertSchmidt(A, K).Im != 0,
			SelectsAmplitudeCoefficient: false,
			PromotableToYukawaAction:    false,
			Verdict:                     "trace metric evaluates the basis but is O(2)-degenerate on {A,K}; it gives a norm, not a coefficient selector",
		},
		{
			FunctionalName:              "ad_tau commutator norm",
			Formula:                     "Tr([tau,X]†[tau,X])",
			AppliesToM3:                 true,
			ExactEvaluation:             true,
			RealBasisValue:              inh.RealCommutatorNormSquared,
			PhaseBasisValue:             inh.PhaseCommutatorNormSquared,
			CrossValue:                  hilbertSchmidt(commutator(tau, A), commutator(tau, K)).Re,
			TauRealCrossValue:           0,
			TauPhaseCrossValue:          0,
			NonZeroOnBasis:              inh.RealCommutatorNormSquared != 0 && inh.PhaseCommutatorNormSquared != 0,
			DistinguishesRealAndPhase:   inh.RealCommutatorNormSquared != inh.PhaseCommutatorNormSquared,
			SelectsAmplitudeCoefficient: false,
			PromotableToYukawaAction:    false,
			Verdict:                     "commutator norm confirms both bases are equally non-commuting with tau_eta; it does not choose real/phase weights",
		},
	}
}

func auditNativeActionCandidates(inh Gate262Inheritance, traces []TraceFunctionalAudit) []NativeActionCandidate {
	_ = inh
	return []NativeActionCandidate{
		{
			Name:                       "M3 trace/Hilbert-Schmidt diagnostic functional",
			Source:                     "Gate 263 local exact trace evaluation on Hom(G_R,G_L)",
			Available:                  true,
			Canonical:                  true,
			ActsOnM3BilinearCarrier:    true,
			HasFiniteTraceOrVariation:  true,
			EvaluatesTrialityBasis:     allTraceAuditsEvaluate(traces),
			AssignsNonzeroCoefficients: false,
			SelectsRelativeAmplitude:   false,
			RequiresMissingIngredient:  "Euler-Lagrange/action coefficient rule on the generation bilinear carrier",
			Verdict:                    "canonical trace metric exists but only supplies diagnostics and norms; it is not a finite Yukawa action",
		},
		{
			Name:                         "canonical scalar/gauge finite variational action",
			Source:                       "Gate 100 audited snapshot",
			Available:                    true,
			Canonical:                    true,
			ActsOnM3BilinearCarrier:      false,
			HasFiniteTraceOrVariation:    true,
			EvaluatesTrialityBasis:       false,
			AssignsNonzeroCoefficients:   false,
			SelectsRelativeAmplitude:     false,
			RequiresMissingIngredient:    "non-commuting generation-texture term or variation over Hom(G_R,G_L)",
			PhysicalYukawaTextureDerived: false,
			Verdict:                      "selects scalar kinetic and gauge Hessian data, but its own ledger leaves non-commuting textures unselected",
		},
		{
			Name:                         "finite spectral action / spectral triple audit",
			Source:                       "Gate 163 audited snapshot",
			Available:                    true,
			Canonical:                    false,
			ActsOnM3BilinearCarrier:      false,
			HasFiniteTraceOrVariation:    false,
			EvaluatesTrialityBasis:       false,
			AssignsNonzeroCoefficients:   false,
			SelectsRelativeAmplitude:     false,
			RequiresMissingIngredient:    "complete finite spectral triple: algebra representation, D_F, J, grading, order-one calculus, cutoff/test function, gauge fluctuation map",
			PhysicalYukawaTextureDerived: false,
			Verdict:                      "spectral pre-data exist, but the spectral action principle is not ready to generate Yukawa coefficients",
		},
		{
			Name:                         "finite Dirac D_F initialization",
			Source:                       "Gate 233 audited snapshot",
			Available:                    true,
			Canonical:                    false,
			ActsOnM3BilinearCarrier:      false,
			HasFiniteTraceOrVariation:    true,
			EvaluatesTrialityBasis:       false,
			AssignsNonzeroCoefficients:   false,
			SelectsRelativeAmplitude:     false,
			RequiresMissingIngredient:    "canonical block M, real structure/order-one calculus, and embedding of tau_eta/triality basis into D_F",
			PhysicalYukawaTextureDerived: false,
			Verdict:                      "odd self-adjoint D_F family exists only as an ansatz; no canonical Yukawa block is selected",
		},
		{
			Name:                         "matter Fock representation action",
			Source:                       "matter/action audited snapshot",
			Available:                    true,
			Canonical:                    true,
			ActsOnM3BilinearCarrier:      false,
			HasFiniteTraceOrVariation:    true,
			EvaluatesTrialityBasis:       false,
			AssignsNonzeroCoefficients:   false,
			SelectsRelativeAmplitude:     false,
			RequiresMissingIngredient:    "canonical eigenvector/charge representation that turns number-operator response into Yukawa texture",
			PhysicalYukawaTextureDerived: false,
			Verdict:                      "acts on the 16-state Fock basis as a number-operator response, not on M3(C) as a flavor texture",
		},
	}
}

func auditScalarPhaseIntegration(c []NativeActionCandidate) ScalarPhaseIntegrationAudit {
	_ = c
	return ScalarPhaseIntegrationAudit{
		BGapAvailableAsScale:              true,
		BGapActionCoefficientDerived:      false,
		BGapGenerationEndomorphismDerived: false,
		BGapCanWeightTrialityBasis:        false,
		HopfPhaseLedgerAvailable:          true,
		HopfProjectionToKTrialityDerived:  false,
		HopfCanFixCPPhase:                 false,
		ScalarPhaseIntegrationDerived:     false,
		Verdict:                           StatusBGapNoActionMap + "; " + StatusHopfNoProjection + "; B_gap and Hopf phase ledgers remain meaningful but representation-free for the M3(C) off-diagonal texture basis",
	}
}

func auditTextureConstruction(inh Gate262Inheritance, traces []TraceFunctionalAudit, candidates []NativeActionCandidate, scalar ScalarPhaseIntegrationAudit) TextureConstructionAudit {
	free := []string{"overall Yukawa scale", "diagonal normalization of tau_eta", "real triality amplitude a", "phase triality amplitude b", "fermion-kind sector map", "left/right basis convention"}
	qualified := anyQualifiedAction(candidates)
	return TextureConstructionAudit{
		DiagonalTauSourceAvailable:      inh.TauEtaDiagonalSourceOpened,
		HermitianOffDiagonalBasisExists: inh.HermitianTrialityBasisExposed,
		TraceMetricAvailable:            allTraceAuditsEvaluate(traces),
		FiniteActionCoefficientRule:     qualified,
		RelativeRealPhaseWeightSelected: qualified && scalar.HopfCanFixCPPhase,
		OverallYukawaScaleSelected:      qualified && scalar.BGapCanWeightTrialityBasis,
		FermionKindDependenceSelected:   false,
		CandidateFormula:                "Y_f = alpha*tau_eta + beta*(C+C^T) + gamma*i(C-C^T)",
		FreeParameters:                  free,
		PhysicalTextureConstructed:      false,
		EmpiricalYukawaSealRequired:     true,
		Verdict:                         StatusNoFiniteYukawaAction + "; the lawful texture ansatz is exposed, but alpha,beta,gamma and fermion-sector dependence are not selected by a finite action",
	}
}

func auditFirewall(inh Gate262Inheritance, traces []TraceFunctionalAudit, candidates []NativeActionCandidate, scalar ScalarPhaseIntegrationAudit, texture TextureConstructionAudit) FirewallAudit {
	return FirewallAudit{
		Gate262RawBasisPreserved:            inh.HermitianTrialityBasisExposed && inh.RawNonCommutingPartnerExists,
		DoesNotPromoteTraceMetricToDynamics: !anyTraceSelectsAmplitude(traces),
		DoesNotPromoteSymmetryToAmplitude:   !anyQualifiedAction(candidates),
		DoesNotUseObservedMasses:            true,
		DoesNotUseObservedMixingAngles:      true,
		DoesNotUseBGapWithoutMap:            !scalar.BGapCanWeightTrialityBasis,
		DoesNotUseHopfWithoutProjection:     !scalar.HopfCanFixCPPhase,
		DoesNotClaimSpectralTripleComplete:  true,
		EmpiricalYukawaSealPreserved:        texture.EmpiricalYukawaSealRequired && !texture.PhysicalTextureConstructed,
		FiniteCorePolluted:                  false,
		Verdict:                             "firewall holds: no trace diagnostic, symmetry matrix, B_gap scalar, or Hopf phase is promoted into a Yukawa amplitude without a derived action map",
	}
}

func summarize(inh Gate262Inheritance, traces []TraceFunctionalAudit, candidates []NativeActionCandidate, scalar ScalarPhaseIntegrationAudit, texture TextureConstructionAudit) Summary {
	qualified := anyQualifiedAction(candidates)
	return Summary{
		Gate262Inherited:              inh.HermitianTrialityBasisExposed,
		TraceFunctionalsEvaluated:     allTraceAuditsEvaluate(traces),
		TraceMetricDegenerate:         !anyTraceDistinguishesRealPhase(traces),
		NativeActionCandidateCount:    len(candidates),
		ActionCandidateQualified:      qualified,
		ScalarPhaseIntegrationDerived: scalar.ScalarPhaseIntegrationDerived,
		FiniteYukawaActionDerived:     false,
		PhysicalYukawaTextureDerived:  false,
		CKMPMNSDerived:                false,
		FermionMassesDerived:          false,
		Status:                        StatusNoFiniteYukawaAction,
		NextGate:                      "Gate 264 — EmpiricalYukawaSeal activation or finite D_F/order-one Yukawa block selector audit",
		Comment:                       "Gate 263 evaluates the available trace/action ledgers but finds no finite coefficient-selection rule for the Hermitian triality mixing basis.",
	}
}

func buildTruth(inh Gate262Inheritance, traces []TraceFunctionalAudit, candidates []NativeActionCandidate, scalar ScalarPhaseIntegrationAudit, texture TextureConstructionAudit) string {
	return fmt.Sprintf("Gate 263 truth: Gate 262's Hermitian triality basis (%s and %s) is a valid off-diagonal M3(C) arena, and exact trace diagnostics evaluate it. However, Tr and ad_tau norms are degenerate diagnostics, not dynamics; none of the %d audited native action candidates acts as a canonical Yukawa amplitude functional on Hom(G_R,G_L). B_gap remains scale-only (%t) and Hopf phases remain without projection (%t). Therefore the lawful ansatz %s is exposed but all coefficients remain unselected; CKM/PMNS and fermion masses stay behind the EmpiricalYukawaSeal.", inh.RealBasisName, inh.PhaseBasisName, len(candidates), !scalar.BGapCanWeightTrialityBasis, !scalar.HopfCanFixCPPhase, texture.CandidateFormula)
}

func allTraceAuditsEvaluate(a []TraceFunctionalAudit) bool {
	if len(a) == 0 {
		return false
	}
	for _, t := range a {
		if !t.AppliesToM3 || !t.ExactEvaluation {
			return false
		}
	}
	return true
}

func anyTraceSelectsAmplitude(a []TraceFunctionalAudit) bool {
	for _, t := range a {
		if t.SelectsAmplitudeCoefficient || t.PromotableToYukawaAction {
			return true
		}
	}
	return false
}

func anyTraceDistinguishesRealPhase(a []TraceFunctionalAudit) bool {
	for _, t := range a {
		if t.DistinguishesRealAndPhase {
			return true
		}
	}
	return false
}

func anyQualifiedAction(a []NativeActionCandidate) bool {
	for _, c := range a {
		if c.AssignsNonzeroCoefficients && c.SelectsRelativeAmplitude && c.PhysicalYukawaTextureDerived {
			return true
		}
	}
	return false
}

func diagonalTau(vals []int) tauetamixingpartner.Matrix3 {
	if len(vals) != 3 {
		vals = []int{2, -2, 1}
	}
	var m tauetamixingpartner.Matrix3
	for i, v := range vals {
		m[i][i] = tauetamixingpartner.GaussianInt{Re: v}
	}
	return m
}

func trace(m tauetamixingpartner.Matrix3) tauetamixingpartner.GaussianInt {
	out := tauetamixingpartner.GaussianInt{}
	for i := 0; i < 3; i++ {
		out = add(out, m[i][i])
	}
	return out
}

func hilbertSchmidt(a, b tauetamixingpartner.Matrix3) tauetamixingpartner.GaussianInt {
	out := tauetamixingpartner.GaussianInt{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out = add(out, mulZ(conj(a[i][j]), b[i][j]))
		}
	}
	return out
}

func commutator(a, b tauetamixingpartner.Matrix3) tauetamixingpartner.Matrix3 {
	return sub(mul(a, b), mul(b, a))
}

func mul(a, b tauetamixingpartner.Matrix3) tauetamixingpartner.Matrix3 {
	var out tauetamixingpartner.Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s := tauetamixingpartner.GaussianInt{}
			for k := 0; k < 3; k++ {
				s = add(s, mulZ(a[i][k], b[k][j]))
			}
			out[i][j] = s
		}
	}
	return out
}

func add(a, b tauetamixingpartner.GaussianInt) tauetamixingpartner.GaussianInt {
	return tauetamixingpartner.GaussianInt{Re: a.Re + b.Re, Im: a.Im + b.Im}
}

func sub(a, b tauetamixingpartner.Matrix3) tauetamixingpartner.Matrix3 {
	var out tauetamixingpartner.Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = tauetamixingpartner.GaussianInt{Re: a[i][j].Re - b[i][j].Re, Im: a[i][j].Im - b[i][j].Im}
		}
	}
	return out
}

func mulZ(a, b tauetamixingpartner.GaussianInt) tauetamixingpartner.GaussianInt {
	return tauetamixingpartner.GaussianInt{Re: a.Re*b.Re - a.Im*b.Im, Im: a.Re*b.Im + a.Im*b.Re}
}

func conj(a tauetamixingpartner.GaussianInt) tauetamixingpartner.GaussianInt {
	return tauetamixingpartner.GaussianInt{Re: a.Re, Im: -a.Im}
}

func candidateNames(c []NativeActionCandidate) []string {
	out := make([]string, 0, len(c))
	for _, x := range c {
		out = append(out, x.Name)
	}
	sort.Strings(out)
	return out
}

func statusList(a Analysis) []string {
	ss := []string{StatusGate262Inherited, StatusTraceFunctionalsEvaluated, StatusEmpiricalYukawaSealPreserved}
	if a.Summary.TraceMetricDegenerate {
		ss = append(ss, StatusTraceMetricDegenerate)
	}
	if !a.ActionCandidates[1].SelectsRelativeAmplitude {
		ss = append(ss, StatusCanonicalActionNoMixing)
	}
	if !a.ActionCandidates[2].SelectsRelativeAmplitude {
		ss = append(ss, StatusSpectralActionNotReady)
	}
	if !a.ScalarPhase.BGapCanWeightTrialityBasis {
		ss = append(ss, StatusBGapNoActionMap)
	}
	if !a.ScalarPhase.HopfCanFixCPPhase {
		ss = append(ss, StatusHopfNoProjection)
	}
	if !a.Summary.FiniteYukawaActionDerived {
		ss = append(ss, StatusNoFiniteYukawaAction, StatusPhysicalTextureStillBlocked, StatusCKMPMNSMassesStillBlocked)
	}
	return ss
}

func StatusLines(a Analysis) string { return strings.Join(statusList(a), "\n") }
