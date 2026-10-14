package progressmessage

import "strings"

func Render(code int, args map[string]string, locale string) (string, bool) {
	ir, err := Catalog()
	if err != nil {
		return "", false
	}
	for _, message := range ir.Messages {
		if message.Code != code {
			continue
		}
		template := message.Locales[strings.TrimSpace(locale)]
		if template == "" {
			template = message.Locales[DefaultLocale]
		}
		if template == "" {
			return "", false
		}
		return renderTemplate(template, args), true
	}
	return "", false
}
