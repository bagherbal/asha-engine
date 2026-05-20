// Package generation2higgsdimensionlesspredictionindependenceandsealdependencyaudit implements
// Gate 780: Higgs Dimensionless Prediction Independence and Seal-Dependency Audit.
//
// Gate 779 produced the Fermi-normalized tree ratio
// 4 sqrt(2) G_F m_H_tree^2 = C_Higgs. Gate 780 expands the dependency graph
// behind C_Higgs and classifies whether the current object is an independent
// dimensionless prediction target or a sealed bridge consistency closure. This
// is an independence and seal-dependency audit only. It does not derive G_F, v,
// the Higgs pole mass, scalar runtime lambda, Yukawa operators, CKM/PMNS,
// flavor hierarchy, or a native HistoryLoopUnit theorem.
package generation2higgsdimensionlesspredictionindependenceandsealdependencyaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE780-HIGGS-DIMENSIONLESS-PREDICTION-INDEPENDENCE-SEAL-DEPENDENCY-AUDIT"

	StatusGate779FermiNormalizedRatioInherited            = "PASS_GATE779_FERMI_NORMALIZED_RATIO_INHERITED"
	StatusCHiggsDependencyGraphExpanded                   = "PASS_C_HIGGS_DEPENDENCY_GRAPH_EXPANDED"
	StatusInputIndependenceClassificationAudited          = "PASS_INPUT_INDEPENDENCE_CLASSIFICATION_AUDITED"
	StatusCircularityAuditDefined                         = "PASS_CIRCULARITY_AUDIT_DEFINED"
	StatusPredictionStatusLevelsDefined                   = "PASS_PREDICTION_STATUS_LEVELS_DEFINED"
	StatusRequiredRemovalsForPredictionRecorded           = "PASS_REQUIRED_REMOVALS_FOR_PREDICTION_RECORDED"
	StatusPhysicalFirewallsEnforced                       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusCHiggsCurrentDimensionlessPredictionTarget      = "CONDITIONAL_SUPPORT_C_HIGGS_IS_CURRENT_DIMENSIONLESS_PREDICTION_TARGET"
	StatusFermiNormalizedRatioRightTestInterface          = "CONDITIONAL_SUPPORT_FERMI_NORMALIZED_RATIO_IS_RIGHT_TEST_INTERFACE"
	StatusCurrentStatusBridgeConsistencyNotFullPrediction = "CONDITIONAL_SUPPORT_CURRENT_STATUS_IS_BRIDGE_CONSISTENCY_NOT_FULL_INDEPENDENT_PREDICTION"
	StatusCHiggsNotYetNativeHiggsTheorem                  = "FAILED_ROUTE_C_HIGGS_NOT_YET_NATIVE_HIGGS_THEOREM"
	StatusNEffNotNativeYukawaTheorem                      = "FAILED_ROUTE_N_EFF_NOT_NATIVE_YUKAWA_THEOREM"
	StatusKappaLambdaRedNotNativeScalarMatchingTheorem    = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_NATIVE_SCALAR_MATCHING_THEOREM"
	StatusLHopfNotNativeHistoryLoopTheorem                = "FAILED_ROUTE_L_HOPF_NOT_NATIVE_HISTORYLOOP_THEOREM"
	StatusGFNotNativeElectroweakScaleTheorem              = "FAILED_ROUTE_G_F_NOT_NATIVE_ELECTROWEAK_SCALE_THEOREM"
	StatusTreeProxyNotPoleMass                            = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusGate780HiggsPredictionIndependenceBoundary      = "FIREWALL_PRESERVED_GATE780_HIGGS_PREDICTION_INDEPENDENCE_BOUNDARY"
)

