//
//  AppDelegate.swift
//  KermitWallet
//
//  Created by Jack Chorley on 27/01/2018.
//  Copyright © 2018 Jack Chorley. All rights reserved.
//

import Cocoa

@NSApplicationMain
class AppDelegate: NSObject, NSApplicationDelegate {


    func applicationDidFinishLaunching(_ aNotification: Notification) {
        // Insert code here to initialize your application
    
        let mainWindow = NSApplication.shared.windows.first!
        
        mainWindow.titlebarAppearsTransparent = true
        mainWindow.titleVisibility = NSWindow.TitleVisibility.hidden
        mainWindow.styleMask = [NSWindow.StyleMask.fullSizeContentView, mainWindow.styleMask]
        mainWindow.isMovableByWindowBackground = true
    }

    func applicationWillTerminate(_ aNotification: Notification) {
        // Insert code here to tear down your application
    }


}

