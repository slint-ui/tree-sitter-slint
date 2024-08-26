import XCTest
import SwiftTreeSitter
import TreeSitterSlint

final class TreeSitterSlintTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_slint())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Slint grammar")
    }
}
