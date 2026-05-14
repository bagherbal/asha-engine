// Package spectralgraphtracenormalization implements Gate 383:
// Spectral Graph Trace / Node-to-Edge Kinetic Normalization Sieve.
//
// Gate 382 isolated the exact 10/7 normalization gap between the old contact
// node ledger f0=7 and the J-doubled finite-Dirac edge denominator 10 that
// reproduces the CCM+Pfaffian Higgs near-closure. Gate 383 asks whether the
// Higgs kinetic term in the product spectral action is mathematically forced
// to use the edge/interactions trace rather than the contact-node trace.
package spectralgraphtracenormalization

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ccmpfaffianf0closure"
	"github.com/bagherbal/asha-engine/pkg/bridge/finitetraceedgemultiplicity"
)

const (
	AuditID = "GATE383-SPECTRAL-GRAPH-TRACE-NODE-TO-EDGE-KINETIC-NORMALIZATION-SIEVE"

	StatusNodeEdgeDomainsFormalized   = "CONDITIONAL_SUPPORT_NODE_EDGE_TRACE_DOMAINS_FORMALIZED"
	StatusKineticTraceSupportAudited  = "CONDITIONAL_SUPPORT_HIGGS_KINETIC_TRACE_SUPPORT_AUDITED"
	StatusEdgeSupportedKineticWitness = "CONDITIONAL_SUPPORT_EDGE_SUPPORTED_KINETIC_TRACE_WITNESS_FOUND"
	StatusTenOverSevenBridgeComputed  = "CONDITIONAL_SUPPORT_TEN_OVER_SEVEN_NODE_TO_EDGE_BRIDGE_COMPUTED"
	StatusEdgeDenominatorNearClosure  = "CONDITIONAL_SUPPORT_EDGE_DENOMINATOR_REPRODUCES_NEAR_HIGGS_CLOSURE"
	StatusPfaffianMassLaneComputed    = "CONDITIONAL_SUPPORT_PFAFFIAN_HIGGS_MASS_LANE_COMPUTED"
	StatusCCMMomentRemainsUnit        = "CONDITIONAL_SUPPORT_CCM_MOMENT_REMAINS_SEPARATE_FROM_GRAPH_MULTIPLICITY"

	StatusTensionKineticUsesEdgesButCCMUsesA       = "CONDITIONAL_TENSION_KINETIC_TERM_IS_EDGE_SUPPORTED_BUT_CCM_USES_A_TRACE"
	StatusTensionEOverA2MayAlreadyBeEdgeNormalized = "CONDITIONAL_TENSION_E_OVER_A2_MAY_ALREADY_INCLUDE_EDGE_TRACE_NORMALIZATION"
	StatusTensionTenOverSevenRequiresRawTraceAudit = "CONDITIONAL_TENSION_TEN_OVER_SEVEN_REQUIRES_RAW_A_AND_E_RECOMPUTATION"
	StatusTensionNodeToEdgeBridgeNotYetATheorem    = "CONDITIONAL_TENSION_NODE_TO_EDGE_BRIDGE_NOT_YET_A_CCM_THEOREM"

	StatusFailedEdgeTraceNormalizationNotDerived = "FAILED_ROUTE_EDGE_TRACE_NORMALIZATION_NOT_DERIVED"
	StatusFailedTenOverSevenNotDerived           = "FAILED_ROUTE_TEN_OVER_SEVEN_NOT_DERIVED"
	StatusFailedHiggsMassNotGeometricallySealed  = "FAILED_ROUTE_HIGGS_MASS_NOT_GEOMETRICALLY_SEALED"
	StatusFailedFullNumericalTOEClosureOpen      = "FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED"
)

const (
	ContactNodeCount  = 7.0
	JDoubledEdgeCount = 10.0
	TraceRatioEOverA2 = 1197.0 / 4624.0
	HiggsTargetGeV    = 125.10
	StandardVEVGeV    = 246.22
)

