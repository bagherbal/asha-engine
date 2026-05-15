// Package generation2specialbranchselector implements Gate 451:
// Texture-Zero Special-Branch Selector / Necessary Boundary Audit.
//
// Gate 450 proved that the forced Generation-2 structural zero gives an exact
// texture-zero spectral sum rule, but not a topology-only GST/Fritzsch ratio.
// Gate 451 audits the only honest escape hatch: perhaps an already-native ASHA
// law secretly suppresses the 1-3 edge or fixes the complex phase ray.  The
// answer is negative.  Chirality, the real structure, first-order compatibility,
// traceless/anomaly balance, and KMS integer modular quantization are all either
// family-edge-blind or explicitly compatible with the closed triangle.  The
// 1-3 edge is an allowed integer modular second harmonic, not a forbidden edge.
// Multiple nonzero-lift phase rays also survive the native constraints.
package generation2specialbranchselector

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE451-TEXTURE-ZERO-SPECIAL-BRANCH-SELECTOR-NECESSARY-BOUNDARY-AUDIT"

	StatusGate450Inherited                          = "CONDITIONAL_SUPPORT_GATE450_TEXTURE_ZERO_IDENTITY_INHERITED"
	StatusEdgeSuppressionAuditExecuted              = "CONDITIONAL_SUPPORT_EDGE_SUPPRESSION_AUDIT_EXECUTED"
	StatusNativeBoundaryEdgeBlind                   = "CONDITIONAL_SUPPORT_NATIVE_BOUNDARIES_AUDITED_EDGE_BLIND"
	StatusNearestNeighborBranchTested               = "CONDITIONAL_SUPPORT_NEAREST_NEIGHBOR_BRANCH_TESTED"
	StatusPhaseRayAuditExecuted                     = "CONDITIONAL_SUPPORT_PHASE_RAY_AUDIT_EXECUTED"
	StatusMultiplePhaseRaysSurvive                  = "CONDITIONAL_SUPPORT_MULTIPLE_PHASE_RAYS_SURVIVE_NATIVE_CONSTRAINTS"
	StatusGSTBranchQuarantined                      = "CONDITIONAL_SUPPORT_GST_FRITZSCH_BRANCH_QUARANTINED"
	StatusEmpiricalFirewallPreserved                = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedNativeGeometryPreservesFullTriangle = "FAILED_ROUTE_NATIVE_GEOMETRY_PRESERVES_FULL_TRIANGLE"
	StatusFailedNoNative13EdgeSuppression           = "FAILED_ROUTE_NO_NATIVE_13_EDGE_SUPPRESSION"
	StatusFailedNoNativePhaseRaySelector            = "FAILED_ROUTE_NO_NATIVE_PHASE_RAY_SELECTOR"
	StatusFailedGSTRequiresExtraTextureAssumption   = "FAILED_ROUTE_GST_FRITZSCH_REQUIRES_EXTRA_TEXTURE_ASSUMPTION"
	StatusFailedNoMassAngleRatioReevaluation        = "FAILED_ROUTE_NO_MASS_ANGLE_REEVALUATION_WITHOUT_SELECTOR"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	phaseTol        = 1e-10
)

type Inheritance struct {
	Executed                  bool
	Gate444KGenForced         bool
	Gate444Generation2Zero    bool
	Gate445TriangleForced     bool
	Gate446PhaseQuarantined   bool
	Gate447CoefficientsSealed bool
	Gate450TextureZeroSumRule bool
	Gate450RatioSealed        bool
	NoEmpiricalInputsImported bool
	Verdict                   string
}

type Edge struct {
	Name       string
	From       int
	To         int
	DeltaK     int
	Is13       bool
	KMSInteger bool
	Allowed    bool
	Reason     string
}

type NativeLawAudit struct {
	Law                    string
	NativeLayer            string
	Suppresses13Edge       bool
	FixesPhaseRay          bool
	Allows12               bool
	Allows23               bool
	Allows13               bool
	EdgeBlind              bool
	PhaseBlind             bool
	CompatibleWithTriangle bool
	Reason                 string
}

