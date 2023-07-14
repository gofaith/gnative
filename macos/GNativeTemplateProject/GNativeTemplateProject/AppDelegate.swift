//
//  AppDelegate.swift
//  GNativeTemplateProject
//
//  Created by StevenZack on 2023/7/12.
//

import Cocoa
import Src

//@main
class AppDelegate: NSObject, NSApplicationDelegate,SrcIBridgeProtocol {
    func executeSwift(_ s: String?) {
        do{
            let msg=try GNativeMessage(s)
            print(s)
            switch msg.method{
            case MessageMethods.showWindow.rawValue:
                let w=GNativeWindow(origin: msg.data[0])
                w.show()
            default:
                print("unsupported message method: \(msg.method)")
            }
        }catch{
            print("executeSwift caught error: \(error), \(s)")
        }
    }
    
    private var window: NSWindow!
    private var app:SrcApplication!
    
    func applicationDidFinishLaunching(_ aNotification: Notification) {
        print("applicationDidFinishLaunching")
        if(app == nil){
            app = SrcMain(self)
        }
    }
    
    func applicationWillTerminate(_ aNotification: Notification) {
        // Insert code here to tear down your application
    }

    func applicationSupportsSecureRestorableState(_ app: NSApplication) -> Bool {
        return true
    }


}