type TraceDomain struct {
	Name                  string
	ObjectCount           float64
	MathematicalSlot      string
	RepresentativeFormula string
	Role                  string
	CanReplaceCCMF0       bool
	Verdict               string
}

type KineticTraceAudit struct {
	KineticTermFormula      string
	UsesDFCommutator        bool
	SupportDomain           string
	MandatesEdgeSupport     bool
	MandatesEdgeDenominator bool
	CCMCanonicalCoefficient string
	Verdict                 string
}

type BridgeExtraction struct {
	NodeCount              float64
	EdgeCount              float64
	NodeToEdgeRatio        float64
	BridgeFormula          string
	CombinatorialNative    bool
	CCMNormalizationNative bool
	Verdict                string
}

type HiggsLane struct {
	Name             string
	Formula          string
	Denominator      float64
	LambdaH          float64
	MassStandardGeV  float64
	MassPfaffianGeV  float64
	ErrorPfaffianGeV float64
	PercentErrorPf   float64
	Native           bool
	Sealed           bool
	Verdict          string
}

type ClosureAudit struct {
	RequiredTheorem  string
	KnownEvidence    []string
	OpenObstructions []string
	Conclusion       string
}

type Calculation struct {
	Executed               bool
	Domains                []TraceDomain
	Kinetic                KineticTraceAudit
	Bridge                 BridgeExtraction
	Lanes                  []HiggsLane
	Closure                ClosureAudit
	Statuses               []string
	EdgeTraceDerived       bool
	TenOverSevenDerived    bool
	HiggsMassSealed        bool
	FullNumericalTOEClosed bool
	Truth                  string
}

