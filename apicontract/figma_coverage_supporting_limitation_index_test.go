package apicontract

type figmaSupportingToolLimitationLookup struct {
	metadataPageList          figmaSupportingToolLimitation
	headlessFileKey           figmaSupportingToolLimitation
	onboardingPageLoadTimeout figmaSupportingToolLimitation
}

func figmaSupportingToolLimitationIndex(limitations []figmaSupportingToolLimitation) figmaSupportingToolLimitationLookup {
	var out figmaSupportingToolLimitationLookup
	for _, limitation := range limitations {
		switch limitation.ID {
		case "figma-metadata-page-list-underreports-pages.v1":
			out.metadataPageList = limitation
		case "figma-headless-file-key-placeholder.v1":
			out.headlessFileKey = limitation
		case "figma-onboarding-page-load-timeout.v1":
			out.onboardingPageLoadTimeout = limitation
		}
	}
	return out
}
