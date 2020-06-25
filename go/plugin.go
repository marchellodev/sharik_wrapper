package sharik_wrapper

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/marchellodev/sharic/wrapper"
)

const channelName = "sharik_wrapper"

// SharikWrapperPlugin implements flutter.Plugin and handles method.
type SharikWrapperPlugin struct{}

var _ flutter.Plugin = &SharikWrapperPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *SharikWrapperPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getLocalIp", p.handleLocalIp)
	return nil
}

func (p *SharikWrapperPlugin) handleLocalIp(arguments interface{}) (reply interface{}, err error) {
	return wrapper.GetIp(), nil
	//return "go-flutter " + flutter.PlatformVersion, nil
}