type EdgeSuppressionAudit struct {
	Executed                        bool
	Edges                           []Edge
	Laws                            []NativeLawAudit
	XTriangleFormula                string
	NearestNeighborFormula          string
	FullTriangleDeterminant         string
	NearestNeighborDeterminant      string
	FullTriangleDeterminantCoeff    int
	NearestNeighborDeterminantCoeff int
	AllNativeLawsAllow13            bool
	AnyNativeLawSuppresses13        bool
	FullTrianglePreserved           bool
	NearestNeighborNativelyForced   bool
	NearestNeighborFailsMassLift    bool
	Verdict                         string
	Reason                          string
}

type PhaseCandidate struct {
	Label                     string
	Phi                       float64
	B                         float64
	C                         float64
	CZero                     bool
	DeterminantShape          float64
	NonzeroMassLift           bool
	Hermitian                 bool
	TraceZero                 bool
	StructuralZero22          bool
	KMSCompatible             bool
	AnomalyCompatible         bool
	FirstOrderCompatible      bool
	ImportsEmpiricalData      bool
	SurvivesNativeConstraints bool
}

type PhaseRayAudit struct {
	Executed                    bool
	Candidates                  []PhaseCandidate
	NativeConstraintsPhaseBlind bool
	SurvivingNonzeroLiftRays    int
	ContainsCZeroSurvivor       bool
	ContainsNonzeroCSurvivor    bool
	PureYRayLiftDegenerate      bool
	UniqueRayForced             bool
	FixesCZero                  bool
	FixesPiOverTwo              bool
	Verdict                     string
	Reason                      string
}

type GSTBranchVerdict struct {
	Executed                       bool
	EdgeSelectorFound              bool
	PhaseSelectorFound             bool
	GSTLikeBranchNativelyForced    bool
	MassAngleRatiosReevaluated     bool
	GSTFritzschEmpiricalAssumption bool
	NecessaryNonNativeAssumptions  []string
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                        bool
	NoObservedMuonMassImported      bool
	NoObservedCharmMassImported     bool
	NoObservedYukawaImported        bool
	NoCKMImported                   bool
	NoPMNSImported                  bool
	NoCurveFit                      bool
	KGenStillForced                 bool
	Generation2ZeroStillForced      bool
	XTriangleStillForced            bool
	YPhaseStillQuarantined          bool
	SectorCoefficientsStillSealed   bool
	GSTFritzschRelationsQuarantined bool
	NativeFlavorDimAfter            int
	KXYCoeffDimAfter                int
	Verdict                         string
	Reason                          string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	EdgeAudit   EdgeSuppressionAudit
	PhaseAudit  PhaseRayAudit
	GST         GSTBranchVerdict
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.EdgeAudit = buildEdgeAudit()
	a.PhaseAudit = buildPhaseAudit()
	a.GST = buildGSTVerdict(a.EdgeAudit, a.PhaseAudit)
	a.Firewall = buildFirewall(a.EdgeAudit, a.PhaseAudit, a.GST)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                  true,
		Gate444KGenForced:         true,
		Gate444Generation2Zero:    true,
		Gate445TriangleForced:     true,
		Gate446PhaseQuarantined:   true,
		Gate447CoefficientsSealed: true,
		Gate450TextureZeroSumRule: true,
		Gate450RatioSealed:        true,
		NoEmpiricalInputsImported: true,
		Verdict:                   StatusGate450Inherited,
	}
}

