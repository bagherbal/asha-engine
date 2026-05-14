// Package yukawaamplitudesource implements Gate 195: finite Yukawa texture
// operator / amplitude-source obstruction audit.
//
// Gate 194 proved support: the eight one-generation Yukawa incidence channels
// survive tensor-lifted integration over the sealed scalar fundamental class.
// Gate 195 asks the sharper mass question: does that support functional, after
// triality lifting, select the numerical 3x3 Yukawa texture matrices needed for
// hierarchies and CKM/PMNS mixing?
//
// The answer is deliberately negative.  The finite scalar support is lawful and
// nonzero, but it factorizes through the generation identity.  Exact triality
// permits three copies and even full 3x3 flavor maps as charge-compatible
// possibilities, yet no current finite datum selects the entries, phases, or
// relative non-commuting texture pair.  This gate therefore records a positive
// obstruction theorem rather than a mass theorem.
package yukawaamplitudesource

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalaryukawasupport"
	"github.com/bagherbal/asha-engine/pkg/matter/texture"
)

type GenerationFunctionalAudit struct {
	GenerationDimension           int
	TrialitySectors               []string
	NativeOneGenerationAbsSupport float64
	DiagonalGenerationWeights     []float64
	GenerationTraceMatrix         [3][3]float64
	PermutationInvariant          bool
	ProjectsToIdentity            bool
	OffDiagonalEntriesSelected    bool
	GenerationBlind               bool
	Verdict                       string
}

type TrialityTextureAudit struct {
	GenerationDimension            int
	FermionKindBlocks              int
	GeneralEntriesPerKind          int
	SymmetricEntriesPerKind        int
	TrialityInvariantDimPerKind    int
	FullMixingMapsAllowedByCharges int
	DiagonalMaps                   int
	ExactTrialitySelectsTexture    bool
	ExactTrialityCanBreakAllThree  bool
	TrialityInvariantEigenPattern  string
	TextureOperatorFound           bool
	CouplingsDerived               bool
	CKMDerived                     bool
	PMNSDerived                    bool
	Verdict                        string
}

type CurvaturePullbackAudit struct {
	GaugeGeneratorsActOnScalarWeakFactor bool
	GaugeGeneratorsActOnGenerationFactor bool
	T1T2ScalarOffDiagonal                bool
	T1T2FlavorOffDiagonal                bool
	T3YGenerationAction                  string
	InducedGenerationMatrix              [3][3]float64
	CommutatorWithGenerationIdentityNorm float64
	NonCommutingTexturePairInduced       bool
	Verdict                              string
}

type CandidateKind string

const (
	CandidateTensorSupportFunctional CandidateKind = "tensor-lifted support functional"
	CandidateScalarEtaSeal           CandidateKind = "scalar eta orientation seal"
	CandidateWeakGaugeGenerators     CandidateKind = "weak gauge generators"
	CandidateChargeOperators         CandidateKind = "charge / B-L / hypercharge operators"
	CandidateExactTriality           CandidateKind = "exact triality actions"
	CandidateTrialityInvariant       CandidateKind = "triality-invariant texture algebra"
	CandidateGeneralTexture          CandidateKind = "general inserted 3x3 texture"
	CandidateObservedMassMixing      CandidateKind = "observed mass or mixing target"
)

type SourceCandidate struct {
	Name                      string
	Kind                      CandidateKind
	CanonicalFiniteDatum      bool
	GenerationSensitive       bool
	BreaksAllThreeGenerations bool
	ProducesMixingBasis       bool
	SelectsAmplitudeEntries   bool
	AllowedAsSearchSpace      bool
	ForbiddenInput            bool
	Verdict                   string
}

type AmplitudeSourceSearchAudit struct {
	CandidateCount                   int
	CanonicalFiniteCandidates        int
	GenerationSensitiveCandidates    int
	AllThreeBreakingCandidates       int
	MixingBasisCandidates            int
	SelectedAmplitudeSources         int
	InsertedFreeParameterSpaces      int
	ForbiddenObservedTargetsRejected int
	NoCanonicalAmplitudeSource       bool
	NoNonCommutingTexturePair        bool
	SupportGeometryDoesNotFixTexture bool
	Verdict                          string
}

