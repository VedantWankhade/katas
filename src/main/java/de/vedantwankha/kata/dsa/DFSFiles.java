package de.vedantwankha.kata.dsa;

import java.io.File;

public class DFSFiles {
    public static void printFiles(String rootPath) {
        File rootFile = new File(rootPath);
        Stack<File> q = new LinkedList<>();
        q.push(rootFile);
        while (!q.isEmpty()) {
            var f = q.pop();
            if (f.isFile()) {
                System.out.println(f.getAbsolutePath());
            } else {
                System.out.println(f.getAbsolutePath());
                var files = f.listFiles();
                if (files != null) {
                    for (var ff : files) {
                        q.push(ff);
                    }
                }
            }
        }
    }

    public static void printFilesRecursive(String rootFile) {
        printFilesRecursive(new File(rootFile));
    }

    private static void printFilesRecursive(File file) {
        if (file.isFile()) {
            System.out.println(file.getAbsolutePath());
        } else {
            System.out.println(file.getAbsolutePath());
            var files = file.listFiles();
            if (files != null)
                for (var f: files) {
                    printFilesRecursive(f);
                }
        }
    }
}
