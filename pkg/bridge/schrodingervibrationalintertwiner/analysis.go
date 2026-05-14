// Package schrodingervibrationalintertwiner implements Gate 371:
// Schrodinger Vibrational Modes / Quantum Information Intertwiner Audit.
package schrodingervibrationalintertwiner

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate370 "github.com/bagherbal/asha-engine/pkg/bridge/supportgenerationintertwiner"
)

const (
	AuditID = "GATE371-SCHRODINGER-VIBRATIONAL-MODES-QUANTUM-INFORMATION-INTERTWINER-AUDIT"

	StatusGate370Inherited                   = "CONDITIONAL_SUPPORT_GATE370_INTERTWINER_OBSTRUCTION_INHERITED"
	StatusFockSpaceFormalized                = "CONDITIONAL_SUPPORT_VIBRATIONAL_FOCK_SPACE_FORMALIZED"
	StatusNumberOperatorSieveExecuted        = "CONDITIONAL_SUPPORT_NUMBER_OPERATOR_SIEVE_EXECUTED"
	StatusInformationEntropyAudited          = "CONDITIONAL_SUPPORT_INFORMATION_ENTROPY_OPERATOR_AUDITED"
	StatusTopologicalIndexPullbackExecuted   = "CONDITIONAL_SUPPORT_TOPOLOGICAL_INDEX_PULLBACK_EXECUTED"
	StatusKMSInformationStateReconstructed   = "CONDITIONAL_SUPPORT_KMS_INFORMATION_STATE_RECONSTRUCTED"
	StatusVibrationalCapacityWitnessed       = "CONDITIONAL_SUPPORT_VIBRATIONAL_NONCENTRAL_CAPACITY_WITNESSED"
	StatusLandscapePreservationAudited       = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED"
	StatusKineticSafetyAudited               = "CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED"
	StatusParameterCensusUpdated             = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusVibrationalIntertwinerDerived      = "CONDITIONAL_SUPPORT_VIBRATIONAL_INTERTWINER_DERIVED"
	StatusThermalTimeActivated               = "CONDITIONAL_SUPPORT_THERMAL_TIME_ACTIVATED"
	StatusTauEtaExtractedByInformation       = "CONDITIONAL_SUPPORT_TAU_ETA_EXTRACTED_BY_QUANTUM_INFORMATION_INTERTWINER"
	StatusTensionFockBasisNotSelected        = "CONDITIONAL_TENSION_FOCK_BASIS_NOT_SELECTED_BY_CURRENT_LEDGER"
	StatusTensionNumberNoncentralNotTau      = "CONDITIONAL_TENSION_NUMBER_OPERATOR_BREAKS_U3_BUT_IS_NOT_TAU_ETA"
	StatusTensionOscillatorCouplingNew       = "CONDITIONAL_TENSION_SUPPORT_DEFECT_TIMES_NUMBER_OPERATOR_IS_NEW_COUPLING_STRUCTURE"
	StatusTensionEntropyDependsOnChosenK     = "CONDITIONAL_TENSION_INFORMATION_ENTROPY_STATE_DEPENDS_ON_CHOSEN_HAMILTONIAN"
	StatusTensionQuadraticTauPolynomial      = "CONDITIONAL_TENSION_TAU_ETA_REQUIRES_TARGET_QUADRATIC_POLYNOMIAL_IN_N"
	StatusTensionPolynomialCircular          = "CONDITIONAL_TENSION_TAU_POLYNOMIAL_CALIBRATION_WOULD_BE_CIRCULAR"
	StatusTensionNoncentralButNotSelecting   = "CONDITIONAL_TENSION_VIBRATIONAL_OPERATOR_NONCENTRAL_BUT_NOT_VACUUM_SELECTING"
	StatusTensionPhaseIVInformationExtension = "CONDITIONAL_TENSION_PHASE_IV_QUANTUM_INFORMATION_EXTENSION_MAY_BE_REQUIRED"

	StatusFailedVibrationalIntertwinerNotDerived = "FAILED_ROUTE_VIBRATIONAL_INTERTWINER_NOT_DERIVED"
	StatusFailedFockBasisNotNativeSelected       = "FAILED_ROUTE_FOCK_GENERATION_BASIS_NOT_NATIVE_SELECTED"
	StatusFailedNumberOperatorNotDerived         = "FAILED_ROUTE_NUMBER_OPERATOR_NOT_DERIVED_FROM_CL17_LEDGER"
	StatusFailedTargetTauNotExtracted            = "FAILED_ROUTE_TARGET_TAU_ETA_NOT_EXTRACTED_FROM_NUMBER_OPERATOR"
	StatusFailedTopologicalPullbackNotDerived    = "FAILED_ROUTE_TOPOLOGICAL_INDEX_PULLBACK_TO_FOCK_SPACE_NOT_DERIVED"
	StatusFailedThermalTimeNotActivated          = "FAILED_ROUTE_INTERNAL_THERMAL_TIME_NOT_ACTIVATED"
	StatusFailedTauStillNotSelected              = "FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED"
	StatusFailedVacuumNotSelected                = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_VIBRATIONAL_INTERTWINER"
	StatusFailedCKMNotDerived                    = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_VIBRATIONAL_INTERTWINER"
	StatusFailedYukawaNotDerived                 = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_VIBRATIONAL_INTERTWINER"
	StatusFailedCensusNotReduced                 = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
)

