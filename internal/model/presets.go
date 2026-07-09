package model

// VisualPolishComponents returns the managed theme/logo components that make up
// the visual polish option across install and uninstall flows.
func VisualPolishComponents() []ComponentID {
	return []ComponentID{ComponentTheme, ComponentClaudeTheme, ComponentOpenCodeGentleLogo}
}

// ComponentsForPreset returns the managed components implied by a preset/persona
// pair. PersonaCustom opts out of managed persona only; preset choice still
// controls visual polish.
func ComponentsForPreset(preset PresetID, persona PersonaID) []ComponentID {
	var components []ComponentID
	switch preset {
	case PresetMinimal:
		components = []ComponentID{ComponentEngram}
	case PresetEcosystemOnly:
		components = []ComponentID{ComponentEngram, ComponentSDD, ComponentSkills, ComponentContext7, ComponentGGA}
	case PresetCustom:
		return nil
	default: // full-gentleman
		components = []ComponentID{
			ComponentEngram,
			ComponentSDD,
			ComponentSkills,
			ComponentContext7,
			ComponentPermission,
			ComponentGGA,
		}
		components = append(components, VisualPolishComponents()...)
	}
	if persona != PersonaCustom {
		components = append(components, ComponentPersona)
	}
	return components
}
