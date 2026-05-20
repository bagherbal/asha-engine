// Package generation2chistorynativesourceindependenceandtransportlawaudit implements
// Gate 781: C_History Native-Source Independence and Transport-Law Audit.
//
// Gate 780 showed that C_Higgs is not yet a full independent prediction because
// the dominant correction factor C_History contains bridge-layer and historically
// runtime-tied ingredients. Gate 781 audits the whole History correction cluster
// at once: L_Hopf source typing, the transport law form, and the runtime
// independence of kappa_lambda_red. This is a History correction independence
// audit only. It does not derive the VEV, G_F, Higgs pole mass, Yukawa
// operators, CKM/PMNS, flavor hierarchy, or a native electroweak scale theorem.
package generation2chistorynativesourceindependenceandtransportlawaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE781-C-HISTORY-NATIVE-SOURCE-INDEPENDENCE-TRANSPORT-LAW-AUDIT"

	StatusGate780PredictionIndependenceAuditInherited     = "PASS_GATE780_PREDICTION_INDEPENDENCE_AUDIT_INHERITED"
	StatusCHistoryDependencyClusterExpanded               = "PASS_C_HISTORY_DEPENDENCY_CLUSTER_EXPANDED"
	StatusLHopfRadialHessianHopfSourceAudited             = "PASS_L_HOPF_RADIAL_HESSIAN_HOPF_SOURCE_AUDITED"
	StatusTransportLawFormAudited                         = "PASS_TRANSPORT_LAW_FORM_AUDITED"
	StatusKappaLambdaRedRuntimeIndependenceAudited        = "PASS_KAPPA_LAMBDA_RED_RUNTIME_INDEPENDENCE_AUDITED"
	StatusBranchOutcomesRecorded                          = "PASS_BRANCH_OUTCOMES_RECORDED"
	StatusPhysicalFirewallsEnforced                       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusCHistoryDominantDimensionlessCorrectionTarget   = "CONDITIONAL_SUPPORT_C_HISTORY_IS_DOMINANT_DIMENSIONLESS_CORRECTION_TARGET"
	StatusLHopfStrongRadialHessianHopfSourceType          = "CONDITIONAL_SUPPORT_L_HOPF_HAS_STRONG_RADIAL_HESSIAN_HOPF_SOURCE_TYPE"
	StatusTransportBracketScalarMatchingComplement        = "CONDITIONAL_SUPPORT_TRANSPORT_BRACKET_IS_SCALAR_MATCHING_COMPLEMENT_CANDIDATE"
	StatusNoNativeHistoryLoopUnitTheorem                  = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoNativeTransportLawTheorem                     = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_LAW_THEOREM"
	StatusKappaLambdaRedNotYetNativeScalarMatchingTheorem = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_YET_NATIVE_SCALAR_MATCHING_THEOREM"
	StatusCHistoryNotYetFullIndependentPrediction         = "FAILED_ROUTE_C_HISTORY_NOT_YET_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusTreeProxyNotPoleMass                            = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem             = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate781CHistoryNativeSourceBoundary             = "FIREWALL_PRESERVED_GATE781_C_HISTORY_NATIVE_SOURCE_BOUNDARY"
)

const (
	// Snapshot ledger inherited from Gates 759, 776, and 780. Gate 781 keeps
	// these local to avoid importing deep predecessor chains while auditing the
	// source-dependency cluster rather than recomputing previous gates.
	cHistorySnapshot       = 1.038025177923625
	lHopfSnapshot          = 0.039788735772973836
	kappaLambdaRedSnapshot = 0.04432304306956136
	deltaHistorySnapshot   = 0.03802517792362492
	epsilonYukawaSnapshot  = 0.0007751811187991509
	cHiggsSnapshot         = 1.0372205204048603
)

type Gate780Inheritance struct {
	Inherited          bool
	CHiggsFormula      string
	CHistoryFormula    string
	PredictionStatus   string
	CHistory           float64
	DeltaHistory       float64
	EpsilonYukawa      float64
	DominantCorrection bool
	Verdict            string
}

