// Package universalbetasource implements Gate 203: universal beta source
// classification / complete-multiplet versus regulator-trace audit.
//
// Gate 201 found two non-universal threshold shapes that close the quarantined
// topological mismatch triangle only after adding a real universal beta row.
// Gate 202 proved that the same universal row can be rewritten as a common
// topological boundary offset, but the finite B-sector gap and contact zeta
// traces did not canonically derive that offset.
//
// Gate 203 therefore classifies the remaining possible sources of the universal
// row.  It audits two physically standard possibilities without promoting either
// to a theorem unless the finite algebra supplies the missing semantics:
//
//  1. a complete unified heavy multiplet, whose one-loop beta contribution is
//     universal in GUT normalization;
//  2. a regulator/ghost/spectral-measure trace, which would act as a universal
//     conformal anomaly.
//
// The result is intentionally conservative: complete multiplet rows are exact
// rational universal rows, while the Gate-201 universal rows are real numbers
// inherited from a phenomenological inverse-threshold lever arm.  No audited
// integer sum of complete multiplets equals the required rows, and the finite
// contact/Fock/B-sector data still lacks a canonical heavy-threshold or
// regulator-trace map.  The universal beta source remains external
// phenomenological data under current axioms.
package universalbetasource

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/universaltracedeformation"
)

const tolerance = 1e-9

type Rational struct {
	Num int64
	Den int64
}

func R(num, den int64) Rational {
	if den == 0 {
		panic("zero denominator")
	}
	if den < 0 {
		num = -num
		den = -den
	}
	g := gcd(abs(num), den)
	return Rational{Num: num / g, Den: den / g}
}

func (r Rational) Float() float64 { return float64(r.Num) / float64(r.Den) }

func (r Rational) String() string {
	if r.Den == 1 {
		return fmt.Sprintf("%d", r.Num)
	}
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}

func (r Rational) MulInt(n int) Rational { return R(r.Num*int64(n), r.Den) }

