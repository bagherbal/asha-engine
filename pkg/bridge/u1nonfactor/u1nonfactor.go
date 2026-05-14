// Package u1nonfactor implements Gate 77: the non-factorized abelian action /
// kinetic-mixing search.
//
// Gate 76 showed that factorized trace sources such as Tr(B-L)Tr(T_phi)
// vanish.  This gate tests the next natural non-factorized candidate already
// present in the finite engine: the Yukawa-incidence correlation between the
// matter-side B-L charge and the scalar/contact T_phi charge.  This is a real
// finite source tensor because it is supported only on gauge-compatible Yukawa
// channels, not on the full tensor product.
//
// The result is again strict: the non-factorized channel support exists and has
// nonzero local/norm content, but its signed total cancels between up/down and
// neutrino/electron branches.  Therefore it does not yet generate a net
// B-L/contact-u1 kinetic mixing Hessian or a physical U(1)_Y coupling.
package u1nonfactor

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/u1source"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type KindContribution struct {
	Kind               yukawaintertwiner.FermionKind
	Channels           int
	FiberEntries       int
	SignedBLContact    float64
	SignedCentral      float64
	AbsoluteBLContact  float64
	QuadraticBLContact float64
}

type Analysis struct {
	Previous u1source.Analysis
	Yukawa   yukawaintertwiner.Analysis

	FullTensorDimension         int
	YukawaSupportEntries        int
	SupportFraction             float64
	NonFactorizedSupportDerived bool

	SignedBLContactMoment      float64
	SignedCentralContactMoment float64
	AbsoluteBLContactMoment    float64
	QuadraticBLContactMoment   float64
	BLContactRMS               float64

	KindContributions []KindContribution

	UpDownCancellation      bool
	LeptonPairCancellation  bool
	TotalSignedCancellation bool
	LocalNonzeroCorrelation bool

	NonFactorizedActionDerived bool
	CrossCarrierSourceDerived  bool
	FullU1HessianDerived       bool
	PhysicalU1CouplingDerived  bool
	FineStructureDerived       bool
	HiddenObservedInputUsed    bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := u1source.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		y, err := yukawaintertwiner.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, y, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev u1source.Analysis, y yukawaintertwiner.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !y.ChargeCompatibleYukawaChannelsDerived {
		return Analysis{}, fmt.Errorf("Gate 77 requires gauge-compatible Yukawa channels")
	}

	// Full tensor domain uses left doublets (8) times scalar fiber (4).  The
	// Yukawa-incidence support has only the allowed scalar-fiber entries.
	scalarFiberDim := 0
	for _, b := range y.ScalarBranches {
		scalarFiberDim += b.Multiplicity
	}
	fullTensor := y.LeftDimension * scalarFiberDim

	byKind := map[yukawaintertwiner.FermionKind]*KindContribution{}
	signedBL := 0.0
	signedCentral := 0.0
	absBL := 0.0
	quadBL := 0.0
	entries := 0

	for _, ch := range y.Channels {
		bl := bMinusL(ch.Left)
		tphi := ch.Scalar.Hypercharge
		c := 1.0
		k := ch.Right.Kind
		if _, ok := byKind[k]; !ok {
			byKind[k] = &KindContribution{Kind: k}
		}
		kc := byKind[k]
		kc.Channels++
		for f := 0; f < ch.Scalar.Multiplicity; f++ {
			v := bl * tphi
			cv := c * tphi
			signedBL += v
			signedCentral += cv
			absBL += math.Abs(v)
			quadBL += v * v
			entries++
			kc.FiberEntries++
			kc.SignedBLContact += v
			kc.SignedCentral += cv
			kc.AbsoluteBLContact += math.Abs(v)
			kc.QuadraticBLContact += v * v
		}
	}

	kinds := make([]KindContribution, 0, len(byKind))
	for _, v := range byKind {
		kinds = append(kinds, *v)
	}
	sort.Slice(kinds, func(i, j int) bool { return kinds[i].Kind < kinds[j].Kind })

	up := byKind[yukawaintertwiner.UpType]
	down := byKind[yukawaintertwiner.DownType]
	nu := byKind[yukawaintertwiner.NeutrinoType]
	e := byKind[yukawaintertwiner.ElectronType]
	upDownCancel := up != nil && down != nil && math.Abs(up.SignedBLContact+down.SignedBLContact) <= eps
	leptonCancel := nu != nil && e != nil && math.Abs(nu.SignedBLContact+e.SignedBLContact) <= eps
	totalCancel := math.Abs(signedBL) <= eps && math.Abs(signedCentral) <= eps
	localNonzero := absBL > eps && quadBL > eps
	supportFraction := 0.0
	if fullTensor > 0 {
		supportFraction = float64(entries) / float64(fullTensor)
	}
	rms := 0.0
	if entries > 0 {
		rms = math.Sqrt(quadBL / float64(entries))
	}

	truth := "Gate 77 constructs the first non-factorized abelian source candidate: the Yukawa-incidence correlation between matter B-L and scalar/contact T_phi. The support is finite and nonzero locally, but the signed moment cancels exactly between up/down quark branches and neutrino/electron lepton branches. Thus the current finite data still does not produce a net B-L/contact-u1 kinetic mixing Hessian."

	return Analysis{
		Previous:                    prev,
		Yukawa:                      y,
		FullTensorDimension:         fullTensor,
		YukawaSupportEntries:        entries,
		SupportFraction:             supportFraction,
		NonFactorizedSupportDerived: entries == y.FiberEntryCount && entries > 0,
		SignedBLContactMoment:       signedBL,
		SignedCentralContactMoment:  signedCentral,
		AbsoluteBLContactMoment:     absBL,
		QuadraticBLContactMoment:    quadBL,
		BLContactRMS:                rms,
		KindContributions:           kinds,
		UpDownCancellation:          upDownCancel,
		LeptonPairCancellation:      leptonCancel,
		TotalSignedCancellation:     totalCancel,
		LocalNonzeroCorrelation:     localNonzero,
		NonFactorizedActionDerived:  true,
		CrossCarrierSourceDerived:   false,
		FullU1HessianDerived:        false,
		PhysicalU1CouplingDerived:   false,
		FineStructureDerived:        false,
		HiddenObservedInputUsed:     false,
		TruthStatement:              truth,
		RecommendedNextGate:         "Gate 78 — Chiral / Orientational Abelian Source Search",
		RemainingUnknowns: []string{
			"U-20D3C1-CHIRAL-ABELIAN-SOURCE: derive a chiral or orientation-weighted source that avoids the up/down and lepton-pair cancellation, if one exists",
			"U-20D3C2-NONFACTORIZED-U1-HESSIAN: convert incidence correlations into an actual kinetic Hessian second variation",
			"U-20D3C3-U1-PHYSICAL-COUPLING: no g_Y or alpha follows until the Hessian and RG scale are derived",
			"U-20D3C4-SIGN-ORIENTATION: determine whether scalar branch orientation is selected by finite contact geometry or by an external convention",
		},
	}, nil
}

func bMinusL(s su2lgauge.LeftDoubletState) float64 {
	if s.Kind == su2lgauge.QuarkDoublet {
		return 1.0 / 3.0
	}
	return -1.0
}
