from __future__ import annotations

from pathlib import Path
from math import log10, pi
import html

from .style import PALETTE, STYLE, FigureStyle
from .readme_neutrinos import FigureContract, L, S_SPLIT


def esc(s: object) -> str:
    return html.escape(str(s), quote=True)


def svg_open(title: str, desc: str, style: FigureStyle = STYLE) -> list[str]:
    p = PALETTE
    w, h = style.width, style.height
    return [f'''<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}" role="img" aria-labelledby="title desc">
<title id="title">{esc(title)}</title>
<desc id="desc">{esc(desc)}</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="72%"><stop offset="0%" stop-color="#10101B"/><stop offset="62%" stop-color="{p.abyss}"/><stop offset="100%" stop-color="#010103"/></radialGradient>
  <linearGradient id="seesaw" x1="0" x2="1"><stop offset="0%" stop-color="{p.cyan}" stop-opacity="0.28"/><stop offset="50%" stop-color="{p.platinum}" stop-opacity="0.86"/><stop offset="100%" stop-color="{p.gold}" stop-opacity="0.42"/></linearGradient>
  <linearGradient id="fire" x1="0" x2="1"><stop offset="0%" stop-color="#050508"/><stop offset="50%" stop-color="#181820"/><stop offset="100%" stop-color="#050508"/></linearGradient>
  <filter id="softGlow" x="-180%" y="-180%" width="460%" height="460%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-220%" y="-220%" width="540%" height="540%"><feGaussianBlur stdDeviation="15" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="10" markerHeight="10" refX="7" refY="3" orient="auto"><path d="M0,0 L0,6 L8,3 z" fill="{p.platinum}" fill-opacity="0.72"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 60px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 34px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 25px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 2.5px; fill: {p.muted}; }}
    .tiny {{ font-family: {style.label_font}; font-size: 14px; letter-spacing: 1.35px; fill: #8B93A5; }}
    .caption {{ font-family: {style.label_font}; font-size: 19px; fill: #AAB3C2; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{w-96}" height="{h-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README · NEUTRINO SECTOR · SEALED RANK-2 LANE</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
<text class="mathSmall" x="{w/2:.0f}" y="278" text-anchor="middle">Type-I seesaw · heavy Majorana bridge · PMNS selector unresolved</text>
''']


def boxed(parts: list[str], x: float, y: float, width: float, height: float, title: str, subtitle: str, stroke: str, fill_opacity: float = 0.44) -> None:
    parts.append(f'<rect x="{x}" y="{y}" width="{width}" height="{height}" rx="34" fill="#071018" fill-opacity="{fill_opacity}" stroke="{stroke}" stroke-opacity="0.38"/>')
    parts.append(f'<text class="label" x="{x+34}" y="{y+50}">{esc(title)}</text>')
    parts.append(f'<text class="tiny" x="{x+34}" y="{y+82}">{esc(subtitle)}</text>')


def mass_node(parts: list[str], x: float, y: float, label: str, value: str, color: str, r: float = 46, dashed: bool = False) -> None:
    dash = ' stroke-dasharray="9 9"' if dashed else ''
    parts.append(f'<circle cx="{x}" cy="{y}" r="{r}" fill="{color}" fill-opacity="0.13" stroke="{color}" stroke-opacity="0.82"{dash} filter="url(#softGlow)"/>')
    parts.append(f'<text class="mathSmall" x="{x}" y="{y-8}" text-anchor="middle" fill="{color}">{esc(label)}</text>')
    parts.append(f'<text class="tiny" x="{x}" y="{y+28}" text-anchor="middle">{esc(value)}</text>')


def draw_scale_axis(parts: list[str], x: float, y0: float, y1: float, vals: list[tuple[str, float, str]], color: str) -> None:
    logs = [log10(v) for _, v, _ in vals]
    mn, mx = min(logs), max(logs)
    parts.append(f'<line x1="{x}" y1="{y0}" x2="{x}" y2="{y1}" stroke="{color}" stroke-opacity="0.42" stroke-width="4"/>')
    for lab, val, note in vals:
        t = (log10(val)-mn)/(mx-mn)
        y = y1 - t*(y1-y0)
        parts.append(f'<circle cx="{x}" cy="{y:.1f}" r="18" fill="{color}" fill-opacity="0.24" stroke="{color}" stroke-opacity="0.88" filter="url(#softGlow)"/>')
        parts.append(f'<text class="mathSmall" x="{x+42}" y="{y+7:.1f}" fill="{color}">{esc(lab)}</text>')
        parts.append(f'<text class="tiny" x="{x+42}" y="{y+38:.1f}">{val:.4g} GeV · {esc(note)}</text>')