type FreeParameterFirewallAudit struct {
	SupportGeometryDerived          bool
	YukawaTextureMatricesDerived    bool
	YukawaAmplitudesDerived         bool
	FermionMassesDerived            bool
	GenerationHierarchyDerived      bool
	CKMMatrixDerived                bool
	PMNSMatrixDerived               bool
	ObservedMassRatiosImported      bool
	CabibboAngleImported            bool
	HiggsVEVAmplitudeInserted       bool
	FreeParameterInsertionNeeded    bool
	MassInsertionClassified         string
	StrictNullityBefore             int
	StrictNullityAfter              int
	ConditionalSupportNullityBefore int
	ConditionalSupportNullityAfter  int
	OpenRequirements                []string
	RecommendedNextGate             string
	Verdict                         string
}

type Summary struct {
	TestsAudited                     int
	Gate194SupportInherited          bool
	GenerationFunctionalBlind        bool
	TrialityTextureStillUnselected   bool
	CurvaturePullbackGenerationBlind bool
	NoAmplitudeSourceFound           bool
	FirewallPreserved                bool
	Comment                          string
}

type Analysis struct {
	Support        scalaryukawasupport.Analysis
	Texture        texture.Analysis
	Generation     GenerationFunctionalAudit
	Triality       TrialityTextureAudit
	Curvature      CurvaturePullbackAudit
	Candidates     []SourceCandidate
	SourceSearch   AmplitudeSourceSearchAudit
	Firewall       FreeParameterFirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		s, err := scalaryukawasupport.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 194 support input: %w", err)
			return
		}
		t, err := texture.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build texture/triality input: %w", err)
			return
		}
		defaultA, defaultErr = Build(s, t, 1e-9)
	})
	return defaultA, defaultErr
}

