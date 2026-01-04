package context

import "strings"

type AppRoute uint8

// This list of views can be expanded
// according to the needs of the application.
const (
	list     string = "LIST"
	settings string = "SETTINGS"
	today    string = "TODAY"
	trash    string = "TRASH"
	invalid  string = "INVALID_ROUTE"
)

const (
	List AppRoute = iota
	Settings
	Today
	Trash
)

func getMapping() map[AppRoute]string {

	return map[AppRoute]string{
		List:     list,
		Settings: settings,
		Today:    today,
		Trash:    trash,
	}
}

func getReverseMapping() map[string]AppRoute {

	return map[string]AppRoute{
		list:     List,
		settings: Settings,
		today:    Today,
		trash:    Trash,
	}
}

func RouteFromString(route string) AppRoute {
	route = strings.ToUpper(route)
	mapping := getReverseMapping()
	if value, ok := mapping[route]; ok {
		return value
	}

	return List
}

func (a AppRoute) String() string {
	mapping := getMapping()
	if value, ok := mapping[a]; ok {
		return value
	}

	return invalid
}
