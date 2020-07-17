/* eslint-disable no-undef */
/* eslint-disable @typescript-eslint/no-var-requires */
const fs = require('fs').promises
const path = require('path')

;(async () => {
  await fs.mkdir(path.resolve(__dirname, '../src/lib/pb'), { recursive: true })
})()
