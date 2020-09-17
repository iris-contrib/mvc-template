import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import babel from '@rollup/plugin-babel';
import { terser } from 'rollup-plugin-terser';


const name = 'RollupTypeScriptBabel';

const extensions = [
    '.js', '.jsx', '.ts', '.tsx',
];

const production = !process.env.ROLLUP_WATCH;

export default {
    input: 'src/index.ts',
    output: {
        sourcemap: true,
        format: 'iife',
        name: 'app',
        file: 'public/build/bundle.js'
    },
    plugins: [
        // Allows node_modules resolution
        resolve({ extensions }),

        // Allow bundling cjs modules. Rollup doesn't understand cjs
        commonjs(),

        // Compile TypeScript/JavaScript files
        babel({ extensions, include: ['src/**/*'] }),
        // If we're building for production (npm run build
        // instead of npm run dev), minify
        production && terser()
    ],
    watch: {
        clearScreen: false
    }
};