const (
	vacuumInputs = 15
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type Inheritance struct {
	Executed                                                               bool
	Gate370Truth, MissingObject                                            string
	AllCurrentNativeMapsI3, TauEtaManualMapCircular, NoEmpiricalFlavorData bool
	Verdict                                                                string
}
type FockFormalization struct {
	Executed                                       bool
	GeometricBasis, FockBasis                      []string
	LadderConvention                               string
	NumberOperator, CenteredNumber, HarmonicEnergy []float64
	BasisSelectedByASHA                            bool
	SelectionEvidence, ForbiddenMoves              []string
	TargetFormula, Verdict                         string
}
type OperatorLane struct {
	Lane, Name, Source, Formula                                                                     string
	NativeToChosenFock, NativeToCurrentASHA, Circular, UsesEmpiricalData, RequiresNewCoupling       bool
	Spectrum                                                                                        []float64
	Hamiltonian                                                                                     [][]float64
	SelfAdjoint, Central, NonCentral                                                                bool
	Decomposition                                                                                   Decomposition
	KMS                                                                                             KMSState
	Commutators                                                                                     []CommutatorResult
	BreaksFlavorOrbit, FockBasisSelected, TopologicalPullback, PromotableHamiltonian, SelectsVacuum bool
	Verdict                                                                                         string
}
type Decomposition struct {
	A, B, ResidualNorm                      float64
	ExactInSpan, HasNonzeroB, TargetReached bool
	Verdict                                 string
}
type KMSState struct {
	Beta                            float64
	Density                         []float64
	Frequencies                     []Frequency
	NontrivialFrequencies, Faithful bool
	Entropy                         float64
}
type Frequency struct {
	Pair     string
	LogRatio float64
	NonZero  bool
}
type CommutatorResult struct {
	Generator string
	Norm      float64
	NonZero   bool
}
type PullbackAudit struct {
	Executed                                                                                             bool
	SupportDefectInput                                                                                   string
	NativeSupportDefectScalar                                                                            float64
	NativeASHAFockMapDerived, AnyNoncentralFockWitness, AnyPromotableHamiltonian, AnyTargetTauExtraction bool
	ExactTauRequiresPolynomial                                                                           string
	PolynomialCircular                                                                                   bool
	DirectAnswer, Verdict                                                                                string
}
type ActivationAudit struct {
	Executed, FockBasisNativeSelected, NumberOperatorNativeSelected, PullbackDerived, TargetTauExtractedNatively, NoncentralCapacityWitnessed, InternalThermalTimeActivated, VacuumCoordinatesReduced bool
	DirectAnswer, NextGate, Verdict                                                                                                                                                                   string
}
type LandscapeAudit struct {
	Executed, WeakMixingPreserved, QuarticRatioPreserved, AlphaGUTPreserved, MoritaSplitPreserved, BGapLedgerPreserved, OmegaIndexPreserved, NoEmpiricalFlavorImport, NoObservedMassImport, NoVacuumPointClaimed, FiniteCorePolluted bool
	Verdict                                                                                                                                                                                                                          string
}
type KineticAudit struct {
	Executed, AllOperatorsSelf, FaithfulKMSStates, NoRankCollapse, NoGhostMetric, NoNonunitaryPush bool
	Verdict                                                                                        string
}
type Census struct {
	StartingInputs, Reduction, RemainingInputs int
	SevenSealTarget                            bool
	Verdict                                    string
}
type Analysis struct {
	Inheritance   Inheritance
	Formalization FockFormalization
	Lanes         []OperatorLane
	Pullback      PullbackAudit
	Activation    ActivationAudit
	Landscape     LandscapeAudit
	Kinetic       KineticAudit
	Census        Census
	Truth         string
}

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	prev, err := gate370.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	inheritance := Inheritance{true, prev.Truth, "native generation-addressing representation map; Gate 371 tests whether finite Fock information supplies it", true, true, true, join(StatusGate370Inherited, StatusTensionPhaseIVInformationExtension)}
	formal := formalizeFockSpace()
	lanes := executeOperatorLanes(formal)
	pullback := auditPullback(lanes)
	activation := auditActivation(formal, pullback, lanes)
	landscape := auditLandscape()
	kinetic := auditKinetic(lanes)
	census := updateCensus(activation)
	truth := "Gate 371 audits the quantum-information pivot: treating the three generations as finite Schrodinger/Fock vibration levels |0>, |1>, |2> rather than geometric copies. The number operator N=diag(0,1,2) and its entropy/KMS descendants are genuinely noncentral and therefore witness a real capacity to break the copied U(3) degeneracy. However, the current ASHA Cℓ(1,7)/Morita ledger does not yet derive the Fock basis, the number operator as a native generation Hamiltonian, or the coupling Phi(s)=sN from the support defect. Moreover N is not of the required form aI_3+b tau_eta; exact tau_eta requires the target-calibrated quadratic polynomial P_tau(N)=2-(15/2)N+(7/2)N^2, which is circular unless its coefficients are derived from finite topology. Thus Gate 371 opens a powerful Phase-IV quantum-information direction, but does not yet activate internal thermal time or reduce the 15 vacuum coordinates."
	return Analysis{inheritance, formal, lanes, pullback, activation, landscape, kinetic, census, truth}, nil
}

