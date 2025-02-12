module github.com/omniboost/go-tilastokeskus

go 1.23.6

require (
	github.com/cydev/zero v0.0.0-20160322155811-4a4535dd56e7
	github.com/elliotchance/pie/v2 v2.9.1
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/omniboost/go-unit4-multivers v0.0.0-20200928111357-29606eb52519
	github.com/pkg/errors v0.9.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require golang.org/x/exp v0.0.0-20241210194714-1829a127f884 // indirect

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