type CHistoryCluster struct {
	Expanded                 bool
	Formula                  string
	LHopf                    float64
	KappaLambdaRed           float64
	Complement               float64
	ComputedCHistory         float64
	MatchesSnapshot          bool
	KappaLambdaRedFormula    string
	KappaLambdaRedComponents []string
	ClusterQuestion          string
	Verdict                  string
}

type LHopfSourceAudit struct {
	Audited                         bool
	Formula                         string
	MaximumEntropyStateRequired     bool
	HessianSupportProjectorRequired bool
	PhaseLoopPayoffRequired         bool
	HistoryEvaluatesEventTheorem    bool
	NativeTheorem                   bool
	ConditionalSourceTyping         bool
	EventWeight                     float64
	PhasePayoff                     float64
	LHopf                           float64
	MissingIngredients              []string
	Verdict                         string
}

type TransportLawAudit struct {
	Audited                bool
	TransportFormula       string
	BaselineInterpretation string
	LHopfInterpretation    string
	BracketFormula         string
	BracketInterpretation  string
	NativeTransportLaw     bool
	BridgeReconstruction   bool
	Verdict                string
}

type KappaLambdaRuntimeIndependenceAudit struct {
	Audited                             bool
	KappaLambdaRedFormula               string
	BoundaryWallCoordinate              string
	BoundaryResponsePolynomial          string
	FlavorWallReducedInput              string
	UsesLambdaRuntimeTarget             bool
	UsesTreeMassTarget                  bool
	UsesPoleMassTarget                  bool
	UsesHiggsTargetClosure              bool
	CanBeEvaluatedWithoutRuntimeClosure bool
	ReducedButNotRuntimeIndependent     bool
	IndependentComponents               []string
	RuntimeDependentOrSealedComponents  []string
	Verdict                             string
}

type BranchOutcomes struct {
	Recorded        bool
	StrongSuccess   string
	PartialSuccess  string
	Failure         string
	SelectedOutcome string
	NextGate        string
	Reason          string
	Verdict         string
}

type PhysicalFirewalls struct {
	Enforced                          bool
	CHistoryNativePredictionComponent bool
	LHopfNativeHistoryLoopTheorem     bool
	TransportLawNativeTheorem         bool
	KappaLambdaNativeScalarTheorem    bool
	TreeProxyPoleMass                 bool
	YukawaNativeTheorem               bool
	Verdict                           string
}

