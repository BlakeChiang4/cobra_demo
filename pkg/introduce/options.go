package introduce

import "github.com/spf13/pflag"

type IntroduceOptions struct {
	name string
	sex string
	age int32
	hometown string
	father string
}

func NewIntroduceOptions() *IntroduceOptions{
	return &IntroduceOptions{
		name: "none",
		sex: "male",
		age: 18,
		hometown: "none",
		father: "BlakeChiang",
	}
}

func (io *IntroduceOptions) AddFlags(fs *pflag.FlagSet){
	fs.StringVar(&io.name,"name",io.name,"your name")
	fs.StringVar(&io.sex,"sex",io.sex,"male or female")
	fs.Int32Var(&io.age,"age",io.age,"your age")
	fs.StringVarP(&io.hometown,"hometown","t",io.hometown,"where are you from")
	fs.StringVar(&io.father,"father",io.father,"who is your father")
}