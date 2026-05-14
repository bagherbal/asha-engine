// Package noncommutingtexturepair implements Gate 173: the finite
// non-commuting Yukawa texture-pair search.
//
// Gates 168-172 converted the mass problem into a precise finite target:
// four 3x3 Yukawa matrices, a scalar-shape moment constraint, and the need
// for at least two non-commuting generation texture operators before CKM/PMNS
// mixing can be claimed. Gate 173 audits every currently derived generation
// operator that could plausibly act on that 3D carrier.
//
// The audit deliberately separates two notions which are easy to confuse:
// raw non-commuting generation maps versus qualified non-commuting Yukawa
// texture sources. Triality permutation generators are canonical raw maps and
// some of them do not commute, but they are symmetry/label actions, not
// Hermitian generation-breaking amplitude operators. The texture pair required
// for masses and mixing must be charge-compatible, self-adjoint or polarizable
// into a Yukawa amplitude block, generation-breaking, canonical as finite data,
// and not merely a relabelling symmetry. No such pair is selected by the gates
// currently available.
package noncommutingtexturepair

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/trialitytexturelift"
)

type OperatorKind string

const (
	KindIdentity             OperatorKind = "generation identity"
	KindTrialityCycle        OperatorKind = "triality cycle generator"
	KindTrialityReflection   OperatorKind = "triality reflection generator"
	KindTrialityInvariant    OperatorKind = "triality-invariant projector/algebra"
	KindDiagonalSpurion      OperatorKind = "Higgs/contact diagonal generation spurion"
	KindBFResidual           OperatorKind = "BF/active-generation curvature residual"
	KindScalarShapeProjector OperatorKind = "scalar-shape contact-kind projector lifted to generation identity"
	KindRealStructure        OperatorKind = "spectral-triple real structure on generation indices"
	KindSourceTensorMinimum  OperatorKind = "source-tensor variational minimum"
)

type GenerationOperator struct {
	Name                   string
	Kind                   OperatorKind
	SourceGate             string
	Matrix                 [3][3]float64
	Canonical              bool
	SelfAdjoint            bool
	LinearTextureCandidate bool
	GenerationBreaking     bool
	ProducesMixingBasis    bool
	ChargeCompatible       bool
	KindSensitive          bool
	RequiresBridge         bool
	PureSymmetryAction     bool
	ZeroOperator           bool
	IdentityLike           bool
	QualifiedTexture       bool
	Disqualification       string
}

type PairAudit struct {
	A                         string
	B                         string
	RawCommutatorNorm         float64
	RawNonCommuting           bool
	BothQualifiedTextures     bool
	QualifiedNonCommutingPair bool
	ReasonNotQualified        string
}

type InventoryAudit struct {
	OperatorCount                 int
	CanonicalOperators            int
	SelfAdjointOperators          int
	GenerationBreakingOperators   int
	LinearTextureCandidates       int
	QualifiedTextureOperators     int
	RawNonCommutingPairs          int
	QualifiedNonCommutingPairs    int
	CanonicalBreakingTextures     int
	CanonicalNonzeroCurvatureMaps int
	CanonicalMixingSources        int
	Verdict                       string
}

type NoGoAudit struct {
	TrialityRawNoncommutationSeen       bool
	TrialityRawMapsAreSymmetries        bool
	TrialityInvariantTexturesTooSmall   bool
	DiagonalSpurionRequiresBridge       bool
	BFResidualZero                      bool
	ScalarShapeProjectorGenerationBlind bool
	RealStructureGenerationBlind        bool
	SourceTensorMinimumZero             bool
	NoQualifiedTexturePair              bool
	MassGenerationSealedAtCurrentStage  bool
	NextRequiredInput                   string
	Verdict                             string
}

type FirewallAudit struct {
	GaugeRatioClosed                   bool
	ScalarShapeTargetAvailable         bool
	MassProblemLocalizedToYukawaMatrix bool
	NonCommutingTexturePairRequired    bool
	NonCommutingTexturePairFound       bool
	YukawaAmplitudesDerived            bool
	FermionMassesDerived               bool
	CKMPMNSDerived                     bool
	PhysicalConstantsDerived           bool
	ResidualNullityBefore              int
	ResidualNullityAfter               int
	RecommendedNextGate                string
	Verdict                            string
}

