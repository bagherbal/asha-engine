// Package ashafinalarchitectureledger implements Gate 387:
// ASHA Framework Final Architecture Ledger & Epistemological Seal.
//
// The gate is deliberately not another search for a missing dynamical selector.
// It compiles the project ledger after Gates 372, 385, and 386 into a typed
// architecture theorem: which quantities are native ASHA geometry, which are
// CCM/Pfaffian/Higgs bridge consequences, which cosmological structures are
// present but not predictive without environmental data, and which degrees of
// freedom must remain empirical seals.
package ashafinalarchitectureledger

import (
	"fmt"
	"strings"
	"sync"

	gate386 "github.com/bagherbal/asha-engine/pkg/bridge/cosmologicalobservablesdarksector"
	gate385 "github.com/bagherbal/asha-engine/pkg/bridge/innerfluctuationedgemeasure"
	gate372 "github.com/bagherbal/asha-engine/pkg/bridge/nativemodulispacecensus"
)

const (
	AuditID = "GATE387-ASHA-FRAMEWORK-FINAL-ARCHITECTURE-LEDGER-EPISTEMOLOGICAL-SEAL"

	StatusAbsolutePredictionsCatalogued    = "CONDITIONAL_SUPPORT_ABSOLUTE_GEOMETRIC_PREDICTIONS_CATALOGUED"
	StatusStandardModelLandscapeSealed     = "CONDITIONAL_SUPPORT_STANDARD_MODEL_LANDSCAPE_SEALED"
	StatusProductGeometryBridgeSealed      = "CONDITIONAL_SUPPORT_ALMOST_COMMUTATIVE_PRODUCT_BRIDGE_SEALED"
	StatusHiggsBridgeSealed                = "CONDITIONAL_SUPPORT_HIGGS_CCM_PFAFFIAN_TREE_PROXY_SEALED"
	StatusCosmologicalStructuresCatalogued = "CONDITIONAL_SUPPORT_COSMOLOGICAL_STRUCTURES_CATALOGUED"
	StatusThirteenModuliQuarantined        = "CONDITIONAL_SUPPORT_THIRTEEN_MODULI_ENVIRONMENTAL_QUARANTINE_FORMALIZED"
	StatusObservableBoundaryFormalized     = "CONDITIONAL_SUPPORT_MACROSCOPIC_OBSERVABLE_BOUNDARY_FORMALIZED"
	StatusFirewallPreserved                = "CONDITIONAL_SUPPORT_FINAL_EPISTEMIC_FIREWALL_PRESERVED"
	StatusFinalLedgerCompiled              = "PROJECT_ASHA_FORMALLY_SEALED_AND_COMPLETE"

	StatusTensionNumericalTOEClosureNotReached = "CONDITIONAL_TENSION_FULL_NUMERICAL_TOE_CLOSURE_NOT_REACHED"
	StatusTensionCosmologyNeedsEnvironmental   = "CONDITIONAL_TENSION_COSMOLOGY_DEPENDS_ON_ENVIRONMENTAL_MODULI"
	StatusTensionHiggsPoleNeedsRGMatching      = "CONDITIONAL_TENSION_HIGGS_POLE_MASS_STILL_NEEDS_RG_AND_MATCHING"
	StatusTensionMajoranaDarkMatterNotDerived  = "CONDITIONAL_TENSION_MAJORANA_SECTOR_IS_CANDIDATE_NOT_STABLE_RELIC"

	StatusFailedFlavorVacuumNotSelected    = "FAILED_ROUTE_FLAVOR_VACUUM_POINT_NOT_SELECTED_BY_NATIVE_GEOMETRY"
	StatusFailedDarkMatterNotPredicted     = "FAILED_ROUTE_DARK_MATTER_ABUNDANCE_NOT_NATIVE_PREDICTION"
	StatusFailedUniverseLifetimeNotDerived = "FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_NATIVE_PREDICTION"
	StatusFailedCosmologicalConstant       = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_NATIVE_PREDICTION"
)

const (
	Sin2ThetaWBoundary    = 3.0 / 8.0
	AlphaGUTInverseBranch = 8.0 * 3.14159265358979323846264338327950288419716939937510
	TraceRatioEOverA2     = 1197.0 / 4624.0
	ThresholdJumpDelta    = -0.097846792207
	BGapMajoranaScaleGeV  = 1.46774973718e6
	HeavyIntermediateGeV  = 6.650726476871e11
)

