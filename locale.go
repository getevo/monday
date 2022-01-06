package monday

// Locale identifies locales supported by 'monday' package.
// Monday uses ICU locale identifiers. See http://userguide.icu-project.org/locale
type Locale string

// Locale constants represent all locales that are currently supported by
// this package.
const (
	LocaleEnUS = "en-US" // English (United States)
	LocaleEnGB = "en-GB" // English (United Kingdom)
	LocaleDaDK = "da-DK" // Danish (Denmark)
	LocaleNlBE = "nl-BE" // Dutch (Belgium)
	LocaleNlNL = "nl-NL" // Dutch (Netherlands)
	LocaleFiFI = "fi-FI" // Finnish (Finland)
	LocaleFrFR = "fr-FR" // French (France)
	LocaleFrCA = "fr-CA" // French (Canada)
	LocaleDeDE = "de-DE" // German (Germany)
	LocaleHuHU = "hu-HU" // Hungarian (Hungary)
	LocaleItIT = "it-IT" // Italian (Italy)
	LocaleNnNO = "nn-NO" // Norwegian Nynorsk (Norway)
	LocaleNbNO = "nb-NO" // Norwegian Bokmål (Norway)
	LocalePlPL = "pl-PL" // Polish (Poland)
	LocalePtPT = "pt-PT" // Portuguese (Portugal)
	LocalePtBR = "pt-BR" // Portuguese (Brazil)
	LocaleRoRO = "ro-RO" // Romanian (Romania)
	LocaleRuRU = "ru-RU" // Russian (Russia)
	LocaleEsES = "es-ES" // Spanish (Spain)
	LocaleCaES = "ca-ES" // Catalan (Spain)
	LocaleSvSE = "sv-SE" // Swedish (Sweden)
	LocaleTrTR = "tr-TR" // Turkish (Turkey)
	LocaleUkUA = "uk-UA" // Ukrainian (Ukraine)
	LocaleBgBG = "bg-BG" // Bulgarian (Bulgaria)
	LocaleZhCN = "zh-CN" // Chinese (Mainland)
	LocaleZhTW = "zh-TW" // Chinese (Taiwan)
	LocaleZhHK = "zh-HK" // Chinese (Hong Kong)
	LocaleKoKR = "ko-KR" // Korean (Korea)
	LocaleJaJP = "ja-JP" // Japanese (Japan)
	LocaleElGR = "el-GR" // Greek (Greece)
	LocaleIdID = "id-ID" // Indonesian (Indonesia)
	LocaleFrGP = "fr-GP" // French (Guadeloupe)
	LocaleFrLU = "fr-LU" // French (Luxembourg)
	LocaleFrMQ = "fr-MQ" // French (Martinique)
	LocaleFrRE = "fr-RE" // French (Reunion)
	LocaleFrGF = "fr-GF" // French (French Guiana)
	LocaleCsCZ = "cs-CZ" // Czech (Czech Republic)
	LocaleSlSI = "sl-SI" // Slovenian (Slovenia)
	LocaleLtLT = "lt-LT" // Lithuanian (Lithuania)
)

// ListLocales returns all locales supported by the package.
func ListLocales() []Locale {
	return []Locale{
		LocaleEnUS,
		LocaleEnGB,
		LocaleDaDK,
		LocaleNlBE,
		LocaleNlNL,
		LocaleFiFI,
		LocaleFrFR,
		LocaleFrCA,
		LocaleDeDE,
		LocaleHuHU,
		LocaleItIT,
		LocaleNnNO,
		LocaleNbNO,
		LocalePlPL,
		LocalePtPT,
		LocalePtBR,
		LocaleRoRO,
		LocaleRuRU,
		LocaleEsES,
		LocaleCaES,
		LocaleSvSE,
		LocaleTrTR,
		LocaleUkUA,
		LocaleBgBG,
		LocaleZhCN,
		LocaleZhTW,
		LocaleZhHK,
		LocaleKoKR,
		LocaleJaJP,
		LocaleElGR,
		LocaleFrGP,
		LocaleFrLU,
		LocaleFrMQ,
		LocaleFrRE,
		LocaleFrGF,
		LocaleCsCZ,
		LocaleSlSI,
		LocaleLtLT,
	}
}
