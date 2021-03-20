package goRest

type IDatabaseHelper interface {
	/*correctOffset: returns an obtained max in int type. Invalid value, value<1, or value>1000 are rejected and replaced by default value */
	ProcessMax(max string) (maxInt int64)
	/*correctOffset: returns an obtained offset in int type. Invalid value or negative are converted to default value:0 */
	ProcessOffset(offset string) (offsetInt int64)
}
