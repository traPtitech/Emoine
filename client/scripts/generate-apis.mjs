import fs from 'fs/promises'
import { exec as _exec } from 'child_process'
import { promisify } from 'util'
import addApis from './add-apis.mjs'
import process from 'process'

const exec = promisify(_exec)

const SWAGGER_PATH = '../docs/openapi.yml'
const GENERATED_DIR = 'src/lib/apis/generated'

const npx = process.platform === 'win32' ? 'npx.cmd' : 'npx'

const generateCmd = [
  npx,
  'openapi-generator-cli',
  'generate',
  '-g',
  'typescript-axios',
  '-i',
  SWAGGER_PATH,
  '-o',
  GENERATED_DIR
]

await fs.mkdir(new URL(`../${GENERATED_DIR}`, import.meta.url), {
  recursive: true
})

await exec(generateCmd.join(' '))

// generate Apis class
await addApis(GENERATED_DIR)