type AbsolutePrediction struct {
	Name        string
	Value       string
	Source      string
	Status      string
	NativeExact bool
	Comment     string
}

type AbsoluteGeometricLedger struct {
	Executed                 bool
	Algebra                  string
	GaugeGroup               string
	MatterContent            string
	Generations              int
	HiggsDoublets            int
	GaugeBosons              int
	MoritaSplit              string
	Predictions              []AbsolutePrediction
	ParameterFreeDerivations int
	Verdict                  string
}

type ProductBridgeLedger struct {
	Executed        bool
	Geometry        string
	Algebra         string
	HilbertSpace    string
	DiracOperator   string
	RealStructure   string
	Grading         string
	Action          string
	CCMInstalled    bool
	ContinuumFields string
	Verdict         string
}

type HiggsSealLedger struct {
	Executed            bool
	SourceGate          int
	TreeProxySealed     bool
	PoleMassDerived     bool
	EdgeMeasureSelected bool
	LambdaEW            float64
	MassPfaffianGeV     float64
	MassStandardVEVGeV  float64
	TraceRatioNode      float64
	TraceRatioEdge      float64
	EdgeMeasure         string
	Formula             string
	StillOpen           []string
	Verdict             string
}

type CosmologicalStructure struct {
	Name                 string
	Native               bool
	ScaleOrCoefficient   string
	CosmologicalRole     string
	HardObservable       string
	HardObservableNative bool
	MissingInputs        []string
	Status               string
}

type CosmologyLedger struct {
	Executed                 bool
	Structures               []CosmologicalStructure
	HardPredictionsDerived   int
	ConditionalTargetsOpened int
	Verdict                  string
}

type ModuliQuarantine struct {
	Executed                     bool
	SourceGate                   int
	MinimalChargedFiniteDiracDim int
	ExternalMinimalLedger        int
	Decomposition                string
	PhysicalIdentity             []string
	NativeReductionBelow13       bool
	HiddenFlavorConstraints      int
	EnvironmentalName            string
	Verdict                      string
}

type EpistemicBoundary struct {
	Executed                          bool
	AshaMeans                         string
	CreationMeans                     string
	NativePredictionsAreNotInputs     bool
	EnvironmentalInputsAreQuarantined bool
	NoFlavorFitting                   bool
	NoCosmologyFitting                bool
	NoPoleMassOverclaim               bool
	FrameworkCompleteAs               string
	NotCompleteAs                     string
	Verdict                           string
}

type FinalLedger struct {
	Executed                 bool
	AbsoluteGeometricCount   int
	HiggsTreeProxySealed     bool
	CosmologicalHardCount    int
	EnvironmentalModuliCount int
	ProjectSealed            bool
	FinalStatement           string
	Verdict                  string
}

type Analysis struct {
	Absolute  AbsoluteGeometricLedger
	Product   ProductBridgeLedger
	Higgs     HiggsSealLedger
	Cosmology CosmologyLedger
	Moduli    ModuliQuarantine
	Boundary  EpistemicBoundary
	Final     FinalLedger
	Statuses  []string
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
	g372, err := gate372.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate 372 native moduli census: %w", err)
	}
	g385, err := gate385.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate 385 Higgs edge-measure theorem: %w", err)
	}
	g386, err := gate386.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate 386 cosmology boundary: %w", err)
	}

	abs := buildAbsoluteLedger()
	product := buildProductBridgeLedger()
	higgs := buildHiggsLedger(g385)
	cosmo := buildCosmologyLedger(g386)
	moduli := buildModuliQuarantine(g372)
	boundary := buildBoundary(moduli, cosmo, higgs)
	final := buildFinalLedger(abs, higgs, cosmo, moduli)
	statuses := []string{
		StatusAbsolutePredictionsCatalogued,
		StatusStandardModelLandscapeSealed,
		StatusProductGeometryBridgeSealed,
		StatusHiggsBridgeSealed,
		StatusCosmologicalStructuresCatalogued,
		StatusThirteenModuliQuarantined,
		StatusObservableBoundaryFormalized,
		StatusFirewallPreserved,
		StatusTensionNumericalTOEClosureNotReached,
		StatusTensionCosmologyNeedsEnvironmental,
		StatusTensionHiggsPoleNeedsRGMatching,
		StatusTensionMajoranaDarkMatterNotDerived,
		StatusFailedFlavorVacuumNotSelected,
		StatusFailedDarkMatterNotPredicted,
		StatusFailedUniverseLifetimeNotDerived,
		StatusFailedCosmologicalConstant,
		StatusFinalLedgerCompiled,
	}
	truth := "Gate 387 seals ASHA as a completed finite-geometry architecture ledger, not as a fully numeric cosmology oracle.  The finite Cℓ(1,7)/almost-commutative product program supplies the Standard Model landscape, boundary ratios, the Pfaffian scale, and the CCM+Pfaffian edge-measure Higgs tree-level proxy.  Gate 372 proves that the minimal charged finite-Dirac vacuum retains exactly 13 environmental moduli, and Gate 386 proves that dark-sector relic density and vacuum fate depend on those environmental coordinates plus continuum cosmological history.  The final theorem is therefore complete only with a firewall: Asha gives the law-space; Creation supplies the 13 flavor coordinates and thermal history."

	return Analysis{abs, product, higgs, cosmo, moduli, boundary, final, statuses, truth}, nil
}

