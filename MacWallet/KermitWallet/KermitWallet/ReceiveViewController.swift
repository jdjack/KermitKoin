//
//  ReceiveViewController.swift
//  KermitWallet
//
//  Created by Jack Chorley on 27/01/2018.
//  Copyright © 2018 Jack Chorley. All rights reserved.
//

import Cocoa

class ReceiveViewController: NSViewController {

    var address: String!
    
    @IBOutlet weak var addressLabel: NSTextField!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        addressLabel.stringValue = address
    }
    
    @IBAction func crossButtonPressed(_ sender: Any) {
        self.dismiss(nil)
    }
}