def angle_chip(parts: list[str], x: float, y: float, label: str, formula: str, value: str, color: str) -> None:
    parts.append(f'<rect x="{x}" y="{y}" width="360" height="118" rx="28" fill="#050508" fill-opacity="0.62" stroke="{color}" stroke-opacity="0.35"/>')
    parts.append(f'<text class="mathSmall" x="{x+180}" y="{y+42}" text-anchor="middle" fill="{color}">{esc(label)}</text>')
    parts.append(f'<text class="tiny" x="{x+180}" y="{y+74}" text-anchor="middle">{esc(formula)}</text>')
    parts.append(f'<text class="tiny" x="{x+180}" y="{y+99}" text-anchor="middle">{esc(value)}</text>')


def render_neutrinos(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p = PALETTE
    q = c.quantities
    lane = q["rank2_normal_order_lane"]
    hs = q["heavy_scale_bridge"]
    pmns = q["pmns_skeleton_rad"]
    parts = svg_open(c.title, "Neutrino Type-I seesaw, rank-2 normal-order bridge, heavy Majorana scale, and PMNS firewall.", style)

    # Top theorem/bridge pipeline.
    boxed(parts, 130, 395, 1540, 520, "TYPE-I SEESAW CORE", "theorem skeleton is allowed; filling choices remain typed", p.platinum, 0.34)
    parts.append(f'<circle cx="320" cy="660" r="92" fill="{p.cyan}" fill-opacity="0.10" stroke="{p.cyan}" stroke-opacity="0.55" filter="url(#bigGlow)"/>')
    parts.append(f'<text class="math" x="320" y="652" text-anchor="middle" fill="{p.cyan}">Y_nu</text>')
    parts.append(f'<text class="tiny" x="320" y="700" text-anchor="middle">Dirac lane</text>')
    parts.append(f'<path d="M 420 660 C 570 540 670 540 820 660" stroke="url(#seesaw)" stroke-width="8" fill="none" marker-end="url(#arrow)" filter="url(#softGlow)"/>')
    parts.append(f'<rect x="700" y="540" width="400" height="240" rx="42" fill="#050508" fill-opacity="0.72" stroke="{p.gold}" stroke-opacity="0.32" filter="url(#softGlow)"/>')
    parts.append(f'<text class="math" x="900" y="625" text-anchor="middle">M_R^-1</text>')
    parts.append(f'<text class="tiny" x="900" y="672" text-anchor="middle">heavy Majorana inverse</text>')
    parts.append(f'<text class="tiny" x="900" y="716" text-anchor="middle">Planck-stiffness bridge, not native MP theorem</text>')
    parts.append(f'<path d="M 1100 660 C 1240 795 1360 795 1500 660" stroke="url(#seesaw)" stroke-width="8" fill="none" marker-end="url(#arrow)" filter="url(#softGlow)"/>')
    parts.append(f'<circle cx="1510" cy="660" r="92" fill="{p.platinum}" fill-opacity="0.08" stroke="{p.platinum}" stroke-opacity="0.58" filter="url(#bigGlow)"/>')
    parts.append(f'<text class="math" x="1510" y="647" text-anchor="middle">M_nu</text>')
    parts.append(f'<text class="tiny" x="1510" y="698" text-anchor="middle">light effective matrix</text>')
    parts.append(f'<text class="mathSmall" x="900" y="850" text-anchor="middle">M_nu = -(v^2/2) Y_nu^T M_R^-1 Y_nu</text>')

    # Rank-2 light spectrum.
    boxed(parts, 130, 1010, 720, 520, "RANK-2 NORMAL-ORDER LANE", "m1 is dark; m2/m3 is locked by 4L+10S", p.cyan, 0.42)
    mass_node(parts, 290, 1270, "m1", "~ 0", p.obsidian, 52, True)
    mass_node(parts, 505, 1270, "m2", f'{lane["m2_eV"]:.6f} eV', p.cyan, 48)
    mass_node(parts, 720, 1270, "m3", f'{lane["m3_eV"]:.6f} eV', p.gold, 58)
    parts.append(f'<path d="M 555 1270 C 610 1205 650 1205 680 1270" stroke="{p.gold}" stroke-opacity="0.58" stroke-width="4" fill="none" marker-end="url(#arrow)"/>')
    parts.append(f'<text class="mathSmall" x="490" y="1410" text-anchor="middle">m2 = (4L + 10S) m3</text>')
    parts.append(f'<text class="mathSmall" x="490" y="1460" text-anchor="middle" fill="{p.gold}">m2/m3 = {lane["m2_over_m3"]:.12f}</text>')

    # Heavy scales.
    boxed(parts, 950, 1010, 720, 520, "HEAVY SCALE BRIDGE", "MR3 and MR2 are scale fillings, not native derivations", p.gold, 0.42)
    draw_scale_axis(parts, 1055, 1165, 1450, [("MR3", hs["M_R3_GeV"], "sqrt(v MP)"), ("MR2", hs["M_R2_GeV"], "depth suppressed")], p.gold)
    parts.append(f'<text class="tiny" x="1335" y="1165" text-anchor="middle">M_R3 = (sqrt(2pi)+49S+90S^2) sqrt(v M_P)</text>')
    parts.append(f'<text class="tiny" x="1335" y="1260" text-anchor="middle">M_R2 = M_R3 exp(-4pi/3)/(4L+10S)</text>')
    parts.append(f'<text class="tiny" x="1335" y="1360" text-anchor="middle">m3 = v^2 exp(-2A_tau)/(2M_R3)</text>')
    parts.append(f'<text class="tiny" x="1335" y="1435" text-anchor="middle">M_R3 ~= {hs["M_R3_GeV"]:.4g} GeV · M_R2 ~= {hs["M_R2_GeV"]:.4g} GeV</text>')

    # PMNS firewall.
    boxed(parts, 130, 1625, 1540, 430, "PMNS SKELETON BEHIND FIREWALL", "angle skeleton exists; CP and Majorana orientation selector remain unresolved", p.warm, 0.38)
    parts.append(f'<rect x="190" y="1725" width="1420" height="110" rx="38" fill="url(#fire)" stroke="{p.obsidian}" stroke-opacity="0.95" stroke-width="10"/>')
    parts.append(f'<text class="caption" x="900" y="1792" text-anchor="middle">OBSIDIAN FIREWALL: PMNS is not locked · Majorana phases unresolved · no native orientation selector</text>')
    angle_chip(parts, 210, 1885, "th23 PMNS", "pi/4 +/- 48S", f'{pmns["theta23_minus"]*180/pi:.3f}° / {pmns["theta23_plus"]*180/pi:.3f}°', p.warm)
    angle_chip(parts, 610, 1885, "th12 PMNS", "pi/6 + 48S", f'{pmns["theta12"]*180/pi:.3f}°', p.cyan)
    angle_chip(parts, 1010, 1885, "th13 PMNS", "4L - 7S + 4S^2", f'{pmns["theta13"]*180/pi:.3f}°', p.gold)
    parts.append(f'<circle cx="1510" cy="1944" r="62" fill="{p.obsidian}" fill-opacity="0.86" stroke="{p.warm}" stroke-opacity="0.55" stroke-dasharray="8 9" filter="url(#softGlow)"/>')
    parts.append(f'<text class="mathSmall" x="1510" y="1937" text-anchor="middle" fill="{p.warm}">αM</text>')
    parts.append(f'<text class="tiny" x="1510" y="1974" text-anchor="middle">unresolved</text>')

    parts.append(f'<rect x="300" y="2135" width="1200" height="92" rx="34" fill="#050508" fill-opacity="0.72" stroke="{p.platinum}" stroke-opacity="0.20"/>')
    parts.append(f'<text class="caption" x="900" y="2188" text-anchor="middle">L = {L:.12f} · S = {S_SPLIT:.16f} · rank-2 lane is sealed diagnostic/filling, not native flavor theorem</text>')
    parts.append(f'<text class="tiny" x="1580" y="2290" text-anchor="end">EXACT CHECK: SEESAW IDENTITY, m2/m3, MR3, MR2, m3 AND PMNS SKELETON MATCH README</text>')
    parts.append(f'<text class="tiny" x="1580" y="2320" text-anchor="end">BOUNDARY: NO PMNS LOCK, NO MAJORANA PHASE DERIVATION, NO NATIVE NEUTRINO SELECTOR</text>')
    parts.append('</svg>')
    Path(out_path).write_text("".join(parts), encoding="utf-8")
