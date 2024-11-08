package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords is a map to check if a literal is a keyword and get its TokenType
var keywords = map[string]TokenType{
	"fn":     FUNCTION, // A function keyword
	"let":    LET,      // A variable declaration keyword
	"if":     IF,       // A conditional keyword
	"else":   ELSE,     // A conditional keyword
	"return": RETURN,   // A return keyword
}

// LookupIdent checks if an identifier is a keyword or a user-defined identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok // It's a keyword
	}
	return IDENT // Otherwise, it's an identifier
}
