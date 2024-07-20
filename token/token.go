package token

type TokenType string;

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEAGAL = "ILLEAGAL"
  EOF      = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y
  INT = "INT" //integer

  // Operators
  ASSIGN = "="
  ADD = "+"

  //Delim
  COMMA = ","
  SEMI = ";"

  LPAREN = "("
  RPAREN = ")"

  LBRACE = "{"
  RBRACE = "}"


  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"

)

