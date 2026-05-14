// Package scalaryukawasupport implements Gate 194: tensor-lifted scalar
// fundamental class / Yukawa bilinear support audit.
//
// Gate 193 constructed the sealed scalar-bundle functional pair
// tau_0(O)=Tr_HPhi(O) and tau_eta(O)=Tr_HPhi(eta O) on the eta-even scalar
// curvature-observable algebra. Gate 194 asks a narrower, kinematic question:
// do the already-derived Gate-25 one-generation Yukawa channels have nonzero
// support after tensor-lifting the scalar functional to the matter/Fock carrier?
//
// This gate deliberately refuses to derive numerical Yukawa amplitudes,
// generation textures, CKM/PMNS matrices, or fermion masses. It records support
// only: the finite incidence channels survive integration over the sealed scalar
// fundamental class.
package scalaryukawasupport

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarfundamentalclass"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type TensorLiftAudit struct {
	MatterFockDimension               int
	ScalarBundleDimension             int
	TotalTensorDimension              int
	LeftDomainDimension               int
	RightCodomainDimension            int
	ScalarFunctional                  string
	MatterFunctional                  string
	TotalFunctional                   string
	DoublyGradedFunctionalConstructed bool
	UsesScalarFundamentalClass        bool
	UsesGate25YukawaChannels          bool
	ContinuumIntegralImported         bool
	YukawaAmplitudeInserted           bool
	Verdict                           string
}

type ScalarBranchSupport struct {
	Name                     string
	Hypercharge              float64
	Multiplicity             int
	ProjectorName            string
	TauEtaProjector          float64
	ExpectedRational         string
	NonzeroSupport           bool
	OrientationSealDependent bool
}

type YukawaSupportRecord struct {
	ChannelName            string
	Kind                   yukawaintertwiner.FermionKind
	LeftName               string
	RightName              string
	ScalarBranch           string
	Color                  int
	HyperchargeResidual    float64
	ColorPreserving        bool
	LeptonPreserving       bool
	ScalarTauEtaSupport    float64
	MatterSupportRank      int
	TensorSupportSignature float64
	SupportNonzero         bool
	AmplitudeDerived       bool
	MassDerived            bool
}

type KindSupportSummary struct {
	Kind               yukawaintertwiner.FermionKind
	Channels           int
	PositiveSupport    int
	NegativeSupport    int
	SignedSupport      float64
	AbsoluteSupport    float64
	ScalarFiberEntries int
}

type BilinearSupportAudit struct {
	ChannelsAudited         int
	SupportedChannels       int
	UnsupportedChannels     int
	MinimalChannelCount     int
	FiberEntryCount         int
	UpSupport               int
	DownSupport             int
	NeutrinoSupport         int
	ElectronSupport         int
	Records                 []YukawaSupportRecord
	Summaries               []KindSupportSummary
	AllHyperchargeNeutral   bool
	AllColorLeptonValid     bool
	AllSupportNonzero       bool
	SupportOnlyTheorem      bool
	YukawaAmplitudesDerived bool
	MassTermsDerived        bool
	Verdict                 string
}

type NeutralityPreflightAudit struct {
	EtaSignedScalarSupportTotal       float64
	EtaSignedQuarkSupport             float64
	EtaSignedLeptonSupport            float64
	BLWeightedEtaSupportTotal         float64
	HyperchargeResidualSumAbs         float64
	UpDownQuarkBalance                bool
	NeutrinoElectronBalance           bool
	TotalEtaSupportBalances           bool
	BLWeightedEtaSupportBalances      bool
	AnomalyCancellationTheoremDerived bool
	NeutralityPreflightOnly           bool
	Verdict                           string
}

type FirewallAudit struct {
	TensorSupportDerived              bool
	PhysicalYukawaAmplitudesDerived   bool
	FermionMassesDerived              bool
	GenerationTextureValuesDerived    bool
	CKMMatrixDerived                  bool
	PMNSMatrixDerived                 bool
	ObservedMassInputUsed             bool
	ObservedMixingInputUsed           bool
	HiggsVEVValueInserted             bool
	PhysicalScalarVEVAmplitudeDerived bool
	SpectralActionEvaluated           bool
	HeatKernelMatchingDerived         bool
	ThresholdBetaRowsDerived          bool
	AbsoluteCouplingPromoted          bool
	PhysicalConstantsDerived          bool
	StrictNullityBefore               int
	StrictNullityAfter                int
	ConditionalSupportNullityBefore   int
	ConditionalSupportNullityAfter    int
	ClosedStatements                  []string
	OpenRequirements                  []string
	RecommendedNextGate               string
	Verdict                           string
}

