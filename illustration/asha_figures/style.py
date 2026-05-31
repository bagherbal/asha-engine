from dataclasses import dataclass

@dataclass(frozen=True)
class Palette:
    abyss: str = "#050508"
    abyss_2: str = "#090912"
    platinum: str = "#E8EDF2"
    frost: str = "#A9F8FF"
    cyan: str = "#67E8F9"
    gold: str = "#D9B45F"
    warm: str = "#FF8F70"
    obsidian: str = "#101016"
    muted: str = "#78808E"
    glass: str = "#D7F7FF"

@dataclass(frozen=True)
class FigureStyle:
    width: int = 1800
    height: int = 2400
    margin_x: int = 230
    tier_top: int = 360
    tier_bottom: int = 1980
    node_radius: float = 4.8
    middle_node_radius: float = 4.1
    terminal_node_radius: float = 8.0
    edge_opacity: float = 0.070
    edge_width: float = 0.80
    title_font: str = "STIXGeneral, Noto Serif, DejaVu Serif, Georgia, serif"
    label_font: str = "Inter, Noto Sans, DejaVu Sans, Avenir, Helvetica, Arial, sans-serif"

PALETTE = Palette()
STYLE = FigureStyle()
