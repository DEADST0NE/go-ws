package api

import "fmt"

func makeRoute(path string) string {
	return fmt.Sprintf("/api/v1/%s", path)
}

type Routes struct {
	API APIRoutes
}

type APIRoutes struct {
	V1 APIVersion1Routes
}

type APIVersion1Routes struct {
	GET_CHART string
}

var routes = Routes{
	API: APIRoutes{
		V1: APIVersion1Routes{
			GET_CHART: makeRoute("chart"),
		},
	},
}
