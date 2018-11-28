package parser

import "fmt"

type ParserType interface {
	Init()
	Parse(line, wholeLine, previousLine string) (changeState string, err error)
	Valid() bool
	GetParserName() string
	String() []string
}

type ParserTypes struct {
	parsers []ParserType
	maxSize int
}

func (p *ParserTypes) Get(attribute string) (ParserType, error) {
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("attribute not found")
}

func (p *ParserTypes) Set(parser ParserType) {
	attribute := parser.GetParserName()
	for index, oldParser := range p.parsers {
		if oldParser.GetParserName() == attribute {
			p.parsers[index] = parser
			return
		}
	}
	p.parsers = append(p.parsers, parser)
}
