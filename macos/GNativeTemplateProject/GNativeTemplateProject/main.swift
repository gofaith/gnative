//
//  main.swift
//  GNativeTemplateProject
//
//  Created by StevenZack on 2023/7/12.
//

import Foundation
import Cocoa

let app = NSApplication.shared
let delegate = AppDelegate()
app.delegate = delegate

_ = NSApplicationMain(CommandLine.argc, CommandLine.unsafeArgv)