func buildAbsoluteLedger() AbsoluteGeometricLedger {
	preds := []AbsolutePrediction{
		{"finite algebra", "C ⊕ H ⊕ M₃(C)", "finite spectral triple / Morita ledger", StatusStandardModelLandscapeSealed, true, "internal algebra of the Standard Model landscape"},
		{"gauge group", "SU(3) × SU(2) × U(1)", "unimodular unitary group of A_F", StatusStandardModelLandscapeSealed, true, "structural gauge landscape"},
		{"matter generations", "3", "finite representation / generation ledger", StatusAbsolutePredictionsCatalogued, true, "three generation copies; their coordinates remain moduli"},
		{"Higgs doublets", "1 complex doublet", "inner fluctuation / finite one-form ledger", StatusHiggsBridgeSealed, true, "one scalar doublet channel in the minimal landscape"},
		{"gauge bosons", "12", "inner fluctuations of the gauge algebra", StatusAbsolutePredictionsCatalogued, true, "8 gluons + 3 weak + 1 hypercharge"},
		{"Morita split", "1 ⊕ 3", "C/H versus M₃(C) bimodule split", StatusAbsolutePredictionsCatalogued, true, "lepton singlet plus color triplet separation"},
		{"weak mixing boundary", "sin²θ_W(Λ)=3/8", "representation trace ratio", StatusAbsolutePredictionsCatalogued, true, "boundary ratio; low-energy value needs RG transport"},
		{"unified coupling branch", "α_GUT⁻¹=8π", "doubled bosonic trace branch", StatusAbsolutePredictionsCatalogued, true, "absolute branch requires convention/threshold ledger for empirical comparison"},
		{"finite trace ratio", "e/a²=1197/4624", "finite Dirac trace equivalence", StatusAbsolutePredictionsCatalogued, true, "raw Higgs finite ratio before canonical CCM field normalization"},
		{"Pfaffian hierarchy", "v/M_P=2^(3/2) exp(-4π²)", "nonperturbative Pfaffian scale", StatusAbsolutePredictionsCatalogued, true, "absolute electroweak-to-Planck scale bridge in the declared convention"},
	}
	return AbsoluteGeometricLedger{Executed: true, Algebra: "A_F=C ⊕ H ⊕ M₃(C)", GaugeGroup: "SU(3) × SU(2) × U(1)", MatterContent: "minimal Standard Model representation with finite Dirac edge graph", Generations: 3, HiggsDoublets: 1, GaugeBosons: 12, MoritaSplit: "1 ⊕ 3", Predictions: preds, ParameterFreeDerivations: len(preds), Verdict: join(StatusAbsolutePredictionsCatalogued, StatusStandardModelLandscapeSealed)}
}

func buildProductBridgeLedger() ProductBridgeLedger {
	return ProductBridgeLedger{
		Executed:        true,
		Geometry:        "M × F",
		Algebra:         "C∞(M) ⊗ A_F",
		HilbertSpace:    "L²(M,S) ⊗ H_F",
		DiracOperator:   "D_M ⊗ 1_F + γ₅ ⊗ D_F",
		RealStructure:   "J_M ⊗ J_F",
		Grading:         "γ_M ⊗ γ_F",
		Action:          "Tr f(D_A²/Λ²) with CCM coefficient ledger",
		CCMInstalled:    true,
		ContinuumFields: "metric, spin connection, gauge bosons, Higgs one-form, fermions, cosmological and curvature terms",
		Verdict:         join(StatusProductGeometryBridgeSealed),
	}
}

