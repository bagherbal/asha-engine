// Package finitespectraltriple implements Gate 217: finite spectral triple /
// heavy-sector gauge-curvature projection audit.
//
// Gate 216 proved that the Gate-215 matching residual cannot be obtained by
// fitting raw finite scalars. Gate 217 audits the stronger machinery required
// by a genuine finite spectral-action derivation: a finite Hilbert carrier for
// the sealed heavy sector, a canonical finite Dirac operator D_F, real and
// grading structures, a gauge-curvature projection/heat-kernel map, and a
// cutoff/subtraction scheme. The gate intentionally refuses to invent any of
// these ingredients by hand.
package finitespectraltriple

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/matchingresidualstructure"
)

const (
	StatusFailedRoute        = "FAILED_ROUTE_FINITE_SPECTRAL_TRIPLE_MATCHING_DERIVATION"
	StatusConditionalSupport = "CONDITIONAL_SUPPORT_FINITE_SPECTRAL_TRIPLE_MATCHING"

	DiracMissingOperator = "FINITE_DIRAC_OPERATOR_NOT_DERIVED"
	ProjectionMissing    = "GAUGE_CURVATURE_PROJECTION_NOT_DERIVED"
	CutoffMissing        = "CUTOFF_SUBTRACTION_SCHEME_NOT_DERIVED"
)

type FloatTriple struct{ U1GUT, SU2L, SU3C float64 }

func (t FloatTriple) String() string {
	return fmt.Sprintf("(%.12g,%.12g,%.12g)", t.U1GUT, t.SU2L, t.SU3C)
}
func (t FloatTriple) MaxAbs() float64 {
	return math.Max(math.Abs(t.U1GUT), math.Max(math.Abs(t.SU2L), math.Abs(t.SU3C)))
}
func (t FloatTriple) NormalizedMax() FloatTriple {
	m := t.MaxAbs()
	if m == 0 || math.IsNaN(m) || math.IsInf(m, 0) {
		return FloatTriple{}
	}
	return FloatTriple{t.U1GUT / m, t.SU2L / m, t.SU3C / m}
}

type Gate216Snapshot struct {
	Gate216Inherited                  bool
	Gate215SingleScaleTargetInherited bool
	ThresholdSpectrumSealInherited    bool
	MatchingEnvelopeInherited         bool
	SpectralResidualTarget            FloatTriple
	SpectralResidualNormalized        FloatTriple
	ResidualSignPattern               string
	BestRows                          string
	MBGeV                             float64
	MStarGeV                          float64
	BGapAvailable                     bool
	ContactModes                      int
	SignOnlyResonances                int
	StructuralMatches                 int
	FiniteDiracOperatorDerived        bool
	HeatKernelMapDerived              bool
	CanonicalCutoffDerived            bool
	FiniteMatchingRowsDerived         bool
	TruthStatement                    string
}

type HeavyRepresentation struct {
	Name                  string
	Representation        string
	Statistics            string
	InternalDimension     int
	DiracChiralCarrierDim int
	BetaRow               FloatTriple
	FromThresholdSeal     bool
	FiniteCoreDerived     bool
	GaugeRepresentationOK bool
	MassFromFiniteCore    bool
	LocalFieldMapDerived  bool
	FiniteHilbertEmbedded bool
	Verdict               string
}

type HeavyHilbertAudit struct {
	SpectrumSealInherited      bool
	RepresentationsAudited     int
	InternalDimensionTotal     int
	DiracChiralCarrierDimTotal int
	FiniteAlgebraCarrierKnown  bool
	FiniteHilbertSpaceDerived  bool
	InnerProductDerived        bool
	RealStructureDerived       bool
	GradingDerived             bool
	HeavyChargeConjugation     bool
	UsesPhenomenologicalRows   bool
	UsesObservedMassScale      bool
	Verdict                    string
}

