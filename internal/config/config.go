package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

//creating config for the server


type Config struct{
	Server		ServerConfig
	Database    DatabaseConfig
	JWT 		JWTConfig
	AWS 		AWSConfig
	Upload 		UploadConfig
}

type ServerConfig struct{
	Port string 
	GinMode string
}

//Database configg 

type DatabaseConfig struct{
	Host string
	Port string
	User string
	Password string
	Name string 
	SSLMode string
	URL 	string
}

//jwt 
type JWTConfig struct{
	Secret					string
	ExpiresIn				time.Duration
	RefreshTokenExpires 	time.Duration
}

//Aws config

type AWSConfig struct{
	Region				string
	AccessKeyId			string
	SecretAccessKeyId	string
	S3Bucket			string
	S3Endpoint			string
}

type UploadConfig struct{
	Path		string
	MaxFileSize int64
}


//function to load environment variables ..

func Load() (*Config,error){
		_=godotenv.Load()//returns and error 
		jwtExpiresIn,_:=time.ParseDuration(getEnv("JWT_EXPIRES_IN","24h"))
		refreshTokenExpires,_:=time.ParseDuration(getEnv("REFRESH_TOKEN_EXPIRES_IN","720h"))
		MaxUploadSize,_:=strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE","10485760"),10,64)	

		return &Config{
			Server: ServerConfig{
				Port: getEnv("PORT","8080"),
				GinMode: getEnv("GIN_MODE","debug"),
			},

			Database: DatabaseConfig{
				Host: getEnv("DB_HOST","localhost"),
				Port: getEnv("DB_PORT","5432"),
				User: getEnv("DB_USER","postgres"),
				Password: getEnv("DB_PASSWORD","password"),
				Name: getEnv("DB_NAME","ECOMMERCEAPP"),
				SSLMode: getEnv("DB_SSL_MODE","disable"),
				URL: getEnv("DB_URL", ""),

			},
			
			JWT: JWTConfig{
				Secret: getEnv("JWT_SECRET","your_jwt_secret_key"),
				ExpiresIn: jwtExpiresIn,
				RefreshTokenExpires: refreshTokenExpires,
			},

			AWS: AWSConfig{

				Region: getEnv("AWS_REGION","us-east-1"),
				AccessKeyId: getEnv("AWS_ACCESS_KEY_ID","test"),
				SecretAccessKeyId: getEnv("AWS_SECRET_ACCESS_KEY","test"),
				S3Bucket: getEnv("AWS_S3_BUCKET","ecommerce-uploads"),
				S3Endpoint: getEnv("AWS_S3_ENDPOINT","http://localhost:9000"),

			},

			Upload: UploadConfig{
				Path: getEnv("UPLOAD_PATH","./uploads"),
				MaxFileSize: MaxUploadSize,
			},
		},nil 

}


//now we will create a function to manage data, like suppose we are accessing port info so first we will lookup environment variables env , if- 
//-its not there we will return a default data 

func getEnv(key,defaultValue string )string{
	if value :=os.Getenv(key);value!=""{
		return value
	}

	return defaultValue
}