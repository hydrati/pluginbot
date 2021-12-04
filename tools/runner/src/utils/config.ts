import { defineConfig, Plugin } from 'rollup'

import alias from '@rollup/plugin-alias'
import cjs from '@rollup/plugin-commonjs'
import json from '@rollup/plugin-json'
import node from '@rollup/plugin-node-resolve'

import esbuild from 'rollup-plugin-esbuild'
import path from 'path'

export const plugins: Plugin[] = [
  esbuild({
    tsconfig: "./tsconfig.json",
    minify: true,
  }),
  cjs(),
  node({
    browser: false,
    preferBuiltins: true,
  }),
  alias({
    entries: [
      { find: '../../src/sleep', replacement: path.resolve(__dirname, "compat", "sleep.ts") },
      { find: '../../src/utils', replacement: path.resolve(__dirname, "compat", "utils.ts") },
      { find: '../../src/class', replacement: path.resolve(__dirname, "compat", "class.ts") },
      { find: '../../src/enum', replacement: path.resolve(__dirname, "compat", "enum.ts") },
      { find: 'axios', replacement: require.resolve("axios") },
      { find: 'cheerio', replacement: require.resolve("cheerio") }
    ]
  }),
  json()
]