grammar Sashimi;

keyValuePair : IDENT EQ IDENT;
keyAtom : ATOM IDENT;
constraintList: HLPAREN (keyAtom | keyValuePair) (SEPERATOR (keyAtom | keyValuePair))* HRPAREN;
unionDecl : UNION LPAREN ALIAS (SEPERATOR ALIAS)* RPAREN;
typeDecl : TYPE ( constraintList )?;
entityRef: ENTITY IDENT;
listDecl : LIST (typeDecl | entityRef);
typeDef : listDecl | unionDecl | typeDecl | entityRef;
aliasDecl : AS ALIAS;
typeIs : IS typeDef;
propDecl : PROP_START IDENT aliasDecl? typeIs;
predicate: IDENT ARROW boolExpression;
qualifier: IDENT(DOT IDENT)*;
loopCall: LOOP LPAREN (GLOBAL | qualifier)  RPAREN (AS alias=IDENT)? (HLPAREN predicate HRPAREN)? (DIRECTIVE blockScope)?;
entityDef : UNIQUE? ENTITY LPAREN IDENT RPAREN OF propDecl (propDecl)*;
scopeBegin: BEGIN (LPAREN SCOPEIDENT RPAREN)?;
scopeEnd: DIRECTIVE END (LPAREN SCOPEIDENT RPAREN)?;
commandCall : COMMAND LPAREN qualifier RPAREN;
blockScope : scopeBegin export* scopeEnd;
export : DIRECTIVE (commandCall | loopCall | entityDef | blockScope);
block : export* EOF;


boolExpression: LPAREN boolExpression RPAREN 
    | NOT boolExpression 
    | left=boolExpression op right=boolExpression
    | truth 
    | qualifier 
    | ALIAS
    | NUMBER;

op : binary | comparator;
comparator : EQ | LEQ | LT | GEQ | GT | NEQ; 
binary: AND | OR;
truth: TRUE | FALSE;

DIRECTIVE : 'sashimi:';
COMMAND : 'display' |  'layout_section' | 'layout' | 'link';
LOOP: 'repeat';
ENTITY: 'entity';
UNIQUE: 'unique';
SEPERATOR : ',';
PROP_START: '-';
OF : 'of';
IS : 'is';
AS : 'as';
BEGIN: 'begin';
END: 'end';
ARROW: '->';
LPAREN : '(';
RPAREN : ')';
HLPAREN : '[';
HRPAREN : ']';
EQ : '=';
LEQ: '<=';
LT: '<';
GEQ: '>=';
GT: '>';
NOT: '!';
NEQ: '<>';
AND: '&';
OR: '|';
ATOM : ':';
DOT : '.';
AT: '@';
TYPE : 'text' | 'picture' | 'number' | 'bool' | 'date';
UNION : 'manyOf' | 'oneOf';
LIST : 'list'; 
ALIAS : '"' ~["\r\n]* '"';
GLOBAL: AT IDENT;
IDENT : [a-zA-Z0-9_]+;
SCOPEIDENT : '\'' [a-zA-Z0-9:]+ '\'';
NUMBER : '-'?[0-9]+('.'[0-9]+)? ;
TRUE: 'true';
FALSE: 'false';
WS : [ \t\r\n]+ -> skip; 