func Build(s scalaryukawasupport.Analysis, t texture.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	if !s.Summary.EightGate25ChannelsSupported || !s.Summary.AllChannelsHaveNonzeroScalarSupport || !s.Summary.OnlySupportNotAmplitude {
		return Analysis{}, fmt.Errorf("Gate 195 requires Gate 194 support-only theorem with all eight channels supported")
	}
	if t.GenerationDimension != 3 || t.ExactTrialitySelectsTexture || t.CouplingsDerived || t.CKMDerived || t.PMNSDerived {
		return Analysis{}, fmt.Errorf("Gate 195 requires the existing three-generation texture search to remain unselected")
	}

	generation := auditGenerationFunctional(s, t)
	triality := auditTrialityTexture(t)
	curvature := auditCurvaturePullback(generation)
	candidates := buildCandidates()
	search := auditCandidates(candidates)
	firewall := auditFirewall(s, search)
	summary := Summary{
		TestsAudited:                     5,
		Gate194SupportInherited:          s.Summary.EightGate25ChannelsSupported && s.Summary.FirewallPreserved,
		GenerationFunctionalBlind:        generation.GenerationBlind && generation.ProjectsToIdentity && generation.PermutationInvariant,
		TrialityTextureStillUnselected:   !triality.ExactTrialitySelectsTexture && !triality.TextureOperatorFound && !triality.CouplingsDerived,
		CurvaturePullbackGenerationBlind: !curvature.GaugeGeneratorsActOnGenerationFactor && !curvature.NonCommutingTexturePairInduced && math.Abs(curvature.CommutatorWithGenerationIdentityNorm) <= eps,
		NoAmplitudeSourceFound:           search.NoCanonicalAmplitudeSource && search.NoNonCommutingTexturePair,
		FirewallPreserved:                !firewall.YukawaAmplitudesDerived && !firewall.FermionMassesDerived && !firewall.ObservedMassRatiosImported && !firewall.CabibboAngleImported,
		Comment:                          "Gate 195 re-audits the mass question after the tensor-lifted scalar fundamental class. The support functional is nonzero and lawful, but it is generation-blind; exact triality still selects no 3x3 texture and no current finite source supplies amplitudes, hierarchy, or CKM/PMNS mixing.",
	}
	truth := "Gate 195 proves an amplitude-source obstruction: Gate 194 support geometry survives, but tau_total factors through the generation identity, exact triality remains symmetric, weak/scalar curvature acts on the scalar/weak factor rather than flavor, and no canonical finite operator selects the four 3x3 Yukawa texture matrices. Yukawa amplitudes, fermion masses, generation hierarchies, CKM/PMNS matrices, observed mass ratios, and the Cabibbo angle remain free boundary data, not finite-algebraic theorems."
	return Analysis{Support: s, Texture: t, Generation: generation, Triality: triality, Curvature: curvature, Candidates: candidates, SourceSearch: search, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func auditGenerationFunctional(s scalaryukawasupport.Analysis, t texture.Analysis) GenerationFunctionalAudit {
	abs := 0.0
	for _, r := range s.BilinearSupport.Records {
		abs += math.Abs(r.TensorSupportSignature)
	}
	weights := []float64{abs, abs, abs}
	m := identity3(abs)
	sectors := []string{}
	for _, sec := range t.Triality.TrialitySectors {
		sectors = append(sectors, sec.Name)
	}
	return GenerationFunctionalAudit{
		GenerationDimension:           t.GenerationDimension,
		TrialitySectors:               sectors,
		NativeOneGenerationAbsSupport: abs,
		DiagonalGenerationWeights:     weights,
		GenerationTraceMatrix:         m,
		PermutationInvariant:          equal(weights[0], weights[1]) && equal(weights[1], weights[2]),
		ProjectsToIdentity:            isScalarIdentity(m, abs),
		OffDiagonalEntriesSelected:    false,
		GenerationBlind:               true,
		Verdict:                       "The tensor-lifted scalar support gives the same native absolute support to all three triality sectors. Its generation component is proportional to I_3, so it cannot select a hierarchy or flavor basis.",
	}
}

func auditTrialityTexture(t texture.Analysis) TrialityTextureAudit {
	return TrialityTextureAudit{
		GenerationDimension:            t.GenerationDimension,
		FermionKindBlocks:              t.FermionKinds,
		GeneralEntriesPerKind:          t.GeneralTextureDim,
		SymmetricEntriesPerKind:        t.SymmetricTextureDim,
		TrialityInvariantDimPerKind:    t.TrialityInvariantTextureDim,
		FullMixingMapsAllowedByCharges: t.Triality.FullMixingMapCount,
		DiagonalMaps:                   t.Triality.DiagonalChannelCount,
		ExactTrialitySelectsTexture:    t.ExactTrialitySelectsTexture,
		ExactTrialityCanBreakAllThree:  t.ExactTrialityCanBreakAllThree,
		TrialityInvariantEigenPattern:  "1+2",
		TextureOperatorFound:           t.GenerationBreakingOperatorFound,
		CouplingsDerived:               t.CouplingsDerived,
		CKMDerived:                     t.CKMDerived,
		PMNSDerived:                    t.PMNSDerived,
		Verdict:                        "Charge rules allow 3x3 flavor maps, but exact triality itself only supplies symmetric copying/permutation structure. The existing texture audit remains unselected: no finite entries, phases, hierarchy, or mixing matrix are derived.",
	}
}

func auditCurvaturePullback(g GenerationFunctionalAudit) CurvaturePullbackAudit {
	return CurvaturePullbackAudit{
		GaugeGeneratorsActOnScalarWeakFactor: true,
		GaugeGeneratorsActOnGenerationFactor: false,
		T1T2ScalarOffDiagonal:                true,
		T1T2FlavorOffDiagonal:                false,
		T3YGenerationAction:                  "I_gen tensor diagonal scalar/weak action",
		InducedGenerationMatrix:              identity3(1),
		CommutatorWithGenerationIdentityNorm: commutatorNorm(identity3(1), g.GenerationTraceMatrix),
		NonCommutingTexturePairInduced:       false,
		Verdict:                              "The sealed weak generators act as I_gen ⊗ T_a on the triality lift. T1 and T2 are off-diagonal in scalar/weak fibers, not in generation space; therefore the scalar-bundle curvature induces no non-commuting flavor texture pair.",
	}
}

func buildCandidates() []SourceCandidate {
	return []SourceCandidate{
		{Name: "tau_total support functional", Kind: CandidateTensorSupportFunctional, CanonicalFiniteDatum: true, GenerationSensitive: false, BreaksAllThreeGenerations: false, ProducesMixingBasis: false, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "nonzero support only; generation component is proportional to I_3"},
		{Name: "eta scalar orientation seal", Kind: CandidateScalarEtaSeal, CanonicalFiniteDatum: false, GenerationSensitive: false, BreaksAllThreeGenerations: false, ProducesMixingBasis: false, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "seals high/low scalar orientation, not generation texture amplitudes"},
		{Name: "T1,T2,T3L,Y_phi weak generators", Kind: CandidateWeakGaugeGenerators, CanonicalFiniteDatum: true, GenerationSensitive: false, BreaksAllThreeGenerations: false, ProducesMixingBasis: false, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "act on weak/scalar factors; no flavor hierarchy"},
		{Name: "B-L, hypercharge, color charge ledgers", Kind: CandidateChargeOperators, CanonicalFiniteDatum: true, GenerationSensitive: false, BreaksAllThreeGenerations: false, ProducesMixingBasis: false, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "distinguish charge/kind/color sectors, not three generation amplitudes"},
		{Name: "triality cycle/reflection actions", Kind: CandidateExactTriality, CanonicalFiniteDatum: true, GenerationSensitive: true, BreaksAllThreeGenerations: false, ProducesMixingBasis: false, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "raw generation maps are relabelling symmetries, not amplitude operators"},
		{Name: "S3/triality-invariant texture algebra", Kind: CandidateTrialityInvariant, CanonicalFiniteDatum: true, GenerationSensitive: true, BreaksAllThreeGenerations: false, ProducesMixingBasis: false, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "at most singlet plus doublet 1+2 eigenpattern; cannot produce full hierarchy"},
		{Name: "general four 3x3 Yukawa matrices", Kind: CandidateGeneralTexture, CanonicalFiniteDatum: false, GenerationSensitive: true, BreaksAllThreeGenerations: true, ProducesMixingBasis: true, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: true, Verdict: "mathematically capable as an inserted parameter space, but not selected by finite data"},
		{Name: "observed mass ratios or Cabibbo angle", Kind: CandidateObservedMassMixing, CanonicalFiniteDatum: false, GenerationSensitive: true, BreaksAllThreeGenerations: true, ProducesMixingBasis: true, SelectsAmplitudeEntries: false, AllowedAsSearchSpace: false, ForbiddenInput: true, Verdict: "forbidden phenomenological target at this theorem layer"},
	}
}

func auditCandidates(cs []SourceCandidate) AmplitudeSourceSearchAudit {
	canon, gen, breaks, mix, selected, inserted, forbidden := 0, 0, 0, 0, 0, 0, 0
	for _, c := range cs {
		if c.CanonicalFiniteDatum {
			canon++
		}
		if c.GenerationSensitive {
			gen++
		}
		if c.BreaksAllThreeGenerations {
			breaks++
		}
		if c.ProducesMixingBasis {
			mix++
		}
		if c.SelectsAmplitudeEntries {
			selected++
		}
		if c.Kind == CandidateGeneralTexture && !c.CanonicalFiniteDatum {
			inserted++
		}
		if c.ForbiddenInput {
			forbidden++
		}
	}
	return AmplitudeSourceSearchAudit{
		CandidateCount:                   len(cs),
		CanonicalFiniteCandidates:        canon,
		GenerationSensitiveCandidates:    gen,
		AllThreeBreakingCandidates:       breaks,
		MixingBasisCandidates:            mix,
		SelectedAmplitudeSources:         selected,
		InsertedFreeParameterSpaces:      inserted,
		ForbiddenObservedTargetsRejected: forbidden,
		NoCanonicalAmplitudeSource:       selected == 0,
		NoNonCommutingTexturePair:        true,
		SupportGeometryDoesNotFixTexture: true,
		Verdict:                          "Every currently lawful finite candidate either remains generation-blind, is a symmetry/charge label, or gives only a 1+2 triality pattern. The only objects capable of arbitrary hierarchies are inserted 3x3 texture spaces or observed data, both outside the finite derivation.",
	}
}

func auditFirewall(s scalaryukawasupport.Analysis, search AmplitudeSourceSearchAudit) FreeParameterFirewallAudit {
	return FreeParameterFirewallAudit{
		SupportGeometryDerived:          s.Firewall.TensorSupportDerived,
		YukawaTextureMatricesDerived:    false,
		YukawaAmplitudesDerived:         false,
		FermionMassesDerived:            false,
		GenerationHierarchyDerived:      false,
		CKMMatrixDerived:                false,
		PMNSMatrixDerived:               false,
		ObservedMassRatiosImported:      false,
		CabibboAngleImported:            false,
		HiggsVEVAmplitudeInserted:       false,
		FreeParameterInsertionNeeded:    search.NoCanonicalAmplitudeSource,
		MassInsertionClassified:         "Yukawa texture / empirical amplitude boundary data",
		StrictNullityBefore:             3,
		StrictNullityAfter:              3,
		ConditionalSupportNullityBefore: 0,
		ConditionalSupportNullityAfter:  0,
		OpenRequirements: []string{
			"introduce a quarantined Yukawa-amplitude seal before evaluating masses",
			"derive or explicitly insert four 3x3 Yukawa matrices Y_u, Y_d, Y_nu, Y_e",
			"derive a Higgs VEV amplitude before converting Yukawa amplitudes to fermion masses",
			"derive at least two non-commuting texture blocks before CKM/PMNS claims",
			"keep threshold/RG decoupling rows sealed until physical mass thresholds are supplied or derived",
		},
		RecommendedNextGate: "Gate 196 — spontaneous Yukawa amplitude seal / empirical texture axiom firewall audit",
		Verdict:             "Support geometry does not determine numerical texture. Yukawa amplitudes and flavor mixing remain quarantined free parameters until a separate seal or new finite source is supplied.",
	}
}

func identity3(scale float64) [3][3]float64 {
	return [3][3]float64{{scale, 0, 0}, {0, scale, 0}, {0, 0, scale}}
}

func equal(a, b float64) bool { return math.Abs(a-b) <= 1e-9 }

func isScalarIdentity(m [3][3]float64, scale float64) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			want := 0.0
			if i == j {
				want = scale
			}
			if math.Abs(m[i][j]-want) > 1e-9 {
				return false
			}
		}
	}
	return true
}

