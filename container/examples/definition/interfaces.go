package definition

type SomeInterface interface {
	SomeInterfaceMethod()
}

type InterfaceImplOne struct{}

func (iio InterfaceImplOne) SomeInterfaceMethod() {

}

type InterfaceImplTwo struct{}

func (iit InterfaceImplTwo) SomeInterfaceMethod() {

}