type DiracCandidate struct {
	Name                          string
	Formula                       string
	ActsOnHeavySector             bool
	SelfAdjointIfChosen           bool
	InvertibleIfChosen            bool
	GaugeIntertwiner              bool
	OddWithGrading                bool
	OrderOneTestable              bool
	OrderOneVerified              bool
	MassScaleFiniteDerived        bool
	StructureCliffordG2Dictated   bool
	RequiresHandChosenMatrix      bool
	RequiresPhenomenologicalScale bool
	PromotableFiniteDirac         bool
	Verdict                       string
}

type FiniteDiracOperatorAudit struct {
	CandidatesAudited          int
	SelfAdjointCandidates      int
	InvertibleCandidates       int
	GaugeIntertwinerCandidates int
	OrderOneVerified           int
	MassScaleFiniteDerived     int
	CliffordG2Dictated         int
	PromotableFiniteDirac      int
	MissingPiece               string
	Verdict                    string
}

type HeatKernelProjectionAudit struct {
	A2A4LanguageAudited                 bool
	FiniteSpectralTripleComplete        bool
	GaugeFluctuationMapDerived          bool
	CurvatureComponents                 []string
	RepresentationTraceRowsKnown        int
	GaugeCurvatureProjectionRowsDerived int
	A2CoefficientsDerived               int
	A4GaugeCoefficientsDerived          int
	CanProduceVectorSignStructure       bool
	CanProduceMatchingMagnitude         bool
	TrDFMinus2Computed                  bool
	TrDFMinus2Promoted                  bool
	RequiredResidual                    FloatTriple
	RequiredResidualNormalized          FloatTriple
	ProjectedDeltaMatchRows             int
	MissingPiece                        string
	Verdict                             string
}

type CutoffSubtractionAudit struct {
	CutoffFunctionDerived             bool
	CutoffMomentsDerived              bool
	RenormalizationSchemeDerived      bool
	ThresholdSubtractionRuleDerived   bool
	MSbarImported                     bool
	DimensionalRegularizationImported bool
	SchemeDependentConstantFixed      bool
	FiniteCountertermFunctional       bool
	PhysicalDeltaMatchRows            int
	MissingPiece                      string
	Verdict                           string
}

type MatchingCorrectionReadinessAudit struct {
	FiniteDiracReady         bool
	GaugeProjectionReady     bool
	CutoffSubtractionReady   bool
	MatchingRowsDerived      bool
	CanEvaluateGate215Target bool
	CanDeriveDeltaMatch      bool
	CanOnlyStateTarget       bool
	RequiredNextIngredient   string
	Verdict                  string
}

type FirewallAudit struct {
	Gate216Inherited                bool
	ThresholdSpectrumSealInherited  bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	EmpiricalLedgerQuarantined      bool
	DFFittedByHand                  bool
	CutoffFunctionInvented          bool
	MSbarSchemeImported             bool
	HeatKernelProjectionFitted      bool
	MatchingResidualPromoted        bool
	MatchingCorrectionsDerived      bool
	HeavyMassesFiniteDerived        bool
	PhysicalUnificationClaimed      bool
	ContactModesPromotedToParticles bool
	BGapPromotedToMass              bool
	ProtonLifetimeComputed          bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Summary struct {
	TestsAudited                int
	Gate216Inherited            bool
	HeavyRepresentationsAudited int
	DiracCandidatesAudited      int
	PromotableFiniteDirac       int
	GaugeProjectionRows         int
	CutoffSchemesDerived        int
	MatchingRowsDerived         int
	BlockingPieces              []string
	Status                      string
	Comment                     string
}

type Analysis struct {
	Gate216         Gate216Snapshot
	Gate216Analysis matchingresidualstructure.Analysis
	Representations []HeavyRepresentation
	Hilbert         HeavyHilbertAudit
	DiracCandidates []DiracCandidate
	DiracAudit      FiniteDiracOperatorAudit
	HeatKernel      HeatKernelProjectionAudit
	Cutoff          CutoffSubtractionAudit
	Readiness       MatchingCorrectionReadinessAudit
	Firewall        FirewallAudit
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
		g216, err := matchingresidualstructure.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 216 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g216)
	})
	return defaultA, defaultErr
}

