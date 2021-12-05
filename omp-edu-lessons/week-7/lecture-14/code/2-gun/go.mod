module github.com/ozonmp/omp-edu-lessons/week-6/lecture-14/1-gun

go 1.16

require (
	github.com/ozonmp/omp-edu-lessons/week-6/lecture-14/1-load-testing/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/spf13/afero v1.6.0
	github.com/yandex/pandora v0.3.5
	go.uber.org/zap v1.19.1
)

replace github.com/ozonmp/omp-edu-lessons/week-6/lecture-14/1-load-testing/api => ../1-load-testing/api
