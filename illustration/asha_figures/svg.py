from __future__ import annotations

import html
from pathlib import Path
from .style import PALETTE, STYLE, FigureStyle
from .boolean_lattice import LatticeNode, LatticeEdge, EXPECTED_TIER_COUNTS


def esc(s: str) -> str:
    return html.escape(s, quote=True)


def path_for_edge(a: LatticeNode, b: LatticeNode) -> str:
    # Slight curvature gives the dense 8-cube Hasse graph an optical crystalline field.
    dy = b.y - a.y
    c1x = a.x
    c1y = a.y + 0.42 * dy
    c2x = b.x
    c2y = b.y - 0.42 * dy
    return f"M {a.x:.2f},{a.y:.2f} C {c1x:.2f},{c1y:.2f} {c2x:.2f},{c2y:.2f} {b.x:.2f},{b.y:.2f}"


def render_section1_svg(nodes: list[LatticeNode], edges: list[LatticeEdge], out_path: str | Path,
                        style: FigureStyle = STYLE) -> None:
    out_path = Path(out_path)
    by_mask = {n.mask: n for n in nodes}
    p = PALETTE
    width, height = style.width, style.height
    tier_y = [style.tier_top + k * (style.tier_bottom - style.tier_top) / 8 for k in range(9)]

    parts: list[str] = []
    parts.append(f'''<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}" role="img" aria-labelledby="title desc">
<title id="title">ASHA Section 1 — Measurement Ladder Cl(1,7)</title>
<desc id="desc">Exact nine-tier exterior-grade Hasse diagram with 256 basis blades in Pascal counts 1, 8, 28, 56, 70, 56, 28, 8, 1. Warm nodes contain the time-like e0 generator; cyan nodes are pure space-like blades.</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="70%">
    <stop offset="0%" stop-color="#10101B"/>
    <stop offset="58%" stop-color="{p.abyss}"/>
    <stop offset="100%" stop-color="#010103"/>
  </radialGradient>
  <radialGradient id="haloGold" cx="50%" cy="50%" r="50%">
    <stop offset="0%" stop-color="#FFE7A3" stop-opacity="0.55"/>
    <stop offset="35%" stop-color="#D9B45F" stop-opacity="0.18"/>
    <stop offset="100%" stop-color="#D9B45F" stop-opacity="0"/>
  </radialGradient>
  <filter id="softGlow" x="-200%" y="-200%" width="500%" height="500%">
    <feGaussianBlur stdDeviation="5" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <filter id="bigGlow" x="-250%" y="-250%" width="600%" height="600%">
    <feGaussianBlur stdDeviation="14" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 66px; letter-spacing: 0.5px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 20px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 38px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 20px; letter-spacing: 3px; fill: {p.muted}; }}
    .tierCount {{ font-family: {style.title_font}; font-size: 34px; fill: {p.gold}; }}
    .caption {{ font-family: {style.label_font}; font-size: 22px; line-height: 1.5; fill: #AAB3C2; }}
    .tiny {{ font-family: {style.label_font}; font-size: 15px; letter-spacing: 1.5px; fill: #8B93A5; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{width-96}" height="{height-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
''')

    # Subtle vertical light shaft and central Λ4 chamber halo.
    parts.append(f'<ellipse cx="{width/2:.1f}" cy="{tier_y[4]:.1f}" rx="730" ry="122" fill="url(#haloGold)" opacity="0.72"/>\n')
    parts.append(f'<line x1="{width/2:.1f}" y1="260" x2="{width/2:.1f}" y2="2060" stroke="{p.platinum}" stroke-opacity="0.055" stroke-width="1"/>\n')

    # Tier guide rails.
    for k, y in enumerate(tier_y):
        op = 0.22 if k == 4 else 0.105
        sw = 1.2 if k == 4 else 0.65
        color = p.gold if k == 4 else p.glass
        parts.append(f'<line x1="145" x2="{width-145}" y1="{y:.2f}" y2="{y:.2f}" stroke="{color}" stroke-opacity="{op}" stroke-width="{sw}"/>\n')

    # Edges, split by warm/cold added generator for signature cue.
    parts.append('<g id="hasse_edges" fill="none" stroke-linecap="round">\n')
    for edge in edges:
        a = by_mask[edge.source]
        b = by_mask[edge.target]
        color = p.warm if edge.warm else p.cyan
        opacity = style.edge_opacity * (1.65 if edge.warm else 1.0)
        parts.append(f'<path d="{path_for_edge(a,b)}" stroke="{color}" stroke-opacity="{opacity:.4f}" stroke-width="{style.edge_width}"/>\n')
    parts.append('</g>\n')

    # Tier labels and counts.
    for k, y in enumerate(tier_y):
        parts.append(f'<text class="label" x="86" y="{y+7:.1f}">Λ^{k}</text>\n')
        parts.append(f'<text class="tierCount" x="{width-112}" y="{y+10:.1f}" text-anchor="end">{EXPECTED_TIER_COUNTS[k]}</text>\n')
    parts.append(f'<text class="tiny" x="{width-112}" y="{tier_y[4]-45:.1f}" text-anchor="end" fill="{p.gold}">MIDDLE CHAMBER</text>\n')
    parts.append(f'<text class="math" x="{width/2:.1f}" y="{tier_y[4]-48:.1f}" text-anchor="middle" fill="{p.gold}">Lambda-4 R8 : 70</text>\n')

    # Nodes with glow. Draw halos first for depth.
    parts.append('<g id="node_halos" filter="url(#bigGlow)">\n')
    for n in nodes:
        if n.grade == 4:
            color = p.gold
            op = 0.30
            rr = n.radius * 2.7
        elif n.has_time:
            color = p.warm
            op = 0.18
            rr = n.radius * 2.15
        else:
            color = p.cyan
            op = 0.13
            rr = n.radius * 1.95
        parts.append(f'<circle cx="{n.x:.2f}" cy="{n.y:.2f}" r="{rr:.2f}" fill="{color}" opacity="{op:.3f}"/>\n')
    parts.append('</g>\n')

    parts.append('<g id="nodes" filter="url(#softGlow)">\n')
    for n in nodes:
        if n.grade == 4:
            fill = "#F7E7B2" if n.has_time else "#D9B45F"
            stroke = p.platinum
            sw = 0.65
        elif n.has_time:
            fill = "#FFB49A"
            stroke = "#FFE0C4"
            sw = 0.7
        else:
            fill = "#8AF7FF"
            stroke = "#E2FDFF"
            sw = 0.55
        parts.append(f'<circle cx="{n.x:.2f}" cy="{n.y:.2f}" r="{n.radius:.2f}" fill="{fill}" stroke="{stroke}" stroke-opacity="0.82" stroke-width="{sw}"/>\n')
    parts.append('</g>\n')

    # Grade-1 labels, because they encode the Cl(1,7) signature directly.
    grade1 = [n for n in nodes if n.grade == 1]
    for n in grade1:
        color = p.warm if n.has_time else p.cyan
        parts.append(f'<text class="tiny" x="{n.x:.2f}" y="{n.y-22:.2f}" text-anchor="middle" fill="{color}">{esc(n.basis)}</text>\n')

    # Header / footer annotation.
    parts.append(f'''<text class="subtitle" x="{width/2:.1f}" y="132" text-anchor="middle">ASHA MANUSCRIPT · SECTION 1 · FINITE MEASUREMENT LANGUAGE</text>
<text class="title" x="{width/2:.1f}" y="214" text-anchor="middle">Measurement Ladder of Cl(1,7)</text>
<text class="math" x="{width/2:.1f}" y="280" text-anchor="middle">1, 8, 28, 56, 70, 56, 28, 8, 1</text>
<g transform="translate(220 2120)">
  <circle cx="0" cy="0" r="7" fill="{p.warm}" filter="url(#softGlow)"/><text class="caption" x="26" y="8">time-like participation: blades containing e0, metric +1</text>
  <circle cx="0" cy="48" r="7" fill="{p.cyan}" filter="url(#softGlow)"/><text class="caption" x="26" y="56">space-like-only blades from e1...e7, metric -1</text>
  <circle cx="0" cy="96" r="7" fill="{p.gold}" filter="url(#softGlow)"/><text class="caption" x="26" y="104">Lambda-4 R8 middle chamber: 70-dimensional arena for later Boolean/G2 comparison</text>
</g>
<text class="tiny" x="{width-220}" y="2240" text-anchor="end">EXACT CHECK: 256 NODES · 1024 COVER EDGES · 9 HORIZONTAL TIERS</text>
<text class="tiny" x="{width-220}" y="2270" text-anchor="end">BOUNDARY: MEASUREMENT BOOKKEEPING, NOT PARTICLE NUMEROLOGY</text>
</svg>''')
    out_path.write_text("".join(parts), encoding="utf-8")
