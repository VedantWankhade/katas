package de.vedantwankha.kata.dsa;

import java.io.File;

public class BFSFiles {
    public static void printFiles(String rootPath) {
        File rootFile = new File(rootPath);
        Queue<File> q = new LinkedList<>();
        q.pushLast(rootFile);
        while (!q.isEmpty()) {
            var f = q.popFirst();
            if (f.isFile()) {
                System.out.println(f.getAbsolutePath());
            } else {
                System.out.println(f.getAbsolutePath() + "/");
                var files = f.listFiles();
                if (files != null) {
                    for (var ff : files) {
                        q.pushLast(ff);
                    }
                }
            }
        }
    }
}
