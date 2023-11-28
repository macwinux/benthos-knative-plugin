package bloblang

import "github.com/benthosdev/benthos/v4/public/bloblang"

func init() {
	lenSpec := bloblang.NewPluginSpec().
		Category("String Manipulation").
		Description("Return the length of the string.")

	if err := bloblang.RegisterMethodV2("len_s", lenSpec,
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return bloblang.StringMethod(func(s string) (any, error) {
				return len(s), nil
			}), nil
		},
	); err != nil {
		panic(err)
	}

}