type Analysis struct {
	Gate780   Gate780Inheritance
	Cluster   CHistoryCluster
	LHopf     LHopfSourceAudit
	Transport TransportLawAudit
	Runtime   KappaLambdaRuntimeIndependenceAudit
	Branches  BranchOutcomes
	Firewalls PhysicalFirewalls
	Truth     string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		return cloneAnalysis(cache), nil
	}
	if !finite(cHistorySnapshot, lHopfSnapshot, kappaLambdaRedSnapshot, deltaHistorySnapshot, epsilonYukawaSnapshot, cHiggsSnapshot) {
		return nil, fmt.Errorf("invalid Gate781 snapshot ledger")
	}
	if cHistorySnapshot <= 0 || lHopfSnapshot <= 0 || kappaLambdaRedSnapshot < 0 || cHiggsSnapshot <= 0 {
		return nil, fmt.Errorf("non-positive Gate781 snapshot ledger")
	}
	complement := 1 - kappaLambdaRedSnapshot
	computedHistory := 1 + lHopfSnapshot*complement
	matchesHistory := closeRel(computedHistory, cHistorySnapshot, 1e-15)
	dominant := deltaHistorySnapshot > 10*epsilonYukawaSnapshot
	a := &Analysis{
		Gate780: Gate780Inheritance{
			Inherited:          true,
			CHiggsFormula:      "C_Higgs=C_Yukawa C_History",
			CHistoryFormula:    "C_History=1+L_Hopf(1-kappa_lambda_red)",
			PredictionStatus:   "Level A/B: current C_Higgs is a bridge consistency and semi-independent target, not a full independent prediction",
			CHistory:           cHistorySnapshot,
			DeltaHistory:       deltaHistorySnapshot,
			EpsilonYukawa:      epsilonYukawaSnapshot,
			DominantCorrection: dominant,
			Verdict:            StatusGate780PredictionIndependenceAuditInherited,
		},
		Cluster: CHistoryCluster{
			Expanded:                 true,
			Formula:                  "C_History=1+L_Hopf(1-kappa_lambda_red)",
			LHopf:                    lHopfSnapshot,
			KappaLambdaRed:           kappaLambdaRedSnapshot,
			Complement:               complement,
			ComputedCHistory:         computedHistory,
			MatchesSnapshot:          matchesHistory,
			KappaLambdaRedFormula:    "kappa_lambda_red=|lambda(Lambda12)|+F_wall_3_red(s)-kappa_e_red",
			KappaLambdaRedComponents: []string{"|lambda(Lambda12)| boundary scalar wall coordinate", "F_wall_3_red(s) cubic boundary response polynomial", "kappa_e_red reduced flavor-wall input"},
			ClusterQuestion:          "Can C_History be sourced without scalar-runtime target closure?",
			Verdict:                  StatusCHistoryDependencyClusterExpanded,
		},
		LHopf: LHopfSourceAudit{
			Audited:                         true,
			Formula:                         "L_Hopf=Tr_K7+(rho_plus[(1/(2*pi))supp(H_V(x0))])=(1/(2*pi))(1/4)=1/(8*pi)",
			MaximumEntropyStateRequired:     true,
			HessianSupportProjectorRequired: true,
			PhaseLoopPayoffRequired:         true,
			HistoryEvaluatesEventTheorem:    false,
			NativeTheorem:                   false,
			ConditionalSourceTyping:         true,
			EventWeight:                     0.25,
			PhasePayoff:                     1 / (2 * math.Pi),
			LHopf:                           lHopfSnapshot,
			MissingIngredients:              []string{"native reason History transport evaluates the Hessian support event", "native HistoryLoopUnit theorem", "native transport law from radial-Hessian event to scalar runtime correction"},
			Verdict:                         StatusLHopfRadialHessianHopfSourceAudited,
		},
		Transport: TransportLawAudit{
			Audited:                true,
			TransportFormula:       "C_History=1+L_Hopf(1-kappa_lambda_red)",
			BaselineInterpretation: "1 is the normalized untransported scalar baseline",
			LHopfInterpretation:    "L_Hopf is the radial-Hessian Hopf event unit",
			BracketFormula:         "1-kappa_lambda_red",
			BracketInterpretation:  "candidate scalar matching complement reconstructed from boundary/flavor wall data",
			NativeTransportLaw:     false,
			BridgeReconstruction:   true,
			Verdict:                StatusTransportLawFormAudited,
		},
		Runtime: KappaLambdaRuntimeIndependenceAudit{
			Audited:                             true,
			KappaLambdaRedFormula:               "kappa_lambda_red=|lambda(Lambda12)|+F_wall_3_red(s)-kappa_e_red",
			BoundaryWallCoordinate:              "|lambda(Lambda12)| is a boundary scalar wall coordinate, not a native scalar theorem",
			BoundaryResponsePolynomial:          "F_wall_3_red(s) is a cubic boundary response polynomial with prior deficit-closure history",
			FlavorWallReducedInput:              "kappa_e_red is a reduced flavor-wall input still depending on flavor/orientation data",
			UsesLambdaRuntimeTarget:             true,
			UsesTreeMassTarget:                  false,
			UsesPoleMassTarget:                  false,
			UsesHiggsTargetClosure:              true,
			CanBeEvaluatedWithoutRuntimeClosure: false,
			ReducedButNotRuntimeIndependent:     true,
			IndependentComponents:               []string{"formal complement rewrite", "radial-Hopf source typing of L_Hopf", "boundary/flavor decomposition syntax once inputs are supplied"},
			RuntimeDependentOrSealedComponents:  []string{"kappa_lambda_red scalar matching deficit", "F_wall_3_red deficit-closure compression history", "kappa_e_red flavor/orientation data", "lambda(Lambda12) boundary scalar coordinate"},
			Verdict:                             StatusKappaLambdaRedRuntimeIndependenceAudited,
		},
		Branches: BranchOutcomes{
			Recorded:        true,
			StrongSuccess:   "L_Hopf, transport law, and kappa_lambda_red are all runtime-independent; C_History becomes an independent dimensionless prediction component; next Gate 782 targets C_Yukawa native participation.",
			PartialSuccess:  "L_Hopf is strongly source-typed but kappa_lambda_red remains bridge-dependent; scalar matching complement is the bottleneck; next Gate 782 targets boundary-flavor scalar matching complement independence.",
			Failure:         "transport form remains target-defined or circular; C_History is bridge consistency only; next Gate 782 rebuilds boundary raw-moment response.",
			SelectedOutcome: "Outcome 2 — partial success",
			NextGate:        "Gate 782 — Boundary-Flavor Scalar Matching Complement Independence Audit",
			Reason:          "L_Hopf has strong radial-Hessian Hopf source typing, but the transport law is not native and kappa_lambda_red is reduced but not runtime-independent.",
			Verdict:         StatusBranchOutcomesRecorded,
		},
		Firewalls: PhysicalFirewalls{
			Enforced:                          true,
			CHistoryNativePredictionComponent: false,
			LHopfNativeHistoryLoopTheorem:     false,
			TransportLawNativeTheorem:         false,
			KappaLambdaNativeScalarTheorem:    false,
			TreeProxyPoleMass:                 false,
			YukawaNativeTheorem:               false,
			Verdict:                           StatusGate781CHistoryNativeSourceBoundary,
		},
		Truth: fmt.Sprintf("Gate781 audits the dominant History correction C_History=%.16f as a macro-gate cluster. L_Hopf has strong radial-Hessian Hopf source typing, but no native HistoryLoop transport theorem or runtime-independent kappa_lambda_red theorem is certified; selected branch is Outcome 2, targeting boundary-flavor scalar matching complement next.", cHistorySnapshot),
	}
	cache = a
	return cloneAnalysis(a), nil
}

