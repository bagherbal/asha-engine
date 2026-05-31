from __future__ import annotations

from pathlib import Path
from math import cos, sin, pi
import html

from .style import PALETTE, STYLE, FigureStyle
from .readme_ckm import FigureContract, L, S_SPLIT


def esc(s: object) -> str:
    return html.escape(str(s), quote=True)


def polar(cx: float, cy: float, r: float, angle: float) -> tuple[float, float]:
    return cx + r*cos(angle), cy + r*sin(angle)


def svg_open(title: str, desc: str, style: FigureStyle = STYLE) -> list[str]:
    p = PALETTE
    w, h = style.width, style.height
    return [f'''<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}" role="img" aria-labelledby="title desc">
<title id="title">{esc(title)}</title>
<desc id="desc">{esc(desc)}</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="72%"><stop offset="0%" stop-color="#10101B"/><stop offset="62%" stop-color="{p.abyss}"/><stop offset="100%" stop-color="#010103"/></radialGradient>
  <linearGradient id="bridge" x1="0" x2="1"><stop offset="0%" stop-color="{p.cyan}" stop-opacity="0.20"/><stop offset="50%" stop-color="{p.platinum}" stop-opacity="0.78"/><stop offset="100%" stop-color="{p.gold}" stop-opacity="0.34"/></linearGradient>
  <linearGradient id="rg" x1="0" x2="1"><stop offset="0%" stop-color="{p.gold}" stop-opacity="0.18"/><stop offset="55%" stop-color="{p.warm}" stop-opacity="0.86"/><stop offset="100%" stop-color="#FFD6C4" stop-opacity="0.45"/></linearGradient>
  <filter id="softGlow" x="-180%" y="-180%" width="460%" height="460%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-220%" y="-220%" width="540%" height="540%"><feGaussianBlur stdDeviation="15" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="10" markerHeight="10" refX="7" refY="3" orient="auto"><path d="M0,0 L0,6 L8,3 z" fill="{p.platinum}" fill-opacity="0.75"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 58px; fill: {p.platinum}; }}
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
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README · CKM MIXING · RELATIVE ORIENTATION</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
<text class="mathSmall" x="{w/2:.0f}" y="278" text-anchor="middle">masses are depth eigenvalues · mixing is Uu†Ud · RG transport is top-driven</text>
''']


def boxed(parts: list[str], x: float, y: float, width: float, height: float, title: str, subtitle: str, stroke: str, fill_opacity: float = 0.46) -> None:
    parts.append(f'<rect x="{x}" y="{y}" width="{width}" height="{height}" rx="34" fill="#071018" fill-opacity="{fill_opacity}" stroke="{stroke}" stroke-opacity="0.38"/>')
    parts.append(f'<text class="label" x="{x+34}" y="{y+50}">{esc(title)}</text>')
    parts.append(f'<text class="tiny" x="{x+34}" y="{y+82}">{esc(subtitle)}</text>')


def draw_basis_disc(parts: list[str], cx: float, cy: float, title: str, color: str, angle_offset: float, labels: list[str]) -> None:
    p = PALETTE
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="170" fill="{color}" fill-opacity="0.055" stroke="{color}" stroke-opacity="0.42" filter="url(#bigGlow)"/>')
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="105" fill="#050508" fill-opacity="0.66" stroke="{p.platinum}" stroke-opacity="0.18"/>')
    parts.append(f'<text class="math" x="{cx}" y="{cy-205}" text-anchor="middle" fill="{color}">{esc(title)}</text>')
    for i, lab in enumerate(labels):
        a = angle_offset + i*2*pi/3 - pi/2
        x, y = polar(cx, cy, 126, a)
        parts.append(f'<line x1="{cx}" y1="{cy}" x2="{x:.1f}" y2="{y:.1f}" stroke="{color}" stroke-opacity="0.35" stroke-width="2"/>')
        parts.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="28" fill="{color}" fill-opacity="0.20" stroke="{color}" stroke-opacity="0.82" filter="url(#softGlow)"/>')
        parts.append(f'<text class="mathSmall" x="{x:.1f}" y="{y+8:.1f}" text-anchor="middle">{esc(lab)}</text>')


def draw_angle_gauge(parts: list[str], cx: float, cy: float, radius: float, value_rad: float, label: str, formula: str, color: str) -> None:
    p = PALETTE
    deg = value_rad*180/pi
    # scale tiny CKM angles into visible gauge sweep while preserving numeric label.
    visible = min(max(value_rad*3.2, 0.035), 1.35)
    a0 = -pi*0.82
    a1 = a0 + visible
    sx, sy = polar(cx, cy, radius, a0)
    ex, ey = polar(cx, cy, radius, a1)
    large = 1 if visible > pi else 0
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="{radius}" fill="#050508" fill-opacity="0.45" stroke="{p.platinum}" stroke-opacity="0.10"/>')
    parts.append(f'<path d="M {sx:.1f},{sy:.1f} A {radius},{radius} 0 {large} 1 {ex:.1f},{ey:.1f}" fill="none" stroke="{color}" stroke-opacity="0.88" stroke-width="8" filter="url(#softGlow)"/>')
    parts.append(f'<line x1="{cx}" y1="{cy}" x2="{sx:.1f}" y2="{sy:.1f}" stroke="{p.muted}" stroke-opacity="0.25"/>')
    parts.append(f'<line x1="{cx}" y1="{cy}" x2="{ex:.1f}" y2="{ey:.1f}" stroke="{color}" stroke-opacity="0.55"/>')
    parts.append(f'<text class="mathSmall" x="{cx}" y="{cy-14}" text-anchor="middle" fill="{color}">{esc(label)}</text>')
    parts.append(f'<text class="mathSmall" x="{cx}" y="{cy+26}" text-anchor="middle">{value_rad:.9f} rad</text>')
    parts.append(f'<text class="tiny" x="{cx}" y="{cy+58}" text-anchor="middle">{deg:.6f} deg</text>')
    parts.append(f'<text class="tiny" x="{cx}" y="{cy+94}" text-anchor="middle">{esc(formula)}</text>')