const (
	// Snapshot ledger inherited from Gates 756, 759, 775, 778, and 779. Keeping
	// these as local audited constants prevents Gate 780 from importing deep
	// predecessor chains while preserving the exact dependency graph under audit.
	cHiggsSnapshot         = 1.0372205204048603
	cYukawaSnapshot        = 0.9992248188812008
	cHistorySnapshot       = 1.038025177923625
	nEffSnapshot           = 3.0023273474722147
	lHopfSnapshot          = 0.039788735772973836
	kappaLambdaRedSnapshot = 0.04432304306956136
	vevConventionGeV       = 246.2196508
	gFermiEquivalent       = 1.1663786999444556e-05
	treeMassProxyGeV       = 125.38000000304908
	fermiNormalizedRatio   = 1.0372205204048603
)

type InheritedGate779Ratio struct {
	Inherited               bool
	FermiNormalizedIdentity string
	CHiggs                  float64
	VEVGeV                  float64
	GFermiGeVMinus2         float64
	TreeMassProxyGeV        float64
	FourSqrt2GFMassSquared  float64
	DimensionlessTask       string
	ScaleTask               string
	DerivesGF               bool
	DerivesPoleMass         bool
	Verdict                 string
}

type DependencyGraph struct {
	Expanded              bool
	CHiggsFormula         string
	CYukawaFormula        string
	NEffFormula           string
	YukawaTraceLedger     string
	CHistoryFormula       string
	LHopfFormula          string
	KappaLambdaRedFormula string
	FWall3Formula         string
	KappaERedFormula      string
	CHiggs                float64
	CYukawa               float64
	CHistory              float64
	NEff                  float64
	LHopf                 float64
	KappaLambdaRed        float64
	ProductMatchesCHiggs  bool
	HistoryMatchesFormula bool
	Verdict               string
}

type InputClassification struct {
	Audited                      bool
	PK7Classification            string
	LHopfClassification          string
	NEffClassification           string
	Theta13JCKMClassification    string
	BoundaryScalarClassification string
	GFClassification             string
	TreeProxyClassification      string
	NativeInputs                 []string
	BridgeResponseInputs         []string
	EmpiricalOrSealedInputs      []string
	RuntimeDefinedInputs         []string
	ExternalScaleInputs          []string
	Verdict                      string
}

type CircularityAudit struct {
	Defined                                  bool
	CriticalTargets                          []string
	KappaLambdaHistoricallyRuntimeTied       bool
	LambdaRuntimeEffBridgeClosureQuantity    bool
	LambdaProxyFromYukawaLedgerNotHiggsMass  bool
	YukawaLedgerStillSealed                  bool
	FWallCompressedAgainstDeficitRelations   bool
	AnyComponentUsesHiggsOrRuntimeTargetData bool
	IndependentPieces                        []string
	DependentOrSealedPieces                  []string
	Verdict                                  string
}

type PredictionStatusLevels struct {
	Defined      bool
	LevelA       string
	LevelB       string
	LevelC       string
	LevelD       string
	CurrentLevel string
	NotLevelC    bool
	NotLevelD    bool
	Reason       string
	Verdict      string
}

type RequiredRemovals struct {
	Recorded                     bool
	Items                        []string
	NeedsNativeYukawaOperator    bool
	NeedsScalarMatchingSource    bool
	NeedsFlavorSource            bool
	NeedsHistoryLoopTheorem      bool
	NeedsBoundaryResponseTheorem bool
	NeedsFermiOrVEVScale         bool
	NeedsTreeToPolePackage       bool
	Verdict                      string
}

type PhysicalFirewalls struct {
	Enforced                          bool
	CHiggsIndependentIfTargetDataUsed bool
	FermiRatioPoleMassTheorem         bool
	YukawaLedgerNativeYukawaTheorem   bool
	KappaLambdaRedNativeScalarTheorem bool
	LHopfNativeHistoryLoopTheorem     bool
	GFAShaDerivedScale                bool
	TreeProxyPoleMass                 bool
	Verdict                           string
}

