/* eslint-disable no-undef */
/* eslint-disable @typescript-eslint/no-var-requires */
const fs = require('fs').promises
const path = require('path')
const { exec: _exec } = require('child_process')
const { promisify } = require('util')

const protoTargets = ['comment', 'reaction', 'state', 'message']

const npx = process.platform === 'win32' ? 'npx.cmd' : 'npx'
const exec = promisify(_exec)

;(async () => {
  await fs.mkdir(path.resolve(__dirname, '../src/lib/pb'), {
    recursive: true
  })

  const promises = protoTargets.map(async p => {
    await exec(
      `${npx} pbjs -t static-module -w es6 ../docs/${p}.proto -o src/lib/pb/${p}.js`
    )
    await exec(`${npx} pbts src/lib/pb/${p}.js -o src/lib/pb/${p}.d.ts`)
  })
  await Promise.all(promises)
})()
