import XCTest
import SwiftTreeSitter
import TreeSitterGritql

final class TreeSitterGritqlTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_gritql())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading GritQL grammar")
    }
}
