package omnijson

type GetNewAddressResult = string

type GetNewAddressCommand struct {
}

func (GetNewAddressCommand) Method() string {
	return "getnewaddress"
}

func (GetNewAddressCommand) ID() string {
	return "1"
}

func (cmd GetNewAddressCommand) Params() []interface{} {
	return []interface{}{}
}