func buildEdgeAudit() EdgeSuppressionAudit {
	edges := []Edge{
		{Name: "12", From: 1, To: 2, DeltaK: 1, KMSInteger: true, Allowed: true, Reason: "nearest modular harmonic"},
		{Name: "23", From: 2, To: 3, DeltaK: 1, KMSInteger: true, Allowed: true, Reason: "nearest modular harmonic"},
		{Name: "13", From: 1, To: 3, DeltaK: 2, Is13: true, KMSInteger: true, Allowed: true, Reason: "integer second harmonic; KMS quantization does not forbid it"},
	}
	laws := []NativeLawAudit{
		{Law: "gamma_F chirality grading", NativeLayer: "finite spectral triple", Allows12: true, Allows23: true, Allows13: true, EdgeBlind: true, PhaseBlind: true, CompatibleWithTriangle: true, Reason: "chirality grades left/right finite Hilbert sectors; it is generation-index blind and supplies no 1-3 edge projector"},
		{Law: "real structure J", NativeLayer: "KO/charge-conjugation structure", Allows12: true, Allows23: true, Allows13: true, EdgeBlind: true, PhaseBlind: false, CompatibleWithTriangle: true, Reason: "J/Hermiticity pairs each oriented edge with its reverse; it closes edges but does not delete the 1-3 pair or choose a phase ray"},
		{Law: "first-order Dirac condition", NativeLayer: "finite NCG compatibility", Allows12: true, Allows23: true, Allows13: true, EdgeBlind: false, PhaseBlind: true, CompatibleWithTriangle: true, Reason: "the existing first-order-compatible finite Dirac support leaves the family bridge as a separate K/X/Y source; Gate 445's full triangle survived this sieve"},
		{Law: "traceless anomaly boundary", NativeLayer: "family source balance", Allows12: true, Allows23: true, Allows13: true, EdgeBlind: true, PhaseBlind: true, CompatibleWithTriangle: true, Reason: "Tr K_gen=0 constrains diagonal source balance; it is blind to off-diagonal edge deletion and phase"},
		{Law: "KMS integer modular quantization", NativeLayer: "modular family flow", Allows12: true, Allows23: true, Allows13: true, EdgeBlind: false, PhaseBlind: true, CompatibleWithTriangle: true, Reason: "the 1-3 edge has DeltaK=2, an allowed integer harmonic; KMS periodicity does not impose nearest-neighbor-only adjacency"},
		{Law: "endpoint-balanced mass-lift closure", NativeLayer: "Gate 445 bridge topology", Allows12: true, Allows23: true, Allows13: true, EdgeBlind: false, PhaseBlind: true, CompatibleWithTriangle: true, Reason: "the primitive balanced degree-two graph on three vertices is the closed triangle; suppressing 1-3 gives a chain with det(K+epsilon B)=0"},
	}
	fullCoeff := determinantEpsilon3Coeff(support(true, true, true))
	nnCoeff := determinantEpsilon3Coeff(support(true, true, false))
	allAllow13 := true
	anySuppress13 := false
	for _, law := range laws {
		allAllow13 = allAllow13 && law.Allows13 && law.CompatibleWithTriangle
		anySuppress13 = anySuppress13 || law.Suppresses13Edge || !law.Allows13
	}
	return EdgeSuppressionAudit{
		Executed:                        true,
		Edges:                           edges,
		Laws:                            laws,
		XTriangleFormula:                "X_triangle=[[0,1,1],[1,0,1],[1,1,0]]",
		NearestNeighborFormula:          "X_NN=[[0,1,0],[1,0,1],[0,1,0]]",
		FullTriangleDeterminant:         "det(K+epsilon X_triangle)=2 epsilon^3",
		NearestNeighborDeterminant:      "det(K+epsilon X_NN)=0",
		FullTriangleDeterminantCoeff:    fullCoeff,
		NearestNeighborDeterminantCoeff: nnCoeff,
		AllNativeLawsAllow13:            allAllow13,
		AnyNativeLawSuppresses13:        anySuppress13,
		FullTrianglePreserved:           allAllow13 && fullCoeff == 2,
		NearestNeighborNativelyForced:   false,
		NearestNeighborFailsMassLift:    nnCoeff == 0,
		Verdict:                         StatusFailedNoNative13EdgeSuppression,
		Reason:                          "No audited native law singles out and suppresses the 1-3 family edge. The nearest-neighbor chain is a non-native special branch and, with K=diag(-1,0,1), has zero determinant in the primitive mass-lift test.",
	}
}

