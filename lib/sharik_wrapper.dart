import 'dart:async';

import 'package:flutter/services.dart';

class SharikWrapper {
  static const MethodChannel _channel = const MethodChannel('sharik_wrapper');

  static Future<String> get getLocalIp async {
    final String version = await _channel.invokeMethod('getLocalIp');
    return version;
  }
}