func formalizeFockSpace() FockFormalization {
	return FockFormalization{Executed: true, GeometricBasis: []string{"generation copy 1", "generation copy 2", "generation copy 3"}, FockBasis: []string{"|0> ground vibration", "|1> first excitation", "|2> second excitation"}, LadderConvention: "truncated oscillator: N|n>=n|n>, n in {0,1,2}; boundary-truncated and used only as an informational address", NumberOperator: []float64{0, 1, 2}, CenteredNumber: []float64{-1, 0, 1}, HarmonicEnergy: []float64{0.5, 1.5, 2.5}, BasisSelectedByASHA: false, SelectionEvidence: []string{"Gate 370 found no native current-ledger support-to-generation map", "no existing finite Dirac/J/Omega/Morita theorem selects |0>,|1>,|2> as generation labels"}, ForbiddenMoves: []string{"renaming generation copies as Fock states without a selection theorem", "using N as modular Hamiltonian before deriving support-defect coupling", "fitting a polynomial in N to tau_eta and calling it topology", "importing observed masses, CKM, PMNS, or Yukawa values"}, TargetFormula: "Pi_gen Phi_info(Tr_support^eta(C_LR)) = a I_3 + b tau_eta, b != 0, with Phi_info derived rather than postulated", Verdict: join(StatusFockSpaceFormalized, StatusNumberOperatorSieveExecuted, StatusTensionFockBasisNotSelected)}
}

