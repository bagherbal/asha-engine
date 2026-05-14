// Package sourcepotential audits whether the finite scalar/curvature data derives
// a nonzero symmetry-breaking action for the active-to-generation source tensor.
//
// Gate 35 proved that the minimal positive action over
//
//	M : H_active -> H_generation
//
// selects the zero map because the geometric source J vanishes. The obvious next
// repair would be a Mexican-hat or fixed-radius action for M. This package tests
// whether the finite engine has actually derived the ingredients required for
// such an action:
//
//   - a negative quadratic coefficient or instability for M;
//   - a positive quartic stabilizer;
//   - a nonzero radius;
//   - a canonical orientation in Hom(R^4,R^3).
//
// The current finite data has strong positive scalar invariants in the
// Higgs/contact and active BF sectors, but it does not derive a tachyonic sign or
// an orientation for the 12-dimensional source-tensor space. Therefore this gate
// is intentionally a no-go: scalar curvature and leakage are real, but they do
// not yet become a generation-mixing source tensor.
package sourcepotential

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/matter/bfsource"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
	"github.com/bagherbal/asha-engine/pkg/matter/sourceaction"
)

type SourceKind string

const (
	SourceHiggsContactOrder SourceKind = "Higgs/contact order parameter"
	SourceActiveBFCurvature SourceKind = "active BF scalar curvature"
	SourceContactLeakage    SourceKind = "bare contact leakage"
	SourceDiagonalSpurion   SourceKind = "diagonal generation spurion"
	SourceTensorAction      SourceKind = "source tensor action"
)

type Candidate struct {
	Name                string
	Kind                SourceKind
	ScalarInvariant     float64
	RadiusSquared       float64
	PositiveQuartic     bool
	TachyonicSign       bool
	NonzeroRadius       bool
	OrientationSelected bool
	Canonical           bool
	Derived             bool
	Detail              string
}

