// Code generated by make generate-models DO NOT EDIT.

package notifications

type BookInstrumentNameIntervalFirst struct {
	Asks           []interface{} `json:"asks" mapstructure:"asks"`
	Bids           [][]float64   `json:"bids" mapstructure:"bids"`
	ChangeID       int64         `json:"change_id" mapstructure:"change_id"`
	InstrumentName string        `json:"instrument_name" mapstructure:"instrument_name"`
	Timestamp      int64         `json:"timestamp" mapstructure:"timestamp"`
}