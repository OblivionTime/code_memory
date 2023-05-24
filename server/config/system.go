package config

type System struct {
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 端口值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                   // 数据库类型:mysql(默认)
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` //  路由全局前缀
}
