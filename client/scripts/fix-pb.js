/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable no-undef */
const path = require('path')
const { Project } = require('ts-morph')

const GENERATED = 'src/lib/pb'
const protoTargets = ['comment', 'reaction', 'state', 'viewer', 'message']

const fixImport = async project => {
  const sourceFile = project.getSourceFileOrThrow('message.js')
  const imports = sourceFile.getImportDeclarations()

  const im = imports.find(
    im => im.getModuleSpecifierValue() === 'protobufjs/minimal'
  )
  const importName = im.getNamespaceImport().getText()
  im.removeNamespaceImport()
  im.setDefaultImport(importName)

  await sourceFile.save()
}

const addImport = async project => {
  const sourceFile = project.getSourceFileOrThrow('message.d.ts')

  protoTargets
    .filter(p => p !== 'message')
    .forEach(p => {
      sourceFile.addImportDeclaration({
        namedImports: [{ name: `I${p[0].toUpperCase()}${p.slice(1)}` }],
        moduleSpecifier: `./${p}`
      })
    })

  await sourceFile.save()
}

;(async () => {
  const dir = path.resolve(__dirname, '../', GENERATED)

  const project = new Project()
  project.addSourceFilesAtPaths(`${dir}/**/*.{js,ts}`)

  await fixImport(project)
  await addImport(project)
})()
