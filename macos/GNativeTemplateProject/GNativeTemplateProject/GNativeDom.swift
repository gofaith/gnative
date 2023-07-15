//
//  GNativeDom.swift
//  GNativeTemplateProject
//
//  Created by StevenZack on 2023/7/15.
//

import Foundation

class GNativeDom {
    var tagName: String = ""
    var attributes: [String: String] = [:]
    var children: [GNativeDom] = []
    var innerText: String = ""
    
    init(_ s: String) {
        
    }
}
