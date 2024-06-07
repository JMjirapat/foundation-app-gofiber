package errors

var postgresErrorMap = map[string]error{
	"21000": NewBadRequestError("cardinality violation"),
	"22001": NewBadRequestError("string data right truncation"),
	"22002": NewBadRequestError("null value not allowed"),
	"22003": NewBadRequestError("numeric value out of range"),
	"22007": NewBadRequestError("invalid date format"),
	"22019": NewBadRequestError("invalid escape character"),
	"2200B": NewBadRequestError("escape character conflict"),
	"23000": NewBadRequestError("integrity constraint violation"),
	"23502": NewBadRequestError("not null violation"),
	"23503": NewBadRequestError("foreign key violation"),
	"23505": NewBadRequestError("unique violation"),
	"42703": NewBadRequestError("undefined column"),
	"42P01": NewBadRequestError("undefined table"),
	"42P07": NewBadRequestError("duplicate table"),
	"42P08": NewBadRequestError("ambiguous alias"),
	"42P10": NewBadRequestError("invalid column reference"),
	"22023": NewBadRequestError("invalid parameter value"),
	"42P02": NewBadRequestError("undefined parameter"),
}
