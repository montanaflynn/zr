// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-23 17:50:57 972508                          [zr/settings_test.go]
// -----------------------------------------------------------------------------

package zr

// # Methods
//   Test_sett_Settings_GetSetting_
//   Test_sett_Settings_HasSetting_
//   Test_sett_Settings_SetSetting_
//
// # Extenders
//   Test_sett_Settings_ExtendGet_
//   Test_sett_Settings_ExtendHas_
//   Test_sett_Settings_ExtendSet_

/*
to test all methods use:
	go test --run Test_sett_

to generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import "testing" // standard

// -----------------------------------------------------------------------------
// # Methods

// go test --run Test_sett_Settings_GetSetting_
func Test_sett_Settings_GetSetting_(t *testing.T) {
	TBegin(t)
	// (ob *Settings) GetSetting(name string) string
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var nulp *Settings
		TEqual(t, nulp.GetSetting("alpha"), (""))
		TCheckError(t, ENilReceiver)
	}
	{
		// zero-length name must log an error
		TBeginError()
		var o Settings
		TEqual(t, o.GetSetting(""), (""))
		TCheckError(t, EInvalidArg)
	}
	{
		// check if settings are read properly
		var o = Settings{
			m: map[string]string{
				"alpha": "111",
				"beta":  "222",
				"blank": "",
			},
		}
		// existing settings are looked-up irrespective of case
		TEqual(t, o.GetSetting("alpha"), ("111"))
		TEqual(t, o.GetSetting("beta"), ("222"))
		TEqual(t, o.GetSetting("ALPHA"), ("111"))
		TEqual(t, o.GetSetting("BETA"), ("222"))
		//
		// non-existent settings just return a blank string
		TEqual(t, o.GetSetting("delta"), (""))
		//
		// must store blank setting values
		TEqual(t, o.GetSetting("blank"), (""))
	}
} //                                              Test_sett_Settings_GetSetting_

// go test --run Test_sett_Settings_HasSetting_
func Test_sett_Settings_HasSetting_(t *testing.T) {
	TBegin(t)
	// (ob *Settings) HasSetting(name string) bool
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var nulp *Settings
		var res = nulp.HasSetting("alpha")
		TEqual(t, res, (false))
		TCheckError(t, ENilReceiver)
	}
	{
		// zero-length name must log an error
		TBeginError()
		var o Settings
		TEqual(t, o.GetSetting(""), (""))
		TCheckError(t, EInvalidArg)
	}
	{
		// check if settings are read properly
		var o = Settings{
			m: map[string]string{
				"alpha": "111",
				"beta":  "222",
				"blank": "",
			},
		}
		TEqual(t, o.HasSetting("ALPHA"), (true))
		TEqual(t, o.HasSetting("BETA"), (true))
		TEqual(t, o.HasSetting("alpha"), (true))
		TEqual(t, o.HasSetting("beta"), (true))
		TEqual(t, o.HasSetting("blank"), (true))
		//
		TEqual(t, o.HasSetting("delta"), (false))
	}
} //                                              Test_sett_Settings_HasSetting_

// go test --run Test_sett_Settings_SetSetting_
func Test_sett_Settings_SetSetting_(t *testing.T) {
	TBegin(t)
	// (ob *Settings) SetSetting(name string, val interface{})
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var nulp *Settings
		nulp.SetSetting("alpha", 1)
		TCheckError(t, ENilReceiver)
	}
	{
		// zero-length name must log an error
		TBeginError()
		var o Settings
		TEqual(t, o.GetSetting(""), (""))
		TCheckError(t, EInvalidArg)
	}
	{
		// check if settings are set properly
		var o = Settings{
			m: map[string]string{
				"alpha": "111",
				"beta":  "222",
				"blank": "",
			},
		}
		{
			o.SetSetting("alpha", "changed")
			TTrue(t, o.GetSetting("alpha") == "changed")
		}
		{
			o.SetSetting("delta", "new")
			TTrue(t, o.GetSetting("delta") == "new")
		}
	}
} //                                              Test_sett_Settings_SetSetting_

// -----------------------------------------------------------------------------
// # Extenders

// go test --run Test_sett_Settings_ExtendGet_
func Test_sett_Settings_ExtendGet_(t *testing.T) {
	TBegin(t)
	// (ob *Settings) ExtendGet(
	//     handler func(name, val string, exists bool) string,
	// )
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var nulp *Settings
		nulp.ExtendGet(nil)
		TCheckError(t, ENilReceiver)
	}
	//TODO: check if extender function is called after this is set
	//TODO: must not call extender function after being set to nil
	//TODO: check if extender is passed correct arguments
	//TODO: check if extender's result is used
} //                                               Test_sett_Settings_ExtendGet_

// go test --run Test_sett_Settings_ExtendHas_
func Test_sett_Settings_ExtendHas_(t *testing.T) {
	TBegin(t)
	// (ob *Settings) ExtendHas(
	//     handler func(name, val string, exists bool) bool,
	// )
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var nulp *Settings
		nulp.ExtendHas(nil)
		TCheckError(t, ENilReceiver)
	}
	//TODO: check if extender function is called after this is set
	//TODO: must not call extender function after being set to nil
	//TODO: check if extender is passed correct arguments
	//TODO: check if extender's result is used
} //                                               Test_sett_Settings_ExtendHas_

// go test --run Test_sett_Settings_ExtendSet_
func Test_sett_Settings_ExtendSet_(t *testing.T) {
	TBegin(t)
	// (ob *Settings) ExtendSet(
	//     handler func(name string, old, val interface{}) *string,
	// )
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var nulp *Settings
		nulp.ExtendSet(nil)
		TCheckError(t, ENilReceiver)
	}
	//TODO: check if extender function is called after this is set
	//TODO: must not call extender function after being set to nil
	//TODO: check if extender is passed correct arguments
	//TODO: check if extender's result is used
} //                                               Test_sett_Settings_ExtendSet_

//end
