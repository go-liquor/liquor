package firebase

import "go.uber.org/fx"

var FirebaseModule = fx.Module("liquor-module-firebase", fx.Provide(
	NewApp,
	NewAuth,
	NewFirestore,
))
