package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class DFSFilesTest {
    @Test
    void testDFSFiles() {
        DFSFiles.printFiles("/home/vedant/Second Brain");
        DFSFiles.printFilesRecursive("/home/vedant/Second Brain");
    }
}