func Build(g216 matchingresidualstructure.Analysis) (Analysis, error) {
	snap := snapshotFromGate216(g216)
	if !snap.Gate216Inherited || !snap.Gate215SingleScaleTargetInherited || !snap.ThresholdSpectrumSealInherited {
		return Analysis{}, fmt.Errorf("Gate 217 requires Gate 216 failed-route residual target under ThresholdSpectrumSeal")
	}
	reps := sealedHeavyRepresentations()
	hilbert := auditHeavyHilbert(snap, reps)
	cands := diracCandidates(hilbert)
	diracAudit := auditDirac(cands)
	heat := auditHeatKernel(snap, diracAudit)
	cutoff := auditCutoffSubtraction()
	ready := auditReadiness(diracAudit, heat, cutoff)
	fw := auditFirewall(snap, diracAudit, heat, cutoff, ready)
	status := StatusFailedRoute
	if ready.CanDeriveDeltaMatch {
		status = StatusConditionalSupport
	}
	summary := Summary{
		TestsAudited:                7,
		Gate216Inherited:            snap.Gate216Inherited,
		HeavyRepresentationsAudited: len(reps),
		DiracCandidatesAudited:      len(cands),
		PromotableFiniteDirac:       diracAudit.PromotableFiniteDirac,
		GaugeProjectionRows:         heat.GaugeCurvatureProjectionRowsDerived,
		CutoffSchemesDerived:        boolToInt(cutoff.CutoffFunctionDerived && cutoff.ThresholdSubtractionRuleDerived),
		MatchingRowsDerived:         cutoff.PhysicalDeltaMatchRows,
		BlockingPieces:              []string{diracAudit.MissingPiece, heat.MissingPiece, cutoff.MissingPiece},
		Status:                      status,
		Comment:                     "Gate 217 audits the finite spectral-triple machinery required to convert the Gate-215 target residual into derived threshold matching constants.",
	}
	truth := buildTruth(summary, snap, ready)
	return Analysis{Gate216: snap, Gate216Analysis: g216, Representations: reps, Hilbert: hilbert, DiracCandidates: cands, DiracAudit: diracAudit, HeatKernel: heat, Cutoff: cutoff, Readiness: ready, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate216(a matchingresidualstructure.Analysis) Gate216Snapshot {
	d := a.Gate215.RequiredDeltaMatch
	target := FloatTriple{d.U1GUT, d.SU2L, d.SU3C}
	return Gate216Snapshot{
		Gate216Inherited:                  a.Summary.Status == matchingresidualstructure.StatusFailedRoute,
		Gate215SingleScaleTargetInherited: a.Gate215.SingleScaleCandidateUnique && a.Gate215.PlausibleClasses == 1,
		ThresholdSpectrumSealInherited:    a.Firewall.ThresholdSpectrumSealInherited,
		MatchingEnvelopeInherited:         a.Gate215.MatchingEnvelopeInherited,
		SpectralResidualTarget:            target,
		SpectralResidualNormalized:        target.NormalizedMax(),
		ResidualSignPattern:               a.Gate215.RequiredSignPattern,
		BestRows:                          a.Gate215.BestRows,
		MBGeV:                             a.Gate215.MBGeV,
		MStarGeV:                          a.Gate215.MStarGeV,
		BGapAvailable:                     a.SpectralData.BGapAvailable,
		ContactModes:                      a.SpectralData.ContactPartialModeCount,
		SignOnlyResonances:                a.HeatKernelMap.SignOnlyResonances,
		StructuralMatches:                 a.HeatKernelMap.FullStructuralMatches,
		FiniteDiracOperatorDerived:        a.HeatKernelMap.FiniteDiracOperatorDerived,
		HeatKernelMapDerived:              a.HeatKernelMap.GaugeKineticTraceMapDerived,
		CanonicalCutoffDerived:            a.HeatKernelMap.CanonicalCutoffMomentsDerived,
		FiniteMatchingRowsDerived:         a.HeatKernelMap.DeltaMatchRowsDerived > 0 || a.SpectralData.FiniteMatchingRowsDerived,
		TruthStatement:                    a.TruthStatement,
	}
}

func sealedHeavyRepresentations() []HeavyRepresentation {
	return []HeavyRepresentation{
		{
			Name: "sealed Dirac weak triplet", Representation: "Dirac (1,3,Y=1)", Statistics: "Dirac fermion", InternalDimension: 3, DiracChiralCarrierDim: 6,
			BetaRow: FloatTriple{12.0 / 5.0, 8.0 / 3.0, 0}, FromThresholdSeal: true, FiniteCoreDerived: false,
			GaugeRepresentationOK: true, MassFromFiniteCore: false, LocalFieldMapDerived: false, FiniteHilbertEmbedded: false,
			Verdict: "representation row is permitted under the ThresholdSpectrumSeal, but its finite Hilbert embedding and mass are not derived",
		},
		{
			Name: "sealed Dirac color-octet weak doublet", Representation: "Dirac (8,2,Y=1/2)", Statistics: "Dirac fermion", InternalDimension: 16, DiracChiralCarrierDim: 32,
			BetaRow: FloatTriple{16.0 / 5.0, 16.0 / 3.0, 8}, FromThresholdSeal: true, FiniteCoreDerived: false,
			GaugeRepresentationOK: true, MassFromFiniteCore: false, LocalFieldMapDerived: false, FiniteHilbertEmbedded: false,
			Verdict: "representation row is permitted under the ThresholdSpectrumSeal, but its finite Hilbert embedding and mass are not derived",
		},
	}
}

func auditHeavyHilbert(s Gate216Snapshot, reps []HeavyRepresentation) HeavyHilbertAudit {
	internal, chiral := 0, 0
	for _, r := range reps {
		internal += r.InternalDimension
		chiral += r.DiracChiralCarrierDim
	}
	return HeavyHilbertAudit{
		SpectrumSealInherited:      s.ThresholdSpectrumSealInherited,
		RepresentationsAudited:     len(reps),
		InternalDimensionTotal:     internal,
		DiracChiralCarrierDimTotal: chiral,
		FiniteAlgebraCarrierKnown:  false,
		FiniteHilbertSpaceDerived:  false,
		InnerProductDerived:        false,
		RealStructureDerived:       false,
		GradingDerived:             false,
		HeavyChargeConjugation:     false,
		UsesPhenomenologicalRows:   true,
		UsesObservedMassScale:      false,
		Verdict:                    "the sealed heavy rows define continuum representation content, but not a finite spectral-triple Hilbert carrier, inner product, real structure, or grading",
	}
}

func diracCandidates(h HeavyHilbertAudit) []DiracCandidate {
	return []DiracCandidate{
		{
			Name: "zero heavy-sector D_F", Formula: "D_F = 0 on H_heavy", ActsOnHeavySector: true, SelfAdjointIfChosen: true, InvertibleIfChosen: false,
			GaugeIntertwiner: true, OddWithGrading: false, OrderOneTestable: false, OrderOneVerified: false, MassScaleFiniteDerived: false,
			StructureCliffordG2Dictated: false, RequiresHandChosenMatrix: false, RequiresPhenomenologicalScale: false, PromotableFiniteDirac: false,
			Verdict: "canonical only in the vacuous sense; it cannot generate Tr(D_F^-2), decoupling, or threshold matching constants",
		},
		{
			Name: "degenerate identity mass ansatz", Formula: "D_F = M_B · I on sealed Dirac carrier", ActsOnHeavySector: h.SpectrumSealInherited, SelfAdjointIfChosen: true, InvertibleIfChosen: true,
			GaugeIntertwiner: true, OddWithGrading: false, OrderOneTestable: false, OrderOneVerified: false, MassScaleFiniteDerived: false,
			StructureCliffordG2Dictated: false, RequiresHandChosenMatrix: true, RequiresPhenomenologicalScale: true, PromotableFiniteDirac: false,
			Verdict: "would be a valid phenomenological mass matrix, but the finite algebra does not dictate M_B or the identity operator as D_F",
		},
		{
			Name: "off-diagonal triplet/octet intertwiner", Formula: "D_F: (1,3,1) ↔ (8,2,1/2)", ActsOnHeavySector: true, SelfAdjointIfChosen: true, InvertibleIfChosen: false,
			GaugeIntertwiner: false, OddWithGrading: false, OrderOneTestable: false, OrderOneVerified: false, MassScaleFiniteDerived: false,
			StructureCliffordG2Dictated: false, RequiresHandChosenMatrix: true, RequiresPhenomenologicalScale: true, PromotableFiniteDirac: false,
			Verdict: "no gauge-equivariant finite intertwiner exists between inequivalent color/weak/hypercharge representations",
		},
		{
			Name: "top-down Fock spectral triple reuse", Formula: "reuse Gate-166 Fock D_F support for heavy rows", ActsOnHeavySector: false, SelfAdjointIfChosen: true, InvertibleIfChosen: false,
			GaugeIntertwiner: false, OddWithGrading: true, OrderOneTestable: true, OrderOneVerified: false, MassScaleFiniteDerived: false,
			StructureCliffordG2Dictated: false, RequiresHandChosenMatrix: true, RequiresPhenomenologicalScale: false, PromotableFiniteDirac: false,
			Verdict: "the Fock support is a representation-trace certificate for the SM seed, not a heavy-sector Dirac operator or matching functional",
		},
	}
}

func auditDirac(cands []DiracCandidate) FiniteDiracOperatorAudit {
	var selfAdj, inv, inter, order, scale, dictated, promotable int
	for _, c := range cands {
		if c.SelfAdjointIfChosen {
			selfAdj++
		}
		if c.InvertibleIfChosen {
			inv++
		}
		if c.GaugeIntertwiner {
			inter++
		}
		if c.OrderOneVerified {
			order++
		}
		if c.MassScaleFiniteDerived {
			scale++
		}
		if c.StructureCliffordG2Dictated {
			dictated++
		}
		if c.PromotableFiniteDirac {
			promotable++
		}
	}
	return FiniteDiracOperatorAudit{
		CandidatesAudited: len(cands), SelfAdjointCandidates: selfAdj, InvertibleCandidates: inv, GaugeIntertwinerCandidates: inter,
		OrderOneVerified: order, MassScaleFiniteDerived: scale, CliffordG2Dictated: dictated, PromotableFiniteDirac: promotable,
		MissingPiece: DiracMissingOperator,
		Verdict:      "no candidate is simultaneously finite-algebra-dictated, nontrivial, gauge-compatible, order-one verified, and equipped with a finite-derived mass scale",
	}
}

func auditHeatKernel(s Gate216Snapshot, d FiniteDiracOperatorAudit) HeatKernelProjectionAudit {
	return HeatKernelProjectionAudit{
		A2A4LanguageAudited:                 true,
		FiniteSpectralTripleComplete:        d.PromotableFiniteDirac > 0,
		GaugeFluctuationMapDerived:          false,
		CurvatureComponents:                 []string{"U(1)_Y curvature squared", "SU(2)_L curvature squared", "SU(3)_C curvature squared"},
		RepresentationTraceRowsKnown:        2,
		GaugeCurvatureProjectionRowsDerived: 0,
		A2CoefficientsDerived:               0,
		A4GaugeCoefficientsDerived:          0,
		CanProduceVectorSignStructure:       false,
		CanProduceMatchingMagnitude:         false,
		TrDFMinus2Computed:                  false,
		TrDFMinus2Promoted:                  false,
		RequiredResidual:                    s.SpectralResidualTarget,
		RequiredResidualNormalized:          s.SpectralResidualNormalized,
		ProjectedDeltaMatchRows:             0,
		MissingPiece:                        ProjectionMissing,
		Verdict:                             "representation traces and beta rows are known, but there is no D_A fluctuation or heat-kernel projection map from finite traces to gauge-specific δ_i^match rows",
	}
}

func auditCutoffSubtraction() CutoffSubtractionAudit {
	return CutoffSubtractionAudit{
		CutoffFunctionDerived:             false,
		CutoffMomentsDerived:              false,
		RenormalizationSchemeDerived:      false,
		ThresholdSubtractionRuleDerived:   false,
		MSbarImported:                     false,
		DimensionalRegularizationImported: false,
		SchemeDependentConstantFixed:      false,
		FiniteCountertermFunctional:       false,
		PhysicalDeltaMatchRows:            0,
		MissingPiece:                      CutoffMissing,
		Verdict:                           "without a canonical cutoff function, moments, and subtraction prescription, heat-kernel traces cannot be converted into physical finite threshold constants",
	}
}

func auditReadiness(d FiniteDiracOperatorAudit, h HeatKernelProjectionAudit, c CutoffSubtractionAudit) MatchingCorrectionReadinessAudit {
	fd := d.PromotableFiniteDirac > 0
	gp := h.GaugeCurvatureProjectionRowsDerived == 3 && h.ProjectedDeltaMatchRows > 0
	cs := c.CutoffFunctionDerived && c.ThresholdSubtractionRuleDerived && c.PhysicalDeltaMatchRows == 3
	return MatchingCorrectionReadinessAudit{
		FiniteDiracReady:         fd,
		GaugeProjectionReady:     gp,
		CutoffSubtractionReady:   cs,
		MatchingRowsDerived:      fd && gp && cs,
		CanEvaluateGate215Target: false,
		CanDeriveDeltaMatch:      fd && gp && cs,
		CanOnlyStateTarget:       true,
		RequiredNextIngredient:   "derive D_F, J, gamma, gauge fluctuation, cutoff moments, and threshold subtraction in one finite spectral-action theorem",
		Verdict:                  "the Gate-215 residual remains a target vector; Gate 217 cannot derive matching corrections with the current finite machinery",
	}
}

func auditFirewall(s Gate216Snapshot, d FiniteDiracOperatorAudit, h HeatKernelProjectionAudit, c CutoffSubtractionAudit, r MatchingCorrectionReadinessAudit) FirewallAudit {
	return FirewallAudit{
		Gate216Inherited:                s.Gate216Inherited,
		ThresholdSpectrumSealInherited:  s.ThresholdSpectrumSealInherited,
		EmpiricalCarrierSealInherited:   true,
		LeptoquarkDynamicsSealInherited: true,
		EmpiricalLedgerQuarantined:      true,
		DFFittedByHand:                  false,
		CutoffFunctionInvented:          false,
		MSbarSchemeImported:             c.MSbarImported,
		HeatKernelProjectionFitted:      false,
		MatchingResidualPromoted:        false,
		MatchingCorrectionsDerived:      r.CanDeriveDeltaMatch,
		HeavyMassesFiniteDerived:        false,
		PhysicalUnificationClaimed:      false,
		ContactModesPromotedToParticles: false,
		BGapPromotedToMass:              false,
		ProtonLifetimeComputed:          false,
		RecommendedNextGate:             "Gate 218 — matching-correction seal / precision-uncertainty ledger, or finite heavy-sector Dirac operator construction attempt with explicit axioms",
		OpenRequirements: []string{
			"construct a finite Hilbert carrier for the sealed heavy representations",
			"derive a nontrivial self-adjoint D_F from Clifford/G2/contact data rather than M_B by hand",
			"derive real structure J, grading gamma, and a nontrivial order-one calculus",
			"derive gauge fluctuation and heat-kernel projection rows for U(1), SU(2), and SU(3)",
			"derive cutoff moments and a finite subtraction scheme before interpreting δ_i^match physically",
		},
		Verdict: "all spectral-action ingredients remain audited as missing; no hand-built D_F, cutoff, or projection is imported",
	}
}

func buildTruth(s Summary, g Gate216Snapshot, r MatchingCorrectionReadinessAudit) string {
	if s.Status == StatusConditionalSupport {
		return "Gate 217 constructed a finite spectral triple and derived matching corrections."
	}
	return fmt.Sprintf("Gate 217 audits the finite spectral-triple machinery needed to derive the Gate-215 residual %s for %s. The sealed heavy rows can be named as continuum representations, but the finite algebra does not yet supply the heavy Hilbert carrier, D_F, J, gamma, gauge fluctuation, heat-kernel projection, cutoff function, or threshold subtraction scheme. Therefore δ_i^match remains an external target, not a derived spectral-action coefficient. Required next ingredient: %s.", g.SpectralResidualTarget.String(), g.BestRows, r.RequiredNextIngredient)
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func signPattern(t FloatTriple) string {
	var b strings.Builder
	for _, v := range []float64{t.U1GUT, t.SU2L, t.SU3C} {
		switch {
		case v > 0:
			b.WriteByte('+')
		case v < 0:
			b.WriteByte('-')
		default:
			b.WriteByte('0')
		}
	}
	return b.String()
}

func FormatGate216(s Gate216Snapshot) string {
	return fmt.Sprintf("inherited=%t targetInherited=%t seal=%t envelope=%t rows=%s target=%s normalized=%s pattern=%s signOnly=%d structural=%d finiteD=%t heatMap=%t cutoff=%t matchRows=%t", s.Gate216Inherited, s.Gate215SingleScaleTargetInherited, s.ThresholdSpectrumSealInherited, s.MatchingEnvelopeInherited, s.BestRows, s.SpectralResidualTarget.String(), s.SpectralResidualNormalized.String(), s.ResidualSignPattern, s.SignOnlyResonances, s.StructuralMatches, s.FiniteDiracOperatorDerived, s.HeatKernelMapDerived, s.CanonicalCutoffDerived, s.FiniteMatchingRowsDerived)
}

func FormatRepresentations(rs []HeavyRepresentation) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, fmt.Sprintf("%s dim=%d chiralDim=%d row=%s finite=%t hilbert=%t mass=%t", r.Representation, r.InternalDimension, r.DiracChiralCarrierDim, r.BetaRow.String(), r.FiniteCoreDerived, r.FiniteHilbertEmbedded, r.MassFromFiniteCore))
	}
	return strings.Join(parts, " | ")
}

