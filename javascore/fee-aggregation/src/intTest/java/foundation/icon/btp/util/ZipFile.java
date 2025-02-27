/*
 * Copyright 2020 ICON Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package foundation.icon.btp.util;

import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.util.zip.ZipEntry;
import java.util.zip.ZipOutputStream;

public class ZipFile {
    public static byte[] zipContent(String path) throws IOException {
        ByteArrayOutputStream outputStream = new ByteArrayOutputStream();
        ZipOutputStream zos = new ZipOutputStream(outputStream);
        recursiveZip(new File(path), null, zos);
        zos.close();
        outputStream.close();
        return outputStream.toByteArray();
    }

    private static void recursiveZip(File source, String zipPath, ZipOutputStream zos) throws IOException {
        if (source.isHidden()) {
            return;
        }
        if (source.isDirectory()) {
            String dir = source.getName();
            if (!dir.endsWith(File.separator)) {
                dir = dir + File.separator;
            }
            zos.putNextEntry(new ZipEntry(dir));
            zos.closeEntry();
            File[] files = source.listFiles();
            if (files == null) {
                return;
            }
            String path = zipPath == null ? dir : zipPath + dir;
            for (File file : files) {
                recursiveZip(file, path, zos);
            }
        } else {
            ZipEntry ze = new ZipEntry(zipPath + source.getName());
            zos.putNextEntry(ze);
            zos.write(Files.readAllBytes(source.toPath()));
            zos.closeEntry();
        }
    }
}