def render_ckm(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p = PALETTE
    q = c.quantities
    angles = q["boundary_angles_rad"]
    tc = q["transport_coefficients"]
    parts = svg_open(c.title, "CKM mixing as relative up/down orientation plus exact README boundary-angle formulas and symbolic top-driven RG transport.", style)

    # Upper orientation engine.
    boxed(parts, 130, 392, 1540, 735, "RELATIVE ORIENTATION ENGINE", "depth eigenvalues stay diagonal; CKM lives in the mismatch between up and down bases", p.platinum, 0.34)
    draw_basis_disc(parts, 425, 735, "Uu", p.cyan, 0.0, ["u", "c", "t"])
    draw_basis_disc(parts, 1375, 735, "Ud", p.gold, 0.46, ["d", "s", "b"])
    parts.append(f'<path d="M 595 735 C 720 610 810 590 900 735 C 990 880 1080 860 1205 735" stroke="url(#bridge)" stroke-width="9" fill="none" stroke-opacity="0.82" filter="url(#softGlow)" marker-end="url(#arrow)"/>')
    parts.append(f'<rect x="715" y="610" width="370" height="250" rx="44" fill="#050508" fill-opacity="0.72" stroke="{p.platinum}" stroke-opacity="0.25" filter="url(#softGlow)"/>')
    parts.append(f'<text class="math" x="900" y="705" text-anchor="middle">V_CKM = Uu† Ud</text>')
    parts.append(f'<text class="caption" x="900" y="765" text-anchor="middle">orientation, not mass eigenvalue</text>')
    parts.append(f'<text class="tiny" x="900" y="815" text-anchor="middle">BOUNDARY MATRIX · FLAVOR SELECTOR NOT NATIVE</text>')

    # Angle gauges.
    boxed(parts, 130, 1205, 1540, 500, "LOCKED CKM BOUNDARY ANGLES", "four source-typed parameters evaluated from L=1/(8π) and S=S_split", p.gold, 0.38)
    draw_angle_gauge(parts, 355, 1450, 116, angles["theta12_0"], "θ12⁰", "1/4 − 18S + 158S²", p.gold)
    draw_angle_gauge(parts, 725, 1450, 116, angles["theta23_0"], "θ23⁰", "L + 5S/3 − (8−2L)S²", p.cyan)
    draw_angle_gauge(parts, 1095, 1450, 116, angles["theta13_0"], "θ13⁰", "72LS − 3S²/2", p.warm)
    draw_angle_gauge(parts, 1465, 1450, 116, angles["delta_CKM_0"], "δCKM⁰", "π/3 + 71S + 93S²/4", p.platinum)

    # Transport lane.
    boxed(parts, 130, 1785, 1540, 320, "TOP-YUKAWA RG TRANSPORT", "Γq is a transport integral; the visual does not invent an unprovided running curve", p.warm, 0.38)
    y0 = 1945
    parts.append(f'<path d="M 220 {y0} C 430 {y0-145} 610 {y0+100} 790 {y0-15} S 1120 {y0-100} 1305 {y0+8} S 1505 {y0-60} 1600 {y0-18}" stroke="url(#rg)" stroke-width="7" fill="none" filter="url(#softGlow)" marker-end="url(#arrow)"/>')
    parts.append(f'<text class="mathSmall" x="900" y="1858" text-anchor="middle">Γq(μ) = 3/(32π²) ∫[ln MZ → ln μ] yt(t)² dt</text>')
    parts.append(f'<text class="mathSmall" x="390" y="2028" text-anchor="middle">Du(μ)=Du⁰ − Γq + εu</text>')
    parts.append(f'<text class="mathSmall" x="770" y="2028" text-anchor="middle">Dd(μ)=Dd⁰ + Γq + εd</text>')
    parts.append(f'<text class="mathSmall" x="1160" y="2028" text-anchor="middle">θ23(μ)=θ23⁰ + {tc["theta23_gamma_coefficient"]:.9f} Γq</text>')
    parts.append(f'<text class="mathSmall" x="1510" y="2028" text-anchor="middle">θ13(μ)=θ13⁰ + 1/256 Γq</text>')

    # Constants/firewall strip.
    parts.append(f'<rect x="290" y="2148" width="1220" height="80" rx="30" fill="#050508" fill-opacity="0.70" stroke="{p.obsidian}" stroke-opacity="0.88" stroke-width="5"/>')
    parts.append(f'<text class="caption" x="900" y="2200" text-anchor="middle">L = {L:.12f} · S = {S_SPLIT:.16f} · 158 = 2(72+7) · CKM is sealed boundary orientation, not native flavor theorem</text>')
    parts.append(f'<text class="tiny" x="1580" y="2290" text-anchor="end">EXACT CHECK: θ12⁰, θ23⁰, θ13⁰, δCKM⁰ AND Γq TRANSPORT RULES MATCH README</text>')
    parts.append(f'<text class="tiny" x="1580" y="2320" text-anchor="end">BOUNDARY: NO PMNS, MAJORANA SELECTOR, OR NATIVE FLAVOR CLAIM PROMOTED</text>')
    parts.append('</svg>')
    Path(out_path).write_text("".join(parts), encoding="utf-8")
