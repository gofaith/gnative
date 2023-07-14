//
//  ViewController.swift
//  GNativeTemplateProject
//
//  Created by StevenZack on 2023/7/12.
//

import Cocoa
import SwiftUI

class ViewController: NSViewController {
    override func loadView() {
        view = NSHostingView(rootView: SwiftUIView())
    }
    override func viewDidLoad() {
        super.viewDidLoad()

        // Do any additional setup after loading the view.
    }

    override var representedObject: Any? {
        didSet {
        // Update the view, if already loaded.
        }
    }


}

