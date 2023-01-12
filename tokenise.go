package readline

import "strings"

// tokeniser - The input line must be splitted according to different rules (split between spaces, brackets, etc ?).
type tokeniser func(line []rune, cursorPos int) (split []string, index int, newPos int)

func tokeniseLine(line []rune, linePos int) ([]string, int, int) {
	if len(line) == 0 {
		return nil, 0, 0
	}

	var index, pos int
	var punc bool

	split := make([]string, 1)

	for i, r := range line {
		switch {
		case isPunctuation(r):
			if i > 0 && line[i-1] != r {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)
			punc = true

		case r == ' ' || r == '\t' || r == '\n':
			split[len(split)-1] += string(r)
			punc = true

		default:
			if punc {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)
			punc = false
		}

		// Not caught when we are appending to the end
		// of the line, where rl.pos = linePos + 1, so...
		if i == linePos {
			index = len(split) - 1
			pos = len(split[index]) - 1
		}
	}

	// ... so we ajust here for this case.
	if linePos == len(line) {
		index = len(split) - 1
		pos = len(split[index])
	}

	return split, index, pos
}

func tokeniseSplitSpaces(line []rune, linePos int) ([]string, int, int) {
	if len(line) == 0 {
		return nil, 0, 0
	}

	var index, pos int
	split := make([]string, 1)

	for i, r := range line {
		switch r {
		case ' ', '\t', '\n':
			split[len(split)-1] += string(r)

		default:
			if i > 0 && (line[i-1] == ' ' || line[i-1] == '\t' || line[i-1] == '\n') {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)
		}

		// Not caught when we are appending to the end
		// of the line, where rl.pos = linePos + 1, so...
		if i == linePos {
			index = len(split) - 1
			pos = len(split[index]) - 1
		}
	}

	// ... so we ajust here for this case.
	if linePos == len(line) {
		index = len(split) - 1
		pos = len(split[index])
	}

	return split, index, pos
}

func tokeniseBrackets(line []rune, linePos int) ([]string, int, int) {
	var (
		open, close    rune
		split          []string
		count          int
		pos            = make(map[int]int)
		match          int
		single, double bool
	)

	switch line[linePos] {
	case '(', ')':
		open = '('
		close = ')'

	case '{', '[':
		open = line[linePos]
		close = line[linePos] + 2

	case '}', ']':
		open = line[linePos] - 2
		close = line[linePos]

	default:
		return nil, 0, 0
	}

	for i := range line {
		switch line[i] {
		case '\'':
			if !single {
				double = !double
			}

		case '"':
			if !double {
				single = !single
			}

		case open:
			if !single && !double {
				count++
				pos[count] = i
				if i == linePos {
					match = count
					split = []string{string(line[:i-1])}
				}

			} else if i == linePos {
				return nil, 0, 0
			}

		case close:
			if !single && !double {
				if match == count {
					split = append(split, string(line[pos[count]:i]))
					return split, 1, 0
				}
				if i == linePos {
					split = []string{
						string(line[:pos[count]-1]),
						string(line[pos[count]:i]),
					}
					return split, 1, len(split[1])
				}
				count--

			} else if i == linePos {
				return nil, 0, 0
			}
		}
	}

	return nil, 0, 0
}

func rTrimWhiteSpace(oldString string) (newString string) {
	return strings.TrimRightFunc(oldString, func(r rune) bool {
		if r == ' ' || r == '\t' || r == '\n' {
			return true
		}
		return false
	})
	// return strings.TrimS(oldString, " ")
}

// isPunctuation returns true if the rune is non-blank word delimiter.
func isPunctuation(r rune) bool {
	if (r >= 33 && 47 >= r) ||
		(r >= 58 && 64 >= r) ||
		(r >= 91 && 94 >= r) ||
		r == 96 ||
		(r >= 123 && 126 >= r) {
		return true
	}

	return false
}
