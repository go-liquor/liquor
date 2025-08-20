package firebase

import "go.uber.org/fx"

var FirebaseModule = fx.Module("lq-module-firebase", fx.Provide(
	NewApp,
	NewAuth,
	NewFirestore,
))
