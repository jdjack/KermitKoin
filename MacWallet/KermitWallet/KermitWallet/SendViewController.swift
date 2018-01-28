//
//  SendViewController.swift
//  KermitWallet
//
//  Created by Jack Chorley on 28/01/2018.
//  Copyright © 2018 Jack Chorley. All rights reserved.
//

import Cocoa

class SendViewController: NSViewController, NSTextFieldDelegate {
    
    var lastAmountText = ""
    var lastHexText = ""
    
    @IBOutlet weak var hexTextField: NSTextField!
    @IBOutlet weak var sendButton: NSButton!
    @IBOutlet weak var amountTextField: NSTextField!
    @IBOutlet weak var availableLable: NSTextField!
    var balance: Double!

    override func viewDidLoad() {
        super.viewDidLoad()
        
        sendButton.isEnabled = false
        
        availableLable.stringValue = "\(balance.decimalFormatted) Available"
        
    }
    
    func updateAvailableLabel(amount: Double) {
        availableLable.stringValue = "\((balance - amount).decimalFormatted) Available"
    }

    func isAmountValid(string: String) -> Bool {
        return string != ""
    }
    
    func isHexValid(string: String) -> Bool {
        
        if string == "" {
            return false
        }
        
        let xChecker = string.split(separator: "x", maxSplits: 3, omittingEmptySubsequences: false)
        return (xChecker.count == 2) && (xChecker[0].count == 1) && string.count == 18 && string.first! == "0"
    }
    
    @IBAction func sendButtonPressed(_ sender: Any) {
        self.dismiss(nil)
    }
}

extension SendViewController: NSControlTextEditingDelegate {
    override func controlTextDidChange(_ notification: Notification) {
        if let textField = notification.object as? NSTextField {
            
            if textField == amountTextField {
                var allowedCharacters = CharacterSet.decimalDigits
                allowedCharacters.insert(charactersIn: ".")
                let characterSet = CharacterSet(charactersIn: textField.stringValue)
                
                if !allowedCharacters.isSuperset(of: characterSet) {
                    textField.stringValue = lastAmountText
                } else if textField.stringValue != "" && Double(textField.stringValue)! > balance {
                    textField.stringValue = "\(balance!)"
                    lastAmountText = textField.stringValue
                } else {
                    lastAmountText = textField.stringValue
                }
                
                updateAvailableLabel(amount: Double(textField.stringValue) ?? 0)
                
                
                sendButton.isEnabled = isAmountValid(string: textField.stringValue) && isHexValid(string: hexTextField.stringValue)
            } else if textField == hexTextField {
                if textField.stringValue.count > 18 {
                    textField.stringValue = lastHexText
                } else {
                    lastHexText = textField.stringValue
                }
                
                sendButton.isEnabled = isAmountValid(string: textField.stringValue) && isHexValid(string: hexTextField.stringValue)
            }
            //do what you need here
        }
    }
}
