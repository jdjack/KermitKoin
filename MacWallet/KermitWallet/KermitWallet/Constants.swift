//
//  Constants.swift
//  KermitWallet
//
//  Created by Jack Chorley on 28/01/2018.
//  Copyright © 2018 Jack Chorley. All rights reserved.
//

import Foundation

extension Int {
    var decimalFormatted:String {
        let formatter = NumberFormatter()
        formatter.numberStyle = .decimal
        formatter.minimumFractionDigits = 0
        
        return formatter.string(from: NSNumber(value:self))!
    }
}

extension Double {
    var decimalFormatted:String {
        let formatter = NumberFormatter()
        formatter.numberStyle = .decimal
        formatter.maximumFractionDigits = 2
        formatter.minimumFractionDigits = 2
        return formatter.string(from: NSNumber(value: self))!
    }
}