func cloneAnalysis(a *Analysis) *Analysis {
	clone := *a
	clone.Cluster.KappaLambdaRedComponents = append([]string(nil), a.Cluster.KappaLambdaRedComponents...)
	clone.LHopf.MissingIngredients = append([]string(nil), a.LHopf.MissingIngredients...)
	clone.Runtime.IndependentComponents = append([]string(nil), a.Runtime.IndependentComponents...)
	clone.Runtime.RuntimeDependentOrSealedComponents = append([]string(nil), a.Runtime.RuntimeDependentOrSealedComponents...)
	return &clone
}

func finite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if want == 0 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}

func Statuses() []string {
	return []string{
		StatusGate780PredictionIndependenceAuditInherited,
		StatusCHistoryDependencyClusterExpanded,
		StatusLHopfRadialHessianHopfSourceAudited,
		StatusTransportLawFormAudited,
		StatusKappaLambdaRedRuntimeIndependenceAudited,
		StatusBranchOutcomesRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusCHistoryDominantDimensionlessCorrectionTarget,
		StatusLHopfStrongRadialHessianHopfSourceType,
		StatusTransportBracketScalarMatchingComplement,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoNativeTransportLawTheorem,
		StatusKappaLambdaRedNotYetNativeScalarMatchingTheorem,
		StatusCHistoryNotYetFullIndependentPrediction,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate781CHistoryNativeSourceBoundary,
	}
}

