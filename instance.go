package readline

import (
	"os"
	"regexp"
)

// Instance is used to encapsulate the parameter group and run time of any given
// readline instance so that you can reuse the readline API for multiple entry
// captures without having to repeatedly unload configuration.
type Instance struct {

	// Public Prompt and Vim parameters/functions
	Multiline       bool   // If set to true, the shell will have a two-line prompt.
	MultilinePrompt string // The second line of the prompt, where input follows.

	// VimModePrompt - If set to true, the MultilinePrompt variable will be erased,
	// and instead will be '[i] >' or '[N] >' for indicating the current Vim mode
	VimModePrompt   bool
	VimModeColorize bool // If set to true, varies colors of the VimModePrompt

	// RefreshMultiline allows the user's program to refresh the input prompt.
	// In this version, the prompt is treated like the input line following it:
	// we can refresh it at any time, like we do with SyntaxHighlighter below.
	RefreshMultiline func([]rune) string

	// PasswordMask is what character to hide password entry behind.
	// Once enabled, set to 0 (zero) to disable the mask again.
	PasswordMask rune

	// SyntaxHighlight is a helper function to provide syntax highlighting.
	// Once enabled, set to nil to disable again.
	SyntaxHighlighter func([]rune) string

	// History is an interface for querying the readline history.
	// This is exposed as an interface to allow you the flexibility to define how
	// you want your history managed (eg file on disk, database, cloud, or even
	// no history at all). By default it uses a dummy interface that only stores
	// historic items in memory.
	History History

	// HistoryAutoWrite defines whether items automatically get written to
	// history.
	// Enabled by default. Set to false to disable.
	HistoryAutoWrite bool // = true

	// TabCompleter is a simple function that offers completion suggestions.
	// It takes the readline line ([]rune) and cursor pos.
	// Returns a prefix string, and several completion groups with their items and description
	TabCompleter func([]rune, int) (string, []*CompletionGroup)

	// MaxTabCompletionRows is the maximum number of rows to display in the tab
	// completion grid.
	MaxTabCompleterRows int // = 4

	// SyntaxCompletion is used to autocomplete code syntax (like braces and
	// quotation marks). If you want to complete words or phrases then you might
	// be better off using the TabCompletion function.
	// SyntaxCompletion takes the line ([]rune) and cursor position, and returns
	// the new line and cursor position.
	SyntaxCompleter func([]rune, int) ([]rune, int)

	// HintText is a helper function which displays hint text the prompt.
	// HintText takes the line input from the promt and the cursor position.
	// It returns the hint text to display.
	HintText func([]rune, int) []rune

	// HintColor any ANSI escape codes you wish to use for hint formatting. By
	// default this will just be blue.
	HintFormatting string

	// TempDirectory is the path to write temporary files when editing a line in
	// $EDITOR. This will default to os.TempDir()
	TempDirectory string

	// GetMultiLine is a callback to your host program. Since multiline support
	// is handled by the application rather than readline itself, this callback
	// is required when calling $EDITOR. However if this function is not set
	// then readline will just use the current line.
	GetMultiLine func([]rune) []rune

	// readline operating parameters
	prompt        string //  = ">>> "
	mlnPrompt     []rune // Our multiline prompt, different from multiline below
	mlnArrow      []rune
	promptLen     int    //= 4
	line          []rune // This is the input line, with entered text: full line = mlnPrompt + line
	pos           int
	multiline     []byte
	multisplit    []string
	skipStdinRead bool

	// history
	lineBuf string
	histPos int

	// hint text
	hintY    int //= 0
	hintText []rune

	// tab completion
	tcGroups          []*CompletionGroup // All of our suggestions tree is in here
	modeTabCompletion bool
	tcPrefix          string
	tcOffset          int
	tcPosX            int
	tcPosY            int
	tcMaxX            int
	tcMaxY            int
	tcUsedY           int
	tcMaxLength       int

	// Tab Find
	modeTabFind  bool           // This does not change, because we will seach in all options, no matter the group
	tfLine       []rune         // The current search pattern entered
	modeAutoFind bool           // for when invoked via ^R or ^F outside of [tab]
	searchMode   FindMode       // Used for varying hints, and underlying functions called
	regexSearch  *regexp.Regexp // Holds the current search regex match

	// vim
	modeViMode       viMode //= vimInsert
	viIteration      string
	viUndoHistory    []undoItem
	viUndoSkipAppend bool
	viYankBuffer     string

	// event

	evtKeyPress map[string]func(string, []rune, int) *EventReturn
}

// NewInstance is used to create a readline instance and initialise it with sane
// defaults.
func NewInstance() *Instance {
	rl := new(Instance)

	//GetTermWidth()

	rl.History = new(ExampleHistory)
	rl.HistoryAutoWrite = true
	rl.MaxTabCompleterRows = 100
	rl.prompt = ">>> "
	// rl.promptLen = len(rl.computePrompt()) // We need
	rl.mlnArrow = []rune{' ', '>', ' '}
	rl.HintFormatting = seqFgBlue
	rl.evtKeyPress = make(map[string]func(string, []rune, int) *EventReturn)

	rl.TempDirectory = os.TempDir()

	return rl
}
