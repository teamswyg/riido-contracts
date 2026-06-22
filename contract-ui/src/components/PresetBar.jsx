import { figmaPresets } from "../figmaPresets.js";

export default function PresetBar({ onSelect }) {
  return (
    <nav className="preset-bar" aria-label="Figma screen presets">
      <span>Figma screen presets</span>
      {figmaPresets.map((preset) => (
        <button key={preset.label} onClick={() => onSelect(preset.query)}>{preset.label}</button>
      ))}
    </nav>
  );
}