func FormatGate780(g Gate780Inheritance) string {
	return fmt.Sprintf("inherited=%t CFormula=%s historyFormula=%s status=%s CH=%.17g deltaH=%.17g epsY=%.17g dominant=%t verdict=%s", g.Inherited, g.CHiggsFormula, g.CHistoryFormula, g.PredictionStatus, g.CHistory, g.DeltaHistory, g.EpsilonYukawa, g.DominantCorrection, g.Verdict)
}

func FormatCluster(c CHistoryCluster) string {
	return fmt.Sprintf("expanded=%t formula=%s L=%.17g kappa=%.17g complement=%.17g computed=%.17g matches=%t kappaFormula=%s components=%s question=%s verdict=%s", c.Expanded, c.Formula, c.LHopf, c.KappaLambdaRed, c.Complement, c.ComputedCHistory, c.MatchesSnapshot, c.KappaLambdaRedFormula, strings.Join(c.KappaLambdaRedComponents, ";"), c.ClusterQuestion, c.Verdict)
}

func FormatLHopf(l LHopfSourceAudit) string {
	return fmt.Sprintf("audited=%t formula=%s rho=%t hessian=%t phase=%t historyTheorem=%t native=%t conditional=%t weight=%.17g payoff=%.17g L=%.17g missing=%d verdict=%s", l.Audited, l.Formula, l.MaximumEntropyStateRequired, l.HessianSupportProjectorRequired, l.PhaseLoopPayoffRequired, l.HistoryEvaluatesEventTheorem, l.NativeTheorem, l.ConditionalSourceTyping, l.EventWeight, l.PhasePayoff, l.LHopf, len(l.MissingIngredients), l.Verdict)
}

func FormatTransport(t TransportLawAudit) string {
	return fmt.Sprintf("audited=%t formula=%s baseline=%s L=%s bracket=%s interp=%s native=%t bridge=%t verdict=%s", t.Audited, t.TransportFormula, t.BaselineInterpretation, t.LHopfInterpretation, t.BracketFormula, t.BracketInterpretation, t.NativeTransportLaw, t.BridgeReconstruction, t.Verdict)
}

func FormatRuntime(r KappaLambdaRuntimeIndependenceAudit) string {
	return fmt.Sprintf("audited=%t formula=%s wall=%s response=%s flavor=%s lambdaTarget=%t treeTarget=%t poleTarget=%t higgsClosure=%t runtimeFree=%t reduced=%t independent=%d dependent=%d verdict=%s", r.Audited, r.KappaLambdaRedFormula, r.BoundaryWallCoordinate, r.BoundaryResponsePolynomial, r.FlavorWallReducedInput, r.UsesLambdaRuntimeTarget, r.UsesTreeMassTarget, r.UsesPoleMassTarget, r.UsesHiggsTargetClosure, r.CanBeEvaluatedWithoutRuntimeClosure, r.ReducedButNotRuntimeIndependent, len(r.IndependentComponents), len(r.RuntimeDependentOrSealedComponents), r.Verdict)
}

func FormatBranches(b BranchOutcomes) string {
	return fmt.Sprintf("recorded=%t strong=%s partial=%s failure=%s selected=%s next=%s reason=%s verdict=%s", b.Recorded, b.StrongSuccess, b.PartialSuccess, b.Failure, b.SelectedOutcome, b.NextGate, b.Reason, b.Verdict)
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("enforced=%t Cnative=%t Lnative=%t transportNative=%t kappaNative=%t treePole=%t yukawaNative=%t verdict=%s", f.Enforced, f.CHistoryNativePredictionComponent, f.LHopfNativeHistoryLoopTheorem, f.TransportLawNativeTheorem, f.KappaLambdaNativeScalarTheorem, f.TreeProxyPoleMass, f.YukawaNativeTheorem, f.Verdict)
}

func containsAll(haystack, needles []string) bool {
	joined := "\x00" + strings.Join(haystack, "\x00") + "\x00"
	for _, n := range needles {
		if !strings.Contains(joined, "\x00"+n+"\x00") {
			return false
		}
	}
	return true
}