func FormatHilbert(h HeavyHilbertAudit) string {
	return fmt.Sprintf("seal=%t reps=%d internalDim=%d chiralDim=%d finiteCarrier=%t H=%t inner=%t J=%t gamma=%t C=%t phenRows=%t verdict=%s", h.SpectrumSealInherited, h.RepresentationsAudited, h.InternalDimensionTotal, h.DiracChiralCarrierDimTotal, h.FiniteAlgebraCarrierKnown, h.FiniteHilbertSpaceDerived, h.InnerProductDerived, h.RealStructureDerived, h.GradingDerived, h.HeavyChargeConjugation, h.UsesPhenomenologicalRows, h.Verdict)
}

func FormatDiracCandidates(cs []DiracCandidate) string {
	parts := make([]string, 0, len(cs))
	for _, c := range cs {
		parts = append(parts, fmt.Sprintf("%s formula=%q selfAdj=%t inv=%t gaugeIntertwiner=%t orderOne=%t finiteMass=%t dictated=%t hand=%t promoted=%t", c.Name, c.Formula, c.SelfAdjointIfChosen, c.InvertibleIfChosen, c.GaugeIntertwiner, c.OrderOneVerified, c.MassScaleFiniteDerived, c.StructureCliffordG2Dictated, c.RequiresHandChosenMatrix, c.PromotableFiniteDirac))
	}
	return strings.Join(parts, " | ")
}

