program     -> declaration* EOF

declaration -> varDecl | stmt
stmt        -> printStmt | exprStmt

printStmt   -> print expression ;
exprStmt    -> expression ;
varDecl     -> "var" IDENTIFIER ("=" expression)? ;

expression  -> equality

equality    -> comparison (("!=" | "==") comparison)*
comparison  -> term ((">" | ">=" | "<" | "<=") term)*
term        -> factor (("+" | "-") factor)*
factor      -> unary (("/" | "*") unary)*
unary       -> ("!" | "-") unary | primary;
primary     -> IDENTIFIER | NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")";
