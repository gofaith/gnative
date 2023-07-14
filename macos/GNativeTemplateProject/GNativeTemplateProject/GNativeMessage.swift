//
//  GNativeMessage.swift
//  GNativeTemplateProject
//
//  Created by StevenZack on 2023/7/14.
//

import Foundation

enum MessageMethods:String{
    case call = "call"
    case returnValue = "return"
    case showWindow = "show"
}

class GNativeMessage{
    var origin:String=""
    var method:String=""
    var requestID:UInt64=0
    var path:String=""
    var data:[String]=[]
    
    enum MessageError: Error {
        case invalidMessage(String)
    }

    init(_ origin:String!)throws{
        self.origin = origin
        let ss=origin.split(separator: " ")
        if ss.count<3{
            throw MessageError.invalidMessage(origin)
        }
        for i in 0...ss.count-1{
            let s=String(ss[i])
            if i==0{
                self.method=s
                continue
            }
            if i==1{
                self.requestID=UInt64(s)!
                continue
            }
            if i==2{
                self.path = s.decodeUrl()!
                continue
            }
            self.data.append(s.decodeUrl()!)
        }
    }
}

extension String
{
    func encodeUrl() -> String?
    {
        return self.addingPercentEncoding( withAllowedCharacters: NSCharacterSet.urlQueryAllowed)
    }
    func decodeUrl() -> String?
    {
        return self.removingPercentEncoding
    }
}
