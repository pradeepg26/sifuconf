grammar Sifu;

sifu
  : ITEM+
  ;

VARIABLE
  : [a-zA-Z][a-zA-Z_.]*[a-zA-Z]
  ;

fragment INT
  : '0' | [1-9] [0-9]*
  ;

fragment EXP
  : [Ee] [+\-]? INT
  ;

NUMBER
  : [+\-]? INT ('.' [0-9] +)? EXP?
  ;

ITEM
  : VARIABLE '=' NUMBER
  ;

ARRAY
  : '[' NUMBER (',' NUMBER)* ']'
  | '[' BOOLEAN (',' BOOLEAN)* ']'
  | '[' ']'
  ;

BOOLEAN
  : 'true'
  | 'false'
  ;

WS
  : [ \t\n\r] + -> skip
  ;