type Analysis struct{ Calculation Calculation }

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	gate382, err := finitetraceedgemultiplicity.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate 382 finite trace audit: %w", err)
	}
	gate380 := ccmpfaffianf0closure.NativeConstants()
	if len(gate380) == 0 {
		return Analysis{}, fmt.Errorf("could not inherit Gate 380 CCM+Pfaffian constants")
	}
	pfVEV := gate380["pfaffian_vev_gev"]
	if pfVEV == 0 {
		pfVEV = ccmpfaffianf0closure.UnreducedPlanckGeV * math.Pow(2, 1.5) * math.Exp(-4*math.Pi*math.Pi)
	}

	if gate382.Calculation.EdgeMultiplicity.JDoubledEdgeCount != int(JDoubledEdgeCount) {
		return Analysis{}, fmt.Errorf("Gate 382 did not expose the expected ten J-doubled edge slots")
	}

	domains := []TraceDomain{
		{
			Name:                  "contact node trace",
			ObjectCount:           ContactNodeCount,
			MathematicalSlot:      "finite contact/vacuum support measure",
			RepresentativeFormula: "Tr_node(P_contact)=7",
			Role:                  "Counts the contact spectral nodes used by the earlier f0=7 ledger.",
			CanReplaceCCMF0:       false,
			Verdict:               "The node ledger is a real finite support count, but it is not automatically the CCM test-function moment and it is not the natural support of D_F interaction terms.",
		},
		{
			Name:                  "J-doubled finite Dirac edge trace",
			ObjectCount:           JDoubledEdgeCount,
			MathematicalSlot:      "finite interaction/endomorphism channel support",
			RepresentativeFormula: "Tr_edge(P_DF)=2×5=10",
			Role:                  "Counts the five allowed finite Dirac edge classes together with their J-conjugates.",
			CanReplaceCCMF0:       false,
			Verdict:               "The edge ledger is the correct-looking support for inner fluctuations and kinetic normalization, but it still cannot replace CCM f0. It must enter through the finite a/e trace channel if it enters at all.",
		},
	}

	kinetic := KineticTraceAudit{
		KineticTermFormula:      "a |D_μ H|² with a=Tr_F(Y†Y); the trace is induced by finite Dirac edges in D_F and the inner fluctuation one-form.",
		UsesDFCommutator:        true,
		SupportDomain:           "edge-supported, because the Higgs field appears as a finite one-form/fluctuation along allowed D_F interaction channels",
		MandatesEdgeSupport:     true,
		MandatesEdgeDenominator: false,
		CCMCanonicalCoefficient: "Canonical normalization divides the quartic trace e by the square of the kinetic trace a, producing λ ∝ e/a² with convention-dependent outer factors.",
		Verdict:                 "Gate 383 finds a valid structural witness that the Higgs kinetic term lives on D_F edges, not on bare contact nodes. However, CCM already packages that support into a=Tr_F(Y†Y). Therefore this witness does not by itself prove an extra replacement 7→10 in the denominator unless a and e are recomputed from raw node and edge measures.",
	}

	bridge := BridgeExtraction{
		NodeCount:              ContactNodeCount,
		EdgeCount:              JDoubledEdgeCount,
		NodeToEdgeRatio:        JDoubledEdgeCount / ContactNodeCount,
		BridgeFormula:          "N_edge,J / N_node = 10/7",
		CombinatorialNative:    true,
		CCMNormalizationNative: false,
		Verdict:                "The ratio 10/7 is an exact native graph conversion between contact nodes and J-doubled Dirac edges. What remains unproven is that the CCM Higgs canonical normalization must apply this graph conversion to the already-recorded e/a²=1197/4624 ratio.",
	}

	lanes := []HiggsLane{
		higgsLane("contact-node denominator", "λ=π²(e/a²)/(2·7)", ContactNodeCount, StandardVEVGeV, pfVEV, true, false, "This is the old contact ledger. It overpredicts the Higgs mass and therefore cannot be the final physical kinetic normalization if the CCM formula is used literally."),
		higgsLane("edge-trace denominator", "λ=π²(e/a²)/(2·10)", JDoubledEdgeCount, StandardVEVGeV, pfVEV, false, false, "This lane implements the node-to-edge correction and reproduces the near-125 GeV Higgs mass. It is the strongest closure witness, but it still requires the missing CCM kinetic normalization theorem."),
		higgsLane("unit sharp-cutoff denominator", "λ=π²(e/a²)/(2·1)", 1, StandardVEVGeV, pfVEV, true, false, "This keeps f0=1 but applies no finite kinetic normalization conversion. It badly overpredicts the Higgs mass."),
	}

	closure := ClosureAudit{
		RequiredTheorem: "SpectralGraphKineticNormalizationTheorem: the canonical Higgs kinetic trace in the ASHA finite spectral triple is edge-normalized with N_edge,J=10 rather than contact-node-normalized with N_node=7, and this replacement is not already contained in the stored full finite trace ratio e/a².",
		KnownEvidence: []string{
			"The Higgs field is an inner fluctuation/finite one-form and therefore is supported on allowed D_F interaction edges.",
			"The finite Dirac graph has exactly five structural edge classes and ten J-doubled edge slots.",
			"Using denominator 10 gives m_H≈124.9 GeV with the Pfaffian VEV lane.",
			"The node-to-edge conversion factor is exactly 10/7.",
		},
		OpenObstructions: []string{
			"The CCM kinetic coefficient uses a=Tr_F(Y†Y), and the ASHA ledger has not yet performed the raw a/e trace recomputation under separate node-measure and edge-measure conventions.",
			"The recorded ratio e/a²=1197/4624 may already be a full finite trace ratio; applying 10/7 after the fact could double-count normalization.",
			"No package currently proves that contact f0=7 is the node normalization of a while edge count 10 is the canonical normalization of the physical Higgs field.",
		},
		Conclusion: "The edge-trace explanation is the correct next theorem target and the strongest Higgs closure lane, but Gate 383 does not honestly seal it without raw a/e trace recomputation.",
	}

	statuses := []string{
		StatusNodeEdgeDomainsFormalized,
		StatusKineticTraceSupportAudited,
		StatusEdgeSupportedKineticWitness,
		StatusTenOverSevenBridgeComputed,
		StatusEdgeDenominatorNearClosure,
		StatusPfaffianMassLaneComputed,
		StatusCCMMomentRemainsUnit,
		StatusTensionKineticUsesEdgesButCCMUsesA,
		StatusTensionEOverA2MayAlreadyBeEdgeNormalized,
		StatusTensionTenOverSevenRequiresRawTraceAudit,
		StatusTensionNodeToEdgeBridgeNotYetATheorem,
		StatusFailedEdgeTraceNormalizationNotDerived,
		StatusFailedTenOverSevenNotDerived,
		StatusFailedHiggsMassNotGeometricallySealed,
		StatusFailedFullNumericalTOEClosureOpen,
	}

	truth := "Gate 383 proves the conceptual correction: the Higgs kinetic term is structurally supported on finite Dirac interaction edges, while the old contact ledger counted seven nodes. This gives an exact native graph conversion 10/7 and the edge-denominator lane reproduces the near-125 GeV Higgs mass. But the gate does not yet seal the Higgs mass, because CCM canonical normalization already uses the finite trace a=Tr_F(Y†Y), and the stored ratio e/a²=1197/4624 may already include the relevant edge support. The missing final theorem is a raw a/e recomputation under node versus edge measures proving that 7 must be replaced by 10 in the kinetic normalization denominator without double-counting."

	return Analysis{Calculation: Calculation{
		Executed:               true,
		Domains:                domains,
		Kinetic:                kinetic,
		Bridge:                 bridge,
		Lanes:                  lanes,
		Closure:                closure,
		Statuses:               statuses,
		EdgeTraceDerived:       false,
		TenOverSevenDerived:    false,
		HiggsMassSealed:        false,
		FullNumericalTOEClosed: false,
		Truth:                  truth,
	}}, nil
}

