package service

import (
	settingmodule "FilesWithTag/setting_module"
	"net/http"
)

// @Summary		Получение настроек
// @Tags		setting
// @Description	Получение настроек запуска
// @ID			getSettings
// @Accept		json
// @Produce		json
// @Success		200		{object}	settingmodule.Setting	"settings"
// @Router		/Settings [get]
func GetSettings(w http.ResponseWriter, r *http.Request) any {
	settings := settingmodule.GetSetting()

	return *settings
}

// @Summary		Получение настроек
// @Tags		setting
// @Description	Получение настроек запуска
// @ID			getSettings
// @Accept		json
// @Produce		json
// @Param		NameSetting		query	string	true	"name setting"
// @Param		ValueSetting	query	string	true	"value setting"
// @Success		200			{object}	Responce	"settings"
// @Failure		400,500		{object}	ResponceError	"error"
// @Router		/Settings [put]
func ChangeSettings(w http.ResponseWriter, r *http.Request) any {
	NameSetting := r.URL.Query().Get("NameSetting")
	ValueSetting := r.URL.Query().Get("ValueSetting")

	if NameSetting == "" || ValueSetting == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	if err := settingmodule.GetSetting().Change(NameSetting, ValueSetting); err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	return Responce{Status: http.StatusOK, Comment: "OK"}
}
