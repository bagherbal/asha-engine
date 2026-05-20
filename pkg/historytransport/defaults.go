package historytransport

func DefaultInputs() InputSet {
	mu0 := 91.1876
	return InputSet{
		TaskName: TaskName,
		Mu0Name:  Mu0Name,
		Mu0GeV:   mu0,
		ASHABoundary: ASHABoundaryLaw{
			KY:                        5.0 / 3.0,
			Sin2ThetaBoundary:         3.0 / 8.0,
			CanonicalBoundaryRelation: "g1(Lambda)=g2(Lambda), with g1^2=(5/3) gY^2",
			FiniteAlgebra:             "A_F = C + H + M_3(C)",
			ScalarCarrier:             "H_phi ~= C^2",
			BridgeOnly:                true,
		},
		Measured: map[string]MeasuredValue{
			"G_F":        {Name: "G_F", Value: 1.1663787e-5, Uncertainty: 0.0000006e-5, Unit: "GeV^-2", SourceID: "NIST_CODATA_2018_GF", Role: "electroweak v extraction", BridgeOnly: true},
			"m_W":        {Name: "m_W", Value: 80.3602, Uncertainty: 0.0099, Unit: "GeV", SourceID: "CMS_2024_W_MASS", Role: "on-shell weak-angle and g2 extraction", BridgeOnly: true},
			"m_Z":        {Name: "m_Z", Value: 91.1876, Uncertainty: 0.0021, Unit: "GeV", SourceID: "PDG_2024_Z_ALPHA", Role: "observation scale and neutral Hessian extraction", BridgeOnly: true},
			"m_H":        {Name: "m_H", Value: 125.38, Uncertainty: 0.14, Unit: "GeV", SourceID: "CMS_2020_HIGGS_MASS", Role: "lambda extraction", BridgeOnly: true},
			"alpha_s_MZ": {Name: "alpha_s(M_Z)", Value: 0.1179, Uncertainty: 0.0010, Unit: "dimensionless", Scale: "M_Z", Scheme: "MS-bar", SourceID: "PDG_2024_Z_ALPHA", Role: "g3 extraction", BridgeOnly: true},
		},
		Fermions: []FermionMassInput{
			{Name: "u", MassGeV: 0.00216, MassUncertainty: 0.00038, InputScaleGeV: 2.0, TargetScaleGeV: mu0, Scheme: "MS-bar reference mass at 2 GeV", Transport: "one-loop QCD mass transport to M_Z; threshold corrections not included", SourceID: "PDG_2024_QUARKS", BridgeOnly: true},
			{Name: "c", MassGeV: 1.27, MassUncertainty: 0.02, InputScaleGeV: 1.27, TargetScaleGeV: mu0, Scheme: "MS-bar mass at mu=m_c", Transport: "one-loop QCD mass transport to M_Z; threshold corrections not included", SourceID: "PDG_2024_QUARKS", BridgeOnly: true},
			{Name: "t", MassGeV: 162.5, MassUncertainty: 1.0, InputScaleGeV: 162.5, TargetScaleGeV: mu0, Scheme: "MS-bar converted top reference", Transport: "one-loop QCD mass transport to M_Z; top threshold convention explicit", SourceID: "ASHA_GATE473_PDG_STYLE_TOP_MSBAR_REFERENCE", BridgeOnly: true},
			{Name: "d", MassGeV: 0.00467, MassUncertainty: 0.00033, InputScaleGeV: 2.0, TargetScaleGeV: mu0, Scheme: "MS-bar reference mass at 2 GeV", Transport: "one-loop QCD mass transport to M_Z; threshold corrections not included", SourceID: "PDG_2024_QUARKS", BridgeOnly: true},
			{Name: "s", MassGeV: 0.0934, MassUncertainty: 0.0060, InputScaleGeV: 2.0, TargetScaleGeV: mu0, Scheme: "MS-bar reference mass at 2 GeV", Transport: "one-loop QCD mass transport to M_Z; threshold corrections not included", SourceID: "PDG_2024_QUARKS", BridgeOnly: true},
			{Name: "b", MassGeV: 4.18, MassUncertainty: 0.03, InputScaleGeV: 4.18, TargetScaleGeV: mu0, Scheme: "MS-bar mass at mu=m_b", Transport: "one-loop QCD mass transport to M_Z; threshold corrections not included", SourceID: "PDG_2024_QUARKS", BridgeOnly: true},
			{Name: "e", MassGeV: 0.000510998950, MassUncertainty: 0, InputScaleGeV: mu0, TargetScaleGeV: mu0, Scheme: "charged-lepton pole mass used as v1 proxy", Transport: "no lepton mass running before Yukawa extraction in v1", SourceID: "PDG_2024_LEPTONS", BridgeOnly: true},
			{Name: "mu", MassGeV: 0.1056583755, MassUncertainty: 0, InputScaleGeV: mu0, TargetScaleGeV: mu0, Scheme: "charged-lepton pole mass used as v1 proxy", Transport: "no lepton mass running before Yukawa extraction in v1", SourceID: "PDG_2024_LEPTONS", BridgeOnly: true},
			{Name: "tau", MassGeV: 1.77686, MassUncertainty: 0.00012, InputScaleGeV: mu0, TargetScaleGeV: mu0, Scheme: "charged-lepton pole mass used as v1 proxy", Transport: "no lepton mass running before Yukawa extraction in v1", SourceID: "PDG_2024_LEPTONS", BridgeOnly: true},
		},
		CKM:       CKMInput{S12: 0.22501, S13: 0.003732, S23: 0.04183, Delta: 1.147, SourceID: "PDG_2024_CKM"},
		Cosmology: CosmologyEndpoint{OmegaCH2: 0.120, OmegaBH2: 0.0224, NS: 0.965, Tau: 0.054, SourceID: "PLANCK_2018_LCDM", BridgeOnly: true},
		Sources: []SourceRef{
			{ID: "CMS_2024_W_MASS", Title: "High-precision measurement of the W boson mass with the CMS experiment at the LHC", URL: "https://arxiv.org/abs/2412.13872", Version: "arXiv:2412.13872", Note: "m_W = 80.3602 +/- 0.0099 GeV used as selectable W endpoint"},
			{ID: "NIST_CODATA_2018_GF", Title: "CODATA value: Fermi coupling constant", URL: "https://physics.nist.gov/cgi-bin/cuu/Value?gf=", Version: "CODATA 2018/NIST", Note: "G_F/(hbar c)^3 = 1.1663787(6)e-5 GeV^-2"},
			{ID: "CMS_2020_HIGGS_MASS", Title: "A measurement of the Higgs boson mass in the diphoton decay channel", URL: "https://arxiv.org/abs/2002.06398", Version: "arXiv:2002.06398v2", Note: "combined CMS value m_H = 125.38 +/- 0.14 GeV"},
			{ID: "PLANCK_2018_LCDM", Title: "Planck 2018 results. VI. Cosmological parameters", URL: "https://arxiv.org/abs/1807.06209", Version: "arXiv:1807.06209", Note: "base-LambdaCDM endpoint values omega_c h^2, omega_b h^2, n_s, tau"},
			{ID: "PDG_2024_CKM", Title: "PDG 2024 Review: CKM Quark-Mixing Matrix", URL: "https://pdg.lbl.gov/2024/reviews/rpp2024-rev-ckm-matrix.pdf", Version: "PDG 2024, Phys. Rev. D 110, 030001", Note: "s12, s13, s23, delta and J_CKM reference"},
			{ID: "PDG_2024_QUARKS", Title: "PDG 2024 Summary Table: Quarks", URL: "https://pdg.lbl.gov/2024/tables/rpp2024-sum-quarks.pdf", Version: "PDG 2024", Note: "reference quark masses and schemes; v1 runs them to M_Z with one-loop QCD and no threshold matching"},
			{ID: "PDG_2024_Z_ALPHA", Title: "PDG 2024 physical constants / Z and alpha_s reference", URL: "https://pdg.lbl.gov/2024/reviews/rpp2024-rev-phys-constants.pdf", Version: "PDG 2024", Note: "m_Z and alpha_s(M_Z) endpoint values"},
			{ID: "PDG_2024_LEPTONS", Title: "PDG 2024 lepton mass references", URL: "https://pdg.lbl.gov/2024/", Version: "PDG 2024", Note: "charged-lepton masses used as v1 pole-mass proxies"},
			{ID: "ASHA_GATE473_PDG_STYLE_TOP_MSBAR_REFERENCE", Title: "ASHA Gate473 raw quark mass hierarchy fixture", URL: "data/pdg_raw_quark_masses_gate473.json", Version: "Gate473 bridge ledger", Note: "top MS-bar reference retained as bridge-only project input"},
		},
		Warnings: []string{
			"Observed inputs are bridge-only endpoints, not ASHA-native derivations.",
			"The g1=g2 crossing is a one-loop Standard Model boundary-normalization test, not full gauge unification.",
			"Quark masses are transported to M_Z with a v1 one-loop QCD approximation; production use requires multi-loop running and threshold matching.",
			"Scalar and flavor upward transport are v1 one-loop approximations; full precision vacuum-stability and flavor transport are not claimed.",
			"Planck LambdaCDM values are optional cosmology endpoints and remain sealed residual data.",
		},
	}
}
