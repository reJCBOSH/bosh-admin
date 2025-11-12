package config

type Config struct {
    Server    Server    `mapstructure:"server" json:"server" yaml:"server"`          // 服务配置
    Log       Log       `mapstructure:"log" json:"log" yaml:"log"`                   // 日志配置
    Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`             // Mysql 配置
    Pgsql     Pgsql     `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`             // Pgsql 配置
    JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                   // JWT 配置
    Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`       // 图片验证码配置
    Local     Local     `mapstructure:"local" json:"local" yaml:"local"`             // 本地文件存储配置
    AliyunOss AliyunOss `mapstructure:"aliyunOss" json:"aliyunOss" yaml:"aliyunOss"` // 阿里云OSS配置
}

// Server 服务配置
type Server struct {
    Env      string `mapstructure:"env" json:"env" yaml:"env"`                // 服务环境
    Name     string `mapstructure:"name" json:"name" yaml:"name"`             // 服务名称
    Url      string `mapstructure:"url" json:"url" yaml:"url"`                // 服务地址
    Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 服务端口
    Database string `mapstructure:"database" json:"database" yaml:"database"` // 数据库类型
    OssType  string `mapstructure:"ossType" json:"ossType" yaml:"ossType"`    // 存储类型
    BlackJWT bool   `mapstructure:"blackJWT" json:"blackJWT" yaml:"blackJWT"` // 是否启用黑名单JWT
}

// Log 日志配置
type Log struct {
    RootDir         string `mapstructure:"rootDir" json:"rootDir" yaml:"rootDir"`                         // 日志根目录
    Format          string `mapstructure:"format" json:"format" yaml:"format"`                            // 写入格式
    TimestampFormat string `mapstructure:"timestampFormat" json:"timestampFormat" yaml:"timestampFormat"` // 日志时间格式
    MaxBackups      int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups"`                // 旧文件的最大个数
    MaxSize         int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`                         // 日志文件最大大小(MB)
    MaxAge          int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`                            // 旧文件的最大保留天数
    Compress        bool   `mapstructure:"compress" json:"compress" yaml:"compress"`                      // 是否压缩
}

// Mysql Mysql 配置
type Mysql struct {
    Username     string `mapstructure:"username" json:"username" yaml:"username"`             // 用户名
    Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 密码
    IP           string `mapstructure:"ip" json:"ip" yaml:"ip"`                               // IP地址
    Port         string `mapstructure:"port" json:"port" yaml:"port"`                         // 端口
    Database     string `mapstructure:"database" json:"database" yaml:"database"`             // 数据库名
    Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 连接配置
    MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"` // 空闲中的最大连接数
    MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"` //  打开的最大连接数
}

// Pgsql Pgsql 配置
type Pgsql struct {
    Host         string `mapstructure:"host" json:"host" yaml:"host"`                         // 主机
    User         string `mapstructure:"user" json:"user" yaml:"user"`                         // 用户名
    Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 密码
    Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                   // 数据库名
    Port         string `mapstructure:"port" json:"port" yaml:"port"`                         // 端口
    Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 连接配置
    MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"` // 空闲中的最大连接数
    MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"` //  打开的最大连接数
}

// JWT JWT配置
type JWT struct {
    AccessSecret    string `mapstructure:"accessSecret" json:"accessSecret" yaml:"accessSecret"`          // access token 签名密钥
    RefreshSecret   string `mapstructure:"refreshSecret" json:"refreshSecret" yaml:"refreshSecret"`       // refresh token 签名密钥
    AccessDuration  int64  `mapstructure:"accessDuration" json:"accessDuration" yaml:"accessDuration"`    // access token 有效时长
    RefreshDuration int64  `mapstructure:"refreshDuration" json:"refreshDuration" yaml:"refreshDuration"` // refresh token 有效时长
    BufferDuration  int64  `mapstructure:"bufferDuration" json:"bufferDuration" yaml:"bufferDuration"`    // 缓冲时长
}

// Captcha 图片验证码配置
type Captcha struct {
    KeyLong   int `mapstructure:"keyLong" json:"keyLong" yaml:"keyLong"`       // 验证码长度
    ImgWidth  int `mapstructure:"imgWidth" json:"imgWidth" yaml:"imgWidth"`    // 图片宽度
    ImgHeight int `mapstructure:"imgHeight" json:"imgHeight" yaml:"imgHeight"` // 图片高度
}

// Local 本地文件配置
type Local struct {
    Path      string `mapstructure:"path" json:"path" yaml:"path"`                // 本地文件访问路径
    StorePath string `mapstructure:"storePath" json:"storePath" yaml:"storePath"` // 本地文件存储路径
}

// AliyunOss 阿里云OSS配置
type AliyunOss struct {
    Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
    AccessKeyId     string `mapstructure:"accessKeyId" json:"accessKeyId" yaml:"accessKeyId"`
    AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret" yaml:"accessKeySecret"`
    BucketName      string `mapstructure:"bucketName" json:"bucketName" yaml:"bucketName"`
    BucketUrl       string `mapstructure:"bucketUrl" json:"bucketUrl" yaml:"bucketUrl"`
    BasePath        string `mapstructure:"basePath" json:"basePath" yaml:"basePath"`
}