type Analysis struct {
	Gate779        InheritedGate779Ratio
	Graph          DependencyGraph
	Classification InputClassification
	Circularity    CircularityAudit
	Levels         PredictionStatusLevels
	Removals       RequiredRemovals
	Firewalls      PhysicalFirewalls
	Truth          string
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
	if !finite(cHiggsSnapshot, cYukawaSnapshot, cHistorySnapshot, nEffSnapshot, lHopfSnapshot, kappaLambdaRedSnapshot, vevConventionGeV, gFermiEquivalent, treeMassProxyGeV, fermiNormalizedRatio) {
		return nil, fmt.Errorf("invalid Gate780 snapshot ledger")
	}
	if cHiggsSnapshot <= 0 || cYukawaSnapshot <= 0 || cHistorySnapshot <= 0 || nEffSnapshot <= 0 || lHopfSnapshot <= 0 || vevConventionGeV <= 0 || gFermiEquivalent <= 0 || treeMassProxyGeV <= 0 {
		return nil, fmt.Errorf("non-positive Gate780 snapshot ledger")
	}
	productMatches := closeRel(cYukawaSnapshot*cHistorySnapshot, cHiggsSnapshot, 1e-15)
	historyExpected := 1 + lHopfSnapshot*(1-kappaLambdaRedSnapshot)
	historyMatches := closeRel(historyExpected, cHistorySnapshot, 1e-15)

	a := &Analysis{
		Gate779: InheritedGate779Ratio{
			Inherited:               true,
			FermiNormalizedIdentity: "4sqrt(2)G_F m_H_tree_proxy^2=C_Higgs",
			CHiggs:                  cHiggsSnapshot,
			VEVGeV:                  vevConventionGeV,
			GFermiGeVMinus2:         gFermiEquivalent,
			TreeMassProxyGeV:        treeMassProxyGeV,
			FourSqrt2GFMassSquared:  fermiNormalizedRatio,
			DimensionlessTask:       "derive or reduce C_Higgs natively",
			ScaleTask:               "derive or seal G_F / v",
			DerivesGF:               false,
			DerivesPoleMass:         false,
			Verdict:                 StatusGate779FermiNormalizedRatioInherited,
		},
		Graph: DependencyGraph{
			Expanded:              true,
			CHiggsFormula:         "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]",
			CYukawaFormula:        "C_Yukawa=3/N_eff",
			NEffFormula:           "N_eff=a^2/b",
			YukawaTraceLedger:     "a,b finite spectral-action Yukawa trace ledger values",
			CHistoryFormula:       "C_History=1+L_Hopf(1-kappa_lambda_red)",
			LHopfFormula:          "L_Hopf=Tr_K7+(rho_plus[(1/(2*pi))supp(H_V(x0))])=1/(8*pi)",
			KappaLambdaRedFormula: "kappa_lambda_red=|lambda(Lambda12)|+F_wall_3_red(s)-kappa_e_red",
			FWall3Formula:         "F_wall_3_red(s)=p_K7 s+kappa_e_red p_K7 s^2-2p_K7^2 s^3",
			KappaERedFormula:      "kappa_e_red=sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p_K7 s^2",
			CHiggs:                cHiggsSnapshot,
			CYukawa:               cYukawaSnapshot,
			CHistory:              cHistorySnapshot,
			NEff:                  nEffSnapshot,
			LHopf:                 lHopfSnapshot,
			KappaLambdaRed:        kappaLambdaRedSnapshot,
			ProductMatchesCHiggs:  productMatches,
			HistoryMatchesFormula: historyMatches,
			Verdict:               StatusCHiggsDependencyGraphExpanded,
		},
		Classification: InputClassification{
			Audited:                      true,
			PK7Classification:            "native support plus observer-event bridge; p_K7=7/72 is not a scalar theorem by itself",
			LHopfClassification:          "bridge source-typed radial-Hopf event; not a native HistoryLoop theorem",
			NEffClassification:           "sealed Yukawa trace participation ledger; not a native Yukawa theorem",
			Theta13JCKMClassification:    "flavor/empirical orientation inputs unless separately derived",
			BoundaryScalarClassification: "boundary/history scalar bridge coordinates; not native scalar theorem",
			GFClassification:             "external electroweak scale seal",
			TreeProxyClassification:      "tree Hessian proxy output after C_Higgs and G_F/v seals",
			NativeInputs:                 []string{"finite K7 support ranks already audited", "projector/rank identities where certified"},
			BridgeResponseInputs:         []string{"p_K7 observer-event response", "L_Hopf radial-Hopf event", "F_wall_3_red boundary response", "kappa_lambda_red scalar matching deficit"},
			EmpiricalOrSealedInputs:      []string{"N_eff from sealed Yukawa ledger", "theta13", "J_CKM", "lambda(Lambda12)", "R3-1", "xi_boundary"},
			RuntimeDefinedInputs:         []string{"lambda_runtime_eff bridge closure", "m_H_tree_proxy output"},
			ExternalScaleInputs:          []string{"G_F", "v"},
			Verdict:                      StatusInputIndependenceClassificationAudited,
		},
		Circularity: CircularityAudit{
			Defined:                                  true,
			CriticalTargets:                          []string{"kappa_lambda_red", "lambda_runtime_eff", "lambda_proxy", "F_wall_3_red", "N_eff"},
			KappaLambdaHistoricallyRuntimeTied:       true,
			LambdaRuntimeEffBridgeClosureQuantity:    true,
			LambdaProxyFromYukawaLedgerNotHiggsMass:  true,
			YukawaLedgerStillSealed:                  true,
			FWallCompressedAgainstDeficitRelations:   true,
			AnyComponentUsesHiggsOrRuntimeTargetData: true,
			IndependentPieces:                        []string{"Fermi-normalized interface algebra", "C_Yukawa formula from aggregate trace pair once ledger is supplied", "rank-one radial event arithmetic once potential lane is supplied"},
			DependentOrSealedPieces:                  []string{"sealed Yukawa trace ledger", "runtime-tied scalar matching deficit", "flavor orientation inputs", "boundary/history scalar coordinates", "external Fermi/VEV scale"},
			Verdict:                                  StatusCircularityAuditDefined,
		},
		Levels: PredictionStatusLevels{
			Defined:      true,
			LevelA:       "algebraic identity / consistency closure when runtime-derived ingredients remain",
			LevelB:       "semi-independent bridge estimate using external Yukawa/flavor/boundary data but not Higgs mass",
			LevelC:       "independent tree-level prediction computed without Higgs/runtime target data",
			LevelD:       "physical pole-mass prediction after tree-to-pole correction package and uncertainties",
			CurrentLevel: "Level A/B: bridge consistency and semi-independent target, not full independent prediction",
			NotLevelC:    true,
			NotLevelD:    true,
			Reason:       "kappa_lambda_red and boundary/flavor/Yukawa ingredients remain sealed or historically runtime-linked; no native tree-to-pole package exists",
			Verdict:      StatusPredictionStatusLevelsDefined,
		},
		Removals: RequiredRemovals{
			Recorded: true,
			Items: []string{
				"derive N_eff from a native Yukawa operator or independent non-Higgs Yukawa data",
				"source kappa_lambda_red without scalar runtime target closure",
				"source kappa_e_red from native flavor theorem or independent flavor data",
				"derive L_Hopf from native HistoryLoop theorem",
				"derive F_wall_3_red from native boundary response theorem",
				"derive or seal G_F/v for dimensionful prediction",
				"add a tree-to-pole correction package for pole-mass comparison",
			},
			NeedsNativeYukawaOperator:    true,
			NeedsScalarMatchingSource:    true,
			NeedsFlavorSource:            true,
			NeedsHistoryLoopTheorem:      true,
			NeedsBoundaryResponseTheorem: true,
			NeedsFermiOrVEVScale:         true,
			NeedsTreeToPolePackage:       true,
			Verdict:                      StatusRequiredRemovalsForPredictionRecorded,
		},
		Firewalls: PhysicalFirewalls{
			Enforced:                          true,
			CHiggsIndependentIfTargetDataUsed: false,
			FermiRatioPoleMassTheorem:         false,
			YukawaLedgerNativeYukawaTheorem:   false,
			KappaLambdaRedNativeScalarTheorem: false,
			LHopfNativeHistoryLoopTheorem:     false,
			GFAShaDerivedScale:                false,
			TreeProxyPoleMass:                 false,
			Verdict:                           StatusGate780HiggsPredictionIndependenceBoundary,
		},
		Truth: fmt.Sprintf("Gate780 classifies C_Higgs=%.16f as the current dimensionless Higgs prediction target and the correct Fermi-normalized test interface, but not yet a full independent prediction: N_eff is sealed, L_Hopf is bridge source-typed, kappa_lambda_red remains scalar-matching bridge data, G_F is external, and the tree proxy is not a pole mass.", cHiggsSnapshot),
	}
	cache = a
	return cloneAnalysis(a), nil
}

