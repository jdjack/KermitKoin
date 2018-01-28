//
//  ViewController.swift
//  KermitWallet
//
//  Created by Jack Chorley on 27/01/2018.
//  Copyright © 2018 Jack Chorley. All rights reserved.
//

import Cocoa

class ViewController: NSViewController {

    @IBOutlet weak var balanceLabel: NSTextField!
    
    override func viewDidLoad() {
        super.viewDidLoad()

        balanceLabel.stringValue = "\(getBalance().decimalFormatted)"
        // Do any additional setup after loading the view.
    }

    override var representedObject: Any? {
        didSet {
        // Update the view, if already loaded.
        }
    }
    
    func getBalance() -> Double {
        return 3500
    }
    
    func getAddress() -> String {
        return "0x4204204204204204"
    }
    
    override func prepare(for segue: NSStoryboardSegue, sender: Any?) {
        if segue.identifier!.rawValue == "presentSend" {
            let dest = segue.destinationController as! SendViewController
            dest.balance = getBalance()
        } else if segue.identifier!.rawValue == "presentAddress" {
            let dest = segue.destinationController as! ReceiveViewController
            dest.address = getAddress()
        }
    }


}