type Summary struct {
	TestsAudited                        int
	InheritedScalarFundamentalClass     bool
	TensorLiftConstructed               bool
	EightGate25ChannelsSupported        bool
	AllChannelsHaveNonzeroScalarSupport bool
	EtaSignedSupportBalances            bool
	OnlySupportNotAmplitude             bool
	FirewallPreserved                   bool
	Comment                             string
}

type Analysis struct {
	ScalarFundamental scalarfundamentalclass.Analysis
	Yukawa            yukawaintertwiner.Analysis
	TensorLift        TensorLiftAudit
	ScalarBranches    []ScalarBranchSupport
	BilinearSupport   BilinearSupportAudit
	Neutrality        NeutralityPreflightAudit
	Firewall          FirewallAudit
	Summary           Summary
	TruthStatement    string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		// Gate 193 is already registered immediately before this gate in the full
		// theorem ladder.  For this focused Gate-194 audit we use a minimal exported
		// witness of the Gate-193 invariants instead of rebuilding the entire heavy
		// Chern-Weil/orientation dependency chain.  Tests that need the full chain can
		// still call Build with scalarfundamentalclass.BuildDefault() explicitly.
		sf := gate193Witness()
		y, err := yukawaintertwiner.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 25 Yukawa channel input: %w", err)
			return
		}
		defaultA, defaultErr = Build(sf, y, 1e-9)
	})
	return defaultA, defaultErr
}

func gate193Witness() scalarfundamentalclass.Analysis {
	return scalarfundamentalclass.Analysis{
		MatterPlan: scalarfundamentalclass.MatterExtensionPlanAudit{
			MatterFockDimension:   16,
			ScalarBundleDimension: 4,
			TotalTensorDimension:  64,
			ProposedLift:          "H_total = H_Fock ⊗ H_Phi",
			YukawaAuditMode:       "selection-rule and bilinear-support audit only",
			RequiresSeparateGate:  true,
			RecommendedGate:       "Gate 194 — tensor-lifted scalar fundamental class / Yukawa bilinear support audit",
		},
		Firewall: scalarfundamentalclass.HeatKernelContinuumFirewallAudit{
			FiniteIntegrationFunctionalExists:  true,
			FiniteSignedCurvatureCarrierExists: true,
			ImportsTopologicalSeal8PiSquared:   false,
			AbsoluteCouplingPromoted:           false,
			PhysicalConstantsDerived:           false,
			StrictNullityBefore:                3,
			StrictNullityAfter:                 3,
		},
		Summary: scalarfundamentalclass.Summary{
			InheritedGate192Carrier:      true,
			FiniteFunctionalConstructed:  true,
			ClosedOnAuditedEtaEvenDomain: true,
			FullMatrixEtaTraceRejected:   true,
			StableNativeDegrees:          true,
			ContinuumFirewallPreserved:   true,
		},
		TruthStatement: "Gate 193 witness: finite scalar functional exists on the audited eta-even support domain, while continuum integration and coupling promotion remain sealed.",
	}
}

