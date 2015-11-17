//
//  main.swift
//  morse
//
//  Created by Gabriel Bremond on 05/06/14.
//  Copyright (c) 2014 photograve. All rights reserved.
//

import Foundation
import AppKit

var data: String
var opts: [MorseOptions]

(data, opts) = IOHelper.lookForData()

// Check data.
if data.isEmpty
{
    IOHelper.exitOnError("Empty data")
}

// Morse computation.
var out = "failure"
if opts.contains(MorseOptions.Mode)
{
    out = Morse.demorse(data)
}
else
{
    out = Morse.morse(data)
}

// Copy to MacOS pastboard.
if opts.contains(MorseOptions.Pastboard)
{
    let pasteboard = NSPasteboard.generalPasteboard()
    pasteboard.declareTypes([NSStringPboardType], owner: nil)
    pasteboard.setString(out, forType: NSStringPboardType)
}

// Output result.
if opts.contains(MorseOptions.Verbose)
{
    print("\n\nResult:\n\n\"\(out)\"\n\n")
}
else
{
    print(out)
}

exit(0)