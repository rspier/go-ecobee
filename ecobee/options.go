package ecobee

type SelectionOption func(s Selection)

func WithIncludeAlerts(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeAlerts = value
	}
}

func WithIncludeAudio(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeAudio = value
	}
}

func WithIncludeDevice(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeDevice = value
	}
}

func WithIncludeElectricity(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeElectricity = value
	}
}

func WithIncludeEquipmentStatus(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeElectricity = value
	}
}

func WithIncludeEvents(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeEvents = value
	}
}

func WithIncludeExtendedRuntime(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeExtendedRuntime = value
	}
}

func WithIncludeHouseDetails(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeHouseDetails = value
	}
}

func WithIncludeLocation(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeLocation = value
	}
}

func WithIncludeManagement(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeManagement = value
	}
}

func WithIncludeNotificationSettings(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeManagement = value
	}
}

func WithIncludeOemCfg(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeOemCfg = value
	}
}

func WithIncludePrivacy(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludePrivacy = value
	}
}

func WithIncludeProgram(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeProgram = value
	}
}

func WithIncludeRuntime(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeRuntime = value
	}
}

func WithIncludeSecuritySettings(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSecuritySettings = value
	}
}

func WithIncludeSensors(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSensors = value
	}
}

func WithIncludeSettings(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSettings = value
	}
}

func WithIncludeTechnician(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeSettings = value
	}
}

func WithIncludeUtility(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeUtility = value
	}
}

func WithIncludeVersion(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeVersion = value
	}
}

func WithIncludeWeather(value bool) SelectionOption {
	return func(s Selection) {
		s.IncludeWeather = value
	}
}