type Analysis struct {
	Previous       trialitytexturelift.Analysis
	Operators      []GenerationOperator
	Pairs          []PairAudit
	Inventory      InventoryAudit
	NoGo           NoGoAudit
	Firewall       FirewallAudit
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := trialitytexturelift.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev trialitytexturelift.Analysis) (Analysis, error) {
	if !prev.MassAudit.AtLeastTwoNoncommutingNeeded || prev.MassAudit.AtLeastTwoNoncommutingFound || prev.OperatorAudit.NonCommutingPairsFound != 0 {
		return Analysis{}, fmt.Errorf("Gate 173 requires Gate 172 to require, but not find, a non-commuting texture pair")
	}
	if prev.Generation.GenerationCarrierDimension != 3 {
		return Analysis{}, fmt.Errorf("Gate 173 requires a 3D generation carrier")
	}

	ops := buildOperators(prev)
	pairs := auditPairs(ops)
	inventory := auditInventory(ops, pairs)
	noGo := buildNoGo(ops, inventory)
	firewall := FirewallAudit{
		GaugeRatioClosed:                   prev.Firewall.GaugeRatioClosed,
		ScalarShapeTargetAvailable:         prev.Firewall.ScalarShapeTargetAvailable,
		MassProblemLocalizedToYukawaMatrix: prev.MassAudit.FourYukawaMatricesRecognized,
		NonCommutingTexturePairRequired:    prev.MassAudit.AtLeastTwoNoncommutingNeeded,
		NonCommutingTexturePairFound:       inventory.QualifiedNonCommutingPairs > 0,
		YukawaAmplitudesDerived:            false,
		FermionMassesDerived:               false,
		CKMPMNSDerived:                     false,
		PhysicalConstantsDerived:           false,
		ResidualNullityBefore:              3,
		ResidualNullityAfter:               3,
		RecommendedNextGate:                "Gate 174 — spectral-action normalization from the topological action seal",
		Verdict:                            "the finite mass-generation route is sealed at the current stage; the next independent problem is absolute gauge-coupling normalization, not another texture probe",
	}

	return Analysis{
		Previous:       prev,
		Operators:      ops,
		Pairs:          pairs,
		Inventory:      inventory,
		NoGo:           noGo,
		Firewall:       firewall,
		TruthStatement: "Gate 173 finds raw non-commuting triality/generation maps, but no canonical non-commuting Yukawa texture pair. Triality permutations are symmetries, the diagonal Higgs/contact spurion requires a bridge, BF/source residuals are zero, and scalar-shape/contact-kind projectors are generation-blind. The CKM/PMNS problem is therefore structurally open at the current finite-data stage.",
	}, nil
}