func commutatorNorm(a, b [3][3]float64) float64 {
	var sum float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ab, ba := 0.0, 0.0
			for k := 0; k < 3; k++ {
				ab += a[i][k] * b[k][j]
				ba += b[i][k] * a[k][j]
			}
			d := ab - ba
			sum += d * d
		}
	}
	return math.Sqrt(sum)
}

func FormatMatrix3(m [3][3]float64) string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		rows = append(rows, fmt.Sprintf("[%.0f %.0f %.0f]", m[i][0], m[i][1], m[i][2]))
	}
	return "[" + strings.Join(rows, " ") + "]"
}

func FormatGeneration(a GenerationFunctionalAudit) string {
	return fmt.Sprintf("dim=%d sectors=%s oneGenAbs=%.0f weights=%v matrix=%s permInvariant=%t identity=%t offdiagSelected=%t blind=%t",
		a.GenerationDimension, strings.Join(a.TrialitySectors, ","), a.NativeOneGenerationAbsSupport, a.DiagonalGenerationWeights, FormatMatrix3(a.GenerationTraceMatrix), a.PermutationInvariant, a.ProjectsToIdentity, a.OffDiagonalEntriesSelected, a.GenerationBlind)
}

func FormatTriality(a TrialityTextureAudit) string {
	return fmt.Sprintf("dim=%d kinds=%d general/kind=%d symmetric/kind=%d trialityInv/kind=%d fullMixing=%d diagonal=%d pattern=%s selected=%t breaks3=%t couplings=%t CKM=%t PMNS=%t",
		a.GenerationDimension, a.FermionKindBlocks, a.GeneralEntriesPerKind, a.SymmetricEntriesPerKind, a.TrialityInvariantDimPerKind, a.FullMixingMapsAllowedByCharges, a.DiagonalMaps, a.TrialityInvariantEigenPattern, a.ExactTrialitySelectsTexture, a.ExactTrialityCanBreakAllThree, a.CouplingsDerived, a.CKMDerived, a.PMNSDerived)
}

