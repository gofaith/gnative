//
//  Dom.swift
//  GNativeTemplateProject
//
//  Created by StevenZack on 2023/7/14.
//

import Foundation
import Cocoa

class GNativeWindow{
    let origin:String!
    private var window:NSWindow!
    
    init(origin: String!) {
        self.origin = origin
        makeWindow()
    }
    
    private func makeWindow(){
        window = NSWindow(
        contentRect: NSRect(x: 0, y: 0, width: 480, height: 270),
            styleMask: [.miniaturizable, .closable, .resizable, .titled],
            backing: .buffered, defer: false)
        window.contentViewController = ViewController()
        window.center()
        window.backgroundColor = .systemPink
        window.isReleasedWhenClosed = false
        window.title = "No Storyboard Window"
    }
    
    func show(){
        if(window != nil){
            window.makeKeyAndOrderFront(nil)
        }
    }
}
