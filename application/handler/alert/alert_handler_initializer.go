package alert

import (
	"hu-design-project/application/handler/alert/handler_impl"
	"hu-design-project/application/repository"
)

func InitializeAlertHandler(alertRepository repository.AlertsRepository, whitelistRepository repository.WhitelistRepository, blacklistRepository repository.BlacklistRepository) *Handler {
	alertHandler := Handler{}
	alertHandler.GetAlerts = handler_impl.NewGetAlertsHandler(alertRepository)
	alertHandler.AddWhitelist = handler_impl.NewAddWhitelistHandler(alertRepository, whitelistRepository)
	alertHandler.AddBlacklist = handler_impl.NewAddBlacklistHandler(alertRepository, blacklistRepository)
	alertHandler.GetWhitelist = handler_impl.NewGetWhitelistHandler(whitelistRepository)
	alertHandler.GetBlacklist = handler_impl.NewGetBlacklistHandler(blacklistRepository)
	alertHandler.DeleteWhitelist = handler_impl.NewDeleteWhitelistHandler(whitelistRepository)
	alertHandler.DeleteBlacklist = handler_impl.NewDeleteBlacklistHandler(blacklistRepository)
	return &alertHandler
}