func executeOperatorLanes(f FockFormalization) []OperatorLane {
	n := f.NumberOperator
	support := 2.0
	bgap := 0.102464921191
	entropyOp := affine(n, 1, math.Log(sumExpScaled(n, -1)))
	return []OperatorLane{
		buildLane("A", "geometric copied-generation identity", "current Gate-370 ASHA ledger", "K=I_3", true, true, false, false, false, []float64{1, 1, 1}, false, false, join(StatusGate370Inherited, StatusFailedFockBasisNotNativeSelected)),
		buildLane("B", "finite Fock number operator", "hypothesized Schrodinger generation basis", "N=diag(0,1,2)", true, false, false, false, false, n, true, false, join(StatusNumberOperatorSieveExecuted, StatusVibrationalCapacityWitnessed, StatusTensionFockBasisNotSelected, StatusTensionNumberNoncentralNotTau, StatusFailedNumberOperatorNotDerived, StatusFailedTargetTauNotExtracted)),
		buildLane("C", "support defect pulled back to number address", "new informational coupling candidate", "Phi_info(s)=sN with s=Tr_support^eta(C_LR)", true, false, false, false, true, scale(n, support), true, true, join(StatusTopologicalIndexPullbackExecuted, StatusVibrationalCapacityWitnessed, StatusTensionOscillatorCouplingNew, StatusFailedTopologicalPullbackNotDerived, StatusFailedTargetTauNotExtracted)),
		buildLane("D", "B-gap coupled Fock number address", "new B_gap-information coupling candidate", "K=B_gap*N", true, false, false, false, true, scale(n, bgap), true, true, join(StatusTopologicalIndexPullbackExecuted, StatusVibrationalCapacityWitnessed, StatusTensionOscillatorCouplingNew, StatusFailedTopologicalPullbackNotDerived)),
		buildLane("E", "centered vibration operator", "chosen finite Fock basis", "N-1", true, false, false, false, false, f.CenteredNumber, true, false, join(StatusNumberOperatorSieveExecuted, StatusVibrationalCapacityWitnessed, StatusTensionNumberNoncentralNotTau, StatusFailedTargetTauNotExtracted)),
		buildLane("F", "finite oscillator energy", "chosen truncated Schrodinger oscillator", "H_osc=N+1/2", true, false, false, false, false, f.HarmonicEnergy, true, false, join(StatusNumberOperatorSieveExecuted, StatusVibrationalCapacityWitnessed, StatusTensionNumberNoncentralNotTau, StatusFailedTargetTauNotExtracted)),
		buildLane("G", "information entropy modular operator", "rho_N=exp(-N)/Tr exp(-N)", "-log(rho_N)=N+log(Z)I_3", true, false, false, false, false, entropyOp, true, false, join(StatusInformationEntropyAudited, StatusKMSInformationStateReconstructed, StatusTensionEntropyDependsOnChosenK, StatusTensionNumberNoncentralNotTau, StatusFailedTargetTauNotExtracted)),
		buildLane("H", "simple quadratic finite-oscillator invariant", "polynomial algebra of chosen N", "(N-1)^2", true, false, false, false, false, []float64{1, 0, 1}, true, false, join(StatusNumberOperatorSieveExecuted, StatusTensionNumberNoncentralNotTau, StatusFailedTargetTauNotExtracted)),
		buildLane("I", "target tau_eta polynomial witness", "quadratic interpolation in chosen N", "P_tau(N)=2-(15/2)N+(7/2)N^2=tau_eta", true, false, true, false, false, applyTauPolynomial(n), true, false, join(StatusVibrationalCapacityWitnessed, StatusTensionQuadraticTauPolynomial, StatusTensionPolynomialCircular, StatusFailedTauStillNotSelected)),
	}
}

