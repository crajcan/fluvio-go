package third_thing

type ThirdThing struct {
	Wrapper *string
}

func ThirdThingHello() *ThirdThing {
	s := "Hello, world."
	t := &s

	return &ThirdThing{
		Wrapper: t,
	}
}