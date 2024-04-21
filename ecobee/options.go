package ecobee

type SelectionOption func(s Selection)

func IncludeAlerts(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeAlerts = value
	}
}

func IncludeAudio(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeAudio = value
	}
}

func IncludeDevice(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeDevice = value
	}
}

func IncludeElectricity(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeElectricity = value
	}
}

func IncludeEquipmentStatus(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeElectricity = value
	}
}

func IncludeEvents(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeEvents = value
	}
}

func IncludeExtendedRuntime(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeExtendedRuntime = value
	}
}

func IncludeHouseDetails(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeHouseDetails = value
	}
}

func IncludeLocation(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeLocation = value
	}
}

func IncludeManagement(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeManagement = value
	}
}

func IncludeNotificationSettings(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeManagement = value
	}
}

func IncludeOemCfg(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeOemCfg = value
	}
}

func IncludePrivacy(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludePrivacy = value
	}
}

func IncludeProgram(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeProgram = value
	}
}

func IncludeRuntime(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeRuntime = value
	}
}

func IncludeSecuritySettings(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSecuritySettings = value
	}
}

func IncludeSensors(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSensors = value
	}
}

func IncludeSettings(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSettings = value
	}
}

func IncludeTechnician(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSettings = value
	}
}

func IncludeUtility(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeUtility = value
	}
}

func IncludeVersion(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeVersion = value
	}
}

func IncludeWeather(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeWeather = value
	}
}
