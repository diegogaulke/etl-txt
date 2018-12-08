package main

import (
	"regexp"
	"strconv"
)

const typeDebtX = "debtX"
const typeDebtY = "debtY"

const regexDebtX = "(.{1})(.{11})(.{10})(.{14})"
const regexDebtY = "(.{1})(.{11})(.{10})(.{14})"

type Item struct {
	debtType string
	data     interface{}
}

type Operation struct {
	op   string
	item *Item
}

type Extractor interface {
	Extract(line string) (*Operation, error)
}

type EDebtX struct {
	Extractor
	re *regexp.Regexp
}

type EDebtY struct {
	Extractor
	re *regexp.Regexp
}

type DebtX struct {
	key      string
	document string
	value    float64
	creditor string
}

type DebtY struct {
	key      string
	document string
	value    float64
	creditor string
}

func NewExtractor(fileType string) Extractor {
	switch fileType {
	case typeDebtX:
		return &EDebtX{re: regexp.MustCompile(regexDebtX)}
	case typeDebtY:
		return &EDebtY{re: regexp.MustCompile(regexDebtY)}
	}

	return nil
}

func (dx *EDebtX) Extract(line string) (*Operation, error) {
	match := dx.re.FindStringSubmatch(line)

	if len(match) > 0 {
		dx := new(DebtX)
		dx.document = match[2]

		item := new(Item)
		item.debtType = typeDebtX
		item.data = dx

		op := new(Operation)
		op.op = match[1]
		op.item = item

		return op, nil
	}

	return nil, nil
}

func formatValue(raw string) float64 {
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0
	}
	return f / 100
}