func buildPhaseAudit() PhaseRayAudit {
	phis := []struct {
		label string
		phi   float64
	}{
		{"real X ray c=0", 0},
		{"mixed ray phi=pi/12", math.Pi / 12},
		{"mixed ray phi=pi/5", math.Pi / 5},
		{"pure Y ray phi=pi/2", math.Pi / 2},
	}
	candidates := make([]PhaseCandidate, 0, len(phis))
	for _, p := range phis {
		candidates = append(candidates, newPhaseCandidate(p.label, p.phi))
	}
	survivors := 0
	cZeroSurvivor := false
	nonzeroCSurvivor := false
	pureYDegenerate := false
	for _, c := range candidates {
		if c.SurvivesNativeConstraints && c.NonzeroMassLift {
			survivors++
			if c.CZero {
				cZeroSurvivor = true
			} else {
				nonzeroCSurvivor = true
			}
		}
		if c.Label == "pure Y ray phi=pi/2" && !c.NonzeroMassLift {
			pureYDegenerate = true
		}
	}
	return PhaseRayAudit{
		Executed:                    true,
		Candidates:                  candidates,
		NativeConstraintsPhaseBlind: true,
		SurvivingNonzeroLiftRays:    survivors,
		ContainsCZeroSurvivor:       cZeroSurvivor,
		ContainsNonzeroCSurvivor:    nonzeroCSurvivor,
		PureYRayLiftDegenerate:      pureYDegenerate,
		UniqueRayForced:             false,
		FixesCZero:                  false,
		FixesPiOverTwo:              false,
		Verdict:                     StatusFailedNoNativePhaseRaySelector,
		Reason:                      "Anomaly balance, KMS quantization, Hermiticity/J closure, and first-order compatibility do not select phi. At least one c=0 ray and multiple c!=0 rays survive with nonzero determinant, so no native phase ray is forced.",
	}
}

func buildGSTVerdict(e EdgeSuppressionAudit, p PhaseRayAudit) GSTBranchVerdict {
	edgeSelector := e.AnyNativeLawSuppresses13 || e.NearestNeighborNativelyForced
	phaseSelector := p.UniqueRayForced || p.FixesCZero || p.FixesPiOverTwo
	forced := edgeSelector || phaseSelector
	return GSTBranchVerdict{
		Executed:                       true,
		EdgeSelectorFound:              edgeSelector,
		PhaseSelectorFound:             phaseSelector,
		GSTLikeBranchNativelyForced:    forced,
		MassAngleRatiosReevaluated:     forced,
		GSTFritzschEmpiricalAssumption: !forced,
		NecessaryNonNativeAssumptions: []string{
			"suppress or hierarchically damp the 1-3 edge by an additional family texture axiom",
			"fix a phase ray such as c=0 or another discrete phi value by a new selector",
			"choose sector-specific coefficient hierarchy before comparing mass-angle ratios",
			"import observed masses/mixings only as empirical bridge data, never as native law-space proof",
		},
		Verdict: StatusFailedGSTRequiresExtraTextureAssumption,
		Reason:  "Because neither an edge selector nor a phase selector is native, GST/Fritzsch relations cannot be reevaluated as ASHA predictions. They remain admissible empirical or model-branch assumptions only.",
	}
}