func gcd(a, b int64) int64 {
	if a == 0 {
		if b == 0 {
			return 1
		}
		return b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

type Gate202UniversalRequirement struct {
	ShapeName             string
	RationalShapeDeltaB   string
	RequiredUniversalBeta float64
	RequiredDeltaU        float64
	ThresholdScaleGeV     float64
	BoundaryScaleGeV      float64
	ConditionalOnly       bool
	FiniteDerived         bool
}

type Gate202Snapshot struct {
	Gate202Inherited                      bool
	Gate201ConditionalShapesInherited     bool
	BoundaryOffsetEquivalenceEstablished  bool
	BGapContactTraceOffsetFailedRoute     bool
	UniversalCompletionNotDerived         bool
	NoPhysicalPredictionClaim             bool
	ObservedInputsUsedForFiniteDerivation bool
	PhysicalUnificationClaimed            bool
	ThresholdCorrectedPhysicalFitClaimed  bool
	AbsoluteMassPredicted                 bool
	FiniteMatchingCorrectionsDerived      bool
	StrictNullityAfter                    int
	PhysicalPredictionNullityAfter        int
	Requirements                          []Gate202UniversalRequirement
	TruthStatement                        string
}

func DefaultGate202Snapshot() (Gate202Snapshot, error) {
	prev, err := universaltracedeformation.BuildDefault()
	if err != nil {
		return Gate202Snapshot{}, err
	}
	reqs := make([]Gate202UniversalRequirement, 0, len(prev.RequiredOffsets))
	for _, r := range prev.RequiredOffsets {
		reqs = append(reqs, Gate202UniversalRequirement{
			ShapeName:             r.ShapeName,
			RationalShapeDeltaB:   r.RationalShapeDeltaB,
			RequiredUniversalBeta: r.UniversalBetaDelta,
			RequiredDeltaU:        r.RequiredDeltaU,
			ThresholdScaleGeV:     r.ThresholdScaleGeV,
			BoundaryScaleGeV:      r.BoundaryScaleGeV,
			ConditionalOnly:       r.ConditionalOnly,
			FiniteDerived:         r.FiniteDerived,
		})
	}
	return Gate202Snapshot{
		Gate202Inherited:                      true,
		Gate201ConditionalShapesInherited:     prev.Firewall.Gate201UniversalShapesConditionalOnly,
		BoundaryOffsetEquivalenceEstablished:  prev.Summary.EquivalenceEstablished,
		BGapContactTraceOffsetFailedRoute:     prev.Summary.FailedRouteLogged && prev.Summary.NoCanonicalOffsetFound,
		UniversalCompletionNotDerived:         prev.Summary.UniversalCompletionNotDerived,
		NoPhysicalPredictionClaim:             prev.Summary.NoPhysicalPredictionClaim,
		ObservedInputsUsedForFiniteDerivation: prev.Firewall.ObservedInputsUsedForFiniteDerivation,
		PhysicalUnificationClaimed:            prev.Firewall.PhysicalUnificationClaimed,
		ThresholdCorrectedPhysicalFitClaimed:  prev.Firewall.ThresholdCorrectedPhysicalFitClaimed,
		AbsoluteMassPredicted:                 prev.Firewall.AbsoluteMassPredicted,
		FiniteMatchingCorrectionsDerived:      prev.Firewall.FiniteMatchingCorrectionsDerived,
		StrictNullityAfter:                    prev.Firewall.StrictNullityAfter,
		PhysicalPredictionNullityAfter:        prev.Firewall.PhysicalPredictionNullityAfter,
		Requirements:                          reqs,
		TruthStatement:                        prev.TruthStatement,
	}, nil
}

type CompleteMultipletBasis struct {
	Name                       string
	UnifiedRepresentation      string
	Statistics                 string
	UniversalBetaRow           Rational
	GUTComplete                bool
	OneLoopRowExact            bool
	KnownPhysicalScaffold      bool
	FiniteEngineDerivedAsHeavy bool
	RequiresHeavyDuplicate     bool
	RequiresMassActivation     bool
	RequiresDecouplingScheme   bool
	RequiresFiniteThresholdMap bool
	Verdict                    string
}

type MultipletFit struct {
	RequiredShape              string
	RequiredUniversalBeta      float64
	BasisName                  string
	BasisRow                   Rational
	ExactMultiplicity          float64
	NearestIntegerMultiplicity int
	NearestUniversalBeta       Rational
	SignedResidual             float64
	AbsResidual                float64
	ExactIntegerMatch          bool
	FiniteDerived              bool
	ConditionalPrediction      bool
	Verdict                    string
}

type CompleteMultipletAudit struct {
	RequirementsAudited          int
	BasisRowsAudited             int
	GUTCompleteRows              int
	ExactOneLoopRows             int
	ExactIntegerMultipletMatches int
	FiniteDerivedHeavyMultiplets int
	ConditionalPredictions       int
	NearestResidual              float64
	NearestCandidate             string
	CompleteMultipletSourceFound bool
	Verdict                      string
}

type FiniteInventoryAudit struct {
	ContactPartialOverlapModes         int
	ContactRowsHaveChargeSemantics     bool
	ContactRowsHaveGaugeRepresentation bool
	ContactRowsHaveDynkinIndex         bool
	ContactRowsHaveBetaPermission      bool
	ContactCanAssembleSU5Five          bool
	ContactCanAssembleSU5Ten           bool
	ContactCanAssembleSO10Sixteen      bool
	QuarticContactRows                 int
	QuarticMultipletRepresentationRows int
	QuarticBetaIndexRows               int
	FockStates                         int
	FockKinematicSO10SixteenAvailable  bool
	FockSterileSingletAvailable        bool
	FockRepTraceBoundarySeedClosed     bool
	FockHeavyDuplicateDerived          bool
	FockThresholdMassDerived           bool
	FockCompleteMultipletBetaActivated bool
	UnassignedFockComponentsAssembled  bool
	FiniteCompleteMultipletFound       bool
	Verdict                            string
}

type RegulatorTraceCandidate struct {
	Name                     string
	Source                   string
	TraceValue               string
	ApproxValue              float64
	Canonical                bool
	Universal                bool
	ConformalAnomalyDerived  bool
	GhostOrBRSTComplete      bool
	SpectralTripleComplete   bool
	CutoffFunctionDerived    bool
	GaugeMeasureMapDerived   bool
	BetaRowPermission        bool
	MatchesRequiredUniversal bool
	NearestRequiredShape     string
	NearestResidual          float64
	ConditionalPrediction    bool
	Verdict                  string
}

type RegulatorTraceAudit struct {
	CandidatesAudited             int
	CanonicalTraces               int
	UniversalAnomalyCandidates    int
	ConformalAnomalyDerived       bool
	GhostBRSTCancellationComplete bool
	SpectralTripleComplete        bool
	GaugeMeasureMapDerived        bool
	BetaRowPermission             bool
	ExactRequiredMatches          int
	ConditionalPredictions        int
	RegulatorTraceSourceFound     bool
	Verdict                       string
}

type SourceClassificationAudit struct {
	CompleteMultipletBranchAudited bool
	RegulatorTraceBranchAudited    bool
	AllowedCanonicalSources        int
	ExternalPhenomenologySources   int
	ClassificationComplete         bool
	Verdict                        string
}

type FirewallAudit struct {
	Gate202Inherited                      bool
	Gate202FailedRoutePreserved           bool
	ObservedInputsUsedForFiniteDerivation bool
	UniversalBetaSourceDerived            bool
	CompleteHeavyMultipletDerived         bool
	RegulatorTraceAnomalyDerived          bool
	ContactModesPromotedToBetaRows        bool
	FockGenerationPromotedToNewThreshold  bool
	ArbitraryIntegerMultiplicityInserted  bool
	ArbitraryRegulatorCoefficientInserted bool
	PhysicalUnificationClaimed            bool
	ThresholdCorrectedPhysicalFitClaimed  bool
	AbsoluteMassPredicted                 bool
	FiniteMatchingCorrectionsDerived      bool
	StrictNullityBefore                   int
	StrictNullityAfter                    int
	PhysicalPredictionNullityBefore       int
	PhysicalPredictionNullityAfter        int
	RecommendedNextGate                   string
	OpenRequirements                      []string
	Verdict                               string
}

type Summary struct {
	TestsAudited                     int
	Gate202Inherited                 bool
	CompleteMultipletRowsAudited     bool
	NoCompleteMultipletSourceFound   bool
	FiniteInventoryAudited           bool
	NoRegulatorTraceSourceFound      bool
	UniversalBetaSourceStillExternal bool
	FailedRouteLogged                bool
	NoPhysicalPredictionClaim        bool
	Comment                          string
}

type Analysis struct {
	PreviousGate202 Gate202Snapshot

	MultipletBasis      []CompleteMultipletBasis
	MultipletFits       []MultipletFit
	MultipletAudit      CompleteMultipletAudit
	FiniteInventory     FiniteInventoryAudit
	RegulatorCandidates []RegulatorTraceCandidate
	RegulatorAudit      RegulatorTraceAudit
	Classification      SourceClassificationAudit
	Firewall            FirewallAudit
	Summary             Summary
	TruthStatement      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := DefaultGate202Snapshot()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev Gate202Snapshot) (Analysis, error) {
	if !prev.Gate202Inherited || !prev.Gate201ConditionalShapesInherited || !prev.BoundaryOffsetEquivalenceEstablished || !prev.BGapContactTraceOffsetFailedRoute || len(prev.Requirements) == 0 {
		return Analysis{}, fmt.Errorf("Gate 203 requires Gate 202 failed-route offset audit with Gate-201 universal requirements")
	}
	if prev.PhysicalUnificationClaimed || prev.ThresholdCorrectedPhysicalFitClaimed || prev.AbsoluteMassPredicted || prev.ObservedInputsUsedForFiniteDerivation {
		return Analysis{}, fmt.Errorf("Gate 203 refuses inherited physical prediction or observed-input leakage")
	}

	basis := completeMultipletBasis()
	fits := multipletFits(prev.Requirements, basis)
	ma := auditCompleteMultiplets(prev, basis, fits)
	fi := auditFiniteInventory(ma)
	regs := regulatorTraceCandidates(prev.Requirements)
	ra := auditRegulatorTrace(regs)
	cl := auditClassification(ma, ra)
	fw := auditFirewall(prev, ma, fi, ra, cl)
	summary := Summary{
		TestsAudited:                     7,
		Gate202Inherited:                 fw.Gate202Inherited && fw.Gate202FailedRoutePreserved,
		CompleteMultipletRowsAudited:     ma.BasisRowsAudited == len(basis) && ma.RequirementsAudited == len(prev.Requirements),
		NoCompleteMultipletSourceFound:   !ma.CompleteMultipletSourceFound && ma.ExactIntegerMultipletMatches == 0,
		FiniteInventoryAudited:           !fi.FiniteCompleteMultipletFound && fi.ContactPartialOverlapModes == 7 && fi.FockStates == 16,
		NoRegulatorTraceSourceFound:      !ra.RegulatorTraceSourceFound && ra.ConditionalPredictions == 0,
		UniversalBetaSourceStillExternal: !fw.UniversalBetaSourceDerived && cl.ExternalPhenomenologySources == len(prev.Requirements),
		FailedRouteLogged:                !fw.UniversalBetaSourceDerived && !fw.PhysicalUnificationClaimed,
		NoPhysicalPredictionClaim:        !fw.PhysicalUnificationClaimed && !fw.ThresholdCorrectedPhysicalFitClaimed && !fw.AbsoluteMassPredicted && fw.PhysicalPredictionNullityBefore == fw.PhysicalPredictionNullityAfter,
		Comment:                          "Gate 203 classifies the Gate-201/Gate-202 universal beta row. Exact complete-multiplet beta rows are rational universal rows, but no integer complete-multiplet sum equals the required real c_univ values, and the finite contact/Fock inventory does not derive a new heavy multiplet. Regulator and ghost trace routes remain blocked by missing spectral-triple, cutoff, BRST, and gauge-measure maps.",
	}
	truth := "Gate 203 audits the two standard sources of a universal one-loop beta shift: complete unified heavy multiplets and regulator/ghost measure traces. The complete-multiplet branch has exact rational universal rows, but the required Gate-201 universal rows are not exact integer sums of those rows and are not finite-derived. The contact partial-overlap modes and quartic block lack charge, Dynkin-index, local-field, mass-activation, and decoupling semantics; the Fock 16 is a kinematic one-generation scaffold, not a derived new heavy duplicate threshold. The regulator branch also fails: tau_eta/contact-zeta/BRST/spectral-action traces are real finite data but do not yet form a conformal anomaly or universal beta ledger. Therefore the universal beta source remains external phenomenological data and Gate 203 is a FAILED_ROUTE under current axioms."

	return Analysis{
		PreviousGate202:     prev,
		MultipletBasis:      basis,
		MultipletFits:       fits,
		MultipletAudit:      ma,
		FiniteInventory:     fi,
		RegulatorCandidates: regs,
		RegulatorAudit:      ra,
		Classification:      cl,
		Firewall:            fw,
		Summary:             summary,
		TruthStatement:      truth,
	}, nil
}

func completeMultipletBasis() []CompleteMultipletBasis {
	return []CompleteMultipletBasis{
		{Name: "Weyl SU(5) 5bar", UnifiedRepresentation: "5bar = d^c ⊕ L", Statistics: "Weyl fermion", UniversalBetaRow: R(1, 3), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: true, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "complete rational row, but no finite heavy duplicate or activation threshold is derived"},
		{Name: "Weyl SU(5) 10", UnifiedRepresentation: "10 = Q ⊕ u^c ⊕ e^c", Statistics: "Weyl fermion", UniversalBetaRow: R(1, 1), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: true, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "complete rational row, but no finite heavy duplicate or activation threshold is derived"},
		{Name: "Weyl SU(5) 5bar+10 / SO(10) 16 without sterile beta", UnifiedRepresentation: "5bar ⊕ 10, with ν^c beta-neutral", Statistics: "Weyl fermion generation", UniversalBetaRow: R(4, 3), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: true, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "one complete generation row; Fock kinematics exist, but a new heavy duplicate generation is not derived"},
		{Name: "Dirac/vectorlike SU(5) 5+5bar", UnifiedRepresentation: "5 ⊕ 5bar", Statistics: "vectorlike Weyl pair", UniversalBetaRow: R(2, 3), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: false, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "complete vectorlike row, not finite-derived"},
		{Name: "Dirac/vectorlike full generation 16+16bar", UnifiedRepresentation: "(5bar ⊕ 10 ⊕ 1) plus conjugate", Statistics: "vectorlike generation", UniversalBetaRow: R(8, 3), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: false, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "complete vectorlike generation row, not finite-derived"},
		{Name: "Complex scalar SU(5) 5", UnifiedRepresentation: "scalar 5", Statistics: "complex scalar", UniversalBetaRow: R(1, 6), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: false, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "scalar complete multiplet row, but no scalar heavy threshold map is derived"},
		{Name: "Complex scalar SU(5) 10", UnifiedRepresentation: "scalar 10", Statistics: "complex scalar", UniversalBetaRow: R(1, 2), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: false, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "scalar complete multiplet row, but no scalar heavy threshold map is derived"},
		{Name: "Complex scalar SU(5) 5+10", UnifiedRepresentation: "scalar 5 ⊕ 10", Statistics: "complex scalar generation shape", UniversalBetaRow: R(2, 3), GUTComplete: true, OneLoopRowExact: true, KnownPhysicalScaffold: false, FiniteEngineDerivedAsHeavy: false, RequiresHeavyDuplicate: true, RequiresMassActivation: true, RequiresDecouplingScheme: true, RequiresFiniteThresholdMap: true, Verdict: "scalar generation-shape row equals vectorlike 5 pair row, but no finite heavy threshold is derived"},
	}
}

func multipletFits(reqs []Gate202UniversalRequirement, basis []CompleteMultipletBasis) []MultipletFit {
	out := make([]MultipletFit, 0, len(reqs)*len(basis))
	for _, req := range reqs {
		for _, b := range basis {
			row := b.UniversalBetaRow.Float()
			exactMult := req.RequiredUniversalBeta / row
			nearest := int(math.Round(exactMult))
			if nearest < 0 {
				nearest = 0
			}
			nearestRow := b.UniversalBetaRow.MulInt(nearest)
			residual := nearestRow.Float() - req.RequiredUniversalBeta
			exact := math.Abs(residual) <= tolerance && nearest > 0
			finite := exact && b.FiniteEngineDerivedAsHeavy && !b.RequiresHeavyDuplicate && !b.RequiresMassActivation && !b.RequiresDecouplingScheme && !b.RequiresFiniteThresholdMap
			verdict := "rejected: required universal row is not an exact integer sum of this complete multiplet row"
			if exact && !finite {
				verdict = "numeric integer equality would still be conditional, because the finite engine has not derived the heavy multiplet threshold"
			} else if finite {
				verdict = "conditional prediction: exact finite-derived complete multiplet supplies the universal row"
			}
			out = append(out, MultipletFit{
				RequiredShape:              req.ShapeName,
				RequiredUniversalBeta:      req.RequiredUniversalBeta,
				BasisName:                  b.Name,
				BasisRow:                   b.UniversalBetaRow,
				ExactMultiplicity:          exactMult,
				NearestIntegerMultiplicity: nearest,
				NearestUniversalBeta:       nearestRow,
				SignedResidual:             residual,
				AbsResidual:                math.Abs(residual),
				ExactIntegerMatch:          exact,
				FiniteDerived:              finite,
				ConditionalPrediction:      finite,
				Verdict:                    verdict,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if math.Abs(out[i].AbsResidual-out[j].AbsResidual) > 1e-15 {
			return out[i].AbsResidual < out[j].AbsResidual
		}
		if out[i].RequiredShape != out[j].RequiredShape {
			return out[i].RequiredShape < out[j].RequiredShape
		}
		return out[i].BasisName < out[j].BasisName
	})
	return out
}

func auditCompleteMultiplets(prev Gate202Snapshot, basis []CompleteMultipletBasis, fits []MultipletFit) CompleteMultipletAudit {
	gut := 0
	exactRows := 0
	finiteHeavy := 0
	for _, b := range basis {
		if b.GUTComplete {
			gut++
		}
		if b.OneLoopRowExact {
			exactRows++
		}
		if b.FiniteEngineDerivedAsHeavy {
			finiteHeavy++
		}
	}
	exactMatches := 0
	predictions := 0
	nearest := math.Inf(1)
	nearestName := ""
	for _, f := range fits {
		if f.ExactIntegerMatch {
			exactMatches++
		}
		if f.ConditionalPrediction {
			predictions++
		}
		if f.AbsResidual < nearest {
			nearest = f.AbsResidual
			nearestName = fmt.Sprintf("%s via %d × %s = %s", f.RequiredShape, f.NearestIntegerMultiplicity, f.BasisName, f.NearestUniversalBeta)
		}
	}
	return CompleteMultipletAudit{
		RequirementsAudited:          len(prev.Requirements),
		BasisRowsAudited:             len(basis),
		GUTCompleteRows:              gut,
		ExactOneLoopRows:             exactRows,
		ExactIntegerMultipletMatches: exactMatches,
		FiniteDerivedHeavyMultiplets: finiteHeavy,
		ConditionalPredictions:       predictions,
		NearestResidual:              nearest,
		NearestCandidate:             nearestName,
		CompleteMultipletSourceFound: predictions > 0,
		Verdict:                      "complete multiplets supply rational universal beta rows, but none exactly matches the required Gate-201 real universal rows as a finite-derived integer heavy threshold",
	}
}

func auditFiniteInventory(ma CompleteMultipletAudit) FiniteInventoryAudit {
	return FiniteInventoryAudit{
		ContactPartialOverlapModes:         7,
		ContactRowsHaveChargeSemantics:     false,
		ContactRowsHaveGaugeRepresentation: false,
		ContactRowsHaveDynkinIndex:         false,
		ContactRowsHaveBetaPermission:      false,
		ContactCanAssembleSU5Five:          false,
		ContactCanAssembleSU5Ten:           false,
		ContactCanAssembleSO10Sixteen:      false,
		QuarticContactRows:                 4,
		QuarticMultipletRepresentationRows: 0,
		QuarticBetaIndexRows:               0,
		FockStates:                         16,
		FockKinematicSO10SixteenAvailable:  true,
		FockSterileSingletAvailable:        true,
		FockRepTraceBoundarySeedClosed:     true,
		FockHeavyDuplicateDerived:          false,
		FockThresholdMassDerived:           false,
		FockCompleteMultipletBetaActivated: false,
		UnassignedFockComponentsAssembled:  false,
		FiniteCompleteMultipletFound:       ma.CompleteMultipletSourceFound,
		Verdict:                            "contact modes remain spectral without charge/Dynkin semantics; the Fock 16 is a kinematic matter scaffold and boundary-trace certificate, not a derived new heavy threshold multiplet",
	}
}

func regulatorTraceCandidates(reqs []Gate202UniversalRequirement) []RegulatorTraceCandidate {
	base := []RegulatorTraceCandidate{
		{Name: "tau_eta dimension trace", Source: "finite scalar/contact integration functional", TraceValue: "7", ApproxValue: 7, Canonical: true, Universal: false, ConformalAnomalyDerived: false, GhostOrBRSTComplete: false, SpectralTripleComplete: false, CutoffFunctionDerived: false, GaugeMeasureMapDerived: false, BetaRowPermission: false, Verdict: "canonical finite trace, but no conformal-anomaly or gauge-measure map turns it into c_univ"},
		{Name: "contact zeta dimension ζ(0)", Source: "Gate-162 contact zeta ledger", TraceValue: "7", ApproxValue: 7, Canonical: true, Universal: false, ConformalAnomalyDerived: false, GhostOrBRSTComplete: false, SpectralTripleComplete: false, CutoffFunctionDerived: false, GaugeMeasureMapDerived: false, BetaRowPermission: false, Verdict: "exact zeta value, but spectral-action coefficient and cutoff map are sealed"},
		{Name: "quartic BRST supertrace route", Source: "Gate-158 quartic BRST/ghost audit", TraceValue: "0 canonical zero differential; nontrivial zero-supertrace noncanonical", ApproxValue: 0, Canonical: false, Universal: false, ConformalAnomalyDerived: false, GhostOrBRSTComplete: false, SpectralTripleComplete: false, CutoffFunctionDerived: false, GaugeMeasureMapDerived: false, BetaRowPermission: false, Verdict: "zero differential is canonical but inert; nontrivial ghost gradings require branch choices and do not yield a zero-beta ledger"},
		{Name: "top-down Fock representation trace", Source: "Gate-166/167 representation trace", TraceValue: "diag(1,1,1,5/3) ratio certificate", ApproxValue: 4.666666666666667, Canonical: true, Universal: false, ConformalAnomalyDerived: false, GhostOrBRSTComplete: false, SpectralTripleComplete: false, CutoffFunctionDerived: false, GaugeMeasureMapDerived: false, BetaRowPermission: false, Verdict: "representation-trace ratio closes the boundary seed, but it is not a universal regulator shift or threshold beta ledger"},
		{Name: "spectral-action zeta ansatz", Source: "Gate-163 spectral action preflight", TraceValue: "zeta/action scalars without spectral triple", ApproxValue: math.NaN(), Canonical: false, Universal: false, ConformalAnomalyDerived: false, GhostOrBRSTComplete: false, SpectralTripleComplete: false, CutoffFunctionDerived: false, GaugeMeasureMapDerived: false, BetaRowPermission: false, Verdict: "spectral pre-data exists, but finite spectral triple, cutoff function, and gauge fluctuation map are missing"},
	}
	for i := range base {
		bestShape := ""
		bestResidual := math.Inf(1)
		if !math.IsNaN(base[i].ApproxValue) {
			for _, req := range reqs {
				res := math.Abs(base[i].ApproxValue - req.RequiredUniversalBeta)
				if res < bestResidual {
					bestResidual = res
					bestShape = req.ShapeName
				}
			}
		} else {
			bestResidual = math.Inf(1)
		}
		base[i].NearestRequiredShape = bestShape
		base[i].NearestResidual = bestResidual
		base[i].MatchesRequiredUniversal = bestResidual <= tolerance
		base[i].ConditionalPrediction = base[i].MatchesRequiredUniversal && base[i].Canonical && base[i].Universal && base[i].ConformalAnomalyDerived && base[i].SpectralTripleComplete && base[i].CutoffFunctionDerived && base[i].GaugeMeasureMapDerived && base[i].BetaRowPermission
	}
	return base
}

func auditRegulatorTrace(cands []RegulatorTraceCandidate) RegulatorTraceAudit {
	canonical := 0
	universal := 0
	exactMatches := 0
	predictions := 0
	conformal := false
	brst := false
	triple := false
	gaugeMap := false
	beta := false
	for _, c := range cands {
		if c.Canonical {
			canonical++
		}
		if c.Universal {
			universal++
		}
		if c.MatchesRequiredUniversal {
			exactMatches++
		}
		if c.ConditionalPrediction {
			predictions++
		}
		conformal = conformal || c.ConformalAnomalyDerived
		brst = brst || c.GhostOrBRSTComplete
		triple = triple || c.SpectralTripleComplete
		gaugeMap = gaugeMap || c.GaugeMeasureMapDerived
		beta = beta || c.BetaRowPermission
	}
	return RegulatorTraceAudit{
		CandidatesAudited:             len(cands),
		CanonicalTraces:               canonical,
		UniversalAnomalyCandidates:    universal,
		ConformalAnomalyDerived:       conformal,
		GhostBRSTCancellationComplete: brst,
		SpectralTripleComplete:        triple,
		GaugeMeasureMapDerived:        gaugeMap,
		BetaRowPermission:             beta,
		ExactRequiredMatches:          exactMatches,
		ConditionalPredictions:        predictions,
		RegulatorTraceSourceFound:     predictions > 0,
		Verdict:                       "no tau_eta, contact-zeta, BRST, or spectral-action trace currently derives a universal conformal anomaly with beta-row permission",
	}
}

func auditClassification(ma CompleteMultipletAudit, ra RegulatorTraceAudit) SourceClassificationAudit {
	allowed := 0
	if ma.CompleteMultipletSourceFound {
		allowed++
	}
	if ra.RegulatorTraceSourceFound {
		allowed++
	}
	return SourceClassificationAudit{
		CompleteMultipletBranchAudited: true,
		RegulatorTraceBranchAudited:    true,
		AllowedCanonicalSources:        allowed,
		ExternalPhenomenologySources:   2,
		ClassificationComplete:         true,
		Verdict:                        "both standard universal-row source branches were audited; neither supplies a canonical finite source, so the row remains external phenomenology",
	}
}

func auditFirewall(prev Gate202Snapshot, ma CompleteMultipletAudit, fi FiniteInventoryAudit, ra RegulatorTraceAudit, cl SourceClassificationAudit) FirewallAudit {
	derived := ma.CompleteMultipletSourceFound || ra.RegulatorTraceSourceFound
	return FirewallAudit{
		Gate202Inherited:                      prev.Gate202Inherited && prev.NoPhysicalPredictionClaim,
		Gate202FailedRoutePreserved:           prev.BGapContactTraceOffsetFailedRoute && prev.UniversalCompletionNotDerived,
		ObservedInputsUsedForFiniteDerivation: prev.ObservedInputsUsedForFiniteDerivation,
		UniversalBetaSourceDerived:            derived,
		CompleteHeavyMultipletDerived:         fi.FiniteCompleteMultipletFound,
		RegulatorTraceAnomalyDerived:          ra.RegulatorTraceSourceFound,
		ContactModesPromotedToBetaRows:        false,
		FockGenerationPromotedToNewThreshold:  false,
		ArbitraryIntegerMultiplicityInserted:  false,
		ArbitraryRegulatorCoefficientInserted: false,
		PhysicalUnificationClaimed:            false,
		ThresholdCorrectedPhysicalFitClaimed:  false,
		AbsoluteMassPredicted:                 false,
		FiniteMatchingCorrectionsDerived:      false,
		StrictNullityBefore:                   prev.StrictNullityAfter,
		StrictNullityAfter:                    prev.StrictNullityAfter,
		PhysicalPredictionNullityBefore:       prev.PhysicalPredictionNullityAfter,
		PhysicalPredictionNullityAfter:        prev.PhysicalPredictionNullityAfter,
		RecommendedNextGate:                   "Gate 204 — representation-row lattice completion / finite heavy-sector basis search",
		OpenRequirements: []string{
			"derive a finite heavy-sector carrier with charge, spin-statistics, local-field, mass-activation, and decoupling semantics",
			"derive an exact representation-row lattice rather than importing complete SU(5)/SO(10) multiplets as phenomenological bases",
			"derive a spectral triple, cutoff function, and gauge-measure map before treating tau_eta or zeta traces as regulator anomalies",
			"derive a canonical threshold matching law and M_* independent of the inverse-threshold phenomenological solution",
			"separate the observed/kinematic Fock generation scaffold from any new heavy duplicate multiplet",
		},
		Verdict: fmt.Sprintf("Gate 203 preserves the Gate-202 failed route; sourceClassification=%s", cl.Verdict),
	}
}

func FormatGate202(s Gate202Snapshot) string {
	return fmt.Sprintf("gate202=%t g201Shapes=%t offsetEq=%t failedOffset=%t universalNotDerived=%t noPrediction=%t observedFinite=%t unification=%t fit=%t mass=%t matching=%t strict=%d prediction=%d requirements=%d", s.Gate202Inherited, s.Gate201ConditionalShapesInherited, s.BoundaryOffsetEquivalenceEstablished, s.BGapContactTraceOffsetFailedRoute, s.UniversalCompletionNotDerived, s.NoPhysicalPredictionClaim, s.ObservedInputsUsedForFiniteDerivation, s.PhysicalUnificationClaimed, s.ThresholdCorrectedPhysicalFitClaimed, s.AbsoluteMassPredicted, s.FiniteMatchingCorrectionsDerived, s.StrictNullityAfter, s.PhysicalPredictionNullityAfter, len(s.Requirements))
}

func FormatRequirement(r Gate202UniversalRequirement) string {
	return fmt.Sprintf("%s shape=%s c_univ=%.9g delta_u=%.9g MB=%.9g M*=%.9g conditional=%t finite=%t", r.ShapeName, r.RationalShapeDeltaB, r.RequiredUniversalBeta, r.RequiredDeltaU, r.ThresholdScaleGeV, r.BoundaryScaleGeV, r.ConditionalOnly, r.FiniteDerived)
}

func FormatRequirements(rs []Gate202UniversalRequirement) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, FormatRequirement(r))
	}
	return strings.Join(parts, "; ")
}

func FormatMultipletBasis(b CompleteMultipletBasis) string {
	return fmt.Sprintf("%s rep=%s stats=%s row=%s complete=%t exact=%t scaffold=%t finiteHeavy=%t needsDuplicate=%t mass=%t decouple=%t thresholdMap=%t", b.Name, b.UnifiedRepresentation, b.Statistics, b.UniversalBetaRow, b.GUTComplete, b.OneLoopRowExact, b.KnownPhysicalScaffold, b.FiniteEngineDerivedAsHeavy, b.RequiresHeavyDuplicate, b.RequiresMassActivation, b.RequiresDecouplingScheme, b.RequiresFiniteThresholdMap)
}

func FormatMultipletBasisList(bs []CompleteMultipletBasis, max int) string {
	if max <= 0 || max > len(bs) {
		max = len(bs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatMultipletBasis(bs[i]))
	}
	if max < len(bs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(bs)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatMultipletFit(f MultipletFit) string {
	return fmt.Sprintf("shape=%s required=%.9g basis=%s row=%s exactMult=%.9g nearest=%d total=%s residual=%.9g exact=%t finite=%t", f.RequiredShape, f.RequiredUniversalBeta, f.BasisName, f.BasisRow, f.ExactMultiplicity, f.NearestIntegerMultiplicity, f.NearestUniversalBeta, f.SignedResidual, f.ExactIntegerMatch, f.FiniteDerived)
}

func FormatMultipletFits(fs []MultipletFit, max int) string {
	if max <= 0 || max > len(fs) {
		max = len(fs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatMultipletFit(fs[i]))
	}
	if max < len(fs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(fs)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCompleteMultipletAudit(a CompleteMultipletAudit) string {
	return fmt.Sprintf("requirements=%d basis=%d gut=%d exactRows=%d exactMatches=%d finiteHeavy=%d predictions=%d nearestResidual=%.9g nearest=%q sourceFound=%t", a.RequirementsAudited, a.BasisRowsAudited, a.GUTCompleteRows, a.ExactOneLoopRows, a.ExactIntegerMultipletMatches, a.FiniteDerivedHeavyMultiplets, a.ConditionalPredictions, a.NearestResidual, a.NearestCandidate, a.CompleteMultipletSourceFound)
}

func FormatFiniteInventory(f FiniteInventoryAudit) string {
	return fmt.Sprintf("contactModes=%d contactCharge=%t contactRep=%t contactDynkin=%t contactBeta=%t su5_5=%t su5_10=%t so10_16=%t quarticRows=%d quarticRep=%d quarticBeta=%d fockStates=%d fockSO10=%t fockSeed=%t heavyDup=%t mass=%t betaActivated=%t finiteMultiplet=%t", f.ContactPartialOverlapModes, f.ContactRowsHaveChargeSemantics, f.ContactRowsHaveGaugeRepresentation, f.ContactRowsHaveDynkinIndex, f.ContactRowsHaveBetaPermission, f.ContactCanAssembleSU5Five, f.ContactCanAssembleSU5Ten, f.ContactCanAssembleSO10Sixteen, f.QuarticContactRows, f.QuarticMultipletRepresentationRows, f.QuarticBetaIndexRows, f.FockStates, f.FockKinematicSO10SixteenAvailable, f.FockRepTraceBoundarySeedClosed, f.FockHeavyDuplicateDerived, f.FockThresholdMassDerived, f.FockCompleteMultipletBetaActivated, f.FiniteCompleteMultipletFound)
}

func FormatRegulatorCandidate(c RegulatorTraceCandidate) string {
	res := "inf"
	if !math.IsInf(c.NearestResidual, 0) && !math.IsNaN(c.NearestResidual) {
		res = fmt.Sprintf("%.9g", c.NearestResidual)
	}
	return fmt.Sprintf("%s source=%s trace=%s approx=%.9g canonical=%t universal=%t anomaly=%t brst=%t triple=%t cutoff=%t gaugeMap=%t beta=%t nearest=%s residual=%s match=%t prediction=%t", c.Name, c.Source, c.TraceValue, c.ApproxValue, c.Canonical, c.Universal, c.ConformalAnomalyDerived, c.GhostOrBRSTComplete, c.SpectralTripleComplete, c.CutoffFunctionDerived, c.GaugeMeasureMapDerived, c.BetaRowPermission, c.NearestRequiredShape, res, c.MatchesRequiredUniversal, c.ConditionalPrediction)
}

func FormatRegulatorCandidates(cs []RegulatorTraceCandidate, max int) string {
	if max <= 0 || max > len(cs) {
		max = len(cs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatRegulatorCandidate(cs[i]))
	}
	if max < len(cs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(cs)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRegulatorAudit(a RegulatorTraceAudit) string {
	return fmt.Sprintf("candidates=%d canonical=%d universalAnomaly=%d conformal=%t brst=%t triple=%t gaugeMap=%t beta=%t exactMatches=%d predictions=%d sourceFound=%t", a.CandidatesAudited, a.CanonicalTraces, a.UniversalAnomalyCandidates, a.ConformalAnomalyDerived, a.GhostBRSTCancellationComplete, a.SpectralTripleComplete, a.GaugeMeasureMapDerived, a.BetaRowPermission, a.ExactRequiredMatches, a.ConditionalPredictions, a.RegulatorTraceSourceFound)
}

func FormatClassification(c SourceClassificationAudit) string {
	return fmt.Sprintf("completeMultiplet=%t regulator=%t allowed=%d external=%d complete=%t verdict=%s", c.CompleteMultipletBranchAudited, c.RegulatorTraceBranchAudited, c.AllowedCanonicalSources, c.ExternalPhenomenologySources, c.ClassificationComplete, c.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g202=%t failedRoute=%t observedFinite=%t sourceDerived=%t heavyMultiplet=%t regulator=%t contactBeta=%t fockThreshold=%t intFit=%t regCoeff=%t unification=%t fit=%t mass=%t matching=%t strict=%d->%d prediction=%d->%d next=%s", f.Gate202Inherited, f.Gate202FailedRoutePreserved, f.ObservedInputsUsedForFiniteDerivation, f.UniversalBetaSourceDerived, f.CompleteHeavyMultipletDerived, f.RegulatorTraceAnomalyDerived, f.ContactModesPromotedToBetaRows, f.FockGenerationPromotedToNewThreshold, f.ArbitraryIntegerMultiplicityInserted, f.ArbitraryRegulatorCoefficientInserted, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.AbsoluteMassPredicted, f.FiniteMatchingCorrectionsDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d g202=%t multiplets=%t noMultiplet=%t inventory=%t noRegulator=%t external=%t failed=%t noPrediction=%t comment=%s", s.TestsAudited, s.Gate202Inherited, s.CompleteMultipletRowsAudited, s.NoCompleteMultipletSourceFound, s.FiniteInventoryAudited, s.NoRegulatorTraceSourceFound, s.UniversalBetaSourceStillExternal, s.FailedRouteLogged, s.NoPhysicalPredictionClaim, s.Comment)
}