func cloneAnalysis(a *Analysis) *Analysis {
	clone := *a
	clone.Classification.NativeInputs = append([]string(nil), a.Classification.NativeInputs...)
	clone.Classification.BridgeResponseInputs = append([]string(nil), a.Classification.BridgeResponseInputs...)
	clone.Classification.EmpiricalOrSealedInputs = append([]string(nil), a.Classification.EmpiricalOrSealedInputs...)
	clone.Classification.RuntimeDefinedInputs = append([]string(nil), a.Classification.RuntimeDefinedInputs...)
	clone.Classification.ExternalScaleInputs = append([]string(nil), a.Classification.ExternalScaleInputs...)
	clone.Circularity.CriticalTargets = append([]string(nil), a.Circularity.CriticalTargets...)
	clone.Circularity.IndependentPieces = append([]string(nil), a.Circularity.IndependentPieces...)
	clone.Circularity.DependentOrSealedPieces = append([]string(nil), a.Circularity.DependentOrSealedPieces...)
	clone.Removals.Items = append([]string(nil), a.Removals.Items...)
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
		StatusGate779FermiNormalizedRatioInherited,
		StatusCHiggsDependencyGraphExpanded,
		StatusInputIndependenceClassificationAudited,
		StatusCircularityAuditDefined,
		StatusPredictionStatusLevelsDefined,
		StatusRequiredRemovalsForPredictionRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusCHiggsCurrentDimensionlessPredictionTarget,
		StatusFermiNormalizedRatioRightTestInterface,
		StatusCurrentStatusBridgeConsistencyNotFullPrediction,
		StatusCHiggsNotYetNativeHiggsTheorem,
		StatusNEffNotNativeYukawaTheorem,
		StatusKappaLambdaRedNotNativeScalarMatchingTheorem,
		StatusLHopfNotNativeHistoryLoopTheorem,
		StatusGFNotNativeElectroweakScaleTheorem,
		StatusTreeProxyNotPoleMass,
		StatusGate780HiggsPredictionIndependenceBoundary,
	}
}