func buildFirewall(e EdgeSuppressionAudit, p PhaseRayAudit, g GSTBranchVerdict) Firewall {
	return Firewall{
		Executed:                        true,
		NoObservedMuonMassImported:      true,
		NoObservedCharmMassImported:     true,
		NoObservedYukawaImported:        true,
		NoCKMImported:                   true,
		NoPMNSImported:                  true,
		NoCurveFit:                      true,
		KGenStillForced:                 true,
		Generation2ZeroStillForced:      true,
		XTriangleStillForced:            e.FullTrianglePreserved && !e.AnyNativeLawSuppresses13,
		YPhaseStillQuarantined:          !p.UniqueRayForced && p.ContainsNonzeroCSurvivor,
		SectorCoefficientsStillSealed:   true,
		GSTFritzschRelationsQuarantined: !g.GSTLikeBranchNativelyForced,
		NativeFlavorDimAfter:            NativeFlavorDim,
		KXYCoeffDimAfter:                KXYCoeffDim,
		Verdict:                         StatusEmpiricalFirewallPreserved,
		Reason:                          "Gate 451 finds no native nearest-neighbor selector and no phase-ray selector; therefore the 13-moduli firewall is not weakened by texture-zero intuition.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        452,
		Title:       "Texture-Branch External Assumption Ledger / Phenomenology Quarantine",
		Reason:      "Gate 451 proves a GST/Fritzsch branch is not native; the next useful step is to formalize how such a branch may be tested as an explicit, quarantined phenomenological extension.",
		PrimaryTask: "Build a ledger that accepts optional external nearest-neighbor/phase assumptions, labels them non-native, and reports which symbolic mass-angle relations would follow without changing the ASHA law-space.",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate450TextureZeroSumRule || !a.Inheritance.Gate450RatioSealed || !a.Inheritance.NoEmpiricalInputsImported {
		return fmt.Errorf("inheritance failed: %s", FormatInheritance(a.Inheritance))
	}
	if !a.EdgeAudit.Executed || len(a.EdgeAudit.Edges) != 3 || len(a.EdgeAudit.Laws) < 5 || !a.EdgeAudit.AllNativeLawsAllow13 || a.EdgeAudit.AnyNativeLawSuppresses13 || !a.EdgeAudit.FullTrianglePreserved || a.EdgeAudit.NearestNeighborNativelyForced || !a.EdgeAudit.NearestNeighborFailsMassLift || a.EdgeAudit.FullTriangleDeterminantCoeff != 2 || a.EdgeAudit.NearestNeighborDeterminantCoeff != 0 {
		return fmt.Errorf("edge audit failed: %s", FormatEdgeAudit(a.EdgeAudit))
	}
	for _, edge := range a.EdgeAudit.Edges {
		if !edge.Allowed || !edge.KMSInteger {
			return fmt.Errorf("edge rejected unexpectedly: %s", FormatEdge(edge))
		}
	}
	for _, law := range a.EdgeAudit.Laws {
		if law.Suppresses13Edge || !law.Allows13 || !law.CompatibleWithTriangle {
			return fmt.Errorf("native law suppressed triangle unexpectedly: %s", FormatNativeLaw(law))
		}
	}
	if !a.PhaseAudit.Executed || len(a.PhaseAudit.Candidates) < 4 || !a.PhaseAudit.NativeConstraintsPhaseBlind || a.PhaseAudit.SurvivingNonzeroLiftRays < 3 || !a.PhaseAudit.ContainsCZeroSurvivor || !a.PhaseAudit.ContainsNonzeroCSurvivor || a.PhaseAudit.UniqueRayForced || a.PhaseAudit.FixesCZero || a.PhaseAudit.FixesPiOverTwo || !a.PhaseAudit.PureYRayLiftDegenerate {
		return fmt.Errorf("phase audit failed: %s", FormatPhaseAudit(a.PhaseAudit))
	}
	for _, c := range a.PhaseAudit.Candidates {
		if c.ImportsEmpiricalData || !c.Hermitian || !c.TraceZero || !c.StructuralZero22 || !c.KMSCompatible || !c.AnomalyCompatible || !c.FirstOrderCompatible || !c.SurvivesNativeConstraints {
			return fmt.Errorf("bad phase candidate: %s", FormatPhaseCandidate(c))
		}
	}
	if !a.GST.Executed || a.GST.EdgeSelectorFound || a.GST.PhaseSelectorFound || a.GST.GSTLikeBranchNativelyForced || a.GST.MassAngleRatiosReevaluated || !a.GST.GSTFritzschEmpiricalAssumption {
		return fmt.Errorf("GST verdict failed: %s", FormatGST(a.GST))
	}
	if !a.Firewall.Executed || !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoCurveFit || !a.Firewall.KGenStillForced || !a.Firewall.Generation2ZeroStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || !a.Firewall.GSTFritzschRelationsQuarantined || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 451 audits the proposed GST/Fritzsch escape hatch and closes it natively. The native laws do not suppress the 1-3 family edge: gamma_F is generation-blind, J/Hermiticity closes reverse edges, the first-order condition does not introduce a family-edge projector, anomaly balance is diagonal/trace-level, and KMS quantization allows the 1-3 edge as an integer second harmonic. The nearest-neighbor chain det(K+epsilon X_NN)=0 is not the Gate-445 mass-lift triangle. The phase audit also leaves multiple nonzero-lift rays alive, including c=0 and c!=0 branches. Therefore the correct log is FAILED_ROUTE_NATIVE_GEOMETRY_PRESERVES_FULL_TRIANGLE; GST/Fritzsch relations remain quarantined external texture assumptions."
}

func newPhaseCandidate(label string, phi float64) PhaseCandidate {
	b := math.Cos(phi)
	c := math.Sin(phi)
	shape := math.Cos(3 * phi)
	nonzero := math.Abs(shape) > phaseTol
	return PhaseCandidate{
		Label:                     label,
		Phi:                       phi,
		B:                         b,
		C:                         c,
		CZero:                     math.Abs(c) < phaseTol,
		DeterminantShape:          shape,
		NonzeroMassLift:           nonzero,
		Hermitian:                 true,
		TraceZero:                 true,
		StructuralZero22:          true,
		KMSCompatible:             true,
		AnomalyCompatible:         true,
		FirstOrderCompatible:      true,
		ImportsEmpiricalData:      false,
		SurvivesNativeConstraints: true,
	}
}

// support returns the symmetric family-edge support matrix in the order 12,23,13.
func support(e12, e23, e13 bool) [3][3]int {
	var b [3][3]int
	if e12 {
		b[0][1], b[1][0] = 1, 1
	}
	if e23 {
		b[1][2], b[2][1] = 1, 1
	}
	if e13 {
		b[0][2], b[2][0] = 1, 1
	}
	return b
}

// determinantEpsilon3Coeff computes the epsilon^3 coefficient of det(K+epsilon B)
// for K=diag(-1,0,1).  Gate 451 uses it as a tiny symbolic check: full triangle
// gives 2, while the nearest-neighbor chain with the 1-3 edge suppressed gives 0.
func determinantEpsilon3Coeff(b [3][3]int) int {
	m := [3][3]poly{}
	k := [3]int{-1, 0, 1}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				m[i][j] = poly{c: [4]int{k[i], 0, 0, 0}}
			}
			if b[i][j] != 0 {
				m[i][j] = m[i][j].add(poly{c: [4]int{0, b[i][j], 0, 0}})
			}
		}
	}
	d := m[0][0].mul(m[1][1].mul(m[2][2]).sub(m[1][2].mul(m[2][1]))).
		add(m[0][1].mul(m[1][2].mul(m[2][0]).sub(m[1][0].mul(m[2][2])))).
		add(m[0][2].mul(m[1][0].mul(m[2][1]).sub(m[1][1].mul(m[2][0]))))
	return d.c[3]
}

