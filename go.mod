module github.com/omniboost/go-tilastokeskus

go 1.24.0

require (
	github.com/elliotchance/pie/v2 v2.9.1
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.46.0
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	golang.org/x/exp v0.0.0-20251017212417-90e834f514db // indirect
	golang.org/x/text v0.30.0 // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
