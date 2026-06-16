package progressmessage

import "strings"

func NormalizeArgsForCode(code int, args map[string]string) map[string]string {
	if len(args) == 0 {
		return args
	}
	label := strings.TrimSpace(args["label"])
	if label == "" {
		return args
	}
	normalized := NormalizeLabelForCode(code, label)
	if normalized == label {
		return args
	}
	out := make(map[string]string, len(args))
	for key, value := range args {
		out[key] = value
	}
	out["label"] = normalized
	return out
}

func NormalizeLabelForCode(code int, label string) string {
	switch code {
	case 1101:
		return trimLabelSuffixes(label, " 수집 중", " 수집", " 조회 중", " 조회")
	case 1102:
		return trimLabelSuffixes(label, " 조회 완료", " 완료", " 조회")
	case 1103:
		return trimLabelSuffixes(label, " 실행 중", " 진행 중", " 처리 중", " 실행", " 진행", " 처리")
	case 1104:
		return trimLabelSuffixes(label, " 조회 완료", " 실행 완료", " 진행 완료", " 처리 완료", " 완료됨", " 완료", " 종료", " 끝남")
	default:
		return strings.TrimSpace(label)
	}
}

func trimLabelSuffixes(label string, suffixes ...string) string {
	label = strings.TrimSpace(label)
	for {
		changed := false
		for _, suffix := range suffixes {
			if !strings.HasSuffix(label, suffix) {
				continue
			}
			next := strings.TrimSpace(strings.TrimSuffix(label, suffix))
			if next == "" {
				continue
			}
			label = next
			changed = true
			break
		}
		if !changed {
			return label
		}
	}
}