func higgsLane(name, formula string, denom, vStd, vPf float64, native, sealed bool, verdict string) HiggsLane {
	lambda := math.Pi * math.Pi * TraceRatioEOverA2 / (2 * denom)
	mStd := vStd * math.Sqrt(2*lambda)
	mPf := vPf * math.Sqrt(2*lambda)
	return HiggsLane{
		Name:             name,
		Formula:          formula,
		Denominator:      denom,
		LambdaH:          lambda,
		MassStandardGeV:  mStd,
		MassPfaffianGeV:  mPf,
		ErrorPfaffianGeV: mPf - HiggsTargetGeV,
		PercentErrorPf:   100 * (mPf - HiggsTargetGeV) / HiggsTargetGeV,
		Native:           native,
		Sealed:           sealed,
		Verdict:          verdict,
	}
}

func StatusLine(c Calculation) string { return strings.Join(c.Statuses, "\n") }

func NativeConstants() map[string]float64 {
	a, err := BuildDefault()
	if err != nil {
		return map[string]float64{}
	}
	c := a.Calculation
	m := map[string]float64{
		"contact_node_count":    c.Bridge.NodeCount,
		"j_doubled_edge_count":  c.Bridge.EdgeCount,
		"node_to_edge_ratio":    c.Bridge.NodeToEdgeRatio,
		"trace_ratio_e_over_a2": TraceRatioEOverA2,
		"edge_trace_derived":    boolFloat(c.EdgeTraceDerived),
		"higgs_mass_sealed":     boolFloat(c.HiggsMassSealed),
	}
	for _, l := range c.Lanes {
		key := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(l.Name, " ", "_"), "-", "_"))
		m[key+"_lambda"] = l.LambdaH
		m[key+"_mh_pfaffian"] = l.MassPfaffianGeV
	}
	return m
}

func boolFloat(v bool) float64 {
	if v {
		return 1
	}
	return 0
}