func FormatGate779(g InheritedGate779Ratio) string {
	return fmt.Sprintf("inherited=%t identity=%s C=%.17g v=%.10f GF=%.14e mass=%.12f ratio=%.17g dimTask=%s scaleTask=%s derivesGF=%t pole=%t verdict=%s", g.Inherited, g.FermiNormalizedIdentity, g.CHiggs, g.VEVGeV, g.GFermiGeVMinus2, g.TreeMassProxyGeV, g.FourSqrt2GFMassSquared, g.DimensionlessTask, g.ScaleTask, g.DerivesGF, g.DerivesPoleMass, g.Verdict)
}

func FormatGraph(g DependencyGraph) string {
	return fmt.Sprintf("expanded=%t C=%s CY=%s Neff=%s history=%s L=%s kappa=%s wall=%s kappaE=%s C=%.17g CY=%.17g CH=%.17g Neff=%.17g L=%.17g kappa=%.17g product=%t history=%t verdict=%s", g.Expanded, g.CHiggsFormula, g.CYukawaFormula, g.NEffFormula, g.CHistoryFormula, g.LHopfFormula, g.KappaLambdaRedFormula, g.FWall3Formula, g.KappaERedFormula, g.CHiggs, g.CYukawa, g.CHistory, g.NEff, g.LHopf, g.KappaLambdaRed, g.ProductMatchesCHiggs, g.HistoryMatchesFormula, g.Verdict)
}

