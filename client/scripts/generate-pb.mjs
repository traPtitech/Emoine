import fs from 'fs/promises'
import { exec as _exec } from 'child_process'
import { promisify } from 'util'
import process from 'process'

const protoTargets = ['comment', 'reaction', 'state', 'viewer', 'message']

const npx = process.platform === 'win32' ? 'npx.cmd' : 'npx'
const exec = promisify(_exec)

await fs.mkdir(new URL(`../src/lib/pb`, import.meta.url), {
  recursive: true
})

const promises = protoTargets.map(async p => {
  await exec(
    `${npx} pbjs -t static-module -w es6 ../docs/${p}.proto -o src/lib/pb/${p}.js`
  )
  await exec(`${npx} pbts src/lib/pb/${p}.js -o src/lib/pb/${p}.d.ts`)
})
await Promise.all(promises)