func buildOperators(prev trialitytexturelift.Analysis) []GenerationOperator {
	spurion := prev.Generation.BestCandidate.Eigenvalues
	if len(spurion) != 3 {
		spurion = []float64{3, 2, 1}
	}
	// Normalization is irrelevant for commutator tests; only eigenspaces matter.
	D := diag(spurion[0], spurion[1], spurion[2])
	I := identity()
	C := [3][3]float64{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	R := [3][3]float64{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	Jdem := [3][3]float64{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	zero := [3][3]float64{}

	ops := []GenerationOperator{
		{
			Name:                   "I_gen",
			Kind:                   KindIdentity,
			SourceGate:             "Gate 26 triality carrier bookkeeping",
			Matrix:                 I,
			Canonical:              true,
			SelfAdjoint:            true,
			LinearTextureCandidate: false,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			IdentityLike:           true,
			QualifiedTexture:       false,
			Disqualification:       "identity-like; cannot split generations or generate mixing",
		},
		{
			Name:                   "C3_cycle",
			Kind:                   KindTrialityCycle,
			SourceGate:             "Gate 26 exact triality action",
			Matrix:                 C,
			Canonical:              true,
			SelfAdjoint:            false,
			LinearTextureCandidate: false,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			PureSymmetryAction:     true,
			QualifiedTexture:       false,
			Disqualification:       "canonical raw map but not self-adjoint and represents a triality symmetry/label action, not a Yukawa amplitude source",
		},
		{
			Name:                   "S3_reflection",
			Kind:                   KindTrialityReflection,
			SourceGate:             "Gate 26 exact triality action",
			Matrix:                 R,
			Canonical:              true,
			SelfAdjoint:            true,
			LinearTextureCandidate: false,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			PureSymmetryAction:     true,
			QualifiedTexture:       false,
			Disqualification:       "canonical raw map but it is a triality reflection symmetry/label action, not a finite Dirac amplitude texture",
		},
		{
			Name:                   "P_triality_singlet",
			Kind:                   KindTrialityInvariant,
			SourceGate:             "Gate 28 exact triality-invariant texture algebra",
			Matrix:                 scale(Jdem, 1.0/3.0),
			Canonical:              true,
			SelfAdjoint:            true,
			LinearTextureCandidate: true,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			QualifiedTexture:       false,
			Disqualification:       "triality-invariant texture algebra has only singlet-plus-doublet structure and cannot produce three distinct masses or CKM/PMNS",
		},
		{
			Name:                   "D_HiggsContact",
			Kind:                   KindDiagonalSpurion,
			SourceGate:             "Gate 29 Higgs/contact anisotropy diagonal spurion",
			Matrix:                 D,
			Canonical:              false,
			SelfAdjoint:            true,
			LinearTextureCandidate: true,
			GenerationBreaking:     true,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			RequiresBridge:         true,
			QualifiedTexture:       false,
			Disqualification:       "splits three diagonal weights but Gate 29 marked it bridge-required, not canonical as a total Yukawa operator; by itself it produces no mixing",
		},
		{
			Name:                   "F_BF_residual",
			Kind:                   KindBFResidual,
			SourceGate:             "Gates 29-35 active-to-generation curvature projection",
			Matrix:                 zero,
			Canonical:              true,
			SelfAdjoint:            true,
			LinearTextureCandidate: true,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			ZeroOperator:           true,
			QualifiedTexture:       false,
			Disqualification:       "canonical residual is zero; zero curvature cannot supply a non-commuting texture source",
		},
		{
			Name:                   "P_shape_kind_lift",
			Kind:                   KindScalarShapeProjector,
			SourceGate:             "Gates 169-171 scalar-shape/contact-kind projector",
			Matrix:                 I,
			Canonical:              false,
			SelfAdjoint:            true,
			LinearTextureCandidate: true,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			KindSensitive:          true,
			RequiresBridge:         true,
			IdentityLike:           true,
			QualifiedTexture:       false,
			Disqualification:       "kind-sensitive but generation-blind; Gate 171 leaves six kind assignments and no generation texture",
		},
		{
			Name:                   "J_generation",
			Kind:                   KindRealStructure,
			SourceGate:             "Gates 22/166 charge-conjugation real structure candidate",
			Matrix:                 I,
			Canonical:              true,
			SelfAdjoint:            true,
			LinearTextureCandidate: false,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			IdentityLike:           true,
			QualifiedTexture:       false,
			Disqualification:       "real structure acts as conjugation/pairing, not as a generation-breaking Yukawa amplitude matrix",
		},
		{
			Name:                   "M_source_min",
			Kind:                   KindSourceTensorMinimum,
			SourceGate:             "Gate 36 source tensor variational action",
			Matrix:                 zero,
			Canonical:              true,
			SelfAdjoint:            true,
			LinearTextureCandidate: true,
			GenerationBreaking:     false,
			ProducesMixingBasis:    false,
			ChargeCompatible:       true,
			ZeroOperator:           true,
			QualifiedTexture:       false,
			Disqualification:       "Gate 36 variational minimum is M=0; no nonzero source texture is selected",
		},
	}
	return ops
}

func auditPairs(ops []GenerationOperator) []PairAudit {
	pairs := make([]PairAudit, 0)
	for i := 0; i < len(ops); i++ {
		for j := i + 1; j < len(ops); j++ {
			norm := commutatorFrobenius(ops[i].Matrix, ops[j].Matrix)
			raw := norm > 1e-9
			bothQualified := ops[i].QualifiedTexture && ops[j].QualifiedTexture
			reason := ""
			if !bothQualified {
				reasons := []string{}
				if !ops[i].QualifiedTexture {
					reasons = append(reasons, ops[i].Name+": "+ops[i].Disqualification)
				}
				if !ops[j].QualifiedTexture {
					reasons = append(reasons, ops[j].Name+": "+ops[j].Disqualification)
				}
				reason = strings.Join(reasons, "; ")
			}
			pairs = append(pairs, PairAudit{
				A:                         ops[i].Name,
				B:                         ops[j].Name,
				RawCommutatorNorm:         norm,
				RawNonCommuting:           raw,
				BothQualifiedTextures:     bothQualified,
				QualifiedNonCommutingPair: bothQualified && raw,
				ReasonNotQualified:        reason,
			})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].RawNonCommuting != pairs[j].RawNonCommuting {
			return pairs[i].RawNonCommuting
		}
		if math.Abs(pairs[i].RawCommutatorNorm-pairs[j].RawCommutatorNorm) > 1e-12 {
			return pairs[i].RawCommutatorNorm > pairs[j].RawCommutatorNorm
		}
		return pairs[i].A+pairs[i].B < pairs[j].A+pairs[j].B
	})
	return pairs
}

func auditInventory(ops []GenerationOperator, pairs []PairAudit) InventoryAudit {
	inv := InventoryAudit{OperatorCount: len(ops)}
	for _, op := range ops {
		if op.Canonical {
			inv.CanonicalOperators++
		}
		if op.SelfAdjoint {
			inv.SelfAdjointOperators++
		}
		if op.GenerationBreaking {
			inv.GenerationBreakingOperators++
		}
		if op.LinearTextureCandidate {
			inv.LinearTextureCandidates++
		}
		if op.QualifiedTexture {
			inv.QualifiedTextureOperators++
		}
		if op.QualifiedTexture && op.Canonical && op.GenerationBreaking {
			inv.CanonicalBreakingTextures++
		}
		if op.Kind == KindBFResidual && !op.ZeroOperator {
			inv.CanonicalNonzeroCurvatureMaps++
		}
		if op.ProducesMixingBasis && op.QualifiedTexture {
			inv.CanonicalMixingSources++
		}
	}
	for _, p := range pairs {
		if p.RawNonCommuting {
			inv.RawNonCommutingPairs++
		}
		if p.QualifiedNonCommutingPair {
			inv.QualifiedNonCommutingPairs++
		}
	}
	inv.Verdict = "raw non-commuting maps exist, but every such pair contains a symmetry action or bridge-required/non-texture/zero/generation-blind object; no qualified non-commuting Yukawa texture pair exists"
	return inv
}

func buildNoGo(ops []GenerationOperator, inv InventoryAudit) NoGoAudit {
	ng := NoGoAudit{
		TrialityRawMapsAreSymmetries:        true,
		TrialityInvariantTexturesTooSmall:   true,
		DiagonalSpurionRequiresBridge:       true,
		BFResidualZero:                      true,
		ScalarShapeProjectorGenerationBlind: true,
		RealStructureGenerationBlind:        true,
		SourceTensorMinimumZero:             true,
		NoQualifiedTexturePair:              inv.QualifiedNonCommutingPairs == 0,
		MassGenerationSealedAtCurrentStage:  inv.QualifiedNonCommutingPairs == 0,
		NextRequiredInput:                   "a new finite source that is simultaneously canonical, generation-breaking, nonzero, charge-compatible, and non-commuting with another qualified texture source",
		Verdict:                             "the CKM/PMNS texture problem is sealed as structurally open at the current stage; current finite data do not contain the required non-commuting mass operators",
	}
	for _, op := range ops {
		if op.Kind == KindTrialityCycle || op.Kind == KindTrialityReflection {
			ng.TrialityRawNoncommutationSeen = true
		}
	}
	return ng
}

func identity() [3][3]float64 { return diag(1, 1, 1) }

func diag(a, b, c float64) [3][3]float64 {
	return [3][3]float64{{a, 0, 0}, {0, b, 0}, {0, 0, c}}
}

func scale(m [3][3]float64, s float64) [3][3]float64 {
	var out [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = s * m[i][j]
		}
	}
	return out
}

func commutatorFrobenius(a, b [3][3]float64) float64 {
	ab := multiply(a, b)
	ba := multiply(b, a)
	sum := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			d := ab[i][j] - ba[i][j]
			sum += d * d
		}
	}
	return math.Sqrt(sum)
}