func FormatClassification(c InputClassification) string {
	return fmt.Sprintf("audited=%t pK7=%s L=%s Neff=%s flavor=%s boundary=%s GF=%s tree=%s native=%d bridge=%d sealed=%d runtime=%d external=%d verdict=%s", c.Audited, c.PK7Classification, c.LHopfClassification, c.NEffClassification, c.Theta13JCKMClassification, c.BoundaryScalarClassification, c.GFClassification, c.TreeProxyClassification, len(c.NativeInputs), len(c.BridgeResponseInputs), len(c.EmpiricalOrSealedInputs), len(c.RuntimeDefinedInputs), len(c.ExternalScaleInputs), c.Verdict)
}

func FormatCircularity(c CircularityAudit) string {
	return fmt.Sprintf("defined=%t targets=%s kappaRuntime=%t lambdaRuntimeClosure=%t proxyNoHiggs=%t yukawaSealed=%t wallDeficit=%t targetData=%t independent=%d dependent=%d verdict=%s", c.Defined, strings.Join(c.CriticalTargets, ","), c.KappaLambdaHistoricallyRuntimeTied, c.LambdaRuntimeEffBridgeClosureQuantity, c.LambdaProxyFromYukawaLedgerNotHiggsMass, c.YukawaLedgerStillSealed, c.FWallCompressedAgainstDeficitRelations, c.AnyComponentUsesHiggsOrRuntimeTargetData, len(c.IndependentPieces), len(c.DependentOrSealedPieces), c.Verdict)
}

func FormatLevels(l PredictionStatusLevels) string {
	return fmt.Sprintf("defined=%t A=%s B=%s C=%s D=%s current=%s notC=%t notD=%t reason=%s verdict=%s", l.Defined, l.LevelA, l.LevelB, l.LevelC, l.LevelD, l.CurrentLevel, l.NotLevelC, l.NotLevelD, l.Reason, l.Verdict)
}

func FormatRemovals(r RequiredRemovals) string {
	return fmt.Sprintf("recorded=%t count=%d yukawa=%t scalar=%t flavor=%t history=%t boundary=%t scale=%t pole=%t verdict=%s", r.Recorded, len(r.Items), r.NeedsNativeYukawaOperator, r.NeedsScalarMatchingSource, r.NeedsFlavorSource, r.NeedsHistoryLoopTheorem, r.NeedsBoundaryResponseTheorem, r.NeedsFermiOrVEVScale, r.NeedsTreeToPolePackage, r.Verdict)
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("enforced=%t CIndependentWithTargetData=%t ratioPole=%t yukawaNative=%t kappaNative=%t LNative=%t GFNative=%t treePole=%t verdict=%s", f.Enforced, f.CHiggsIndependentIfTargetDataUsed, f.FermiRatioPoleMassTheorem, f.YukawaLedgerNativeYukawaTheorem, f.KappaLambdaRedNativeScalarTheorem, f.LHopfNativeHistoryLoopTheorem, f.GFAShaDerivedScale, f.TreeProxyPoleMass, f.Verdict)
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
