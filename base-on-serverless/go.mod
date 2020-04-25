module serverless

go 1.12

require github.com/tencentyun/scf-go-lib v0.0.0-20200116145541-9a6ea1bf75b8

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/offcn-jl/chaos-go-scf v0.0.0-20200422072856-dfa76f029278
	github.com/stretchr/objx v0.2.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/offcn-jl/chaos-go-scf => ../../chaos-go-scf