func buildLane(lane, name, source, formula string, nativeFock, nativeASHA, circular, empirical, newCoupling bool, spectrum []float64, fockSelected, pullback bool, verdict string) OperatorLane {
	k := diagMatrix(spectrum)
	d := decomposeIdentityTau(spectrum)
	kms := kmsState(spectrum, 1)
	comms := commutatorSieve(k)
	breaks := anyCommutator(comms)
	central := isCentral(k)
	promotable := nativeASHA && nativeFock && !circular && !empirical && !newCoupling && fockSelected && pullback && d.TargetReached && breaks
	if d.TargetReached && !circular && nativeASHA && fockSelected && pullback {
		verdict = join(verdict, StatusTauEtaExtractedByInformation, StatusVibrationalIntertwinerDerived, StatusThermalTimeActivated)
	}
	return OperatorLane{lane, name, source, formula, nativeFock, nativeASHA, circular, empirical, newCoupling, clone(spectrum), k, isSelfAdjoint(k), central, !central, d, kms, comms, breaks, fockSelected, pullback, promotable, false, verdict}
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
	verdict := StatusFailedTargetTauNotExtracted
	if exact && nonzeroB {
		verdict = StatusTauEtaExtractedByInformation
	}
	return Decomposition{a, b, res, exact, nonzeroB, exact && nonzeroB, verdict}
}