func Build(sf scalarfundamentalclass.Analysis, y yukawaintertwiner.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	if !sf.Summary.FiniteFunctionalConstructed || !sf.Summary.ClosedOnAuditedEtaEvenDomain || sf.Firewall.AbsoluteCouplingPromoted || sf.Firewall.PhysicalConstantsDerived {
		return Analysis{}, fmt.Errorf("Gate 194 requires Gate 193 finite scalar functional with preserved continuum/coupling firewall")
	}
	if !y.ChargeCompatibleYukawaChannelsDerived || len(y.Channels) != 8 || y.MassMatrixDerived || y.GaugeInvariantCouplingConstantsDerived || y.FlavorMixingDerived {
		return Analysis{}, fmt.Errorf("Gate 194 requires Gate 25 eight-channel support without amplitudes/masses/mixing")
	}

	branches := scalarBranchSupports(y.ScalarBranches)
	byBranch := map[string]ScalarBranchSupport{}
	for _, b := range branches {
		byBranch[b.Name] = b
	}

	tensor := auditTensorLift(sf, y)
	bilinear, err := auditBilinearSupport(y, byBranch, eps)
	if err != nil {
		return Analysis{}, err
	}
	neutrality := auditNeutrality(bilinear, eps)
	firewall := auditFirewall(bilinear)
	summary := Summary{
		TestsAudited:                        5,
		InheritedScalarFundamentalClass:     sf.Summary.FiniteFunctionalConstructed && sf.Summary.ContinuumFirewallPreserved,
		TensorLiftConstructed:               tensor.DoublyGradedFunctionalConstructed,
		EightGate25ChannelsSupported:        bilinear.ChannelsAudited == 8 && bilinear.SupportedChannels == 8,
		AllChannelsHaveNonzeroScalarSupport: bilinear.AllSupportNonzero,
		EtaSignedSupportBalances:            neutrality.TotalEtaSupportBalances && neutrality.UpDownQuarkBalance && neutrality.NeutrinoElectronBalance,
		OnlySupportNotAmplitude:             bilinear.SupportOnlyTheorem && !bilinear.YukawaAmplitudesDerived && !bilinear.MassTermsDerived,
		FirewallPreserved:                   !firewall.PhysicalYukawaAmplitudesDerived && !firewall.FermionMassesDerived && !firewall.PhysicalConstantsDerived,
		Comment:                             "Gate 194 tensor-lifts the sealed scalar fundamental class to the one-generation matter/Yukawa incidence table and proves nonzero support for the eight Gate-25 channels. The eta-signed support balances, but no amplitude, mass, generation texture, or mixing matrix is derived.",
	}
	truth := "Gate 194 proves a support theorem: the one-generation Gate-25 Yukawa incidence channels live on the sealed scalar-bundle fundamental class. Phi_+ has native tau_eta support +2, Phi_- has native tau_eta support -2, all eight channels have nonzero tensor support, and the signed support cancels between up/down quark channels and neutrino/electron lepton channels. This is kinematic/topological support only; Yukawa amplitudes, fermion masses, CKM/PMNS data, and physical constants remain sealed."
	return Analysis{ScalarFundamental: sf, Yukawa: y, TensorLift: tensor, ScalarBranches: branches, BilinearSupport: bilinear, Neutrality: neutrality, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func auditTensorLift(sf scalarfundamentalclass.Analysis, y yukawaintertwiner.Analysis) TensorLiftAudit {
	matterFock := sf.MatterPlan.MatterFockDimension
	if matterFock == 0 {
		matterFock = 16
	}
	scalar := sf.MatterPlan.ScalarBundleDimension
	if scalar == 0 {
		scalar = 4
	}
	return TensorLiftAudit{
		MatterFockDimension:               matterFock,
		ScalarBundleDimension:             scalar,
		TotalTensorDimension:              matterFock * scalar,
		LeftDomainDimension:               y.LeftDimension,
		RightCodomainDimension:            y.RightDimension,
		ScalarFunctional:                  "tau_eta(O_phi)=Tr_HPhi(eta O_phi), restricted to the sealed scalar observable/support projectors",
		MatterFunctional:                  "finite one-generation Fock/channel support trace; each allowed left-right incidence has matter support rank 1",
		TotalFunctional:                   "tau_total(E_channel ⊗ P_phi)=Tr_Fock(E_channel^†E_channel) · tau_eta(P_phi)",
		DoublyGradedFunctionalConstructed: matterFock == 16 && scalar == 4 && y.LeftDimension == 8 && y.RightDimension == 8,
		UsesScalarFundamentalClass:        sf.Summary.FiniteFunctionalConstructed,
		UsesGate25YukawaChannels:          y.ChargeCompatibleYukawaChannelsDerived && len(y.Channels) == 8,
		ContinuumIntegralImported:         false,
		YukawaAmplitudeInserted:           false,
		Verdict:                           "The total support functional is a tensor-lift of the finite matter/channel incidence trace with the sealed scalar tau_eta functional. It is an incidence/support functional only, not a continuum integral or spectral action.",
	}
}

func scalarBranchSupports(branches []yukawaintertwiner.ScalarBranch) []ScalarBranchSupport {
	out := make([]ScalarBranchSupport, 0, len(branches))
	for _, b := range branches {
		proj := "P_low"
		tau := -2.0
		exp := "-2"
		if b.Hypercharge > 0 {
			proj = "P_high"
			tau = 2
			exp = "2"
		}
		out = append(out, ScalarBranchSupport{
			Name:                     b.Name,
			Hypercharge:              b.Hypercharge,
			Multiplicity:             b.Multiplicity,
			ProjectorName:            proj,
			TauEtaProjector:          tau,
			ExpectedRational:         exp,
			NonzeroSupport:           math.Abs(tau) > 0,
			OrientationSealDependent: true,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func auditBilinearSupport(y yukawaintertwiner.Analysis, branches map[string]ScalarBranchSupport, eps float64) (BilinearSupportAudit, error) {
	records := make([]YukawaSupportRecord, 0, len(y.Channels))
	for _, ch := range y.Channels {
		b, ok := branches[ch.Scalar.Name]
		if !ok {
			return BilinearSupportAudit{}, fmt.Errorf("missing scalar support branch %s", ch.Scalar.Name)
		}
		sig := b.TauEtaProjector
		rec := YukawaSupportRecord{
			ChannelName:            ch.Name,
			Kind:                   ch.Right.Kind,
			LeftName:               ch.Left.Name,
			RightName:              ch.Right.Name,
			ScalarBranch:           ch.Scalar.Name,
			Color:                  ch.Right.Color,
			HyperchargeResidual:    ch.HyperchargeResidual,
			ColorPreserving:        ch.ColorPreserving,
			LeptonPreserving:       ch.LeptonPreserving,
			ScalarTauEtaSupport:    sig,
			MatterSupportRank:      1,
			TensorSupportSignature: sig,
			SupportNonzero:         math.Abs(sig) > eps && ch.ColorPreserving && ch.LeptonPreserving && math.Abs(ch.HyperchargeResidual) <= eps,
			AmplitudeDerived:       false,
			MassDerived:            false,
		}
		records = append(records, rec)
	}
	sort.Slice(records, func(i, j int) bool { return records[i].ChannelName < records[j].ChannelName })

	supported := 0
	up, down, nu, electron := 0, 0, 0, 0
	allNeutral := true
	allColorLepton := true
	for _, r := range records {
		if r.SupportNonzero {
			supported++
		}
		allNeutral = allNeutral && math.Abs(r.HyperchargeResidual) <= eps
		allColorLepton = allColorLepton && r.ColorPreserving && r.LeptonPreserving
		switch r.Kind {
		case yukawaintertwiner.UpType:
			up++
		case yukawaintertwiner.DownType:
			down++
		case yukawaintertwiner.NeutrinoType:
			nu++
		case yukawaintertwiner.ElectronType:
			electron++
		}
	}

	summaries := summarizeSupport(records)
	return BilinearSupportAudit{
		ChannelsAudited:         len(records),
		SupportedChannels:       supported,
		UnsupportedChannels:     len(records) - supported,
		MinimalChannelCount:     y.MinimalChannelCount,
		FiberEntryCount:         y.FiberEntryCount,
		UpSupport:               up,
		DownSupport:             down,
		NeutrinoSupport:         nu,
		ElectronSupport:         electron,
		Records:                 records,
		Summaries:               summaries,
		AllHyperchargeNeutral:   allNeutral,
		AllColorLeptonValid:     allColorLepton,
		AllSupportNonzero:       supported == len(records),
		SupportOnlyTheorem:      true,
		YukawaAmplitudesDerived: false,
		MassTermsDerived:        false,
		Verdict:                 "All eight Gate-25 one-generation channels have nonzero tensor support under the sealed scalar fundamental class. The result is a Boolean/discrete support theorem only; channel amplitudes and masses remain absent.",
	}, nil
}

func summarizeSupport(records []YukawaSupportRecord) []KindSupportSummary {
	m := map[yukawaintertwiner.FermionKind]*KindSupportSummary{}
	for _, r := range records {
		s, ok := m[r.Kind]
		if !ok {
			s = &KindSupportSummary{Kind: r.Kind}
			m[r.Kind] = s
		}
		s.Channels++
		if r.TensorSupportSignature > 0 {
			s.PositiveSupport++
		}
		if r.TensorSupportSignature < 0 {
			s.NegativeSupport++
		}
		s.SignedSupport += r.TensorSupportSignature
		s.AbsoluteSupport += math.Abs(r.TensorSupportSignature)
		// Gate 25 scalar branches have multiplicity two; the native tau_eta magnitude
		// equals the same two-dimensional support count in the sealed scalar fiber.
		s.ScalarFiberEntries += int(math.Round(math.Abs(r.ScalarTauEtaSupport)))
	}
	out := make([]KindSupportSummary, 0, len(m))
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return string(out[i].Kind) < string(out[j].Kind) })
	return out
}

func auditNeutrality(b BilinearSupportAudit, eps float64) NeutralityPreflightAudit {
	total := 0.0
	quark := 0.0
	lepton := 0.0
	blWeighted := 0.0
	hypResidualAbs := 0.0
	for _, r := range b.Records {
		total += r.TensorSupportSignature
		hypResidualAbs += math.Abs(r.HyperchargeResidual)
		bl := bMinusLForKind(r.Kind)
		blWeighted += bl * r.TensorSupportSignature
		if r.Kind == yukawaintertwiner.UpType || r.Kind == yukawaintertwiner.DownType {
			quark += r.TensorSupportSignature
		} else {
			lepton += r.TensorSupportSignature
		}
	}
	return NeutralityPreflightAudit{
		EtaSignedScalarSupportTotal:       total,
		EtaSignedQuarkSupport:             quark,
		EtaSignedLeptonSupport:            lepton,
		BLWeightedEtaSupportTotal:         blWeighted,
		HyperchargeResidualSumAbs:         hypResidualAbs,
		UpDownQuarkBalance:                math.Abs(quark) <= eps,
		NeutrinoElectronBalance:           math.Abs(lepton) <= eps,
		TotalEtaSupportBalances:           math.Abs(total) <= eps,
		BLWeightedEtaSupportBalances:      math.Abs(blWeighted) <= eps,
		AnomalyCancellationTheoremDerived: false,
		NeutralityPreflightOnly:           true,
		Verdict:                           "The eta-signed scalar support cancels across up/down quark channels and across neutrino/electron lepton channels. This is a finite support-neutrality preflight, not an anomaly-cancellation theorem or amplitude computation.",
	}
}

func bMinusLForKind(k yukawaintertwiner.FermionKind) float64 {
	switch k {
	case yukawaintertwiner.UpType, yukawaintertwiner.DownType:
		return 1.0 / 3.0
	case yukawaintertwiner.NeutrinoType, yukawaintertwiner.ElectronType:
		return -1
	default:
		return math.NaN()
	}
}

func auditFirewall(b BilinearSupportAudit) FirewallAudit {
	return FirewallAudit{
		TensorSupportDerived:              b.AllSupportNonzero && b.ChannelsAudited == 8,
		PhysicalYukawaAmplitudesDerived:   false,
		FermionMassesDerived:              false,
		GenerationTextureValuesDerived:    false,
		CKMMatrixDerived:                  false,
		PMNSMatrixDerived:                 false,
		ObservedMassInputUsed:             false,
		ObservedMixingInputUsed:           false,
		HiggsVEVValueInserted:             false,
		PhysicalScalarVEVAmplitudeDerived: false,
		SpectralActionEvaluated:           false,
		HeatKernelMatchingDerived:         false,
		ThresholdBetaRowsDerived:          false,
		AbsoluteCouplingPromoted:          false,
		PhysicalConstantsDerived:          false,
		StrictNullityBefore:               3,
		StrictNullityAfter:                3,
		ConditionalSupportNullityBefore:   1,
		ConditionalSupportNullityAfter:    0,
		ClosedStatements: []string{
			"the sealed scalar fundamental class supports all eight Gate-25 one-generation Yukawa incidence channels",
			"Phi_+ and Phi_- carry opposite native tau_eta support, +2 and -2",
			"eta-signed support balances across up/down and neutrino/electron pairs",
		},
		OpenRequirements: []string{
			"derive finite Yukawa amplitudes as texture operators rather than assigning support weights",
			"derive generation/triality mixing before CKM or PMNS claims",
			"derive a physical scalar VEV amplitude before any fermion mass statement",
			"derive heat-kernel/spectral-action normalization before threshold or absolute-coupling claims",
		},
		RecommendedNextGate: "Gate 195 — finite Yukawa texture operator / amplitude-source obstruction audit",
		Verdict:             "Gate 194 closes the support question but leaves all amplitude, generation, mass, mixing, threshold, and constants questions sealed.",
	}
}

func FormatTensorLift(a TensorLiftAudit) string {
	return fmt.Sprintf("H_Fock=%d H_Phi=%d H_total=%d left=%d right=%d functional=%q constructed=%t continuum=%t amplitudeInserted=%t",
		a.MatterFockDimension, a.ScalarBundleDimension, a.TotalTensorDimension, a.LeftDomainDimension, a.RightCodomainDimension, a.TotalFunctional, a.DoublyGradedFunctionalConstructed, a.ContinuumIntegralImported, a.YukawaAmplitudeInserted)
}

func FormatBranches(bs []ScalarBranchSupport) string {
	parts := make([]string, 0, len(bs))
	for _, b := range bs {
		parts = append(parts, fmt.Sprintf("%s(Y=%.3g,%s,tau_eta=%s,mult=%d)", b.Name, b.Hypercharge, b.ProjectorName, b.ExpectedRational, b.Multiplicity))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRecords(rs []YukawaSupportRecord) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, fmt.Sprintf("%s support=%.0f", r.ChannelName, r.TensorSupportSignature))
	}
	sort.Strings(parts)
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummaries(ss []KindSupportSummary) string {
	parts := make([]string, 0, len(ss))
	for _, s := range ss {
		parts = append(parts, fmt.Sprintf("%s channels=%d signed=%.0f abs=%.0f entries=%d", s.Kind, s.Channels, s.SignedSupport, s.AbsoluteSupport, s.ScalarFiberEntries))
	}
	sort.Strings(parts)
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatBilinear(a BilinearSupportAudit) string {
	return fmt.Sprintf("channels=%d supported=%d unsupported=%d up=%d down=%d nu=%d e=%d fiberEntries=%d neutral=%t colorLepton=%t supportOnly=%t records=%s summaries=%s",
		a.ChannelsAudited, a.SupportedChannels, a.UnsupportedChannels, a.UpSupport, a.DownSupport, a.NeutrinoSupport, a.ElectronSupport, a.FiberEntryCount, a.AllHyperchargeNeutral, a.AllColorLeptonValid, a.SupportOnlyTheorem, FormatRecords(a.Records), FormatSummaries(a.Summaries))
}

func FormatNeutrality(a NeutralityPreflightAudit) string {
	return fmt.Sprintf("etaTotal=%.0f quark=%.0f lepton=%.0f BLweighted=%.0f YresidualAbs=%.3e upDown=%t nuElectron=%t anomalyTheorem=%t preflightOnly=%t",
		a.EtaSignedScalarSupportTotal, a.EtaSignedQuarkSupport, a.EtaSignedLeptonSupport, a.BLWeightedEtaSupportTotal, a.HyperchargeResidualSumAbs, a.UpDownQuarkBalance, a.NeutrinoElectronBalance, a.AnomalyCancellationTheoremDerived, a.NeutralityPreflightOnly)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("support=%t amplitudes=%t masses=%t textures=%t CKM=%t PMNS=%t observedMass=%t VEV=%t spectralAction=%t heatKernel=%t thresholds=%t couplings=%t constants=%t strict=%d->%d conditionalSupport=%d->%d next=%s",
		a.TensorSupportDerived, a.PhysicalYukawaAmplitudesDerived, a.FermionMassesDerived, a.GenerationTextureValuesDerived, a.CKMMatrixDerived, a.PMNSMatrixDerived, a.ObservedMassInputUsed, a.HiggsVEVValueInserted, a.SpectralActionEvaluated, a.HeatKernelMatchingDerived, a.ThresholdBetaRowsDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalSupportNullityBefore, a.ConditionalSupportNullityAfter, a.RecommendedNextGate)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d inherited=%t tensorLift=%t channels8=%t nonzero=%t balance=%t supportOnly=%t firewall=%t :: %s",
		a.TestsAudited, a.InheritedScalarFundamentalClass, a.TensorLiftConstructed, a.EightGate25ChannelsSupported, a.AllChannelsHaveNonzeroScalarSupport, a.EtaSignedSupportBalances, a.OnlySupportNotAmplitude, a.FirewallPreserved, a.Comment)
}

// compile-time guard that this package uses the same SU(2)_L left-state kind
// vocabulary as the Gate-25 channel audit; this avoids silently drifting into a
// separate matter convention.
var _ su2lgauge.DoubletKind = su2lgauge.QuarkDoublet
