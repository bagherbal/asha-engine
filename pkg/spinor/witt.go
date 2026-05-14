package spinor

import "fmt"

// WittPair records the native bookkeeping bridge from one complex Fock mode to
// a real two-plane in the eight-dimensional Clifford carrier.  The engine keeps
// this as typed structure so later gates do not identify a mode label with a
// Lie-algebra coordinate by name alone.
type WittPair struct {
	ModeIndex           int
	CreationName        string
	AnnihilationName    string
	Kind                ModeKind
	RealBasisA          int
	RealBasisB          int
	BivectorLabel       string
	CreationFormula     string
	AnnihilationFormula string
}

// WittDecomposition is the explicit four-mode / eight-real-axis dictionary
// used by the finite Fock bookkeeping.  It is a representation dictionary, not
// a fitted particle-physics input.
type WittDecomposition struct {
	RealDimension    int
	ComplexModeCount int
	Pairs            []WittPair
	Convention       string
}

// NativeWittDecomposition returns the current native mode pairing.  It is kept
// deliberately small and explicit: mode k is paired with the real two-plane
// (e_{2k}, e_{2k+1}).
func NativeWittDecomposition(spacetimeDimension int) (WittDecomposition, error) {
	f, err := NewCovariantPhaseFockSpace(spacetimeDimension)
	if err != nil {
		return WittDecomposition{}, err
	}
	pairs := make([]WittPair, 0, f.ModeCount())
	for _, mode := range f.Modes {
		a := 2 * mode.Index
		b := a + 1
		pairs = append(pairs, WittPair{
			ModeIndex:           mode.Index,
			CreationName:        mode.Name,
			AnnihilationName:    fmt.Sprintf("a_%d", mode.Index),
			Kind:                mode.Kind,
			RealBasisA:          a,
			RealBasisB:          b,
			BivectorLabel:       fmt.Sprintf("e%d∧e%d", a, b),
			CreationFormula:     fmt.Sprintf("a†_%d = 1/2(e%d - i e%d)", mode.Index, a, b),
			AnnihilationFormula: fmt.Sprintf("a_%d = 1/2(e%d + i e%d)", mode.Index, a, b),
		})
	}
	return WittDecomposition{
		RealDimension:    2 * f.ModeCount(),
		ComplexModeCount: f.ModeCount(),
		Pairs:            pairs,
		Convention:       "native mode k ↔ real two-plane (e_{2k},e_{2k+1}); N_k contributes only its grade-2 Cartan coordinate to so(8), while scalar identity shifts are outside the Lie algebra",
	}, nil
}
