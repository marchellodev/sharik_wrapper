import Flutter
import UIKit
import Wrapper

public class SwiftSharikWrapperPlugin: NSObject, FlutterPlugin {
  public static func register(with registrar: FlutterPluginRegistrar) {
    let channel = FlutterMethodChannel(name: "sharik_wrapper", binaryMessenger: registrar.messenger())
    let instance = SwiftSharikWrapperPlugin()
    registrar.addMethodCallDelegate(instance, channel: channel)
  }

  public func handle(_ call: FlutterMethodCall, result: @escaping FlutterResult) {
    result(Wrapper.WrapperGetIp())
  }
}