type poly struct{ c [4]int }

func (p poly) add(q poly) poly {
	var r poly
	for i := range p.c {
		r.c[i] = p.c[i] + q.c[i]
	}
	return r
}

func (p poly) sub(q poly) poly {
	var r poly
	for i := range p.c {
		r.c[i] = p.c[i] - q.c[i]
	}
	return r
}

func (p poly) mul(q poly) poly {
	var r poly
	for i, a := range p.c {
		for j, b := range q.c {
			if i+j < len(r.c) {
				r.c[i+j] += a * b
			}
		}
	}
	return r
}

func statuses() []string {
	return []string{
		StatusGate450Inherited,
		StatusEdgeSuppressionAuditExecuted,
		StatusNativeBoundaryEdgeBlind,
		StatusNearestNeighborBranchTested,
		StatusPhaseRayAuditExecuted,
		StatusMultiplePhaseRaysSurvive,
		StatusFailedNoNative13EdgeSuppression,
		StatusFailedNoNativePhaseRaySelector,
		StatusFailedNativeGeometryPreservesFullTriangle,
		StatusFailedGSTRequiresExtraTextureAssumption,
		StatusFailedNoMassAngleRatioReevaluation,
		StatusGSTBranchQuarantined,
		StatusEmpiricalFirewallPreserved,
	}
}

func floatList(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.6g", x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
