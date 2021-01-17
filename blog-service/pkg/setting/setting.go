package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

/**
初始化setting本项目配置的基础属性
*/
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/") // 调用AddConfigPath可设置多个配置路径
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
