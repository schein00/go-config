package configuration


import (
	"time"
 )

type ConfigurationEntity struct {
	ID                      int64     `json:"id"`
	IDEmpresa               int64     `json:"id_empresa"`
	URLBackendAPI           string    `json:"url_backend_api"`
	URLBackendAuthentication string   `json:"url_backend_authentication"`
	URLBackendConfiguration string    `json:"url_backend_configuration"`
	InicioVigencia          time.Time `json:"inicio_vigencia"`
	Status                  string    `json:"status"`
 }

func (config *ConfigurationEntity) GetId() int64 {
	return config.id
}

func (config *ConfigurationEntity) SetId(id int64) {
	config.id = id
}

func (config *ConfigurationEntity) GetIdEmpresa() int64 {
	return config.id_empresa
}

func (config *ConfigurationEntity) SetIdEmpresa(id_empresa int64) {
	config.id_empresa = id_empresa
}

func (config *ConfigurationEntity) GetUrlBackendApi() string {
	return config.url_backend_api
}

func (config *ConfigurationEntity) SetUrlBackendApi(url_backend_api string) {
	config.url_backend_api = url_backend_api
}

func (config *ConfigurationEntity) GetUrlBackendAuthentication() string {
	return config.url_backend_authentication
}

func (config *ConfigurationEntity) SetUrlBackendAuthentication(url_backend_authentication string) {
	config.url_backend_authentication = url_backend_authentication
}

func (config *ConfigurationEntity) GetUrlBackendConfiguration() string {
	return config.url_backend_configuration
}

func (config *ConfigurationEntity) SetUrlBackendConfiguration(url_backend_configuration string) {
	config.url_backend_configuration = url_backend_configuration
}

func (config *ConfigurationEntity) GetInicioVigencia() time.Time {
	return config.inicio_vigencia
}

func (config *ConfigurationEntity) SetInicioVigencia(inicio_vigencia time.Time) {
	config.inicio_vigencia = inicio_vigencia
}

func (config *ConfigurationEntity) GetStatus() string {
	return config.status
}

func (config *ConfigurationEntity) SetStatus(status string) {
	config.status = status
}


