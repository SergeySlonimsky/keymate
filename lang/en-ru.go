package lang

import "unicode"

var EnRu = map[rune]rune{
	'q':  'й',
	'a':  'ф',
	'z':  'я',
	'w':  'ц',
	's':  'ы',
	'x':  'ч',
	'e':  'у',
	'd':  'в',
	'c':  'с',
	'r':  'к',
	'f':  'а',
	'v':  'м',
	't':  'е',
	'g':  'п',
	'b':  'и',
	'y':  'н',
	'h':  'р',
	'n':  'т',
	'u':  'г',
	'j':  'о',
	'm':  'ь',
	'i':  'ш',
	'k':  'л',
	',':  'б',
	'o':  'щ',
	'l':  'д',
	'.':  'ю',
	'p':  'з',
	';':  'ж',
	'/':  '/',
	'[':  'х',
	'\'': 'э',
	']':  'ъ',
	'\\': 'ё',
	'<':  'Б',
	'>':  'Ю',
	':':  'Ж',
	'?':  '?',
	'{':  'Х',
	'"':  'Э',
	'}':  'Ъ',
	'|':  'Ё',
	'^':  ',',
	'&':  '.',
	'*':  ';',
}

func Convert(text string, set map[rune]rune) string {
	var result string
	for _, rune := range text {
		upper := false
		if unicode.IsUpper(rune) {
			upper = true
			rune = unicode.ToLower(rune)
		}
		if char, ok := set[rune]; ok {
			if upper {
				char = unicode.ToUpper(char)
			}
			result = result + string(char)
		} else {
			if upper {
				rune = unicode.ToUpper(rune)
			}
			result = result + string(rune)
		}
	}
	return result
}