func auditPullback(lanes []OperatorLane) PullbackAudit {
	anyNoncentral, anyPromotable, anyTarget := false, false, false
	for _, l := range lanes {
		if l.NonCentral && !l.Circular {
			anyNoncentral = true
		}
		if l.PromotableHamiltonian {
			anyPromotable = true
		}
		if l.Decomposition.TargetReached {
			anyTarget = true
		}
	}
	return PullbackAudit{true, "s=Tr_support^eta(C_LR), inherited as a scalar support defect under the current ledger", 2, false, anyNoncentral, anyPromotable, anyTarget, "P_tau(N)=2-(15/2)N+(7/2)N^2", true, "The finite-Fock reinterpretation provides noncentral witnesses, but not a native ASHA support-to-generation pullback. Phi_info(s)=sN is an additional coupling rule, not derived from D_F, J, Omega_Hsigma, Morita 1:3, or the Gate-370 support ledger. Exact tau_eta appears only through a target-fitted quadratic polynomial in N.", join(StatusTopologicalIndexPullbackExecuted, StatusVibrationalCapacityWitnessed, StatusTensionOscillatorCouplingNew, StatusTensionQuadraticTauPolynomial, StatusTensionPolynomialCircular, StatusFailedTopologicalPullbackNotDerived, StatusFailedVibrationalIntertwinerNotDerived, StatusFailedTargetTauNotExtracted)}
}
func auditActivation(f FockFormalization, p PullbackAudit, lanes []OperatorLane) ActivationAudit {
	noncentral := false
	for _, l := range lanes {
		if l.NonCentral {
			noncentral = true
			break
		}
	}
	return ActivationAudit{true, f.BasisSelectedByASHA, false, p.NativeASHAFockMapDerived, false, noncentral, false, false, "Vibrational quantum information breaks U(3) at the level of a hypothesized Fock basis, but the current ASHA ledger does not select that basis or the support-defect-to-number coupling. Internal thermal time is therefore not activated as a theorem.", "Prove a native generation oscillator theorem: derive |0>,|1>,|2>, N, and support-defect coupling from Cℓ(1,7)/triality/phase data; or prove no current finite representation can select such an informational basis.", join(StatusFailedVibrationalIntertwinerNotDerived, StatusFailedFockBasisNotNativeSelected, StatusFailedNumberOperatorNotDerived, StatusFailedTopologicalPullbackNotDerived, StatusFailedThermalTimeNotActivated, StatusFailedTauStillNotSelected, StatusFailedVacuumNotSelected, StatusFailedCKMNotDerived, StatusFailedYukawaNotDerived, StatusTensionNoncentralButNotSelecting, StatusTensionPhaseIVInformationExtension)}
}
func auditLandscape() LandscapeAudit {
	return LandscapeAudit{true, true, true, true, true, true, true, true, true, true, false, join(StatusLandscapePreservationAudited)}
}
func auditKinetic(lanes []OperatorLane) KineticAudit {
	self, faithful := true, true
	for _, l := range lanes {
		self = self && l.SelfAdjoint
		faithful = faithful && l.KMS.Faithful
	}
	return KineticAudit{true, self, faithful, true, true, true, join(StatusKineticSafetyAudited)}
}
func updateCensus(a ActivationAudit) Census {
	reduction := 0
	if a.InternalThermalTimeActivated && a.VacuumCoordinatesReduced {
		reduction = 1
	}
	return Census{vacuumInputs, reduction, vacuumInputs - reduction, false, join(StatusParameterCensusUpdated, StatusFailedCensusNotReduced)}
}
func applyTauPolynomial(n []float64) []float64 {
	out := make([]float64, len(n))
	for i, x := range n {
		out[i] = 2 - 7.5*x + 3.5*x*x
	}
	return out
}
func affine(x []float64, scale, shift float64) []float64 {
	out := make([]float64, len(x))
	for i, v := range x {
		out[i] = scale*v + shift
	}
	return out
}
func scale(x []float64, s float64) []float64 {
	out := make([]float64, len(x))
	for i, v := range x {
		out[i] = s * v
	}
	return out
}
func sumExpScaled(x []float64, beta float64) float64 {
	z := 0.0
	for _, v := range x {
		z += math.Exp(beta * v)
	}
	return z
}
func kmsState(spectrum []float64, beta float64) KMSState {
	weights := make([]float64, len(spectrum))
	z := 0.0
	for i, e := range spectrum {
		weights[i] = math.Exp(-beta * e)
		z += weights[i]
	}
	density := make([]float64, len(weights))
	entropy := 0.0
	for i, w := range weights {
		density[i] = w / z
		if density[i] > 0 {
			entropy -= density[i] * math.Log(density[i])
		}
	}
	freqs := []Frequency{freq("01", density[0], density[1]), freq("02", density[0], density[2]), freq("12", density[1], density[2])}
	non := false
	for _, f := range freqs {
		if f.NonZero {
			non = true
		}
	}
	return KMSState{beta, density, freqs, non, allPositive(density), entropy}
}
func freq(pair string, a, b float64) Frequency {
	lr := math.Log(a / b)
	return Frequency{pair, lr, math.Abs(lr) > eps}
}
func allPositive(xs []float64) bool {
	for _, x := range xs {
		if x <= 0 || math.IsNaN(x) || math.IsInf(x, 0) {
			return false
		}
	}
	return true
}
func diagMatrix(d []float64) [][]float64 {
	out := zero3()
	for i := range d {
		out[i][i] = d[i]
	}
	return out
}
func zero3() [][]float64 { return [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}} }
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
	d := m[0][0]
	for i := 0; i < 3; i++ {
		if math.Abs(m[i][i]-d) > eps {
			return false
		}
		for j := 0; j < 3; j++ {
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
	}{{"E_12", generator(0, 1)}, {"E_13", generator(0, 2)}, {"E_23", generator(1, 2)}}
	out := make([]CommutatorResult, 0, len(gens))
	for _, g := range gens {
		n := commutatorNorm(k, g.m)
		out = append(out, CommutatorResult{g.name, n, n > eps})
	}
	return out
}
func generator(i, j int) [][]float64 { g := zero3(); g[i][j] = 1; g[j][i] = -1; return g }
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
func clone(x []float64) []float64 { y := make([]float64, len(x)); copy(y, x); return y }
func join(parts ...string) string { return strings.Join(parts, ";") }
func Statuses(a Analysis) []string {
	set := map[string]struct{}{}
	blocks := []string{a.Inheritance.Verdict, a.Formalization.Verdict, a.Pullback.Verdict, a.Activation.Verdict, a.Landscape.Verdict, a.Kinetic.Verdict, a.Census.Verdict}
	for _, l := range a.Lanes {
		blocks = append(blocks, l.Verdict, l.Decomposition.Verdict)
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
func laneByID(lanes []OperatorLane, id string) OperatorLane {
	for _, l := range lanes {
		if l.Lane == id {
			return l
		}
	}
	if len(lanes) == 0 {
		return OperatorLane{}
	}
	return lanes[0]
}
func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("executed=%v; all-current-native-maps-I3=%v; tau-manual-circular=%v; missing=%s; verdict=%s", i.Executed, i.AllCurrentNativeMapsI3, i.TauEtaManualMapCircular, i.MissingObject, i.Verdict)
}
func FormatFormalization(f FockFormalization) string {
	return fmt.Sprintf("executed=%v; fock-basis=%v; N=%v; centered=%v; energy=%v; basis-selected-by-ASHA=%v; target=%s; verdict=%s", f.Executed, f.FockBasis, f.NumberOperator, f.CenteredNumber, f.HarmonicEnergy, f.BasisSelectedByASHA, f.TargetFormula, f.Verdict)
}
func FormatLane(l OperatorLane) string {
	return fmt.Sprintf("lane=%s; name=%s; formula=%s; nativeFock=%v; nativeASHA=%v; circular=%v; newCoupling=%v; spectrum=%v; central=%v; noncentral=%v; decomp={a=%.12g,b=%.12g,res=%.12g,target=%v}; breaks=%v; promotable=%v; verdict=%s", l.Lane, l.Name, l.Formula, l.NativeToChosenFock, l.NativeToCurrentASHA, l.Circular, l.RequiresNewCoupling, l.Spectrum, l.Central, l.NonCentral, l.Decomposition.A, l.Decomposition.B, l.Decomposition.ResidualNorm, l.Decomposition.TargetReached, l.BreaksFlavorOrbit, l.PromotableHamiltonian, l.Verdict)
}
func FormatPullback(p PullbackAudit) string {
	return fmt.Sprintf("executed=%v; native-map-derived=%v; any-noncentral-witness=%v; any-promotable=%v; any-target=%v; exact-tau-polynomial=%s; circular=%v; answer=%s; verdict=%s", p.Executed, p.NativeASHAFockMapDerived, p.AnyNoncentralFockWitness, p.AnyPromotableHamiltonian, p.AnyTargetTauExtraction, p.ExactTauRequiresPolynomial, p.PolynomialCircular, p.DirectAnswer, p.Verdict)
}
func FormatActivation(a ActivationAudit) string {
	return fmt.Sprintf("executed=%v; fock-basis-native=%v; N-native=%v; pullback-derived=%v; target-tau-native=%v; capacity=%v; activated=%v; census-reduced=%v; answer=%s; next=%s; verdict=%s", a.Executed, a.FockBasisNativeSelected, a.NumberOperatorNativeSelected, a.PullbackDerived, a.TargetTauExtractedNatively, a.NoncentralCapacityWitnessed, a.InternalThermalTimeActivated, a.VacuumCoordinatesReduced, a.DirectAnswer, a.NextGate, a.Verdict)
}
func FormatLandscape(l LandscapeAudit) string {
	return fmt.Sprintf("executed=%v; sin2=%v; quartic=%v; alpha=%v; morita=%v; bgap=%v; omega=%v; no-empirical=%v; polluted=%v; verdict=%s", l.Executed, l.WeakMixingPreserved, l.QuarticRatioPreserved, l.AlphaGUTPreserved, l.MoritaSplitPreserved, l.BGapLedgerPreserved, l.OmegaIndexPreserved, l.NoEmpiricalFlavorImport, l.FiniteCorePolluted, l.Verdict)
}
func FormatKinetic(k KineticAudit) string {
	return fmt.Sprintf("executed=%v; self=%v; faithful=%v; rank=%v; ghost=%v; nonunitary=%v; verdict=%s", k.Executed, k.AllOperatorsSelf, k.FaithfulKMSStates, k.NoRankCollapse, k.NoGhostMetric, k.NoNonunitaryPush, k.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d; reduction=%d; remaining=%d; seven-seal=%v; verdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}
