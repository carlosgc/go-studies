// Package designerpdfviewer provides a solution implementation for the Designer PDF Viewer problem.
//
// Statement: In this PDF viewer, each word is highlighted independently. Using the letter heights given,
// determine the area of the rectangle highlight in mm² assuming all letters are 1mm wide.
//
// A complete description of this problem can be found here: https://www.hackerrank.com/challenges/designer-pdf-viewer
package designerpdfviewer

// DesignerPdfViewer determines the area of a word rectangle highlight
func DesignerPdfViewer(h []int32, word string) int32 {

	maxH := int32(0)
	wordLen := int32(len(word))

	// Create the visitedLetter array to store a boolean to each letter
	// That boolean indicates whether the letter height is known
	visitedLetter := [26]bool{}

	// Finding the highest height
	for _, charCode := range word {

		// Using the character ASCII code to calculate the index of visitedLetter
		idl := charCode - 97

		// Only test heights once
		letterNotVisited := !visitedLetter[idl]
		if letterNotVisited {
			visitedLetter[idl] = true
			if maxH < h[idl] {
				maxH = h[idl]
			}
		}
	}

	return int32(wordLen * maxH)
}
