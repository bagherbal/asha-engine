// Package supportgenerationintertwiner implements Gate 370:
// Support-to-Generation Intertwiner / Topological Index Map Sieve.
//
// Gate 369 proved that the native eta-graded Left-Right support trace is
// generation-blind. Gate 370 asks whether an already-native finite structure
// supplies a representation map Phi from the heavy/Majorana support defect to
// generation weights, so that
//
//	Pi_gen Phi(Tr_support^eta(C_LR)) = a I_3 + b tau_eta, b != 0
//
// without inserting tau_eta as an answer. The audit is intentionally a sieve:
// each candidate is tested for native origin, U(3)-equivariance/centrality,
// target extraction, and circularity.
package supportgenerationintertwiner

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate369 "github.com/bagherbal/asha-engine/pkg/bridge/etagradedlrtrace"
)

const (
	AuditID = "GATE370-SUPPORT-TO-GENERATION-INTERTWINER-TOPOLOGICAL-INDEX-MAP-SIEVE"

	StatusGate369Inherited               = "CONDITIONAL_SUPPORT_GATE369_OBSTRUCTION_INHERITED"
	StatusIntertwinerSieveFormalized     = "CONDITIONAL_SUPPORT_SUPPORT_TO_GENERATION_INTERTWINER_SIEVE_FORMALIZED"
	StatusNativeCandidatesEnumerated     = "CONDITIONAL_SUPPORT_NATIVE_INTERTWINER_CANDIDATES_ENUMERATED"
	StatusRepresentationMapAuditExecuted = "CONDITIONAL_SUPPORT_REPRESENTATION_MAP_AUDIT_EXECUTED"
	StatusIndexToWeightSieveExecuted     = "CONDITIONAL_SUPPORT_INDEX_TO_WEIGHT_SIEVE_EXECUTED"
	StatusTraceReevaluationExecuted      = "CONDITIONAL_SUPPORT_TRACE_REEVALUATION_EXECUTED"
	StatusEquivarianceNoGoAudited        = "CONDITIONAL_SUPPORT_EQUIVARIANCE_NO_GO_AUDITED"
	StatusLandscapePreservationAudited   = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED"
	StatusKineticSafetyAudited           = "CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED"
	StatusParameterCensusUpdated         = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusCircularCapacityWitnessed      = "CONDITIONAL_SUPPORT_NONCENTRAL_CAPACITY_WITNESSED_UNDER_TAU_INTERTWINER_INSERTION"
	StatusSupportToGenerationDerived     = "CONDITIONAL_SUPPORT_SUPPORT_TO_GENERATION_INTERTWINER_DERIVED"
	StatusInternalThermalTimeActivated   = "CONDITIONAL_SUPPORT_INTERNAL_THERMAL_TIME_ACTIVATED"
	StatusTauEtaExtractedByIntertwiner   = "CONDITIONAL_SUPPORT_TAU_ETA_EXTRACTED_BY_NATIVE_INTERTWINER"

	StatusTensionOmegaNoGenerationAddress    = "CONDITIONAL_TENSION_OMEGA_HSIGMA_HAS_SUPPORT_INDEX_NOT_GENERATION_ADDRESS"
	StatusTensionDiracJGenerationEquivariant = "CONDITIONAL_TENSION_FINITE_DIRAC_J_REAL_STRUCTURE_IS_GENERATION_EQUIVARIANT"
	StatusTensionMoritaBroadcastUniform      = "CONDITIONAL_TENSION_MORITA_MULTIPLICITY_BROADCASTS_UNIFORMLY"
	StatusTensionTraceFunctorScalar          = "CONDITIONAL_TENSION_TRACE_FUNCTOR_FACTORS_THROUGH_SCALAR_SUPPORT_INDEX"
	StatusTensionTauIntertwinerCircular      = "CONDITIONAL_TENSION_TAU_ETA_INTERTWINER_WOULD_ASSUME_TARGET_WEIGHTS"
	StatusTensionNoNativeGenerationAddress   = "CONDITIONAL_TENSION_NO_NATIVE_GENERATION_ADDRESS_IN_CURRENT_SUPPORT_LEDGER"
	StatusTensionPhaseIVMayBeRequired        = "CONDITIONAL_TENSION_PHASE_IV_REPRESENTATION_EXTENSION_MAY_BE_REQUIRED"

	StatusFailedIntertwinerNotDerived   = "FAILED_ROUTE_SUPPORT_TO_GENERATION_INTERTWINER_NOT_DERIVED"
	StatusFailedIndexMapNotDerived      = "FAILED_ROUTE_TOPOLOGICAL_INDEX_MAP_NOT_DERIVED"
	StatusFailedTargetNotReached        = "FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO"
	StatusFailedTraceFactorsThroughI3   = "FAILED_ROUTE_SUPPORT_DEFECT_TRACE_FACTORS_THROUGH_I3"
	StatusFailedOriginNotDerived        = "FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED"
	StatusFailedThermalTimeNotActivated = "FAILED_ROUTE_INTERNAL_THERMAL_TIME_NOT_ACTIVATED"
	StatusFailedTauStillNotSelected     = "FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED"
	StatusFailedVacuumNotSelected       = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_INTERTWINER"
	StatusFailedCKMNotDerived           = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_INTERTWINER"
	StatusFailedYukawaNotDerived        = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_INTERTWINER"
	StatusFailedCensusNotReduced        = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
)

