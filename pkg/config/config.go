package config

var UploadLocation string = "C:\\Users\\User\\Desktop\\GoLang\\file_upload"

var AllowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

var Port string = "8000"

const MaxUploadSize = 10 << 20