func FormatDiracAudit(d FiniteDiracOperatorAudit) string {
	return fmt.Sprintf("candidates=%d selfAdj=%d invertible=%d intertwiners=%d orderOne=%d finiteMass=%d dictated=%d promotable=%d missing=%s verdict=%s", d.CandidatesAudited, d.SelfAdjointCandidates, d.InvertibleCandidates, d.GaugeIntertwinerCandidates, d.OrderOneVerified, d.MassScaleFiniteDerived, d.CliffordG2Dictated, d.PromotableFiniteDirac, d.MissingPiece, d.Verdict)
}

func FormatHeatKernel(h HeatKernelProjectionAudit) string {
	return fmt.Sprintf("a2a4=%t triple=%t fluctuation=%t components=%s repRows=%d projectionRows=%d a2=%d a4=%d sign=%t mag=%t trDminus2=%t promoted=%t deltaRows=%d missing=%s verdict=%s", h.A2A4LanguageAudited, h.FiniteSpectralTripleComplete, h.GaugeFluctuationMapDerived, strings.Join(h.CurvatureComponents, ","), h.RepresentationTraceRowsKnown, h.GaugeCurvatureProjectionRowsDerived, h.A2CoefficientsDerived, h.A4GaugeCoefficientsDerived, h.CanProduceVectorSignStructure, h.CanProduceMatchingMagnitude, h.TrDFMinus2Computed, h.TrDFMinus2Promoted, h.ProjectedDeltaMatchRows, h.MissingPiece, h.Verdict)
}