func multiply(a, b [3][3]float64) [3][3]float64 {
	var out [3][3]float64
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			for j := 0; j < 3; j++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func FormatOperator(op GenerationOperator) string {
	flags := []string{}
	if op.Canonical {
		flags = append(flags, "canonical")
	}
	if op.SelfAdjoint {
		flags = append(flags, "self-adjoint")
	}
	if op.GenerationBreaking {
		flags = append(flags, "generation-breaking")
	}
	if op.PureSymmetryAction {
		flags = append(flags, "symmetry-action")
	}
	if op.RequiresBridge {
		flags = append(flags, "bridge-required")
	}
	if op.ZeroOperator {
		flags = append(flags, "zero")
	}
	if op.IdentityLike {
		flags = append(flags, "identity-like")
	}
	if op.QualifiedTexture {
		flags = append(flags, "qualified-texture")
	} else {
		flags = append(flags, "not-qualified")
	}
	return fmt.Sprintf("%s[%s] source=%s reason=%s", op.Name, strings.Join(flags, ","), op.SourceGate, op.Disqualification)
}

func FormatOperators(ops []GenerationOperator) string {
	parts := make([]string, 0, len(ops))
	for _, op := range ops {
		parts = append(parts, FormatOperator(op))
	}
	return strings.Join(parts, " | ")
}

func FormatPair(p PairAudit) string {
	state := "commuting"
	if p.RawNonCommuting {
		state = fmt.Sprintf("raw-noncommuting(norm=%.6g)", p.RawCommutatorNorm)
	}
	qual := "not-qualified"
	if p.QualifiedNonCommutingPair {
		qual = "qualified-noncommuting"
	}
	return fmt.Sprintf("%s vs %s: %s, %s; %s", p.A, p.B, state, qual, p.ReasonNotQualified)
}

func FormatTopRawPairs(pairs []PairAudit, n int) string {
	parts := []string{}
	for _, p := range pairs {
		if !p.RawNonCommuting {
			continue
		}
		parts = append(parts, FormatPair(p))
		if len(parts) >= n {
			break
		}
	}
	if len(parts) == 0 {
		return "no raw non-commuting pairs"
	}
	return strings.Join(parts, " | ")
}

func FormatInventory(inv InventoryAudit) string {
	return fmt.Sprintf("operators=%d canonical=%d selfAdjoint=%d genBreaking=%d linearTextureCandidates=%d qualifiedTextures=%d rawNoncommutingPairs=%d qualifiedNoncommutingPairs=%d canonicalBreakingTextures=%d nonzeroCurvatureMaps=%d mixingSources=%d verdict=%s",
		inv.OperatorCount,
		inv.CanonicalOperators,
		inv.SelfAdjointOperators,
		inv.GenerationBreakingOperators,
		inv.LinearTextureCandidates,
		inv.QualifiedTextureOperators,
		inv.RawNonCommutingPairs,
		inv.QualifiedNonCommutingPairs,
		inv.CanonicalBreakingTextures,
		inv.CanonicalNonzeroCurvatureMaps,
		inv.CanonicalMixingSources,
		inv.Verdict)
}

func FormatNoGo(ng NoGoAudit) string {
	return fmt.Sprintf("trialityRaw=%t symmetryOnly=%t invariantTooSmall=%t diagonalBridge=%t bfZero=%t shapeGenBlind=%t realStructureGenBlind=%t sourceZero=%t noQualifiedPair=%t sealed=%t next=%s verdict=%s",
		ng.TrialityRawNoncommutationSeen,
		ng.TrialityRawMapsAreSymmetries,
		ng.TrialityInvariantTexturesTooSmall,
		ng.DiagonalSpurionRequiresBridge,
		ng.BFResidualZero,
		ng.ScalarShapeProjectorGenerationBlind,
		ng.RealStructureGenerationBlind,
		ng.SourceTensorMinimumZero,
		ng.NoQualifiedTexturePair,
		ng.MassGenerationSealedAtCurrentStage,
		ng.NextRequiredInput,
		ng.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gaugeClosed=%t scalarTarget=%t massYukawa=%t pairRequired=%t pairFound=%t yukawa=%t masses=%t ckmPmns=%t constants=%t nullity=%d->%d next=%s verdict=%s",
		f.GaugeRatioClosed,
		f.ScalarShapeTargetAvailable,
		f.MassProblemLocalizedToYukawaMatrix,
		f.NonCommutingTexturePairRequired,
		f.NonCommutingTexturePairFound,
		f.YukawaAmplitudesDerived,
		f.FermionMassesDerived,
		f.CKMPMNSDerived,
		f.PhysicalConstantsDerived,
		f.ResidualNullityBefore,
		f.ResidualNullityAfter,
		f.RecommendedNextGate,
		f.Verdict)
}