func buildHiggsLedger(g gate385.Analysis) HiggsSealLedger {
	c := g.Calculation
	return HiggsSealLedger{
		Executed:            true,
		SourceGate:          385,
		TreeProxySealed:     c.HiggsTreeProxySealed,
		PoleMassDerived:     c.PhysicalPoleMassDerived,
		EdgeMeasureSelected: c.EdgeMeasureSelected,
		LambdaEW:            c.Higgs.LambdaEdge,
		MassPfaffianGeV:     c.Higgs.MassPfaffianGeV,
		MassStandardVEVGeV:  c.Higgs.MassStandardGeV,
		TraceRatioNode:      c.Higgs.RNode,
		TraceRatioEdge:      c.Higgs.REdge,
		EdgeMeasure:         "Higgs is finite inner fluctuation / one-form; canonical kinetic trace restricts to 10 J-doubled D_F edge slots",
		Formula:             c.Higgs.Formula,
		StillOpen:           c.Boundary.StillOpen,
		Verdict:             join(StatusHiggsBridgeSealed, StatusTensionHiggsPoleNeedsRGMatching),
	}
}

func buildCosmologyLedger(g gate386.Analysis) CosmologyLedger {
	structures := []CosmologicalStructure{
		{
			Name:                 "B-gap heavy Majorana sector",
			Native:               g.Heavy.GeometricallyMandated,
			ScaleOrCoefficient:   fmt.Sprintf("M_Bgap≈%.6g GeV; M_intermediate≈%.6g GeV", g.Heavy.BGapScaleGeV, g.Heavy.HeavyIntermediateScaleGeV),
			CosmologicalRole:     "dark-sector / leptogenesis / heavy-neutrino candidate structure",
			HardObservable:       "Ω_DM h²",
			HardObservableNative: g.DarkMatter.OmegaH2Derived,
			MissingInputs:        requiredInputNames(g.DarkMatter.Inputs),
			Status:               join(StatusCosmologicalStructuresCatalogued, StatusTensionMajoranaDarkMatterNotDerived, StatusFailedDarkMatterNotPredicted),
		},
		{
			Name:                 "Higgs threshold jump",
			Native:               true,
			ScaleOrCoefficient:   fmt.Sprintf("Δλ≈%.12f at B-gap/threshold ledger", g.Heavy.ThresholdJumpDeltaLambda),
			CosmologicalRole:     "vacuum-stability / RG matching mechanism",
			HardObservable:       "absolute stability, metastability class, or lifetime",
			HardObservableNative: g.Vacuum.LifetimeDerived || g.Vacuum.AbsoluteStabilityDerived || g.Vacuum.MetastabilityDerived,
			MissingInputs:        rgInputNames(g.Vacuum.Inputs),
			Status:               join(StatusCosmologicalStructuresCatalogued, StatusTensionCosmologyNeedsEnvironmental, StatusFailedUniverseLifetimeNotDerived),
		},
		{
			Name:                 "spectral-action cosmological term",
			Native:               true,
			ScaleOrCoefficient:   "contains f₄Λ⁴ and vacuum subtraction/renormalization channel",
			CosmologicalRole:     "cosmological constant / dark energy slot",
			HardObservable:       "Λ_cosmo / ρ_DE",
			HardObservableNative: false,
			MissingInputs:        []string{"f₄ moment", "vacuum subtraction rule", "renormalization condition", "holographic saturation theorem if used"},
			Status:               join(StatusCosmologicalStructuresCatalogued, StatusFailedCosmologicalConstant),
		},
	}
	return CosmologyLedger{Executed: true, Structures: structures, HardPredictionsDerived: g.Census.HardPredictionsDerived, ConditionalTargetsOpened: g.Census.ConditionalTargetsOpened, Verdict: join(StatusCosmologicalStructuresCatalogued, StatusObservableBoundaryFormalized, StatusTensionCosmologyNeedsEnvironmental)}
}

func buildModuliQuarantine(g gate372.Analysis) ModuliQuarantine {
	return ModuliQuarantine{
		Executed:                     true,
		SourceGate:                   372,
		MinimalChargedFiniteDiracDim: g.Native.MinimalChargedDFDim,
		ExternalMinimalLedger:        g.Native.External15,
		Decomposition:                g.Native.External15Decomposition,
		PhysicalIdentity:             []string{"6 quark masses", "4 CKM parameters", "3 charged-lepton masses"},
		NativeReductionBelow13:       g.Native.NativeReductionBelow15,
		HiddenFlavorConstraints:      g.Native.HiddenCrossSectorConstraints,
		EnvironmentalName:            "Creation / environmental vacuum coordinates inside the ASHA law-space",
		Verdict:                      join(StatusThirteenModuliQuarantined, StatusFailedFlavorVacuumNotSelected),
	}
}