func FormatCutoff(c CutoffSubtractionAudit) string {
	return fmt.Sprintf("cutoff=%t moments=%t scheme=%t subtraction=%t msbar=%t dimreg=%t constants=%t counterterm=%t deltaRows=%d missing=%s verdict=%s", c.CutoffFunctionDerived, c.CutoffMomentsDerived, c.RenormalizationSchemeDerived, c.ThresholdSubtractionRuleDerived, c.MSbarImported, c.DimensionalRegularizationImported, c.SchemeDependentConstantFixed, c.FiniteCountertermFunctional, c.PhysicalDeltaMatchRows, c.MissingPiece, c.Verdict)
}

func FormatReadiness(r MatchingCorrectionReadinessAudit) string {
	return fmt.Sprintf("D=%t projection=%t cutoff=%t rows=%t evalTarget=%t derive=%t targetOnly=%t next=%s verdict=%s", r.FiniteDiracReady, r.GaugeProjectionReady, r.CutoffSubtractionReady, r.MatchingRowsDerived, r.CanEvaluateGate215Target, r.CanDeriveDeltaMatch, r.CanOnlyStateTarget, r.RequiredNextIngredient, r.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate216=%t spectrumSeal=%t carrierSeal=%t lqSeal=%t ledger=%t handD=%t inventedCutoff=%t msbar=%t fittedProjection=%t residualPromoted=%t corrections=%t finiteMasses=%t unification=%t contactParticles=%t bGapMass=%t proton=%t next=%s", f.Gate216Inherited, f.ThresholdSpectrumSealInherited, f.EmpiricalCarrierSealInherited, f.LeptoquarkDynamicsSealInherited, f.EmpiricalLedgerQuarantined, f.DFFittedByHand, f.CutoffFunctionInvented, f.MSbarSchemeImported, f.HeatKernelProjectionFitted, f.MatchingResidualPromoted, f.MatchingCorrectionsDerived, f.HeavyMassesFiniteDerived, f.PhysicalUnificationClaimed, f.ContactModesPromotedToParticles, f.BGapPromotedToMass, f.ProtonLifetimeComputed, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d inherited=%t reps=%d diracCandidates=%d promotableD=%d projectionRows=%d cutoffSchemes=%d matchingRows=%d blocks=%s status=%s", s.TestsAudited, s.Gate216Inherited, s.HeavyRepresentationsAudited, s.DiracCandidatesAudited, s.PromotableFiniteDirac, s.GaugeProjectionRows, s.CutoffSchemesDerived, s.MatchingRowsDerived, strings.Join(s.BlockingPieces, ","), s.Status)
}