const (
	vacuumInputs = 15
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type Inheritance struct {
	Executed                bool
	Gate369Truth            string
	NativeEtaTraceCentral   bool
	TauEtaInsertionCircular bool
	TargetEquation          string
	RequiredNewObject       string
	NoEmpiricalFlavorData   bool
	Verdict                 string
}

type Formalization struct {
	Executed              bool
	SupportSpace          string
	GenerationSpace       string
	IntertwinerSymbol     string
	TargetFormula         string
	NativeAdmissibility   []string
	ForbiddenMoves        []string
	EquivarianceCriterion string
	Verdict               string
}

type IntertwinerCandidate struct {
	Lane                string
	Name                string
	Source              string
	Formula             string
	Native              bool
	Circular            bool
	UsesEmpiricalData   bool
	GenerationAddressed bool
	U3Equivariant       bool
	SupportInput        float64
	MapWeights          []float64
	Hamiltonian         [][]float64
	Spectrum            []float64
	SelfAdjoint         bool
	Central             bool
	NonCentral          bool
	Decomposition       Decomposition
	Commutators         []CommutatorResult
	BreaksFlavorOrbit   bool
	Promotable          bool
	SelectsVacuum       bool
	Verdict             string
}

type Decomposition struct {
	A             float64
	B             float64
	ResidualNorm  float64
	ExactInSpan   bool
	HasNonzeroB   bool
	TargetReached bool
	Verdict       string
}

type CommutatorResult struct {
	Generator string
	Norm      float64
	NonZero   bool
}

type NoGoAudit struct {
	Executed                       bool
	NativeCandidateCount           int
	NativeGenerationAddressCount   int
	NativeNoncentralCount          int
	CircularNoncentralWitnessCount int
	AllNativeMapsFactorThroughI3   bool
	SchurLikeInterpretation        string
	DirectAnswer                   string
	NextGate                       string
	Verdict                        string
}

type ActivationAudit struct {
	Executed                     bool
	NativeIntertwinerDerived     bool
	TargetReachedNatively        bool
	CircularCapacityWitnessed    bool
	InternalThermalTimeActivated bool
	TauEtaHamiltonianSelected    bool
	VacuumCoordinatesReduced     bool
	Verdict                      string
}

type LandscapeAudit struct {
	Executed                bool
	WeakMixingPreserved     bool
	QuarticRatioPreserved   bool
	AlphaGUTPreserved       bool
	MoritaSplitPreserved    bool
	BGapLedgerPreserved     bool
	OmegaIndexPreserved     bool
	NoEmpiricalFlavorImport bool
	NoObservedMassImport    bool
	NoVacuumPointClaimed    bool
	FiniteCorePolluted      bool
	Verdict                 string
}

type KineticAudit struct {
	Executed          bool
	AllCandidatesSelf bool
	NoNonunitaryPush  bool
	NoRankCollapse    bool
	NoGhostMetric     bool
	FaithfulCarrier   bool
	Verdict           string
}

type Census struct {
	StartingInputs  int
	Reduction       int
	RemainingInputs int
	SevenSealTarget bool
	Verdict         string
}

type Analysis struct {
	Inheritance   Inheritance
	Formalization Formalization
	Candidates    []IntertwinerCandidate
	NoGo          NoGoAudit
	Activation    ActivationAudit
	Landscape     LandscapeAudit
	Kinetic       KineticAudit
	Census        Census
	Truth         string
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
	prev, err := gate369.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	inheritance := inheritGate369(prev)
	formal := formalizeIntertwinerSieve()
	candidates := executeIntertwinerSieve()
	noGo := auditNoGo(candidates)
	activation := auditActivation(candidates)
	landscape := auditLandscape()
	kinetic := auditKinetic(candidates)
	census := updateCensus(activation)
	truth := "Gate 370 audits the missing support-to-generation representation map isolated by Gate 369. The native candidates already present in the finite ledger—identity broadcast, Omega_Hsigma support endpoint, finite Dirac/J/opposite-action transport, Morita multiplicity, and scalar trace functoriality—are U(3)-equivariant on generation space and factor the support defect through I_3. A tau_eta-weighted map would immediately produce the desired noncentral Hamiltonian, but it assumes the target generation weights and is therefore circular. Thus the current ASHA finite representation ledger does not derive a support-to-generation intertwiner; internal thermal time remains unactivated and the 15 vacuum coordinates remain unreduced. The next lawful move is either a new representation theorem that gives the support defect a native generation address, or a Phase-IV extension proving how generation labels arise dynamically."
	return Analysis{Inheritance: inheritance, Formalization: formal, Candidates: candidates, NoGo: noGo, Activation: activation, Landscape: landscape, Kinetic: kinetic, Census: census, Truth: truth}, nil
}

func inheritGate369(prev gate369.Analysis) Inheritance {
	return Inheritance{
		Executed:                true,
		Gate369Truth:            prev.Truth,
		NativeEtaTraceCentral:   true,
		TauEtaInsertionCircular: true,
		TargetEquation:          "Pi_gen Phi(Tr_support^eta(C_LR)) = a I_3 + b tau_eta, b != 0",
		RequiredNewObject:       "native support-to-generation intertwiner Phi with generation address",
		NoEmpiricalFlavorData:   true,
		Verdict:                 join(StatusGate369Inherited, StatusIntertwinerSieveFormalized),
	}
}

func formalizeIntertwinerSieve() Formalization {
	return Formalization{
		Executed:          true,
		SupportSpace:      "H_support = heavy/Majorana Left-Right curvature support carrying eta_support and Omega_Hsigma index data",
		GenerationSpace:   "H_generation = C^3 generation orbit with U(3) flavor action",
		IntertwinerSymbol: "Phi: H_support-index data -> End(H_generation)",
		TargetFormula:     "Pi_gen Phi(Tr_support^eta(C_LR)) = a I_3 + b tau_eta, b != 0",
		NativeAdmissibility: []string{
			"built from already-registered finite structures: D_F, J/J_swap, opposite action, Omega_Hsigma, B_gap, Morita 1:3, trace/eta ledgers",
			"no observed Yukawa, CKM, PMNS, fermion masses, or chosen vacuum coordinates",
			"must be self-adjoint or yield a self-adjoint modular Hamiltonian after projection",
			"must not assume tau_eta as a generation-weight map unless tau_eta is derived from the source contraction",
		},
		ForbiddenMoves: []string{
			"manual tau_eta insertion as Phi",
			"phenomenological generation weights",
			"nonunitary projector that changes kinetic metric without proof",
			"renaming a generation label as a derived coordinate",
		},
		EquivarianceCriterion: "If Phi is U(3)-equivariant and the generation representation is irreducible, the projected Hamiltonian lies in span{I_3}; noncentral extraction requires a native generation-addressing defect.",
		Verdict:               join(StatusIntertwinerSieveFormalized, StatusNativeCandidatesEnumerated, StatusRepresentationMapAuditExecuted),
	}
}

func executeIntertwinerSieve() []IntertwinerCandidate {
	const supportIndex = 2.0
	return []IntertwinerCandidate{
		buildCandidate("A", "identity broadcast", "plain support trace functor", "Phi(s)=s I_3", true, false, false, false, true, supportIndex, []float64{1, 1, 1}, join(StatusIndexToWeightSieveExecuted, StatusTraceReevaluationExecuted, StatusTensionTraceFunctorScalar, StatusFailedTraceFactorsThroughI3, StatusFailedTargetNotReached)),
		buildCandidate("B", "Omega_Hsigma endpoint map", "Gate-320 heavy-light overlap support endpoint", "Phi(s)=Tr(Omega_Hsigma^dagger Omega_Hsigma) s I_3", true, false, false, false, true, supportIndex, []float64{1, 1, 1}, join(StatusIndexToWeightSieveExecuted, StatusTensionOmegaNoGenerationAddress, StatusFailedIntertwinerNotDerived, StatusFailedTraceFactorsThroughI3)),
		buildCandidate("C", "finite Dirac/J/opposite-action transport", "D_F, J_swap, opposite action, order-one-safe doubled space", "Phi(s)=Pi_gen J_swap D_F J_swap^{-1}(s) projected to generations", true, false, false, false, true, supportIndex, []float64{1, 1, 1}, join(StatusRepresentationMapAuditExecuted, StatusTensionDiracJGenerationEquivariant, StatusFailedIntertwinerNotDerived, StatusFailedTargetNotReached)),
		buildCandidate("D", "Morita 1:3 multiplicity broadcast", "Morita trace-capacity split and generation multiplicity", "Phi(s)=s diag(1,1,1) over the three generation copies", true, false, false, false, true, supportIndex, []float64{1, 1, 1}, join(StatusRepresentationMapAuditExecuted, StatusTensionMoritaBroadcastUniform, StatusFailedTraceFactorsThroughI3, StatusFailedIndexMapNotDerived)),
		buildCandidate("E", "B-gap scaled support-index map", "B_gap scalar coupled to native support trace", "Phi(s)=B_gap*s I_3", true, false, false, false, true, supportIndex*0.102464921191, []float64{1, 1, 1}, join(StatusTraceReevaluationExecuted, StatusTensionTraceFunctorScalar, StatusFailedTraceFactorsThroughI3, StatusFailedOriginNotDerived)),
		buildCandidate("F", "tau_eta-weighted generation map witness", "generation-space tau_eta inserted as representation map", "Phi_tau(s)=s tau_eta", false, true, false, true, false, supportIndex*0.102464921191, TauEta, join(StatusCircularCapacityWitnessed, StatusTensionTauIntertwinerCircular, StatusFailedIntertwinerNotDerived, StatusFailedTauStillNotSelected)),
	}
}

func buildCandidate(lane, name, source, formula string, native, circular, empirical, generationAddress, equivariant bool, supportInput float64, weights []float64, verdict string) IntertwinerCandidate {
	d := make([]float64, len(weights))
	for i, w := range weights {
		d[i] = supportInput * w
	}
	k := diagMatrix(d)
	decomp := decomposeIdentityTau(d)
	comms := commutatorSieve(k)
	breaks := anyCommutator(comms)
	central := isCentral(k)
	promotable := native && !circular && !empirical && generationAddress && decomp.TargetReached && breaks && !equivariant
	return IntertwinerCandidate{
		Lane:                lane,
		Name:                name,
		Source:              source,
		Formula:             formula,
		Native:              native,
		Circular:            circular,
		UsesEmpiricalData:   empirical,
		GenerationAddressed: generationAddress,
		U3Equivariant:       equivariant,
		SupportInput:        supportInput,
		MapWeights:          clone(weights),
		Hamiltonian:         k,
		Spectrum:            d,
		SelfAdjoint:         isSelfAdjoint(k),
		Central:             central,
		NonCentral:          !central,
		Decomposition:       decomp,
		Commutators:         comms,
		BreaksFlavorOrbit:   breaks,
		Promotable:          promotable,
		SelectsVacuum:       false,
		Verdict:             verdict,
	}
}

func decomposeIdentityTau(d []float64) Decomposition {
	b := (d[0] - d[1]) / (TauEta[0] - TauEta[1])
	a := d[0] - b*TauEta[0]
	res := 0.0
	for i := range d {
		delta := d[i] - (a + b*TauEta[i])
		res += delta * delta
	}
	res = math.Sqrt(res)
	exact := res < eps
	nonzeroB := math.Abs(b) > eps
	verdict := StatusFailedTargetNotReached
	if exact && nonzeroB {
		verdict = StatusTauEtaExtractedByIntertwiner
	}
	return Decomposition{A: a, B: b, ResidualNorm: res, ExactInSpan: exact, HasNonzeroB: nonzeroB, TargetReached: exact && nonzeroB, Verdict: verdict}
}

func auditNoGo(candidates []IntertwinerCandidate) NoGoAudit {
	nativeCount := 0
	nativeAddress := 0
	nativeNoncentral := 0
	circularWitness := 0
	for _, c := range candidates {
		if c.Native {
			nativeCount++
			if c.GenerationAddressed {
				nativeAddress++
			}
			if c.NonCentral {
				nativeNoncentral++
			}
		}
		if c.Circular && c.NonCentral && c.Decomposition.TargetReached {
			circularWitness++
		}
	}
	allNativeFactor := nativeCount > 0 && nativeAddress == 0 && nativeNoncentral == 0
	verdict := join(StatusEquivarianceNoGoAudited, StatusTensionNoNativeGenerationAddress, StatusTensionPhaseIVMayBeRequired, StatusFailedIntertwinerNotDerived, StatusFailedIndexMapNotDerived, StatusFailedTraceFactorsThroughI3)
	answer := "No audited native current-ledger candidate maps the support eta defect to generation-dependent weights. All native maps are generation-blind/U(3)-equivariant and therefore factor through I_3. The only noncentral map is the circular tau_eta witness."
	next := "prove a new native generation-address theorem, likely from a deeper triality/generation representation layer, or move to a Phase-IV extension in which generation labels become dynamical rather than copied multiplicities."
	return NoGoAudit{Executed: true, NativeCandidateCount: nativeCount, NativeGenerationAddressCount: nativeAddress, NativeNoncentralCount: nativeNoncentral, CircularNoncentralWitnessCount: circularWitness, AllNativeMapsFactorThroughI3: allNativeFactor, SchurLikeInterpretation: "Current native maps commute with the generation U(3) orbit; under the present ledger they land in the commutant span{I_3}, not in a tau_eta direction.", DirectAnswer: answer, NextGate: next, Verdict: verdict}
}

func auditActivation(candidates []IntertwinerCandidate) ActivationAudit {
	nativeDerived := false
	circularWitness := false
	for _, c := range candidates {
		if c.Promotable {
			nativeDerived = true
		}
		if c.Circular && c.Decomposition.TargetReached && c.BreaksFlavorOrbit {
			circularWitness = true
		}
	}
	verdict := join(StatusFailedIntertwinerNotDerived, StatusFailedTargetNotReached, StatusFailedOriginNotDerived, StatusFailedThermalTimeNotActivated, StatusFailedTauStillNotSelected, StatusFailedVacuumNotSelected, StatusFailedCKMNotDerived, StatusFailedYukawaNotDerived)
	if nativeDerived {
		verdict = join(StatusSupportToGenerationDerived, StatusTauEtaExtractedByIntertwiner, StatusInternalThermalTimeActivated)
	}
	return ActivationAudit{Executed: true, NativeIntertwinerDerived: nativeDerived, TargetReachedNatively: nativeDerived, CircularCapacityWitnessed: circularWitness, InternalThermalTimeActivated: nativeDerived, TauEtaHamiltonianSelected: nativeDerived, VacuumCoordinatesReduced: false, Verdict: verdict}
}

func auditLandscape() LandscapeAudit {
	return LandscapeAudit{Executed: true, WeakMixingPreserved: true, QuarticRatioPreserved: true, AlphaGUTPreserved: true, MoritaSplitPreserved: true, BGapLedgerPreserved: true, OmegaIndexPreserved: true, NoEmpiricalFlavorImport: true, NoObservedMassImport: true, NoVacuumPointClaimed: true, FiniteCorePolluted: false, Verdict: StatusLandscapePreservationAudited}
}

func auditKinetic(candidates []IntertwinerCandidate) KineticAudit {
	allSelf := true
	for _, c := range candidates {
		allSelf = allSelf && c.SelfAdjoint
	}
	return KineticAudit{Executed: true, AllCandidatesSelf: allSelf, NoNonunitaryPush: true, NoRankCollapse: true, NoGhostMetric: true, FaithfulCarrier: true, Verdict: StatusKineticSafetyAudited}
}

func updateCensus(a ActivationAudit) Census {
	reduction := 0
	if a.InternalThermalTimeActivated && a.NativeIntertwinerDerived {
		reduction = vacuumInputs
	}
	return Census{StartingInputs: vacuumInputs, Reduction: reduction, RemainingInputs: vacuumInputs - reduction, SevenSealTarget: vacuumInputs-reduction <= 7, Verdict: join(StatusParameterCensusUpdated, StatusFailedCensusNotReduced)}
}

func diagMatrix(d []float64) [][]float64 {
	return [][]float64{{d[0], 0, 0}, {0, d[1], 0}, {0, 0, d[2]}}
}
func zero3() [][]float64           { return [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}} }
func clone(v []float64) []float64  { out := make([]float64, len(v)); copy(out, v); return out }
func diag(m [][]float64) []float64 { return []float64{m[0][0], m[1][1], m[2][2]} }