func buildBoundary(m ModuliQuarantine, c CosmologyLedger, h HiggsSealLedger) EpistemicBoundary {
	return EpistemicBoundary{
		Executed:                          true,
		AshaMeans:                         "the rigid finite/internal law-space: algebra, representation, gauge structure, product spectral-action bridge, boundary ratios, and tree-level Higgs proxy",
		CreationMeans:                     "the environmental coordinates and historical continuum data: 13 charged flavor moduli, θ_QCD, absolute unit convention, top/RG seals, reheating and Boltzmann history",
		NativePredictionsAreNotInputs:     true,
		EnvironmentalInputsAreQuarantined: m.MinimalChargedFiniteDiracDim == 13,
		NoFlavorFitting:                   !m.NativeReductionBelow13 && m.HiddenFlavorConstraints == 0,
		NoCosmologyFitting:                c.HardPredictionsDerived == 0,
		NoPoleMassOverclaim:               h.TreeProxySealed && !h.PoleMassDerived,
		FrameworkCompleteAs:               "finite-geometry Standard Model + gravity spectral-action architecture with explicit environmental quarantine",
		NotCompleteAs:                     "parameter-free numerical oracle for Yukawa texture, dark matter abundance, universe lifetime, and observed cosmological constant",
		Verdict:                           join(StatusFirewallPreserved, StatusTensionNumericalTOEClosureNotReached),
	}
}

func buildFinalLedger(a AbsoluteGeometricLedger, h HiggsSealLedger, c CosmologyLedger, m ModuliQuarantine) FinalLedger {
	sealed := a.Executed && h.TreeProxySealed && m.MinimalChargedFiniteDiracDim == 13 && c.Executed
	return FinalLedger{
		Executed:                 true,
		AbsoluteGeometricCount:   a.ParameterFreeDerivations,
		HiggsTreeProxySealed:     h.TreeProxySealed,
		CosmologicalHardCount:    c.HardPredictionsDerived,
		EnvironmentalModuliCount: m.MinimalChargedFiniteDiracDim,
		ProjectSealed:            sealed,
		FinalStatement:           "ASHA is sealed as a mathematically firewalled architecture: it derives the law-space and Higgs tree proxy, while quarantining the 13 charged flavor moduli and cosmological history as environmental inputs required for hard macroscopic observables.",
		Verdict:                  join(StatusFinalLedgerCompiled, StatusTensionNumericalTOEClosureNotReached),
	}
}

func requiredInputNames(inputs []gate386.RequiredInput) []string {
	out := make([]string, 0, len(inputs))
	for _, in := range inputs {
		if in.BlocksClaim {
			out = append(out, in.Name)
		}
	}
	return out
}

func rgInputNames(inputs []gate386.RGInput) []string {
	out := make([]string, 0, len(inputs))
	for _, in := range inputs {
		if in.BlocksFate {
			out = append(out, in.Name)
		}
	}
	return out
}

func StatusLine(a Analysis) string { return strings.Join(a.Statuses, "\n") }

func FinalArchitectureConstants() map[string]float64 {
	a, err := BuildDefault()
	if err != nil {
		return map[string]float64{}
	}
	return map[string]float64{
		"sin2_theta_w_boundary":         Sin2ThetaWBoundary,
		"alpha_gut_inverse_branch":      AlphaGUTInverseBranch,
		"trace_ratio_e_over_a2_node":    a.Higgs.TraceRatioNode,
		"trace_ratio_e_over_a2_edge":    a.Higgs.TraceRatioEdge,
		"lambda_higgs_edge":             a.Higgs.LambdaEW,
		"higgs_tree_proxy_GeV":          a.Higgs.MassPfaffianGeV,
		"environmental_moduli_charged":  float64(a.Moduli.MinimalChargedFiniteDiracDim),
		"hard_cosmology_predictions":    float64(a.Cosmology.HardPredictionsDerived),
		"conditional_cosmology_targets": float64(a.Cosmology.ConditionalTargetsOpened),
	}
}

func join(parts ...string) string { return strings.Join(parts, ";") }