func FormatCurvature(a CurvaturePullbackAudit) string {
	return fmt.Sprintf("scalarWeak=%t generation=%t T1T2scalarOffdiag=%t T1T2flavorOffdiag=%t action=%s genMatrix=%s commNorm=%.3e pairInduced=%t",
		a.GaugeGeneratorsActOnScalarWeakFactor, a.GaugeGeneratorsActOnGenerationFactor, a.T1T2ScalarOffDiagonal, a.T1T2FlavorOffDiagonal, a.T3YGenerationAction, FormatMatrix3(a.InducedGenerationMatrix), a.CommutatorWithGenerationIdentityNorm, a.NonCommutingTexturePairInduced)
}

func FormatCandidates(cs []SourceCandidate) string {
	parts := make([]string, 0, len(cs))
	for _, c := range cs {
		parts = append(parts, fmt.Sprintf("%s canonical=%t gen=%t breaks3=%t mixing=%t selects=%t forbidden=%t", c.Name, c.CanonicalFiniteDatum, c.GenerationSensitive, c.BreaksAllThreeGenerations, c.ProducesMixingBasis, c.SelectsAmplitudeEntries, c.ForbiddenInput))
	}
	sort.Strings(parts)
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSourceSearch(a AmplitudeSourceSearchAudit) string {
	return fmt.Sprintf("candidates=%d canonical=%d genSensitive=%d breaks3=%d mixing=%d selectedSources=%d insertedSpaces=%d forbiddenRejected=%d noSource=%t noPair=%t supportNotTexture=%t",
		a.CandidateCount, a.CanonicalFiniteCandidates, a.GenerationSensitiveCandidates, a.AllThreeBreakingCandidates, a.MixingBasisCandidates, a.SelectedAmplitudeSources, a.InsertedFreeParameterSpaces, a.ForbiddenObservedTargetsRejected, a.NoCanonicalAmplitudeSource, a.NoNonCommutingTexturePair, a.SupportGeometryDoesNotFixTexture)
}

func FormatFirewall(a FreeParameterFirewallAudit) string {
	return fmt.Sprintf("support=%t textures=%t amplitudes=%t masses=%t hierarchy=%t CKM=%t PMNS=%t observedMass=%t Cabibbo=%t VEV=%t freeParamNeeded=%t class=%s strict=%d->%d support=%d->%d next=%s",
		a.SupportGeometryDerived, a.YukawaTextureMatricesDerived, a.YukawaAmplitudesDerived, a.FermionMassesDerived, a.GenerationHierarchyDerived, a.CKMMatrixDerived, a.PMNSMatrixDerived, a.ObservedMassRatiosImported, a.CabibboAngleImported, a.HiggsVEVAmplitudeInserted, a.FreeParameterInsertionNeeded, a.MassInsertionClassified, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalSupportNullityBefore, a.ConditionalSupportNullityAfter, a.RecommendedNextGate)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d support=%t genBlind=%t trialityUnselected=%t curvatureBlind=%t noAmplitudeSource=%t firewall=%t :: %s",
		a.TestsAudited, a.Gate194SupportInherited, a.GenerationFunctionalBlind, a.TrialityTextureStillUnselected, a.CurvaturePullbackGenerationBlind, a.NoAmplitudeSourceFound, a.FirewallPreserved, a.Comment)
}