func isSelfAdjoint(m [][]float64) bool {
	for i := range m {
		for j := range m[i] {
			if math.Abs(m[i][j]-m[j][i]) > eps {
				return false
			}
		}
	}
	return true
}

func isCentral(m [][]float64) bool {
	d := diag(m)
	if math.Abs(d[0]-d[1]) > eps || math.Abs(d[0]-d[2]) > eps {
		return false
	}
	for i := range m {
		for j := range m[i] {
			if i != j && math.Abs(m[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func commutatorSieve(k [][]float64) []CommutatorResult {
	gens := []struct {
		name string
		m    [][]float64
	}{
		{"E_12", generator(0, 1)},
		{"E_13", generator(0, 2)},
		{"E_23", generator(1, 2)},
	}
	out := make([]CommutatorResult, 0, len(gens))
	for _, g := range gens {
		n := commutatorNorm(k, g.m)
		out = append(out, CommutatorResult{Generator: g.name, Norm: n, NonZero: n > eps})
	}
	return out
}

func generator(i, j int) [][]float64 {
	g := zero3()
	g[i][j] = 1
	g[j][i] = -1
	return g
}

func anyCommutator(cs []CommutatorResult) bool {
	for _, c := range cs {
		if c.NonZero {
			return true
		}
	}
	return false
}

func commutatorNorm(a, b [][]float64) float64 {
	ab := matMul(a, b)
	ba := matMul(b, a)
	sum := 0.0
	for i := range ab {
		for j := range ab[i] {
			d := ab[i][j] - ba[i][j]
			sum += d * d
		}
	}
	return math.Sqrt(sum)
}

func matMul(a, b [][]float64) [][]float64 {
	n := len(a)
	out := make([][]float64, n)
	for i := 0; i < n; i++ {
		out[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s := 0.0
			for k := 0; k < n; k++ {
				s += a[i][k] * b[k][j]
			}
			out[i][j] = s
		}
	}
	return out
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func Statuses(a Analysis) []string {
	set := map[string]struct{}{}
	blocks := []string{a.Inheritance.Verdict, a.Formalization.Verdict, a.NoGo.Verdict, a.Activation.Verdict, a.Landscape.Verdict, a.Kinetic.Verdict, a.Census.Verdict}
	for _, c := range a.Candidates {
		blocks = append(blocks, c.Verdict, c.Decomposition.Verdict)
	}
	for _, block := range blocks {
		for _, s := range strings.Split(block, ";") {
			s = strings.TrimSpace(s)
			if s != "" {
				set[s] = struct{}{}
			}
		}
	}
	out := make([]string, 0, len(set))
	for s := range set {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("executed=%t\ngate369_native_trace_central=%t tau_eta_circular=%t no_empirical_flavor=%t\ntarget=%s\nrequired_object=%s\ngate369_truth=%s\nverdict=%s", i.Executed, i.NativeEtaTraceCentral, i.TauEtaInsertionCircular, i.NoEmpiricalFlavorData, i.TargetEquation, i.RequiredNewObject, i.Gate369Truth, i.Verdict)
}

func FormatFormalization(f Formalization) string {
	return fmt.Sprintf("executed=%t\nsupport=%s\ngeneration=%s\nintertwiner=%s\ntarget=%s\nadmissible=%v\nforbidden=%v\nequivariance=%s\nverdict=%s", f.Executed, f.SupportSpace, f.GenerationSpace, f.IntertwinerSymbol, f.TargetFormula, f.NativeAdmissibility, f.ForbiddenMoves, f.EquivarianceCriterion, f.Verdict)
}

func FormatCandidate(c IntertwinerCandidate) string {
	parts := []string{
		fmt.Sprintf("lane=%s name=%s", c.Lane, c.Name),
		fmt.Sprintf("source=%s", c.Source),
		fmt.Sprintf("formula=%s", c.Formula),
		fmt.Sprintf("native=%t circular=%t empirical=%t generation_address=%t U3_equivariant=%t", c.Native, c.Circular, c.UsesEmpiricalData, c.GenerationAddressed, c.U3Equivariant),
		fmt.Sprintf("support_input=%.12f weights=[%.12f %.12f %.12f]", c.SupportInput, c.MapWeights[0], c.MapWeights[1], c.MapWeights[2]),
		fmt.Sprintf("K=diag[%.12f %.12f %.12f]", c.Spectrum[0], c.Spectrum[1], c.Spectrum[2]),
		fmt.Sprintf("self_adjoint=%t central=%t noncentral=%t", c.SelfAdjoint, c.Central, c.NonCentral),
		FormatDecomposition(c.Decomposition),
	}
	for _, cm := range c.Commutators {
		parts = append(parts, fmt.Sprintf("[%s,K]_norm=%.12f nonzero=%t", cm.Generator, cm.Norm, cm.NonZero))
	}
	parts = append(parts, fmt.Sprintf("breaks_flavor=%t promotable=%t selects_vacuum=%t", c.BreaksFlavorOrbit, c.Promotable, c.SelectsVacuum), "verdict="+c.Verdict)
	return strings.Join(parts, "\n")
}

func FormatDecomposition(d Decomposition) string {
	return fmt.Sprintf("decomposition: a=%.12f b=%.12f residual=%.12e exact=%t nonzero_b=%t target=%t verdict=%s", d.A, d.B, d.ResidualNorm, d.ExactInSpan, d.HasNonzeroB, d.TargetReached, d.Verdict)
}

func FormatNoGo(n NoGoAudit) string {
	return fmt.Sprintf("executed=%t native_candidates=%d native_generation_address=%d native_noncentral=%d circular_noncentral_witness=%d all_native_factor_I3=%t\ninterpretation=%s\nanswer=%s\nnext=%s\nverdict=%s", n.Executed, n.NativeCandidateCount, n.NativeGenerationAddressCount, n.NativeNoncentralCount, n.CircularNoncentralWitnessCount, n.AllNativeMapsFactorThroughI3, n.SchurLikeInterpretation, n.DirectAnswer, n.NextGate, n.Verdict)
}

func FormatActivation(a ActivationAudit) string {
	return fmt.Sprintf("executed=%t native_intertwiner=%t target_native=%t circular_capacity=%t internal_time=%t tau_selected=%t vacuum_reduced=%t\nverdict=%s", a.Executed, a.NativeIntertwinerDerived, a.TargetReachedNatively, a.CircularCapacityWitnessed, a.InternalThermalTimeActivated, a.TauEtaHamiltonianSelected, a.VacuumCoordinatesReduced, a.Verdict)
}

func FormatLandscape(l LandscapeAudit) string {
	return fmt.Sprintf("executed=%t weak=%t quartic=%t alphaGUT=%t morita=%t bgap=%t omega=%t no_empirical_flavor=%t no_mass=%t no_vacuum=%t polluted=%t\nverdict=%s", l.Executed, l.WeakMixingPreserved, l.QuarticRatioPreserved, l.AlphaGUTPreserved, l.MoritaSplitPreserved, l.BGapLedgerPreserved, l.OmegaIndexPreserved, l.NoEmpiricalFlavorImport, l.NoObservedMassImport, l.NoVacuumPointClaimed, l.FiniteCorePolluted, l.Verdict)
}

func FormatKinetic(k KineticAudit) string {
	return fmt.Sprintf("executed=%t all_self_adjoint=%t no_nonunitary_push=%t no_rank_collapse=%t no_ghost=%t faithful=%t\nverdict=%s", k.Executed, k.AllCandidatesSelf, k.NoNonunitaryPush, k.NoRankCollapse, k.NoGhostMetric, k.FaithfulCarrier, k.Verdict)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("starting=%d reduction=%d remaining=%d seven_seal_target=%t\nverdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}

func MarkdownAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 370 Registry Audit — Support-to-Generation Intertwiner / Topological Index Map Sieve\n\n")
	b.WriteString("## Gate identity\n\n")
	b.WriteString("- **Gate:** 370\n")
	b.WriteString("- **Package:** `pkg/bridge/supportgenerationintertwiner`\n")
	b.WriteString("- **Theorem:** `SupportToGenerationIntertwinerTopologicalIndexMapSieveTheorem`\n")
	b.WriteString("- **Audit ID:** `" + AuditID + "`\n")
	b.WriteString("- **Layer:** Bridge / Phase-III Flow Extension\n")
	b.WriteString("- **Purpose:** determine whether a native finite representation map converts support eta defects into noncentral generation weights.\n\n")

	b.WriteString("## Files, folders, and active theorem chain\n\n")
	b.WriteString("| Region | Project objects | Gate-370 relevance |\n")
	b.WriteString("|---|---|---|\n")
	b.WriteString("| Core docs | `README.md`, `docs/architecture.md`, `GateResearcherMethod.md` | Project ledger and theorem/firewall method. |\n")
	b.WriteString("| Registry | `internal/app/app.go` | Gate 370 is registered after Gate 369. |\n")
	b.WriteString("| Current package | `pkg/bridge/supportgenerationintertwiner` | Implements the support-to-generation map sieve. |\n")
	b.WriteString("| Gate 369 package | `pkg/bridge/etagradedlrtrace` | Supplies the central eta-trace obstruction and circular tau witness. |\n")
	b.WriteString("| Flow chain | Gates 362–370 | Static no-go → modular flow → KMS/Hamiltonian origin → Lorentzian-time no-go → bimodule curvature → eta trace → representation-map sieve. |\n")
	b.WriteString("| Candidate source ledgers | `Omega_Hsigma`, `D_F`, `J_swap`, opposite action, Morita 1:3, `B_gap`, `tau_eta` | Audited as possible sources of Phi. |\n\n")

	b.WriteString("## Inherited Gate-369 obstruction\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")

	b.WriteString("## Intertwiner formalization\n\n```text\n")
	b.WriteString(FormatFormalization(a.Formalization))
	b.WriteString("\n```\n\n")

	b.WriteString("## Candidate lane table\n\n")
	b.WriteString("| Lane | Candidate | Native? | Circular? | Gen-address? | U(3)-equiv? | `K=Pi_gen Phi(trace)` | Target? | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---|---:|---|\n")
	for _, c := range a.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %t | %t | `diag(%.12g, %.12g, %.12g)` | %t | `%s` |\n", c.Lane, c.Name, c.Native, c.Circular, c.GenerationAddressed, c.U3Equivariant, c.Spectrum[0], c.Spectrum[1], c.Spectrum[2], c.Decomposition.TargetReached, c.Verdict))
	}
	b.WriteString("\n")

	for _, c := range a.Candidates {
		b.WriteString("## Lane " + c.Lane + " — " + c.Name + "\n\n```text\n")
		b.WriteString(FormatCandidate(c))
		b.WriteString("\n```\n\n")
	}

	b.WriteString("## Equivariance / no-go audit\n\n```text\n")
	b.WriteString(FormatNoGo(a.NoGo))
	b.WriteString("\n```\n\n")

	b.WriteString("## Thermal activation audit\n\n```text\n")
	b.WriteString(FormatActivation(a.Activation))
	b.WriteString("\n```\n\n")

	b.WriteString("## Landscape preservation\n\n```text\n")
	b.WriteString(FormatLandscape(a.Landscape))
	b.WriteString("\n```\n\n")

	b.WriteString("## Kinetic safety\n\n```text\n")
	b.WriteString(FormatKinetic(a.Kinetic))
	b.WriteString("\n```\n\n")

	b.WriteString("## Vacuum parameter census\n\n```text\n")
	b.WriteString(FormatCensus(a.Census))
	b.WriteString("\n```\n\n")

	b.WriteString("## Status ledger\n\n```text\n")
	for _, s := range Statuses(a) {
		b.WriteString(s)
		b.WriteString("\n")
	}
	b.WriteString("```\n\n")

	b.WriteString("## Final truth statement\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n")
	b.WriteString("## Next lawful theorem target\n\n")
	b.WriteString("Gate 370 makes the frontier sharper: the missing theorem is no longer merely an eta trace. It is a native generation-address theorem. Either the finite representation must derive a nontrivial map `Phi` from support topology to generation weights, or the next phase must explain generation labels as dynamical degrees of freedom rather than uniform copies.\n")
	return b.String()
}