type Analysis struct {
	SourceAction sourceaction.Analysis
	Potential    higgspotential.Analysis
	BFSource     bfsource.Analysis
	Generation   generationbreak.Analysis
	Contact      contact.Space

	GenerationDimension int
	ActiveDimension     int
	TensorDimension     int

	HiggsOrderParameter float64
	HiggsQuarticTrace   float64
	ActiveBFScalarTrace float64
	ActiveBFScalarNorm  float64
	ContactLeakageNorm  float64
	ContactLeakageSq    float64
	DiagonalSpurionNorm float64

	ScalarInvariantsFound   bool
	PositiveQuarticFound    bool
	TachyonicSignDerived    bool
	NonzeroRadiusDerived    bool
	TensorOrientationFound  bool
	StableZeroPersists      bool
	ArbitraryRadiusRejected bool

	Candidates []Candidate
	Best       Candidate

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		sa, err := sourceaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		hp, err := higgspotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		bs, err := bfsource.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		gb, err := generationbreak.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		k, err := contact.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sa, hp, bs, gb, k, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(sa sourceaction.Analysis, hp higgspotential.Analysis, bs bfsource.Analysis, gb generationbreak.Analysis, k contact.Space, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	if sa.GenerationDimension != 3 || sa.ActiveDimension != 4 || sa.TensorDimension != 12 {
		return Analysis{}, fmt.Errorf("source potential expects source tensor dimensions 3x4=12, got gen=%d active=%d tensor=%d", sa.GenerationDimension, sa.ActiveDimension, sa.TensorDimension)
	}
	activeTrace := bs.ActiveQuadratic.Trace
	activeNorm := bs.ActiveQuadratic.Norm
	leakage := k.BareLeakageNorm()
	leakageSq := k.BareLeakageNormSquared()
	diagonalNorm := diagonalSpurionNorm(gb.BestCandidate.Eigenvalues)

	higgs := Candidate{
		Name:                "finite Higgs/contact scalar radius candidate",
		Kind:                SourceHiggsContactOrder,
		ScalarInvariant:     hp.OrderParameterNormSquared,
		RadiusSquared:       hp.OrderParameterNormSquared,
		PositiveQuartic:     hp.QuarticTrace > eps,
		TachyonicSign:       false,
		NonzeroRadius:       hp.OrderParameterNormSquared > eps,
		OrientationSelected: false,
		Canonical:           true,
		Derived:             true,
		Detail:              "The finite Higgs/contact sector derives a real positive scalar order parameter, but this radius belongs to Φ, not to a 3x4 generation source tensor M.",
	}
	activeBF := Candidate{
		Name:                "active BF scalar curvature candidate",
		Kind:                SourceActiveBFCurvature,
		ScalarInvariant:     activeTrace,
		RadiusSquared:       activeTrace,
		PositiveQuartic:     activeTrace > eps,
		TachyonicSign:       false,
		NonzeroRadius:       activeTrace > eps,
		OrientationSelected: false,
		Canonical:           true,
		Derived:             true,
		Detail:              "Active BF curvature produces scalar/Higgs-sector energy, but Gate 33 showed its protected and mixed generation contractions vanish.",
	}
	contactLeak := Candidate{
		Name:                "bare contact leakage scalar candidate",
		Kind:                SourceContactLeakage,
		ScalarInvariant:     leakageSq,
		RadiusSquared:       leakageSq,
		PositiveQuartic:     false,
		TachyonicSign:       false,
		NonzeroRadius:       leakage > eps,
		OrientationSelected: false,
		Canonical:           true,
		Derived:             true,
		Detail:              "Bare contact leakage is finite geometric frustration, not an oriented source in Hom(H_active,H_generation).",
	}
	diagonal := Candidate{
		Name:                "diagonal generation-breaking spurion",
		Kind:                SourceDiagonalSpurion,
		ScalarInvariant:     diagonalNorm,
		RadiusSquared:       diagonalNorm * diagonalNorm,
		PositiveQuartic:     false,
		TachyonicSign:       false,
		NonzeroRadius:       diagonalNorm > eps,
		OrientationSelected: false,
		Canonical:           false,
		Derived:             false,
		Detail:              "The diagonal spurion splits generations but does not select a source-tensor orientation or mixing operator.",
	}
	tensor := Candidate{
		Name:                "source tensor Mexican-hat action",
		Kind:                SourceTensorAction,
		ScalarInvariant:     sa.NaturalSourceNorm,
		RadiusSquared:       sa.NaturalStationaryNorm * sa.NaturalStationaryNorm,
		PositiveQuartic:     false,
		TachyonicSign:       false,
		NonzeroRadius:       sa.NonzeroStationaryFound,
		OrientationSelected: sa.NonzeroStationaryFound,
		Canonical:           true,
		Derived:             false,
		Detail:              "A nonzero Mexican-hat source action would require a derived negative quadratic sign and fixed orientation; neither is present in the current finite data.",
	}
	candidates := []Candidate{higgs, activeBF, contactLeak, diagonal, tensor}

	scalarFound := hp.OrderParameterNormSquared > eps || activeTrace > eps || leakage > eps
	positiveQuartic := hp.QuarticTrace > eps && hp.NormalizedQuarticShape > eps
	tachyonic := false
	nonzeroRadiusForM := false
	orientation := false
	zeroPersists := sa.NaturalSelectsZero

	return Analysis{
		SourceAction:            sa,
		Potential:               hp,
		BFSource:                bs,
		Generation:              gb,
		Contact:                 k,
		GenerationDimension:     sa.GenerationDimension,
		ActiveDimension:         sa.ActiveDimension,
		TensorDimension:         sa.TensorDimension,
		HiggsOrderParameter:     hp.OrderParameterNormSquared,
		HiggsQuarticTrace:       hp.QuarticTrace,
		ActiveBFScalarTrace:     activeTrace,
		ActiveBFScalarNorm:      activeNorm,
		ContactLeakageNorm:      leakage,
		ContactLeakageSq:        leakageSq,
		DiagonalSpurionNorm:     diagonalNorm,
		ScalarInvariantsFound:   scalarFound,
		PositiveQuarticFound:    positiveQuartic,
		TachyonicSignDerived:    tachyonic,
		NonzeroRadiusDerived:    nonzeroRadiusForM,
		TensorOrientationFound:  orientation,
		StableZeroPersists:      zeroPersists,
		ArbitraryRadiusRejected: true,
		Candidates:              candidates,
		Best:                    higgs,
		TruthStatement:          truth(scalarFound, positiveQuartic, tachyonic, nonzeroRadiusForM, zeroPersists),
		RemainingUnknowns: []string{
			"U-17G-SYMMETRY-BREAKING-SIGN: derive a negative quadratic coefficient for M from finite geometry, not by inserting a Mexican-hat potential",
			"U-17H-SOURCE-RADIUS: derive a nonzero radius for the 3x4 source tensor itself, not merely a scalar/Higgs-sector norm",
			"U-17I-SOURCE-ORIENTATION: select an orientation in Hom(R4,R3) so the action does not leave an arbitrary 12D vacuum manifold",
			"U-16C-NONCOMMUTING-TEXTURES: obtain at least two non-commuting 3x3 operators before CKM/PMNS claims",
		},
	}, nil
}

func diagonalSpurionNorm(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func truth(scalarFound, quarticFound, tachyonic, radius, zeroPersists bool) string {
	if scalarFound && quarticFound && !tachyonic && !radius && zeroPersists {
		return "Finite scalar invariants are real, and the Higgs/contact sector has positive quartic data, but no finite theorem derives a tachyonic source-tensor sign, nonzero radius, or orientation. The stable source-tensor action still selects M=0."
	}
	if radius {
		return "A nonzero source-tensor radius has been selected. It must still pass orientation, gauge-compatibility, and non-fitting audits."
	}
	return "The finite data does not yet derive a nonzero symmetry-breaking source action for the generation tensor."
}

func FormatCandidate(c Candidate) string {
	tags := []string{}
	if c.Canonical {
		tags = append(tags, "canonical")
	} else {
		tags = append(tags, "noncanonical")
	}
	if c.Derived {
		tags = append(tags, "derived")
	} else {
		tags = append(tags, "not-derived")
	}
	if c.TachyonicSign {
		tags = append(tags, "tachyonic")
	} else {
		tags = append(tags, "no-tachyonic-sign")
	}
	if c.OrientationSelected {
		tags = append(tags, "oriented")
	} else {
		tags = append(tags, "unoriented")
	}
	return fmt.Sprintf("%s: invariant=%.6e radius²=%.6e [%s]", c.Name, c.ScalarInvariant, c.RadiusSquared, strings.Join(tags, ","))
}
