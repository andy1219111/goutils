package iniparser

import (
	"errors"
	"log"

	"github.com/go-ini/ini"
)

//IniParser ini解析器
type IniParser struct {
	INIReader *ini.File
}

//Load 加载ini配置文件
func (i *IniParser) Load(iniFile string) error {
	if iniFile == "" {
		return errors.New("the conf file is is not exist")
	}

	iniFile, err := ini.Load(iniFile)
	if err != nil {
		return err
	}
	i.INIReader = iniFile
	return nil
}

//GetString 获取字符串参数
func (i *IniParser) GetString(section, key string) string {
	if i.INIReader == nil {
		return ""
	}

	s := i.INIReader.Section(section)
	if s == nil {
		return ""
	}
	return s.Key(key).String()
}

//GetInt 获取整形参数
func (i *IniParser) GetInt(section, key string) int {
	if i.INIReader == nil {
		return 0
	}

	s := i.INIReader.Section(section)
	if s == nil {
		return 0
	}
	valueInt, err := s.Key(key).Int()
	if err != nil {
		log.Printf("get ini value error,section:%s,key:%s,%v\n", section, key, err)
		return 0
	}
	return valueInt
}

//GetInt64 获取整形参数
func (i *IniParser) GetInt64(section, key string) int64 {
	if i.INIReader == nil {
		return 0
	}

	s := i.INIReader.Section(section)
	if s == nil {
		return 0
	}
	valueInt, err := s.Key(key).Int64()
	if err != nil {
		log.Println("get ini value error,", err)
		return 0
	}
	return valueInt
}

//GetFloat32 获取float32参数
func (i *IniParser) GetFloat32(section, key string) float32 {
	if i.INIReader == nil {
		return 0
	}

	s := i.INIReader.Section(section)
	if s == nil {
		return 0
	}
	valueFloat64, err := s.Key(key).Float64()
	if err != nil {
		log.Println("get ini value error,", err)
		return 0
	}
	return float32(valueFloat64)
}

//GetFloat64 获取float64参数
func (i *IniParser) GetFloat64(section, key string) float64 {
	if i.INIReader == nil {
		return 0
	}

	s := i.INIReader.Section(section)
	if s == nil {
		return 0
	}
	valueFloat64, err := s.Key(key).Float64()
	if err != nil {
		log.Println("get ini value error,", err)
		return 0
	}
	return valueFloat64
}

//GetBool 获取布尔值
func (i *IniParser) GetBool(section, key string) bool {
	if i.INIReader == nil {
		return false
	}

	s := i.INIReader.Section(section)
	if s == nil {
		return false
	}
	return "true" == s.Key(key).String()
}
