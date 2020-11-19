/*
 * Copyright (c) 2020 fuzamei-33cn Group Holding Limited. All Rights Reserved.
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
 *
 * This code is free software; you can redistribute it and/or modify it
 * under the terms of the GNU General Public License version 2 only, as
 * published by the Free Software Foundation. fuzamei designates this
 * particular file as subject to the "Classpath" exception as provided
 * by Oracle in the LICENSE file that accompanied this code.
 *
 * This code is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * version 2 for more details (a copy is included in the LICENSE file that
 * accompanied this code).
 *
 * You should have received a copy of the GNU General Public License version
 * 2 along with this work; if not, write to the Free Software Foundation,
 * Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package com.fuzamei.chain33;

import java.security.AccessController;

/**
 * All chain33 state database operations are in this class
 */
public class StateDB {

    static {
        AccessController.doPrivileged(
                new java.security.PrivilegedAction<Void>() {
                    public Void run() {
                        System.loadLibrary("java");
                        return null;
                    }
                });
        registerNatives0();
    }

    // set value to state db
    public static native boolean setState(byte[] key, byte[] value);

    // get value from statedb
    public static native byte[] getFromState(byte[] key);

    // set value to state db in the format of string
    public static native boolean setStateInStr(String key, String value);

    // get value from statedb in the format of string
    public static native String getFromStateInStr(String key);

    private static native void registerNatives0();
}