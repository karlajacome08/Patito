grammar Patito;

// Parser rules

program
  : PROGRAMA ID ';' vars? funcDef* INICIO body FIN EOF
  ;

vars
  : VARS varDecl*
  ;

varDecl
  : idList ':' type_ ';'
  ;

idList
  : ID (',' ID)*
  ;

type_
  : ENTERO
  | FLOTANTE
  ;

funcDef
  : (NULA | type_) ID '(' params? ')' '{' vars? body '}' ';'
  ;

params
  : param (',' param)*
  ;

param
  : ID ':' type_
  ;

body
  : '{' stmt* '}'
  ;

stmt
  : assign 
  | print
  | ifStmt
  | whileStmt
  | '[' stmt* ']' 
  | call ';'
  ;

assign
  : ID '=' expr ';'
  ;

print
  : ESCRIBE '(' printItemList? ')' ';'
  ;

printItemList
  : printItem (',' printItem)*
  ;

printItem
  : STRING
  | expr
  ;

ifStmt
  : SI '(' expr ')' ifMarkThen body (SINO elseMark body)? ';'
  ;

ifMarkThen
  : 
  ;

elseMark
  : 
  ;

whileStmt
  : MIENTRAS whileStart '(' expr ')' whileCond HAZ body ';'
  ;

whileStart
  : 
  ;

whileCond
  : 
  ;


call
  : ID '(' (expr (',' expr)*)? ')'
  ;

expr
  : sum (relop sum)? # RelExpr
  ;

relop
  : '==' | '!=' | '<' | '>'
  ;

sum
  : prod (('+'|'-') prod)*            # AddSub
  ;

prod
  : unary (('*'|'/') unary)*          # MulDiv
  ;

unary
  : ('+'|'-'|'!') unary               # Prefix
  | primary                           # ToPrimary
  ;

primary
  : '(' expr ')'                      # Parens
  | call                               # CallPrim
  | ID                                 # Ident
  | INT                                # Int
  | FLOAT                              # Float
  ;



// Lexer rules

// Palabras clave
PROGRAMA : 'programa';
VARS     : 'vars';
INICIO   : 'inicio';
FIN      : 'fin';
SI       : 'si';
SINO     : 'sino';
MIENTRAS : 'mientras';
HAZ      : 'haz';
ESCRIBE  : 'escribe';
NULA     : 'nula';
ENTERO   : 'entero';
FLOTANTE : 'flotante';

// Identificadores 
ID       : [a-zA-Z_] [a-zA-Z0-9_]* ;
INT      : '0' | [1-9] [0-9]* ;
FLOAT    : [0-9]+ '.' [0-9]+ ;

STRING
  : '"' ( '\\' ['"\\ntr] | ~["\\\r\n] )* '"'
  ;

// Espacios y comentarios
WS              : [ \t\r\n]+ -> skip ;
LINE_COMMENT    : '//' ~[\r\n]* -> skip ;
BODY_COMMENT   : '/*' .*? '*/' -> skip ;
