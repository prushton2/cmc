package main

const (
    // Node types
    PrimaryExpressionNode = 1
    PostfixExpressionNode = 2
    UnaryExpressionNode   = 3
    CastExpressionNode    = 4
)

type Node struct {
    nodeType int;
    value string
}

func primaryExpression(Node* node, tokenPointer *[]Token) {
		
}

func postfixExpression(Node* node, tokenPointer *[]Token) {
        
}

func unaryExpression(Node* node, tokenPointer *[]Token) {
        
}

func castExpression(Node* node, tokenPointer *[]Token) {
        
}

func parse(tokenPointer *[]Token) {

}